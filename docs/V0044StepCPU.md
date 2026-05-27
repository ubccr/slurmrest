# V0044StepCPU

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestedFrequency** | Pointer to [**V0044StepCPURequestedFrequency**](V0044StepCPURequestedFrequency.md) |  | [optional] 
**Governor** | Pointer to **string** | Requested CPU frequency governor in kHz | [optional] 

## Methods

### NewV0044StepCPU

`func NewV0044StepCPU() *V0044StepCPU`

NewV0044StepCPU instantiates a new V0044StepCPU object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0044StepCPUWithDefaults

`func NewV0044StepCPUWithDefaults() *V0044StepCPU`

NewV0044StepCPUWithDefaults instantiates a new V0044StepCPU object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestedFrequency

`func (o *V0044StepCPU) GetRequestedFrequency() V0044StepCPURequestedFrequency`

GetRequestedFrequency returns the RequestedFrequency field if non-nil, zero value otherwise.

### GetRequestedFrequencyOk

`func (o *V0044StepCPU) GetRequestedFrequencyOk() (*V0044StepCPURequestedFrequency, bool)`

GetRequestedFrequencyOk returns a tuple with the RequestedFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedFrequency

`func (o *V0044StepCPU) SetRequestedFrequency(v V0044StepCPURequestedFrequency)`

SetRequestedFrequency sets RequestedFrequency field to given value.

### HasRequestedFrequency

`func (o *V0044StepCPU) HasRequestedFrequency() bool`

HasRequestedFrequency returns a boolean if a field has been set.

### GetGovernor

`func (o *V0044StepCPU) GetGovernor() string`

GetGovernor returns the Governor field if non-nil, zero value otherwise.

### GetGovernorOk

`func (o *V0044StepCPU) GetGovernorOk() (*string, bool)`

GetGovernorOk returns a tuple with the Governor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGovernor

`func (o *V0044StepCPU) SetGovernor(v string)`

SetGovernor sets Governor field to given value.

### HasGovernor

`func (o *V0044StepCPU) HasGovernor() bool`

HasGovernor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


