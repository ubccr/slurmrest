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

// checks if the V0039Pings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0039Pings{}

// V0039Pings struct for V0039Pings
type V0039Pings struct {
	Meta *V0039Meta `json:"meta,omitempty"`
	// Slurm errors
	Errors []V0039Error `json:"errors,omitempty"`
	// Slurm warnings
	Warnings []V0039Warning `json:"warnings,omitempty"`
	Pings []V0039ControllerPing `json:"pings,omitempty"`
}

// NewV0039Pings instantiates a new V0039Pings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0039Pings() *V0039Pings {
	this := V0039Pings{}
	return &this
}

// NewV0039PingsWithDefaults instantiates a new V0039Pings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0039PingsWithDefaults() *V0039Pings {
	this := V0039Pings{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *V0039Pings) GetMeta() V0039Meta {
	if o == nil || IsNil(o.Meta) {
		var ret V0039Meta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039Pings) GetMetaOk() (*V0039Meta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *V0039Pings) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given V0039Meta and assigns it to the Meta field.
func (o *V0039Pings) SetMeta(v V0039Meta) {
	o.Meta = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *V0039Pings) GetErrors() []V0039Error {
	if o == nil || IsNil(o.Errors) {
		var ret []V0039Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039Pings) GetErrorsOk() ([]V0039Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *V0039Pings) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []V0039Error and assigns it to the Errors field.
func (o *V0039Pings) SetErrors(v []V0039Error) {
	o.Errors = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *V0039Pings) GetWarnings() []V0039Warning {
	if o == nil || IsNil(o.Warnings) {
		var ret []V0039Warning
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039Pings) GetWarningsOk() ([]V0039Warning, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *V0039Pings) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []V0039Warning and assigns it to the Warnings field.
func (o *V0039Pings) SetWarnings(v []V0039Warning) {
	o.Warnings = v
}

// GetPings returns the Pings field value if set, zero value otherwise.
func (o *V0039Pings) GetPings() []V0039ControllerPing {
	if o == nil || IsNil(o.Pings) {
		var ret []V0039ControllerPing
		return ret
	}
	return o.Pings
}

// GetPingsOk returns a tuple with the Pings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039Pings) GetPingsOk() ([]V0039ControllerPing, bool) {
	if o == nil || IsNil(o.Pings) {
		return nil, false
	}
	return o.Pings, true
}

// HasPings returns a boolean if a field has been set.
func (o *V0039Pings) HasPings() bool {
	if o != nil && !IsNil(o.Pings) {
		return true
	}

	return false
}

// SetPings gets a reference to the given []V0039ControllerPing and assigns it to the Pings field.
func (o *V0039Pings) SetPings(v []V0039ControllerPing) {
	o.Pings = v
}

func (o V0039Pings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0039Pings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	if !IsNil(o.Pings) {
		toSerialize["pings"] = o.Pings
	}
	return toSerialize, nil
}

type NullableV0039Pings struct {
	value *V0039Pings
	isSet bool
}

func (v NullableV0039Pings) Get() *V0039Pings {
	return v.value
}

func (v *NullableV0039Pings) Set(val *V0039Pings) {
	v.value = val
	v.isSet = true
}

func (v NullableV0039Pings) IsSet() bool {
	return v.isSet
}

func (v *NullableV0039Pings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0039Pings(val *V0039Pings) *NullableV0039Pings {
	return &NullableV0039Pings{value: val, isSet: true}
}

func (v NullableV0039Pings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0039Pings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


