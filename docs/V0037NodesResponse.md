# V0037NodesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to [**[]V0037Error**](V0037Error.md) | slurm errors | [optional] 
**Nodes** | Pointer to [**[]V0037Node**](V0037Node.md) | nodes info | [optional] 

## Methods

### NewV0037NodesResponse

`func NewV0037NodesResponse() *V0037NodesResponse`

NewV0037NodesResponse instantiates a new V0037NodesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0037NodesResponseWithDefaults

`func NewV0037NodesResponseWithDefaults() *V0037NodesResponse`

NewV0037NodesResponseWithDefaults instantiates a new V0037NodesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *V0037NodesResponse) GetErrors() []V0037Error`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V0037NodesResponse) GetErrorsOk() (*[]V0037Error, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V0037NodesResponse) SetErrors(v []V0037Error)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V0037NodesResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetNodes

`func (o *V0037NodesResponse) GetNodes() []V0037Node`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *V0037NodesResponse) GetNodesOk() (*[]V0037Node, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *V0037NodesResponse) SetNodes(v []V0037Node)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *V0037NodesResponse) HasNodes() bool`

HasNodes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


