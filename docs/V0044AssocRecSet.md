# V0044AssocRecSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | Arbitrary comment | [optional] 
**Defaultqos** | Pointer to **string** | Default QOS | [optional] 
**Grpjobs** | Pointer to [**V0044Uint32NoValStruct**](V0044Uint32NoValStruct.md) |  | [optional] 
**Grpjobsaccrue** | Pointer to [**V0044Uint32NoValStruct**](V0044Uint32NoValStruct.md) |  | [optional] 
**Grpsubmitjobs** | Pointer to [**V0044Uint32NoValStruct**](V0044Uint32NoValStruct.md) |  | [optional] 
**Grptres** | Pointer to [**[]V0044Tres**](V0044Tres.md) |  | [optional] 
**Grptresmins** | Pointer to [**[]V0044Tres**](V0044Tres.md) |  | [optional] 
**Grptresrunmins** | Pointer to [**[]V0044Tres**](V0044Tres.md) |  | [optional] 
**Grpwall** | Pointer to [**V0044Uint32NoValStruct**](V0044Uint32NoValStruct.md) |  | [optional] 
**Maxjobs** | Pointer to [**V0044Uint32NoValStruct**](V0044Uint32NoValStruct.md) |  | [optional] 
**Maxjobsaccrue** | Pointer to [**V0044Uint32NoValStruct**](V0044Uint32NoValStruct.md) |  | [optional] 
**Maxsubmitjobs** | Pointer to [**V0044Uint32NoValStruct**](V0044Uint32NoValStruct.md) |  | [optional] 
**Maxtresminsperjob** | Pointer to [**[]V0044Tres**](V0044Tres.md) |  | [optional] 
**Maxtresrunmins** | Pointer to [**[]V0044Tres**](V0044Tres.md) |  | [optional] 
**Maxtresperjob** | Pointer to [**[]V0044Tres**](V0044Tres.md) |  | [optional] 
**Maxtrespernode** | Pointer to [**[]V0044Tres**](V0044Tres.md) |  | [optional] 
**Maxwalldurationperjob** | Pointer to [**V0044Uint32NoValStruct**](V0044Uint32NoValStruct.md) |  | [optional] 
**Minpriothresh** | Pointer to [**V0044Uint32NoValStruct**](V0044Uint32NoValStruct.md) |  | [optional] 
**Parent** | Pointer to **string** | Name of parent account | [optional] 
**Priority** | Pointer to [**V0044Uint32NoValStruct**](V0044Uint32NoValStruct.md) |  | [optional] 
**Qoslevel** | Pointer to **[]string** | List of QOS names | [optional] 
**Fairshare** | Pointer to **int32** | Allocated shares used for fairshare calculation | [optional] 

## Methods

### NewV0044AssocRecSet

`func NewV0044AssocRecSet() *V0044AssocRecSet`

NewV0044AssocRecSet instantiates a new V0044AssocRecSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0044AssocRecSetWithDefaults

`func NewV0044AssocRecSetWithDefaults() *V0044AssocRecSet`

NewV0044AssocRecSetWithDefaults instantiates a new V0044AssocRecSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *V0044AssocRecSet) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *V0044AssocRecSet) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *V0044AssocRecSet) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *V0044AssocRecSet) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDefaultqos

`func (o *V0044AssocRecSet) GetDefaultqos() string`

GetDefaultqos returns the Defaultqos field if non-nil, zero value otherwise.

### GetDefaultqosOk

`func (o *V0044AssocRecSet) GetDefaultqosOk() (*string, bool)`

GetDefaultqosOk returns a tuple with the Defaultqos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultqos

`func (o *V0044AssocRecSet) SetDefaultqos(v string)`

SetDefaultqos sets Defaultqos field to given value.

### HasDefaultqos

`func (o *V0044AssocRecSet) HasDefaultqos() bool`

HasDefaultqos returns a boolean if a field has been set.

### GetGrpjobs

`func (o *V0044AssocRecSet) GetGrpjobs() V0044Uint32NoValStruct`

GetGrpjobs returns the Grpjobs field if non-nil, zero value otherwise.

### GetGrpjobsOk

`func (o *V0044AssocRecSet) GetGrpjobsOk() (*V0044Uint32NoValStruct, bool)`

GetGrpjobsOk returns a tuple with the Grpjobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrpjobs

`func (o *V0044AssocRecSet) SetGrpjobs(v V0044Uint32NoValStruct)`

SetGrpjobs sets Grpjobs field to given value.

### HasGrpjobs

`func (o *V0044AssocRecSet) HasGrpjobs() bool`

HasGrpjobs returns a boolean if a field has been set.

