# V0044InstanceTime

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimeEnd** | Pointer to **int64** | When the instance will end (UNIX timestamp) (UNIX timestamp or time string recognized by Slurm (e.g., &#39;[MM/DD[/YY]-]HH:MM[:SS]&#39;)) | [optional] 
**TimeStart** | Pointer to **int64** | When the instance will start (UNIX timestamp) (UNIX timestamp or time string recognized by Slurm (e.g., &#39;[MM/DD[/YY]-]HH:MM[:SS]&#39;)) | [optional] 

## Methods

### NewV0044InstanceTime

`func NewV0044InstanceTime() *V0044InstanceTime`

NewV0044InstanceTime instantiates a new V0044InstanceTime object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0044InstanceTimeWithDefaults

`func NewV0044InstanceTimeWithDefaults() *V0044InstanceTime`

NewV0044InstanceTimeWithDefaults instantiates a new V0044InstanceTime object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeEnd

`func (o *V0044InstanceTime) GetTimeEnd() int64`

GetTimeEnd returns the TimeEnd field if non-nil, zero value otherwise.

### GetTimeEndOk

`func (o *V0044InstanceTime) GetTimeEndOk() (*int64, bool)`

GetTimeEndOk returns a tuple with the TimeEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeEnd

`func (o *V0044InstanceTime) SetTimeEnd(v int64)`

SetTimeEnd sets TimeEnd field to given value.

### HasTimeEnd

`func (o *V0044InstanceTime) HasTimeEnd() bool`

HasTimeEnd returns a boolean if a field has been set.

### GetTimeStart

`func (o *V0044InstanceTime) GetTimeStart() int64`

GetTimeStart returns the TimeStart field if non-nil, zero value otherwise.

### GetTimeStartOk

`func (o *V0044InstanceTime) GetTimeStartOk() (*int64, bool)`

GetTimeStartOk returns a tuple with the TimeStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStart

`func (o *V0044InstanceTime) SetTimeStart(v int64)`

SetTimeStart sets TimeStart field to given value.

### HasTimeStart

`func (o *V0044InstanceTime) HasTimeStart() bool`

HasTimeStart returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


