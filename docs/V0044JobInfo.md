# V0044JobInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** | Account associated with the job | [optional] 
**AccrueTime** | Pointer to [**V0044Uint64NoValStruct**](V0044Uint64NoValStruct.md) |  | [optional] 
**AdminComment** | Pointer to **string** | Arbitrary comment made by administrator | [optional] 
**AllocatingNode** | Pointer to **string** | Local node making the resource allocation | [optional] 
**ArrayJobId** | Pointer to [**V0044Uint32NoValStruct**](V0044Uint32NoValStruct.md) |  | [optional] 
**ArrayTaskId** | Pointer to [**V0044Uint32NoValStruct**](V0044Uint32NoValStruct.md) |  | [optional] 
**ArrayMaxTasks** | Pointer to [**V0044Uint32NoValStruct**](V0044Uint32NoValStruct.md) |  | [optional] 
**ArrayTaskString** | Pointer to **string** | String expression of task IDs in this record | [optional] 
**AssociationId** | Pointer to **int32** | Unique identifier for the association | [optional] 
**BatchFeatures** | Pointer to **string** | Features required for batch script&#39;s node | [optional] 
**BatchFlag** | Pointer to **bool** | True if batch job | [optional] 
**BatchHost** | Pointer to **string** | Name of host running batch script | [optional] 
**Flags** | Pointer to **[]string** | Job flags | [optional] 
**BurstBuffer** | Pointer to **string** | Burst buffer specifications | [optional] 
**BurstBufferState** | Pointer to **string** | Burst buffer state details | [optional] 
**Cluster** | Pointer to **string** | Cluster name | [optional] 
**ClusterFeatures** | Pointer to **string** | List of required cluster features | [optional] 
**Command** | Pointer to **string** | Executed command | [optional] 
**Comment** | Pointer to **string** | Arbitrary comment | [optional] 
**Container** | Pointer to **string** | Absolute path to OCI container bundle | [optional] 
**ContainerId** | Pointer to **string** | OCI container ID | [optional] 
**Contiguous** | Pointer to **bool** | True if job requires contiguous nodes | [optional] 
**CoreSpec** | Pointer to **int32** | Specialized core count | [optional] 
**ThreadSpec** | Pointer to **int32** | Specialized thread count | [optional] 
**CoresPerSocket** | Pointer to [**V0044Uint16NoValStruct**](V0044Uint16NoValStruct.md) |  | [optional] 
**BillableTres** | Pointer to [**V0044Float64NoValStruct**](V0044Float64NoValStruct.md) |  | [optional] 
**CpusPerTask** | Pointer to [**V0044Uint16NoValStruct**](V0044Uint16NoValStruct.md) |  | [optional] 
**CpuFrequencyMinimum** | Pointer to [**V0044Uint32NoValStruct**](V0044Uint32NoValStruct.md) |  | [optional] 
**CpuFrequencyMaximum** | Pointer to [**V0044Uint32NoValStruct**](V0044Uint32NoValStruct.md) |  | [optional] 
**CpuFrequencyGovernor** | Pointer to [**V0044Uint32NoValStruct**](V0044Uint32NoValStruct.md) |  | [optional] 
**CpusPerTres** | Pointer to **string** | Semicolon delimited list of TRES&#x3D;# values indicating how many CPUs should be allocated for each specified TRES (currently only used for gres/gpu) | [optional] 
**Cron** | Pointer to **string** | Time specification for scrontab job | [optional] 
**Deadline** | Pointer to [**V0044Uint64NoValStruct**](V0044Uint64NoValStruct.md) |  | [optional] 
**DelayBoot** | Pointer to [**V0044Uint32NoValStruct**](V0044Uint32NoValStruct.md) |  | [optional] 
**Dependency** | Pointer to **string** | Other jobs that must meet certain criteria before this job can start | [optional] 
**DerivedExitCode** | Pointer to [**V0044ProcessExitCodeVerbose**](V0044ProcessExitCodeVerbose.md) |  | [optional] 
**EligibleTime** | Pointer to [**V0044Uint64NoValStruct**](V0044Uint64NoValStruct.md) |  | [optional] 
**EndTime** | Pointer to [**V0044Uint64NoValStruct**](V0044Uint64NoValStruct.md) |  | [optional] 
**ExcludedNodes** | Pointer to **string** | Comma-separated list of nodes that may not be used | [optional] 
**ExitCode** | Pointer to [**V0044ProcessExitCodeVerbose**](V0044ProcessExitCodeVerbose.md) |  | [optional] 
**Extra** | Pointer to **string** | Arbitrary string used for node filtering if extra constraints are enabled | [optional] 
**FailedNode** | Pointer to **string** | Name of node that caused job failure | [optional] 
**Features** | Pointer to **string** | Comma-separated list of features that are required | [optional] 
**FederationOrigin** | Pointer to **string** | Origin cluster&#39;s name (when using federation) | [optional] 
**FederationSiblingsActive** | Pointer to **string** | Active sibling job names | [optional] 
**FederationSiblingsViable** | Pointer to **string** | Viable sibling job names | [optional] 
**GresDetail** | Pointer to **[]string** |  | [optional] 
**GroupId** | Pointer to **int32** | Group ID of the user that owns the job | [optional] 
**GroupName** | Pointer to **string** | Group name of the user that owns the job | [optional] 
**HetJobId** | Pointer to [**V0044Uint32NoValStruct**](V0044Uint32NoValStruct.md) |  | [optional] 
**HetJobIdSet** | Pointer to **string** | Job ID range for all heterogeneous job components | [optional] 
**HetJobOffset** | Pointer to [**V0044Uint32NoValStruct**](V0044Uint32NoValStruct.md) |  | [optional] 
**JobId** | Pointer to **int32** | Job ID | [optional] 
**JobResources** | Pointer to [**V0044JobRes**](V0044JobRes.md) |  | [optional] 
**JobSizeStr** | Pointer to **[]string** |  | [optional] 
**JobState** | Pointer to **[]string** | Current state | [optional] 
**LastSchedEvaluation** | Pointer to [**V0044Uint64NoValStruct**](V0044Uint64NoValStruct.md) |  | [optional] 
**Licenses** | Pointer to **string** | License(s) required by the job | [optional] 
**LicensesAllocated** | Pointer to **string** | License(s) allocated to the job | [optional] 
**MailType** | Pointer to **[]string** | Mail event type(s) | [optional] 
**MailUser** | Pointer to **string** | User to receive email notifications | [optional] 
**MaxCpus** | Pointer to [**V0044Uint32NoValStruct**](V0044Uint32NoValStruct.md) |  | [optional] 
**MaxNodes** | Pointer to [**V0044Uint32NoValStruct**](V0044Uint32NoValStruct.md) |  | [optional] 
**McsLabel** | Pointer to **string** | Multi-Category Security label on the job | [optional] 
**MemoryPerTres** | Pointer to **string** | Semicolon delimited list of TRES&#x3D;# values indicating how much memory in megabytes should be allocated for each specified TRES (currently only used for gres/gpu) | [optional] 
**Name** | Pointer to **string** | Job name | [optional] 
**Network** | Pointer to **string** | Network specs for the job | [optional] 
**Nodes** | Pointer to **string** | Node(s) allocated to the job | [optional] 
**Nice** | Pointer to **int32** | Requested job priority change | [optional] 
**TasksPerCore** | Pointer to [**V0044Uint16NoValStruct**](V0044Uint16NoValStruct.md) |  | [optional] 
**TasksPerTres** | Pointer to [**V0044Uint16NoValStruct**](V0044Uint16NoValStruct.md) |  | [optional] 
**TasksPerNode** | Pointer to [**V0044Uint16NoValStruct**](V0044Uint16NoValStruct.md) |  | [optional] 
**TasksPerSocket** | Pointer to [**V0044Uint16NoValStruct**](V0044Uint16NoValStruct.md) |  | [optional] 
**TasksPerBoard** | Pointer to [**V0044Uint16NoValStruct**](V0044Uint16NoValStruct.md) |  | [optional] 
**Cpus** | Pointer to [**V0044Uint32NoValStruct**](V0044Uint32NoValStruct.md) |  | [optional] 
**NodeCount** | Pointer to [**V0044Uint32NoValStruct**](V0044Uint32NoValStruct.md) |  | [optional] 
**Tasks** | Pointer to [**V0044Uint32NoValStruct**](V0044Uint32NoValStruct.md) |  | [optional] 
**Partition** | Pointer to **string** | Partition assigned to the job | [optional] 
**Prefer** | Pointer to **string** | Feature(s) the job requested but that are not required | [optional] 
**MemoryPerCpu** | Pointer to [**V0044Uint64NoValStruct**](V0044Uint64NoValStruct.md) |  | [optional] 
**MemoryPerNode** | Pointer to [**V0044Uint64NoValStruct**](V0044Uint64NoValStruct.md) |  | [optional] 
**MinimumCpusPerNode** | Pointer to [**V0044Uint16NoValStruct**](V0044Uint16NoValStruct.md) |  | [optional] 
**MinimumTmpDiskPerNode** | Pointer to [**V0044Uint32NoValStruct**](V0044Uint32NoValStruct.md) |  | [optional] 
**Power** | Pointer to [**V0044JobInfoPower**](V0044JobInfoPower.md) |  | [optional] 
**PreemptTime** | Pointer to [**V0044Uint64NoValStruct**](V0044Uint64NoValStruct.md) |  | [optional] 
**PreemptableTime** | Pointer to [**V0044Uint64NoValStruct**](V0044Uint64NoValStruct.md) |  | [optional] 
**PreSusTime** | Pointer to [**V0044Uint64NoValStruct**](V0044Uint64NoValStruct.md) |  | [optional] 
**Hold** | Pointer to **bool** | Hold (true) or release (false) job (Job held) | [optional] 
**Priority** | Pointer to [**V0044Uint32NoValStruct**](V0044Uint32NoValStruct.md) |  | [optional] 
**PriorityByPartition** | Pointer to [**[]V0044PartPrio**](V0044PartPrio.md) |  | [optional] 
**Profile** | Pointer to **[]string** | Profile used by the acct_gather_profile plugin | [optional] 
**Qos** | Pointer to **string** | Quality of Service assigned to the job, if pending the QOS requested | [optional] 
**Reboot** | Pointer to **bool** | Node reboot requested before start | [optional] 
**RequiredNodes** | Pointer to **string** | Comma-separated list of required nodes | [optional] 
**RequiredSwitches** | Pointer to **int32** | Maximum number of switches | [optional] 
**Requeue** | Pointer to **bool** | Determines whether the job may be requeued | [optional] 
**ResizeTime** | Pointer to [**V0044Uint64NoValStruct**](V0044Uint64NoValStruct.md) |  | [optional] 
**RestartCnt** | Pointer to **int32** | Number of job restarts | [optional] 
**ResvName** | Pointer to **string** | Name of reservation to use | [optional] 
**ScheduledNodes** | Pointer to **string** | List of nodes scheduled to be used for the job | [optional] 
**SegmentSize** | Pointer to **int32** | Requested segment size | [optional] 
**SelinuxContext** | Pointer to **string** | SELinux context | [optional] 
**Shared** | Pointer to **[]string** | How the job can share resources with other jobs, if at all | [optional] 
**StepId** | Pointer to [**V0044SlurmStepId**](V0044SlurmStepId.md) |  | [optional] 
**SocketsPerBoard** | Pointer to **int32** | Number of sockets per board required | [optional] 
**SocketsPerNode** | Pointer to [**V0044Uint16NoValStruct**](V0044Uint16NoValStruct.md) |  | [optional] 
**StartTime** | Pointer to [**V0044Uint64NoValStruct**](V0044Uint64NoValStruct.md) |  | [optional] 
**StateDescription** | Pointer to **string** | Optional details for state_reason | [optional] 
**StateReason** | Pointer to **string** | Reason for current Pending or Failed state | [optional] 
**StandardInput** | Pointer to **string** | Path to stdin file | [optional] 
**StandardOutput** | Pointer to **string** | Path to stdout file | [optional] 
**StandardError** | Pointer to **string** | Path to stderr file | [optional] 
**StdinExpanded** | Pointer to **string** | Job stdin with expanded fields | [optional] 
**StdoutExpanded** | Pointer to **string** | Job stdout with expanded fields | [optional] 
**StderrExpanded** | Pointer to **string** | Job stderr with expanded fields | [optional] 
**SubmitTime** | Pointer to [**V0044Uint64NoValStruct**](V0044Uint64NoValStruct.md) |  | [optional] 
**SubmitLine** | Pointer to **string** | Job submit line (e.g. &#39;sbatch -N3 job.sh job_arg&#39; | [optional] 
**SuspendTime** | Pointer to [**V0044Uint64NoValStruct**](V0044Uint64NoValStruct.md) |  | [optional] 
**SystemComment** | Pointer to **string** | Arbitrary comment from slurmctld | [optional] 
**TimeLimit** | Pointer to [**V0044Uint32NoValStruct**](V0044Uint32NoValStruct.md) |  | [optional] 
**TimeMinimum** | Pointer to [**V0044Uint32NoValStruct**](V0044Uint32NoValStruct.md) |  | [optional] 
**ThreadsPerCore** | Pointer to [**V0044Uint16NoValStruct**](V0044Uint16NoValStruct.md) |  | [optional] 
**TresBind** | Pointer to **string** | Task to TRES binding directives | [optional] 
**TresFreq** | Pointer to **string** | TRES frequency directives | [optional] 
**TresPerJob** | Pointer to **string** | Comma-separated list of TRES&#x3D;# values to be allocated per job | [optional] 
**TresPerNode** | Pointer to **string** | Comma-separated list of TRES&#x3D;# values to be allocated per node | [optional] 
**TresPerSocket** | Pointer to **string** | Comma-separated list of TRES&#x3D;# values to be allocated per socket | [optional] 
**TresPerTask** | Pointer to **string** | Comma-separated list of TRES&#x3D;# values to be allocated per task | [optional] 
**TresReqStr** | Pointer to **string** | TRES requested by the job | [optional] 
**TresAllocStr** | Pointer to **string** | TRES used by the job | [optional] 
**UserId** | Pointer to **int32** | User ID that owns the job | [optional] 
**UserName** | Pointer to **string** | User name that owns the job | [optional] 
**MaximumSwitchWaitTime** | Pointer to **int32** | Maximum time to wait for switches in seconds | [optional] 
**Wckey** | Pointer to **string** | Workload characterization key | [optional] 
**CurrentWorkingDirectory** | Pointer to **string** | Working directory to use for the job | [optional] 

## Methods

### NewV0044JobInfo

`func NewV0044JobInfo() *V0044JobInfo`

NewV0044JobInfo instantiates a new V0044JobInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0044JobInfoWithDefaults

`func NewV0044JobInfoWithDefaults() *V0044JobInfo`

NewV0044JobInfoWithDefaults instantiates a new V0044JobInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *V0044JobInfo) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *V0044JobInfo) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *V0044JobInfo) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *V0044JobInfo) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetAccrueTime

