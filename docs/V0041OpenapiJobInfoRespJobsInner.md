# V0041OpenapiJobInfoRespJobsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** |  | [optional] 
**AccrueTime** | Pointer to [**V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu**](V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu.md) |  | [optional] 
**AdminComment** | Pointer to **string** |  | [optional] 
**AllocatingNode** | Pointer to **string** |  | [optional] 
**ArrayJobId** | Pointer to [**V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId**](V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId.md) |  | [optional] 
**ArrayTaskId** | Pointer to [**V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId**](V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId.md) |  | [optional] 
**ArrayMaxTasks** | Pointer to [**V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId**](V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId.md) |  | [optional] 
**ArrayTaskString** | Pointer to **string** |  | [optional] 
**AssociationId** | Pointer to **int32** |  | [optional] 
**BatchFeatures** | Pointer to **string** |  | [optional] 
**BatchFlag** | Pointer to **bool** |  | [optional] 
**BatchHost** | Pointer to **string** |  | [optional] 
**Flags** | Pointer to **[]string** |  | [optional] 
**BurstBuffer** | Pointer to **string** |  | [optional] 
**BurstBufferState** | Pointer to **string** |  | [optional] 
**Cluster** | Pointer to **string** |  | [optional] 
**ClusterFeatures** | Pointer to **string** |  | [optional] 
**Command** | Pointer to **string** |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 
**Container** | Pointer to **string** |  | [optional] 
**ContainerId** | Pointer to **string** |  | [optional] 
**Contiguous** | Pointer to **bool** |  | [optional] 
**CoreSpec** | Pointer to **int32** |  | [optional] 
**ThreadSpec** | Pointer to **int32** |  | [optional] 
**CoresPerSocket** | Pointer to [**V0041JobDescMsgDistributionPlaneSize**](V0041JobDescMsgDistributionPlaneSize.md) |  | [optional] 
**BillableTres** | Pointer to [**V0041OpenapiSlurmdbdConfigRespQosInnerLimitsFactor**](V0041OpenapiSlurmdbdConfigRespQosInnerLimitsFactor.md) |  | [optional] 
**CpusPerTask** | Pointer to [**V0041JobDescMsgDistributionPlaneSize**](V0041JobDescMsgDistributionPlaneSize.md) |  | [optional] 
**CpuFrequencyMinimum** | Pointer to [**V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId**](V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId.md) |  | [optional] 
**CpuFrequencyMaximum** | Pointer to [**V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId**](V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId.md) |  | [optional] 
**CpuFrequencyGovernor** | Pointer to [**V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId**](V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId.md) |  | [optional] 
**CpusPerTres** | Pointer to **string** |  | [optional] 
**Cron** | Pointer to **string** |  | [optional] 
**Deadline** | Pointer to [**V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu**](V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu.md) |  | [optional] 
**DelayBoot** | Pointer to [**V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId**](V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId.md) |  | [optional] 
**Dependency** | Pointer to **string** |  | [optional] 
**DerivedExitCode** | Pointer to [**V0041OpenapiSlurmdbdJobsRespJobsInnerDerivedExitCode**](V0041OpenapiSlurmdbdJobsRespJobsInnerDerivedExitCode.md) |  | [optional] 
**EligibleTime** | Pointer to [**V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu**](V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu.md) |  | [optional] 
**EndTime** | Pointer to [**V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu**](V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu.md) |  | [optional] 
**ExcludedNodes** | Pointer to **string** |  | [optional] 
**ExitCode** | Pointer to [**V0041OpenapiSlurmdbdJobsRespJobsInnerDerivedExitCode**](V0041OpenapiSlurmdbdJobsRespJobsInnerDerivedExitCode.md) |  | [optional] 
**Extra** | Pointer to **string** |  | [optional] 
**FailedNode** | Pointer to **string** |  | [optional] 
**Features** | Pointer to **string** |  | [optional] 
**FederationOrigin** | Pointer to **string** |  | [optional] 
**FederationSiblingsActive** | Pointer to **string** |  | [optional] 
**FederationSiblingsViable** | Pointer to **string** |  | [optional] 
**GresDetail** | Pointer to **[]string** |  | [optional] 
**GroupId** | Pointer to **int32** |  | [optional] 
**GroupName** | Pointer to **string** |  | [optional] 
**HetJobId** | Pointer to [**V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId**](V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId.md) |  | [optional] 
**HetJobIdSet** | Pointer to **string** |  | [optional] 
**HetJobOffset** | Pointer to [**V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId**](V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId.md) |  | [optional] 
**JobId** | Pointer to **int32** |  | [optional] 
**JobResources** | Pointer to [**V0041OpenapiJobInfoRespJobsInnerJobResources**](V0041OpenapiJobInfoRespJobsInnerJobResources.md) |  | [optional] 
**JobSizeStr** | Pointer to **[]string** |  | [optional] 
**JobState** | Pointer to **[]string** |  | [optional] 
**LastSchedEvaluation** | Pointer to [**V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu**](V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu.md) |  | [optional] 
**Licenses** | Pointer to **string** |  | [optional] 
**MailType** | Pointer to **[]string** |  | [optional] 
**MailUser** | Pointer to **string** |  | [optional] 
**MaxCpus** | Pointer to [**V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId**](V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId.md) |  | [optional] 
**MaxNodes** | Pointer to [**V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId**](V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId.md) |  | [optional] 
**McsLabel** | Pointer to **string** |  | [optional] 
**MemoryPerTres** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Network** | Pointer to **string** |  | [optional] 
**Nodes** | Pointer to **string** |  | [optional] 
**Nice** | Pointer to **int32** |  | [optional] 
**TasksPerCore** | Pointer to [**V0041JobDescMsgDistributionPlaneSize**](V0041JobDescMsgDistributionPlaneSize.md) |  | [optional] 
**TasksPerTres** | Pointer to [**V0041JobDescMsgDistributionPlaneSize**](V0041JobDescMsgDistributionPlaneSize.md) |  | [optional] 
**TasksPerNode** | Pointer to [**V0041JobDescMsgDistributionPlaneSize**](V0041JobDescMsgDistributionPlaneSize.md) |  | [optional] 
**TasksPerSocket** | Pointer to [**V0041JobDescMsgDistributionPlaneSize**](V0041JobDescMsgDistributionPlaneSize.md) |  | [optional] 
**TasksPerBoard** | Pointer to [**V0041JobDescMsgDistributionPlaneSize**](V0041JobDescMsgDistributionPlaneSize.md) |  | [optional] 
**Cpus** | Pointer to [**V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId**](V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId.md) |  | [optional] 
**NodeCount** | Pointer to [**V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId**](V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId.md) |  | [optional] 
**Tasks** | Pointer to [**V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId**](V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId.md) |  | [optional] 
**Partition** | Pointer to **string** |  | [optional] 
**Prefer** | Pointer to **string** |  | [optional] 
**MemoryPerCpu** | Pointer to [**V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu**](V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu.md) |  | [optional] 
**MemoryPerNode** | Pointer to [**V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu**](V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu.md) |  | [optional] 
**MinimumCpusPerNode** | Pointer to [**V0041JobDescMsgDistributionPlaneSize**](V0041JobDescMsgDistributionPlaneSize.md) |  | [optional] 
**MinimumTmpDiskPerNode** | Pointer to [**V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId**](V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId.md) |  | [optional] 
**Power** | Pointer to [**V0041OpenapiJobInfoRespJobsInnerPower**](V0041OpenapiJobInfoRespJobsInnerPower.md) |  | [optional] 
**PreemptTime** | Pointer to [**V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu**](V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu.md) |  | [optional] 
**PreemptableTime** | Pointer to [**V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu**](V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu.md) |  | [optional] 
**PreSusTime** | Pointer to [**V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu**](V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu.md) |  | [optional] 
**Hold** | Pointer to **bool** | Hold (true) or release (false) job | [optional] 
**Priority** | Pointer to [**V0041OpenapiSlurmdbdJobsRespJobsInnerPriority**](V0041OpenapiSlurmdbdJobsRespJobsInnerPriority.md) |  | [optional] 
**Profile** | Pointer to **[]string** |  | [optional] 
**Qos** | Pointer to **string** |  | [optional] 
**Reboot** | Pointer to **bool** |  | [optional] 
**RequiredNodes** | Pointer to **string** |  | [optional] 
**MinimumSwitches** | Pointer to **int32** |  | [optional] 
**Requeue** | Pointer to **bool** |  | [optional] 
**ResizeTime** | Pointer to [**V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu**](V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu.md) |  | [optional] 
**RestartCnt** | Pointer to **int32** |  | [optional] 
**ResvName** | Pointer to **string** |  | [optional] 
**ScheduledNodes** | Pointer to **string** |  | [optional] 
**SelinuxContext** | Pointer to **string** |  | [optional] 
**Shared** | Pointer to **[]string** |  | [optional] 
**Exclusive** | Pointer to **[]string** |  | [optional] 
**Oversubscribe** | Pointer to **bool** |  | [optional] 
**ShowFlags** | Pointer to **[]string** |  | [optional] 
**SocketsPerBoard** | Pointer to **int32** |  | [optional] 
**SocketsPerNode** | Pointer to [**V0041JobDescMsgDistributionPlaneSize**](V0041JobDescMsgDistributionPlaneSize.md) |  | [optional] 
**StartTime** | Pointer to [**V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu**](V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu.md) |  | [optional] 
**StateDescription** | Pointer to **string** |  | [optional] 
**StateReason** | Pointer to **string** |  | [optional] 
**StandardError** | Pointer to **string** |  | [optional] 
**StandardInput** | Pointer to **string** |  | [optional] 
**StandardOutput** | Pointer to **string** |  | [optional] 
**SubmitTime** | Pointer to [**V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu**](V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu.md) |  | [optional] 
**SuspendTime** | Pointer to [**V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu**](V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu.md) |  | [optional] 
**SystemComment** | Pointer to **string** |  | [optional] 
**TimeLimit** | Pointer to [**V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId**](V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId.md) |  | [optional] 
**TimeMinimum** | Pointer to [**V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId**](V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId.md) |  | [optional] 
**ThreadsPerCore** | Pointer to [**V0041JobDescMsgDistributionPlaneSize**](V0041JobDescMsgDistributionPlaneSize.md) |  | [optional] 
**TresBind** | Pointer to **string** |  | [optional] 
**TresFreq** | Pointer to **string** |  | [optional] 
**TresPerJob** | Pointer to **string** |  | [optional] 
**TresPerNode** | Pointer to **string** |  | [optional] 
**TresPerSocket** | Pointer to **string** |  | [optional] 
**TresPerTask** | Pointer to **string** |  | [optional] 
**TresReqStr** | Pointer to **string** |  | [optional] 
**TresAllocStr** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **int32** |  | [optional] 
**UserName** | Pointer to **string** |  | [optional] 
**MaximumSwitchWaitTime** | Pointer to **int32** |  | [optional] 
**Wckey** | Pointer to **string** |  | [optional] 
**CurrentWorkingDirectory** | Pointer to **string** |  | [optional] 

## Methods

### NewV0041OpenapiJobInfoRespJobsInner

`func NewV0041OpenapiJobInfoRespJobsInner() *V0041OpenapiJobInfoRespJobsInner`

NewV0041OpenapiJobInfoRespJobsInner instantiates a new V0041OpenapiJobInfoRespJobsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0041OpenapiJobInfoRespJobsInnerWithDefaults

`func NewV0041OpenapiJobInfoRespJobsInnerWithDefaults() *V0041OpenapiJobInfoRespJobsInner`

NewV0041OpenapiJobInfoRespJobsInnerWithDefaults instantiates a new V0041OpenapiJobInfoRespJobsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *V0041OpenapiJobInfoRespJobsInner) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *V0041OpenapiJobInfoRespJobsInner) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *V0041OpenapiJobInfoRespJobsInner) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetAccrueTime

