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

// checks if the V0041OpenapiSlurmdbdJobsRespJobsInnerTimeSystem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0041OpenapiSlurmdbdJobsRespJobsInnerTimeSystem{}

// V0041OpenapiSlurmdbdJobsRespJobsInnerTimeSystem struct for V0041OpenapiSlurmdbdJobsRespJobsInnerTimeSystem
type V0041OpenapiSlurmdbdJobsRespJobsInnerTimeSystem struct {
	Seconds *int64 `json:"seconds,omitempty"`
	Microseconds *int64 `json:"microseconds,omitempty"`
}

// NewV0041OpenapiSlurmdbdJobsRespJobsInnerTimeSystem instantiates a new V0041OpenapiSlurmdbdJobsRespJobsInnerTimeSystem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0041OpenapiSlurmdbdJobsRespJobsInnerTimeSystem() *V0041OpenapiSlurmdbdJobsRespJobsInnerTimeSystem {
	this := V0041OpenapiSlurmdbdJobsRespJobsInnerTimeSystem{}
	return &this
}

// NewV0041OpenapiSlurmdbdJobsRespJobsInnerTimeSystemWithDefaults instantiates a new V0041OpenapiSlurmdbdJobsRespJobsInnerTimeSystem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0041OpenapiSlurmdbdJobsRespJobsInnerTimeSystemWithDefaults() *V0041OpenapiSlurmdbdJobsRespJobsInnerTimeSystem {
	this := V0041OpenapiSlurmdbdJobsRespJobsInnerTimeSystem{}
	return &this
}

// GetSeconds returns the Seconds field value if set, zero value otherwise.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerTimeSystem) GetSeconds() int64 {
	if o == nil || IsNil(o.Seconds) {
		var ret int64
		return ret
	}
	return *o.Seconds
}

// GetSecondsOk returns a tuple with the Seconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerTimeSystem) GetSecondsOk() (*int64, bool) {
	if o == nil || IsNil(o.Seconds) {
		return nil, false
	}
	return o.Seconds, true
}

// HasSeconds returns a boolean if a field has been set.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerTimeSystem) HasSeconds() bool {
	if o != nil && !IsNil(o.Seconds) {
		return true
	}

	return false
}

// SetSeconds gets a reference to the given int64 and assigns it to the Seconds field.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerTimeSystem) SetSeconds(v int64) {
	o.Seconds = &v
}

// GetMicroseconds returns the Microseconds field value if set, zero value otherwise.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerTimeSystem) GetMicroseconds() int64 {
	if o == nil || IsNil(o.Microseconds) {
		var ret int64
		return ret
	}
	return *o.Microseconds
}

// GetMicrosecondsOk returns a tuple with the Microseconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerTimeSystem) GetMicrosecondsOk() (*int64, bool) {
	if o == nil || IsNil(o.Microseconds) {
		return nil, false
	}
	return o.Microseconds, true
}

// HasMicroseconds returns a boolean if a field has been set.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerTimeSystem) HasMicroseconds() bool {
	if o != nil && !IsNil(o.Microseconds) {
		return true
	}

	return false
}

// SetMicroseconds gets a reference to the given int64 and assigns it to the Microseconds field.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerTimeSystem) SetMicroseconds(v int64) {
	o.Microseconds = &v
}

func (o V0041OpenapiSlurmdbdJobsRespJobsInnerTimeSystem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0041OpenapiSlurmdbdJobsRespJobsInnerTimeSystem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Seconds) {
		toSerialize["seconds"] = o.Seconds
	}
	if !IsNil(o.Microseconds) {
		toSerialize["microseconds"] = o.Microseconds
	}
	return toSerialize, nil
}

type NullableV0041OpenapiSlurmdbdJobsRespJobsInnerTimeSystem struct {
	value *V0041OpenapiSlurmdbdJobsRespJobsInnerTimeSystem
	isSet bool
}

func (v NullableV0041OpenapiSlurmdbdJobsRespJobsInnerTimeSystem) Get() *V0041OpenapiSlurmdbdJobsRespJobsInnerTimeSystem {
	return v.value
}

func (v *NullableV0041OpenapiSlurmdbdJobsRespJobsInnerTimeSystem) Set(val *V0041OpenapiSlurmdbdJobsRespJobsInnerTimeSystem) {
	v.value = val
	v.isSet = true
}

func (v NullableV0041OpenapiSlurmdbdJobsRespJobsInnerTimeSystem) IsSet() bool {
	return v.isSet
}

func (v *NullableV0041OpenapiSlurmdbdJobsRespJobsInnerTimeSystem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0041OpenapiSlurmdbdJobsRespJobsInnerTimeSystem(val *V0041OpenapiSlurmdbdJobsRespJobsInnerTimeSystem) *NullableV0041OpenapiSlurmdbdJobsRespJobsInnerTimeSystem {
	return &NullableV0041OpenapiSlurmdbdJobsRespJobsInnerTimeSystem{value: val, isSet: true}
}

func (v NullableV0041OpenapiSlurmdbdJobsRespJobsInnerTimeSystem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0041OpenapiSlurmdbdJobsRespJobsInnerTimeSystem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


