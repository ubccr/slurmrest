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

// checks if the V0041JobDescMsgCrontabLine type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0041JobDescMsgCrontabLine{}

// V0041JobDescMsgCrontabLine struct for V0041JobDescMsgCrontabLine
type V0041JobDescMsgCrontabLine struct {
	Start *int32 `json:"start,omitempty"`
	End *int32 `json:"end,omitempty"`
}

// NewV0041JobDescMsgCrontabLine instantiates a new V0041JobDescMsgCrontabLine object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0041JobDescMsgCrontabLine() *V0041JobDescMsgCrontabLine {
	this := V0041JobDescMsgCrontabLine{}
	return &this
}

// NewV0041JobDescMsgCrontabLineWithDefaults instantiates a new V0041JobDescMsgCrontabLine object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0041JobDescMsgCrontabLineWithDefaults() *V0041JobDescMsgCrontabLine {
	this := V0041JobDescMsgCrontabLine{}
	return &this
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *V0041JobDescMsgCrontabLine) GetStart() int32 {
	if o == nil || IsNil(o.Start) {
		var ret int32
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041JobDescMsgCrontabLine) GetStartOk() (*int32, bool) {
	if o == nil || IsNil(o.Start) {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *V0041JobDescMsgCrontabLine) HasStart() bool {
	if o != nil && !IsNil(o.Start) {
		return true
	}

	return false
}

// SetStart gets a reference to the given int32 and assigns it to the Start field.
func (o *V0041JobDescMsgCrontabLine) SetStart(v int32) {
	o.Start = &v
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *V0041JobDescMsgCrontabLine) GetEnd() int32 {
	if o == nil || IsNil(o.End) {
		var ret int32
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041JobDescMsgCrontabLine) GetEndOk() (*int32, bool) {
	if o == nil || IsNil(o.End) {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *V0041JobDescMsgCrontabLine) HasEnd() bool {
	if o != nil && !IsNil(o.End) {
		return true
	}

	return false
}

// SetEnd gets a reference to the given int32 and assigns it to the End field.
func (o *V0041JobDescMsgCrontabLine) SetEnd(v int32) {
	o.End = &v
}

func (o V0041JobDescMsgCrontabLine) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0041JobDescMsgCrontabLine) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Start) {
		toSerialize["start"] = o.Start
	}
	if !IsNil(o.End) {
		toSerialize["end"] = o.End
	}
	return toSerialize, nil
}

type NullableV0041JobDescMsgCrontabLine struct {
	value *V0041JobDescMsgCrontabLine
	isSet bool
}

func (v NullableV0041JobDescMsgCrontabLine) Get() *V0041JobDescMsgCrontabLine {
	return v.value
}

func (v *NullableV0041JobDescMsgCrontabLine) Set(val *V0041JobDescMsgCrontabLine) {
	v.value = val
	v.isSet = true
}

func (v NullableV0041JobDescMsgCrontabLine) IsSet() bool {
	return v.isSet
}

func (v *NullableV0041JobDescMsgCrontabLine) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0041JobDescMsgCrontabLine(val *V0041JobDescMsgCrontabLine) *NullableV0041JobDescMsgCrontabLine {
	return &NullableV0041JobDescMsgCrontabLine{value: val, isSet: true}
}

func (v NullableV0041JobDescMsgCrontabLine) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0041JobDescMsgCrontabLine) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

