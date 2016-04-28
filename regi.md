## Registration Service

### Service Registration APIs

##### ListAllServices

ListAllServices lists all services.

gRPC:

```
rpc ListAllServices(DistributedServiceInfo) returns (DistributedServiceInfo) {}
```

HTTP:

```
Method: GET/POST
Path: api/biz/listallservices
Body: distributedserviceinfo=[json format of DistributedServiceInfo]
Return: {"code":0,msg:"success",redirect:"",data:[json format of DistributedServiceInfo]}
Errors:

```

##### RegisterService

RegisterService registers the service info.

gRPC:

```
rpc RegisterService(DistributedServiceInfo) returns (NullMessage) {}
```

HTTP:

```
Method: GET/POST
Path: api/biz/registerservice
Body: distributedserviceinfo=[json format of DistributedServiceInfo]
Return: {"code":0,msg:"success",redirect:"",data:[json format of NullMessage]}
Errors:
  41401: Invalid service.
```

##### RemoveService

RemoveService removes the service.

gRPC:

```
rpc RemoveService(DistributedServiceInfo) returns (NullMessage) {}
```

HTTP:

```
Method: GET/POST
Path: api/biz/removeservice
Body: distributedserviceinfo=[json format of DistributedServiceInfo]
Return: {"code":0,msg:"success",redirect:"",data:[json format of NullMessage]}
Errors:
  41402: Service doesn't exist.
```

## Related Data Structure
```
message DistributedServiceInfo {
	int32 Id = 1;
	int32 InstanceId = 2;
	ClusterNode ClusterNode = 3;
	// TODO
}
```

```
message DistributedServiceInfos {
	repeated DistributedServiceInfo DistributedServiceInfo = 1;
}
```

