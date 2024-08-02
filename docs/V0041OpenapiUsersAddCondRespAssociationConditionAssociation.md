# V0041OpenapiUsersAddCondRespAssociationConditionAssociation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | Comment for the association | [optional] 
**Defaultqos** | Pointer to **string** | Which QOS id is this association default | [optional] 
**Grpjobs** | Pointer to [**V0041OpenapiUsersAddCondRespAssociationConditionAssociationGrpjobs**](V0041OpenapiUsersAddCondRespAssociationConditionAssociationGrpjobs.md) |  | [optional] 
**Grpjobsaccrue** | Pointer to [**V0041OpenapiUsersAddCondRespAssociationConditionAssociationGrpjobsaccrue**](V0041OpenapiUsersAddCondRespAssociationConditionAssociationGrpjobsaccrue.md) |  | [optional] 
**Grpsubmitjobs** | Pointer to [**V0041OpenapiUsersAddCondRespAssociationConditionAssociationGrpsubmitjobs**](V0041OpenapiUsersAddCondRespAssociationConditionAssociationGrpsubmitjobs.md) |  | [optional] 
**Grptres** | Pointer to [**[]V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTresRequestedMaxInner**](V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTresRequestedMaxInner.md) |  | [optional] 
**Grptresmins** | Pointer to [**[]V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTresRequestedMaxInner**](V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTresRequestedMaxInner.md) | Max number of cpu minutes the underlying group of associations can run for | [optional] 
**Grptresrunmins** | Pointer to [**[]V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTresRequestedMaxInner**](V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTresRequestedMaxInner.md) | Max number of cpu minutes the underlying group of associations can having running at one time | [optional] 
**Grpwall** | Pointer to [**V0041OpenapiUsersAddCondRespAssociationConditionAssociationGrpwall**](V0041OpenapiUsersAddCondRespAssociationConditionAssociationGrpwall.md) |  | [optional] 
**Maxjobs** | Pointer to [**V0041OpenapiUsersAddCondRespAssociationConditionAssociationMaxjobs**](V0041OpenapiUsersAddCondRespAssociationConditionAssociationMaxjobs.md) |  | [optional] 
**Maxjobsaccrue** | Pointer to [**V0041OpenapiUsersAddCondRespAssociationConditionAssociationMaxjobsaccrue**](V0041OpenapiUsersAddCondRespAssociationConditionAssociationMaxjobsaccrue.md) |  | [optional] 
**Maxsubmitjobs** | Pointer to [**V0041OpenapiUsersAddCondRespAssociationConditionAssociationMaxsubmitjobs**](V0041OpenapiUsersAddCondRespAssociationConditionAssociationMaxsubmitjobs.md) |  | [optional] 
**Maxtresminsperjob** | Pointer to [**[]V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTresRequestedMaxInner**](V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTresRequestedMaxInner.md) | Max number of cpu minutes this association can have per job | [optional] 
**Maxtresrunmins** | Pointer to [**[]V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTresRequestedMaxInner**](V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTresRequestedMaxInner.md) | Max number of cpu minutes this association can having running at one time | [optional] 
**Maxtresperjob** | Pointer to [**[]V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTresRequestedMaxInner**](V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTresRequestedMaxInner.md) | Max number of cpus this association can allocate per job | [optional] 
**Maxtrespernode** | Pointer to [**[]V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTresRequestedMaxInner**](V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTresRequestedMaxInner.md) | Max number of TRES this association can allocate per node | [optional] 
**Maxwalldurationperjob** | Pointer to [**V0041OpenapiUsersAddCondRespAssociationConditionAssociationMaxwalldurationperjob**](V0041OpenapiUsersAddCondRespAssociationConditionAssociationMaxwalldurationperjob.md) |  | [optional] 
**Minpriothresh** | Pointer to [**V0041OpenapiUsersAddCondRespAssociationConditionAssociationMinpriothresh**](V0041OpenapiUsersAddCondRespAssociationConditionAssociationMinpriothresh.md) |  | [optional] 
**Parent** | Pointer to **string** | Name of parent account | [optional] 
**Priority** | Pointer to [**V0041OpenapiUsersAddCondRespAssociationConditionAssociationPriority**](V0041OpenapiUsersAddCondRespAssociationConditionAssociationPriority.md) |  | [optional] 
**Qoslevel** | Pointer to **[]string** | Default QoS&#39; that jobs are able to run at for this association | [optional] 
**Fairshare** | Pointer to **int32** | Number of shares allocated to this association | [optional] 

