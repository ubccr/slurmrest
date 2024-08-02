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

// checks if the V0041OpenapiPartitionRespPartitionsInnerCpus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0041OpenapiPartitionRespPartitionsInnerCpus{}

// V0041OpenapiPartitionRespPartitionsInnerCpus struct for V0041OpenapiPartitionRespPartitionsInnerCpus
type V0041OpenapiPartitionRespPartitionsInnerCpus struct {
	TaskBinding *int32 `json:"task_binding,omitempty"`
	Total *int32 `json:"total,omitempty"`
}

// NewV0041OpenapiPartitionRespPartitionsInnerCpus instantiates a new V0041OpenapiPartitionRespPartitionsInnerCpus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0041OpenapiPartitionRespPartitionsInnerCpus() *V0041OpenapiPartitionRespPartitionsInnerCpus {
	this := V0041OpenapiPartitionRespPartitionsInnerCpus{}
	return &this
}

// NewV0041OpenapiPartitionRespPartitionsInnerCpusWithDefaults instantiates a new V0041OpenapiPartitionRespPartitionsInnerCpus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0041OpenapiPartitionRespPartitionsInnerCpusWithDefaults() *V0041OpenapiPartitionRespPartitionsInnerCpus {
	this := V0041OpenapiPartitionRespPartitionsInnerCpus{}
	return &this
}

// GetTaskBinding returns the TaskBinding field value if set, zero value otherwise.
func (o *V0041OpenapiPartitionRespPartitionsInnerCpus) GetTaskBinding() int32 {
	if o == nil || IsNil(o.TaskBinding) {
		var ret int32
		return ret
	}
	return *o.TaskBinding
}

// GetTaskBindingOk returns a tuple with the TaskBinding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiPartitionRespPartitionsInnerCpus) GetTaskBindingOk() (*int32, bool) {
	if o == nil || IsNil(o.TaskBinding) {
		return nil, false
	}
	return o.TaskBinding, true
}

// HasTaskBinding returns a boolean if a field has been set.
func (o *V0041OpenapiPartitionRespPartitionsInnerCpus) HasTaskBinding() bool {
	if o != nil && !IsNil(o.TaskBinding) {
		return true
	}

	return false
}

// SetTaskBinding gets a reference to the given int32 and assigns it to the TaskBinding field.
func (o *V0041OpenapiPartitionRespPartitionsInnerCpus) SetTaskBinding(v int32) {
	o.TaskBinding = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *V0041OpenapiPartitionRespPartitionsInnerCpus) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiPartitionRespPartitionsInnerCpus) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *V0041OpenapiPartitionRespPartitionsInnerCpus) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *V0041OpenapiPartitionRespPartitionsInnerCpus) SetTotal(v int32) {
	o.Total = &v
}

func (o V0041OpenapiPartitionRespPartitionsInnerCpus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0041OpenapiPartitionRespPartitionsInnerCpus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TaskBinding) {
		toSerialize["task_binding"] = o.TaskBinding
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	return toSerialize, nil
}

type NullableV0041OpenapiPartitionRespPartitionsInnerCpus struct {
	value *V0041OpenapiPartitionRespPartitionsInnerCpus
	isSet bool
}

func (v NullableV0041OpenapiPartitionRespPartitionsInnerCpus) Get() *V0041OpenapiPartitionRespPartitionsInnerCpus {
	return v.value
}

func (v *NullableV0041OpenapiPartitionRespPartitionsInnerCpus) Set(val *V0041OpenapiPartitionRespPartitionsInnerCpus) {
	v.value = val
	v.isSet = true
}

func (v NullableV0041OpenapiPartitionRespPartitionsInnerCpus) IsSet() bool {
	return v.isSet
}

func (v *NullableV0041OpenapiPartitionRespPartitionsInnerCpus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0041OpenapiPartitionRespPartitionsInnerCpus(val *V0041OpenapiPartitionRespPartitionsInnerCpus) *NullableV0041OpenapiPartitionRespPartitionsInnerCpus {
	return &NullableV0041OpenapiPartitionRespPartitionsInnerCpus{value: val, isSet: true}
}

func (v NullableV0041OpenapiPartitionRespPartitionsInnerCpus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0041OpenapiPartitionRespPartitionsInnerCpus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


