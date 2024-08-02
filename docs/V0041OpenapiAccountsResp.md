# V0041OpenapiAccountsResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accounts** | [**[]V0041OpenapiSlurmdbdConfigRespAccountsInner**](V0041OpenapiSlurmdbdConfigRespAccountsInner.md) | accounts | 
**Meta** | Pointer to [**V0041OpenapiSlurmdbdJobsRespMeta**](V0041OpenapiSlurmdbdJobsRespMeta.md) |  | [optional] 
**Errors** | Pointer to [**[]V0041OpenapiSlurmdbdJobsRespErrorsInner**](V0041OpenapiSlurmdbdJobsRespErrorsInner.md) | Query errors | [optional] 
**Warnings** | Pointer to [**[]V0041OpenapiSlurmdbdJobsRespWarningsInner**](V0041OpenapiSlurmdbdJobsRespWarningsInner.md) | Query warnings | [optional] 

## Methods

### NewV0041OpenapiAccountsResp

`func NewV0041OpenapiAccountsResp(accounts []V0041OpenapiSlurmdbdConfigRespAccountsInner, ) *V0041OpenapiAccountsResp`

NewV0041OpenapiAccountsResp instantiates a new V0041OpenapiAccountsResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0041OpenapiAccountsRespWithDefaults

`func NewV0041OpenapiAccountsRespWithDefaults() *V0041OpenapiAccountsResp`

NewV0041OpenapiAccountsRespWithDefaults instantiates a new V0041OpenapiAccountsResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccounts

`func (o *V0041OpenapiAccountsResp) GetAccounts() []V0041OpenapiSlurmdbdConfigRespAccountsInner`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *V0041OpenapiAccountsResp) GetAccountsOk() (*[]V0041OpenapiSlurmdbdConfigRespAccountsInner, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *V0041OpenapiAccountsResp) SetAccounts(v []V0041OpenapiSlurmdbdConfigRespAccountsInner)`

SetAccounts sets Accounts field to given value.


### GetMeta

`func (o *V0041OpenapiAccountsResp) GetMeta() V0041OpenapiSlurmdbdJobsRespMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *V0041OpenapiAccountsResp) GetMetaOk() (*V0041OpenapiSlurmdbdJobsRespMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *V0041OpenapiAccountsResp) SetMeta(v V0041OpenapiSlurmdbdJobsRespMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *V0041OpenapiAccountsResp) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *V0041OpenapiAccountsResp) GetErrors() []V0041OpenapiSlurmdbdJobsRespErrorsInner`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V0041OpenapiAccountsResp) GetErrorsOk() (*[]V0041OpenapiSlurmdbdJobsRespErrorsInner, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V0041OpenapiAccountsResp) SetErrors(v []V0041OpenapiSlurmdbdJobsRespErrorsInner)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V0041OpenapiAccountsResp) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *V0041OpenapiAccountsResp) GetWarnings() []V0041OpenapiSlurmdbdJobsRespWarningsInner`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *V0041OpenapiAccountsResp) GetWarningsOk() (*[]V0041OpenapiSlurmdbdJobsRespWarningsInner, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *V0041OpenapiAccountsResp) SetWarnings(v []V0041OpenapiSlurmdbdJobsRespWarningsInner)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *V0041OpenapiAccountsResp) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


