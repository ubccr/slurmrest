# V0037Diag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to [**[]V0037Error**](V0037Error.md) | slurm errors | [optional] 
**Statistics** | Pointer to [**V0037DiagStatistics**](V0037DiagStatistics.md) |  | [optional] 

## Methods

### NewV0037Diag

`func NewV0037Diag() *V0037Diag`

NewV0037Diag instantiates a new V0037Diag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0037DiagWithDefaults

`func NewV0037DiagWithDefaults() *V0037Diag`

NewV0037DiagWithDefaults instantiates a new V0037Diag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *V0037Diag) GetErrors() []V0037Error`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V0037Diag) GetErrorsOk() (*[]V0037Error, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V0037Diag) SetErrors(v []V0037Error)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V0037Diag) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetStatistics

`func (o *V0037Diag) GetStatistics() V0037DiagStatistics`

GetStatistics returns the Statistics field if non-nil, zero value otherwise.

### GetStatisticsOk

`func (o *V0037Diag) GetStatisticsOk() (*V0037DiagStatistics, bool)`

GetStatisticsOk returns a tuple with the Statistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatistics

`func (o *V0037Diag) SetStatistics(v V0037DiagStatistics)`

SetStatistics sets Statistics field to given value.

### HasStatistics

`func (o *V0037Diag) HasStatistics() bool`

HasStatistics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


