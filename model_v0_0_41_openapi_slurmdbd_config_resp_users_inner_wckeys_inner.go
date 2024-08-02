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

// checks if the V0041OpenapiSlurmdbdConfigRespUsersInnerWckeysInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0041OpenapiSlurmdbdConfigRespUsersInnerWckeysInner{}

// V0041OpenapiSlurmdbdConfigRespUsersInnerWckeysInner struct for V0041OpenapiSlurmdbdConfigRespUsersInnerWckeysInner
type V0041OpenapiSlurmdbdConfigRespUsersInnerWckeysInner struct {
	Accounting []V0041OpenapiSlurmdbdConfigRespUsersInnerWckeysInnerAccountingInner `json:"accounting,omitempty"`
	Cluster string `json:"cluster"`
	Id *int32 `json:"id,omitempty"`
	Name string `json:"name"`
	User string `json:"user"`
	Flags []string `json:"flags,omitempty"`
}

type _V0041OpenapiSlurmdbdConfigRespUsersInnerWckeysInner V0041OpenapiSlurmdbdConfigRespUsersInnerWckeysInner

// NewV0041OpenapiSlurmdbdConfigRespUsersInnerWckeysInner instantiates a new V0041OpenapiSlurmdbdConfigRespUsersInnerWckeysInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0041OpenapiSlurmdbdConfigRespUsersInnerWckeysInner(cluster string, name string, user string) *V0041OpenapiSlurmdbdConfigRespUsersInnerWckeysInner {
	this := V0041OpenapiSlurmdbdConfigRespUsersInnerWckeysInner{}
	this.Cluster = cluster
	this.Name = name
	this.User = user
	return &this
}

// NewV0041OpenapiSlurmdbdConfigRespUsersInnerWckeysInnerWithDefaults instantiates a new V0041OpenapiSlurmdbdConfigRespUsersInnerWckeysInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0041OpenapiSlurmdbdConfigRespUsersInnerWckeysInnerWithDefaults() *V0041OpenapiSlurmdbdConfigRespUsersInnerWckeysInner {
	this := V0041OpenapiSlurmdbdConfigRespUsersInnerWckeysInner{}
	return &this
}

// GetAccounting returns the Accounting field value if set, zero value otherwise.
func (o *V0041OpenapiSlurmdbdConfigRespUsersInnerWckeysInner) GetAccounting() []V0041OpenapiSlurmdbdConfigRespUsersInnerWckeysInnerAccountingInner {
	if o == nil || IsNil(o.Accounting) {
		var ret []V0041OpenapiSlurmdbdConfigRespUsersInnerWckeysInnerAccountingInner
		return ret
	}
	return o.Accounting
}

// GetAccountingOk returns a tuple with the Accounting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiSlurmdbdConfigRespUsersInnerWckeysInner) GetAccountingOk() ([]V0041OpenapiSlurmdbdConfigRespUsersInnerWckeysInnerAccountingInner, bool) {
	if o == nil || IsNil(o.Accounting) {
		return nil, false
	}
	return o.Accounting, true
}

// HasAccounting returns a boolean if a field has been set.
func (o *V0041OpenapiSlurmdbdConfigRespUsersInnerWckeysInner) HasAccounting() bool {
	if o != nil && !IsNil(o.Accounting) {
		return true
	}

	return false
}

// SetAccounting gets a reference to the given []V0041OpenapiSlurmdbdConfigRespUsersInnerWckeysInnerAccountingInner and assigns it to the Accounting field.
func (o *V0041OpenapiSlurmdbdConfigRespUsersInnerWckeysInner) SetAccounting(v []V0041OpenapiSlurmdbdConfigRespUsersInnerWckeysInnerAccountingInner) {
	o.Accounting = v
}

// GetCluster returns the Cluster field value
func (o *V0041OpenapiSlurmdbdConfigRespUsersInnerWckeysInner) GetCluster() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value
// and a boolean to check if the value has been set.
func (o *V0041OpenapiSlurmdbdConfigRespUsersInnerWckeysInner) GetClusterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cluster, true
}

