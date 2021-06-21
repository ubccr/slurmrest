# V0036JobSubmissionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to [**[]V0036Error**](V0036Error.md) | slurm errors | [optional] 
**JobId** | Pointer to **int32** | new job ID | [optional] 
**StepId** | Pointer to **string** | new job step ID | [optional] 
**JobSubmitUserMsg** | Pointer to **string** | Message to user from job_submit plugin | [optional] 

## Methods

### NewV0036JobSubmissionResponse

`func NewV0036JobSubmissionResponse() *V0036JobSubmissionResponse`

NewV0036JobSubmissionResponse instantiates a new V0036JobSubmissionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0036JobSubmissionResponseWithDefaults

`func NewV0036JobSubmissionResponseWithDefaults() *V0036JobSubmissionResponse`

NewV0036JobSubmissionResponseWithDefaults instantiates a new V0036JobSubmissionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *V0036JobSubmissionResponse) GetErrors() []V0036Error`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V0036JobSubmissionResponse) GetErrorsOk() (*[]V0036Error, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V0036JobSubmissionResponse) SetErrors(v []V0036Error)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V0036JobSubmissionResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetJobId

`func (o *V0036JobSubmissionResponse) GetJobId() int32`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *V0036JobSubmissionResponse) GetJobIdOk() (*int32, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *V0036JobSubmissionResponse) SetJobId(v int32)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *V0036JobSubmissionResponse) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetStepId

`func (o *V0036JobSubmissionResponse) GetStepId() string`

GetStepId returns the StepId field if non-nil, zero value otherwise.

### GetStepIdOk

`func (o *V0036JobSubmissionResponse) GetStepIdOk() (*string, bool)`

GetStepIdOk returns a tuple with the StepId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepId

`func (o *V0036JobSubmissionResponse) SetStepId(v string)`

SetStepId sets StepId field to given value.

### HasStepId

`func (o *V0036JobSubmissionResponse) HasStepId() bool`

HasStepId returns a boolean if a field has been set.

### GetJobSubmitUserMsg

`func (o *V0036JobSubmissionResponse) GetJobSubmitUserMsg() string`

GetJobSubmitUserMsg returns the JobSubmitUserMsg field if non-nil, zero value otherwise.

### GetJobSubmitUserMsgOk

`func (o *V0036JobSubmissionResponse) GetJobSubmitUserMsgOk() (*string, bool)`

GetJobSubmitUserMsgOk returns a tuple with the JobSubmitUserMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobSubmitUserMsg

`func (o *V0036JobSubmissionResponse) SetJobSubmitUserMsg(v string)`

SetJobSubmitUserMsg sets JobSubmitUserMsg field to given value.

### HasJobSubmitUserMsg

`func (o *V0036JobSubmissionResponse) HasJobSubmitUserMsg() bool`

HasJobSubmitUserMsg returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


