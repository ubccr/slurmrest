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

// checks if the V0039ClusterRec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0039ClusterRec{}

// V0039ClusterRec struct for V0039ClusterRec
type V0039ClusterRec struct {
	Controller *V0041OpenapiSlurmdbdConfigRespClustersInnerController `json:"controller,omitempty"`
	Flags []string `json:"flags,omitempty"`
	Name *string `json:"name,omitempty"`
	Nodes *string `json:"nodes,omitempty"`
	SelectPlugin *string `json:"select_plugin,omitempty"`
	Associations *V0039ClusterRecAssociations `json:"associations,omitempty"`
	RpcVersion *int32 `json:"rpc_version,omitempty"`
	Tres []V0039Tres `json:"tres,omitempty"`
}

// NewV0039ClusterRec instantiates a new V0039ClusterRec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0039ClusterRec() *V0039ClusterRec {
	this := V0039ClusterRec{}
	return &this
}

// NewV0039ClusterRecWithDefaults instantiates a new V0039ClusterRec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0039ClusterRecWithDefaults() *V0039ClusterRec {
	this := V0039ClusterRec{}
	return &this
}

// GetController returns the Controller field value if set, zero value otherwise.
func (o *V0039ClusterRec) GetController() V0041OpenapiSlurmdbdConfigRespClustersInnerController {
	if o == nil || IsNil(o.Controller) {
		var ret V0041OpenapiSlurmdbdConfigRespClustersInnerController
		return ret
	}
	return *o.Controller
}

// GetControllerOk returns a tuple with the Controller field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039ClusterRec) GetControllerOk() (*V0041OpenapiSlurmdbdConfigRespClustersInnerController, bool) {
	if o == nil || IsNil(o.Controller) {
		return nil, false
	}
	return o.Controller, true
}

// HasController returns a boolean if a field has been set.
func (o *V0039ClusterRec) HasController() bool {
	if o != nil && !IsNil(o.Controller) {
		return true
	}

	return false
}

// SetController gets a reference to the given V0041OpenapiSlurmdbdConfigRespClustersInnerController and assigns it to the Controller field.
func (o *V0039ClusterRec) SetController(v V0041OpenapiSlurmdbdConfigRespClustersInnerController) {
	o.Controller = &v
}

// GetFlags returns the Flags field value if set, zero value otherwise.
func (o *V0039ClusterRec) GetFlags() []string {
	if o == nil || IsNil(o.Flags) {
		var ret []string
		return ret
	}
	return o.Flags
}

// GetFlagsOk returns a tuple with the Flags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039ClusterRec) GetFlagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Flags) {
		return nil, false
	}
	return o.Flags, true
}

// HasFlags returns a boolean if a field has been set.
func (o *V0039ClusterRec) HasFlags() bool {
	if o != nil && !IsNil(o.Flags) {
		return true
	}

	return false
}

// SetFlags gets a reference to the given []string and assigns it to the Flags field.
func (o *V0039ClusterRec) SetFlags(v []string) {
	o.Flags = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *V0039ClusterRec) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039ClusterRec) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *V0039ClusterRec) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *V0039ClusterRec) SetName(v string) {
	o.Name = &v
}

// GetNodes returns the Nodes field value if set, zero value otherwise.
func (o *V0039ClusterRec) GetNodes() string {
	if o == nil || IsNil(o.Nodes) {
		var ret string
		return ret
	}
	return *o.Nodes
}

// GetNodesOk returns a tuple with the Nodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039ClusterRec) GetNodesOk() (*string, bool) {
	if o == nil || IsNil(o.Nodes) {
		return nil, false
	}
	return o.Nodes, true
}

// HasNodes returns a boolean if a field has been set.
func (o *V0039ClusterRec) HasNodes() bool {
	if o != nil && !IsNil(o.Nodes) {
		return true
	}

	return false
}

// SetNodes gets a reference to the given string and assigns it to the Nodes field.
func (o *V0039ClusterRec) SetNodes(v string) {
	o.Nodes = &v
}

// GetSelectPlugin returns the SelectPlugin field value if set, zero value otherwise.
func (o *V0039ClusterRec) GetSelectPlugin() string {
	if o == nil || IsNil(o.SelectPlugin) {
		var ret string
		return ret
	}
	return *o.SelectPlugin
}

// GetSelectPluginOk returns a tuple with the SelectPlugin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039ClusterRec) GetSelectPluginOk() (*string, bool) {
	if o == nil || IsNil(o.SelectPlugin) {
		return nil, false
	}
	return o.SelectPlugin, true
}

