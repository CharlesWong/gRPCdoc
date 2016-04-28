## Business Intelligence APIs

### Metadata init APIs

##### UpdateIdTextMapping

UpdateIdTextMapping updates Id mapping table of mapping.Type with mapping.Mapping.

gRPC:

```
rpc UpdateIdTextMapping(IdTextMapping) returns (IdTextMappingChange) {}
```

HTTP:

```
Method: GET/POST
Path: api/biz/updateidtextmapping
Body: idtextmapping=[json format of IdTextMapping]
Return: {"code":0,msg:"success",redirect:"",data:[json format of IdTextMappingChange]}
Errors:
  40101 "Id/Text array lengthes are not matching"
  40102 "Duplicated Ids"
```

### Search APIs

##### SearchVehicle

SeaarchVehicle searches vehicles based on the given query.

gRPC:

```
rpc SearchVehicle(VehicleQuery) returns (VehicleResult) {}
```

HTTP:

```
Method: GET/POST
Path: api/biz/searchvehicle
Body: vehiclequery=[json format of VehicleQuery]
Return: {"code":0,msg:"success",redirect:"",data:[json format of VehicleResult]}
Errors:

```

##### SearchInvalidPlate

SearchInvalidPlate searches vehicles based on the given query and validate the results with DMV's database.

gRPC:

```
rpc SearchInvalidPlate(VehicleQuery) returns (VehicleResult) {}
```

HTTP:

```
Method: GET/POST
Path: api/biz/searchinvalidplate
Body: vehiclequery=[json format of VehicleQuery]
Return: {"code":0,msg:"success",redirect:"",data:[json format of VehicleResult]}
Errors:
   40202: No DMV database found.
```

##### SearchMultiSpatialTemporalAppearences

SearchMultiSpatialTemporalAppearences searches vehicles in multiple spatial/temporal queries and get the intersection of the results.

gRPC:

```
rpc SearchMultiSpatialTemporalAppearences(VehicleQueries) returns (VehicleResult) {}
```

HTTP:

```
Method: GET/POST
Path: api/biz/searchmultispatialtemporalappearences
Body: vehiclequeries=[json format of VehicleQueries]
Return: {"code":0,msg:"success",redirect:"",data:[json format of VehicleResult]}
Errors:

```

##### SearchMultiAppearences

SearchMultiAppearences searches vehicles in a single spatial/temporal queries and get the multiple appearence of the same vehicles.

gRPC:

```
rpc SearchMultiAppearences(VehicleQuery) returns (VehicleResult) {}
```

HTTP:

```
Method: GET/POST
Path: api/biz/searchmultiappearences
Body: vehiclequery=[json format of VehicleQuery]
Return: {"code":0,msg:"success",redirect:"",data:[json format of VehicleResult]}
Errors:

```

##### SearchSimilarPlate

SearchSimilarPlate searches vehicles that possess the similar plate with the plates in the given query.

gRPC:

```
rpc SearchSimilarPlate(VehicleQuery) returns (VehicleResult) {}
```

HTTP:

```
Method: GET/POST
Path: api/biz/searchsimilarplate
Body: vehiclequery=[json format of VehicleQuery]
Return: {"code":0,msg:"success",redirect:"",data:[json format of VehicleResult]}
Errors:
  40201: Required parameters missing.
  40203: Invalid plate format.
```

##### SearchFirstAppearence

SearchFirstAppearence searches vehicles that appeared in the given query and didn't appear for at least query.SilencePeriod before the appearence.

gRPC:

```
rpc SearchFirstAppearence(VehicleQuery) returns (VehicleResult) {}
```

HTTP:

```
Method: GET/POST
Path: api/biz/searchfirstappearence
Body: vehiclequery=[json format of VehicleQuery]
Return: {"code":0,msg:"success",redirect:"",data:[json format of VehicleResult]}
Errors:
  40201: Required parameters missing.
```

##### SearchFrequentAppearence

SearchFrequentAppearence searches vehicles that have more than MaxAppearenceTimes records within query.AppearencePeriod in the given query scope.

