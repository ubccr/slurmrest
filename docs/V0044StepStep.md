# V0044StepStep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Step ID (Slurm job step ID) | [optional] 
**Name** | Pointer to **string** | Step name | [optional] 
**Stderr** | Pointer to **string** | Path to stderr file | [optional] 
**Stdin** | Pointer to **string** | Path to stdin file | [optional] 
**Stdout** | Pointer to **string** | Path to stdout file | [optional] 
**StderrExpanded** | Pointer to **string** | Step stderr with expanded fields | [optional] 
**StdinExpanded** | Pointer to **string** | Step stdin with expanded fields | [optional] 
**StdoutExpanded** | Pointer to **string** | Step stdout with expanded fields | [optional] 

## Methods

### NewV0044StepStep

`func NewV0044StepStep() *V0044StepStep`

NewV0044StepStep instantiates a new V0044StepStep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0044StepStepWithDefaults

`func NewV0044StepStepWithDefaults() *V0044StepStep`

NewV0044StepStepWithDefaults instantiates a new V0044StepStep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *V0044StepStep) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V0044StepStep) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V0044StepStep) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *V0044StepStep) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *V0044StepStep) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V0044StepStep) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V0044StepStep) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V0044StepStep) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStderr

`func (o *V0044StepStep) GetStderr() string`

GetStderr returns the Stderr field if non-nil, zero value otherwise.

### GetStderrOk

`func (o *V0044StepStep) GetStderrOk() (*string, bool)`

GetStderrOk returns a tuple with the Stderr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStderr

`func (o *V0044StepStep) SetStderr(v string)`

SetStderr sets Stderr field to given value.

### HasStderr

`func (o *V0044StepStep) HasStderr() bool`

HasStderr returns a boolean if a field has been set.

### GetStdin

`func (o *V0044StepStep) GetStdin() string`

GetStdin returns the Stdin field if non-nil, zero value otherwise.

### GetStdinOk

`func (o *V0044StepStep) GetStdinOk() (*string, bool)`

GetStdinOk returns a tuple with the Stdin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStdin

`func (o *V0044StepStep) SetStdin(v string)`

SetStdin sets Stdin field to given value.

### HasStdin

`func (o *V0044StepStep) HasStdin() bool`

HasStdin returns a boolean if a field has been set.

### GetStdout

`func (o *V0044StepStep) GetStdout() string`

GetStdout returns the Stdout field if non-nil, zero value otherwise.

### GetStdoutOk

`func (o *V0044StepStep) GetStdoutOk() (*string, bool)`

GetStdoutOk returns a tuple with the Stdout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStdout

`func (o *V0044StepStep) SetStdout(v string)`

SetStdout sets Stdout field to given value.

### HasStdout

`func (o *V0044StepStep) HasStdout() bool`

HasStdout returns a boolean if a field has been set.

### GetStderrExpanded

`func (o *V0044StepStep) GetStderrExpanded() string`

GetStderrExpanded returns the StderrExpanded field if non-nil, zero value otherwise.

### GetStderrExpandedOk

`func (o *V0044StepStep) GetStderrExpandedOk() (*string, bool)`

GetStderrExpandedOk returns a tuple with the StderrExpanded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStderrExpanded

`func (o *V0044StepStep) SetStderrExpanded(v string)`

SetStderrExpanded sets StderrExpanded field to given value.

### HasStderrExpanded

`func (o *V0044StepStep) HasStderrExpanded() bool`

HasStderrExpanded returns a boolean if a field has been set.

### GetStdinExpanded

`func (o *V0044StepStep) GetStdinExpanded() string`

GetStdinExpanded returns the StdinExpanded field if non-nil, zero value otherwise.

### GetStdinExpandedOk

`func (o *V0044StepStep) GetStdinExpandedOk() (*string, bool)`

GetStdinExpandedOk returns a tuple with the StdinExpanded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStdinExpanded

`func (o *V0044StepStep) SetStdinExpanded(v string)`

SetStdinExpanded sets StdinExpanded field to given value.

### HasStdinExpanded

`func (o *V0044StepStep) HasStdinExpanded() bool`

HasStdinExpanded returns a boolean if a field has been set.

### GetStdoutExpanded

`func (o *V0044StepStep) GetStdoutExpanded() string`

GetStdoutExpanded returns the StdoutExpanded field if non-nil, zero value otherwise.

### GetStdoutExpandedOk

`func (o *V0044StepStep) GetStdoutExpandedOk() (*string, bool)`

GetStdoutExpandedOk returns a tuple with the StdoutExpanded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStdoutExpanded

`func (o *V0044StepStep) SetStdoutExpanded(v string)`

SetStdoutExpanded sets StdoutExpanded field to given value.

### HasStdoutExpanded

`func (o *V0044StepStep) HasStdoutExpanded() bool`

HasStdoutExpanded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