## Methods

### NewV0041OpenapiUsersAddCondRespAssociationConditionAssociation

`func NewV0041OpenapiUsersAddCondRespAssociationConditionAssociation() *V0041OpenapiUsersAddCondRespAssociationConditionAssociation`

NewV0041OpenapiUsersAddCondRespAssociationConditionAssociation instantiates a new V0041OpenapiUsersAddCondRespAssociationConditionAssociation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0041OpenapiUsersAddCondRespAssociationConditionAssociationWithDefaults

`func NewV0041OpenapiUsersAddCondRespAssociationConditionAssociationWithDefaults() *V0041OpenapiUsersAddCondRespAssociationConditionAssociation`

NewV0041OpenapiUsersAddCondRespAssociationConditionAssociationWithDefaults instantiates a new V0041OpenapiUsersAddCondRespAssociationConditionAssociation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *V0041OpenapiUsersAddCondRespAssociationConditionAssociation) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *V0041OpenapiUsersAddCondRespAssociationConditionAssociation) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *V0041OpenapiUsersAddCondRespAssociationConditionAssociation) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *V0041OpenapiUsersAddCondRespAssociationConditionAssociation) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDefaultqos

`func (o *V0041OpenapiUsersAddCondRespAssociationConditionAssociation) GetDefaultqos() string`

GetDefaultqos returns the Defaultqos field if non-nil, zero value otherwise.

### GetDefaultqosOk

`func (o *V0041OpenapiUsersAddCondRespAssociationConditionAssociation) GetDefaultqosOk() (*string, bool)`

GetDefaultqosOk returns a tuple with the Defaultqos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultqos

`func (o *V0041OpenapiUsersAddCondRespAssociationConditionAssociation) SetDefaultqos(v string)`

SetDefaultqos sets Defaultqos field to given value.

### HasDefaultqos

`func (o *V0041OpenapiUsersAddCondRespAssociationConditionAssociation) HasDefaultqos() bool`

HasDefaultqos returns a boolean if a field has been set.

### GetGrpjobs

`func (o *V0041OpenapiUsersAddCondRespAssociationConditionAssociation) GetGrpjobs() V0041OpenapiUsersAddCondRespAssociationConditionAssociationGrpjobs`

GetGrpjobs returns the Grpjobs field if non-nil, zero value otherwise.

### GetGrpjobsOk

`func (o *V0041OpenapiUsersAddCondRespAssociationConditionAssociation) GetGrpjobsOk() (*V0041OpenapiUsersAddCondRespAssociationConditionAssociationGrpjobs, bool)`

GetGrpjobsOk returns a tuple with the Grpjobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrpjobs

`func (o *V0041OpenapiUsersAddCondRespAssociationConditionAssociation) SetGrpjobs(v V0041OpenapiUsersAddCondRespAssociationConditionAssociationGrpjobs)`

SetGrpjobs sets Grpjobs field to given value.

### HasGrpjobs

`func (o *V0041OpenapiUsersAddCondRespAssociationConditionAssociation) HasGrpjobs() bool`

HasGrpjobs returns a boolean if a field has been set.

### GetGrpjobsaccrue

`func (o *V0041OpenapiUsersAddCondRespAssociationConditionAssociation) GetGrpjobsaccrue() V0041OpenapiUsersAddCondRespAssociationConditionAssociationGrpjobsaccrue`

GetGrpjobsaccrue returns the Grpjobsaccrue field if non-nil, zero value otherwise.

### GetGrpjobsaccrueOk

`func (o *V0041OpenapiUsersAddCondRespAssociationConditionAssociation) GetGrpjobsaccrueOk() (*V0041OpenapiUsersAddCondRespAssociationConditionAssociationGrpjobsaccrue, bool)`

GetGrpjobsaccrueOk returns a tuple with the Grpjobsaccrue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrpjobsaccrue

`func (o *V0041OpenapiUsersAddCondRespAssociationConditionAssociation) SetGrpjobsaccrue(v V0041OpenapiUsersAddCondRespAssociationConditionAssociationGrpjobsaccrue)`

SetGrpjobsaccrue sets Grpjobsaccrue field to given value.

### HasGrpjobsaccrue

`func (o *V0041OpenapiUsersAddCondRespAssociationConditionAssociation) HasGrpjobsaccrue() bool`

HasGrpjobsaccrue returns a boolean if a field has been set.

### GetGrpsubmitjobs

