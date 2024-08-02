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

// checks if the V0039JobExitCode type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0039JobExitCode{}

// V0039JobExitCode job exit details
type V0039JobExitCode struct {
	// exit status
	Status *string `json:"status,omitempty"`
	// return code (numeric)
	ReturnCode *int32 `json:"return_code,omitempty"`
	Signal *V0039JobExitCodeSignal `json:"signal,omitempty"`
}

// NewV0039JobExitCode instantiates a new V0039JobExitCode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0039JobExitCode() *V0039JobExitCode {
	this := V0039JobExitCode{}
	return &this
}

// NewV0039JobExitCodeWithDefaults instantiates a new V0039JobExitCode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0039JobExitCodeWithDefaults() *V0039JobExitCode {
	this := V0039JobExitCode{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *V0039JobExitCode) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039JobExitCode) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *V0039JobExitCode) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *V0039JobExitCode) SetStatus(v string) {
	o.Status = &v
}

// GetReturnCode returns the ReturnCode field value if set, zero value otherwise.
func (o *V0039JobExitCode) GetReturnCode() int32 {
	if o == nil || IsNil(o.ReturnCode) {
		var ret int32
		return ret
	}
	return *o.ReturnCode
}

// GetReturnCodeOk returns a tuple with the ReturnCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039JobExitCode) GetReturnCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.ReturnCode) {
		return nil, false
	}
	return o.ReturnCode, true
}

// HasReturnCode returns a boolean if a field has been set.
func (o *V0039JobExitCode) HasReturnCode() bool {
	if o != nil && !IsNil(o.ReturnCode) {
		return true
	}

	return false
}

// SetReturnCode gets a reference to the given int32 and assigns it to the ReturnCode field.
func (o *V0039JobExitCode) SetReturnCode(v int32) {
	o.ReturnCode = &v
}

// GetSignal returns the Signal field value if set, zero value otherwise.
func (o *V0039JobExitCode) GetSignal() V0039JobExitCodeSignal {
	if o == nil || IsNil(o.Signal) {
		var ret V0039JobExitCodeSignal
		return ret
	}
	return *o.Signal
}

// GetSignalOk returns a tuple with the Signal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039JobExitCode) GetSignalOk() (*V0039JobExitCodeSignal, bool) {
	if o == nil || IsNil(o.Signal) {
		return nil, false
	}
	return o.Signal, true
}

// HasSignal returns a boolean if a field has been set.
func (o *V0039JobExitCode) HasSignal() bool {
	if o != nil && !IsNil(o.Signal) {
		return true
	}

	return false
}

// SetSignal gets a reference to the given V0039JobExitCodeSignal and assigns it to the Signal field.
func (o *V0039JobExitCode) SetSignal(v V0039JobExitCodeSignal) {
	o.Signal = &v
}

func (o V0039JobExitCode) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0039JobExitCode) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.ReturnCode) {
		toSerialize["return_code"] = o.ReturnCode
	}
	if !IsNil(o.Signal) {
		toSerialize["signal"] = o.Signal
	}
	return toSerialize, nil
}

type NullableV0039JobExitCode struct {
	value *V0039JobExitCode
	isSet bool
}

func (v NullableV0039JobExitCode) Get() *V0039JobExitCode {
	return v.value
}

func (v *NullableV0039JobExitCode) Set(val *V0039JobExitCode) {
	v.value = val
	v.isSet = true
}

func (v NullableV0039JobExitCode) IsSet() bool {
	return v.isSet
}

func (v *NullableV0039JobExitCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0039JobExitCode(val *V0039JobExitCode) *NullableV0039JobExitCode {
	return &NullableV0039JobExitCode{value: val, isSet: true}
}

func (v NullableV0039JobExitCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0039JobExitCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


