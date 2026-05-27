# V0044UsersAddCond

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accounts** | Pointer to **[]string** |  | [optional] 
**Association** | Pointer to [**V0044AssocRecSet**](V0044AssocRecSet.md) |  | [optional] 
**Clusters** | Pointer to **[]string** |  | [optional] 
**Partitions** | Pointer to **[]string** |  | [optional] 
**Users** | **[]string** |  | 
**Wckeys** | Pointer to **[]string** |  | [optional] 

## Methods

### NewV0044UsersAddCond

`func NewV0044UsersAddCond(users []string, ) *V0044UsersAddCond`

NewV0044UsersAddCond instantiates a new V0044UsersAddCond object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0044UsersAddCondWithDefaults

`func NewV0044UsersAddCondWithDefaults() *V0044UsersAddCond`

NewV0044UsersAddCondWithDefaults instantiates a new V0044UsersAddCond object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccounts

`func (o *V0044UsersAddCond) GetAccounts() []string`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *V0044UsersAddCond) GetAccountsOk() (*[]string, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *V0044UsersAddCond) SetAccounts(v []string)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *V0044UsersAddCond) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetAssociation

`func (o *V0044UsersAddCond) GetAssociation() V0044AssocRecSet`

GetAssociation returns the Association field if non-nil, zero value otherwise.

### GetAssociationOk

`func (o *V0044UsersAddCond) GetAssociationOk() (*V0044AssocRecSet, bool)`

GetAssociationOk returns a tuple with the Association field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociation

`func (o *V0044UsersAddCond) SetAssociation(v V0044AssocRecSet)`

SetAssociation sets Association field to given value.

### HasAssociation

`func (o *V0044UsersAddCond) HasAssociation() bool`

HasAssociation returns a boolean if a field has been set.

### GetClusters

`func (o *V0044UsersAddCond) GetClusters() []string`

GetClusters returns the Clusters field if non-nil, zero value otherwise.

### GetClustersOk

`func (o *V0044UsersAddCond) GetClustersOk() (*[]string, bool)`

GetClustersOk returns a tuple with the Clusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusters

`func (o *V0044UsersAddCond) SetClusters(v []string)`

SetClusters sets Clusters field to given value.

### HasClusters

`func (o *V0044UsersAddCond) HasClusters() bool`

HasClusters returns a boolean if a field has been set.

### GetPartitions

`func (o *V0044UsersAddCond) GetPartitions() []string`

GetPartitions returns the Partitions field if non-nil, zero value otherwise.

### GetPartitionsOk

`func (o *V0044UsersAddCond) GetPartitionsOk() (*[]string, bool)`

GetPartitionsOk returns a tuple with the Partitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitions

`func (o *V0044UsersAddCond) SetPartitions(v []string)`

SetPartitions sets Partitions field to given value.

### HasPartitions

`func (o *V0044UsersAddCond) HasPartitions() bool`

HasPartitions returns a boolean if a field has been set.

### GetUsers

`func (o *V0044UsersAddCond) GetUsers() []string`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *V0044UsersAddCond) GetUsersOk() (*[]string, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *V0044UsersAddCond) SetUsers(v []string)`

SetUsers sets Users field to given value.


### GetWckeys

`func (o *V0044UsersAddCond) GetWckeys() []string`

GetWckeys returns the Wckeys field if non-nil, zero value otherwise.

### GetWckeysOk

`func (o *V0044UsersAddCond) GetWckeysOk() (*[]string, bool)`

GetWckeysOk returns a tuple with the Wckeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWckeys

`func (o *V0044UsersAddCond) SetWckeys(v []string)`

SetWckeys sets Wckeys field to given value.

### HasWckeys

`func (o *V0044UsersAddCond) HasWckeys() bool`

HasWckeys returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


