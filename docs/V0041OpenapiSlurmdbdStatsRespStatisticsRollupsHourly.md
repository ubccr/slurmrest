# V0041OpenapiSlurmdbdStatsRespStatisticsRollupsHourly

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | number of hourly rollups since last_run | [optional] 
**LastRun** | Pointer to **int64** | Last time hourly rollup ran (UNIX timestamp) | [optional] 
**Duration** | Pointer to [**V0041OpenapiSlurmdbdStatsRespStatisticsRollupsHourlyDuration**](V0041OpenapiSlurmdbdStatsRespStatisticsRollupsHourlyDuration.md) |  | [optional] 

## Methods

### NewV0041OpenapiSlurmdbdStatsRespStatisticsRollupsHourly

`func NewV0041OpenapiSlurmdbdStatsRespStatisticsRollupsHourly() *V0041OpenapiSlurmdbdStatsRespStatisticsRollupsHourly`

NewV0041OpenapiSlurmdbdStatsRespStatisticsRollupsHourly instantiates a new V0041OpenapiSlurmdbdStatsRespStatisticsRollupsHourly object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0041OpenapiSlurmdbdStatsRespStatisticsRollupsHourlyWithDefaults

`func NewV0041OpenapiSlurmdbdStatsRespStatisticsRollupsHourlyWithDefaults() *V0041OpenapiSlurmdbdStatsRespStatisticsRollupsHourly`

NewV0041OpenapiSlurmdbdStatsRespStatisticsRollupsHourlyWithDefaults instantiates a new V0041OpenapiSlurmdbdStatsRespStatisticsRollupsHourly object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *V0041OpenapiSlurmdbdStatsRespStatisticsRollupsHourly) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *V0041OpenapiSlurmdbdStatsRespStatisticsRollupsHourly) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *V0041OpenapiSlurmdbdStatsRespStatisticsRollupsHourly) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *V0041OpenapiSlurmdbdStatsRespStatisticsRollupsHourly) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetLastRun

`func (o *V0041OpenapiSlurmdbdStatsRespStatisticsRollupsHourly) GetLastRun() int64`

GetLastRun returns the LastRun field if non-nil, zero value otherwise.

### GetLastRunOk

`func (o *V0041OpenapiSlurmdbdStatsRespStatisticsRollupsHourly) GetLastRunOk() (*int64, bool)`

GetLastRunOk returns a tuple with the LastRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRun

`func (o *V0041OpenapiSlurmdbdStatsRespStatisticsRollupsHourly) SetLastRun(v int64)`

SetLastRun sets LastRun field to given value.

### HasLastRun

`func (o *V0041OpenapiSlurmdbdStatsRespStatisticsRollupsHourly) HasLastRun() bool`

HasLastRun returns a boolean if a field has been set.

### GetDuration

`func (o *V0041OpenapiSlurmdbdStatsRespStatisticsRollupsHourly) GetDuration() V0041OpenapiSlurmdbdStatsRespStatisticsRollupsHourlyDuration`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *V0041OpenapiSlurmdbdStatsRespStatisticsRollupsHourly) GetDurationOk() (*V0041OpenapiSlurmdbdStatsRespStatisticsRollupsHourlyDuration, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *V0041OpenapiSlurmdbdStatsRespStatisticsRollupsHourly) SetDuration(v V0041OpenapiSlurmdbdStatsRespStatisticsRollupsHourlyDuration)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *V0041OpenapiSlurmdbdStatsRespStatisticsRollupsHourly) HasDuration() bool`

HasDuration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


