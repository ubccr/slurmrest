# V0044OpenapiAccountsAddCondResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssociationCondition** | [**V0044AccountsAddCond**](V0044AccountsAddCond.md) |  | 
**Account** | Pointer to [**V0044AccountShort**](V0044AccountShort.md) |  | [optional] 
**Meta** | Pointer to [**V0044OpenapiMeta**](V0044OpenapiMeta.md) |  | [optional] 
**Errors** | Pointer to [**[]V0044OpenapiError**](V0044OpenapiError.md) |  | [optional] 
**Warnings** | Pointer to [**[]V0044OpenapiWarning**](V0044OpenapiWarning.md) |  | [optional] 

## Methods

### NewV0044OpenapiAccountsAddCondResp

`func NewV0044OpenapiAccountsAddCondResp(associationCondition V0044AccountsAddCond, ) *V0044OpenapiAccountsAddCondResp`

NewV0044OpenapiAccountsAddCondResp instantiates a new V0044OpenapiAccountsAddCondResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0044OpenapiAccountsAddCondRespWithDefaults

`func NewV0044OpenapiAccountsAddCondRespWithDefaults() *V0044OpenapiAccountsAddCondResp`

NewV0044OpenapiAccountsAddCondRespWithDefaults instantiates a new V0044OpenapiAccountsAddCondResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssociationCondition

`func (o *V0044OpenapiAccountsAddCondResp) GetAssociationCondition() V0044AccountsAddCond`

GetAssociationCondition returns the AssociationCondition field if non-nil, zero value otherwise.

### GetAssociationConditionOk

`func (o *V0044OpenapiAccountsAddCondResp) GetAssociationConditionOk() (*V0044AccountsAddCond, bool)`

GetAssociationConditionOk returns a tuple with the AssociationCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociationCondition

`func (o *V0044OpenapiAccountsAddCondResp) SetAssociationCondition(v V0044AccountsAddCond)`

SetAssociationCondition sets AssociationCondition field to given value.


### GetAccount

`func (o *V0044OpenapiAccountsAddCondResp) GetAccount() V0044AccountShort`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *V0044OpenapiAccountsAddCondResp) GetAccountOk() (*V0044AccountShort, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *V0044OpenapiAccountsAddCondResp) SetAccount(v V0044AccountShort)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *V0044OpenapiAccountsAddCondResp) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetMeta

`func (o *V0044OpenapiAccountsAddCondResp) GetMeta() V0044OpenapiMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *V0044OpenapiAccountsAddCondResp) GetMetaOk() (*V0044OpenapiMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *V0044OpenapiAccountsAddCondResp) SetMeta(v V0044OpenapiMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *V0044OpenapiAccountsAddCondResp) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *V0044OpenapiAccountsAddCondResp) GetErrors() []V0044OpenapiError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V0044OpenapiAccountsAddCondResp) GetErrorsOk() (*[]V0044OpenapiError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V0044OpenapiAccountsAddCondResp) SetErrors(v []V0044OpenapiError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V0044OpenapiAccountsAddCondResp) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *V0044OpenapiAccountsAddCondResp) GetWarnings() []V0044OpenapiWarning`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *V0044OpenapiAccountsAddCondResp) GetWarningsOk() (*[]V0044OpenapiWarning, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *V0044OpenapiAccountsAddCondResp) SetWarnings(v []V0044OpenapiWarning)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *V0044OpenapiAccountsAddCondResp) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


