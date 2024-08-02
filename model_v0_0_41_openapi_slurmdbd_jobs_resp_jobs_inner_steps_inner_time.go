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

// checks if the V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTime type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTime{}

// V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTime struct for V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTime
type V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTime struct {
	Elapsed *int32 `json:"elapsed,omitempty"`
	End *V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu `json:"end,omitempty"`
	Start *V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu `json:"start,omitempty"`
	Suspended *int32 `json:"suspended,omitempty"`
	System *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTimeSystem `json:"system,omitempty"`
	Total *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTimeSystem `json:"total,omitempty"`
	User *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTimeSystem `json:"user,omitempty"`
}

// NewV0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTime instantiates a new V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTime object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTime() *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTime {
	this := V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTime{}
	return &this
}

// NewV0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTimeWithDefaults instantiates a new V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTime object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTimeWithDefaults() *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTime {
	this := V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTime{}
	return &this
}

// GetElapsed returns the Elapsed field value if set, zero value otherwise.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTime) GetElapsed() int32 {
	if o == nil || IsNil(o.Elapsed) {
		var ret int32
		return ret
	}
	return *o.Elapsed
}

// GetElapsedOk returns a tuple with the Elapsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTime) GetElapsedOk() (*int32, bool) {
	if o == nil || IsNil(o.Elapsed) {
		return nil, false
	}
	return o.Elapsed, true
}

// HasElapsed returns a boolean if a field has been set.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTime) HasElapsed() bool {
	if o != nil && !IsNil(o.Elapsed) {
		return true
	}

	return false
}

// SetElapsed gets a reference to the given int32 and assigns it to the Elapsed field.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTime) SetElapsed(v int32) {
	o.Elapsed = &v
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTime) GetEnd() V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu {
	if o == nil || IsNil(o.End) {
		var ret V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTime) GetEndOk() (*V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu, bool) {
	if o == nil || IsNil(o.End) {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTime) HasEnd() bool {
	if o != nil && !IsNil(o.End) {
		return true
	}

	return false
}

// SetEnd gets a reference to the given V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu and assigns it to the End field.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTime) SetEnd(v V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu) {
	o.End = &v
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTime) GetStart() V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu {
	if o == nil || IsNil(o.Start) {
		var ret V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTime) GetStartOk() (*V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu, bool) {
	if o == nil || IsNil(o.Start) {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTime) HasStart() bool {
	if o != nil && !IsNil(o.Start) {
		return true
	}

	return false
}

// SetStart gets a reference to the given V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu and assigns it to the Start field.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTime) SetStart(v V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu) {
	o.Start = &v
}

// GetSuspended returns the Suspended field value if set, zero value otherwise.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTime) GetSuspended() int32 {
	if o == nil || IsNil(o.Suspended) {
		var ret int32
		return ret
	}
	return *o.Suspended
}

// GetSuspendedOk returns a tuple with the Suspended field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTime) GetSuspendedOk() (*int32, bool) {
	if o == nil || IsNil(o.Suspended) {
		return nil, false
	}
	return o.Suspended, true
}

// HasSuspended returns a boolean if a field has been set.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTime) HasSuspended() bool {
	if o != nil && !IsNil(o.Suspended) {
		return true
	}

	return false
}

// SetSuspended gets a reference to the given int32 and assigns it to the Suspended field.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTime) SetSuspended(v int32) {
	o.Suspended = &v
}

// GetSystem returns the System field value if set, zero value otherwise.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTime) GetSystem() V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTimeSystem {
	if o == nil || IsNil(o.System) {
		var ret V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTimeSystem
		return ret
	}
	return *o.System
}

// GetSystemOk returns a tuple with the System field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTime) GetSystemOk() (*V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTimeSystem, bool) {
	if o == nil || IsNil(o.System) {
		return nil, false
	}
	return o.System, true
}

// HasSystem returns a boolean if a field has been set.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTime) HasSystem() bool {
	if o != nil && !IsNil(o.System) {
		return true
	}

	return false
}

// SetSystem gets a reference to the given V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTimeSystem and assigns it to the System field.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTime) SetSystem(v V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTimeSystem) {
	o.System = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTime) GetTotal() V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTimeSystem {
	if o == nil || IsNil(o.Total) {
		var ret V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTimeSystem
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTime) GetTotalOk() (*V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTimeSystem, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTime) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTimeSystem and assigns it to the Total field.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTime) SetTotal(v V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTimeSystem) {
	o.Total = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTime) GetUser() V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTimeSystem {
	if o == nil || IsNil(o.User) {
		var ret V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTimeSystem
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTime) GetUserOk() (*V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTimeSystem, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTime) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTimeSystem and assigns it to the User field.
func (o *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTime) SetUser(v V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTimeSystem) {
	o.User = &v
}

func (o V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTime) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTime) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Elapsed) {
		toSerialize["elapsed"] = o.Elapsed
	}
	if !IsNil(o.End) {
		toSerialize["end"] = o.End
	}
	if !IsNil(o.Start) {
		toSerialize["start"] = o.Start
	}
	if !IsNil(o.Suspended) {
		toSerialize["suspended"] = o.Suspended
	}
	if !IsNil(o.System) {
		toSerialize["system"] = o.System
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	return toSerialize, nil
}

type NullableV0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTime struct {
	value *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTime
	isSet bool
}

func (v NullableV0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTime) Get() *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTime {
	return v.value
}

func (v *NullableV0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTime) Set(val *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTime) {
	v.value = val
	v.isSet = true
}

func (v NullableV0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTime) IsSet() bool {
	return v.isSet
}

func (v *NullableV0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTime) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTime(val *V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTime) *NullableV0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTime {
	return &NullableV0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTime{value: val, isSet: true}
}

func (v NullableV0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTime) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

