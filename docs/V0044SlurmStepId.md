# V0044SlurmStepId

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sluid** | Pointer to **string** | SLUID (Slurm Lexicographically-sortable Unique ID) | [optional] 
**JobId** | Pointer to [**V0044Uint32NoValStruct**](V0044Uint32NoValStruct.md) |  | [optional] 
**StepHetComponent** | Pointer to [**V0044Uint32NoValStruct**](V0044Uint32NoValStruct.md) |  | [optional] 
**StepId** | Pointer to **string** | Job step ID | [optional] 

## Methods

### NewV0044SlurmStepId

`func NewV0044SlurmStepId() *V0044SlurmStepId`

NewV0044SlurmStepId instantiates a new V0044SlurmStepId object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0044SlurmStepIdWithDefaults

`func NewV0044SlurmStepIdWithDefaults() *V0044SlurmStepId`

NewV0044SlurmStepIdWithDefaults instantiates a new V0044SlurmStepId object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSluid

`func (o *V0044SlurmStepId) GetSluid() string`

GetSluid returns the Sluid field if non-nil, zero value otherwise.

### GetSluidOk

`func (o *V0044SlurmStepId) GetSluidOk() (*string, bool)`

GetSluidOk returns a tuple with the Sluid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSluid

`func (o *V0044SlurmStepId) SetSluid(v string)`

SetSluid sets Sluid field to given value.

### HasSluid

`func (o *V0044SlurmStepId) HasSluid() bool`

HasSluid returns a boolean if a field has been set.

### GetJobId

`func (o *V0044SlurmStepId) GetJobId() V0044Uint32NoValStruct`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *V0044SlurmStepId) GetJobIdOk() (*V0044Uint32NoValStruct, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *V0044SlurmStepId) SetJobId(v V0044Uint32NoValStruct)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *V0044SlurmStepId) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetStepHetComponent

`func (o *V0044SlurmStepId) GetStepHetComponent() V0044Uint32NoValStruct`

GetStepHetComponent returns the StepHetComponent field if non-nil, zero value otherwise.

### GetStepHetComponentOk

`func (o *V0044SlurmStepId) GetStepHetComponentOk() (*V0044Uint32NoValStruct, bool)`

GetStepHetComponentOk returns a tuple with the StepHetComponent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepHetComponent

`func (o *V0044SlurmStepId) SetStepHetComponent(v V0044Uint32NoValStruct)`

SetStepHetComponent sets StepHetComponent field to given value.

### HasStepHetComponent

`func (o *V0044SlurmStepId) HasStepHetComponent() bool`

HasStepHetComponent returns a boolean if a field has been set.

### GetStepId

`func (o *V0044SlurmStepId) GetStepId() string`

GetStepId returns the StepId field if non-nil, zero value otherwise.

### GetStepIdOk

`func (o *V0044SlurmStepId) GetStepIdOk() (*string, bool)`

GetStepIdOk returns a tuple with the StepId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepId

`func (o *V0044SlurmStepId) SetStepId(v string)`

SetStepId sets StepId field to given value.

### HasStepId

`func (o *V0044SlurmStepId) HasStepId() bool`

HasStepId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


