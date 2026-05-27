# V0044JobSubmitReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Script** | Pointer to **string** | Job batch script contents; Same as the script field in jobs[0] or job. | [optional] 
**Jobs** | Pointer to [**[]V0044JobDescMsg**](V0044JobDescMsg.md) |  | [optional] 
**Job** | Pointer to [**V0044JobDescMsg**](V0044JobDescMsg.md) |  | [optional] 

## Methods

### NewV0044JobSubmitReq

`func NewV0044JobSubmitReq() *V0044JobSubmitReq`

NewV0044JobSubmitReq instantiates a new V0044JobSubmitReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0044JobSubmitReqWithDefaults

`func NewV0044JobSubmitReqWithDefaults() *V0044JobSubmitReq`

NewV0044JobSubmitReqWithDefaults instantiates a new V0044JobSubmitReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScript

`func (o *V0044JobSubmitReq) GetScript() string`

GetScript returns the Script field if non-nil, zero value otherwise.

### GetScriptOk

`func (o *V0044JobSubmitReq) GetScriptOk() (*string, bool)`

GetScriptOk returns a tuple with the Script field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScript

`func (o *V0044JobSubmitReq) SetScript(v string)`

SetScript sets Script field to given value.

### HasScript

`func (o *V0044JobSubmitReq) HasScript() bool`

HasScript returns a boolean if a field has been set.

### GetJobs

`func (o *V0044JobSubmitReq) GetJobs() []V0044JobDescMsg`

GetJobs returns the Jobs field if non-nil, zero value otherwise.

### GetJobsOk

`func (o *V0044JobSubmitReq) GetJobsOk() (*[]V0044JobDescMsg, bool)`

GetJobsOk returns a tuple with the Jobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobs

`func (o *V0044JobSubmitReq) SetJobs(v []V0044JobDescMsg)`

SetJobs sets Jobs field to given value.

### HasJobs

`func (o *V0044JobSubmitReq) HasJobs() bool`

HasJobs returns a boolean if a field has been set.

### GetJob

`func (o *V0044JobSubmitReq) GetJob() V0044JobDescMsg`

GetJob returns the Job field if non-nil, zero value otherwise.

### GetJobOk

`func (o *V0044JobSubmitReq) GetJobOk() (*V0044JobDescMsg, bool)`

GetJobOk returns a tuple with the Job field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJob

`func (o *V0044JobSubmitReq) SetJob(v V0044JobDescMsg)`

SetJob sets Job field to given value.

### HasJob

`func (o *V0044JobSubmitReq) HasJob() bool`

HasJob returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


