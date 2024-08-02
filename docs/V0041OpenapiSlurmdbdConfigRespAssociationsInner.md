# V0041OpenapiSlurmdbdConfigRespAssociationsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accounting** | Pointer to [**[]V0041OpenapiSlurmdbdConfigRespUsersInnerWckeysInnerAccountingInner**](V0041OpenapiSlurmdbdConfigRespUsersInnerWckeysInnerAccountingInner.md) | Usage accounting | [optional] 
**Account** | Pointer to **string** |  | [optional] 
**Cluster** | Pointer to **string** |  | [optional] 
**Comment** | Pointer to **string** | comment for the association | [optional] 
**Default** | Pointer to [**V0041OpenapiSlurmdbdConfigRespAssociationsInnerDefault**](V0041OpenapiSlurmdbdConfigRespAssociationsInnerDefault.md) |  | [optional] 
**Flags** | Pointer to **[]string** | Active flags | [optional] 
**Max** | Pointer to [**V0041OpenapiSlurmdbdConfigRespAssociationsInnerMax**](V0041OpenapiSlurmdbdConfigRespAssociationsInnerMax.md) |  | [optional] 
**Id** | Pointer to **int32** | Association ID | [optional] 
**IsDefault** | Pointer to **bool** |  | [optional] 
**Lineage** | Pointer to **string** | Complete path up the hierarchy to the root association | [optional] 
**Min** | Pointer to [**V0041OpenapiSlurmdbdConfigRespAssociationsInnerMin**](V0041OpenapiSlurmdbdConfigRespAssociationsInnerMin.md) |  | [optional] 
**ParentAccount** | Pointer to **string** |  | [optional] 
**Partition** | Pointer to **string** |  | [optional] 
**Priority** | Pointer to [**V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId**](V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId.md) |  | [optional] 
**Qos** | Pointer to **[]string** | List of QOS names | [optional] 
**SharesRaw** | Pointer to **int32** |  | [optional] 
**User** | **string** |  | 

## Methods

### NewV0041OpenapiSlurmdbdConfigRespAssociationsInner

`func NewV0041OpenapiSlurmdbdConfigRespAssociationsInner(user string, ) *V0041OpenapiSlurmdbdConfigRespAssociationsInner`

NewV0041OpenapiSlurmdbdConfigRespAssociationsInner instantiates a new V0041OpenapiSlurmdbdConfigRespAssociationsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0041OpenapiSlurmdbdConfigRespAssociationsInnerWithDefaults

`func NewV0041OpenapiSlurmdbdConfigRespAssociationsInnerWithDefaults() *V0041OpenapiSlurmdbdConfigRespAssociationsInner`

NewV0041OpenapiSlurmdbdConfigRespAssociationsInnerWithDefaults instantiates a new V0041OpenapiSlurmdbdConfigRespAssociationsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccounting

`func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInner) GetAccounting() []V0041OpenapiSlurmdbdConfigRespUsersInnerWckeysInnerAccountingInner`

GetAccounting returns the Accounting field if non-nil, zero value otherwise.

### GetAccountingOk

`func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInner) GetAccountingOk() (*[]V0041OpenapiSlurmdbdConfigRespUsersInnerWckeysInnerAccountingInner, bool)`

GetAccountingOk returns a tuple with the Accounting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounting

`func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInner) SetAccounting(v []V0041OpenapiSlurmdbdConfigRespUsersInnerWckeysInnerAccountingInner)`

SetAccounting sets Accounting field to given value.

### HasAccounting

`func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInner) HasAccounting() bool`

HasAccounting returns a boolean if a field has been set.

### GetAccount

`func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInner) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInner) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInner) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInner) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetCluster

`func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInner) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInner) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInner) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInner) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetComment

