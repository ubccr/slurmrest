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

// checks if the V0041OpenapiSharesRespSharesSharesInnerTresRunSecondsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0041OpenapiSharesRespSharesSharesInnerTresRunSecondsInner{}

// V0041OpenapiSharesRespSharesSharesInnerTresRunSecondsInner struct for V0041OpenapiSharesRespSharesSharesInnerTresRunSecondsInner
type V0041OpenapiSharesRespSharesSharesInnerTresRunSecondsInner struct {
	// TRES name
	Name *string `json:"name,omitempty"`
	Value *V0041OpenapiSharesRespSharesSharesInnerTresRunSecondsInnerValue `json:"value,omitempty"`
}

// NewV0041OpenapiSharesRespSharesSharesInnerTresRunSecondsInner instantiates a new V0041OpenapiSharesRespSharesSharesInnerTresRunSecondsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0041OpenapiSharesRespSharesSharesInnerTresRunSecondsInner() *V0041OpenapiSharesRespSharesSharesInnerTresRunSecondsInner {
	this := V0041OpenapiSharesRespSharesSharesInnerTresRunSecondsInner{}
	return &this
}

// NewV0041OpenapiSharesRespSharesSharesInnerTresRunSecondsInnerWithDefaults instantiates a new V0041OpenapiSharesRespSharesSharesInnerTresRunSecondsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0041OpenapiSharesRespSharesSharesInnerTresRunSecondsInnerWithDefaults() *V0041OpenapiSharesRespSharesSharesInnerTresRunSecondsInner {
	this := V0041OpenapiSharesRespSharesSharesInnerTresRunSecondsInner{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *V0041OpenapiSharesRespSharesSharesInnerTresRunSecondsInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiSharesRespSharesSharesInnerTresRunSecondsInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *V0041OpenapiSharesRespSharesSharesInnerTresRunSecondsInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *V0041OpenapiSharesRespSharesSharesInnerTresRunSecondsInner) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *V0041OpenapiSharesRespSharesSharesInnerTresRunSecondsInner) GetValue() V0041OpenapiSharesRespSharesSharesInnerTresRunSecondsInnerValue {
	if o == nil || IsNil(o.Value) {
		var ret V0041OpenapiSharesRespSharesSharesInnerTresRunSecondsInnerValue
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiSharesRespSharesSharesInnerTresRunSecondsInner) GetValueOk() (*V0041OpenapiSharesRespSharesSharesInnerTresRunSecondsInnerValue, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *V0041OpenapiSharesRespSharesSharesInnerTresRunSecondsInner) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given V0041OpenapiSharesRespSharesSharesInnerTresRunSecondsInnerValue and assigns it to the Value field.
func (o *V0041OpenapiSharesRespSharesSharesInnerTresRunSecondsInner) SetValue(v V0041OpenapiSharesRespSharesSharesInnerTresRunSecondsInnerValue) {
	o.Value = &v
}

func (o V0041OpenapiSharesRespSharesSharesInnerTresRunSecondsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0041OpenapiSharesRespSharesSharesInnerTresRunSecondsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableV0041OpenapiSharesRespSharesSharesInnerTresRunSecondsInner struct {
	value *V0041OpenapiSharesRespSharesSharesInnerTresRunSecondsInner
	isSet bool
}

func (v NullableV0041OpenapiSharesRespSharesSharesInnerTresRunSecondsInner) Get() *V0041OpenapiSharesRespSharesSharesInnerTresRunSecondsInner {
	return v.value
}

func (v *NullableV0041OpenapiSharesRespSharesSharesInnerTresRunSecondsInner) Set(val *V0041OpenapiSharesRespSharesSharesInnerTresRunSecondsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableV0041OpenapiSharesRespSharesSharesInnerTresRunSecondsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableV0041OpenapiSharesRespSharesSharesInnerTresRunSecondsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0041OpenapiSharesRespSharesSharesInnerTresRunSecondsInner(val *V0041OpenapiSharesRespSharesSharesInnerTresRunSecondsInner) *NullableV0041OpenapiSharesRespSharesSharesInnerTresRunSecondsInner {
	return &NullableV0041OpenapiSharesRespSharesSharesInnerTresRunSecondsInner{value: val, isSet: true}
}

func (v NullableV0041OpenapiSharesRespSharesSharesInnerTresRunSecondsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0041OpenapiSharesRespSharesSharesInnerTresRunSecondsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

