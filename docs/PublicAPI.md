# \PublicAPI

All URIs are relative to *https://nitrozen.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PublicProjectsSlugEntriesGet**](PublicAPI.md#PublicProjectsSlugEntriesGet) | **Get** /public/projects/{slug}/entries | List published entries (public)



## PublicProjectsSlugEntriesGet

> PublicProjectsSlugEntriesGet200Response PublicProjectsSlugEntriesGet(ctx, slug).Execute()

List published entries (public)



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
	slug := "slug_example" // string | Project slug

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicAPI.PublicProjectsSlugEntriesGet(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicAPI.PublicProjectsSlugEntriesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PublicProjectsSlugEntriesGet`: PublicProjectsSlugEntriesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `PublicAPI.PublicProjectsSlugEntriesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** | Project slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiPublicProjectsSlugEntriesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PublicProjectsSlugEntriesGet200Response**](PublicProjectsSlugEntriesGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

