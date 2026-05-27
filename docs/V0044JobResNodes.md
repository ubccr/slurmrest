# V0044JobResNodes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | Number of allocated nodes | [optional] 
**SelectType** | Pointer to **[]string** | Node scheduling selection method | [optional] 
**List** | Pointer to **string** | Node(s) allocated to the job | [optional] 
**Whole** | Pointer to **bool** | Whether whole nodes were allocated | [optional] 
**Allocation** | Pointer to [**[]V0044JobResNode**](V0044JobResNode.md) | Job resources for a node | [optional] 

## Methods

### NewV0044JobResNodes

`func NewV0044JobResNodes() *V0044JobResNodes`

NewV0044JobResNodes instantiates a new V0044JobResNodes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0044JobResNodesWithDefaults

`func NewV0044JobResNodesWithDefaults() *V0044JobResNodes`

NewV0044JobResNodesWithDefaults instantiates a new V0044JobResNodes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *V0044JobResNodes) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *V0044JobResNodes) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *V0044JobResNodes) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *V0044JobResNodes) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetSelectType

`func (o *V0044JobResNodes) GetSelectType() []string`

GetSelectType returns the SelectType field if non-nil, zero value otherwise.

### GetSelectTypeOk

`func (o *V0044JobResNodes) GetSelectTypeOk() (*[]string, bool)`

GetSelectTypeOk returns a tuple with the SelectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectType

`func (o *V0044JobResNodes) SetSelectType(v []string)`

SetSelectType sets SelectType field to given value.

### HasSelectType

`func (o *V0044JobResNodes) HasSelectType() bool`

HasSelectType returns a boolean if a field has been set.

### GetList

`func (o *V0044JobResNodes) GetList() string`

GetList returns the List field if non-nil, zero value otherwise.

### GetListOk

`func (o *V0044JobResNodes) GetListOk() (*string, bool)`

GetListOk returns a tuple with the List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetList

`func (o *V0044JobResNodes) SetList(v string)`

SetList sets List field to given value.

### HasList

`func (o *V0044JobResNodes) HasList() bool`

HasList returns a boolean if a field has been set.

### GetWhole

`func (o *V0044JobResNodes) GetWhole() bool`

GetWhole returns the Whole field if non-nil, zero value otherwise.

### GetWholeOk

`func (o *V0044JobResNodes) GetWholeOk() (*bool, bool)`

GetWholeOk returns a tuple with the Whole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhole

`func (o *V0044JobResNodes) SetWhole(v bool)`

SetWhole sets Whole field to given value.

### HasWhole

`func (o *V0044JobResNodes) HasWhole() bool`

HasWhole returns a boolean if a field has been set.

### GetAllocation

`func (o *V0044JobResNodes) GetAllocation() []V0044JobResNode`

GetAllocation returns the Allocation field if non-nil, zero value otherwise.

### GetAllocationOk

`func (o *V0044JobResNodes) GetAllocationOk() (*[]V0044JobResNode, bool)`

GetAllocationOk returns a tuple with the Allocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocation

`func (o *V0044JobResNodes) SetAllocation(v []V0044JobResNode)`

SetAllocation sets Allocation field to given value.

### HasAllocation

`func (o *V0044JobResNodes) HasAllocation() bool`

HasAllocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