`func (o *V0041OpenapiUsersAddCondRespAssociationConditionAssociation) GetGrpsubmitjobs() V0041OpenapiUsersAddCondRespAssociationConditionAssociationGrpsubmitjobs`

GetGrpsubmitjobs returns the Grpsubmitjobs field if non-nil, zero value otherwise.

### GetGrpsubmitjobsOk

`func (o *V0041OpenapiUsersAddCondRespAssociationConditionAssociation) GetGrpsubmitjobsOk() (*V0041OpenapiUsersAddCondRespAssociationConditionAssociationGrpsubmitjobs, bool)`

GetGrpsubmitjobsOk returns a tuple with the Grpsubmitjobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrpsubmitjobs

`func (o *V0041OpenapiUsersAddCondRespAssociationConditionAssociation) SetGrpsubmitjobs(v V0041OpenapiUsersAddCondRespAssociationConditionAssociationGrpsubmitjobs)`

SetGrpsubmitjobs sets Grpsubmitjobs field to given value.

### HasGrpsubmitjobs

`func (o *V0041OpenapiUsersAddCondRespAssociationConditionAssociation) HasGrpsubmitjobs() bool`

HasGrpsubmitjobs returns a boolean if a field has been set.

### GetGrptres

`func (o *V0041OpenapiUsersAddCondRespAssociationConditionAssociation) GetGrptres() []V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTresRequestedMaxInner`

GetGrptres returns the Grptres field if non-nil, zero value otherwise.

### GetGrptresOk

`func (o *V0041OpenapiUsersAddCondRespAssociationConditionAssociation) GetGrptresOk() (*[]V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTresRequestedMaxInner, bool)`

GetGrptresOk returns a tuple with the Grptres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrptres

`func (o *V0041OpenapiUsersAddCondRespAssociationConditionAssociation) SetGrptres(v []V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTresRequestedMaxInner)`

SetGrptres sets Grptres field to given value.

### HasGrptres

`func (o *V0041OpenapiUsersAddCondRespAssociationConditionAssociation) HasGrptres() bool`

HasGrptres returns a boolean if a field has been set.

### GetGrptresmins

`func (o *V0041OpenapiUsersAddCondRespAssociationConditionAssociation) GetGrptresmins() []V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTresRequestedMaxInner`

GetGrptresmins returns the Grptresmins field if non-nil, zero value otherwise.

### GetGrptresminsOk

`func (o *V0041OpenapiUsersAddCondRespAssociationConditionAssociation) GetGrptresminsOk() (*[]V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTresRequestedMaxInner, bool)`

GetGrptresminsOk returns a tuple with the Grptresmins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrptresmins

`func (o *V0041OpenapiUsersAddCondRespAssociationConditionAssociation) SetGrptresmins(v []V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTresRequestedMaxInner)`

SetGrptresmins sets Grptresmins field to given value.

### HasGrptresmins

`func (o *V0041OpenapiUsersAddCondRespAssociationConditionAssociation) HasGrptresmins() bool`

HasGrptresmins returns a boolean if a field has been set.

### GetGrptresrunmins

`func (o *V0041OpenapiUsersAddCondRespAssociationConditionAssociation) GetGrptresrunmins() []V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTresRequestedMaxInner`

GetGrptresrunmins returns the Grptresrunmins field if non-nil, zero value otherwise.

### GetGrptresrunminsOk

`func (o *V0041OpenapiUsersAddCondRespAssociationConditionAssociation) GetGrptresrunminsOk() (*[]V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTresRequestedMaxInner, bool)`

GetGrptresrunminsOk returns a tuple with the Grptresrunmins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrptresrunmins

`func (o *V0041OpenapiUsersAddCondRespAssociationConditionAssociation) SetGrptresrunmins(v []V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTresRequestedMaxInner)`

SetGrptresrunmins sets Grptresrunmins field to given value.

### HasGrptresrunmins

`func (o *V0041OpenapiUsersAddCondRespAssociationConditionAssociation) HasGrptresrunmins() bool`

HasGrptresrunmins returns a boolean if a field has been set.

### GetGrpwall

`func (o *V0041OpenapiUsersAddCondRespAssociationConditionAssociation) GetGrpwall() V0041OpenapiUsersAddCondRespAssociationConditionAssociationGrpwall`

GetGrpwall returns the Grpwall field if non-nil, zero value otherwise.

### GetGrpwallOk

`func (o *V0041OpenapiUsersAddCondRespAssociationConditionAssociation) GetGrpwallOk() (*V0041OpenapiUsersAddCondRespAssociationConditionAssociationGrpwall, bool)`

