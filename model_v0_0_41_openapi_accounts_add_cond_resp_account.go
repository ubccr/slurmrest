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

// checks if the V0041OpenapiAccountsAddCondRespAccount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0041OpenapiAccountsAddCondRespAccount{}

// V0041OpenapiAccountsAddCondRespAccount Account organization and description
type V0041OpenapiAccountsAddCondRespAccount struct {
	// An arbitrary string describing an account
	Description *string `json:"description,omitempty"`
	// Organization to which the account belongs
	Organization *string `json:"organization,omitempty"`
}

// NewV0041OpenapiAccountsAddCondRespAccount instantiates a new V0041OpenapiAccountsAddCondRespAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0041OpenapiAccountsAddCondRespAccount() *V0041OpenapiAccountsAddCondRespAccount {
	this := V0041OpenapiAccountsAddCondRespAccount{}
	return &this
}

// NewV0041OpenapiAccountsAddCondRespAccountWithDefaults instantiates a new V0041OpenapiAccountsAddCondRespAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0041OpenapiAccountsAddCondRespAccountWithDefaults() *V0041OpenapiAccountsAddCondRespAccount {
	this := V0041OpenapiAccountsAddCondRespAccount{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *V0041OpenapiAccountsAddCondRespAccount) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiAccountsAddCondRespAccount) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *V0041OpenapiAccountsAddCondRespAccount) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *V0041OpenapiAccountsAddCondRespAccount) SetDescription(v string) {
	o.Description = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *V0041OpenapiAccountsAddCondRespAccount) GetOrganization() string {
	if o == nil || IsNil(o.Organization) {
		var ret string
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiAccountsAddCondRespAccount) GetOrganizationOk() (*string, bool) {
	if o == nil || IsNil(o.Organization) {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *V0041OpenapiAccountsAddCondRespAccount) HasOrganization() bool {
	if o != nil && !IsNil(o.Organization) {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given string and assigns it to the Organization field.
func (o *V0041OpenapiAccountsAddCondRespAccount) SetOrganization(v string) {
	o.Organization = &v
}

func (o V0041OpenapiAccountsAddCondRespAccount) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0041OpenapiAccountsAddCondRespAccount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Organization) {
		toSerialize["organization"] = o.Organization
	}
	return toSerialize, nil
}

type NullableV0041OpenapiAccountsAddCondRespAccount struct {
	value *V0041OpenapiAccountsAddCondRespAccount
	isSet bool
}

func (v NullableV0041OpenapiAccountsAddCondRespAccount) Get() *V0041OpenapiAccountsAddCondRespAccount {
	return v.value
}

func (v *NullableV0041OpenapiAccountsAddCondRespAccount) Set(val *V0041OpenapiAccountsAddCondRespAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableV0041OpenapiAccountsAddCondRespAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableV0041OpenapiAccountsAddCondRespAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0041OpenapiAccountsAddCondRespAccount(val *V0041OpenapiAccountsAddCondRespAccount) *NullableV0041OpenapiAccountsAddCondRespAccount {
	return &NullableV0041OpenapiAccountsAddCondRespAccount{value: val, isSet: true}
}

func (v NullableV0041OpenapiAccountsAddCondRespAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0041OpenapiAccountsAddCondRespAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

