# V0041JobAllocReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hetjob** | Pointer to [**[]V0041JobDescMsg**](V0041JobDescMsg.md) | HetJob description | [optional] 
**Job** | Pointer to [**V0041JobDescMsg**](V0041JobDescMsg.md) |  | [optional] 

## Methods

### NewV0041JobAllocReq

`func NewV0041JobAllocReq() *V0041JobAllocReq`

NewV0041JobAllocReq instantiates a new V0041JobAllocReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0041JobAllocReqWithDefaults

`func NewV0041JobAllocReqWithDefaults() *V0041JobAllocReq`

NewV0041JobAllocReqWithDefaults instantiates a new V0041JobAllocReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHetjob

`func (o *V0041JobAllocReq) GetHetjob() []V0041JobDescMsg`

GetHetjob returns the Hetjob field if non-nil, zero value otherwise.

### GetHetjobOk

`func (o *V0041JobAllocReq) GetHetjobOk() (*[]V0041JobDescMsg, bool)`

GetHetjobOk returns a tuple with the Hetjob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHetjob

`func (o *V0041JobAllocReq) SetHetjob(v []V0041JobDescMsg)`

SetHetjob sets Hetjob field to given value.

### HasHetjob

`func (o *V0041JobAllocReq) HasHetjob() bool`

HasHetjob returns a boolean if a field has been set.

### GetJob

`func (o *V0041JobAllocReq) GetJob() V0041JobDescMsg`

GetJob returns the Job field if non-nil, zero value otherwise.

### GetJobOk

`func (o *V0041JobAllocReq) GetJobOk() (*V0041JobDescMsg, bool)`

GetJobOk returns a tuple with the Job field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJob

`func (o *V0041JobAllocReq) SetJob(v V0041JobDescMsg)`

SetJob sets Job field to given value.

### HasJob

`func (o *V0041JobAllocReq) HasJob() bool`

HasJob returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