gRPC:

```
rpc SearchFrequentAppearence(VehicleQuery) returns (VehicleResult) {}
```

HTTP:

```
Method: GET/POST
Path: api/biz/searchfrequentappearence
Body: vehiclequery=[json format of VehicleQuery]
Return: {"code":0,msg:"success",redirect:"",data:[json format of VehicleResult]}
Errors:
  40201: Required parameters missing.
```

##### SearchFollowers

SearchFollowers searches vehicles that have followers for more than query.FollowingPeriod in the given query scope.

gRPC:

```
rpc SearchFollowers(VehicleQuery) returns (VehicleResult) {}
```

HTTP:

```
Method: GET/POST
Path: api/biz/searchfollowers
Body: vehiclequery=[json format of VehicleQuery]
Return: {"code":0,msg:"success",redirect:"",data:[json format of VehicleResult]}
Errors:
  40201: Required parameters missing.
```

##### SearchCustomizedQuery

SearchFollowers searches vehicles with customized query.

gRPC:

```
rpc SearchCustomizedQuery(CustomizedQuery) returns (VehicleResult) {}
```

HTTP:

```
Method: GET/POST
Path: api/biz/searchcustomizedquery
Body: customizedquery=[json format of CustomizedQuery]
Return: {"code":0,msg:"success",redirect:"",data:[json format of VehicleResult]}
Errors:
  40201: Required parameters missing.
```

##### AddCustomizedQuery

AddCustomizedQuery adds customized query.

gRPC:

```
rpc AddCustomizedQuery(CustomizedQuery) returns (NullMessage) {}
```

HTTP:

```
Method: GET/POST
Path: api/biz/addcustomizedquery
Body: customizedquery=[json format of CustomizedQuery]
Return: {"code":0,msg:"success",redirect:"",data:[json format of NullMessage]}
Errors:
  40201: Required parameters missing.
```

##### RemoveCustomizedQuery

RemoveCustomizedQuery removes customized query.

gRPC:

```
rpc RemoveCustomizedQuery(CustomizedQuery) returns (NullMessage) {}
```

HTTP:

```
Method: GET/POST
Path: api/biz/removecustomizedquery
Body: customizedquery=[json format of CustomizedQuery]
Return: {"code":0,msg:"success",redirect:"",data:[json format of NullMessage]}
Errors:
  40201: Required parameters missing.
  40204: No given query found.
```

##### GetCustomizedQuery

GetCustomizedQuery gets customized query.

gRPC:

```
rpc GetCustomizedQuery(CustomizedQuery) returns (CustomizedQuery) {}
```

HTTP:

```
Method: GET/POST
Path: api/biz/getcustomizedquery
Body: customizedquery=[json format of CustomizedQuery]
Return: {"code":0,msg:"success",redirect:"",data:[json format of CustomizedQuery]}
Errors:
  40201: Required parameters missing.
  40204: No given query found.
```

##### SetCustomizedQuery

SetCustomizedQuery sets customized query.

gRPC:

```
rpc SetCustomizedQuery(CustomizedQuery) returns (CustomizedQuery) {}
```

HTTP:

```
Method: GET/POST
Path: api/biz/setcustomizedquery
Body: customizedquery=[json format of CustomizedQuery]
Return: {"code":0,msg:"success",redirect:"",data:[json format of CustomizedQuery]}
Errors:
  40201: Required parameters missing.
  40204: No given query found.
```

##### GetCustomizedQueryList

GetCustomizedQueryList gets all customized queries.

gRPC:

```
rpc GetCustomizedQueryList(NullMessage) returns (CustomizedQueries) {}
```

HTTP:

```
Method: GET/POST
Path: api/biz/getcustomizedquerylist
Body: nullmessage=[json format of NullMessage]
Return: {"code":0,msg:"success",redirect:"",data:[json format of CustomizedQueries]}
Errors:
  40201: Required parameters missing.
```

### Index APIs

##### IndexVehicle

IndexVehicle stores vehicle entity.

gRPC:

```
rpc IndexVehicle(Vehicle) returns (NullMessage) {}
```

