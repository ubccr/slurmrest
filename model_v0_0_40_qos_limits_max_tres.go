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

// checks if the V0040QosLimitsMaxTres type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0040QosLimitsMaxTres{}

// V0040QosLimitsMaxTres struct for V0040QosLimitsMaxTres
type V0040QosLimitsMaxTres struct {
	Total []V0040Tres `json:"total,omitempty"`
	Minutes *V0040QosLimitsMaxTresMinutes `json:"minutes,omitempty"`
	Per *V0040QosLimitsMaxTresPer `json:"per,omitempty"`
}

// NewV0040QosLimitsMaxTres instantiates a new V0040QosLimitsMaxTres object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0040QosLimitsMaxTres() *V0040QosLimitsMaxTres {
	this := V0040QosLimitsMaxTres{}
	return &this
}

// NewV0040QosLimitsMaxTresWithDefaults instantiates a new V0040QosLimitsMaxTres object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0040QosLimitsMaxTresWithDefaults() *V0040QosLimitsMaxTres {
	this := V0040QosLimitsMaxTres{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *V0040QosLimitsMaxTres) GetTotal() []V0040Tres {
	if o == nil || IsNil(o.Total) {
		var ret []V0040Tres
		return ret
	}
	return o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040QosLimitsMaxTres) GetTotalOk() ([]V0040Tres, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *V0040QosLimitsMaxTres) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given []V0040Tres and assigns it to the Total field.
func (o *V0040QosLimitsMaxTres) SetTotal(v []V0040Tres) {
	o.Total = v
}

// GetMinutes returns the Minutes field value if set, zero value otherwise.
func (o *V0040QosLimitsMaxTres) GetMinutes() V0040QosLimitsMaxTresMinutes {
	if o == nil || IsNil(o.Minutes) {
		var ret V0040QosLimitsMaxTresMinutes
		return ret
	}
	return *o.Minutes
}

// GetMinutesOk returns a tuple with the Minutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040QosLimitsMaxTres) GetMinutesOk() (*V0040QosLimitsMaxTresMinutes, bool) {
	if o == nil || IsNil(o.Minutes) {
		return nil, false
	}
	return o.Minutes, true
}

// HasMinutes returns a boolean if a field has been set.
func (o *V0040QosLimitsMaxTres) HasMinutes() bool {
	if o != nil && !IsNil(o.Minutes) {
		return true
	}

	return false
}

// SetMinutes gets a reference to the given V0040QosLimitsMaxTresMinutes and assigns it to the Minutes field.
func (o *V0040QosLimitsMaxTres) SetMinutes(v V0040QosLimitsMaxTresMinutes) {
	o.Minutes = &v
}

// GetPer returns the Per field value if set, zero value otherwise.
func (o *V0040QosLimitsMaxTres) GetPer() V0040QosLimitsMaxTresPer {
	if o == nil || IsNil(o.Per) {
		var ret V0040QosLimitsMaxTresPer
		return ret
	}
	return *o.Per
}

// GetPerOk returns a tuple with the Per field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040QosLimitsMaxTres) GetPerOk() (*V0040QosLimitsMaxTresPer, bool) {
	if o == nil || IsNil(o.Per) {
		return nil, false
	}
	return o.Per, true
}

// HasPer returns a boolean if a field has been set.
func (o *V0040QosLimitsMaxTres) HasPer() bool {
	if o != nil && !IsNil(o.Per) {
		return true
	}

	return false
}

// SetPer gets a reference to the given V0040QosLimitsMaxTresPer and assigns it to the Per field.
func (o *V0040QosLimitsMaxTres) SetPer(v V0040QosLimitsMaxTresPer) {
	o.Per = &v
}

func (o V0040QosLimitsMaxTres) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0040QosLimitsMaxTres) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !IsNil(o.Minutes) {
		toSerialize["minutes"] = o.Minutes
	}
	if !IsNil(o.Per) {
		toSerialize["per"] = o.Per
	}
	return toSerialize, nil
}

type NullableV0040QosLimitsMaxTres struct {
	value *V0040QosLimitsMaxTres
	isSet bool
}

func (v NullableV0040QosLimitsMaxTres) Get() *V0040QosLimitsMaxTres {
	return v.value
}

func (v *NullableV0040QosLimitsMaxTres) Set(val *V0040QosLimitsMaxTres) {
	v.value = val
	v.isSet = true
}

func (v NullableV0040QosLimitsMaxTres) IsSet() bool {
	return v.isSet
}

func (v *NullableV0040QosLimitsMaxTres) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0040QosLimitsMaxTres(val *V0040QosLimitsMaxTres) *NullableV0040QosLimitsMaxTres {
	return &NullableV0040QosLimitsMaxTres{value: val, isSet: true}
}

func (v NullableV0040QosLimitsMaxTres) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0040QosLimitsMaxTres) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


