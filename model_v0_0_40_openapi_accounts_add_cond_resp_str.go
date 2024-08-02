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

// checks if the V0040OpenapiAccountsAddCondRespStr type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0040OpenapiAccountsAddCondRespStr{}

// V0040OpenapiAccountsAddCondRespStr struct for V0040OpenapiAccountsAddCondRespStr
type V0040OpenapiAccountsAddCondRespStr struct {
	// added_accounts
	AddedAccounts string `json:"added_accounts"`
	Meta *V0040OpenapiMeta `json:"meta,omitempty"`
	Errors []V0040OpenapiError `json:"errors,omitempty"`
	Warnings []V0040OpenapiWarning `json:"warnings,omitempty"`
}

type _V0040OpenapiAccountsAddCondRespStr V0040OpenapiAccountsAddCondRespStr

// NewV0040OpenapiAccountsAddCondRespStr instantiates a new V0040OpenapiAccountsAddCondRespStr object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0040OpenapiAccountsAddCondRespStr(addedAccounts string) *V0040OpenapiAccountsAddCondRespStr {
	this := V0040OpenapiAccountsAddCondRespStr{}
	this.AddedAccounts = addedAccounts
	return &this
}

// NewV0040OpenapiAccountsAddCondRespStrWithDefaults instantiates a new V0040OpenapiAccountsAddCondRespStr object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0040OpenapiAccountsAddCondRespStrWithDefaults() *V0040OpenapiAccountsAddCondRespStr {
	this := V0040OpenapiAccountsAddCondRespStr{}
	return &this
}

// GetAddedAccounts returns the AddedAccounts field value
func (o *V0040OpenapiAccountsAddCondRespStr) GetAddedAccounts() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AddedAccounts
}

// GetAddedAccountsOk returns a tuple with the AddedAccounts field value
// and a boolean to check if the value has been set.
func (o *V0040OpenapiAccountsAddCondRespStr) GetAddedAccountsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AddedAccounts, true
}

// SetAddedAccounts sets field value
func (o *V0040OpenapiAccountsAddCondRespStr) SetAddedAccounts(v string) {
	o.AddedAccounts = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *V0040OpenapiAccountsAddCondRespStr) GetMeta() V0040OpenapiMeta {
	if o == nil || IsNil(o.Meta) {
		var ret V0040OpenapiMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040OpenapiAccountsAddCondRespStr) GetMetaOk() (*V0040OpenapiMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *V0040OpenapiAccountsAddCondRespStr) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given V0040OpenapiMeta and assigns it to the Meta field.
func (o *V0040OpenapiAccountsAddCondRespStr) SetMeta(v V0040OpenapiMeta) {
	o.Meta = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *V0040OpenapiAccountsAddCondRespStr) GetErrors() []V0040OpenapiError {
	if o == nil || IsNil(o.Errors) {
		var ret []V0040OpenapiError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040OpenapiAccountsAddCondRespStr) GetErrorsOk() ([]V0040OpenapiError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *V0040OpenapiAccountsAddCondRespStr) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []V0040OpenapiError and assigns it to the Errors field.
func (o *V0040OpenapiAccountsAddCondRespStr) SetErrors(v []V0040OpenapiError) {
	o.Errors = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *V0040OpenapiAccountsAddCondRespStr) GetWarnings() []V0040OpenapiWarning {
	if o == nil || IsNil(o.Warnings) {
		var ret []V0040OpenapiWarning
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0040OpenapiAccountsAddCondRespStr) GetWarningsOk() ([]V0040OpenapiWarning, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *V0040OpenapiAccountsAddCondRespStr) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []V0040OpenapiWarning and assigns it to the Warnings field.
func (o *V0040OpenapiAccountsAddCondRespStr) SetWarnings(v []V0040OpenapiWarning) {
	o.Warnings = v
}

func (o V0040OpenapiAccountsAddCondRespStr) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0040OpenapiAccountsAddCondRespStr) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["added_accounts"] = o.AddedAccounts
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

func (o *V0040OpenapiAccountsAddCondRespStr) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"added_accounts",
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

	varV0040OpenapiAccountsAddCondRespStr := _V0040OpenapiAccountsAddCondRespStr{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV0040OpenapiAccountsAddCondRespStr)

	if err != nil {
		return err
	}

	*o = V0040OpenapiAccountsAddCondRespStr(varV0040OpenapiAccountsAddCondRespStr)

	return err
}

type NullableV0040OpenapiAccountsAddCondRespStr struct {
	value *V0040OpenapiAccountsAddCondRespStr
	isSet bool
}

func (v NullableV0040OpenapiAccountsAddCondRespStr) Get() *V0040OpenapiAccountsAddCondRespStr {
	return v.value
}

func (v *NullableV0040OpenapiAccountsAddCondRespStr) Set(val *V0040OpenapiAccountsAddCondRespStr) {
	v.value = val
	v.isSet = true
}

func (v NullableV0040OpenapiAccountsAddCondRespStr) IsSet() bool {
	return v.isSet
}

func (v *NullableV0040OpenapiAccountsAddCondRespStr) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0040OpenapiAccountsAddCondRespStr(val *V0040OpenapiAccountsAddCondRespStr) *NullableV0040OpenapiAccountsAddCondRespStr {
	return &NullableV0040OpenapiAccountsAddCondRespStr{value: val, isSet: true}
}

func (v NullableV0040OpenapiAccountsAddCondRespStr) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0040OpenapiAccountsAddCondRespStr) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


