# V0044StepTime

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Elapsed** | Pointer to **int32** | Elapsed time in seconds | [optional] 
**End** | Pointer to [**V0044Uint64NoValStruct**](V0044Uint64NoValStruct.md) |  | [optional] 
**Start** | Pointer to [**V0044Uint64NoValStruct**](V0044Uint64NoValStruct.md) |  | [optional] 
**Suspended** | Pointer to **int32** | Total time in suspended state in seconds | [optional] 
**System** | Pointer to [**V0044StepTimeSystem**](V0044StepTimeSystem.md) |  | [optional] 
**Limit** | Pointer to [**V0044Uint32NoValStruct**](V0044Uint32NoValStruct.md) |  | [optional] 
**Total** | Pointer to [**V0044StepTimeTotal**](V0044StepTimeTotal.md) |  | [optional] 
**User** | Pointer to [**V0044StepTimeUser**](V0044StepTimeUser.md) |  | [optional] 

## Methods

### NewV0044StepTime

`func NewV0044StepTime() *V0044StepTime`

NewV0044StepTime instantiates a new V0044StepTime object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0044StepTimeWithDefaults

`func NewV0044StepTimeWithDefaults() *V0044StepTime`

NewV0044StepTimeWithDefaults instantiates a new V0044StepTime object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetElapsed

`func (o *V0044StepTime) GetElapsed() int32`

GetElapsed returns the Elapsed field if non-nil, zero value otherwise.

### GetElapsedOk

`func (o *V0044StepTime) GetElapsedOk() (*int32, bool)`

GetElapsedOk returns a tuple with the Elapsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElapsed

`func (o *V0044StepTime) SetElapsed(v int32)`

SetElapsed sets Elapsed field to given value.

### HasElapsed

`func (o *V0044StepTime) HasElapsed() bool`

HasElapsed returns a boolean if a field has been set.

### GetEnd

`func (o *V0044StepTime) GetEnd() V0044Uint64NoValStruct`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *V0044StepTime) GetEndOk() (*V0044Uint64NoValStruct, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *V0044StepTime) SetEnd(v V0044Uint64NoValStruct)`

SetEnd sets End field to given value.

### HasEnd

`func (o *V0044StepTime) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetStart

`func (o *V0044StepTime) GetStart() V0044Uint64NoValStruct`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *V0044StepTime) GetStartOk() (*V0044Uint64NoValStruct, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *V0044StepTime) SetStart(v V0044Uint64NoValStruct)`

SetStart sets Start field to given value.

### HasStart

`func (o *V0044StepTime) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetSuspended

`func (o *V0044StepTime) GetSuspended() int32`

GetSuspended returns the Suspended field if non-nil, zero value otherwise.

### GetSuspendedOk

`func (o *V0044StepTime) GetSuspendedOk() (*int32, bool)`

GetSuspendedOk returns a tuple with the Suspended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspended

`func (o *V0044StepTime) SetSuspended(v int32)`

SetSuspended sets Suspended field to given value.

### HasSuspended

`func (o *V0044StepTime) HasSuspended() bool`

HasSuspended returns a boolean if a field has been set.

### GetSystem

`func (o *V0044StepTime) GetSystem() V0044StepTimeSystem`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *V0044StepTime) GetSystemOk() (*V0044StepTimeSystem, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *V0044StepTime) SetSystem(v V0044StepTimeSystem)`

SetSystem sets System field to given value.

### HasSystem

`func (o *V0044StepTime) HasSystem() bool`

HasSystem returns a boolean if a field has been set.

### GetLimit

`func (o *V0044StepTime) GetLimit() V0044Uint32NoValStruct`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *V0044StepTime) GetLimitOk() (*V0044Uint32NoValStruct, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *V0044StepTime) SetLimit(v V0044Uint32NoValStruct)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *V0044StepTime) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetTotal

`func (o *V0044StepTime) GetTotal() V0044StepTimeTotal`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *V0044StepTime) GetTotalOk() (*V0044StepTimeTotal, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *V0044StepTime) SetTotal(v V0044StepTimeTotal)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *V0044StepTime) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetUser

`func (o *V0044StepTime) GetUser() V0044StepTimeUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *V0044StepTime) GetUserOk() (*V0044StepTimeUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *V0044StepTime) SetUser(v V0044StepTimeUser)`

SetUser sets User field to given value.

### HasUser

`func (o *V0044StepTime) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


