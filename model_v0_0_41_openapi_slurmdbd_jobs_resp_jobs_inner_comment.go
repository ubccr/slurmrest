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

// checks if the V0041OpenapiSlurmdbdJobsRespJobsInnerComment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0041OpenapiSlurmdbdJobsRespJobsInnerComment{}

// V0041OpenapiSlurmdbdJobsRespJobsInnerComment struct for V0041OpenapiSlurmdbdJobsRespJobsInnerComment
type V0041OpenapiSlurmdbdJobsRespJobsInnerComment struct {
	Administrator *string `json:"administrator,omitempty"`
	Job *string `json:"job,omitempty"`
	System *string `json:"system,omitempty"`
}

// NewV0041OpenapiSlurmdbdJobsRespJobsInnerComment instantiates a new V0041OpenapiSlurmdbdJobsRespJobsInnerComment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0041OpenapiSlurmdbdJobsRespJobsInnerComment() *V0041OpenapiSlurmdbdJobsRespJobsInnerComment {
	this := V0041OpenapiSlurmdbdJobsRespJobsInnerComment{}
	return &this
}

// NewV0041OpenapiSlurmdbdJobsRespJobsInnerCommentWithDefaults instantiates a new V0041OpenapiSlurmdbdJobsRespJobsInnerComment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0041OpenapiSlurmdbdJobsRespJobsInnerCommentWithDefaults() *V0041OpenapiSlurmdbdJobsRespJobsInnerComment {
	this := V0041OpenapiSlurmdbdJobsRespJobsInnerComment{}
	return &this
}

// GetAdministrator returns the Administrator field value if set, zero value otherwise.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerComment) GetAdministrator() string {
	if o == nil || IsNil(o.Administrator) {
		var ret string
		return ret
	}
	return *o.Administrator
}

// GetAdministratorOk returns a tuple with the Administrator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerComment) GetAdministratorOk() (*string, bool) {
	if o == nil || IsNil(o.Administrator) {
		return nil, false
	}
	return o.Administrator, true
}

// HasAdministrator returns a boolean if a field has been set.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerComment) HasAdministrator() bool {
	if o != nil && !IsNil(o.Administrator) {
		return true
	}

	return false
}

// SetAdministrator gets a reference to the given string and assigns it to the Administrator field.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerComment) SetAdministrator(v string) {
	o.Administrator = &v
}

// GetJob returns the Job field value if set, zero value otherwise.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerComment) GetJob() string {
	if o == nil || IsNil(o.Job) {
		var ret string
		return ret
	}
	return *o.Job
}

// GetJobOk returns a tuple with the Job field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerComment) GetJobOk() (*string, bool) {
	if o == nil || IsNil(o.Job) {
		return nil, false
	}
	return o.Job, true
}

// HasJob returns a boolean if a field has been set.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerComment) HasJob() bool {
	if o != nil && !IsNil(o.Job) {
		return true
	}

	return false
}

// SetJob gets a reference to the given string and assigns it to the Job field.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerComment) SetJob(v string) {
	o.Job = &v
}

// GetSystem returns the System field value if set, zero value otherwise.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerComment) GetSystem() string {
	if o == nil || IsNil(o.System) {
		var ret string
		return ret
	}
	return *o.System
}

// GetSystemOk returns a tuple with the System field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerComment) GetSystemOk() (*string, bool) {
	if o == nil || IsNil(o.System) {
		return nil, false
	}
	return o.System, true
}

// HasSystem returns a boolean if a field has been set.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerComment) HasSystem() bool {
	if o != nil && !IsNil(o.System) {
		return true
	}

	return false
}

// SetSystem gets a reference to the given string and assigns it to the System field.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerComment) SetSystem(v string) {
	o.System = &v
}

func (o V0041OpenapiSlurmdbdJobsRespJobsInnerComment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0041OpenapiSlurmdbdJobsRespJobsInnerComment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Administrator) {
		toSerialize["administrator"] = o.Administrator
	}
	if !IsNil(o.Job) {
		toSerialize["job"] = o.Job
	}
	if !IsNil(o.System) {
		toSerialize["system"] = o.System
	}
	return toSerialize, nil
}

type NullableV0041OpenapiSlurmdbdJobsRespJobsInnerComment struct {
	value *V0041OpenapiSlurmdbdJobsRespJobsInnerComment
	isSet bool
}

func (v NullableV0041OpenapiSlurmdbdJobsRespJobsInnerComment) Get() *V0041OpenapiSlurmdbdJobsRespJobsInnerComment {
	return v.value
}

func (v *NullableV0041OpenapiSlurmdbdJobsRespJobsInnerComment) Set(val *V0041OpenapiSlurmdbdJobsRespJobsInnerComment) {
	v.value = val
	v.isSet = true
}

func (v NullableV0041OpenapiSlurmdbdJobsRespJobsInnerComment) IsSet() bool {
	return v.isSet
}

func (v *NullableV0041OpenapiSlurmdbdJobsRespJobsInnerComment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0041OpenapiSlurmdbdJobsRespJobsInnerComment(val *V0041OpenapiSlurmdbdJobsRespJobsInnerComment) *NullableV0041OpenapiSlurmdbdJobsRespJobsInnerComment {
	return &NullableV0041OpenapiSlurmdbdJobsRespJobsInnerComment{value: val, isSet: true}
}

func (v NullableV0041OpenapiSlurmdbdJobsRespJobsInnerComment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0041OpenapiSlurmdbdJobsRespJobsInnerComment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


