# V0041OpenapiSlurmdbdConfigResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Clusters** | Pointer to [**[]V0041OpenapiSlurmdbdConfigRespClustersInner**](V0041OpenapiSlurmdbdConfigRespClustersInner.md) | clusters | [optional] 
**Tres** | Pointer to [**[]V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTresRequestedMaxInner**](V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTresRequestedMaxInner.md) | tres | [optional] 
**Accounts** | Pointer to [**[]V0041OpenapiSlurmdbdConfigRespAccountsInner**](V0041OpenapiSlurmdbdConfigRespAccountsInner.md) | accounts | [optional] 
**Users** | Pointer to [**[]V0041OpenapiSlurmdbdConfigRespUsersInner**](V0041OpenapiSlurmdbdConfigRespUsersInner.md) | users | [optional] 
**Qos** | Pointer to [**[]V0041OpenapiSlurmdbdConfigRespQosInner**](V0041OpenapiSlurmdbdConfigRespQosInner.md) | qos | [optional] 
**Wckeys** | Pointer to [**[]V0041OpenapiSlurmdbdConfigRespUsersInnerWckeysInner**](V0041OpenapiSlurmdbdConfigRespUsersInnerWckeysInner.md) | wckeys | [optional] 
**Associations** | Pointer to [**[]V0041OpenapiSlurmdbdConfigRespAssociationsInner**](V0041OpenapiSlurmdbdConfigRespAssociationsInner.md) | associations | [optional] 
**Instances** | Pointer to [**[]V0041OpenapiSlurmdbdConfigRespInstancesInner**](V0041OpenapiSlurmdbdConfigRespInstancesInner.md) | instances | [optional] 
**Meta** | Pointer to [**V0041OpenapiSlurmdbdJobsRespMeta**](V0041OpenapiSlurmdbdJobsRespMeta.md) |  | [optional] 
**Errors** | Pointer to [**[]V0041OpenapiSlurmdbdJobsRespErrorsInner**](V0041OpenapiSlurmdbdJobsRespErrorsInner.md) | Query errors | [optional] 
**Warnings** | Pointer to [**[]V0041OpenapiSlurmdbdJobsRespWarningsInner**](V0041OpenapiSlurmdbdJobsRespWarningsInner.md) | Query warnings | [optional] 

## Methods

### NewV0041OpenapiSlurmdbdConfigResp

`func NewV0041OpenapiSlurmdbdConfigResp() *V0041OpenapiSlurmdbdConfigResp`

NewV0041OpenapiSlurmdbdConfigResp instantiates a new V0041OpenapiSlurmdbdConfigResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0041OpenapiSlurmdbdConfigRespWithDefaults

`func NewV0041OpenapiSlurmdbdConfigRespWithDefaults() *V0041OpenapiSlurmdbdConfigResp`

NewV0041OpenapiSlurmdbdConfigRespWithDefaults instantiates a new V0041OpenapiSlurmdbdConfigResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusters

`func (o *V0041OpenapiSlurmdbdConfigResp) GetClusters() []V0041OpenapiSlurmdbdConfigRespClustersInner`

GetClusters returns the Clusters field if non-nil, zero value otherwise.

### GetClustersOk

`func (o *V0041OpenapiSlurmdbdConfigResp) GetClustersOk() (*[]V0041OpenapiSlurmdbdConfigRespClustersInner, bool)`

GetClustersOk returns a tuple with the Clusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusters

`func (o *V0041OpenapiSlurmdbdConfigResp) SetClusters(v []V0041OpenapiSlurmdbdConfigRespClustersInner)`

SetClusters sets Clusters field to given value.

### HasClusters

`func (o *V0041OpenapiSlurmdbdConfigResp) HasClusters() bool`

HasClusters returns a boolean if a field has been set.

### GetTres

`func (o *V0041OpenapiSlurmdbdConfigResp) GetTres() []V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTresRequestedMaxInner`

GetTres returns the Tres field if non-nil, zero value otherwise.

### GetTresOk

`func (o *V0041OpenapiSlurmdbdConfigResp) GetTresOk() (*[]V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTresRequestedMaxInner, bool)`

GetTresOk returns a tuple with the Tres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTres

`func (o *V0041OpenapiSlurmdbdConfigResp) SetTres(v []V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTresRequestedMaxInner)`

SetTres sets Tres field to given value.

### HasTres

`func (o *V0041OpenapiSlurmdbdConfigResp) HasTres() bool`

HasTres returns a boolean if a field has been set.

### GetAccounts

`func (o *V0041OpenapiSlurmdbdConfigResp) GetAccounts() []V0041OpenapiSlurmdbdConfigRespAccountsInner`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *V0041OpenapiSlurmdbdConfigResp) GetAccountsOk() (*[]V0041OpenapiSlurmdbdConfigRespAccountsInner, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *V0041OpenapiSlurmdbdConfigResp) SetAccounts(v []V0041OpenapiSlurmdbdConfigRespAccountsInner)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *V0041OpenapiSlurmdbdConfigResp) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetUsers

