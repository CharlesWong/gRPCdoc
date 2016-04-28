## Log Statistics APIs

### Matrix Engine Log APIs

##### GetMatrixStats

GetMatrixStats gets the service statistics information.

gRPC:

```
rpc GetMatrixStats(NullMessage) returns (MatrixStats) {}
```

HTTP:

```
Method: GET/POST
Path: api/biz/getmatrixstats
Body: nullmessage=[json format of NullMessage]
Return: {"code":0,msg:"success",redirect:"",data:[json format of MatrixStats]}
Errors:

```

### Data Engine Log APIs

##### GetDataEngineStats

GetDataEngineStats gets the service statistics information.

gRPC:

```
rpc GetDataEngineStats(NullMessage) returns (DataEngineStats) {}
```

HTTP:

```
Method: GET/POST
Path: api/biz/getdataenginestats
Body: nullmessage=[json format of NullMessage]
Return: {"code":0,msg:"success",redirect:"",data:[json format of DataEngineStats]}
Errors:

```

##### GetDataEngineUsage

GetDataEngineUsage gets the resources usage information.

gRPC:

```
rpc GetDataEngineUsage(NullMessage) returns (DataEngineUsage) {}
```

HTTP:

```
Method: GET/POST
Path: api/biz/getdataengineusage
Body: nullmessage=[json format of NullMessage]
Return: {"code":0,msg:"success",redirect:"",data:[json format of DataEngineUsage]}
Errors:

```

## Related Data Structure
```
message MatrixStats {
	// TODO
}
```

```
message DataEngineStats {
	// TODO
}
```

```
message DataEngineUsage {
	// TODO
}
```

