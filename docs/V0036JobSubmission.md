# V0036JobSubmission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Script** | **string** | Executable script (full contents) to run in batch step | 
**Job** | Pointer to [**V0036JobProperties**](V0036JobProperties.md) |  | [optional] 
**Jobs** | Pointer to [**[]V0036JobProperties**](V0036JobProperties.md) | Properties of an HetJob | [optional] 

## Methods

### NewV0036JobSubmission

`func NewV0036JobSubmission(script string, ) *V0036JobSubmission`

NewV0036JobSubmission instantiates a new V0036JobSubmission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0036JobSubmissionWithDefaults

`func NewV0036JobSubmissionWithDefaults() *V0036JobSubmission`

NewV0036JobSubmissionWithDefaults instantiates a new V0036JobSubmission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScript

`func (o *V0036JobSubmission) GetScript() string`

GetScript returns the Script field if non-nil, zero value otherwise.

### GetScriptOk

`func (o *V0036JobSubmission) GetScriptOk() (*string, bool)`

GetScriptOk returns a tuple with the Script field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScript

`func (o *V0036JobSubmission) SetScript(v string)`

SetScript sets Script field to given value.


### GetJob

`func (o *V0036JobSubmission) GetJob() V0036JobProperties`

GetJob returns the Job field if non-nil, zero value otherwise.

### GetJobOk

`func (o *V0036JobSubmission) GetJobOk() (*V0036JobProperties, bool)`

GetJobOk returns a tuple with the Job field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJob

`func (o *V0036JobSubmission) SetJob(v V0036JobProperties)`

SetJob sets Job field to given value.

### HasJob

`func (o *V0036JobSubmission) HasJob() bool`

HasJob returns a boolean if a field has been set.

### GetJobs

`func (o *V0036JobSubmission) GetJobs() []V0036JobProperties`

GetJobs returns the Jobs field if non-nil, zero value otherwise.

### GetJobsOk

`func (o *V0036JobSubmission) GetJobsOk() (*[]V0036JobProperties, bool)`

GetJobsOk returns a tuple with the Jobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobs

`func (o *V0036JobSubmission) SetJobs(v []V0036JobProperties)`

SetJobs sets Jobs field to given value.

### HasJobs

`func (o *V0036JobSubmission) HasJobs() bool`

HasJobs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


