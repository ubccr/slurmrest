# V0044OpenapiSlurmdbdConfigResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Clusters** | Pointer to [**[]V0044ClusterRec**](V0044ClusterRec.md) |  | [optional] 
**Tres** | Pointer to [**[]V0044Tres**](V0044Tres.md) |  | [optional] 
**Accounts** | Pointer to [**[]V0044Account**](V0044Account.md) |  | [optional] 
**Users** | Pointer to [**[]V0044User**](V0044User.md) |  | [optional] 
**Qos** | Pointer to [**[]V0044Qos**](V0044Qos.md) |  | [optional] 
**Wckeys** | Pointer to [**[]V0044Wckey**](V0044Wckey.md) |  | [optional] 
**Associations** | Pointer to [**[]V0044Assoc**](V0044Assoc.md) |  | [optional] 
**Instances** | Pointer to [**[]V0044Instance**](V0044Instance.md) |  | [optional] 
**Meta** | Pointer to [**V0044OpenapiMeta**](V0044OpenapiMeta.md) |  | [optional] 
**Errors** | Pointer to [**[]V0044OpenapiError**](V0044OpenapiError.md) |  | [optional] 
**Warnings** | Pointer to [**[]V0044OpenapiWarning**](V0044OpenapiWarning.md) |  | [optional] 

## Methods

### NewV0044OpenapiSlurmdbdConfigResp

`func NewV0044OpenapiSlurmdbdConfigResp() *V0044OpenapiSlurmdbdConfigResp`

NewV0044OpenapiSlurmdbdConfigResp instantiates a new V0044OpenapiSlurmdbdConfigResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0044OpenapiSlurmdbdConfigRespWithDefaults

`func NewV0044OpenapiSlurmdbdConfigRespWithDefaults() *V0044OpenapiSlurmdbdConfigResp`

NewV0044OpenapiSlurmdbdConfigRespWithDefaults instantiates a new V0044OpenapiSlurmdbdConfigResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusters

`func (o *V0044OpenapiSlurmdbdConfigResp) GetClusters() []V0044ClusterRec`

GetClusters returns the Clusters field if non-nil, zero value otherwise.

### GetClustersOk

`func (o *V0044OpenapiSlurmdbdConfigResp) GetClustersOk() (*[]V0044ClusterRec, bool)`

GetClustersOk returns a tuple with the Clusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusters

`func (o *V0044OpenapiSlurmdbdConfigResp) SetClusters(v []V0044ClusterRec)`

SetClusters sets Clusters field to given value.

### HasClusters

`func (o *V0044OpenapiSlurmdbdConfigResp) HasClusters() bool`

HasClusters returns a boolean if a field has been set.

### GetTres

`func (o *V0044OpenapiSlurmdbdConfigResp) GetTres() []V0044Tres`

GetTres returns the Tres field if non-nil, zero value otherwise.

### GetTresOk

`func (o *V0044OpenapiSlurmdbdConfigResp) GetTresOk() (*[]V0044Tres, bool)`

GetTresOk returns a tuple with the Tres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTres

`func (o *V0044OpenapiSlurmdbdConfigResp) SetTres(v []V0044Tres)`

SetTres sets Tres field to given value.

### HasTres

`func (o *V0044OpenapiSlurmdbdConfigResp) HasTres() bool`

HasTres returns a boolean if a field has been set.

### GetAccounts

`func (o *V0044OpenapiSlurmdbdConfigResp) GetAccounts() []V0044Account`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *V0044OpenapiSlurmdbdConfigResp) GetAccountsOk() (*[]V0044Account, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *V0044OpenapiSlurmdbdConfigResp) SetAccounts(v []V0044Account)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *V0044OpenapiSlurmdbdConfigResp) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetUsers

`func (o *V0044OpenapiSlurmdbdConfigResp) GetUsers() []V0044User`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *V0044OpenapiSlurmdbdConfigResp) GetUsersOk() (*[]V0044User, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *V0044OpenapiSlurmdbdConfigResp) SetUsers(v []V0044User)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *V0044OpenapiSlurmdbdConfigResp) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetQos

