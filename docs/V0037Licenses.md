# V0037Licenses

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to [**[]V0037Error**](V0037Error.md) | slurm errors | [optional] 
**Licenses** | Pointer to [**[]V0037License**](V0037License.md) | licenses info | [optional] 

## Methods

### NewV0037Licenses

`func NewV0037Licenses() *V0037Licenses`

NewV0037Licenses instantiates a new V0037Licenses object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0037LicensesWithDefaults

`func NewV0037LicensesWithDefaults() *V0037Licenses`

NewV0037LicensesWithDefaults instantiates a new V0037Licenses object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *V0037Licenses) GetErrors() []V0037Error`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V0037Licenses) GetErrorsOk() (*[]V0037Error, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V0037Licenses) SetErrors(v []V0037Error)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V0037Licenses) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetLicenses

`func (o *V0037Licenses) GetLicenses() []V0037License`

GetLicenses returns the Licenses field if non-nil, zero value otherwise.

### GetLicensesOk

`func (o *V0037Licenses) GetLicensesOk() (*[]V0037License, bool)`

GetLicensesOk returns a tuple with the Licenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenses

`func (o *V0037Licenses) SetLicenses(v []V0037License)`

SetLicenses sets Licenses field to given value.

### HasLicenses

`func (o *V0037Licenses) HasLicenses() bool`

HasLicenses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


