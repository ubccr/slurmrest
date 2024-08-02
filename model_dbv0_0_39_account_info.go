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

// checks if the Dbv0039AccountInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dbv0039AccountInfo{}

// Dbv0039AccountInfo struct for Dbv0039AccountInfo
type Dbv0039AccountInfo struct {
	Meta *Dbv0039Meta `json:"meta,omitempty"`
	// Slurm errors
	Errors []Dbv0039Error `json:"errors,omitempty"`
	// Slurm warnings
	Warnings []Dbv0039Warning `json:"warnings,omitempty"`
	Accounts []V0039Account `json:"accounts,omitempty"`
}

// NewDbv0039AccountInfo instantiates a new Dbv0039AccountInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0039AccountInfo() *Dbv0039AccountInfo {
	this := Dbv0039AccountInfo{}
	return &this
}

// NewDbv0039AccountInfoWithDefaults instantiates a new Dbv0039AccountInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0039AccountInfoWithDefaults() *Dbv0039AccountInfo {
	this := Dbv0039AccountInfo{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *Dbv0039AccountInfo) GetMeta() Dbv0039Meta {
	if o == nil || IsNil(o.Meta) {
		var ret Dbv0039Meta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0039AccountInfo) GetMetaOk() (*Dbv0039Meta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *Dbv0039AccountInfo) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given Dbv0039Meta and assigns it to the Meta field.
func (o *Dbv0039AccountInfo) SetMeta(v Dbv0039Meta) {
	o.Meta = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *Dbv0039AccountInfo) GetErrors() []Dbv0039Error {
	if o == nil || IsNil(o.Errors) {
		var ret []Dbv0039Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0039AccountInfo) GetErrorsOk() ([]Dbv0039Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *Dbv0039AccountInfo) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Dbv0039Error and assigns it to the Errors field.
func (o *Dbv0039AccountInfo) SetErrors(v []Dbv0039Error) {
	o.Errors = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *Dbv0039AccountInfo) GetWarnings() []Dbv0039Warning {
	if o == nil || IsNil(o.Warnings) {
		var ret []Dbv0039Warning
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0039AccountInfo) GetWarningsOk() ([]Dbv0039Warning, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *Dbv0039AccountInfo) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []Dbv0039Warning and assigns it to the Warnings field.
func (o *Dbv0039AccountInfo) SetWarnings(v []Dbv0039Warning) {
	o.Warnings = v
}

// GetAccounts returns the Accounts field value if set, zero value otherwise.
func (o *Dbv0039AccountInfo) GetAccounts() []V0039Account {
	if o == nil || IsNil(o.Accounts) {
		var ret []V0039Account
		return ret
	}
	return o.Accounts
}

// GetAccountsOk returns a tuple with the Accounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0039AccountInfo) GetAccountsOk() ([]V0039Account, bool) {
	if o == nil || IsNil(o.Accounts) {
		return nil, false
	}
	return o.Accounts, true
}

// HasAccounts returns a boolean if a field has been set.
func (o *Dbv0039AccountInfo) HasAccounts() bool {
	if o != nil && !IsNil(o.Accounts) {
		return true
	}

	return false
}

// SetAccounts gets a reference to the given []V0039Account and assigns it to the Accounts field.
func (o *Dbv0039AccountInfo) SetAccounts(v []V0039Account) {
	o.Accounts = v
}

func (o Dbv0039AccountInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dbv0039AccountInfo) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Accounts) {
		toSerialize["accounts"] = o.Accounts
	}
	return toSerialize, nil
}

type NullableDbv0039AccountInfo struct {
	value *Dbv0039AccountInfo
	isSet bool
}

func (v NullableDbv0039AccountInfo) Get() *Dbv0039AccountInfo {
	return v.value
}

func (v *NullableDbv0039AccountInfo) Set(val *Dbv0039AccountInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0039AccountInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0039AccountInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0039AccountInfo(val *Dbv0039AccountInfo) *NullableDbv0039AccountInfo {
	return &NullableDbv0039AccountInfo{value: val, isSet: true}
}

func (v NullableDbv0039AccountInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0039AccountInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


