# V0036JobsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to [**[]V0036Error**](V0036Error.md) | slurm errors | [optional] 
**Jobs** | Pointer to [**[]V0036JobResponseProperties**](V0036JobResponseProperties.md) | job descriptions | [optional] 

## Methods

### NewV0036JobsResponse

`func NewV0036JobsResponse() *V0036JobsResponse`

NewV0036JobsResponse instantiates a new V0036JobsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0036JobsResponseWithDefaults

`func NewV0036JobsResponseWithDefaults() *V0036JobsResponse`

NewV0036JobsResponseWithDefaults instantiates a new V0036JobsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *V0036JobsResponse) GetErrors() []V0036Error`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V0036JobsResponse) GetErrorsOk() (*[]V0036Error, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V0036JobsResponse) SetErrors(v []V0036Error)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V0036JobsResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetJobs

`func (o *V0036JobsResponse) GetJobs() []V0036JobResponseProperties`

GetJobs returns the Jobs field if non-nil, zero value otherwise.

### GetJobsOk

`func (o *V0036JobsResponse) GetJobsOk() (*[]V0036JobResponseProperties, bool)`

GetJobsOk returns a tuple with the Jobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobs

`func (o *V0036JobsResponse) SetJobs(v []V0036JobResponseProperties)`

SetJobs sets Jobs field to given value.

### HasJobs

`func (o *V0036JobsResponse) HasJobs() bool`

HasJobs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


