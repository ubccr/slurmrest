# V0044StatsMsgRpcDump

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TypeId** | **int32** | Message type as integer | 
**MessageType** | **string** | Message type as string (Slurm RPC message type) | 
**Count** | **[]string** |  | 

## Methods

### NewV0044StatsMsgRpcDump

`func NewV0044StatsMsgRpcDump(typeId int32, messageType string, count []string, ) *V0044StatsMsgRpcDump`

NewV0044StatsMsgRpcDump instantiates a new V0044StatsMsgRpcDump object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0044StatsMsgRpcDumpWithDefaults

`func NewV0044StatsMsgRpcDumpWithDefaults() *V0044StatsMsgRpcDump`

NewV0044StatsMsgRpcDumpWithDefaults instantiates a new V0044StatsMsgRpcDump object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTypeId

`func (o *V0044StatsMsgRpcDump) GetTypeId() int32`

GetTypeId returns the TypeId field if non-nil, zero value otherwise.

### GetTypeIdOk

`func (o *V0044StatsMsgRpcDump) GetTypeIdOk() (*int32, bool)`

GetTypeIdOk returns a tuple with the TypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeId

`func (o *V0044StatsMsgRpcDump) SetTypeId(v int32)`

SetTypeId sets TypeId field to given value.


### GetMessageType

`func (o *V0044StatsMsgRpcDump) GetMessageType() string`

GetMessageType returns the MessageType field if non-nil, zero value otherwise.

### GetMessageTypeOk

`func (o *V0044StatsMsgRpcDump) GetMessageTypeOk() (*string, bool)`

GetMessageTypeOk returns a tuple with the MessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageType

`func (o *V0044StatsMsgRpcDump) SetMessageType(v string)`

SetMessageType sets MessageType field to given value.


### GetCount

`func (o *V0044StatsMsgRpcDump) GetCount() []string`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *V0044StatsMsgRpcDump) GetCountOk() (*[]string, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *V0044StatsMsgRpcDump) SetCount(v []string)`

SetCount sets Count field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