### GetGrpjobsaccrue

`func (o *V0044AssocRecSet) GetGrpjobsaccrue() V0044Uint32NoValStruct`

GetGrpjobsaccrue returns the Grpjobsaccrue field if non-nil, zero value otherwise.

### GetGrpjobsaccrueOk

`func (o *V0044AssocRecSet) GetGrpjobsaccrueOk() (*V0044Uint32NoValStruct, bool)`

GetGrpjobsaccrueOk returns a tuple with the Grpjobsaccrue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrpjobsaccrue

`func (o *V0044AssocRecSet) SetGrpjobsaccrue(v V0044Uint32NoValStruct)`

SetGrpjobsaccrue sets Grpjobsaccrue field to given value.

### HasGrpjobsaccrue

`func (o *V0044AssocRecSet) HasGrpjobsaccrue() bool`

HasGrpjobsaccrue returns a boolean if a field has been set.

### GetGrpsubmitjobs

`func (o *V0044AssocRecSet) GetGrpsubmitjobs() V0044Uint32NoValStruct`

GetGrpsubmitjobs returns the Grpsubmitjobs field if non-nil, zero value otherwise.

### GetGrpsubmitjobsOk

`func (o *V0044AssocRecSet) GetGrpsubmitjobsOk() (*V0044Uint32NoValStruct, bool)`

GetGrpsubmitjobsOk returns a tuple with the Grpsubmitjobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrpsubmitjobs

`func (o *V0044AssocRecSet) SetGrpsubmitjobs(v V0044Uint32NoValStruct)`

SetGrpsubmitjobs sets Grpsubmitjobs field to given value.

### HasGrpsubmitjobs

`func (o *V0044AssocRecSet) HasGrpsubmitjobs() bool`

HasGrpsubmitjobs returns a boolean if a field has been set.

### GetGrptres

`func (o *V0044AssocRecSet) GetGrptres() []V0044Tres`

GetGrptres returns the Grptres field if non-nil, zero value otherwise.

### GetGrptresOk

`func (o *V0044AssocRecSet) GetGrptresOk() (*[]V0044Tres, bool)`

GetGrptresOk returns a tuple with the Grptres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrptres

`func (o *V0044AssocRecSet) SetGrptres(v []V0044Tres)`

SetGrptres sets Grptres field to given value.

### HasGrptres

`func (o *V0044AssocRecSet) HasGrptres() bool`

HasGrptres returns a boolean if a field has been set.

### GetGrptresmins

`func (o *V0044AssocRecSet) GetGrptresmins() []V0044Tres`

GetGrptresmins returns the Grptresmins field if non-nil, zero value otherwise.

### GetGrptresminsOk

`func (o *V0044AssocRecSet) GetGrptresminsOk() (*[]V0044Tres, bool)`

GetGrptresminsOk returns a tuple with the Grptresmins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrptresmins

`func (o *V0044AssocRecSet) SetGrptresmins(v []V0044Tres)`

SetGrptresmins sets Grptresmins field to given value.

### HasGrptresmins

`func (o *V0044AssocRecSet) HasGrptresmins() bool`

HasGrptresmins returns a boolean if a field has been set.

### GetGrptresrunmins

`func (o *V0044AssocRecSet) GetGrptresrunmins() []V0044Tres`

GetGrptresrunmins returns the Grptresrunmins field if non-nil, zero value otherwise.

### GetGrptresrunminsOk

`func (o *V0044AssocRecSet) GetGrptresrunminsOk() (*[]V0044Tres, bool)`

GetGrptresrunminsOk returns a tuple with the Grptresrunmins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrptresrunmins

`func (o *V0044AssocRecSet) SetGrptresrunmins(v []V0044Tres)`

SetGrptresrunmins sets Grptresrunmins field to given value.

### HasGrptresrunmins

`func (o *V0044AssocRecSet) HasGrptresrunmins() bool`

HasGrptresrunmins returns a boolean if a field has been set.

### GetGrpwall

`func (o *V0044AssocRecSet) GetGrpwall() V0044Uint32NoValStruct`

GetGrpwall returns the Grpwall field if non-nil, zero value otherwise.

### GetGrpwallOk

`func (o *V0044AssocRecSet) GetGrpwallOk() (*V0044Uint32NoValStruct, bool)`

GetGrpwallOk returns a tuple with the Grpwall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrpwall

`func (o *V0044AssocRecSet) SetGrpwall(v V0044Uint32NoValStruct)`

SetGrpwall sets Grpwall field to given value.

### HasGrpwall

`func (o *V0044AssocRecSet) HasGrpwall() bool`

