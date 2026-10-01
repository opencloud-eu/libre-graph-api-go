# SearchQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QueryString** | **string** | The search query string in KQL (Keyword Query Language) format. The query string can contain free-text keywords and property filters.  Examples: - &#x60;budget report&#x60;: free text search - &#x60;mediatype:audio&#x60;: filter by media type - &#x60;audio.artist:\&quot;Saxon\&quot;&#x60;: filter by audio metadata - &#x60;audio.genre:Rock AND audio.year:1979&#x60;: combined filters  | 

## Methods

### NewSearchQuery

`func NewSearchQuery(queryString string, ) *SearchQuery`

NewSearchQuery instantiates a new SearchQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchQueryWithDefaults

`func NewSearchQueryWithDefaults() *SearchQuery`

NewSearchQueryWithDefaults instantiates a new SearchQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueryString

`func (o *SearchQuery) GetQueryString() string`

GetQueryString returns the QueryString field if non-nil, zero value otherwise.

### GetQueryStringOk

`func (o *SearchQuery) GetQueryStringOk() (*string, bool)`

GetQueryStringOk returns a tuple with the QueryString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryString

`func (o *SearchQuery) SetQueryString(v string)`

SetQueryString sets QueryString field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


