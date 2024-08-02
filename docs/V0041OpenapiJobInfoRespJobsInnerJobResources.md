# V0041OpenapiJobInfoRespJobsInnerJobResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SelectType** | **[]string** | Scheduling consumption resource selection type | 
**Nodes** | Pointer to [**V0041OpenapiJobInfoRespJobsInnerJobResourcesNodes**](V0041OpenapiJobInfoRespJobsInnerJobResourcesNodes.md) |  | [optional] 
**Cpus** | **int32** | Number of processors in the allocation | 
**ThreadsPerCore** | [**V0041OpenapiJobInfoRespJobsInnerJobResourcesThreadsPerCore**](V0041OpenapiJobInfoRespJobsInnerJobResourcesThreadsPerCore.md) |  | 

## Methods

### NewV0041OpenapiJobInfoRespJobsInnerJobResources

`func NewV0041OpenapiJobInfoRespJobsInnerJobResources(selectType []string, cpus int32, threadsPerCore V0041OpenapiJobInfoRespJobsInnerJobResourcesThreadsPerCore, ) *V0041OpenapiJobInfoRespJobsInnerJobResources`

NewV0041OpenapiJobInfoRespJobsInnerJobResources instantiates a new V0041OpenapiJobInfoRespJobsInnerJobResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0041OpenapiJobInfoRespJobsInnerJobResourcesWithDefaults

`func NewV0041OpenapiJobInfoRespJobsInnerJobResourcesWithDefaults() *V0041OpenapiJobInfoRespJobsInnerJobResources`

NewV0041OpenapiJobInfoRespJobsInnerJobResourcesWithDefaults instantiates a new V0041OpenapiJobInfoRespJobsInnerJobResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelectType

`func (o *V0041OpenapiJobInfoRespJobsInnerJobResources) GetSelectType() []string`

GetSelectType returns the SelectType field if non-nil, zero value otherwise.

### GetSelectTypeOk

`func (o *V0041OpenapiJobInfoRespJobsInnerJobResources) GetSelectTypeOk() (*[]string, bool)`

GetSelectTypeOk returns a tuple with the SelectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectType

`func (o *V0041OpenapiJobInfoRespJobsInnerJobResources) SetSelectType(v []string)`

SetSelectType sets SelectType field to given value.


### GetNodes

`func (o *V0041OpenapiJobInfoRespJobsInnerJobResources) GetNodes() V0041OpenapiJobInfoRespJobsInnerJobResourcesNodes`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *V0041OpenapiJobInfoRespJobsInnerJobResources) GetNodesOk() (*V0041OpenapiJobInfoRespJobsInnerJobResourcesNodes, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *V0041OpenapiJobInfoRespJobsInnerJobResources) SetNodes(v V0041OpenapiJobInfoRespJobsInnerJobResourcesNodes)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *V0041OpenapiJobInfoRespJobsInnerJobResources) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetCpus

`func (o *V0041OpenapiJobInfoRespJobsInnerJobResources) GetCpus() int32`

GetCpus returns the Cpus field if non-nil, zero value otherwise.

### GetCpusOk

`func (o *V0041OpenapiJobInfoRespJobsInnerJobResources) GetCpusOk() (*int32, bool)`

GetCpusOk returns a tuple with the Cpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpus

`func (o *V0041OpenapiJobInfoRespJobsInnerJobResources) SetCpus(v int32)`

SetCpus sets Cpus field to given value.


### GetThreadsPerCore

`func (o *V0041OpenapiJobInfoRespJobsInnerJobResources) GetThreadsPerCore() V0041OpenapiJobInfoRespJobsInnerJobResourcesThreadsPerCore`

GetThreadsPerCore returns the ThreadsPerCore field if non-nil, zero value otherwise.

### GetThreadsPerCoreOk

`func (o *V0041OpenapiJobInfoRespJobsInnerJobResources) GetThreadsPerCoreOk() (*V0041OpenapiJobInfoRespJobsInnerJobResourcesThreadsPerCore, bool)`

GetThreadsPerCoreOk returns a tuple with the ThreadsPerCore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadsPerCore

`func (o *V0041OpenapiJobInfoRespJobsInnerJobResources) SetThreadsPerCore(v V0041OpenapiJobInfoRespJobsInnerJobResourcesThreadsPerCore)`

SetThreadsPerCore sets ThreadsPerCore field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


