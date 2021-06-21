# V0036NodesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to [**[]V0036Error**](V0036Error.md) | slurm errors | [optional] 
**Nodes** | Pointer to [**[]V0036Node**](V0036Node.md) | nodes info | [optional] 

## Methods

### NewV0036NodesResponse

`func NewV0036NodesResponse() *V0036NodesResponse`

NewV0036NodesResponse instantiates a new V0036NodesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0036NodesResponseWithDefaults

`func NewV0036NodesResponseWithDefaults() *V0036NodesResponse`

NewV0036NodesResponseWithDefaults instantiates a new V0036NodesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *V0036NodesResponse) GetErrors() []V0036Error`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V0036NodesResponse) GetErrorsOk() (*[]V0036Error, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V0036NodesResponse) SetErrors(v []V0036Error)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V0036NodesResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetNodes

`func (o *V0036NodesResponse) GetNodes() []V0036Node`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *V0036NodesResponse) GetNodesOk() (*[]V0036Node, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *V0036NodesResponse) SetNodes(v []V0036Node)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *V0036NodesResponse) HasNodes() bool`

HasNodes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


