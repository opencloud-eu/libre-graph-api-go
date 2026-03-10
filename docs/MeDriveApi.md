# \MeDriveApi

All URIs are relative to *https://localhost:9200/graph*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FollowDriveItem**](MeDriveApi.md#FollowDriveItem) | **Post** /v1.0/me/drive/items/{item-id}/follow | Follow a DriveItem
[**GetHome**](MeDriveApi.md#GetHome) | **Get** /v1.0/me/drive | Get personal space for user
[**ListSharedByMe**](MeDriveApi.md#ListSharedByMe) | **Get** /v1beta1/me/drive/sharedByMe | Get a list of driveItem objects shared by the current user.
[**ListSharedWithMe**](MeDriveApi.md#ListSharedWithMe) | **Get** /v1beta1/me/drive/sharedWithMe | Get a list of driveItem objects shared with the owner of a drive.
[**UnfollowDriveItem**](MeDriveApi.md#UnfollowDriveItem) | **Delete** /v1.0/me/drive/following/{item-id} | Unfollow a DriveItem



## FollowDriveItem

> DriveItem FollowDriveItem(ctx, itemId).Execute()

Follow a DriveItem



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/opencloud-eu/libre-graph-api-go"
)

func main() {
	itemId := "itemId_example" // string | key: id of item

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MeDriveApi.FollowDriveItem(context.Background(), itemId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeDriveApi.FollowDriveItem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FollowDriveItem`: DriveItem
	fmt.Fprintf(os.Stdout, "Response from `MeDriveApi.FollowDriveItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** | key: id of item | 

### Other Parameters

Other parameters are passed through a pointer to a apiFollowDriveItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DriveItem**](DriveItem.md)

### Authorization

[openId](../README.md#openId), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHome

> Drive GetHome(ctx).Execute()

Get personal space for user

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/opencloud-eu/libre-graph-api-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MeDriveApi.GetHome(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeDriveApi.GetHome``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHome`: Drive
	fmt.Fprintf(os.Stdout, "Response from `MeDriveApi.GetHome`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetHomeRequest struct via the builder pattern


### Return type

[**Drive**](Drive.md)

### Authorization

[openId](../README.md#openId), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSharedByMe

> CollectionOfDriveItems1 ListSharedByMe(ctx).Expand(expand).Execute()

Get a list of driveItem objects shared by the current user.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/opencloud-eu/libre-graph-api-go"
)

func main() {
	expand := []string{"Expand_example"} // []string | Expand related entities (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MeDriveApi.ListSharedByMe(context.Background()).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeDriveApi.ListSharedByMe``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSharedByMe`: CollectionOfDriveItems1
	fmt.Fprintf(os.Stdout, "Response from `MeDriveApi.ListSharedByMe`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSharedByMeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **expand** | **[]string** | Expand related entities | 

### Return type

[**CollectionOfDriveItems1**](CollectionOfDriveItems1.md)

### Authorization

[openId](../README.md#openId), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSharedWithMe

> CollectionOfDriveItems1 ListSharedWithMe(ctx).Expand(expand).Execute()

Get a list of driveItem objects shared with the owner of a drive.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/opencloud-eu/libre-graph-api-go"
)

func main() {
	expand := []string{"Expand_example"} // []string | Expand related entities (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MeDriveApi.ListSharedWithMe(context.Background()).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeDriveApi.ListSharedWithMe``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSharedWithMe`: CollectionOfDriveItems1
	fmt.Fprintf(os.Stdout, "Response from `MeDriveApi.ListSharedWithMe`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSharedWithMeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **expand** | **[]string** | Expand related entities | 

### Return type

[**CollectionOfDriveItems1**](CollectionOfDriveItems1.md)

### Authorization

[openId](../README.md#openId), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnfollowDriveItem

> UnfollowDriveItem(ctx, itemId).Execute()

Unfollow a DriveItem



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/opencloud-eu/libre-graph-api-go"
)

func main() {
	itemId := "itemId_example" // string | key: id of item

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MeDriveApi.UnfollowDriveItem(context.Background(), itemId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeDriveApi.UnfollowDriveItem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** | key: id of item | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnfollowDriveItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[openId](../README.md#openId), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

