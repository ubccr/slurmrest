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

// checks if the V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMin type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMin{}

// V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMin struct for V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMin
type V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMin struct {
	PriorityThreshold *V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId `json:"priority_threshold,omitempty"`
	Tres *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMinTres `json:"tres,omitempty"`
}

// NewV0041OpenapiSlurmdbdConfigRespQosInnerLimitsMin instantiates a new V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMin object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0041OpenapiSlurmdbdConfigRespQosInnerLimitsMin() *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMin {
	this := V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMin{}
	return &this
}

// NewV0041OpenapiSlurmdbdConfigRespQosInnerLimitsMinWithDefaults instantiates a new V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMin object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0041OpenapiSlurmdbdConfigRespQosInnerLimitsMinWithDefaults() *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMin {
	this := V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMin{}
	return &this
}

// GetPriorityThreshold returns the PriorityThreshold field value if set, zero value otherwise.
func (o *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMin) GetPriorityThreshold() V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId {
	if o == nil || IsNil(o.PriorityThreshold) {
		var ret V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId
		return ret
	}
	return *o.PriorityThreshold
}

// GetPriorityThresholdOk returns a tuple with the PriorityThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMin) GetPriorityThresholdOk() (*V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId, bool) {
	if o == nil || IsNil(o.PriorityThreshold) {
		return nil, false
	}
	return o.PriorityThreshold, true
}

// HasPriorityThreshold returns a boolean if a field has been set.
func (o *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMin) HasPriorityThreshold() bool {
	if o != nil && !IsNil(o.PriorityThreshold) {
		return true
	}

	return false
}

// SetPriorityThreshold gets a reference to the given V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId and assigns it to the PriorityThreshold field.
func (o *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMin) SetPriorityThreshold(v V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId) {
	o.PriorityThreshold = &v
}

// GetTres returns the Tres field value if set, zero value otherwise.
func (o *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMin) GetTres() V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMinTres {
	if o == nil || IsNil(o.Tres) {
		var ret V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMinTres
		return ret
	}
	return *o.Tres
}

// GetTresOk returns a tuple with the Tres field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMin) GetTresOk() (*V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMinTres, bool) {
	if o == nil || IsNil(o.Tres) {
		return nil, false
	}
	return o.Tres, true
}

// HasTres returns a boolean if a field has been set.
func (o *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMin) HasTres() bool {
	if o != nil && !IsNil(o.Tres) {
		return true
	}

	return false
}

// SetTres gets a reference to the given V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMinTres and assigns it to the Tres field.
func (o *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMin) SetTres(v V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMinTres) {
	o.Tres = &v
}

func (o V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMin) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMin) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PriorityThreshold) {
		toSerialize["priority_threshold"] = o.PriorityThreshold
	}
	if !IsNil(o.Tres) {
		toSerialize["tres"] = o.Tres
	}
	return toSerialize, nil
}

type NullableV0041OpenapiSlurmdbdConfigRespQosInnerLimitsMin struct {
	value *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMin
	isSet bool
}

func (v NullableV0041OpenapiSlurmdbdConfigRespQosInnerLimitsMin) Get() *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMin {
	return v.value
}

func (v *NullableV0041OpenapiSlurmdbdConfigRespQosInnerLimitsMin) Set(val *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMin) {
	v.value = val
	v.isSet = true
}

func (v NullableV0041OpenapiSlurmdbdConfigRespQosInnerLimitsMin) IsSet() bool {
	return v.isSet
}

func (v *NullableV0041OpenapiSlurmdbdConfigRespQosInnerLimitsMin) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0041OpenapiSlurmdbdConfigRespQosInnerLimitsMin(val *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMin) *NullableV0041OpenapiSlurmdbdConfigRespQosInnerLimitsMin {
	return &NullableV0041OpenapiSlurmdbdConfigRespQosInnerLimitsMin{value: val, isSet: true}
}

func (v NullableV0041OpenapiSlurmdbdConfigRespQosInnerLimitsMin) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0041OpenapiSlurmdbdConfigRespQosInnerLimitsMin) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

