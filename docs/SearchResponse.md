# SearchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SearchTerms** | Pointer to **[]string** | Contains the search terms sent in the initial search query. | [optional] 
**HitsContainers** | Pointer to [**[]SearchHitsContainer**](SearchHitsContainer.md) | A collection of search result sets. One for each entity type that was queried.  | [optional] 

## Methods

### NewSearchResponse

`func NewSearchResponse() *SearchResponse`

NewSearchResponse instantiates a new SearchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchResponseWithDefaults

`func NewSearchResponseWithDefaults() *SearchResponse`

NewSearchResponseWithDefaults instantiates a new SearchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSearchTerms

`func (o *SearchResponse) GetSearchTerms() []string`

GetSearchTerms returns the SearchTerms field if non-nil, zero value otherwise.

### GetSearchTermsOk

`func (o *SearchResponse) GetSearchTermsOk() (*[]string, bool)`

GetSearchTermsOk returns a tuple with the SearchTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchTerms

`func (o *SearchResponse) SetSearchTerms(v []string)`

SetSearchTerms sets SearchTerms field to given value.

### HasSearchTerms

`func (o *SearchResponse) HasSearchTerms() bool`

HasSearchTerms returns a boolean if a field has been set.

### GetHitsContainers

`func (o *SearchResponse) GetHitsContainers() []SearchHitsContainer`

GetHitsContainers returns the HitsContainers field if non-nil, zero value otherwise.

### GetHitsContainersOk

`func (o *SearchResponse) GetHitsContainersOk() (*[]SearchHitsContainer, bool)`

GetHitsContainersOk returns a tuple with the HitsContainers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHitsContainers

`func (o *SearchResponse) SetHitsContainers(v []SearchHitsContainer)`

SetHitsContainers sets HitsContainers field to given value.

### HasHitsContainers

`func (o *SearchResponse) HasHitsContainers() bool`

HasHitsContainers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


