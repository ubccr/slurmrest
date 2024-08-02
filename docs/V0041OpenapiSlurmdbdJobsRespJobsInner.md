# V0041OpenapiSlurmdbdJobsRespJobsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** |  | [optional] 
**Comment** | Pointer to [**V0041OpenapiSlurmdbdJobsRespJobsInnerComment**](V0041OpenapiSlurmdbdJobsRespJobsInnerComment.md) |  | [optional] 
**AllocationNodes** | Pointer to **int32** |  | [optional] 
**Array** | Pointer to [**V0041OpenapiSlurmdbdJobsRespJobsInnerArray**](V0041OpenapiSlurmdbdJobsRespJobsInnerArray.md) |  | [optional] 
**Association** | Pointer to [**V0041OpenapiSlurmdbdJobsRespJobsInnerAssociation**](V0041OpenapiSlurmdbdJobsRespJobsInnerAssociation.md) |  | [optional] 
**Block** | Pointer to **string** |  | [optional] 
**Cluster** | Pointer to **string** |  | [optional] 
**Constraints** | Pointer to **string** |  | [optional] 
**Container** | Pointer to **string** |  | [optional] 
**DerivedExitCode** | Pointer to [**V0041OpenapiSlurmdbdJobsRespJobsInnerDerivedExitCode**](V0041OpenapiSlurmdbdJobsRespJobsInnerDerivedExitCode.md) |  | [optional] 
**Time** | Pointer to [**V0041OpenapiSlurmdbdJobsRespJobsInnerTime**](V0041OpenapiSlurmdbdJobsRespJobsInnerTime.md) |  | [optional] 
**ExitCode** | Pointer to [**V0041OpenapiSlurmdbdJobsRespJobsInnerDerivedExitCode**](V0041OpenapiSlurmdbdJobsRespJobsInnerDerivedExitCode.md) |  | [optional] 
**Extra** | Pointer to **string** |  | [optional] 
**FailedNode** | Pointer to **string** |  | [optional] 
**Flags** | Pointer to **[]string** |  | [optional] 
**Group** | Pointer to **string** |  | [optional] 
**Het** | Pointer to [**V0041OpenapiSlurmdbdJobsRespJobsInnerHet**](V0041OpenapiSlurmdbdJobsRespJobsInnerHet.md) |  | [optional] 
**JobId** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Licenses** | Pointer to **string** |  | [optional] 
**Mcs** | Pointer to [**V0041OpenapiSlurmdbdJobsRespJobsInnerMcs**](V0041OpenapiSlurmdbdJobsRespJobsInnerMcs.md) |  | [optional] 
**Nodes** | Pointer to **string** |  | [optional] 
**Partition** | Pointer to **string** |  | [optional] 
**Hold** | Pointer to **bool** | Hold (true) or release (false) job | [optional] 
**Priority** | Pointer to [**V0041OpenapiSlurmdbdJobsRespJobsInnerPriority**](V0041OpenapiSlurmdbdJobsRespJobsInnerPriority.md) |  | [optional] 
**Qos** | Pointer to **string** |  | [optional] 
**Required** | Pointer to [**V0041OpenapiSlurmdbdJobsRespJobsInnerRequired**](V0041OpenapiSlurmdbdJobsRespJobsInnerRequired.md) |  | [optional] 
**KillRequestUser** | Pointer to **string** |  | [optional] 
**Reservation** | Pointer to [**V0041OpenapiSlurmdbdJobsRespJobsInnerReservation**](V0041OpenapiSlurmdbdJobsRespJobsInnerReservation.md) |  | [optional] 
**Script** | Pointer to **string** |  | [optional] 
**StdinExpanded** | Pointer to **string** | Job stdin with expanded fields | [optional] 
**StdoutExpanded** | Pointer to **string** | Job stdout with expanded fields | [optional] 
**StderrExpanded** | Pointer to **string** | Job stderr with expanded fields | [optional] 
**Stdout** | Pointer to **string** |  | [optional] 
**Stderr** | Pointer to **string** |  | [optional] 
**Stdin** | Pointer to **string** |  | [optional] 
**State** | Pointer to [**V0041OpenapiSlurmdbdJobsRespJobsInnerState**](V0041OpenapiSlurmdbdJobsRespJobsInnerState.md) |  | [optional] 
**Steps** | Pointer to [**[]V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInner**](V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInner.md) |  | [optional] 
**SubmitLine** | Pointer to **string** |  | [optional] 
**Tres** | Pointer to [**V0041OpenapiSlurmdbdJobsRespJobsInnerTres**](V0041OpenapiSlurmdbdJobsRespJobsInnerTres.md) |  | [optional] 
**UsedGres** | Pointer to **string** |  | [optional] 
**User** | Pointer to **string** |  | [optional] 
**Wckey** | Pointer to [**V0041OpenapiSlurmdbdJobsRespJobsInnerWckey**](V0041OpenapiSlurmdbdJobsRespJobsInnerWckey.md) |  | [optional] 
**WorkingDirectory** | Pointer to **string** |  | [optional] 