`func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInner) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInner) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInner) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInner) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDefault

`func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInner) GetDefault() V0041OpenapiSlurmdbdConfigRespAssociationsInnerDefault`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInner) GetDefaultOk() (*V0041OpenapiSlurmdbdConfigRespAssociationsInnerDefault, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInner) SetDefault(v V0041OpenapiSlurmdbdConfigRespAssociationsInnerDefault)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInner) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetFlags

`func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInner) GetFlags() []string`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInner) GetFlagsOk() (*[]string, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInner) SetFlags(v []string)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInner) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetMax

`func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInner) GetMax() V0041OpenapiSlurmdbdConfigRespAssociationsInnerMax`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInner) GetMaxOk() (*V0041OpenapiSlurmdbdConfigRespAssociationsInnerMax, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInner) SetMax(v V0041OpenapiSlurmdbdConfigRespAssociationsInnerMax)`

SetMax sets Max field to given value.

### HasMax

`func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInner) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetId

`func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsDefault

`func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInner) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInner) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInner) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInner) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetLineage

`func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInner) GetLineage() string`

GetLineage returns the Lineage field if non-nil, zero value otherwise.

### GetLineageOk

`func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInner) GetLineageOk() (*string, bool)`

GetLineageOk returns a tuple with the Lineage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineage

`func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInner) SetLineage(v string)`

SetLineage sets Lineage field to given value.

### HasLineage

`func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInner) HasLineage() bool`

HasLineage returns a boolean if a field has been set.

### GetMin

`func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInner) GetMin() V0041OpenapiSlurmdbdConfigRespAssociationsInnerMin`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInner) GetMinOk() (*V0041OpenapiSlurmdbdConfigRespAssociationsInnerMin, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInner) SetMin(v V0041OpenapiSlurmdbdConfigRespAssociationsInnerMin)`

SetMin sets Min field to given value.

### HasMin

`func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInner) HasMin() bool`

HasMin returns a boolean if a field has been set.

### GetParentAccount

`func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInner) GetParentAccount() string`

GetParentAccount returns the ParentAccount field if non-nil, zero value otherwise.

### GetParentAccountOk

`func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInner) GetParentAccountOk() (*string, bool)`

GetParentAccountOk returns a tuple with the ParentAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentAccount

`func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInner) SetParentAccount(v string)`

SetParentAccount sets ParentAccount field to given value.

### HasParentAccount

`func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInner) HasParentAccount() bool`

HasParentAccount returns a boolean if a field has been set.

### GetPartition

`func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInner) GetPartition() string`

GetPartition returns the Partition field if non-nil, zero value otherwise.

### GetPartitionOk

`func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInner) GetPartitionOk() (*string, bool)`

GetPartitionOk returns a tuple with the Partition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartition

`func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInner) SetPartition(v string)`

SetPartition sets Partition field to given value.

### HasPartition

`func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInner) HasPartition() bool`

HasPartition returns a boolean if a field has been set.

### GetPriority

`func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInner) GetPriority() V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInner) GetPriorityOk() (*V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInner) SetPriority(v V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInner) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetQos

`func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInner) GetQos() []string`

GetQos returns the Qos field if non-nil, zero value otherwise.

### GetQosOk

`func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInner) GetQosOk() (*[]string, bool)`

GetQosOk returns a tuple with the Qos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQos

`func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInner) SetQos(v []string)`

SetQos sets Qos field to given value.

### HasQos

`func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInner) HasQos() bool`

HasQos returns a boolean if a field has been set.

### GetSharesRaw

`func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInner) GetSharesRaw() int32`

GetSharesRaw returns the SharesRaw field if non-nil, zero value otherwise.

### GetSharesRawOk

`func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInner) GetSharesRawOk() (*int32, bool)`

GetSharesRawOk returns a tuple with the SharesRaw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharesRaw

`func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInner) SetSharesRaw(v int32)`

SetSharesRaw sets SharesRaw field to given value.

### HasSharesRaw

`func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInner) HasSharesRaw() bool`

HasSharesRaw returns a boolean if a field has been set.

### GetUser

`func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInner) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInner) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *V0041OpenapiSlurmdbdConfigRespAssociationsInner) SetUser(v string)`

SetUser sets User field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


