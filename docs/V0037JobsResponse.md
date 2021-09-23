# V0037JobsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to [**[]V0037Error**](V0037Error.md) | slurm errors | [optional] 
**Jobs** | Pointer to [**[]V0037JobResponseProperties**](V0037JobResponseProperties.md) | job descriptions | [optional] 

## Methods

### NewV0037JobsResponse

`func NewV0037JobsResponse() *V0037JobsResponse`

NewV0037JobsResponse instantiates a new V0037JobsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0037JobsResponseWithDefaults

`func NewV0037JobsResponseWithDefaults() *V0037JobsResponse`

NewV0037JobsResponseWithDefaults instantiates a new V0037JobsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *V0037JobsResponse) GetErrors() []V0037Error`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V0037JobsResponse) GetErrorsOk() (*[]V0037Error, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V0037JobsResponse) SetErrors(v []V0037Error)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V0037JobsResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetJobs

`func (o *V0037JobsResponse) GetJobs() []V0037JobResponseProperties`

GetJobs returns the Jobs field if non-nil, zero value otherwise.

### GetJobsOk

`func (o *V0037JobsResponse) GetJobsOk() (*[]V0037JobResponseProperties, bool)`

GetJobsOk returns a tuple with the Jobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobs

`func (o *V0037JobsResponse) SetJobs(v []V0037JobResponseProperties)`

SetJobs sets Jobs field to given value.

### HasJobs

`func (o *V0037JobsResponse) HasJobs() bool`

HasJobs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


