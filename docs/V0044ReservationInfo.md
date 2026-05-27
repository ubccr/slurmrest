# V0044ReservationInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accounts** | Pointer to **string** | Comma-separated list of permitted accounts | [optional] 
**BurstBuffer** | Pointer to **string** | BurstBuffer - Burst buffer resources reserved | [optional] 
**CoreCount** | Pointer to **int32** | CoreCnt - Number of cores reserved | [optional] 
**CoreSpecializations** | Pointer to [**[]V0044ReservationCoreSpec**](V0044ReservationCoreSpec.md) |  | [optional] 
**EndTime** | Pointer to [**V0044Uint64NoValStruct**](V0044Uint64NoValStruct.md) |  | [optional] 
**Features** | Pointer to **string** | Features - Expression describing the reservation&#39;s required node features | [optional] 
**Flags** | Pointer to **[]string** | Flags associated with this reservation | [optional] 
**Groups** | Pointer to **string** | Groups - Comma-separated list of permitted groups | [optional] 
**Licenses** | Pointer to **string** | Licenses - Comma-separated list of licenses reserved | [optional] 
**MaxStartDelay** | Pointer to **int32** | MaxStartDelay - Maximum time an eligible job not requesting this reservation can delay a job requesting it in seconds | [optional] 
**Name** | Pointer to **string** | ReservationName - Name of the reservation | [optional] 
**NodeCount** | Pointer to **int32** | NodeCnt - Number of nodes reserved | [optional] 
**NodeList** | Pointer to **string** | Nodes - Comma-separated list of node names and/or node ranges reserved | [optional] 
**Partition** | Pointer to **string** | PartitionName - Partition used to reserve nodes from | [optional] 
**PurgeCompleted** | Pointer to [**V0044ReservationInfoPurgeCompleted**](V0044ReservationInfoPurgeCompleted.md) |  | [optional] 
**StartTime** | Pointer to [**V0044Uint64NoValStruct**](V0044Uint64NoValStruct.md) |  | [optional] 
**Watts** | Pointer to [**V0044Uint32NoValStruct**](V0044Uint32NoValStruct.md) |  | [optional] 
**Tres** | Pointer to **string** | Comma-separated list of required TRES | [optional] 
**Users** | Pointer to **string** | Comma-separated list of permitted users | [optional] 

## Methods

### NewV0044ReservationInfo

`func NewV0044ReservationInfo() *V0044ReservationInfo`

NewV0044ReservationInfo instantiates a new V0044ReservationInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0044ReservationInfoWithDefaults

`func NewV0044ReservationInfoWithDefaults() *V0044ReservationInfo`

NewV0044ReservationInfoWithDefaults instantiates a new V0044ReservationInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccounts

`func (o *V0044ReservationInfo) GetAccounts() string`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *V0044ReservationInfo) GetAccountsOk() (*string, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *V0044ReservationInfo) SetAccounts(v string)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *V0044ReservationInfo) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetBurstBuffer

`func (o *V0044ReservationInfo) GetBurstBuffer() string`

GetBurstBuffer returns the BurstBuffer field if non-nil, zero value otherwise.

### GetBurstBufferOk

`func (o *V0044ReservationInfo) GetBurstBufferOk() (*string, bool)`

GetBurstBufferOk returns a tuple with the BurstBuffer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurstBuffer

`func (o *V0044ReservationInfo) SetBurstBuffer(v string)`

SetBurstBuffer sets BurstBuffer field to given value.

### HasBurstBuffer

`func (o *V0044ReservationInfo) HasBurstBuffer() bool`

HasBurstBuffer returns a boolean if a field has been set.

### GetCoreCount

`func (o *V0044ReservationInfo) GetCoreCount() int32`

GetCoreCount returns the CoreCount field if non-nil, zero value otherwise.

### GetCoreCountOk

`func (o *V0044ReservationInfo) GetCoreCountOk() (*int32, bool)`

