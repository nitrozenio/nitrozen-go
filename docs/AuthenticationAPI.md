# \AuthenticationAPI

All URIs are relative to *https://nitrozen.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TokensGet**](AuthenticationAPI.md#TokensGet) | **Get** /tokens | List API tokens
[**TokensPost**](AuthenticationAPI.md#TokensPost) | **Post** /tokens | Create an API token
[**TokensTokenDelete**](AuthenticationAPI.md#TokensTokenDelete) | **Delete** /tokens/{token} | Revoke an API token



## TokensGet

> TokensGet200Response TokensGet(ctx).Execute()

List API tokens



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/nitrozenio/nitrozen-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.TokensGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.TokensGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TokensGet`: TokensGet200Response
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.TokensGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTokensGetRequest struct via the builder pattern


### Return type

[**TokensGet200Response**](TokensGet200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TokensPost

> TokensPost201Response TokensPost(ctx).TokensPostRequest(tokensPostRequest).Execute()

Create an API token



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/nitrozenio/nitrozen-go"
)

func main() {
	tokensPostRequest := *openapiclient.NewTokensPostRequest("Name_example") // TokensPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.TokensPost(context.Background()).TokensPostRequest(tokensPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.TokensPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TokensPost`: TokensPost201Response
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.TokensPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTokensPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokensPostRequest** | [**TokensPostRequest**](TokensPostRequest.md) |  | 

### Return type

[**TokensPost201Response**](TokensPost201Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TokensTokenDelete

> TokensTokenDelete(ctx, token).Execute()

Revoke an API token



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/nitrozenio/nitrozen-go"
)

func main() {
	token := int32(56) // int32 | Token ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AuthenticationAPI.TokensTokenDelete(context.Background(), token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.TokensTokenDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **int32** | Token ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiTokensTokenDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

