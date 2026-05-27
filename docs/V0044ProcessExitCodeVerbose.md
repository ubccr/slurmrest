# V0044ProcessExitCodeVerbose

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **[]string** | Status given by return code | [optional] 
**ReturnCode** | Pointer to [**V0044Uint32NoValStruct**](V0044Uint32NoValStruct.md) |  | [optional] 
**Signal** | Pointer to [**V0044ProcessExitCodeVerboseSignal**](V0044ProcessExitCodeVerboseSignal.md) |  | [optional] 

## Methods

### NewV0044ProcessExitCodeVerbose

`func NewV0044ProcessExitCodeVerbose() *V0044ProcessExitCodeVerbose`

NewV0044ProcessExitCodeVerbose instantiates a new V0044ProcessExitCodeVerbose object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0044ProcessExitCodeVerboseWithDefaults

`func NewV0044ProcessExitCodeVerboseWithDefaults() *V0044ProcessExitCodeVerbose`

NewV0044ProcessExitCodeVerboseWithDefaults instantiates a new V0044ProcessExitCodeVerbose object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *V0044ProcessExitCodeVerbose) GetStatus() []string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *V0044ProcessExitCodeVerbose) GetStatusOk() (*[]string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *V0044ProcessExitCodeVerbose) SetStatus(v []string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *V0044ProcessExitCodeVerbose) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetReturnCode

`func (o *V0044ProcessExitCodeVerbose) GetReturnCode() V0044Uint32NoValStruct`

GetReturnCode returns the ReturnCode field if non-nil, zero value otherwise.

### GetReturnCodeOk

`func (o *V0044ProcessExitCodeVerbose) GetReturnCodeOk() (*V0044Uint32NoValStruct, bool)`

GetReturnCodeOk returns a tuple with the ReturnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnCode

`func (o *V0044ProcessExitCodeVerbose) SetReturnCode(v V0044Uint32NoValStruct)`

SetReturnCode sets ReturnCode field to given value.

### HasReturnCode

`func (o *V0044ProcessExitCodeVerbose) HasReturnCode() bool`

HasReturnCode returns a boolean if a field has been set.

### GetSignal

`func (o *V0044ProcessExitCodeVerbose) GetSignal() V0044ProcessExitCodeVerboseSignal`

GetSignal returns the Signal field if non-nil, zero value otherwise.

### GetSignalOk

`func (o *V0044ProcessExitCodeVerbose) GetSignalOk() (*V0044ProcessExitCodeVerboseSignal, bool)`

GetSignalOk returns a tuple with the Signal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignal

`func (o *V0044ProcessExitCodeVerbose) SetSignal(v V0044ProcessExitCodeVerboseSignal)`

SetSignal sets Signal field to given value.

### HasSignal

`func (o *V0044ProcessExitCodeVerbose) HasSignal() bool`

HasSignal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