HasGrpwall returns a boolean if a field has been set.

### GetMaxjobs

`func (o *V0044AssocRecSet) GetMaxjobs() V0044Uint32NoValStruct`

GetMaxjobs returns the Maxjobs field if non-nil, zero value otherwise.

### GetMaxjobsOk

`func (o *V0044AssocRecSet) GetMaxjobsOk() (*V0044Uint32NoValStruct, bool)`

GetMaxjobsOk returns a tuple with the Maxjobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxjobs

`func (o *V0044AssocRecSet) SetMaxjobs(v V0044Uint32NoValStruct)`

SetMaxjobs sets Maxjobs field to given value.

### HasMaxjobs

`func (o *V0044AssocRecSet) HasMaxjobs() bool`

HasMaxjobs returns a boolean if a field has been set.

### GetMaxjobsaccrue

`func (o *V0044AssocRecSet) GetMaxjobsaccrue() V0044Uint32NoValStruct`

GetMaxjobsaccrue returns the Maxjobsaccrue field if non-nil, zero value otherwise.

### GetMaxjobsaccrueOk

`func (o *V0044AssocRecSet) GetMaxjobsaccrueOk() (*V0044Uint32NoValStruct, bool)`

GetMaxjobsaccrueOk returns a tuple with the Maxjobsaccrue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxjobsaccrue

`func (o *V0044AssocRecSet) SetMaxjobsaccrue(v V0044Uint32NoValStruct)`

SetMaxjobsaccrue sets Maxjobsaccrue field to given value.

### HasMaxjobsaccrue

`func (o *V0044AssocRecSet) HasMaxjobsaccrue() bool`

HasMaxjobsaccrue returns a boolean if a field has been set.

### GetMaxsubmitjobs

`func (o *V0044AssocRecSet) GetMaxsubmitjobs() V0044Uint32NoValStruct`

GetMaxsubmitjobs returns the Maxsubmitjobs field if non-nil, zero value otherwise.

### GetMaxsubmitjobsOk

`func (o *V0044AssocRecSet) GetMaxsubmitjobsOk() (*V0044Uint32NoValStruct, bool)`

GetMaxsubmitjobsOk returns a tuple with the Maxsubmitjobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxsubmitjobs

`func (o *V0044AssocRecSet) SetMaxsubmitjobs(v V0044Uint32NoValStruct)`

SetMaxsubmitjobs sets Maxsubmitjobs field to given value.

### HasMaxsubmitjobs

`func (o *V0044AssocRecSet) HasMaxsubmitjobs() bool`

HasMaxsubmitjobs returns a boolean if a field has been set.

### GetMaxtresminsperjob

`func (o *V0044AssocRecSet) GetMaxtresminsperjob() []V0044Tres`

GetMaxtresminsperjob returns the Maxtresminsperjob field if non-nil, zero value otherwise.

### GetMaxtresminsperjobOk

`func (o *V0044AssocRecSet) GetMaxtresminsperjobOk() (*[]V0044Tres, bool)`

GetMaxtresminsperjobOk returns a tuple with the Maxtresminsperjob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxtresminsperjob

`func (o *V0044AssocRecSet) SetMaxtresminsperjob(v []V0044Tres)`

SetMaxtresminsperjob sets Maxtresminsperjob field to given value.

### HasMaxtresminsperjob

`func (o *V0044AssocRecSet) HasMaxtresminsperjob() bool`

HasMaxtresminsperjob returns a boolean if a field has been set.

### GetMaxtresrunmins

`func (o *V0044AssocRecSet) GetMaxtresrunmins() []V0044Tres`

GetMaxtresrunmins returns the Maxtresrunmins field if non-nil, zero value otherwise.

### GetMaxtresrunminsOk

`func (o *V0044AssocRecSet) GetMaxtresrunminsOk() (*[]V0044Tres, bool)`

GetMaxtresrunminsOk returns a tuple with the Maxtresrunmins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxtresrunmins

`func (o *V0044AssocRecSet) SetMaxtresrunmins(v []V0044Tres)`

SetMaxtresrunmins sets Maxtresrunmins field to given value.

### HasMaxtresrunmins

`func (o *V0044AssocRecSet) HasMaxtresrunmins() bool`

HasMaxtresrunmins returns a boolean if a field has been set.

### GetMaxtresperjob

`func (o *V0044AssocRecSet) GetMaxtresperjob() []V0044Tres`

GetMaxtresperjob returns the Maxtresperjob field if non-nil, zero value otherwise.

### GetMaxtresperjobOk

`func (o *V0044AssocRecSet) GetMaxtresperjobOk() (*[]V0044Tres, bool)`

