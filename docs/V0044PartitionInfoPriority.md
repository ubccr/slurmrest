# V0044PartitionInfoPriority

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobFactor** | Pointer to **int32** | PriorityJobFactor - Partition factor used by priority/multifactor plugin in calculating job priority | [optional] 
**Tier** | Pointer to **int32** | PriorityTier - Controls the order in which the scheduler evaluates jobs from different partitions | [optional] 

## Methods

### NewV0044PartitionInfoPriority

`func NewV0044PartitionInfoPriority() *V0044PartitionInfoPriority`

NewV0044PartitionInfoPriority instantiates a new V0044PartitionInfoPriority object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0044PartitionInfoPriorityWithDefaults

`func NewV0044PartitionInfoPriorityWithDefaults() *V0044PartitionInfoPriority`

NewV0044PartitionInfoPriorityWithDefaults instantiates a new V0044PartitionInfoPriority object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobFactor

`func (o *V0044PartitionInfoPriority) GetJobFactor() int32`

GetJobFactor returns the JobFactor field if non-nil, zero value otherwise.

### GetJobFactorOk

`func (o *V0044PartitionInfoPriority) GetJobFactorOk() (*int32, bool)`

GetJobFactorOk returns a tuple with the JobFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobFactor

`func (o *V0044PartitionInfoPriority) SetJobFactor(v int32)`

SetJobFactor sets JobFactor field to given value.

### HasJobFactor

`func (o *V0044PartitionInfoPriority) HasJobFactor() bool`

HasJobFactor returns a boolean if a field has been set.

### GetTier

`func (o *V0044PartitionInfoPriority) GetTier() int32`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *V0044PartitionInfoPriority) GetTierOk() (*int32, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *V0044PartitionInfoPriority) SetTier(v int32)`

SetTier sets Tier field to given value.

### HasTier

`func (o *V0044PartitionInfoPriority) HasTier() bool`

HasTier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


