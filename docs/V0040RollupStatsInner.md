# V0040RollupStatsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | type | [optional] 
**LastRun** | Pointer to **int32** | Last time rollup ran (UNIX timestamp) | [optional] 
**MaxCycle** | Pointer to **int64** | longest rollup time (seconds) | [optional] 
**TotalTime** | Pointer to **int64** | total time spent doing rollups (seconds) | [optional] 
**TotalCycles** | Pointer to **int64** | number of rollups since last_run | [optional] 
**MeanCycles** | Pointer to **int64** | average time for rollup (seconds) | [optional] 

## Methods

### NewV0040RollupStatsInner

`func NewV0040RollupStatsInner() *V0040RollupStatsInner`

NewV0040RollupStatsInner instantiates a new V0040RollupStatsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0040RollupStatsInnerWithDefaults

`func NewV0040RollupStatsInnerWithDefaults() *V0040RollupStatsInner`

NewV0040RollupStatsInnerWithDefaults instantiates a new V0040RollupStatsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *V0040RollupStatsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V0040RollupStatsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V0040RollupStatsInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *V0040RollupStatsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLastRun

`func (o *V0040RollupStatsInner) GetLastRun() int32`

GetLastRun returns the LastRun field if non-nil, zero value otherwise.

### GetLastRunOk

`func (o *V0040RollupStatsInner) GetLastRunOk() (*int32, bool)`

GetLastRunOk returns a tuple with the LastRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRun

`func (o *V0040RollupStatsInner) SetLastRun(v int32)`

SetLastRun sets LastRun field to given value.

### HasLastRun

`func (o *V0040RollupStatsInner) HasLastRun() bool`

HasLastRun returns a boolean if a field has been set.

### GetMaxCycle

`func (o *V0040RollupStatsInner) GetMaxCycle() int64`

GetMaxCycle returns the MaxCycle field if non-nil, zero value otherwise.

### GetMaxCycleOk

`func (o *V0040RollupStatsInner) GetMaxCycleOk() (*int64, bool)`

GetMaxCycleOk returns a tuple with the MaxCycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCycle

`func (o *V0040RollupStatsInner) SetMaxCycle(v int64)`

SetMaxCycle sets MaxCycle field to given value.

### HasMaxCycle

`func (o *V0040RollupStatsInner) HasMaxCycle() bool`

HasMaxCycle returns a boolean if a field has been set.

### GetTotalTime

`func (o *V0040RollupStatsInner) GetTotalTime() int64`

GetTotalTime returns the TotalTime field if non-nil, zero value otherwise.

### GetTotalTimeOk

`func (o *V0040RollupStatsInner) GetTotalTimeOk() (*int64, bool)`

GetTotalTimeOk returns a tuple with the TotalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTime

`func (o *V0040RollupStatsInner) SetTotalTime(v int64)`

SetTotalTime sets TotalTime field to given value.

### HasTotalTime

`func (o *V0040RollupStatsInner) HasTotalTime() bool`

HasTotalTime returns a boolean if a field has been set.

### GetTotalCycles

`func (o *V0040RollupStatsInner) GetTotalCycles() int64`

GetTotalCycles returns the TotalCycles field if non-nil, zero value otherwise.

### GetTotalCyclesOk

`func (o *V0040RollupStatsInner) GetTotalCyclesOk() (*int64, bool)`

GetTotalCyclesOk returns a tuple with the TotalCycles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCycles

`func (o *V0040RollupStatsInner) SetTotalCycles(v int64)`

SetTotalCycles sets TotalCycles field to given value.

### HasTotalCycles

`func (o *V0040RollupStatsInner) HasTotalCycles() bool`

HasTotalCycles returns a boolean if a field has been set.

### GetMeanCycles

`func (o *V0040RollupStatsInner) GetMeanCycles() int64`

GetMeanCycles returns the MeanCycles field if non-nil, zero value otherwise.

### GetMeanCyclesOk

`func (o *V0040RollupStatsInner) GetMeanCyclesOk() (*int64, bool)`

GetMeanCyclesOk returns a tuple with the MeanCycles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeanCycles

`func (o *V0040RollupStatsInner) SetMeanCycles(v int64)`

SetMeanCycles sets MeanCycles field to given value.

### HasMeanCycles

`func (o *V0040RollupStatsInner) HasMeanCycles() bool`

HasMeanCycles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


