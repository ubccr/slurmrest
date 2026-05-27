# V0044UpdateNodeMsg

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | Arbitrary comment | [optional] 
**CpuBind** | Pointer to **int32** | Default method for binding tasks to allocated CPUs | [optional] 
**Extra** | Pointer to **string** | Arbitrary string used for node filtering if extra constraints are enabled | [optional] 
**Features** | Pointer to **[]string** |  | [optional] 
**FeaturesAct** | Pointer to **[]string** |  | [optional] 
**Gres** | Pointer to **string** | Generic resources | [optional] 
**Address** | Pointer to **[]string** |  | [optional] 
**Hostname** | Pointer to **[]string** |  | [optional] 
**Name** | Pointer to **[]string** |  | [optional] 
**State** | Pointer to **[]string** | New state to assign to the node | [optional] 
**Reason** | Pointer to **string** | Reason for node being DOWN or DRAINING | [optional] 
**ReasonUid** | Pointer to **string** | User ID to associate with the reason (needed if user root is sending message) | [optional] 
**ResumeAfter** | Pointer to [**V0044Uint32NoValStruct**](V0044Uint32NoValStruct.md) |  | [optional] 
**TopologyStr** | Pointer to **string** | Topology | [optional] 
**Weight** | Pointer to [**V0044Uint32NoValStruct**](V0044Uint32NoValStruct.md) |  | [optional] 

## Methods

### NewV0044UpdateNodeMsg

`func NewV0044UpdateNodeMsg() *V0044UpdateNodeMsg`

NewV0044UpdateNodeMsg instantiates a new V0044UpdateNodeMsg object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0044UpdateNodeMsgWithDefaults

`func NewV0044UpdateNodeMsgWithDefaults() *V0044UpdateNodeMsg`

NewV0044UpdateNodeMsgWithDefaults instantiates a new V0044UpdateNodeMsg object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *V0044UpdateNodeMsg) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *V0044UpdateNodeMsg) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *V0044UpdateNodeMsg) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *V0044UpdateNodeMsg) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCpuBind

`func (o *V0044UpdateNodeMsg) GetCpuBind() int32`

GetCpuBind returns the CpuBind field if non-nil, zero value otherwise.

### GetCpuBindOk

`func (o *V0044UpdateNodeMsg) GetCpuBindOk() (*int32, bool)`

GetCpuBindOk returns a tuple with the CpuBind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuBind

`func (o *V0044UpdateNodeMsg) SetCpuBind(v int32)`

SetCpuBind sets CpuBind field to given value.

### HasCpuBind

`func (o *V0044UpdateNodeMsg) HasCpuBind() bool`

HasCpuBind returns a boolean if a field has been set.

### GetExtra

`func (o *V0044UpdateNodeMsg) GetExtra() string`

GetExtra returns the Extra field if non-nil, zero value otherwise.

### GetExtraOk

`func (o *V0044UpdateNodeMsg) GetExtraOk() (*string, bool)`

GetExtraOk returns a tuple with the Extra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtra

`func (o *V0044UpdateNodeMsg) SetExtra(v string)`

SetExtra sets Extra field to given value.

### HasExtra

`func (o *V0044UpdateNodeMsg) HasExtra() bool`

HasExtra returns a boolean if a field has been set.

### GetFeatures

`func (o *V0044UpdateNodeMsg) GetFeatures() []string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *V0044UpdateNodeMsg) GetFeaturesOk() (*[]string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *V0044UpdateNodeMsg) SetFeatures(v []string)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *V0044UpdateNodeMsg) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetFeaturesAct

`func (o *V0044UpdateNodeMsg) GetFeaturesAct() []string`

GetFeaturesAct returns the FeaturesAct field if non-nil, zero value otherwise.

### GetFeaturesActOk

`func (o *V0044UpdateNodeMsg) GetFeaturesActOk() (*[]string, bool)`

GetFeaturesActOk returns a tuple with the FeaturesAct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeaturesAct

`func (o *V0044UpdateNodeMsg) SetFeaturesAct(v []string)`

SetFeaturesAct sets FeaturesAct field to given value.

### HasFeaturesAct

`func (o *V0044UpdateNodeMsg) HasFeaturesAct() bool`

HasFeaturesAct returns a boolean if a field has been set.

### GetGres

`func (o *V0044UpdateNodeMsg) GetGres() string`

GetGres returns the Gres field if non-nil, zero value otherwise.

### GetGresOk

`func (o *V0044UpdateNodeMsg) GetGresOk() (*string, bool)`

GetGresOk returns a tuple with the Gres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGres

`func (o *V0044UpdateNodeMsg) SetGres(v string)`

SetGres sets Gres field to given value.

### HasGres

`func (o *V0044UpdateNodeMsg) HasGres() bool`

HasGres returns a boolean if a field has been set.

### GetAddress

`func (o *V0044UpdateNodeMsg) GetAddress() []string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *V0044UpdateNodeMsg) GetAddressOk() (*[]string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *V0044UpdateNodeMsg) SetAddress(v []string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *V0044UpdateNodeMsg) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetHostname

`func (o *V0044UpdateNodeMsg) GetHostname() []string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *V0044UpdateNodeMsg) GetHostnameOk() (*[]string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *V0044UpdateNodeMsg) SetHostname(v []string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *V0044UpdateNodeMsg) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetName

`func (o *V0044UpdateNodeMsg) GetName() []string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V0044UpdateNodeMsg) GetNameOk() (*[]string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V0044UpdateNodeMsg) SetName(v []string)`

SetName sets Name field to given value.

### HasName

`func (o *V0044UpdateNodeMsg) HasName() bool`

HasName returns a boolean if a field has been set.

### GetState

`func (o *V0044UpdateNodeMsg) GetState() []string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *V0044UpdateNodeMsg) GetStateOk() (*[]string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *V0044UpdateNodeMsg) SetState(v []string)`

SetState sets State field to given value.

### HasState

`func (o *V0044UpdateNodeMsg) HasState() bool`

HasState returns a boolean if a field has been set.

### GetReason

`func (o *V0044UpdateNodeMsg) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *V0044UpdateNodeMsg) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *V0044UpdateNodeMsg) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *V0044UpdateNodeMsg) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetReasonUid

`func (o *V0044UpdateNodeMsg) GetReasonUid() string`

GetReasonUid returns the ReasonUid field if non-nil, zero value otherwise.

### GetReasonUidOk

`func (o *V0044UpdateNodeMsg) GetReasonUidOk() (*string, bool)`

GetReasonUidOk returns a tuple with the ReasonUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonUid

`func (o *V0044UpdateNodeMsg) SetReasonUid(v string)`

SetReasonUid sets ReasonUid field to given value.

### HasReasonUid

`func (o *V0044UpdateNodeMsg) HasReasonUid() bool`

HasReasonUid returns a boolean if a field has been set.

### GetResumeAfter

`func (o *V0044UpdateNodeMsg) GetResumeAfter() V0044Uint32NoValStruct`

GetResumeAfter returns the ResumeAfter field if non-nil, zero value otherwise.

### GetResumeAfterOk

`func (o *V0044UpdateNodeMsg) GetResumeAfterOk() (*V0044Uint32NoValStruct, bool)`

GetResumeAfterOk returns a tuple with the ResumeAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResumeAfter

`func (o *V0044UpdateNodeMsg) SetResumeAfter(v V0044Uint32NoValStruct)`

SetResumeAfter sets ResumeAfter field to given value.

### HasResumeAfter

`func (o *V0044UpdateNodeMsg) HasResumeAfter() bool`

HasResumeAfter returns a boolean if a field has been set.

### GetTopologyStr

`func (o *V0044UpdateNodeMsg) GetTopologyStr() string`

GetTopologyStr returns the TopologyStr field if non-nil, zero value otherwise.

### GetTopologyStrOk

`func (o *V0044UpdateNodeMsg) GetTopologyStrOk() (*string, bool)`

GetTopologyStrOk returns a tuple with the TopologyStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopologyStr

`func (o *V0044UpdateNodeMsg) SetTopologyStr(v string)`

SetTopologyStr sets TopologyStr field to given value.

### HasTopologyStr

`func (o *V0044UpdateNodeMsg) HasTopologyStr() bool`

HasTopologyStr returns a boolean if a field has been set.

### GetWeight

`func (o *V0044UpdateNodeMsg) GetWeight() V0044Uint32NoValStruct`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *V0044UpdateNodeMsg) GetWeightOk() (*V0044Uint32NoValStruct, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *V0044UpdateNodeMsg) SetWeight(v V0044Uint32NoValStruct)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *V0044UpdateNodeMsg) HasWeight() bool`

HasWeight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