`func (o *V0041OpenapiJobInfoRespJobsInner) GetAccrueTime() V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu`

GetAccrueTime returns the AccrueTime field if non-nil, zero value otherwise.

### GetAccrueTimeOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetAccrueTimeOk() (*V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu, bool)`

GetAccrueTimeOk returns a tuple with the AccrueTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccrueTime

`func (o *V0041OpenapiJobInfoRespJobsInner) SetAccrueTime(v V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu)`

SetAccrueTime sets AccrueTime field to given value.

### HasAccrueTime

`func (o *V0041OpenapiJobInfoRespJobsInner) HasAccrueTime() bool`

HasAccrueTime returns a boolean if a field has been set.

### GetAdminComment

`func (o *V0041OpenapiJobInfoRespJobsInner) GetAdminComment() string`

GetAdminComment returns the AdminComment field if non-nil, zero value otherwise.

### GetAdminCommentOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetAdminCommentOk() (*string, bool)`

GetAdminCommentOk returns a tuple with the AdminComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminComment

`func (o *V0041OpenapiJobInfoRespJobsInner) SetAdminComment(v string)`

SetAdminComment sets AdminComment field to given value.

### HasAdminComment

`func (o *V0041OpenapiJobInfoRespJobsInner) HasAdminComment() bool`

HasAdminComment returns a boolean if a field has been set.

### GetAllocatingNode

`func (o *V0041OpenapiJobInfoRespJobsInner) GetAllocatingNode() string`

GetAllocatingNode returns the AllocatingNode field if non-nil, zero value otherwise.

### GetAllocatingNodeOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetAllocatingNodeOk() (*string, bool)`

GetAllocatingNodeOk returns a tuple with the AllocatingNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatingNode

`func (o *V0041OpenapiJobInfoRespJobsInner) SetAllocatingNode(v string)`

SetAllocatingNode sets AllocatingNode field to given value.

### HasAllocatingNode

`func (o *V0041OpenapiJobInfoRespJobsInner) HasAllocatingNode() bool`

HasAllocatingNode returns a boolean if a field has been set.

### GetArrayJobId

`func (o *V0041OpenapiJobInfoRespJobsInner) GetArrayJobId() V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId`

GetArrayJobId returns the ArrayJobId field if non-nil, zero value otherwise.

### GetArrayJobIdOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetArrayJobIdOk() (*V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId, bool)`

GetArrayJobIdOk returns a tuple with the ArrayJobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrayJobId

`func (o *V0041OpenapiJobInfoRespJobsInner) SetArrayJobId(v V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId)`

SetArrayJobId sets ArrayJobId field to given value.

### HasArrayJobId

`func (o *V0041OpenapiJobInfoRespJobsInner) HasArrayJobId() bool`

HasArrayJobId returns a boolean if a field has been set.

### GetArrayTaskId

`func (o *V0041OpenapiJobInfoRespJobsInner) GetArrayTaskId() V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId`

GetArrayTaskId returns the ArrayTaskId field if non-nil, zero value otherwise.

### GetArrayTaskIdOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetArrayTaskIdOk() (*V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId, bool)`

GetArrayTaskIdOk returns a tuple with the ArrayTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrayTaskId

`func (o *V0041OpenapiJobInfoRespJobsInner) SetArrayTaskId(v V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId)`

SetArrayTaskId sets ArrayTaskId field to given value.

### HasArrayTaskId

`func (o *V0041OpenapiJobInfoRespJobsInner) HasArrayTaskId() bool`

HasArrayTaskId returns a boolean if a field has been set.

### GetArrayMaxTasks

`func (o *V0041OpenapiJobInfoRespJobsInner) GetArrayMaxTasks() V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId`

GetArrayMaxTasks returns the ArrayMaxTasks field if non-nil, zero value otherwise.

### GetArrayMaxTasksOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetArrayMaxTasksOk() (*V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId, bool)`

GetArrayMaxTasksOk returns a tuple with the ArrayMaxTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrayMaxTasks

`func (o *V0041OpenapiJobInfoRespJobsInner) SetArrayMaxTasks(v V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId)`

SetArrayMaxTasks sets ArrayMaxTasks field to given value.

### HasArrayMaxTasks

`func (o *V0041OpenapiJobInfoRespJobsInner) HasArrayMaxTasks() bool`

HasArrayMaxTasks returns a boolean if a field has been set.

### GetArrayTaskString

`func (o *V0041OpenapiJobInfoRespJobsInner) GetArrayTaskString() string`

GetArrayTaskString returns the ArrayTaskString field if non-nil, zero value otherwise.

### GetArrayTaskStringOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetArrayTaskStringOk() (*string, bool)`

GetArrayTaskStringOk returns a tuple with the ArrayTaskString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrayTaskString

`func (o *V0041OpenapiJobInfoRespJobsInner) SetArrayTaskString(v string)`

SetArrayTaskString sets ArrayTaskString field to given value.

### HasArrayTaskString

`func (o *V0041OpenapiJobInfoRespJobsInner) HasArrayTaskString() bool`

HasArrayTaskString returns a boolean if a field has been set.

### GetAssociationId

`func (o *V0041OpenapiJobInfoRespJobsInner) GetAssociationId() int32`

GetAssociationId returns the AssociationId field if non-nil, zero value otherwise.

### GetAssociationIdOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetAssociationIdOk() (*int32, bool)`

GetAssociationIdOk returns a tuple with the AssociationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociationId

`func (o *V0041OpenapiJobInfoRespJobsInner) SetAssociationId(v int32)`

SetAssociationId sets AssociationId field to given value.

### HasAssociationId

`func (o *V0041OpenapiJobInfoRespJobsInner) HasAssociationId() bool`

HasAssociationId returns a boolean if a field has been set.

### GetBatchFeatures

`func (o *V0041OpenapiJobInfoRespJobsInner) GetBatchFeatures() string`

GetBatchFeatures returns the BatchFeatures field if non-nil, zero value otherwise.

### GetBatchFeaturesOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetBatchFeaturesOk() (*string, bool)`

GetBatchFeaturesOk returns a tuple with the BatchFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchFeatures

`func (o *V0041OpenapiJobInfoRespJobsInner) SetBatchFeatures(v string)`

SetBatchFeatures sets BatchFeatures field to given value.

### HasBatchFeatures

`func (o *V0041OpenapiJobInfoRespJobsInner) HasBatchFeatures() bool`

HasBatchFeatures returns a boolean if a field has been set.

### GetBatchFlag

`func (o *V0041OpenapiJobInfoRespJobsInner) GetBatchFlag() bool`

GetBatchFlag returns the BatchFlag field if non-nil, zero value otherwise.

### GetBatchFlagOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetBatchFlagOk() (*bool, bool)`

GetBatchFlagOk returns a tuple with the BatchFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchFlag

`func (o *V0041OpenapiJobInfoRespJobsInner) SetBatchFlag(v bool)`

SetBatchFlag sets BatchFlag field to given value.

### HasBatchFlag

`func (o *V0041OpenapiJobInfoRespJobsInner) HasBatchFlag() bool`

HasBatchFlag returns a boolean if a field has been set.

### GetBatchHost

`func (o *V0041OpenapiJobInfoRespJobsInner) GetBatchHost() string`

GetBatchHost returns the BatchHost field if non-nil, zero value otherwise.

### GetBatchHostOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetBatchHostOk() (*string, bool)`

GetBatchHostOk returns a tuple with the BatchHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchHost

`func (o *V0041OpenapiJobInfoRespJobsInner) SetBatchHost(v string)`

SetBatchHost sets BatchHost field to given value.

### HasBatchHost

`func (o *V0041OpenapiJobInfoRespJobsInner) HasBatchHost() bool`

HasBatchHost returns a boolean if a field has been set.

### GetFlags

`func (o *V0041OpenapiJobInfoRespJobsInner) GetFlags() []string`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetFlagsOk() (*[]string, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *V0041OpenapiJobInfoRespJobsInner) SetFlags(v []string)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *V0041OpenapiJobInfoRespJobsInner) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetBurstBuffer

`func (o *V0041OpenapiJobInfoRespJobsInner) GetBurstBuffer() string`

GetBurstBuffer returns the BurstBuffer field if non-nil, zero value otherwise.

### GetBurstBufferOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetBurstBufferOk() (*string, bool)`

GetBurstBufferOk returns a tuple with the BurstBuffer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurstBuffer

`func (o *V0041OpenapiJobInfoRespJobsInner) SetBurstBuffer(v string)`

SetBurstBuffer sets BurstBuffer field to given value.

### HasBurstBuffer

`func (o *V0041OpenapiJobInfoRespJobsInner) HasBurstBuffer() bool`

HasBurstBuffer returns a boolean if a field has been set.

### GetBurstBufferState

`func (o *V0041OpenapiJobInfoRespJobsInner) GetBurstBufferState() string`

GetBurstBufferState returns the BurstBufferState field if non-nil, zero value otherwise.

### GetBurstBufferStateOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetBurstBufferStateOk() (*string, bool)`

GetBurstBufferStateOk returns a tuple with the BurstBufferState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurstBufferState

`func (o *V0041OpenapiJobInfoRespJobsInner) SetBurstBufferState(v string)`

SetBurstBufferState sets BurstBufferState field to given value.

### HasBurstBufferState

`func (o *V0041OpenapiJobInfoRespJobsInner) HasBurstBufferState() bool`

HasBurstBufferState returns a boolean if a field has been set.

### GetCluster

`func (o *V0041OpenapiJobInfoRespJobsInner) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *V0041OpenapiJobInfoRespJobsInner) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *V0041OpenapiJobInfoRespJobsInner) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetClusterFeatures

`func (o *V0041OpenapiJobInfoRespJobsInner) GetClusterFeatures() string`

GetClusterFeatures returns the ClusterFeatures field if non-nil, zero value otherwise.

### GetClusterFeaturesOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetClusterFeaturesOk() (*string, bool)`

GetClusterFeaturesOk returns a tuple with the ClusterFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterFeatures

`func (o *V0041OpenapiJobInfoRespJobsInner) SetClusterFeatures(v string)`

SetClusterFeatures sets ClusterFeatures field to given value.

### HasClusterFeatures

`func (o *V0041OpenapiJobInfoRespJobsInner) HasClusterFeatures() bool`

HasClusterFeatures returns a boolean if a field has been set.

### GetCommand

`func (o *V0041OpenapiJobInfoRespJobsInner) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *V0041OpenapiJobInfoRespJobsInner) SetCommand(v string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *V0041OpenapiJobInfoRespJobsInner) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### GetComment

`func (o *V0041OpenapiJobInfoRespJobsInner) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *V0041OpenapiJobInfoRespJobsInner) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *V0041OpenapiJobInfoRespJobsInner) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetContainer

`func (o *V0041OpenapiJobInfoRespJobsInner) GetContainer() string`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetContainerOk() (*string, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *V0041OpenapiJobInfoRespJobsInner) SetContainer(v string)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *V0041OpenapiJobInfoRespJobsInner) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### GetContainerId

`func (o *V0041OpenapiJobInfoRespJobsInner) GetContainerId() string`

GetContainerId returns the ContainerId field if non-nil, zero value otherwise.

### GetContainerIdOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetContainerIdOk() (*string, bool)`

GetContainerIdOk returns a tuple with the ContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerId

`func (o *V0041OpenapiJobInfoRespJobsInner) SetContainerId(v string)`

SetContainerId sets ContainerId field to given value.

### HasContainerId

`func (o *V0041OpenapiJobInfoRespJobsInner) HasContainerId() bool`

HasContainerId returns a boolean if a field has been set.

### GetContiguous

`func (o *V0041OpenapiJobInfoRespJobsInner) GetContiguous() bool`

GetContiguous returns the Contiguous field if non-nil, zero value otherwise.

### GetContiguousOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetContiguousOk() (*bool, bool)`