HTTP:

```
Method: GET/POST
Path: api/biz/indexvehicle
Body: vehicle=[json format of Vehicle]
Return: {"code":0,msg:"success",redirect:"",data:[json format of NullMessage]}
Errors:
  40301: Invalid entity.
```

##### IndexBatchVehicle

IndexBatchVehicle stores batch vehicle entities.

gRPC:

```
rpc IndexBatchVehicle(Vehicles) returns (NullMessage) {}
```

HTTP:

```
Method: GET/POST
Path: api/biz/indexbatchvehicle
Body: vehicles=[json format of Vehicles]
Return: {"code":0,msg:"success",redirect:"",data:[json format of NullMessage]}
Errors:
  40301: Invalid entity.
```

##### ImportImageRepo

ImportImageRepo stores external image repository.

gRPC:

```
rpc ImportImageRepo(ImageRepo) returns (NullMessage) {}
```

HTTP:

```
Method: GET/POST
Path: api/biz/importimagerepo
Body: imagerepo=[json format of ImageRepo]
Return: {"code":0,msg:"success",redirect:"",data:[json format of NullMessage]}
Errors:
  40302: Invalid repo.
```

##### RemoveImageRepo

RemoveImageRepo removes external image repository.

gRPC:

```
rpc RemoveImageRepo(ImageRepo) returns (NullMessage) {}
```

HTTP:

```
Method: GET/POST
Path: api/biz/removeimagerepo
Body: imagerepo=[json format of ImageRepo]
Return: {"code":0,msg:"success",redirect:"",data:[json format of NullMessage]}
Errors:
  40303: Repo doesn't exist.
```

##### ModifyImageRepo

ModifyImageRepo modifies external image repository.

gRPC:

```
rpc ModifyImageRepo(ImageRepo) returns (NullMessage) {}
```

HTTP:

```
Method: GET/POST
Path: api/biz/modifyimagerepo
Body: imagerepo=[json format of ImageRepo]
Return: {"code":0,msg:"success",redirect:"",data:[json format of NullMessage]}
Errors:
  40303: Repo doesn't exist.
```

##### GetAllImageRepo

GetAllImageRepo lists all external image repositories.

gRPC:

```
rpc GetAllImageRepo(NullMessage) returns (ImageRepos) {}
```

HTTP:

```
Method: GET/POST
Path: api/biz/getallimagerepo
Body: nullmessage=[json format of NullMessage]
Return: {"code":0,msg:"success",redirect:"",data:[json format of ImageRepos]}
Errors:

```

##### ImportExternalData

ImportExternalData imports external data like DMV's data.

gRPC:

```
rpc ImportExternalData(ExternalData) returns (NullMessage) {}
```

HTTP:

```
Method: GET/POST
Path: api/biz/importexternaldata
Body: externaldata=[json format of ExternalData]
Return: {"code":0,msg:"success",redirect:"",data:[json format of NullMessage]}
Errors:
  40304: Invalid external data.
```

## Related Data Structure
```
enum Lang {
	UNKNOWN_LANG = 0;
	EN_US = 1;
	ZH_CN = 2;
}
```

```
enum MappingType {
	UNKNOWN_MAPPING = 0;
	BIZ = 1;
	BRAND = 2;
	SUBBRAND = 3;
	MODELYEAR = 4;
	MODELTYPE = 5;
	COLOR = 6;
	PLATETYPE = 7;
	SIDES = 8;
}
```

```
message Text {
	string Text = 1;
	Lang Lang = 2;
}
```

```
message TextArray {
	repeated string Text = 1;
	Lang Lang = 2;
}
```

```
message IdTextMapping {
	MappingType Type = 1;
	map<int64, TextArray> Mapping = 2;
}
```

```
message IdTextMappingChange {
	int32 MappingChanged = 1;
}
```

```
message ImageRepo {
	int32 Id = 1;
	string Name = 2;
	// TODO
}
```

```
message ImageRepos {
	repeated ImageRepo ImageRepo = 1;
}
```

```
message ExternalData {
	// TODO
}
```

