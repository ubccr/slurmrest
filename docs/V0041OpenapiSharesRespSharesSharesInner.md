# V0041OpenapiSharesRespSharesSharesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | assocation id | [optional] 
**Cluster** | Pointer to **string** | cluster name | [optional] 
**Name** | Pointer to **string** | share name | [optional] 
**Parent** | Pointer to **string** | parent name | [optional] 
**Partition** | Pointer to **string** | partition name | [optional] 
**SharesNormalized** | Pointer to [**V0041OpenapiSharesRespSharesSharesInnerSharesNormalized**](V0041OpenapiSharesRespSharesSharesInnerSharesNormalized.md) |  | [optional] 
**Shares** | Pointer to [**V0041OpenapiSharesRespSharesSharesInnerShares**](V0041OpenapiSharesRespSharesSharesInnerShares.md) |  | [optional] 
**Tres** | Pointer to [**V0041OpenapiSharesRespSharesSharesInnerTres**](V0041OpenapiSharesRespSharesSharesInnerTres.md) |  | [optional] 
**EffectiveUsage** | Pointer to **float64** | effective, normalized usage | [optional] 
**UsageNormalized** | Pointer to [**V0041OpenapiSharesRespSharesSharesInnerUsageNormalized**](V0041OpenapiSharesRespSharesSharesInnerUsageNormalized.md) |  | [optional] 
**Usage** | Pointer to **int64** | measure of tresbillableunits usage | [optional] 
**Fairshare** | Pointer to [**V0041OpenapiSharesRespSharesSharesInnerFairshare**](V0041OpenapiSharesRespSharesSharesInnerFairshare.md) |  | [optional] 
**Type** | Pointer to **[]string** | user or account association | [optional] 

## Methods

### NewV0041OpenapiSharesRespSharesSharesInner

`func NewV0041OpenapiSharesRespSharesSharesInner() *V0041OpenapiSharesRespSharesSharesInner`

NewV0041OpenapiSharesRespSharesSharesInner instantiates a new V0041OpenapiSharesRespSharesSharesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0041OpenapiSharesRespSharesSharesInnerWithDefaults

`func NewV0041OpenapiSharesRespSharesSharesInnerWithDefaults() *V0041OpenapiSharesRespSharesSharesInner`

NewV0041OpenapiSharesRespSharesSharesInnerWithDefaults instantiates a new V0041OpenapiSharesRespSharesSharesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *V0041OpenapiSharesRespSharesSharesInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V0041OpenapiSharesRespSharesSharesInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V0041OpenapiSharesRespSharesSharesInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *V0041OpenapiSharesRespSharesSharesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCluster

`func (o *V0041OpenapiSharesRespSharesSharesInner) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *V0041OpenapiSharesRespSharesSharesInner) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *V0041OpenapiSharesRespSharesSharesInner) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *V0041OpenapiSharesRespSharesSharesInner) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetName

`func (o *V0041OpenapiSharesRespSharesSharesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V0041OpenapiSharesRespSharesSharesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V0041OpenapiSharesRespSharesSharesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V0041OpenapiSharesRespSharesSharesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParent

`func (o *V0041OpenapiSharesRespSharesSharesInner) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *V0041OpenapiSharesRespSharesSharesInner) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *V0041OpenapiSharesRespSharesSharesInner) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *V0041OpenapiSharesRespSharesSharesInner) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPartition

`func (o *V0041OpenapiSharesRespSharesSharesInner) GetPartition() string`

GetPartition returns the Partition field if non-nil, zero value otherwise.

### GetPartitionOk

`func (o *V0041OpenapiSharesRespSharesSharesInner) GetPartitionOk() (*string, bool)`

GetPartitionOk returns a tuple with the Partition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartition

`func (o *V0041OpenapiSharesRespSharesSharesInner) SetPartition(v string)`

SetPartition sets Partition field to given value.

### HasPartition

`func (o *V0041OpenapiSharesRespSharesSharesInner) HasPartition() bool`

HasPartition returns a boolean if a field has been set.

### GetSharesNormalized

`func (o *V0041OpenapiSharesRespSharesSharesInner) GetSharesNormalized() V0041OpenapiSharesRespSharesSharesInnerSharesNormalized`

GetSharesNormalized returns the SharesNormalized field if non-nil, zero value otherwise.

### GetSharesNormalizedOk

`func (o *V0041OpenapiSharesRespSharesSharesInner) GetSharesNormalizedOk() (*V0041OpenapiSharesRespSharesSharesInnerSharesNormalized, bool)`

GetSharesNormalizedOk returns a tuple with the SharesNormalized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharesNormalized

`func (o *V0041OpenapiSharesRespSharesSharesInner) SetSharesNormalized(v V0041OpenapiSharesRespSharesSharesInnerSharesNormalized)`

SetSharesNormalized sets SharesNormalized field to given value.

### HasSharesNormalized

`func (o *V0041OpenapiSharesRespSharesSharesInner) HasSharesNormalized() bool`

HasSharesNormalized returns a boolean if a field has been set.

### GetShares

`func (o *V0041OpenapiSharesRespSharesSharesInner) GetShares() V0041OpenapiSharesRespSharesSharesInnerShares`

GetShares returns the Shares field if non-nil, zero value otherwise.

### GetSharesOk

`func (o *V0041OpenapiSharesRespSharesSharesInner) GetSharesOk() (*V0041OpenapiSharesRespSharesSharesInnerShares, bool)`

GetSharesOk returns a tuple with the Shares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShares

`func (o *V0041OpenapiSharesRespSharesSharesInner) SetShares(v V0041OpenapiSharesRespSharesSharesInnerShares)`

SetShares sets Shares field to given value.

### HasShares

`func (o *V0041OpenapiSharesRespSharesSharesInner) HasShares() bool`

HasShares returns a boolean if a field has been set.

### GetTres

`func (o *V0041OpenapiSharesRespSharesSharesInner) GetTres() V0041OpenapiSharesRespSharesSharesInnerTres`

GetTres returns the Tres field if non-nil, zero value otherwise.

### GetTresOk

`func (o *V0041OpenapiSharesRespSharesSharesInner) GetTresOk() (*V0041OpenapiSharesRespSharesSharesInnerTres, bool)`

GetTresOk returns a tuple with the Tres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTres

`func (o *V0041OpenapiSharesRespSharesSharesInner) SetTres(v V0041OpenapiSharesRespSharesSharesInnerTres)`

SetTres sets Tres field to given value.

### HasTres

`func (o *V0041OpenapiSharesRespSharesSharesInner) HasTres() bool`

HasTres returns a boolean if a field has been set.

### GetEffectiveUsage

`func (o *V0041OpenapiSharesRespSharesSharesInner) GetEffectiveUsage() float64`

GetEffectiveUsage returns the EffectiveUsage field if non-nil, zero value otherwise.

### GetEffectiveUsageOk

`func (o *V0041OpenapiSharesRespSharesSharesInner) GetEffectiveUsageOk() (*float64, bool)`

GetEffectiveUsageOk returns a tuple with the EffectiveUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveUsage

`func (o *V0041OpenapiSharesRespSharesSharesInner) SetEffectiveUsage(v float64)`

SetEffectiveUsage sets EffectiveUsage field to given value.

### HasEffectiveUsage

`func (o *V0041OpenapiSharesRespSharesSharesInner) HasEffectiveUsage() bool`

HasEffectiveUsage returns a boolean if a field has been set.

### GetUsageNormalized

`func (o *V0041OpenapiSharesRespSharesSharesInner) GetUsageNormalized() V0041OpenapiSharesRespSharesSharesInnerUsageNormalized`

GetUsageNormalized returns the UsageNormalized field if non-nil, zero value otherwise.

### GetUsageNormalizedOk

`func (o *V0041OpenapiSharesRespSharesSharesInner) GetUsageNormalizedOk() (*V0041OpenapiSharesRespSharesSharesInnerUsageNormalized, bool)`

GetUsageNormalizedOk returns a tuple with the UsageNormalized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageNormalized

`func (o *V0041OpenapiSharesRespSharesSharesInner) SetUsageNormalized(v V0041OpenapiSharesRespSharesSharesInnerUsageNormalized)`

SetUsageNormalized sets UsageNormalized field to given value.

### HasUsageNormalized

`func (o *V0041OpenapiSharesRespSharesSharesInner) HasUsageNormalized() bool`

HasUsageNormalized returns a boolean if a field has been set.

### GetUsage

`func (o *V0041OpenapiSharesRespSharesSharesInner) GetUsage() int64`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *V0041OpenapiSharesRespSharesSharesInner) GetUsageOk() (*int64, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *V0041OpenapiSharesRespSharesSharesInner) SetUsage(v int64)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *V0041OpenapiSharesRespSharesSharesInner) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### GetFairshare

`func (o *V0041OpenapiSharesRespSharesSharesInner) GetFairshare() V0041OpenapiSharesRespSharesSharesInnerFairshare`

GetFairshare returns the Fairshare field if non-nil, zero value otherwise.

### GetFairshareOk

`func (o *V0041OpenapiSharesRespSharesSharesInner) GetFairshareOk() (*V0041OpenapiSharesRespSharesSharesInnerFairshare, bool)`

GetFairshareOk returns a tuple with the Fairshare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFairshare

`func (o *V0041OpenapiSharesRespSharesSharesInner) SetFairshare(v V0041OpenapiSharesRespSharesSharesInnerFairshare)`

SetFairshare sets Fairshare field to given value.

### HasFairshare

`func (o *V0041OpenapiSharesRespSharesSharesInner) HasFairshare() bool`

HasFairshare returns a boolean if a field has been set.

### GetType

`func (o *V0041OpenapiSharesRespSharesSharesInner) GetType() []string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V0041OpenapiSharesRespSharesSharesInner) GetTypeOk() (*[]string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V0041OpenapiSharesRespSharesSharesInner) SetType(v []string)`

SetType sets Type field to given value.

### HasType

`func (o *V0041OpenapiSharesRespSharesSharesInner) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


