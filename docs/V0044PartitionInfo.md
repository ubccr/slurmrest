# V0044PartitionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nodes** | Pointer to [**V0044PartitionInfoNodes**](V0044PartitionInfoNodes.md) |  | [optional] 
**Accounts** | Pointer to [**V0044PartitionInfoAccounts**](V0044PartitionInfoAccounts.md) |  | [optional] 
**Groups** | Pointer to [**V0044PartitionInfoGroups**](V0044PartitionInfoGroups.md) |  | [optional] 
**Qos** | Pointer to [**V0044PartitionInfoQos**](V0044PartitionInfoQos.md) |  | [optional] 
**Alternate** | Pointer to **string** | Alternate - Partition name of alternate partition to be used if the state of this partition is DRAIN or INACTIVE | [optional] 
**Tres** | Pointer to [**V0044PartitionInfoTres**](V0044PartitionInfoTres.md) |  | [optional] 
**Cluster** | Pointer to **string** | Cluster name | [optional] 
**SelectType** | Pointer to **[]string** | Scheduler consumable resource selection type | [optional] 
**Cpus** | Pointer to [**V0044PartitionInfoCpus**](V0044PartitionInfoCpus.md) |  | [optional] 
**Defaults** | Pointer to [**V0044PartitionInfoDefaults**](V0044PartitionInfoDefaults.md) |  | [optional] 
**GraceTime** | Pointer to **int32** | GraceTime - Grace time in seconds to be extended to a job which has been selected for preemption | [optional] 
**Maximums** | Pointer to [**V0044PartitionInfoMaximums**](V0044PartitionInfoMaximums.md) |  | [optional] 
**Minimums** | Pointer to [**V0044PartitionInfoMinimums**](V0044PartitionInfoMinimums.md) |  | [optional] 
**Name** | Pointer to **string** | PartitionName - Name by which the partition may be referenced | [optional] 
**NodeSets** | Pointer to **string** | NodeSets - Comma-separated list of nodesets which are associated with this partition | [optional] 
**Priority** | Pointer to [**V0044PartitionInfoPriority**](V0044PartitionInfoPriority.md) |  | [optional] 
**Timeouts** | Pointer to [**V0044PartitionInfoTimeouts**](V0044PartitionInfoTimeouts.md) |  | [optional] 
**Topology** | Pointer to **string** | Topology - Name of the topology, defined in topology.yaml, used by jobs in this partition | [optional] 
**Partition** | Pointer to [**V0044PartitionInfoPartition**](V0044PartitionInfoPartition.md) |  | [optional] 
**SuspendTime** | Pointer to [**V0044Uint32NoValStruct**](V0044Uint32NoValStruct.md) |  | [optional] 

## Methods

### NewV0044PartitionInfo

`func NewV0044PartitionInfo() *V0044PartitionInfo`

NewV0044PartitionInfo instantiates a new V0044PartitionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0044PartitionInfoWithDefaults

`func NewV0044PartitionInfoWithDefaults() *V0044PartitionInfo`

NewV0044PartitionInfoWithDefaults instantiates a new V0044PartitionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodes

`func (o *V0044PartitionInfo) GetNodes() V0044PartitionInfoNodes`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *V0044PartitionInfo) GetNodesOk() (*V0044PartitionInfoNodes, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *V0044PartitionInfo) SetNodes(v V0044PartitionInfoNodes)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *V0044PartitionInfo) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetAccounts

`func (o *V0044PartitionInfo) GetAccounts() V0044PartitionInfoAccounts`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *V0044PartitionInfo) GetAccountsOk() (*V0044PartitionInfoAccounts, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *V0044PartitionInfo) SetAccounts(v V0044PartitionInfoAccounts)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *V0044PartitionInfo) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetGroups

`func (o *V0044PartitionInfo) GetGroups() V0044PartitionInfoGroups`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *V0044PartitionInfo) GetGroupsOk() (*V0044PartitionInfoGroups, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *V0044PartitionInfo) SetGroups(v V0044PartitionInfoGroups)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *V0044PartitionInfo) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetQos

`func (o *V0044PartitionInfo) GetQos() V0044PartitionInfoQos`

GetQos returns the Qos field if non-nil, zero value otherwise.

### GetQosOk

`func (o *V0044PartitionInfo) GetQosOk() (*V0044PartitionInfoQos, bool)`

GetQosOk returns a tuple with the Qos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQos

`func (o *V0044PartitionInfo) SetQos(v V0044PartitionInfoQos)`

SetQos sets Qos field to given value.

### HasQos

