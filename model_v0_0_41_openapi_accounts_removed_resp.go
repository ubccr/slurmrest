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

// checks if the V0041OpenapiAccountsRemovedResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0041OpenapiAccountsRemovedResp{}

// V0041OpenapiAccountsRemovedResp struct for V0041OpenapiAccountsRemovedResp
type V0041OpenapiAccountsRemovedResp struct {
	// removed_accounts
	RemovedAccounts []string `json:"removed_accounts"`
	Meta *V0041OpenapiSlurmdbdJobsRespMeta `json:"meta,omitempty"`
	// Query errors
	Errors []V0041OpenapiSlurmdbdJobsRespErrorsInner `json:"errors,omitempty"`
	// Query warnings
	Warnings []V0041OpenapiSlurmdbdJobsRespWarningsInner `json:"warnings,omitempty"`
}

type _V0041OpenapiAccountsRemovedResp V0041OpenapiAccountsRemovedResp

// NewV0041OpenapiAccountsRemovedResp instantiates a new V0041OpenapiAccountsRemovedResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0041OpenapiAccountsRemovedResp(removedAccounts []string) *V0041OpenapiAccountsRemovedResp {
	this := V0041OpenapiAccountsRemovedResp{}
	this.RemovedAccounts = removedAccounts
	return &this
}

// NewV0041OpenapiAccountsRemovedRespWithDefaults instantiates a new V0041OpenapiAccountsRemovedResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0041OpenapiAccountsRemovedRespWithDefaults() *V0041OpenapiAccountsRemovedResp {
	this := V0041OpenapiAccountsRemovedResp{}
	return &this
}

// GetRemovedAccounts returns the RemovedAccounts field value
func (o *V0041OpenapiAccountsRemovedResp) GetRemovedAccounts() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.RemovedAccounts
}

// GetRemovedAccountsOk returns a tuple with the RemovedAccounts field value
// and a boolean to check if the value has been set.
func (o *V0041OpenapiAccountsRemovedResp) GetRemovedAccountsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RemovedAccounts, true
}

// SetRemovedAccounts sets field value
func (o *V0041OpenapiAccountsRemovedResp) SetRemovedAccounts(v []string) {
	o.RemovedAccounts = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *V0041OpenapiAccountsRemovedResp) GetMeta() V0041OpenapiSlurmdbdJobsRespMeta {
	if o == nil || IsNil(o.Meta) {
		var ret V0041OpenapiSlurmdbdJobsRespMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiAccountsRemovedResp) GetMetaOk() (*V0041OpenapiSlurmdbdJobsRespMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *V0041OpenapiAccountsRemovedResp) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given V0041OpenapiSlurmdbdJobsRespMeta and assigns it to the Meta field.
func (o *V0041OpenapiAccountsRemovedResp) SetMeta(v V0041OpenapiSlurmdbdJobsRespMeta) {
	o.Meta = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *V0041OpenapiAccountsRemovedResp) GetErrors() []V0041OpenapiSlurmdbdJobsRespErrorsInner {
	if o == nil || IsNil(o.Errors) {
		var ret []V0041OpenapiSlurmdbdJobsRespErrorsInner
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiAccountsRemovedResp) GetErrorsOk() ([]V0041OpenapiSlurmdbdJobsRespErrorsInner, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *V0041OpenapiAccountsRemovedResp) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []V0041OpenapiSlurmdbdJobsRespErrorsInner and assigns it to the Errors field.
func (o *V0041OpenapiAccountsRemovedResp) SetErrors(v []V0041OpenapiSlurmdbdJobsRespErrorsInner) {
	o.Errors = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *V0041OpenapiAccountsRemovedResp) GetWarnings() []V0041OpenapiSlurmdbdJobsRespWarningsInner {
	if o == nil || IsNil(o.Warnings) {
		var ret []V0041OpenapiSlurmdbdJobsRespWarningsInner
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiAccountsRemovedResp) GetWarningsOk() ([]V0041OpenapiSlurmdbdJobsRespWarningsInner, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *V0041OpenapiAccountsRemovedResp) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []V0041OpenapiSlurmdbdJobsRespWarningsInner and assigns it to the Warnings field.
func (o *V0041OpenapiAccountsRemovedResp) SetWarnings(v []V0041OpenapiSlurmdbdJobsRespWarningsInner) {
	o.Warnings = v
}

func (o V0041OpenapiAccountsRemovedResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0041OpenapiAccountsRemovedResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["removed_accounts"] = o.RemovedAccounts
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

func (o *V0041OpenapiAccountsRemovedResp) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"removed_accounts",
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

	varV0041OpenapiAccountsRemovedResp := _V0041OpenapiAccountsRemovedResp{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV0041OpenapiAccountsRemovedResp)

	if err != nil {
		return err
	}

	*o = V0041OpenapiAccountsRemovedResp(varV0041OpenapiAccountsRemovedResp)

	return err
}

type NullableV0041OpenapiAccountsRemovedResp struct {
	value *V0041OpenapiAccountsRemovedResp
	isSet bool
}

func (v NullableV0041OpenapiAccountsRemovedResp) Get() *V0041OpenapiAccountsRemovedResp {
	return v.value
}

func (v *NullableV0041OpenapiAccountsRemovedResp) Set(val *V0041OpenapiAccountsRemovedResp) {
	v.value = val
	v.isSet = true
}

func (v NullableV0041OpenapiAccountsRemovedResp) IsSet() bool {
	return v.isSet
}

func (v *NullableV0041OpenapiAccountsRemovedResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0041OpenapiAccountsRemovedResp(val *V0041OpenapiAccountsRemovedResp) *NullableV0041OpenapiAccountsRemovedResp {
	return &NullableV0041OpenapiAccountsRemovedResp{value: val, isSet: true}
}

func (v NullableV0041OpenapiAccountsRemovedResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0041OpenapiAccountsRemovedResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


