# V0044StatsMsgRpcType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TypeId** | **int32** | Message type as integer | 
**MessageType** | **string** | Message type as string (Slurm RPC message type) | 
**Count** | **int32** | Number of RPCs received | 
**Queued** | **int32** | Number of RPCs queued | 
**Dropped** | **int64** | Number of RPCs dropped | 
**CycleLast** | **int32** | Number of RPCs processed within the last RPC queue cycle | 
**CycleMax** | **int32** | Maximum number of RPCs processed within a RPC queue cycle since start | 
**TotalTime** | **int64** | Total time spent processing RPC in seconds | 
**AverageTime** | [**V0044Uint64NoValStruct**](V0044Uint64NoValStruct.md) |  | 

## Methods

### NewV0044StatsMsgRpcType

`func NewV0044StatsMsgRpcType(typeId int32, messageType string, count int32, queued int32, dropped int64, cycleLast int32, cycleMax int32, totalTime int64, averageTime V0044Uint64NoValStruct, ) *V0044StatsMsgRpcType`

NewV0044StatsMsgRpcType instantiates a new V0044StatsMsgRpcType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0044StatsMsgRpcTypeWithDefaults

`func NewV0044StatsMsgRpcTypeWithDefaults() *V0044StatsMsgRpcType`

NewV0044StatsMsgRpcTypeWithDefaults instantiates a new V0044StatsMsgRpcType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTypeId

`func (o *V0044StatsMsgRpcType) GetTypeId() int32`

GetTypeId returns the TypeId field if non-nil, zero value otherwise.

### GetTypeIdOk

`func (o *V0044StatsMsgRpcType) GetTypeIdOk() (*int32, bool)`

GetTypeIdOk returns a tuple with the TypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeId

`func (o *V0044StatsMsgRpcType) SetTypeId(v int32)`

SetTypeId sets TypeId field to given value.


### GetMessageType

`func (o *V0044StatsMsgRpcType) GetMessageType() string`

GetMessageType returns the MessageType field if non-nil, zero value otherwise.

### GetMessageTypeOk

`func (o *V0044StatsMsgRpcType) GetMessageTypeOk() (*string, bool)`

GetMessageTypeOk returns a tuple with the MessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageType

`func (o *V0044StatsMsgRpcType) SetMessageType(v string)`

SetMessageType sets MessageType field to given value.


### GetCount

`func (o *V0044StatsMsgRpcType) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *V0044StatsMsgRpcType) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *V0044StatsMsgRpcType) SetCount(v int32)`

SetCount sets Count field to given value.


### GetQueued

`func (o *V0044StatsMsgRpcType) GetQueued() int32`

GetQueued returns the Queued field if non-nil, zero value otherwise.

### GetQueuedOk

`func (o *V0044StatsMsgRpcType) GetQueuedOk() (*int32, bool)`

GetQueuedOk returns a tuple with the Queued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueued

`func (o *V0044StatsMsgRpcType) SetQueued(v int32)`

SetQueued sets Queued field to given value.


### GetDropped

`func (o *V0044StatsMsgRpcType) GetDropped() int64`

GetDropped returns the Dropped field if non-nil, zero value otherwise.

### GetDroppedOk

`func (o *V0044StatsMsgRpcType) GetDroppedOk() (*int64, bool)`

GetDroppedOk returns a tuple with the Dropped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropped

`func (o *V0044StatsMsgRpcType) SetDropped(v int64)`

SetDropped sets Dropped field to given value.


### GetCycleLast

`func (o *V0044StatsMsgRpcType) GetCycleLast() int32`

GetCycleLast returns the CycleLast field if non-nil, zero value otherwise.

### GetCycleLastOk

`func (o *V0044StatsMsgRpcType) GetCycleLastOk() (*int32, bool)`

GetCycleLastOk returns a tuple with the CycleLast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCycleLast

`func (o *V0044StatsMsgRpcType) SetCycleLast(v int32)`

SetCycleLast sets CycleLast field to given value.


### GetCycleMax

`func (o *V0044StatsMsgRpcType) GetCycleMax() int32`

GetCycleMax returns the CycleMax field if non-nil, zero value otherwise.

### GetCycleMaxOk

`func (o *V0044StatsMsgRpcType) GetCycleMaxOk() (*int32, bool)`

GetCycleMaxOk returns a tuple with the CycleMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCycleMax

`func (o *V0044StatsMsgRpcType) SetCycleMax(v int32)`

SetCycleMax sets CycleMax field to given value.


### GetTotalTime

`func (o *V0044StatsMsgRpcType) GetTotalTime() int64`

GetTotalTime returns the TotalTime field if non-nil, zero value otherwise.

### GetTotalTimeOk

`func (o *V0044StatsMsgRpcType) GetTotalTimeOk() (*int64, bool)`

GetTotalTimeOk returns a tuple with the TotalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTime

`func (o *V0044StatsMsgRpcType) SetTotalTime(v int64)`

SetTotalTime sets TotalTime field to given value.


### GetAverageTime

`func (o *V0044StatsMsgRpcType) GetAverageTime() V0044Uint64NoValStruct`

GetAverageTime returns the AverageTime field if non-nil, zero value otherwise.

### GetAverageTimeOk

`func (o *V0044StatsMsgRpcType) GetAverageTimeOk() (*V0044Uint64NoValStruct, bool)`

GetAverageTimeOk returns a tuple with the AverageTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageTime

`func (o *V0044StatsMsgRpcType) SetAverageTime(v V0044Uint64NoValStruct)`

SetAverageTime sets AverageTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


