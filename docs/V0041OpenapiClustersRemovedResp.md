# V0041OpenapiClustersRemovedResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeletedClusters** | **[]string** | deleted_clusters | 
**Meta** | Pointer to [**V0041OpenapiSlurmdbdJobsRespMeta**](V0041OpenapiSlurmdbdJobsRespMeta.md) |  | [optional] 
**Errors** | Pointer to [**[]V0041OpenapiSlurmdbdJobsRespErrorsInner**](V0041OpenapiSlurmdbdJobsRespErrorsInner.md) | Query errors | [optional] 
**Warnings** | Pointer to [**[]V0041OpenapiSlurmdbdJobsRespWarningsInner**](V0041OpenapiSlurmdbdJobsRespWarningsInner.md) | Query warnings | [optional] 

## Methods

### NewV0041OpenapiClustersRemovedResp

`func NewV0041OpenapiClustersRemovedResp(deletedClusters []string, ) *V0041OpenapiClustersRemovedResp`

NewV0041OpenapiClustersRemovedResp instantiates a new V0041OpenapiClustersRemovedResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0041OpenapiClustersRemovedRespWithDefaults

`func NewV0041OpenapiClustersRemovedRespWithDefaults() *V0041OpenapiClustersRemovedResp`

NewV0041OpenapiClustersRemovedRespWithDefaults instantiates a new V0041OpenapiClustersRemovedResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeletedClusters

`func (o *V0041OpenapiClustersRemovedResp) GetDeletedClusters() []string`

GetDeletedClusters returns the DeletedClusters field if non-nil, zero value otherwise.

### GetDeletedClustersOk

`func (o *V0041OpenapiClustersRemovedResp) GetDeletedClustersOk() (*[]string, bool)`

GetDeletedClustersOk returns a tuple with the DeletedClusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedClusters

`func (o *V0041OpenapiClustersRemovedResp) SetDeletedClusters(v []string)`

SetDeletedClusters sets DeletedClusters field to given value.


### GetMeta

`func (o *V0041OpenapiClustersRemovedResp) GetMeta() V0041OpenapiSlurmdbdJobsRespMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *V0041OpenapiClustersRemovedResp) GetMetaOk() (*V0041OpenapiSlurmdbdJobsRespMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *V0041OpenapiClustersRemovedResp) SetMeta(v V0041OpenapiSlurmdbdJobsRespMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *V0041OpenapiClustersRemovedResp) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *V0041OpenapiClustersRemovedResp) GetErrors() []V0041OpenapiSlurmdbdJobsRespErrorsInner`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V0041OpenapiClustersRemovedResp) GetErrorsOk() (*[]V0041OpenapiSlurmdbdJobsRespErrorsInner, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V0041OpenapiClustersRemovedResp) SetErrors(v []V0041OpenapiSlurmdbdJobsRespErrorsInner)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V0041OpenapiClustersRemovedResp) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *V0041OpenapiClustersRemovedResp) GetWarnings() []V0041OpenapiSlurmdbdJobsRespWarningsInner`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *V0041OpenapiClustersRemovedResp) GetWarningsOk() (*[]V0041OpenapiSlurmdbdJobsRespWarningsInner, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *V0041OpenapiClustersRemovedResp) SetWarnings(v []V0041OpenapiSlurmdbdJobsRespWarningsInner)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *V0041OpenapiClustersRemovedResp) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


