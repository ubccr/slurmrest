# V0037Pings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to [**[]V0037Error**](V0037Error.md) | slurm errors | [optional] 
**Pings** | Pointer to [**[]V0037Ping**](V0037Ping.md) | slurm controller pings | [optional] 

## Methods

### NewV0037Pings

`func NewV0037Pings() *V0037Pings`

NewV0037Pings instantiates a new V0037Pings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0037PingsWithDefaults

`func NewV0037PingsWithDefaults() *V0037Pings`

NewV0037PingsWithDefaults instantiates a new V0037Pings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *V0037Pings) GetErrors() []V0037Error`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V0037Pings) GetErrorsOk() (*[]V0037Error, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V0037Pings) SetErrors(v []V0037Error)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V0037Pings) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetPings

`func (o *V0037Pings) GetPings() []V0037Ping`

GetPings returns the Pings field if non-nil, zero value otherwise.

### GetPingsOk

`func (o *V0037Pings) GetPingsOk() (*[]V0037Ping, bool)`

GetPingsOk returns a tuple with the Pings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPings

`func (o *V0037Pings) SetPings(v []V0037Ping)`

SetPings sets Pings field to given value.

### HasPings

`func (o *V0037Pings) HasPings() bool`

HasPings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


