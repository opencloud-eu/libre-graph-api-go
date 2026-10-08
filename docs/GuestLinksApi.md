# \GuestLinksApi

All URIs are relative to *https://localhost:9200/graph*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RenewGuestLink**](GuestLinksApi.md#RenewGuestLink) | **Post** /v1beta1/extensions/org.libregraph/guestLinks/renew | Renew a guest link
[**VerifyGuestLinkPin**](GuestLinksApi.md#VerifyGuestLinkPin) | **Post** /v1beta1/extensions/org.libregraph/guestLinks/verify/pin | Verify a guest link PIN
[**VerifyGuestLinkToken**](GuestLinksApi.md#VerifyGuestLinkToken) | **Post** /v1beta1/extensions/org.libregraph/guestLinks/verify/token | Verify a guest link token



## RenewGuestLink

> RenewGuestLink(ctx).GuestLinkRenewRequest(guestLinkRenewRequest).Execute()

Renew a guest link



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
	guestLinkRenewRequest := *openapiclient.NewGuestLinkRenewRequest("PermissionId_example") // GuestLinkRenewRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GuestLinksApi.RenewGuestLink(context.Background()).GuestLinkRenewRequest(guestLinkRenewRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GuestLinksApi.RenewGuestLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRenewGuestLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **guestLinkRenewRequest** | [**GuestLinkRenewRequest**](GuestLinkRenewRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[openId](../README.md#openId), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifyGuestLinkPin

> GuestLinkSessionResponse VerifyGuestLinkPin(ctx).GuestLinkVerifyPinRequest(guestLinkVerifyPinRequest).Execute()

Verify a guest link PIN



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
	guestLinkVerifyPinRequest := *openapiclient.NewGuestLinkVerifyPinRequest("Pin_example", "PermissionId_example") // GuestLinkVerifyPinRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GuestLinksApi.VerifyGuestLinkPin(context.Background()).GuestLinkVerifyPinRequest(guestLinkVerifyPinRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GuestLinksApi.VerifyGuestLinkPin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VerifyGuestLinkPin`: GuestLinkSessionResponse
	fmt.Fprintf(os.Stdout, "Response from `GuestLinksApi.VerifyGuestLinkPin`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVerifyGuestLinkPinRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **guestLinkVerifyPinRequest** | [**GuestLinkVerifyPinRequest**](GuestLinkVerifyPinRequest.md) |  | 

### Return type

[**GuestLinkSessionResponse**](GuestLinkSessionResponse.md)

### Authorization

[openId](../README.md#openId), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifyGuestLinkToken

> GuestLinkSessionResponse VerifyGuestLinkToken(ctx).GuestLinkVerifyTokenRequest(guestLinkVerifyTokenRequest).Execute()

Verify a guest link token



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
	guestLinkVerifyTokenRequest := *openapiclient.NewGuestLinkVerifyTokenRequest("Token_example") // GuestLinkVerifyTokenRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GuestLinksApi.VerifyGuestLinkToken(context.Background()).GuestLinkVerifyTokenRequest(guestLinkVerifyTokenRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GuestLinksApi.VerifyGuestLinkToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VerifyGuestLinkToken`: GuestLinkSessionResponse
	fmt.Fprintf(os.Stdout, "Response from `GuestLinksApi.VerifyGuestLinkToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVerifyGuestLinkTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **guestLinkVerifyTokenRequest** | [**GuestLinkVerifyTokenRequest**](GuestLinkVerifyTokenRequest.md) |  | 

### Return type

[**GuestLinkSessionResponse**](GuestLinkSessionResponse.md)

### Authorization

[openId](../README.md#openId), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

