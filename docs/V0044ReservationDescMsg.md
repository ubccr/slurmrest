# V0044ReservationDescMsg

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accounts** | Pointer to **[]string** |  | [optional] 
**BurstBuffer** | Pointer to **string** | BurstBuffer | [optional] 
**Comment** | Pointer to **string** | Arbitrary string | [optional] 
**CoreCount** | Pointer to [**V0044Uint32NoValStruct**](V0044Uint32NoValStruct.md) |  | [optional] 
**Duration** | Pointer to [**V0044Uint32NoValStruct**](V0044Uint32NoValStruct.md) |  | [optional] 
**EndTime** | Pointer to [**V0044Uint64NoValStruct**](V0044Uint64NoValStruct.md) |  | [optional] 
**Features** | Pointer to **string** | Requested node features. Multiple values may be \&quot;&amp;\&quot; separated if all features are required (AND operation) or separated by \&quot;|\&quot; if any of the specified features are required (OR operation). Parenthesis are also supported for features to be ANDed together with counts of nodes having the specified features. | [optional] 
**Flags** | Pointer to **[]string** | Flags associated with this reservation. Note, to remove flags use \&quot;NO_\&quot; prefixed flag excluding NO_HOLD_JOBS_AFTER_END | [optional] 
**Groups** | Pointer to **[]string** |  | [optional] 
**Licenses** | Pointer to **[]string** |  | [optional] 
**MaxStartDelay** | Pointer to [**V0044Uint32NoValStruct**](V0044Uint32NoValStruct.md) |  | [optional] 
**Name** | Pointer to **string** | ReservationName | [optional] 
**NodeCount** | Pointer to [**V0044Uint32NoValStruct**](V0044Uint32NoValStruct.md) |  | [optional] 
**NodeList** | Pointer to **[]string** |  | [optional] 
**Partition** | Pointer to **string** | Partition used to reserve nodes from. This will attempt to allocate all nodes in the specified partition unless you request fewer resources than are available with core_cnt, node_cnt or tres. | [optional] 
**PurgeCompleted** | Pointer to [**V0044ReservationInfoPurgeCompleted**](V0044ReservationInfoPurgeCompleted.md) |  | [optional] 
**StartTime** | Pointer to [**V0044Uint64NoValStruct**](V0044Uint64NoValStruct.md) |  | [optional] 
**Tres** | Pointer to [**[]V0044Tres**](V0044Tres.md) |  | [optional] 
**Users** | Pointer to **[]string** |  | [optional] 

## Methods

### NewV0044ReservationDescMsg

`func NewV0044ReservationDescMsg() *V0044ReservationDescMsg`

NewV0044ReservationDescMsg instantiates a new V0044ReservationDescMsg object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0044ReservationDescMsgWithDefaults

`func NewV0044ReservationDescMsgWithDefaults() *V0044ReservationDescMsg`

NewV0044ReservationDescMsgWithDefaults instantiates a new V0044ReservationDescMsg object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccounts

`func (o *V0044ReservationDescMsg) GetAccounts() []string`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *V0044ReservationDescMsg) GetAccountsOk() (*[]string, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *V0044ReservationDescMsg) SetAccounts(v []string)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *V0044ReservationDescMsg) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetBurstBuffer

`func (o *V0044ReservationDescMsg) GetBurstBuffer() string`

GetBurstBuffer returns the BurstBuffer field if non-nil, zero value otherwise.

### GetBurstBufferOk

`func (o *V0044ReservationDescMsg) GetBurstBufferOk() (*string, bool)`

GetBurstBufferOk returns a tuple with the BurstBuffer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurstBuffer

`func (o *V0044ReservationDescMsg) SetBurstBuffer(v string)`

SetBurstBuffer sets BurstBuffer field to given value.

### HasBurstBuffer

`func (o *V0044ReservationDescMsg) HasBurstBuffer() bool`