GetGrpwallOk returns a tuple with the Grpwall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrpwall

`func (o *V0041OpenapiUsersAddCondRespAssociationConditionAssociation) SetGrpwall(v V0041OpenapiUsersAddCondRespAssociationConditionAssociationGrpwall)`

SetGrpwall sets Grpwall field to given value.

### HasGrpwall

`func (o *V0041OpenapiUsersAddCondRespAssociationConditionAssociation) HasGrpwall() bool`

HasGrpwall returns a boolean if a field has been set.

### GetMaxjobs

`func (o *V0041OpenapiUsersAddCondRespAssociationConditionAssociation) GetMaxjobs() V0041OpenapiUsersAddCondRespAssociationConditionAssociationMaxjobs`

GetMaxjobs returns the Maxjobs field if non-nil, zero value otherwise.

### GetMaxjobsOk

`func (o *V0041OpenapiUsersAddCondRespAssociationConditionAssociation) GetMaxjobsOk() (*V0041OpenapiUsersAddCondRespAssociationConditionAssociationMaxjobs, bool)`

GetMaxjobsOk returns a tuple with the Maxjobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxjobs

`func (o *V0041OpenapiUsersAddCondRespAssociationConditionAssociation) SetMaxjobs(v V0041OpenapiUsersAddCondRespAssociationConditionAssociationMaxjobs)`

SetMaxjobs sets Maxjobs field to given value.

### HasMaxjobs

`func (o *V0041OpenapiUsersAddCondRespAssociationConditionAssociation) HasMaxjobs() bool`

HasMaxjobs returns a boolean if a field has been set.

### GetMaxjobsaccrue

`func (o *V0041OpenapiUsersAddCondRespAssociationConditionAssociation) GetMaxjobsaccrue() V0041OpenapiUsersAddCondRespAssociationConditionAssociationMaxjobsaccrue`

GetMaxjobsaccrue returns the Maxjobsaccrue field if non-nil, zero value otherwise.

### GetMaxjobsaccrueOk

`func (o *V0041OpenapiUsersAddCondRespAssociationConditionAssociation) GetMaxjobsaccrueOk() (*V0041OpenapiUsersAddCondRespAssociationConditionAssociationMaxjobsaccrue, bool)`

GetMaxjobsaccrueOk returns a tuple with the Maxjobsaccrue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxjobsaccrue

`func (o *V0041OpenapiUsersAddCondRespAssociationConditionAssociation) SetMaxjobsaccrue(v V0041OpenapiUsersAddCondRespAssociationConditionAssociationMaxjobsaccrue)`

SetMaxjobsaccrue sets Maxjobsaccrue field to given value.

### HasMaxjobsaccrue

`func (o *V0041OpenapiUsersAddCondRespAssociationConditionAssociation) HasMaxjobsaccrue() bool`

HasMaxjobsaccrue returns a boolean if a field has been set.

### GetMaxsubmitjobs

`func (o *V0041OpenapiUsersAddCondRespAssociationConditionAssociation) GetMaxsubmitjobs() V0041OpenapiUsersAddCondRespAssociationConditionAssociationMaxsubmitjobs`

GetMaxsubmitjobs returns the Maxsubmitjobs field if non-nil, zero value otherwise.

### GetMaxsubmitjobsOk

`func (o *V0041OpenapiUsersAddCondRespAssociationConditionAssociation) GetMaxsubmitjobsOk() (*V0041OpenapiUsersAddCondRespAssociationConditionAssociationMaxsubmitjobs, bool)`

GetMaxsubmitjobsOk returns a tuple with the Maxsubmitjobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxsubmitjobs

`func (o *V0041OpenapiUsersAddCondRespAssociationConditionAssociation) SetMaxsubmitjobs(v V0041OpenapiUsersAddCondRespAssociationConditionAssociationMaxsubmitjobs)`

SetMaxsubmitjobs sets Maxsubmitjobs field to given value.

### HasMaxsubmitjobs

`func (o *V0041OpenapiUsersAddCondRespAssociationConditionAssociation) HasMaxsubmitjobs() bool`

HasMaxsubmitjobs returns a boolean if a field has been set.

### GetMaxtresminsperjob

`func (o *V0041OpenapiUsersAddCondRespAssociationConditionAssociation) GetMaxtresminsperjob() []V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTresRequestedMaxInner`

GetMaxtresminsperjob returns the Maxtresminsperjob field if non-nil, zero value otherwise.

