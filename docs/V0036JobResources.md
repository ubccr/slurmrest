# V0036JobResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nodes** | Pointer to **string** | list of assigned job nodes | [optional] 
**AllocatedCpus** | Pointer to **int32** | number of assigned job cpus | [optional] 
**AllocatedHosts** | Pointer to **int32** | number of assigned job hosts | [optional] 
**AllocatedNodes** | Pointer to [**map[string]V0036NodeAllocation**](V0036NodeAllocation.md) | node allocations | [optional] 

## Methods

### NewV0036JobResources

`func NewV0036JobResources() *V0036JobResources`

NewV0036JobResources instantiates a new V0036JobResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0036JobResourcesWithDefaults

`func NewV0036JobResourcesWithDefaults() *V0036JobResources`

NewV0036JobResourcesWithDefaults instantiates a new V0036JobResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodes

`func (o *V0036JobResources) GetNodes() string`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *V0036JobResources) GetNodesOk() (*string, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *V0036JobResources) SetNodes(v string)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *V0036JobResources) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetAllocatedCpus

`func (o *V0036JobResources) GetAllocatedCpus() int32`

GetAllocatedCpus returns the AllocatedCpus field if non-nil, zero value otherwise.

### GetAllocatedCpusOk

`func (o *V0036JobResources) GetAllocatedCpusOk() (*int32, bool)`

GetAllocatedCpusOk returns a tuple with the AllocatedCpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatedCpus

`func (o *V0036JobResources) SetAllocatedCpus(v int32)`

SetAllocatedCpus sets AllocatedCpus field to given value.

### HasAllocatedCpus

`func (o *V0036JobResources) HasAllocatedCpus() bool`

HasAllocatedCpus returns a boolean if a field has been set.

### GetAllocatedHosts

`func (o *V0036JobResources) GetAllocatedHosts() int32`

GetAllocatedHosts returns the AllocatedHosts field if non-nil, zero value otherwise.

### GetAllocatedHostsOk

`func (o *V0036JobResources) GetAllocatedHostsOk() (*int32, bool)`

GetAllocatedHostsOk returns a tuple with the AllocatedHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatedHosts

`func (o *V0036JobResources) SetAllocatedHosts(v int32)`

SetAllocatedHosts sets AllocatedHosts field to given value.

### HasAllocatedHosts

`func (o *V0036JobResources) HasAllocatedHosts() bool`

HasAllocatedHosts returns a boolean if a field has been set.

### GetAllocatedNodes

`func (o *V0036JobResources) GetAllocatedNodes() map[string]V0036NodeAllocation`

GetAllocatedNodes returns the AllocatedNodes field if non-nil, zero value otherwise.

### GetAllocatedNodesOk

`func (o *V0036JobResources) GetAllocatedNodesOk() (*map[string]V0036NodeAllocation, bool)`

GetAllocatedNodesOk returns a tuple with the AllocatedNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatedNodes

`func (o *V0036JobResources) SetAllocatedNodes(v map[string]V0036NodeAllocation)`

SetAllocatedNodes sets AllocatedNodes field to given value.

### HasAllocatedNodes

`func (o *V0036JobResources) HasAllocatedNodes() bool`

HasAllocatedNodes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


