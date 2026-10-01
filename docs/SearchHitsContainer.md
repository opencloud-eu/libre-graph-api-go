# SearchHitsContainer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hits** | Pointer to [**[]SearchHit**](SearchHit.md) | A collection of the search results, ordered by relevance or, when the request specifies &#x60;sortProperties&#x60;, by those properties.  | [optional] 
**Total** | Pointer to **int64** | The total number of results. Note this is not the number of results on the page, but the total number of results satisfying the query.  | [optional] [readonly] 
**MoreResultsAvailable** | Pointer to **bool** | Provides information if more results are available. Based on this information, you can adjust the &#x60;from&#x60; and &#x60;size&#x60; properties of the &#x60;searchRequest&#x60; accordingly.  | [optional] [readonly] 
**Aggregations** | Pointer to [**[]SearchAggregation**](SearchAggregation.md) | Contains the collection of aggregations computed based on the provided &#x60;aggregationOption&#x60; definitions in the request.  | [optional] 

## Methods

### NewSearchHitsContainer

`func NewSearchHitsContainer() *SearchHitsContainer`

NewSearchHitsContainer instantiates a new SearchHitsContainer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchHitsContainerWithDefaults

`func NewSearchHitsContainerWithDefaults() *SearchHitsContainer`

NewSearchHitsContainerWithDefaults instantiates a new SearchHitsContainer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHits

`func (o *SearchHitsContainer) GetHits() []SearchHit`

GetHits returns the Hits field if non-nil, zero value otherwise.

### GetHitsOk

`func (o *SearchHitsContainer) GetHitsOk() (*[]SearchHit, bool)`

GetHitsOk returns a tuple with the Hits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHits

`func (o *SearchHitsContainer) SetHits(v []SearchHit)`

SetHits sets Hits field to given value.

### HasHits

`func (o *SearchHitsContainer) HasHits() bool`

HasHits returns a boolean if a field has been set.

### GetTotal

`func (o *SearchHitsContainer) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *SearchHitsContainer) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *SearchHitsContainer) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *SearchHitsContainer) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetMoreResultsAvailable

`func (o *SearchHitsContainer) GetMoreResultsAvailable() bool`

GetMoreResultsAvailable returns the MoreResultsAvailable field if non-nil, zero value otherwise.

### GetMoreResultsAvailableOk

`func (o *SearchHitsContainer) GetMoreResultsAvailableOk() (*bool, bool)`

GetMoreResultsAvailableOk returns a tuple with the MoreResultsAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoreResultsAvailable

`func (o *SearchHitsContainer) SetMoreResultsAvailable(v bool)`

SetMoreResultsAvailable sets MoreResultsAvailable field to given value.

### HasMoreResultsAvailable

`func (o *SearchHitsContainer) HasMoreResultsAvailable() bool`

HasMoreResultsAvailable returns a boolean if a field has been set.

### GetAggregations

`func (o *SearchHitsContainer) GetAggregations() []SearchAggregation`

GetAggregations returns the Aggregations field if non-nil, zero value otherwise.

### GetAggregationsOk

`func (o *SearchHitsContainer) GetAggregationsOk() (*[]SearchAggregation, bool)`

GetAggregationsOk returns a tuple with the Aggregations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregations

`func (o *SearchHitsContainer) SetAggregations(v []SearchAggregation)`

SetAggregations sets Aggregations field to given value.

### HasAggregations

`func (o *SearchHitsContainer) HasAggregations() bool`

HasAggregations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


