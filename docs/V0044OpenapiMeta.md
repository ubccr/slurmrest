# V0044OpenapiMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Plugin** | Pointer to [**V0044OpenapiMetaPlugin**](V0044OpenapiMetaPlugin.md) |  | [optional] 
**Client** | Pointer to [**V0044OpenapiMetaClient**](V0044OpenapiMetaClient.md) |  | [optional] 
**Command** | Pointer to **[]string** |  | [optional] 
**Slurm** | Pointer to [**V0044OpenapiMetaSlurm**](V0044OpenapiMetaSlurm.md) |  | [optional] 

## Methods

### NewV0044OpenapiMeta

`func NewV0044OpenapiMeta() *V0044OpenapiMeta`

NewV0044OpenapiMeta instantiates a new V0044OpenapiMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0044OpenapiMetaWithDefaults

`func NewV0044OpenapiMetaWithDefaults() *V0044OpenapiMeta`

NewV0044OpenapiMetaWithDefaults instantiates a new V0044OpenapiMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlugin

`func (o *V0044OpenapiMeta) GetPlugin() V0044OpenapiMetaPlugin`

GetPlugin returns the Plugin field if non-nil, zero value otherwise.

### GetPluginOk

`func (o *V0044OpenapiMeta) GetPluginOk() (*V0044OpenapiMetaPlugin, bool)`

GetPluginOk returns a tuple with the Plugin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlugin

`func (o *V0044OpenapiMeta) SetPlugin(v V0044OpenapiMetaPlugin)`

SetPlugin sets Plugin field to given value.

### HasPlugin

`func (o *V0044OpenapiMeta) HasPlugin() bool`

HasPlugin returns a boolean if a field has been set.

### GetClient

`func (o *V0044OpenapiMeta) GetClient() V0044OpenapiMetaClient`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *V0044OpenapiMeta) GetClientOk() (*V0044OpenapiMetaClient, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *V0044OpenapiMeta) SetClient(v V0044OpenapiMetaClient)`

SetClient sets Client field to given value.

### HasClient

`func (o *V0044OpenapiMeta) HasClient() bool`

HasClient returns a boolean if a field has been set.

### GetCommand

`func (o *V0044OpenapiMeta) GetCommand() []string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *V0044OpenapiMeta) GetCommandOk() (*[]string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *V0044OpenapiMeta) SetCommand(v []string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *V0044OpenapiMeta) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### GetSlurm

`func (o *V0044OpenapiMeta) GetSlurm() V0044OpenapiMetaSlurm`

GetSlurm returns the Slurm field if non-nil, zero value otherwise.

### GetSlurmOk

`func (o *V0044OpenapiMeta) GetSlurmOk() (*V0044OpenapiMetaSlurm, bool)`

GetSlurmOk returns a tuple with the Slurm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlurm

`func (o *V0044OpenapiMeta) SetSlurm(v V0044OpenapiMetaSlurm)`

SetSlurm sets Slurm field to given value.

### HasSlurm

`func (o *V0044OpenapiMeta) HasSlurm() bool`

HasSlurm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


