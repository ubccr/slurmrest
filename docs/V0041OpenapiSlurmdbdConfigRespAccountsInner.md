# V0041OpenapiSlurmdbdConfigRespAccountsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Associations** | Pointer to [**[]V0041OpenapiSlurmdbdJobsRespJobsInnerAssociation**](V0041OpenapiSlurmdbdJobsRespJobsInnerAssociation.md) |  | [optional] 
**Coordinators** | Pointer to [**[]V0041OpenapiSlurmdbdConfigRespAccountsInnerCoordinatorsInner**](V0041OpenapiSlurmdbdConfigRespAccountsInnerCoordinatorsInner.md) |  | [optional] 
**Description** | **string** |  | 
**Name** | **string** |  | 
**Organization** | **string** |  | 
**Flags** | Pointer to **[]string** |  | [optional] 

## Methods

### NewV0041OpenapiSlurmdbdConfigRespAccountsInner

`func NewV0041OpenapiSlurmdbdConfigRespAccountsInner(description string, name string, organization string, ) *V0041OpenapiSlurmdbdConfigRespAccountsInner`

NewV0041OpenapiSlurmdbdConfigRespAccountsInner instantiates a new V0041OpenapiSlurmdbdConfigRespAccountsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0041OpenapiSlurmdbdConfigRespAccountsInnerWithDefaults

`func NewV0041OpenapiSlurmdbdConfigRespAccountsInnerWithDefaults() *V0041OpenapiSlurmdbdConfigRespAccountsInner`

NewV0041OpenapiSlurmdbdConfigRespAccountsInnerWithDefaults instantiates a new V0041OpenapiSlurmdbdConfigRespAccountsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssociations

`func (o *V0041OpenapiSlurmdbdConfigRespAccountsInner) GetAssociations() []V0041OpenapiSlurmdbdJobsRespJobsInnerAssociation`

GetAssociations returns the Associations field if non-nil, zero value otherwise.

### GetAssociationsOk

`func (o *V0041OpenapiSlurmdbdConfigRespAccountsInner) GetAssociationsOk() (*[]V0041OpenapiSlurmdbdJobsRespJobsInnerAssociation, bool)`

GetAssociationsOk returns a tuple with the Associations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociations

`func (o *V0041OpenapiSlurmdbdConfigRespAccountsInner) SetAssociations(v []V0041OpenapiSlurmdbdJobsRespJobsInnerAssociation)`

SetAssociations sets Associations field to given value.

### HasAssociations

`func (o *V0041OpenapiSlurmdbdConfigRespAccountsInner) HasAssociations() bool`

HasAssociations returns a boolean if a field has been set.

### GetCoordinators

`func (o *V0041OpenapiSlurmdbdConfigRespAccountsInner) GetCoordinators() []V0041OpenapiSlurmdbdConfigRespAccountsInnerCoordinatorsInner`

GetCoordinators returns the Coordinators field if non-nil, zero value otherwise.

### GetCoordinatorsOk

`func (o *V0041OpenapiSlurmdbdConfigRespAccountsInner) GetCoordinatorsOk() (*[]V0041OpenapiSlurmdbdConfigRespAccountsInnerCoordinatorsInner, bool)`

GetCoordinatorsOk returns a tuple with the Coordinators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoordinators

`func (o *V0041OpenapiSlurmdbdConfigRespAccountsInner) SetCoordinators(v []V0041OpenapiSlurmdbdConfigRespAccountsInnerCoordinatorsInner)`

SetCoordinators sets Coordinators field to given value.

### HasCoordinators

`func (o *V0041OpenapiSlurmdbdConfigRespAccountsInner) HasCoordinators() bool`

HasCoordinators returns a boolean if a field has been set.

### GetDescription

`func (o *V0041OpenapiSlurmdbdConfigRespAccountsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *V0041OpenapiSlurmdbdConfigRespAccountsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *V0041OpenapiSlurmdbdConfigRespAccountsInner) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetName

`func (o *V0041OpenapiSlurmdbdConfigRespAccountsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V0041OpenapiSlurmdbdConfigRespAccountsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V0041OpenapiSlurmdbdConfigRespAccountsInner) SetName(v string)`

SetName sets Name field to given value.


### GetOrganization

`func (o *V0041OpenapiSlurmdbdConfigRespAccountsInner) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *V0041OpenapiSlurmdbdConfigRespAccountsInner) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *V0041OpenapiSlurmdbdConfigRespAccountsInner) SetOrganization(v string)`

SetOrganization sets Organization field to given value.


### GetFlags

`func (o *V0041OpenapiSlurmdbdConfigRespAccountsInner) GetFlags() []string`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *V0041OpenapiSlurmdbdConfigRespAccountsInner) GetFlagsOk() (*[]string, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *V0041OpenapiSlurmdbdConfigRespAccountsInner) SetFlags(v []string)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *V0041OpenapiSlurmdbdConfigRespAccountsInner) HasFlags() bool`

HasFlags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


