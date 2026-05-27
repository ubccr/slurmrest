# V0044QosLimits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GraceTime** | Pointer to **int32** | GraceTime - Preemption grace time in seconds to be extended to a job which has been selected for preemption | [optional] 
**Max** | Pointer to [**V0044QosLimitsMax**](V0044QosLimitsMax.md) |  | [optional] 
**Factor** | Pointer to [**V0044Float64NoValStruct**](V0044Float64NoValStruct.md) |  | [optional] 
**Min** | Pointer to [**V0044QosLimitsMin**](V0044QosLimitsMin.md) |  | [optional] 

## Methods

### NewV0044QosLimits

`func NewV0044QosLimits() *V0044QosLimits`

NewV0044QosLimits instantiates a new V0044QosLimits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0044QosLimitsWithDefaults

`func NewV0044QosLimitsWithDefaults() *V0044QosLimits`

NewV0044QosLimitsWithDefaults instantiates a new V0044QosLimits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGraceTime

`func (o *V0044QosLimits) GetGraceTime() int32`

GetGraceTime returns the GraceTime field if non-nil, zero value otherwise.

### GetGraceTimeOk

`func (o *V0044QosLimits) GetGraceTimeOk() (*int32, bool)`

GetGraceTimeOk returns a tuple with the GraceTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraceTime

`func (o *V0044QosLimits) SetGraceTime(v int32)`

SetGraceTime sets GraceTime field to given value.

### HasGraceTime

`func (o *V0044QosLimits) HasGraceTime() bool`

HasGraceTime returns a boolean if a field has been set.

### GetMax

`func (o *V0044QosLimits) GetMax() V0044QosLimitsMax`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *V0044QosLimits) GetMaxOk() (*V0044QosLimitsMax, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *V0044QosLimits) SetMax(v V0044QosLimitsMax)`

SetMax sets Max field to given value.

### HasMax

`func (o *V0044QosLimits) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetFactor

`func (o *V0044QosLimits) GetFactor() V0044Float64NoValStruct`

GetFactor returns the Factor field if non-nil, zero value otherwise.

### GetFactorOk

`func (o *V0044QosLimits) GetFactorOk() (*V0044Float64NoValStruct, bool)`

GetFactorOk returns a tuple with the Factor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactor

`func (o *V0044QosLimits) SetFactor(v V0044Float64NoValStruct)`

SetFactor sets Factor field to given value.

### HasFactor

`func (o *V0044QosLimits) HasFactor() bool`

HasFactor returns a boolean if a field has been set.

### GetMin

`func (o *V0044QosLimits) GetMin() V0044QosLimitsMin`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *V0044QosLimits) GetMinOk() (*V0044QosLimitsMin, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *V0044QosLimits) SetMin(v V0044QosLimitsMin)`

SetMin sets Min field to given value.

### HasMin

`func (o *V0044QosLimits) HasMin() bool`

HasMin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


