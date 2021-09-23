# V0037Ping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hostname** | Pointer to **string** | slurm controller hostname | [optional] 
**Ping** | Pointer to **string** | slurm controller host up | [optional] 
**Mode** | Pointer to **string** | slurm controller mode | [optional] 
**Status** | Pointer to **int32** | slurm controller status | [optional] 

## Methods

### NewV0037Ping

`func NewV0037Ping() *V0037Ping`

NewV0037Ping instantiates a new V0037Ping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0037PingWithDefaults

`func NewV0037PingWithDefaults() *V0037Ping`

NewV0037PingWithDefaults instantiates a new V0037Ping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostname

`func (o *V0037Ping) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *V0037Ping) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *V0037Ping) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *V0037Ping) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetPing

`func (o *V0037Ping) GetPing() string`

GetPing returns the Ping field if non-nil, zero value otherwise.

### GetPingOk

`func (o *V0037Ping) GetPingOk() (*string, bool)`

GetPingOk returns a tuple with the Ping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPing

`func (o *V0037Ping) SetPing(v string)`

SetPing sets Ping field to given value.

### HasPing

`func (o *V0037Ping) HasPing() bool`

HasPing returns a boolean if a field has been set.

### GetMode

`func (o *V0037Ping) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *V0037Ping) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *V0037Ping) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *V0037Ping) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetStatus

`func (o *V0037Ping) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *V0037Ping) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *V0037Ping) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *V0037Ping) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


