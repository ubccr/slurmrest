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

// V0037Reservation struct for V0037Reservation
type V0037Reservation struct {
	// Allowed accounts
	Accounts *string `json:"accounts,omitempty"`
	// Reserved burst buffer
	BurstBuffer *string `json:"burst_buffer,omitempty"`
	// Number of reserved cores
	CoreCount *int32 `json:"core_count,omitempty"`
	// Number of reserved specialized cores
	CoreSpecCnt *int32 `json:"core_spec_cnt,omitempty"`
	// End time of the reservation
	EndTime *int32 `json:"end_time,omitempty"`
	// List of features
	Features *string `json:"features,omitempty"`
	// Reservation options
	Flags *[]string `json:"flags,omitempty"`
	// List of groups permitted to use the reserved nodes
	Groups *string `json:"groups,omitempty"`
	// List of licenses
	Licenses *string `json:"licenses,omitempty"`
	// Maximum delay in which jobs outside of the reservation will be permitted to overlap once any jobs are queued for the reservation
	MaxStartDelay *int32 `json:"max_start_delay,omitempty"`
	// Reservationn name
	Name *string `json:"name,omitempty"`
	// Count of nodes reserved
	NodeCount *int32 `json:"node_count,omitempty"`
	// List of reserved nodes
	NodeList *string `json:"node_list,omitempty"`
	// Partition
	Partition *string `json:"partition,omitempty"`
	PurgeCompleted *V0037ReservationPurgeCompleted `json:"purge_completed,omitempty"`
	// Start time of reservation
	StartTime *int32 `json:"start_time,omitempty"`
	// amount of power to reserve in watts
	Watts *int32 `json:"watts,omitempty"`
	// List of TRES
	Tres *string `json:"tres,omitempty"`
	// List of users
	Users *string `json:"users,omitempty"`
}

// NewV0037Reservation instantiates a new V0037Reservation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0037Reservation() *V0037Reservation {
	this := V0037Reservation{}
	return &this
}

// NewV0037ReservationWithDefaults instantiates a new V0037Reservation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0037ReservationWithDefaults() *V0037Reservation {
	this := V0037Reservation{}
	return &this
}

// GetAccounts returns the Accounts field value if set, zero value otherwise.
func (o *V0037Reservation) GetAccounts() string {
	if o == nil || o.Accounts == nil {
		var ret string
		return ret
	}
	return *o.Accounts
}

// GetAccountsOk returns a tuple with the Accounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0037Reservation) GetAccountsOk() (*string, bool) {
	if o == nil || o.Accounts == nil {
		return nil, false
	}
	return o.Accounts, true
}

// HasAccounts returns a boolean if a field has been set.
func (o *V0037Reservation) HasAccounts() bool {
	if o != nil && o.Accounts != nil {
		return true
	}

	return false
}

// SetAccounts gets a reference to the given string and assigns it to the Accounts field.
func (o *V0037Reservation) SetAccounts(v string) {
	o.Accounts = &v
}

// GetBurstBuffer returns the BurstBuffer field value if set, zero value otherwise.
func (o *V0037Reservation) GetBurstBuffer() string {
	if o == nil || o.BurstBuffer == nil {
		var ret string
		return ret
	}
	return *o.BurstBuffer
}

// GetBurstBufferOk returns a tuple with the BurstBuffer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0037Reservation) GetBurstBufferOk() (*string, bool) {
	if o == nil || o.BurstBuffer == nil {
		return nil, false
	}
	return o.BurstBuffer, true
}

// HasBurstBuffer returns a boolean if a field has been set.
func (o *V0037Reservation) HasBurstBuffer() bool {
	if o != nil && o.BurstBuffer != nil {
		return true
	}

	return false
}

// SetBurstBuffer gets a reference to the given string and assigns it to the BurstBuffer field.
func (o *V0037Reservation) SetBurstBuffer(v string) {
	o.BurstBuffer = &v
}

// GetCoreCount returns the CoreCount field value if set, zero value otherwise.
func (o *V0037Reservation) GetCoreCount() int32 {
	if o == nil || o.CoreCount == nil {
		var ret int32
		return ret
	}
	return *o.CoreCount
}

// GetCoreCountOk returns a tuple with the CoreCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0037Reservation) GetCoreCountOk() (*int32, bool) {
	if o == nil || o.CoreCount == nil {
		return nil, false
	}
	return o.CoreCount, true
}

// HasCoreCount returns a boolean if a field has been set.
func (o *V0037Reservation) HasCoreCount() bool {
	if o != nil && o.CoreCount != nil {
		return true
	}

	return false
}

// SetCoreCount gets a reference to the given int32 and assigns it to the CoreCount field.
func (o *V0037Reservation) SetCoreCount(v int32) {
	o.CoreCount = &v
}

