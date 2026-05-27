# V0044PartitionInfoCpus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaskBinding** | Pointer to **int32** | CpuBind - Default method controlling how tasks are bound to allocated resources | [optional] 
**Total** | Pointer to **int32** | TotalCPUs - Number of CPUs available in this partition | [optional] 

## Methods

### NewV0044PartitionInfoCpus

`func NewV0044PartitionInfoCpus() *V0044PartitionInfoCpus`

NewV0044PartitionInfoCpus instantiates a new V0044PartitionInfoCpus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0044PartitionInfoCpusWithDefaults

`func NewV0044PartitionInfoCpusWithDefaults() *V0044PartitionInfoCpus`

NewV0044PartitionInfoCpusWithDefaults instantiates a new V0044PartitionInfoCpus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaskBinding

`func (o *V0044PartitionInfoCpus) GetTaskBinding() int32`

GetTaskBinding returns the TaskBinding field if non-nil, zero value otherwise.

### GetTaskBindingOk

`func (o *V0044PartitionInfoCpus) GetTaskBindingOk() (*int32, bool)`

GetTaskBindingOk returns a tuple with the TaskBinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskBinding

`func (o *V0044PartitionInfoCpus) SetTaskBinding(v int32)`

SetTaskBinding sets TaskBinding field to given value.

### HasTaskBinding

`func (o *V0044PartitionInfoCpus) HasTaskBinding() bool`

HasTaskBinding returns a boolean if a field has been set.

### GetTotal

`func (o *V0044PartitionInfoCpus) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *V0044PartitionInfoCpus) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *V0044PartitionInfoCpus) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *V0044PartitionInfoCpus) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


