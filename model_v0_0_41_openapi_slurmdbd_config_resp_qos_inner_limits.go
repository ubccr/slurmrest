/*
Slurm REST API

API to access and control Slurm

API version: Slurm-24.05.2&openapi/dbv0.0.39&openapi/v0.0.39&openapi/slurmdbd&openapi/slurmctld
Contact: sales@schedmd.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package slurmrest

import (
	"encoding/json"
)

// checks if the V0041OpenapiSlurmdbdConfigRespQosInnerLimits type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0041OpenapiSlurmdbdConfigRespQosInnerLimits{}

// V0041OpenapiSlurmdbdConfigRespQosInnerLimits struct for V0041OpenapiSlurmdbdConfigRespQosInnerLimits
type V0041OpenapiSlurmdbdConfigRespQosInnerLimits struct {
	GraceTime *int32 `json:"grace_time,omitempty"`
	Max *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMax `json:"max,omitempty"`
	Factor *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsFactor `json:"factor,omitempty"`
	Min *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMin `json:"min,omitempty"`
}

// NewV0041OpenapiSlurmdbdConfigRespQosInnerLimits instantiates a new V0041OpenapiSlurmdbdConfigRespQosInnerLimits object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0041OpenapiSlurmdbdConfigRespQosInnerLimits() *V0041OpenapiSlurmdbdConfigRespQosInnerLimits {
	this := V0041OpenapiSlurmdbdConfigRespQosInnerLimits{}
	return &this
}

// NewV0041OpenapiSlurmdbdConfigRespQosInnerLimitsWithDefaults instantiates a new V0041OpenapiSlurmdbdConfigRespQosInnerLimits object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0041OpenapiSlurmdbdConfigRespQosInnerLimitsWithDefaults() *V0041OpenapiSlurmdbdConfigRespQosInnerLimits {
	this := V0041OpenapiSlurmdbdConfigRespQosInnerLimits{}
	return &this
}

// GetGraceTime returns the GraceTime field value if set, zero value otherwise.
func (o *V0041OpenapiSlurmdbdConfigRespQosInnerLimits) GetGraceTime() int32 {
	if o == nil || IsNil(o.GraceTime) {
		var ret int32
		return ret
	}
	return *o.GraceTime
}

// GetGraceTimeOk returns a tuple with the GraceTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiSlurmdbdConfigRespQosInnerLimits) GetGraceTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.GraceTime) {
		return nil, false
	}
	return o.GraceTime, true
}

// HasGraceTime returns a boolean if a field has been set.
func (o *V0041OpenapiSlurmdbdConfigRespQosInnerLimits) HasGraceTime() bool {
	if o != nil && !IsNil(o.GraceTime) {
		return true
	}

	return false
}

// SetGraceTime gets a reference to the given int32 and assigns it to the GraceTime field.
func (o *V0041OpenapiSlurmdbdConfigRespQosInnerLimits) SetGraceTime(v int32) {
	o.GraceTime = &v
}

// GetMax returns the Max field value if set, zero value otherwise.
func (o *V0041OpenapiSlurmdbdConfigRespQosInnerLimits) GetMax() V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMax {
	if o == nil || IsNil(o.Max) {
		var ret V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMax
		return ret
	}
	return *o.Max
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiSlurmdbdConfigRespQosInnerLimits) GetMaxOk() (*V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMax, bool) {
	if o == nil || IsNil(o.Max) {
		return nil, false
	}
	return o.Max, true
}

// HasMax returns a boolean if a field has been set.
func (o *V0041OpenapiSlurmdbdConfigRespQosInnerLimits) HasMax() bool {
	if o != nil && !IsNil(o.Max) {
		return true
	}

	return false
}

// SetMax gets a reference to the given V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMax and assigns it to the Max field.
func (o *V0041OpenapiSlurmdbdConfigRespQosInnerLimits) SetMax(v V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMax) {
	o.Max = &v
}

// GetFactor returns the Factor field value if set, zero value otherwise.
func (o *V0041OpenapiSlurmdbdConfigRespQosInnerLimits) GetFactor() V0041OpenapiSlurmdbdConfigRespQosInnerLimitsFactor {
	if o == nil || IsNil(o.Factor) {
		var ret V0041OpenapiSlurmdbdConfigRespQosInnerLimitsFactor
		return ret
	}
	return *o.Factor
}

// GetFactorOk returns a tuple with the Factor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiSlurmdbdConfigRespQosInnerLimits) GetFactorOk() (*V0041OpenapiSlurmdbdConfigRespQosInnerLimitsFactor, bool) {
	if o == nil || IsNil(o.Factor) {
		return nil, false
	}
	return o.Factor, true
}

// HasFactor returns a boolean if a field has been set.
func (o *V0041OpenapiSlurmdbdConfigRespQosInnerLimits) HasFactor() bool {
	if o != nil && !IsNil(o.Factor) {
		return true
	}

	return false
}

// SetFactor gets a reference to the given V0041OpenapiSlurmdbdConfigRespQosInnerLimitsFactor and assigns it to the Factor field.
func (o *V0041OpenapiSlurmdbdConfigRespQosInnerLimits) SetFactor(v V0041OpenapiSlurmdbdConfigRespQosInnerLimitsFactor) {
	o.Factor = &v
}

// GetMin returns the Min field value if set, zero value otherwise.
func (o *V0041OpenapiSlurmdbdConfigRespQosInnerLimits) GetMin() V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMin {
	if o == nil || IsNil(o.Min) {
		var ret V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMin
		return ret
	}
	return *o.Min
}

// GetMinOk returns a tuple with the Min field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiSlurmdbdConfigRespQosInnerLimits) GetMinOk() (*V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMin, bool) {
	if o == nil || IsNil(o.Min) {
		return nil, false
	}
	return o.Min, true
}

// HasMin returns a boolean if a field has been set.
func (o *V0041OpenapiSlurmdbdConfigRespQosInnerLimits) HasMin() bool {
	if o != nil && !IsNil(o.Min) {
		return true
	}

	return false
}

// SetMin gets a reference to the given V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMin and assigns it to the Min field.
func (o *V0041OpenapiSlurmdbdConfigRespQosInnerLimits) SetMin(v V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMin) {
	o.Min = &v
}

func (o V0041OpenapiSlurmdbdConfigRespQosInnerLimits) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0041OpenapiSlurmdbdConfigRespQosInnerLimits) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GraceTime) {
		toSerialize["grace_time"] = o.GraceTime
	}
	if !IsNil(o.Max) {
		toSerialize["max"] = o.Max
	}
	if !IsNil(o.Factor) {
		toSerialize["factor"] = o.Factor
	}
	if !IsNil(o.Min) {
		toSerialize["min"] = o.Min
	}
	return toSerialize, nil
}

type NullableV0041OpenapiSlurmdbdConfigRespQosInnerLimits struct {
	value *V0041OpenapiSlurmdbdConfigRespQosInnerLimits
	isSet bool
}

func (v NullableV0041OpenapiSlurmdbdConfigRespQosInnerLimits) Get() *V0041OpenapiSlurmdbdConfigRespQosInnerLimits {
	return v.value
}

func (v *NullableV0041OpenapiSlurmdbdConfigRespQosInnerLimits) Set(val *V0041OpenapiSlurmdbdConfigRespQosInnerLimits) {
	v.value = val
	v.isSet = true
}

func (v NullableV0041OpenapiSlurmdbdConfigRespQosInnerLimits) IsSet() bool {
	return v.isSet
}

func (v *NullableV0041OpenapiSlurmdbdConfigRespQosInnerLimits) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0041OpenapiSlurmdbdConfigRespQosInnerLimits(val *V0041OpenapiSlurmdbdConfigRespQosInnerLimits) *NullableV0041OpenapiSlurmdbdConfigRespQosInnerLimits {
	return &NullableV0041OpenapiSlurmdbdConfigRespQosInnerLimits{value: val, isSet: true}
}

func (v NullableV0041OpenapiSlurmdbdConfigRespQosInnerLimits) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0041OpenapiSlurmdbdConfigRespQosInnerLimits) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

