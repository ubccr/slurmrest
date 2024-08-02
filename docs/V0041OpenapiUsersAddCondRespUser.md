# V0041OpenapiUsersAddCondRespUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Adminlevel** | Pointer to **[]string** | Admin level of user.  Valid levels are None, Operator, and Admin. | [optional] 
**Defaultaccount** | Pointer to **string** | Identify the default bank account name to be used for a job if none is specified at submission time. | [optional] 
**Defaultwckey** | Pointer to **string** | Identify the default Workload Characterization Key. | [optional] 

## Methods

### NewV0041OpenapiUsersAddCondRespUser

`func NewV0041OpenapiUsersAddCondRespUser() *V0041OpenapiUsersAddCondRespUser`

NewV0041OpenapiUsersAddCondRespUser instantiates a new V0041OpenapiUsersAddCondRespUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0041OpenapiUsersAddCondRespUserWithDefaults

`func NewV0041OpenapiUsersAddCondRespUserWithDefaults() *V0041OpenapiUsersAddCondRespUser`

NewV0041OpenapiUsersAddCondRespUserWithDefaults instantiates a new V0041OpenapiUsersAddCondRespUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminlevel

`func (o *V0041OpenapiUsersAddCondRespUser) GetAdminlevel() []string`

GetAdminlevel returns the Adminlevel field if non-nil, zero value otherwise.

### GetAdminlevelOk

`func (o *V0041OpenapiUsersAddCondRespUser) GetAdminlevelOk() (*[]string, bool)`

GetAdminlevelOk returns a tuple with the Adminlevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminlevel

`func (o *V0041OpenapiUsersAddCondRespUser) SetAdminlevel(v []string)`

SetAdminlevel sets Adminlevel field to given value.

### HasAdminlevel

`func (o *V0041OpenapiUsersAddCondRespUser) HasAdminlevel() bool`

HasAdminlevel returns a boolean if a field has been set.

### GetDefaultaccount

`func (o *V0041OpenapiUsersAddCondRespUser) GetDefaultaccount() string`

GetDefaultaccount returns the Defaultaccount field if non-nil, zero value otherwise.

### GetDefaultaccountOk

`func (o *V0041OpenapiUsersAddCondRespUser) GetDefaultaccountOk() (*string, bool)`

GetDefaultaccountOk returns a tuple with the Defaultaccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultaccount

`func (o *V0041OpenapiUsersAddCondRespUser) SetDefaultaccount(v string)`

SetDefaultaccount sets Defaultaccount field to given value.

### HasDefaultaccount

`func (o *V0041OpenapiUsersAddCondRespUser) HasDefaultaccount() bool`

HasDefaultaccount returns a boolean if a field has been set.

### GetDefaultwckey

`func (o *V0041OpenapiUsersAddCondRespUser) GetDefaultwckey() string`

GetDefaultwckey returns the Defaultwckey field if non-nil, zero value otherwise.

### GetDefaultwckeyOk

`func (o *V0041OpenapiUsersAddCondRespUser) GetDefaultwckeyOk() (*string, bool)`

GetDefaultwckeyOk returns a tuple with the Defaultwckey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultwckey

`func (o *V0041OpenapiUsersAddCondRespUser) SetDefaultwckey(v string)`

SetDefaultwckey sets Defaultwckey field to given value.

### HasDefaultwckey

`func (o *V0041OpenapiUsersAddCondRespUser) HasDefaultwckey() bool`

HasDefaultwckey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


