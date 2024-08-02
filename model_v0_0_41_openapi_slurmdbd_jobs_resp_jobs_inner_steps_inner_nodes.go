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

// checks if the V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerNodes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerNodes{}

// V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerNodes struct for V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerNodes
type V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerNodes struct {
	Count *int32 `json:"count,omitempty"`
	Range *string `json:"range,omitempty"`
	List []string `json:"list,omitempty"`
}

// NewV0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerNodes instantiates a new V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerNodes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerNodes() *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerNodes {
	this := V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerNodes{}
	return &this
}

// NewV0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerNodesWithDefaults instantiates a new V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerNodes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerNodesWithDefaults() *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerNodes {
	this := V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerNodes{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerNodes) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerNodes) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerNodes) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerNodes) SetCount(v int32) {
	o.Count = &v
}

// GetRange returns the Range field value if set, zero value otherwise.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerNodes) GetRange() string {
	if o == nil || IsNil(o.Range) {
		var ret string
		return ret
	}
	return *o.Range
}

// GetRangeOk returns a tuple with the Range field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerNodes) GetRangeOk() (*string, bool) {
	if o == nil || IsNil(o.Range) {
		return nil, false
	}
	return o.Range, true
}

// HasRange returns a boolean if a field has been set.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerNodes) HasRange() bool {
	if o != nil && !IsNil(o.Range) {
		return true
	}

	return false
}

// SetRange gets a reference to the given string and assigns it to the Range field.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerNodes) SetRange(v string) {
	o.Range = &v
}

// GetList returns the List field value if set, zero value otherwise.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerNodes) GetList() []string {
	if o == nil || IsNil(o.List) {
		var ret []string
		return ret
	}
	return o.List
}

// GetListOk returns a tuple with the List field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerNodes) GetListOk() ([]string, bool) {
	if o == nil || IsNil(o.List) {
		return nil, false
	}
	return o.List, true
}

// HasList returns a boolean if a field has been set.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerNodes) HasList() bool {
	if o != nil && !IsNil(o.List) {
		return true
	}

	return false
}

// SetList gets a reference to the given []string and assigns it to the List field.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerNodes) SetList(v []string) {
	o.List = v
}

func (o V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerNodes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerNodes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.Range) {
		toSerialize["range"] = o.Range
	}
	if !IsNil(o.List) {
		toSerialize["list"] = o.List
	}
	return toSerialize, nil
}

type NullableV0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerNodes struct {
	value *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerNodes
	isSet bool
}

func (v NullableV0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerNodes) Get() *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerNodes {
	return v.value
}

func (v *NullableV0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerNodes) Set(val *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerNodes) {
	v.value = val
	v.isSet = true
}

func (v NullableV0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerNodes) IsSet() bool {
	return v.isSet
}

func (v *NullableV0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerNodes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerNodes(val *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerNodes) *NullableV0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerNodes {
	return &NullableV0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerNodes{value: val, isSet: true}
}

func (v NullableV0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerNodes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerNodes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