GetContiguousOk returns a tuple with the Contiguous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContiguous

`func (o *V0041OpenapiJobInfoRespJobsInner) SetContiguous(v bool)`

SetContiguous sets Contiguous field to given value.

### HasContiguous

`func (o *V0041OpenapiJobInfoRespJobsInner) HasContiguous() bool`

HasContiguous returns a boolean if a field has been set.

### GetCoreSpec

`func (o *V0041OpenapiJobInfoRespJobsInner) GetCoreSpec() int32`

GetCoreSpec returns the CoreSpec field if non-nil, zero value otherwise.

### GetCoreSpecOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetCoreSpecOk() (*int32, bool)`

GetCoreSpecOk returns a tuple with the CoreSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoreSpec

`func (o *V0041OpenapiJobInfoRespJobsInner) SetCoreSpec(v int32)`

SetCoreSpec sets CoreSpec field to given value.

### HasCoreSpec

`func (o *V0041OpenapiJobInfoRespJobsInner) HasCoreSpec() bool`

HasCoreSpec returns a boolean if a field has been set.

### GetThreadSpec

`func (o *V0041OpenapiJobInfoRespJobsInner) GetThreadSpec() int32`

GetThreadSpec returns the ThreadSpec field if non-nil, zero value otherwise.

### GetThreadSpecOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetThreadSpecOk() (*int32, bool)`

GetThreadSpecOk returns a tuple with the ThreadSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadSpec

`func (o *V0041OpenapiJobInfoRespJobsInner) SetThreadSpec(v int32)`

SetThreadSpec sets ThreadSpec field to given value.

### HasThreadSpec

`func (o *V0041OpenapiJobInfoRespJobsInner) HasThreadSpec() bool`

HasThreadSpec returns a boolean if a field has been set.

### GetCoresPerSocket

`func (o *V0041OpenapiJobInfoRespJobsInner) GetCoresPerSocket() V0041JobDescMsgDistributionPlaneSize`

GetCoresPerSocket returns the CoresPerSocket field if non-nil, zero value otherwise.

### GetCoresPerSocketOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetCoresPerSocketOk() (*V0041JobDescMsgDistributionPlaneSize, bool)`

GetCoresPerSocketOk returns a tuple with the CoresPerSocket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoresPerSocket

`func (o *V0041OpenapiJobInfoRespJobsInner) SetCoresPerSocket(v V0041JobDescMsgDistributionPlaneSize)`

SetCoresPerSocket sets CoresPerSocket field to given value.

### HasCoresPerSocket

`func (o *V0041OpenapiJobInfoRespJobsInner) HasCoresPerSocket() bool`

HasCoresPerSocket returns a boolean if a field has been set.

### GetBillableTres

`func (o *V0041OpenapiJobInfoRespJobsInner) GetBillableTres() V0041OpenapiSlurmdbdConfigRespQosInnerLimitsFactor`

GetBillableTres returns the BillableTres field if non-nil, zero value otherwise.

### GetBillableTresOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetBillableTresOk() (*V0041OpenapiSlurmdbdConfigRespQosInnerLimitsFactor, bool)`

GetBillableTresOk returns a tuple with the BillableTres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillableTres

`func (o *V0041OpenapiJobInfoRespJobsInner) SetBillableTres(v V0041OpenapiSlurmdbdConfigRespQosInnerLimitsFactor)`

SetBillableTres sets BillableTres field to given value.

### HasBillableTres

`func (o *V0041OpenapiJobInfoRespJobsInner) HasBillableTres() bool`

HasBillableTres returns a boolean if a field has been set.

### GetCpusPerTask

`func (o *V0041OpenapiJobInfoRespJobsInner) GetCpusPerTask() V0041JobDescMsgDistributionPlaneSize`

GetCpusPerTask returns the CpusPerTask field if non-nil, zero value otherwise.

### GetCpusPerTaskOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetCpusPerTaskOk() (*V0041JobDescMsgDistributionPlaneSize, bool)`

GetCpusPerTaskOk returns a tuple with the CpusPerTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpusPerTask

`func (o *V0041OpenapiJobInfoRespJobsInner) SetCpusPerTask(v V0041JobDescMsgDistributionPlaneSize)`

SetCpusPerTask sets CpusPerTask field to given value.

### HasCpusPerTask

`func (o *V0041OpenapiJobInfoRespJobsInner) HasCpusPerTask() bool`

HasCpusPerTask returns a boolean if a field has been set.

### GetCpuFrequencyMinimum

`func (o *V0041OpenapiJobInfoRespJobsInner) GetCpuFrequencyMinimum() V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId`

GetCpuFrequencyMinimum returns the CpuFrequencyMinimum field if non-nil, zero value otherwise.

### GetCpuFrequencyMinimumOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetCpuFrequencyMinimumOk() (*V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId, bool)`

GetCpuFrequencyMinimumOk returns a tuple with the CpuFrequencyMinimum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuFrequencyMinimum

`func (o *V0041OpenapiJobInfoRespJobsInner) SetCpuFrequencyMinimum(v V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId)`

SetCpuFrequencyMinimum sets CpuFrequencyMinimum field to given value.

### HasCpuFrequencyMinimum

`func (o *V0041OpenapiJobInfoRespJobsInner) HasCpuFrequencyMinimum() bool`

HasCpuFrequencyMinimum returns a boolean if a field has been set.

### GetCpuFrequencyMaximum

`func (o *V0041OpenapiJobInfoRespJobsInner) GetCpuFrequencyMaximum() V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId`

GetCpuFrequencyMaximum returns the CpuFrequencyMaximum field if non-nil, zero value otherwise.

### GetCpuFrequencyMaximumOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetCpuFrequencyMaximumOk() (*V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId, bool)`

GetCpuFrequencyMaximumOk returns a tuple with the CpuFrequencyMaximum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuFrequencyMaximum

`func (o *V0041OpenapiJobInfoRespJobsInner) SetCpuFrequencyMaximum(v V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId)`

SetCpuFrequencyMaximum sets CpuFrequencyMaximum field to given value.

### HasCpuFrequencyMaximum

`func (o *V0041OpenapiJobInfoRespJobsInner) HasCpuFrequencyMaximum() bool`

HasCpuFrequencyMaximum returns a boolean if a field has been set.

### GetCpuFrequencyGovernor

`func (o *V0041OpenapiJobInfoRespJobsInner) GetCpuFrequencyGovernor() V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId`

GetCpuFrequencyGovernor returns the CpuFrequencyGovernor field if non-nil, zero value otherwise.

### GetCpuFrequencyGovernorOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetCpuFrequencyGovernorOk() (*V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId, bool)`

GetCpuFrequencyGovernorOk returns a tuple with the CpuFrequencyGovernor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuFrequencyGovernor

`func (o *V0041OpenapiJobInfoRespJobsInner) SetCpuFrequencyGovernor(v V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId)`

SetCpuFrequencyGovernor sets CpuFrequencyGovernor field to given value.

### HasCpuFrequencyGovernor

`func (o *V0041OpenapiJobInfoRespJobsInner) HasCpuFrequencyGovernor() bool`

HasCpuFrequencyGovernor returns a boolean if a field has been set.

### GetCpusPerTres

`func (o *V0041OpenapiJobInfoRespJobsInner) GetCpusPerTres() string`

GetCpusPerTres returns the CpusPerTres field if non-nil, zero value otherwise.

### GetCpusPerTresOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetCpusPerTresOk() (*string, bool)`

GetCpusPerTresOk returns a tuple with the CpusPerTres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpusPerTres

`func (o *V0041OpenapiJobInfoRespJobsInner) SetCpusPerTres(v string)`

SetCpusPerTres sets CpusPerTres field to given value.

### HasCpusPerTres

`func (o *V0041OpenapiJobInfoRespJobsInner) HasCpusPerTres() bool`

HasCpusPerTres returns a boolean if a field has been set.

### GetCron

`func (o *V0041OpenapiJobInfoRespJobsInner) GetCron() string`

GetCron returns the Cron field if non-nil, zero value otherwise.

### GetCronOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetCronOk() (*string, bool)`

GetCronOk returns a tuple with the Cron field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCron

`func (o *V0041OpenapiJobInfoRespJobsInner) SetCron(v string)`

SetCron sets Cron field to given value.

### HasCron

`func (o *V0041OpenapiJobInfoRespJobsInner) HasCron() bool`

HasCron returns a boolean if a field has been set.

### GetDeadline

`func (o *V0041OpenapiJobInfoRespJobsInner) GetDeadline() V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu`

GetDeadline returns the Deadline field if non-nil, zero value otherwise.

### GetDeadlineOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetDeadlineOk() (*V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu, bool)`

GetDeadlineOk returns a tuple with the Deadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadline

`func (o *V0041OpenapiJobInfoRespJobsInner) SetDeadline(v V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu)`

SetDeadline sets Deadline field to given value.

### HasDeadline

`func (o *V0041OpenapiJobInfoRespJobsInner) HasDeadline() bool`

HasDeadline returns a boolean if a field has been set.

### GetDelayBoot

`func (o *V0041OpenapiJobInfoRespJobsInner) GetDelayBoot() V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId`

GetDelayBoot returns the DelayBoot field if non-nil, zero value otherwise.

### GetDelayBootOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetDelayBootOk() (*V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId, bool)`

GetDelayBootOk returns a tuple with the DelayBoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayBoot

`func (o *V0041OpenapiJobInfoRespJobsInner) SetDelayBoot(v V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId)`

SetDelayBoot sets DelayBoot field to given value.

### HasDelayBoot

`func (o *V0041OpenapiJobInfoRespJobsInner) HasDelayBoot() bool`

HasDelayBoot returns a boolean if a field has been set.

### GetDependency

`func (o *V0041OpenapiJobInfoRespJobsInner) GetDependency() string`

GetDependency returns the Dependency field if non-nil, zero value otherwise.

### GetDependencyOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetDependencyOk() (*string, bool)`

GetDependencyOk returns a tuple with the Dependency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependency

`func (o *V0041OpenapiJobInfoRespJobsInner) SetDependency(v string)`

SetDependency sets Dependency field to given value.

### HasDependency

`func (o *V0041OpenapiJobInfoRespJobsInner) HasDependency() bool`

HasDependency returns a boolean if a field has been set.

### GetDerivedExitCode

`func (o *V0041OpenapiJobInfoRespJobsInner) GetDerivedExitCode() V0041OpenapiSlurmdbdJobsRespJobsInnerDerivedExitCode`

GetDerivedExitCode returns the DerivedExitCode field if non-nil, zero value otherwise.

### GetDerivedExitCodeOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetDerivedExitCodeOk() (*V0041OpenapiSlurmdbdJobsRespJobsInnerDerivedExitCode, bool)`

GetDerivedExitCodeOk returns a tuple with the DerivedExitCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerivedExitCode

`func (o *V0041OpenapiJobInfoRespJobsInner) SetDerivedExitCode(v V0041OpenapiSlurmdbdJobsRespJobsInnerDerivedExitCode)`

SetDerivedExitCode sets DerivedExitCode field to given value.

### HasDerivedExitCode

`func (o *V0041OpenapiJobInfoRespJobsInner) HasDerivedExitCode() bool`

HasDerivedExitCode returns a boolean if a field has been set.

### GetEligibleTime

`func (o *V0041OpenapiJobInfoRespJobsInner) GetEligibleTime() V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu`

GetEligibleTime returns the EligibleTime field if non-nil, zero value otherwise.

### GetEligibleTimeOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetEligibleTimeOk() (*V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu, bool)`

GetEligibleTimeOk returns a tuple with the EligibleTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEligibleTime

`func (o *V0041OpenapiJobInfoRespJobsInner) SetEligibleTime(v V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu)`

SetEligibleTime sets EligibleTime field to given value.

### HasEligibleTime

`func (o *V0041OpenapiJobInfoRespJobsInner) HasEligibleTime() bool`

