# V0044OpenapiJobModifyReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobIdList** | Pointer to **[]string** |  | [optional] 
**JobRec** | Pointer to [**V0044JobModify**](V0044JobModify.md) |  | [optional] 
**Meta** | Pointer to [**V0044OpenapiMeta**](V0044OpenapiMeta.md) |  | [optional] 
**Errors** | Pointer to [**[]V0044OpenapiError**](V0044OpenapiError.md) |  | [optional] 
**Warnings** | Pointer to [**[]V0044OpenapiWarning**](V0044OpenapiWarning.md) |  | [optional] 

## Methods

### NewV0044OpenapiJobModifyReq

`func NewV0044OpenapiJobModifyReq() *V0044OpenapiJobModifyReq`

NewV0044OpenapiJobModifyReq instantiates a new V0044OpenapiJobModifyReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0044OpenapiJobModifyReqWithDefaults

`func NewV0044OpenapiJobModifyReqWithDefaults() *V0044OpenapiJobModifyReq`

NewV0044OpenapiJobModifyReqWithDefaults instantiates a new V0044OpenapiJobModifyReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobIdList

`func (o *V0044OpenapiJobModifyReq) GetJobIdList() []string`

GetJobIdList returns the JobIdList field if non-nil, zero value otherwise.

### GetJobIdListOk

`func (o *V0044OpenapiJobModifyReq) GetJobIdListOk() (*[]string, bool)`

GetJobIdListOk returns a tuple with the JobIdList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobIdList

`func (o *V0044OpenapiJobModifyReq) SetJobIdList(v []string)`

SetJobIdList sets JobIdList field to given value.

### HasJobIdList

`func (o *V0044OpenapiJobModifyReq) HasJobIdList() bool`

HasJobIdList returns a boolean if a field has been set.

### GetJobRec

`func (o *V0044OpenapiJobModifyReq) GetJobRec() V0044JobModify`

GetJobRec returns the JobRec field if non-nil, zero value otherwise.

### GetJobRecOk

`func (o *V0044OpenapiJobModifyReq) GetJobRecOk() (*V0044JobModify, bool)`

GetJobRecOk returns a tuple with the JobRec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobRec

`func (o *V0044OpenapiJobModifyReq) SetJobRec(v V0044JobModify)`

SetJobRec sets JobRec field to given value.

### HasJobRec

`func (o *V0044OpenapiJobModifyReq) HasJobRec() bool`

HasJobRec returns a boolean if a field has been set.

### GetMeta

`func (o *V0044OpenapiJobModifyReq) GetMeta() V0044OpenapiMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *V0044OpenapiJobModifyReq) GetMetaOk() (*V0044OpenapiMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *V0044OpenapiJobModifyReq) SetMeta(v V0044OpenapiMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *V0044OpenapiJobModifyReq) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *V0044OpenapiJobModifyReq) GetErrors() []V0044OpenapiError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V0044OpenapiJobModifyReq) GetErrorsOk() (*[]V0044OpenapiError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V0044OpenapiJobModifyReq) SetErrors(v []V0044OpenapiError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V0044OpenapiJobModifyReq) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *V0044OpenapiJobModifyReq) GetWarnings() []V0044OpenapiWarning`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *V0044OpenapiJobModifyReq) GetWarningsOk() (*[]V0044OpenapiWarning, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *V0044OpenapiJobModifyReq) SetWarnings(v []V0044OpenapiWarning)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *V0044OpenapiJobModifyReq) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


