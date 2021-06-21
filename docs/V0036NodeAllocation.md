# V0036NodeAllocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Memory** | Pointer to **int32** | amount of assigned job memory | [optional] 
**Cpus** | Pointer to **map[string]interface{}** | amount of assigned job CPUs | [optional] 
**Sockets** | Pointer to **map[string]interface{}** | assignment status of each socket by socket id | [optional] 
**Cores** | Pointer to **map[string]interface{}** | assignment status of each core by core id | [optional] 

## Methods

### NewV0036NodeAllocation

`func NewV0036NodeAllocation() *V0036NodeAllocation`

NewV0036NodeAllocation instantiates a new V0036NodeAllocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0036NodeAllocationWithDefaults

`func NewV0036NodeAllocationWithDefaults() *V0036NodeAllocation`

NewV0036NodeAllocationWithDefaults instantiates a new V0036NodeAllocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMemory

`func (o *V0036NodeAllocation) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *V0036NodeAllocation) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *V0036NodeAllocation) SetMemory(v int32)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *V0036NodeAllocation) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetCpus

`func (o *V0036NodeAllocation) GetCpus() map[string]interface{}`

GetCpus returns the Cpus field if non-nil, zero value otherwise.

### GetCpusOk

`func (o *V0036NodeAllocation) GetCpusOk() (*map[string]interface{}, bool)`

GetCpusOk returns a tuple with the Cpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpus

`func (o *V0036NodeAllocation) SetCpus(v map[string]interface{})`

SetCpus sets Cpus field to given value.

### HasCpus

`func (o *V0036NodeAllocation) HasCpus() bool`

HasCpus returns a boolean if a field has been set.

### GetSockets

`func (o *V0036NodeAllocation) GetSockets() map[string]interface{}`

GetSockets returns the Sockets field if non-nil, zero value otherwise.

### GetSocketsOk

`func (o *V0036NodeAllocation) GetSocketsOk() (*map[string]interface{}, bool)`

GetSocketsOk returns a tuple with the Sockets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSockets

`func (o *V0036NodeAllocation) SetSockets(v map[string]interface{})`

SetSockets sets Sockets field to given value.

### HasSockets

`func (o *V0036NodeAllocation) HasSockets() bool`

HasSockets returns a boolean if a field has been set.

### GetCores

`func (o *V0036NodeAllocation) GetCores() map[string]interface{}`

GetCores returns the Cores field if non-nil, zero value otherwise.

### GetCoresOk

`func (o *V0036NodeAllocation) GetCoresOk() (*map[string]interface{}, bool)`

GetCoresOk returns a tuple with the Cores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCores

`func (o *V0036NodeAllocation) SetCores(v map[string]interface{})`

SetCores sets Cores field to given value.

### HasCores

`func (o *V0036NodeAllocation) HasCores() bool`

HasCores returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


