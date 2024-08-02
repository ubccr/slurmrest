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

// checks if the V0039QosLimitsMaxJobs type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0039QosLimitsMaxJobs{}

// V0039QosLimitsMaxJobs struct for V0039QosLimitsMaxJobs
type V0039QosLimitsMaxJobs struct {
	ActiveJobs *V0039QosLimitsMaxJobsActiveJobs `json:"active_jobs,omitempty"`
	Per *V0039QosLimitsMaxJobsActiveJobsPer `json:"per,omitempty"`
}

// NewV0039QosLimitsMaxJobs instantiates a new V0039QosLimitsMaxJobs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0039QosLimitsMaxJobs() *V0039QosLimitsMaxJobs {
	this := V0039QosLimitsMaxJobs{}
	return &this
}

// NewV0039QosLimitsMaxJobsWithDefaults instantiates a new V0039QosLimitsMaxJobs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0039QosLimitsMaxJobsWithDefaults() *V0039QosLimitsMaxJobs {
	this := V0039QosLimitsMaxJobs{}
	return &this
}

// GetActiveJobs returns the ActiveJobs field value if set, zero value otherwise.
func (o *V0039QosLimitsMaxJobs) GetActiveJobs() V0039QosLimitsMaxJobsActiveJobs {
	if o == nil || IsNil(o.ActiveJobs) {
		var ret V0039QosLimitsMaxJobsActiveJobs
		return ret
	}
	return *o.ActiveJobs
}

// GetActiveJobsOk returns a tuple with the ActiveJobs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039QosLimitsMaxJobs) GetActiveJobsOk() (*V0039QosLimitsMaxJobsActiveJobs, bool) {
	if o == nil || IsNil(o.ActiveJobs) {
		return nil, false
	}
	return o.ActiveJobs, true
}

// HasActiveJobs returns a boolean if a field has been set.
func (o *V0039QosLimitsMaxJobs) HasActiveJobs() bool {
	if o != nil && !IsNil(o.ActiveJobs) {
		return true
	}

	return false
}

// SetActiveJobs gets a reference to the given V0039QosLimitsMaxJobsActiveJobs and assigns it to the ActiveJobs field.
func (o *V0039QosLimitsMaxJobs) SetActiveJobs(v V0039QosLimitsMaxJobsActiveJobs) {
	o.ActiveJobs = &v
}

// GetPer returns the Per field value if set, zero value otherwise.
func (o *V0039QosLimitsMaxJobs) GetPer() V0039QosLimitsMaxJobsActiveJobsPer {
	if o == nil || IsNil(o.Per) {
		var ret V0039QosLimitsMaxJobsActiveJobsPer
		return ret
	}
	return *o.Per
}

// GetPerOk returns a tuple with the Per field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039QosLimitsMaxJobs) GetPerOk() (*V0039QosLimitsMaxJobsActiveJobsPer, bool) {
	if o == nil || IsNil(o.Per) {
		return nil, false
	}
	return o.Per, true
}

// HasPer returns a boolean if a field has been set.
func (o *V0039QosLimitsMaxJobs) HasPer() bool {
	if o != nil && !IsNil(o.Per) {
		return true
	}

	return false
}

// SetPer gets a reference to the given V0039QosLimitsMaxJobsActiveJobsPer and assigns it to the Per field.
func (o *V0039QosLimitsMaxJobs) SetPer(v V0039QosLimitsMaxJobsActiveJobsPer) {
	o.Per = &v
}

func (o V0039QosLimitsMaxJobs) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0039QosLimitsMaxJobs) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ActiveJobs) {
		toSerialize["active_jobs"] = o.ActiveJobs
	}
	if !IsNil(o.Per) {
		toSerialize["per"] = o.Per
	}
	return toSerialize, nil
}

type NullableV0039QosLimitsMaxJobs struct {
	value *V0039QosLimitsMaxJobs
	isSet bool
}

func (v NullableV0039QosLimitsMaxJobs) Get() *V0039QosLimitsMaxJobs {
	return v.value
}

func (v *NullableV0039QosLimitsMaxJobs) Set(val *V0039QosLimitsMaxJobs) {
	v.value = val
	v.isSet = true
}

func (v NullableV0039QosLimitsMaxJobs) IsSet() bool {
	return v.isSet
}

func (v *NullableV0039QosLimitsMaxJobs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0039QosLimitsMaxJobs(val *V0039QosLimitsMaxJobs) *NullableV0039QosLimitsMaxJobs {
	return &NullableV0039QosLimitsMaxJobs{value: val, isSet: true}
}

func (v NullableV0039QosLimitsMaxJobs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0039QosLimitsMaxJobs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

