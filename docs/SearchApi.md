# \SearchApi

All URIs are relative to *https://localhost:9200/graph*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchQuery**](SearchApi.md#SearchQuery) | **Post** /v1beta1/search/query | Search for resources



## SearchQuery

> SearchQuery200Response SearchQuery(ctx).SearchQueryRequest(searchQueryRequest).Expand(expand).Execute()

Search for resources



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
	searchQueryRequest := *openapiclient.NewSearchQueryRequest([]openapiclient.SearchRequest{*openapiclient.NewSearchRequest([]string{"EntityTypes_example"}, *openapiclient.NewSearchQuery("QueryString_example"))}) // SearchQueryRequest | 
	expand := []string{"Expand_example"} // []string | Relationships to expand inline on each hit's driveItem. Only `thumbnails` is supported, attaching a preview thumbnail set for thumbnailable mime types. Libregraph extension: MS Graph search has no $expand and returns no thumbnails on search hits.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchApi.SearchQuery(context.Background()).SearchQueryRequest(searchQueryRequest).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.SearchQuery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchQuery`: SearchQuery200Response
	fmt.Fprintf(os.Stdout, "Response from `SearchApi.SearchQuery`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchQueryRequest** | [**SearchQueryRequest**](SearchQueryRequest.md) |  | 
 **expand** | **[]string** | Relationships to expand inline on each hit&#39;s driveItem. Only &#x60;thumbnails&#x60; is supported, attaching a preview thumbnail set for thumbnailable mime types. Libregraph extension: MS Graph search has no $expand and returns no thumbnails on search hits.  | 

### Return type

[**SearchQuery200Response**](SearchQuery200Response.md)

### Authorization

[openId](../README.md#openId), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

