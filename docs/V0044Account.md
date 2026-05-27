# V0044Account

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Associations** | Pointer to [**[]V0044AssocShort**](V0044AssocShort.md) |  | [optional] 
**Coordinators** | Pointer to [**[]V0044Coord**](V0044Coord.md) |  | [optional] 
**Description** | **string** | Arbitrary string describing the account | 
**Name** | **string** | Account name | 
**Organization** | **string** | Organization to which the account belongs | 
**Flags** | Pointer to **[]string** | Flags associated with this account | [optional] 

## Methods

### NewV0044Account

`func NewV0044Account(description string, name string, organization string, ) *V0044Account`

NewV0044Account instantiates a new V0044Account object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0044AccountWithDefaults

`func NewV0044AccountWithDefaults() *V0044Account`

NewV0044AccountWithDefaults instantiates a new V0044Account object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssociations

`func (o *V0044Account) GetAssociations() []V0044AssocShort`

GetAssociations returns the Associations field if non-nil, zero value otherwise.

### GetAssociationsOk

`func (o *V0044Account) GetAssociationsOk() (*[]V0044AssocShort, bool)`

GetAssociationsOk returns a tuple with the Associations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociations

`func (o *V0044Account) SetAssociations(v []V0044AssocShort)`

SetAssociations sets Associations field to given value.

### HasAssociations

`func (o *V0044Account) HasAssociations() bool`

HasAssociations returns a boolean if a field has been set.

### GetCoordinators

`func (o *V0044Account) GetCoordinators() []V0044Coord`

GetCoordinators returns the Coordinators field if non-nil, zero value otherwise.

### GetCoordinatorsOk

`func (o *V0044Account) GetCoordinatorsOk() (*[]V0044Coord, bool)`

GetCoordinatorsOk returns a tuple with the Coordinators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoordinators

`func (o *V0044Account) SetCoordinators(v []V0044Coord)`

SetCoordinators sets Coordinators field to given value.

### HasCoordinators

`func (o *V0044Account) HasCoordinators() bool`

HasCoordinators returns a boolean if a field has been set.

### GetDescription

`func (o *V0044Account) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *V0044Account) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *V0044Account) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetName

`func (o *V0044Account) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V0044Account) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V0044Account) SetName(v string)`

SetName sets Name field to given value.


### GetOrganization

`func (o *V0044Account) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *V0044Account) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *V0044Account) SetOrganization(v string)`

SetOrganization sets Organization field to given value.


### GetFlags

`func (o *V0044Account) GetFlags() []string`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *V0044Account) GetFlagsOk() (*[]string, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *V0044Account) SetFlags(v []string)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *V0044Account) HasFlags() bool`

HasFlags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


