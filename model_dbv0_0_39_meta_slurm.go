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

// checks if the Dbv0039MetaSlurm type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dbv0039MetaSlurm{}

// Dbv0039MetaSlurm Slurm information
type Dbv0039MetaSlurm struct {
	Version *Dbv0039MetaSlurmVersion `json:"version,omitempty"`
	// version specifier
	Release *string `json:"release,omitempty"`
}

// NewDbv0039MetaSlurm instantiates a new Dbv0039MetaSlurm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0039MetaSlurm() *Dbv0039MetaSlurm {
	this := Dbv0039MetaSlurm{}
	return &this
}

// NewDbv0039MetaSlurmWithDefaults instantiates a new Dbv0039MetaSlurm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0039MetaSlurmWithDefaults() *Dbv0039MetaSlurm {
	this := Dbv0039MetaSlurm{}
	return &this
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *Dbv0039MetaSlurm) GetVersion() Dbv0039MetaSlurmVersion {
	if o == nil || IsNil(o.Version) {
		var ret Dbv0039MetaSlurmVersion
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0039MetaSlurm) GetVersionOk() (*Dbv0039MetaSlurmVersion, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *Dbv0039MetaSlurm) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given Dbv0039MetaSlurmVersion and assigns it to the Version field.
func (o *Dbv0039MetaSlurm) SetVersion(v Dbv0039MetaSlurmVersion) {
	o.Version = &v
}

// GetRelease returns the Release field value if set, zero value otherwise.
func (o *Dbv0039MetaSlurm) GetRelease() string {
	if o == nil || IsNil(o.Release) {
		var ret string
		return ret
	}
	return *o.Release
}

// GetReleaseOk returns a tuple with the Release field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0039MetaSlurm) GetReleaseOk() (*string, bool) {
	if o == nil || IsNil(o.Release) {
		return nil, false
	}
	return o.Release, true
}

// HasRelease returns a boolean if a field has been set.
func (o *Dbv0039MetaSlurm) HasRelease() bool {
	if o != nil && !IsNil(o.Release) {
		return true
	}

	return false
}

// SetRelease gets a reference to the given string and assigns it to the Release field.
func (o *Dbv0039MetaSlurm) SetRelease(v string) {
	o.Release = &v
}

func (o Dbv0039MetaSlurm) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dbv0039MetaSlurm) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.Release) {
		toSerialize["release"] = o.Release
	}
	return toSerialize, nil
}

type NullableDbv0039MetaSlurm struct {
	value *Dbv0039MetaSlurm
	isSet bool
}

func (v NullableDbv0039MetaSlurm) Get() *Dbv0039MetaSlurm {
	return v.value
}

func (v *NullableDbv0039MetaSlurm) Set(val *Dbv0039MetaSlurm) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0039MetaSlurm) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0039MetaSlurm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0039MetaSlurm(val *Dbv0039MetaSlurm) *NullableDbv0039MetaSlurm {
	return &NullableDbv0039MetaSlurm{value: val, isSet: true}
}

func (v NullableDbv0039MetaSlurm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0039MetaSlurm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


