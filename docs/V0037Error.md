# V0037Error

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** | error message | [optional] 
**Errno** | Pointer to **int32** | error number | [optional] 

## Methods

### NewV0037Error

`func NewV0037Error() *V0037Error`

NewV0037Error instantiates a new V0037Error object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0037ErrorWithDefaults

`func NewV0037ErrorWithDefaults() *V0037Error`

NewV0037ErrorWithDefaults instantiates a new V0037Error object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *V0037Error) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *V0037Error) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *V0037Error) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *V0037Error) HasError() bool`

HasError returns a boolean if a field has been set.

### GetErrno

`func (o *V0037Error) GetErrno() int32`

GetErrno returns the Errno field if non-nil, zero value otherwise.

### GetErrnoOk

`func (o *V0037Error) GetErrnoOk() (*int32, bool)`

GetErrnoOk returns a tuple with the Errno field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrno

`func (o *V0037Error) SetErrno(v int32)`

SetErrno sets Errno field to given value.

### HasErrno

`func (o *V0037Error) HasErrno() bool`

HasErrno returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


