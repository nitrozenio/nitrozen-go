# \EntriesAPI

All URIs are relative to *https://nitrozen.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProjectsProjectEntriesEntryDelete**](EntriesAPI.md#ProjectsProjectEntriesEntryDelete) | **Delete** /projects/{project}/entries/{entry} | Delete an entry
[**ProjectsProjectEntriesEntryGet**](EntriesAPI.md#ProjectsProjectEntriesEntryGet) | **Get** /projects/{project}/entries/{entry} | Get an entry
[**ProjectsProjectEntriesEntryPut**](EntriesAPI.md#ProjectsProjectEntriesEntryPut) | **Put** /projects/{project}/entries/{entry} | Update an entry
[**ProjectsProjectEntriesGet**](EntriesAPI.md#ProjectsProjectEntriesGet) | **Get** /projects/{project}/entries | List entries
[**ProjectsProjectEntriesPost**](EntriesAPI.md#ProjectsProjectEntriesPost) | **Post** /projects/{project}/entries | Create an entry



## ProjectsProjectEntriesEntryDelete

> ProjectsProjectEntriesEntryDelete(ctx, project, entry).Execute()

Delete an entry



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
	project := int32(56) // int32 | Project ID
	entry := int32(56) // int32 | Entry ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EntriesAPI.ProjectsProjectEntriesEntryDelete(context.Background(), project, entry).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntriesAPI.ProjectsProjectEntriesEntryDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **int32** | Project ID | 
**entry** | **int32** | Entry ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsProjectEntriesEntryDeleteRequest struct via the builder pattern


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


## ProjectsProjectEntriesEntryGet

> ProjectsProjectEntriesPost201Response ProjectsProjectEntriesEntryGet(ctx, project, entry).Execute()

Get an entry



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
	project := int32(56) // int32 | Project ID
	entry := int32(56) // int32 | Entry ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntriesAPI.ProjectsProjectEntriesEntryGet(context.Background(), project, entry).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntriesAPI.ProjectsProjectEntriesEntryGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsProjectEntriesEntryGet`: ProjectsProjectEntriesPost201Response
	fmt.Fprintf(os.Stdout, "Response from `EntriesAPI.ProjectsProjectEntriesEntryGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **int32** | Project ID | 
**entry** | **int32** | Entry ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsProjectEntriesEntryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ProjectsProjectEntriesPost201Response**](ProjectsProjectEntriesPost201Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsProjectEntriesEntryPut

> ProjectsProjectEntriesPost201Response ProjectsProjectEntriesEntryPut(ctx, project, entry).EntryInput(entryInput).Execute()

Update an entry



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
	project := int32(56) // int32 | Project ID
	entry := int32(56) // int32 | Entry ID
	entryInput := *openapiclient.NewEntryInput("Title_example", "Content_example", "Category_example") // EntryInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntriesAPI.ProjectsProjectEntriesEntryPut(context.Background(), project, entry).EntryInput(entryInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntriesAPI.ProjectsProjectEntriesEntryPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsProjectEntriesEntryPut`: ProjectsProjectEntriesPost201Response
	fmt.Fprintf(os.Stdout, "Response from `EntriesAPI.ProjectsProjectEntriesEntryPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **int32** | Project ID | 
**entry** | **int32** | Entry ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsProjectEntriesEntryPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **entryInput** | [**EntryInput**](EntryInput.md) |  | 

### Return type

[**ProjectsProjectEntriesPost201Response**](ProjectsProjectEntriesPost201Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsProjectEntriesGet

> ProjectsProjectEntriesGet200Response ProjectsProjectEntriesGet(ctx, project).Category(category).Status(status).Page(page).PerPage(perPage).Execute()

List entries



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
	project := int32(56) // int32 | Project ID
	category := "category_example" // string | Filter by category (optional)
	status := "status_example" // string | Filter by publish status (optional)
	page := int32(56) // int32 | Page number (optional) (default to 1)
	perPage := int32(56) // int32 | Items per page (optional) (default to 15)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntriesAPI.ProjectsProjectEntriesGet(context.Background(), project).Category(category).Status(status).Page(page).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntriesAPI.ProjectsProjectEntriesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsProjectEntriesGet`: ProjectsProjectEntriesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `EntriesAPI.ProjectsProjectEntriesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **int32** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsProjectEntriesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **category** | **string** | Filter by category | 
 **status** | **string** | Filter by publish status | 
 **page** | **int32** | Page number | [default to 1]
 **perPage** | **int32** | Items per page | [default to 15]

### Return type

[**ProjectsProjectEntriesGet200Response**](ProjectsProjectEntriesGet200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsProjectEntriesPost

> ProjectsProjectEntriesPost201Response ProjectsProjectEntriesPost(ctx, project).EntryInput(entryInput).Execute()

Create an entry



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
	project := int32(56) // int32 | Project ID
	entryInput := *openapiclient.NewEntryInput("Title_example", "Content_example", "Category_example") // EntryInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntriesAPI.ProjectsProjectEntriesPost(context.Background(), project).EntryInput(entryInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntriesAPI.ProjectsProjectEntriesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsProjectEntriesPost`: ProjectsProjectEntriesPost201Response
	fmt.Fprintf(os.Stdout, "Response from `EntriesAPI.ProjectsProjectEntriesPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **int32** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsProjectEntriesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **entryInput** | [**EntryInput**](EntryInput.md) |  | 

### Return type

[**ProjectsProjectEntriesPost201Response**](ProjectsProjectEntriesPost201Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

