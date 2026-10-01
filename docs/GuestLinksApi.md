# \GuestLinksApi

All URIs are relative to *https://localhost:9200/graph*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RedeemGuestLink**](GuestLinksApi.md#RedeemGuestLink) | **Post** /v1beta1/extensions/org.libregraph/guestLinks/redeem | Redeem a guest link token



## RedeemGuestLink

> GuestLinkRedeemResponse RedeemGuestLink(ctx).GuestLinkRedeemRequest(guestLinkRedeemRequest).Execute()

Redeem a guest link token



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
	guestLinkRedeemRequest := *openapiclient.NewGuestLinkRedeemRequest("Token_example") // GuestLinkRedeemRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GuestLinksApi.RedeemGuestLink(context.Background()).GuestLinkRedeemRequest(guestLinkRedeemRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GuestLinksApi.RedeemGuestLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RedeemGuestLink`: GuestLinkRedeemResponse
	fmt.Fprintf(os.Stdout, "Response from `GuestLinksApi.RedeemGuestLink`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRedeemGuestLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **guestLinkRedeemRequest** | [**GuestLinkRedeemRequest**](GuestLinkRedeemRequest.md) |  | 

### Return type

[**GuestLinkRedeemResponse**](GuestLinkRedeemResponse.md)

### Authorization

[openId](../README.md#openId), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

