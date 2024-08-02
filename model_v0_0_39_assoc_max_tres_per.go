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

// checks if the V0039AssocMaxTresPer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0039AssocMaxTresPer{}

// V0039AssocMaxTresPer struct for V0039AssocMaxTresPer
type V0039AssocMaxTresPer struct {
	Job []V0039Tres `json:"job,omitempty"`
	Node []V0039Tres `json:"node,omitempty"`
}

// NewV0039AssocMaxTresPer instantiates a new V0039AssocMaxTresPer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0039AssocMaxTresPer() *V0039AssocMaxTresPer {
	this := V0039AssocMaxTresPer{}
	return &this
}

// NewV0039AssocMaxTresPerWithDefaults instantiates a new V0039AssocMaxTresPer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0039AssocMaxTresPerWithDefaults() *V0039AssocMaxTresPer {
	this := V0039AssocMaxTresPer{}
	return &this
}

// GetJob returns the Job field value if set, zero value otherwise.
func (o *V0039AssocMaxTresPer) GetJob() []V0039Tres {
	if o == nil || IsNil(o.Job) {
		var ret []V0039Tres
		return ret
	}
	return o.Job
}

// GetJobOk returns a tuple with the Job field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039AssocMaxTresPer) GetJobOk() ([]V0039Tres, bool) {
	if o == nil || IsNil(o.Job) {
		return nil, false
	}
	return o.Job, true
}

// HasJob returns a boolean if a field has been set.
func (o *V0039AssocMaxTresPer) HasJob() bool {
	if o != nil && !IsNil(o.Job) {
		return true
	}

	return false
}

// SetJob gets a reference to the given []V0039Tres and assigns it to the Job field.
func (o *V0039AssocMaxTresPer) SetJob(v []V0039Tres) {
	o.Job = v
}

// GetNode returns the Node field value if set, zero value otherwise.
func (o *V0039AssocMaxTresPer) GetNode() []V0039Tres {
	if o == nil || IsNil(o.Node) {
		var ret []V0039Tres
		return ret
	}
	return o.Node
}

// GetNodeOk returns a tuple with the Node field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039AssocMaxTresPer) GetNodeOk() ([]V0039Tres, bool) {
	if o == nil || IsNil(o.Node) {
		return nil, false
	}
	return o.Node, true
}

// HasNode returns a boolean if a field has been set.
func (o *V0039AssocMaxTresPer) HasNode() bool {
	if o != nil && !IsNil(o.Node) {
		return true
	}

	return false
}

// SetNode gets a reference to the given []V0039Tres and assigns it to the Node field.
func (o *V0039AssocMaxTresPer) SetNode(v []V0039Tres) {
	o.Node = v
}

func (o V0039AssocMaxTresPer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0039AssocMaxTresPer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Job) {
		toSerialize["job"] = o.Job
	}
	if !IsNil(o.Node) {
		toSerialize["node"] = o.Node
	}
	return toSerialize, nil
}

type NullableV0039AssocMaxTresPer struct {
	value *V0039AssocMaxTresPer
	isSet bool
}

func (v NullableV0039AssocMaxTresPer) Get() *V0039AssocMaxTresPer {
	return v.value
}

func (v *NullableV0039AssocMaxTresPer) Set(val *V0039AssocMaxTresPer) {
	v.value = val
	v.isSet = true
}

func (v NullableV0039AssocMaxTresPer) IsSet() bool {
	return v.isSet
}

func (v *NullableV0039AssocMaxTresPer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0039AssocMaxTresPer(val *V0039AssocMaxTresPer) *NullableV0039AssocMaxTresPer {
	return &NullableV0039AssocMaxTresPer{value: val, isSet: true}
}

func (v NullableV0039AssocMaxTresPer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0039AssocMaxTresPer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


