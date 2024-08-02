# V0041OpenapiJobInfoRespJobsInnerJobResourcesNodes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | Number of nodes assigned to job | [optional] 
**SelectType** | Pointer to **[]string** | Node scheduling selection request | [optional] 
**List** | Pointer to **string** | host list for job | [optional] 
**Whole** | Pointer to **bool** | Job allocated full nodes | [optional] 
**Allocation** | Pointer to [**[]V0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInner**](V0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInner.md) | resource allocations by node | [optional] 

## Methods

### NewV0041OpenapiJobInfoRespJobsInnerJobResourcesNodes

`func NewV0041OpenapiJobInfoRespJobsInnerJobResourcesNodes() *V0041OpenapiJobInfoRespJobsInnerJobResourcesNodes`

NewV0041OpenapiJobInfoRespJobsInnerJobResourcesNodes instantiates a new V0041OpenapiJobInfoRespJobsInnerJobResourcesNodes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0041OpenapiJobInfoRespJobsInnerJobResourcesNodesWithDefaults

`func NewV0041OpenapiJobInfoRespJobsInnerJobResourcesNodesWithDefaults() *V0041OpenapiJobInfoRespJobsInnerJobResourcesNodes`

NewV0041OpenapiJobInfoRespJobsInnerJobResourcesNodesWithDefaults instantiates a new V0041OpenapiJobInfoRespJobsInnerJobResourcesNodes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *V0041OpenapiJobInfoRespJobsInnerJobResourcesNodes) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *V0041OpenapiJobInfoRespJobsInnerJobResourcesNodes) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *V0041OpenapiJobInfoRespJobsInnerJobResourcesNodes) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *V0041OpenapiJobInfoRespJobsInnerJobResourcesNodes) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetSelectType

`func (o *V0041OpenapiJobInfoRespJobsInnerJobResourcesNodes) GetSelectType() []string`

GetSelectType returns the SelectType field if non-nil, zero value otherwise.

### GetSelectTypeOk

`func (o *V0041OpenapiJobInfoRespJobsInnerJobResourcesNodes) GetSelectTypeOk() (*[]string, bool)`

GetSelectTypeOk returns a tuple with the SelectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectType

`func (o *V0041OpenapiJobInfoRespJobsInnerJobResourcesNodes) SetSelectType(v []string)`

SetSelectType sets SelectType field to given value.

### HasSelectType

`func (o *V0041OpenapiJobInfoRespJobsInnerJobResourcesNodes) HasSelectType() bool`

HasSelectType returns a boolean if a field has been set.

### GetList

`func (o *V0041OpenapiJobInfoRespJobsInnerJobResourcesNodes) GetList() string`

GetList returns the List field if non-nil, zero value otherwise.

### GetListOk

`func (o *V0041OpenapiJobInfoRespJobsInnerJobResourcesNodes) GetListOk() (*string, bool)`

GetListOk returns a tuple with the List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetList

`func (o *V0041OpenapiJobInfoRespJobsInnerJobResourcesNodes) SetList(v string)`

SetList sets List field to given value.

### HasList

`func (o *V0041OpenapiJobInfoRespJobsInnerJobResourcesNodes) HasList() bool`

HasList returns a boolean if a field has been set.

### GetWhole

`func (o *V0041OpenapiJobInfoRespJobsInnerJobResourcesNodes) GetWhole() bool`

GetWhole returns the Whole field if non-nil, zero value otherwise.

### GetWholeOk

`func (o *V0041OpenapiJobInfoRespJobsInnerJobResourcesNodes) GetWholeOk() (*bool, bool)`

GetWholeOk returns a tuple with the Whole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhole

`func (o *V0041OpenapiJobInfoRespJobsInnerJobResourcesNodes) SetWhole(v bool)`

SetWhole sets Whole field to given value.

### HasWhole

`func (o *V0041OpenapiJobInfoRespJobsInnerJobResourcesNodes) HasWhole() bool`

HasWhole returns a boolean if a field has been set.

### GetAllocation

`func (o *V0041OpenapiJobInfoRespJobsInnerJobResourcesNodes) GetAllocation() []V0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInner`

GetAllocation returns the Allocation field if non-nil, zero value otherwise.

### GetAllocationOk

`func (o *V0041OpenapiJobInfoRespJobsInnerJobResourcesNodes) GetAllocationOk() (*[]V0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInner, bool)`

GetAllocationOk returns a tuple with the Allocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocation

`func (o *V0041OpenapiJobInfoRespJobsInnerJobResourcesNodes) SetAllocation(v []V0041OpenapiJobInfoRespJobsInnerJobResourcesNodesAllocationInner)`

SetAllocation sets Allocation field to given value.

### HasAllocation

`func (o *V0041OpenapiJobInfoRespJobsInnerJobResourcesNodes) HasAllocation() bool`

HasAllocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


