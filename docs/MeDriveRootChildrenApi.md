# \MeDriveRootChildrenApi

All URIs are relative to *https://localhost:9200/graph*

Method | HTTP request | Description
------------- | ------------- | -------------
[**HomeGetChildren**](MeDriveRootChildrenApi.md#HomeGetChildren) | **Get** /v1.0/me/drive/root/children | Get children from drive



## HomeGetChildren

> CollectionOfDriveItems HomeGetChildren(ctx).Select_(select_).Execute()

Get children from drive

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
	select_ := []string{"Select_example"} // []string | Select additional properties to be returned. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MeDriveRootChildrenApi.HomeGetChildren(context.Background()).Select_(select_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeDriveRootChildrenApi.HomeGetChildren``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HomeGetChildren`: CollectionOfDriveItems
	fmt.Fprintf(os.Stdout, "Response from `MeDriveRootChildrenApi.HomeGetChildren`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiHomeGetChildrenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **select_** | **[]string** | Select additional properties to be returned. | 

### Return type

[**CollectionOfDriveItems**](CollectionOfDriveItems.md)

### Authorization

[openId](../README.md#openId), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

