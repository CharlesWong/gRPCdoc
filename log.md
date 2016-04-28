## Log Service

### Log Service APIs

##### ListAllLogs

ListAllLogs lists all log entries of all services on all cluster nodes.

gRPC:

```
rpc ListAllLogs(NullMessage) returns (LogInfos) {}
```

HTTP:

```
Method: GET/POST
Path: api/biz/listalllogs
Body: nullmessage=[json format of NullMessage]
Return: {"code":0,msg:"success",redirect:"",data:[json format of LogInfos]}
Errors:

```

##### ListLogs

ListLogs lists all log entries info (including log file pathes) of given logInfo.ServiceId and/or logInfo.ClusterNode.

gRPC:

```
rpc ListLogs(LogInfo) returns (LogInfos) {}
```

HTTP:

```
Method: GET/POST
Path: api/biz/listlogs
Body: loginfo=[json format of LogInfo]
Return: {"code":0,msg:"success",redirect:"",data:[json format of LogInfos]}
Errors:
  41302: No given service Id.
  41303: No given cluster node.
```

## Related Data Structure
```
message LogInfo {
	int32 Id = 1;
	int32 ServiceId = 2;
	ClusterNode ClusterNode = 3;
	repeated string Path = 4;
	// TODO
}
```

```
message LogInfos {
	repeated LogInfo LogInfo = 1;
}
```

