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

// checks if the V0041OpenapiAccountsAddCondResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0041OpenapiAccountsAddCondResp{}

// V0041OpenapiAccountsAddCondResp struct for V0041OpenapiAccountsAddCondResp
type V0041OpenapiAccountsAddCondResp struct {
	AssociationCondition *V0041OpenapiAccountsAddCondRespAssociationCondition `json:"association_condition,omitempty"`
	Account *V0041OpenapiAccountsAddCondRespAccount `json:"account,omitempty"`
	Meta *V0041OpenapiSlurmdbdJobsRespMeta `json:"meta,omitempty"`
	// Query errors
	Errors []V0041OpenapiSlurmdbdJobsRespErrorsInner `json:"errors,omitempty"`
	// Query warnings
	Warnings []V0041OpenapiSlurmdbdJobsRespWarningsInner `json:"warnings,omitempty"`
}

// NewV0041OpenapiAccountsAddCondResp instantiates a new V0041OpenapiAccountsAddCondResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0041OpenapiAccountsAddCondResp() *V0041OpenapiAccountsAddCondResp {
	this := V0041OpenapiAccountsAddCondResp{}
	return &this
}

// NewV0041OpenapiAccountsAddCondRespWithDefaults instantiates a new V0041OpenapiAccountsAddCondResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0041OpenapiAccountsAddCondRespWithDefaults() *V0041OpenapiAccountsAddCondResp {
	this := V0041OpenapiAccountsAddCondResp{}
	return &this
}

// GetAssociationCondition returns the AssociationCondition field value if set, zero value otherwise.
func (o *V0041OpenapiAccountsAddCondResp) GetAssociationCondition() V0041OpenapiAccountsAddCondRespAssociationCondition {
	if o == nil || IsNil(o.AssociationCondition) {
		var ret V0041OpenapiAccountsAddCondRespAssociationCondition
		return ret
	}
	return *o.AssociationCondition
}

// GetAssociationConditionOk returns a tuple with the AssociationCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiAccountsAddCondResp) GetAssociationConditionOk() (*V0041OpenapiAccountsAddCondRespAssociationCondition, bool) {
	if o == nil || IsNil(o.AssociationCondition) {
		return nil, false
	}
	return o.AssociationCondition, true
}

// HasAssociationCondition returns a boolean if a field has been set.
func (o *V0041OpenapiAccountsAddCondResp) HasAssociationCondition() bool {
	if o != nil && !IsNil(o.AssociationCondition) {
		return true
	}

	return false
}

// SetAssociationCondition gets a reference to the given V0041OpenapiAccountsAddCondRespAssociationCondition and assigns it to the AssociationCondition field.
func (o *V0041OpenapiAccountsAddCondResp) SetAssociationCondition(v V0041OpenapiAccountsAddCondRespAssociationCondition) {
	o.AssociationCondition = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *V0041OpenapiAccountsAddCondResp) GetAccount() V0041OpenapiAccountsAddCondRespAccount {
	if o == nil || IsNil(o.Account) {
		var ret V0041OpenapiAccountsAddCondRespAccount
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiAccountsAddCondResp) GetAccountOk() (*V0041OpenapiAccountsAddCondRespAccount, bool) {
	if o == nil || IsNil(o.Account) {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *V0041OpenapiAccountsAddCondResp) HasAccount() bool {
	if o != nil && !IsNil(o.Account) {
		return true
	}

	return false
}

// SetAccount gets a reference to the given V0041OpenapiAccountsAddCondRespAccount and assigns it to the Account field.
func (o *V0041OpenapiAccountsAddCondResp) SetAccount(v V0041OpenapiAccountsAddCondRespAccount) {
	o.Account = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *V0041OpenapiAccountsAddCondResp) GetMeta() V0041OpenapiSlurmdbdJobsRespMeta {
	if o == nil || IsNil(o.Meta) {
		var ret V0041OpenapiSlurmdbdJobsRespMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiAccountsAddCondResp) GetMetaOk() (*V0041OpenapiSlurmdbdJobsRespMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *V0041OpenapiAccountsAddCondResp) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given V0041OpenapiSlurmdbdJobsRespMeta and assigns it to the Meta field.
func (o *V0041OpenapiAccountsAddCondResp) SetMeta(v V0041OpenapiSlurmdbdJobsRespMeta) {
	o.Meta = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *V0041OpenapiAccountsAddCondResp) GetErrors() []V0041OpenapiSlurmdbdJobsRespErrorsInner {
	if o == nil || IsNil(o.Errors) {
		var ret []V0041OpenapiSlurmdbdJobsRespErrorsInner
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiAccountsAddCondResp) GetErrorsOk() ([]V0041OpenapiSlurmdbdJobsRespErrorsInner, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *V0041OpenapiAccountsAddCondResp) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []V0041OpenapiSlurmdbdJobsRespErrorsInner and assigns it to the Errors field.
func (o *V0041OpenapiAccountsAddCondResp) SetErrors(v []V0041OpenapiSlurmdbdJobsRespErrorsInner) {
	o.Errors = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *V0041OpenapiAccountsAddCondResp) GetWarnings() []V0041OpenapiSlurmdbdJobsRespWarningsInner {
	if o == nil || IsNil(o.Warnings) {
		var ret []V0041OpenapiSlurmdbdJobsRespWarningsInner
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiAccountsAddCondResp) GetWarningsOk() ([]V0041OpenapiSlurmdbdJobsRespWarningsInner, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *V0041OpenapiAccountsAddCondResp) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []V0041OpenapiSlurmdbdJobsRespWarningsInner and assigns it to the Warnings field.
func (o *V0041OpenapiAccountsAddCondResp) SetWarnings(v []V0041OpenapiSlurmdbdJobsRespWarningsInner) {
	o.Warnings = v
}

func (o V0041OpenapiAccountsAddCondResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0041OpenapiAccountsAddCondResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AssociationCondition) {
		toSerialize["association_condition"] = o.AssociationCondition
	}
	if !IsNil(o.Account) {
		toSerialize["account"] = o.Account
	}
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

type NullableV0041OpenapiAccountsAddCondResp struct {
	value *V0041OpenapiAccountsAddCondResp
	isSet bool
}

func (v NullableV0041OpenapiAccountsAddCondResp) Get() *V0041OpenapiAccountsAddCondResp {
	return v.value
}

func (v *NullableV0041OpenapiAccountsAddCondResp) Set(val *V0041OpenapiAccountsAddCondResp) {
	v.value = val
	v.isSet = true
}

func (v NullableV0041OpenapiAccountsAddCondResp) IsSet() bool {
	return v.isSet
}

func (v *NullableV0041OpenapiAccountsAddCondResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0041OpenapiAccountsAddCondResp(val *V0041OpenapiAccountsAddCondResp) *NullableV0041OpenapiAccountsAddCondResp {
	return &NullableV0041OpenapiAccountsAddCondResp{value: val, isSet: true}
}

func (v NullableV0041OpenapiAccountsAddCondResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0041OpenapiAccountsAddCondResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


