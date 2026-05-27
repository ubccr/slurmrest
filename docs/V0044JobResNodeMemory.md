# V0044JobResNodeMemory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Used** | Pointer to **int64** | Total memory (MiB) used by job | [optional] 
**Allocated** | Pointer to **int64** | Total memory (MiB) allocated to job | [optional] 

## Methods

### NewV0044JobResNodeMemory

`func NewV0044JobResNodeMemory() *V0044JobResNodeMemory`

NewV0044JobResNodeMemory instantiates a new V0044JobResNodeMemory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0044JobResNodeMemoryWithDefaults

`func NewV0044JobResNodeMemoryWithDefaults() *V0044JobResNodeMemory`

NewV0044JobResNodeMemoryWithDefaults instantiates a new V0044JobResNodeMemory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsed

`func (o *V0044JobResNodeMemory) GetUsed() int64`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *V0044JobResNodeMemory) GetUsedOk() (*int64, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *V0044JobResNodeMemory) SetUsed(v int64)`

SetUsed sets Used field to given value.

### HasUsed

`func (o *V0044JobResNodeMemory) HasUsed() bool`

HasUsed returns a boolean if a field has been set.

### GetAllocated

`func (o *V0044JobResNodeMemory) GetAllocated() int64`

GetAllocated returns the Allocated field if non-nil, zero value otherwise.

### GetAllocatedOk

`func (o *V0044JobResNodeMemory) GetAllocatedOk() (*int64, bool)`

GetAllocatedOk returns a tuple with the Allocated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocated

`func (o *V0044JobResNodeMemory) SetAllocated(v int64)`

SetAllocated sets Allocated field to given value.

### HasAllocated

`func (o *V0044JobResNodeMemory) HasAllocated() bool`

HasAllocated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


