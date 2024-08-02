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

// checks if the V0040QosLimitsMin type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0040QosLimitsMin{}

// V0040QosLimitsMin struct for V0040QosLimitsMin
type V0040QosLimitsMin struct {
	PriorityThreshold *V0040Uint32NoVal `json:"priority_threshold,omitempty"`
	Tres *V0040QosLimitsMinTres `json:"tres,omitempty"`
}

// NewV0040QosLimitsMin instantiates a new V0040QosLimitsMin object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0040QosLimitsMin() *V0040QosLimitsMin {
	this := V0040QosLimitsMin{}
	return &this
}

// NewV0040QosLimitsMinWithDefaults instantiates a new V0040QosLimitsMin object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0040QosLimitsMinWithDefaults() *V0040QosLimitsMin {
	this := V0040QosLimitsMin{}
	return &this
}

// GetPriorityThreshold returns the PriorityThreshold field value if set, zero value otherwise.
func (o *V0040QosLimitsMin) GetPriorityThreshold() V0040Uint32NoVal {
	if o == nil || IsNil(o.PriorityThreshold) {
		var ret V0040Uint32NoVal
		return ret
	}
	return *o.PriorityThreshold
}

// GetPriorityThresholdOk returns a tuple with the PriorityThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040QosLimitsMin) GetPriorityThresholdOk() (*V0040Uint32NoVal, bool) {
	if o == nil || IsNil(o.PriorityThreshold) {
		return nil, false
	}
	return o.PriorityThreshold, true
}

// HasPriorityThreshold returns a boolean if a field has been set.
func (o *V0040QosLimitsMin) HasPriorityThreshold() bool {
	if o != nil && !IsNil(o.PriorityThreshold) {
		return true
	}

	return false
}

// SetPriorityThreshold gets a reference to the given V0040Uint32NoVal and assigns it to the PriorityThreshold field.
func (o *V0040QosLimitsMin) SetPriorityThreshold(v V0040Uint32NoVal) {
	o.PriorityThreshold = &v
}

// GetTres returns the Tres field value if set, zero value otherwise.
func (o *V0040QosLimitsMin) GetTres() V0040QosLimitsMinTres {
	if o == nil || IsNil(o.Tres) {
		var ret V0040QosLimitsMinTres
		return ret
	}
	return *o.Tres
}

// GetTresOk returns a tuple with the Tres field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040QosLimitsMin) GetTresOk() (*V0040QosLimitsMinTres, bool) {
	if o == nil || IsNil(o.Tres) {
		return nil, false
	}
	return o.Tres, true
}

// HasTres returns a boolean if a field has been set.
func (o *V0040QosLimitsMin) HasTres() bool {
	if o != nil && !IsNil(o.Tres) {
		return true
	}

	return false
}

// SetTres gets a reference to the given V0040QosLimitsMinTres and assigns it to the Tres field.
func (o *V0040QosLimitsMin) SetTres(v V0040QosLimitsMinTres) {
	o.Tres = &v
}

func (o V0040QosLimitsMin) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0040QosLimitsMin) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PriorityThreshold) {
		toSerialize["priority_threshold"] = o.PriorityThreshold
	}
	if !IsNil(o.Tres) {
		toSerialize["tres"] = o.Tres
	}
	return toSerialize, nil
}

type NullableV0040QosLimitsMin struct {
	value *V0040QosLimitsMin
	isSet bool
}

func (v NullableV0040QosLimitsMin) Get() *V0040QosLimitsMin {
	return v.value
}

func (v *NullableV0040QosLimitsMin) Set(val *V0040QosLimitsMin) {
	v.value = val
	v.isSet = true
}

func (v NullableV0040QosLimitsMin) IsSet() bool {
	return v.isSet
}

func (v *NullableV0040QosLimitsMin) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0040QosLimitsMin(val *V0040QosLimitsMin) *NullableV0040QosLimitsMin {
	return &NullableV0040QosLimitsMin{value: val, isSet: true}
}

func (v NullableV0040QosLimitsMin) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0040QosLimitsMin) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

