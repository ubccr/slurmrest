# V0044OpenapiLicensesResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Licenses** | [**[]V0044License**](V0044License.md) |  | 
**LastUpdate** | [**V0044Uint64NoValStruct**](V0044Uint64NoValStruct.md) |  | 
**Meta** | Pointer to [**V0044OpenapiMeta**](V0044OpenapiMeta.md) |  | [optional] 
**Errors** | Pointer to [**[]V0044OpenapiError**](V0044OpenapiError.md) |  | [optional] 
**Warnings** | Pointer to [**[]V0044OpenapiWarning**](V0044OpenapiWarning.md) |  | [optional] 

## Methods

### NewV0044OpenapiLicensesResp

`func NewV0044OpenapiLicensesResp(licenses []V0044License, lastUpdate V0044Uint64NoValStruct, ) *V0044OpenapiLicensesResp`

NewV0044OpenapiLicensesResp instantiates a new V0044OpenapiLicensesResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0044OpenapiLicensesRespWithDefaults

`func NewV0044OpenapiLicensesRespWithDefaults() *V0044OpenapiLicensesResp`

NewV0044OpenapiLicensesRespWithDefaults instantiates a new V0044OpenapiLicensesResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLicenses

`func (o *V0044OpenapiLicensesResp) GetLicenses() []V0044License`

GetLicenses returns the Licenses field if non-nil, zero value otherwise.

### GetLicensesOk

`func (o *V0044OpenapiLicensesResp) GetLicensesOk() (*[]V0044License, bool)`

GetLicensesOk returns a tuple with the Licenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenses

`func (o *V0044OpenapiLicensesResp) SetLicenses(v []V0044License)`

SetLicenses sets Licenses field to given value.


### GetLastUpdate

`func (o *V0044OpenapiLicensesResp) GetLastUpdate() V0044Uint64NoValStruct`

GetLastUpdate returns the LastUpdate field if non-nil, zero value otherwise.

### GetLastUpdateOk

`func (o *V0044OpenapiLicensesResp) GetLastUpdateOk() (*V0044Uint64NoValStruct, bool)`

GetLastUpdateOk returns a tuple with the LastUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdate

`func (o *V0044OpenapiLicensesResp) SetLastUpdate(v V0044Uint64NoValStruct)`

SetLastUpdate sets LastUpdate field to given value.


### GetMeta

`func (o *V0044OpenapiLicensesResp) GetMeta() V0044OpenapiMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *V0044OpenapiLicensesResp) GetMetaOk() (*V0044OpenapiMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *V0044OpenapiLicensesResp) SetMeta(v V0044OpenapiMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *V0044OpenapiLicensesResp) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *V0044OpenapiLicensesResp) GetErrors() []V0044OpenapiError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V0044OpenapiLicensesResp) GetErrorsOk() (*[]V0044OpenapiError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V0044OpenapiLicensesResp) SetErrors(v []V0044OpenapiError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V0044OpenapiLicensesResp) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *V0044OpenapiLicensesResp) GetWarnings() []V0044OpenapiWarning`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *V0044OpenapiLicensesResp) GetWarningsOk() (*[]V0044OpenapiWarning, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *V0044OpenapiLicensesResp) SetWarnings(v []V0044OpenapiWarning)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *V0044OpenapiLicensesResp) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