HasBurstBuffer returns a boolean if a field has been set.

### GetComment

`func (o *V0044ReservationDescMsg) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *V0044ReservationDescMsg) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *V0044ReservationDescMsg) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *V0044ReservationDescMsg) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCoreCount

`func (o *V0044ReservationDescMsg) GetCoreCount() V0044Uint32NoValStruct`

GetCoreCount returns the CoreCount field if non-nil, zero value otherwise.

### GetCoreCountOk

`func (o *V0044ReservationDescMsg) GetCoreCountOk() (*V0044Uint32NoValStruct, bool)`

GetCoreCountOk returns a tuple with the CoreCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoreCount

`func (o *V0044ReservationDescMsg) SetCoreCount(v V0044Uint32NoValStruct)`

SetCoreCount sets CoreCount field to given value.

### HasCoreCount

`func (o *V0044ReservationDescMsg) HasCoreCount() bool`

HasCoreCount returns a boolean if a field has been set.

### GetDuration

`func (o *V0044ReservationDescMsg) GetDuration() V0044Uint32NoValStruct`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *V0044ReservationDescMsg) GetDurationOk() (*V0044Uint32NoValStruct, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *V0044ReservationDescMsg) SetDuration(v V0044Uint32NoValStruct)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *V0044ReservationDescMsg) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetEndTime

`func (o *V0044ReservationDescMsg) GetEndTime() V0044Uint64NoValStruct`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *V0044ReservationDescMsg) GetEndTimeOk() (*V0044Uint64NoValStruct, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *V0044ReservationDescMsg) SetEndTime(v V0044Uint64NoValStruct)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *V0044ReservationDescMsg) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetFeatures

`func (o *V0044ReservationDescMsg) GetFeatures() string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *V0044ReservationDescMsg) GetFeaturesOk() (*string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *V0044ReservationDescMsg) SetFeatures(v string)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *V0044ReservationDescMsg) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetFlags

`func (o *V0044ReservationDescMsg) GetFlags() []string`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *V0044ReservationDescMsg) GetFlagsOk() (*[]string, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *V0044ReservationDescMsg) SetFlags(v []string)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *V0044ReservationDescMsg) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetGroups

`func (o *V0044ReservationDescMsg) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *V0044ReservationDescMsg) GetGroupsOk() (*[]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *V0044ReservationDescMsg) SetGroups(v []string)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *V0044ReservationDescMsg) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetLicenses

`func (o *V0044ReservationDescMsg) GetLicenses() []string`

GetLicenses returns the Licenses field if non-nil, zero value otherwise.

### GetLicensesOk

`func (o *V0044ReservationDescMsg) GetLicensesOk() (*[]string, bool)`

GetLicensesOk returns a tuple with the Licenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenses

`func (o *V0044ReservationDescMsg) SetLicenses(v []string)`

SetLicenses sets Licenses field to given value.

### HasLicenses

`func (o *V0044ReservationDescMsg) HasLicenses() bool`

HasLicenses returns a boolean if a field has been set.

### GetMaxStartDelay

`func (o *V0044ReservationDescMsg) GetMaxStartDelay() V0044Uint32NoValStruct`

GetMaxStartDelay returns the MaxStartDelay field if non-nil, zero value otherwise.

### GetMaxStartDelayOk

`func (o *V0044ReservationDescMsg) GetMaxStartDelayOk() (*V0044Uint32NoValStruct, bool)`

GetMaxStartDelayOk returns a tuple with the MaxStartDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStartDelay

`func (o *V0044ReservationDescMsg) SetMaxStartDelay(v V0044Uint32NoValStruct)`

SetMaxStartDelay sets MaxStartDelay field to given value.

### HasMaxStartDelay

`func (o *V0044ReservationDescMsg) HasMaxStartDelay() bool`

HasMaxStartDelay returns a boolean if a field has been set.

### GetName

`func (o *V0044ReservationDescMsg) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V0044ReservationDescMsg) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V0044ReservationDescMsg) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V0044ReservationDescMsg) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNodeCount

