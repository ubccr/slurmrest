# V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Time** | Pointer to [**V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTime**](V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTime.md) |  | [optional] 
**ExitCode** | Pointer to [**V0041OpenapiSlurmdbdJobsRespJobsInnerDerivedExitCode**](V0041OpenapiSlurmdbdJobsRespJobsInnerDerivedExitCode.md) |  | [optional] 
**Nodes** | Pointer to [**V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerNodes**](V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerNodes.md) |  | [optional] 
**Tasks** | Pointer to [**V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTasks**](V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTasks.md) |  | [optional] 
**Pid** | Pointer to **string** |  | [optional] 
**CPU** | Pointer to [**V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerCPU**](V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerCPU.md) |  | [optional] 
**KillRequestUser** | Pointer to **string** |  | [optional] 
**State** | Pointer to **[]string** |  | [optional] 
**Statistics** | Pointer to [**V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerStatistics**](V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerStatistics.md) |  | [optional] 
**Step** | Pointer to [**V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerStep**](V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerStep.md) |  | [optional] 
**Task** | Pointer to [**V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTask**](V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTask.md) |  | [optional] 
**Tres** | Pointer to [**V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTres**](V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTres.md) |  | [optional] 

## Methods

### NewV0041OpenapiSlurmdbdJobsRespJobsInnerStepsInner

`func NewV0041OpenapiSlurmdbdJobsRespJobsInnerStepsInner() *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInner`

NewV0041OpenapiSlurmdbdJobsRespJobsInnerStepsInner instantiates a new V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerWithDefaults

`func NewV0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerWithDefaults() *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInner`

NewV0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerWithDefaults instantiates a new V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTime

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInner) GetTime() V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTime`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInner) GetTimeOk() (*V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTime, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInner) SetTime(v V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTime)`

SetTime sets Time field to given value.

### HasTime

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInner) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetExitCode

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInner) GetExitCode() V0041OpenapiSlurmdbdJobsRespJobsInnerDerivedExitCode`

GetExitCode returns the ExitCode field if non-nil, zero value otherwise.

### GetExitCodeOk

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInner) GetExitCodeOk() (*V0041OpenapiSlurmdbdJobsRespJobsInnerDerivedExitCode, bool)`

GetExitCodeOk returns a tuple with the ExitCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitCode

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInner) SetExitCode(v V0041OpenapiSlurmdbdJobsRespJobsInnerDerivedExitCode)`

SetExitCode sets ExitCode field to given value.

### HasExitCode

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInner) HasExitCode() bool`

HasExitCode returns a boolean if a field has been set.

### GetNodes

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInner) GetNodes() V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerNodes`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInner) GetNodesOk() (*V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerNodes, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInner) SetNodes(v V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerNodes)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInner) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetTasks

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInner) GetTasks() V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTasks`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInner) GetTasksOk() (*V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTasks, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInner) SetTasks(v V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTasks)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInner) HasTasks() bool`

HasTasks returns a boolean if a field has been set.

### GetPid

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInner) GetPid() string`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInner) GetPidOk() (*string, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInner) SetPid(v string)`

SetPid sets Pid field to given value.

### HasPid

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInner) HasPid() bool`

HasPid returns a boolean if a field has been set.

### GetCPU

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInner) GetCPU() V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerCPU`

GetCPU returns the CPU field if non-nil, zero value otherwise.

### GetCPUOk

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInner) GetCPUOk() (*V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerCPU, bool)`

GetCPUOk returns a tuple with the CPU field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCPU

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInner) SetCPU(v V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerCPU)`

SetCPU sets CPU field to given value.

### HasCPU

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInner) HasCPU() bool`

HasCPU returns a boolean if a field has been set.

### GetKillRequestUser

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInner) GetKillRequestUser() string`

GetKillRequestUser returns the KillRequestUser field if non-nil, zero value otherwise.

### GetKillRequestUserOk

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInner) GetKillRequestUserOk() (*string, bool)`

GetKillRequestUserOk returns a tuple with the KillRequestUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKillRequestUser

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInner) SetKillRequestUser(v string)`

SetKillRequestUser sets KillRequestUser field to given value.

### HasKillRequestUser

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInner) HasKillRequestUser() bool`

HasKillRequestUser returns a boolean if a field has been set.

### GetState

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInner) GetState() []string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInner) GetStateOk() (*[]string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInner) SetState(v []string)`

SetState sets State field to given value.

### HasState

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInner) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStatistics

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInner) GetStatistics() V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerStatistics`

GetStatistics returns the Statistics field if non-nil, zero value otherwise.

### GetStatisticsOk

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInner) GetStatisticsOk() (*V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerStatistics, bool)`

GetStatisticsOk returns a tuple with the Statistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatistics

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInner) SetStatistics(v V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerStatistics)`

SetStatistics sets Statistics field to given value.

### HasStatistics

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInner) HasStatistics() bool`

HasStatistics returns a boolean if a field has been set.

### GetStep

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInner) GetStep() V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerStep`

GetStep returns the Step field if non-nil, zero value otherwise.

### GetStepOk

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInner) GetStepOk() (*V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerStep, bool)`

GetStepOk returns a tuple with the Step field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStep

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInner) SetStep(v V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerStep)`

SetStep sets Step field to given value.

### HasStep

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInner) HasStep() bool`

HasStep returns a boolean if a field has been set.

### GetTask

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInner) GetTask() V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTask`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInner) GetTaskOk() (*V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTask, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInner) SetTask(v V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTask)`

SetTask sets Task field to given value.

### HasTask

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInner) HasTask() bool`

HasTask returns a boolean if a field has been set.

### GetTres

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInner) GetTres() V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTres`

GetTres returns the Tres field if non-nil, zero value otherwise.

### GetTresOk

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInner) GetTresOk() (*V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTres, bool)`

GetTresOk returns a tuple with the Tres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTres

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInner) SetTres(v V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTres)`

SetTres sets Tres field to given value.

### HasTres

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInner) HasTres() bool`

HasTres returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