`func (o *V0044PartitionInfo) HasQos() bool`

HasQos returns a boolean if a field has been set.

### GetAlternate

`func (o *V0044PartitionInfo) GetAlternate() string`

GetAlternate returns the Alternate field if non-nil, zero value otherwise.

### GetAlternateOk

`func (o *V0044PartitionInfo) GetAlternateOk() (*string, bool)`

GetAlternateOk returns a tuple with the Alternate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternate

`func (o *V0044PartitionInfo) SetAlternate(v string)`

SetAlternate sets Alternate field to given value.

### HasAlternate

`func (o *V0044PartitionInfo) HasAlternate() bool`

HasAlternate returns a boolean if a field has been set.

### GetTres

`func (o *V0044PartitionInfo) GetTres() V0044PartitionInfoTres`

GetTres returns the Tres field if non-nil, zero value otherwise.

### GetTresOk

`func (o *V0044PartitionInfo) GetTresOk() (*V0044PartitionInfoTres, bool)`

GetTresOk returns a tuple with the Tres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTres

`func (o *V0044PartitionInfo) SetTres(v V0044PartitionInfoTres)`

SetTres sets Tres field to given value.

### HasTres

`func (o *V0044PartitionInfo) HasTres() bool`

HasTres returns a boolean if a field has been set.

### GetCluster

`func (o *V0044PartitionInfo) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *V0044PartitionInfo) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *V0044PartitionInfo) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *V0044PartitionInfo) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetSelectType

`func (o *V0044PartitionInfo) GetSelectType() []string`

GetSelectType returns the SelectType field if non-nil, zero value otherwise.

### GetSelectTypeOk

`func (o *V0044PartitionInfo) GetSelectTypeOk() (*[]string, bool)`

GetSelectTypeOk returns a tuple with the SelectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectType

`func (o *V0044PartitionInfo) SetSelectType(v []string)`

SetSelectType sets SelectType field to given value.

### HasSelectType

`func (o *V0044PartitionInfo) HasSelectType() bool`

HasSelectType returns a boolean if a field has been set.

### GetCpus

`func (o *V0044PartitionInfo) GetCpus() V0044PartitionInfoCpus`

GetCpus returns the Cpus field if non-nil, zero value otherwise.

### GetCpusOk

`func (o *V0044PartitionInfo) GetCpusOk() (*V0044PartitionInfoCpus, bool)`

GetCpusOk returns a tuple with the Cpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpus

`func (o *V0044PartitionInfo) SetCpus(v V0044PartitionInfoCpus)`

SetCpus sets Cpus field to given value.

### HasCpus

`func (o *V0044PartitionInfo) HasCpus() bool`

HasCpus returns a boolean if a field has been set.

### GetDefaults

`func (o *V0044PartitionInfo) GetDefaults() V0044PartitionInfoDefaults`

GetDefaults returns the Defaults field if non-nil, zero value otherwise.

### GetDefaultsOk

`func (o *V0044PartitionInfo) GetDefaultsOk() (*V0044PartitionInfoDefaults, bool)`

GetDefaultsOk returns a tuple with the Defaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaults

`func (o *V0044PartitionInfo) SetDefaults(v V0044PartitionInfoDefaults)`

SetDefaults sets Defaults field to given value.

### HasDefaults

`func (o *V0044PartitionInfo) HasDefaults() bool`

HasDefaults returns a boolean if a field has been set.

### GetGraceTime

`func (o *V0044PartitionInfo) GetGraceTime() int32`

GetGraceTime returns the GraceTime field if non-nil, zero value otherwise.

### GetGraceTimeOk

`func (o *V0044PartitionInfo) GetGraceTimeOk() (*int32, bool)`

GetGraceTimeOk returns a tuple with the GraceTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraceTime

`func (o *V0044PartitionInfo) SetGraceTime(v int32)`

SetGraceTime sets GraceTime field to given value.

### HasGraceTime

`func (o *V0044PartitionInfo) HasGraceTime() bool`

HasGraceTime returns a boolean if a field has been set.

### GetMaximums

`func (o *V0044PartitionInfo) GetMaximums() V0044PartitionInfoMaximums`

GetMaximums returns the Maximums field if non-nil, zero value otherwise.

### GetMaximumsOk

`func (o *V0044PartitionInfo) GetMaximumsOk() (*V0044PartitionInfoMaximums, bool)`

GetMaximumsOk returns a tuple with the Maximums field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximums

`func (o *V0044PartitionInfo) SetMaximums(v V0044PartitionInfoMaximums)`

SetMaximums sets Maximums field to given value.

### HasMaximums

`func (o *V0044PartitionInfo) HasMaximums() bool`

HasMaximums returns a boolean if a field has been set.

### GetMinimums

`func (o *V0044PartitionInfo) GetMinimums() V0044PartitionInfoMinimums`

GetMinimums returns the Minimums field if non-nil, zero value otherwise.

### GetMinimumsOk

`func (o *V0044PartitionInfo) GetMinimumsOk() (*V0044PartitionInfoMinimums, bool)`

GetMinimumsOk returns a tuple with the Minimums field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimums

`func (o *V0044PartitionInfo) SetMinimums(v V0044PartitionInfoMinimums)`

SetMinimums sets Minimums field to given value.

### HasMinimums

`func (o *V0044PartitionInfo) HasMinimums() bool`

HasMinimums returns a boolean if a field has been set.

### GetName

`func (o *V0044PartitionInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V0044PartitionInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V0044PartitionInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V0044PartitionInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNodeSets

`func (o *V0044PartitionInfo) GetNodeSets() string`

GetNodeSets returns the NodeSets field if non-nil, zero value otherwise.

### GetNodeSetsOk

`func (o *V0044PartitionInfo) GetNodeSetsOk() (*string, bool)`

GetNodeSetsOk returns a tuple with the NodeSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeSets

`func (o *V0044PartitionInfo) SetNodeSets(v string)`

SetNodeSets sets NodeSets field to given value.

### HasNodeSets

`func (o *V0044PartitionInfo) HasNodeSets() bool`

HasNodeSets returns a boolean if a field has been set.

### GetPriority

`func (o *V0044PartitionInfo) GetPriority() V0044PartitionInfoPriority`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *V0044PartitionInfo) GetPriorityOk() (*V0044PartitionInfoPriority, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *V0044PartitionInfo) SetPriority(v V0044PartitionInfoPriority)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *V0044PartitionInfo) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetTimeouts

`func (o *V0044PartitionInfo) GetTimeouts() V0044PartitionInfoTimeouts`

GetTimeouts returns the Timeouts field if non-nil, zero value otherwise.

### GetTimeoutsOk

`func (o *V0044PartitionInfo) GetTimeoutsOk() (*V0044PartitionInfoTimeouts, bool)`

GetTimeoutsOk returns a tuple with the Timeouts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeouts

`func (o *V0044PartitionInfo) SetTimeouts(v V0044PartitionInfoTimeouts)`

SetTimeouts sets Timeouts field to given value.

### HasTimeouts

`func (o *V0044PartitionInfo) HasTimeouts() bool`

HasTimeouts returns a boolean if a field has been set.

### GetTopology

`func (o *V0044PartitionInfo) GetTopology() string`

GetTopology returns the Topology field if non-nil, zero value otherwise.

### GetTopologyOk

`func (o *V0044PartitionInfo) GetTopologyOk() (*string, bool)`

GetTopologyOk returns a tuple with the Topology field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopology

`func (o *V0044PartitionInfo) SetTopology(v string)`

SetTopology sets Topology field to given value.

### HasTopology

`func (o *V0044PartitionInfo) HasTopology() bool`

HasTopology returns a boolean if a field has been set.

### GetPartition

`func (o *V0044PartitionInfo) GetPartition() V0044PartitionInfoPartition`

GetPartition returns the Partition field if non-nil, zero value otherwise.

### GetPartitionOk

`func (o *V0044PartitionInfo) GetPartitionOk() (*V0044PartitionInfoPartition, bool)`

GetPartitionOk returns a tuple with the Partition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartition

`func (o *V0044PartitionInfo) SetPartition(v V0044PartitionInfoPartition)`

SetPartition sets Partition field to given value.

### HasPartition

`func (o *V0044PartitionInfo) HasPartition() bool`

HasPartition returns a boolean if a field has been set.

### GetSuspendTime

`func (o *V0044PartitionInfo) GetSuspendTime() V0044Uint32NoValStruct`

GetSuspendTime returns the SuspendTime field if non-nil, zero value otherwise.

### GetSuspendTimeOk

`func (o *V0044PartitionInfo) GetSuspendTimeOk() (*V0044Uint32NoValStruct, bool)`

GetSuspendTimeOk returns a tuple with the SuspendTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspendTime

`func (o *V0044PartitionInfo) SetSuspendTime(v V0044Uint32NoValStruct)`

SetSuspendTime sets SuspendTime field to given value.

### HasSuspendTime

`func (o *V0044PartitionInfo) HasSuspendTime() bool`

HasSuspendTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


