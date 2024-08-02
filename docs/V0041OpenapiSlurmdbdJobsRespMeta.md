# V0041OpenapiSlurmdbdJobsRespMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Plugin** | Pointer to [**V0041OpenapiSlurmdbdJobsRespMetaPlugin**](V0041OpenapiSlurmdbdJobsRespMetaPlugin.md) |  | [optional] 
**Client** | Pointer to [**V0041OpenapiSlurmdbdJobsRespMetaClient**](V0041OpenapiSlurmdbdJobsRespMetaClient.md) |  | [optional] 
**Command** | Pointer to **[]string** | CLI command (if applicable) | [optional] 
**Slurm** | Pointer to [**V0041OpenapiSlurmdbdJobsRespMetaSlurm**](V0041OpenapiSlurmdbdJobsRespMetaSlurm.md) |  | [optional] 

## Methods

### NewV0041OpenapiSlurmdbdJobsRespMeta

`func NewV0041OpenapiSlurmdbdJobsRespMeta() *V0041OpenapiSlurmdbdJobsRespMeta`

NewV0041OpenapiSlurmdbdJobsRespMeta instantiates a new V0041OpenapiSlurmdbdJobsRespMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0041OpenapiSlurmdbdJobsRespMetaWithDefaults

`func NewV0041OpenapiSlurmdbdJobsRespMetaWithDefaults() *V0041OpenapiSlurmdbdJobsRespMeta`

NewV0041OpenapiSlurmdbdJobsRespMetaWithDefaults instantiates a new V0041OpenapiSlurmdbdJobsRespMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlugin

`func (o *V0041OpenapiSlurmdbdJobsRespMeta) GetPlugin() V0041OpenapiSlurmdbdJobsRespMetaPlugin`

GetPlugin returns the Plugin field if non-nil, zero value otherwise.

### GetPluginOk

`func (o *V0041OpenapiSlurmdbdJobsRespMeta) GetPluginOk() (*V0041OpenapiSlurmdbdJobsRespMetaPlugin, bool)`

GetPluginOk returns a tuple with the Plugin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlugin

`func (o *V0041OpenapiSlurmdbdJobsRespMeta) SetPlugin(v V0041OpenapiSlurmdbdJobsRespMetaPlugin)`

SetPlugin sets Plugin field to given value.

### HasPlugin

`func (o *V0041OpenapiSlurmdbdJobsRespMeta) HasPlugin() bool`

HasPlugin returns a boolean if a field has been set.

### GetClient

`func (o *V0041OpenapiSlurmdbdJobsRespMeta) GetClient() V0041OpenapiSlurmdbdJobsRespMetaClient`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *V0041OpenapiSlurmdbdJobsRespMeta) GetClientOk() (*V0041OpenapiSlurmdbdJobsRespMetaClient, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *V0041OpenapiSlurmdbdJobsRespMeta) SetClient(v V0041OpenapiSlurmdbdJobsRespMetaClient)`

SetClient sets Client field to given value.

### HasClient

`func (o *V0041OpenapiSlurmdbdJobsRespMeta) HasClient() bool`

HasClient returns a boolean if a field has been set.

### GetCommand

`func (o *V0041OpenapiSlurmdbdJobsRespMeta) GetCommand() []string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *V0041OpenapiSlurmdbdJobsRespMeta) GetCommandOk() (*[]string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *V0041OpenapiSlurmdbdJobsRespMeta) SetCommand(v []string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *V0041OpenapiSlurmdbdJobsRespMeta) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### GetSlurm

`func (o *V0041OpenapiSlurmdbdJobsRespMeta) GetSlurm() V0041OpenapiSlurmdbdJobsRespMetaSlurm`

GetSlurm returns the Slurm field if non-nil, zero value otherwise.

### GetSlurmOk

`func (o *V0041OpenapiSlurmdbdJobsRespMeta) GetSlurmOk() (*V0041OpenapiSlurmdbdJobsRespMetaSlurm, bool)`

GetSlurmOk returns a tuple with the Slurm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlurm

`func (o *V0041OpenapiSlurmdbdJobsRespMeta) SetSlurm(v V0041OpenapiSlurmdbdJobsRespMetaSlurm)`

SetSlurm sets Slurm field to given value.

### HasSlurm

`func (o *V0041OpenapiSlurmdbdJobsRespMeta) HasSlurm() bool`

HasSlurm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


