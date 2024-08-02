# V0041OpenapiReservationRespReservationsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accounts** | Pointer to **string** |  | [optional] 
**BurstBuffer** | Pointer to **string** |  | [optional] 
**CoreCount** | Pointer to **int32** |  | [optional] 
**CoreSpecializations** | Pointer to [**[]V0041OpenapiReservationRespReservationsInnerCoreSpecializationsInner**](V0041OpenapiReservationRespReservationsInnerCoreSpecializationsInner.md) |  | [optional] 
**EndTime** | Pointer to [**V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu**](V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu.md) |  | [optional] 
**Features** | Pointer to **string** |  | [optional] 
**Flags** | Pointer to **[]string** |  | [optional] 
**Groups** | Pointer to **string** |  | [optional] 
**Licenses** | Pointer to **string** |  | [optional] 
**MaxStartDelay** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NodeCount** | Pointer to **int32** |  | [optional] 
**NodeList** | Pointer to **string** |  | [optional] 
**Partition** | Pointer to **string** |  | [optional] 
**PurgeCompleted** | Pointer to [**V0041OpenapiReservationRespReservationsInnerPurgeCompleted**](V0041OpenapiReservationRespReservationsInnerPurgeCompleted.md) |  | [optional] 
**StartTime** | Pointer to [**V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu**](V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu.md) |  | [optional] 
**Watts** | Pointer to [**V0041OpenapiReservationRespReservationsInnerWatts**](V0041OpenapiReservationRespReservationsInnerWatts.md) |  | [optional] 
**Tres** | Pointer to **string** |  | [optional] 
**Users** | Pointer to **string** |  | [optional] 

## Methods

### NewV0041OpenapiReservationRespReservationsInner

`func NewV0041OpenapiReservationRespReservationsInner() *V0041OpenapiReservationRespReservationsInner`

NewV0041OpenapiReservationRespReservationsInner instantiates a new V0041OpenapiReservationRespReservationsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0041OpenapiReservationRespReservationsInnerWithDefaults

`func NewV0041OpenapiReservationRespReservationsInnerWithDefaults() *V0041OpenapiReservationRespReservationsInner`

NewV0041OpenapiReservationRespReservationsInnerWithDefaults instantiates a new V0041OpenapiReservationRespReservationsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccounts

`func (o *V0041OpenapiReservationRespReservationsInner) GetAccounts() string`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *V0041OpenapiReservationRespReservationsInner) GetAccountsOk() (*string, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *V0041OpenapiReservationRespReservationsInner) SetAccounts(v string)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *V0041OpenapiReservationRespReservationsInner) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetBurstBuffer

`func (o *V0041OpenapiReservationRespReservationsInner) GetBurstBuffer() string`

GetBurstBuffer returns the BurstBuffer field if non-nil, zero value otherwise.

### GetBurstBufferOk

`func (o *V0041OpenapiReservationRespReservationsInner) GetBurstBufferOk() (*string, bool)`

GetBurstBufferOk returns a tuple with the BurstBuffer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurstBuffer

`func (o *V0041OpenapiReservationRespReservationsInner) SetBurstBuffer(v string)`

SetBurstBuffer sets BurstBuffer field to given value.

### HasBurstBuffer

`func (o *V0041OpenapiReservationRespReservationsInner) HasBurstBuffer() bool`

HasBurstBuffer returns a boolean if a field has been set.

### GetCoreCount

`func (o *V0041OpenapiReservationRespReservationsInner) GetCoreCount() int32`

GetCoreCount returns the CoreCount field if non-nil, zero value otherwise.

### GetCoreCountOk

`func (o *V0041OpenapiReservationRespReservationsInner) GetCoreCountOk() (*int32, bool)`

GetCoreCountOk returns a tuple with the CoreCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoreCount

`func (o *V0041OpenapiReservationRespReservationsInner) SetCoreCount(v int32)`

SetCoreCount sets CoreCount field to given value.

### HasCoreCount

`func (o *V0041OpenapiReservationRespReservationsInner) HasCoreCount() bool`

HasCoreCount returns a boolean if a field has been set.

### GetCoreSpecializations

`func (o *V0041OpenapiReservationRespReservationsInner) GetCoreSpecializations() []V0041OpenapiReservationRespReservationsInnerCoreSpecializationsInner`

GetCoreSpecializations returns the CoreSpecializations field if non-nil, zero value otherwise.

