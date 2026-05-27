# V0044StatsRec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimeStart** | Pointer to **int64** | When data collection started (UNIX timestamp) (UNIX timestamp or time string recognized by Slurm (e.g., &#39;[MM/DD[/YY]-]HH:MM[:SS]&#39;)) | [optional] 
**Rollups** | Pointer to [**V0044RollupStats**](V0044RollupStats.md) |  | [optional] 
**RPCs** | Pointer to [**[]V0044StatsRpc**](V0044StatsRpc.md) |  | [optional] 
**Users** | Pointer to [**[]V0044StatsUser**](V0044StatsUser.md) |  | [optional] 

## Methods

### NewV0044StatsRec

`func NewV0044StatsRec() *V0044StatsRec`

NewV0044StatsRec instantiates a new V0044StatsRec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0044StatsRecWithDefaults

`func NewV0044StatsRecWithDefaults() *V0044StatsRec`

NewV0044StatsRecWithDefaults instantiates a new V0044StatsRec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeStart

`func (o *V0044StatsRec) GetTimeStart() int64`

GetTimeStart returns the TimeStart field if non-nil, zero value otherwise.

### GetTimeStartOk

`func (o *V0044StatsRec) GetTimeStartOk() (*int64, bool)`

GetTimeStartOk returns a tuple with the TimeStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStart

`func (o *V0044StatsRec) SetTimeStart(v int64)`

SetTimeStart sets TimeStart field to given value.

### HasTimeStart

`func (o *V0044StatsRec) HasTimeStart() bool`

HasTimeStart returns a boolean if a field has been set.

### GetRollups

`func (o *V0044StatsRec) GetRollups() V0044RollupStats`

GetRollups returns the Rollups field if non-nil, zero value otherwise.

### GetRollupsOk

`func (o *V0044StatsRec) GetRollupsOk() (*V0044RollupStats, bool)`

GetRollupsOk returns a tuple with the Rollups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollups

`func (o *V0044StatsRec) SetRollups(v V0044RollupStats)`

SetRollups sets Rollups field to given value.

### HasRollups

`func (o *V0044StatsRec) HasRollups() bool`

HasRollups returns a boolean if a field has been set.

### GetRPCs

`func (o *V0044StatsRec) GetRPCs() []V0044StatsRpc`

GetRPCs returns the RPCs field if non-nil, zero value otherwise.

### GetRPCsOk

`func (o *V0044StatsRec) GetRPCsOk() (*[]V0044StatsRpc, bool)`

GetRPCsOk returns a tuple with the RPCs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRPCs

`func (o *V0044StatsRec) SetRPCs(v []V0044StatsRpc)`

SetRPCs sets RPCs field to given value.

### HasRPCs

`func (o *V0044StatsRec) HasRPCs() bool`

HasRPCs returns a boolean if a field has been set.

### GetUsers

`func (o *V0044StatsRec) GetUsers() []V0044StatsUser`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *V0044StatsRec) GetUsersOk() (*[]V0044StatsUser, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *V0044StatsRec) SetUsers(v []V0044StatsUser)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *V0044StatsRec) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


