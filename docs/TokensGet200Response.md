# TokensGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]TokenInfo**](TokenInfo.md) |  | [optional] 

## Methods

### NewTokensGet200Response

`func NewTokensGet200Response() *TokensGet200Response`

NewTokensGet200Response instantiates a new TokensGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokensGet200ResponseWithDefaults

`func NewTokensGet200ResponseWithDefaults() *TokensGet200Response`

NewTokensGet200ResponseWithDefaults instantiates a new TokensGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *TokensGet200Response) GetData() []TokenInfo`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TokensGet200Response) GetDataOk() (*[]TokenInfo, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TokensGet200Response) SetData(v []TokenInfo)`

SetData sets Data field to given value.

### HasData

`func (o *TokensGet200Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


