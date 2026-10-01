# SearchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityTypes** | **[]string** | One or more types of resources expected in the response. Currently only &#x60;driveItem&#x60; is supported.  | 
**Query** | [**SearchQuery**](SearchQuery.md) |  | 
**From** | Pointer to **int32** | Specifies the offset for the search results. Offset 0 returns the very first result. Used together with the &#x60;size&#x60; property for pagination.  | [optional] [default to 0]
**Size** | Pointer to **int32** | The size of the page to be retrieved. The maximum value is 500. Set to 0 to return only aggregations without any hits.  | [optional] [default to 25]
**Aggregations** | Pointer to [**[]AggregationOption**](AggregationOption.md) | Specifies aggregations (also known as refiners or facets) to be returned alongside the search results. Optional.  | [optional] 
**AggregationFilters** | Pointer to **[]string** | Contains one or more filters to narrow search results to specific buckets of a prior aggregation. Build each filter from the response of a prior search that aggregated on the same field: take the &#x60;aggregationFilterToken&#x60; of the wanted &#x60;searchBucket&#x60; and combine it with the field as &#x60;{field}:{aggregationFilterToken}&#x60;, e.g. &#x60;audio.artist:\&quot;ǂǂ5361786f6e\&quot;&#x60; for a terms bucket or &#x60;audio.year:range(1980, 1990)&#x60; for a range bucket. Several buckets of the same field are combined with &#x60;{field}:or({aggregationFilterToken},{aggregationFilterToken})&#x60;. Whitespace after the commas of &#x60;range(...)&#x60; and &#x60;or(...)&#x60; is optional.  Multiple filters can be provided as separate array items. This results in a logical AND between the filters. Filters that are not built from server-issued tokens are rejected with &#x60;invalidRequest&#x60;.  | [optional] 
**SortProperties** | Pointer to [**[]SortProperty**](SortProperty.md) | Contains the ordered collection of fields to sort the results on, primary sort key first. At most 5 sort properties. If absent, the results are sorted by relevance. See &#x60;sortProperty.name&#x60; for the set of sortable fields. Ties are broken by relevance, and results missing the sort property are placed last. Optional.  | [optional] 

## Methods

### NewSearchRequest

`func NewSearchRequest(entityTypes []string, query SearchQuery, ) *SearchRequest`

NewSearchRequest instantiates a new SearchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchRequestWithDefaults

`func NewSearchRequestWithDefaults() *SearchRequest`

NewSearchRequestWithDefaults instantiates a new SearchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityTypes

`func (o *SearchRequest) GetEntityTypes() []string`

GetEntityTypes returns the EntityTypes field if non-nil, zero value otherwise.

### GetEntityTypesOk

`func (o *SearchRequest) GetEntityTypesOk() (*[]string, bool)`

GetEntityTypesOk returns a tuple with the EntityTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityTypes

`func (o *SearchRequest) SetEntityTypes(v []string)`

SetEntityTypes sets EntityTypes field to given value.


### GetQuery

`func (o *SearchRequest) GetQuery() SearchQuery`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *SearchRequest) GetQueryOk() (*SearchQuery, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *SearchRequest) SetQuery(v SearchQuery)`

SetQuery sets Query field to given value.


### GetFrom

`func (o *SearchRequest) GetFrom() int32`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *SearchRequest) GetFromOk() (*int32, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *SearchRequest) SetFrom(v int32)`

SetFrom sets From field to given value.

### HasFrom

`func (o *SearchRequest) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetSize

`func (o *SearchRequest) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *SearchRequest) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *SearchRequest) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *SearchRequest) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetAggregations

`func (o *SearchRequest) GetAggregations() []AggregationOption`

GetAggregations returns the Aggregations field if non-nil, zero value otherwise.

### GetAggregationsOk

`func (o *SearchRequest) GetAggregationsOk() (*[]AggregationOption, bool)`

GetAggregationsOk returns a tuple with the Aggregations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregations

`func (o *SearchRequest) SetAggregations(v []AggregationOption)`

SetAggregations sets Aggregations field to given value.

### HasAggregations

`func (o *SearchRequest) HasAggregations() bool`

HasAggregations returns a boolean if a field has been set.

### GetAggregationFilters

`func (o *SearchRequest) GetAggregationFilters() []string`

GetAggregationFilters returns the AggregationFilters field if non-nil, zero value otherwise.

### GetAggregationFiltersOk

`func (o *SearchRequest) GetAggregationFiltersOk() (*[]string, bool)`

GetAggregationFiltersOk returns a tuple with the AggregationFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationFilters

`func (o *SearchRequest) SetAggregationFilters(v []string)`

SetAggregationFilters sets AggregationFilters field to given value.

### HasAggregationFilters

`func (o *SearchRequest) HasAggregationFilters() bool`

HasAggregationFilters returns a boolean if a field has been set.

### GetSortProperties

`func (o *SearchRequest) GetSortProperties() []SortProperty`

GetSortProperties returns the SortProperties field if non-nil, zero value otherwise.

### GetSortPropertiesOk

`func (o *SearchRequest) GetSortPropertiesOk() (*[]SortProperty, bool)`

GetSortPropertiesOk returns a tuple with the SortProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortProperties

`func (o *SearchRequest) SetSortProperties(v []SortProperty)`

SetSortProperties sets SortProperties field to given value.

### HasSortProperties

`func (o *SearchRequest) HasSortProperties() bool`

HasSortProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