GetMaxtresperjobOk returns a tuple with the Maxtresperjob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxtresperjob

`func (o *V0044AssocRecSet) SetMaxtresperjob(v []V0044Tres)`

SetMaxtresperjob sets Maxtresperjob field to given value.

### HasMaxtresperjob

`func (o *V0044AssocRecSet) HasMaxtresperjob() bool`

HasMaxtresperjob returns a boolean if a field has been set.

### GetMaxtrespernode

`func (o *V0044AssocRecSet) GetMaxtrespernode() []V0044Tres`

GetMaxtrespernode returns the Maxtrespernode field if non-nil, zero value otherwise.

### GetMaxtrespernodeOk

`func (o *V0044AssocRecSet) GetMaxtrespernodeOk() (*[]V0044Tres, bool)`

GetMaxtrespernodeOk returns a tuple with the Maxtrespernode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxtrespernode

`func (o *V0044AssocRecSet) SetMaxtrespernode(v []V0044Tres)`

SetMaxtrespernode sets Maxtrespernode field to given value.

### HasMaxtrespernode

`func (o *V0044AssocRecSet) HasMaxtrespernode() bool`

HasMaxtrespernode returns a boolean if a field has been set.

### GetMaxwalldurationperjob

`func (o *V0044AssocRecSet) GetMaxwalldurationperjob() V0044Uint32NoValStruct`

GetMaxwalldurationperjob returns the Maxwalldurationperjob field if non-nil, zero value otherwise.

### GetMaxwalldurationperjobOk

`func (o *V0044AssocRecSet) GetMaxwalldurationperjobOk() (*V0044Uint32NoValStruct, bool)`

GetMaxwalldurationperjobOk returns a tuple with the Maxwalldurationperjob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxwalldurationperjob

`func (o *V0044AssocRecSet) SetMaxwalldurationperjob(v V0044Uint32NoValStruct)`

SetMaxwalldurationperjob sets Maxwalldurationperjob field to given value.

### HasMaxwalldurationperjob

`func (o *V0044AssocRecSet) HasMaxwalldurationperjob() bool`

HasMaxwalldurationperjob returns a boolean if a field has been set.

### GetMinpriothresh

`func (o *V0044AssocRecSet) GetMinpriothresh() V0044Uint32NoValStruct`

GetMinpriothresh returns the Minpriothresh field if non-nil, zero value otherwise.

### GetMinpriothreshOk

`func (o *V0044AssocRecSet) GetMinpriothreshOk() (*V0044Uint32NoValStruct, bool)`

GetMinpriothreshOk returns a tuple with the Minpriothresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinpriothresh

`func (o *V0044AssocRecSet) SetMinpriothresh(v V0044Uint32NoValStruct)`

SetMinpriothresh sets Minpriothresh field to given value.

### HasMinpriothresh

`func (o *V0044AssocRecSet) HasMinpriothresh() bool`

HasMinpriothresh returns a boolean if a field has been set.

### GetParent

`func (o *V0044AssocRecSet) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *V0044AssocRecSet) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *V0044AssocRecSet) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *V0044AssocRecSet) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPriority

`func (o *V0044AssocRecSet) GetPriority() V0044Uint32NoValStruct`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *V0044AssocRecSet) GetPriorityOk() (*V0044Uint32NoValStruct, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *V0044AssocRecSet) SetPriority(v V0044Uint32NoValStruct)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *V0044AssocRecSet) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetQoslevel

`func (o *V0044AssocRecSet) GetQoslevel() []string`

GetQoslevel returns the Qoslevel field if non-nil, zero value otherwise.

### GetQoslevelOk

`func (o *V0044AssocRecSet) GetQoslevelOk() (*[]string, bool)`

GetQoslevelOk returns a tuple with the Qoslevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQoslevel

`func (o *V0044AssocRecSet) SetQoslevel(v []string)`

SetQoslevel sets Qoslevel field to given value.

### HasQoslevel

`func (o *V0044AssocRecSet) HasQoslevel() bool`

HasQoslevel returns a boolean if a field has been set.

### GetFairshare

`func (o *V0044AssocRecSet) GetFairshare() int32`

GetFairshare returns the Fairshare field if non-nil, zero value otherwise.

### GetFairshareOk

`func (o *V0044AssocRecSet) GetFairshareOk() (*int32, bool)`

GetFairshareOk returns a tuple with the Fairshare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFairshare

`func (o *V0044AssocRecSet) SetFairshare(v int32)`

SetFairshare sets Fairshare field to given value.

### HasFairshare

`func (o *V0044AssocRecSet) HasFairshare() bool`

HasFairshare returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


