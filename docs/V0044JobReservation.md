# V0044JobReservation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Unique identifier of requested reservation | [optional] 
**Name** | Pointer to **string** | Name of reservation to use | [optional] 
**Requested** | Pointer to **string** | Comma-separated list of requested reservation names | [optional] 

## Methods

### NewV0044JobReservation

`func NewV0044JobReservation() *V0044JobReservation`

NewV0044JobReservation instantiates a new V0044JobReservation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0044JobReservationWithDefaults

`func NewV0044JobReservationWithDefaults() *V0044JobReservation`

NewV0044JobReservationWithDefaults instantiates a new V0044JobReservation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *V0044JobReservation) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V0044JobReservation) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V0044JobReservation) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *V0044JobReservation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *V0044JobReservation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V0044JobReservation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V0044JobReservation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V0044JobReservation) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRequested

`func (o *V0044JobReservation) GetRequested() string`

GetRequested returns the Requested field if non-nil, zero value otherwise.

### GetRequestedOk

`func (o *V0044JobReservation) GetRequestedOk() (*string, bool)`

GetRequestedOk returns a tuple with the Requested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequested

`func (o *V0044JobReservation) SetRequested(v string)`

SetRequested sets Requested field to given value.

### HasRequested

`func (o *V0044JobReservation) HasRequested() bool`

HasRequested returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