HasEligibleTime returns a boolean if a field has been set.

### GetEndTime

`func (o *V0041OpenapiJobInfoRespJobsInner) GetEndTime() V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetEndTimeOk() (*V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *V0041OpenapiJobInfoRespJobsInner) SetEndTime(v V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *V0041OpenapiJobInfoRespJobsInner) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetExcludedNodes

`func (o *V0041OpenapiJobInfoRespJobsInner) GetExcludedNodes() string`

GetExcludedNodes returns the ExcludedNodes field if non-nil, zero value otherwise.

### GetExcludedNodesOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetExcludedNodesOk() (*string, bool)`

GetExcludedNodesOk returns a tuple with the ExcludedNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedNodes

`func (o *V0041OpenapiJobInfoRespJobsInner) SetExcludedNodes(v string)`

SetExcludedNodes sets ExcludedNodes field to given value.

### HasExcludedNodes

`func (o *V0041OpenapiJobInfoRespJobsInner) HasExcludedNodes() bool`

HasExcludedNodes returns a boolean if a field has been set.

### GetExitCode

`func (o *V0041OpenapiJobInfoRespJobsInner) GetExitCode() V0041OpenapiSlurmdbdJobsRespJobsInnerDerivedExitCode`

GetExitCode returns the ExitCode field if non-nil, zero value otherwise.

### GetExitCodeOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetExitCodeOk() (*V0041OpenapiSlurmdbdJobsRespJobsInnerDerivedExitCode, bool)`

GetExitCodeOk returns a tuple with the ExitCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitCode

`func (o *V0041OpenapiJobInfoRespJobsInner) SetExitCode(v V0041OpenapiSlurmdbdJobsRespJobsInnerDerivedExitCode)`

SetExitCode sets ExitCode field to given value.

### HasExitCode

`func (o *V0041OpenapiJobInfoRespJobsInner) HasExitCode() bool`

HasExitCode returns a boolean if a field has been set.

### GetExtra

`func (o *V0041OpenapiJobInfoRespJobsInner) GetExtra() string`

GetExtra returns the Extra field if non-nil, zero value otherwise.

### GetExtraOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetExtraOk() (*string, bool)`

GetExtraOk returns a tuple with the Extra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtra

`func (o *V0041OpenapiJobInfoRespJobsInner) SetExtra(v string)`

SetExtra sets Extra field to given value.

### HasExtra

`func (o *V0041OpenapiJobInfoRespJobsInner) HasExtra() bool`

HasExtra returns a boolean if a field has been set.

### GetFailedNode

`func (o *V0041OpenapiJobInfoRespJobsInner) GetFailedNode() string`

GetFailedNode returns the FailedNode field if non-nil, zero value otherwise.

### GetFailedNodeOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetFailedNodeOk() (*string, bool)`

GetFailedNodeOk returns a tuple with the FailedNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedNode

`func (o *V0041OpenapiJobInfoRespJobsInner) SetFailedNode(v string)`

SetFailedNode sets FailedNode field to given value.

### HasFailedNode

`func (o *V0041OpenapiJobInfoRespJobsInner) HasFailedNode() bool`

HasFailedNode returns a boolean if a field has been set.

### GetFeatures

`func (o *V0041OpenapiJobInfoRespJobsInner) GetFeatures() string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetFeaturesOk() (*string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *V0041OpenapiJobInfoRespJobsInner) SetFeatures(v string)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *V0041OpenapiJobInfoRespJobsInner) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetFederationOrigin

`func (o *V0041OpenapiJobInfoRespJobsInner) GetFederationOrigin() string`

GetFederationOrigin returns the FederationOrigin field if non-nil, zero value otherwise.

### GetFederationOriginOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetFederationOriginOk() (*string, bool)`

GetFederationOriginOk returns a tuple with the FederationOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederationOrigin

`func (o *V0041OpenapiJobInfoRespJobsInner) SetFederationOrigin(v string)`

SetFederationOrigin sets FederationOrigin field to given value.

### HasFederationOrigin

`func (o *V0041OpenapiJobInfoRespJobsInner) HasFederationOrigin() bool`

HasFederationOrigin returns a boolean if a field has been set.

### GetFederationSiblingsActive

`func (o *V0041OpenapiJobInfoRespJobsInner) GetFederationSiblingsActive() string`

GetFederationSiblingsActive returns the FederationSiblingsActive field if non-nil, zero value otherwise.

### GetFederationSiblingsActiveOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetFederationSiblingsActiveOk() (*string, bool)`

GetFederationSiblingsActiveOk returns a tuple with the FederationSiblingsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederationSiblingsActive

`func (o *V0041OpenapiJobInfoRespJobsInner) SetFederationSiblingsActive(v string)`

SetFederationSiblingsActive sets FederationSiblingsActive field to given value.

### HasFederationSiblingsActive

`func (o *V0041OpenapiJobInfoRespJobsInner) HasFederationSiblingsActive() bool`

HasFederationSiblingsActive returns a boolean if a field has been set.

### GetFederationSiblingsViable

`func (o *V0041OpenapiJobInfoRespJobsInner) GetFederationSiblingsViable() string`

GetFederationSiblingsViable returns the FederationSiblingsViable field if non-nil, zero value otherwise.

### GetFederationSiblingsViableOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetFederationSiblingsViableOk() (*string, bool)`

GetFederationSiblingsViableOk returns a tuple with the FederationSiblingsViable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederationSiblingsViable

`func (o *V0041OpenapiJobInfoRespJobsInner) SetFederationSiblingsViable(v string)`

SetFederationSiblingsViable sets FederationSiblingsViable field to given value.

### HasFederationSiblingsViable

`func (o *V0041OpenapiJobInfoRespJobsInner) HasFederationSiblingsViable() bool`

HasFederationSiblingsViable returns a boolean if a field has been set.

### GetGresDetail

`func (o *V0041OpenapiJobInfoRespJobsInner) GetGresDetail() []string`

GetGresDetail returns the GresDetail field if non-nil, zero value otherwise.

### GetGresDetailOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetGresDetailOk() (*[]string, bool)`

GetGresDetailOk returns a tuple with the GresDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGresDetail

`func (o *V0041OpenapiJobInfoRespJobsInner) SetGresDetail(v []string)`

SetGresDetail sets GresDetail field to given value.

### HasGresDetail

`func (o *V0041OpenapiJobInfoRespJobsInner) HasGresDetail() bool`

HasGresDetail returns a boolean if a field has been set.

### GetGroupId

`func (o *V0041OpenapiJobInfoRespJobsInner) GetGroupId() int32`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetGroupIdOk() (*int32, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *V0041OpenapiJobInfoRespJobsInner) SetGroupId(v int32)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *V0041OpenapiJobInfoRespJobsInner) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetGroupName

`func (o *V0041OpenapiJobInfoRespJobsInner) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *V0041OpenapiJobInfoRespJobsInner) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *V0041OpenapiJobInfoRespJobsInner) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetHetJobId

`func (o *V0041OpenapiJobInfoRespJobsInner) GetHetJobId() V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId`

GetHetJobId returns the HetJobId field if non-nil, zero value otherwise.

### GetHetJobIdOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetHetJobIdOk() (*V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId, bool)`

GetHetJobIdOk returns a tuple with the HetJobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHetJobId

`func (o *V0041OpenapiJobInfoRespJobsInner) SetHetJobId(v V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId)`

SetHetJobId sets HetJobId field to given value.

### HasHetJobId

`func (o *V0041OpenapiJobInfoRespJobsInner) HasHetJobId() bool`

HasHetJobId returns a boolean if a field has been set.

### GetHetJobIdSet

`func (o *V0041OpenapiJobInfoRespJobsInner) GetHetJobIdSet() string`

GetHetJobIdSet returns the HetJobIdSet field if non-nil, zero value otherwise.

### GetHetJobIdSetOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetHetJobIdSetOk() (*string, bool)`

GetHetJobIdSetOk returns a tuple with the HetJobIdSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHetJobIdSet

`func (o *V0041OpenapiJobInfoRespJobsInner) SetHetJobIdSet(v string)`

SetHetJobIdSet sets HetJobIdSet field to given value.

### HasHetJobIdSet

`func (o *V0041OpenapiJobInfoRespJobsInner) HasHetJobIdSet() bool`

HasHetJobIdSet returns a boolean if a field has been set.

### GetHetJobOffset

`func (o *V0041OpenapiJobInfoRespJobsInner) GetHetJobOffset() V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId`

GetHetJobOffset returns the HetJobOffset field if non-nil, zero value otherwise.

### GetHetJobOffsetOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetHetJobOffsetOk() (*V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId, bool)`

GetHetJobOffsetOk returns a tuple with the HetJobOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHetJobOffset

`func (o *V0041OpenapiJobInfoRespJobsInner) SetHetJobOffset(v V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId)`

SetHetJobOffset sets HetJobOffset field to given value.

### HasHetJobOffset

`func (o *V0041OpenapiJobInfoRespJobsInner) HasHetJobOffset() bool`

HasHetJobOffset returns a boolean if a field has been set.

### GetJobId

`func (o *V0041OpenapiJobInfoRespJobsInner) GetJobId() int32`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetJobIdOk() (*int32, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *V0041OpenapiJobInfoRespJobsInner) SetJobId(v int32)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *V0041OpenapiJobInfoRespJobsInner) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetJobResources

`func (o *V0041OpenapiJobInfoRespJobsInner) GetJobResources() V0041OpenapiJobInfoRespJobsInnerJobResources`

GetJobResources returns the JobResources field if non-nil, zero value otherwise.

### GetJobResourcesOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetJobResourcesOk() (*V0041OpenapiJobInfoRespJobsInnerJobResources, bool)`

GetJobResourcesOk returns a tuple with the JobResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobResources

`func (o *V0041OpenapiJobInfoRespJobsInner) SetJobResources(v V0041OpenapiJobInfoRespJobsInnerJobResources)`

SetJobResources sets JobResources field to given value.

### HasJobResources

`func (o *V0041OpenapiJobInfoRespJobsInner) HasJobResources() bool`

HasJobResources returns a boolean if a field has been set.

### GetJobSizeStr

`func (o *V0041OpenapiJobInfoRespJobsInner) GetJobSizeStr() []string`

GetJobSizeStr returns the JobSizeStr field if non-nil, zero value otherwise.

### GetJobSizeStrOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetJobSizeStrOk() (*[]string, bool)`

GetJobSizeStrOk returns a tuple with the JobSizeStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobSizeStr

`func (o *V0041OpenapiJobInfoRespJobsInner) SetJobSizeStr(v []string)`

SetJobSizeStr sets JobSizeStr field to given value.

### HasJobSizeStr

`func (o *V0041OpenapiJobInfoRespJobsInner) HasJobSizeStr() bool`

HasJobSizeStr returns a boolean if a field has been set.

### GetJobState

`func (o *V0041OpenapiJobInfoRespJobsInner) GetJobState() []string`

GetJobState returns the JobState field if non-nil, zero value otherwise.

### GetJobStateOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetJobStateOk() (*[]string, bool)`

GetJobStateOk returns a tuple with the JobState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobState

`func (o *V0041OpenapiJobInfoRespJobsInner) SetJobState(v []string)`

SetJobState sets JobState field to given value.

### HasJobState

`func (o *V0041OpenapiJobInfoRespJobsInner) HasJobState() bool`

HasJobState returns a boolean if a field has been set.

### GetLastSchedEvaluation

`func (o *V0041OpenapiJobInfoRespJobsInner) GetLastSchedEvaluation() V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu`

GetLastSchedEvaluation returns the LastSchedEvaluation field if non-nil, zero value otherwise.

### GetLastSchedEvaluationOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetLastSchedEvaluationOk() (*V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu, bool)`

GetLastSchedEvaluationOk returns a tuple with the LastSchedEvaluation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSchedEvaluation

