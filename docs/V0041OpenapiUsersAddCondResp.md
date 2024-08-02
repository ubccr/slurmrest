# V0041OpenapiUsersAddCondResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssociationCondition** | [**V0041OpenapiUsersAddCondRespAssociationCondition**](V0041OpenapiUsersAddCondRespAssociationCondition.md) |  | 
**User** | [**V0041OpenapiUsersAddCondRespUser**](V0041OpenapiUsersAddCondRespUser.md) |  | 
**Meta** | Pointer to [**V0041OpenapiSlurmdbdJobsRespMeta**](V0041OpenapiSlurmdbdJobsRespMeta.md) |  | [optional] 
**Errors** | Pointer to [**[]V0041OpenapiSlurmdbdJobsRespErrorsInner**](V0041OpenapiSlurmdbdJobsRespErrorsInner.md) | Query errors | [optional] 
**Warnings** | Pointer to [**[]V0041OpenapiSlurmdbdJobsRespWarningsInner**](V0041OpenapiSlurmdbdJobsRespWarningsInner.md) | Query warnings | [optional] 

## Methods

### NewV0041OpenapiUsersAddCondResp

`func NewV0041OpenapiUsersAddCondResp(associationCondition V0041OpenapiUsersAddCondRespAssociationCondition, user V0041OpenapiUsersAddCondRespUser, ) *V0041OpenapiUsersAddCondResp`

NewV0041OpenapiUsersAddCondResp instantiates a new V0041OpenapiUsersAddCondResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0041OpenapiUsersAddCondRespWithDefaults

`func NewV0041OpenapiUsersAddCondRespWithDefaults() *V0041OpenapiUsersAddCondResp`

NewV0041OpenapiUsersAddCondRespWithDefaults instantiates a new V0041OpenapiUsersAddCondResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssociationCondition

`func (o *V0041OpenapiUsersAddCondResp) GetAssociationCondition() V0041OpenapiUsersAddCondRespAssociationCondition`

GetAssociationCondition returns the AssociationCondition field if non-nil, zero value otherwise.

### GetAssociationConditionOk

`func (o *V0041OpenapiUsersAddCondResp) GetAssociationConditionOk() (*V0041OpenapiUsersAddCondRespAssociationCondition, bool)`

GetAssociationConditionOk returns a tuple with the AssociationCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociationCondition

`func (o *V0041OpenapiUsersAddCondResp) SetAssociationCondition(v V0041OpenapiUsersAddCondRespAssociationCondition)`

SetAssociationCondition sets AssociationCondition field to given value.


### GetUser

`func (o *V0041OpenapiUsersAddCondResp) GetUser() V0041OpenapiUsersAddCondRespUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *V0041OpenapiUsersAddCondResp) GetUserOk() (*V0041OpenapiUsersAddCondRespUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *V0041OpenapiUsersAddCondResp) SetUser(v V0041OpenapiUsersAddCondRespUser)`

SetUser sets User field to given value.


### GetMeta

`func (o *V0041OpenapiUsersAddCondResp) GetMeta() V0041OpenapiSlurmdbdJobsRespMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *V0041OpenapiUsersAddCondResp) GetMetaOk() (*V0041OpenapiSlurmdbdJobsRespMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *V0041OpenapiUsersAddCondResp) SetMeta(v V0041OpenapiSlurmdbdJobsRespMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *V0041OpenapiUsersAddCondResp) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *V0041OpenapiUsersAddCondResp) GetErrors() []V0041OpenapiSlurmdbdJobsRespErrorsInner`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V0041OpenapiUsersAddCondResp) GetErrorsOk() (*[]V0041OpenapiSlurmdbdJobsRespErrorsInner, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V0041OpenapiUsersAddCondResp) SetErrors(v []V0041OpenapiSlurmdbdJobsRespErrorsInner)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V0041OpenapiUsersAddCondResp) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *V0041OpenapiUsersAddCondResp) GetWarnings() []V0041OpenapiSlurmdbdJobsRespWarningsInner`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *V0041OpenapiUsersAddCondResp) GetWarningsOk() (*[]V0041OpenapiSlurmdbdJobsRespWarningsInner, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *V0041OpenapiUsersAddCondResp) SetWarnings(v []V0041OpenapiSlurmdbdJobsRespWarningsInner)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *V0041OpenapiUsersAddCondResp) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


