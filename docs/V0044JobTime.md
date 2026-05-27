# V0044JobTime

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Elapsed** | Pointer to **int32** | Elapsed time in seconds | [optional] 
**Eligible** | Pointer to **int64** | Time when the job became eligible to run (UNIX timestamp) (UNIX timestamp or time string recognized by Slurm (e.g., &#39;[MM/DD[/YY]-]HH:MM[:SS]&#39;)) | [optional] 
**End** | Pointer to **int64** | End time (UNIX timestamp) (UNIX timestamp or time string recognized by Slurm (e.g., &#39;[MM/DD[/YY]-]HH:MM[:SS]&#39;)) | [optional] 
**Planned** | Pointer to [**V0044Uint64NoValStruct**](V0044Uint64NoValStruct.md) |  | [optional] 
**Start** | Pointer to **int64** | Time execution began (UNIX timestamp) (UNIX timestamp or time string recognized by Slurm (e.g., &#39;[MM/DD[/YY]-]HH:MM[:SS]&#39;)) | [optional] 
**Submission** | Pointer to **int64** | Time when the job was submitted (UNIX timestamp) (UNIX timestamp or time string recognized by Slurm (e.g., &#39;[MM/DD[/YY]-]HH:MM[:SS]&#39;)) | [optional] 
**Suspended** | Pointer to **int32** | Total time in suspended state in seconds | [optional] 
**System** | Pointer to [**V0044JobTimeSystem**](V0044JobTimeSystem.md) |  | [optional] 
**Limit** | Pointer to [**V0044Uint32NoValStruct**](V0044Uint32NoValStruct.md) |  | [optional] 
**Total** | Pointer to [**V0044JobTimeTotal**](V0044JobTimeTotal.md) |  | [optional] 
**User** | Pointer to [**V0044JobTimeUser**](V0044JobTimeUser.md) |  | [optional] 

## Methods

### NewV0044JobTime

`func NewV0044JobTime() *V0044JobTime`

NewV0044JobTime instantiates a new V0044JobTime object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0044JobTimeWithDefaults

`func NewV0044JobTimeWithDefaults() *V0044JobTime`

NewV0044JobTimeWithDefaults instantiates a new V0044JobTime object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetElapsed

`func (o *V0044JobTime) GetElapsed() int32`

GetElapsed returns the Elapsed field if non-nil, zero value otherwise.

### GetElapsedOk

`func (o *V0044JobTime) GetElapsedOk() (*int32, bool)`

GetElapsedOk returns a tuple with the Elapsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElapsed

`func (o *V0044JobTime) SetElapsed(v int32)`

SetElapsed sets Elapsed field to given value.

### HasElapsed

`func (o *V0044JobTime) HasElapsed() bool`

HasElapsed returns a boolean if a field has been set.

### GetEligible

`func (o *V0044JobTime) GetEligible() int64`

GetEligible returns the Eligible field if non-nil, zero value otherwise.

### GetEligibleOk

`func (o *V0044JobTime) GetEligibleOk() (*int64, bool)`

GetEligibleOk returns a tuple with the Eligible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEligible

`func (o *V0044JobTime) SetEligible(v int64)`

SetEligible sets Eligible field to given value.

### HasEligible

`func (o *V0044JobTime) HasEligible() bool`

HasEligible returns a boolean if a field has been set.

### GetEnd

`func (o *V0044JobTime) GetEnd() int64`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *V0044JobTime) GetEndOk() (*int64, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *V0044JobTime) SetEnd(v int64)`

SetEnd sets End field to given value.

### HasEnd

`func (o *V0044JobTime) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetPlanned

`func (o *V0044JobTime) GetPlanned() V0044Uint64NoValStruct`

GetPlanned returns the Planned field if non-nil, zero value otherwise.

### GetPlannedOk

`func (o *V0044JobTime) GetPlannedOk() (*V0044Uint64NoValStruct, bool)`

GetPlannedOk returns a tuple with the Planned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanned

`func (o *V0044JobTime) SetPlanned(v V0044Uint64NoValStruct)`

SetPlanned sets Planned field to given value.

### HasPlanned

`func (o *V0044JobTime) HasPlanned() bool`

HasPlanned returns a boolean if a field has been set.

### GetStart

`func (o *V0044JobTime) GetStart() int64`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *V0044JobTime) GetStartOk() (*int64, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *V0044JobTime) SetStart(v int64)`

SetStart sets Start field to given value.

### HasStart

`func (o *V0044JobTime) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetSubmission

`func (o *V0044JobTime) GetSubmission() int64`

GetSubmission returns the Submission field if non-nil, zero value otherwise.

### GetSubmissionOk

`func (o *V0044JobTime) GetSubmissionOk() (*int64, bool)`

GetSubmissionOk returns a tuple with the Submission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmission

`func (o *V0044JobTime) SetSubmission(v int64)`

SetSubmission sets Submission field to given value.

### HasSubmission

`func (o *V0044JobTime) HasSubmission() bool`

HasSubmission returns a boolean if a field has been set.

### GetSuspended

`func (o *V0044JobTime) GetSuspended() int32`

GetSuspended returns the Suspended field if non-nil, zero value otherwise.

### GetSuspendedOk

`func (o *V0044JobTime) GetSuspendedOk() (*int32, bool)`

GetSuspendedOk returns a tuple with the Suspended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspended

`func (o *V0044JobTime) SetSuspended(v int32)`

SetSuspended sets Suspended field to given value.

### HasSuspended

`func (o *V0044JobTime) HasSuspended() bool`

HasSuspended returns a boolean if a field has been set.

### GetSystem

`func (o *V0044JobTime) GetSystem() V0044JobTimeSystem`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *V0044JobTime) GetSystemOk() (*V0044JobTimeSystem, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *V0044JobTime) SetSystem(v V0044JobTimeSystem)`

SetSystem sets System field to given value.

### HasSystem

`func (o *V0044JobTime) HasSystem() bool`

HasSystem returns a boolean if a field has been set.

### GetLimit

`func (o *V0044JobTime) GetLimit() V0044Uint32NoValStruct`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *V0044JobTime) GetLimitOk() (*V0044Uint32NoValStruct, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *V0044JobTime) SetLimit(v V0044Uint32NoValStruct)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *V0044JobTime) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetTotal

`func (o *V0044JobTime) GetTotal() V0044JobTimeTotal`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *V0044JobTime) GetTotalOk() (*V0044JobTimeTotal, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *V0044JobTime) SetTotal(v V0044JobTimeTotal)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *V0044JobTime) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetUser

`func (o *V0044JobTime) GetUser() V0044JobTimeUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *V0044JobTime) GetUserOk() (*V0044JobTimeUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *V0044JobTime) SetUser(v V0044JobTimeUser)`

SetUser sets User field to given value.

### HasUser

`func (o *V0044JobTime) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


