# V0044Tres

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | TRES type (CPU, MEM, etc) | 
**Name** | Pointer to **string** | TRES name (if applicable) | [optional] 
**Id** | Pointer to **int32** | ID used in the database | [optional] 
**Count** | Pointer to **int64** | TRES count (0 if listed generically) | [optional] 

## Methods

### NewV0044Tres

`func NewV0044Tres(type_ string, ) *V0044Tres`

NewV0044Tres instantiates a new V0044Tres object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0044TresWithDefaults

`func NewV0044TresWithDefaults() *V0044Tres`

NewV0044TresWithDefaults instantiates a new V0044Tres object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *V0044Tres) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V0044Tres) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V0044Tres) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *V0044Tres) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V0044Tres) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V0044Tres) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V0044Tres) HasName() bool`

HasName returns a boolean if a field has been set.

### GetId

`func (o *V0044Tres) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V0044Tres) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V0044Tres) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *V0044Tres) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCount

`func (o *V0044Tres) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *V0044Tres) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *V0044Tres) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *V0044Tres) HasCount() bool`

HasCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


