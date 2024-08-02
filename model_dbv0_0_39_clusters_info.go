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

// checks if the Dbv0039ClustersInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dbv0039ClustersInfo{}

// Dbv0039ClustersInfo struct for Dbv0039ClustersInfo
type Dbv0039ClustersInfo struct {
	Meta *Dbv0039Meta `json:"meta,omitempty"`
	// Slurm errors
	Errors []Dbv0039Error `json:"errors,omitempty"`
	// Slurm warnings
	Warnings []Dbv0039Warning `json:"warnings,omitempty"`
	Clusters []V0039ClusterRec `json:"clusters,omitempty"`
}

// NewDbv0039ClustersInfo instantiates a new Dbv0039ClustersInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0039ClustersInfo() *Dbv0039ClustersInfo {
	this := Dbv0039ClustersInfo{}
	return &this
}

// NewDbv0039ClustersInfoWithDefaults instantiates a new Dbv0039ClustersInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0039ClustersInfoWithDefaults() *Dbv0039ClustersInfo {
	this := Dbv0039ClustersInfo{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *Dbv0039ClustersInfo) GetMeta() Dbv0039Meta {
	if o == nil || IsNil(o.Meta) {
		var ret Dbv0039Meta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0039ClustersInfo) GetMetaOk() (*Dbv0039Meta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *Dbv0039ClustersInfo) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given Dbv0039Meta and assigns it to the Meta field.
func (o *Dbv0039ClustersInfo) SetMeta(v Dbv0039Meta) {
	o.Meta = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *Dbv0039ClustersInfo) GetErrors() []Dbv0039Error {
	if o == nil || IsNil(o.Errors) {
		var ret []Dbv0039Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0039ClustersInfo) GetErrorsOk() ([]Dbv0039Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *Dbv0039ClustersInfo) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Dbv0039Error and assigns it to the Errors field.
func (o *Dbv0039ClustersInfo) SetErrors(v []Dbv0039Error) {
	o.Errors = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *Dbv0039ClustersInfo) GetWarnings() []Dbv0039Warning {
	if o == nil || IsNil(o.Warnings) {
		var ret []Dbv0039Warning
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0039ClustersInfo) GetWarningsOk() ([]Dbv0039Warning, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *Dbv0039ClustersInfo) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []Dbv0039Warning and assigns it to the Warnings field.
func (o *Dbv0039ClustersInfo) SetWarnings(v []Dbv0039Warning) {
	o.Warnings = v
}

// GetClusters returns the Clusters field value if set, zero value otherwise.
func (o *Dbv0039ClustersInfo) GetClusters() []V0039ClusterRec {
	if o == nil || IsNil(o.Clusters) {
		var ret []V0039ClusterRec
		return ret
	}
	return o.Clusters
}

// GetClustersOk returns a tuple with the Clusters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0039ClustersInfo) GetClustersOk() ([]V0039ClusterRec, bool) {
	if o == nil || IsNil(o.Clusters) {
		return nil, false
	}
	return o.Clusters, true
}

// HasClusters returns a boolean if a field has been set.
func (o *Dbv0039ClustersInfo) HasClusters() bool {
	if o != nil && !IsNil(o.Clusters) {
		return true
	}

	return false
}

// SetClusters gets a reference to the given []V0039ClusterRec and assigns it to the Clusters field.
func (o *Dbv0039ClustersInfo) SetClusters(v []V0039ClusterRec) {
	o.Clusters = v
}

func (o Dbv0039ClustersInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dbv0039ClustersInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	if !IsNil(o.Clusters) {
		toSerialize["clusters"] = o.Clusters
	}
	return toSerialize, nil
}

type NullableDbv0039ClustersInfo struct {
	value *Dbv0039ClustersInfo
	isSet bool
}

func (v NullableDbv0039ClustersInfo) Get() *Dbv0039ClustersInfo {
	return v.value
}

func (v *NullableDbv0039ClustersInfo) Set(val *Dbv0039ClustersInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0039ClustersInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0039ClustersInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0039ClustersInfo(val *Dbv0039ClustersInfo) *NullableDbv0039ClustersInfo {
	return &NullableDbv0039ClustersInfo{value: val, isSet: true}
}

func (v NullableDbv0039ClustersInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0039ClustersInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