`func (o *V0041OpenapiJobInfoRespJobsInner) SetLastSchedEvaluation(v V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu)`

SetLastSchedEvaluation sets LastSchedEvaluation field to given value.

### HasLastSchedEvaluation

`func (o *V0041OpenapiJobInfoRespJobsInner) HasLastSchedEvaluation() bool`

HasLastSchedEvaluation returns a boolean if a field has been set.

### GetLicenses

`func (o *V0041OpenapiJobInfoRespJobsInner) GetLicenses() string`

GetLicenses returns the Licenses field if non-nil, zero value otherwise.

### GetLicensesOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetLicensesOk() (*string, bool)`

GetLicensesOk returns a tuple with the Licenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenses

`func (o *V0041OpenapiJobInfoRespJobsInner) SetLicenses(v string)`

SetLicenses sets Licenses field to given value.

### HasLicenses

`func (o *V0041OpenapiJobInfoRespJobsInner) HasLicenses() bool`

HasLicenses returns a boolean if a field has been set.

### GetMailType

`func (o *V0041OpenapiJobInfoRespJobsInner) GetMailType() []string`

GetMailType returns the MailType field if non-nil, zero value otherwise.

### GetMailTypeOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetMailTypeOk() (*[]string, bool)`

GetMailTypeOk returns a tuple with the MailType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailType

`func (o *V0041OpenapiJobInfoRespJobsInner) SetMailType(v []string)`

SetMailType sets MailType field to given value.

### HasMailType

`func (o *V0041OpenapiJobInfoRespJobsInner) HasMailType() bool`

HasMailType returns a boolean if a field has been set.

### GetMailUser

`func (o *V0041OpenapiJobInfoRespJobsInner) GetMailUser() string`

GetMailUser returns the MailUser field if non-nil, zero value otherwise.

### GetMailUserOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetMailUserOk() (*string, bool)`

GetMailUserOk returns a tuple with the MailUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailUser

`func (o *V0041OpenapiJobInfoRespJobsInner) SetMailUser(v string)`

SetMailUser sets MailUser field to given value.

### HasMailUser

`func (o *V0041OpenapiJobInfoRespJobsInner) HasMailUser() bool`

HasMailUser returns a boolean if a field has been set.

### GetMaxCpus

`func (o *V0041OpenapiJobInfoRespJobsInner) GetMaxCpus() V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId`

GetMaxCpus returns the MaxCpus field if non-nil, zero value otherwise.

### GetMaxCpusOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetMaxCpusOk() (*V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId, bool)`

GetMaxCpusOk returns a tuple with the MaxCpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCpus

`func (o *V0041OpenapiJobInfoRespJobsInner) SetMaxCpus(v V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId)`

SetMaxCpus sets MaxCpus field to given value.

### HasMaxCpus

`func (o *V0041OpenapiJobInfoRespJobsInner) HasMaxCpus() bool`

HasMaxCpus returns a boolean if a field has been set.

### GetMaxNodes

`func (o *V0041OpenapiJobInfoRespJobsInner) GetMaxNodes() V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId`

GetMaxNodes returns the MaxNodes field if non-nil, zero value otherwise.

### GetMaxNodesOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetMaxNodesOk() (*V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId, bool)`

GetMaxNodesOk returns a tuple with the MaxNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNodes

`func (o *V0041OpenapiJobInfoRespJobsInner) SetMaxNodes(v V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId)`

SetMaxNodes sets MaxNodes field to given value.

### HasMaxNodes

`func (o *V0041OpenapiJobInfoRespJobsInner) HasMaxNodes() bool`

HasMaxNodes returns a boolean if a field has been set.

### GetMcsLabel

`func (o *V0041OpenapiJobInfoRespJobsInner) GetMcsLabel() string`

GetMcsLabel returns the McsLabel field if non-nil, zero value otherwise.

### GetMcsLabelOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetMcsLabelOk() (*string, bool)`

GetMcsLabelOk returns a tuple with the McsLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcsLabel

`func (o *V0041OpenapiJobInfoRespJobsInner) SetMcsLabel(v string)`

SetMcsLabel sets McsLabel field to given value.

### HasMcsLabel

`func (o *V0041OpenapiJobInfoRespJobsInner) HasMcsLabel() bool`

HasMcsLabel returns a boolean if a field has been set.

### GetMemoryPerTres

`func (o *V0041OpenapiJobInfoRespJobsInner) GetMemoryPerTres() string`

GetMemoryPerTres returns the MemoryPerTres field if non-nil, zero value otherwise.

### GetMemoryPerTresOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetMemoryPerTresOk() (*string, bool)`

GetMemoryPerTresOk returns a tuple with the MemoryPerTres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryPerTres

`func (o *V0041OpenapiJobInfoRespJobsInner) SetMemoryPerTres(v string)`

SetMemoryPerTres sets MemoryPerTres field to given value.

### HasMemoryPerTres

`func (o *V0041OpenapiJobInfoRespJobsInner) HasMemoryPerTres() bool`

HasMemoryPerTres returns a boolean if a field has been set.

### GetName

`func (o *V0041OpenapiJobInfoRespJobsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V0041OpenapiJobInfoRespJobsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V0041OpenapiJobInfoRespJobsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetwork

`func (o *V0041OpenapiJobInfoRespJobsInner) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *V0041OpenapiJobInfoRespJobsInner) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *V0041OpenapiJobInfoRespJobsInner) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetNodes

`func (o *V0041OpenapiJobInfoRespJobsInner) GetNodes() string`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetNodesOk() (*string, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *V0041OpenapiJobInfoRespJobsInner) SetNodes(v string)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *V0041OpenapiJobInfoRespJobsInner) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetNice

`func (o *V0041OpenapiJobInfoRespJobsInner) GetNice() int32`

GetNice returns the Nice field if non-nil, zero value otherwise.

### GetNiceOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetNiceOk() (*int32, bool)`

GetNiceOk returns a tuple with the Nice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNice

`func (o *V0041OpenapiJobInfoRespJobsInner) SetNice(v int32)`

SetNice sets Nice field to given value.

### HasNice

`func (o *V0041OpenapiJobInfoRespJobsInner) HasNice() bool`

HasNice returns a boolean if a field has been set.

### GetTasksPerCore

`func (o *V0041OpenapiJobInfoRespJobsInner) GetTasksPerCore() V0041JobDescMsgDistributionPlaneSize`

GetTasksPerCore returns the TasksPerCore field if non-nil, zero value otherwise.

### GetTasksPerCoreOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetTasksPerCoreOk() (*V0041JobDescMsgDistributionPlaneSize, bool)`

GetTasksPerCoreOk returns a tuple with the TasksPerCore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasksPerCore

`func (o *V0041OpenapiJobInfoRespJobsInner) SetTasksPerCore(v V0041JobDescMsgDistributionPlaneSize)`

SetTasksPerCore sets TasksPerCore field to given value.

### HasTasksPerCore

`func (o *V0041OpenapiJobInfoRespJobsInner) HasTasksPerCore() bool`

HasTasksPerCore returns a boolean if a field has been set.

### GetTasksPerTres

`func (o *V0041OpenapiJobInfoRespJobsInner) GetTasksPerTres() V0041JobDescMsgDistributionPlaneSize`

GetTasksPerTres returns the TasksPerTres field if non-nil, zero value otherwise.

### GetTasksPerTresOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetTasksPerTresOk() (*V0041JobDescMsgDistributionPlaneSize, bool)`

GetTasksPerTresOk returns a tuple with the TasksPerTres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasksPerTres

`func (o *V0041OpenapiJobInfoRespJobsInner) SetTasksPerTres(v V0041JobDescMsgDistributionPlaneSize)`

SetTasksPerTres sets TasksPerTres field to given value.

### HasTasksPerTres

`func (o *V0041OpenapiJobInfoRespJobsInner) HasTasksPerTres() bool`

HasTasksPerTres returns a boolean if a field has been set.

### GetTasksPerNode

`func (o *V0041OpenapiJobInfoRespJobsInner) GetTasksPerNode() V0041JobDescMsgDistributionPlaneSize`

GetTasksPerNode returns the TasksPerNode field if non-nil, zero value otherwise.

### GetTasksPerNodeOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetTasksPerNodeOk() (*V0041JobDescMsgDistributionPlaneSize, bool)`

GetTasksPerNodeOk returns a tuple with the TasksPerNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasksPerNode

`func (o *V0041OpenapiJobInfoRespJobsInner) SetTasksPerNode(v V0041JobDescMsgDistributionPlaneSize)`

SetTasksPerNode sets TasksPerNode field to given value.

### HasTasksPerNode

`func (o *V0041OpenapiJobInfoRespJobsInner) HasTasksPerNode() bool`

HasTasksPerNode returns a boolean if a field has been set.

### GetTasksPerSocket

`func (o *V0041OpenapiJobInfoRespJobsInner) GetTasksPerSocket() V0041JobDescMsgDistributionPlaneSize`

GetTasksPerSocket returns the TasksPerSocket field if non-nil, zero value otherwise.

### GetTasksPerSocketOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetTasksPerSocketOk() (*V0041JobDescMsgDistributionPlaneSize, bool)`

GetTasksPerSocketOk returns a tuple with the TasksPerSocket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasksPerSocket

`func (o *V0041OpenapiJobInfoRespJobsInner) SetTasksPerSocket(v V0041JobDescMsgDistributionPlaneSize)`

SetTasksPerSocket sets TasksPerSocket field to given value.

### HasTasksPerSocket

`func (o *V0041OpenapiJobInfoRespJobsInner) HasTasksPerSocket() bool`

HasTasksPerSocket returns a boolean if a field has been set.

### GetTasksPerBoard

`func (o *V0041OpenapiJobInfoRespJobsInner) GetTasksPerBoard() V0041JobDescMsgDistributionPlaneSize`

GetTasksPerBoard returns the TasksPerBoard field if non-nil, zero value otherwise.

### GetTasksPerBoardOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetTasksPerBoardOk() (*V0041JobDescMsgDistributionPlaneSize, bool)`

GetTasksPerBoardOk returns a tuple with the TasksPerBoard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasksPerBoard

`func (o *V0041OpenapiJobInfoRespJobsInner) SetTasksPerBoard(v V0041JobDescMsgDistributionPlaneSize)`

SetTasksPerBoard sets TasksPerBoard field to given value.

### HasTasksPerBoard

`func (o *V0041OpenapiJobInfoRespJobsInner) HasTasksPerBoard() bool`

HasTasksPerBoard returns a boolean if a field has been set.

### GetCpus

`func (o *V0041OpenapiJobInfoRespJobsInner) GetCpus() V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId`

GetCpus returns the Cpus field if non-nil, zero value otherwise.

### GetCpusOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetCpusOk() (*V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId, bool)`

GetCpusOk returns a tuple with the Cpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpus

`func (o *V0041OpenapiJobInfoRespJobsInner) SetCpus(v V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId)`

SetCpus sets Cpus field to given value.

### HasCpus

`func (o *V0041OpenapiJobInfoRespJobsInner) HasCpus() bool`

HasCpus returns a boolean if a field has been set.

### GetNodeCount

`func (o *V0041OpenapiJobInfoRespJobsInner) GetNodeCount() V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetNodeCountOk() (*V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *V0041OpenapiJobInfoRespJobsInner) SetNodeCount(v V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId)`

SetNodeCount sets NodeCount field to given value.

### HasNodeCount

`func (o *V0041OpenapiJobInfoRespJobsInner) HasNodeCount() bool`

HasNodeCount returns a boolean if a field has been set.

### GetTasks

`func (o *V0041OpenapiJobInfoRespJobsInner) GetTasks() V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetTasksOk() (*V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *V0041OpenapiJobInfoRespJobsInner) SetTasks(v V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *V0041OpenapiJobInfoRespJobsInner) HasTasks() bool`

HasTasks returns a boolean if a field has been set.

### GetPartition

