# \DriveItemApi

All URIs are relative to *https://localhost:9200/graph*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateChildDriveItem**](DriveItemApi.md#CreateChildDriveItem) | **Post** /v1beta1/drives/{drive-id}/items/{item-id}/children | Create a new DriveItem under a parent item
[**DeleteDriveItem**](DriveItemApi.md#DeleteDriveItem) | **Delete** /v1beta1/drives/{drive-id}/items/{item-id} | Delete a DriveItem.
[**GetDriveItem**](DriveItemApi.md#GetDriveItem) | **Get** /v1beta1/drives/{drive-id}/items/{item-id} | Get a DriveItem.
[**GetDriveItemChildren**](DriveItemApi.md#GetDriveItemChildren) | **Get** /v1.0/drives/{drive-id}/items/{item-id}/children | List children of a DriveItem
[**GetDriveItemContent**](DriveItemApi.md#GetDriveItemContent) | **Get** /v1beta1/drives/{drive-id}/items/{item-id}/content | Download the content of a DriveItem
[**UpdateDriveItem**](DriveItemApi.md#UpdateDriveItem) | **Patch** /v1beta1/drives/{drive-id}/items/{item-id} | Update a DriveItem.



## CreateChildDriveItem

> DriveItem CreateChildDriveItem(ctx, driveId, itemId).DriveItem(driveItem).LibreGraphConflictBehavior(libreGraphConflictBehavior).LibreGraphMissingParentsBehavior(libreGraphMissingParentsBehavior).Execute()

Create a new DriveItem under a parent item



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
	driveId := "a0ca6a90-a365-4782-871e-d44447bbc668$a0ca6a90-a365-4782-871e-d44447bbc668" // string | key: id of drive
	itemId := "a0ca6a90-a365-4782-871e-d44447bbc668$a0ca6a90-a365-4782-871e-d44447bbc668!share-id" // string | key: id of item
	driveItem := *openapiclient.NewDriveItem() // DriveItem | In the request body, provide a JSON object describing the new driveItem. Must specify exactly one of `folder` or `file`.
	libreGraphConflictBehavior := "libreGraphConflictBehavior_example" // string | Controls what happens when a child with the same name already exists. `fail` (default) returns 409; `replace` overwrites the existing item. MS Graph's `rename` value is not supported.  (optional) (default to "fail")
	libreGraphMissingParentsBehavior := "libreGraphMissingParentsBehavior_example" // string | Controls what happens when a colon-syntax URL refers to a path whose intermediate folders don't all exist yet. `fail` (default) returns 404; `create` creates the missing intermediate folders before creating the final item. Only meaningful for colon-syntax URLs; ignored otherwise.  (optional) (default to "fail")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DriveItemApi.CreateChildDriveItem(context.Background(), driveId, itemId).DriveItem(driveItem).LibreGraphConflictBehavior(libreGraphConflictBehavior).LibreGraphMissingParentsBehavior(libreGraphMissingParentsBehavior).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DriveItemApi.CreateChildDriveItem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateChildDriveItem`: DriveItem
	fmt.Fprintf(os.Stdout, "Response from `DriveItemApi.CreateChildDriveItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveId** | **string** | key: id of drive | 
**itemId** | **string** | key: id of item | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateChildDriveItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **driveItem** | [**DriveItem**](DriveItem.md) | In the request body, provide a JSON object describing the new driveItem. Must specify exactly one of &#x60;folder&#x60; or &#x60;file&#x60;. | 
 **libreGraphConflictBehavior** | **string** | Controls what happens when a child with the same name already exists. &#x60;fail&#x60; (default) returns 409; &#x60;replace&#x60; overwrites the existing item. MS Graph&#39;s &#x60;rename&#x60; value is not supported.  | [default to &quot;fail&quot;]
 **libreGraphMissingParentsBehavior** | **string** | Controls what happens when a colon-syntax URL refers to a path whose intermediate folders don&#39;t all exist yet. &#x60;fail&#x60; (default) returns 404; &#x60;create&#x60; creates the missing intermediate folders before creating the final item. Only meaningful for colon-syntax URLs; ignored otherwise.  | [default to &quot;fail&quot;]

### Return type

[**DriveItem**](DriveItem.md)

### Authorization

[openId](../README.md#openId), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDriveItem

> DeleteDriveItem(ctx, driveId, itemId).Execute()

Delete a DriveItem.



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
	driveId := "a0ca6a90-a365-4782-871e-d44447bbc668$a0ca6a90-a365-4782-871e-d44447bbc668" // string | key: id of drive
	itemId := "a0ca6a90-a365-4782-871e-d44447bbc668$a0ca6a90-a365-4782-871e-d44447bbc668!share-id" // string | key: id of item

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DriveItemApi.DeleteDriveItem(context.Background(), driveId, itemId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DriveItemApi.DeleteDriveItem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveId** | **string** | key: id of drive | 
**itemId** | **string** | key: id of item | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDriveItemRequest struct via the builder pattern


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


## GetDriveItem

> DriveItem GetDriveItem(ctx, driveId, itemId).Select_(select_).Execute()

Get a DriveItem.



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
	driveId := "a0ca6a90-a365-4782-871e-d44447bbc668$a0ca6a90-a365-4782-871e-d44447bbc668" // string | key: id of drive
	itemId := "a0ca6a90-a365-4782-871e-d44447bbc668$a0ca6a90-a365-4782-871e-d44447bbc668!share-id" // string | key: id of item
	select_ := []string{"Select_example"} // []string | Select additional properties to be returned. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DriveItemApi.GetDriveItem(context.Background(), driveId, itemId).Select_(select_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DriveItemApi.GetDriveItem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDriveItem`: DriveItem
	fmt.Fprintf(os.Stdout, "Response from `DriveItemApi.GetDriveItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveId** | **string** | key: id of drive | 
**itemId** | **string** | key: id of item | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDriveItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | **[]string** | Select additional properties to be returned. | 

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


## GetDriveItemChildren

> CollectionOfDriveItems GetDriveItemChildren(ctx, driveId, itemId).Execute()

List children of a DriveItem



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
	driveId := "a0ca6a90-a365-4782-871e-d44447bbc668$a0ca6a90-a365-4782-871e-d44447bbc668" // string | key: id of drive
	itemId := "a0ca6a90-a365-4782-871e-d44447bbc668$a0ca6a90-a365-4782-871e-d44447bbc668!share-id" // string | key: id of item

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DriveItemApi.GetDriveItemChildren(context.Background(), driveId, itemId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DriveItemApi.GetDriveItemChildren``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDriveItemChildren`: CollectionOfDriveItems
	fmt.Fprintf(os.Stdout, "Response from `DriveItemApi.GetDriveItemChildren`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveId** | **string** | key: id of drive | 
**itemId** | **string** | key: id of item | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDriveItemChildrenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## GetDriveItemContent

> OdataError GetDriveItemContent(ctx, driveId, itemId).Execute()

Download the content of a DriveItem



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
	driveId := "driveId_example" // string | key: id of drive
	itemId := "itemId_example" // string | key: id of item

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DriveItemApi.GetDriveItemContent(context.Background(), driveId, itemId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DriveItemApi.GetDriveItemContent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDriveItemContent`: OdataError
	fmt.Fprintf(os.Stdout, "Response from `DriveItemApi.GetDriveItemContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveId** | **string** | key: id of drive | 
**itemId** | **string** | key: id of item | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDriveItemContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OdataError**](OdataError.md)

### Authorization

[openId](../README.md#openId), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDriveItem

> DriveItem UpdateDriveItem(ctx, driveId, itemId).DriveItem(driveItem).Execute()

Update a DriveItem.



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
	driveId := "a0ca6a90-a365-4782-871e-d44447bbc668$a0ca6a90-a365-4782-871e-d44447bbc668" // string | key: id of drive
	itemId := "a0ca6a90-a365-4782-871e-d44447bbc668$a0ca6a90-a365-4782-871e-d44447bbc668!share-id" // string | key: id of item
	driveItem := *openapiclient.NewDriveItem() // DriveItem | DriveItem properties to update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DriveItemApi.UpdateDriveItem(context.Background(), driveId, itemId).DriveItem(driveItem).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DriveItemApi.UpdateDriveItem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDriveItem`: DriveItem
	fmt.Fprintf(os.Stdout, "Response from `DriveItemApi.UpdateDriveItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveId** | **string** | key: id of drive | 
**itemId** | **string** | key: id of item | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDriveItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **driveItem** | [**DriveItem**](DriveItem.md) | DriveItem properties to update | 

### Return type

[**DriveItem**](DriveItem.md)

### Authorization

[openId](../README.md#openId), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

