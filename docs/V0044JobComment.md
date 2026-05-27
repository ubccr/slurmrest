# V0044JobComment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Administrator** | Pointer to **string** | Arbitrary comment made by administrator | [optional] 
**Job** | Pointer to **string** | Arbitrary comment made by user | [optional] 
**System** | Pointer to **string** | Arbitrary comment from slurmctld | [optional] 

## Methods

### NewV0044JobComment

`func NewV0044JobComment() *V0044JobComment`

NewV0044JobComment instantiates a new V0044JobComment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0044JobCommentWithDefaults

`func NewV0044JobCommentWithDefaults() *V0044JobComment`

NewV0044JobCommentWithDefaults instantiates a new V0044JobComment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdministrator

`func (o *V0044JobComment) GetAdministrator() string`

GetAdministrator returns the Administrator field if non-nil, zero value otherwise.

### GetAdministratorOk

`func (o *V0044JobComment) GetAdministratorOk() (*string, bool)`

GetAdministratorOk returns a tuple with the Administrator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministrator

`func (o *V0044JobComment) SetAdministrator(v string)`

SetAdministrator sets Administrator field to given value.

### HasAdministrator

`func (o *V0044JobComment) HasAdministrator() bool`

HasAdministrator returns a boolean if a field has been set.

### GetJob

`func (o *V0044JobComment) GetJob() string`

GetJob returns the Job field if non-nil, zero value otherwise.

### GetJobOk

`func (o *V0044JobComment) GetJobOk() (*string, bool)`

GetJobOk returns a tuple with the Job field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJob

`func (o *V0044JobComment) SetJob(v string)`

SetJob sets Job field to given value.

### HasJob

`func (o *V0044JobComment) HasJob() bool`

HasJob returns a boolean if a field has been set.

### GetSystem

`func (o *V0044JobComment) GetSystem() string`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *V0044JobComment) GetSystemOk() (*string, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *V0044JobComment) SetSystem(v string)`

SetSystem sets System field to given value.

### HasSystem

`func (o *V0044JobComment) HasSystem() bool`

HasSystem returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