## Methods

### NewV0041OpenapiSlurmdbdJobsRespJobsInner

`func NewV0041OpenapiSlurmdbdJobsRespJobsInner() *V0041OpenapiSlurmdbdJobsRespJobsInner`

NewV0041OpenapiSlurmdbdJobsRespJobsInner instantiates a new V0041OpenapiSlurmdbdJobsRespJobsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0041OpenapiSlurmdbdJobsRespJobsInnerWithDefaults

`func NewV0041OpenapiSlurmdbdJobsRespJobsInnerWithDefaults() *V0041OpenapiSlurmdbdJobsRespJobsInner`

NewV0041OpenapiSlurmdbdJobsRespJobsInnerWithDefaults instantiates a new V0041OpenapiSlurmdbdJobsRespJobsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetComment

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) GetComment() V0041OpenapiSlurmdbdJobsRespJobsInnerComment`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) GetCommentOk() (*V0041OpenapiSlurmdbdJobsRespJobsInnerComment, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) SetComment(v V0041OpenapiSlurmdbdJobsRespJobsInnerComment)`

SetComment sets Comment field to given value.

### HasComment

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetAllocationNodes

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) GetAllocationNodes() int32`

GetAllocationNodes returns the AllocationNodes field if non-nil, zero value otherwise.

### GetAllocationNodesOk

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) GetAllocationNodesOk() (*int32, bool)`

GetAllocationNodesOk returns a tuple with the AllocationNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocationNodes

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) SetAllocationNodes(v int32)`

SetAllocationNodes sets AllocationNodes field to given value.

### HasAllocationNodes

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) HasAllocationNodes() bool`

HasAllocationNodes returns a boolean if a field has been set.

### GetArray

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) GetArray() V0041OpenapiSlurmdbdJobsRespJobsInnerArray`

GetArray returns the Array field if non-nil, zero value otherwise.

### GetArrayOk

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) GetArrayOk() (*V0041OpenapiSlurmdbdJobsRespJobsInnerArray, bool)`

GetArrayOk returns a tuple with the Array field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArray

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) SetArray(v V0041OpenapiSlurmdbdJobsRespJobsInnerArray)`

SetArray sets Array field to given value.

### HasArray

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) HasArray() bool`

HasArray returns a boolean if a field has been set.

### GetAssociation

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) GetAssociation() V0041OpenapiSlurmdbdJobsRespJobsInnerAssociation`

GetAssociation returns the Association field if non-nil, zero value otherwise.

### GetAssociationOk

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) GetAssociationOk() (*V0041OpenapiSlurmdbdJobsRespJobsInnerAssociation, bool)`

GetAssociationOk returns a tuple with the Association field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociation

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) SetAssociation(v V0041OpenapiSlurmdbdJobsRespJobsInnerAssociation)`

SetAssociation sets Association field to given value.

### HasAssociation

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) HasAssociation() bool`

