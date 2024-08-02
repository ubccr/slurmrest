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

// checks if the V0041OpenapiKillJobsRespStatusInnerError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0041OpenapiKillJobsRespStatusInnerError{}

// V0041OpenapiKillJobsRespStatusInnerError struct for V0041OpenapiKillJobsRespStatusInnerError
type V0041OpenapiKillJobsRespStatusInnerError struct {
	// String error encountered signaling job
	String *string `json:"string,omitempty"`
	// Numeric error encountered signaling job
	Code *int32 `json:"code,omitempty"`
	// Error message why signaling job failed
	Message *string `json:"message,omitempty"`
}

// NewV0041OpenapiKillJobsRespStatusInnerError instantiates a new V0041OpenapiKillJobsRespStatusInnerError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0041OpenapiKillJobsRespStatusInnerError() *V0041OpenapiKillJobsRespStatusInnerError {
	this := V0041OpenapiKillJobsRespStatusInnerError{}
	return &this
}

// NewV0041OpenapiKillJobsRespStatusInnerErrorWithDefaults instantiates a new V0041OpenapiKillJobsRespStatusInnerError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0041OpenapiKillJobsRespStatusInnerErrorWithDefaults() *V0041OpenapiKillJobsRespStatusInnerError {
	this := V0041OpenapiKillJobsRespStatusInnerError{}
	return &this
}

// GetString returns the String field value if set, zero value otherwise.
func (o *V0041OpenapiKillJobsRespStatusInnerError) GetString() string {
	if o == nil || IsNil(o.String) {
		var ret string
		return ret
	}
	return *o.String
}

// GetStringOk returns a tuple with the String field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiKillJobsRespStatusInnerError) GetStringOk() (*string, bool) {
	if o == nil || IsNil(o.String) {
		return nil, false
	}
	return o.String, true
}

// HasString returns a boolean if a field has been set.
func (o *V0041OpenapiKillJobsRespStatusInnerError) HasString() bool {
	if o != nil && !IsNil(o.String) {
		return true
	}

	return false
}

// SetString gets a reference to the given string and assigns it to the String field.
func (o *V0041OpenapiKillJobsRespStatusInnerError) SetString(v string) {
	o.String = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *V0041OpenapiKillJobsRespStatusInnerError) GetCode() int32 {
	if o == nil || IsNil(o.Code) {
		var ret int32
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiKillJobsRespStatusInnerError) GetCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *V0041OpenapiKillJobsRespStatusInnerError) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given int32 and assigns it to the Code field.
func (o *V0041OpenapiKillJobsRespStatusInnerError) SetCode(v int32) {
	o.Code = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *V0041OpenapiKillJobsRespStatusInnerError) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiKillJobsRespStatusInnerError) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *V0041OpenapiKillJobsRespStatusInnerError) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *V0041OpenapiKillJobsRespStatusInnerError) SetMessage(v string) {
	o.Message = &v
}

func (o V0041OpenapiKillJobsRespStatusInnerError) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0041OpenapiKillJobsRespStatusInnerError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.String) {
		toSerialize["string"] = o.String
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableV0041OpenapiKillJobsRespStatusInnerError struct {
	value *V0041OpenapiKillJobsRespStatusInnerError
	isSet bool
}

func (v NullableV0041OpenapiKillJobsRespStatusInnerError) Get() *V0041OpenapiKillJobsRespStatusInnerError {
	return v.value
}

func (v *NullableV0041OpenapiKillJobsRespStatusInnerError) Set(val *V0041OpenapiKillJobsRespStatusInnerError) {
	v.value = val
	v.isSet = true
}

func (v NullableV0041OpenapiKillJobsRespStatusInnerError) IsSet() bool {
	return v.isSet
}

func (v *NullableV0041OpenapiKillJobsRespStatusInnerError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0041OpenapiKillJobsRespStatusInnerError(val *V0041OpenapiKillJobsRespStatusInnerError) *NullableV0041OpenapiKillJobsRespStatusInnerError {
	return &NullableV0041OpenapiKillJobsRespStatusInnerError{value: val, isSet: true}
}

func (v NullableV0041OpenapiKillJobsRespStatusInnerError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0041OpenapiKillJobsRespStatusInnerError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

