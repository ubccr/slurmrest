# V0041OpenapiPartitionRespPartitionsInnerMaximums

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CpusPerNode** | Pointer to [**V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId**](V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId.md) |  | [optional] 
**CpusPerSocket** | Pointer to [**V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId**](V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId.md) |  | [optional] 
**MemoryPerCpu** | Pointer to **int64** |  | [optional] 
**PartitionMemoryPerCpu** | Pointer to [**V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu**](V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu.md) |  | [optional] 
**PartitionMemoryPerNode** | Pointer to [**V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu**](V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu.md) |  | [optional] 
**Nodes** | Pointer to [**V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId**](V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId.md) |  | [optional] 
**Shares** | Pointer to **int32** |  | [optional] 
**Oversubscribe** | Pointer to [**V0041OpenapiPartitionRespPartitionsInnerMaximumsOversubscribe**](V0041OpenapiPartitionRespPartitionsInnerMaximumsOversubscribe.md) |  | [optional] 
**Time** | Pointer to [**V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId**](V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId.md) |  | [optional] 
**OverTimeLimit** | Pointer to [**V0041JobDescMsgDistributionPlaneSize**](V0041JobDescMsgDistributionPlaneSize.md) |  | [optional] 

## Methods

### NewV0041OpenapiPartitionRespPartitionsInnerMaximums

`func NewV0041OpenapiPartitionRespPartitionsInnerMaximums() *V0041OpenapiPartitionRespPartitionsInnerMaximums`

NewV0041OpenapiPartitionRespPartitionsInnerMaximums instantiates a new V0041OpenapiPartitionRespPartitionsInnerMaximums object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0041OpenapiPartitionRespPartitionsInnerMaximumsWithDefaults

`func NewV0041OpenapiPartitionRespPartitionsInnerMaximumsWithDefaults() *V0041OpenapiPartitionRespPartitionsInnerMaximums`

NewV0041OpenapiPartitionRespPartitionsInnerMaximumsWithDefaults instantiates a new V0041OpenapiPartitionRespPartitionsInnerMaximums object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpusPerNode

`func (o *V0041OpenapiPartitionRespPartitionsInnerMaximums) GetCpusPerNode() V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId`

GetCpusPerNode returns the CpusPerNode field if non-nil, zero value otherwise.

### GetCpusPerNodeOk

`func (o *V0041OpenapiPartitionRespPartitionsInnerMaximums) GetCpusPerNodeOk() (*V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId, bool)`

GetCpusPerNodeOk returns a tuple with the CpusPerNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpusPerNode

`func (o *V0041OpenapiPartitionRespPartitionsInnerMaximums) SetCpusPerNode(v V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId)`

SetCpusPerNode sets CpusPerNode field to given value.

### HasCpusPerNode

`func (o *V0041OpenapiPartitionRespPartitionsInnerMaximums) HasCpusPerNode() bool`

HasCpusPerNode returns a boolean if a field has been set.

### GetCpusPerSocket

`func (o *V0041OpenapiPartitionRespPartitionsInnerMaximums) GetCpusPerSocket() V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId`

GetCpusPerSocket returns the CpusPerSocket field if non-nil, zero value otherwise.

### GetCpusPerSocketOk

`func (o *V0041OpenapiPartitionRespPartitionsInnerMaximums) GetCpusPerSocketOk() (*V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId, bool)`

GetCpusPerSocketOk returns a tuple with the CpusPerSocket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpusPerSocket

`func (o *V0041OpenapiPartitionRespPartitionsInnerMaximums) SetCpusPerSocket(v V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId)`

SetCpusPerSocket sets CpusPerSocket field to given value.

### HasCpusPerSocket

`func (o *V0041OpenapiPartitionRespPartitionsInnerMaximums) HasCpusPerSocket() bool`

HasCpusPerSocket returns a boolean if a field has been set.

### GetMemoryPerCpu

`func (o *V0041OpenapiPartitionRespPartitionsInnerMaximums) GetMemoryPerCpu() int64`

GetMemoryPerCpu returns the MemoryPerCpu field if non-nil, zero value otherwise.

### GetMemoryPerCpuOk

`func (o *V0041OpenapiPartitionRespPartitionsInnerMaximums) GetMemoryPerCpuOk() (*int64, bool)`

GetMemoryPerCpuOk returns a tuple with the MemoryPerCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryPerCpu

`func (o *V0041OpenapiPartitionRespPartitionsInnerMaximums) SetMemoryPerCpu(v int64)`

SetMemoryPerCpu sets MemoryPerCpu field to given value.

### HasMemoryPerCpu

`func (o *V0041OpenapiPartitionRespPartitionsInnerMaximums) HasMemoryPerCpu() bool`

HasMemoryPerCpu returns a boolean if a field has been set.

### GetPartitionMemoryPerCpu

`func (o *V0041OpenapiPartitionRespPartitionsInnerMaximums) GetPartitionMemoryPerCpu() V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu`

GetPartitionMemoryPerCpu returns the PartitionMemoryPerCpu field if non-nil, zero value otherwise.

### GetPartitionMemoryPerCpuOk

`func (o *V0041OpenapiPartitionRespPartitionsInnerMaximums) GetPartitionMemoryPerCpuOk() (*V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu, bool)`