`func (o *V0041OpenapiSlurmdbdConfigResp) GetUsers() []V0041OpenapiSlurmdbdConfigRespUsersInner`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *V0041OpenapiSlurmdbdConfigResp) GetUsersOk() (*[]V0041OpenapiSlurmdbdConfigRespUsersInner, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *V0041OpenapiSlurmdbdConfigResp) SetUsers(v []V0041OpenapiSlurmdbdConfigRespUsersInner)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *V0041OpenapiSlurmdbdConfigResp) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetQos

`func (o *V0041OpenapiSlurmdbdConfigResp) GetQos() []V0041OpenapiSlurmdbdConfigRespQosInner`

GetQos returns the Qos field if non-nil, zero value otherwise.

### GetQosOk

`func (o *V0041OpenapiSlurmdbdConfigResp) GetQosOk() (*[]V0041OpenapiSlurmdbdConfigRespQosInner, bool)`

GetQosOk returns a tuple with the Qos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQos

`func (o *V0041OpenapiSlurmdbdConfigResp) SetQos(v []V0041OpenapiSlurmdbdConfigRespQosInner)`

SetQos sets Qos field to given value.

### HasQos

`func (o *V0041OpenapiSlurmdbdConfigResp) HasQos() bool`

HasQos returns a boolean if a field has been set.

### GetWckeys

`func (o *V0041OpenapiSlurmdbdConfigResp) GetWckeys() []V0041OpenapiSlurmdbdConfigRespUsersInnerWckeysInner`

GetWckeys returns the Wckeys field if non-nil, zero value otherwise.

### GetWckeysOk

`func (o *V0041OpenapiSlurmdbdConfigResp) GetWckeysOk() (*[]V0041OpenapiSlurmdbdConfigRespUsersInnerWckeysInner, bool)`

GetWckeysOk returns a tuple with the Wckeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWckeys

`func (o *V0041OpenapiSlurmdbdConfigResp) SetWckeys(v []V0041OpenapiSlurmdbdConfigRespUsersInnerWckeysInner)`

SetWckeys sets Wckeys field to given value.

### HasWckeys

`func (o *V0041OpenapiSlurmdbdConfigResp) HasWckeys() bool`

HasWckeys returns a boolean if a field has been set.

### GetAssociations

`func (o *V0041OpenapiSlurmdbdConfigResp) GetAssociations() []V0041OpenapiSlurmdbdConfigRespAssociationsInner`

GetAssociations returns the Associations field if non-nil, zero value otherwise.

### GetAssociationsOk

`func (o *V0041OpenapiSlurmdbdConfigResp) GetAssociationsOk() (*[]V0041OpenapiSlurmdbdConfigRespAssociationsInner, bool)`

GetAssociationsOk returns a tuple with the Associations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociations

`func (o *V0041OpenapiSlurmdbdConfigResp) SetAssociations(v []V0041OpenapiSlurmdbdConfigRespAssociationsInner)`

SetAssociations sets Associations field to given value.

### HasAssociations

`func (o *V0041OpenapiSlurmdbdConfigResp) HasAssociations() bool`

HasAssociations returns a boolean if a field has been set.

### GetInstances

`func (o *V0041OpenapiSlurmdbdConfigResp) GetInstances() []V0041OpenapiSlurmdbdConfigRespInstancesInner`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *V0041OpenapiSlurmdbdConfigResp) GetInstancesOk() (*[]V0041OpenapiSlurmdbdConfigRespInstancesInner, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *V0041OpenapiSlurmdbdConfigResp) SetInstances(v []V0041OpenapiSlurmdbdConfigRespInstancesInner)`

SetInstances sets Instances field to given value.

### HasInstances

`func (o *V0041OpenapiSlurmdbdConfigResp) HasInstances() bool`

HasInstances returns a boolean if a field has been set.

### GetMeta

`func (o *V0041OpenapiSlurmdbdConfigResp) GetMeta() V0041OpenapiSlurmdbdJobsRespMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *V0041OpenapiSlurmdbdConfigResp) GetMetaOk() (*V0041OpenapiSlurmdbdJobsRespMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *V0041OpenapiSlurmdbdConfigResp) SetMeta(v V0041OpenapiSlurmdbdJobsRespMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *V0041OpenapiSlurmdbdConfigResp) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *V0041OpenapiSlurmdbdConfigResp) GetErrors() []V0041OpenapiSlurmdbdJobsRespErrorsInner`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V0041OpenapiSlurmdbdConfigResp) GetErrorsOk() (*[]V0041OpenapiSlurmdbdJobsRespErrorsInner, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V0041OpenapiSlurmdbdConfigResp) SetErrors(v []V0041OpenapiSlurmdbdJobsRespErrorsInner)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V0041OpenapiSlurmdbdConfigResp) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *V0041OpenapiSlurmdbdConfigResp) GetWarnings() []V0041OpenapiSlurmdbdJobsRespWarningsInner`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *V0041OpenapiSlurmdbdConfigResp) GetWarningsOk() (*[]V0041OpenapiSlurmdbdJobsRespWarningsInner, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *V0041OpenapiSlurmdbdConfigResp) SetWarnings(v []V0041OpenapiSlurmdbdJobsRespWarningsInner)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *V0041OpenapiSlurmdbdConfigResp) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


