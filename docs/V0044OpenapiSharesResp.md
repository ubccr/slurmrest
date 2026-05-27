# V0044OpenapiSharesResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Shares** | [**V0044SharesRespMsg**](V0044SharesRespMsg.md) |  | 
**Meta** | Pointer to [**V0044OpenapiMeta**](V0044OpenapiMeta.md) |  | [optional] 
**Errors** | Pointer to [**[]V0044OpenapiError**](V0044OpenapiError.md) |  | [optional] 
**Warnings** | Pointer to [**[]V0044OpenapiWarning**](V0044OpenapiWarning.md) |  | [optional] 

## Methods

### NewV0044OpenapiSharesResp

`func NewV0044OpenapiSharesResp(shares V0044SharesRespMsg, ) *V0044OpenapiSharesResp`

NewV0044OpenapiSharesResp instantiates a new V0044OpenapiSharesResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0044OpenapiSharesRespWithDefaults

`func NewV0044OpenapiSharesRespWithDefaults() *V0044OpenapiSharesResp`

NewV0044OpenapiSharesRespWithDefaults instantiates a new V0044OpenapiSharesResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShares

`func (o *V0044OpenapiSharesResp) GetShares() V0044SharesRespMsg`

GetShares returns the Shares field if non-nil, zero value otherwise.

### GetSharesOk

`func (o *V0044OpenapiSharesResp) GetSharesOk() (*V0044SharesRespMsg, bool)`

GetSharesOk returns a tuple with the Shares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShares

`func (o *V0044OpenapiSharesResp) SetShares(v V0044SharesRespMsg)`

SetShares sets Shares field to given value.


### GetMeta

`func (o *V0044OpenapiSharesResp) GetMeta() V0044OpenapiMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *V0044OpenapiSharesResp) GetMetaOk() (*V0044OpenapiMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *V0044OpenapiSharesResp) SetMeta(v V0044OpenapiMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *V0044OpenapiSharesResp) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *V0044OpenapiSharesResp) GetErrors() []V0044OpenapiError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V0044OpenapiSharesResp) GetErrorsOk() (*[]V0044OpenapiError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V0044OpenapiSharesResp) SetErrors(v []V0044OpenapiError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V0044OpenapiSharesResp) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *V0044OpenapiSharesResp) GetWarnings() []V0044OpenapiWarning`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *V0044OpenapiSharesResp) GetWarningsOk() (*[]V0044OpenapiWarning, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *V0044OpenapiSharesResp) SetWarnings(v []V0044OpenapiWarning)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *V0044OpenapiSharesResp) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


