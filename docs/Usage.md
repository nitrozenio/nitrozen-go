# Usage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectsCount** | Pointer to **int32** |  | [optional] 
**EntriesThisMonth** | Pointer to **int32** |  | [optional] 
**Limits** | Pointer to [**UsageLimits**](UsageLimits.md) |  | [optional] 
**PlanName** | Pointer to **string** |  | [optional] 

## Methods

### NewUsage

`func NewUsage() *Usage`

NewUsage instantiates a new Usage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageWithDefaults

`func NewUsageWithDefaults() *Usage`

NewUsageWithDefaults instantiates a new Usage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectsCount

`func (o *Usage) GetProjectsCount() int32`

GetProjectsCount returns the ProjectsCount field if non-nil, zero value otherwise.

### GetProjectsCountOk

`func (o *Usage) GetProjectsCountOk() (*int32, bool)`

GetProjectsCountOk returns a tuple with the ProjectsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectsCount

`func (o *Usage) SetProjectsCount(v int32)`

SetProjectsCount sets ProjectsCount field to given value.

### HasProjectsCount

`func (o *Usage) HasProjectsCount() bool`

HasProjectsCount returns a boolean if a field has been set.

### GetEntriesThisMonth

`func (o *Usage) GetEntriesThisMonth() int32`

GetEntriesThisMonth returns the EntriesThisMonth field if non-nil, zero value otherwise.

### GetEntriesThisMonthOk

`func (o *Usage) GetEntriesThisMonthOk() (*int32, bool)`

GetEntriesThisMonthOk returns a tuple with the EntriesThisMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntriesThisMonth

`func (o *Usage) SetEntriesThisMonth(v int32)`

SetEntriesThisMonth sets EntriesThisMonth field to given value.

### HasEntriesThisMonth

`func (o *Usage) HasEntriesThisMonth() bool`

HasEntriesThisMonth returns a boolean if a field has been set.

### GetLimits

`func (o *Usage) GetLimits() UsageLimits`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *Usage) GetLimitsOk() (*UsageLimits, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *Usage) SetLimits(v UsageLimits)`

SetLimits sets Limits field to given value.

### HasLimits

`func (o *Usage) HasLimits() bool`

HasLimits returns a boolean if a field has been set.

### GetPlanName

`func (o *Usage) GetPlanName() string`

GetPlanName returns the PlanName field if non-nil, zero value otherwise.

### GetPlanNameOk

`func (o *Usage) GetPlanNameOk() (*string, bool)`

GetPlanNameOk returns a tuple with the PlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanName

`func (o *Usage) SetPlanName(v string)`

SetPlanName sets PlanName field to given value.

### HasPlanName

`func (o *Usage) HasPlanName() bool`

HasPlanName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