`func (o *V0044JobInfo) GetAccrueTime() V0044Uint64NoValStruct`

GetAccrueTime returns the AccrueTime field if non-nil, zero value otherwise.

### GetAccrueTimeOk

`func (o *V0044JobInfo) GetAccrueTimeOk() (*V0044Uint64NoValStruct, bool)`

GetAccrueTimeOk returns a tuple with the AccrueTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccrueTime

`func (o *V0044JobInfo) SetAccrueTime(v V0044Uint64NoValStruct)`

SetAccrueTime sets AccrueTime field to given value.

### HasAccrueTime

`func (o *V0044JobInfo) HasAccrueTime() bool`

HasAccrueTime returns a boolean if a field has been set.

### GetAdminComment

`func (o *V0044JobInfo) GetAdminComment() string`

GetAdminComment returns the AdminComment field if non-nil, zero value otherwise.

### GetAdminCommentOk

`func (o *V0044JobInfo) GetAdminCommentOk() (*string, bool)`

GetAdminCommentOk returns a tuple with the AdminComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminComment

`func (o *V0044JobInfo) SetAdminComment(v string)`

SetAdminComment sets AdminComment field to given value.

### HasAdminComment

`func (o *V0044JobInfo) HasAdminComment() bool`

HasAdminComment returns a boolean if a field has been set.

### GetAllocatingNode

`func (o *V0044JobInfo) GetAllocatingNode() string`

GetAllocatingNode returns the AllocatingNode field if non-nil, zero value otherwise.

### GetAllocatingNodeOk

`func (o *V0044JobInfo) GetAllocatingNodeOk() (*string, bool)`

GetAllocatingNodeOk returns a tuple with the AllocatingNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatingNode

`func (o *V0044JobInfo) SetAllocatingNode(v string)`

SetAllocatingNode sets AllocatingNode field to given value.

### HasAllocatingNode

`func (o *V0044JobInfo) HasAllocatingNode() bool`

HasAllocatingNode returns a boolean if a field has been set.

### GetArrayJobId

`func (o *V0044JobInfo) GetArrayJobId() V0044Uint32NoValStruct`

GetArrayJobId returns the ArrayJobId field if non-nil, zero value otherwise.

### GetArrayJobIdOk

`func (o *V0044JobInfo) GetArrayJobIdOk() (*V0044Uint32NoValStruct, bool)`

GetArrayJobIdOk returns a tuple with the ArrayJobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrayJobId

`func (o *V0044JobInfo) SetArrayJobId(v V0044Uint32NoValStruct)`

SetArrayJobId sets ArrayJobId field to given value.

### HasArrayJobId

`func (o *V0044JobInfo) HasArrayJobId() bool`

HasArrayJobId returns a boolean if a field has been set.

### GetArrayTaskId

`func (o *V0044JobInfo) GetArrayTaskId() V0044Uint32NoValStruct`

GetArrayTaskId returns the ArrayTaskId field if non-nil, zero value otherwise.

### GetArrayTaskIdOk

`func (o *V0044JobInfo) GetArrayTaskIdOk() (*V0044Uint32NoValStruct, bool)`

GetArrayTaskIdOk returns a tuple with the ArrayTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrayTaskId

`func (o *V0044JobInfo) SetArrayTaskId(v V0044Uint32NoValStruct)`

SetArrayTaskId sets ArrayTaskId field to given value.

### HasArrayTaskId

`func (o *V0044JobInfo) HasArrayTaskId() bool`

HasArrayTaskId returns a boolean if a field has been set.

### GetArrayMaxTasks

`func (o *V0044JobInfo) GetArrayMaxTasks() V0044Uint32NoValStruct`

GetArrayMaxTasks returns the ArrayMaxTasks field if non-nil, zero value otherwise.

### GetArrayMaxTasksOk

`func (o *V0044JobInfo) GetArrayMaxTasksOk() (*V0044Uint32NoValStruct, bool)`

GetArrayMaxTasksOk returns a tuple with the ArrayMaxTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrayMaxTasks

`func (o *V0044JobInfo) SetArrayMaxTasks(v V0044Uint32NoValStruct)`

SetArrayMaxTasks sets ArrayMaxTasks field to given value.

### HasArrayMaxTasks

`func (o *V0044JobInfo) HasArrayMaxTasks() bool`

HasArrayMaxTasks returns a boolean if a field has been set.

### GetArrayTaskString

`func (o *V0044JobInfo) GetArrayTaskString() string`

GetArrayTaskString returns the ArrayTaskString field if non-nil, zero value otherwise.

### GetArrayTaskStringOk

`func (o *V0044JobInfo) GetArrayTaskStringOk() (*string, bool)`

GetArrayTaskStringOk returns a tuple with the ArrayTaskString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrayTaskString

`func (o *V0044JobInfo) SetArrayTaskString(v string)`

SetArrayTaskString sets ArrayTaskString field to given value.

### HasArrayTaskString

`func (o *V0044JobInfo) HasArrayTaskString() bool`

HasArrayTaskString returns a boolean if a field has been set.

### GetAssociationId

`func (o *V0044JobInfo) GetAssociationId() int32`

GetAssociationId returns the AssociationId field if non-nil, zero value otherwise.

### GetAssociationIdOk

`func (o *V0044JobInfo) GetAssociationIdOk() (*int32, bool)`

GetAssociationIdOk returns a tuple with the AssociationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociationId

`func (o *V0044JobInfo) SetAssociationId(v int32)`

SetAssociationId sets AssociationId field to given value.

### HasAssociationId

`func (o *V0044JobInfo) HasAssociationId() bool`

HasAssociationId returns a boolean if a field has been set.

### GetBatchFeatures

`func (o *V0044JobInfo) GetBatchFeatures() string`

GetBatchFeatures returns the BatchFeatures field if non-nil, zero value otherwise.

### GetBatchFeaturesOk

`func (o *V0044JobInfo) GetBatchFeaturesOk() (*string, bool)`

GetBatchFeaturesOk returns a tuple with the BatchFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchFeatures

`func (o *V0044JobInfo) SetBatchFeatures(v string)`

SetBatchFeatures sets BatchFeatures field to given value.

### HasBatchFeatures

`func (o *V0044JobInfo) HasBatchFeatures() bool`

HasBatchFeatures returns a boolean if a field has been set.

### GetBatchFlag

`func (o *V0044JobInfo) GetBatchFlag() bool`

GetBatchFlag returns the BatchFlag field if non-nil, zero value otherwise.

### GetBatchFlagOk

`func (o *V0044JobInfo) GetBatchFlagOk() (*bool, bool)`

GetBatchFlagOk returns a tuple with the BatchFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchFlag

`func (o *V0044JobInfo) SetBatchFlag(v bool)`

SetBatchFlag sets BatchFlag field to given value.

### HasBatchFlag

`func (o *V0044JobInfo) HasBatchFlag() bool`

HasBatchFlag returns a boolean if a field has been set.

### GetBatchHost

`func (o *V0044JobInfo) GetBatchHost() string`

GetBatchHost returns the BatchHost field if non-nil, zero value otherwise.

### GetBatchHostOk

`func (o *V0044JobInfo) GetBatchHostOk() (*string, bool)`

GetBatchHostOk returns a tuple with the BatchHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchHost

`func (o *V0044JobInfo) SetBatchHost(v string)`

SetBatchHost sets BatchHost field to given value.

### HasBatchHost

`func (o *V0044JobInfo) HasBatchHost() bool`

HasBatchHost returns a boolean if a field has been set.

### GetFlags

`func (o *V0044JobInfo) GetFlags() []string`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *V0044JobInfo) GetFlagsOk() (*[]string, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *V0044JobInfo) SetFlags(v []string)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *V0044JobInfo) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetBurstBuffer

`func (o *V0044JobInfo) GetBurstBuffer() string`

GetBurstBuffer returns the BurstBuffer field if non-nil, zero value otherwise.

### GetBurstBufferOk

`func (o *V0044JobInfo) GetBurstBufferOk() (*string, bool)`

GetBurstBufferOk returns a tuple with the BurstBuffer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurstBuffer

`func (o *V0044JobInfo) SetBurstBuffer(v string)`

SetBurstBuffer sets BurstBuffer field to given value.

### HasBurstBuffer

`func (o *V0044JobInfo) HasBurstBuffer() bool`

HasBurstBuffer returns a boolean if a field has been set.

### GetBurstBufferState

`func (o *V0044JobInfo) GetBurstBufferState() string`

GetBurstBufferState returns the BurstBufferState field if non-nil, zero value otherwise.

### GetBurstBufferStateOk

`func (o *V0044JobInfo) GetBurstBufferStateOk() (*string, bool)`

GetBurstBufferStateOk returns a tuple with the BurstBufferState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurstBufferState

`func (o *V0044JobInfo) SetBurstBufferState(v string)`

SetBurstBufferState sets BurstBufferState field to given value.

### HasBurstBufferState

`func (o *V0044JobInfo) HasBurstBufferState() bool`

HasBurstBufferState returns a boolean if a field has been set.

