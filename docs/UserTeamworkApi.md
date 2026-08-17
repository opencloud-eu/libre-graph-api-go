# \UserTeamworkApi

All URIs are relative to *https://localhost:9200/graph*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SendActivityNotification**](UserTeamworkApi.md#SendActivityNotification) | **Post** /v1.0/users/{user-id}/teamwork/sendActivityNotification | Send an activity notification to a user



## SendActivityNotification

> SendActivityNotification(ctx, userId).ActivityNotification(activityNotification).Execute()

Send an activity notification to a user



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
	activityNotification := *openapiclient.NewActivityNotification(*openapiclient.NewActivityTopic("Source_example", "Value_example"), "ActivityType_example", "TeamsAppId_example") // ActivityNotification | The activity the user is notified about.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserTeamworkApi.SendActivityNotification(context.Background(), userId).ActivityNotification(activityNotification).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserTeamworkApi.SendActivityNotification``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id or name of user | 

### Other Parameters

Other parameters are passed through a pointer to a apiSendActivityNotificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **activityNotification** | [**ActivityNotification**](ActivityNotification.md) | The activity the user is notified about. | 

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

