# V0041OpenapiPartitionRespPartitionsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nodes** | Pointer to [**V0041OpenapiPartitionRespPartitionsInnerNodes**](V0041OpenapiPartitionRespPartitionsInnerNodes.md) |  | [optional] 
**Accounts** | Pointer to [**V0041OpenapiPartitionRespPartitionsInnerAccounts**](V0041OpenapiPartitionRespPartitionsInnerAccounts.md) |  | [optional] 
**Groups** | Pointer to [**V0041OpenapiPartitionRespPartitionsInnerGroups**](V0041OpenapiPartitionRespPartitionsInnerGroups.md) |  | [optional] 
**Qos** | Pointer to [**V0041OpenapiPartitionRespPartitionsInnerQos**](V0041OpenapiPartitionRespPartitionsInnerQos.md) |  | [optional] 
**Alternate** | Pointer to **string** |  | [optional] 
**Tres** | Pointer to [**V0041OpenapiPartitionRespPartitionsInnerTres**](V0041OpenapiPartitionRespPartitionsInnerTres.md) |  | [optional] 
**Cluster** | Pointer to **string** |  | [optional] 
**SelectType** | Pointer to **[]string** | Scheduler consumable resource selection types | [optional] 
**Cpus** | Pointer to [**V0041OpenapiPartitionRespPartitionsInnerCpus**](V0041OpenapiPartitionRespPartitionsInnerCpus.md) |  | [optional] 
**Defaults** | Pointer to [**V0041OpenapiPartitionRespPartitionsInnerDefaults**](V0041OpenapiPartitionRespPartitionsInnerDefaults.md) |  | [optional] 
**GraceTime** | Pointer to **int32** |  | [optional] 
**Maximums** | Pointer to [**V0041OpenapiPartitionRespPartitionsInnerMaximums**](V0041OpenapiPartitionRespPartitionsInnerMaximums.md) |  | [optional] 
**Minimums** | Pointer to [**V0041OpenapiPartitionRespPartitionsInnerMinimums**](V0041OpenapiPartitionRespPartitionsInnerMinimums.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NodeSets** | Pointer to **string** |  | [optional] 
**Priority** | Pointer to [**V0041OpenapiPartitionRespPartitionsInnerPriority**](V0041OpenapiPartitionRespPartitionsInnerPriority.md) |  | [optional] 
**Timeouts** | Pointer to [**V0041OpenapiPartitionRespPartitionsInnerTimeouts**](V0041OpenapiPartitionRespPartitionsInnerTimeouts.md) |  | [optional] 
**Partition** | Pointer to [**V0041OpenapiPartitionRespPartitionsInnerPartition**](V0041OpenapiPartitionRespPartitionsInnerPartition.md) |  | [optional] 
**SuspendTime** | Pointer to [**V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId**](V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId.md) |  | [optional] 

## Methods

### NewV0041OpenapiPartitionRespPartitionsInner

`func NewV0041OpenapiPartitionRespPartitionsInner() *V0041OpenapiPartitionRespPartitionsInner`

NewV0041OpenapiPartitionRespPartitionsInner instantiates a new V0041OpenapiPartitionRespPartitionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0041OpenapiPartitionRespPartitionsInnerWithDefaults

`func NewV0041OpenapiPartitionRespPartitionsInnerWithDefaults() *V0041OpenapiPartitionRespPartitionsInner`

NewV0041OpenapiPartitionRespPartitionsInnerWithDefaults instantiates a new V0041OpenapiPartitionRespPartitionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodes

`func (o *V0041OpenapiPartitionRespPartitionsInner) GetNodes() V0041OpenapiPartitionRespPartitionsInnerNodes`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *V0041OpenapiPartitionRespPartitionsInner) GetNodesOk() (*V0041OpenapiPartitionRespPartitionsInnerNodes, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *V0041OpenapiPartitionRespPartitionsInner) SetNodes(v V0041OpenapiPartitionRespPartitionsInnerNodes)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *V0041OpenapiPartitionRespPartitionsInner) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetAccounts

`func (o *V0041OpenapiPartitionRespPartitionsInner) GetAccounts() V0041OpenapiPartitionRespPartitionsInnerAccounts`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *V0041OpenapiPartitionRespPartitionsInner) GetAccountsOk() (*V0041OpenapiPartitionRespPartitionsInnerAccounts, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *V0041OpenapiPartitionRespPartitionsInner) SetAccounts(v V0041OpenapiPartitionRespPartitionsInnerAccounts)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *V0041OpenapiPartitionRespPartitionsInner) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetGroups

`func (o *V0041OpenapiPartitionRespPartitionsInner) GetGroups() V0041OpenapiPartitionRespPartitionsInnerGroups`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *V0041OpenapiPartitionRespPartitionsInner) GetGroupsOk() (*V0041OpenapiPartitionRespPartitionsInnerGroups, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *V0041OpenapiPartitionRespPartitionsInner) SetGroups(v V0041OpenapiPartitionRespPartitionsInnerGroups)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *V0041OpenapiPartitionRespPartitionsInner) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetQos

`func (o *V0041OpenapiPartitionRespPartitionsInner) GetQos() V0041OpenapiPartitionRespPartitionsInnerQos`

GetQos returns the Qos field if non-nil, zero value otherwise.

### GetQosOk

`func (o *V0041OpenapiPartitionRespPartitionsInner) GetQosOk() (*V0041OpenapiPartitionRespPartitionsInnerQos, bool)`

GetQosOk returns a tuple with the Qos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQos

`func (o *V0041OpenapiPartitionRespPartitionsInner) SetQos(v V0041OpenapiPartitionRespPartitionsInnerQos)`

SetQos sets Qos field to given value.

### HasQos

`func (o *V0041OpenapiPartitionRespPartitionsInner) HasQos() bool`

HasQos returns a boolean if a field has been set.

### GetAlternate

`func (o *V0041OpenapiPartitionRespPartitionsInner) GetAlternate() string`

GetAlternate returns the Alternate field if non-nil, zero value otherwise.

### GetAlternateOk

`func (o *V0041OpenapiPartitionRespPartitionsInner) GetAlternateOk() (*string, bool)`

GetAlternateOk returns a tuple with the Alternate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternate

`func (o *V0041OpenapiPartitionRespPartitionsInner) SetAlternate(v string)`

SetAlternate sets Alternate field to given value.

### HasAlternate

`func (o *V0041OpenapiPartitionRespPartitionsInner) HasAlternate() bool`

HasAlternate returns a boolean if a field has been set.

### GetTres

`func (o *V0041OpenapiPartitionRespPartitionsInner) GetTres() V0041OpenapiPartitionRespPartitionsInnerTres`

GetTres returns the Tres field if non-nil, zero value otherwise.

### GetTresOk

`func (o *V0041OpenapiPartitionRespPartitionsInner) GetTresOk() (*V0041OpenapiPartitionRespPartitionsInnerTres, bool)`

GetTresOk returns a tuple with the Tres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTres

`func (o *V0041OpenapiPartitionRespPartitionsInner) SetTres(v V0041OpenapiPartitionRespPartitionsInnerTres)`

SetTres sets Tres field to given value.

### HasTres

`func (o *V0041OpenapiPartitionRespPartitionsInner) HasTres() bool`

HasTres returns a boolean if a field has been set.

### GetCluster

