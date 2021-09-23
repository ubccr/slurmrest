# V0037ReservationsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to [**[]V0037Error**](V0037Error.md) | slurm errors | [optional] 
**Reservations** | Pointer to [**[]V0037Reservation**](V0037Reservation.md) | reservation info | [optional] 

## Methods

### NewV0037ReservationsResponse

`func NewV0037ReservationsResponse() *V0037ReservationsResponse`

NewV0037ReservationsResponse instantiates a new V0037ReservationsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0037ReservationsResponseWithDefaults

`func NewV0037ReservationsResponseWithDefaults() *V0037ReservationsResponse`

NewV0037ReservationsResponseWithDefaults instantiates a new V0037ReservationsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *V0037ReservationsResponse) GetErrors() []V0037Error`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V0037ReservationsResponse) GetErrorsOk() (*[]V0037Error, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V0037ReservationsResponse) SetErrors(v []V0037Error)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V0037ReservationsResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetReservations

`func (o *V0037ReservationsResponse) GetReservations() []V0037Reservation`

GetReservations returns the Reservations field if non-nil, zero value otherwise.

### GetReservationsOk

`func (o *V0037ReservationsResponse) GetReservationsOk() (*[]V0037Reservation, bool)`

GetReservationsOk returns a tuple with the Reservations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservations

`func (o *V0037ReservationsResponse) SetReservations(v []V0037Reservation)`

SetReservations sets Reservations field to given value.

### HasReservations

`func (o *V0037ReservationsResponse) HasReservations() bool`

HasReservations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


