# V0041OpenapiSlurmdbdConfigRespClustersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Controller** | Pointer to [**V0041OpenapiSlurmdbdConfigRespClustersInnerController**](V0041OpenapiSlurmdbdConfigRespClustersInnerController.md) |  | [optional] 
**Flags** | Pointer to **[]string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Nodes** | Pointer to **string** |  | [optional] 
**SelectPlugin** | Pointer to **string** |  | [optional] 
**Associations** | Pointer to [**V0041OpenapiSlurmdbdConfigRespClustersInnerAssociations**](V0041OpenapiSlurmdbdConfigRespClustersInnerAssociations.md) |  | [optional] 
**RpcVersion** | Pointer to **int32** |  | [optional] 
**Tres** | Pointer to [**[]V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTresRequestedMaxInner**](V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTresRequestedMaxInner.md) |  | [optional] 

## Methods

### NewV0041OpenapiSlurmdbdConfigRespClustersInner

`func NewV0041OpenapiSlurmdbdConfigRespClustersInner() *V0041OpenapiSlurmdbdConfigRespClustersInner`

NewV0041OpenapiSlurmdbdConfigRespClustersInner instantiates a new V0041OpenapiSlurmdbdConfigRespClustersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0041OpenapiSlurmdbdConfigRespClustersInnerWithDefaults

`func NewV0041OpenapiSlurmdbdConfigRespClustersInnerWithDefaults() *V0041OpenapiSlurmdbdConfigRespClustersInner`

NewV0041OpenapiSlurmdbdConfigRespClustersInnerWithDefaults instantiates a new V0041OpenapiSlurmdbdConfigRespClustersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetController

`func (o *V0041OpenapiSlurmdbdConfigRespClustersInner) GetController() V0041OpenapiSlurmdbdConfigRespClustersInnerController`

GetController returns the Controller field if non-nil, zero value otherwise.

### GetControllerOk

`func (o *V0041OpenapiSlurmdbdConfigRespClustersInner) GetControllerOk() (*V0041OpenapiSlurmdbdConfigRespClustersInnerController, bool)`

GetControllerOk returns a tuple with the Controller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetController

`func (o *V0041OpenapiSlurmdbdConfigRespClustersInner) SetController(v V0041OpenapiSlurmdbdConfigRespClustersInnerController)`

SetController sets Controller field to given value.

### HasController

`func (o *V0041OpenapiSlurmdbdConfigRespClustersInner) HasController() bool`

HasController returns a boolean if a field has been set.

### GetFlags

`func (o *V0041OpenapiSlurmdbdConfigRespClustersInner) GetFlags() []string`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *V0041OpenapiSlurmdbdConfigRespClustersInner) GetFlagsOk() (*[]string, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *V0041OpenapiSlurmdbdConfigRespClustersInner) SetFlags(v []string)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *V0041OpenapiSlurmdbdConfigRespClustersInner) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetName

`func (o *V0041OpenapiSlurmdbdConfigRespClustersInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V0041OpenapiSlurmdbdConfigRespClustersInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V0041OpenapiSlurmdbdConfigRespClustersInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V0041OpenapiSlurmdbdConfigRespClustersInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNodes

`func (o *V0041OpenapiSlurmdbdConfigRespClustersInner) GetNodes() string`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *V0041OpenapiSlurmdbdConfigRespClustersInner) GetNodesOk() (*string, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *V0041OpenapiSlurmdbdConfigRespClustersInner) SetNodes(v string)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *V0041OpenapiSlurmdbdConfigRespClustersInner) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetSelectPlugin

`func (o *V0041OpenapiSlurmdbdConfigRespClustersInner) GetSelectPlugin() string`

GetSelectPlugin returns the SelectPlugin field if non-nil, zero value otherwise.

### GetSelectPluginOk

`func (o *V0041OpenapiSlurmdbdConfigRespClustersInner) GetSelectPluginOk() (*string, bool)`

GetSelectPluginOk returns a tuple with the SelectPlugin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectPlugin

`func (o *V0041OpenapiSlurmdbdConfigRespClustersInner) SetSelectPlugin(v string)`

SetSelectPlugin sets SelectPlugin field to given value.

### HasSelectPlugin

`func (o *V0041OpenapiSlurmdbdConfigRespClustersInner) HasSelectPlugin() bool`

HasSelectPlugin returns a boolean if a field has been set.

### GetAssociations

`func (o *V0041OpenapiSlurmdbdConfigRespClustersInner) GetAssociations() V0041OpenapiSlurmdbdConfigRespClustersInnerAssociations`

GetAssociations returns the Associations field if non-nil, zero value otherwise.

### GetAssociationsOk

`func (o *V0041OpenapiSlurmdbdConfigRespClustersInner) GetAssociationsOk() (*V0041OpenapiSlurmdbdConfigRespClustersInnerAssociations, bool)`

GetAssociationsOk returns a tuple with the Associations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociations

`func (o *V0041OpenapiSlurmdbdConfigRespClustersInner) SetAssociations(v V0041OpenapiSlurmdbdConfigRespClustersInnerAssociations)`

SetAssociations sets Associations field to given value.

### HasAssociations

`func (o *V0041OpenapiSlurmdbdConfigRespClustersInner) HasAssociations() bool`

HasAssociations returns a boolean if a field has been set.

### GetRpcVersion

`func (o *V0041OpenapiSlurmdbdConfigRespClustersInner) GetRpcVersion() int32`

GetRpcVersion returns the RpcVersion field if non-nil, zero value otherwise.

### GetRpcVersionOk

`func (o *V0041OpenapiSlurmdbdConfigRespClustersInner) GetRpcVersionOk() (*int32, bool)`

GetRpcVersionOk returns a tuple with the RpcVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpcVersion

`func (o *V0041OpenapiSlurmdbdConfigRespClustersInner) SetRpcVersion(v int32)`

SetRpcVersion sets RpcVersion field to given value.

### HasRpcVersion

`func (o *V0041OpenapiSlurmdbdConfigRespClustersInner) HasRpcVersion() bool`

HasRpcVersion returns a boolean if a field has been set.

### GetTres

`func (o *V0041OpenapiSlurmdbdConfigRespClustersInner) GetTres() []V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTresRequestedMaxInner`

GetTres returns the Tres field if non-nil, zero value otherwise.

### GetTresOk

`func (o *V0041OpenapiSlurmdbdConfigRespClustersInner) GetTresOk() (*[]V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTresRequestedMaxInner, bool)`

GetTresOk returns a tuple with the Tres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTres

`func (o *V0041OpenapiSlurmdbdConfigRespClustersInner) SetTres(v []V0041OpenapiSlurmdbdJobsRespJobsInnerStepsInnerTresRequestedMaxInner)`

SetTres sets Tres field to given value.

### HasTres

`func (o *V0041OpenapiSlurmdbdConfigRespClustersInner) HasTres() bool`

HasTres returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