`func (o *V0041OpenapiPartitionRespPartitionsInner) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *V0041OpenapiPartitionRespPartitionsInner) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *V0041OpenapiPartitionRespPartitionsInner) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *V0041OpenapiPartitionRespPartitionsInner) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetSelectType

`func (o *V0041OpenapiPartitionRespPartitionsInner) GetSelectType() []string`

GetSelectType returns the SelectType field if non-nil, zero value otherwise.

### GetSelectTypeOk

`func (o *V0041OpenapiPartitionRespPartitionsInner) GetSelectTypeOk() (*[]string, bool)`

GetSelectTypeOk returns a tuple with the SelectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectType

`func (o *V0041OpenapiPartitionRespPartitionsInner) SetSelectType(v []string)`

SetSelectType sets SelectType field to given value.

### HasSelectType

`func (o *V0041OpenapiPartitionRespPartitionsInner) HasSelectType() bool`

HasSelectType returns a boolean if a field has been set.

### GetCpus

`func (o *V0041OpenapiPartitionRespPartitionsInner) GetCpus() V0041OpenapiPartitionRespPartitionsInnerCpus`

GetCpus returns the Cpus field if non-nil, zero value otherwise.

### GetCpusOk

`func (o *V0041OpenapiPartitionRespPartitionsInner) GetCpusOk() (*V0041OpenapiPartitionRespPartitionsInnerCpus, bool)`

GetCpusOk returns a tuple with the Cpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpus

`func (o *V0041OpenapiPartitionRespPartitionsInner) SetCpus(v V0041OpenapiPartitionRespPartitionsInnerCpus)`

SetCpus sets Cpus field to given value.

### HasCpus

`func (o *V0041OpenapiPartitionRespPartitionsInner) HasCpus() bool`

HasCpus returns a boolean if a field has been set.

### GetDefaults

`func (o *V0041OpenapiPartitionRespPartitionsInner) GetDefaults() V0041OpenapiPartitionRespPartitionsInnerDefaults`

GetDefaults returns the Defaults field if non-nil, zero value otherwise.

### GetDefaultsOk

`func (o *V0041OpenapiPartitionRespPartitionsInner) GetDefaultsOk() (*V0041OpenapiPartitionRespPartitionsInnerDefaults, bool)`

GetDefaultsOk returns a tuple with the Defaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaults

`func (o *V0041OpenapiPartitionRespPartitionsInner) SetDefaults(v V0041OpenapiPartitionRespPartitionsInnerDefaults)`

SetDefaults sets Defaults field to given value.

### HasDefaults

`func (o *V0041OpenapiPartitionRespPartitionsInner) HasDefaults() bool`

HasDefaults returns a boolean if a field has been set.

### GetGraceTime

`func (o *V0041OpenapiPartitionRespPartitionsInner) GetGraceTime() int32`

GetGraceTime returns the GraceTime field if non-nil, zero value otherwise.

### GetGraceTimeOk

`func (o *V0041OpenapiPartitionRespPartitionsInner) GetGraceTimeOk() (*int32, bool)`

GetGraceTimeOk returns a tuple with the GraceTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraceTime

`func (o *V0041OpenapiPartitionRespPartitionsInner) SetGraceTime(v int32)`

SetGraceTime sets GraceTime field to given value.

### HasGraceTime

`func (o *V0041OpenapiPartitionRespPartitionsInner) HasGraceTime() bool`

HasGraceTime returns a boolean if a field has been set.

### GetMaximums

`func (o *V0041OpenapiPartitionRespPartitionsInner) GetMaximums() V0041OpenapiPartitionRespPartitionsInnerMaximums`

GetMaximums returns the Maximums field if non-nil, zero value otherwise.

### GetMaximumsOk

`func (o *V0041OpenapiPartitionRespPartitionsInner) GetMaximumsOk() (*V0041OpenapiPartitionRespPartitionsInnerMaximums, bool)`

GetMaximumsOk returns a tuple with the Maximums field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximums

`func (o *V0041OpenapiPartitionRespPartitionsInner) SetMaximums(v V0041OpenapiPartitionRespPartitionsInnerMaximums)`

SetMaximums sets Maximums field to given value.

### HasMaximums

`func (o *V0041OpenapiPartitionRespPartitionsInner) HasMaximums() bool`

HasMaximums returns a boolean if a field has been set.

### GetMinimums

`func (o *V0041OpenapiPartitionRespPartitionsInner) GetMinimums() V0041OpenapiPartitionRespPartitionsInnerMinimums`

GetMinimums returns the Minimums field if non-nil, zero value otherwise.

### GetMinimumsOk

`func (o *V0041OpenapiPartitionRespPartitionsInner) GetMinimumsOk() (*V0041OpenapiPartitionRespPartitionsInnerMinimums, bool)`

GetMinimumsOk returns a tuple with the Minimums field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimums

`func (o *V0041OpenapiPartitionRespPartitionsInner) SetMinimums(v V0041OpenapiPartitionRespPartitionsInnerMinimums)`

SetMinimums sets Minimums field to given value.

### HasMinimums

`func (o *V0041OpenapiPartitionRespPartitionsInner) HasMinimums() bool`

HasMinimums returns a boolean if a field has been set.

### GetName

`func (o *V0041OpenapiPartitionRespPartitionsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V0041OpenapiPartitionRespPartitionsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V0041OpenapiPartitionRespPartitionsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V0041OpenapiPartitionRespPartitionsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNodeSets

`func (o *V0041OpenapiPartitionRespPartitionsInner) GetNodeSets() string`

GetNodeSets returns the NodeSets field if non-nil, zero value otherwise.

### GetNodeSetsOk

`func (o *V0041OpenapiPartitionRespPartitionsInner) GetNodeSetsOk() (*string, bool)`

GetNodeSetsOk returns a tuple with the NodeSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeSets

`func (o *V0041OpenapiPartitionRespPartitionsInner) SetNodeSets(v string)`

SetNodeSets sets NodeSets field to given value.

### HasNodeSets

`func (o *V0041OpenapiPartitionRespPartitionsInner) HasNodeSets() bool`

HasNodeSets returns a boolean if a field has been set.

### GetPriority

`func (o *V0041OpenapiPartitionRespPartitionsInner) GetPriority() V0041OpenapiPartitionRespPartitionsInnerPriority`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *V0041OpenapiPartitionRespPartitionsInner) GetPriorityOk() (*V0041OpenapiPartitionRespPartitionsInnerPriority, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *V0041OpenapiPartitionRespPartitionsInner) SetPriority(v V0041OpenapiPartitionRespPartitionsInnerPriority)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *V0041OpenapiPartitionRespPartitionsInner) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetTimeouts

