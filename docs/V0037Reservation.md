# V0037Reservation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accounts** | Pointer to **string** | Allowed accounts | [optional] 
**BurstBuffer** | Pointer to **string** | Reserved burst buffer | [optional] 
**CoreCount** | Pointer to **int32** | Number of reserved cores | [optional] 
**CoreSpecCnt** | Pointer to **int32** | Number of reserved specialized cores | [optional] 
**EndTime** | Pointer to **int32** | End time of the reservation | [optional] 
**Features** | Pointer to **string** | List of features | [optional] 
**Flags** | Pointer to **[]string** | Reservation options | [optional] 
**Groups** | Pointer to **string** | List of groups permitted to use the reserved nodes | [optional] 
**Licenses** | Pointer to **string** | List of licenses | [optional] 
**MaxStartDelay** | Pointer to **int32** | Maximum delay in which jobs outside of the reservation will be permitted to overlap once any jobs are queued for the reservation | [optional] 
**Name** | Pointer to **string** | Reservationn name | [optional] 
**NodeCount** | Pointer to **int32** | Count of nodes reserved | [optional] 
**NodeList** | Pointer to **string** | List of reserved nodes | [optional] 
**Partition** | Pointer to **string** | Partition | [optional] 
**PurgeCompleted** | Pointer to [**V0037ReservationPurgeCompleted**](V0037ReservationPurgeCompleted.md) |  | [optional] 
**StartTime** | Pointer to **int32** | Start time of reservation | [optional] 
**Watts** | Pointer to **int32** | amount of power to reserve in watts | [optional] 
**Tres** | Pointer to **string** | List of TRES | [optional] 
**Users** | Pointer to **string** | List of users | [optional] 

## Methods

### NewV0037Reservation

`func NewV0037Reservation() *V0037Reservation`

NewV0037Reservation instantiates a new V0037Reservation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0037ReservationWithDefaults

`func NewV0037ReservationWithDefaults() *V0037Reservation`

NewV0037ReservationWithDefaults instantiates a new V0037Reservation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccounts

`func (o *V0037Reservation) GetAccounts() string`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *V0037Reservation) GetAccountsOk() (*string, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *V0037Reservation) SetAccounts(v string)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *V0037Reservation) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetBurstBuffer

`func (o *V0037Reservation) GetBurstBuffer() string`

GetBurstBuffer returns the BurstBuffer field if non-nil, zero value otherwise.

### GetBurstBufferOk

`func (o *V0037Reservation) GetBurstBufferOk() (*string, bool)`

GetBurstBufferOk returns a tuple with the BurstBuffer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurstBuffer

`func (o *V0037Reservation) SetBurstBuffer(v string)`

SetBurstBuffer sets BurstBuffer field to given value.

### HasBurstBuffer

`func (o *V0037Reservation) HasBurstBuffer() bool`

HasBurstBuffer returns a boolean if a field has been set.

### GetCoreCount

`func (o *V0037Reservation) GetCoreCount() int32`

GetCoreCount returns the CoreCount field if non-nil, zero value otherwise.

### GetCoreCountOk

`func (o *V0037Reservation) GetCoreCountOk() (*int32, bool)`

GetCoreCountOk returns a tuple with the CoreCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoreCount

`func (o *V0037Reservation) SetCoreCount(v int32)`

SetCoreCount sets CoreCount field to given value.

### HasCoreCount

`func (o *V0037Reservation) HasCoreCount() bool`

HasCoreCount returns a boolean if a field has been set.

### GetCoreSpecCnt

`func (o *V0037Reservation) GetCoreSpecCnt() int32`

GetCoreSpecCnt returns the CoreSpecCnt field if non-nil, zero value otherwise.

### GetCoreSpecCntOk

`func (o *V0037Reservation) GetCoreSpecCntOk() (*int32, bool)`

GetCoreSpecCntOk returns a tuple with the CoreSpecCnt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoreSpecCnt

`func (o *V0037Reservation) SetCoreSpecCnt(v int32)`

SetCoreSpecCnt sets CoreSpecCnt field to given value.

### HasCoreSpecCnt

`func (o *V0037Reservation) HasCoreSpecCnt() bool`

HasCoreSpecCnt returns a boolean if a field has been set.

### GetEndTime

`func (o *V0037Reservation) GetEndTime() int32`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *V0037Reservation) GetEndTimeOk() (*int32, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *V0037Reservation) SetEndTime(v int32)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *V0037Reservation) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetFeatures

`func (o *V0037Reservation) GetFeatures() string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *V0037Reservation) GetFeaturesOk() (*string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *V0037Reservation) SetFeatures(v string)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *V0037Reservation) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetFlags

`func (o *V0037Reservation) GetFlags() []string`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *V0037Reservation) GetFlagsOk() (*[]string, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *V0037Reservation) SetFlags(v []string)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *V0037Reservation) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetGroups

`func (o *V0037Reservation) GetGroups() string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *V0037Reservation) GetGroupsOk() (*string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *V0037Reservation) SetGroups(v string)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *V0037Reservation) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetLicenses

`func (o *V0037Reservation) GetLicenses() string`

GetLicenses returns the Licenses field if non-nil, zero value otherwise.

### GetLicensesOk

`func (o *V0037Reservation) GetLicensesOk() (*string, bool)`

GetLicensesOk returns a tuple with the Licenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenses

`func (o *V0037Reservation) SetLicenses(v string)`

SetLicenses sets Licenses field to given value.

### HasLicenses

`func (o *V0037Reservation) HasLicenses() bool`

HasLicenses returns a boolean if a field has been set.

### GetMaxStartDelay

`func (o *V0037Reservation) GetMaxStartDelay() int32`

GetMaxStartDelay returns the MaxStartDelay field if non-nil, zero value otherwise.

### GetMaxStartDelayOk

`func (o *V0037Reservation) GetMaxStartDelayOk() (*int32, bool)`

GetMaxStartDelayOk returns a tuple with the MaxStartDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStartDelay

`func (o *V0037Reservation) SetMaxStartDelay(v int32)`

SetMaxStartDelay sets MaxStartDelay field to given value.

### HasMaxStartDelay

`func (o *V0037Reservation) HasMaxStartDelay() bool`

HasMaxStartDelay returns a boolean if a field has been set.

### GetName

`func (o *V0037Reservation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V0037Reservation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V0037Reservation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V0037Reservation) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNodeCount

`func (o *V0037Reservation) GetNodeCount() int32`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *V0037Reservation) GetNodeCountOk() (*int32, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *V0037Reservation) SetNodeCount(v int32)`

SetNodeCount sets NodeCount field to given value.

### HasNodeCount

`func (o *V0037Reservation) HasNodeCount() bool`

HasNodeCount returns a boolean if a field has been set.

### GetNodeList

`func (o *V0037Reservation) GetNodeList() string`

GetNodeList returns the NodeList field if non-nil, zero value otherwise.

### GetNodeListOk

`func (o *V0037Reservation) GetNodeListOk() (*string, bool)`

GetNodeListOk returns a tuple with the NodeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeList

`func (o *V0037Reservation) SetNodeList(v string)`

SetNodeList sets NodeList field to given value.

### HasNodeList

`func (o *V0037Reservation) HasNodeList() bool`

HasNodeList returns a boolean if a field has been set.

### GetPartition

`func (o *V0037Reservation) GetPartition() string`

GetPartition returns the Partition field if non-nil, zero value otherwise.

### GetPartitionOk

`func (o *V0037Reservation) GetPartitionOk() (*string, bool)`

GetPartitionOk returns a tuple with the Partition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartition

`func (o *V0037Reservation) SetPartition(v string)`

SetPartition sets Partition field to given value.

### HasPartition

`func (o *V0037Reservation) HasPartition() bool`

HasPartition returns a boolean if a field has been set.

### GetPurgeCompleted

`func (o *V0037Reservation) GetPurgeCompleted() V0037ReservationPurgeCompleted`

GetPurgeCompleted returns the PurgeCompleted field if non-nil, zero value otherwise.

### GetPurgeCompletedOk

`func (o *V0037Reservation) GetPurgeCompletedOk() (*V0037ReservationPurgeCompleted, bool)`

GetPurgeCompletedOk returns a tuple with the PurgeCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurgeCompleted

`func (o *V0037Reservation) SetPurgeCompleted(v V0037ReservationPurgeCompleted)`

SetPurgeCompleted sets PurgeCompleted field to given value.

### HasPurgeCompleted

`func (o *V0037Reservation) HasPurgeCompleted() bool`

HasPurgeCompleted returns a boolean if a field has been set.

### GetStartTime

`func (o *V0037Reservation) GetStartTime() int32`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *V0037Reservation) GetStartTimeOk() (*int32, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *V0037Reservation) SetStartTime(v int32)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *V0037Reservation) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetWatts

`func (o *V0037Reservation) GetWatts() int32`

GetWatts returns the Watts field if non-nil, zero value otherwise.

### GetWattsOk

`func (o *V0037Reservation) GetWattsOk() (*int32, bool)`

GetWattsOk returns a tuple with the Watts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatts

`func (o *V0037Reservation) SetWatts(v int32)`

SetWatts sets Watts field to given value.

### HasWatts

`func (o *V0037Reservation) HasWatts() bool`

HasWatts returns a boolean if a field has been set.

### GetTres

`func (o *V0037Reservation) GetTres() string`

GetTres returns the Tres field if non-nil, zero value otherwise.

### GetTresOk

`func (o *V0037Reservation) GetTresOk() (*string, bool)`

GetTresOk returns a tuple with the Tres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTres

`func (o *V0037Reservation) SetTres(v string)`

SetTres sets Tres field to given value.

### HasTres

`func (o *V0037Reservation) HasTres() bool`

HasTres returns a boolean if a field has been set.

### GetUsers

`func (o *V0037Reservation) GetUsers() string`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *V0037Reservation) GetUsersOk() (*string, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *V0037Reservation) SetUsers(v string)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *V0037Reservation) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