`func (o *V0041OpenapiJobInfoRespJobsInner) GetPartition() string`

GetPartition returns the Partition field if non-nil, zero value otherwise.

### GetPartitionOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetPartitionOk() (*string, bool)`

GetPartitionOk returns a tuple with the Partition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartition

`func (o *V0041OpenapiJobInfoRespJobsInner) SetPartition(v string)`

SetPartition sets Partition field to given value.

### HasPartition

`func (o *V0041OpenapiJobInfoRespJobsInner) HasPartition() bool`

HasPartition returns a boolean if a field has been set.

### GetPrefer

`func (o *V0041OpenapiJobInfoRespJobsInner) GetPrefer() string`

GetPrefer returns the Prefer field if non-nil, zero value otherwise.

### GetPreferOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetPreferOk() (*string, bool)`

GetPreferOk returns a tuple with the Prefer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefer

`func (o *V0041OpenapiJobInfoRespJobsInner) SetPrefer(v string)`

SetPrefer sets Prefer field to given value.

### HasPrefer

`func (o *V0041OpenapiJobInfoRespJobsInner) HasPrefer() bool`

HasPrefer returns a boolean if a field has been set.

### GetMemoryPerCpu

`func (o *V0041OpenapiJobInfoRespJobsInner) GetMemoryPerCpu() V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu`

GetMemoryPerCpu returns the MemoryPerCpu field if non-nil, zero value otherwise.

### GetMemoryPerCpuOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetMemoryPerCpuOk() (*V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu, bool)`

GetMemoryPerCpuOk returns a tuple with the MemoryPerCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryPerCpu

`func (o *V0041OpenapiJobInfoRespJobsInner) SetMemoryPerCpu(v V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu)`

SetMemoryPerCpu sets MemoryPerCpu field to given value.

### HasMemoryPerCpu

`func (o *V0041OpenapiJobInfoRespJobsInner) HasMemoryPerCpu() bool`

HasMemoryPerCpu returns a boolean if a field has been set.

### GetMemoryPerNode

`func (o *V0041OpenapiJobInfoRespJobsInner) GetMemoryPerNode() V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu`

GetMemoryPerNode returns the MemoryPerNode field if non-nil, zero value otherwise.

### GetMemoryPerNodeOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetMemoryPerNodeOk() (*V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu, bool)`

GetMemoryPerNodeOk returns a tuple with the MemoryPerNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryPerNode

`func (o *V0041OpenapiJobInfoRespJobsInner) SetMemoryPerNode(v V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu)`

SetMemoryPerNode sets MemoryPerNode field to given value.

### HasMemoryPerNode

`func (o *V0041OpenapiJobInfoRespJobsInner) HasMemoryPerNode() bool`

HasMemoryPerNode returns a boolean if a field has been set.

### GetMinimumCpusPerNode

`func (o *V0041OpenapiJobInfoRespJobsInner) GetMinimumCpusPerNode() V0041JobDescMsgDistributionPlaneSize`

GetMinimumCpusPerNode returns the MinimumCpusPerNode field if non-nil, zero value otherwise.

### GetMinimumCpusPerNodeOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetMinimumCpusPerNodeOk() (*V0041JobDescMsgDistributionPlaneSize, bool)`

GetMinimumCpusPerNodeOk returns a tuple with the MinimumCpusPerNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumCpusPerNode

`func (o *V0041OpenapiJobInfoRespJobsInner) SetMinimumCpusPerNode(v V0041JobDescMsgDistributionPlaneSize)`

SetMinimumCpusPerNode sets MinimumCpusPerNode field to given value.

### HasMinimumCpusPerNode

`func (o *V0041OpenapiJobInfoRespJobsInner) HasMinimumCpusPerNode() bool`

HasMinimumCpusPerNode returns a boolean if a field has been set.

### GetMinimumTmpDiskPerNode

`func (o *V0041OpenapiJobInfoRespJobsInner) GetMinimumTmpDiskPerNode() V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId`

GetMinimumTmpDiskPerNode returns the MinimumTmpDiskPerNode field if non-nil, zero value otherwise.

### GetMinimumTmpDiskPerNodeOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetMinimumTmpDiskPerNodeOk() (*V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId, bool)`

GetMinimumTmpDiskPerNodeOk returns a tuple with the MinimumTmpDiskPerNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumTmpDiskPerNode

`func (o *V0041OpenapiJobInfoRespJobsInner) SetMinimumTmpDiskPerNode(v V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId)`

SetMinimumTmpDiskPerNode sets MinimumTmpDiskPerNode field to given value.

### HasMinimumTmpDiskPerNode

`func (o *V0041OpenapiJobInfoRespJobsInner) HasMinimumTmpDiskPerNode() bool`

HasMinimumTmpDiskPerNode returns a boolean if a field has been set.

### GetPower

`func (o *V0041OpenapiJobInfoRespJobsInner) GetPower() V0041OpenapiJobInfoRespJobsInnerPower`

GetPower returns the Power field if non-nil, zero value otherwise.

### GetPowerOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetPowerOk() (*V0041OpenapiJobInfoRespJobsInnerPower, bool)`

GetPowerOk returns a tuple with the Power field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPower

`func (o *V0041OpenapiJobInfoRespJobsInner) SetPower(v V0041OpenapiJobInfoRespJobsInnerPower)`

SetPower sets Power field to given value.

### HasPower

`func (o *V0041OpenapiJobInfoRespJobsInner) HasPower() bool`

HasPower returns a boolean if a field has been set.

### GetPreemptTime

`func (o *V0041OpenapiJobInfoRespJobsInner) GetPreemptTime() V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu`

GetPreemptTime returns the PreemptTime field if non-nil, zero value otherwise.

### GetPreemptTimeOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetPreemptTimeOk() (*V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu, bool)`

GetPreemptTimeOk returns a tuple with the PreemptTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreemptTime

`func (o *V0041OpenapiJobInfoRespJobsInner) SetPreemptTime(v V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu)`

SetPreemptTime sets PreemptTime field to given value.

### HasPreemptTime

`func (o *V0041OpenapiJobInfoRespJobsInner) HasPreemptTime() bool`

HasPreemptTime returns a boolean if a field has been set.

### GetPreemptableTime

`func (o *V0041OpenapiJobInfoRespJobsInner) GetPreemptableTime() V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu`

GetPreemptableTime returns the PreemptableTime field if non-nil, zero value otherwise.

### GetPreemptableTimeOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetPreemptableTimeOk() (*V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu, bool)`

GetPreemptableTimeOk returns a tuple with the PreemptableTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreemptableTime

`func (o *V0041OpenapiJobInfoRespJobsInner) SetPreemptableTime(v V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu)`

SetPreemptableTime sets PreemptableTime field to given value.

### HasPreemptableTime

`func (o *V0041OpenapiJobInfoRespJobsInner) HasPreemptableTime() bool`

HasPreemptableTime returns a boolean if a field has been set.

### GetPreSusTime

`func (o *V0041OpenapiJobInfoRespJobsInner) GetPreSusTime() V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu`

GetPreSusTime returns the PreSusTime field if non-nil, zero value otherwise.

### GetPreSusTimeOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetPreSusTimeOk() (*V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu, bool)`

GetPreSusTimeOk returns a tuple with the PreSusTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreSusTime

`func (o *V0041OpenapiJobInfoRespJobsInner) SetPreSusTime(v V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu)`

SetPreSusTime sets PreSusTime field to given value.

### HasPreSusTime

`func (o *V0041OpenapiJobInfoRespJobsInner) HasPreSusTime() bool`

HasPreSusTime returns a boolean if a field has been set.

### GetHold

`func (o *V0041OpenapiJobInfoRespJobsInner) GetHold() bool`

GetHold returns the Hold field if non-nil, zero value otherwise.

### GetHoldOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetHoldOk() (*bool, bool)`

GetHoldOk returns a tuple with the Hold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHold

`func (o *V0041OpenapiJobInfoRespJobsInner) SetHold(v bool)`

SetHold sets Hold field to given value.

### HasHold

`func (o *V0041OpenapiJobInfoRespJobsInner) HasHold() bool`

HasHold returns a boolean if a field has been set.

### GetPriority

`func (o *V0041OpenapiJobInfoRespJobsInner) GetPriority() V0041OpenapiSlurmdbdJobsRespJobsInnerPriority`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetPriorityOk() (*V0041OpenapiSlurmdbdJobsRespJobsInnerPriority, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *V0041OpenapiJobInfoRespJobsInner) SetPriority(v V0041OpenapiSlurmdbdJobsRespJobsInnerPriority)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *V0041OpenapiJobInfoRespJobsInner) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetProfile

`func (o *V0041OpenapiJobInfoRespJobsInner) GetProfile() []string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetProfileOk() (*[]string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *V0041OpenapiJobInfoRespJobsInner) SetProfile(v []string)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *V0041OpenapiJobInfoRespJobsInner) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetQos

`func (o *V0041OpenapiJobInfoRespJobsInner) GetQos() string`

GetQos returns the Qos field if non-nil, zero value otherwise.

### GetQosOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetQosOk() (*string, bool)`

GetQosOk returns a tuple with the Qos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQos

`func (o *V0041OpenapiJobInfoRespJobsInner) SetQos(v string)`

SetQos sets Qos field to given value.

### HasQos

`func (o *V0041OpenapiJobInfoRespJobsInner) HasQos() bool`

HasQos returns a boolean if a field has been set.

### GetReboot

`func (o *V0041OpenapiJobInfoRespJobsInner) GetReboot() bool`

GetReboot returns the Reboot field if non-nil, zero value otherwise.

### GetRebootOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetRebootOk() (*bool, bool)`

GetRebootOk returns a tuple with the Reboot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReboot

`func (o *V0041OpenapiJobInfoRespJobsInner) SetReboot(v bool)`

SetReboot sets Reboot field to given value.

### HasReboot

`func (o *V0041OpenapiJobInfoRespJobsInner) HasReboot() bool`

HasReboot returns a boolean if a field has been set.

### GetRequiredNodes

`func (o *V0041OpenapiJobInfoRespJobsInner) GetRequiredNodes() string`

GetRequiredNodes returns the RequiredNodes field if non-nil, zero value otherwise.

### GetRequiredNodesOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetRequiredNodesOk() (*string, bool)`

GetRequiredNodesOk returns a tuple with the RequiredNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredNodes

`func (o *V0041OpenapiJobInfoRespJobsInner) SetRequiredNodes(v string)`

SetRequiredNodes sets RequiredNodes field to given value.

### HasRequiredNodes

`func (o *V0041OpenapiJobInfoRespJobsInner) HasRequiredNodes() bool`

HasRequiredNodes returns a boolean if a field has been set.

### GetMinimumSwitches

`func (o *V0041OpenapiJobInfoRespJobsInner) GetMinimumSwitches() int32`

GetMinimumSwitches returns the MinimumSwitches field if non-nil, zero value otherwise.

### GetMinimumSwitchesOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetMinimumSwitchesOk() (*int32, bool)`

GetMinimumSwitchesOk returns a tuple with the MinimumSwitches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumSwitches

`func (o *V0041OpenapiJobInfoRespJobsInner) SetMinimumSwitches(v int32)`

SetMinimumSwitches sets MinimumSwitches field to given value.

### HasMinimumSwitches

`func (o *V0041OpenapiJobInfoRespJobsInner) HasMinimumSwitches() bool`

HasMinimumSwitches returns a boolean if a field has been set.

### GetRequeue

`func (o *V0041OpenapiJobInfoRespJobsInner) GetRequeue() bool`

GetRequeue returns the Requeue field if non-nil, zero value otherwise.

### GetRequeueOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetRequeueOk() (*bool, bool)`

GetRequeueOk returns a tuple with the Requeue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequeue

`func (o *V0041OpenapiJobInfoRespJobsInner) SetRequeue(v bool)`

SetRequeue sets Requeue field to given value.

### HasRequeue

