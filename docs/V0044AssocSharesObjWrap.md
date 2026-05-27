# V0044AssocSharesObjWrap

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Association ID | [optional] 
**Cluster** | Pointer to **string** | Cluster name | [optional] 
**Name** | Pointer to **string** | Share name | [optional] 
**Parent** | Pointer to **string** | Parent name | [optional] 
**Partition** | Pointer to **string** | Partition name | [optional] 
**SharesNormalized** | Pointer to [**V0044Float64NoValStruct**](V0044Float64NoValStruct.md) |  | [optional] 
**Shares** | Pointer to [**V0044Uint32NoValStruct**](V0044Uint32NoValStruct.md) |  | [optional] 
**Tres** | Pointer to [**V0044AssocSharesObjWrapTres**](V0044AssocSharesObjWrapTres.md) |  | [optional] 
**EffectiveUsage** | Pointer to [**V0044Float64NoValStruct**](V0044Float64NoValStruct.md) |  | [optional] 
**UsageNormalized** | Pointer to [**V0044Float64NoValStruct**](V0044Float64NoValStruct.md) |  | [optional] 
**Usage** | Pointer to **int64** | Measure of tresbillableunits usage | [optional] 
**Fairshare** | Pointer to [**V0044AssocSharesObjWrapFairshare**](V0044AssocSharesObjWrapFairshare.md) |  | [optional] 
**Type** | Pointer to **[]string** | User or account association | [optional] 

## Methods

### NewV0044AssocSharesObjWrap

`func NewV0044AssocSharesObjWrap() *V0044AssocSharesObjWrap`

NewV0044AssocSharesObjWrap instantiates a new V0044AssocSharesObjWrap object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0044AssocSharesObjWrapWithDefaults

`func NewV0044AssocSharesObjWrapWithDefaults() *V0044AssocSharesObjWrap`

NewV0044AssocSharesObjWrapWithDefaults instantiates a new V0044AssocSharesObjWrap object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *V0044AssocSharesObjWrap) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V0044AssocSharesObjWrap) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V0044AssocSharesObjWrap) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *V0044AssocSharesObjWrap) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCluster

`func (o *V0044AssocSharesObjWrap) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *V0044AssocSharesObjWrap) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *V0044AssocSharesObjWrap) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *V0044AssocSharesObjWrap) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetName

`func (o *V0044AssocSharesObjWrap) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V0044AssocSharesObjWrap) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V0044AssocSharesObjWrap) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V0044AssocSharesObjWrap) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParent

