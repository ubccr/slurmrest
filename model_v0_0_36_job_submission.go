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

// V0036JobSubmission struct for V0036JobSubmission
type V0036JobSubmission struct {
	// Executable script (full contents) to run in batch step
	Script string `json:"script"`
	Job *V0036JobProperties `json:"job,omitempty"`
	// Properties of an HetJob
	Jobs *[]V0036JobProperties `json:"jobs,omitempty"`
}

// NewV0036JobSubmission instantiates a new V0036JobSubmission object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0036JobSubmission(script string) *V0036JobSubmission {
	this := V0036JobSubmission{}
	this.Script = script
	return &this
}

// NewV0036JobSubmissionWithDefaults instantiates a new V0036JobSubmission object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0036JobSubmissionWithDefaults() *V0036JobSubmission {
	this := V0036JobSubmission{}
	return &this
}

// GetScript returns the Script field value
func (o *V0036JobSubmission) GetScript() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Script
}

// GetScriptOk returns a tuple with the Script field value
// and a boolean to check if the value has been set.
func (o *V0036JobSubmission) GetScriptOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Script, true
}

// SetScript sets field value
func (o *V0036JobSubmission) SetScript(v string) {
	o.Script = v
}

// GetJob returns the Job field value if set, zero value otherwise.
func (o *V0036JobSubmission) GetJob() V0036JobProperties {
	if o == nil || o.Job == nil {
		var ret V0036JobProperties
		return ret
	}
	return *o.Job
}

// GetJobOk returns a tuple with the Job field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0036JobSubmission) GetJobOk() (*V0036JobProperties, bool) {
	if o == nil || o.Job == nil {
		return nil, false
	}
	return o.Job, true
}

// HasJob returns a boolean if a field has been set.
func (o *V0036JobSubmission) HasJob() bool {
	if o != nil && o.Job != nil {
		return true
	}

	return false
}

// SetJob gets a reference to the given V0036JobProperties and assigns it to the Job field.
func (o *V0036JobSubmission) SetJob(v V0036JobProperties) {
	o.Job = &v
}

// GetJobs returns the Jobs field value if set, zero value otherwise.
func (o *V0036JobSubmission) GetJobs() []V0036JobProperties {
	if o == nil || o.Jobs == nil {
		var ret []V0036JobProperties
		return ret
	}
	return *o.Jobs
}

// GetJobsOk returns a tuple with the Jobs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0036JobSubmission) GetJobsOk() (*[]V0036JobProperties, bool) {
	if o == nil || o.Jobs == nil {
		return nil, false
	}
	return o.Jobs, true
}

// HasJobs returns a boolean if a field has been set.
func (o *V0036JobSubmission) HasJobs() bool {
	if o != nil && o.Jobs != nil {
		return true
	}

	return false
}

// SetJobs gets a reference to the given []V0036JobProperties and assigns it to the Jobs field.
func (o *V0036JobSubmission) SetJobs(v []V0036JobProperties) {
	o.Jobs = &v
}

func (o V0036JobSubmission) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["script"] = o.Script
	}
	if o.Job != nil {
		toSerialize["job"] = o.Job
	}
	if o.Jobs != nil {
		toSerialize["jobs"] = o.Jobs
	}
	return json.Marshal(toSerialize)
}

type NullableV0036JobSubmission struct {
	value *V0036JobSubmission
	isSet bool
}

func (v NullableV0036JobSubmission) Get() *V0036JobSubmission {
	return v.value
}

func (v *NullableV0036JobSubmission) Set(val *V0036JobSubmission) {
	v.value = val
	v.isSet = true
}

func (v NullableV0036JobSubmission) IsSet() bool {
	return v.isSet
}

func (v *NullableV0036JobSubmission) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0036JobSubmission(val *V0036JobSubmission) *NullableV0036JobSubmission {
	return &NullableV0036JobSubmission{value: val, isSet: true}
}

func (v NullableV0036JobSubmission) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0036JobSubmission) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


