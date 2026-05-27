# V0044JobRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SelectType** | **[]string** | Scheduler consumable resource selection type | 
**Nodes** | Pointer to [**V0044JobResNodes**](V0044JobResNodes.md) |  | [optional] 
**Cpus** | **int32** | Number of allocated CPUs | 
**ThreadsPerCore** | [**V0044Uint16NoValStruct**](V0044Uint16NoValStruct.md) |  | 

## Methods

### NewV0044JobRes

`func NewV0044JobRes(selectType []string, cpus int32, threadsPerCore V0044Uint16NoValStruct, ) *V0044JobRes`

NewV0044JobRes instantiates a new V0044JobRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0044JobResWithDefaults

`func NewV0044JobResWithDefaults() *V0044JobRes`

NewV0044JobResWithDefaults instantiates a new V0044JobRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelectType

`func (o *V0044JobRes) GetSelectType() []string`

GetSelectType returns the SelectType field if non-nil, zero value otherwise.

### GetSelectTypeOk

`func (o *V0044JobRes) GetSelectTypeOk() (*[]string, bool)`

GetSelectTypeOk returns a tuple with the SelectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectType

`func (o *V0044JobRes) SetSelectType(v []string)`

SetSelectType sets SelectType field to given value.


### GetNodes

`func (o *V0044JobRes) GetNodes() V0044JobResNodes`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *V0044JobRes) GetNodesOk() (*V0044JobResNodes, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *V0044JobRes) SetNodes(v V0044JobResNodes)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *V0044JobRes) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetCpus

`func (o *V0044JobRes) GetCpus() int32`

GetCpus returns the Cpus field if non-nil, zero value otherwise.

### GetCpusOk

`func (o *V0044JobRes) GetCpusOk() (*int32, bool)`

GetCpusOk returns a tuple with the Cpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpus

`func (o *V0044JobRes) SetCpus(v int32)`

SetCpus sets Cpus field to given value.


### GetThreadsPerCore

`func (o *V0044JobRes) GetThreadsPerCore() V0044Uint16NoValStruct`

GetThreadsPerCore returns the ThreadsPerCore field if non-nil, zero value otherwise.

### GetThreadsPerCoreOk

`func (o *V0044JobRes) GetThreadsPerCoreOk() (*V0044Uint16NoValStruct, bool)`

GetThreadsPerCoreOk returns a tuple with the ThreadsPerCore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadsPerCore

`func (o *V0044JobRes) SetThreadsPerCore(v V0044Uint16NoValStruct)`

SetThreadsPerCore sets ThreadsPerCore field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


