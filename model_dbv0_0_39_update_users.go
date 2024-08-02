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

// checks if the Dbv0039UpdateUsers type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dbv0039UpdateUsers{}

// Dbv0039UpdateUsers struct for Dbv0039UpdateUsers
type Dbv0039UpdateUsers struct {
	Users []V0039User `json:"users,omitempty"`
}

// NewDbv0039UpdateUsers instantiates a new Dbv0039UpdateUsers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0039UpdateUsers() *Dbv0039UpdateUsers {
	this := Dbv0039UpdateUsers{}
	return &this
}

// NewDbv0039UpdateUsersWithDefaults instantiates a new Dbv0039UpdateUsers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0039UpdateUsersWithDefaults() *Dbv0039UpdateUsers {
	this := Dbv0039UpdateUsers{}
	return &this
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *Dbv0039UpdateUsers) GetUsers() []V0039User {
	if o == nil || IsNil(o.Users) {
		var ret []V0039User
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0039UpdateUsers) GetUsersOk() ([]V0039User, bool) {
	if o == nil || IsNil(o.Users) {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *Dbv0039UpdateUsers) HasUsers() bool {
	if o != nil && !IsNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []V0039User and assigns it to the Users field.
func (o *Dbv0039UpdateUsers) SetUsers(v []V0039User) {
	o.Users = v
}

func (o Dbv0039UpdateUsers) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dbv0039UpdateUsers) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Users) {
		toSerialize["users"] = o.Users
	}
	return toSerialize, nil
}

type NullableDbv0039UpdateUsers struct {
	value *Dbv0039UpdateUsers
	isSet bool
}

func (v NullableDbv0039UpdateUsers) Get() *Dbv0039UpdateUsers {
	return v.value
}

func (v *NullableDbv0039UpdateUsers) Set(val *Dbv0039UpdateUsers) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0039UpdateUsers) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0039UpdateUsers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0039UpdateUsers(val *Dbv0039UpdateUsers) *NullableDbv0039UpdateUsers {
	return &NullableDbv0039UpdateUsers{value: val, isSet: true}
}

func (v NullableDbv0039UpdateUsers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0039UpdateUsers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


