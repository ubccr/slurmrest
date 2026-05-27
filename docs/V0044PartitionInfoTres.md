# V0044PartitionInfoTres

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingWeights** | Pointer to **string** | TRESBillingWeights - Billing weights of each tracked TRES type that will be used in calculating the usage of a job | [optional] 
**Configured** | Pointer to **string** | TRES - Number of each applicable TRES type available in this partition | [optional] 

## Methods

### NewV0044PartitionInfoTres

`func NewV0044PartitionInfoTres() *V0044PartitionInfoTres`

NewV0044PartitionInfoTres instantiates a new V0044PartitionInfoTres object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0044PartitionInfoTresWithDefaults

`func NewV0044PartitionInfoTresWithDefaults() *V0044PartitionInfoTres`

NewV0044PartitionInfoTresWithDefaults instantiates a new V0044PartitionInfoTres object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingWeights

`func (o *V0044PartitionInfoTres) GetBillingWeights() string`

GetBillingWeights returns the BillingWeights field if non-nil, zero value otherwise.

### GetBillingWeightsOk

`func (o *V0044PartitionInfoTres) GetBillingWeightsOk() (*string, bool)`

GetBillingWeightsOk returns a tuple with the BillingWeights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingWeights

`func (o *V0044PartitionInfoTres) SetBillingWeights(v string)`

SetBillingWeights sets BillingWeights field to given value.

### HasBillingWeights

`func (o *V0044PartitionInfoTres) HasBillingWeights() bool`

HasBillingWeights returns a boolean if a field has been set.

### GetConfigured

`func (o *V0044PartitionInfoTres) GetConfigured() string`

GetConfigured returns the Configured field if non-nil, zero value otherwise.

### GetConfiguredOk

`func (o *V0044PartitionInfoTres) GetConfiguredOk() (*string, bool)`

GetConfiguredOk returns a tuple with the Configured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigured

`func (o *V0044PartitionInfoTres) SetConfigured(v string)`

SetConfigured sets Configured field to given value.

### HasConfigured

`func (o *V0044PartitionInfoTres) HasConfigured() bool`

HasConfigured returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


