# \UserPhotoApi

All URIs are relative to *https://localhost:9200/graph*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUserPhoto**](UserPhotoApi.md#GetUserPhoto) | **Get** /v1.0/users/{user-id}/photo/$value | Get the photo of a user



## GetUserPhoto

> *os.File GetUserPhoto(ctx, userId).Execute()

Get the photo of a user

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
	userId := "userId_example" // string | key: id or name of user

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserPhotoApi.GetUserPhoto(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserPhotoApi.GetUserPhoto``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserPhoto`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `UserPhotoApi.GetUserPhoto`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id or name of user | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserPhotoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[openId](../README.md#openId), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/jpeg, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