GetPartitionMemoryPerCpuOk returns a tuple with the PartitionMemoryPerCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionMemoryPerCpu

`func (o *V0041OpenapiPartitionRespPartitionsInnerMaximums) SetPartitionMemoryPerCpu(v V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu)`

SetPartitionMemoryPerCpu sets PartitionMemoryPerCpu field to given value.

### HasPartitionMemoryPerCpu

`func (o *V0041OpenapiPartitionRespPartitionsInnerMaximums) HasPartitionMemoryPerCpu() bool`

HasPartitionMemoryPerCpu returns a boolean if a field has been set.

### GetPartitionMemoryPerNode

`func (o *V0041OpenapiPartitionRespPartitionsInnerMaximums) GetPartitionMemoryPerNode() V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu`

GetPartitionMemoryPerNode returns the PartitionMemoryPerNode field if non-nil, zero value otherwise.

### GetPartitionMemoryPerNodeOk

`func (o *V0041OpenapiPartitionRespPartitionsInnerMaximums) GetPartitionMemoryPerNodeOk() (*V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu, bool)`

GetPartitionMemoryPerNodeOk returns a tuple with the PartitionMemoryPerNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionMemoryPerNode

`func (o *V0041OpenapiPartitionRespPartitionsInnerMaximums) SetPartitionMemoryPerNode(v V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu)`

SetPartitionMemoryPerNode sets PartitionMemoryPerNode field to given value.

### HasPartitionMemoryPerNode

`func (o *V0041OpenapiPartitionRespPartitionsInnerMaximums) HasPartitionMemoryPerNode() bool`

HasPartitionMemoryPerNode returns a boolean if a field has been set.

### GetNodes

`func (o *V0041OpenapiPartitionRespPartitionsInnerMaximums) GetNodes() V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *V0041OpenapiPartitionRespPartitionsInnerMaximums) GetNodesOk() (*V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *V0041OpenapiPartitionRespPartitionsInnerMaximums) SetNodes(v V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *V0041OpenapiPartitionRespPartitionsInnerMaximums) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetShares

`func (o *V0041OpenapiPartitionRespPartitionsInnerMaximums) GetShares() int32`

GetShares returns the Shares field if non-nil, zero value otherwise.

### GetSharesOk

`func (o *V0041OpenapiPartitionRespPartitionsInnerMaximums) GetSharesOk() (*int32, bool)`

GetSharesOk returns a tuple with the Shares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShares

`func (o *V0041OpenapiPartitionRespPartitionsInnerMaximums) SetShares(v int32)`

SetShares sets Shares field to given value.

### HasShares

`func (o *V0041OpenapiPartitionRespPartitionsInnerMaximums) HasShares() bool`

HasShares returns a boolean if a field has been set.

### GetOversubscribe

`func (o *V0041OpenapiPartitionRespPartitionsInnerMaximums) GetOversubscribe() V0041OpenapiPartitionRespPartitionsInnerMaximumsOversubscribe`

GetOversubscribe returns the Oversubscribe field if non-nil, zero value otherwise.

### GetOversubscribeOk

`func (o *V0041OpenapiPartitionRespPartitionsInnerMaximums) GetOversubscribeOk() (*V0041OpenapiPartitionRespPartitionsInnerMaximumsOversubscribe, bool)`

GetOversubscribeOk returns a tuple with the Oversubscribe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOversubscribe

`func (o *V0041OpenapiPartitionRespPartitionsInnerMaximums) SetOversubscribe(v V0041OpenapiPartitionRespPartitionsInnerMaximumsOversubscribe)`

SetOversubscribe sets Oversubscribe field to given value.

### HasOversubscribe

`func (o *V0041OpenapiPartitionRespPartitionsInnerMaximums) HasOversubscribe() bool`

HasOversubscribe returns a boolean if a field has been set.

### GetTime

`func (o *V0041OpenapiPartitionRespPartitionsInnerMaximums) GetTime() V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *V0041OpenapiPartitionRespPartitionsInnerMaximums) GetTimeOk() (*V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *V0041OpenapiPartitionRespPartitionsInnerMaximums) SetTime(v V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId)`

SetTime sets Time field to given value.

### HasTime

`func (o *V0041OpenapiPartitionRespPartitionsInnerMaximums) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetOverTimeLimit

`func (o *V0041OpenapiPartitionRespPartitionsInnerMaximums) GetOverTimeLimit() V0041JobDescMsgDistributionPlaneSize`

GetOverTimeLimit returns the OverTimeLimit field if non-nil, zero value otherwise.

### GetOverTimeLimitOk

`func (o *V0041OpenapiPartitionRespPartitionsInnerMaximums) GetOverTimeLimitOk() (*V0041JobDescMsgDistributionPlaneSize, bool)`

GetOverTimeLimitOk returns a tuple with the OverTimeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverTimeLimit

`func (o *V0041OpenapiPartitionRespPartitionsInnerMaximums) SetOverTimeLimit(v V0041JobDescMsgDistributionPlaneSize)`

SetOverTimeLimit sets OverTimeLimit field to given value.

### HasOverTimeLimit

`func (o *V0041OpenapiPartitionRespPartitionsInnerMaximums) HasOverTimeLimit() bool`

HasOverTimeLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