### GetCluster

`func (o *V0044JobInfo) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *V0044JobInfo) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *V0044JobInfo) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *V0044JobInfo) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetClusterFeatures

`func (o *V0044JobInfo) GetClusterFeatures() string`

GetClusterFeatures returns the ClusterFeatures field if non-nil, zero value otherwise.

### GetClusterFeaturesOk

`func (o *V0044JobInfo) GetClusterFeaturesOk() (*string, bool)`

GetClusterFeaturesOk returns a tuple with the ClusterFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterFeatures

`func (o *V0044JobInfo) SetClusterFeatures(v string)`

SetClusterFeatures sets ClusterFeatures field to given value.

### HasClusterFeatures

`func (o *V0044JobInfo) HasClusterFeatures() bool`

HasClusterFeatures returns a boolean if a field has been set.

### GetCommand

`func (o *V0044JobInfo) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *V0044JobInfo) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *V0044JobInfo) SetCommand(v string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *V0044JobInfo) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### GetComment

`func (o *V0044JobInfo) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *V0044JobInfo) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *V0044JobInfo) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *V0044JobInfo) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetContainer

`func (o *V0044JobInfo) GetContainer() string`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *V0044JobInfo) GetContainerOk() (*string, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *V0044JobInfo) SetContainer(v string)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *V0044JobInfo) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### GetContainerId

`func (o *V0044JobInfo) GetContainerId() string`

GetContainerId returns the ContainerId field if non-nil, zero value otherwise.

### GetContainerIdOk

`func (o *V0044JobInfo) GetContainerIdOk() (*string, bool)`

GetContainerIdOk returns a tuple with the ContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerId

`func (o *V0044JobInfo) SetContainerId(v string)`

SetContainerId sets ContainerId field to given value.

### HasContainerId

`func (o *V0044JobInfo) HasContainerId() bool`

HasContainerId returns a boolean if a field has been set.

### GetContiguous

`func (o *V0044JobInfo) GetContiguous() bool`

GetContiguous returns the Contiguous field if non-nil, zero value otherwise.

### GetContiguousOk

`func (o *V0044JobInfo) GetContiguousOk() (*bool, bool)`

GetContiguousOk returns a tuple with the Contiguous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContiguous

`func (o *V0044JobInfo) SetContiguous(v bool)`

SetContiguous sets Contiguous field to given value.

### HasContiguous

`func (o *V0044JobInfo) HasContiguous() bool`

HasContiguous returns a boolean if a field has been set.

### GetCoreSpec

`func (o *V0044JobInfo) GetCoreSpec() int32`

GetCoreSpec returns the CoreSpec field if non-nil, zero value otherwise.

### GetCoreSpecOk

`func (o *V0044JobInfo) GetCoreSpecOk() (*int32, bool)`

GetCoreSpecOk returns a tuple with the CoreSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoreSpec

`func (o *V0044JobInfo) SetCoreSpec(v int32)`

SetCoreSpec sets CoreSpec field to given value.

### HasCoreSpec

`func (o *V0044JobInfo) HasCoreSpec() bool`

HasCoreSpec returns a boolean if a field has been set.

### GetThreadSpec

`func (o *V0044JobInfo) GetThreadSpec() int32`

GetThreadSpec returns the ThreadSpec field if non-nil, zero value otherwise.

### GetThreadSpecOk

`func (o *V0044JobInfo) GetThreadSpecOk() (*int32, bool)`

GetThreadSpecOk returns a tuple with the ThreadSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadSpec

`func (o *V0044JobInfo) SetThreadSpec(v int32)`

SetThreadSpec sets ThreadSpec field to given value.

### HasThreadSpec

`func (o *V0044JobInfo) HasThreadSpec() bool`

HasThreadSpec returns a boolean if a field has been set.

### GetCoresPerSocket

`func (o *V0044JobInfo) GetCoresPerSocket() V0044Uint16NoValStruct`

GetCoresPerSocket returns the CoresPerSocket field if non-nil, zero value otherwise.

### GetCoresPerSocketOk

`func (o *V0044JobInfo) GetCoresPerSocketOk() (*V0044Uint16NoValStruct, bool)`

GetCoresPerSocketOk returns a tuple with the CoresPerSocket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoresPerSocket

`func (o *V0044JobInfo) SetCoresPerSocket(v V0044Uint16NoValStruct)`

SetCoresPerSocket sets CoresPerSocket field to given value.

### HasCoresPerSocket

`func (o *V0044JobInfo) HasCoresPerSocket() bool`

HasCoresPerSocket returns a boolean if a field has been set.

### GetBillableTres

`func (o *V0044JobInfo) GetBillableTres() V0044Float64NoValStruct`

GetBillableTres returns the BillableTres field if non-nil, zero value otherwise.

### GetBillableTresOk

`func (o *V0044JobInfo) GetBillableTresOk() (*V0044Float64NoValStruct, bool)`

GetBillableTresOk returns a tuple with the BillableTres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillableTres

`func (o *V0044JobInfo) SetBillableTres(v V0044Float64NoValStruct)`

SetBillableTres sets BillableTres field to given value.

### HasBillableTres

`func (o *V0044JobInfo) HasBillableTres() bool`

HasBillableTres returns a boolean if a field has been set.

### GetCpusPerTask

`func (o *V0044JobInfo) GetCpusPerTask() V0044Uint16NoValStruct`

GetCpusPerTask returns the CpusPerTask field if non-nil, zero value otherwise.

### GetCpusPerTaskOk

`func (o *V0044JobInfo) GetCpusPerTaskOk() (*V0044Uint16NoValStruct, bool)`

GetCpusPerTaskOk returns a tuple with the CpusPerTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpusPerTask

`func (o *V0044JobInfo) SetCpusPerTask(v V0044Uint16NoValStruct)`

SetCpusPerTask sets CpusPerTask field to given value.

### HasCpusPerTask

`func (o *V0044JobInfo) HasCpusPerTask() bool`

HasCpusPerTask returns a boolean if a field has been set.

### GetCpuFrequencyMinimum

`func (o *V0044JobInfo) GetCpuFrequencyMinimum() V0044Uint32NoValStruct`

GetCpuFrequencyMinimum returns the CpuFrequencyMinimum field if non-nil, zero value otherwise.

### GetCpuFrequencyMinimumOk

`func (o *V0044JobInfo) GetCpuFrequencyMinimumOk() (*V0044Uint32NoValStruct, bool)`

GetCpuFrequencyMinimumOk returns a tuple with the CpuFrequencyMinimum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuFrequencyMinimum

`func (o *V0044JobInfo) SetCpuFrequencyMinimum(v V0044Uint32NoValStruct)`

SetCpuFrequencyMinimum sets CpuFrequencyMinimum field to given value.

### HasCpuFrequencyMinimum

`func (o *V0044JobInfo) HasCpuFrequencyMinimum() bool`

HasCpuFrequencyMinimum returns a boolean if a field has been set.

### GetCpuFrequencyMaximum

`func (o *V0044JobInfo) GetCpuFrequencyMaximum() V0044Uint32NoValStruct`

GetCpuFrequencyMaximum returns the CpuFrequencyMaximum field if non-nil, zero value otherwise.

### GetCpuFrequencyMaximumOk

`func (o *V0044JobInfo) GetCpuFrequencyMaximumOk() (*V0044Uint32NoValStruct, bool)`

GetCpuFrequencyMaximumOk returns a tuple with the CpuFrequencyMaximum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuFrequencyMaximum

`func (o *V0044JobInfo) SetCpuFrequencyMaximum(v V0044Uint32NoValStruct)`

SetCpuFrequencyMaximum sets CpuFrequencyMaximum field to given value.

### HasCpuFrequencyMaximum

`func (o *V0044JobInfo) HasCpuFrequencyMaximum() bool`

HasCpuFrequencyMaximum returns a boolean if a field has been set.

### GetCpuFrequencyGovernor

`func (o *V0044JobInfo) GetCpuFrequencyGovernor() V0044Uint32NoValStruct`

GetCpuFrequencyGovernor returns the CpuFrequencyGovernor field if non-nil, zero value otherwise.

### GetCpuFrequencyGovernorOk

`func (o *V0044JobInfo) GetCpuFrequencyGovernorOk() (*V0044Uint32NoValStruct, bool)`

GetCpuFrequencyGovernorOk returns a tuple with the CpuFrequencyGovernor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuFrequencyGovernor

`func (o *V0044JobInfo) SetCpuFrequencyGovernor(v V0044Uint32NoValStruct)`

SetCpuFrequencyGovernor sets CpuFrequencyGovernor field to given value.

### HasCpuFrequencyGovernor

`func (o *V0044JobInfo) HasCpuFrequencyGovernor() bool`

HasCpuFrequencyGovernor returns a boolean if a field has been set.

### GetCpusPerTres

`func (o *V0044JobInfo) GetCpusPerTres() string`

GetCpusPerTres returns the CpusPerTres field if non-nil, zero value otherwise.

### GetCpusPerTresOk

`func (o *V0044JobInfo) GetCpusPerTresOk() (*string, bool)`

GetCpusPerTresOk returns a tuple with the CpusPerTres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpusPerTres

`func (o *V0044JobInfo) SetCpusPerTres(v string)`

SetCpusPerTres sets CpusPerTres field to given value.

### HasCpusPerTres

`func (o *V0044JobInfo) HasCpusPerTres() bool`

HasCpusPerTres returns a boolean if a field has been set.

### GetCron

`func (o *V0044JobInfo) GetCron() string`

GetCron returns the Cron field if non-nil, zero value otherwise.

### GetCronOk

`func (o *V0044JobInfo) GetCronOk() (*string, bool)`

GetCronOk returns a tuple with the Cron field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCron

`func (o *V0044JobInfo) SetCron(v string)`

SetCron sets Cron field to given value.

### HasCron

`func (o *V0044JobInfo) HasCron() bool`

HasCron returns a boolean if a field has been set.

### GetDeadline

`func (o *V0044JobInfo) GetDeadline() V0044Uint64NoValStruct`

GetDeadline returns the Deadline field if non-nil, zero value otherwise.

### GetDeadlineOk

`func (o *V0044JobInfo) GetDeadlineOk() (*V0044Uint64NoValStruct, bool)`

GetDeadlineOk returns a tuple with the Deadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadline

`func (o *V0044JobInfo) SetDeadline(v V0044Uint64NoValStruct)`

SetDeadline sets Deadline field to given value.

### HasDeadline

`func (o *V0044JobInfo) HasDeadline() bool`

HasDeadline returns a boolean if a field has been set.

### GetDelayBoot

`func (o *V0044JobInfo) GetDelayBoot() V0044Uint32NoValStruct`

GetDelayBoot returns the DelayBoot field if non-nil, zero value otherwise.

### GetDelayBootOk

`func (o *V0044JobInfo) GetDelayBootOk() (*V0044Uint32NoValStruct, bool)`

GetDelayBootOk returns a tuple with the DelayBoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayBoot

`func (o *V0044JobInfo) SetDelayBoot(v V0044Uint32NoValStruct)`

SetDelayBoot sets DelayBoot field to given value.

### HasDelayBoot

`func (o *V0044JobInfo) HasDelayBoot() bool`

HasDelayBoot returns a boolean if a field has been set.

### GetDependency

`func (o *V0044JobInfo) GetDependency() string`

GetDependency returns the Dependency field if non-nil, zero value otherwise.

### GetDependencyOk

`func (o *V0044JobInfo) GetDependencyOk() (*string, bool)`

GetDependencyOk returns a tuple with the Dependency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependency

`func (o *V0044JobInfo) SetDependency(v string)`

SetDependency sets Dependency field to given value.

### HasDependency

`func (o *V0044JobInfo) HasDependency() bool`

HasDependency returns a boolean if a field has been set.

### GetDerivedExitCode

`func (o *V0044JobInfo) GetDerivedExitCode() V0044ProcessExitCodeVerbose`

GetDerivedExitCode returns the DerivedExitCode field if non-nil, zero value otherwise.

### GetDerivedExitCodeOk

`func (o *V0044JobInfo) GetDerivedExitCodeOk() (*V0044ProcessExitCodeVerbose, bool)`

GetDerivedExitCodeOk returns a tuple with the DerivedExitCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerivedExitCode

`func (o *V0044JobInfo) SetDerivedExitCode(v V0044ProcessExitCodeVerbose)`

SetDerivedExitCode sets DerivedExitCode field to given value.

### HasDerivedExitCode

`func (o *V0044JobInfo) HasDerivedExitCode() bool`

HasDerivedExitCode returns a boolean if a field has been set.

### GetEligibleTime

`func (o *V0044JobInfo) GetEligibleTime() V0044Uint64NoValStruct`

GetEligibleTime returns the EligibleTime field if non-nil, zero value otherwise.

### GetEligibleTimeOk

`func (o *V0044JobInfo) GetEligibleTimeOk() (*V0044Uint64NoValStruct, bool)`

GetEligibleTimeOk returns a tuple with the EligibleTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEligibleTime

`func (o *V0044JobInfo) SetEligibleTime(v V0044Uint64NoValStruct)`

SetEligibleTime sets EligibleTime field to given value.

### HasEligibleTime

`func (o *V0044JobInfo) HasEligibleTime() bool`

HasEligibleTime returns a boolean if a field has been set.

### GetEndTime

`func (o *V0044JobInfo) GetEndTime() V0044Uint64NoValStruct`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *V0044JobInfo) GetEndTimeOk() (*V0044Uint64NoValStruct, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *V0044JobInfo) SetEndTime(v V0044Uint64NoValStruct)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *V0044JobInfo) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetExcludedNodes

`func (o *V0044JobInfo) GetExcludedNodes() string`

GetExcludedNodes returns the ExcludedNodes field if non-nil, zero value otherwise.

### GetExcludedNodesOk

`func (o *V0044JobInfo) GetExcludedNodesOk() (*string, bool)`

GetExcludedNodesOk returns a tuple with the ExcludedNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedNodes

`func (o *V0044JobInfo) SetExcludedNodes(v string)`

SetExcludedNodes sets ExcludedNodes field to given value.

### HasExcludedNodes

`func (o *V0044JobInfo) HasExcludedNodes() bool`

HasExcludedNodes returns a boolean if a field has been set.

### GetExitCode

`func (o *V0044JobInfo) GetExitCode() V0044ProcessExitCodeVerbose`

GetExitCode returns the ExitCode field if non-nil, zero value otherwise.

### GetExitCodeOk

`func (o *V0044JobInfo) GetExitCodeOk() (*V0044ProcessExitCodeVerbose, bool)`

GetExitCodeOk returns a tuple with the ExitCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitCode

`func (o *V0044JobInfo) SetExitCode(v V0044ProcessExitCodeVerbose)`

SetExitCode sets ExitCode field to given value.

### HasExitCode

`func (o *V0044JobInfo) HasExitCode() bool`

HasExitCode returns a boolean if a field has been set.

### GetExtra

`func (o *V0044JobInfo) GetExtra() string`

GetExtra returns the Extra field if non-nil, zero value otherwise.

### GetExtraOk

`func (o *V0044JobInfo) GetExtraOk() (*string, bool)`

GetExtraOk returns a tuple with the Extra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtra

`func (o *V0044JobInfo) SetExtra(v string)`

SetExtra sets Extra field to given value.

### HasExtra

`func (o *V0044JobInfo) HasExtra() bool`

HasExtra returns a boolean if a field has been set.

### GetFailedNode

`func (o *V0044JobInfo) GetFailedNode() string`

GetFailedNode returns the FailedNode field if non-nil, zero value otherwise.

### GetFailedNodeOk

`func (o *V0044JobInfo) GetFailedNodeOk() (*string, bool)`

GetFailedNodeOk returns a tuple with the FailedNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedNode

`func (o *V0044JobInfo) SetFailedNode(v string)`

SetFailedNode sets FailedNode field to given value.

### HasFailedNode

`func (o *V0044JobInfo) HasFailedNode() bool`

HasFailedNode returns a boolean if a field has been set.

### GetFeatures

`func (o *V0044JobInfo) GetFeatures() string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *V0044JobInfo) GetFeaturesOk() (*string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *V0044JobInfo) SetFeatures(v string)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *V0044JobInfo) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetFederationOrigin

`func (o *V0044JobInfo) GetFederationOrigin() string`

GetFederationOrigin returns the FederationOrigin field if non-nil, zero value otherwise.

### GetFederationOriginOk

`func (o *V0044JobInfo) GetFederationOriginOk() (*string, bool)`

GetFederationOriginOk returns a tuple with the FederationOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederationOrigin

`func (o *V0044JobInfo) SetFederationOrigin(v string)`

SetFederationOrigin sets FederationOrigin field to given value.

### HasFederationOrigin

`func (o *V0044JobInfo) HasFederationOrigin() bool`

HasFederationOrigin returns a boolean if a field has been set.

### GetFederationSiblingsActive

`func (o *V0044JobInfo) GetFederationSiblingsActive() string`

GetFederationSiblingsActive returns the FederationSiblingsActive field if non-nil, zero value otherwise.

### GetFederationSiblingsActiveOk

`func (o *V0044JobInfo) GetFederationSiblingsActiveOk() (*string, bool)`

GetFederationSiblingsActiveOk returns a tuple with the FederationSiblingsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederationSiblingsActive

`func (o *V0044JobInfo) SetFederationSiblingsActive(v string)`

SetFederationSiblingsActive sets FederationSiblingsActive field to given value.

### HasFederationSiblingsActive

`func (o *V0044JobInfo) HasFederationSiblingsActive() bool`

HasFederationSiblingsActive returns a boolean if a field has been set.

### GetFederationSiblingsViable

`func (o *V0044JobInfo) GetFederationSiblingsViable() string`

GetFederationSiblingsViable returns the FederationSiblingsViable field if non-nil, zero value otherwise.

### GetFederationSiblingsViableOk

`func (o *V0044JobInfo) GetFederationSiblingsViableOk() (*string, bool)`

GetFederationSiblingsViableOk returns a tuple with the FederationSiblingsViable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederationSiblingsViable

`func (o *V0044JobInfo) SetFederationSiblingsViable(v string)`

SetFederationSiblingsViable sets FederationSiblingsViable field to given value.

### HasFederationSiblingsViable

`func (o *V0044JobInfo) HasFederationSiblingsViable() bool`

HasFederationSiblingsViable returns a boolean if a field has been set.

### GetGresDetail

`func (o *V0044JobInfo) GetGresDetail() []string`

GetGresDetail returns the GresDetail field if non-nil, zero value otherwise.

### GetGresDetailOk

`func (o *V0044JobInfo) GetGresDetailOk() (*[]string, bool)`

GetGresDetailOk returns a tuple with the GresDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGresDetail

`func (o *V0044JobInfo) SetGresDetail(v []string)`

SetGresDetail sets GresDetail field to given value.

### HasGresDetail

`func (o *V0044JobInfo) HasGresDetail() bool`

HasGresDetail returns a boolean if a field has been set.

### GetGroupId

`func (o *V0044JobInfo) GetGroupId() int32`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *V0044JobInfo) GetGroupIdOk() (*int32, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *V0044JobInfo) SetGroupId(v int32)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *V0044JobInfo) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetGroupName

`func (o *V0044JobInfo) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *V0044JobInfo) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *V0044JobInfo) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *V0044JobInfo) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetHetJobId

`func (o *V0044JobInfo) GetHetJobId() V0044Uint32NoValStruct`

GetHetJobId returns the HetJobId field if non-nil, zero value otherwise.

### GetHetJobIdOk

`func (o *V0044JobInfo) GetHetJobIdOk() (*V0044Uint32NoValStruct, bool)`

GetHetJobIdOk returns a tuple with the HetJobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHetJobId

`func (o *V0044JobInfo) SetHetJobId(v V0044Uint32NoValStruct)`

SetHetJobId sets HetJobId field to given value.

### HasHetJobId

`func (o *V0044JobInfo) HasHetJobId() bool`

HasHetJobId returns a boolean if a field has been set.

### GetHetJobIdSet

`func (o *V0044JobInfo) GetHetJobIdSet() string`

GetHetJobIdSet returns the HetJobIdSet field if non-nil, zero value otherwise.

### GetHetJobIdSetOk

`func (o *V0044JobInfo) GetHetJobIdSetOk() (*string, bool)`

GetHetJobIdSetOk returns a tuple with the HetJobIdSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHetJobIdSet

`func (o *V0044JobInfo) SetHetJobIdSet(v string)`

SetHetJobIdSet sets HetJobIdSet field to given value.

### HasHetJobIdSet

`func (o *V0044JobInfo) HasHetJobIdSet() bool`

HasHetJobIdSet returns a boolean if a field has been set.

### GetHetJobOffset

`func (o *V0044JobInfo) GetHetJobOffset() V0044Uint32NoValStruct`

GetHetJobOffset returns the HetJobOffset field if non-nil, zero value otherwise.

### GetHetJobOffsetOk

`func (o *V0044JobInfo) GetHetJobOffsetOk() (*V0044Uint32NoValStruct, bool)`

GetHetJobOffsetOk returns a tuple with the HetJobOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHetJobOffset

`func (o *V0044JobInfo) SetHetJobOffset(v V0044Uint32NoValStruct)`

SetHetJobOffset sets HetJobOffset field to given value.

### HasHetJobOffset

`func (o *V0044JobInfo) HasHetJobOffset() bool`

HasHetJobOffset returns a boolean if a field has been set.

### GetJobId

`func (o *V0044JobInfo) GetJobId() int32`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *V0044JobInfo) GetJobIdOk() (*int32, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *V0044JobInfo) SetJobId(v int32)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *V0044JobInfo) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetJobResources

`func (o *V0044JobInfo) GetJobResources() V0044JobRes`

GetJobResources returns the JobResources field if non-nil, zero value otherwise.

### GetJobResourcesOk

`func (o *V0044JobInfo) GetJobResourcesOk() (*V0044JobRes, bool)`

GetJobResourcesOk returns a tuple with the JobResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobResources

`func (o *V0044JobInfo) SetJobResources(v V0044JobRes)`

SetJobResources sets JobResources field to given value.

### HasJobResources

`func (o *V0044JobInfo) HasJobResources() bool`

HasJobResources returns a boolean if a field has been set.

### GetJobSizeStr

`func (o *V0044JobInfo) GetJobSizeStr() []string`

GetJobSizeStr returns the JobSizeStr field if non-nil, zero value otherwise.

### GetJobSizeStrOk

`func (o *V0044JobInfo) GetJobSizeStrOk() (*[]string, bool)`

GetJobSizeStrOk returns a tuple with the JobSizeStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobSizeStr

`func (o *V0044JobInfo) SetJobSizeStr(v []string)`

SetJobSizeStr sets JobSizeStr field to given value.

### HasJobSizeStr

`func (o *V0044JobInfo) HasJobSizeStr() bool`

HasJobSizeStr returns a boolean if a field has been set.

### GetJobState

`func (o *V0044JobInfo) GetJobState() []string`

GetJobState returns the JobState field if non-nil, zero value otherwise.

### GetJobStateOk

`func (o *V0044JobInfo) GetJobStateOk() (*[]string, bool)`

GetJobStateOk returns a tuple with the JobState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobState

`func (o *V0044JobInfo) SetJobState(v []string)`

SetJobState sets JobState field to given value.

### HasJobState

`func (o *V0044JobInfo) HasJobState() bool`

HasJobState returns a boolean if a field has been set.

### GetLastSchedEvaluation

`func (o *V0044JobInfo) GetLastSchedEvaluation() V0044Uint64NoValStruct`

GetLastSchedEvaluation returns the LastSchedEvaluation field if non-nil, zero value otherwise.

### GetLastSchedEvaluationOk

`func (o *V0044JobInfo) GetLastSchedEvaluationOk() (*V0044Uint64NoValStruct, bool)`

GetLastSchedEvaluationOk returns a tuple with the LastSchedEvaluation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSchedEvaluation

`func (o *V0044JobInfo) SetLastSchedEvaluation(v V0044Uint64NoValStruct)`

SetLastSchedEvaluation sets LastSchedEvaluation field to given value.

### HasLastSchedEvaluation

`func (o *V0044JobInfo) HasLastSchedEvaluation() bool`

HasLastSchedEvaluation returns a boolean if a field has been set.

### GetLicenses

`func (o *V0044JobInfo) GetLicenses() string`

GetLicenses returns the Licenses field if non-nil, zero value otherwise.

### GetLicensesOk

`func (o *V0044JobInfo) GetLicensesOk() (*string, bool)`

GetLicensesOk returns a tuple with the Licenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenses

`func (o *V0044JobInfo) SetLicenses(v string)`

SetLicenses sets Licenses field to given value.

### HasLicenses

`func (o *V0044JobInfo) HasLicenses() bool`

HasLicenses returns a boolean if a field has been set.

### GetLicensesAllocated

`func (o *V0044JobInfo) GetLicensesAllocated() string`

GetLicensesAllocated returns the LicensesAllocated field if non-nil, zero value otherwise.

### GetLicensesAllocatedOk

`func (o *V0044JobInfo) GetLicensesAllocatedOk() (*string, bool)`

GetLicensesAllocatedOk returns a tuple with the LicensesAllocated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensesAllocated

`func (o *V0044JobInfo) SetLicensesAllocated(v string)`

SetLicensesAllocated sets LicensesAllocated field to given value.

### HasLicensesAllocated

`func (o *V0044JobInfo) HasLicensesAllocated() bool`

HasLicensesAllocated returns a boolean if a field has been set.

### GetMailType

`func (o *V0044JobInfo) GetMailType() []string`

GetMailType returns the MailType field if non-nil, zero value otherwise.

### GetMailTypeOk

`func (o *V0044JobInfo) GetMailTypeOk() (*[]string, bool)`

GetMailTypeOk returns a tuple with the MailType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailType

`func (o *V0044JobInfo) SetMailType(v []string)`

SetMailType sets MailType field to given value.

### HasMailType

`func (o *V0044JobInfo) HasMailType() bool`

HasMailType returns a boolean if a field has been set.

### GetMailUser

`func (o *V0044JobInfo) GetMailUser() string`

GetMailUser returns the MailUser field if non-nil, zero value otherwise.

### GetMailUserOk

`func (o *V0044JobInfo) GetMailUserOk() (*string, bool)`

GetMailUserOk returns a tuple with the MailUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailUser

`func (o *V0044JobInfo) SetMailUser(v string)`

SetMailUser sets MailUser field to given value.

### HasMailUser

`func (o *V0044JobInfo) HasMailUser() bool`

HasMailUser returns a boolean if a field has been set.

### GetMaxCpus

`func (o *V0044JobInfo) GetMaxCpus() V0044Uint32NoValStruct`

GetMaxCpus returns the MaxCpus field if non-nil, zero value otherwise.

### GetMaxCpusOk

`func (o *V0044JobInfo) GetMaxCpusOk() (*V0044Uint32NoValStruct, bool)`

GetMaxCpusOk returns a tuple with the MaxCpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCpus

`func (o *V0044JobInfo) SetMaxCpus(v V0044Uint32NoValStruct)`

SetMaxCpus sets MaxCpus field to given value.

### HasMaxCpus

`func (o *V0044JobInfo) HasMaxCpus() bool`

HasMaxCpus returns a boolean if a field has been set.

### GetMaxNodes

`func (o *V0044JobInfo) GetMaxNodes() V0044Uint32NoValStruct`

GetMaxNodes returns the MaxNodes field if non-nil, zero value otherwise.

### GetMaxNodesOk

`func (o *V0044JobInfo) GetMaxNodesOk() (*V0044Uint32NoValStruct, bool)`

GetMaxNodesOk returns a tuple with the MaxNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNodes

`func (o *V0044JobInfo) SetMaxNodes(v V0044Uint32NoValStruct)`

SetMaxNodes sets MaxNodes field to given value.

### HasMaxNodes

`func (o *V0044JobInfo) HasMaxNodes() bool`

HasMaxNodes returns a boolean if a field has been set.

### GetMcsLabel

`func (o *V0044JobInfo) GetMcsLabel() string`

GetMcsLabel returns the McsLabel field if non-nil, zero value otherwise.

### GetMcsLabelOk

`func (o *V0044JobInfo) GetMcsLabelOk() (*string, bool)`

GetMcsLabelOk returns a tuple with the McsLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcsLabel

`func (o *V0044JobInfo) SetMcsLabel(v string)`

SetMcsLabel sets McsLabel field to given value.

### HasMcsLabel

`func (o *V0044JobInfo) HasMcsLabel() bool`

HasMcsLabel returns a boolean if a field has been set.

### GetMemoryPerTres

`func (o *V0044JobInfo) GetMemoryPerTres() string`

GetMemoryPerTres returns the MemoryPerTres field if non-nil, zero value otherwise.

### GetMemoryPerTresOk

`func (o *V0044JobInfo) GetMemoryPerTresOk() (*string, bool)`

GetMemoryPerTresOk returns a tuple with the MemoryPerTres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryPerTres

`func (o *V0044JobInfo) SetMemoryPerTres(v string)`

SetMemoryPerTres sets MemoryPerTres field to given value.

### HasMemoryPerTres

`func (o *V0044JobInfo) HasMemoryPerTres() bool`

HasMemoryPerTres returns a boolean if a field has been set.

### GetName

`func (o *V0044JobInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V0044JobInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V0044JobInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V0044JobInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetwork

`func (o *V0044JobInfo) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *V0044JobInfo) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *V0044JobInfo) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *V0044JobInfo) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetNodes

`func (o *V0044JobInfo) GetNodes() string`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *V0044JobInfo) GetNodesOk() (*string, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *V0044JobInfo) SetNodes(v string)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *V0044JobInfo) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetNice

`func (o *V0044JobInfo) GetNice() int32`

GetNice returns the Nice field if non-nil, zero value otherwise.

### GetNiceOk

`func (o *V0044JobInfo) GetNiceOk() (*int32, bool)`

GetNiceOk returns a tuple with the Nice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNice

`func (o *V0044JobInfo) SetNice(v int32)`

SetNice sets Nice field to given value.

### HasNice

`func (o *V0044JobInfo) HasNice() bool`

HasNice returns a boolean if a field has been set.

### GetTasksPerCore

`func (o *V0044JobInfo) GetTasksPerCore() V0044Uint16NoValStruct`

GetTasksPerCore returns the TasksPerCore field if non-nil, zero value otherwise.

### GetTasksPerCoreOk

`func (o *V0044JobInfo) GetTasksPerCoreOk() (*V0044Uint16NoValStruct, bool)`

GetTasksPerCoreOk returns a tuple with the TasksPerCore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasksPerCore

`func (o *V0044JobInfo) SetTasksPerCore(v V0044Uint16NoValStruct)`

SetTasksPerCore sets TasksPerCore field to given value.

### HasTasksPerCore

`func (o *V0044JobInfo) HasTasksPerCore() bool`

HasTasksPerCore returns a boolean if a field has been set.

### GetTasksPerTres

`func (o *V0044JobInfo) GetTasksPerTres() V0044Uint16NoValStruct`

GetTasksPerTres returns the TasksPerTres field if non-nil, zero value otherwise.

### GetTasksPerTresOk

`func (o *V0044JobInfo) GetTasksPerTresOk() (*V0044Uint16NoValStruct, bool)`

GetTasksPerTresOk returns a tuple with the TasksPerTres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasksPerTres

`func (o *V0044JobInfo) SetTasksPerTres(v V0044Uint16NoValStruct)`

SetTasksPerTres sets TasksPerTres field to given value.

### HasTasksPerTres

`func (o *V0044JobInfo) HasTasksPerTres() bool`

HasTasksPerTres returns a boolean if a field has been set.

### GetTasksPerNode

`func (o *V0044JobInfo) GetTasksPerNode() V0044Uint16NoValStruct`

GetTasksPerNode returns the TasksPerNode field if non-nil, zero value otherwise.

### GetTasksPerNodeOk

`func (o *V0044JobInfo) GetTasksPerNodeOk() (*V0044Uint16NoValStruct, bool)`

GetTasksPerNodeOk returns a tuple with the TasksPerNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasksPerNode

`func (o *V0044JobInfo) SetTasksPerNode(v V0044Uint16NoValStruct)`

SetTasksPerNode sets TasksPerNode field to given value.

### HasTasksPerNode

`func (o *V0044JobInfo) HasTasksPerNode() bool`

HasTasksPerNode returns a boolean if a field has been set.

### GetTasksPerSocket

`func (o *V0044JobInfo) GetTasksPerSocket() V0044Uint16NoValStruct`

GetTasksPerSocket returns the TasksPerSocket field if non-nil, zero value otherwise.

### GetTasksPerSocketOk

`func (o *V0044JobInfo) GetTasksPerSocketOk() (*V0044Uint16NoValStruct, bool)`

GetTasksPerSocketOk returns a tuple with the TasksPerSocket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasksPerSocket

`func (o *V0044JobInfo) SetTasksPerSocket(v V0044Uint16NoValStruct)`

SetTasksPerSocket sets TasksPerSocket field to given value.

### HasTasksPerSocket

`func (o *V0044JobInfo) HasTasksPerSocket() bool`

HasTasksPerSocket returns a boolean if a field has been set.

### GetTasksPerBoard

`func (o *V0044JobInfo) GetTasksPerBoard() V0044Uint16NoValStruct`

GetTasksPerBoard returns the TasksPerBoard field if non-nil, zero value otherwise.

### GetTasksPerBoardOk

`func (o *V0044JobInfo) GetTasksPerBoardOk() (*V0044Uint16NoValStruct, bool)`

GetTasksPerBoardOk returns a tuple with the TasksPerBoard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasksPerBoard

`func (o *V0044JobInfo) SetTasksPerBoard(v V0044Uint16NoValStruct)`

SetTasksPerBoard sets TasksPerBoard field to given value.

### HasTasksPerBoard

`func (o *V0044JobInfo) HasTasksPerBoard() bool`

HasTasksPerBoard returns a boolean if a field has been set.

### GetCpus

`func (o *V0044JobInfo) GetCpus() V0044Uint32NoValStruct`

GetCpus returns the Cpus field if non-nil, zero value otherwise.

### GetCpusOk

`func (o *V0044JobInfo) GetCpusOk() (*V0044Uint32NoValStruct, bool)`

GetCpusOk returns a tuple with the Cpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpus

`func (o *V0044JobInfo) SetCpus(v V0044Uint32NoValStruct)`

SetCpus sets Cpus field to given value.

### HasCpus

`func (o *V0044JobInfo) HasCpus() bool`

HasCpus returns a boolean if a field has been set.

### GetNodeCount

`func (o *V0044JobInfo) GetNodeCount() V0044Uint32NoValStruct`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *V0044JobInfo) GetNodeCountOk() (*V0044Uint32NoValStruct, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *V0044JobInfo) SetNodeCount(v V0044Uint32NoValStruct)`

SetNodeCount sets NodeCount field to given value.

### HasNodeCount

`func (o *V0044JobInfo) HasNodeCount() bool`

HasNodeCount returns a boolean if a field has been set.

### GetTasks

`func (o *V0044JobInfo) GetTasks() V0044Uint32NoValStruct`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *V0044JobInfo) GetTasksOk() (*V0044Uint32NoValStruct, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *V0044JobInfo) SetTasks(v V0044Uint32NoValStruct)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *V0044JobInfo) HasTasks() bool`

HasTasks returns a boolean if a field has been set.

### GetPartition

`func (o *V0044JobInfo) GetPartition() string`

GetPartition returns the Partition field if non-nil, zero value otherwise.

### GetPartitionOk

`func (o *V0044JobInfo) GetPartitionOk() (*string, bool)`

GetPartitionOk returns a tuple with the Partition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartition

`func (o *V0044JobInfo) SetPartition(v string)`

SetPartition sets Partition field to given value.

### HasPartition

`func (o *V0044JobInfo) HasPartition() bool`

HasPartition returns a boolean if a field has been set.

### GetPrefer

`func (o *V0044JobInfo) GetPrefer() string`

GetPrefer returns the Prefer field if non-nil, zero value otherwise.

### GetPreferOk

`func (o *V0044JobInfo) GetPreferOk() (*string, bool)`

GetPreferOk returns a tuple with the Prefer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefer

`func (o *V0044JobInfo) SetPrefer(v string)`

SetPrefer sets Prefer field to given value.

### HasPrefer

`func (o *V0044JobInfo) HasPrefer() bool`

HasPrefer returns a boolean if a field has been set.

### GetMemoryPerCpu

`func (o *V0044JobInfo) GetMemoryPerCpu() V0044Uint64NoValStruct`

GetMemoryPerCpu returns the MemoryPerCpu field if non-nil, zero value otherwise.

### GetMemoryPerCpuOk

`func (o *V0044JobInfo) GetMemoryPerCpuOk() (*V0044Uint64NoValStruct, bool)`

GetMemoryPerCpuOk returns a tuple with the MemoryPerCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryPerCpu

`func (o *V0044JobInfo) SetMemoryPerCpu(v V0044Uint64NoValStruct)`

SetMemoryPerCpu sets MemoryPerCpu field to given value.

### HasMemoryPerCpu

`func (o *V0044JobInfo) HasMemoryPerCpu() bool`

HasMemoryPerCpu returns a boolean if a field has been set.

### GetMemoryPerNode

`func (o *V0044JobInfo) GetMemoryPerNode() V0044Uint64NoValStruct`

GetMemoryPerNode returns the MemoryPerNode field if non-nil, zero value otherwise.

### GetMemoryPerNodeOk

`func (o *V0044JobInfo) GetMemoryPerNodeOk() (*V0044Uint64NoValStruct, bool)`

GetMemoryPerNodeOk returns a tuple with the MemoryPerNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryPerNode

`func (o *V0044JobInfo) SetMemoryPerNode(v V0044Uint64NoValStruct)`

SetMemoryPerNode sets MemoryPerNode field to given value.

### HasMemoryPerNode

`func (o *V0044JobInfo) HasMemoryPerNode() bool`

HasMemoryPerNode returns a boolean if a field has been set.

### GetMinimumCpusPerNode

`func (o *V0044JobInfo) GetMinimumCpusPerNode() V0044Uint16NoValStruct`

GetMinimumCpusPerNode returns the MinimumCpusPerNode field if non-nil, zero value otherwise.

### GetMinimumCpusPerNodeOk

`func (o *V0044JobInfo) GetMinimumCpusPerNodeOk() (*V0044Uint16NoValStruct, bool)`

GetMinimumCpusPerNodeOk returns a tuple with the MinimumCpusPerNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumCpusPerNode

`func (o *V0044JobInfo) SetMinimumCpusPerNode(v V0044Uint16NoValStruct)`

SetMinimumCpusPerNode sets MinimumCpusPerNode field to given value.

### HasMinimumCpusPerNode

`func (o *V0044JobInfo) HasMinimumCpusPerNode() bool`

HasMinimumCpusPerNode returns a boolean if a field has been set.

### GetMinimumTmpDiskPerNode

`func (o *V0044JobInfo) GetMinimumTmpDiskPerNode() V0044Uint32NoValStruct`

GetMinimumTmpDiskPerNode returns the MinimumTmpDiskPerNode field if non-nil, zero value otherwise.

### GetMinimumTmpDiskPerNodeOk

`func (o *V0044JobInfo) GetMinimumTmpDiskPerNodeOk() (*V0044Uint32NoValStruct, bool)`

GetMinimumTmpDiskPerNodeOk returns a tuple with the MinimumTmpDiskPerNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumTmpDiskPerNode

`func (o *V0044JobInfo) SetMinimumTmpDiskPerNode(v V0044Uint32NoValStruct)`

SetMinimumTmpDiskPerNode sets MinimumTmpDiskPerNode field to given value.

### HasMinimumTmpDiskPerNode

`func (o *V0044JobInfo) HasMinimumTmpDiskPerNode() bool`

HasMinimumTmpDiskPerNode returns a boolean if a field has been set.

### GetPower

`func (o *V0044JobInfo) GetPower() V0044JobInfoPower`

GetPower returns the Power field if non-nil, zero value otherwise.

### GetPowerOk

`func (o *V0044JobInfo) GetPowerOk() (*V0044JobInfoPower, bool)`

GetPowerOk returns a tuple with the Power field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPower

`func (o *V0044JobInfo) SetPower(v V0044JobInfoPower)`

SetPower sets Power field to given value.

### HasPower

`func (o *V0044JobInfo) HasPower() bool`

HasPower returns a boolean if a field has been set.

### GetPreemptTime

`func (o *V0044JobInfo) GetPreemptTime() V0044Uint64NoValStruct`

GetPreemptTime returns the PreemptTime field if non-nil, zero value otherwise.

### GetPreemptTimeOk

`func (o *V0044JobInfo) GetPreemptTimeOk() (*V0044Uint64NoValStruct, bool)`

GetPreemptTimeOk returns a tuple with the PreemptTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreemptTime

`func (o *V0044JobInfo) SetPreemptTime(v V0044Uint64NoValStruct)`

SetPreemptTime sets PreemptTime field to given value.

### HasPreemptTime

`func (o *V0044JobInfo) HasPreemptTime() bool`

HasPreemptTime returns a boolean if a field has been set.

### GetPreemptableTime

`func (o *V0044JobInfo) GetPreemptableTime() V0044Uint64NoValStruct`

GetPreemptableTime returns the PreemptableTime field if non-nil, zero value otherwise.

### GetPreemptableTimeOk

`func (o *V0044JobInfo) GetPreemptableTimeOk() (*V0044Uint64NoValStruct, bool)`

GetPreemptableTimeOk returns a tuple with the PreemptableTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreemptableTime

`func (o *V0044JobInfo) SetPreemptableTime(v V0044Uint64NoValStruct)`

SetPreemptableTime sets PreemptableTime field to given value.

### HasPreemptableTime

`func (o *V0044JobInfo) HasPreemptableTime() bool`

HasPreemptableTime returns a boolean if a field has been set.

### GetPreSusTime

`func (o *V0044JobInfo) GetPreSusTime() V0044Uint64NoValStruct`

GetPreSusTime returns the PreSusTime field if non-nil, zero value otherwise.

### GetPreSusTimeOk

`func (o *V0044JobInfo) GetPreSusTimeOk() (*V0044Uint64NoValStruct, bool)`

GetPreSusTimeOk returns a tuple with the PreSusTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreSusTime

`func (o *V0044JobInfo) SetPreSusTime(v V0044Uint64NoValStruct)`

SetPreSusTime sets PreSusTime field to given value.

### HasPreSusTime

`func (o *V0044JobInfo) HasPreSusTime() bool`

HasPreSusTime returns a boolean if a field has been set.

### GetHold

`func (o *V0044JobInfo) GetHold() bool`

GetHold returns the Hold field if non-nil, zero value otherwise.

### GetHoldOk

`func (o *V0044JobInfo) GetHoldOk() (*bool, bool)`

GetHoldOk returns a tuple with the Hold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHold

`func (o *V0044JobInfo) SetHold(v bool)`

SetHold sets Hold field to given value.

### HasHold

`func (o *V0044JobInfo) HasHold() bool`

HasHold returns a boolean if a field has been set.

### GetPriority

`func (o *V0044JobInfo) GetPriority() V0044Uint32NoValStruct`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *V0044JobInfo) GetPriorityOk() (*V0044Uint32NoValStruct, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *V0044JobInfo) SetPriority(v V0044Uint32NoValStruct)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *V0044JobInfo) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetPriorityByPartition

`func (o *V0044JobInfo) GetPriorityByPartition() []V0044PartPrio`

GetPriorityByPartition returns the PriorityByPartition field if non-nil, zero value otherwise.

### GetPriorityByPartitionOk

`func (o *V0044JobInfo) GetPriorityByPartitionOk() (*[]V0044PartPrio, bool)`

GetPriorityByPartitionOk returns a tuple with the PriorityByPartition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorityByPartition

`func (o *V0044JobInfo) SetPriorityByPartition(v []V0044PartPrio)`

SetPriorityByPartition sets PriorityByPartition field to given value.

### HasPriorityByPartition

`func (o *V0044JobInfo) HasPriorityByPartition() bool`

HasPriorityByPartition returns a boolean if a field has been set.

### GetProfile

`func (o *V0044JobInfo) GetProfile() []string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *V0044JobInfo) GetProfileOk() (*[]string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *V0044JobInfo) SetProfile(v []string)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *V0044JobInfo) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetQos

`func (o *V0044JobInfo) GetQos() string`

GetQos returns the Qos field if non-nil, zero value otherwise.

### GetQosOk

`func (o *V0044JobInfo) GetQosOk() (*string, bool)`

GetQosOk returns a tuple with the Qos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQos

`func (o *V0044JobInfo) SetQos(v string)`

SetQos sets Qos field to given value.

### HasQos

`func (o *V0044JobInfo) HasQos() bool`

HasQos returns a boolean if a field has been set.

### GetReboot

`func (o *V0044JobInfo) GetReboot() bool`

GetReboot returns the Reboot field if non-nil, zero value otherwise.

### GetRebootOk

`func (o *V0044JobInfo) GetRebootOk() (*bool, bool)`

GetRebootOk returns a tuple with the Reboot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReboot

`func (o *V0044JobInfo) SetReboot(v bool)`

SetReboot sets Reboot field to given value.

### HasReboot

`func (o *V0044JobInfo) HasReboot() bool`

HasReboot returns a boolean if a field has been set.

### GetRequiredNodes

`func (o *V0044JobInfo) GetRequiredNodes() string`

GetRequiredNodes returns the RequiredNodes field if non-nil, zero value otherwise.

### GetRequiredNodesOk

`func (o *V0044JobInfo) GetRequiredNodesOk() (*string, bool)`

GetRequiredNodesOk returns a tuple with the RequiredNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredNodes

`func (o *V0044JobInfo) SetRequiredNodes(v string)`

SetRequiredNodes sets RequiredNodes field to given value.

### HasRequiredNodes

`func (o *V0044JobInfo) HasRequiredNodes() bool`

HasRequiredNodes returns a boolean if a field has been set.

### GetRequiredSwitches

`func (o *V0044JobInfo) GetRequiredSwitches() int32`

GetRequiredSwitches returns the RequiredSwitches field if non-nil, zero value otherwise.

### GetRequiredSwitchesOk

`func (o *V0044JobInfo) GetRequiredSwitchesOk() (*int32, bool)`

GetRequiredSwitchesOk returns a tuple with the RequiredSwitches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredSwitches

`func (o *V0044JobInfo) SetRequiredSwitches(v int32)`

SetRequiredSwitches sets RequiredSwitches field to given value.

### HasRequiredSwitches

`func (o *V0044JobInfo) HasRequiredSwitches() bool`

HasRequiredSwitches returns a boolean if a field has been set.

### GetRequeue

`func (o *V0044JobInfo) GetRequeue() bool`

GetRequeue returns the Requeue field if non-nil, zero value otherwise.

### GetRequeueOk

`func (o *V0044JobInfo) GetRequeueOk() (*bool, bool)`

GetRequeueOk returns a tuple with the Requeue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequeue

`func (o *V0044JobInfo) SetRequeue(v bool)`

SetRequeue sets Requeue field to given value.

### HasRequeue

`func (o *V0044JobInfo) HasRequeue() bool`

HasRequeue returns a boolean if a field has been set.

### GetResizeTime

`func (o *V0044JobInfo) GetResizeTime() V0044Uint64NoValStruct`

GetResizeTime returns the ResizeTime field if non-nil, zero value otherwise.