// SetCluster sets field value
func (o *V0041OpenapiSlurmdbdConfigRespUsersInnerWckeysInner) SetCluster(v string) {
	o.Cluster = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *V0041OpenapiSlurmdbdConfigRespUsersInnerWckeysInner) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiSlurmdbdConfigRespUsersInnerWckeysInner) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *V0041OpenapiSlurmdbdConfigRespUsersInnerWckeysInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *V0041OpenapiSlurmdbdConfigRespUsersInnerWckeysInner) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *V0041OpenapiSlurmdbdConfigRespUsersInnerWckeysInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *V0041OpenapiSlurmdbdConfigRespUsersInnerWckeysInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *V0041OpenapiSlurmdbdConfigRespUsersInnerWckeysInner) SetName(v string) {
	o.Name = v
}

// GetUser returns the User field value
func (o *V0041OpenapiSlurmdbdConfigRespUsersInnerWckeysInner) GetUser() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *V0041OpenapiSlurmdbdConfigRespUsersInnerWckeysInner) GetUserOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *V0041OpenapiSlurmdbdConfigRespUsersInnerWckeysInner) SetUser(v string) {
	o.User = v
}

// GetFlags returns the Flags field value if set, zero value otherwise.
func (o *V0041OpenapiSlurmdbdConfigRespUsersInnerWckeysInner) GetFlags() []string {
	if o == nil || IsNil(o.Flags) {
		var ret []string
		return ret
	}
	return o.Flags
}

// GetFlagsOk returns a tuple with the Flags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiSlurmdbdConfigRespUsersInnerWckeysInner) GetFlagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Flags) {
		return nil, false
	}
	return o.Flags, true
}

// HasFlags returns a boolean if a field has been set.
func (o *V0041OpenapiSlurmdbdConfigRespUsersInnerWckeysInner) HasFlags() bool {
	if o != nil && !IsNil(o.Flags) {
		return true
	}

	return false
}

// SetFlags gets a reference to the given []string and assigns it to the Flags field.
func (o *V0041OpenapiSlurmdbdConfigRespUsersInnerWckeysInner) SetFlags(v []string) {
	o.Flags = v
}

func (o V0041OpenapiSlurmdbdConfigRespUsersInnerWckeysInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0041OpenapiSlurmdbdConfigRespUsersInnerWckeysInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Accounting) {
		toSerialize["accounting"] = o.Accounting
	}
	toSerialize["cluster"] = o.Cluster
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["name"] = o.Name
	toSerialize["user"] = o.User
	if !IsNil(o.Flags) {
		toSerialize["flags"] = o.Flags
	}
	return toSerialize, nil
}

func (o *V0041OpenapiSlurmdbdConfigRespUsersInnerWckeysInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cluster",
		"name",
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

	varV0041OpenapiSlurmdbdConfigRespUsersInnerWckeysInner := _V0041OpenapiSlurmdbdConfigRespUsersInnerWckeysInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV0041OpenapiSlurmdbdConfigRespUsersInnerWckeysInner)

	if err != nil {
		return err
	}

	*o = V0041OpenapiSlurmdbdConfigRespUsersInnerWckeysInner(varV0041OpenapiSlurmdbdConfigRespUsersInnerWckeysInner)

	return err
}

type NullableV0041OpenapiSlurmdbdConfigRespUsersInnerWckeysInner struct {
	value *V0041OpenapiSlurmdbdConfigRespUsersInnerWckeysInner
	isSet bool
}

func (v NullableV0041OpenapiSlurmdbdConfigRespUsersInnerWckeysInner) Get() *V0041OpenapiSlurmdbdConfigRespUsersInnerWckeysInner {
	return v.value
}

func (v *NullableV0041OpenapiSlurmdbdConfigRespUsersInnerWckeysInner) Set(val *V0041OpenapiSlurmdbdConfigRespUsersInnerWckeysInner) {
	v.value = val
	v.isSet = true
}

func (v NullableV0041OpenapiSlurmdbdConfigRespUsersInnerWckeysInner) IsSet() bool {
	return v.isSet
}

func (v *NullableV0041OpenapiSlurmdbdConfigRespUsersInnerWckeysInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0041OpenapiSlurmdbdConfigRespUsersInnerWckeysInner(val *V0041OpenapiSlurmdbdConfigRespUsersInnerWckeysInner) *NullableV0041OpenapiSlurmdbdConfigRespUsersInnerWckeysInner {
	return &NullableV0041OpenapiSlurmdbdConfigRespUsersInnerWckeysInner{value: val, isSet: true}
}

func (v NullableV0041OpenapiSlurmdbdConfigRespUsersInnerWckeysInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0041OpenapiSlurmdbdConfigRespUsersInnerWckeysInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