`func (o *V0041OpenapiPartitionRespPartitionsInner) GetTimeouts() V0041OpenapiPartitionRespPartitionsInnerTimeouts`

GetTimeouts returns the Timeouts field if non-nil, zero value otherwise.

### GetTimeoutsOk

`func (o *V0041OpenapiPartitionRespPartitionsInner) GetTimeoutsOk() (*V0041OpenapiPartitionRespPartitionsInnerTimeouts, bool)`

GetTimeoutsOk returns a tuple with the Timeouts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeouts

`func (o *V0041OpenapiPartitionRespPartitionsInner) SetTimeouts(v V0041OpenapiPartitionRespPartitionsInnerTimeouts)`

SetTimeouts sets Timeouts field to given value.

### HasTimeouts

`func (o *V0041OpenapiPartitionRespPartitionsInner) HasTimeouts() bool`

HasTimeouts returns a boolean if a field has been set.

### GetPartition

`func (o *V0041OpenapiPartitionRespPartitionsInner) GetPartition() V0041OpenapiPartitionRespPartitionsInnerPartition`

GetPartition returns the Partition field if non-nil, zero value otherwise.

### GetPartitionOk

`func (o *V0041OpenapiPartitionRespPartitionsInner) GetPartitionOk() (*V0041OpenapiPartitionRespPartitionsInnerPartition, bool)`

GetPartitionOk returns a tuple with the Partition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartition

`func (o *V0041OpenapiPartitionRespPartitionsInner) SetPartition(v V0041OpenapiPartitionRespPartitionsInnerPartition)`

SetPartition sets Partition field to given value.

### HasPartition

`func (o *V0041OpenapiPartitionRespPartitionsInner) HasPartition() bool`

HasPartition returns a boolean if a field has been set.

### GetSuspendTime

`func (o *V0041OpenapiPartitionRespPartitionsInner) GetSuspendTime() V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId`

GetSuspendTime returns the SuspendTime field if non-nil, zero value otherwise.

### GetSuspendTimeOk

`func (o *V0041OpenapiPartitionRespPartitionsInner) GetSuspendTimeOk() (*V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId, bool)`

GetSuspendTimeOk returns a tuple with the SuspendTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspendTime

`func (o *V0041OpenapiPartitionRespPartitionsInner) SetSuspendTime(v V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId)`

SetSuspendTime sets SuspendTime field to given value.

### HasSuspendTime

`func (o *V0041OpenapiPartitionRespPartitionsInner) HasSuspendTime() bool`

HasSuspendTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


