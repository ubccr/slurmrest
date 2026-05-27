# V0044RollupStatsHourly

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | Number of hourly rollups since last_run | [optional] 
**LastRun** | Pointer to **int64** | Last time hourly rollup ran (UNIX timestamp) (UNIX timestamp or time string recognized by Slurm (e.g., &#39;[MM/DD[/YY]-]HH:MM[:SS]&#39;)) | [optional] 
**Duration** | Pointer to [**V0044RollupStatsHourlyDuration**](V0044RollupStatsHourlyDuration.md) |  | [optional] 

## Methods

### NewV0044RollupStatsHourly

`func NewV0044RollupStatsHourly() *V0044RollupStatsHourly`

NewV0044RollupStatsHourly instantiates a new V0044RollupStatsHourly object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0044RollupStatsHourlyWithDefaults

`func NewV0044RollupStatsHourlyWithDefaults() *V0044RollupStatsHourly`

NewV0044RollupStatsHourlyWithDefaults instantiates a new V0044RollupStatsHourly object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *V0044RollupStatsHourly) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *V0044RollupStatsHourly) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *V0044RollupStatsHourly) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *V0044RollupStatsHourly) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetLastRun

`func (o *V0044RollupStatsHourly) GetLastRun() int64`

GetLastRun returns the LastRun field if non-nil, zero value otherwise.

### GetLastRunOk

`func (o *V0044RollupStatsHourly) GetLastRunOk() (*int64, bool)`

GetLastRunOk returns a tuple with the LastRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRun

`func (o *V0044RollupStatsHourly) SetLastRun(v int64)`

SetLastRun sets LastRun field to given value.

### HasLastRun

`func (o *V0044RollupStatsHourly) HasLastRun() bool`

HasLastRun returns a boolean if a field has been set.

### GetDuration

`func (o *V0044RollupStatsHourly) GetDuration() V0044RollupStatsHourlyDuration`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *V0044RollupStatsHourly) GetDurationOk() (*V0044RollupStatsHourlyDuration, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *V0044RollupStatsHourly) SetDuration(v V0044RollupStatsHourlyDuration)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *V0044RollupStatsHourly) HasDuration() bool`

HasDuration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