`func (o *V0044AssocSharesObjWrap) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *V0044AssocSharesObjWrap) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *V0044AssocSharesObjWrap) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *V0044AssocSharesObjWrap) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPartition

`func (o *V0044AssocSharesObjWrap) GetPartition() string`

GetPartition returns the Partition field if non-nil, zero value otherwise.

### GetPartitionOk

`func (o *V0044AssocSharesObjWrap) GetPartitionOk() (*string, bool)`

GetPartitionOk returns a tuple with the Partition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartition

`func (o *V0044AssocSharesObjWrap) SetPartition(v string)`

SetPartition sets Partition field to given value.

### HasPartition

`func (o *V0044AssocSharesObjWrap) HasPartition() bool`

HasPartition returns a boolean if a field has been set.

### GetSharesNormalized

`func (o *V0044AssocSharesObjWrap) GetSharesNormalized() V0044Float64NoValStruct`

GetSharesNormalized returns the SharesNormalized field if non-nil, zero value otherwise.

### GetSharesNormalizedOk

`func (o *V0044AssocSharesObjWrap) GetSharesNormalizedOk() (*V0044Float64NoValStruct, bool)`

GetSharesNormalizedOk returns a tuple with the SharesNormalized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharesNormalized

`func (o *V0044AssocSharesObjWrap) SetSharesNormalized(v V0044Float64NoValStruct)`

SetSharesNormalized sets SharesNormalized field to given value.

### HasSharesNormalized

`func (o *V0044AssocSharesObjWrap) HasSharesNormalized() bool`

HasSharesNormalized returns a boolean if a field has been set.

### GetShares

`func (o *V0044AssocSharesObjWrap) GetShares() V0044Uint32NoValStruct`

GetShares returns the Shares field if non-nil, zero value otherwise.

### GetSharesOk

`func (o *V0044AssocSharesObjWrap) GetSharesOk() (*V0044Uint32NoValStruct, bool)`

GetSharesOk returns a tuple with the Shares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShares

`func (o *V0044AssocSharesObjWrap) SetShares(v V0044Uint32NoValStruct)`

SetShares sets Shares field to given value.

### HasShares

`func (o *V0044AssocSharesObjWrap) HasShares() bool`

HasShares returns a boolean if a field has been set.

### GetTres

`func (o *V0044AssocSharesObjWrap) GetTres() V0044AssocSharesObjWrapTres`

GetTres returns the Tres field if non-nil, zero value otherwise.

### GetTresOk

`func (o *V0044AssocSharesObjWrap) GetTresOk() (*V0044AssocSharesObjWrapTres, bool)`

GetTresOk returns a tuple with the Tres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTres

`func (o *V0044AssocSharesObjWrap) SetTres(v V0044AssocSharesObjWrapTres)`

SetTres sets Tres field to given value.

### HasTres

`func (o *V0044AssocSharesObjWrap) HasTres() bool`

HasTres returns a boolean if a field has been set.

### GetEffectiveUsage

`func (o *V0044AssocSharesObjWrap) GetEffectiveUsage() V0044Float64NoValStruct`

GetEffectiveUsage returns the EffectiveUsage field if non-nil, zero value otherwise.

### GetEffectiveUsageOk

`func (o *V0044AssocSharesObjWrap) GetEffectiveUsageOk() (*V0044Float64NoValStruct, bool)`

GetEffectiveUsageOk returns a tuple with the EffectiveUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveUsage

`func (o *V0044AssocSharesObjWrap) SetEffectiveUsage(v V0044Float64NoValStruct)`

SetEffectiveUsage sets EffectiveUsage field to given value.

### HasEffectiveUsage

`func (o *V0044AssocSharesObjWrap) HasEffectiveUsage() bool`

HasEffectiveUsage returns a boolean if a field has been set.

### GetUsageNormalized

`func (o *V0044AssocSharesObjWrap) GetUsageNormalized() V0044Float64NoValStruct`

GetUsageNormalized returns the UsageNormalized field if non-nil, zero value otherwise.

### GetUsageNormalizedOk

`func (o *V0044AssocSharesObjWrap) GetUsageNormalizedOk() (*V0044Float64NoValStruct, bool)`

GetUsageNormalizedOk returns a tuple with the UsageNormalized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageNormalized

`func (o *V0044AssocSharesObjWrap) SetUsageNormalized(v V0044Float64NoValStruct)`

SetUsageNormalized sets UsageNormalized field to given value.

### HasUsageNormalized

`func (o *V0044AssocSharesObjWrap) HasUsageNormalized() bool`

HasUsageNormalized returns a boolean if a field has been set.

### GetUsage

`func (o *V0044AssocSharesObjWrap) GetUsage() int64`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *V0044AssocSharesObjWrap) GetUsageOk() (*int64, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *V0044AssocSharesObjWrap) SetUsage(v int64)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *V0044AssocSharesObjWrap) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### GetFairshare

`func (o *V0044AssocSharesObjWrap) GetFairshare() V0044AssocSharesObjWrapFairshare`

GetFairshare returns the Fairshare field if non-nil, zero value otherwise.

### GetFairshareOk

`func (o *V0044AssocSharesObjWrap) GetFairshareOk() (*V0044AssocSharesObjWrapFairshare, bool)`

GetFairshareOk returns a tuple with the Fairshare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFairshare

`func (o *V0044AssocSharesObjWrap) SetFairshare(v V0044AssocSharesObjWrapFairshare)`

SetFairshare sets Fairshare field to given value.

### HasFairshare

`func (o *V0044AssocSharesObjWrap) HasFairshare() bool`

HasFairshare returns a boolean if a field has been set.

### GetType

`func (o *V0044AssocSharesObjWrap) GetType() []string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V0044AssocSharesObjWrap) GetTypeOk() (*[]string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V0044AssocSharesObjWrap) SetType(v []string)`

SetType sets Type field to given value.

### HasType

`func (o *V0044AssocSharesObjWrap) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


