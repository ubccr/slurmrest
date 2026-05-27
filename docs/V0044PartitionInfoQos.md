# V0044PartitionInfoQos

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Allowed** | Pointer to **string** | AllowQOS - Comma-separated list of Qos which may execute jobs in the partition | [optional] 
**Deny** | Pointer to **string** | DenyQOS - Comma-separated list of Qos which may not execute jobs in the partition | [optional] 
**Assigned** | Pointer to **string** | QOS - QOS name containing limits that will apply to all jobs in this partition | [optional] 

## Methods

### NewV0044PartitionInfoQos

`func NewV0044PartitionInfoQos() *V0044PartitionInfoQos`

NewV0044PartitionInfoQos instantiates a new V0044PartitionInfoQos object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0044PartitionInfoQosWithDefaults

`func NewV0044PartitionInfoQosWithDefaults() *V0044PartitionInfoQos`

NewV0044PartitionInfoQosWithDefaults instantiates a new V0044PartitionInfoQos object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowed

`func (o *V0044PartitionInfoQos) GetAllowed() string`

GetAllowed returns the Allowed field if non-nil, zero value otherwise.

### GetAllowedOk

`func (o *V0044PartitionInfoQos) GetAllowedOk() (*string, bool)`

GetAllowedOk returns a tuple with the Allowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowed

`func (o *V0044PartitionInfoQos) SetAllowed(v string)`

SetAllowed sets Allowed field to given value.

### HasAllowed

`func (o *V0044PartitionInfoQos) HasAllowed() bool`

HasAllowed returns a boolean if a field has been set.

### GetDeny

`func (o *V0044PartitionInfoQos) GetDeny() string`

GetDeny returns the Deny field if non-nil, zero value otherwise.

### GetDenyOk

`func (o *V0044PartitionInfoQos) GetDenyOk() (*string, bool)`

GetDenyOk returns a tuple with the Deny field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeny

`func (o *V0044PartitionInfoQos) SetDeny(v string)`

SetDeny sets Deny field to given value.

### HasDeny

`func (o *V0044PartitionInfoQos) HasDeny() bool`

HasDeny returns a boolean if a field has been set.

### GetAssigned

`func (o *V0044PartitionInfoQos) GetAssigned() string`

GetAssigned returns the Assigned field if non-nil, zero value otherwise.

### GetAssignedOk

`func (o *V0044PartitionInfoQos) GetAssignedOk() (*string, bool)`

GetAssignedOk returns a tuple with the Assigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssigned

`func (o *V0044PartitionInfoQos) SetAssigned(v string)`

SetAssigned sets Assigned field to given value.

### HasAssigned

`func (o *V0044PartitionInfoQos) HasAssigned() bool`

HasAssigned returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


