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

// checks if the V0041OpenapiAccountsAddCondRespAssociationCondition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0041OpenapiAccountsAddCondRespAssociationCondition{}

// V0041OpenapiAccountsAddCondRespAssociationCondition CSV list of accounts, association limits and options, CSV list of clusters
type V0041OpenapiAccountsAddCondRespAssociationCondition struct {
	// CSV accounts list
	Accounts []string `json:"accounts"`
	Association *V0041OpenapiUsersAddCondRespAssociationConditionAssociation `json:"association,omitempty"`
	// CSV clusters list
	Clusters []string `json:"clusters,omitempty"`
}

type _V0041OpenapiAccountsAddCondRespAssociationCondition V0041OpenapiAccountsAddCondRespAssociationCondition

// NewV0041OpenapiAccountsAddCondRespAssociationCondition instantiates a new V0041OpenapiAccountsAddCondRespAssociationCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0041OpenapiAccountsAddCondRespAssociationCondition(accounts []string) *V0041OpenapiAccountsAddCondRespAssociationCondition {
	this := V0041OpenapiAccountsAddCondRespAssociationCondition{}
	this.Accounts = accounts
	return &this
}

// NewV0041OpenapiAccountsAddCondRespAssociationConditionWithDefaults instantiates a new V0041OpenapiAccountsAddCondRespAssociationCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0041OpenapiAccountsAddCondRespAssociationConditionWithDefaults() *V0041OpenapiAccountsAddCondRespAssociationCondition {
	this := V0041OpenapiAccountsAddCondRespAssociationCondition{}
	return &this
}

// GetAccounts returns the Accounts field value
func (o *V0041OpenapiAccountsAddCondRespAssociationCondition) GetAccounts() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Accounts
}

// GetAccountsOk returns a tuple with the Accounts field value
// and a boolean to check if the value has been set.
func (o *V0041OpenapiAccountsAddCondRespAssociationCondition) GetAccountsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Accounts, true
}

// SetAccounts sets field value
func (o *V0041OpenapiAccountsAddCondRespAssociationCondition) SetAccounts(v []string) {
	o.Accounts = v
}

// GetAssociation returns the Association field value if set, zero value otherwise.
func (o *V0041OpenapiAccountsAddCondRespAssociationCondition) GetAssociation() V0041OpenapiUsersAddCondRespAssociationConditionAssociation {
	if o == nil || IsNil(o.Association) {
		var ret V0041OpenapiUsersAddCondRespAssociationConditionAssociation
		return ret
	}
	return *o.Association
}

// GetAssociationOk returns a tuple with the Association field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiAccountsAddCondRespAssociationCondition) GetAssociationOk() (*V0041OpenapiUsersAddCondRespAssociationConditionAssociation, bool) {
	if o == nil || IsNil(o.Association) {
		return nil, false
	}
	return o.Association, true
}

// HasAssociation returns a boolean if a field has been set.
func (o *V0041OpenapiAccountsAddCondRespAssociationCondition) HasAssociation() bool {
	if o != nil && !IsNil(o.Association) {
		return true
	}

	return false
}

// SetAssociation gets a reference to the given V0041OpenapiUsersAddCondRespAssociationConditionAssociation and assigns it to the Association field.
func (o *V0041OpenapiAccountsAddCondRespAssociationCondition) SetAssociation(v V0041OpenapiUsersAddCondRespAssociationConditionAssociation) {
	o.Association = &v
}

// GetClusters returns the Clusters field value if set, zero value otherwise.
func (o *V0041OpenapiAccountsAddCondRespAssociationCondition) GetClusters() []string {
	if o == nil || IsNil(o.Clusters) {
		var ret []string
		return ret
	}
	return o.Clusters
}

// GetClustersOk returns a tuple with the Clusters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiAccountsAddCondRespAssociationCondition) GetClustersOk() ([]string, bool) {
	if o == nil || IsNil(o.Clusters) {
		return nil, false
	}
	return o.Clusters, true
}

// HasClusters returns a boolean if a field has been set.
func (o *V0041OpenapiAccountsAddCondRespAssociationCondition) HasClusters() bool {
	if o != nil && !IsNil(o.Clusters) {
		return true
	}

	return false
}

// SetClusters gets a reference to the given []string and assigns it to the Clusters field.
func (o *V0041OpenapiAccountsAddCondRespAssociationCondition) SetClusters(v []string) {
	o.Clusters = v
}

func (o V0041OpenapiAccountsAddCondRespAssociationCondition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0041OpenapiAccountsAddCondRespAssociationCondition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["accounts"] = o.Accounts
	if !IsNil(o.Association) {
		toSerialize["association"] = o.Association
	}
	if !IsNil(o.Clusters) {
		toSerialize["clusters"] = o.Clusters
	}
	return toSerialize, nil
}

func (o *V0041OpenapiAccountsAddCondRespAssociationCondition) UnmarshalJSON(data []byte) (err error) {
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

	varV0041OpenapiAccountsAddCondRespAssociationCondition := _V0041OpenapiAccountsAddCondRespAssociationCondition{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV0041OpenapiAccountsAddCondRespAssociationCondition)

	if err != nil {
		return err
	}

	*o = V0041OpenapiAccountsAddCondRespAssociationCondition(varV0041OpenapiAccountsAddCondRespAssociationCondition)

	return err
}

type NullableV0041OpenapiAccountsAddCondRespAssociationCondition struct {
	value *V0041OpenapiAccountsAddCondRespAssociationCondition
	isSet bool
}

func (v NullableV0041OpenapiAccountsAddCondRespAssociationCondition) Get() *V0041OpenapiAccountsAddCondRespAssociationCondition {
	return v.value
}

func (v *NullableV0041OpenapiAccountsAddCondRespAssociationCondition) Set(val *V0041OpenapiAccountsAddCondRespAssociationCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableV0041OpenapiAccountsAddCondRespAssociationCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableV0041OpenapiAccountsAddCondRespAssociationCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0041OpenapiAccountsAddCondRespAssociationCondition(val *V0041OpenapiAccountsAddCondRespAssociationCondition) *NullableV0041OpenapiAccountsAddCondRespAssociationCondition {
	return &NullableV0041OpenapiAccountsAddCondRespAssociationCondition{value: val, isSet: true}
}

func (v NullableV0041OpenapiAccountsAddCondRespAssociationCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0041OpenapiAccountsAddCondRespAssociationCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

