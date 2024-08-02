# V0041OpenapiSlurmdbdJobsRespJobsInnerTime

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Elapsed** | Pointer to **int32** |  | [optional] 
**Eligible** | Pointer to **int64** |  | [optional] 
**End** | Pointer to **int64** |  | [optional] 
**Planned** | Pointer to [**V0041OpenapiSlurmdbdJobsRespJobsInnerTimePlanned**](V0041OpenapiSlurmdbdJobsRespJobsInnerTimePlanned.md) |  | [optional] 
**Start** | Pointer to **int64** |  | [optional] 
**Submission** | Pointer to **int64** |  | [optional] 
**Suspended** | Pointer to **int32** |  | [optional] 
**System** | Pointer to [**V0041OpenapiSlurmdbdJobsRespJobsInnerTimeSystem**](V0041OpenapiSlurmdbdJobsRespJobsInnerTimeSystem.md) |  | [optional] 
**Limit** | Pointer to [**V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId**](V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId.md) |  | [optional] 
**Total** | Pointer to [**V0041OpenapiSlurmdbdJobsRespJobsInnerTimeSystem**](V0041OpenapiSlurmdbdJobsRespJobsInnerTimeSystem.md) |  | [optional] 
**User** | Pointer to [**V0041OpenapiSlurmdbdJobsRespJobsInnerTimeSystem**](V0041OpenapiSlurmdbdJobsRespJobsInnerTimeSystem.md) |  | [optional] 

## Methods

### NewV0041OpenapiSlurmdbdJobsRespJobsInnerTime

`func NewV0041OpenapiSlurmdbdJobsRespJobsInnerTime() *V0041OpenapiSlurmdbdJobsRespJobsInnerTime`

NewV0041OpenapiSlurmdbdJobsRespJobsInnerTime instantiates a new V0041OpenapiSlurmdbdJobsRespJobsInnerTime object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0041OpenapiSlurmdbdJobsRespJobsInnerTimeWithDefaults

`func NewV0041OpenapiSlurmdbdJobsRespJobsInnerTimeWithDefaults() *V0041OpenapiSlurmdbdJobsRespJobsInnerTime`

NewV0041OpenapiSlurmdbdJobsRespJobsInnerTimeWithDefaults instantiates a new V0041OpenapiSlurmdbdJobsRespJobsInnerTime object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetElapsed

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerTime) GetElapsed() int32`

GetElapsed returns the Elapsed field if non-nil, zero value otherwise.

### GetElapsedOk

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerTime) GetElapsedOk() (*int32, bool)`

GetElapsedOk returns a tuple with the Elapsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElapsed

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerTime) SetElapsed(v int32)`

SetElapsed sets Elapsed field to given value.

### HasElapsed

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerTime) HasElapsed() bool`

HasElapsed returns a boolean if a field has been set.

### GetEligible

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerTime) GetEligible() int64`

GetEligible returns the Eligible field if non-nil, zero value otherwise.

### GetEligibleOk

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerTime) GetEligibleOk() (*int64, bool)`

GetEligibleOk returns a tuple with the Eligible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEligible

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerTime) SetEligible(v int64)`

SetEligible sets Eligible field to given value.

### HasEligible

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerTime) HasEligible() bool`

HasEligible returns a boolean if a field has been set.

### GetEnd

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerTime) GetEnd() int64`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerTime) GetEndOk() (*int64, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerTime) SetEnd(v int64)`

SetEnd sets End field to given value.

### HasEnd

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerTime) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetPlanned

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerTime) GetPlanned() V0041OpenapiSlurmdbdJobsRespJobsInnerTimePlanned`

GetPlanned returns the Planned field if non-nil, zero value otherwise.

### GetPlannedOk

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerTime) GetPlannedOk() (*V0041OpenapiSlurmdbdJobsRespJobsInnerTimePlanned, bool)`

GetPlannedOk returns a tuple with the Planned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanned

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerTime) SetPlanned(v V0041OpenapiSlurmdbdJobsRespJobsInnerTimePlanned)`

SetPlanned sets Planned field to given value.

### HasPlanned

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerTime) HasPlanned() bool`

HasPlanned returns a boolean if a field has been set.

### GetStart

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerTime) GetStart() int64`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerTime) GetStartOk() (*int64, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerTime) SetStart(v int64)`

SetStart sets Start field to given value.

### HasStart

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerTime) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetSubmission

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerTime) GetSubmission() int64`

GetSubmission returns the Submission field if non-nil, zero value otherwise.

### GetSubmissionOk

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerTime) GetSubmissionOk() (*int64, bool)`

GetSubmissionOk returns a tuple with the Submission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmission

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerTime) SetSubmission(v int64)`

SetSubmission sets Submission field to given value.

### HasSubmission

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerTime) HasSubmission() bool`

HasSubmission returns a boolean if a field has been set.

### GetSuspended

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerTime) GetSuspended() int32`

GetSuspended returns the Suspended field if non-nil, zero value otherwise.

### GetSuspendedOk

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerTime) GetSuspendedOk() (*int32, bool)`

GetSuspendedOk returns a tuple with the Suspended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspended

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerTime) SetSuspended(v int32)`

SetSuspended sets Suspended field to given value.

### HasSuspended

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerTime) HasSuspended() bool`

HasSuspended returns a boolean if a field has been set.

### GetSystem

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerTime) GetSystem() V0041OpenapiSlurmdbdJobsRespJobsInnerTimeSystem`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerTime) GetSystemOk() (*V0041OpenapiSlurmdbdJobsRespJobsInnerTimeSystem, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerTime) SetSystem(v V0041OpenapiSlurmdbdJobsRespJobsInnerTimeSystem)`

SetSystem sets System field to given value.

### HasSystem

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerTime) HasSystem() bool`

HasSystem returns a boolean if a field has been set.

### GetLimit

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerTime) GetLimit() V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerTime) GetLimitOk() (*V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerTime) SetLimit(v V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerTime) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetTotal

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerTime) GetTotal() V0041OpenapiSlurmdbdJobsRespJobsInnerTimeSystem`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerTime) GetTotalOk() (*V0041OpenapiSlurmdbdJobsRespJobsInnerTimeSystem, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerTime) SetTotal(v V0041OpenapiSlurmdbdJobsRespJobsInnerTimeSystem)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerTime) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetUser

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerTime) GetUser() V0041OpenapiSlurmdbdJobsRespJobsInnerTimeSystem`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerTime) GetUserOk() (*V0041OpenapiSlurmdbdJobsRespJobsInnerTimeSystem, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerTime) SetUser(v V0041OpenapiSlurmdbdJobsRespJobsInnerTimeSystem)`

SetUser sets User field to given value.

### HasUser

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerTime) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


