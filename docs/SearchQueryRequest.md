# SearchQueryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Requests** | [**[]SearchRequest**](SearchRequest.md) | A collection of one or more search requests. | 

## Methods

### NewSearchQueryRequest

`func NewSearchQueryRequest(requests []SearchRequest, ) *SearchQueryRequest`

NewSearchQueryRequest instantiates a new SearchQueryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchQueryRequestWithDefaults

`func NewSearchQueryRequestWithDefaults() *SearchQueryRequest`

NewSearchQueryRequestWithDefaults instantiates a new SearchQueryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequests

`func (o *SearchQueryRequest) GetRequests() []SearchRequest`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *SearchQueryRequest) GetRequestsOk() (*[]SearchRequest, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *SearchQueryRequest) SetRequests(v []SearchRequest)`

SetRequests sets Requests field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


