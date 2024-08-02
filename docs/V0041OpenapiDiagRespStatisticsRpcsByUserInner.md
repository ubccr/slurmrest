# V0041OpenapiDiagRespStatisticsRpcsByUserInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **int32** | user id (numeric) | 
**User** | **string** | user name | 
**Count** | **int32** | Number of RPCs received | 
**TotalTime** | **int64** | Total time spent processing RPC in seconds | 
**AverageTime** | [**V0041OpenapiDiagRespStatisticsRpcsByMessageTypeInnerAverageTime**](V0041OpenapiDiagRespStatisticsRpcsByMessageTypeInnerAverageTime.md) |  | 

## Methods

### NewV0041OpenapiDiagRespStatisticsRpcsByUserInner

`func NewV0041OpenapiDiagRespStatisticsRpcsByUserInner(userId int32, user string, count int32, totalTime int64, averageTime V0041OpenapiDiagRespStatisticsRpcsByMessageTypeInnerAverageTime, ) *V0041OpenapiDiagRespStatisticsRpcsByUserInner`

NewV0041OpenapiDiagRespStatisticsRpcsByUserInner instantiates a new V0041OpenapiDiagRespStatisticsRpcsByUserInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0041OpenapiDiagRespStatisticsRpcsByUserInnerWithDefaults

`func NewV0041OpenapiDiagRespStatisticsRpcsByUserInnerWithDefaults() *V0041OpenapiDiagRespStatisticsRpcsByUserInner`

NewV0041OpenapiDiagRespStatisticsRpcsByUserInnerWithDefaults instantiates a new V0041OpenapiDiagRespStatisticsRpcsByUserInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *V0041OpenapiDiagRespStatisticsRpcsByUserInner) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *V0041OpenapiDiagRespStatisticsRpcsByUserInner) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *V0041OpenapiDiagRespStatisticsRpcsByUserInner) SetUserId(v int32)`

SetUserId sets UserId field to given value.


### GetUser

`func (o *V0041OpenapiDiagRespStatisticsRpcsByUserInner) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *V0041OpenapiDiagRespStatisticsRpcsByUserInner) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *V0041OpenapiDiagRespStatisticsRpcsByUserInner) SetUser(v string)`

SetUser sets User field to given value.


### GetCount

`func (o *V0041OpenapiDiagRespStatisticsRpcsByUserInner) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *V0041OpenapiDiagRespStatisticsRpcsByUserInner) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *V0041OpenapiDiagRespStatisticsRpcsByUserInner) SetCount(v int32)`

SetCount sets Count field to given value.


### GetTotalTime

`func (o *V0041OpenapiDiagRespStatisticsRpcsByUserInner) GetTotalTime() int64`

GetTotalTime returns the TotalTime field if non-nil, zero value otherwise.

### GetTotalTimeOk

`func (o *V0041OpenapiDiagRespStatisticsRpcsByUserInner) GetTotalTimeOk() (*int64, bool)`

GetTotalTimeOk returns a tuple with the TotalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTime

`func (o *V0041OpenapiDiagRespStatisticsRpcsByUserInner) SetTotalTime(v int64)`

SetTotalTime sets TotalTime field to given value.


### GetAverageTime

`func (o *V0041OpenapiDiagRespStatisticsRpcsByUserInner) GetAverageTime() V0041OpenapiDiagRespStatisticsRpcsByMessageTypeInnerAverageTime`

GetAverageTime returns the AverageTime field if non-nil, zero value otherwise.

### GetAverageTimeOk

`func (o *V0041OpenapiDiagRespStatisticsRpcsByUserInner) GetAverageTimeOk() (*V0041OpenapiDiagRespStatisticsRpcsByMessageTypeInnerAverageTime, bool)`

GetAverageTimeOk returns a tuple with the AverageTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageTime

`func (o *V0041OpenapiDiagRespStatisticsRpcsByUserInner) SetAverageTime(v V0041OpenapiDiagRespStatisticsRpcsByMessageTypeInnerAverageTime)`

SetAverageTime sets AverageTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


