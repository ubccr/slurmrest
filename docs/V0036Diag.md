# V0036Diag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to [**[]V0036Error**](V0036Error.md) | slurm errors | [optional] 
**Statistics** | Pointer to [**V0036DiagStatistics**](V0036DiagStatistics.md) |  | [optional] 

## Methods

### NewV0036Diag

`func NewV0036Diag() *V0036Diag`

NewV0036Diag instantiates a new V0036Diag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0036DiagWithDefaults

`func NewV0036DiagWithDefaults() *V0036Diag`

NewV0036DiagWithDefaults instantiates a new V0036Diag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *V0036Diag) GetErrors() []V0036Error`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V0036Diag) GetErrorsOk() (*[]V0036Error, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V0036Diag) SetErrors(v []V0036Error)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V0036Diag) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetStatistics

`func (o *V0036Diag) GetStatistics() V0036DiagStatistics`

GetStatistics returns the Statistics field if non-nil, zero value otherwise.

### GetStatisticsOk

`func (o *V0036Diag) GetStatisticsOk() (*V0036DiagStatistics, bool)`

GetStatisticsOk returns a tuple with the Statistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatistics

`func (o *V0036Diag) SetStatistics(v V0036DiagStatistics)`

SetStatistics sets Statistics field to given value.

### HasStatistics

`func (o *V0036Diag) HasStatistics() bool`

HasStatistics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


