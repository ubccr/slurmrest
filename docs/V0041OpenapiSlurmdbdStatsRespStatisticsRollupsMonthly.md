# V0041OpenapiSlurmdbdStatsRespStatisticsRollupsMonthly

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | number of monthly rollups since last_run | [optional] 
**LastRun** | Pointer to **int64** | Last time monthly rollup ran (UNIX timestamp) | [optional] 
**Duration** | Pointer to [**V0041OpenapiSlurmdbdStatsRespStatisticsRollupsMonthlyDuration**](V0041OpenapiSlurmdbdStatsRespStatisticsRollupsMonthlyDuration.md) |  | [optional] 

## Methods

### NewV0041OpenapiSlurmdbdStatsRespStatisticsRollupsMonthly

`func NewV0041OpenapiSlurmdbdStatsRespStatisticsRollupsMonthly() *V0041OpenapiSlurmdbdStatsRespStatisticsRollupsMonthly`

NewV0041OpenapiSlurmdbdStatsRespStatisticsRollupsMonthly instantiates a new V0041OpenapiSlurmdbdStatsRespStatisticsRollupsMonthly object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0041OpenapiSlurmdbdStatsRespStatisticsRollupsMonthlyWithDefaults

`func NewV0041OpenapiSlurmdbdStatsRespStatisticsRollupsMonthlyWithDefaults() *V0041OpenapiSlurmdbdStatsRespStatisticsRollupsMonthly`

NewV0041OpenapiSlurmdbdStatsRespStatisticsRollupsMonthlyWithDefaults instantiates a new V0041OpenapiSlurmdbdStatsRespStatisticsRollupsMonthly object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *V0041OpenapiSlurmdbdStatsRespStatisticsRollupsMonthly) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *V0041OpenapiSlurmdbdStatsRespStatisticsRollupsMonthly) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *V0041OpenapiSlurmdbdStatsRespStatisticsRollupsMonthly) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *V0041OpenapiSlurmdbdStatsRespStatisticsRollupsMonthly) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetLastRun

`func (o *V0041OpenapiSlurmdbdStatsRespStatisticsRollupsMonthly) GetLastRun() int64`

GetLastRun returns the LastRun field if non-nil, zero value otherwise.

### GetLastRunOk

`func (o *V0041OpenapiSlurmdbdStatsRespStatisticsRollupsMonthly) GetLastRunOk() (*int64, bool)`

GetLastRunOk returns a tuple with the LastRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRun

`func (o *V0041OpenapiSlurmdbdStatsRespStatisticsRollupsMonthly) SetLastRun(v int64)`

SetLastRun sets LastRun field to given value.

### HasLastRun

`func (o *V0041OpenapiSlurmdbdStatsRespStatisticsRollupsMonthly) HasLastRun() bool`

HasLastRun returns a boolean if a field has been set.

### GetDuration

`func (o *V0041OpenapiSlurmdbdStatsRespStatisticsRollupsMonthly) GetDuration() V0041OpenapiSlurmdbdStatsRespStatisticsRollupsMonthlyDuration`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *V0041OpenapiSlurmdbdStatsRespStatisticsRollupsMonthly) GetDurationOk() (*V0041OpenapiSlurmdbdStatsRespStatisticsRollupsMonthlyDuration, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *V0041OpenapiSlurmdbdStatsRespStatisticsRollupsMonthly) SetDuration(v V0041OpenapiSlurmdbdStatsRespStatisticsRollupsMonthlyDuration)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *V0041OpenapiSlurmdbdStatsRespStatisticsRollupsMonthly) HasDuration() bool`

HasDuration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


