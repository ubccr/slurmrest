# V0044JobAllocReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hetjob** | Pointer to [**[]V0044JobDescMsg**](V0044JobDescMsg.md) |  | [optional] 
**Job** | Pointer to [**V0044JobDescMsg**](V0044JobDescMsg.md) |  | [optional] 

## Methods

### NewV0044JobAllocReq

`func NewV0044JobAllocReq() *V0044JobAllocReq`

NewV0044JobAllocReq instantiates a new V0044JobAllocReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0044JobAllocReqWithDefaults

`func NewV0044JobAllocReqWithDefaults() *V0044JobAllocReq`

NewV0044JobAllocReqWithDefaults instantiates a new V0044JobAllocReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHetjob

`func (o *V0044JobAllocReq) GetHetjob() []V0044JobDescMsg`

GetHetjob returns the Hetjob field if non-nil, zero value otherwise.

### GetHetjobOk

`func (o *V0044JobAllocReq) GetHetjobOk() (*[]V0044JobDescMsg, bool)`

GetHetjobOk returns a tuple with the Hetjob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHetjob

`func (o *V0044JobAllocReq) SetHetjob(v []V0044JobDescMsg)`

SetHetjob sets Hetjob field to given value.

### HasHetjob

`func (o *V0044JobAllocReq) HasHetjob() bool`

HasHetjob returns a boolean if a field has been set.

### GetJob

`func (o *V0044JobAllocReq) GetJob() V0044JobDescMsg`

GetJob returns the Job field if non-nil, zero value otherwise.

### GetJobOk

`func (o *V0044JobAllocReq) GetJobOk() (*V0044JobDescMsg, bool)`

GetJobOk returns a tuple with the Job field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJob

`func (o *V0044JobAllocReq) SetJob(v V0044JobDescMsg)`

SetJob sets Job field to given value.

### HasJob

`func (o *V0044JobAllocReq) HasJob() bool`

HasJob returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


