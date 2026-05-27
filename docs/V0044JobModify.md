# V0044JobModify

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to [**V0044JobComment**](V0044JobComment.md) |  | [optional] 
**DerivedExitCode** | Pointer to [**V0044ProcessExitCodeVerbose**](V0044ProcessExitCodeVerbose.md) |  | [optional] 
**Extra** | Pointer to **string** | Arbitrary string used for node filtering if extra constraints are enabled | [optional] 
**Tres** | Pointer to [**V0044JobModifyTres**](V0044JobModifyTres.md) |  | [optional] 
**Wckey** | Pointer to **string** | Workload characterization key | [optional] 

## Methods

### NewV0044JobModify

`func NewV0044JobModify() *V0044JobModify`

NewV0044JobModify instantiates a new V0044JobModify object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0044JobModifyWithDefaults

`func NewV0044JobModifyWithDefaults() *V0044JobModify`

NewV0044JobModifyWithDefaults instantiates a new V0044JobModify object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *V0044JobModify) GetComment() V0044JobComment`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *V0044JobModify) GetCommentOk() (*V0044JobComment, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *V0044JobModify) SetComment(v V0044JobComment)`

SetComment sets Comment field to given value.

### HasComment

`func (o *V0044JobModify) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDerivedExitCode

`func (o *V0044JobModify) GetDerivedExitCode() V0044ProcessExitCodeVerbose`

GetDerivedExitCode returns the DerivedExitCode field if non-nil, zero value otherwise.

### GetDerivedExitCodeOk

`func (o *V0044JobModify) GetDerivedExitCodeOk() (*V0044ProcessExitCodeVerbose, bool)`

GetDerivedExitCodeOk returns a tuple with the DerivedExitCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerivedExitCode

`func (o *V0044JobModify) SetDerivedExitCode(v V0044ProcessExitCodeVerbose)`

SetDerivedExitCode sets DerivedExitCode field to given value.

### HasDerivedExitCode

`func (o *V0044JobModify) HasDerivedExitCode() bool`

HasDerivedExitCode returns a boolean if a field has been set.

### GetExtra

`func (o *V0044JobModify) GetExtra() string`

GetExtra returns the Extra field if non-nil, zero value otherwise.

### GetExtraOk

`func (o *V0044JobModify) GetExtraOk() (*string, bool)`

GetExtraOk returns a tuple with the Extra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtra

`func (o *V0044JobModify) SetExtra(v string)`

SetExtra sets Extra field to given value.

### HasExtra

`func (o *V0044JobModify) HasExtra() bool`

HasExtra returns a boolean if a field has been set.

### GetTres

`func (o *V0044JobModify) GetTres() V0044JobModifyTres`

GetTres returns the Tres field if non-nil, zero value otherwise.

### GetTresOk

`func (o *V0044JobModify) GetTresOk() (*V0044JobModifyTres, bool)`

GetTresOk returns a tuple with the Tres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTres

`func (o *V0044JobModify) SetTres(v V0044JobModifyTres)`

SetTres sets Tres field to given value.

### HasTres

`func (o *V0044JobModify) HasTres() bool`

HasTres returns a boolean if a field has been set.

### GetWckey

`func (o *V0044JobModify) GetWckey() string`

GetWckey returns the Wckey field if non-nil, zero value otherwise.

### GetWckeyOk

`func (o *V0044JobModify) GetWckeyOk() (*string, bool)`

GetWckeyOk returns a tuple with the Wckey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWckey

`func (o *V0044JobModify) SetWckey(v string)`

SetWckey sets Wckey field to given value.

### HasWckey

`func (o *V0044JobModify) HasWckey() bool`

HasWckey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


