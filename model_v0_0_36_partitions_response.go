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

// V0036PartitionsResponse struct for V0036PartitionsResponse
type V0036PartitionsResponse struct {
	// slurm errors
	Errors *[]V0036Error `json:"errors,omitempty"`
	// partition info
	Partitions *[]V0036Partition `json:"partitions,omitempty"`
}

// NewV0036PartitionsResponse instantiates a new V0036PartitionsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0036PartitionsResponse() *V0036PartitionsResponse {
	this := V0036PartitionsResponse{}
	return &this
}

// NewV0036PartitionsResponseWithDefaults instantiates a new V0036PartitionsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0036PartitionsResponseWithDefaults() *V0036PartitionsResponse {
	this := V0036PartitionsResponse{}
	return &this
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *V0036PartitionsResponse) GetErrors() []V0036Error {
	if o == nil || o.Errors == nil {
		var ret []V0036Error
		return ret
	}
	return *o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0036PartitionsResponse) GetErrorsOk() (*[]V0036Error, bool) {
	if o == nil || o.Errors == nil {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *V0036PartitionsResponse) HasErrors() bool {
	if o != nil && o.Errors != nil {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []V0036Error and assigns it to the Errors field.
func (o *V0036PartitionsResponse) SetErrors(v []V0036Error) {
	o.Errors = &v
}

// GetPartitions returns the Partitions field value if set, zero value otherwise.
func (o *V0036PartitionsResponse) GetPartitions() []V0036Partition {
	if o == nil || o.Partitions == nil {
		var ret []V0036Partition
		return ret
	}
	return *o.Partitions
}

// GetPartitionsOk returns a tuple with the Partitions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0036PartitionsResponse) GetPartitionsOk() (*[]V0036Partition, bool) {
	if o == nil || o.Partitions == nil {
		return nil, false
	}
	return o.Partitions, true
}

// HasPartitions returns a boolean if a field has been set.
func (o *V0036PartitionsResponse) HasPartitions() bool {
	if o != nil && o.Partitions != nil {
		return true
	}

	return false
}

// SetPartitions gets a reference to the given []V0036Partition and assigns it to the Partitions field.
func (o *V0036PartitionsResponse) SetPartitions(v []V0036Partition) {
	o.Partitions = &v
}

func (o V0036PartitionsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Errors != nil {
		toSerialize["errors"] = o.Errors
	}
	if o.Partitions != nil {
		toSerialize["partitions"] = o.Partitions
	}
	return json.Marshal(toSerialize)
}

type NullableV0036PartitionsResponse struct {
	value *V0036PartitionsResponse
	isSet bool
}

func (v NullableV0036PartitionsResponse) Get() *V0036PartitionsResponse {
	return v.value
}

func (v *NullableV0036PartitionsResponse) Set(val *V0036PartitionsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableV0036PartitionsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableV0036PartitionsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0036PartitionsResponse(val *V0036PartitionsResponse) *NullableV0036PartitionsResponse {
	return &NullableV0036PartitionsResponse{value: val, isSet: true}
}

func (v NullableV0036PartitionsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0036PartitionsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

