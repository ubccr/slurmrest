# V0044RollupStatsMonthly

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | Number of monthly rollups since last_run | [optional] 
**LastRun** | Pointer to **int64** | Last time monthly rollup ran (UNIX timestamp) (UNIX timestamp or time string recognized by Slurm (e.g., &#39;[MM/DD[/YY]-]HH:MM[:SS]&#39;)) | [optional] 
**Duration** | Pointer to [**V0044RollupStatsMonthlyDuration**](V0044RollupStatsMonthlyDuration.md) |  | [optional] 

## Methods

### NewV0044RollupStatsMonthly

`func NewV0044RollupStatsMonthly() *V0044RollupStatsMonthly`

NewV0044RollupStatsMonthly instantiates a new V0044RollupStatsMonthly object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0044RollupStatsMonthlyWithDefaults

`func NewV0044RollupStatsMonthlyWithDefaults() *V0044RollupStatsMonthly`

NewV0044RollupStatsMonthlyWithDefaults instantiates a new V0044RollupStatsMonthly object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *V0044RollupStatsMonthly) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *V0044RollupStatsMonthly) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *V0044RollupStatsMonthly) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *V0044RollupStatsMonthly) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetLastRun

`func (o *V0044RollupStatsMonthly) GetLastRun() int64`

GetLastRun returns the LastRun field if non-nil, zero value otherwise.

### GetLastRunOk

`func (o *V0044RollupStatsMonthly) GetLastRunOk() (*int64, bool)`

GetLastRunOk returns a tuple with the LastRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRun

`func (o *V0044RollupStatsMonthly) SetLastRun(v int64)`

SetLastRun sets LastRun field to given value.

### HasLastRun

`func (o *V0044RollupStatsMonthly) HasLastRun() bool`

HasLastRun returns a boolean if a field has been set.

### GetDuration

`func (o *V0044RollupStatsMonthly) GetDuration() V0044RollupStatsMonthlyDuration`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *V0044RollupStatsMonthly) GetDurationOk() (*V0044RollupStatsMonthlyDuration, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *V0044RollupStatsMonthly) SetDuration(v V0044RollupStatsMonthlyDuration)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *V0044RollupStatsMonthly) HasDuration() bool`

HasDuration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


