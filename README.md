# gRPCdoc

This tool can auto-generate API doc for gRPC definitions. It consists of two parts:

* parser for parsing gRPC proto file.
* generator for generating API doc in given format (Only support .md currently).

The parser and generator needs to be modified to meet customized gRPC definition annotations and API doc format. This tool only privodes one way of generating API doc.
