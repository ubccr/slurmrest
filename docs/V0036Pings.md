# V0036Pings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to [**[]V0036Error**](V0036Error.md) | slurm errors | [optional] 
**Pings** | Pointer to [**[]V0036Ping**](V0036Ping.md) | slurm controller pings | [optional] 

## Methods

### NewV0036Pings

`func NewV0036Pings() *V0036Pings`

NewV0036Pings instantiates a new V0036Pings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0036PingsWithDefaults

`func NewV0036PingsWithDefaults() *V0036Pings`

NewV0036PingsWithDefaults instantiates a new V0036Pings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *V0036Pings) GetErrors() []V0036Error`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V0036Pings) GetErrorsOk() (*[]V0036Error, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V0036Pings) SetErrors(v []V0036Error)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V0036Pings) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetPings

`func (o *V0036Pings) GetPings() []V0036Ping`

GetPings returns the Pings field if non-nil, zero value otherwise.

### GetPingsOk

`func (o *V0036Pings) GetPingsOk() (*[]V0036Ping, bool)`

GetPingsOk returns a tuple with the Pings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPings

`func (o *V0036Pings) SetPings(v []V0036Ping)`

SetPings sets Pings field to given value.

### HasPings

`func (o *V0036Pings) HasPings() bool`

HasPings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


