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

// checks if the V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTres type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTres{}

// V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTres struct for V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTres
type V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTres struct {
	Requested *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTresRequested `json:"requested,omitempty"`
	Consumed *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTresRequested `json:"consumed,omitempty"`
	Allocated []V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTresRequestedMaxInner `json:"allocated,omitempty"`
}

// NewV0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTres instantiates a new V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTres object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTres() *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTres {
	this := V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTres{}
	return &this
}

// NewV0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTresWithDefaults instantiates a new V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTres object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTresWithDefaults() *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTres {
	this := V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTres{}
	return &this
}

// GetRequested returns the Requested field value if set, zero value otherwise.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTres) GetRequested() V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTresRequested {
	if o == nil || IsNil(o.Requested) {
		var ret V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTresRequested
		return ret
	}
	return *o.Requested
}

// GetRequestedOk returns a tuple with the Requested field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTres) GetRequestedOk() (*V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTresRequested, bool) {
	if o == nil || IsNil(o.Requested) {
		return nil, false
	}
	return o.Requested, true
}

// HasRequested returns a boolean if a field has been set.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTres) HasRequested() bool {
	if o != nil && !IsNil(o.Requested) {
		return true
	}

	return false
}

// SetRequested gets a reference to the given V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTresRequested and assigns it to the Requested field.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTres) SetRequested(v V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTresRequested) {
	o.Requested = &v
}

// GetConsumed returns the Consumed field value if set, zero value otherwise.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTres) GetConsumed() V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTresRequested {
	if o == nil || IsNil(o.Consumed) {
		var ret V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTresRequested
		return ret
	}
	return *o.Consumed
}

// GetConsumedOk returns a tuple with the Consumed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTres) GetConsumedOk() (*V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTresRequested, bool) {
	if o == nil || IsNil(o.Consumed) {
		return nil, false
	}
	return o.Consumed, true
}

// HasConsumed returns a boolean if a field has been set.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTres) HasConsumed() bool {
	if o != nil && !IsNil(o.Consumed) {
		return true
	}

	return false
}

// SetConsumed gets a reference to the given V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTresRequested and assigns it to the Consumed field.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTres) SetConsumed(v V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTresRequested) {
	o.Consumed = &v
}

// GetAllocated returns the Allocated field value if set, zero value otherwise.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTres) GetAllocated() []V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTresRequestedMaxInner {
	if o == nil || IsNil(o.Allocated) {
		var ret []V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTresRequestedMaxInner
		return ret
	}
	return o.Allocated
}

// GetAllocatedOk returns a tuple with the Allocated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTres) GetAllocatedOk() ([]V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTresRequestedMaxInner, bool) {
	if o == nil || IsNil(o.Allocated) {
		return nil, false
	}
	return o.Allocated, true
}

// HasAllocated returns a boolean if a field has been set.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTres) HasAllocated() bool {
	if o != nil && !IsNil(o.Allocated) {
		return true
	}

	return false
}

// SetAllocated gets a reference to the given []V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTresRequestedMaxInner and assigns it to the Allocated field.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTres) SetAllocated(v []V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTresRequestedMaxInner) {
	o.Allocated = v
}

func (o V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTres) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTres) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Requested) {
		toSerialize["requested"] = o.Requested
	}
	if !IsNil(o.Consumed) {
		toSerialize["consumed"] = o.Consumed
	}
	if !IsNil(o.Allocated) {
		toSerialize["allocated"] = o.Allocated
	}
	return toSerialize, nil
}

type NullableV0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTres struct {
	value *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTres
	isSet bool
}

func (v NullableV0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTres) Get() *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTres {
	return v.value
}

func (v *NullableV0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTres) Set(val *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTres) {
	v.value = val
	v.isSet = true
}

func (v NullableV0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTres) IsSet() bool {
	return v.isSet
}

func (v *NullableV0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTres) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTres(val *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTres) *NullableV0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTres {
	return &NullableV0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTres{value: val, isSet: true}
}

func (v NullableV0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTres) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTres) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


