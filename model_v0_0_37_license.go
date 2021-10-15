/*
 * Slurm Rest API
 *
 * API to access and control Slurm.
 *
 * API version: 0.0.37
 * Contact: sales@schedmd.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package slurmrest

import (
	"encoding/json"
)

// V0037License struct for V0037License
type V0037License struct {
	// name of license
	Name *string `json:"name,omitempty"`
	// total number of licenses
	Total *int32 `json:"total,omitempty"`
	// number of licenses in use
	InUse *int32 `json:"in_use,omitempty"`
	// number of licenses available
	Available *int32 `json:"available,omitempty"`
	// number of licenses reserved
	Reserved *int32 `json:"reserved,omitempty"`
	// license is remote
	Remote *bool `json:"remote,omitempty"`
}

// NewV0037License instantiates a new V0037License object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0037License() *V0037License {
	this := V0037License{}
	return &this
}

// NewV0037LicenseWithDefaults instantiates a new V0037License object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0037LicenseWithDefaults() *V0037License {
	this := V0037License{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *V0037License) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0037License) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *V0037License) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *V0037License) SetName(v string) {
	o.Name = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *V0037License) GetTotal() int32 {
	if o == nil || o.Total == nil {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0037License) GetTotalOk() (*int32, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *V0037License) HasTotal() bool {
	if o != nil && o.Total != nil {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *V0037License) SetTotal(v int32) {
	o.Total = &v
}

// GetInUse returns the InUse field value if set, zero value otherwise.
func (o *V0037License) GetInUse() int32 {
	if o == nil || o.InUse == nil {
		var ret int32
		return ret
	}
	return *o.InUse
}

// GetInUseOk returns a tuple with the InUse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0037License) GetInUseOk() (*int32, bool) {
	if o == nil || o.InUse == nil {
		return nil, false
	}
	return o.InUse, true
}

// HasInUse returns a boolean if a field has been set.
func (o *V0037License) HasInUse() bool {
	if o != nil && o.InUse != nil {
		return true
	}

	return false
}

// SetInUse gets a reference to the given int32 and assigns it to the InUse field.
func (o *V0037License) SetInUse(v int32) {
	o.InUse = &v
}

// GetAvailable returns the Available field value if set, zero value otherwise.
func (o *V0037License) GetAvailable() int32 {
	if o == nil || o.Available == nil {
		var ret int32
		return ret
	}
	return *o.Available
}

// GetAvailableOk returns a tuple with the Available field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0037License) GetAvailableOk() (*int32, bool) {
	if o == nil || o.Available == nil {
		return nil, false
	}
	return o.Available, true
}

// HasAvailable returns a boolean if a field has been set.
func (o *V0037License) HasAvailable() bool {
	if o != nil && o.Available != nil {
		return true
	}

	return false
}

// SetAvailable gets a reference to the given int32 and assigns it to the Available field.
func (o *V0037License) SetAvailable(v int32) {
	o.Available = &v
}

// GetReserved returns the Reserved field value if set, zero value otherwise.
func (o *V0037License) GetReserved() int32 {
	if o == nil || o.Reserved == nil {
		var ret int32
		return ret
	}
	return *o.Reserved
}

// GetReservedOk returns a tuple with the Reserved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0037License) GetReservedOk() (*int32, bool) {
	if o == nil || o.Reserved == nil {
		return nil, false
	}
	return o.Reserved, true
}

// HasReserved returns a boolean if a field has been set.
func (o *V0037License) HasReserved() bool {
	if o != nil && o.Reserved != nil {
		return true
	}

	return false
}

// SetReserved gets a reference to the given int32 and assigns it to the Reserved field.
func (o *V0037License) SetReserved(v int32) {
	o.Reserved = &v
}

// GetRemote returns the Remote field value if set, zero value otherwise.
func (o *V0037License) GetRemote() bool {
	if o == nil || o.Remote == nil {
		var ret bool
		return ret
	}
	return *o.Remote
}

// GetRemoteOk returns a tuple with the Remote field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0037License) GetRemoteOk() (*bool, bool) {
	if o == nil || o.Remote == nil {
		return nil, false
	}
	return o.Remote, true
}

// HasRemote returns a boolean if a field has been set.
func (o *V0037License) HasRemote() bool {
	if o != nil && o.Remote != nil {
		return true
	}

	return false
}

// SetRemote gets a reference to the given bool and assigns it to the Remote field.
func (o *V0037License) SetRemote(v bool) {
	o.Remote = &v
}

func (o V0037License) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Total != nil {
		toSerialize["total"] = o.Total
	}
	if o.InUse != nil {
		toSerialize["in_use"] = o.InUse
	}
	if o.Available != nil {
		toSerialize["available"] = o.Available
	}
	if o.Reserved != nil {
		toSerialize["reserved"] = o.Reserved
	}
	if o.Remote != nil {
		toSerialize["remote"] = o.Remote
	}
	return json.Marshal(toSerialize)
}

type NullableV0037License struct {
	value *V0037License
	isSet bool
}

func (v NullableV0037License) Get() *V0037License {
	return v.value
}

func (v *NullableV0037License) Set(val *V0037License) {
	v.value = val
	v.isSet = true
}

func (v NullableV0037License) IsSet() bool {
	return v.isSet
}

func (v *NullableV0037License) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0037License(val *V0037License) *NullableV0037License {
	return &NullableV0037License{value: val, isSet: true}
}

func (v NullableV0037License) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0037License) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


