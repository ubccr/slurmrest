# V0044JobState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Current** | Pointer to **[]string** | Current state | [optional] 
**Reason** | Pointer to **string** | Reason for previous Pending or Failed state | [optional] 

## Methods

### NewV0044JobState

`func NewV0044JobState() *V0044JobState`

NewV0044JobState instantiates a new V0044JobState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0044JobStateWithDefaults

`func NewV0044JobStateWithDefaults() *V0044JobState`

NewV0044JobStateWithDefaults instantiates a new V0044JobState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrent

`func (o *V0044JobState) GetCurrent() []string`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *V0044JobState) GetCurrentOk() (*[]string, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *V0044JobState) SetCurrent(v []string)`

SetCurrent sets Current field to given value.

### HasCurrent

`func (o *V0044JobState) HasCurrent() bool`

HasCurrent returns a boolean if a field has been set.

### GetReason

`func (o *V0044JobState) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *V0044JobState) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *V0044JobState) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *V0044JobState) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


