# V0041OpenapiSlurmdbdStatsRespStatistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimeStart** | Pointer to **int64** |  | [optional] 
**Rollups** | Pointer to [**V0041OpenapiSlurmdbdStatsRespStatisticsRollups**](V0041OpenapiSlurmdbdStatsRespStatisticsRollups.md) |  | [optional] 
**RPCs** | Pointer to [**[]V0041OpenapiSlurmdbdStatsRespStatisticsRPCsInner**](V0041OpenapiSlurmdbdStatsRespStatisticsRPCsInner.md) |  | [optional] 
**Users** | Pointer to [**[]V0041OpenapiSlurmdbdStatsRespStatisticsUsersInner**](V0041OpenapiSlurmdbdStatsRespStatisticsUsersInner.md) |  | [optional] 

## Methods

### NewV0041OpenapiSlurmdbdStatsRespStatistics

`func NewV0041OpenapiSlurmdbdStatsRespStatistics() *V0041OpenapiSlurmdbdStatsRespStatistics`

NewV0041OpenapiSlurmdbdStatsRespStatistics instantiates a new V0041OpenapiSlurmdbdStatsRespStatistics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0041OpenapiSlurmdbdStatsRespStatisticsWithDefaults

`func NewV0041OpenapiSlurmdbdStatsRespStatisticsWithDefaults() *V0041OpenapiSlurmdbdStatsRespStatistics`

NewV0041OpenapiSlurmdbdStatsRespStatisticsWithDefaults instantiates a new V0041OpenapiSlurmdbdStatsRespStatistics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeStart

`func (o *V0041OpenapiSlurmdbdStatsRespStatistics) GetTimeStart() int64`

GetTimeStart returns the TimeStart field if non-nil, zero value otherwise.

### GetTimeStartOk

`func (o *V0041OpenapiSlurmdbdStatsRespStatistics) GetTimeStartOk() (*int64, bool)`

GetTimeStartOk returns a tuple with the TimeStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStart

`func (o *V0041OpenapiSlurmdbdStatsRespStatistics) SetTimeStart(v int64)`

SetTimeStart sets TimeStart field to given value.

### HasTimeStart

`func (o *V0041OpenapiSlurmdbdStatsRespStatistics) HasTimeStart() bool`

HasTimeStart returns a boolean if a field has been set.

### GetRollups

`func (o *V0041OpenapiSlurmdbdStatsRespStatistics) GetRollups() V0041OpenapiSlurmdbdStatsRespStatisticsRollups`

GetRollups returns the Rollups field if non-nil, zero value otherwise.

### GetRollupsOk

`func (o *V0041OpenapiSlurmdbdStatsRespStatistics) GetRollupsOk() (*V0041OpenapiSlurmdbdStatsRespStatisticsRollups, bool)`

GetRollupsOk returns a tuple with the Rollups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollups

`func (o *V0041OpenapiSlurmdbdStatsRespStatistics) SetRollups(v V0041OpenapiSlurmdbdStatsRespStatisticsRollups)`

SetRollups sets Rollups field to given value.

### HasRollups

`func (o *V0041OpenapiSlurmdbdStatsRespStatistics) HasRollups() bool`

HasRollups returns a boolean if a field has been set.

### GetRPCs

`func (o *V0041OpenapiSlurmdbdStatsRespStatistics) GetRPCs() []V0041OpenapiSlurmdbdStatsRespStatisticsRPCsInner`

GetRPCs returns the RPCs field if non-nil, zero value otherwise.

### GetRPCsOk

`func (o *V0041OpenapiSlurmdbdStatsRespStatistics) GetRPCsOk() (*[]V0041OpenapiSlurmdbdStatsRespStatisticsRPCsInner, bool)`

GetRPCsOk returns a tuple with the RPCs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRPCs

`func (o *V0041OpenapiSlurmdbdStatsRespStatistics) SetRPCs(v []V0041OpenapiSlurmdbdStatsRespStatisticsRPCsInner)`

SetRPCs sets RPCs field to given value.

### HasRPCs

`func (o *V0041OpenapiSlurmdbdStatsRespStatistics) HasRPCs() bool`

HasRPCs returns a boolean if a field has been set.

### GetUsers

`func (o *V0041OpenapiSlurmdbdStatsRespStatistics) GetUsers() []V0041OpenapiSlurmdbdStatsRespStatisticsUsersInner`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *V0041OpenapiSlurmdbdStatsRespStatistics) GetUsersOk() (*[]V0041OpenapiSlurmdbdStatsRespStatisticsUsersInner, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *V0041OpenapiSlurmdbdStatsRespStatistics) SetUsers(v []V0041OpenapiSlurmdbdStatsRespStatisticsUsersInner)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *V0041OpenapiSlurmdbdStatsRespStatistics) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