### GetMaxtresminsperjobOk

`func (o *V0041OpenapiUsersAddCondRespAssociationConditionAssociation) GetMaxtresminsperjobOk() (*[]V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTresRequestedMaxInner, bool)`

GetMaxtresminsperjobOk returns a tuple with the Maxtresminsperjob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxtresminsperjob

`func (o *V0041OpenapiUsersAddCondRespAssociationConditionAssociation) SetMaxtresminsperjob(v []V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTresRequestedMaxInner)`

SetMaxtresminsperjob sets Maxtresminsperjob field to given value.

### HasMaxtresminsperjob

`func (o *V0041OpenapiUsersAddCondRespAssociationConditionAssociation) HasMaxtresminsperjob() bool`

HasMaxtresminsperjob returns a boolean if a field has been set.

### GetMaxtresrunmins

`func (o *V0041OpenapiUsersAddCondRespAssociationConditionAssociation) GetMaxtresrunmins() []V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTresRequestedMaxInner`

GetMaxtresrunmins returns the Maxtresrunmins field if non-nil, zero value otherwise.

### GetMaxtresrunminsOk

`func (o *V0041OpenapiUsersAddCondRespAssociationConditionAssociation) GetMaxtresrunminsOk() (*[]V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTresRequestedMaxInner, bool)`

GetMaxtresrunminsOk returns a tuple with the Maxtresrunmins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxtresrunmins

`func (o *V0041OpenapiUsersAddCondRespAssociationConditionAssociation) SetMaxtresrunmins(v []V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTresRequestedMaxInner)`

SetMaxtresrunmins sets Maxtresrunmins field to given value.

### HasMaxtresrunmins

`func (o *V0041OpenapiUsersAddCondRespAssociationConditionAssociation) HasMaxtresrunmins() bool`

HasMaxtresrunmins returns a boolean if a field has been set.

### GetMaxtresperjob

`func (o *V0041OpenapiUsersAddCondRespAssociationConditionAssociation) GetMaxtresperjob() []V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTresRequestedMaxInner`

GetMaxtresperjob returns the Maxtresperjob field if non-nil, zero value otherwise.

### GetMaxtresperjobOk

`func (o *V0041OpenapiUsersAddCondRespAssociationConditionAssociation) GetMaxtresperjobOk() (*[]V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTresRequestedMaxInner, bool)`

GetMaxtresperjobOk returns a tuple with the Maxtresperjob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxtresperjob

`func (o *V0041OpenapiUsersAddCondRespAssociationConditionAssociation) SetMaxtresperjob(v []V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTresRequestedMaxInner)`

SetMaxtresperjob sets Maxtresperjob field to given value.

### HasMaxtresperjob

`func (o *V0041OpenapiUsersAddCondRespAssociationConditionAssociation) HasMaxtresperjob() bool`

HasMaxtresperjob returns a boolean if a field has been set.

### GetMaxtrespernode

`func (o *V0041OpenapiUsersAddCondRespAssociationConditionAssociation) GetMaxtrespernode() []V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTresRequestedMaxInner`

GetMaxtrespernode returns the Maxtrespernode field if non-nil, zero value otherwise.

### GetMaxtrespernodeOk

`func (o *V0041OpenapiUsersAddCondRespAssociationConditionAssociation) GetMaxtrespernodeOk() (*[]V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTresRequestedMaxInner, bool)`

GetMaxtrespernodeOk returns a tuple with the Maxtrespernode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxtrespernode

`func (o *V0041OpenapiUsersAddCondRespAssociationConditionAssociation) SetMaxtrespernode(v []V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTresRequestedMaxInner)`

SetMaxtrespernode sets Maxtrespernode field to given value.

### HasMaxtrespernode

`func (o *V0041OpenapiUsersAddCondRespAssociationConditionAssociation) HasMaxtrespernode() bool`

HasMaxtrespernode returns a boolean if a field has been set.

### GetMaxwalldurationperjob

`func (o *V0041OpenapiUsersAddCondRespAssociationConditionAssociation) GetMaxwalldurationperjob() V0041OpenapiUsersAddCondRespAssociationConditionAssociationMaxwalldurationperjob`

GetMaxwalldurationperjob returns the Maxwalldurationperjob field if non-nil, zero value otherwise.

### GetMaxwalldurationperjobOk

