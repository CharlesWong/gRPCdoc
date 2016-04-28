## Computation API

### Query/Batch Computation

##### GetAllBatchTasks

GetAllBatchTasks gets all batch tasks (running and stopped tasks).

gRPC:

```
rpc GetAllBatchTasks(NullMessage) returns (BatchTasks) {}
```

HTTP:

```
Method: GET/POST
Path: api/biz/getallbatchtasks
Body: nullmessage=[json format of NullMessage]
Return: {"code":0,msg:"success",redirect:"",data:[json format of BatchTasks]}
Errors:

```

##### AddBatchTask

AddBatchTask adds new task into task repo.
No need to assign task Id in client side.

gRPC:

```
rpc AddBatchTask(BatchTask) returns (BatchTaskStatus) {}
```

HTTP:

```
Method: GET/POST
Path: api/biz/addbatchtask
Body: batchtask=[json format of BatchTask]
Return: {"code":0,msg:"success",redirect:"",data:[json format of BatchTaskStatus]}
Errors:
  40301: Invalid task.
```

##### RemoveBatchTask

RemoveBatchTask removes existing task of task.Id from task repo.

gRPC:

```
rpc RemoveBatchTask(BatchTask) returns (BatchTaskStatus) {}
```

HTTP:

```
Method: GET/POST
Path: api/biz/removebatchtask
Body: batchtask=[json format of BatchTask]
Return: {"code":0,msg:"success",redirect:"",data:[json format of BatchTaskStatus]}
Errors:
  40302: Task doesn't exist.
```

##### SubmitBatchTask

SubmitBatchTask submits existing task of task.Id from task repo and stages it for execution.

gRPC:

```
rpc SubmitBatchTask(BatchTask) returns (BatchTaskStatus) {}
```

HTTP:

```
Method: GET/POST
Path: api/biz/submitbatchtask
Body: batchtask=[json format of BatchTask]
Return: {"code":0,msg:"success",redirect:"",data:[json format of BatchTaskStatus]}
Errors:
  40302: Task doesn't exist.
  40303: Insufficient resources.
```

##### CancelBatchTask

CancelBatchTask cancels existing running task of task.Id.

gRPC:

```
rpc CancelBatchTask(BatchTask) returns (BatchTaskStatus) {}
```

HTTP:

```
Method: GET/POST
Path: api/biz/cancelbatchtask
Body: batchtask=[json format of BatchTask]
Return: {"code":0,msg:"success",redirect:"",data:[json format of BatchTaskStatus]}
Errors:
  40302: Task doesn't exist.
  40304: Task isn't running.
```

##### GetBatchTaskStatus

GetBatchTaskStatus get status of existing task of task.Id.

gRPC:

```
rpc GetBatchTaskStatus(BatchTask) returns (BatchTaskStatus) {}
```

HTTP:

```
Method: GET/POST
Path: api/biz/getbatchtaskstatus
Body: batchtask=[json format of BatchTask]
Return: {"code":0,msg:"success",redirect:"",data:[json format of BatchTaskStatus]}
Errors:
  40302: Task doesn't exist.
```

### Alerts/Streaming Computation

##### GetAllStreamingTasks

GetAllStreamingTasks gets all streaming tasks (running and stopped tasks).

gRPC:

```
rpc GetAllStreamingTasks(NullMessage) returns (StreamingTasks) {}
```

HTTP:

```
Method: GET/POST
Path: api/biz/getallstreamingtasks
Body: nullmessage=[json format of NullMessage]
Return: {"code":0,msg:"success",redirect:"",data:[json format of StreamingTasks]}
Errors:

```

##### AddStreamingTask

AddStreamingTask adds new task into task repo.
No need to assign task Id in client side.

gRPC:

```
rpc AddStreamingTask(StreamingTask) returns (StreamingTaskStatus) {}
```

HTTP:

```
Method: GET/POST
Path: api/biz/addstreamingtask
Body: streamingtask=[json format of StreamingTask]
Return: {"code":0,msg:"success",redirect:"",data:[json format of StreamingTaskStatus]}
Errors:
  40301: Invalid task.
```

