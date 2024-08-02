# V0040KillJobsRespJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to [**V0041OpenapiKillJobsRespStatusInnerError**](V0041OpenapiKillJobsRespStatusInnerError.md) |  | [optional] 
**StepId** | **string** | Job or Step ID that signaling failed | 
**JobId** | [**V0040Uint32NoVal**](V0040Uint32NoVal.md) |  | 
**Federation** | Pointer to [**V0041OpenapiKillJobsRespStatusInnerFederation**](V0041OpenapiKillJobsRespStatusInnerFederation.md) |  | [optional] 

## Methods

### NewV0040KillJobsRespJob

`func NewV0040KillJobsRespJob(stepId string, jobId V0040Uint32NoVal, ) *V0040KillJobsRespJob`

NewV0040KillJobsRespJob instantiates a new V0040KillJobsRespJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0040KillJobsRespJobWithDefaults

`func NewV0040KillJobsRespJobWithDefaults() *V0040KillJobsRespJob`

NewV0040KillJobsRespJobWithDefaults instantiates a new V0040KillJobsRespJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *V0040KillJobsRespJob) GetError() V0041OpenapiKillJobsRespStatusInnerError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *V0040KillJobsRespJob) GetErrorOk() (*V0041OpenapiKillJobsRespStatusInnerError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *V0040KillJobsRespJob) SetError(v V0041OpenapiKillJobsRespStatusInnerError)`

SetError sets Error field to given value.

### HasError

`func (o *V0040KillJobsRespJob) HasError() bool`

HasError returns a boolean if a field has been set.

### GetStepId

`func (o *V0040KillJobsRespJob) GetStepId() string`

GetStepId returns the StepId field if non-nil, zero value otherwise.

### GetStepIdOk

`func (o *V0040KillJobsRespJob) GetStepIdOk() (*string, bool)`

GetStepIdOk returns a tuple with the StepId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepId

`func (o *V0040KillJobsRespJob) SetStepId(v string)`

SetStepId sets StepId field to given value.


### GetJobId

`func (o *V0040KillJobsRespJob) GetJobId() V0040Uint32NoVal`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *V0040KillJobsRespJob) GetJobIdOk() (*V0040Uint32NoVal, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *V0040KillJobsRespJob) SetJobId(v V0040Uint32NoVal)`

SetJobId sets JobId field to given value.


### GetFederation

`func (o *V0040KillJobsRespJob) GetFederation() V0041OpenapiKillJobsRespStatusInnerFederation`

GetFederation returns the Federation field if non-nil, zero value otherwise.

### GetFederationOk

`func (o *V0040KillJobsRespJob) GetFederationOk() (*V0041OpenapiKillJobsRespStatusInnerFederation, bool)`

GetFederationOk returns a tuple with the Federation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederation

`func (o *V0040KillJobsRespJob) SetFederation(v V0041OpenapiKillJobsRespStatusInnerFederation)`

SetFederation sets Federation field to given value.

### HasFederation

`func (o *V0040KillJobsRespJob) HasFederation() bool`

HasFederation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