HasAssociation returns a boolean if a field has been set.

### GetBlock

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) GetBlock() string`

GetBlock returns the Block field if non-nil, zero value otherwise.

### GetBlockOk

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) GetBlockOk() (*string, bool)`

GetBlockOk returns a tuple with the Block field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlock

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) SetBlock(v string)`

SetBlock sets Block field to given value.

### HasBlock

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) HasBlock() bool`

HasBlock returns a boolean if a field has been set.

### GetCluster

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetConstraints

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) GetConstraints() string`

GetConstraints returns the Constraints field if non-nil, zero value otherwise.

### GetConstraintsOk

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) GetConstraintsOk() (*string, bool)`

GetConstraintsOk returns a tuple with the Constraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraints

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) SetConstraints(v string)`

SetConstraints sets Constraints field to given value.

### HasConstraints

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) HasConstraints() bool`

HasConstraints returns a boolean if a field has been set.

### GetContainer

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) GetContainer() string`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) GetContainerOk() (*string, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) SetContainer(v string)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### GetDerivedExitCode

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) GetDerivedExitCode() V0041OpenapiSlurmdbdJobsRespJobsInnerDerivedExitCode`

GetDerivedExitCode returns the DerivedExitCode field if non-nil, zero value otherwise.

### GetDerivedExitCodeOk

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) GetDerivedExitCodeOk() (*V0041OpenapiSlurmdbdJobsRespJobsInnerDerivedExitCode, bool)`

GetDerivedExitCodeOk returns a tuple with the DerivedExitCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerivedExitCode

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) SetDerivedExitCode(v V0041OpenapiSlurmdbdJobsRespJobsInnerDerivedExitCode)`

SetDerivedExitCode sets DerivedExitCode field to given value.

### HasDerivedExitCode

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) HasDerivedExitCode() bool`

HasDerivedExitCode returns a boolean if a field has been set.

### GetTime

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) GetTime() V0041OpenapiSlurmdbdJobsRespJobsInnerTime`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) GetTimeOk() (*V0041OpenapiSlurmdbdJobsRespJobsInnerTime, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) SetTime(v V0041OpenapiSlurmdbdJobsRespJobsInnerTime)`

SetTime sets Time field to given value.

### HasTime

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetExitCode

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) GetExitCode() V0041OpenapiSlurmdbdJobsRespJobsInnerDerivedExitCode`

GetExitCode returns the ExitCode field if non-nil, zero value otherwise.

### GetExitCodeOk

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) GetExitCodeOk() (*V0041OpenapiSlurmdbdJobsRespJobsInnerDerivedExitCode, bool)`

GetExitCodeOk returns a tuple with the ExitCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitCode

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) SetExitCode(v V0041OpenapiSlurmdbdJobsRespJobsInnerDerivedExitCode)`

SetExitCode sets ExitCode field to given value.

### HasExitCode

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) HasExitCode() bool`

HasExitCode returns a boolean if a field has been set.

### GetExtra

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) GetExtra() string`

GetExtra returns the Extra field if non-nil, zero value otherwise.

### GetExtraOk

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) GetExtraOk() (*string, bool)`

GetExtraOk returns a tuple with the Extra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtra

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) SetExtra(v string)`

SetExtra sets Extra field to given value.

### HasExtra

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) HasExtra() bool`

HasExtra returns a boolean if a field has been set.

### GetFailedNode

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) GetFailedNode() string`

GetFailedNode returns the FailedNode field if non-nil, zero value otherwise.

### GetFailedNodeOk

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) GetFailedNodeOk() (*string, bool)`

GetFailedNodeOk returns a tuple with the FailedNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedNode

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) SetFailedNode(v string)`

SetFailedNode sets FailedNode field to given value.

