package gen

import (
	"bytes"
	"fmt"
	"github.com/CharlesWong/grpcdoc/model"
	"strings"
)

func GenMd(services []*model.Service, dataStructures []*model.DataStructure) string {
	buf := new(bytes.Buffer)

	for _, s := range services {
		fmt.Fprintf(buf, "%s\n\n", strings.Join(s.Annotation, "\n"))
		for _, sec := range s.Section {
			fmt.Fprintf(buf, "%s\n\n", strings.Join(sec.Annotation, "\n"))
			for _, m := range sec.Method {
				fmt.Fprintf(buf, "##### %s\n\n", m.MethodName)
				fmt.Fprintf(buf, "%s\n\n", strings.Join(m.Annotation, "\n"))
				fmt.Fprintf(buf, "gRPC:\n\n```\n")
				fmt.Fprintf(buf, "rpc %s(%s) returns (%s) {}\n", m.MethodName, m.ReqName, m.RespName)
				fmt.Fprintf(buf, "```\n\n")

				fmt.Fprintf(buf, "HTTP:\n\n```\n")
				fmt.Fprintf(buf, "Method: GET/POST\n")
				fmt.Fprintf(buf, "Path: api/biz/%s\n", strings.ToLower(m.MethodName))
				fmt.Fprintf(buf, "Body: %s=[json format of %s]\n", strings.ToLower(m.ReqName), m.ReqName)
				fmt.Fprintf(buf, "Return: {\"code\":0,msg:\"success\",redirect:\"\",data:[json format of %s]}\n", m.RespName)
				fmt.Fprintf(buf, "Errors:\n")
				fmt.Fprintf(buf, "%s\n", strings.Join(m.ErrorCode, "\n"))
				fmt.Fprintf(buf, "```\n\n")
			}
		}
	}

	fmt.Fprintf(buf, "## Related Data Structure\n")
	for _, ds := range dataStructures {
		fmt.Fprintf(buf, "```\n%s\n```\n\n", ds.Raw)
	}
	return buf.String()
}
