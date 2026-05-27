# V0044RollupStatsDaily

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | Number of daily rollups since last_run | [optional] 
**LastRun** | Pointer to **int64** | Last time daily rollup ran (UNIX timestamp) (UNIX timestamp or time string recognized by Slurm (e.g., &#39;[MM/DD[/YY]-]HH:MM[:SS]&#39;)) | [optional] 
**Duration** | Pointer to [**V0044RollupStatsDailyDuration**](V0044RollupStatsDailyDuration.md) |  | [optional] 

## Methods

### NewV0044RollupStatsDaily

`func NewV0044RollupStatsDaily() *V0044RollupStatsDaily`

NewV0044RollupStatsDaily instantiates a new V0044RollupStatsDaily object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0044RollupStatsDailyWithDefaults

`func NewV0044RollupStatsDailyWithDefaults() *V0044RollupStatsDaily`

NewV0044RollupStatsDailyWithDefaults instantiates a new V0044RollupStatsDaily object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *V0044RollupStatsDaily) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *V0044RollupStatsDaily) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *V0044RollupStatsDaily) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *V0044RollupStatsDaily) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetLastRun

`func (o *V0044RollupStatsDaily) GetLastRun() int64`

GetLastRun returns the LastRun field if non-nil, zero value otherwise.

### GetLastRunOk

`func (o *V0044RollupStatsDaily) GetLastRunOk() (*int64, bool)`

GetLastRunOk returns a tuple with the LastRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRun

`func (o *V0044RollupStatsDaily) SetLastRun(v int64)`

SetLastRun sets LastRun field to given value.

### HasLastRun

`func (o *V0044RollupStatsDaily) HasLastRun() bool`

HasLastRun returns a boolean if a field has been set.

### GetDuration

`func (o *V0044RollupStatsDaily) GetDuration() V0044RollupStatsDailyDuration`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *V0044RollupStatsDaily) GetDurationOk() (*V0044RollupStatsDailyDuration, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *V0044RollupStatsDaily) SetDuration(v V0044RollupStatsDailyDuration)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *V0044RollupStatsDaily) HasDuration() bool`

HasDuration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


