## Cluster Service

### Cluster Management APIs

##### ListAllNodes

ListAllNodes lists all cluster nodes.

gRPC:

```
rpc ListAllNodes(NullMessage) returns (ClusterNode) {}
```

HTTP:

```
Method: GET/POST
Path: api/biz/listallnodes
Body: nullmessage=[json format of NullMessage]
Return: {"code":0,msg:"success",redirect:"",data:[json format of ClusterNode]}
Errors:

```

##### AddNode

AddNode adds the given cluster node.

gRPC:

```
rpc AddNode(ClusterNode) returns (NullMessage) {}
```

HTTP:

```
Method: GET/POST
Path: api/biz/addnode
Body: clusternode=[json format of ClusterNode]
Return: {"code":0,msg:"success",redirect:"",data:[json format of NullMessage]}
Errors:
  41201: Invalid cluster node.
```

##### RemoveNode

RemoveNode removes the given cluster node.

gRPC:

```
rpc RemoveNode(ClusterNode) returns (NullMessage) {}
```

HTTP:

```
Method: GET/POST
Path: api/biz/removenode
Body: clusternode=[json format of ClusterNode]
Return: {"code":0,msg:"success",redirect:"",data:[json format of NullMessage]}
Errors:
  41202: Cluster node doesn't exist.
```

##### SetEntryIpAddr

SetEntryIpAddr sets the entry Ip addr of the cluster.

gRPC:

```
rpc SetEntryIpAddr(ClusterConfig) returns (NullMessage) {}
```

HTTP:

```
Method: GET/POST
Path: api/biz/setentryipaddr
Body: clusterconfig=[json format of ClusterConfig]
Return: {"code":0,msg:"success",redirect:"",data:[json format of NullMessage]}
Errors:

```

##### GetEntryIpAddr

GetEntryIpAddr gets the entry Ip addr of the cluster.

gRPC:

```
rpc GetEntryIpAddr(NullMessage) returns (ClusterConfig) {}
```

HTTP:

```
Method: GET/POST
Path: api/biz/getentryipaddr
Body: nullmessage=[json format of NullMessage]
Return: {"code":0,msg:"success",redirect:"",data:[json format of ClusterConfig]}
Errors:

```

## Related Data Structure
```
message ClusterNode {
	string InternalIPv4Addr = 1;
	// TODO
}
```

```
message ClusterNodes {
	repeated ClusterNode ClusterNode = 1;
}
```

```
message ClusterConfig {
	string EntryIpAddr = 1;
	// TODO
}
```

