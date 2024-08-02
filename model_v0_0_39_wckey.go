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

// checks if the V0039Wckey type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0039Wckey{}

// V0039Wckey struct for V0039Wckey
type V0039Wckey struct {
	Accounting []V0039Accounting `json:"accounting,omitempty"`
	Cluster string `json:"cluster"`
	Id *int32 `json:"id,omitempty"`
	Name string `json:"name"`
	User string `json:"user"`
	Flags []string `json:"flags,omitempty"`
}

type _V0039Wckey V0039Wckey

// NewV0039Wckey instantiates a new V0039Wckey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0039Wckey(cluster string, name string, user string) *V0039Wckey {
	this := V0039Wckey{}
	this.Cluster = cluster
	this.Name = name
	this.User = user
	return &this
}

// NewV0039WckeyWithDefaults instantiates a new V0039Wckey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0039WckeyWithDefaults() *V0039Wckey {
	this := V0039Wckey{}
	return &this
}

// GetAccounting returns the Accounting field value if set, zero value otherwise.
func (o *V0039Wckey) GetAccounting() []V0039Accounting {
	if o == nil || IsNil(o.Accounting) {
		var ret []V0039Accounting
		return ret
	}
	return o.Accounting
}

// GetAccountingOk returns a tuple with the Accounting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039Wckey) GetAccountingOk() ([]V0039Accounting, bool) {
	if o == nil || IsNil(o.Accounting) {
		return nil, false
	}
	return o.Accounting, true
}

// HasAccounting returns a boolean if a field has been set.
func (o *V0039Wckey) HasAccounting() bool {
	if o != nil && !IsNil(o.Accounting) {
		return true
	}

	return false
}

// SetAccounting gets a reference to the given []V0039Accounting and assigns it to the Accounting field.
func (o *V0039Wckey) SetAccounting(v []V0039Accounting) {
	o.Accounting = v
}

// GetCluster returns the Cluster field value
func (o *V0039Wckey) GetCluster() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value
// and a boolean to check if the value has been set.
func (o *V0039Wckey) GetClusterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cluster, true
}

// SetCluster sets field value
func (o *V0039Wckey) SetCluster(v string) {
	o.Cluster = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *V0039Wckey) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039Wckey) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *V0039Wckey) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *V0039Wckey) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *V0039Wckey) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *V0039Wckey) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *V0039Wckey) SetName(v string) {
	o.Name = v
}

// GetUser returns the User field value
func (o *V0039Wckey) GetUser() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *V0039Wckey) GetUserOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *V0039Wckey) SetUser(v string) {
	o.User = v
}

// GetFlags returns the Flags field value if set, zero value otherwise.
func (o *V0039Wckey) GetFlags() []string {
	if o == nil || IsNil(o.Flags) {
		var ret []string
		return ret
	}
	return o.Flags
}

// GetFlagsOk returns a tuple with the Flags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039Wckey) GetFlagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Flags) {
		return nil, false
	}
	return o.Flags, true
}

// HasFlags returns a boolean if a field has been set.
func (o *V0039Wckey) HasFlags() bool {
	if o != nil && !IsNil(o.Flags) {
		return true
	}

	return false
}

// SetFlags gets a reference to the given []string and assigns it to the Flags field.
func (o *V0039Wckey) SetFlags(v []string) {
	o.Flags = v
}

func (o V0039Wckey) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0039Wckey) ToMap() (map[string]interface{}, error) {
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

func (o *V0039Wckey) UnmarshalJSON(data []byte) (err error) {
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

	varV0039Wckey := _V0039Wckey{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV0039Wckey)

	if err != nil {
		return err
	}

	*o = V0039Wckey(varV0039Wckey)

	return err
}

type NullableV0039Wckey struct {
	value *V0039Wckey
	isSet bool
}

func (v NullableV0039Wckey) Get() *V0039Wckey {
	return v.value
}

func (v *NullableV0039Wckey) Set(val *V0039Wckey) {
	v.value = val
	v.isSet = true
}

func (v NullableV0039Wckey) IsSet() bool {
	return v.isSet
}

func (v *NullableV0039Wckey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0039Wckey(val *V0039Wckey) *NullableV0039Wckey {
	return &NullableV0039Wckey{value: val, isSet: true}
}

func (v NullableV0039Wckey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0039Wckey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