`func (o *V0044OpenapiSlurmdbdConfigResp) GetQos() []V0044Qos`

GetQos returns the Qos field if non-nil, zero value otherwise.

### GetQosOk

`func (o *V0044OpenapiSlurmdbdConfigResp) GetQosOk() (*[]V0044Qos, bool)`

GetQosOk returns a tuple with the Qos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQos

`func (o *V0044OpenapiSlurmdbdConfigResp) SetQos(v []V0044Qos)`

SetQos sets Qos field to given value.

### HasQos

`func (o *V0044OpenapiSlurmdbdConfigResp) HasQos() bool`

HasQos returns a boolean if a field has been set.

### GetWckeys

`func (o *V0044OpenapiSlurmdbdConfigResp) GetWckeys() []V0044Wckey`

GetWckeys returns the Wckeys field if non-nil, zero value otherwise.

### GetWckeysOk

`func (o *V0044OpenapiSlurmdbdConfigResp) GetWckeysOk() (*[]V0044Wckey, bool)`

GetWckeysOk returns a tuple with the Wckeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWckeys

`func (o *V0044OpenapiSlurmdbdConfigResp) SetWckeys(v []V0044Wckey)`

SetWckeys sets Wckeys field to given value.

### HasWckeys

`func (o *V0044OpenapiSlurmdbdConfigResp) HasWckeys() bool`

HasWckeys returns a boolean if a field has been set.

### GetAssociations

`func (o *V0044OpenapiSlurmdbdConfigResp) GetAssociations() []V0044Assoc`

GetAssociations returns the Associations field if non-nil, zero value otherwise.

### GetAssociationsOk

`func (o *V0044OpenapiSlurmdbdConfigResp) GetAssociationsOk() (*[]V0044Assoc, bool)`

GetAssociationsOk returns a tuple with the Associations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociations

`func (o *V0044OpenapiSlurmdbdConfigResp) SetAssociations(v []V0044Assoc)`

SetAssociations sets Associations field to given value.

### HasAssociations

`func (o *V0044OpenapiSlurmdbdConfigResp) HasAssociations() bool`

HasAssociations returns a boolean if a field has been set.

### GetInstances

`func (o *V0044OpenapiSlurmdbdConfigResp) GetInstances() []V0044Instance`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *V0044OpenapiSlurmdbdConfigResp) GetInstancesOk() (*[]V0044Instance, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *V0044OpenapiSlurmdbdConfigResp) SetInstances(v []V0044Instance)`

SetInstances sets Instances field to given value.

### HasInstances

`func (o *V0044OpenapiSlurmdbdConfigResp) HasInstances() bool`

HasInstances returns a boolean if a field has been set.

### GetMeta

`func (o *V0044OpenapiSlurmdbdConfigResp) GetMeta() V0044OpenapiMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *V0044OpenapiSlurmdbdConfigResp) GetMetaOk() (*V0044OpenapiMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *V0044OpenapiSlurmdbdConfigResp) SetMeta(v V0044OpenapiMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *V0044OpenapiSlurmdbdConfigResp) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *V0044OpenapiSlurmdbdConfigResp) GetErrors() []V0044OpenapiError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V0044OpenapiSlurmdbdConfigResp) GetErrorsOk() (*[]V0044OpenapiError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V0044OpenapiSlurmdbdConfigResp) SetErrors(v []V0044OpenapiError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V0044OpenapiSlurmdbdConfigResp) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *V0044OpenapiSlurmdbdConfigResp) GetWarnings() []V0044OpenapiWarning`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *V0044OpenapiSlurmdbdConfigResp) GetWarningsOk() (*[]V0044OpenapiWarning, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *V0044OpenapiSlurmdbdConfigResp) SetWarnings(v []V0044OpenapiWarning)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *V0044OpenapiSlurmdbdConfigResp) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


