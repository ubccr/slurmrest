# V0044PartitionInfoNodes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedAllocation** | Pointer to **string** | AllocNodes - Comma-separated list of nodes from which users can submit jobs in the partition | [optional] 
**Configured** | Pointer to **string** | Nodes - Comma-separated list of nodes which are associated with this partition | [optional] 
**Total** | Pointer to **int32** | TotalNodes - Number of nodes available in this partition | [optional] 

## Methods

### NewV0044PartitionInfoNodes

`func NewV0044PartitionInfoNodes() *V0044PartitionInfoNodes`

NewV0044PartitionInfoNodes instantiates a new V0044PartitionInfoNodes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0044PartitionInfoNodesWithDefaults

`func NewV0044PartitionInfoNodesWithDefaults() *V0044PartitionInfoNodes`

NewV0044PartitionInfoNodesWithDefaults instantiates a new V0044PartitionInfoNodes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedAllocation

`func (o *V0044PartitionInfoNodes) GetAllowedAllocation() string`

GetAllowedAllocation returns the AllowedAllocation field if non-nil, zero value otherwise.

### GetAllowedAllocationOk

`func (o *V0044PartitionInfoNodes) GetAllowedAllocationOk() (*string, bool)`

GetAllowedAllocationOk returns a tuple with the AllowedAllocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedAllocation

`func (o *V0044PartitionInfoNodes) SetAllowedAllocation(v string)`

SetAllowedAllocation sets AllowedAllocation field to given value.

### HasAllowedAllocation

`func (o *V0044PartitionInfoNodes) HasAllowedAllocation() bool`

HasAllowedAllocation returns a boolean if a field has been set.

### GetConfigured

`func (o *V0044PartitionInfoNodes) GetConfigured() string`

GetConfigured returns the Configured field if non-nil, zero value otherwise.

### GetConfiguredOk

`func (o *V0044PartitionInfoNodes) GetConfiguredOk() (*string, bool)`

GetConfiguredOk returns a tuple with the Configured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigured

`func (o *V0044PartitionInfoNodes) SetConfigured(v string)`

SetConfigured sets Configured field to given value.

### HasConfigured

`func (o *V0044PartitionInfoNodes) HasConfigured() bool`

HasConfigured returns a boolean if a field has been set.

### GetTotal

`func (o *V0044PartitionInfoNodes) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *V0044PartitionInfoNodes) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *V0044PartitionInfoNodes) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *V0044PartitionInfoNodes) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


