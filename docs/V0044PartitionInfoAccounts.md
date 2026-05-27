# V0044PartitionInfoAccounts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Allowed** | Pointer to **string** | AllowAccounts - Comma-separated list of accounts which may execute jobs in the partition | [optional] 
**Deny** | Pointer to **string** | DenyAccounts - Comma-separated list of accounts which may not execute jobs in the partition | [optional] 

## Methods

### NewV0044PartitionInfoAccounts

`func NewV0044PartitionInfoAccounts() *V0044PartitionInfoAccounts`

NewV0044PartitionInfoAccounts instantiates a new V0044PartitionInfoAccounts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0044PartitionInfoAccountsWithDefaults

`func NewV0044PartitionInfoAccountsWithDefaults() *V0044PartitionInfoAccounts`

NewV0044PartitionInfoAccountsWithDefaults instantiates a new V0044PartitionInfoAccounts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowed

`func (o *V0044PartitionInfoAccounts) GetAllowed() string`

GetAllowed returns the Allowed field if non-nil, zero value otherwise.

### GetAllowedOk

`func (o *V0044PartitionInfoAccounts) GetAllowedOk() (*string, bool)`

GetAllowedOk returns a tuple with the Allowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowed

`func (o *V0044PartitionInfoAccounts) SetAllowed(v string)`

SetAllowed sets Allowed field to given value.

### HasAllowed

`func (o *V0044PartitionInfoAccounts) HasAllowed() bool`

HasAllowed returns a boolean if a field has been set.

### GetDeny

`func (o *V0044PartitionInfoAccounts) GetDeny() string`

GetDeny returns the Deny field if non-nil, zero value otherwise.

### GetDenyOk

`func (o *V0044PartitionInfoAccounts) GetDenyOk() (*string, bool)`

GetDenyOk returns a tuple with the Deny field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeny

`func (o *V0044PartitionInfoAccounts) SetDeny(v string)`

SetDeny sets Deny field to given value.

### HasDeny

`func (o *V0044PartitionInfoAccounts) HasDeny() bool`

HasDeny returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


