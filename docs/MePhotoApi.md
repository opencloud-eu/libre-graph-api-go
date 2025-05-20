# \MePhotoApi

All URIs are relative to *https://localhost:9200/graph*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteOwnUserPhoto**](MePhotoApi.md#DeleteOwnUserPhoto) | **Delete** /v1.0/me/photo/$value | Delete the current user&#39;s profile photo
[**GetOwnUserPhoto**](MePhotoApi.md#GetOwnUserPhoto) | **Get** /v1.0/me/photo/$value | Get the current user&#39;s profile photo
[**UpdateOwnUserPhotoPatch**](MePhotoApi.md#UpdateOwnUserPhotoPatch) | **Patch** /v1.0/me/photo/$value | Update the current user&#39;s profile photo
[**UpdateOwnUserPhotoPut**](MePhotoApi.md#UpdateOwnUserPhotoPut) | **Put** /v1.0/me/photo/$value | Update the current user&#39;s profile photo



## DeleteOwnUserPhoto

> DeleteOwnUserPhoto(ctx).Execute()

Delete the current user's profile photo

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
	r, err := apiClient.MePhotoApi.DeleteOwnUserPhoto(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MePhotoApi.DeleteOwnUserPhoto``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOwnUserPhotoRequest struct via the builder pattern


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


## GetOwnUserPhoto

> *os.File GetOwnUserPhoto(ctx).Execute()

Get the current user's profile photo

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
	resp, r, err := apiClient.MePhotoApi.GetOwnUserPhoto(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MePhotoApi.GetOwnUserPhoto``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOwnUserPhoto`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `MePhotoApi.GetOwnUserPhoto`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetOwnUserPhotoRequest struct via the builder pattern


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


## UpdateOwnUserPhotoPatch

> UpdateOwnUserPhotoPatch(ctx).Body(body).Execute()

Update the current user's profile photo

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
	body := os.NewFile(1234, "some_file") // *os.File | New user photo (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MePhotoApi.UpdateOwnUserPhotoPatch(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MePhotoApi.UpdateOwnUserPhotoPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOwnUserPhotoPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | ***os.File** | New user photo | 

### Return type

 (empty response body)

### Authorization

[openId](../README.md#openId), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: image/jpeg
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOwnUserPhotoPut

> UpdateOwnUserPhotoPut(ctx).Body(body).Execute()

Update the current user's profile photo

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
	body := os.NewFile(1234, "some_file") // *os.File | New user photo (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MePhotoApi.UpdateOwnUserPhotoPut(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MePhotoApi.UpdateOwnUserPhotoPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOwnUserPhotoPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | ***os.File** | New user photo | 

### Return type

 (empty response body)

### Authorization

[openId](../README.md#openId), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: image/jpeg
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