// GetCoreSpecCnt returns the CoreSpecCnt field value if set, zero value otherwise.
func (o *V0037Reservation) GetCoreSpecCnt() int32 {
	if o == nil || o.CoreSpecCnt == nil {
		var ret int32
		return ret
	}
	return *o.CoreSpecCnt
}

// GetCoreSpecCntOk returns a tuple with the CoreSpecCnt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0037Reservation) GetCoreSpecCntOk() (*int32, bool) {
	if o == nil || o.CoreSpecCnt == nil {
		return nil, false
	}
	return o.CoreSpecCnt, true
}

// HasCoreSpecCnt returns a boolean if a field has been set.
func (o *V0037Reservation) HasCoreSpecCnt() bool {
	if o != nil && o.CoreSpecCnt != nil {
		return true
	}

	return false
}

// SetCoreSpecCnt gets a reference to the given int32 and assigns it to the CoreSpecCnt field.
func (o *V0037Reservation) SetCoreSpecCnt(v int32) {
	o.CoreSpecCnt = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *V0037Reservation) GetEndTime() int32 {
	if o == nil || o.EndTime == nil {
		var ret int32
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0037Reservation) GetEndTimeOk() (*int32, bool) {
	if o == nil || o.EndTime == nil {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *V0037Reservation) HasEndTime() bool {
	if o != nil && o.EndTime != nil {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given int32 and assigns it to the EndTime field.
func (o *V0037Reservation) SetEndTime(v int32) {
	o.EndTime = &v
}

// GetFeatures returns the Features field value if set, zero value otherwise.
func (o *V0037Reservation) GetFeatures() string {
	if o == nil || o.Features == nil {
		var ret string
		return ret
	}
	return *o.Features
}

// GetFeaturesOk returns a tuple with the Features field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0037Reservation) GetFeaturesOk() (*string, bool) {
	if o == nil || o.Features == nil {
		return nil, false
	}
	return o.Features, true
}

// HasFeatures returns a boolean if a field has been set.
func (o *V0037Reservation) HasFeatures() bool {
	if o != nil && o.Features != nil {
		return true
	}

	return false
}

// SetFeatures gets a reference to the given string and assigns it to the Features field.
func (o *V0037Reservation) SetFeatures(v string) {
	o.Features = &v
}

// GetFlags returns the Flags field value if set, zero value otherwise.
func (o *V0037Reservation) GetFlags() []string {
	if o == nil || o.Flags == nil {
		var ret []string
		return ret
	}
	return *o.Flags
}

// GetFlagsOk returns a tuple with the Flags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0037Reservation) GetFlagsOk() (*[]string, bool) {
	if o == nil || o.Flags == nil {
		return nil, false
	}
	return o.Flags, true
}

// HasFlags returns a boolean if a field has been set.
func (o *V0037Reservation) HasFlags() bool {
	if o != nil && o.Flags != nil {
		return true
	}

	return false
}

// SetFlags gets a reference to the given []string and assigns it to the Flags field.
func (o *V0037Reservation) SetFlags(v []string) {
	o.Flags = &v
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *V0037Reservation) GetGroups() string {
	if o == nil || o.Groups == nil {
		var ret string
		return ret
	}
	return *o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0037Reservation) GetGroupsOk() (*string, bool) {
	if o == nil || o.Groups == nil {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *V0037Reservation) HasGroups() bool {
	if o != nil && o.Groups != nil {
		return true
	}

	return false
}

// SetGroups gets a reference to the given string and assigns it to the Groups field.
func (o *V0037Reservation) SetGroups(v string) {
	o.Groups = &v
}

// GetLicenses returns the Licenses field value if set, zero value otherwise.
func (o *V0037Reservation) GetLicenses() string {
	if o == nil || o.Licenses == nil {
		var ret string
		return ret
	}
	return *o.Licenses
}

// GetLicensesOk returns a tuple with the Licenses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0037Reservation) GetLicensesOk() (*string, bool) {
	if o == nil || o.Licenses == nil {
		return nil, false
	}
	return o.Licenses, true
}

// HasLicenses returns a boolean if a field has been set.
func (o *V0037Reservation) HasLicenses() bool {
	if o != nil && o.Licenses != nil {
		return true
	}

	return false
}

// SetLicenses gets a reference to the given string and assigns it to the Licenses field.
func (o *V0037Reservation) SetLicenses(v string) {
	o.Licenses = &v
}

// GetMaxStartDelay returns the MaxStartDelay field value if set, zero value otherwise.
func (o *V0037Reservation) GetMaxStartDelay() int32 {
	if o == nil || o.MaxStartDelay == nil {
		var ret int32
		return ret
	}
	return *o.MaxStartDelay
}

// GetMaxStartDelayOk returns a tuple with the MaxStartDelay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0037Reservation) GetMaxStartDelayOk() (*int32, bool) {
	if o == nil || o.MaxStartDelay == nil {
		return nil, false
	}
	return o.MaxStartDelay, true
}

// HasMaxStartDelay returns a boolean if a field has been set.
func (o *V0037Reservation) HasMaxStartDelay() bool {
	if o != nil && o.MaxStartDelay != nil {
		return true
	}

	return false
}

// SetMaxStartDelay gets a reference to the given int32 and assigns it to the MaxStartDelay field.
func (o *V0037Reservation) SetMaxStartDelay(v int32) {
	o.MaxStartDelay = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *V0037Reservation) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0037Reservation) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *V0037Reservation) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *V0037Reservation) SetName(v string) {
	o.Name = &v
}

