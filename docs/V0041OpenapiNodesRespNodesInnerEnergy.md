# V0041OpenapiNodesRespNodesInnerEnergy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AverageWatts** | Pointer to **int32** |  | [optional] 
**BaseConsumedEnergy** | Pointer to **int64** |  | [optional] 
**ConsumedEnergy** | Pointer to **int64** |  | [optional] 
**CurrentWatts** | Pointer to [**V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId**](V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId.md) |  | [optional] 
**PreviousConsumedEnergy** | Pointer to **int64** |  | [optional] 
**LastCollected** | Pointer to **int64** |  | [optional] 

## Methods

### NewV0041OpenapiNodesRespNodesInnerEnergy

`func NewV0041OpenapiNodesRespNodesInnerEnergy() *V0041OpenapiNodesRespNodesInnerEnergy`

NewV0041OpenapiNodesRespNodesInnerEnergy instantiates a new V0041OpenapiNodesRespNodesInnerEnergy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0041OpenapiNodesRespNodesInnerEnergyWithDefaults

`func NewV0041OpenapiNodesRespNodesInnerEnergyWithDefaults() *V0041OpenapiNodesRespNodesInnerEnergy`

NewV0041OpenapiNodesRespNodesInnerEnergyWithDefaults instantiates a new V0041OpenapiNodesRespNodesInnerEnergy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAverageWatts

`func (o *V0041OpenapiNodesRespNodesInnerEnergy) GetAverageWatts() int32`

GetAverageWatts returns the AverageWatts field if non-nil, zero value otherwise.

### GetAverageWattsOk

`func (o *V0041OpenapiNodesRespNodesInnerEnergy) GetAverageWattsOk() (*int32, bool)`

GetAverageWattsOk returns a tuple with the AverageWatts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageWatts

`func (o *V0041OpenapiNodesRespNodesInnerEnergy) SetAverageWatts(v int32)`

SetAverageWatts sets AverageWatts field to given value.

### HasAverageWatts

`func (o *V0041OpenapiNodesRespNodesInnerEnergy) HasAverageWatts() bool`

HasAverageWatts returns a boolean if a field has been set.

### GetBaseConsumedEnergy

`func (o *V0041OpenapiNodesRespNodesInnerEnergy) GetBaseConsumedEnergy() int64`

GetBaseConsumedEnergy returns the BaseConsumedEnergy field if non-nil, zero value otherwise.

### GetBaseConsumedEnergyOk

`func (o *V0041OpenapiNodesRespNodesInnerEnergy) GetBaseConsumedEnergyOk() (*int64, bool)`

GetBaseConsumedEnergyOk returns a tuple with the BaseConsumedEnergy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseConsumedEnergy

`func (o *V0041OpenapiNodesRespNodesInnerEnergy) SetBaseConsumedEnergy(v int64)`

SetBaseConsumedEnergy sets BaseConsumedEnergy field to given value.

### HasBaseConsumedEnergy

`func (o *V0041OpenapiNodesRespNodesInnerEnergy) HasBaseConsumedEnergy() bool`

HasBaseConsumedEnergy returns a boolean if a field has been set.

### GetConsumedEnergy

`func (o *V0041OpenapiNodesRespNodesInnerEnergy) GetConsumedEnergy() int64`

GetConsumedEnergy returns the ConsumedEnergy field if non-nil, zero value otherwise.

### GetConsumedEnergyOk

`func (o *V0041OpenapiNodesRespNodesInnerEnergy) GetConsumedEnergyOk() (*int64, bool)`

GetConsumedEnergyOk returns a tuple with the ConsumedEnergy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumedEnergy

`func (o *V0041OpenapiNodesRespNodesInnerEnergy) SetConsumedEnergy(v int64)`

SetConsumedEnergy sets ConsumedEnergy field to given value.

### HasConsumedEnergy

`func (o *V0041OpenapiNodesRespNodesInnerEnergy) HasConsumedEnergy() bool`

HasConsumedEnergy returns a boolean if a field has been set.

### GetCurrentWatts

`func (o *V0041OpenapiNodesRespNodesInnerEnergy) GetCurrentWatts() V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId`

GetCurrentWatts returns the CurrentWatts field if non-nil, zero value otherwise.

### GetCurrentWattsOk

`func (o *V0041OpenapiNodesRespNodesInnerEnergy) GetCurrentWattsOk() (*V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId, bool)`

GetCurrentWattsOk returns a tuple with the CurrentWatts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentWatts

`func (o *V0041OpenapiNodesRespNodesInnerEnergy) SetCurrentWatts(v V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId)`

SetCurrentWatts sets CurrentWatts field to given value.

### HasCurrentWatts

`func (o *V0041OpenapiNodesRespNodesInnerEnergy) HasCurrentWatts() bool`

HasCurrentWatts returns a boolean if a field has been set.

### GetPreviousConsumedEnergy

`func (o *V0041OpenapiNodesRespNodesInnerEnergy) GetPreviousConsumedEnergy() int64`

GetPreviousConsumedEnergy returns the PreviousConsumedEnergy field if non-nil, zero value otherwise.

### GetPreviousConsumedEnergyOk

`func (o *V0041OpenapiNodesRespNodesInnerEnergy) GetPreviousConsumedEnergyOk() (*int64, bool)`

GetPreviousConsumedEnergyOk returns a tuple with the PreviousConsumedEnergy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousConsumedEnergy

`func (o *V0041OpenapiNodesRespNodesInnerEnergy) SetPreviousConsumedEnergy(v int64)`

SetPreviousConsumedEnergy sets PreviousConsumedEnergy field to given value.

### HasPreviousConsumedEnergy

`func (o *V0041OpenapiNodesRespNodesInnerEnergy) HasPreviousConsumedEnergy() bool`

HasPreviousConsumedEnergy returns a boolean if a field has been set.

### GetLastCollected

`func (o *V0041OpenapiNodesRespNodesInnerEnergy) GetLastCollected() int64`

GetLastCollected returns the LastCollected field if non-nil, zero value otherwise.

### GetLastCollectedOk

`func (o *V0041OpenapiNodesRespNodesInnerEnergy) GetLastCollectedOk() (*int64, bool)`

GetLastCollectedOk returns a tuple with the LastCollected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCollected

`func (o *V0041OpenapiNodesRespNodesInnerEnergy) SetLastCollected(v int64)`

SetLastCollected sets LastCollected field to given value.

### HasLastCollected

`func (o *V0041OpenapiNodesRespNodesInnerEnergy) HasLastCollected() bool`

HasLastCollected returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