### GetResizeTimeOk

`func (o *V0044JobInfo) GetResizeTimeOk() (*V0044Uint64NoValStruct, bool)`

GetResizeTimeOk returns a tuple with the ResizeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResizeTime

`func (o *V0044JobInfo) SetResizeTime(v V0044Uint64NoValStruct)`

SetResizeTime sets ResizeTime field to given value.

### HasResizeTime

`func (o *V0044JobInfo) HasResizeTime() bool`

HasResizeTime returns a boolean if a field has been set.

### GetRestartCnt

`func (o *V0044JobInfo) GetRestartCnt() int32`

GetRestartCnt returns the RestartCnt field if non-nil, zero value otherwise.

### GetRestartCntOk

`func (o *V0044JobInfo) GetRestartCntOk() (*int32, bool)`

GetRestartCntOk returns a tuple with the RestartCnt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartCnt

`func (o *V0044JobInfo) SetRestartCnt(v int32)`

SetRestartCnt sets RestartCnt field to given value.

### HasRestartCnt

`func (o *V0044JobInfo) HasRestartCnt() bool`

HasRestartCnt returns a boolean if a field has been set.

### GetResvName

`func (o *V0044JobInfo) GetResvName() string`

GetResvName returns the ResvName field if non-nil, zero value otherwise.

### GetResvNameOk

`func (o *V0044JobInfo) GetResvNameOk() (*string, bool)`

GetResvNameOk returns a tuple with the ResvName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResvName

`func (o *V0044JobInfo) SetResvName(v string)`

SetResvName sets ResvName field to given value.

### HasResvName

`func (o *V0044JobInfo) HasResvName() bool`

HasResvName returns a boolean if a field has been set.

### GetScheduledNodes

`func (o *V0044JobInfo) GetScheduledNodes() string`

GetScheduledNodes returns the ScheduledNodes field if non-nil, zero value otherwise.

### GetScheduledNodesOk

`func (o *V0044JobInfo) GetScheduledNodesOk() (*string, bool)`

GetScheduledNodesOk returns a tuple with the ScheduledNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledNodes

`func (o *V0044JobInfo) SetScheduledNodes(v string)`

SetScheduledNodes sets ScheduledNodes field to given value.

### HasScheduledNodes

`func (o *V0044JobInfo) HasScheduledNodes() bool`

HasScheduledNodes returns a boolean if a field has been set.

### GetSegmentSize

`func (o *V0044JobInfo) GetSegmentSize() int32`

GetSegmentSize returns the SegmentSize field if non-nil, zero value otherwise.

### GetSegmentSizeOk

`func (o *V0044JobInfo) GetSegmentSizeOk() (*int32, bool)`

GetSegmentSizeOk returns a tuple with the SegmentSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentSize

`func (o *V0044JobInfo) SetSegmentSize(v int32)`

SetSegmentSize sets SegmentSize field to given value.

### HasSegmentSize

`func (o *V0044JobInfo) HasSegmentSize() bool`

HasSegmentSize returns a boolean if a field has been set.

### GetSelinuxContext

`func (o *V0044JobInfo) GetSelinuxContext() string`

GetSelinuxContext returns the SelinuxContext field if non-nil, zero value otherwise.

### GetSelinuxContextOk

`func (o *V0044JobInfo) GetSelinuxContextOk() (*string, bool)`

GetSelinuxContextOk returns a tuple with the SelinuxContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelinuxContext

`func (o *V0044JobInfo) SetSelinuxContext(v string)`

SetSelinuxContext sets SelinuxContext field to given value.

### HasSelinuxContext

`func (o *V0044JobInfo) HasSelinuxContext() bool`

HasSelinuxContext returns a boolean if a field has been set.

### GetShared

`func (o *V0044JobInfo) GetShared() []string`

GetShared returns the Shared field if non-nil, zero value otherwise.

### GetSharedOk

`func (o *V0044JobInfo) GetSharedOk() (*[]string, bool)`

GetSharedOk returns a tuple with the Shared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShared

`func (o *V0044JobInfo) SetShared(v []string)`

SetShared sets Shared field to given value.

### HasShared

`func (o *V0044JobInfo) HasShared() bool`

HasShared returns a boolean if a field has been set.

### GetStepId

`func (o *V0044JobInfo) GetStepId() V0044SlurmStepId`

GetStepId returns the StepId field if non-nil, zero value otherwise.

### GetStepIdOk

`func (o *V0044JobInfo) GetStepIdOk() (*V0044SlurmStepId, bool)`

GetStepIdOk returns a tuple with the StepId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepId

`func (o *V0044JobInfo) SetStepId(v V0044SlurmStepId)`

SetStepId sets StepId field to given value.

### HasStepId

`func (o *V0044JobInfo) HasStepId() bool`

HasStepId returns a boolean if a field has been set.

### GetSocketsPerBoard

`func (o *V0044JobInfo) GetSocketsPerBoard() int32`

GetSocketsPerBoard returns the SocketsPerBoard field if non-nil, zero value otherwise.

### GetSocketsPerBoardOk

`func (o *V0044JobInfo) GetSocketsPerBoardOk() (*int32, bool)`

GetSocketsPerBoardOk returns a tuple with the SocketsPerBoard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocketsPerBoard

`func (o *V0044JobInfo) SetSocketsPerBoard(v int32)`

SetSocketsPerBoard sets SocketsPerBoard field to given value.

### HasSocketsPerBoard

`func (o *V0044JobInfo) HasSocketsPerBoard() bool`

HasSocketsPerBoard returns a boolean if a field has been set.

### GetSocketsPerNode

`func (o *V0044JobInfo) GetSocketsPerNode() V0044Uint16NoValStruct`

GetSocketsPerNode returns the SocketsPerNode field if non-nil, zero value otherwise.

### GetSocketsPerNodeOk

`func (o *V0044JobInfo) GetSocketsPerNodeOk() (*V0044Uint16NoValStruct, bool)`

GetSocketsPerNodeOk returns a tuple with the SocketsPerNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocketsPerNode

`func (o *V0044JobInfo) SetSocketsPerNode(v V0044Uint16NoValStruct)`

SetSocketsPerNode sets SocketsPerNode field to given value.

### HasSocketsPerNode

`func (o *V0044JobInfo) HasSocketsPerNode() bool`

HasSocketsPerNode returns a boolean if a field has been set.

### GetStartTime

`func (o *V0044JobInfo) GetStartTime() V0044Uint64NoValStruct`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *V0044JobInfo) GetStartTimeOk() (*V0044Uint64NoValStruct, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *V0044JobInfo) SetStartTime(v V0044Uint64NoValStruct)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *V0044JobInfo) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStateDescription

`func (o *V0044JobInfo) GetStateDescription() string`

GetStateDescription returns the StateDescription field if non-nil, zero value otherwise.

### GetStateDescriptionOk

`func (o *V0044JobInfo) GetStateDescriptionOk() (*string, bool)`

GetStateDescriptionOk returns a tuple with the StateDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateDescription

`func (o *V0044JobInfo) SetStateDescription(v string)`

SetStateDescription sets StateDescription field to given value.

### HasStateDescription

`func (o *V0044JobInfo) HasStateDescription() bool`

HasStateDescription returns a boolean if a field has been set.

### GetStateReason

`func (o *V0044JobInfo) GetStateReason() string`

GetStateReason returns the StateReason field if non-nil, zero value otherwise.

### GetStateReasonOk

`func (o *V0044JobInfo) GetStateReasonOk() (*string, bool)`

GetStateReasonOk returns a tuple with the StateReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateReason

`func (o *V0044JobInfo) SetStateReason(v string)`

SetStateReason sets StateReason field to given value.

### HasStateReason

`func (o *V0044JobInfo) HasStateReason() bool`

HasStateReason returns a boolean if a field has been set.

### GetStandardInput

`func (o *V0044JobInfo) GetStandardInput() string`

GetStandardInput returns the StandardInput field if non-nil, zero value otherwise.

### GetStandardInputOk

`func (o *V0044JobInfo) GetStandardInputOk() (*string, bool)`

GetStandardInputOk returns a tuple with the StandardInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardInput

`func (o *V0044JobInfo) SetStandardInput(v string)`

SetStandardInput sets StandardInput field to given value.

### HasStandardInput

`func (o *V0044JobInfo) HasStandardInput() bool`

HasStandardInput returns a boolean if a field has been set.

### GetStandardOutput

`func (o *V0044JobInfo) GetStandardOutput() string`

GetStandardOutput returns the StandardOutput field if non-nil, zero value otherwise.

### GetStandardOutputOk

`func (o *V0044JobInfo) GetStandardOutputOk() (*string, bool)`

GetStandardOutputOk returns a tuple with the StandardOutput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardOutput

`func (o *V0044JobInfo) SetStandardOutput(v string)`

SetStandardOutput sets StandardOutput field to given value.

### HasStandardOutput

`func (o *V0044JobInfo) HasStandardOutput() bool`

HasStandardOutput returns a boolean if a field has been set.

### GetStandardError

`func (o *V0044JobInfo) GetStandardError() string`

GetStandardError returns the StandardError field if non-nil, zero value otherwise.

### GetStandardErrorOk

`func (o *V0044JobInfo) GetStandardErrorOk() (*string, bool)`

GetStandardErrorOk returns a tuple with the StandardError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardError

`func (o *V0044JobInfo) SetStandardError(v string)`

SetStandardError sets StandardError field to given value.

### HasStandardError

`func (o *V0044JobInfo) HasStandardError() bool`

HasStandardError returns a boolean if a field has been set.

### GetStdinExpanded

`func (o *V0044JobInfo) GetStdinExpanded() string`

GetStdinExpanded returns the StdinExpanded field if non-nil, zero value otherwise.

### GetStdinExpandedOk

`func (o *V0044JobInfo) GetStdinExpandedOk() (*string, bool)`

GetStdinExpandedOk returns a tuple with the StdinExpanded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStdinExpanded

`func (o *V0044JobInfo) SetStdinExpanded(v string)`

SetStdinExpanded sets StdinExpanded field to given value.

### HasStdinExpanded

`func (o *V0044JobInfo) HasStdinExpanded() bool`

HasStdinExpanded returns a boolean if a field has been set.

### GetStdoutExpanded

`func (o *V0044JobInfo) GetStdoutExpanded() string`

GetStdoutExpanded returns the StdoutExpanded field if non-nil, zero value otherwise.

### GetStdoutExpandedOk

`func (o *V0044JobInfo) GetStdoutExpandedOk() (*string, bool)`

GetStdoutExpandedOk returns a tuple with the StdoutExpanded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStdoutExpanded

`func (o *V0044JobInfo) SetStdoutExpanded(v string)`

SetStdoutExpanded sets StdoutExpanded field to given value.

### HasStdoutExpanded

`func (o *V0044JobInfo) HasStdoutExpanded() bool`

HasStdoutExpanded returns a boolean if a field has been set.

### GetStderrExpanded

`func (o *V0044JobInfo) GetStderrExpanded() string`

GetStderrExpanded returns the StderrExpanded field if non-nil, zero value otherwise.

### GetStderrExpandedOk

`func (o *V0044JobInfo) GetStderrExpandedOk() (*string, bool)`

GetStderrExpandedOk returns a tuple with the StderrExpanded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStderrExpanded

`func (o *V0044JobInfo) SetStderrExpanded(v string)`

SetStderrExpanded sets StderrExpanded field to given value.

### HasStderrExpanded

`func (o *V0044JobInfo) HasStderrExpanded() bool`

HasStderrExpanded returns a boolean if a field has been set.

### GetSubmitTime

`func (o *V0044JobInfo) GetSubmitTime() V0044Uint64NoValStruct`

GetSubmitTime returns the SubmitTime field if non-nil, zero value otherwise.

### GetSubmitTimeOk

`func (o *V0044JobInfo) GetSubmitTimeOk() (*V0044Uint64NoValStruct, bool)`

GetSubmitTimeOk returns a tuple with the SubmitTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmitTime

