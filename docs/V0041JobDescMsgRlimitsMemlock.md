# V0041JobDescMsgRlimitsMemlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Set** | Pointer to **bool** | True if number has been set. False if number is unset | [optional] 
**Infinite** | Pointer to **bool** | True if number has been set to infinite. \&quot;set\&quot; and \&quot;number\&quot; will be ignored. | [optional] 
**Number** | Pointer to **int64** | If set is True the number will be set with value. Otherwise ignore number contents. | [optional] 

## Methods

### NewV0041JobDescMsgRlimitsMemlock

`func NewV0041JobDescMsgRlimitsMemlock() *V0041JobDescMsgRlimitsMemlock`

NewV0041JobDescMsgRlimitsMemlock instantiates a new V0041JobDescMsgRlimitsMemlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0041JobDescMsgRlimitsMemlockWithDefaults

`func NewV0041JobDescMsgRlimitsMemlockWithDefaults() *V0041JobDescMsgRlimitsMemlock`

NewV0041JobDescMsgRlimitsMemlockWithDefaults instantiates a new V0041JobDescMsgRlimitsMemlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSet

`func (o *V0041JobDescMsgRlimitsMemlock) GetSet() bool`

GetSet returns the Set field if non-nil, zero value otherwise.

### GetSetOk

`func (o *V0041JobDescMsgRlimitsMemlock) GetSetOk() (*bool, bool)`

GetSetOk returns a tuple with the Set field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSet

`func (o *V0041JobDescMsgRlimitsMemlock) SetSet(v bool)`

SetSet sets Set field to given value.

### HasSet

`func (o *V0041JobDescMsgRlimitsMemlock) HasSet() bool`

HasSet returns a boolean if a field has been set.

### GetInfinite

`func (o *V0041JobDescMsgRlimitsMemlock) GetInfinite() bool`

GetInfinite returns the Infinite field if non-nil, zero value otherwise.

### GetInfiniteOk

`func (o *V0041JobDescMsgRlimitsMemlock) GetInfiniteOk() (*bool, bool)`

GetInfiniteOk returns a tuple with the Infinite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfinite

`func (o *V0041JobDescMsgRlimitsMemlock) SetInfinite(v bool)`

SetInfinite sets Infinite field to given value.

### HasInfinite

`func (o *V0041JobDescMsgRlimitsMemlock) HasInfinite() bool`

HasInfinite returns a boolean if a field has been set.

### GetNumber

`func (o *V0041JobDescMsgRlimitsMemlock) GetNumber() int64`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *V0041JobDescMsgRlimitsMemlock) GetNumberOk() (*int64, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *V0041JobDescMsgRlimitsMemlock) SetNumber(v int64)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *V0041JobDescMsgRlimitsMemlock) HasNumber() bool`

HasNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


