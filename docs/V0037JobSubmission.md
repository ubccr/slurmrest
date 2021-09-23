# V0037JobSubmission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Script** | **string** | Executable script (full contents) to run in batch step | 
**Job** | Pointer to [**V0037JobProperties**](V0037JobProperties.md) |  | [optional] 
**Jobs** | Pointer to [**[]V0037JobProperties**](V0037JobProperties.md) | Properties of an HetJob | [optional] 

## Methods

### NewV0037JobSubmission

`func NewV0037JobSubmission(script string, ) *V0037JobSubmission`

NewV0037JobSubmission instantiates a new V0037JobSubmission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0037JobSubmissionWithDefaults

`func NewV0037JobSubmissionWithDefaults() *V0037JobSubmission`

NewV0037JobSubmissionWithDefaults instantiates a new V0037JobSubmission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScript

`func (o *V0037JobSubmission) GetScript() string`

GetScript returns the Script field if non-nil, zero value otherwise.

### GetScriptOk

`func (o *V0037JobSubmission) GetScriptOk() (*string, bool)`

GetScriptOk returns a tuple with the Script field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScript

`func (o *V0037JobSubmission) SetScript(v string)`

SetScript sets Script field to given value.


### GetJob

`func (o *V0037JobSubmission) GetJob() V0037JobProperties`

GetJob returns the Job field if non-nil, zero value otherwise.

### GetJobOk

`func (o *V0037JobSubmission) GetJobOk() (*V0037JobProperties, bool)`

GetJobOk returns a tuple with the Job field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJob

`func (o *V0037JobSubmission) SetJob(v V0037JobProperties)`

SetJob sets Job field to given value.

### HasJob

`func (o *V0037JobSubmission) HasJob() bool`

HasJob returns a boolean if a field has been set.

### GetJobs

`func (o *V0037JobSubmission) GetJobs() []V0037JobProperties`

GetJobs returns the Jobs field if non-nil, zero value otherwise.

### GetJobsOk

`func (o *V0037JobSubmission) GetJobsOk() (*[]V0037JobProperties, bool)`

GetJobsOk returns a tuple with the Jobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobs

`func (o *V0037JobSubmission) SetJobs(v []V0037JobProperties)`

SetJobs sets Jobs field to given value.

### HasJobs

`func (o *V0037JobSubmission) HasJobs() bool`

HasJobs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