`func (o *V0044JobInfo) SetSubmitTime(v V0044Uint64NoValStruct)`

SetSubmitTime sets SubmitTime field to given value.

### HasSubmitTime

`func (o *V0044JobInfo) HasSubmitTime() bool`

HasSubmitTime returns a boolean if a field has been set.

### GetSubmitLine

`func (o *V0044JobInfo) GetSubmitLine() string`

GetSubmitLine returns the SubmitLine field if non-nil, zero value otherwise.

### GetSubmitLineOk

`func (o *V0044JobInfo) GetSubmitLineOk() (*string, bool)`

GetSubmitLineOk returns a tuple with the SubmitLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmitLine

`func (o *V0044JobInfo) SetSubmitLine(v string)`

SetSubmitLine sets SubmitLine field to given value.

### HasSubmitLine

`func (o *V0044JobInfo) HasSubmitLine() bool`

HasSubmitLine returns a boolean if a field has been set.

### GetSuspendTime

`func (o *V0044JobInfo) GetSuspendTime() V0044Uint64NoValStruct`

GetSuspendTime returns the SuspendTime field if non-nil, zero value otherwise.

### GetSuspendTimeOk

`func (o *V0044JobInfo) GetSuspendTimeOk() (*V0044Uint64NoValStruct, bool)`

GetSuspendTimeOk returns a tuple with the SuspendTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspendTime

`func (o *V0044JobInfo) SetSuspendTime(v V0044Uint64NoValStruct)`

SetSuspendTime sets SuspendTime field to given value.

### HasSuspendTime

`func (o *V0044JobInfo) HasSuspendTime() bool`

HasSuspendTime returns a boolean if a field has been set.

### GetSystemComment

`func (o *V0044JobInfo) GetSystemComment() string`

GetSystemComment returns the SystemComment field if non-nil, zero value otherwise.

### GetSystemCommentOk

`func (o *V0044JobInfo) GetSystemCommentOk() (*string, bool)`

GetSystemCommentOk returns a tuple with the SystemComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemComment

`func (o *V0044JobInfo) SetSystemComment(v string)`

SetSystemComment sets SystemComment field to given value.

### HasSystemComment

`func (o *V0044JobInfo) HasSystemComment() bool`

HasSystemComment returns a boolean if a field has been set.

### GetTimeLimit

`func (o *V0044JobInfo) GetTimeLimit() V0044Uint32NoValStruct`

GetTimeLimit returns the TimeLimit field if non-nil, zero value otherwise.

### GetTimeLimitOk

`func (o *V0044JobInfo) GetTimeLimitOk() (*V0044Uint32NoValStruct, bool)`

GetTimeLimitOk returns a tuple with the TimeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeLimit

`func (o *V0044JobInfo) SetTimeLimit(v V0044Uint32NoValStruct)`

SetTimeLimit sets TimeLimit field to given value.

### HasTimeLimit

`func (o *V0044JobInfo) HasTimeLimit() bool`

HasTimeLimit returns a boolean if a field has been set.

### GetTimeMinimum

`func (o *V0044JobInfo) GetTimeMinimum() V0044Uint32NoValStruct`

GetTimeMinimum returns the TimeMinimum field if non-nil, zero value otherwise.

### GetTimeMinimumOk

`func (o *V0044JobInfo) GetTimeMinimumOk() (*V0044Uint32NoValStruct, bool)`

GetTimeMinimumOk returns a tuple with the TimeMinimum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeMinimum

`func (o *V0044JobInfo) SetTimeMinimum(v V0044Uint32NoValStruct)`

SetTimeMinimum sets TimeMinimum field to given value.

### HasTimeMinimum

`func (o *V0044JobInfo) HasTimeMinimum() bool`

HasTimeMinimum returns a boolean if a field has been set.

### GetThreadsPerCore

`func (o *V0044JobInfo) GetThreadsPerCore() V0044Uint16NoValStruct`

GetThreadsPerCore returns the ThreadsPerCore field if non-nil, zero value otherwise.

### GetThreadsPerCoreOk

`func (o *V0044JobInfo) GetThreadsPerCoreOk() (*V0044Uint16NoValStruct, bool)`

GetThreadsPerCoreOk returns a tuple with the ThreadsPerCore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadsPerCore

`func (o *V0044JobInfo) SetThreadsPerCore(v V0044Uint16NoValStruct)`

SetThreadsPerCore sets ThreadsPerCore field to given value.

### HasThreadsPerCore

`func (o *V0044JobInfo) HasThreadsPerCore() bool`

HasThreadsPerCore returns a boolean if a field has been set.

### GetTresBind

`func (o *V0044JobInfo) GetTresBind() string`

GetTresBind returns the TresBind field if non-nil, zero value otherwise.

### GetTresBindOk

`func (o *V0044JobInfo) GetTresBindOk() (*string, bool)`

GetTresBindOk returns a tuple with the TresBind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTresBind

`func (o *V0044JobInfo) SetTresBind(v string)`

SetTresBind sets TresBind field to given value.

### HasTresBind

`func (o *V0044JobInfo) HasTresBind() bool`

HasTresBind returns a boolean if a field has been set.

### GetTresFreq

`func (o *V0044JobInfo) GetTresFreq() string`

GetTresFreq returns the TresFreq field if non-nil, zero value otherwise.

### GetTresFreqOk

`func (o *V0044JobInfo) GetTresFreqOk() (*string, bool)`

GetTresFreqOk returns a tuple with the TresFreq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTresFreq

`func (o *V0044JobInfo) SetTresFreq(v string)`

SetTresFreq sets TresFreq field to given value.

### HasTresFreq

`func (o *V0044JobInfo) HasTresFreq() bool`

HasTresFreq returns a boolean if a field has been set.

### GetTresPerJob

`func (o *V0044JobInfo) GetTresPerJob() string`

GetTresPerJob returns the TresPerJob field if non-nil, zero value otherwise.

### GetTresPerJobOk

`func (o *V0044JobInfo) GetTresPerJobOk() (*string, bool)`

GetTresPerJobOk returns a tuple with the TresPerJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTresPerJob

`func (o *V0044JobInfo) SetTresPerJob(v string)`

SetTresPerJob sets TresPerJob field to given value.

### HasTresPerJob

`func (o *V0044JobInfo) HasTresPerJob() bool`

HasTresPerJob returns a boolean if a field has been set.

### GetTresPerNode

`func (o *V0044JobInfo) GetTresPerNode() string`

GetTresPerNode returns the TresPerNode field if non-nil, zero value otherwise.

### GetTresPerNodeOk

`func (o *V0044JobInfo) GetTresPerNodeOk() (*string, bool)`

GetTresPerNodeOk returns a tuple with the TresPerNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTresPerNode

`func (o *V0044JobInfo) SetTresPerNode(v string)`

SetTresPerNode sets TresPerNode field to given value.

### HasTresPerNode

`func (o *V0044JobInfo) HasTresPerNode() bool`

HasTresPerNode returns a boolean if a field has been set.

### GetTresPerSocket

`func (o *V0044JobInfo) GetTresPerSocket() string`

GetTresPerSocket returns the TresPerSocket field if non-nil, zero value otherwise.

### GetTresPerSocketOk

`func (o *V0044JobInfo) GetTresPerSocketOk() (*string, bool)`

GetTresPerSocketOk returns a tuple with the TresPerSocket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTresPerSocket

`func (o *V0044JobInfo) SetTresPerSocket(v string)`

SetTresPerSocket sets TresPerSocket field to given value.

### HasTresPerSocket

`func (o *V0044JobInfo) HasTresPerSocket() bool`

HasTresPerSocket returns a boolean if a field has been set.

### GetTresPerTask

`func (o *V0044JobInfo) GetTresPerTask() string`

GetTresPerTask returns the TresPerTask field if non-nil, zero value otherwise.

### GetTresPerTaskOk

`func (o *V0044JobInfo) GetTresPerTaskOk() (*string, bool)`

GetTresPerTaskOk returns a tuple with the TresPerTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTresPerTask

`func (o *V0044JobInfo) SetTresPerTask(v string)`

SetTresPerTask sets TresPerTask field to given value.

### HasTresPerTask

`func (o *V0044JobInfo) HasTresPerTask() bool`

HasTresPerTask returns a boolean if a field has been set.

### GetTresReqStr

`func (o *V0044JobInfo) GetTresReqStr() string`

GetTresReqStr returns the TresReqStr field if non-nil, zero value otherwise.

### GetTresReqStrOk

`func (o *V0044JobInfo) GetTresReqStrOk() (*string, bool)`

GetTresReqStrOk returns a tuple with the TresReqStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTresReqStr

`func (o *V0044JobInfo) SetTresReqStr(v string)`

SetTresReqStr sets TresReqStr field to given value.

### HasTresReqStr

`func (o *V0044JobInfo) HasTresReqStr() bool`

HasTresReqStr returns a boolean if a field has been set.

### GetTresAllocStr

`func (o *V0044JobInfo) GetTresAllocStr() string`

GetTresAllocStr returns the TresAllocStr field if non-nil, zero value otherwise.

### GetTresAllocStrOk

`func (o *V0044JobInfo) GetTresAllocStrOk() (*string, bool)`

GetTresAllocStrOk returns a tuple with the TresAllocStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTresAllocStr

`func (o *V0044JobInfo) SetTresAllocStr(v string)`

SetTresAllocStr sets TresAllocStr field to given value.

### HasTresAllocStr

`func (o *V0044JobInfo) HasTresAllocStr() bool`

HasTresAllocStr returns a boolean if a field has been set.

### GetUserId

`func (o *V0044JobInfo) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *V0044JobInfo) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *V0044JobInfo) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *V0044JobInfo) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUserName

`func (o *V0044JobInfo) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *V0044JobInfo) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *V0044JobInfo) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *V0044JobInfo) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetMaximumSwitchWaitTime

`func (o *V0044JobInfo) GetMaximumSwitchWaitTime() int32`

GetMaximumSwitchWaitTime returns the MaximumSwitchWaitTime field if non-nil, zero value otherwise.

### GetMaximumSwitchWaitTimeOk

`func (o *V0044JobInfo) GetMaximumSwitchWaitTimeOk() (*int32, bool)`

GetMaximumSwitchWaitTimeOk returns a tuple with the MaximumSwitchWaitTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumSwitchWaitTime

`func (o *V0044JobInfo) SetMaximumSwitchWaitTime(v int32)`

SetMaximumSwitchWaitTime sets MaximumSwitchWaitTime field to given value.

### HasMaximumSwitchWaitTime

`func (o *V0044JobInfo) HasMaximumSwitchWaitTime() bool`

HasMaximumSwitchWaitTime returns a boolean if a field has been set.

### GetWckey

`func (o *V0044JobInfo) GetWckey() string`

GetWckey returns the Wckey field if non-nil, zero value otherwise.

### GetWckeyOk

`func (o *V0044JobInfo) GetWckeyOk() (*string, bool)`

GetWckeyOk returns a tuple with the Wckey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWckey

`func (o *V0044JobInfo) SetWckey(v string)`

SetWckey sets Wckey field to given value.

### HasWckey

`func (o *V0044JobInfo) HasWckey() bool`

HasWckey returns a boolean if a field has been set.

### GetCurrentWorkingDirectory

`func (o *V0044JobInfo) GetCurrentWorkingDirectory() string`

GetCurrentWorkingDirectory returns the CurrentWorkingDirectory field if non-nil, zero value otherwise.

### GetCurrentWorkingDirectoryOk

`func (o *V0044JobInfo) GetCurrentWorkingDirectoryOk() (*string, bool)`

GetCurrentWorkingDirectoryOk returns a tuple with the CurrentWorkingDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentWorkingDirectory

`func (o *V0044JobInfo) SetCurrentWorkingDirectory(v string)`

SetCurrentWorkingDirectory sets CurrentWorkingDirectory field to given value.

### HasCurrentWorkingDirectory

`func (o *V0044JobInfo) HasCurrentWorkingDirectory() bool`

HasCurrentWorkingDirectory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


