## Generic Storage API

### Storage/Retrival

##### GenUniqueID

Generate an unique id for the given table.
Generate a new ID in req.Table and return it in resp.Id.

gRPC:

```
rpc GenUniqueID(GenericStorageRequest) returns (GenericStorageResponse) {}
```

HTTP:

```
Method: GET/POST
Path: api/biz/genuniqueid
Body: genericstoragerequest=[json format of GenericStorageRequest]
Return: {"code":0,msg:"success",redirect:"",data:[json format of GenericStorageResponse]}
Errors:
  41001 "Table doesn't exist."
```

##### CreateTable

Create a new table.
Generate a new table as in req.Table.

gRPC:

```
rpc CreateTable(GenericStorageRequest) returns (GenericStorageResponse) {}
```

HTTP:

```
Method: GET/POST
Path: api/biz/createtable
Body: genericstoragerequest=[json format of GenericStorageRequest]
Return: {"code":0,msg:"success",redirect:"",data:[json format of GenericStorageResponse]}
Errors:
  41002 "Table already exist."
```

##### DropTable

Drop an existing table.
Drop table as in req.Table.

gRPC:

```
rpc DropTable(GenericStorageRequest) returns (GenericStorageResponse) {}
```

HTTP:

```
Method: GET/POST
Path: api/biz/droptable
Body: genericstoragerequest=[json format of GenericStorageRequest]
Return: {"code":0,msg:"success",redirect:"",data:[json format of GenericStorageResponse]}
Errors:
  41001 "Table doesn't exist."
```

##### GetTableValues

Get the key/val map.
Get all key/val pairs from req.Table and return it in resp.Id and resp.Val.

gRPC:

```
rpc GetTableValues(GenericStorageRequest) returns (GenericStorageResponse) {}
```

HTTP:

```
Method: GET/POST
Path: api/biz/gettablevalues
Body: genericstoragerequest=[json format of GenericStorageRequest]
Return: {"code":0,msg:"success",redirect:"",data:[json format of GenericStorageResponse]}
Errors:
  41001 "Table doesn't exist."
```

##### AddValue

Add new value into the given table, returns the new ID.
Add new value into req.Table and return the ID in resp.Id.

gRPC:

```
rpc AddValue(GenericStorageRequest) returns (GenericStorageResponse) {}
```

HTTP:

```
Method: GET/POST
Path: api/biz/addvalue
Body: genericstoragerequest=[json format of GenericStorageRequest]
Return: {"code":0,msg:"success",redirect:"",data:[json format of GenericStorageResponse]}
Errors:
  41001 "Table doesn't exist."
```

##### DelValue

Delete key/value based on the given key.
Delete req.Id.

gRPC:

```
rpc DelValue(GenericStorageRequest) returns (GenericStorageResponse) {}
```

HTTP:

```
Method: GET/POST
Path: api/biz/delvalue
Body: genericstoragerequest=[json format of GenericStorageRequest]
Return: {"code":0,msg:"success",redirect:"",data:[json format of GenericStorageResponse]}
Errors:
  41001 "Table doesn't exist."
  41003 "Key doesn't exist."
```

##### SetValue

Set the given key/value pair.
Set the req.Id/req.Val in req.Table.

gRPC:

```
rpc SetValue(GenericStorageRequest) returns (GenericStorageResponse) {}
```

HTTP:

```
Method: GET/POST
Path: api/biz/setvalue
Body: genericstoragerequest=[json format of GenericStorageRequest]
Return: {"code":0,msg:"success",redirect:"",data:[json format of GenericStorageResponse]}
Errors:
  41001 "Table doesn't exist."
```

##### IfTableExist

Check if the given table exists.
Check if req.Table exists and set the result in resp.Exists.

gRPC:

```
rpc IfTableExist(GenericStorageRequest) returns (GenericStorageResponse) {}
```

HTTP:

```
Method: GET/POST
Path: api/biz/iftableexist
Body: genericstoragerequest=[json format of GenericStorageRequest]
Return: {"code":0,msg:"success",redirect:"",data:[json format of GenericStorageResponse]}
Errors:

```

##### IfKeyExist

Check if the given key exists.
Check if req.Id exists in req.Table and set the result in resp.Exists.

gRPC:

```
rpc IfKeyExist(GenericStorageRequest) returns (GenericStorageResponse) {}
```

HTTP:

```
Method: GET/POST
Path: api/biz/ifkeyexist
Body: genericstoragerequest=[json format of GenericStorageRequest]
Return: {"code":0,msg:"success",redirect:"",data:[json format of GenericStorageResponse]}
Errors:
  41001 "Table doesn't exist."
```

## Related Data Structure
```
message GenericStorageRequest {
	string Table = 1; // Required field in request.
	repeated int64 Id = 2;
	repeated bytes Val = 3;
}
```

```
message GenericStorageResponse {
	bool Exists = 1;
	// string Table = 2; // NOT USED.
	repeated int64 Id = 3;
	repeated bytes Val = 4;
}
```

