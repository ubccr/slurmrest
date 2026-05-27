# V0044OpenapiMetaSlurm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to [**V0044OpenapiMetaSlurmVersion**](V0044OpenapiMetaSlurmVersion.md) |  | [optional] 
**Release** | Pointer to **string** | Slurm release string | [optional] 
**Cluster** | Pointer to **string** | Slurm cluster name | [optional] 

## Methods

### NewV0044OpenapiMetaSlurm

`func NewV0044OpenapiMetaSlurm() *V0044OpenapiMetaSlurm`

NewV0044OpenapiMetaSlurm instantiates a new V0044OpenapiMetaSlurm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0044OpenapiMetaSlurmWithDefaults

`func NewV0044OpenapiMetaSlurmWithDefaults() *V0044OpenapiMetaSlurm`

NewV0044OpenapiMetaSlurmWithDefaults instantiates a new V0044OpenapiMetaSlurm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *V0044OpenapiMetaSlurm) GetVersion() V0044OpenapiMetaSlurmVersion`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *V0044OpenapiMetaSlurm) GetVersionOk() (*V0044OpenapiMetaSlurmVersion, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *V0044OpenapiMetaSlurm) SetVersion(v V0044OpenapiMetaSlurmVersion)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *V0044OpenapiMetaSlurm) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetRelease

`func (o *V0044OpenapiMetaSlurm) GetRelease() string`

GetRelease returns the Release field if non-nil, zero value otherwise.

### GetReleaseOk

`func (o *V0044OpenapiMetaSlurm) GetReleaseOk() (*string, bool)`

GetReleaseOk returns a tuple with the Release field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelease

`func (o *V0044OpenapiMetaSlurm) SetRelease(v string)`

SetRelease sets Release field to given value.

### HasRelease

`func (o *V0044OpenapiMetaSlurm) HasRelease() bool`

HasRelease returns a boolean if a field has been set.

### GetCluster

`func (o *V0044OpenapiMetaSlurm) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *V0044OpenapiMetaSlurm) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *V0044OpenapiMetaSlurm) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *V0044OpenapiMetaSlurm) HasCluster() bool`

HasCluster returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


