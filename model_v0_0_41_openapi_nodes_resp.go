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
	"bytes"
	"fmt"
)

// checks if the V0041OpenapiNodesResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0041OpenapiNodesResp{}

// V0041OpenapiNodesResp struct for V0041OpenapiNodesResp
type V0041OpenapiNodesResp struct {
	// list of nodes
	Nodes []V0041OpenapiNodesRespNodesInner `json:"nodes"`
	LastUpdate V0041OpenapiNodesRespLastUpdate `json:"last_update"`
	Meta *V0041OpenapiSlurmdbdJobsRespMeta `json:"meta,omitempty"`
	// Query errors
	Errors []V0041OpenapiSlurmdbdJobsRespErrorsInner `json:"errors,omitempty"`
	// Query warnings
	Warnings []V0041OpenapiSlurmdbdJobsRespWarningsInner `json:"warnings,omitempty"`
}

type _V0041OpenapiNodesResp V0041OpenapiNodesResp

// NewV0041OpenapiNodesResp instantiates a new V0041OpenapiNodesResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0041OpenapiNodesResp(nodes []V0041OpenapiNodesRespNodesInner, lastUpdate V0041OpenapiNodesRespLastUpdate) *V0041OpenapiNodesResp {
	this := V0041OpenapiNodesResp{}
	this.Nodes = nodes
	this.LastUpdate = lastUpdate
	return &this
}

// NewV0041OpenapiNodesRespWithDefaults instantiates a new V0041OpenapiNodesResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0041OpenapiNodesRespWithDefaults() *V0041OpenapiNodesResp {
	this := V0041OpenapiNodesResp{}
	return &this
}

// GetNodes returns the Nodes field value
func (o *V0041OpenapiNodesResp) GetNodes() []V0041OpenapiNodesRespNodesInner {
	if o == nil {
		var ret []V0041OpenapiNodesRespNodesInner
		return ret
	}

	return o.Nodes
}

// GetNodesOk returns a tuple with the Nodes field value
// and a boolean to check if the value has been set.
func (o *V0041OpenapiNodesResp) GetNodesOk() ([]V0041OpenapiNodesRespNodesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Nodes, true
}

// SetNodes sets field value
func (o *V0041OpenapiNodesResp) SetNodes(v []V0041OpenapiNodesRespNodesInner) {
	o.Nodes = v
}

// GetLastUpdate returns the LastUpdate field value
func (o *V0041OpenapiNodesResp) GetLastUpdate() V0041OpenapiNodesRespLastUpdate {
	if o == nil {
		var ret V0041OpenapiNodesRespLastUpdate
		return ret
	}

	return o.LastUpdate
}

// GetLastUpdateOk returns a tuple with the LastUpdate field value
// and a boolean to check if the value has been set.
func (o *V0041OpenapiNodesResp) GetLastUpdateOk() (*V0041OpenapiNodesRespLastUpdate, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdate, true
}

// SetLastUpdate sets field value
func (o *V0041OpenapiNodesResp) SetLastUpdate(v V0041OpenapiNodesRespLastUpdate) {
	o.LastUpdate = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *V0041OpenapiNodesResp) GetMeta() V0041OpenapiSlurmdbdJobsRespMeta {
	if o == nil || IsNil(o.Meta) {
		var ret V0041OpenapiSlurmdbdJobsRespMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiNodesResp) GetMetaOk() (*V0041OpenapiSlurmdbdJobsRespMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *V0041OpenapiNodesResp) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given V0041OpenapiSlurmdbdJobsRespMeta and assigns it to the Meta field.
func (o *V0041OpenapiNodesResp) SetMeta(v V0041OpenapiSlurmdbdJobsRespMeta) {
	o.Meta = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *V0041OpenapiNodesResp) GetErrors() []V0041OpenapiSlurmdbdJobsRespErrorsInner {
	if o == nil || IsNil(o.Errors) {
		var ret []V0041OpenapiSlurmdbdJobsRespErrorsInner
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiNodesResp) GetErrorsOk() ([]V0041OpenapiSlurmdbdJobsRespErrorsInner, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *V0041OpenapiNodesResp) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []V0041OpenapiSlurmdbdJobsRespErrorsInner and assigns it to the Errors field.
func (o *V0041OpenapiNodesResp) SetErrors(v []V0041OpenapiSlurmdbdJobsRespErrorsInner) {
	o.Errors = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *V0041OpenapiNodesResp) GetWarnings() []V0041OpenapiSlurmdbdJobsRespWarningsInner {
	if o == nil || IsNil(o.Warnings) {
		var ret []V0041OpenapiSlurmdbdJobsRespWarningsInner
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiNodesResp) GetWarningsOk() ([]V0041OpenapiSlurmdbdJobsRespWarningsInner, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *V0041OpenapiNodesResp) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []V0041OpenapiSlurmdbdJobsRespWarningsInner and assigns it to the Warnings field.
func (o *V0041OpenapiNodesResp) SetWarnings(v []V0041OpenapiSlurmdbdJobsRespWarningsInner) {
	o.Warnings = v
}

func (o V0041OpenapiNodesResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0041OpenapiNodesResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["nodes"] = o.Nodes
	toSerialize["last_update"] = o.LastUpdate
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

func (o *V0041OpenapiNodesResp) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"nodes",
		"last_update",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varV0041OpenapiNodesResp := _V0041OpenapiNodesResp{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV0041OpenapiNodesResp)

	if err != nil {
		return err
	}

	*o = V0041OpenapiNodesResp(varV0041OpenapiNodesResp)

	return err
}

type NullableV0041OpenapiNodesResp struct {
	value *V0041OpenapiNodesResp
	isSet bool
}

func (v NullableV0041OpenapiNodesResp) Get() *V0041OpenapiNodesResp {
	return v.value
}

func (v *NullableV0041OpenapiNodesResp) Set(val *V0041OpenapiNodesResp) {
	v.value = val
	v.isSet = true
}

func (v NullableV0041OpenapiNodesResp) IsSet() bool {
	return v.isSet
}

func (v *NullableV0041OpenapiNodesResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0041OpenapiNodesResp(val *V0041OpenapiNodesResp) *NullableV0041OpenapiNodesResp {
	return &NullableV0041OpenapiNodesResp{value: val, isSet: true}
}

func (v NullableV0041OpenapiNodesResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0041OpenapiNodesResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


