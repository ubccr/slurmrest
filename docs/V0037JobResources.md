# V0037JobResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nodes** | Pointer to **string** | list of assigned job nodes | [optional] 
**AllocatedCpus** | Pointer to **int32** | number of assigned job cpus | [optional] 
**AllocatedHosts** | Pointer to **int32** | number of assigned job hosts | [optional] 
**AllocatedNodes** | Pointer to [**[]V0037NodeAllocation**](V0037NodeAllocation.md) | node allocations | [optional] 

## Methods

### NewV0037JobResources

`func NewV0037JobResources() *V0037JobResources`

NewV0037JobResources instantiates a new V0037JobResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0037JobResourcesWithDefaults

`func NewV0037JobResourcesWithDefaults() *V0037JobResources`

NewV0037JobResourcesWithDefaults instantiates a new V0037JobResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodes

`func (o *V0037JobResources) GetNodes() string`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *V0037JobResources) GetNodesOk() (*string, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *V0037JobResources) SetNodes(v string)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *V0037JobResources) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetAllocatedCpus

`func (o *V0037JobResources) GetAllocatedCpus() int32`

GetAllocatedCpus returns the AllocatedCpus field if non-nil, zero value otherwise.

### GetAllocatedCpusOk

`func (o *V0037JobResources) GetAllocatedCpusOk() (*int32, bool)`

GetAllocatedCpusOk returns a tuple with the AllocatedCpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatedCpus

`func (o *V0037JobResources) SetAllocatedCpus(v int32)`

SetAllocatedCpus sets AllocatedCpus field to given value.

### HasAllocatedCpus

`func (o *V0037JobResources) HasAllocatedCpus() bool`

HasAllocatedCpus returns a boolean if a field has been set.

### GetAllocatedHosts

`func (o *V0037JobResources) GetAllocatedHosts() int32`

GetAllocatedHosts returns the AllocatedHosts field if non-nil, zero value otherwise.

### GetAllocatedHostsOk

`func (o *V0037JobResources) GetAllocatedHostsOk() (*int32, bool)`

GetAllocatedHostsOk returns a tuple with the AllocatedHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatedHosts

`func (o *V0037JobResources) SetAllocatedHosts(v int32)`

SetAllocatedHosts sets AllocatedHosts field to given value.

### HasAllocatedHosts

`func (o *V0037JobResources) HasAllocatedHosts() bool`

HasAllocatedHosts returns a boolean if a field has been set.

### GetAllocatedNodes

`func (o *V0037JobResources) GetAllocatedNodes() []V0037NodeAllocation`

GetAllocatedNodes returns the AllocatedNodes field if non-nil, zero value otherwise.

### GetAllocatedNodesOk

`func (o *V0037JobResources) GetAllocatedNodesOk() (*[]V0037NodeAllocation, bool)`

GetAllocatedNodesOk returns a tuple with the AllocatedNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatedNodes

`func (o *V0037JobResources) SetAllocatedNodes(v []V0037NodeAllocation)`

SetAllocatedNodes sets AllocatedNodes field to given value.

### HasAllocatedNodes

`func (o *V0037JobResources) HasAllocatedNodes() bool`

HasAllocatedNodes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


