# V0044Step

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Time** | Pointer to [**V0044StepTime**](V0044StepTime.md) |  | [optional] 
**ExitCode** | Pointer to [**V0044ProcessExitCodeVerbose**](V0044ProcessExitCodeVerbose.md) |  | [optional] 
**Nodes** | Pointer to [**V0044StepNodes**](V0044StepNodes.md) |  | [optional] 
**Tasks** | Pointer to [**V0044StepTasks**](V0044StepTasks.md) |  | [optional] 
**Pid** | Pointer to **string** | Deprecated; Process ID | [optional] 
**CPU** | Pointer to [**V0044StepCPU**](V0044StepCPU.md) |  | [optional] 
**KillRequestUser** | Pointer to **string** | User ID that requested termination of the step | [optional] 
**State** | Pointer to **[]string** | Current state | [optional] 
**Statistics** | Pointer to [**V0044StepStatistics**](V0044StepStatistics.md) |  | [optional] 
**Step** | Pointer to [**V0044StepStep**](V0044StepStep.md) |  | [optional] 
**Task** | Pointer to [**V0044StepTask**](V0044StepTask.md) |  | [optional] 
**Tres** | Pointer to [**V0044StepTres**](V0044StepTres.md) |  | [optional] 

## Methods

### NewV0044Step

`func NewV0044Step() *V0044Step`

NewV0044Step instantiates a new V0044Step object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0044StepWithDefaults

`func NewV0044StepWithDefaults() *V0044Step`

NewV0044StepWithDefaults instantiates a new V0044Step object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTime

`func (o *V0044Step) GetTime() V0044StepTime`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *V0044Step) GetTimeOk() (*V0044StepTime, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *V0044Step) SetTime(v V0044StepTime)`

SetTime sets Time field to given value.

### HasTime

`func (o *V0044Step) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetExitCode

`func (o *V0044Step) GetExitCode() V0044ProcessExitCodeVerbose`

GetExitCode returns the ExitCode field if non-nil, zero value otherwise.

### GetExitCodeOk

`func (o *V0044Step) GetExitCodeOk() (*V0044ProcessExitCodeVerbose, bool)`

GetExitCodeOk returns a tuple with the ExitCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitCode

`func (o *V0044Step) SetExitCode(v V0044ProcessExitCodeVerbose)`

SetExitCode sets ExitCode field to given value.

### HasExitCode

`func (o *V0044Step) HasExitCode() bool`

HasExitCode returns a boolean if a field has been set.

### GetNodes

`func (o *V0044Step) GetNodes() V0044StepNodes`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *V0044Step) GetNodesOk() (*V0044StepNodes, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *V0044Step) SetNodes(v V0044StepNodes)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *V0044Step) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetTasks

`func (o *V0044Step) GetTasks() V0044StepTasks`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *V0044Step) GetTasksOk() (*V0044StepTasks, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *V0044Step) SetTasks(v V0044StepTasks)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *V0044Step) HasTasks() bool`

HasTasks returns a boolean if a field has been set.

### GetPid

`func (o *V0044Step) GetPid() string`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *V0044Step) GetPidOk() (*string, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *V0044Step) SetPid(v string)`

SetPid sets Pid field to given value.

### HasPid

`func (o *V0044Step) HasPid() bool`

HasPid returns a boolean if a field has been set.

### GetCPU

`func (o *V0044Step) GetCPU() V0044StepCPU`

GetCPU returns the CPU field if non-nil, zero value otherwise.

### GetCPUOk

`func (o *V0044Step) GetCPUOk() (*V0044StepCPU, bool)`

GetCPUOk returns a tuple with the CPU field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCPU

`func (o *V0044Step) SetCPU(v V0044StepCPU)`

SetCPU sets CPU field to given value.

### HasCPU

`func (o *V0044Step) HasCPU() bool`

HasCPU returns a boolean if a field has been set.

### GetKillRequestUser

`func (o *V0044Step) GetKillRequestUser() string`

GetKillRequestUser returns the KillRequestUser field if non-nil, zero value otherwise.

### GetKillRequestUserOk

`func (o *V0044Step) GetKillRequestUserOk() (*string, bool)`

GetKillRequestUserOk returns a tuple with the KillRequestUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKillRequestUser

`func (o *V0044Step) SetKillRequestUser(v string)`

SetKillRequestUser sets KillRequestUser field to given value.

### HasKillRequestUser

`func (o *V0044Step) HasKillRequestUser() bool`

HasKillRequestUser returns a boolean if a field has been set.

### GetState

`func (o *V0044Step) GetState() []string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *V0044Step) GetStateOk() (*[]string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *V0044Step) SetState(v []string)`

SetState sets State field to given value.

### HasState

`func (o *V0044Step) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStatistics

`func (o *V0044Step) GetStatistics() V0044StepStatistics`

GetStatistics returns the Statistics field if non-nil, zero value otherwise.

### GetStatisticsOk

`func (o *V0044Step) GetStatisticsOk() (*V0044StepStatistics, bool)`

GetStatisticsOk returns a tuple with the Statistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatistics

`func (o *V0044Step) SetStatistics(v V0044StepStatistics)`

SetStatistics sets Statistics field to given value.

### HasStatistics

`func (o *V0044Step) HasStatistics() bool`

HasStatistics returns a boolean if a field has been set.

### GetStep

`func (o *V0044Step) GetStep() V0044StepStep`

GetStep returns the Step field if non-nil, zero value otherwise.

### GetStepOk

`func (o *V0044Step) GetStepOk() (*V0044StepStep, bool)`

GetStepOk returns a tuple with the Step field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStep

`func (o *V0044Step) SetStep(v V0044StepStep)`

SetStep sets Step field to given value.

### HasStep

`func (o *V0044Step) HasStep() bool`

HasStep returns a boolean if a field has been set.

### GetTask

`func (o *V0044Step) GetTask() V0044StepTask`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *V0044Step) GetTaskOk() (*V0044StepTask, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *V0044Step) SetTask(v V0044StepTask)`

SetTask sets Task field to given value.

### HasTask

`func (o *V0044Step) HasTask() bool`

HasTask returns a boolean if a field has been set.

### GetTres

`func (o *V0044Step) GetTres() V0044StepTres`

GetTres returns the Tres field if non-nil, zero value otherwise.

### GetTresOk

`func (o *V0044Step) GetTresOk() (*V0044StepTres, bool)`

GetTresOk returns a tuple with the Tres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTres

`func (o *V0044Step) SetTres(v V0044StepTres)`

SetTres sets Tres field to given value.

### HasTres

`func (o *V0044Step) HasTres() bool`

HasTres returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


