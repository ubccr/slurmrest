# V0044Accounting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Allocated** | Pointer to [**V0044AccountingAllocated**](V0044AccountingAllocated.md) |  | [optional] 
**Id** | Pointer to **int32** | Association ID or Workload characterization key ID | [optional] 
**IdAlt** | Pointer to **int32** | Alternate ID (not currently used) | [optional] 
**Start** | Pointer to **int64** | When the record was started (UNIX timestamp) (UNIX timestamp or time string recognized by Slurm (e.g., &#39;[MM/DD[/YY]-]HH:MM[:SS]&#39;)) | [optional] 
**TRES** | Pointer to [**V0044Tres**](V0044Tres.md) |  | [optional] 

## Methods

### NewV0044Accounting

`func NewV0044Accounting() *V0044Accounting`

NewV0044Accounting instantiates a new V0044Accounting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0044AccountingWithDefaults

`func NewV0044AccountingWithDefaults() *V0044Accounting`

NewV0044AccountingWithDefaults instantiates a new V0044Accounting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllocated

`func (o *V0044Accounting) GetAllocated() V0044AccountingAllocated`

GetAllocated returns the Allocated field if non-nil, zero value otherwise.

### GetAllocatedOk

`func (o *V0044Accounting) GetAllocatedOk() (*V0044AccountingAllocated, bool)`

GetAllocatedOk returns a tuple with the Allocated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocated

`func (o *V0044Accounting) SetAllocated(v V0044AccountingAllocated)`

SetAllocated sets Allocated field to given value.

### HasAllocated

`func (o *V0044Accounting) HasAllocated() bool`

HasAllocated returns a boolean if a field has been set.

### GetId

`func (o *V0044Accounting) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V0044Accounting) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V0044Accounting) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *V0044Accounting) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdAlt

`func (o *V0044Accounting) GetIdAlt() int32`

GetIdAlt returns the IdAlt field if non-nil, zero value otherwise.

### GetIdAltOk

`func (o *V0044Accounting) GetIdAltOk() (*int32, bool)`

GetIdAltOk returns a tuple with the IdAlt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdAlt

`func (o *V0044Accounting) SetIdAlt(v int32)`

SetIdAlt sets IdAlt field to given value.

### HasIdAlt

`func (o *V0044Accounting) HasIdAlt() bool`

HasIdAlt returns a boolean if a field has been set.

### GetStart

`func (o *V0044Accounting) GetStart() int64`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *V0044Accounting) GetStartOk() (*int64, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *V0044Accounting) SetStart(v int64)`

SetStart sets Start field to given value.

### HasStart

`func (o *V0044Accounting) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetTRES

`func (o *V0044Accounting) GetTRES() V0044Tres`

GetTRES returns the TRES field if non-nil, zero value otherwise.

### GetTRESOk

`func (o *V0044Accounting) GetTRESOk() (*V0044Tres, bool)`

GetTRESOk returns a tuple with the TRES field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTRES

`func (o *V0044Accounting) SetTRES(v V0044Tres)`

SetTRES sets TRES field to given value.

### HasTRES

`func (o *V0044Accounting) HasTRES() bool`

HasTRES returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


