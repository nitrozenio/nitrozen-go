# Project

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Slug** | Pointer to **string** |  | [optional] 
**PublicUrl** | Pointer to **string** | Public changelog page URL. &#x60;https://&lt;custom_domain&gt;&#x60; when the project&#39;s custom domain is verified, otherwise &#x60;https://nitrozen.io/c/&lt;slug&gt;&#x60;. Read-only; ignored on create and update. | [optional] [readonly] 
**Description** | Pointer to **NullableString** |  | [optional] 
**CustomDomain** | Pointer to **NullableString** |  | [optional] 
**DomainVerificationStatus** | Pointer to **NullableString** |  | [optional] 
**EntriesCount** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewProject

`func NewProject() *Project`

NewProject instantiates a new Project object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectWithDefaults

`func NewProjectWithDefaults() *Project`

NewProjectWithDefaults instantiates a new Project object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Project) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Project) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Project) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Project) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Project) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Project) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Project) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Project) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSlug

`func (o *Project) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *Project) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *Project) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *Project) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetPublicUrl

`func (o *Project) GetPublicUrl() string`

GetPublicUrl returns the PublicUrl field if non-nil, zero value otherwise.

### GetPublicUrlOk

`func (o *Project) GetPublicUrlOk() (*string, bool)`

GetPublicUrlOk returns a tuple with the PublicUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicUrl

`func (o *Project) SetPublicUrl(v string)`

SetPublicUrl sets PublicUrl field to given value.

### HasPublicUrl

`func (o *Project) HasPublicUrl() bool`

HasPublicUrl returns a boolean if a field has been set.

### GetDescription

`func (o *Project) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Project) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Project) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Project) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Project) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Project) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCustomDomain

`func (o *Project) GetCustomDomain() string`

GetCustomDomain returns the CustomDomain field if non-nil, zero value otherwise.

### GetCustomDomainOk

`func (o *Project) GetCustomDomainOk() (*string, bool)`

GetCustomDomainOk returns a tuple with the CustomDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDomain

`func (o *Project) SetCustomDomain(v string)`

SetCustomDomain sets CustomDomain field to given value.

### HasCustomDomain

`func (o *Project) HasCustomDomain() bool`

HasCustomDomain returns a boolean if a field has been set.

### SetCustomDomainNil

`func (o *Project) SetCustomDomainNil(b bool)`

 SetCustomDomainNil sets the value for CustomDomain to be an explicit nil

### UnsetCustomDomain
`func (o *Project) UnsetCustomDomain()`

UnsetCustomDomain ensures that no value is present for CustomDomain, not even an explicit nil
### GetDomainVerificationStatus

`func (o *Project) GetDomainVerificationStatus() string`

GetDomainVerificationStatus returns the DomainVerificationStatus field if non-nil, zero value otherwise.

### GetDomainVerificationStatusOk

`func (o *Project) GetDomainVerificationStatusOk() (*string, bool)`

GetDomainVerificationStatusOk returns a tuple with the DomainVerificationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainVerificationStatus

`func (o *Project) SetDomainVerificationStatus(v string)`

SetDomainVerificationStatus sets DomainVerificationStatus field to given value.

### HasDomainVerificationStatus

`func (o *Project) HasDomainVerificationStatus() bool`

HasDomainVerificationStatus returns a boolean if a field has been set.

### SetDomainVerificationStatusNil

`func (o *Project) SetDomainVerificationStatusNil(b bool)`

 SetDomainVerificationStatusNil sets the value for DomainVerificationStatus to be an explicit nil

### UnsetDomainVerificationStatus
`func (o *Project) UnsetDomainVerificationStatus()`

UnsetDomainVerificationStatus ensures that no value is present for DomainVerificationStatus, not even an explicit nil
### GetEntriesCount

`func (o *Project) GetEntriesCount() int32`

GetEntriesCount returns the EntriesCount field if non-nil, zero value otherwise.

### GetEntriesCountOk

`func (o *Project) GetEntriesCountOk() (*int32, bool)`

GetEntriesCountOk returns a tuple with the EntriesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntriesCount

`func (o *Project) SetEntriesCount(v int32)`

SetEntriesCount sets EntriesCount field to given value.

### HasEntriesCount

`func (o *Project) HasEntriesCount() bool`

HasEntriesCount returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Project) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Project) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Project) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Project) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Project) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Project) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Project) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Project) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


