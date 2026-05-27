# V0044CronEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Flags** | Pointer to **[]string** | Flags | [optional] 
**Minute** | Pointer to **string** | Ranged string specifying eligible minute values (e.g. 0-10,50) | [optional] 
**Hour** | Pointer to **string** | Ranged string specifying eligible hour values (e.g. 0-5,23) | [optional] 
**DayOfMonth** | Pointer to **string** | Ranged string specifying eligible day of month values (e.g. 0-10,29) | [optional] 
**Month** | Pointer to **string** | Ranged string specifying eligible month values (e.g. 0-5,12) | [optional] 
**DayOfWeek** | Pointer to **string** | Ranged string specifying eligible day of week values (e.g.0-3,7) | [optional] 
**Specification** | Pointer to **string** | Complete time specification (* means valid for all allowed values) - minute hour day_of_month month day_of_week | [optional] 
**Command** | Pointer to **string** | Command to run | [optional] 
**Line** | Pointer to [**V0044CronEntryLine**](V0044CronEntryLine.md) |  | [optional] 

## Methods

### NewV0044CronEntry

`func NewV0044CronEntry() *V0044CronEntry`

NewV0044CronEntry instantiates a new V0044CronEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0044CronEntryWithDefaults

`func NewV0044CronEntryWithDefaults() *V0044CronEntry`

NewV0044CronEntryWithDefaults instantiates a new V0044CronEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlags

`func (o *V0044CronEntry) GetFlags() []string`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *V0044CronEntry) GetFlagsOk() (*[]string, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *V0044CronEntry) SetFlags(v []string)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *V0044CronEntry) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetMinute

`func (o *V0044CronEntry) GetMinute() string`

GetMinute returns the Minute field if non-nil, zero value otherwise.

### GetMinuteOk

`func (o *V0044CronEntry) GetMinuteOk() (*string, bool)`

GetMinuteOk returns a tuple with the Minute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinute

`func (o *V0044CronEntry) SetMinute(v string)`

SetMinute sets Minute field to given value.

### HasMinute

`func (o *V0044CronEntry) HasMinute() bool`

HasMinute returns a boolean if a field has been set.

### GetHour

`func (o *V0044CronEntry) GetHour() string`

GetHour returns the Hour field if non-nil, zero value otherwise.

### GetHourOk

`func (o *V0044CronEntry) GetHourOk() (*string, bool)`

GetHourOk returns a tuple with the Hour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHour

`func (o *V0044CronEntry) SetHour(v string)`

SetHour sets Hour field to given value.

### HasHour

`func (o *V0044CronEntry) HasHour() bool`

HasHour returns a boolean if a field has been set.

### GetDayOfMonth

`func (o *V0044CronEntry) GetDayOfMonth() string`

GetDayOfMonth returns the DayOfMonth field if non-nil, zero value otherwise.

### GetDayOfMonthOk

`func (o *V0044CronEntry) GetDayOfMonthOk() (*string, bool)`

GetDayOfMonthOk returns a tuple with the DayOfMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfMonth

`func (o *V0044CronEntry) SetDayOfMonth(v string)`

SetDayOfMonth sets DayOfMonth field to given value.

### HasDayOfMonth

`func (o *V0044CronEntry) HasDayOfMonth() bool`

HasDayOfMonth returns a boolean if a field has been set.

### GetMonth

`func (o *V0044CronEntry) GetMonth() string`

GetMonth returns the Month field if non-nil, zero value otherwise.

### GetMonthOk

`func (o *V0044CronEntry) GetMonthOk() (*string, bool)`

GetMonthOk returns a tuple with the Month field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonth

`func (o *V0044CronEntry) SetMonth(v string)`

SetMonth sets Month field to given value.

### HasMonth

`func (o *V0044CronEntry) HasMonth() bool`

HasMonth returns a boolean if a field has been set.

### GetDayOfWeek

`func (o *V0044CronEntry) GetDayOfWeek() string`

GetDayOfWeek returns the DayOfWeek field if non-nil, zero value otherwise.

### GetDayOfWeekOk

`func (o *V0044CronEntry) GetDayOfWeekOk() (*string, bool)`

GetDayOfWeekOk returns a tuple with the DayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfWeek

`func (o *V0044CronEntry) SetDayOfWeek(v string)`

SetDayOfWeek sets DayOfWeek field to given value.

### HasDayOfWeek

`func (o *V0044CronEntry) HasDayOfWeek() bool`

HasDayOfWeek returns a boolean if a field has been set.

### GetSpecification

`func (o *V0044CronEntry) GetSpecification() string`

GetSpecification returns the Specification field if non-nil, zero value otherwise.

### GetSpecificationOk

`func (o *V0044CronEntry) GetSpecificationOk() (*string, bool)`

GetSpecificationOk returns a tuple with the Specification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecification

`func (o *V0044CronEntry) SetSpecification(v string)`

SetSpecification sets Specification field to given value.

### HasSpecification

`func (o *V0044CronEntry) HasSpecification() bool`

HasSpecification returns a boolean if a field has been set.

### GetCommand

`func (o *V0044CronEntry) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *V0044CronEntry) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *V0044CronEntry) SetCommand(v string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *V0044CronEntry) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### GetLine

`func (o *V0044CronEntry) GetLine() V0044CronEntryLine`

GetLine returns the Line field if non-nil, zero value otherwise.

### GetLineOk

`func (o *V0044CronEntry) GetLineOk() (*V0044CronEntryLine, bool)`

GetLineOk returns a tuple with the Line field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine

`func (o *V0044CronEntry) SetLine(v V0044CronEntryLine)`

SetLine sets Line field to given value.

### HasLine

`func (o *V0044CronEntry) HasLine() bool`

HasLine returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


