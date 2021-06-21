/*
 * Slurm Rest API
 *
 * API to access and control Slurm.
 *
 * API version: 0.0.36
 * Contact: sales@schedmd.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package slurmrest

import (
	"encoding/json"
	"fmt"
)

// V0036Signal POSIX signal name
type V0036Signal string

// List of v0.0.36_signal
const (
	HUP V0036Signal = "HUP"
	INT V0036Signal = "INT"
	QUIT V0036Signal = "QUIT"
	ABRT V0036Signal = "ABRT"
	KILL V0036Signal = "KILL"
	ALRM V0036Signal = "ALRM"
	TERM V0036Signal = "TERM"
	USR1 V0036Signal = "USR1"
	USR2 V0036Signal = "USR2"
	URG V0036Signal = "URG"
	CONT V0036Signal = "CONT"
	STOP V0036Signal = "STOP"
	TSTP V0036Signal = "TSTP"
	TTIN V0036Signal = "TTIN"
	TTOU V0036Signal = "TTOU"
)

func (v *V0036Signal) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := V0036Signal(value)
	for _, existing := range []V0036Signal{ "HUP", "INT", "QUIT", "ABRT", "KILL", "ALRM", "TERM", "USR1", "USR2", "URG", "CONT", "STOP", "TSTP", "TTIN", "TTOU",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid V0036Signal", value)
}

// Ptr returns reference to v0.0.36_signal value
func (v V0036Signal) Ptr() *V0036Signal {
	return &v
}

type NullableV0036Signal struct {
	value *V0036Signal
	isSet bool
}

func (v NullableV0036Signal) Get() *V0036Signal {
	return v.value
}

func (v *NullableV0036Signal) Set(val *V0036Signal) {
	v.value = val
	v.isSet = true
}

func (v NullableV0036Signal) IsSet() bool {
	return v.isSet
}

func (v *NullableV0036Signal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0036Signal(val *V0036Signal) *NullableV0036Signal {
	return &NullableV0036Signal{value: val, isSet: true}
}

func (v NullableV0036Signal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0036Signal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