`func (o *V0041OpenapiJobInfoRespJobsInner) HasRequeue() bool`

HasRequeue returns a boolean if a field has been set.

### GetResizeTime

`func (o *V0041OpenapiJobInfoRespJobsInner) GetResizeTime() V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu`

GetResizeTime returns the ResizeTime field if non-nil, zero value otherwise.

### GetResizeTimeOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetResizeTimeOk() (*V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu, bool)`

GetResizeTimeOk returns a tuple with the ResizeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResizeTime

`func (o *V0041OpenapiJobInfoRespJobsInner) SetResizeTime(v V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu)`

SetResizeTime sets ResizeTime field to given value.

### HasResizeTime

`func (o *V0041OpenapiJobInfoRespJobsInner) HasResizeTime() bool`

HasResizeTime returns a boolean if a field has been set.

### GetRestartCnt

`func (o *V0041OpenapiJobInfoRespJobsInner) GetRestartCnt() int32`

GetRestartCnt returns the RestartCnt field if non-nil, zero value otherwise.

### GetRestartCntOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetRestartCntOk() (*int32, bool)`

GetRestartCntOk returns a tuple with the RestartCnt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartCnt

`func (o *V0041OpenapiJobInfoRespJobsInner) SetRestartCnt(v int32)`

SetRestartCnt sets RestartCnt field to given value.

### HasRestartCnt

`func (o *V0041OpenapiJobInfoRespJobsInner) HasRestartCnt() bool`

HasRestartCnt returns a boolean if a field has been set.

### GetResvName

`func (o *V0041OpenapiJobInfoRespJobsInner) GetResvName() string`

GetResvName returns the ResvName field if non-nil, zero value otherwise.

### GetResvNameOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetResvNameOk() (*string, bool)`

GetResvNameOk returns a tuple with the ResvName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResvName

`func (o *V0041OpenapiJobInfoRespJobsInner) SetResvName(v string)`

SetResvName sets ResvName field to given value.

### HasResvName

`func (o *V0041OpenapiJobInfoRespJobsInner) HasResvName() bool`

HasResvName returns a boolean if a field has been set.

### GetScheduledNodes

`func (o *V0041OpenapiJobInfoRespJobsInner) GetScheduledNodes() string`

GetScheduledNodes returns the ScheduledNodes field if non-nil, zero value otherwise.

### GetScheduledNodesOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetScheduledNodesOk() (*string, bool)`

GetScheduledNodesOk returns a tuple with the ScheduledNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledNodes

`func (o *V0041OpenapiJobInfoRespJobsInner) SetScheduledNodes(v string)`

SetScheduledNodes sets ScheduledNodes field to given value.

### HasScheduledNodes

`func (o *V0041OpenapiJobInfoRespJobsInner) HasScheduledNodes() bool`

HasScheduledNodes returns a boolean if a field has been set.

### GetSelinuxContext

`func (o *V0041OpenapiJobInfoRespJobsInner) GetSelinuxContext() string`

GetSelinuxContext returns the SelinuxContext field if non-nil, zero value otherwise.

### GetSelinuxContextOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetSelinuxContextOk() (*string, bool)`

GetSelinuxContextOk returns a tuple with the SelinuxContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelinuxContext

`func (o *V0041OpenapiJobInfoRespJobsInner) SetSelinuxContext(v string)`

SetSelinuxContext sets SelinuxContext field to given value.

### HasSelinuxContext

`func (o *V0041OpenapiJobInfoRespJobsInner) HasSelinuxContext() bool`

HasSelinuxContext returns a boolean if a field has been set.

### GetShared

`func (o *V0041OpenapiJobInfoRespJobsInner) GetShared() []string`

GetShared returns the Shared field if non-nil, zero value otherwise.

### GetSharedOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetSharedOk() (*[]string, bool)`

GetSharedOk returns a tuple with the Shared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShared

`func (o *V0041OpenapiJobInfoRespJobsInner) SetShared(v []string)`

SetShared sets Shared field to given value.

### HasShared

`func (o *V0041OpenapiJobInfoRespJobsInner) HasShared() bool`

HasShared returns a boolean if a field has been set.

### GetExclusive

`func (o *V0041OpenapiJobInfoRespJobsInner) GetExclusive() []string`

GetExclusive returns the Exclusive field if non-nil, zero value otherwise.

### GetExclusiveOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetExclusiveOk() (*[]string, bool)`

GetExclusiveOk returns a tuple with the Exclusive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclusive

`func (o *V0041OpenapiJobInfoRespJobsInner) SetExclusive(v []string)`

SetExclusive sets Exclusive field to given value.

### HasExclusive

`func (o *V0041OpenapiJobInfoRespJobsInner) HasExclusive() bool`

HasExclusive returns a boolean if a field has been set.

### GetOversubscribe

`func (o *V0041OpenapiJobInfoRespJobsInner) GetOversubscribe() bool`

GetOversubscribe returns the Oversubscribe field if non-nil, zero value otherwise.

### GetOversubscribeOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetOversubscribeOk() (*bool, bool)`

GetOversubscribeOk returns a tuple with the Oversubscribe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOversubscribe

`func (o *V0041OpenapiJobInfoRespJobsInner) SetOversubscribe(v bool)`

SetOversubscribe sets Oversubscribe field to given value.

### HasOversubscribe

`func (o *V0041OpenapiJobInfoRespJobsInner) HasOversubscribe() bool`

HasOversubscribe returns a boolean if a field has been set.

### GetShowFlags

`func (o *V0041OpenapiJobInfoRespJobsInner) GetShowFlags() []string`

GetShowFlags returns the ShowFlags field if non-nil, zero value otherwise.

### GetShowFlagsOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetShowFlagsOk() (*[]string, bool)`

GetShowFlagsOk returns a tuple with the ShowFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowFlags

`func (o *V0041OpenapiJobInfoRespJobsInner) SetShowFlags(v []string)`

SetShowFlags sets ShowFlags field to given value.

### HasShowFlags

`func (o *V0041OpenapiJobInfoRespJobsInner) HasShowFlags() bool`

HasShowFlags returns a boolean if a field has been set.

### GetSocketsPerBoard

`func (o *V0041OpenapiJobInfoRespJobsInner) GetSocketsPerBoard() int32`

GetSocketsPerBoard returns the SocketsPerBoard field if non-nil, zero value otherwise.

### GetSocketsPerBoardOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetSocketsPerBoardOk() (*int32, bool)`

GetSocketsPerBoardOk returns a tuple with the SocketsPerBoard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocketsPerBoard

`func (o *V0041OpenapiJobInfoRespJobsInner) SetSocketsPerBoard(v int32)`

SetSocketsPerBoard sets SocketsPerBoard field to given value.

### HasSocketsPerBoard

`func (o *V0041OpenapiJobInfoRespJobsInner) HasSocketsPerBoard() bool`

HasSocketsPerBoard returns a boolean if a field has been set.

### GetSocketsPerNode

`func (o *V0041OpenapiJobInfoRespJobsInner) GetSocketsPerNode() V0041JobDescMsgDistributionPlaneSize`

GetSocketsPerNode returns the SocketsPerNode field if non-nil, zero value otherwise.

### GetSocketsPerNodeOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetSocketsPerNodeOk() (*V0041JobDescMsgDistributionPlaneSize, bool)`

GetSocketsPerNodeOk returns a tuple with the SocketsPerNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocketsPerNode

`func (o *V0041OpenapiJobInfoRespJobsInner) SetSocketsPerNode(v V0041JobDescMsgDistributionPlaneSize)`

SetSocketsPerNode sets SocketsPerNode field to given value.

### HasSocketsPerNode

`func (o *V0041OpenapiJobInfoRespJobsInner) HasSocketsPerNode() bool`

HasSocketsPerNode returns a boolean if a field has been set.

### GetStartTime

`func (o *V0041OpenapiJobInfoRespJobsInner) GetStartTime() V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetStartTimeOk() (*V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *V0041OpenapiJobInfoRespJobsInner) SetStartTime(v V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *V0041OpenapiJobInfoRespJobsInner) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStateDescription

`func (o *V0041OpenapiJobInfoRespJobsInner) GetStateDescription() string`

GetStateDescription returns the StateDescription field if non-nil, zero value otherwise.

### GetStateDescriptionOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetStateDescriptionOk() (*string, bool)`

GetStateDescriptionOk returns a tuple with the StateDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateDescription

`func (o *V0041OpenapiJobInfoRespJobsInner) SetStateDescription(v string)`

SetStateDescription sets StateDescription field to given value.

### HasStateDescription

`func (o *V0041OpenapiJobInfoRespJobsInner) HasStateDescription() bool`

HasStateDescription returns a boolean if a field has been set.

### GetStateReason

`func (o *V0041OpenapiJobInfoRespJobsInner) GetStateReason() string`

GetStateReason returns the StateReason field if non-nil, zero value otherwise.

### GetStateReasonOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetStateReasonOk() (*string, bool)`

GetStateReasonOk returns a tuple with the StateReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateReason

`func (o *V0041OpenapiJobInfoRespJobsInner) SetStateReason(v string)`

SetStateReason sets StateReason field to given value.

### HasStateReason

`func (o *V0041OpenapiJobInfoRespJobsInner) HasStateReason() bool`

HasStateReason returns a boolean if a field has been set.

### GetStandardError

`func (o *V0041OpenapiJobInfoRespJobsInner) GetStandardError() string`

GetStandardError returns the StandardError field if non-nil, zero value otherwise.

### GetStandardErrorOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetStandardErrorOk() (*string, bool)`

GetStandardErrorOk returns a tuple with the StandardError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardError

`func (o *V0041OpenapiJobInfoRespJobsInner) SetStandardError(v string)`

SetStandardError sets StandardError field to given value.

### HasStandardError

`func (o *V0041OpenapiJobInfoRespJobsInner) HasStandardError() bool`

HasStandardError returns a boolean if a field has been set.

### GetStandardInput

`func (o *V0041OpenapiJobInfoRespJobsInner) GetStandardInput() string`

GetStandardInput returns the StandardInput field if non-nil, zero value otherwise.

### GetStandardInputOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetStandardInputOk() (*string, bool)`

GetStandardInputOk returns a tuple with the StandardInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardInput

`func (o *V0041OpenapiJobInfoRespJobsInner) SetStandardInput(v string)`

SetStandardInput sets StandardInput field to given value.

### HasStandardInput

`func (o *V0041OpenapiJobInfoRespJobsInner) HasStandardInput() bool`

HasStandardInput returns a boolean if a field has been set.

### GetStandardOutput

`func (o *V0041OpenapiJobInfoRespJobsInner) GetStandardOutput() string`

GetStandardOutput returns the StandardOutput field if non-nil, zero value otherwise.

### GetStandardOutputOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetStandardOutputOk() (*string, bool)`

GetStandardOutputOk returns a tuple with the StandardOutput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardOutput

`func (o *V0041OpenapiJobInfoRespJobsInner) SetStandardOutput(v string)`

SetStandardOutput sets StandardOutput field to given value.

### HasStandardOutput

`func (o *V0041OpenapiJobInfoRespJobsInner) HasStandardOutput() bool`

HasStandardOutput returns a boolean if a field has been set.

### GetSubmitTime

`func (o *V0041OpenapiJobInfoRespJobsInner) GetSubmitTime() V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu`

GetSubmitTime returns the SubmitTime field if non-nil, zero value otherwise.

### GetSubmitTimeOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetSubmitTimeOk() (*V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu, bool)`

GetSubmitTimeOk returns a tuple with the SubmitTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmitTime

`func (o *V0041OpenapiJobInfoRespJobsInner) SetSubmitTime(v V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu)`

SetSubmitTime sets SubmitTime field to given value.

### HasSubmitTime

`func (o *V0041OpenapiJobInfoRespJobsInner) HasSubmitTime() bool`

HasSubmitTime returns a boolean if a field has been set.

### GetSuspendTime

