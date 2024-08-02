# V0041OpenapiPartitionRespPartitionsInnerDefaults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MemoryPerCpu** | Pointer to **int64** |  | [optional] 
**PartitionMemoryPerCpu** | Pointer to [**V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu**](V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu.md) |  | [optional] 
**PartitionMemoryPerNode** | Pointer to [**V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu**](V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu.md) |  | [optional] 
**Time** | Pointer to [**V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId**](V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId.md) |  | [optional] 
**Job** | Pointer to **string** |  | [optional] 

## Methods

### NewV0041OpenapiPartitionRespPartitionsInnerDefaults

`func NewV0041OpenapiPartitionRespPartitionsInnerDefaults() *V0041OpenapiPartitionRespPartitionsInnerDefaults`

NewV0041OpenapiPartitionRespPartitionsInnerDefaults instantiates a new V0041OpenapiPartitionRespPartitionsInnerDefaults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0041OpenapiPartitionRespPartitionsInnerDefaultsWithDefaults

`func NewV0041OpenapiPartitionRespPartitionsInnerDefaultsWithDefaults() *V0041OpenapiPartitionRespPartitionsInnerDefaults`

NewV0041OpenapiPartitionRespPartitionsInnerDefaultsWithDefaults instantiates a new V0041OpenapiPartitionRespPartitionsInnerDefaults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMemoryPerCpu

`func (o *V0041OpenapiPartitionRespPartitionsInnerDefaults) GetMemoryPerCpu() int64`

GetMemoryPerCpu returns the MemoryPerCpu field if non-nil, zero value otherwise.

### GetMemoryPerCpuOk

`func (o *V0041OpenapiPartitionRespPartitionsInnerDefaults) GetMemoryPerCpuOk() (*int64, bool)`

GetMemoryPerCpuOk returns a tuple with the MemoryPerCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryPerCpu

`func (o *V0041OpenapiPartitionRespPartitionsInnerDefaults) SetMemoryPerCpu(v int64)`

SetMemoryPerCpu sets MemoryPerCpu field to given value.

### HasMemoryPerCpu

`func (o *V0041OpenapiPartitionRespPartitionsInnerDefaults) HasMemoryPerCpu() bool`

HasMemoryPerCpu returns a boolean if a field has been set.

### GetPartitionMemoryPerCpu

`func (o *V0041OpenapiPartitionRespPartitionsInnerDefaults) GetPartitionMemoryPerCpu() V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu`

GetPartitionMemoryPerCpu returns the PartitionMemoryPerCpu field if non-nil, zero value otherwise.

### GetPartitionMemoryPerCpuOk

`func (o *V0041OpenapiPartitionRespPartitionsInnerDefaults) GetPartitionMemoryPerCpuOk() (*V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu, bool)`

GetPartitionMemoryPerCpuOk returns a tuple with the PartitionMemoryPerCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionMemoryPerCpu

`func (o *V0041OpenapiPartitionRespPartitionsInnerDefaults) SetPartitionMemoryPerCpu(v V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu)`

SetPartitionMemoryPerCpu sets PartitionMemoryPerCpu field to given value.

### HasPartitionMemoryPerCpu

`func (o *V0041OpenapiPartitionRespPartitionsInnerDefaults) HasPartitionMemoryPerCpu() bool`

HasPartitionMemoryPerCpu returns a boolean if a field has been set.

### GetPartitionMemoryPerNode

`func (o *V0041OpenapiPartitionRespPartitionsInnerDefaults) GetPartitionMemoryPerNode() V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu`

GetPartitionMemoryPerNode returns the PartitionMemoryPerNode field if non-nil, zero value otherwise.

### GetPartitionMemoryPerNodeOk

`func (o *V0041OpenapiPartitionRespPartitionsInnerDefaults) GetPartitionMemoryPerNodeOk() (*V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu, bool)`

GetPartitionMemoryPerNodeOk returns a tuple with the PartitionMemoryPerNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionMemoryPerNode

`func (o *V0041OpenapiPartitionRespPartitionsInnerDefaults) SetPartitionMemoryPerNode(v V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu)`

SetPartitionMemoryPerNode sets PartitionMemoryPerNode field to given value.

### HasPartitionMemoryPerNode

`func (o *V0041OpenapiPartitionRespPartitionsInnerDefaults) HasPartitionMemoryPerNode() bool`

HasPartitionMemoryPerNode returns a boolean if a field has been set.

### GetTime

`func (o *V0041OpenapiPartitionRespPartitionsInnerDefaults) GetTime() V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *V0041OpenapiPartitionRespPartitionsInnerDefaults) GetTimeOk() (*V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *V0041OpenapiPartitionRespPartitionsInnerDefaults) SetTime(v V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId)`

SetTime sets Time field to given value.

### HasTime

`func (o *V0041OpenapiPartitionRespPartitionsInnerDefaults) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetJob

`func (o *V0041OpenapiPartitionRespPartitionsInnerDefaults) GetJob() string`

GetJob returns the Job field if non-nil, zero value otherwise.

### GetJobOk

`func (o *V0041OpenapiPartitionRespPartitionsInnerDefaults) GetJobOk() (*string, bool)`

GetJobOk returns a tuple with the Job field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJob

`func (o *V0041OpenapiPartitionRespPartitionsInnerDefaults) SetJob(v string)`

SetJob sets Job field to given value.

### HasJob

`func (o *V0041OpenapiPartitionRespPartitionsInnerDefaults) HasJob() bool`

HasJob returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


