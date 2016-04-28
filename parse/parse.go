package parse

import (
	"github.com/CharlesWong/grpcdoc/model"
	"log"
	"regexp"
	"strings"
)

func SplitLines(s string) []string {
	newLines := make([]string, 0)
	lines := strings.Split(s, "\n")
	for _, l := range lines {
		l = strings.TrimSpace(l)
		newLines = append(newLines, l)
	}
	return newLines
}

func ExtractDataStructures(proto string) []*model.DataStructure {
	regexpDs := regexp.MustCompile(`(?im)(message|enum)\s+(?P<MsgName>[^\s{]+)\s*{[^}]*}`)

	dsStrs := regexpDs.FindAllString(proto, -1)

	dataStructures := make([]*model.DataStructure, 0)

	for _, s := range dsStrs {
		ds := new(model.DataStructure)
		ds.Raw = s
		dataStructures = append(dataStructures, ds)
	}
	return dataStructures
}

func ExtractService(proto string) []*model.Service {
	regexpService := regexp.MustCompile(`(?im)(//.*\s*)*service (?P<ServiceName>.+)\s+{(\s+(//.+\s+)*(.+{}\s+)+)+}`)

	serviceStrs := regexpService.FindAllString(proto, -1)

	services := make([]*model.Service, 0)

	for _, s := range serviceStrs {
		service := new(model.Service)

		lines := SplitLines(s)

		annotation, body := ExtractAnnotation(lines)
		service.Annotation = annotation

		service.Name = ExtractServiceName(body)

		service.Section = ExtractSection(body)

		log.Println(service)
		services = append(services, service)
	}
	return services
}

func RemoveCommentSign(lines []string) []string {
	newLines := make([]string, 0)
	for _, l := range lines {
		l = strings.TrimSpace(l)
		l = strings.TrimPrefix(l, "// ")
		newLines = append(newLines, l)
	}
	return newLines
}

func ExtractAnnotation(lines []string) ([]string, []string) {
	annotationIndex := -1
	for i, l := range lines {
		if strings.HasPrefix(l, "// ") {
			annotationIndex = i
		} else {
			break
		}
	}
	return RemoveCommentSign(lines[:annotationIndex+1]), lines[annotationIndex+1:]
}

func ExtractServiceName(lines []string) string {
	regexpMethod := regexp.MustCompile(`(?im)service\s+(?P<ServiceName>[^\s]+)\s+{`)

	for _, l := range lines {
		if len(l) > 0 {
			match := regexpMethod.FindStringSubmatch(l)
			if len(match) == 0 {
				continue
			}
			result := make(map[string]string)
			for i, name := range regexpMethod.SubexpNames() {
				if i != 0 {
					result[name] = match[i]
				}
			}
			return result["ServiceName"]
		}
	}
	return ""
}

func ExtractSection(lines []string) []*model.Section {
	annotation := make([]string, 0)
	methods := make([]string, 0)
	isAnnotation := false

	sections := make([]*model.Section, 0)
	for _, l := range lines {
		if len(strings.TrimSpace(l)) == 0 {
			continue
		}
		if strings.HasPrefix(l, "// #") {
			if !isAnnotation {
				// Encountered new section. So need to end the previous section.
				if len(methods) > 0 || len(annotation) > 0 {
					section := new(model.Section)
					section.Annotation = RemoveCommentSign(annotation)
					section.Method = ExtractMethod(methods)
					sections = append(sections, section)
					annotation = make([]string, 0)
					methods = make([]string, 0)
				}
			}
			annotation = append(annotation, l)
			isAnnotation = true
		} else if strings.HasPrefix(l, "rpc") || strings.HasPrefix(l, "// ") {
			methods = append(methods, l)
			isAnnotation = false
		}
	}
	// End current section.
	if len(methods) > 0 || len(annotation) > 0 {
		section := new(model.Section)
		section.Annotation = RemoveCommentSign(annotation)
		section.Method = ExtractMethod(methods)
		sections = append(sections, section)
		annotation = make([]string, 0)
		methods = make([]string, 0)
	}
	return sections
}

func ExtractMethod(lines []string) []*model.Method {
	annotation := make([]string, 0)
	method := make([]string, 0)
	isAnnotation := false

	methods := make([]*model.Method, 0)
	for _, l := range lines {
		if len(strings.TrimSpace(l)) == 0 {
			continue
		}
		if strings.HasPrefix(l, "// ") {
			if !isAnnotation {
				// Encountered new section. So need to end the previous section.
				if len(method) > 0 {
					m := ExtractMethodAnnotation(annotation)
					m.MethodName, m.ReqName, m.RespName = ExtractMethodParam(method)
					methods = append(methods, m)
					annotation = make([]string, 0)
					method = make([]string, 0)
				}
			}
			annotation = append(annotation, strings.TrimPrefix(l, "// "))
			isAnnotation = true
		} else if strings.HasPrefix(l, "rpc") {
			if len(method) > 0 {
				m := ExtractMethodAnnotation(annotation)
				m.MethodName, m.ReqName, m.RespName = ExtractMethodParam(method)
				methods = append(methods, m)
				annotation = make([]string, 0)
				method = make([]string, 0)
			}
			method = append(method, l)
			isAnnotation = false
		}
	}
	// End current method.
	if len(method) > 0 {
		m := ExtractMethodAnnotation(annotation)
		m.MethodName, m.ReqName, m.RespName = ExtractMethodParam(method)
		methods = append(methods, m)
		annotation = make([]string, 0)
		method = make([]string, 0)
	}
	return methods
}

func ExtractMethodAnnotation(lines []string) *model.Method {
	isErr := false
	annotation := make([]string, 0)
	errs := make([]string, 0)
	method := new(model.Method)
	for _, l := range lines {
		if len(l) == 0 {
			continue
		}
		if strings.HasPrefix(l, "Errors:") {
			isErr = true
			continue
		}
		if !isErr {
			annotation = append(annotation, l)
		} else {
			errs = append(errs, l)
		}
	}
	method.Annotation = annotation
	method.ErrorCode = errs
	return method
}

func ExtractMethodParam(lines []string) (string, string, string) {
	regexpMethod := regexp.MustCompile(`(?im)rpc\s+(?P<MethodName>[^\(]+)\((?P<ReqName>[^\)]+)\)\s+returns\s+\((?P<RespName>[^\)]+)\)\s+{\s*}`)

	for _, l := range lines {
		if len(l) > 0 {
			match := regexpMethod.FindStringSubmatch(l)
			if len(match) == 0 {
				continue
			}
			result := make(map[string]string)
			for i, name := range regexpMethod.SubexpNames() {
				if i != 0 {
					result[name] = match[i]
				}
			}
			return result["MethodName"], result["ReqName"], result["RespName"]
		}
	}
	return "", "", ""
}
