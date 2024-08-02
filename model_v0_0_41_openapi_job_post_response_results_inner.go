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

// checks if the V0041OpenapiJobPostResponseResultsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0041OpenapiJobPostResponseResultsInner{}

// V0041OpenapiJobPostResponseResultsInner struct for V0041OpenapiJobPostResponseResultsInner
type V0041OpenapiJobPostResponseResultsInner struct {
	// JobId for updated Job
	JobId *int32 `json:"job_id,omitempty"`
	// StepId for updated Job
	StepId *string `json:"step_id,omitempty"`
	// Verbose update status or error
	Error *string `json:"error,omitempty"`
	// Verbose update status or error
	ErrorCode *int32 `json:"error_code,omitempty"`
	// Update response message
	Why *string `json:"why,omitempty"`
}

// NewV0041OpenapiJobPostResponseResultsInner instantiates a new V0041OpenapiJobPostResponseResultsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0041OpenapiJobPostResponseResultsInner() *V0041OpenapiJobPostResponseResultsInner {
	this := V0041OpenapiJobPostResponseResultsInner{}
	return &this
}

// NewV0041OpenapiJobPostResponseResultsInnerWithDefaults instantiates a new V0041OpenapiJobPostResponseResultsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0041OpenapiJobPostResponseResultsInnerWithDefaults() *V0041OpenapiJobPostResponseResultsInner {
	this := V0041OpenapiJobPostResponseResultsInner{}
	return &this
}

// GetJobId returns the JobId field value if set, zero value otherwise.
func (o *V0041OpenapiJobPostResponseResultsInner) GetJobId() int32 {
	if o == nil || IsNil(o.JobId) {
		var ret int32
		return ret
	}
	return *o.JobId
}

// GetJobIdOk returns a tuple with the JobId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiJobPostResponseResultsInner) GetJobIdOk() (*int32, bool) {
	if o == nil || IsNil(o.JobId) {
		return nil, false
	}
	return o.JobId, true
}

// HasJobId returns a boolean if a field has been set.
func (o *V0041OpenapiJobPostResponseResultsInner) HasJobId() bool {
	if o != nil && !IsNil(o.JobId) {
		return true
	}

	return false
}

// SetJobId gets a reference to the given int32 and assigns it to the JobId field.
func (o *V0041OpenapiJobPostResponseResultsInner) SetJobId(v int32) {
	o.JobId = &v
}

// GetStepId returns the StepId field value if set, zero value otherwise.
func (o *V0041OpenapiJobPostResponseResultsInner) GetStepId() string {
	if o == nil || IsNil(o.StepId) {
		var ret string
		return ret
	}
	return *o.StepId
}

// GetStepIdOk returns a tuple with the StepId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiJobPostResponseResultsInner) GetStepIdOk() (*string, bool) {
	if o == nil || IsNil(o.StepId) {
		return nil, false
	}
	return o.StepId, true
}

// HasStepId returns a boolean if a field has been set.
func (o *V0041OpenapiJobPostResponseResultsInner) HasStepId() bool {
	if o != nil && !IsNil(o.StepId) {
		return true
	}

	return false
}

// SetStepId gets a reference to the given string and assigns it to the StepId field.
func (o *V0041OpenapiJobPostResponseResultsInner) SetStepId(v string) {
	o.StepId = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *V0041OpenapiJobPostResponseResultsInner) GetError() string {
	if o == nil || IsNil(o.Error) {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiJobPostResponseResultsInner) GetErrorOk() (*string, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *V0041OpenapiJobPostResponseResultsInner) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *V0041OpenapiJobPostResponseResultsInner) SetError(v string) {
	o.Error = &v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *V0041OpenapiJobPostResponseResultsInner) GetErrorCode() int32 {
	if o == nil || IsNil(o.ErrorCode) {
		var ret int32
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiJobPostResponseResultsInner) GetErrorCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.ErrorCode) {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *V0041OpenapiJobPostResponseResultsInner) HasErrorCode() bool {
	if o != nil && !IsNil(o.ErrorCode) {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given int32 and assigns it to the ErrorCode field.
func (o *V0041OpenapiJobPostResponseResultsInner) SetErrorCode(v int32) {
	o.ErrorCode = &v
}

// GetWhy returns the Why field value if set, zero value otherwise.
func (o *V0041OpenapiJobPostResponseResultsInner) GetWhy() string {
	if o == nil || IsNil(o.Why) {
		var ret string
		return ret
	}
	return *o.Why
}

// GetWhyOk returns a tuple with the Why field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiJobPostResponseResultsInner) GetWhyOk() (*string, bool) {
	if o == nil || IsNil(o.Why) {
		return nil, false
	}
	return o.Why, true
}

// HasWhy returns a boolean if a field has been set.
func (o *V0041OpenapiJobPostResponseResultsInner) HasWhy() bool {
	if o != nil && !IsNil(o.Why) {
		return true
	}

	return false
}

// SetWhy gets a reference to the given string and assigns it to the Why field.
func (o *V0041OpenapiJobPostResponseResultsInner) SetWhy(v string) {
	o.Why = &v
}

func (o V0041OpenapiJobPostResponseResultsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0041OpenapiJobPostResponseResultsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.JobId) {
		toSerialize["job_id"] = o.JobId
	}
	if !IsNil(o.StepId) {
		toSerialize["step_id"] = o.StepId
	}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	if !IsNil(o.ErrorCode) {
		toSerialize["error_code"] = o.ErrorCode
	}
	if !IsNil(o.Why) {
		toSerialize["why"] = o.Why
	}
	return toSerialize, nil
}

type NullableV0041OpenapiJobPostResponseResultsInner struct {
	value *V0041OpenapiJobPostResponseResultsInner
	isSet bool
}

func (v NullableV0041OpenapiJobPostResponseResultsInner) Get() *V0041OpenapiJobPostResponseResultsInner {
	return v.value
}

func (v *NullableV0041OpenapiJobPostResponseResultsInner) Set(val *V0041OpenapiJobPostResponseResultsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableV0041OpenapiJobPostResponseResultsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableV0041OpenapiJobPostResponseResultsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0041OpenapiJobPostResponseResultsInner(val *V0041OpenapiJobPostResponseResultsInner) *NullableV0041OpenapiJobPostResponseResultsInner {
	return &NullableV0041OpenapiJobPostResponseResultsInner{value: val, isSet: true}
}

func (v NullableV0041OpenapiJobPostResponseResultsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0041OpenapiJobPostResponseResultsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


