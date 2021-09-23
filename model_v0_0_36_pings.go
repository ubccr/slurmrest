/*
 * Slurm Rest API
 *
 * API to access and control Slurm.
 *
 * API version: 0.0.36
 * Contact: sales@schedmd.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package slurmrest

import (
	"encoding/json"
)

// V0036Pings struct for V0036Pings
type V0036Pings struct {
	// slurm errors
	Errors *[]V0036Error `json:"errors,omitempty"`
	// slurm controller pings
	Pings *[]V0036Ping `json:"pings,omitempty"`
}

// NewV0036Pings instantiates a new V0036Pings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0036Pings() *V0036Pings {
	this := V0036Pings{}
	return &this
}

// NewV0036PingsWithDefaults instantiates a new V0036Pings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0036PingsWithDefaults() *V0036Pings {
	this := V0036Pings{}
	return &this
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *V0036Pings) GetErrors() []V0036Error {
	if o == nil || o.Errors == nil {
		var ret []V0036Error
		return ret
	}
	return *o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0036Pings) GetErrorsOk() (*[]V0036Error, bool) {
	if o == nil || o.Errors == nil {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *V0036Pings) HasErrors() bool {
	if o != nil && o.Errors != nil {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []V0036Error and assigns it to the Errors field.
func (o *V0036Pings) SetErrors(v []V0036Error) {
	o.Errors = &v
}

// GetPings returns the Pings field value if set, zero value otherwise.
func (o *V0036Pings) GetPings() []V0036Ping {
	if o == nil || o.Pings == nil {
		var ret []V0036Ping
		return ret
	}
	return *o.Pings
}

// GetPingsOk returns a tuple with the Pings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0036Pings) GetPingsOk() (*[]V0036Ping, bool) {
	if o == nil || o.Pings == nil {
		return nil, false
	}
	return o.Pings, true
}

// HasPings returns a boolean if a field has been set.
func (o *V0036Pings) HasPings() bool {
	if o != nil && o.Pings != nil {
		return true
	}

	return false
}

// SetPings gets a reference to the given []V0036Ping and assigns it to the Pings field.
func (o *V0036Pings) SetPings(v []V0036Ping) {
	o.Pings = &v
}

func (o V0036Pings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Errors != nil {
		toSerialize["errors"] = o.Errors
	}
	if o.Pings != nil {
		toSerialize["pings"] = o.Pings
	}
	return json.Marshal(toSerialize)
}

type NullableV0036Pings struct {
	value *V0036Pings
	isSet bool
}

func (v NullableV0036Pings) Get() *V0036Pings {
	return v.value
}

func (v *NullableV0036Pings) Set(val *V0036Pings) {
	v.value = val
	v.isSet = true
}

func (v NullableV0036Pings) IsSet() bool {
	return v.isSet
}

func (v *NullableV0036Pings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0036Pings(val *V0036Pings) *NullableV0036Pings {
	return &NullableV0036Pings{value: val, isSet: true}
}

func (v NullableV0036Pings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0036Pings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