### GetCoreSpecializationsOk

`func (o *V0041OpenapiReservationRespReservationsInner) GetCoreSpecializationsOk() (*[]V0041OpenapiReservationRespReservationsInnerCoreSpecializationsInner, bool)`

GetCoreSpecializationsOk returns a tuple with the CoreSpecializations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoreSpecializations

`func (o *V0041OpenapiReservationRespReservationsInner) SetCoreSpecializations(v []V0041OpenapiReservationRespReservationsInnerCoreSpecializationsInner)`

SetCoreSpecializations sets CoreSpecializations field to given value.

### HasCoreSpecializations

`func (o *V0041OpenapiReservationRespReservationsInner) HasCoreSpecializations() bool`

HasCoreSpecializations returns a boolean if a field has been set.

### GetEndTime

`func (o *V0041OpenapiReservationRespReservationsInner) GetEndTime() V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *V0041OpenapiReservationRespReservationsInner) GetEndTimeOk() (*V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *V0041OpenapiReservationRespReservationsInner) SetEndTime(v V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *V0041OpenapiReservationRespReservationsInner) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetFeatures

`func (o *V0041OpenapiReservationRespReservationsInner) GetFeatures() string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *V0041OpenapiReservationRespReservationsInner) GetFeaturesOk() (*string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *V0041OpenapiReservationRespReservationsInner) SetFeatures(v string)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *V0041OpenapiReservationRespReservationsInner) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetFlags

`func (o *V0041OpenapiReservationRespReservationsInner) GetFlags() []string`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *V0041OpenapiReservationRespReservationsInner) GetFlagsOk() (*[]string, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *V0041OpenapiReservationRespReservationsInner) SetFlags(v []string)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *V0041OpenapiReservationRespReservationsInner) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetGroups

`func (o *V0041OpenapiReservationRespReservationsInner) GetGroups() string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *V0041OpenapiReservationRespReservationsInner) GetGroupsOk() (*string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *V0041OpenapiReservationRespReservationsInner) SetGroups(v string)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *V0041OpenapiReservationRespReservationsInner) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetLicenses

`func (o *V0041OpenapiReservationRespReservationsInner) GetLicenses() string`

GetLicenses returns the Licenses field if non-nil, zero value otherwise.

### GetLicensesOk

`func (o *V0041OpenapiReservationRespReservationsInner) GetLicensesOk() (*string, bool)`

GetLicensesOk returns a tuple with the Licenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenses

`func (o *V0041OpenapiReservationRespReservationsInner) SetLicenses(v string)`

SetLicenses sets Licenses field to given value.

### HasLicenses

`func (o *V0041OpenapiReservationRespReservationsInner) HasLicenses() bool`

HasLicenses returns a boolean if a field has been set.

### GetMaxStartDelay

`func (o *V0041OpenapiReservationRespReservationsInner) GetMaxStartDelay() int32`

GetMaxStartDelay returns the MaxStartDelay field if non-nil, zero value otherwise.

### GetMaxStartDelayOk

`func (o *V0041OpenapiReservationRespReservationsInner) GetMaxStartDelayOk() (*int32, bool)`

GetMaxStartDelayOk returns a tuple with the MaxStartDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStartDelay

`func (o *V0041OpenapiReservationRespReservationsInner) SetMaxStartDelay(v int32)`

SetMaxStartDelay sets MaxStartDelay field to given value.

### HasMaxStartDelay

`func (o *V0041OpenapiReservationRespReservationsInner) HasMaxStartDelay() bool`

HasMaxStartDelay returns a boolean if a field has been set.

### GetName

`func (o *V0041OpenapiReservationRespReservationsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V0041OpenapiReservationRespReservationsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V0041OpenapiReservationRespReservationsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V0041OpenapiReservationRespReservationsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNodeCount

`func (o *V0041OpenapiReservationRespReservationsInner) GetNodeCount() int32`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *V0041OpenapiReservationRespReservationsInner) GetNodeCountOk() (*int32, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *V0041OpenapiReservationRespReservationsInner) SetNodeCount(v int32)`

SetNodeCount sets NodeCount field to given value.

### HasNodeCount

`func (o *V0041OpenapiReservationRespReservationsInner) HasNodeCount() bool`

HasNodeCount returns a boolean if a field has been set.

### GetNodeList

`func (o *V0041OpenapiReservationRespReservationsInner) GetNodeList() string`

GetNodeList returns the NodeList field if non-nil, zero value otherwise.

### GetNodeListOk

`func (o *V0041OpenapiReservationRespReservationsInner) GetNodeListOk() (*string, bool)`

GetNodeListOk returns a tuple with the NodeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeList

`func (o *V0041OpenapiReservationRespReservationsInner) SetNodeList(v string)`

SetNodeList sets NodeList field to given value.

### HasNodeList

`func (o *V0041OpenapiReservationRespReservationsInner) HasNodeList() bool`

HasNodeList returns a boolean if a field has been set.

### GetPartition

`func (o *V0041OpenapiReservationRespReservationsInner) GetPartition() string`

GetPartition returns the Partition field if non-nil, zero value otherwise.

### GetPartitionOk

`func (o *V0041OpenapiReservationRespReservationsInner) GetPartitionOk() (*string, bool)`

GetPartitionOk returns a tuple with the Partition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartition

`func (o *V0041OpenapiReservationRespReservationsInner) SetPartition(v string)`

SetPartition sets Partition field to given value.

### HasPartition

`func (o *V0041OpenapiReservationRespReservationsInner) HasPartition() bool`

HasPartition returns a boolean if a field has been set.

### GetPurgeCompleted

`func (o *V0041OpenapiReservationRespReservationsInner) GetPurgeCompleted() V0041OpenapiReservationRespReservationsInnerPurgeCompleted`

GetPurgeCompleted returns the PurgeCompleted field if non-nil, zero value otherwise.

### GetPurgeCompletedOk

`func (o *V0041OpenapiReservationRespReservationsInner) GetPurgeCompletedOk() (*V0041OpenapiReservationRespReservationsInnerPurgeCompleted, bool)`

GetPurgeCompletedOk returns a tuple with the PurgeCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurgeCompleted

`func (o *V0041OpenapiReservationRespReservationsInner) SetPurgeCompleted(v V0041OpenapiReservationRespReservationsInnerPurgeCompleted)`

SetPurgeCompleted sets PurgeCompleted field to given value.

### HasPurgeCompleted

`func (o *V0041OpenapiReservationRespReservationsInner) HasPurgeCompleted() bool`

HasPurgeCompleted returns a boolean if a field has been set.

### GetStartTime

`func (o *V0041OpenapiReservationRespReservationsInner) GetStartTime() V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *V0041OpenapiReservationRespReservationsInner) GetStartTimeOk() (*V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *V0041OpenapiReservationRespReservationsInner) SetStartTime(v V0041OpenapiSlurmdbdJobsRespJobsInnerRequiredMemoryPerCpu)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *V0041OpenapiReservationRespReservationsInner) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetWatts

`func (o *V0041OpenapiReservationRespReservationsInner) GetWatts() V0041OpenapiReservationRespReservationsInnerWatts`

GetWatts returns the Watts field if non-nil, zero value otherwise.

### GetWattsOk

`func (o *V0041OpenapiReservationRespReservationsInner) GetWattsOk() (*V0041OpenapiReservationRespReservationsInnerWatts, bool)`

GetWattsOk returns a tuple with the Watts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatts

`func (o *V0041OpenapiReservationRespReservationsInner) SetWatts(v V0041OpenapiReservationRespReservationsInnerWatts)`

SetWatts sets Watts field to given value.

### HasWatts

`func (o *V0041OpenapiReservationRespReservationsInner) HasWatts() bool`

HasWatts returns a boolean if a field has been set.

### GetTres

`func (o *V0041OpenapiReservationRespReservationsInner) GetTres() string`

GetTres returns the Tres field if non-nil, zero value otherwise.

### GetTresOk

`func (o *V0041OpenapiReservationRespReservationsInner) GetTresOk() (*string, bool)`

GetTresOk returns a tuple with the Tres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTres

`func (o *V0041OpenapiReservationRespReservationsInner) SetTres(v string)`

SetTres sets Tres field to given value.

### HasTres

`func (o *V0041OpenapiReservationRespReservationsInner) HasTres() bool`

HasTres returns a boolean if a field has been set.

### GetUsers

`func (o *V0041OpenapiReservationRespReservationsInner) GetUsers() string`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *V0041OpenapiReservationRespReservationsInner) GetUsersOk() (*string, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *V0041OpenapiReservationRespReservationsInner) SetUsers(v string)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *V0041OpenapiReservationRespReservationsInner) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


