# V0041OpenapiAccountsAddCondResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssociationCondition** | Pointer to [**V0041OpenapiAccountsAddCondRespAssociationCondition**](V0041OpenapiAccountsAddCondRespAssociationCondition.md) |  | [optional] 
**Account** | Pointer to [**V0041OpenapiAccountsAddCondRespAccount**](V0041OpenapiAccountsAddCondRespAccount.md) |  | [optional] 
**Meta** | Pointer to [**V0041OpenapiSlurmdbdJobsRespMeta**](V0041OpenapiSlurmdbdJobsRespMeta.md) |  | [optional] 
**Errors** | Pointer to [**[]V0041OpenapiSlurmdbdJobsRespErrorsInner**](V0041OpenapiSlurmdbdJobsRespErrorsInner.md) | Query errors | [optional] 
**Warnings** | Pointer to [**[]V0041OpenapiSlurmdbdJobsRespWarningsInner**](V0041OpenapiSlurmdbdJobsRespWarningsInner.md) | Query warnings | [optional] 

## Methods

### NewV0041OpenapiAccountsAddCondResp

`func NewV0041OpenapiAccountsAddCondResp() *V0041OpenapiAccountsAddCondResp`

NewV0041OpenapiAccountsAddCondResp instantiates a new V0041OpenapiAccountsAddCondResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0041OpenapiAccountsAddCondRespWithDefaults

`func NewV0041OpenapiAccountsAddCondRespWithDefaults() *V0041OpenapiAccountsAddCondResp`

NewV0041OpenapiAccountsAddCondRespWithDefaults instantiates a new V0041OpenapiAccountsAddCondResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssociationCondition

`func (o *V0041OpenapiAccountsAddCondResp) GetAssociationCondition() V0041OpenapiAccountsAddCondRespAssociationCondition`

GetAssociationCondition returns the AssociationCondition field if non-nil, zero value otherwise.

### GetAssociationConditionOk

`func (o *V0041OpenapiAccountsAddCondResp) GetAssociationConditionOk() (*V0041OpenapiAccountsAddCondRespAssociationCondition, bool)`

GetAssociationConditionOk returns a tuple with the AssociationCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociationCondition

`func (o *V0041OpenapiAccountsAddCondResp) SetAssociationCondition(v V0041OpenapiAccountsAddCondRespAssociationCondition)`

SetAssociationCondition sets AssociationCondition field to given value.

### HasAssociationCondition

`func (o *V0041OpenapiAccountsAddCondResp) HasAssociationCondition() bool`

HasAssociationCondition returns a boolean if a field has been set.

### GetAccount

`func (o *V0041OpenapiAccountsAddCondResp) GetAccount() V0041OpenapiAccountsAddCondRespAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *V0041OpenapiAccountsAddCondResp) GetAccountOk() (*V0041OpenapiAccountsAddCondRespAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *V0041OpenapiAccountsAddCondResp) SetAccount(v V0041OpenapiAccountsAddCondRespAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *V0041OpenapiAccountsAddCondResp) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetMeta

`func (o *V0041OpenapiAccountsAddCondResp) GetMeta() V0041OpenapiSlurmdbdJobsRespMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *V0041OpenapiAccountsAddCondResp) GetMetaOk() (*V0041OpenapiSlurmdbdJobsRespMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *V0041OpenapiAccountsAddCondResp) SetMeta(v V0041OpenapiSlurmdbdJobsRespMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *V0041OpenapiAccountsAddCondResp) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *V0041OpenapiAccountsAddCondResp) GetErrors() []V0041OpenapiSlurmdbdJobsRespErrorsInner`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V0041OpenapiAccountsAddCondResp) GetErrorsOk() (*[]V0041OpenapiSlurmdbdJobsRespErrorsInner, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V0041OpenapiAccountsAddCondResp) SetErrors(v []V0041OpenapiSlurmdbdJobsRespErrorsInner)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V0041OpenapiAccountsAddCondResp) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *V0041OpenapiAccountsAddCondResp) GetWarnings() []V0041OpenapiSlurmdbdJobsRespWarningsInner`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *V0041OpenapiAccountsAddCondResp) GetWarningsOk() (*[]V0041OpenapiSlurmdbdJobsRespWarningsInner, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *V0041OpenapiAccountsAddCondResp) SetWarnings(v []V0041OpenapiSlurmdbdJobsRespWarningsInner)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *V0041OpenapiAccountsAddCondResp) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


