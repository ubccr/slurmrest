# V0044NodeResourceLayout

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Node** | **string** | Node name | 
**SocketsPerNode** | Pointer to **int32** | Sockets per node | [optional] 
**CoresPerSocket** | Pointer to **int32** | Cores per socket | [optional] 
**MemAlloc** | Pointer to **int64** | Allocated memory | [optional] 
**CoreBitmap** | Pointer to **string** | Abstract core bitmap | [optional] 
**Channel** | Pointer to [**V0044Uint32NoValStruct**](V0044Uint32NoValStruct.md) |  | [optional] 
**Gres** | Pointer to [**[]V0044NodeGresLayout**](V0044NodeGresLayout.md) |  | [optional] 

## Methods

### NewV0044NodeResourceLayout

`func NewV0044NodeResourceLayout(node string, ) *V0044NodeResourceLayout`

NewV0044NodeResourceLayout instantiates a new V0044NodeResourceLayout object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0044NodeResourceLayoutWithDefaults

`func NewV0044NodeResourceLayoutWithDefaults() *V0044NodeResourceLayout`

NewV0044NodeResourceLayoutWithDefaults instantiates a new V0044NodeResourceLayout object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNode

`func (o *V0044NodeResourceLayout) GetNode() string`

GetNode returns the Node field if non-nil, zero value otherwise.

### GetNodeOk

`func (o *V0044NodeResourceLayout) GetNodeOk() (*string, bool)`

GetNodeOk returns a tuple with the Node field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNode

`func (o *V0044NodeResourceLayout) SetNode(v string)`

SetNode sets Node field to given value.


### GetSocketsPerNode

`func (o *V0044NodeResourceLayout) GetSocketsPerNode() int32`

GetSocketsPerNode returns the SocketsPerNode field if non-nil, zero value otherwise.

### GetSocketsPerNodeOk

`func (o *V0044NodeResourceLayout) GetSocketsPerNodeOk() (*int32, bool)`

GetSocketsPerNodeOk returns a tuple with the SocketsPerNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocketsPerNode

`func (o *V0044NodeResourceLayout) SetSocketsPerNode(v int32)`

SetSocketsPerNode sets SocketsPerNode field to given value.

### HasSocketsPerNode

`func (o *V0044NodeResourceLayout) HasSocketsPerNode() bool`

HasSocketsPerNode returns a boolean if a field has been set.

### GetCoresPerSocket

`func (o *V0044NodeResourceLayout) GetCoresPerSocket() int32`

GetCoresPerSocket returns the CoresPerSocket field if non-nil, zero value otherwise.

### GetCoresPerSocketOk

`func (o *V0044NodeResourceLayout) GetCoresPerSocketOk() (*int32, bool)`

GetCoresPerSocketOk returns a tuple with the CoresPerSocket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoresPerSocket

`func (o *V0044NodeResourceLayout) SetCoresPerSocket(v int32)`

SetCoresPerSocket sets CoresPerSocket field to given value.

### HasCoresPerSocket

`func (o *V0044NodeResourceLayout) HasCoresPerSocket() bool`

HasCoresPerSocket returns a boolean if a field has been set.

### GetMemAlloc

`func (o *V0044NodeResourceLayout) GetMemAlloc() int64`

GetMemAlloc returns the MemAlloc field if non-nil, zero value otherwise.

### GetMemAllocOk

`func (o *V0044NodeResourceLayout) GetMemAllocOk() (*int64, bool)`

GetMemAllocOk returns a tuple with the MemAlloc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemAlloc

`func (o *V0044NodeResourceLayout) SetMemAlloc(v int64)`

SetMemAlloc sets MemAlloc field to given value.

### HasMemAlloc

`func (o *V0044NodeResourceLayout) HasMemAlloc() bool`

HasMemAlloc returns a boolean if a field has been set.

### GetCoreBitmap

`func (o *V0044NodeResourceLayout) GetCoreBitmap() string`

GetCoreBitmap returns the CoreBitmap field if non-nil, zero value otherwise.

### GetCoreBitmapOk

`func (o *V0044NodeResourceLayout) GetCoreBitmapOk() (*string, bool)`

GetCoreBitmapOk returns a tuple with the CoreBitmap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoreBitmap

`func (o *V0044NodeResourceLayout) SetCoreBitmap(v string)`

SetCoreBitmap sets CoreBitmap field to given value.

### HasCoreBitmap

`func (o *V0044NodeResourceLayout) HasCoreBitmap() bool`

HasCoreBitmap returns a boolean if a field has been set.

### GetChannel

`func (o *V0044NodeResourceLayout) GetChannel() V0044Uint32NoValStruct`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *V0044NodeResourceLayout) GetChannelOk() (*V0044Uint32NoValStruct, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *V0044NodeResourceLayout) SetChannel(v V0044Uint32NoValStruct)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *V0044NodeResourceLayout) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetGres

`func (o *V0044NodeResourceLayout) GetGres() []V0044NodeGresLayout`

GetGres returns the Gres field if non-nil, zero value otherwise.

### GetGresOk

`func (o *V0044NodeResourceLayout) GetGresOk() (*[]V0044NodeGresLayout, bool)`

GetGresOk returns a tuple with the Gres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGres

`func (o *V0044NodeResourceLayout) SetGres(v []V0044NodeGresLayout)`

SetGres sets Gres field to given value.

### HasGres

`func (o *V0044NodeResourceLayout) HasGres() bool`

HasGres returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