GetCoreCountOk returns a tuple with the CoreCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoreCount

`func (o *V0044ReservationInfo) SetCoreCount(v int32)`

SetCoreCount sets CoreCount field to given value.

### HasCoreCount

`func (o *V0044ReservationInfo) HasCoreCount() bool`

HasCoreCount returns a boolean if a field has been set.

### GetCoreSpecializations

`func (o *V0044ReservationInfo) GetCoreSpecializations() []V0044ReservationCoreSpec`

GetCoreSpecializations returns the CoreSpecializations field if non-nil, zero value otherwise.

### GetCoreSpecializationsOk

`func (o *V0044ReservationInfo) GetCoreSpecializationsOk() (*[]V0044ReservationCoreSpec, bool)`

GetCoreSpecializationsOk returns a tuple with the CoreSpecializations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoreSpecializations

`func (o *V0044ReservationInfo) SetCoreSpecializations(v []V0044ReservationCoreSpec)`

SetCoreSpecializations sets CoreSpecializations field to given value.

### HasCoreSpecializations

`func (o *V0044ReservationInfo) HasCoreSpecializations() bool`

HasCoreSpecializations returns a boolean if a field has been set.

### GetEndTime

`func (o *V0044ReservationInfo) GetEndTime() V0044Uint64NoValStruct`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *V0044ReservationInfo) GetEndTimeOk() (*V0044Uint64NoValStruct, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *V0044ReservationInfo) SetEndTime(v V0044Uint64NoValStruct)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *V0044ReservationInfo) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetFeatures

`func (o *V0044ReservationInfo) GetFeatures() string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *V0044ReservationInfo) GetFeaturesOk() (*string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *V0044ReservationInfo) SetFeatures(v string)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *V0044ReservationInfo) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetFlags

`func (o *V0044ReservationInfo) GetFlags() []string`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *V0044ReservationInfo) GetFlagsOk() (*[]string, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *V0044ReservationInfo) SetFlags(v []string)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *V0044ReservationInfo) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetGroups

`func (o *V0044ReservationInfo) GetGroups() string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *V0044ReservationInfo) GetGroupsOk() (*string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *V0044ReservationInfo) SetGroups(v string)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *V0044ReservationInfo) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetLicenses

`func (o *V0044ReservationInfo) GetLicenses() string`

GetLicenses returns the Licenses field if non-nil, zero value otherwise.

### GetLicensesOk

`func (o *V0044ReservationInfo) GetLicensesOk() (*string, bool)`

GetLicensesOk returns a tuple with the Licenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenses

`func (o *V0044ReservationInfo) SetLicenses(v string)`

SetLicenses sets Licenses field to given value.

### HasLicenses

`func (o *V0044ReservationInfo) HasLicenses() bool`

HasLicenses returns a boolean if a field has been set.

### GetMaxStartDelay

`func (o *V0044ReservationInfo) GetMaxStartDelay() int32`

GetMaxStartDelay returns the MaxStartDelay field if non-nil, zero value otherwise.

### GetMaxStartDelayOk

`func (o *V0044ReservationInfo) GetMaxStartDelayOk() (*int32, bool)`

GetMaxStartDelayOk returns a tuple with the MaxStartDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStartDelay

`func (o *V0044ReservationInfo) SetMaxStartDelay(v int32)`

SetMaxStartDelay sets MaxStartDelay field to given value.

### HasMaxStartDelay

`func (o *V0044ReservationInfo) HasMaxStartDelay() bool`

HasMaxStartDelay returns a boolean if a field has been set.

### GetName

`func (o *V0044ReservationInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V0044ReservationInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V0044ReservationInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V0044ReservationInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNodeCount

`func (o *V0044ReservationInfo) GetNodeCount() int32`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *V0044ReservationInfo) GetNodeCountOk() (*int32, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *V0044ReservationInfo) SetNodeCount(v int32)`

