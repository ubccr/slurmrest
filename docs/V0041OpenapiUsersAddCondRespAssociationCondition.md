# V0041OpenapiUsersAddCondRespAssociationCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accounts** | Pointer to **[]string** | CSV accounts list | [optional] 
**Association** | Pointer to [**V0041OpenapiUsersAddCondRespAssociationConditionAssociation**](V0041OpenapiUsersAddCondRespAssociationConditionAssociation.md) |  | [optional] 
**Clusters** | Pointer to **[]string** | CSV clusters list | [optional] 
**Partitions** | Pointer to **[]string** | CSV partitions list | [optional] 
**Users** | **[]string** | CSV users list | 
**Wckeys** | Pointer to **[]string** | CSV WCKeys list | [optional] 

## Methods

### NewV0041OpenapiUsersAddCondRespAssociationCondition

`func NewV0041OpenapiUsersAddCondRespAssociationCondition(users []string, ) *V0041OpenapiUsersAddCondRespAssociationCondition`

NewV0041OpenapiUsersAddCondRespAssociationCondition instantiates a new V0041OpenapiUsersAddCondRespAssociationCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0041OpenapiUsersAddCondRespAssociationConditionWithDefaults

`func NewV0041OpenapiUsersAddCondRespAssociationConditionWithDefaults() *V0041OpenapiUsersAddCondRespAssociationCondition`

NewV0041OpenapiUsersAddCondRespAssociationConditionWithDefaults instantiates a new V0041OpenapiUsersAddCondRespAssociationCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccounts

`func (o *V0041OpenapiUsersAddCondRespAssociationCondition) GetAccounts() []string`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *V0041OpenapiUsersAddCondRespAssociationCondition) GetAccountsOk() (*[]string, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *V0041OpenapiUsersAddCondRespAssociationCondition) SetAccounts(v []string)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *V0041OpenapiUsersAddCondRespAssociationCondition) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetAssociation

`func (o *V0041OpenapiUsersAddCondRespAssociationCondition) GetAssociation() V0041OpenapiUsersAddCondRespAssociationConditionAssociation`

GetAssociation returns the Association field if non-nil, zero value otherwise.

### GetAssociationOk

`func (o *V0041OpenapiUsersAddCondRespAssociationCondition) GetAssociationOk() (*V0041OpenapiUsersAddCondRespAssociationConditionAssociation, bool)`

GetAssociationOk returns a tuple with the Association field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociation

`func (o *V0041OpenapiUsersAddCondRespAssociationCondition) SetAssociation(v V0041OpenapiUsersAddCondRespAssociationConditionAssociation)`

SetAssociation sets Association field to given value.

### HasAssociation

`func (o *V0041OpenapiUsersAddCondRespAssociationCondition) HasAssociation() bool`

HasAssociation returns a boolean if a field has been set.

### GetClusters

`func (o *V0041OpenapiUsersAddCondRespAssociationCondition) GetClusters() []string`

GetClusters returns the Clusters field if non-nil, zero value otherwise.

### GetClustersOk

`func (o *V0041OpenapiUsersAddCondRespAssociationCondition) GetClustersOk() (*[]string, bool)`

GetClustersOk returns a tuple with the Clusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusters

`func (o *V0041OpenapiUsersAddCondRespAssociationCondition) SetClusters(v []string)`

SetClusters sets Clusters field to given value.

### HasClusters

`func (o *V0041OpenapiUsersAddCondRespAssociationCondition) HasClusters() bool`

HasClusters returns a boolean if a field has been set.

### GetPartitions

`func (o *V0041OpenapiUsersAddCondRespAssociationCondition) GetPartitions() []string`

GetPartitions returns the Partitions field if non-nil, zero value otherwise.

### GetPartitionsOk

`func (o *V0041OpenapiUsersAddCondRespAssociationCondition) GetPartitionsOk() (*[]string, bool)`

GetPartitionsOk returns a tuple with the Partitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitions

`func (o *V0041OpenapiUsersAddCondRespAssociationCondition) SetPartitions(v []string)`

SetPartitions sets Partitions field to given value.

### HasPartitions

`func (o *V0041OpenapiUsersAddCondRespAssociationCondition) HasPartitions() bool`

HasPartitions returns a boolean if a field has been set.

### GetUsers

`func (o *V0041OpenapiUsersAddCondRespAssociationCondition) GetUsers() []string`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *V0041OpenapiUsersAddCondRespAssociationCondition) GetUsersOk() (*[]string, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *V0041OpenapiUsersAddCondRespAssociationCondition) SetUsers(v []string)`

SetUsers sets Users field to given value.


### GetWckeys

`func (o *V0041OpenapiUsersAddCondRespAssociationCondition) GetWckeys() []string`

GetWckeys returns the Wckeys field if non-nil, zero value otherwise.

### GetWckeysOk

`func (o *V0041OpenapiUsersAddCondRespAssociationCondition) GetWckeysOk() (*[]string, bool)`

GetWckeysOk returns a tuple with the Wckeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWckeys

`func (o *V0041OpenapiUsersAddCondRespAssociationCondition) SetWckeys(v []string)`

SetWckeys sets Wckeys field to given value.

### HasWckeys

`func (o *V0041OpenapiUsersAddCondRespAssociationCondition) HasWckeys() bool`

HasWckeys returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


