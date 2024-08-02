# V0041OpenapiJobInfoResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Jobs** | [**[]V0041OpenapiJobInfoRespJobsInner**](V0041OpenapiJobInfoRespJobsInner.md) | list of jobs | 
**LastBackfill** | [**V0041OpenapiJobInfoRespLastBackfill**](V0041OpenapiJobInfoRespLastBackfill.md) |  | 
**LastUpdate** | [**V0041OpenapiJobInfoRespLastUpdate**](V0041OpenapiJobInfoRespLastUpdate.md) |  | 
**Meta** | Pointer to [**V0041OpenapiSlurmdbdJobsRespMeta**](V0041OpenapiSlurmdbdJobsRespMeta.md) |  | [optional] 
**Errors** | Pointer to [**[]V0041OpenapiSlurmdbdJobsRespErrorsInner**](V0041OpenapiSlurmdbdJobsRespErrorsInner.md) | Query errors | [optional] 
**Warnings** | Pointer to [**[]V0041OpenapiSlurmdbdJobsRespWarningsInner**](V0041OpenapiSlurmdbdJobsRespWarningsInner.md) | Query warnings | [optional] 

## Methods

### NewV0041OpenapiJobInfoResp

`func NewV0041OpenapiJobInfoResp(jobs []V0041OpenapiJobInfoRespJobsInner, lastBackfill V0041OpenapiJobInfoRespLastBackfill, lastUpdate V0041OpenapiJobInfoRespLastUpdate, ) *V0041OpenapiJobInfoResp`

NewV0041OpenapiJobInfoResp instantiates a new V0041OpenapiJobInfoResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0041OpenapiJobInfoRespWithDefaults

`func NewV0041OpenapiJobInfoRespWithDefaults() *V0041OpenapiJobInfoResp`

NewV0041OpenapiJobInfoRespWithDefaults instantiates a new V0041OpenapiJobInfoResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobs

`func (o *V0041OpenapiJobInfoResp) GetJobs() []V0041OpenapiJobInfoRespJobsInner`

GetJobs returns the Jobs field if non-nil, zero value otherwise.

### GetJobsOk

`func (o *V0041OpenapiJobInfoResp) GetJobsOk() (*[]V0041OpenapiJobInfoRespJobsInner, bool)`

GetJobsOk returns a tuple with the Jobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobs

`func (o *V0041OpenapiJobInfoResp) SetJobs(v []V0041OpenapiJobInfoRespJobsInner)`

SetJobs sets Jobs field to given value.


### GetLastBackfill

`func (o *V0041OpenapiJobInfoResp) GetLastBackfill() V0041OpenapiJobInfoRespLastBackfill`

GetLastBackfill returns the LastBackfill field if non-nil, zero value otherwise.

### GetLastBackfillOk

`func (o *V0041OpenapiJobInfoResp) GetLastBackfillOk() (*V0041OpenapiJobInfoRespLastBackfill, bool)`

GetLastBackfillOk returns a tuple with the LastBackfill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastBackfill

`func (o *V0041OpenapiJobInfoResp) SetLastBackfill(v V0041OpenapiJobInfoRespLastBackfill)`

SetLastBackfill sets LastBackfill field to given value.


### GetLastUpdate

`func (o *V0041OpenapiJobInfoResp) GetLastUpdate() V0041OpenapiJobInfoRespLastUpdate`

GetLastUpdate returns the LastUpdate field if non-nil, zero value otherwise.

### GetLastUpdateOk

`func (o *V0041OpenapiJobInfoResp) GetLastUpdateOk() (*V0041OpenapiJobInfoRespLastUpdate, bool)`

GetLastUpdateOk returns a tuple with the LastUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdate

`func (o *V0041OpenapiJobInfoResp) SetLastUpdate(v V0041OpenapiJobInfoRespLastUpdate)`

SetLastUpdate sets LastUpdate field to given value.


### GetMeta

`func (o *V0041OpenapiJobInfoResp) GetMeta() V0041OpenapiSlurmdbdJobsRespMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *V0041OpenapiJobInfoResp) GetMetaOk() (*V0041OpenapiSlurmdbdJobsRespMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *V0041OpenapiJobInfoResp) SetMeta(v V0041OpenapiSlurmdbdJobsRespMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *V0041OpenapiJobInfoResp) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetErrors

`func (o *V0041OpenapiJobInfoResp) GetErrors() []V0041OpenapiSlurmdbdJobsRespErrorsInner`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V0041OpenapiJobInfoResp) GetErrorsOk() (*[]V0041OpenapiSlurmdbdJobsRespErrorsInner, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V0041OpenapiJobInfoResp) SetErrors(v []V0041OpenapiSlurmdbdJobsRespErrorsInner)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V0041OpenapiJobInfoResp) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *V0041OpenapiJobInfoResp) GetWarnings() []V0041OpenapiSlurmdbdJobsRespWarningsInner`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *V0041OpenapiJobInfoResp) GetWarningsOk() (*[]V0041OpenapiSlurmdbdJobsRespWarningsInner, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *V0041OpenapiJobInfoResp) SetWarnings(v []V0041OpenapiSlurmdbdJobsRespWarningsInner)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *V0041OpenapiJobInfoResp) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


