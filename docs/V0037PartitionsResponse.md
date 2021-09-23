# V0037PartitionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to [**[]V0037Error**](V0037Error.md) | slurm errors | [optional] 
**Partitions** | Pointer to [**[]V0037Partition**](V0037Partition.md) | partition info | [optional] 

## Methods

### NewV0037PartitionsResponse

`func NewV0037PartitionsResponse() *V0037PartitionsResponse`

NewV0037PartitionsResponse instantiates a new V0037PartitionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0037PartitionsResponseWithDefaults

`func NewV0037PartitionsResponseWithDefaults() *V0037PartitionsResponse`

NewV0037PartitionsResponseWithDefaults instantiates a new V0037PartitionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *V0037PartitionsResponse) GetErrors() []V0037Error`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V0037PartitionsResponse) GetErrorsOk() (*[]V0037Error, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V0037PartitionsResponse) SetErrors(v []V0037Error)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V0037PartitionsResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetPartitions

`func (o *V0037PartitionsResponse) GetPartitions() []V0037Partition`

GetPartitions returns the Partitions field if non-nil, zero value otherwise.

### GetPartitionsOk

`func (o *V0037PartitionsResponse) GetPartitionsOk() (*[]V0037Partition, bool)`

GetPartitionsOk returns a tuple with the Partitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitions

`func (o *V0037PartitionsResponse) SetPartitions(v []V0037Partition)`

SetPartitions sets Partitions field to given value.

### HasPartitions

`func (o *V0037PartitionsResponse) HasPartitions() bool`

HasPartitions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


