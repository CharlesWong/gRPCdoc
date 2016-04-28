package main

import (
	"flag"
	"github.com/CharlesWong/grpcdoc/gen"
	"github.com/CharlesWong/grpcdoc/parse"
	"io/ioutil"
	"log"
)

type GRpcService struct {
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	protoFile := flag.String("proto", "", "gRPC proto file.")
	outputFile := flag.String("output", "", "API doc file.")
	flag.Parse()

	bytes, err := ioutil.ReadFile(*protoFile)
	if err != nil {
		log.Fatal(err)
	}

	services := parse.ExtractService(string(bytes))

	dataStructures := parse.ExtractDataStructures(string(bytes))

	log.Println(services)

	log.Println("\n", gen.GenMd(services, dataStructures))

	ioutil.WriteFile(*outputFile, []byte(gen.GenMd(services, dataStructures)), 0666)
}
