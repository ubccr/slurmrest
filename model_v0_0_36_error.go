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

// V0036Error struct for V0036Error
type V0036Error struct {
	// error message
	Error *string `json:"error,omitempty"`
	// error number
	Errno *int32 `json:"errno,omitempty"`
}

// NewV0036Error instantiates a new V0036Error object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0036Error() *V0036Error {
	this := V0036Error{}
	return &this
}

// NewV0036ErrorWithDefaults instantiates a new V0036Error object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0036ErrorWithDefaults() *V0036Error {
	this := V0036Error{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *V0036Error) GetError() string {
	if o == nil || o.Error == nil {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0036Error) GetErrorOk() (*string, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *V0036Error) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *V0036Error) SetError(v string) {
	o.Error = &v
}

// GetErrno returns the Errno field value if set, zero value otherwise.
func (o *V0036Error) GetErrno() int32 {
	if o == nil || o.Errno == nil {
		var ret int32
		return ret
	}
	return *o.Errno
}

// GetErrnoOk returns a tuple with the Errno field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0036Error) GetErrnoOk() (*int32, bool) {
	if o == nil || o.Errno == nil {
		return nil, false
	}
	return o.Errno, true
}

// HasErrno returns a boolean if a field has been set.
func (o *V0036Error) HasErrno() bool {
	if o != nil && o.Errno != nil {
		return true
	}

	return false
}

// SetErrno gets a reference to the given int32 and assigns it to the Errno field.
func (o *V0036Error) SetErrno(v int32) {
	o.Errno = &v
}

func (o V0036Error) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	if o.Errno != nil {
		toSerialize["errno"] = o.Errno
	}
	return json.Marshal(toSerialize)
}

type NullableV0036Error struct {
	value *V0036Error
	isSet bool
}

func (v NullableV0036Error) Get() *V0036Error {
	return v.value
}

func (v *NullableV0036Error) Set(val *V0036Error) {
	v.value = val
	v.isSet = true
}

func (v NullableV0036Error) IsSet() bool {
	return v.isSet
}

func (v *NullableV0036Error) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0036Error(val *V0036Error) *NullableV0036Error {
	return &NullableV0036Error{value: val, isSet: true}
}

func (v NullableV0036Error) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0036Error) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