### HasFailedNode

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) HasFailedNode() bool`

HasFailedNode returns a boolean if a field has been set.

### GetFlags

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) GetFlags() []string`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) GetFlagsOk() (*[]string, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) SetFlags(v []string)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetGroup

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetHet

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) GetHet() V0041OpenapiSlurmdbdJobsRespJobsInnerHet`

GetHet returns the Het field if non-nil, zero value otherwise.

### GetHetOk

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) GetHetOk() (*V0041OpenapiSlurmdbdJobsRespJobsInnerHet, bool)`

GetHetOk returns a tuple with the Het field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHet

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) SetHet(v V0041OpenapiSlurmdbdJobsRespJobsInnerHet)`

SetHet sets Het field to given value.

### HasHet

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) HasHet() bool`

HasHet returns a boolean if a field has been set.

### GetJobId

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) GetJobId() int32`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) GetJobIdOk() (*int32, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) SetJobId(v int32)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetName

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLicenses

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) GetLicenses() string`

GetLicenses returns the Licenses field if non-nil, zero value otherwise.

### GetLicensesOk

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) GetLicensesOk() (*string, bool)`

GetLicensesOk returns a tuple with the Licenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenses

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) SetLicenses(v string)`

SetLicenses sets Licenses field to given value.

### HasLicenses

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) HasLicenses() bool`

HasLicenses returns a boolean if a field has been set.

### GetMcs

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) GetMcs() V0041OpenapiSlurmdbdJobsRespJobsInnerMcs`

GetMcs returns the Mcs field if non-nil, zero value otherwise.

### GetMcsOk

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) GetMcsOk() (*V0041OpenapiSlurmdbdJobsRespJobsInnerMcs, bool)`

GetMcsOk returns a tuple with the Mcs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcs

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) SetMcs(v V0041OpenapiSlurmdbdJobsRespJobsInnerMcs)`

SetMcs sets Mcs field to given value.

### HasMcs

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) HasMcs() bool`

HasMcs returns a boolean if a field has been set.

### GetNodes

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) GetNodes() string`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) GetNodesOk() (*string, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) SetNodes(v string)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetPartition

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) GetPartition() string`

GetPartition returns the Partition field if non-nil, zero value otherwise.

### GetPartitionOk

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) GetPartitionOk() (*string, bool)`

GetPartitionOk returns a tuple with the Partition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartition

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) SetPartition(v string)`

SetPartition sets Partition field to given value.

### HasPartition

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) HasPartition() bool`

HasPartition returns a boolean if a field has been set.

### GetHold

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) GetHold() bool`

GetHold returns the Hold field if non-nil, zero value otherwise.

### GetHoldOk

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) GetHoldOk() (*bool, bool)`

GetHoldOk returns a tuple with the Hold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHold

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) SetHold(v bool)`

SetHold sets Hold field to given value.

### HasHold

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) HasHold() bool`

HasHold returns a boolean if a field has been set.

### GetPriority

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) GetPriority() V0041OpenapiSlurmdbdJobsRespJobsInnerPriority`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) GetPriorityOk() (*V0041OpenapiSlurmdbdJobsRespJobsInnerPriority, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) SetPriority(v V0041OpenapiSlurmdbdJobsRespJobsInnerPriority)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetQos

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) GetQos() string`

GetQos returns the Qos field if non-nil, zero value otherwise.

### GetQosOk

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) GetQosOk() (*string, bool)`

GetQosOk returns a tuple with the Qos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQos

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) SetQos(v string)`

SetQos sets Qos field to given value.

### HasQos

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) HasQos() bool`

HasQos returns a boolean if a field has been set.

### GetRequired

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) GetRequired() V0041OpenapiSlurmdbdJobsRespJobsInnerRequired`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) GetRequiredOk() (*V0041OpenapiSlurmdbdJobsRespJobsInnerRequired, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) SetRequired(v V0041OpenapiSlurmdbdJobsRespJobsInnerRequired)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetKillRequestUser

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) GetKillRequestUser() string`

GetKillRequestUser returns the KillRequestUser field if non-nil, zero value otherwise.

### GetKillRequestUserOk

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) GetKillRequestUserOk() (*string, bool)`

