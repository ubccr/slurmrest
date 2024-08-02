# V0041JobDescMsgRlimits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cpu** | Pointer to [**V0041JobDescMsgRlimitsCpu**](V0041JobDescMsgRlimitsCpu.md) |  | [optional] 
**Fsize** | Pointer to [**V0041JobDescMsgRlimitsFsize**](V0041JobDescMsgRlimitsFsize.md) |  | [optional] 
**Data** | Pointer to [**V0041JobDescMsgRlimitsData**](V0041JobDescMsgRlimitsData.md) |  | [optional] 
**Stack** | Pointer to [**V0041JobDescMsgRlimitsStack**](V0041JobDescMsgRlimitsStack.md) |  | [optional] 
**Core** | Pointer to [**V0041JobDescMsgRlimitsCore**](V0041JobDescMsgRlimitsCore.md) |  | [optional] 
**Rss** | Pointer to [**V0041JobDescMsgRlimitsRss**](V0041JobDescMsgRlimitsRss.md) |  | [optional] 
**Nproc** | Pointer to [**V0041JobDescMsgRlimitsNproc**](V0041JobDescMsgRlimitsNproc.md) |  | [optional] 
**Nofile** | Pointer to [**V0041JobDescMsgRlimitsNofile**](V0041JobDescMsgRlimitsNofile.md) |  | [optional] 
**Memlock** | Pointer to [**V0041JobDescMsgRlimitsMemlock**](V0041JobDescMsgRlimitsMemlock.md) |  | [optional] 
**As** | Pointer to [**V0041JobDescMsgRlimitsAs**](V0041JobDescMsgRlimitsAs.md) |  | [optional] 

## Methods

### NewV0041JobDescMsgRlimits

`func NewV0041JobDescMsgRlimits() *V0041JobDescMsgRlimits`

NewV0041JobDescMsgRlimits instantiates a new V0041JobDescMsgRlimits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0041JobDescMsgRlimitsWithDefaults

`func NewV0041JobDescMsgRlimitsWithDefaults() *V0041JobDescMsgRlimits`

NewV0041JobDescMsgRlimitsWithDefaults instantiates a new V0041JobDescMsgRlimits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpu

`func (o *V0041JobDescMsgRlimits) GetCpu() V0041JobDescMsgRlimitsCpu`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *V0041JobDescMsgRlimits) GetCpuOk() (*V0041JobDescMsgRlimitsCpu, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *V0041JobDescMsgRlimits) SetCpu(v V0041JobDescMsgRlimitsCpu)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *V0041JobDescMsgRlimits) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetFsize

`func (o *V0041JobDescMsgRlimits) GetFsize() V0041JobDescMsgRlimitsFsize`

GetFsize returns the Fsize field if non-nil, zero value otherwise.

### GetFsizeOk

`func (o *V0041JobDescMsgRlimits) GetFsizeOk() (*V0041JobDescMsgRlimitsFsize, bool)`

GetFsizeOk returns a tuple with the Fsize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsize

`func (o *V0041JobDescMsgRlimits) SetFsize(v V0041JobDescMsgRlimitsFsize)`

SetFsize sets Fsize field to given value.

### HasFsize

`func (o *V0041JobDescMsgRlimits) HasFsize() bool`

HasFsize returns a boolean if a field has been set.

### GetData

`func (o *V0041JobDescMsgRlimits) GetData() V0041JobDescMsgRlimitsData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *V0041JobDescMsgRlimits) GetDataOk() (*V0041JobDescMsgRlimitsData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *V0041JobDescMsgRlimits) SetData(v V0041JobDescMsgRlimitsData)`

SetData sets Data field to given value.

### HasData

`func (o *V0041JobDescMsgRlimits) HasData() bool`

HasData returns a boolean if a field has been set.

### GetStack

`func (o *V0041JobDescMsgRlimits) GetStack() V0041JobDescMsgRlimitsStack`

GetStack returns the Stack field if non-nil, zero value otherwise.

### GetStackOk

`func (o *V0041JobDescMsgRlimits) GetStackOk() (*V0041JobDescMsgRlimitsStack, bool)`

GetStackOk returns a tuple with the Stack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStack

`func (o *V0041JobDescMsgRlimits) SetStack(v V0041JobDescMsgRlimitsStack)`

SetStack sets Stack field to given value.

### HasStack

