# V0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | **int32** | Node index | 
**Name** | **string** | Node name | 
**Cpus** | Pointer to [**V0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInnerCpus**](V0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInnerCpus.md) |  | [optional] 
**Memory** | Pointer to [**V0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInnerMemory**](V0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInnerMemory.md) |  | [optional] 
**Sockets** | [**[]V0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInnerSocketsInner**](V0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInnerSocketsInner.md) | Socket allocations in node | 

## Methods

### NewV0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInner

`func NewV0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInner(index int32, name string, sockets []V0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInnerSocketsInner, ) *V0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInner`

NewV0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInner instantiates a new V0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInnerWithDefaults

`func NewV0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInnerWithDefaults() *V0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInner`

NewV0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInnerWithDefaults instantiates a new V0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *V0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInner) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *V0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInner) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *V0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInner) SetIndex(v int32)`

SetIndex sets Index field to given value.


### GetName

`func (o *V0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInner) SetName(v string)`

SetName sets Name field to given value.


### GetCpus

`func (o *V0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInner) GetCpus() V0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInnerCpus`

GetCpus returns the Cpus field if non-nil, zero value otherwise.

### GetCpusOk

`func (o *V0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInner) GetCpusOk() (*V0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInnerCpus, bool)`

GetCpusOk returns a tuple with the Cpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpus

`func (o *V0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInner) SetCpus(v V0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInnerCpus)`

SetCpus sets Cpus field to given value.

### HasCpus

`func (o *V0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInner) HasCpus() bool`

HasCpus returns a boolean if a field has been set.

### GetMemory

`func (o *V0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInner) GetMemory() V0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInnerMemory`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *V0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInner) GetMemoryOk() (*V0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInnerMemory, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *V0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInner) SetMemory(v V0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInnerMemory)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *V0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInner) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetSockets

`func (o *V0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInner) GetSockets() []V0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInnerSocketsInner`

GetSockets returns the Sockets field if non-nil, zero value otherwise.

### GetSocketsOk

`func (o *V0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInner) GetSocketsOk() (*[]V0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInnerSocketsInner, bool)`

GetSocketsOk returns a tuple with the Sockets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSockets

`func (o *V0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInner) SetSockets(v []V0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInnerSocketsInner)`

SetSockets sets Sockets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