GetKillRequestUserOk returns a tuple with the KillRequestUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKillRequestUser

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) SetKillRequestUser(v string)`

SetKillRequestUser sets KillRequestUser field to given value.

### HasKillRequestUser

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) HasKillRequestUser() bool`

HasKillRequestUser returns a boolean if a field has been set.

### GetReservation

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) GetReservation() V0041OpenapiSlurmdbdJobsRespJobsInnerReservation`

GetReservation returns the Reservation field if non-nil, zero value otherwise.

### GetReservationOk

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) GetReservationOk() (*V0041OpenapiSlurmdbdJobsRespJobsInnerReservation, bool)`

GetReservationOk returns a tuple with the Reservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservation

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) SetReservation(v V0041OpenapiSlurmdbdJobsRespJobsInnerReservation)`

SetReservation sets Reservation field to given value.

### HasReservation

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) HasReservation() bool`

HasReservation returns a boolean if a field has been set.

### GetScript

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) GetScript() string`

GetScript returns the Script field if non-nil, zero value otherwise.

### GetScriptOk

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) GetScriptOk() (*string, bool)`

GetScriptOk returns a tuple with the Script field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScript

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) SetScript(v string)`

SetScript sets Script field to given value.

### HasScript

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) HasScript() bool`

HasScript returns a boolean if a field has been set.

### GetStdinExpanded

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) GetStdinExpanded() string`

GetStdinExpanded returns the StdinExpanded field if non-nil, zero value otherwise.

### GetStdinExpandedOk

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) GetStdinExpandedOk() (*string, bool)`

GetStdinExpandedOk returns a tuple with the StdinExpanded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStdinExpanded

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) SetStdinExpanded(v string)`

SetStdinExpanded sets StdinExpanded field to given value.

### HasStdinExpanded

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) HasStdinExpanded() bool`

HasStdinExpanded returns a boolean if a field has been set.

### GetStdoutExpanded

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) GetStdoutExpanded() string`

GetStdoutExpanded returns the StdoutExpanded field if non-nil, zero value otherwise.

### GetStdoutExpandedOk

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) GetStdoutExpandedOk() (*string, bool)`

GetStdoutExpandedOk returns a tuple with the StdoutExpanded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStdoutExpanded

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) SetStdoutExpanded(v string)`

SetStdoutExpanded sets StdoutExpanded field to given value.

### HasStdoutExpanded

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) HasStdoutExpanded() bool`

HasStdoutExpanded returns a boolean if a field has been set.

### GetStderrExpanded

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) GetStderrExpanded() string`

GetStderrExpanded returns the StderrExpanded field if non-nil, zero value otherwise.

### GetStderrExpandedOk

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) GetStderrExpandedOk() (*string, bool)`

GetStderrExpandedOk returns a tuple with the StderrExpanded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStderrExpanded

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) SetStderrExpanded(v string)`

SetStderrExpanded sets StderrExpanded field to given value.

### HasStderrExpanded

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) HasStderrExpanded() bool`

HasStderrExpanded returns a boolean if a field has been set.

### GetStdout

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) GetStdout() string`

GetStdout returns the Stdout field if non-nil, zero value otherwise.

### GetStdoutOk

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) GetStdoutOk() (*string, bool)`

GetStdoutOk returns a tuple with the Stdout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStdout

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) SetStdout(v string)`

SetStdout sets Stdout field to given value.

### HasStdout

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) HasStdout() bool`

HasStdout returns a boolean if a field has been set.

### GetStderr

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) GetStderr() string`

GetStderr returns the Stderr field if non-nil, zero value otherwise.

### GetStderrOk

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) GetStderrOk() (*string, bool)`

GetStderrOk returns a tuple with the Stderr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStderr

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) SetStderr(v string)`

SetStderr sets Stderr field to given value.

### HasStderr

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) HasStderr() bool`

