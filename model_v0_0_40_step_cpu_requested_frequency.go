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

// checks if the V0040StepCPURequestedFrequency type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0040StepCPURequestedFrequency{}

// V0040StepCPURequestedFrequency struct for V0040StepCPURequestedFrequency
type V0040StepCPURequestedFrequency struct {
	Min *V0040Uint32NoVal `json:"min,omitempty"`
	Max *V0040Uint32NoVal `json:"max,omitempty"`
}

// NewV0040StepCPURequestedFrequency instantiates a new V0040StepCPURequestedFrequency object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0040StepCPURequestedFrequency() *V0040StepCPURequestedFrequency {
	this := V0040StepCPURequestedFrequency{}
	return &this
}

// NewV0040StepCPURequestedFrequencyWithDefaults instantiates a new V0040StepCPURequestedFrequency object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0040StepCPURequestedFrequencyWithDefaults() *V0040StepCPURequestedFrequency {
	this := V0040StepCPURequestedFrequency{}
	return &this
}

// GetMin returns the Min field value if set, zero value otherwise.
func (o *V0040StepCPURequestedFrequency) GetMin() V0040Uint32NoVal {
	if o == nil || IsNil(o.Min) {
		var ret V0040Uint32NoVal
		return ret
	}
	return *o.Min
}

// GetMinOk returns a tuple with the Min field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040StepCPURequestedFrequency) GetMinOk() (*V0040Uint32NoVal, bool) {
	if o == nil || IsNil(o.Min) {
		return nil, false
	}
	return o.Min, true
}

// HasMin returns a boolean if a field has been set.
func (o *V0040StepCPURequestedFrequency) HasMin() bool {
	if o != nil && !IsNil(o.Min) {
		return true
	}

	return false
}

// SetMin gets a reference to the given V0040Uint32NoVal and assigns it to the Min field.
func (o *V0040StepCPURequestedFrequency) SetMin(v V0040Uint32NoVal) {
	o.Min = &v
}

// GetMax returns the Max field value if set, zero value otherwise.
func (o *V0040StepCPURequestedFrequency) GetMax() V0040Uint32NoVal {
	if o == nil || IsNil(o.Max) {
		var ret V0040Uint32NoVal
		return ret
	}
	return *o.Max
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040StepCPURequestedFrequency) GetMaxOk() (*V0040Uint32NoVal, bool) {
	if o == nil || IsNil(o.Max) {
		return nil, false
	}
	return o.Max, true
}

// HasMax returns a boolean if a field has been set.
func (o *V0040StepCPURequestedFrequency) HasMax() bool {
	if o != nil && !IsNil(o.Max) {
		return true
	}

	return false
}

// SetMax gets a reference to the given V0040Uint32NoVal and assigns it to the Max field.
func (o *V0040StepCPURequestedFrequency) SetMax(v V0040Uint32NoVal) {
	o.Max = &v
}

func (o V0040StepCPURequestedFrequency) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0040StepCPURequestedFrequency) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Min) {
		toSerialize["min"] = o.Min
	}
	if !IsNil(o.Max) {
		toSerialize["max"] = o.Max
	}
	return toSerialize, nil
}

type NullableV0040StepCPURequestedFrequency struct {
	value *V0040StepCPURequestedFrequency
	isSet bool
}

func (v NullableV0040StepCPURequestedFrequency) Get() *V0040StepCPURequestedFrequency {
	return v.value
}

func (v *NullableV0040StepCPURequestedFrequency) Set(val *V0040StepCPURequestedFrequency) {
	v.value = val
	v.isSet = true
}

func (v NullableV0040StepCPURequestedFrequency) IsSet() bool {
	return v.isSet
}

func (v *NullableV0040StepCPURequestedFrequency) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0040StepCPURequestedFrequency(val *V0040StepCPURequestedFrequency) *NullableV0040StepCPURequestedFrequency {
	return &NullableV0040StepCPURequestedFrequency{value: val, isSet: true}
}

func (v NullableV0040StepCPURequestedFrequency) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0040StepCPURequestedFrequency) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

