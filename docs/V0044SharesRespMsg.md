# V0044SharesRespMsg

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Shares** | Pointer to [**[]V0044AssocSharesObjWrap**](V0044AssocSharesObjWrap.md) |  | [optional] 
**TotalShares** | Pointer to **int64** | Total number of shares | [optional] 

## Methods

### NewV0044SharesRespMsg

`func NewV0044SharesRespMsg() *V0044SharesRespMsg`

NewV0044SharesRespMsg instantiates a new V0044SharesRespMsg object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0044SharesRespMsgWithDefaults

`func NewV0044SharesRespMsgWithDefaults() *V0044SharesRespMsg`

NewV0044SharesRespMsgWithDefaults instantiates a new V0044SharesRespMsg object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShares

`func (o *V0044SharesRespMsg) GetShares() []V0044AssocSharesObjWrap`

GetShares returns the Shares field if non-nil, zero value otherwise.

### GetSharesOk

`func (o *V0044SharesRespMsg) GetSharesOk() (*[]V0044AssocSharesObjWrap, bool)`

GetSharesOk returns a tuple with the Shares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShares

`func (o *V0044SharesRespMsg) SetShares(v []V0044AssocSharesObjWrap)`

SetShares sets Shares field to given value.

### HasShares

`func (o *V0044SharesRespMsg) HasShares() bool`

HasShares returns a boolean if a field has been set.

### GetTotalShares

`func (o *V0044SharesRespMsg) GetTotalShares() int64`

GetTotalShares returns the TotalShares field if non-nil, zero value otherwise.

### GetTotalSharesOk

`func (o *V0044SharesRespMsg) GetTotalSharesOk() (*int64, bool)`

GetTotalSharesOk returns a tuple with the TotalShares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalShares

`func (o *V0044SharesRespMsg) SetTotalShares(v int64)`

SetTotalShares sets TotalShares field to given value.

### HasTotalShares

`func (o *V0044SharesRespMsg) HasTotalShares() bool`

HasTotalShares returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


