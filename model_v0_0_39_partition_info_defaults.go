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

// checks if the V0039PartitionInfoDefaults type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0039PartitionInfoDefaults{}

// V0039PartitionInfoDefaults struct for V0039PartitionInfoDefaults
type V0039PartitionInfoDefaults struct {
	MemoryPerCpu *int64 `json:"memory_per_cpu,omitempty"`
	Time *V0039Uint32NoVal `json:"time,omitempty"`
	Job *string `json:"job,omitempty"`
}

// NewV0039PartitionInfoDefaults instantiates a new V0039PartitionInfoDefaults object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0039PartitionInfoDefaults() *V0039PartitionInfoDefaults {
	this := V0039PartitionInfoDefaults{}
	return &this
}

// NewV0039PartitionInfoDefaultsWithDefaults instantiates a new V0039PartitionInfoDefaults object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0039PartitionInfoDefaultsWithDefaults() *V0039PartitionInfoDefaults {
	this := V0039PartitionInfoDefaults{}
	return &this
}

// GetMemoryPerCpu returns the MemoryPerCpu field value if set, zero value otherwise.
func (o *V0039PartitionInfoDefaults) GetMemoryPerCpu() int64 {
	if o == nil || IsNil(o.MemoryPerCpu) {
		var ret int64
		return ret
	}
	return *o.MemoryPerCpu
}

// GetMemoryPerCpuOk returns a tuple with the MemoryPerCpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039PartitionInfoDefaults) GetMemoryPerCpuOk() (*int64, bool) {
	if o == nil || IsNil(o.MemoryPerCpu) {
		return nil, false
	}
	return o.MemoryPerCpu, true
}

// HasMemoryPerCpu returns a boolean if a field has been set.
func (o *V0039PartitionInfoDefaults) HasMemoryPerCpu() bool {
	if o != nil && !IsNil(o.MemoryPerCpu) {
		return true
	}

	return false
}

// SetMemoryPerCpu gets a reference to the given int64 and assigns it to the MemoryPerCpu field.
func (o *V0039PartitionInfoDefaults) SetMemoryPerCpu(v int64) {
	o.MemoryPerCpu = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *V0039PartitionInfoDefaults) GetTime() V0039Uint32NoVal {
	if o == nil || IsNil(o.Time) {
		var ret V0039Uint32NoVal
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039PartitionInfoDefaults) GetTimeOk() (*V0039Uint32NoVal, bool) {
	if o == nil || IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *V0039PartitionInfoDefaults) HasTime() bool {
	if o != nil && !IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given V0039Uint32NoVal and assigns it to the Time field.
func (o *V0039PartitionInfoDefaults) SetTime(v V0039Uint32NoVal) {
	o.Time = &v
}

// GetJob returns the Job field value if set, zero value otherwise.
func (o *V0039PartitionInfoDefaults) GetJob() string {
	if o == nil || IsNil(o.Job) {
		var ret string
		return ret
	}
	return *o.Job
}

// GetJobOk returns a tuple with the Job field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039PartitionInfoDefaults) GetJobOk() (*string, bool) {
	if o == nil || IsNil(o.Job) {
		return nil, false
	}
	return o.Job, true
}

// HasJob returns a boolean if a field has been set.
func (o *V0039PartitionInfoDefaults) HasJob() bool {
	if o != nil && !IsNil(o.Job) {
		return true
	}

	return false
}

// SetJob gets a reference to the given string and assigns it to the Job field.
func (o *V0039PartitionInfoDefaults) SetJob(v string) {
	o.Job = &v
}

func (o V0039PartitionInfoDefaults) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0039PartitionInfoDefaults) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MemoryPerCpu) {
		toSerialize["memory_per_cpu"] = o.MemoryPerCpu
	}
	if !IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	if !IsNil(o.Job) {
		toSerialize["job"] = o.Job
	}
	return toSerialize, nil
}

type NullableV0039PartitionInfoDefaults struct {
	value *V0039PartitionInfoDefaults
	isSet bool
}

func (v NullableV0039PartitionInfoDefaults) Get() *V0039PartitionInfoDefaults {
	return v.value
}

func (v *NullableV0039PartitionInfoDefaults) Set(val *V0039PartitionInfoDefaults) {
	v.value = val
	v.isSet = true
}

func (v NullableV0039PartitionInfoDefaults) IsSet() bool {
	return v.isSet
}

func (v *NullableV0039PartitionInfoDefaults) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0039PartitionInfoDefaults(val *V0039PartitionInfoDefaults) *NullableV0039PartitionInfoDefaults {
	return &NullableV0039PartitionInfoDefaults{value: val, isSet: true}
}

func (v NullableV0039PartitionInfoDefaults) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0039PartitionInfoDefaults) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