// HasSelectPlugin returns a boolean if a field has been set.
func (o *V0039ClusterRec) HasSelectPlugin() bool {
	if o != nil && !IsNil(o.SelectPlugin) {
		return true
	}

	return false
}

// SetSelectPlugin gets a reference to the given string and assigns it to the SelectPlugin field.
func (o *V0039ClusterRec) SetSelectPlugin(v string) {
	o.SelectPlugin = &v
}

// GetAssociations returns the Associations field value if set, zero value otherwise.
func (o *V0039ClusterRec) GetAssociations() V0039ClusterRecAssociations {
	if o == nil || IsNil(o.Associations) {
		var ret V0039ClusterRecAssociations
		return ret
	}
	return *o.Associations
}

// GetAssociationsOk returns a tuple with the Associations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039ClusterRec) GetAssociationsOk() (*V0039ClusterRecAssociations, bool) {
	if o == nil || IsNil(o.Associations) {
		return nil, false
	}
	return o.Associations, true
}

// HasAssociations returns a boolean if a field has been set.
func (o *V0039ClusterRec) HasAssociations() bool {
	if o != nil && !IsNil(o.Associations) {
		return true
	}

	return false
}

// SetAssociations gets a reference to the given V0039ClusterRecAssociations and assigns it to the Associations field.
func (o *V0039ClusterRec) SetAssociations(v V0039ClusterRecAssociations) {
	o.Associations = &v
}

// GetRpcVersion returns the RpcVersion field value if set, zero value otherwise.
func (o *V0039ClusterRec) GetRpcVersion() int32 {
	if o == nil || IsNil(o.RpcVersion) {
		var ret int32
		return ret
	}
	return *o.RpcVersion
}

// GetRpcVersionOk returns a tuple with the RpcVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039ClusterRec) GetRpcVersionOk() (*int32, bool) {
	if o == nil || IsNil(o.RpcVersion) {
		return nil, false
	}
	return o.RpcVersion, true
}

// HasRpcVersion returns a boolean if a field has been set.
func (o *V0039ClusterRec) HasRpcVersion() bool {
	if o != nil && !IsNil(o.RpcVersion) {
		return true
	}

	return false
}

// SetRpcVersion gets a reference to the given int32 and assigns it to the RpcVersion field.
func (o *V0039ClusterRec) SetRpcVersion(v int32) {
	o.RpcVersion = &v
}

// GetTres returns the Tres field value if set, zero value otherwise.
func (o *V0039ClusterRec) GetTres() []V0039Tres {
	if o == nil || IsNil(o.Tres) {
		var ret []V0039Tres
		return ret
	}
	return o.Tres
}

// GetTresOk returns a tuple with the Tres field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039ClusterRec) GetTresOk() ([]V0039Tres, bool) {
	if o == nil || IsNil(o.Tres) {
		return nil, false
	}
	return o.Tres, true
}

// HasTres returns a boolean if a field has been set.
func (o *V0039ClusterRec) HasTres() bool {
	if o != nil && !IsNil(o.Tres) {
		return true
	}

	return false
}

// SetTres gets a reference to the given []V0039Tres and assigns it to the Tres field.
func (o *V0039ClusterRec) SetTres(v []V0039Tres) {
	o.Tres = v
}

func (o V0039ClusterRec) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0039ClusterRec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Controller) {
		toSerialize["controller"] = o.Controller
	}
	if !IsNil(o.Flags) {
		toSerialize["flags"] = o.Flags
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Nodes) {
		toSerialize["nodes"] = o.Nodes
	}
	if !IsNil(o.SelectPlugin) {
		toSerialize["select_plugin"] = o.SelectPlugin
	}
	if !IsNil(o.Associations) {
		toSerialize["associations"] = o.Associations
	}
	if !IsNil(o.RpcVersion) {
		toSerialize["rpc_version"] = o.RpcVersion
	}
	if !IsNil(o.Tres) {
		toSerialize["tres"] = o.Tres
	}
	return toSerialize, nil
}

type NullableV0039ClusterRec struct {
	value *V0039ClusterRec
	isSet bool
}

func (v NullableV0039ClusterRec) Get() *V0039ClusterRec {
	return v.value
}

func (v *NullableV0039ClusterRec) Set(val *V0039ClusterRec) {
	v.value = val
	v.isSet = true
}

func (v NullableV0039ClusterRec) IsSet() bool {
	return v.isSet
}

func (v *NullableV0039ClusterRec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0039ClusterRec(val *V0039ClusterRec) *NullableV0039ClusterRec {
	return &NullableV0039ClusterRec{value: val, isSet: true}
}

func (v NullableV0039ClusterRec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0039ClusterRec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


