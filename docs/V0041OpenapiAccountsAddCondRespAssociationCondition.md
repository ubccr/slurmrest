# V0041OpenapiAccountsAddCondRespAssociationCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accounts** | **[]string** | CSV accounts list | 
**Association** | Pointer to [**V0041OpenapiUsersAddCondRespAssociationConditionAssociation**](V0041OpenapiUsersAddCondRespAssociationConditionAssociation.md) |  | [optional] 
**Clusters** | Pointer to **[]string** | CSV clusters list | [optional] 

## Methods

### NewV0041OpenapiAccountsAddCondRespAssociationCondition

`func NewV0041OpenapiAccountsAddCondRespAssociationCondition(accounts []string, ) *V0041OpenapiAccountsAddCondRespAssociationCondition`

NewV0041OpenapiAccountsAddCondRespAssociationCondition instantiates a new V0041OpenapiAccountsAddCondRespAssociationCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0041OpenapiAccountsAddCondRespAssociationConditionWithDefaults

`func NewV0041OpenapiAccountsAddCondRespAssociationConditionWithDefaults() *V0041OpenapiAccountsAddCondRespAssociationCondition`

NewV0041OpenapiAccountsAddCondRespAssociationConditionWithDefaults instantiates a new V0041OpenapiAccountsAddCondRespAssociationCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccounts

`func (o *V0041OpenapiAccountsAddCondRespAssociationCondition) GetAccounts() []string`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *V0041OpenapiAccountsAddCondRespAssociationCondition) GetAccountsOk() (*[]string, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *V0041OpenapiAccountsAddCondRespAssociationCondition) SetAccounts(v []string)`

SetAccounts sets Accounts field to given value.


### GetAssociation

`func (o *V0041OpenapiAccountsAddCondRespAssociationCondition) GetAssociation() V0041OpenapiUsersAddCondRespAssociationConditionAssociation`

GetAssociation returns the Association field if non-nil, zero value otherwise.

### GetAssociationOk

`func (o *V0041OpenapiAccountsAddCondRespAssociationCondition) GetAssociationOk() (*V0041OpenapiUsersAddCondRespAssociationConditionAssociation, bool)`

GetAssociationOk returns a tuple with the Association field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociation

`func (o *V0041OpenapiAccountsAddCondRespAssociationCondition) SetAssociation(v V0041OpenapiUsersAddCondRespAssociationConditionAssociation)`

SetAssociation sets Association field to given value.

### HasAssociation

`func (o *V0041OpenapiAccountsAddCondRespAssociationCondition) HasAssociation() bool`

HasAssociation returns a boolean if a field has been set.

### GetClusters

`func (o *V0041OpenapiAccountsAddCondRespAssociationCondition) GetClusters() []string`

GetClusters returns the Clusters field if non-nil, zero value otherwise.

### GetClustersOk

`func (o *V0041OpenapiAccountsAddCondRespAssociationCondition) GetClustersOk() (*[]string, bool)`

GetClustersOk returns a tuple with the Clusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusters

`func (o *V0041OpenapiAccountsAddCondRespAssociationCondition) SetClusters(v []string)`

SetClusters sets Clusters field to given value.

### HasClusters

`func (o *V0041OpenapiAccountsAddCondRespAssociationCondition) HasClusters() bool`

HasClusters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