// GetNodeCount returns the NodeCount field value if set, zero value otherwise.
func (o *V0037Reservation) GetNodeCount() int32 {
	if o == nil || o.NodeCount == nil {
		var ret int32
		return ret
	}
	return *o.NodeCount
}

// GetNodeCountOk returns a tuple with the NodeCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0037Reservation) GetNodeCountOk() (*int32, bool) {
	if o == nil || o.NodeCount == nil {
		return nil, false
	}
	return o.NodeCount, true
}

// HasNodeCount returns a boolean if a field has been set.
func (o *V0037Reservation) HasNodeCount() bool {
	if o != nil && o.NodeCount != nil {
		return true
	}

	return false
}

// SetNodeCount gets a reference to the given int32 and assigns it to the NodeCount field.
func (o *V0037Reservation) SetNodeCount(v int32) {
	o.NodeCount = &v
}

// GetNodeList returns the NodeList field value if set, zero value otherwise.
func (o *V0037Reservation) GetNodeList() string {
	if o == nil || o.NodeList == nil {
		var ret string
		return ret
	}
	return *o.NodeList
}

// GetNodeListOk returns a tuple with the NodeList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0037Reservation) GetNodeListOk() (*string, bool) {
	if o == nil || o.NodeList == nil {
		return nil, false
	}
	return o.NodeList, true
}

// HasNodeList returns a boolean if a field has been set.
func (o *V0037Reservation) HasNodeList() bool {
	if o != nil && o.NodeList != nil {
		return true
	}

	return false
}

// SetNodeList gets a reference to the given string and assigns it to the NodeList field.
func (o *V0037Reservation) SetNodeList(v string) {
	o.NodeList = &v
}

// GetPartition returns the Partition field value if set, zero value otherwise.
func (o *V0037Reservation) GetPartition() string {
	if o == nil || o.Partition == nil {
		var ret string
		return ret
	}
	return *o.Partition
}

// GetPartitionOk returns a tuple with the Partition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0037Reservation) GetPartitionOk() (*string, bool) {
	if o == nil || o.Partition == nil {
		return nil, false
	}
	return o.Partition, true
}

// HasPartition returns a boolean if a field has been set.
func (o *V0037Reservation) HasPartition() bool {
	if o != nil && o.Partition != nil {
		return true
	}

	return false
}

// SetPartition gets a reference to the given string and assigns it to the Partition field.
func (o *V0037Reservation) SetPartition(v string) {
	o.Partition = &v
}

// GetPurgeCompleted returns the PurgeCompleted field value if set, zero value otherwise.
func (o *V0037Reservation) GetPurgeCompleted() V0037ReservationPurgeCompleted {
	if o == nil || o.PurgeCompleted == nil {
		var ret V0037ReservationPurgeCompleted
		return ret
	}
	return *o.PurgeCompleted
}

// GetPurgeCompletedOk returns a tuple with the PurgeCompleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0037Reservation) GetPurgeCompletedOk() (*V0037ReservationPurgeCompleted, bool) {
	if o == nil || o.PurgeCompleted == nil {
		return nil, false
	}
	return o.PurgeCompleted, true
}

// HasPurgeCompleted returns a boolean if a field has been set.
func (o *V0037Reservation) HasPurgeCompleted() bool {
	if o != nil && o.PurgeCompleted != nil {
		return true
	}

	return false
}

