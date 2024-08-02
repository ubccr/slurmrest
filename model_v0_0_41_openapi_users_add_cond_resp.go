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

// checks if the V0041OpenapiUsersAddCondResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0041OpenapiUsersAddCondResp{}

// V0041OpenapiUsersAddCondResp struct for V0041OpenapiUsersAddCondResp
type V0041OpenapiUsersAddCondResp struct {
	AssociationCondition V0041OpenapiUsersAddCondRespAssociationCondition `json:"association_condition"`
	User V0041OpenapiUsersAddCondRespUser `json:"user"`
	Meta *V0041OpenapiSlurmdbdJobsRespMeta `json:"meta,omitempty"`
	// Query errors
	Errors []V0041OpenapiSlurmdbdJobsRespErrorsInner `json:"errors,omitempty"`
	// Query warnings
	Warnings []V0041OpenapiSlurmdbdJobsRespWarningsInner `json:"warnings,omitempty"`
}

type _V0041OpenapiUsersAddCondResp V0041OpenapiUsersAddCondResp

// NewV0041OpenapiUsersAddCondResp instantiates a new V0041OpenapiUsersAddCondResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0041OpenapiUsersAddCondResp(associationCondition V0041OpenapiUsersAddCondRespAssociationCondition, user V0041OpenapiUsersAddCondRespUser) *V0041OpenapiUsersAddCondResp {
	this := V0041OpenapiUsersAddCondResp{}
	this.AssociationCondition = associationCondition
	this.User = user
	return &this
}

// NewV0041OpenapiUsersAddCondRespWithDefaults instantiates a new V0041OpenapiUsersAddCondResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0041OpenapiUsersAddCondRespWithDefaults() *V0041OpenapiUsersAddCondResp {
	this := V0041OpenapiUsersAddCondResp{}
	return &this
}

// GetAssociationCondition returns the AssociationCondition field value
func (o *V0041OpenapiUsersAddCondResp) GetAssociationCondition() V0041OpenapiUsersAddCondRespAssociationCondition {
	if o == nil {
		var ret V0041OpenapiUsersAddCondRespAssociationCondition
		return ret
	}

	return o.AssociationCondition
}

// GetAssociationConditionOk returns a tuple with the AssociationCondition field value
// and a boolean to check if the value has been set.
func (o *V0041OpenapiUsersAddCondResp) GetAssociationConditionOk() (*V0041OpenapiUsersAddCondRespAssociationCondition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssociationCondition, true
}

// SetAssociationCondition sets field value
func (o *V0041OpenapiUsersAddCondResp) SetAssociationCondition(v V0041OpenapiUsersAddCondRespAssociationCondition) {
	o.AssociationCondition = v
}

// GetUser returns the User field value
func (o *V0041OpenapiUsersAddCondResp) GetUser() V0041OpenapiUsersAddCondRespUser {
	if o == nil {
		var ret V0041OpenapiUsersAddCondRespUser
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *V0041OpenapiUsersAddCondResp) GetUserOk() (*V0041OpenapiUsersAddCondRespUser, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *V0041OpenapiUsersAddCondResp) SetUser(v V0041OpenapiUsersAddCondRespUser) {
	o.User = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *V0041OpenapiUsersAddCondResp) GetMeta() V0041OpenapiSlurmdbdJobsRespMeta {
	if o == nil || IsNil(o.Meta) {
		var ret V0041OpenapiSlurmdbdJobsRespMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiUsersAddCondResp) GetMetaOk() (*V0041OpenapiSlurmdbdJobsRespMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *V0041OpenapiUsersAddCondResp) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given V0041OpenapiSlurmdbdJobsRespMeta and assigns it to the Meta field.
func (o *V0041OpenapiUsersAddCondResp) SetMeta(v V0041OpenapiSlurmdbdJobsRespMeta) {
	o.Meta = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *V0041OpenapiUsersAddCondResp) GetErrors() []V0041OpenapiSlurmdbdJobsRespErrorsInner {
	if o == nil || IsNil(o.Errors) {
		var ret []V0041OpenapiSlurmdbdJobsRespErrorsInner
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiUsersAddCondResp) GetErrorsOk() ([]V0041OpenapiSlurmdbdJobsRespErrorsInner, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *V0041OpenapiUsersAddCondResp) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []V0041OpenapiSlurmdbdJobsRespErrorsInner and assigns it to the Errors field.
func (o *V0041OpenapiUsersAddCondResp) SetErrors(v []V0041OpenapiSlurmdbdJobsRespErrorsInner) {
	o.Errors = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *V0041OpenapiUsersAddCondResp) GetWarnings() []V0041OpenapiSlurmdbdJobsRespWarningsInner {
	if o == nil || IsNil(o.Warnings) {
		var ret []V0041OpenapiSlurmdbdJobsRespWarningsInner
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiUsersAddCondResp) GetWarningsOk() ([]V0041OpenapiSlurmdbdJobsRespWarningsInner, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *V0041OpenapiUsersAddCondResp) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []V0041OpenapiSlurmdbdJobsRespWarningsInner and assigns it to the Warnings field.
func (o *V0041OpenapiUsersAddCondResp) SetWarnings(v []V0041OpenapiSlurmdbdJobsRespWarningsInner) {
	o.Warnings = v
}

func (o V0041OpenapiUsersAddCondResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0041OpenapiUsersAddCondResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["association_condition"] = o.AssociationCondition
	toSerialize["user"] = o.User
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

func (o *V0041OpenapiUsersAddCondResp) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"association_condition",
		"user",
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

	varV0041OpenapiUsersAddCondResp := _V0041OpenapiUsersAddCondResp{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV0041OpenapiUsersAddCondResp)

	if err != nil {
		return err
	}

	*o = V0041OpenapiUsersAddCondResp(varV0041OpenapiUsersAddCondResp)

	return err
}

type NullableV0041OpenapiUsersAddCondResp struct {
	value *V0041OpenapiUsersAddCondResp
	isSet bool
}

func (v NullableV0041OpenapiUsersAddCondResp) Get() *V0041OpenapiUsersAddCondResp {
	return v.value
}

func (v *NullableV0041OpenapiUsersAddCondResp) Set(val *V0041OpenapiUsersAddCondResp) {
	v.value = val
	v.isSet = true
}

func (v NullableV0041OpenapiUsersAddCondResp) IsSet() bool {
	return v.isSet
}

func (v *NullableV0041OpenapiUsersAddCondResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0041OpenapiUsersAddCondResp(val *V0041OpenapiUsersAddCondResp) *NullableV0041OpenapiUsersAddCondResp {
	return &NullableV0041OpenapiUsersAddCondResp{value: val, isSet: true}
}

func (v NullableV0041OpenapiUsersAddCondResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0041OpenapiUsersAddCondResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