`func (o *V0041OpenapiUsersAddCondRespAssociationConditionAssociation) GetMaxwalldurationperjobOk() (*V0041OpenapiUsersAddCondRespAssociationConditionAssociationMaxwalldurationperjob, bool)`

GetMaxwalldurationperjobOk returns a tuple with the Maxwalldurationperjob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxwalldurationperjob

`func (o *V0041OpenapiUsersAddCondRespAssociationConditionAssociation) SetMaxwalldurationperjob(v V0041OpenapiUsersAddCondRespAssociationConditionAssociationMaxwalldurationperjob)`

SetMaxwalldurationperjob sets Maxwalldurationperjob field to given value.

### HasMaxwalldurationperjob

`func (o *V0041OpenapiUsersAddCondRespAssociationConditionAssociation) HasMaxwalldurationperjob() bool`

HasMaxwalldurationperjob returns a boolean if a field has been set.

### GetMinpriothresh

`func (o *V0041OpenapiUsersAddCondRespAssociationConditionAssociation) GetMinpriothresh() V0041OpenapiUsersAddCondRespAssociationConditionAssociationMinpriothresh`

GetMinpriothresh returns the Minpriothresh field if non-nil, zero value otherwise.

### GetMinpriothreshOk

`func (o *V0041OpenapiUsersAddCondRespAssociationConditionAssociation) GetMinpriothreshOk() (*V0041OpenapiUsersAddCondRespAssociationConditionAssociationMinpriothresh, bool)`

GetMinpriothreshOk returns a tuple with the Minpriothresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinpriothresh

`func (o *V0041OpenapiUsersAddCondRespAssociationConditionAssociation) SetMinpriothresh(v V0041OpenapiUsersAddCondRespAssociationConditionAssociationMinpriothresh)`

SetMinpriothresh sets Minpriothresh field to given value.

### HasMinpriothresh

`func (o *V0041OpenapiUsersAddCondRespAssociationConditionAssociation) HasMinpriothresh() bool`

HasMinpriothresh returns a boolean if a field has been set.

### GetParent

`func (o *V0041OpenapiUsersAddCondRespAssociationConditionAssociation) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *V0041OpenapiUsersAddCondRespAssociationConditionAssociation) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *V0041OpenapiUsersAddCondRespAssociationConditionAssociation) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *V0041OpenapiUsersAddCondRespAssociationConditionAssociation) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPriority

`func (o *V0041OpenapiUsersAddCondRespAssociationConditionAssociation) GetPriority() V0041OpenapiUsersAddCondRespAssociationConditionAssociationPriority`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *V0041OpenapiUsersAddCondRespAssociationConditionAssociation) GetPriorityOk() (*V0041OpenapiUsersAddCondRespAssociationConditionAssociationPriority, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *V0041OpenapiUsersAddCondRespAssociationConditionAssociation) SetPriority(v V0041OpenapiUsersAddCondRespAssociationConditionAssociationPriority)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *V0041OpenapiUsersAddCondRespAssociationConditionAssociation) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetQoslevel

`func (o *V0041OpenapiUsersAddCondRespAssociationConditionAssociation) GetQoslevel() []string`

GetQoslevel returns the Qoslevel field if non-nil, zero value otherwise.

### GetQoslevelOk

`func (o *V0041OpenapiUsersAddCondRespAssociationConditionAssociation) GetQoslevelOk() (*[]string, bool)`

GetQoslevelOk returns a tuple with the Qoslevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQoslevel

`func (o *V0041OpenapiUsersAddCondRespAssociationConditionAssociation) SetQoslevel(v []string)`

SetQoslevel sets Qoslevel field to given value.

### HasQoslevel

`func (o *V0041OpenapiUsersAddCondRespAssociationConditionAssociation) HasQoslevel() bool`

HasQoslevel returns a boolean if a field has been set.

### GetFairshare

`func (o *V0041OpenapiUsersAddCondRespAssociationConditionAssociation) GetFairshare() int32`

GetFairshare returns the Fairshare field if non-nil, zero value otherwise.

### GetFairshareOk

`func (o *V0041OpenapiUsersAddCondRespAssociationConditionAssociation) GetFairshareOk() (*int32, bool)`

GetFairshareOk returns a tuple with the Fairshare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFairshare

`func (o *V0041OpenapiUsersAddCondRespAssociationConditionAssociation) SetFairshare(v int32)`

SetFairshare sets Fairshare field to given value.

### HasFairshare

`func (o *V0041OpenapiUsersAddCondRespAssociationConditionAssociation) HasFairshare() bool`

HasFairshare returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


