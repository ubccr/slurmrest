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

// checks if the V0041OpenapiSlurmdbdJobsRespJobsInnerHet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0041OpenapiSlurmdbdJobsRespJobsInnerHet{}

// V0041OpenapiSlurmdbdJobsRespJobsInnerHet struct for V0041OpenapiSlurmdbdJobsRespJobsInnerHet
type V0041OpenapiSlurmdbdJobsRespJobsInnerHet struct {
	JobId *int32 `json:"job_id,omitempty"`
	JobOffset *V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId `json:"job_offset,omitempty"`
}

// NewV0041OpenapiSlurmdbdJobsRespJobsInnerHet instantiates a new V0041OpenapiSlurmdbdJobsRespJobsInnerHet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0041OpenapiSlurmdbdJobsRespJobsInnerHet() *V0041OpenapiSlurmdbdJobsRespJobsInnerHet {
	this := V0041OpenapiSlurmdbdJobsRespJobsInnerHet{}
	return &this
}

// NewV0041OpenapiSlurmdbdJobsRespJobsInnerHetWithDefaults instantiates a new V0041OpenapiSlurmdbdJobsRespJobsInnerHet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0041OpenapiSlurmdbdJobsRespJobsInnerHetWithDefaults() *V0041OpenapiSlurmdbdJobsRespJobsInnerHet {
	this := V0041OpenapiSlurmdbdJobsRespJobsInnerHet{}
	return &this
}

// GetJobId returns the JobId field value if set, zero value otherwise.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerHet) GetJobId() int32 {
	if o == nil || IsNil(o.JobId) {
		var ret int32
		return ret
	}
	return *o.JobId
}

// GetJobIdOk returns a tuple with the JobId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerHet) GetJobIdOk() (*int32, bool) {
	if o == nil || IsNil(o.JobId) {
		return nil, false
	}
	return o.JobId, true
}

// HasJobId returns a boolean if a field has been set.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerHet) HasJobId() bool {
	if o != nil && !IsNil(o.JobId) {
		return true
	}

	return false
}

// SetJobId gets a reference to the given int32 and assigns it to the JobId field.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerHet) SetJobId(v int32) {
	o.JobId = &v
}

// GetJobOffset returns the JobOffset field value if set, zero value otherwise.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerHet) GetJobOffset() V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId {
	if o == nil || IsNil(o.JobOffset) {
		var ret V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId
		return ret
	}
	return *o.JobOffset
}

// GetJobOffsetOk returns a tuple with the JobOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerHet) GetJobOffsetOk() (*V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId, bool) {
	if o == nil || IsNil(o.JobOffset) {
		return nil, false
	}
	return o.JobOffset, true
}

// HasJobOffset returns a boolean if a field has been set.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerHet) HasJobOffset() bool {
	if o != nil && !IsNil(o.JobOffset) {
		return true
	}

	return false
}

// SetJobOffset gets a reference to the given V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId and assigns it to the JobOffset field.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerHet) SetJobOffset(v V0041OpenapiSlurmdbdJobsRespJobsInnerArrayTaskId) {
	o.JobOffset = &v
}

func (o V0041OpenapiSlurmdbdJobsRespJobsInnerHet) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0041OpenapiSlurmdbdJobsRespJobsInnerHet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.JobId) {
		toSerialize["job_id"] = o.JobId
	}
	if !IsNil(o.JobOffset) {
		toSerialize["job_offset"] = o.JobOffset
	}
	return toSerialize, nil
}

type NullableV0041OpenapiSlurmdbdJobsRespJobsInnerHet struct {
	value *V0041OpenapiSlurmdbdJobsRespJobsInnerHet
	isSet bool
}

func (v NullableV0041OpenapiSlurmdbdJobsRespJobsInnerHet) Get() *V0041OpenapiSlurmdbdJobsRespJobsInnerHet {
	return v.value
}

func (v *NullableV0041OpenapiSlurmdbdJobsRespJobsInnerHet) Set(val *V0041OpenapiSlurmdbdJobsRespJobsInnerHet) {
	v.value = val
	v.isSet = true
}

func (v NullableV0041OpenapiSlurmdbdJobsRespJobsInnerHet) IsSet() bool {
	return v.isSet
}

func (v *NullableV0041OpenapiSlurmdbdJobsRespJobsInnerHet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0041OpenapiSlurmdbdJobsRespJobsInnerHet(val *V0041OpenapiSlurmdbdJobsRespJobsInnerHet) *NullableV0041OpenapiSlurmdbdJobsRespJobsInnerHet {
	return &NullableV0041OpenapiSlurmdbdJobsRespJobsInnerHet{value: val, isSet: true}
}

func (v NullableV0041OpenapiSlurmdbdJobsRespJobsInnerHet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0041OpenapiSlurmdbdJobsRespJobsInnerHet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

