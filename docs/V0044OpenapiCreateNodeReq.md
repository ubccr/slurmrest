# V0044OpenapiCreateNodeReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodeConf** | **string** | Node configuration line | 
**Meta** | Pointer to [**V0044OpenapiMeta**](V0044OpenapiMeta.md) |  | [optional] 
**Errors** | Pointer to [**[]V0044OpenapiError**](V0044OpenapiError.md) |  | [optional] 
**Warnings** | Pointer to [**[]V0044OpenapiWarning**](V0044OpenapiWarning.md) |  | [optional] 

## Methods

### NewV0044OpenapiCreateNodeReq

`func NewV0044OpenapiCreateNodeReq(nodeConf string, ) *V0044OpenapiCreateNodeReq`

NewV0044OpenapiCreateNodeReq instantiates a new V0044OpenapiCreateNodeReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0044OpenapiCreateNodeReqWithDefaults

`func NewV0044OpenapiCreateNodeReqWithDefaults() *V0044OpenapiCreateNodeReq`

NewV0044OpenapiCreateNodeReqWithDefaults instantiates a new V0044OpenapiCreateNodeReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodeConf

`func (o *V0044OpenapiCreateNodeReq) GetNodeConf() string`

GetNodeConf returns the NodeConf field if non-nil, zero value otherwise.

### GetNodeConfOk

`func (o *V0044OpenapiCreateNodeReq) GetNodeConfOk() (*string, bool)`

GetNodeConfOk returns a tuple with the NodeConf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeConf

`func (o *V0044OpenapiCreateNodeReq) SetNodeConf(v string)`

SetNodeConf sets NodeConf field to given value.


### GetMeta

`func (o *V0044OpenapiCreateNodeReq) GetMeta() V0044OpenapiMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *V0044OpenapiCreateNodeReq) GetMetaOk() (*V0044OpenapiMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *V0044OpenapiCreateNodeReq) SetMeta(v V0044OpenapiMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *V0044OpenapiCreateNodeReq) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *V0044OpenapiCreateNodeReq) GetErrors() []V0044OpenapiError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V0044OpenapiCreateNodeReq) GetErrorsOk() (*[]V0044OpenapiError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V0044OpenapiCreateNodeReq) SetErrors(v []V0044OpenapiError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V0044OpenapiCreateNodeReq) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *V0044OpenapiCreateNodeReq) GetWarnings() []V0044OpenapiWarning`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *V0044OpenapiCreateNodeReq) GetWarningsOk() (*[]V0044OpenapiWarning, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *V0044OpenapiCreateNodeReq) SetWarnings(v []V0044OpenapiWarning)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *V0044OpenapiCreateNodeReq) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


