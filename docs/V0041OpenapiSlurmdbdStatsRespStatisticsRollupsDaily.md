# V0041OpenapiSlurmdbdStatsRespStatisticsRollupsDaily

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | number of daily rollups since last_run | [optional] 
**LastRun** | Pointer to **int64** | Last time daily rollup ran (UNIX timestamp) | [optional] 
**Duration** | Pointer to [**V0041OpenapiSlurmdbdStatsRespStatisticsRollupsDailyDuration**](V0041OpenapiSlurmdbdStatsRespStatisticsRollupsDailyDuration.md) |  | [optional] 

## Methods

### NewV0041OpenapiSlurmdbdStatsRespStatisticsRollupsDaily

`func NewV0041OpenapiSlurmdbdStatsRespStatisticsRollupsDaily() *V0041OpenapiSlurmdbdStatsRespStatisticsRollupsDaily`

NewV0041OpenapiSlurmdbdStatsRespStatisticsRollupsDaily instantiates a new V0041OpenapiSlurmdbdStatsRespStatisticsRollupsDaily object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0041OpenapiSlurmdbdStatsRespStatisticsRollupsDailyWithDefaults

`func NewV0041OpenapiSlurmdbdStatsRespStatisticsRollupsDailyWithDefaults() *V0041OpenapiSlurmdbdStatsRespStatisticsRollupsDaily`

NewV0041OpenapiSlurmdbdStatsRespStatisticsRollupsDailyWithDefaults instantiates a new V0041OpenapiSlurmdbdStatsRespStatisticsRollupsDaily object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *V0041OpenapiSlurmdbdStatsRespStatisticsRollupsDaily) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *V0041OpenapiSlurmdbdStatsRespStatisticsRollupsDaily) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *V0041OpenapiSlurmdbdStatsRespStatisticsRollupsDaily) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *V0041OpenapiSlurmdbdStatsRespStatisticsRollupsDaily) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetLastRun

`func (o *V0041OpenapiSlurmdbdStatsRespStatisticsRollupsDaily) GetLastRun() int64`

GetLastRun returns the LastRun field if non-nil, zero value otherwise.

### GetLastRunOk

`func (o *V0041OpenapiSlurmdbdStatsRespStatisticsRollupsDaily) GetLastRunOk() (*int64, bool)`

GetLastRunOk returns a tuple with the LastRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRun

`func (o *V0041OpenapiSlurmdbdStatsRespStatisticsRollupsDaily) SetLastRun(v int64)`

SetLastRun sets LastRun field to given value.

### HasLastRun

`func (o *V0041OpenapiSlurmdbdStatsRespStatisticsRollupsDaily) HasLastRun() bool`

HasLastRun returns a boolean if a field has been set.

### GetDuration

`func (o *V0041OpenapiSlurmdbdStatsRespStatisticsRollupsDaily) GetDuration() V0041OpenapiSlurmdbdStatsRespStatisticsRollupsDailyDuration`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *V0041OpenapiSlurmdbdStatsRespStatisticsRollupsDaily) GetDurationOk() (*V0041OpenapiSlurmdbdStatsRespStatisticsRollupsDailyDuration, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *V0041OpenapiSlurmdbdStatsRespStatisticsRollupsDaily) SetDuration(v V0041OpenapiSlurmdbdStatsRespStatisticsRollupsDailyDuration)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *V0041OpenapiSlurmdbdStatsRespStatisticsRollupsDaily) HasDuration() bool`

HasDuration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