`func (o *V0041JobDescMsgRlimits) HasStack() bool`

HasStack returns a boolean if a field has been set.

### GetCore

`func (o *V0041JobDescMsgRlimits) GetCore() V0041JobDescMsgRlimitsCore`

GetCore returns the Core field if non-nil, zero value otherwise.

### GetCoreOk

`func (o *V0041JobDescMsgRlimits) GetCoreOk() (*V0041JobDescMsgRlimitsCore, bool)`

GetCoreOk returns a tuple with the Core field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCore

`func (o *V0041JobDescMsgRlimits) SetCore(v V0041JobDescMsgRlimitsCore)`

SetCore sets Core field to given value.

### HasCore

`func (o *V0041JobDescMsgRlimits) HasCore() bool`

HasCore returns a boolean if a field has been set.

### GetRss

`func (o *V0041JobDescMsgRlimits) GetRss() V0041JobDescMsgRlimitsRss`

GetRss returns the Rss field if non-nil, zero value otherwise.

### GetRssOk

`func (o *V0041JobDescMsgRlimits) GetRssOk() (*V0041JobDescMsgRlimitsRss, bool)`

GetRssOk returns a tuple with the Rss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRss

`func (o *V0041JobDescMsgRlimits) SetRss(v V0041JobDescMsgRlimitsRss)`

SetRss sets Rss field to given value.

### HasRss

`func (o *V0041JobDescMsgRlimits) HasRss() bool`

HasRss returns a boolean if a field has been set.

### GetNproc

`func (o *V0041JobDescMsgRlimits) GetNproc() V0041JobDescMsgRlimitsNproc`

GetNproc returns the Nproc field if non-nil, zero value otherwise.

### GetNprocOk

`func (o *V0041JobDescMsgRlimits) GetNprocOk() (*V0041JobDescMsgRlimitsNproc, bool)`

GetNprocOk returns a tuple with the Nproc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNproc

`func (o *V0041JobDescMsgRlimits) SetNproc(v V0041JobDescMsgRlimitsNproc)`

SetNproc sets Nproc field to given value.

### HasNproc

`func (o *V0041JobDescMsgRlimits) HasNproc() bool`

HasNproc returns a boolean if a field has been set.

### GetNofile

`func (o *V0041JobDescMsgRlimits) GetNofile() V0041JobDescMsgRlimitsNofile`

GetNofile returns the Nofile field if non-nil, zero value otherwise.

### GetNofileOk

`func (o *V0041JobDescMsgRlimits) GetNofileOk() (*V0041JobDescMsgRlimitsNofile, bool)`

GetNofileOk returns a tuple with the Nofile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNofile

`func (o *V0041JobDescMsgRlimits) SetNofile(v V0041JobDescMsgRlimitsNofile)`

SetNofile sets Nofile field to given value.

### HasNofile

`func (o *V0041JobDescMsgRlimits) HasNofile() bool`

HasNofile returns a boolean if a field has been set.

### GetMemlock

`func (o *V0041JobDescMsgRlimits) GetMemlock() V0041JobDescMsgRlimitsMemlock`

GetMemlock returns the Memlock field if non-nil, zero value otherwise.

### GetMemlockOk

`func (o *V0041JobDescMsgRlimits) GetMemlockOk() (*V0041JobDescMsgRlimitsMemlock, bool)`

GetMemlockOk returns a tuple with the Memlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemlock

`func (o *V0041JobDescMsgRlimits) SetMemlock(v V0041JobDescMsgRlimitsMemlock)`

SetMemlock sets Memlock field to given value.

### HasMemlock

`func (o *V0041JobDescMsgRlimits) HasMemlock() bool`

HasMemlock returns a boolean if a field has been set.

### GetAs

`func (o *V0041JobDescMsgRlimits) GetAs() V0041JobDescMsgRlimitsAs`

GetAs returns the As field if non-nil, zero value otherwise.

### GetAsOk

`func (o *V0041JobDescMsgRlimits) GetAsOk() (*V0041JobDescMsgRlimitsAs, bool)`

GetAsOk returns a tuple with the As field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAs

`func (o *V0041JobDescMsgRlimits) SetAs(v V0041JobDescMsgRlimitsAs)`

SetAs sets As field to given value.

### HasAs

`func (o *V0041JobDescMsgRlimits) HasAs() bool`

HasAs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


