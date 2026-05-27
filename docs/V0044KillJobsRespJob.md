# V0044KillJobsRespJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to [**V0044KillJobsRespJobError**](V0044KillJobsRespJobError.md) |  | [optional] 
**StepId** | **string** | Job or Step ID that signaling failed | 
**JobId** | [**V0044Uint32NoValStruct**](V0044Uint32NoValStruct.md) |  | 
**Federation** | Pointer to [**V0044KillJobsRespJobFederation**](V0044KillJobsRespJobFederation.md) |  | [optional] 

## Methods

### NewV0044KillJobsRespJob

`func NewV0044KillJobsRespJob(stepId string, jobId V0044Uint32NoValStruct, ) *V0044KillJobsRespJob`

NewV0044KillJobsRespJob instantiates a new V0044KillJobsRespJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0044KillJobsRespJobWithDefaults

`func NewV0044KillJobsRespJobWithDefaults() *V0044KillJobsRespJob`

NewV0044KillJobsRespJobWithDefaults instantiates a new V0044KillJobsRespJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *V0044KillJobsRespJob) GetError() V0044KillJobsRespJobError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *V0044KillJobsRespJob) GetErrorOk() (*V0044KillJobsRespJobError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *V0044KillJobsRespJob) SetError(v V0044KillJobsRespJobError)`

SetError sets Error field to given value.

### HasError

`func (o *V0044KillJobsRespJob) HasError() bool`

HasError returns a boolean if a field has been set.

### GetStepId

`func (o *V0044KillJobsRespJob) GetStepId() string`

GetStepId returns the StepId field if non-nil, zero value otherwise.

### GetStepIdOk

`func (o *V0044KillJobsRespJob) GetStepIdOk() (*string, bool)`

GetStepIdOk returns a tuple with the StepId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepId

`func (o *V0044KillJobsRespJob) SetStepId(v string)`

SetStepId sets StepId field to given value.


### GetJobId

`func (o *V0044KillJobsRespJob) GetJobId() V0044Uint32NoValStruct`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *V0044KillJobsRespJob) GetJobIdOk() (*V0044Uint32NoValStruct, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *V0044KillJobsRespJob) SetJobId(v V0044Uint32NoValStruct)`

SetJobId sets JobId field to given value.


### GetFederation

`func (o *V0044KillJobsRespJob) GetFederation() V0044KillJobsRespJobFederation`

GetFederation returns the Federation field if non-nil, zero value otherwise.

### GetFederationOk

`func (o *V0044KillJobsRespJob) GetFederationOk() (*V0044KillJobsRespJobFederation, bool)`

GetFederationOk returns a tuple with the Federation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederation

`func (o *V0044KillJobsRespJob) SetFederation(v V0044KillJobsRespJobFederation)`

SetFederation sets Federation field to given value.

### HasFederation

`func (o *V0044KillJobsRespJob) HasFederation() bool`

HasFederation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


