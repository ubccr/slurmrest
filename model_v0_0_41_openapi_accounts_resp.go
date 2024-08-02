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

// checks if the V0041OpenapiAccountsResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0041OpenapiAccountsResp{}

// V0041OpenapiAccountsResp struct for V0041OpenapiAccountsResp
type V0041OpenapiAccountsResp struct {
	// accounts
	Accounts []V0041OpenapiSlurmdbdConfigRespAccountsInner `json:"accounts"`
	Meta *V0041OpenapiSlurmdbdJobsRespMeta `json:"meta,omitempty"`
	// Query errors
	Errors []V0041OpenapiSlurmdbdJobsRespErrorsInner `json:"errors,omitempty"`
	// Query warnings
	Warnings []V0041OpenapiSlurmdbdJobsRespWarningsInner `json:"warnings,omitempty"`
}

type _V0041OpenapiAccountsResp V0041OpenapiAccountsResp

// NewV0041OpenapiAccountsResp instantiates a new V0041OpenapiAccountsResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0041OpenapiAccountsResp(accounts []V0041OpenapiSlurmdbdConfigRespAccountsInner) *V0041OpenapiAccountsResp {
	this := V0041OpenapiAccountsResp{}
	this.Accounts = accounts
	return &this
}

// NewV0041OpenapiAccountsRespWithDefaults instantiates a new V0041OpenapiAccountsResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0041OpenapiAccountsRespWithDefaults() *V0041OpenapiAccountsResp {
	this := V0041OpenapiAccountsResp{}
	return &this
}

// GetAccounts returns the Accounts field value
func (o *V0041OpenapiAccountsResp) GetAccounts() []V0041OpenapiSlurmdbdConfigRespAccountsInner {
	if o == nil {
		var ret []V0041OpenapiSlurmdbdConfigRespAccountsInner
		return ret
	}

	return o.Accounts
}

// GetAccountsOk returns a tuple with the Accounts field value
// and a boolean to check if the value has been set.
func (o *V0041OpenapiAccountsResp) GetAccountsOk() ([]V0041OpenapiSlurmdbdConfigRespAccountsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Accounts, true
}

// SetAccounts sets field value
func (o *V0041OpenapiAccountsResp) SetAccounts(v []V0041OpenapiSlurmdbdConfigRespAccountsInner) {
	o.Accounts = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *V0041OpenapiAccountsResp) GetMeta() V0041OpenapiSlurmdbdJobsRespMeta {
	if o == nil || IsNil(o.Meta) {
		var ret V0041OpenapiSlurmdbdJobsRespMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiAccountsResp) GetMetaOk() (*V0041OpenapiSlurmdbdJobsRespMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *V0041OpenapiAccountsResp) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given V0041OpenapiSlurmdbdJobsRespMeta and assigns it to the Meta field.
func (o *V0041OpenapiAccountsResp) SetMeta(v V0041OpenapiSlurmdbdJobsRespMeta) {
	o.Meta = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *V0041OpenapiAccountsResp) GetErrors() []V0041OpenapiSlurmdbdJobsRespErrorsInner {
	if o == nil || IsNil(o.Errors) {
		var ret []V0041OpenapiSlurmdbdJobsRespErrorsInner
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiAccountsResp) GetErrorsOk() ([]V0041OpenapiSlurmdbdJobsRespErrorsInner, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *V0041OpenapiAccountsResp) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []V0041OpenapiSlurmdbdJobsRespErrorsInner and assigns it to the Errors field.
func (o *V0041OpenapiAccountsResp) SetErrors(v []V0041OpenapiSlurmdbdJobsRespErrorsInner) {
	o.Errors = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *V0041OpenapiAccountsResp) GetWarnings() []V0041OpenapiSlurmdbdJobsRespWarningsInner {
	if o == nil || IsNil(o.Warnings) {
		var ret []V0041OpenapiSlurmdbdJobsRespWarningsInner
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiAccountsResp) GetWarningsOk() ([]V0041OpenapiSlurmdbdJobsRespWarningsInner, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *V0041OpenapiAccountsResp) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []V0041OpenapiSlurmdbdJobsRespWarningsInner and assigns it to the Warnings field.
func (o *V0041OpenapiAccountsResp) SetWarnings(v []V0041OpenapiSlurmdbdJobsRespWarningsInner) {
	o.Warnings = v
}

func (o V0041OpenapiAccountsResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0041OpenapiAccountsResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["accounts"] = o.Accounts
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

func (o *V0041OpenapiAccountsResp) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"accounts",
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

	varV0041OpenapiAccountsResp := _V0041OpenapiAccountsResp{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV0041OpenapiAccountsResp)

	if err != nil {
		return err
	}

	*o = V0041OpenapiAccountsResp(varV0041OpenapiAccountsResp)

	return err
}

type NullableV0041OpenapiAccountsResp struct {
	value *V0041OpenapiAccountsResp
	isSet bool
}

func (v NullableV0041OpenapiAccountsResp) Get() *V0041OpenapiAccountsResp {
	return v.value
}

func (v *NullableV0041OpenapiAccountsResp) Set(val *V0041OpenapiAccountsResp) {
	v.value = val
	v.isSet = true
}

func (v NullableV0041OpenapiAccountsResp) IsSet() bool {
	return v.isSet
}

func (v *NullableV0041OpenapiAccountsResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0041OpenapiAccountsResp(val *V0041OpenapiAccountsResp) *NullableV0041OpenapiAccountsResp {
	return &NullableV0041OpenapiAccountsResp{value: val, isSet: true}
}

func (v NullableV0041OpenapiAccountsResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0041OpenapiAccountsResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