`func (o *V0041OpenapiJobInfoRespJobsInner) GetSuspendTime() V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu`

GetSuspendTime returns the SuspendTime field if non-nil, zero value otherwise.

### GetSuspendTimeOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetSuspendTimeOk() (*V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu, bool)`

GetSuspendTimeOk returns a tuple with the SuspendTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspendTime

`func (o *V0041OpenapiJobInfoRespJobsInner) SetSuspendTime(v V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu)`

SetSuspendTime sets SuspendTime field to given value.

### HasSuspendTime

`func (o *V0041OpenapiJobInfoRespJobsInner) HasSuspendTime() bool`

HasSuspendTime returns a boolean if a field has been set.

### GetSystemComment

`func (o *V0041OpenapiJobInfoRespJobsInner) GetSystemComment() string`

GetSystemComment returns the SystemComment field if non-nil, zero value otherwise.

### GetSystemCommentOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetSystemCommentOk() (*string, bool)`

GetSystemCommentOk returns a tuple with the SystemComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemComment

`func (o *V0041OpenapiJobInfoRespJobsInner) SetSystemComment(v string)`

SetSystemComment sets SystemComment field to given value.

### HasSystemComment

`func (o *V0041OpenapiJobInfoRespJobsInner) HasSystemComment() bool`

HasSystemComment returns a boolean if a field has been set.

### GetTimeLimit

`func (o *V0041OpenapiJobInfoRespJobsInner) GetTimeLimit() V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId`

GetTimeLimit returns the TimeLimit field if non-nil, zero value otherwise.

### GetTimeLimitOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetTimeLimitOk() (*V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId, bool)`

GetTimeLimitOk returns a tuple with the TimeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeLimit

`func (o *V0041OpenapiJobInfoRespJobsInner) SetTimeLimit(v V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId)`

SetTimeLimit sets TimeLimit field to given value.

### HasTimeLimit

`func (o *V0041OpenapiJobInfoRespJobsInner) HasTimeLimit() bool`

HasTimeLimit returns a boolean if a field has been set.

### GetTimeMinimum

`func (o *V0041OpenapiJobInfoRespJobsInner) GetTimeMinimum() V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId`

GetTimeMinimum returns the TimeMinimum field if non-nil, zero value otherwise.

### GetTimeMinimumOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetTimeMinimumOk() (*V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId, bool)`

GetTimeMinimumOk returns a tuple with the TimeMinimum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeMinimum

`func (o *V0041OpenapiJobInfoRespJobsInner) SetTimeMinimum(v V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId)`

SetTimeMinimum sets TimeMinimum field to given value.

### HasTimeMinimum

`func (o *V0041OpenapiJobInfoRespJobsInner) HasTimeMinimum() bool`

HasTimeMinimum returns a boolean if a field has been set.

### GetThreadsPerCore

`func (o *V0041OpenapiJobInfoRespJobsInner) GetThreadsPerCore() V0041JobDescMsgDistributionPlaneSize`

GetThreadsPerCore returns the ThreadsPerCore field if non-nil, zero value otherwise.

### GetThreadsPerCoreOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetThreadsPerCoreOk() (*V0041JobDescMsgDistributionPlaneSize, bool)`

GetThreadsPerCoreOk returns a tuple with the ThreadsPerCore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadsPerCore

`func (o *V0041OpenapiJobInfoRespJobsInner) SetThreadsPerCore(v V0041JobDescMsgDistributionPlaneSize)`

SetThreadsPerCore sets ThreadsPerCore field to given value.

### HasThreadsPerCore

`func (o *V0041OpenapiJobInfoRespJobsInner) HasThreadsPerCore() bool`

HasThreadsPerCore returns a boolean if a field has been set.

### GetTresBind

`func (o *V0041OpenapiJobInfoRespJobsInner) GetTresBind() string`

GetTresBind returns the TresBind field if non-nil, zero value otherwise.

### GetTresBindOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetTresBindOk() (*string, bool)`

GetTresBindOk returns a tuple with the TresBind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTresBind

`func (o *V0041OpenapiJobInfoRespJobsInner) SetTresBind(v string)`

SetTresBind sets TresBind field to given value.

### HasTresBind

`func (o *V0041OpenapiJobInfoRespJobsInner) HasTresBind() bool`

HasTresBind returns a boolean if a field has been set.

### GetTresFreq

`func (o *V0041OpenapiJobInfoRespJobsInner) GetTresFreq() string`

GetTresFreq returns the TresFreq field if non-nil, zero value otherwise.

### GetTresFreqOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetTresFreqOk() (*string, bool)`

GetTresFreqOk returns a tuple with the TresFreq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTresFreq

`func (o *V0041OpenapiJobInfoRespJobsInner) SetTresFreq(v string)`

SetTresFreq sets TresFreq field to given value.

### HasTresFreq

`func (o *V0041OpenapiJobInfoRespJobsInner) HasTresFreq() bool`

HasTresFreq returns a boolean if a field has been set.

### GetTresPerJob

`func (o *V0041OpenapiJobInfoRespJobsInner) GetTresPerJob() string`

GetTresPerJob returns the TresPerJob field if non-nil, zero value otherwise.

### GetTresPerJobOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetTresPerJobOk() (*string, bool)`

GetTresPerJobOk returns a tuple with the TresPerJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTresPerJob

`func (o *V0041OpenapiJobInfoRespJobsInner) SetTresPerJob(v string)`

SetTresPerJob sets TresPerJob field to given value.

### HasTresPerJob

`func (o *V0041OpenapiJobInfoRespJobsInner) HasTresPerJob() bool`

HasTresPerJob returns a boolean if a field has been set.

### GetTresPerNode

`func (o *V0041OpenapiJobInfoRespJobsInner) GetTresPerNode() string`

GetTresPerNode returns the TresPerNode field if non-nil, zero value otherwise.

### GetTresPerNodeOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetTresPerNodeOk() (*string, bool)`

GetTresPerNodeOk returns a tuple with the TresPerNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTresPerNode

`func (o *V0041OpenapiJobInfoRespJobsInner) SetTresPerNode(v string)`

SetTresPerNode sets TresPerNode field to given value.

### HasTresPerNode

`func (o *V0041OpenapiJobInfoRespJobsInner) HasTresPerNode() bool`

HasTresPerNode returns a boolean if a field has been set.

### GetTresPerSocket

`func (o *V0041OpenapiJobInfoRespJobsInner) GetTresPerSocket() string`

GetTresPerSocket returns the TresPerSocket field if non-nil, zero value otherwise.

### GetTresPerSocketOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetTresPerSocketOk() (*string, bool)`

GetTresPerSocketOk returns a tuple with the TresPerSocket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTresPerSocket

`func (o *V0041OpenapiJobInfoRespJobsInner) SetTresPerSocket(v string)`

SetTresPerSocket sets TresPerSocket field to given value.

### HasTresPerSocket

`func (o *V0041OpenapiJobInfoRespJobsInner) HasTresPerSocket() bool`

HasTresPerSocket returns a boolean if a field has been set.

### GetTresPerTask

`func (o *V0041OpenapiJobInfoRespJobsInner) GetTresPerTask() string`

GetTresPerTask returns the TresPerTask field if non-nil, zero value otherwise.

### GetTresPerTaskOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetTresPerTaskOk() (*string, bool)`

GetTresPerTaskOk returns a tuple with the TresPerTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTresPerTask

`func (o *V0041OpenapiJobInfoRespJobsInner) SetTresPerTask(v string)`

SetTresPerTask sets TresPerTask field to given value.

### HasTresPerTask

`func (o *V0041OpenapiJobInfoRespJobsInner) HasTresPerTask() bool`

HasTresPerTask returns a boolean if a field has been set.

### GetTresReqStr

`func (o *V0041OpenapiJobInfoRespJobsInner) GetTresReqStr() string`

GetTresReqStr returns the TresReqStr field if non-nil, zero value otherwise.

### GetTresReqStrOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetTresReqStrOk() (*string, bool)`

GetTresReqStrOk returns a tuple with the TresReqStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTresReqStr

`func (o *V0041OpenapiJobInfoRespJobsInner) SetTresReqStr(v string)`

SetTresReqStr sets TresReqStr field to given value.

### HasTresReqStr

`func (o *V0041OpenapiJobInfoRespJobsInner) HasTresReqStr() bool`

HasTresReqStr returns a boolean if a field has been set.

### GetTresAllocStr

`func (o *V0041OpenapiJobInfoRespJobsInner) GetTresAllocStr() string`

GetTresAllocStr returns the TresAllocStr field if non-nil, zero value otherwise.

### GetTresAllocStrOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetTresAllocStrOk() (*string, bool)`

GetTresAllocStrOk returns a tuple with the TresAllocStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTresAllocStr

`func (o *V0041OpenapiJobInfoRespJobsInner) SetTresAllocStr(v string)`

SetTresAllocStr sets TresAllocStr field to given value.

### HasTresAllocStr

`func (o *V0041OpenapiJobInfoRespJobsInner) HasTresAllocStr() bool`

HasTresAllocStr returns a boolean if a field has been set.

### GetUserId

`func (o *V0041OpenapiJobInfoRespJobsInner) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *V0041OpenapiJobInfoRespJobsInner) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *V0041OpenapiJobInfoRespJobsInner) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUserName

`func (o *V0041OpenapiJobInfoRespJobsInner) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *V0041OpenapiJobInfoRespJobsInner) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *V0041OpenapiJobInfoRespJobsInner) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetMaximumSwitchWaitTime

`func (o *V0041OpenapiJobInfoRespJobsInner) GetMaximumSwitchWaitTime() int32`

GetMaximumSwitchWaitTime returns the MaximumSwitchWaitTime field if non-nil, zero value otherwise.

### GetMaximumSwitchWaitTimeOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetMaximumSwitchWaitTimeOk() (*int32, bool)`

GetMaximumSwitchWaitTimeOk returns a tuple with the MaximumSwitchWaitTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumSwitchWaitTime

`func (o *V0041OpenapiJobInfoRespJobsInner) SetMaximumSwitchWaitTime(v int32)`

SetMaximumSwitchWaitTime sets MaximumSwitchWaitTime field to given value.

### HasMaximumSwitchWaitTime

`func (o *V0041OpenapiJobInfoRespJobsInner) HasMaximumSwitchWaitTime() bool`

HasMaximumSwitchWaitTime returns a boolean if a field has been set.

### GetWckey

`func (o *V0041OpenapiJobInfoRespJobsInner) GetWckey() string`

GetWckey returns the Wckey field if non-nil, zero value otherwise.

### GetWckeyOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetWckeyOk() (*string, bool)`

GetWckeyOk returns a tuple with the Wckey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWckey

`func (o *V0041OpenapiJobInfoRespJobsInner) SetWckey(v string)`

SetWckey sets Wckey field to given value.

### HasWckey

`func (o *V0041OpenapiJobInfoRespJobsInner) HasWckey() bool`

HasWckey returns a boolean if a field has been set.

### GetCurrentWorkingDirectory

`func (o *V0041OpenapiJobInfoRespJobsInner) GetCurrentWorkingDirectory() string`

GetCurrentWorkingDirectory returns the CurrentWorkingDirectory field if non-nil, zero value otherwise.

### GetCurrentWorkingDirectoryOk

`func (o *V0041OpenapiJobInfoRespJobsInner) GetCurrentWorkingDirectoryOk() (*string, bool)`

GetCurrentWorkingDirectoryOk returns a tuple with the CurrentWorkingDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentWorkingDirectory

`func (o *V0041OpenapiJobInfoRespJobsInner) SetCurrentWorkingDirectory(v string)`

SetCurrentWorkingDirectory sets CurrentWorkingDirectory field to given value.

### HasCurrentWorkingDirectory

`func (o *V0041OpenapiJobInfoRespJobsInner) HasCurrentWorkingDirectory() bool`

HasCurrentWorkingDirectory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


