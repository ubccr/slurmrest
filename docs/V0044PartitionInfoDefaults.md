# V0044PartitionInfoDefaults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MemoryPerCpu** | Pointer to **int64** | Raw value for DefMemPerCPU or DefMemPerNode | [optional] 
**PartitionMemoryPerCpu** | Pointer to [**V0044Uint64NoValStruct**](V0044Uint64NoValStruct.md) |  | [optional] 
**PartitionMemoryPerNode** | Pointer to [**V0044Uint64NoValStruct**](V0044Uint64NoValStruct.md) |  | [optional] 
**Time** | Pointer to [**V0044Uint32NoValStruct**](V0044Uint32NoValStruct.md) |  | [optional] 
**Job** | Pointer to **string** | JobDefaults - Comma-separated list of job default values (this field is only used to set new defaults) | [optional] 

## Methods

### NewV0044PartitionInfoDefaults

`func NewV0044PartitionInfoDefaults() *V0044PartitionInfoDefaults`

NewV0044PartitionInfoDefaults instantiates a new V0044PartitionInfoDefaults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0044PartitionInfoDefaultsWithDefaults

`func NewV0044PartitionInfoDefaultsWithDefaults() *V0044PartitionInfoDefaults`

NewV0044PartitionInfoDefaultsWithDefaults instantiates a new V0044PartitionInfoDefaults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMemoryPerCpu

`func (o *V0044PartitionInfoDefaults) GetMemoryPerCpu() int64`

GetMemoryPerCpu returns the MemoryPerCpu field if non-nil, zero value otherwise.

### GetMemoryPerCpuOk

`func (o *V0044PartitionInfoDefaults) GetMemoryPerCpuOk() (*int64, bool)`

GetMemoryPerCpuOk returns a tuple with the MemoryPerCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryPerCpu

`func (o *V0044PartitionInfoDefaults) SetMemoryPerCpu(v int64)`

SetMemoryPerCpu sets MemoryPerCpu field to given value.

### HasMemoryPerCpu

`func (o *V0044PartitionInfoDefaults) HasMemoryPerCpu() bool`

HasMemoryPerCpu returns a boolean if a field has been set.

### GetPartitionMemoryPerCpu

`func (o *V0044PartitionInfoDefaults) GetPartitionMemoryPerCpu() V0044Uint64NoValStruct`

GetPartitionMemoryPerCpu returns the PartitionMemoryPerCpu field if non-nil, zero value otherwise.

### GetPartitionMemoryPerCpuOk

`func (o *V0044PartitionInfoDefaults) GetPartitionMemoryPerCpuOk() (*V0044Uint64NoValStruct, bool)`

GetPartitionMemoryPerCpuOk returns a tuple with the PartitionMemoryPerCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionMemoryPerCpu

`func (o *V0044PartitionInfoDefaults) SetPartitionMemoryPerCpu(v V0044Uint64NoValStruct)`

SetPartitionMemoryPerCpu sets PartitionMemoryPerCpu field to given value.

### HasPartitionMemoryPerCpu

`func (o *V0044PartitionInfoDefaults) HasPartitionMemoryPerCpu() bool`

HasPartitionMemoryPerCpu returns a boolean if a field has been set.

### GetPartitionMemoryPerNode

`func (o *V0044PartitionInfoDefaults) GetPartitionMemoryPerNode() V0044Uint64NoValStruct`

GetPartitionMemoryPerNode returns the PartitionMemoryPerNode field if non-nil, zero value otherwise.

### GetPartitionMemoryPerNodeOk

`func (o *V0044PartitionInfoDefaults) GetPartitionMemoryPerNodeOk() (*V0044Uint64NoValStruct, bool)`

GetPartitionMemoryPerNodeOk returns a tuple with the PartitionMemoryPerNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionMemoryPerNode

`func (o *V0044PartitionInfoDefaults) SetPartitionMemoryPerNode(v V0044Uint64NoValStruct)`

SetPartitionMemoryPerNode sets PartitionMemoryPerNode field to given value.

### HasPartitionMemoryPerNode

`func (o *V0044PartitionInfoDefaults) HasPartitionMemoryPerNode() bool`

HasPartitionMemoryPerNode returns a boolean if a field has been set.

### GetTime

`func (o *V0044PartitionInfoDefaults) GetTime() V0044Uint32NoValStruct`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *V0044PartitionInfoDefaults) GetTimeOk() (*V0044Uint32NoValStruct, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *V0044PartitionInfoDefaults) SetTime(v V0044Uint32NoValStruct)`

SetTime sets Time field to given value.

### HasTime

`func (o *V0044PartitionInfoDefaults) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetJob

`func (o *V0044PartitionInfoDefaults) GetJob() string`

GetJob returns the Job field if non-nil, zero value otherwise.

### GetJobOk

`func (o *V0044PartitionInfoDefaults) GetJobOk() (*string, bool)`

GetJobOk returns a tuple with the Job field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJob

`func (o *V0044PartitionInfoDefaults) SetJob(v string)`

SetJob sets Job field to given value.

### HasJob

`func (o *V0044PartitionInfoDefaults) HasJob() bool`

HasJob returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


