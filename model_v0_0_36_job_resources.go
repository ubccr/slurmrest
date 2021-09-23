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

// V0036JobResources struct for V0036JobResources
type V0036JobResources struct {
	// list of assigned job nodes
	Nodes *string `json:"nodes,omitempty"`
	// number of assigned job cpus
	AllocatedCpus *int32 `json:"allocated_cpus,omitempty"`
	// number of assigned job hosts
	AllocatedHosts *int32 `json:"allocated_hosts,omitempty"`
	// node allocations
	AllocatedNodes *map[string]V0036NodeAllocation `json:"allocated_nodes,omitempty"`
}

// NewV0036JobResources instantiates a new V0036JobResources object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0036JobResources() *V0036JobResources {
	this := V0036JobResources{}
	return &this
}

// NewV0036JobResourcesWithDefaults instantiates a new V0036JobResources object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0036JobResourcesWithDefaults() *V0036JobResources {
	this := V0036JobResources{}
	return &this
}

// GetNodes returns the Nodes field value if set, zero value otherwise.
func (o *V0036JobResources) GetNodes() string {
	if o == nil || o.Nodes == nil {
		var ret string
		return ret
	}
	return *o.Nodes
}

// GetNodesOk returns a tuple with the Nodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0036JobResources) GetNodesOk() (*string, bool) {
	if o == nil || o.Nodes == nil {
		return nil, false
	}
	return o.Nodes, true
}

// HasNodes returns a boolean if a field has been set.
func (o *V0036JobResources) HasNodes() bool {
	if o != nil && o.Nodes != nil {
		return true
	}

	return false
}

// SetNodes gets a reference to the given string and assigns it to the Nodes field.
func (o *V0036JobResources) SetNodes(v string) {
	o.Nodes = &v
}

// GetAllocatedCpus returns the AllocatedCpus field value if set, zero value otherwise.
func (o *V0036JobResources) GetAllocatedCpus() int32 {
	if o == nil || o.AllocatedCpus == nil {
		var ret int32
		return ret
	}
	return *o.AllocatedCpus
}

// GetAllocatedCpusOk returns a tuple with the AllocatedCpus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0036JobResources) GetAllocatedCpusOk() (*int32, bool) {
	if o == nil || o.AllocatedCpus == nil {
		return nil, false
	}
	return o.AllocatedCpus, true
}

// HasAllocatedCpus returns a boolean if a field has been set.
func (o *V0036JobResources) HasAllocatedCpus() bool {
	if o != nil && o.AllocatedCpus != nil {
		return true
	}

	return false
}

// SetAllocatedCpus gets a reference to the given int32 and assigns it to the AllocatedCpus field.
func (o *V0036JobResources) SetAllocatedCpus(v int32) {
	o.AllocatedCpus = &v
}

// GetAllocatedHosts returns the AllocatedHosts field value if set, zero value otherwise.
func (o *V0036JobResources) GetAllocatedHosts() int32 {
	if o == nil || o.AllocatedHosts == nil {
		var ret int32
		return ret
	}
	return *o.AllocatedHosts
}

// GetAllocatedHostsOk returns a tuple with the AllocatedHosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0036JobResources) GetAllocatedHostsOk() (*int32, bool) {
	if o == nil || o.AllocatedHosts == nil {
		return nil, false
	}
	return o.AllocatedHosts, true
}

// HasAllocatedHosts returns a boolean if a field has been set.
func (o *V0036JobResources) HasAllocatedHosts() bool {
	if o != nil && o.AllocatedHosts != nil {
		return true
	}

	return false
}

// SetAllocatedHosts gets a reference to the given int32 and assigns it to the AllocatedHosts field.
func (o *V0036JobResources) SetAllocatedHosts(v int32) {
	o.AllocatedHosts = &v
}

// GetAllocatedNodes returns the AllocatedNodes field value if set, zero value otherwise.
func (o *V0036JobResources) GetAllocatedNodes() map[string]V0036NodeAllocation {
	if o == nil || o.AllocatedNodes == nil {
		var ret map[string]V0036NodeAllocation
		return ret
	}
	return *o.AllocatedNodes
}

// GetAllocatedNodesOk returns a tuple with the AllocatedNodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0036JobResources) GetAllocatedNodesOk() (*map[string]V0036NodeAllocation, bool) {
	if o == nil || o.AllocatedNodes == nil {
		return nil, false
	}
	return o.AllocatedNodes, true
}

// HasAllocatedNodes returns a boolean if a field has been set.
func (o *V0036JobResources) HasAllocatedNodes() bool {
	if o != nil && o.AllocatedNodes != nil {
		return true
	}

	return false
}

// SetAllocatedNodes gets a reference to the given map[string]V0036NodeAllocation and assigns it to the AllocatedNodes field.
func (o *V0036JobResources) SetAllocatedNodes(v map[string]V0036NodeAllocation) {
	o.AllocatedNodes = &v
}

func (o V0036JobResources) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Nodes != nil {
		toSerialize["nodes"] = o.Nodes
	}
	if o.AllocatedCpus != nil {
		toSerialize["allocated_cpus"] = o.AllocatedCpus
	}
	if o.AllocatedHosts != nil {
		toSerialize["allocated_hosts"] = o.AllocatedHosts
	}
	if o.AllocatedNodes != nil {
		toSerialize["allocated_nodes"] = o.AllocatedNodes
	}
	return json.Marshal(toSerialize)
}

type NullableV0036JobResources struct {
	value *V0036JobResources
	isSet bool
}

func (v NullableV0036JobResources) Get() *V0036JobResources {
	return v.value
}

func (v *NullableV0036JobResources) Set(val *V0036JobResources) {
	v.value = val
	v.isSet = true
}

func (v NullableV0036JobResources) IsSet() bool {
	return v.isSet
}

func (v *NullableV0036JobResources) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0036JobResources(val *V0036JobResources) *NullableV0036JobResources {
	return &NullableV0036JobResources{value: val, isSet: true}
}

func (v NullableV0036JobResources) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0036JobResources) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


