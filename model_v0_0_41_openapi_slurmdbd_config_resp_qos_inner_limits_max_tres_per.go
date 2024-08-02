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

// checks if the V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresPer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresPer{}

// V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresPer struct for V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresPer
type V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresPer struct {
	Account []V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTresRequestedMaxInner `json:"account,omitempty"`
	Job []V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTresRequestedMaxInner `json:"job,omitempty"`
	Node []V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTresRequestedMaxInner `json:"node,omitempty"`
	User []V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTresRequestedMaxInner `json:"user,omitempty"`
}

// NewV0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresPer instantiates a new V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresPer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresPer() *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresPer {
	this := V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresPer{}
	return &this
}

// NewV0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresPerWithDefaults instantiates a new V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresPer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresPerWithDefaults() *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresPer {
	this := V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresPer{}
	return &this
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresPer) GetAccount() []V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTresRequestedMaxInner {
	if o == nil || IsNil(o.Account) {
		var ret []V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTresRequestedMaxInner
		return ret
	}
	return o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresPer) GetAccountOk() ([]V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTresRequestedMaxInner, bool) {
	if o == nil || IsNil(o.Account) {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresPer) HasAccount() bool {
	if o != nil && !IsNil(o.Account) {
		return true
	}

	return false
}

// SetAccount gets a reference to the given []V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTresRequestedMaxInner and assigns it to the Account field.
func (o *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresPer) SetAccount(v []V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTresRequestedMaxInner) {
	o.Account = v
}

// GetJob returns the Job field value if set, zero value otherwise.
func (o *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresPer) GetJob() []V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTresRequestedMaxInner {
	if o == nil || IsNil(o.Job) {
		var ret []V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTresRequestedMaxInner
		return ret
	}
	return o.Job
}

// GetJobOk returns a tuple with the Job field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresPer) GetJobOk() ([]V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTresRequestedMaxInner, bool) {
	if o == nil || IsNil(o.Job) {
		return nil, false
	}
	return o.Job, true
}

// HasJob returns a boolean if a field has been set.
func (o *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresPer) HasJob() bool {
	if o != nil && !IsNil(o.Job) {
		return true
	}

	return false
}

// SetJob gets a reference to the given []V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTresRequestedMaxInner and assigns it to the Job field.
func (o *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresPer) SetJob(v []V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTresRequestedMaxInner) {
	o.Job = v
}

// GetNode returns the Node field value if set, zero value otherwise.
func (o *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresPer) GetNode() []V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTresRequestedMaxInner {
	if o == nil || IsNil(o.Node) {
		var ret []V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTresRequestedMaxInner
		return ret
	}
	return o.Node
}

// GetNodeOk returns a tuple with the Node field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresPer) GetNodeOk() ([]V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTresRequestedMaxInner, bool) {
	if o == nil || IsNil(o.Node) {
		return nil, false
	}
	return o.Node, true
}

// HasNode returns a boolean if a field has been set.
func (o *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresPer) HasNode() bool {
	if o != nil && !IsNil(o.Node) {
		return true
	}

	return false
}

// SetNode gets a reference to the given []V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTresRequestedMaxInner and assigns it to the Node field.
func (o *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresPer) SetNode(v []V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTresRequestedMaxInner) {
	o.Node = v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresPer) GetUser() []V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTresRequestedMaxInner {
	if o == nil || IsNil(o.User) {
		var ret []V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTresRequestedMaxInner
		return ret
	}
	return o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresPer) GetUserOk() ([]V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTresRequestedMaxInner, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresPer) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given []V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTresRequestedMaxInner and assigns it to the User field.
func (o *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresPer) SetUser(v []V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTresRequestedMaxInner) {
	o.User = v
}

func (o V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresPer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresPer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Account) {
		toSerialize["account"] = o.Account
	}
	if !IsNil(o.Job) {
		toSerialize["job"] = o.Job
	}
	if !IsNil(o.Node) {
		toSerialize["node"] = o.Node
	}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	return toSerialize, nil
}

type NullableV0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresPer struct {
	value *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresPer
	isSet bool
}

func (v NullableV0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresPer) Get() *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresPer {
	return v.value
}

func (v *NullableV0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresPer) Set(val *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresPer) {
	v.value = val
	v.isSet = true
}

func (v NullableV0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresPer) IsSet() bool {
	return v.isSet
}

func (v *NullableV0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresPer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresPer(val *V0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresPer) *NullableV0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresPer {
	return &NullableV0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresPer{value: val, isSet: true}
}

func (v NullableV0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresPer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0041OpenapiSlurmdbdConfigRespQosInnerLimitsMaxTresPer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


