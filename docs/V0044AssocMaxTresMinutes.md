# V0044AssocMaxTresMinutes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to [**[]V0044Tres**](V0044Tres.md) |  | [optional] 
**Per** | Pointer to [**V0044QosLimitsMinTresPer**](V0044QosLimitsMinTresPer.md) |  | [optional] 

## Methods

### NewV0044AssocMaxTresMinutes

`func NewV0044AssocMaxTresMinutes() *V0044AssocMaxTresMinutes`

NewV0044AssocMaxTresMinutes instantiates a new V0044AssocMaxTresMinutes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0044AssocMaxTresMinutesWithDefaults

`func NewV0044AssocMaxTresMinutesWithDefaults() *V0044AssocMaxTresMinutes`

NewV0044AssocMaxTresMinutesWithDefaults instantiates a new V0044AssocMaxTresMinutes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *V0044AssocMaxTresMinutes) GetTotal() []V0044Tres`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *V0044AssocMaxTresMinutes) GetTotalOk() (*[]V0044Tres, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *V0044AssocMaxTresMinutes) SetTotal(v []V0044Tres)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *V0044AssocMaxTresMinutes) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetPer

`func (o *V0044AssocMaxTresMinutes) GetPer() V0044QosLimitsMinTresPer`

GetPer returns the Per field if non-nil, zero value otherwise.

### GetPerOk

`func (o *V0044AssocMaxTresMinutes) GetPerOk() (*V0044QosLimitsMinTresPer, bool)`

GetPerOk returns a tuple with the Per field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPer

`func (o *V0044AssocMaxTresMinutes) SetPer(v V0044QosLimitsMinTresPer)`

SetPer sets Per field to given value.

### HasPer

`func (o *V0044AssocMaxTresMinutes) HasPer() bool`

HasPer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


