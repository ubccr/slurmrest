# V0044QosLimitsMaxTres

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to [**[]V0044Tres**](V0044Tres.md) |  | [optional] 
**Minutes** | Pointer to [**V0044QosLimitsMaxTresMinutes**](V0044QosLimitsMaxTresMinutes.md) |  | [optional] 
**Per** | Pointer to [**V0044QosLimitsMaxTresPer**](V0044QosLimitsMaxTresPer.md) |  | [optional] 

## Methods

### NewV0044QosLimitsMaxTres

`func NewV0044QosLimitsMaxTres() *V0044QosLimitsMaxTres`

NewV0044QosLimitsMaxTres instantiates a new V0044QosLimitsMaxTres object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0044QosLimitsMaxTresWithDefaults

`func NewV0044QosLimitsMaxTresWithDefaults() *V0044QosLimitsMaxTres`

NewV0044QosLimitsMaxTresWithDefaults instantiates a new V0044QosLimitsMaxTres object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *V0044QosLimitsMaxTres) GetTotal() []V0044Tres`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *V0044QosLimitsMaxTres) GetTotalOk() (*[]V0044Tres, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *V0044QosLimitsMaxTres) SetTotal(v []V0044Tres)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *V0044QosLimitsMaxTres) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetMinutes

`func (o *V0044QosLimitsMaxTres) GetMinutes() V0044QosLimitsMaxTresMinutes`

GetMinutes returns the Minutes field if non-nil, zero value otherwise.

### GetMinutesOk

`func (o *V0044QosLimitsMaxTres) GetMinutesOk() (*V0044QosLimitsMaxTresMinutes, bool)`

GetMinutesOk returns a tuple with the Minutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinutes

`func (o *V0044QosLimitsMaxTres) SetMinutes(v V0044QosLimitsMaxTresMinutes)`

SetMinutes sets Minutes field to given value.

### HasMinutes

`func (o *V0044QosLimitsMaxTres) HasMinutes() bool`

HasMinutes returns a boolean if a field has been set.

### GetPer

`func (o *V0044QosLimitsMaxTres) GetPer() V0044QosLimitsMaxTresPer`

GetPer returns the Per field if non-nil, zero value otherwise.

### GetPerOk

`func (o *V0044QosLimitsMaxTres) GetPerOk() (*V0044QosLimitsMaxTresPer, bool)`

GetPerOk returns a tuple with the Per field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPer

`func (o *V0044QosLimitsMaxTres) SetPer(v V0044QosLimitsMaxTresPer)`

SetPer sets Per field to given value.

### HasPer

`func (o *V0044QosLimitsMaxTres) HasPer() bool`

HasPer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