`func (o *V0044ReservationDescMsg) GetNodeCount() V0044Uint32NoValStruct`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *V0044ReservationDescMsg) GetNodeCountOk() (*V0044Uint32NoValStruct, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *V0044ReservationDescMsg) SetNodeCount(v V0044Uint32NoValStruct)`

SetNodeCount sets NodeCount field to given value.

### HasNodeCount

`func (o *V0044ReservationDescMsg) HasNodeCount() bool`

HasNodeCount returns a boolean if a field has been set.

### GetNodeList

`func (o *V0044ReservationDescMsg) GetNodeList() []string`

GetNodeList returns the NodeList field if non-nil, zero value otherwise.

### GetNodeListOk

`func (o *V0044ReservationDescMsg) GetNodeListOk() (*[]string, bool)`

GetNodeListOk returns a tuple with the NodeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeList

`func (o *V0044ReservationDescMsg) SetNodeList(v []string)`

SetNodeList sets NodeList field to given value.

### HasNodeList

`func (o *V0044ReservationDescMsg) HasNodeList() bool`

HasNodeList returns a boolean if a field has been set.

### GetPartition

`func (o *V0044ReservationDescMsg) GetPartition() string`

GetPartition returns the Partition field if non-nil, zero value otherwise.

### GetPartitionOk

`func (o *V0044ReservationDescMsg) GetPartitionOk() (*string, bool)`

GetPartitionOk returns a tuple with the Partition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartition

`func (o *V0044ReservationDescMsg) SetPartition(v string)`

SetPartition sets Partition field to given value.

### HasPartition

`func (o *V0044ReservationDescMsg) HasPartition() bool`

HasPartition returns a boolean if a field has been set.

### GetPurgeCompleted

`func (o *V0044ReservationDescMsg) GetPurgeCompleted() V0044ReservationInfoPurgeCompleted`

GetPurgeCompleted returns the PurgeCompleted field if non-nil, zero value otherwise.

### GetPurgeCompletedOk

`func (o *V0044ReservationDescMsg) GetPurgeCompletedOk() (*V0044ReservationInfoPurgeCompleted, bool)`

GetPurgeCompletedOk returns a tuple with the PurgeCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurgeCompleted

`func (o *V0044ReservationDescMsg) SetPurgeCompleted(v V0044ReservationInfoPurgeCompleted)`

SetPurgeCompleted sets PurgeCompleted field to given value.

### HasPurgeCompleted

`func (o *V0044ReservationDescMsg) HasPurgeCompleted() bool`

HasPurgeCompleted returns a boolean if a field has been set.

### GetStartTime

`func (o *V0044ReservationDescMsg) GetStartTime() V0044Uint64NoValStruct`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *V0044ReservationDescMsg) GetStartTimeOk() (*V0044Uint64NoValStruct, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *V0044ReservationDescMsg) SetStartTime(v V0044Uint64NoValStruct)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *V0044ReservationDescMsg) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetTres

`func (o *V0044ReservationDescMsg) GetTres() []V0044Tres`

GetTres returns the Tres field if non-nil, zero value otherwise.

### GetTresOk

`func (o *V0044ReservationDescMsg) GetTresOk() (*[]V0044Tres, bool)`

GetTresOk returns a tuple with the Tres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTres

`func (o *V0044ReservationDescMsg) SetTres(v []V0044Tres)`

SetTres sets Tres field to given value.

### HasTres

`func (o *V0044ReservationDescMsg) HasTres() bool`

HasTres returns a boolean if a field has been set.

### GetUsers

`func (o *V0044ReservationDescMsg) GetUsers() []string`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *V0044ReservationDescMsg) GetUsersOk() (*[]string, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *V0044ReservationDescMsg) SetUsers(v []string)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *V0044ReservationDescMsg) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


