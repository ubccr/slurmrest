# V0044User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdministratorLevel** | Pointer to **[]string** | AdminLevel granted to the user | [optional] 
**Associations** | Pointer to [**[]V0044AssocShort**](V0044AssocShort.md) |  | [optional] 
**Coordinators** | Pointer to [**[]V0044Coord**](V0044Coord.md) |  | [optional] 
**Default** | Pointer to [**V0044UserDefault**](V0044UserDefault.md) |  | [optional] 
**Flags** | Pointer to **[]string** | Flags associated with this user | [optional] 
**Name** | **string** | User name | 
**OldName** | Pointer to **string** | Previous user name | [optional] 
**Wckeys** | Pointer to [**[]V0044Wckey**](V0044Wckey.md) |  | [optional] 

## Methods

### NewV0044User

`func NewV0044User(name string, ) *V0044User`

NewV0044User instantiates a new V0044User object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0044UserWithDefaults

`func NewV0044UserWithDefaults() *V0044User`

NewV0044UserWithDefaults instantiates a new V0044User object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdministratorLevel

`func (o *V0044User) GetAdministratorLevel() []string`

GetAdministratorLevel returns the AdministratorLevel field if non-nil, zero value otherwise.

### GetAdministratorLevelOk

`func (o *V0044User) GetAdministratorLevelOk() (*[]string, bool)`

GetAdministratorLevelOk returns a tuple with the AdministratorLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministratorLevel

`func (o *V0044User) SetAdministratorLevel(v []string)`

SetAdministratorLevel sets AdministratorLevel field to given value.

### HasAdministratorLevel

`func (o *V0044User) HasAdministratorLevel() bool`

HasAdministratorLevel returns a boolean if a field has been set.

### GetAssociations

`func (o *V0044User) GetAssociations() []V0044AssocShort`

GetAssociations returns the Associations field if non-nil, zero value otherwise.

### GetAssociationsOk

`func (o *V0044User) GetAssociationsOk() (*[]V0044AssocShort, bool)`

GetAssociationsOk returns a tuple with the Associations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociations

`func (o *V0044User) SetAssociations(v []V0044AssocShort)`

SetAssociations sets Associations field to given value.

### HasAssociations

`func (o *V0044User) HasAssociations() bool`

HasAssociations returns a boolean if a field has been set.

### GetCoordinators

`func (o *V0044User) GetCoordinators() []V0044Coord`

GetCoordinators returns the Coordinators field if non-nil, zero value otherwise.

### GetCoordinatorsOk

`func (o *V0044User) GetCoordinatorsOk() (*[]V0044Coord, bool)`

GetCoordinatorsOk returns a tuple with the Coordinators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoordinators

`func (o *V0044User) SetCoordinators(v []V0044Coord)`

SetCoordinators sets Coordinators field to given value.

### HasCoordinators

`func (o *V0044User) HasCoordinators() bool`

HasCoordinators returns a boolean if a field has been set.

### GetDefault

`func (o *V0044User) GetDefault() V0044UserDefault`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *V0044User) GetDefaultOk() (*V0044UserDefault, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *V0044User) SetDefault(v V0044UserDefault)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *V0044User) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetFlags

`func (o *V0044User) GetFlags() []string`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *V0044User) GetFlagsOk() (*[]string, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *V0044User) SetFlags(v []string)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *V0044User) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetName

`func (o *V0044User) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V0044User) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V0044User) SetName(v string)`

SetName sets Name field to given value.


### GetOldName

`func (o *V0044User) GetOldName() string`

GetOldName returns the OldName field if non-nil, zero value otherwise.

### GetOldNameOk

`func (o *V0044User) GetOldNameOk() (*string, bool)`

GetOldNameOk returns a tuple with the OldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldName

`func (o *V0044User) SetOldName(v string)`

SetOldName sets OldName field to given value.

### HasOldName

`func (o *V0044User) HasOldName() bool`

HasOldName returns a boolean if a field has been set.

### GetWckeys

`func (o *V0044User) GetWckeys() []V0044Wckey`

GetWckeys returns the Wckeys field if non-nil, zero value otherwise.

### GetWckeysOk

`func (o *V0044User) GetWckeysOk() (*[]V0044Wckey, bool)`

GetWckeysOk returns a tuple with the Wckeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWckeys

`func (o *V0044User) SetWckeys(v []V0044Wckey)`

SetWckeys sets Wckeys field to given value.

### HasWckeys

`func (o *V0044User) HasWckeys() bool`

HasWckeys returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


