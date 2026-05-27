# V0044PartPrio

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Partition** | Pointer to **string** | Partition name | [optional] 
**Priority** | Pointer to **int32** | Prospective job priority if it runs in this partition | [optional] 

## Methods

### NewV0044PartPrio

`func NewV0044PartPrio() *V0044PartPrio`

NewV0044PartPrio instantiates a new V0044PartPrio object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0044PartPrioWithDefaults

`func NewV0044PartPrioWithDefaults() *V0044PartPrio`

NewV0044PartPrioWithDefaults instantiates a new V0044PartPrio object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPartition

`func (o *V0044PartPrio) GetPartition() string`

GetPartition returns the Partition field if non-nil, zero value otherwise.

### GetPartitionOk

`func (o *V0044PartPrio) GetPartitionOk() (*string, bool)`

GetPartitionOk returns a tuple with the Partition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartition

`func (o *V0044PartPrio) SetPartition(v string)`

SetPartition sets Partition field to given value.

### HasPartition

`func (o *V0044PartPrio) HasPartition() bool`

HasPartition returns a boolean if a field has been set.

### GetPriority

`func (o *V0044PartPrio) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *V0044PartPrio) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *V0044PartPrio) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *V0044PartPrio) HasPriority() bool`

HasPriority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


