# V0044JobResNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | **int32** | Node index | 
**Name** | **string** | Node name | 
**Cpus** | Pointer to [**V0044JobResNodeCpus**](V0044JobResNodeCpus.md) |  | [optional] 
**Memory** | Pointer to [**V0044JobResNodeMemory**](V0044JobResNodeMemory.md) |  | [optional] 
**Sockets** | [**[]V0044JobResSocket**](V0044JobResSocket.md) |  | 

## Methods

### NewV0044JobResNode

`func NewV0044JobResNode(index int32, name string, sockets []V0044JobResSocket, ) *V0044JobResNode`

NewV0044JobResNode instantiates a new V0044JobResNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0044JobResNodeWithDefaults

`func NewV0044JobResNodeWithDefaults() *V0044JobResNode`

NewV0044JobResNodeWithDefaults instantiates a new V0044JobResNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *V0044JobResNode) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *V0044JobResNode) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *V0044JobResNode) SetIndex(v int32)`

SetIndex sets Index field to given value.


### GetName

`func (o *V0044JobResNode) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V0044JobResNode) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V0044JobResNode) SetName(v string)`

SetName sets Name field to given value.


### GetCpus

`func (o *V0044JobResNode) GetCpus() V0044JobResNodeCpus`

GetCpus returns the Cpus field if non-nil, zero value otherwise.

### GetCpusOk

`func (o *V0044JobResNode) GetCpusOk() (*V0044JobResNodeCpus, bool)`

GetCpusOk returns a tuple with the Cpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpus

`func (o *V0044JobResNode) SetCpus(v V0044JobResNodeCpus)`

SetCpus sets Cpus field to given value.

### HasCpus

`func (o *V0044JobResNode) HasCpus() bool`

HasCpus returns a boolean if a field has been set.

### GetMemory

`func (o *V0044JobResNode) GetMemory() V0044JobResNodeMemory`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *V0044JobResNode) GetMemoryOk() (*V0044JobResNodeMemory, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *V0044JobResNode) SetMemory(v V0044JobResNodeMemory)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *V0044JobResNode) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetSockets

`func (o *V0044JobResNode) GetSockets() []V0044JobResSocket`

GetSockets returns the Sockets field if non-nil, zero value otherwise.

### GetSocketsOk

`func (o *V0044JobResNode) GetSocketsOk() (*[]V0044JobResSocket, bool)`

GetSocketsOk returns a tuple with the Sockets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSockets

`func (o *V0044JobResNode) SetSockets(v []V0044JobResSocket)`

SetSockets sets Sockets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


