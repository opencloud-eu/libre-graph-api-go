# \DrivesRootApi

All URIs are relative to *https://localhost:9200/graph*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDriveItem**](DrivesRootApi.md#CreateDriveItem) | **Post** /v1beta1/drives/{drive-id}/root/children | Create a new DriveItem at the drive root
[**CreateLinkSpaceRoot**](DrivesRootApi.md#CreateLinkSpaceRoot) | **Post** /v1beta1/drives/{drive-id}/root/createLink | Create a sharing link for the root item of a Drive
[**DeletePermissionSpaceRoot**](DrivesRootApi.md#DeletePermissionSpaceRoot) | **Delete** /v1beta1/drives/{drive-id}/root/permissions/{perm-id} | Remove access to a Drive
[**GetPermissionSpaceRoot**](DrivesRootApi.md#GetPermissionSpaceRoot) | **Get** /v1beta1/drives/{drive-id}/root/permissions/{perm-id} | Get a single sharing permission for the root item of a drive
[**GetRoot**](DrivesRootApi.md#GetRoot) | **Get** /v1.0/drives/{drive-id}/root | Get root from arbitrary space
[**InviteSpaceRoot**](DrivesRootApi.md#InviteSpaceRoot) | **Post** /v1beta1/drives/{drive-id}/root/invite | Send a sharing invitation
[**ListPermissionsSpaceRoot**](DrivesRootApi.md#ListPermissionsSpaceRoot) | **Get** /v1beta1/drives/{drive-id}/root/permissions | List the effective permissions on the root item of a drive.
[**SetPermissionPasswordSpaceRoot**](DrivesRootApi.md#SetPermissionPasswordSpaceRoot) | **Post** /v1beta1/drives/{drive-id}/root/permissions/{perm-id}/setPassword | Set sharing link password for the root item of a drive
[**UpdatePermissionSpaceRoot**](DrivesRootApi.md#UpdatePermissionSpaceRoot) | **Patch** /v1beta1/drives/{drive-id}/root/permissions/{perm-id} | Update sharing permission



## CreateDriveItem

> DriveItem CreateDriveItem(ctx, driveId).LibreGraphConflictBehavior(libreGraphConflictBehavior).LibreGraphMissingParentsBehavior(libreGraphMissingParentsBehavior).DriveItem(driveItem).Execute()

Create a new DriveItem at the drive root



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
	libreGraphConflictBehavior := "libreGraphConflictBehavior_example" // string | Controls what happens when a child with the same name already exists. `fail` (default) returns 409; `replace` overwrites the existing item. MS Graph's `rename` value is not supported.  (optional) (default to "fail")
	libreGraphMissingParentsBehavior := "libreGraphMissingParentsBehavior_example" // string | Controls what happens when a colon-syntax URL refers to a path whose intermediate folders don't all exist yet. `fail` (default) returns 404; `create` creates the missing intermediate folders before creating the final item. Only meaningful for colon-syntax URLs; ignored otherwise.  (optional) (default to "fail")
	driveItem := *openapiclient.NewDriveItem() // DriveItem | In the request body, provide a JSON object describing the new driveItem. Must specify exactly one of `folder`, `file`, or `remoteItem`. For mount-share, see [sharedWithMe](#/me.drive/ListSharedWithMe) for obtaining the source `remoteItem.id` and `permission` id. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DrivesRootApi.CreateDriveItem(context.Background(), driveId).LibreGraphConflictBehavior(libreGraphConflictBehavior).LibreGraphMissingParentsBehavior(libreGraphMissingParentsBehavior).DriveItem(driveItem).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DrivesRootApi.CreateDriveItem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDriveItem`: DriveItem
	fmt.Fprintf(os.Stdout, "Response from `DrivesRootApi.CreateDriveItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveId** | **string** | key: id of drive | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDriveItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **libreGraphConflictBehavior** | **string** | Controls what happens when a child with the same name already exists. &#x60;fail&#x60; (default) returns 409; &#x60;replace&#x60; overwrites the existing item. MS Graph&#39;s &#x60;rename&#x60; value is not supported.  | [default to &quot;fail&quot;]
 **libreGraphMissingParentsBehavior** | **string** | Controls what happens when a colon-syntax URL refers to a path whose intermediate folders don&#39;t all exist yet. &#x60;fail&#x60; (default) returns 404; &#x60;create&#x60; creates the missing intermediate folders before creating the final item. Only meaningful for colon-syntax URLs; ignored otherwise.  | [default to &quot;fail&quot;]
 **driveItem** | [**DriveItem**](DriveItem.md) | In the request body, provide a JSON object describing the new driveItem. Must specify exactly one of &#x60;folder&#x60;, &#x60;file&#x60;, or &#x60;remoteItem&#x60;. For mount-share, see [sharedWithMe](#/me.drive/ListSharedWithMe) for obtaining the source &#x60;remoteItem.id&#x60; and &#x60;permission&#x60; id. | 

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


## CreateLinkSpaceRoot

> Permission CreateLinkSpaceRoot(ctx, driveId).DriveItemCreateLink(driveItemCreateLink).Execute()

Create a sharing link for the root item of a Drive



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
	driveItemCreateLink := *openapiclient.NewDriveItemCreateLink() // DriveItemCreateLink | In the request body, provide a JSON object with the following parameters. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DrivesRootApi.CreateLinkSpaceRoot(context.Background(), driveId).DriveItemCreateLink(driveItemCreateLink).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DrivesRootApi.CreateLinkSpaceRoot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateLinkSpaceRoot`: Permission
	fmt.Fprintf(os.Stdout, "Response from `DrivesRootApi.CreateLinkSpaceRoot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveId** | **string** | key: id of drive | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLinkSpaceRootRequest struct via the builder pattern


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


## DeletePermissionSpaceRoot

> DeletePermissionSpaceRoot(ctx, driveId, permId).Execute()

Remove access to a Drive



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
	permId := "permId_example" // string | key: id of permission

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DrivesRootApi.DeletePermissionSpaceRoot(context.Background(), driveId, permId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DrivesRootApi.DeletePermissionSpaceRoot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveId** | **string** | key: id of drive | 
**permId** | **string** | key: id of permission | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePermissionSpaceRootRequest struct via the builder pattern


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


## GetPermissionSpaceRoot

> Permission GetPermissionSpaceRoot(ctx, driveId, permId).Execute()

Get a single sharing permission for the root item of a drive



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
	permId := "permId_example" // string | key: id of permission

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DrivesRootApi.GetPermissionSpaceRoot(context.Background(), driveId, permId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DrivesRootApi.GetPermissionSpaceRoot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPermissionSpaceRoot`: Permission
	fmt.Fprintf(os.Stdout, "Response from `DrivesRootApi.GetPermissionSpaceRoot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveId** | **string** | key: id of drive | 
**permId** | **string** | key: id of permission | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPermissionSpaceRootRequest struct via the builder pattern


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


## GetRoot

> DriveItem GetRoot(ctx, driveId).Select_(select_).Expand(expand).Execute()

Get root from arbitrary space

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
	select_ := []string{"Select_example"} // []string | Select additional properties to be returned. (optional)
	expand := []string{"Expand_example"} // []string | Expand related entities to be returned. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DrivesRootApi.GetRoot(context.Background(), driveId).Select_(select_).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DrivesRootApi.GetRoot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRoot`: DriveItem
	fmt.Fprintf(os.Stdout, "Response from `DrivesRootApi.GetRoot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveId** | **string** | key: id of drive | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRootRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select additional properties to be returned. | 
 **expand** | **[]string** | Expand related entities to be returned. | 

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


## InviteSpaceRoot

> CollectionOfPermissions InviteSpaceRoot(ctx, driveId).DriveItemInvite(driveItemInvite).Execute()

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
	driveItemInvite := *openapiclient.NewDriveItemInvite() // DriveItemInvite | In the request body, provide a JSON object with the following parameters. To create a custom role submit a list of actions instead of roles. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DrivesRootApi.InviteSpaceRoot(context.Background(), driveId).DriveItemInvite(driveItemInvite).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DrivesRootApi.InviteSpaceRoot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InviteSpaceRoot`: CollectionOfPermissions
	fmt.Fprintf(os.Stdout, "Response from `DrivesRootApi.InviteSpaceRoot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveId** | **string** | key: id of drive | 

### Other Parameters

Other parameters are passed through a pointer to a apiInviteSpaceRootRequest struct via the builder pattern


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


## ListPermissionsSpaceRoot

> CollectionOfPermissionsWithAllowedValues ListPermissionsSpaceRoot(ctx, driveId).Filter(filter).Select_(select_).Count(count).Top(top).Execute()

List the effective permissions on the root item of a drive.



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
	filter := "@libre.graph.permissions.roles.allowedValues/rolePermissions/any(p:contains(p/condition, '@Subject.UserType==\"Federated\"'))" // string | Filter items by property values. By default all permissions are returned and the avalable sharing roles are limited to normal users. To get a list of sharing roles applicable to federated users use the example $select query and combine it with $filter to omit the list of permissions. (optional)
	select_ := []string{"Select_example"} // []string | Select properties to be returned. By default all properties are returned. Select the roles property to fetch the available sharing roles without resolving all the permissions. Combine this with the $filter parameter to fetch the actions applicable to federated users. (optional)
	count := true // bool | Include count of items (optional)
	top := int32(50) // int32 | Show only the first n items (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DrivesRootApi.ListPermissionsSpaceRoot(context.Background(), driveId).Filter(filter).Select_(select_).Count(count).Top(top).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DrivesRootApi.ListPermissionsSpaceRoot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPermissionsSpaceRoot`: CollectionOfPermissionsWithAllowedValues
	fmt.Fprintf(os.Stdout, "Response from `DrivesRootApi.ListPermissionsSpaceRoot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveId** | **string** | key: id of drive | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPermissionsSpaceRootRequest struct via the builder pattern


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


## SetPermissionPasswordSpaceRoot

> Permission SetPermissionPasswordSpaceRoot(ctx, driveId, permId).SharingLinkPassword(sharingLinkPassword).Execute()

Set sharing link password for the root item of a drive



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
	permId := "permId_example" // string | key: id of permission
	sharingLinkPassword := *openapiclient.NewSharingLinkPassword() // SharingLinkPassword | New password value

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DrivesRootApi.SetPermissionPasswordSpaceRoot(context.Background(), driveId, permId).SharingLinkPassword(sharingLinkPassword).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DrivesRootApi.SetPermissionPasswordSpaceRoot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetPermissionPasswordSpaceRoot`: Permission
	fmt.Fprintf(os.Stdout, "Response from `DrivesRootApi.SetPermissionPasswordSpaceRoot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveId** | **string** | key: id of drive | 
**permId** | **string** | key: id of permission | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetPermissionPasswordSpaceRootRequest struct via the builder pattern


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


## UpdatePermissionSpaceRoot

> Permission UpdatePermissionSpaceRoot(ctx, driveId, permId).Permission(permission).Execute()

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
	permId := "permId_example" // string | key: id of permission
	permission := *openapiclient.NewPermission() // Permission | New property values

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DrivesRootApi.UpdatePermissionSpaceRoot(context.Background(), driveId, permId).Permission(permission).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DrivesRootApi.UpdatePermissionSpaceRoot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePermissionSpaceRoot`: Permission
	fmt.Fprintf(os.Stdout, "Response from `DrivesRootApi.UpdatePermissionSpaceRoot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveId** | **string** | key: id of drive | 
**permId** | **string** | key: id of permission | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePermissionSpaceRootRequest struct via the builder pattern


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