HasStderr returns a boolean if a field has been set.

### GetStdin

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) GetStdin() string`

GetStdin returns the Stdin field if non-nil, zero value otherwise.

### GetStdinOk

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) GetStdinOk() (*string, bool)`

GetStdinOk returns a tuple with the Stdin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStdin

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) SetStdin(v string)`

SetStdin sets Stdin field to given value.

### HasStdin

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) HasStdin() bool`

HasStdin returns a boolean if a field has been set.

### GetState

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) GetState() V0041OpenapiSlurmdbdJobsRespJobsInnerState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) GetStateOk() (*V0041OpenapiSlurmdbdJobsRespJobsInnerState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) SetState(v V0041OpenapiSlurmdbdJobsRespJobsInnerState)`

SetState sets State field to given value.

### HasState

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSteps

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) GetSteps() []V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInner`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) GetStepsOk() (*[]V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInner, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) SetSteps(v []V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInner)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) HasSteps() bool`

HasSteps returns a boolean if a field has been set.

### GetSubmitLine

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) GetSubmitLine() string`

GetSubmitLine returns the SubmitLine field if non-nil, zero value otherwise.

### GetSubmitLineOk

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) GetSubmitLineOk() (*string, bool)`

GetSubmitLineOk returns a tuple with the SubmitLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmitLine

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) SetSubmitLine(v string)`

SetSubmitLine sets SubmitLine field to given value.

### HasSubmitLine

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) HasSubmitLine() bool`

HasSubmitLine returns a boolean if a field has been set.

### GetTres

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) GetTres() V0041OpenapiSlurmdbdJobsRespJobsInnerTres`

GetTres returns the Tres field if non-nil, zero value otherwise.

### GetTresOk

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) GetTresOk() (*V0041OpenapiSlurmdbdJobsRespJobsInnerTres, bool)`

GetTresOk returns a tuple with the Tres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTres

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) SetTres(v V0041OpenapiSlurmdbdJobsRespJobsInnerTres)`

SetTres sets Tres field to given value.

### HasTres

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) HasTres() bool`

HasTres returns a boolean if a field has been set.

### GetUsedGres

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) GetUsedGres() string`

GetUsedGres returns the UsedGres field if non-nil, zero value otherwise.

### GetUsedGresOk

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) GetUsedGresOk() (*string, bool)`

GetUsedGresOk returns a tuple with the UsedGres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedGres

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) SetUsedGres(v string)`

SetUsedGres sets UsedGres field to given value.

### HasUsedGres

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) HasUsedGres() bool`

HasUsedGres returns a boolean if a field has been set.

### GetUser

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetWckey

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) GetWckey() V0041OpenapiSlurmdbdJobsRespJobsInnerWckey`

GetWckey returns the Wckey field if non-nil, zero value otherwise.

### GetWckeyOk

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) GetWckeyOk() (*V0041OpenapiSlurmdbdJobsRespJobsInnerWckey, bool)`

GetWckeyOk returns a tuple with the Wckey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWckey

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) SetWckey(v V0041OpenapiSlurmdbdJobsRespJobsInnerWckey)`

SetWckey sets Wckey field to given value.

### HasWckey

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) HasWckey() bool`

HasWckey returns a boolean if a field has been set.

### GetWorkingDirectory

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) GetWorkingDirectory() string`

GetWorkingDirectory returns the WorkingDirectory field if non-nil, zero value otherwise.

### GetWorkingDirectoryOk

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) GetWorkingDirectoryOk() (*string, bool)`

GetWorkingDirectoryOk returns a tuple with the WorkingDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingDirectory

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) SetWorkingDirectory(v string)`

SetWorkingDirectory sets WorkingDirectory field to given value.

### HasWorkingDirectory

`func (o *V0041OpenapiSlurmdbdJobsRespJobsInner) HasWorkingDirectory() bool`

HasWorkingDirectory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


