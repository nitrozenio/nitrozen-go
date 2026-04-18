# UserPlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**PlanType** | Pointer to **string** |  | [optional] 
**MaxProjects** | Pointer to **int32** |  | [optional] 
**MaxEntriesPerMonth** | Pointer to **int32** |  | [optional] 
**CustomDomainAllowed** | Pointer to **bool** |  | [optional] 
**AnalyticsEnabled** | Pointer to **bool** |  | [optional] 
**ApiAccessEnabled** | Pointer to **bool** |  | [optional] 
**ApiRateLimit** | Pointer to **int32** | Maximum API requests per minute for this plan | [optional] 

## Methods

### NewUserPlan

`func NewUserPlan() *UserPlan`

NewUserPlan instantiates a new UserPlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserPlanWithDefaults

`func NewUserPlanWithDefaults() *UserPlan`

NewUserPlanWithDefaults instantiates a new UserPlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserPlan) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserPlan) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserPlan) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *UserPlan) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *UserPlan) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserPlan) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserPlan) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UserPlan) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPlanType

`func (o *UserPlan) GetPlanType() string`

GetPlanType returns the PlanType field if non-nil, zero value otherwise.

### GetPlanTypeOk

`func (o *UserPlan) GetPlanTypeOk() (*string, bool)`

GetPlanTypeOk returns a tuple with the PlanType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanType

`func (o *UserPlan) SetPlanType(v string)`

SetPlanType sets PlanType field to given value.

### HasPlanType

`func (o *UserPlan) HasPlanType() bool`

HasPlanType returns a boolean if a field has been set.

### GetMaxProjects

`func (o *UserPlan) GetMaxProjects() int32`

GetMaxProjects returns the MaxProjects field if non-nil, zero value otherwise.

### GetMaxProjectsOk

`func (o *UserPlan) GetMaxProjectsOk() (*int32, bool)`

GetMaxProjectsOk returns a tuple with the MaxProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxProjects

`func (o *UserPlan) SetMaxProjects(v int32)`

SetMaxProjects sets MaxProjects field to given value.

### HasMaxProjects

`func (o *UserPlan) HasMaxProjects() bool`

HasMaxProjects returns a boolean if a field has been set.

### GetMaxEntriesPerMonth

`func (o *UserPlan) GetMaxEntriesPerMonth() int32`

GetMaxEntriesPerMonth returns the MaxEntriesPerMonth field if non-nil, zero value otherwise.

### GetMaxEntriesPerMonthOk

`func (o *UserPlan) GetMaxEntriesPerMonthOk() (*int32, bool)`

GetMaxEntriesPerMonthOk returns a tuple with the MaxEntriesPerMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxEntriesPerMonth

`func (o *UserPlan) SetMaxEntriesPerMonth(v int32)`

SetMaxEntriesPerMonth sets MaxEntriesPerMonth field to given value.

### HasMaxEntriesPerMonth

`func (o *UserPlan) HasMaxEntriesPerMonth() bool`

HasMaxEntriesPerMonth returns a boolean if a field has been set.

### GetCustomDomainAllowed

`func (o *UserPlan) GetCustomDomainAllowed() bool`

GetCustomDomainAllowed returns the CustomDomainAllowed field if non-nil, zero value otherwise.

### GetCustomDomainAllowedOk

`func (o *UserPlan) GetCustomDomainAllowedOk() (*bool, bool)`

GetCustomDomainAllowedOk returns a tuple with the CustomDomainAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDomainAllowed

`func (o *UserPlan) SetCustomDomainAllowed(v bool)`

SetCustomDomainAllowed sets CustomDomainAllowed field to given value.

### HasCustomDomainAllowed

`func (o *UserPlan) HasCustomDomainAllowed() bool`

HasCustomDomainAllowed returns a boolean if a field has been set.

### GetAnalyticsEnabled

`func (o *UserPlan) GetAnalyticsEnabled() bool`

GetAnalyticsEnabled returns the AnalyticsEnabled field if non-nil, zero value otherwise.

### GetAnalyticsEnabledOk

`func (o *UserPlan) GetAnalyticsEnabledOk() (*bool, bool)`

GetAnalyticsEnabledOk returns a tuple with the AnalyticsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyticsEnabled

`func (o *UserPlan) SetAnalyticsEnabled(v bool)`

SetAnalyticsEnabled sets AnalyticsEnabled field to given value.

### HasAnalyticsEnabled

`func (o *UserPlan) HasAnalyticsEnabled() bool`

HasAnalyticsEnabled returns a boolean if a field has been set.

### GetApiAccessEnabled

`func (o *UserPlan) GetApiAccessEnabled() bool`

GetApiAccessEnabled returns the ApiAccessEnabled field if non-nil, zero value otherwise.

### GetApiAccessEnabledOk

`func (o *UserPlan) GetApiAccessEnabledOk() (*bool, bool)`

GetApiAccessEnabledOk returns a tuple with the ApiAccessEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiAccessEnabled

`func (o *UserPlan) SetApiAccessEnabled(v bool)`

SetApiAccessEnabled sets ApiAccessEnabled field to given value.

### HasApiAccessEnabled

`func (o *UserPlan) HasApiAccessEnabled() bool`

HasApiAccessEnabled returns a boolean if a field has been set.

### GetApiRateLimit

`func (o *UserPlan) GetApiRateLimit() int32`

GetApiRateLimit returns the ApiRateLimit field if non-nil, zero value otherwise.

### GetApiRateLimitOk

`func (o *UserPlan) GetApiRateLimitOk() (*int32, bool)`

GetApiRateLimitOk returns a tuple with the ApiRateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiRateLimit

`func (o *UserPlan) SetApiRateLimit(v int32)`

SetApiRateLimit sets ApiRateLimit field to given value.

### HasApiRateLimit

`func (o *UserPlan) HasApiRateLimit() bool`

HasApiRateLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


