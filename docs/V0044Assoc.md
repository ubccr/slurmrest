# V0044Assoc

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accounting** | Pointer to [**[]V0044Accounting**](V0044Accounting.md) |  | [optional] 
**Account** | Pointer to **string** | Account name | [optional] 
**Cluster** | Pointer to **string** | Cluster name | [optional] 
**Comment** | Pointer to **string** | Arbitrary comment | [optional] 
**Default** | Pointer to [**V0044AssocDefault**](V0044AssocDefault.md) |  | [optional] 
**Flags** | Pointer to **[]string** | Flags on the association | [optional] 
**Max** | Pointer to [**V0044AssocMax**](V0044AssocMax.md) |  | [optional] 
**Id** | Pointer to **int32** | Unique ID (Association ID) | [optional] 
**IsDefault** | Pointer to **bool** | Is default association for user | [optional] 
**Lineage** | Pointer to **string** | Complete path up the hierarchy to the root association | [optional] 
**Min** | Pointer to [**V0044AssocMin**](V0044AssocMin.md) |  | [optional] 
**ParentAccount** | Pointer to **string** | Name of parent account | [optional] 
**Partition** | Pointer to **string** | Partition name | [optional] 
**Priority** | Pointer to [**V0044Uint32NoValStruct**](V0044Uint32NoValStruct.md) |  | [optional] 
**Qos** | Pointer to **[]string** | List of QOS names | [optional] 
**SharesRaw** | Pointer to **int32** | Allocated shares used for fairshare calculation | [optional] 
**User** | **string** | User name | 

## Methods

### NewV0044Assoc

`func NewV0044Assoc(user string, ) *V0044Assoc`

NewV0044Assoc instantiates a new V0044Assoc object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0044AssocWithDefaults

`func NewV0044AssocWithDefaults() *V0044Assoc`

NewV0044AssocWithDefaults instantiates a new V0044Assoc object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccounting

`func (o *V0044Assoc) GetAccounting() []V0044Accounting`

GetAccounting returns the Accounting field if non-nil, zero value otherwise.

### GetAccountingOk

`func (o *V0044Assoc) GetAccountingOk() (*[]V0044Accounting, bool)`

GetAccountingOk returns a tuple with the Accounting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounting

`func (o *V0044Assoc) SetAccounting(v []V0044Accounting)`

SetAccounting sets Accounting field to given value.

### HasAccounting

`func (o *V0044Assoc) HasAccounting() bool`

HasAccounting returns a boolean if a field has been set.

### GetAccount

`func (o *V0044Assoc) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *V0044Assoc) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *V0044Assoc) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *V0044Assoc) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetCluster

`func (o *V0044Assoc) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *V0044Assoc) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *V0044Assoc) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *V0044Assoc) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetComment

`func (o *V0044Assoc) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *V0044Assoc) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *V0044Assoc) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *V0044Assoc) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDefault

`func (o *V0044Assoc) GetDefault() V0044AssocDefault`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *V0044Assoc) GetDefaultOk() (*V0044AssocDefault, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *V0044Assoc) SetDefault(v V0044AssocDefault)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *V0044Assoc) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetFlags

`func (o *V0044Assoc) GetFlags() []string`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *V0044Assoc) GetFlagsOk() (*[]string, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *V0044Assoc) SetFlags(v []string)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *V0044Assoc) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetMax

`func (o *V0044Assoc) GetMax() V0044AssocMax`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *V0044Assoc) GetMaxOk() (*V0044AssocMax, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *V0044Assoc) SetMax(v V0044AssocMax)`

SetMax sets Max field to given value.

### HasMax

`func (o *V0044Assoc) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetId

`func (o *V0044Assoc) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V0044Assoc) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V0044Assoc) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *V0044Assoc) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsDefault

`func (o *V0044Assoc) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *V0044Assoc) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *V0044Assoc) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *V0044Assoc) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetLineage

`func (o *V0044Assoc) GetLineage() string`

GetLineage returns the Lineage field if non-nil, zero value otherwise.

### GetLineageOk

`func (o *V0044Assoc) GetLineageOk() (*string, bool)`

GetLineageOk returns a tuple with the Lineage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineage

`func (o *V0044Assoc) SetLineage(v string)`

SetLineage sets Lineage field to given value.

### HasLineage

`func (o *V0044Assoc) HasLineage() bool`

HasLineage returns a boolean if a field has been set.

### GetMin

`func (o *V0044Assoc) GetMin() V0044AssocMin`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *V0044Assoc) GetMinOk() (*V0044AssocMin, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *V0044Assoc) SetMin(v V0044AssocMin)`

SetMin sets Min field to given value.

### HasMin

`func (o *V0044Assoc) HasMin() bool`

HasMin returns a boolean if a field has been set.

### GetParentAccount

`func (o *V0044Assoc) GetParentAccount() string`

GetParentAccount returns the ParentAccount field if non-nil, zero value otherwise.

### GetParentAccountOk

`func (o *V0044Assoc) GetParentAccountOk() (*string, bool)`

GetParentAccountOk returns a tuple with the ParentAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentAccount

`func (o *V0044Assoc) SetParentAccount(v string)`

SetParentAccount sets ParentAccount field to given value.

### HasParentAccount

`func (o *V0044Assoc) HasParentAccount() bool`

HasParentAccount returns a boolean if a field has been set.

### GetPartition

`func (o *V0044Assoc) GetPartition() string`

GetPartition returns the Partition field if non-nil, zero value otherwise.

### GetPartitionOk

`func (o *V0044Assoc) GetPartitionOk() (*string, bool)`

GetPartitionOk returns a tuple with the Partition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartition

`func (o *V0044Assoc) SetPartition(v string)`

SetPartition sets Partition field to given value.

### HasPartition

`func (o *V0044Assoc) HasPartition() bool`

HasPartition returns a boolean if a field has been set.

### GetPriority

`func (o *V0044Assoc) GetPriority() V0044Uint32NoValStruct`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *V0044Assoc) GetPriorityOk() (*V0044Uint32NoValStruct, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *V0044Assoc) SetPriority(v V0044Uint32NoValStruct)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *V0044Assoc) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetQos

`func (o *V0044Assoc) GetQos() []string`

GetQos returns the Qos field if non-nil, zero value otherwise.

### GetQosOk

`func (o *V0044Assoc) GetQosOk() (*[]string, bool)`

GetQosOk returns a tuple with the Qos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQos

`func (o *V0044Assoc) SetQos(v []string)`

SetQos sets Qos field to given value.

### HasQos

`func (o *V0044Assoc) HasQos() bool`

HasQos returns a boolean if a field has been set.

### GetSharesRaw

`func (o *V0044Assoc) GetSharesRaw() int32`

GetSharesRaw returns the SharesRaw field if non-nil, zero value otherwise.

### GetSharesRawOk

`func (o *V0044Assoc) GetSharesRawOk() (*int32, bool)`

GetSharesRawOk returns a tuple with the SharesRaw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharesRaw

`func (o *V0044Assoc) SetSharesRaw(v int32)`

SetSharesRaw sets SharesRaw field to given value.

### HasSharesRaw

`func (o *V0044Assoc) HasSharesRaw() bool`

HasSharesRaw returns a boolean if a field has been set.

### GetUser

`func (o *V0044Assoc) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *V0044Assoc) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *V0044Assoc) SetUser(v string)`

SetUser sets User field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


