# V0037License

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | name of license | [optional] 
**Total** | Pointer to **int32** | total number of licenses | [optional] 
**InUse** | Pointer to **int32** | number of licenses in use | [optional] 
**Available** | Pointer to **int32** | number of licenses available | [optional] 
**Reserved** | Pointer to **int32** | number of licenses reserved | [optional] 
**Remote** | Pointer to **bool** | license is remote | [optional] 

## Methods

### NewV0037License

`func NewV0037License() *V0037License`

NewV0037License instantiates a new V0037License object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0037LicenseWithDefaults

`func NewV0037LicenseWithDefaults() *V0037License`

NewV0037LicenseWithDefaults instantiates a new V0037License object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *V0037License) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V0037License) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V0037License) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V0037License) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTotal

`func (o *V0037License) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *V0037License) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *V0037License) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *V0037License) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetInUse

`func (o *V0037License) GetInUse() int32`

GetInUse returns the InUse field if non-nil, zero value otherwise.

### GetInUseOk

`func (o *V0037License) GetInUseOk() (*int32, bool)`

GetInUseOk returns a tuple with the InUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInUse

`func (o *V0037License) SetInUse(v int32)`

SetInUse sets InUse field to given value.

### HasInUse

`func (o *V0037License) HasInUse() bool`

HasInUse returns a boolean if a field has been set.

### GetAvailable

`func (o *V0037License) GetAvailable() int32`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *V0037License) GetAvailableOk() (*int32, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *V0037License) SetAvailable(v int32)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *V0037License) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetReserved

`func (o *V0037License) GetReserved() int32`

GetReserved returns the Reserved field if non-nil, zero value otherwise.

### GetReservedOk

`func (o *V0037License) GetReservedOk() (*int32, bool)`

GetReservedOk returns a tuple with the Reserved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReserved

`func (o *V0037License) SetReserved(v int32)`

SetReserved sets Reserved field to given value.

### HasReserved

`func (o *V0037License) HasReserved() bool`

HasReserved returns a boolean if a field has been set.

### GetRemote

`func (o *V0037License) GetRemote() bool`

GetRemote returns the Remote field if non-nil, zero value otherwise.

### GetRemoteOk

`func (o *V0037License) GetRemoteOk() (*bool, bool)`

GetRemoteOk returns a tuple with the Remote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemote

`func (o *V0037License) SetRemote(v bool)`

SetRemote sets Remote field to given value.

### HasRemote

`func (o *V0037License) HasRemote() bool`

HasRemote returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