##### RemoveStreamingTask

RemoveStreamingTask removes existing task of task.Id from task repo.

gRPC:

```
rpc RemoveStreamingTask(StreamingTask) returns (StreamingTaskStatus) {}
```

HTTP:

```
Method: GET/POST
Path: api/biz/removestreamingtask
Body: streamingtask=[json format of StreamingTask]
Return: {"code":0,msg:"success",redirect:"",data:[json format of StreamingTaskStatus]}
Errors:
  40302: Task doesn't exist.
```

##### SubmitStreamingTask

SubmitStreamingTask submits existing task of task.Id from task repo and stages it for execution.

gRPC:

```
rpc SubmitStreamingTask(StreamingTask) returns (StreamingTaskStatus) {}
```

HTTP:

```
Method: GET/POST
Path: api/biz/submitstreamingtask
Body: streamingtask=[json format of StreamingTask]
Return: {"code":0,msg:"success",redirect:"",data:[json format of StreamingTaskStatus]}
Errors:
  40302: Task doesn't exist.
  40303: Insufficient resources.
```

##### CancelStreamingTask

CancelStreamingTask cancels existing running task of task.Id.

gRPC:

```
rpc CancelStreamingTask(StreamingTask) returns (StreamingTaskStatus) {}
```

HTTP:

```
Method: GET/POST
Path: api/biz/cancelstreamingtask
Body: streamingtask=[json format of StreamingTask]
Return: {"code":0,msg:"success",redirect:"",data:[json format of StreamingTaskStatus]}
Errors:
  40302: Task doesn't exist.
  40304: Task isn't running.
```

##### GetStreamingTaskStatus

GetStreamingTaskStatus get status of existing task of task.Id.

gRPC:

```
rpc GetStreamingTaskStatus(StreamingTask) returns (StreamingTaskStatus) {}
```

HTTP:

```
Method: GET/POST
Path: api/biz/getstreamingtaskstatus
Body: streamingtask=[json format of StreamingTask]
Return: {"code":0,msg:"success",redirect:"",data:[json format of StreamingTaskStatus]}
Errors:
  40302: Task doesn't exist.
```

## Related Data Structure
```
message Resource {
	int32 Core = 1;
	int32 Mem = 2;
	int32 Disk = 3;
	int32 Bandwidth = 4;

	// TODO: added other resources.
}
```

```
enum TaskStatus {
	UNKNOWN_TASKSTATUS = 0;
	STAGED = 1;
	RUNNING = 2;
	FINISHED = 3;
	FAILED = 4;
}
```

```
message BaseTask {
	int32 Category = 1;
	int32 Id = 2;
	repeated int32 Owner = 3;
	repeated int32 Access = 4;
	Resource ResReq = 5;
}
```

```
message BatchQuery {
	string SqlQuery = 1;
	// TODO
}
```

```
message BatchTask {
	BaseTask BaseTask = 1;
	BatchQuery BatchQuery = 2;
	BatchTaskStatus BatchTaskStatus = 3; 
}
```

```
message BatchTasks {
	repeated BatchTask BatchTask = 1;
}
```

```
message BatchTaskStatus {
	TaskStatus TaskStatus = 1;
	int32 TotalTask = 2;
	int32 FinishedTask = 3;
	int32 RunningTask = 4;
	int32 FailedTask = 5;
	// TODO
}
```

```
message StreamingTopology {
	// TODO
}
```

```
message StreamingTask {
	BaseTask BaseTask = 1;
	StreamingTopology StreamingTopology = 2;
	StreamingTaskStatus StreamingTaskStatus = 3;
}
```

```
message StreamingTasks {
	repeated StreamingTask StreamingTask = 1;
}
```

```
message StreamingTaskStatus {
	TaskStatus TaskStatus = 1;
	int32 FinishedTask = 2;
	int32 RunningTask = 3;
	int32 FailedTask = 4;
	// TODO
}
```