// SetPurgeCompleted gets a reference to the given V0037ReservationPurgeCompleted and assigns it to the PurgeCompleted field.
func (o *V0037Reservation) SetPurgeCompleted(v V0037ReservationPurgeCompleted) {
	o.PurgeCompleted = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *V0037Reservation) GetStartTime() int32 {
	if o == nil || o.StartTime == nil {
		var ret int32
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0037Reservation) GetStartTimeOk() (*int32, bool) {
	if o == nil || o.StartTime == nil {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *V0037Reservation) HasStartTime() bool {
	if o != nil && o.StartTime != nil {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given int32 and assigns it to the StartTime field.
func (o *V0037Reservation) SetStartTime(v int32) {
	o.StartTime = &v
}

// GetWatts returns the Watts field value if set, zero value otherwise.
func (o *V0037Reservation) GetWatts() int32 {
	if o == nil || o.Watts == nil {
		var ret int32
		return ret
	}
	return *o.Watts
}

// GetWattsOk returns a tuple with the Watts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0037Reservation) GetWattsOk() (*int32, bool) {
	if o == nil || o.Watts == nil {
		return nil, false
	}
	return o.Watts, true
}

// HasWatts returns a boolean if a field has been set.
func (o *V0037Reservation) HasWatts() bool {
	if o != nil && o.Watts != nil {
		return true
	}

	return false
}

// SetWatts gets a reference to the given int32 and assigns it to the Watts field.
func (o *V0037Reservation) SetWatts(v int32) {
	o.Watts = &v
}

// GetTres returns the Tres field value if set, zero value otherwise.
func (o *V0037Reservation) GetTres() string {
	if o == nil || o.Tres == nil {
		var ret string
		return ret
	}
	return *o.Tres
}

// GetTresOk returns a tuple with the Tres field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0037Reservation) GetTresOk() (*string, bool) {
	if o == nil || o.Tres == nil {
		return nil, false
	}
	return o.Tres, true
}

// HasTres returns a boolean if a field has been set.
func (o *V0037Reservation) HasTres() bool {
	if o != nil && o.Tres != nil {
		return true
	}

	return false
}

// SetTres gets a reference to the given string and assigns it to the Tres field.
func (o *V0037Reservation) SetTres(v string) {
	o.Tres = &v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *V0037Reservation) GetUsers() string {
	if o == nil || o.Users == nil {
		var ret string
		return ret
	}
	return *o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0037Reservation) GetUsersOk() (*string, bool) {
	if o == nil || o.Users == nil {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *V0037Reservation) HasUsers() bool {
	if o != nil && o.Users != nil {
		return true
	}

	return false
}

// SetUsers gets a reference to the given string and assigns it to the Users field.
func (o *V0037Reservation) SetUsers(v string) {
	o.Users = &v
}

func (o V0037Reservation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Accounts != nil {
		toSerialize["accounts"] = o.Accounts
	}
	if o.BurstBuffer != nil {
		toSerialize["burst_buffer"] = o.BurstBuffer
	}
	if o.CoreCount != nil {
		toSerialize["core_count"] = o.CoreCount
	}
	if o.CoreSpecCnt != nil {
		toSerialize["core_spec_cnt"] = o.CoreSpecCnt
	}
	if o.EndTime != nil {
		toSerialize["end_time"] = o.EndTime
	}
	if o.Features != nil {
		toSerialize["features"] = o.Features
	}
	if o.Flags != nil {
		toSerialize["flags"] = o.Flags
	}
	if o.Groups != nil {
		toSerialize["groups"] = o.Groups
	}
	if o.Licenses != nil {
		toSerialize["licenses"] = o.Licenses
	}
	if o.MaxStartDelay != nil {
		toSerialize["max_start_delay"] = o.MaxStartDelay
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.NodeCount != nil {
		toSerialize["node_count"] = o.NodeCount
	}
	if o.NodeList != nil {
		toSerialize["node_list"] = o.NodeList
	}
	if o.Partition != nil {
		toSerialize["partition"] = o.Partition
	}
	if o.PurgeCompleted != nil {
		toSerialize["purge_completed"] = o.PurgeCompleted
	}
	if o.StartTime != nil {
		toSerialize["start_time"] = o.StartTime
	}
	if o.Watts != nil {
		toSerialize["watts"] = o.Watts
	}
	if o.Tres != nil {
		toSerialize["tres"] = o.Tres
	}
	if o.Users != nil {
		toSerialize["users"] = o.Users
	}
	return json.Marshal(toSerialize)
}

type NullableV0037Reservation struct {
	value *V0037Reservation
	isSet bool
}

func (v NullableV0037Reservation) Get() *V0037Reservation {
	return v.value
}

func (v *NullableV0037Reservation) Set(val *V0037Reservation) {
	v.value = val
	v.isSet = true
}

func (v NullableV0037Reservation) IsSet() bool {
	return v.isSet
}

func (v *NullableV0037Reservation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0037Reservation(val *V0037Reservation) *NullableV0037Reservation {
	return &NullableV0037Reservation{value: val, isSet: true}
}

func (v NullableV0037Reservation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0037Reservation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


