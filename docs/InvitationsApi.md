# \InvitationsApi

All URIs are relative to *https://localhost:9200/graph*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateInvitation**](InvitationsApi.md#CreateInvitation) | **Post** /v1.0/invitations | Create a new invitation
[**GetInvitation**](InvitationsApi.md#GetInvitation) | **Get** /v1.0/invitations/{invitation-id} | Get an invitation by key
[**ListInvitations**](InvitationsApi.md#ListInvitations) | **Get** /v1.0/invitations | Get a list of invitations



## CreateInvitation

> Invitation CreateInvitation(ctx).Invitation(invitation).Execute()

Create a new invitation

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
	invitation := *openapiclient.NewInvitation() // Invitation | New invitation (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvitationsApi.CreateInvitation(context.Background()).Invitation(invitation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvitationsApi.CreateInvitation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateInvitation`: Invitation
	fmt.Fprintf(os.Stdout, "Response from `InvitationsApi.CreateInvitation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateInvitationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **invitation** | [**Invitation**](Invitation.md) | New invitation | 

### Return type

[**Invitation**](Invitation.md)

### Authorization

[openId](../README.md#openId), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInvitation

> Invitation GetInvitation(ctx, invitationId).Execute()

Get an invitation by key

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
	invitationId := "invitationId_example" // string | key: id of invitation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvitationsApi.GetInvitation(context.Background(), invitationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvitationsApi.GetInvitation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInvitation`: Invitation
	fmt.Fprintf(os.Stdout, "Response from `InvitationsApi.GetInvitation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invitationId** | **string** | key: id of invitation | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInvitationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Invitation**](Invitation.md)

### Authorization

[openId](../README.md#openId), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInvitations

> CollectionOfInvitations ListInvitations(ctx).Execute()

Get a list of invitations

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
	resp, r, err := apiClient.InvitationsApi.ListInvitations(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvitationsApi.ListInvitations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListInvitations`: CollectionOfInvitations
	fmt.Fprintf(os.Stdout, "Response from `InvitationsApi.ListInvitations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListInvitationsRequest struct via the builder pattern


### Return type

[**CollectionOfInvitations**](CollectionOfInvitations.md)

### Authorization

[openId](../README.md#openId), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

