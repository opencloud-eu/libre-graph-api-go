# \DrivesPermissionsApi

All URIs are relative to *https://localhost:9200/graph*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLink**](DrivesPermissionsApi.md#CreateLink) | **Post** /v1beta1/drives/{drive-id}/items/{item-id}/createLink | Create a sharing link for a DriveItem
[**DeletePermission**](DrivesPermissionsApi.md#DeletePermission) | **Delete** /v1beta1/drives/{drive-id}/items/{item-id}/permissions/{perm-id} | Remove access to a DriveItem
[**GetPermission**](DrivesPermissionsApi.md#GetPermission) | **Get** /v1beta1/drives/{drive-id}/items/{item-id}/permissions/{perm-id} | Get sharing permission for a file or folder
[**Invite**](DrivesPermissionsApi.md#Invite) | **Post** /v1beta1/drives/{drive-id}/items/{item-id}/invite | Send a sharing invitation
[**ListPermissions**](DrivesPermissionsApi.md#ListPermissions) | **Get** /v1beta1/drives/{drive-id}/items/{item-id}/permissions | List the effective sharing permissions on a driveItem.
[**SetPermissionPassword**](DrivesPermissionsApi.md#SetPermissionPassword) | **Post** /v1beta1/drives/{drive-id}/items/{item-id}/permissions/{perm-id}/setPassword | Set sharing link password
[**UpdatePermission**](DrivesPermissionsApi.md#UpdatePermission) | **Patch** /v1beta1/drives/{drive-id}/items/{item-id}/permissions/{perm-id} | Update sharing permission



## CreateLink

> Permission CreateLink(ctx, driveId, itemId).DriveItemCreateLink(driveItemCreateLink).Execute()

Create a sharing link for a DriveItem



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
	driveItemCreateLink := *openapiclient.NewDriveItemCreateLink() // DriveItemCreateLink | In the request body, provide a JSON object with the following parameters. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DrivesPermissionsApi.CreateLink(context.Background(), driveId, itemId).DriveItemCreateLink(driveItemCreateLink).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DrivesPermissionsApi.CreateLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateLink`: Permission
	fmt.Fprintf(os.Stdout, "Response from `DrivesPermissionsApi.CreateLink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveId** | **string** | key: id of drive | 
**itemId** | **string** | key: id of item | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **driveItemCreateLink** | [**DriveItemCreateLink**](DriveItemCreateLink.md) | In the request body, provide a JSON object with the following parameters. | 

### Return type

[**Permission**](Permission.md)

### Authorization

[openId](../README.md#openId), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePermission

> DeletePermission(ctx, driveId, itemId, permId).Execute()

Remove access to a DriveItem



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
	permId := "permId_example" // string | key: id of permission

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DrivesPermissionsApi.DeletePermission(context.Background(), driveId, itemId, permId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DrivesPermissionsApi.DeletePermission``: %v\n", err)
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
**permId** | **string** | key: id of permission | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePermissionRequest struct via the builder pattern


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


## GetPermission

> Permission GetPermission(ctx, driveId, itemId, permId).Execute()

Get sharing permission for a file or folder



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
	permId := "permId_example" // string | key: id of permission

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DrivesPermissionsApi.GetPermission(context.Background(), driveId, itemId, permId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DrivesPermissionsApi.GetPermission``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPermission`: Permission
	fmt.Fprintf(os.Stdout, "Response from `DrivesPermissionsApi.GetPermission`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveId** | **string** | key: id of drive | 
**itemId** | **string** | key: id of item | 
**permId** | **string** | key: id of permission | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPermissionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Permission**](Permission.md)

### Authorization

[openId](../README.md#openId), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Invite

> CollectionOfPermissions Invite(ctx, driveId, itemId).DriveItemInvite(driveItemInvite).Execute()

Send a sharing invitation



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
	driveItemInvite := *openapiclient.NewDriveItemInvite() // DriveItemInvite | In the request body, provide a JSON object with the following parameters. To create a custom role submit a list of actions instead of roles. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DrivesPermissionsApi.Invite(context.Background(), driveId, itemId).DriveItemInvite(driveItemInvite).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DrivesPermissionsApi.Invite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Invite`: CollectionOfPermissions
	fmt.Fprintf(os.Stdout, "Response from `DrivesPermissionsApi.Invite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveId** | **string** | key: id of drive | 
**itemId** | **string** | key: id of item | 

### Other Parameters

Other parameters are passed through a pointer to a apiInviteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **driveItemInvite** | [**DriveItemInvite**](DriveItemInvite.md) | In the request body, provide a JSON object with the following parameters. To create a custom role submit a list of actions instead of roles. | 

### Return type

[**CollectionOfPermissions**](CollectionOfPermissions.md)

### Authorization

[openId](../README.md#openId), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPermissions

> CollectionOfPermissionsWithAllowedValues ListPermissions(ctx, driveId, itemId).Filter(filter).Select_(select_).Count(count).Top(top).Execute()

List the effective sharing permissions on a driveItem.



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
	filter := "@libre.graph.permissions.roles.allowedValues/rolePermissions/any(p:contains(p/condition, '@Subject.UserType=="Federated"'))" // string | Filter items by property values. By default all permissions are returned and the avalable sharing roles are limited to normal users. To get a list of sharing roles applicable to federated users use the example $select query and combine it with $filter to omit the list of permissions. (optional)
	select_ := []string{"Select_example"} // []string | Select properties to be returned. By default all properties are returned. Select the roles property to fetch the available sharing roles without resolving all the permissions. Combine this with the $filter parameter to fetch the actions applicable to federated users. (optional)
	count := true // bool | Include count of items (optional)
	top := int32(50) // int32 | Show only the first n items (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DrivesPermissionsApi.ListPermissions(context.Background(), driveId, itemId).Filter(filter).Select_(select_).Count(count).Top(top).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DrivesPermissionsApi.ListPermissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPermissions`: CollectionOfPermissionsWithAllowedValues
	fmt.Fprintf(os.Stdout, "Response from `DrivesPermissionsApi.ListPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveId** | **string** | key: id of drive | 
**itemId** | **string** | key: id of item | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **string** | Filter items by property values. By default all permissions are returned and the avalable sharing roles are limited to normal users. To get a list of sharing roles applicable to federated users use the example $select query and combine it with $filter to omit the list of permissions. | 
 **select_** | **[]string** | Select properties to be returned. By default all properties are returned. Select the roles property to fetch the available sharing roles without resolving all the permissions. Combine this with the $filter parameter to fetch the actions applicable to federated users. | 
 **count** | **bool** | Include count of items | 
 **top** | **int32** | Show only the first n items | 

### Return type

[**CollectionOfPermissionsWithAllowedValues**](CollectionOfPermissionsWithAllowedValues.md)

### Authorization

[openId](../README.md#openId), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetPermissionPassword

> Permission SetPermissionPassword(ctx, driveId, itemId, permId).SharingLinkPassword(sharingLinkPassword).Execute()

Set sharing link password



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
	permId := "permId_example" // string | key: id of permission
	sharingLinkPassword := *openapiclient.NewSharingLinkPassword() // SharingLinkPassword | New password value

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DrivesPermissionsApi.SetPermissionPassword(context.Background(), driveId, itemId, permId).SharingLinkPassword(sharingLinkPassword).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DrivesPermissionsApi.SetPermissionPassword``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetPermissionPassword`: Permission
	fmt.Fprintf(os.Stdout, "Response from `DrivesPermissionsApi.SetPermissionPassword`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveId** | **string** | key: id of drive | 
**itemId** | **string** | key: id of item | 
**permId** | **string** | key: id of permission | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetPermissionPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **sharingLinkPassword** | [**SharingLinkPassword**](SharingLinkPassword.md) | New password value | 

### Return type

[**Permission**](Permission.md)

### Authorization

[openId](../README.md#openId), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePermission

> Permission UpdatePermission(ctx, driveId, itemId, permId).Permission(permission).Execute()

Update sharing permission



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
	permId := "permId_example" // string | key: id of permission
	permission := *openapiclient.NewPermission() // Permission | New property values

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DrivesPermissionsApi.UpdatePermission(context.Background(), driveId, itemId, permId).Permission(permission).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DrivesPermissionsApi.UpdatePermission``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePermission`: Permission
	fmt.Fprintf(os.Stdout, "Response from `DrivesPermissionsApi.UpdatePermission`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveId** | **string** | key: id of drive | 
**itemId** | **string** | key: id of item | 
**permId** | **string** | key: id of permission | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePermissionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **permission** | [**Permission**](Permission.md) | New property values | 

### Return type

[**Permission**](Permission.md)

### Authorization

[openId](../README.md#openId), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