SetNodeCount sets NodeCount field to given value.

### HasNodeCount

`func (o *V0044ReservationInfo) HasNodeCount() bool`

HasNodeCount returns a boolean if a field has been set.

### GetNodeList

`func (o *V0044ReservationInfo) GetNodeList() string`

GetNodeList returns the NodeList field if non-nil, zero value otherwise.

### GetNodeListOk

`func (o *V0044ReservationInfo) GetNodeListOk() (*string, bool)`

GetNodeListOk returns a tuple with the NodeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeList

`func (o *V0044ReservationInfo) SetNodeList(v string)`

SetNodeList sets NodeList field to given value.

### HasNodeList

`func (o *V0044ReservationInfo) HasNodeList() bool`

HasNodeList returns a boolean if a field has been set.

### GetPartition

`func (o *V0044ReservationInfo) GetPartition() string`

GetPartition returns the Partition field if non-nil, zero value otherwise.

### GetPartitionOk

`func (o *V0044ReservationInfo) GetPartitionOk() (*string, bool)`

GetPartitionOk returns a tuple with the Partition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartition

`func (o *V0044ReservationInfo) SetPartition(v string)`

SetPartition sets Partition field to given value.

### HasPartition

`func (o *V0044ReservationInfo) HasPartition() bool`

HasPartition returns a boolean if a field has been set.

### GetPurgeCompleted

`func (o *V0044ReservationInfo) GetPurgeCompleted() V0044ReservationInfoPurgeCompleted`

GetPurgeCompleted returns the PurgeCompleted field if non-nil, zero value otherwise.

### GetPurgeCompletedOk

`func (o *V0044ReservationInfo) GetPurgeCompletedOk() (*V0044ReservationInfoPurgeCompleted, bool)`

GetPurgeCompletedOk returns a tuple with the PurgeCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurgeCompleted

`func (o *V0044ReservationInfo) SetPurgeCompleted(v V0044ReservationInfoPurgeCompleted)`

SetPurgeCompleted sets PurgeCompleted field to given value.

### HasPurgeCompleted

`func (o *V0044ReservationInfo) HasPurgeCompleted() bool`

HasPurgeCompleted returns a boolean if a field has been set.

### GetStartTime

`func (o *V0044ReservationInfo) GetStartTime() V0044Uint64NoValStruct`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *V0044ReservationInfo) GetStartTimeOk() (*V0044Uint64NoValStruct, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *V0044ReservationInfo) SetStartTime(v V0044Uint64NoValStruct)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *V0044ReservationInfo) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetWatts

`func (o *V0044ReservationInfo) GetWatts() V0044Uint32NoValStruct`

GetWatts returns the Watts field if non-nil, zero value otherwise.

### GetWattsOk

`func (o *V0044ReservationInfo) GetWattsOk() (*V0044Uint32NoValStruct, bool)`

GetWattsOk returns a tuple with the Watts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatts

`func (o *V0044ReservationInfo) SetWatts(v V0044Uint32NoValStruct)`

SetWatts sets Watts field to given value.

### HasWatts

`func (o *V0044ReservationInfo) HasWatts() bool`

HasWatts returns a boolean if a field has been set.

### GetTres

`func (o *V0044ReservationInfo) GetTres() string`

GetTres returns the Tres field if non-nil, zero value otherwise.

### GetTresOk

`func (o *V0044ReservationInfo) GetTresOk() (*string, bool)`

GetTresOk returns a tuple with the Tres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTres

`func (o *V0044ReservationInfo) SetTres(v string)`

SetTres sets Tres field to given value.

### HasTres

`func (o *V0044ReservationInfo) HasTres() bool`

HasTres returns a boolean if a field has been set.

### GetUsers

`func (o *V0044ReservationInfo) GetUsers() string`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *V0044ReservationInfo) GetUsersOk() (*string, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *V0044ReservationInfo) SetUsers(v string)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *V0044ReservationInfo) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


