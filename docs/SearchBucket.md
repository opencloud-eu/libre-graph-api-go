# SearchBucket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** | The discrete value of the field that was used to compute the aggregation. For terms aggregations this is the field value. For range aggregations this is a string representation of the range. For geohash aggregations this is a geohash cell.  | [optional] 
**Count** | Pointer to **int64** | The approximate number of search matches that share the same value specified in the &#x60;key&#x60; property.  | [optional] 
**AggregationFilterToken** | Pointer to **string** | A token containing the encoded filter that narrows search matches to this bucket. To use it, pass it as part of the &#x60;aggregationFilters&#x60; property of a subsequent &#x60;searchRequest&#x60; in the format &#x60;{field}:{aggregationFilterToken}&#x60;. The filter matches the bucket &#x60;key&#x60; exactly and case-sensitively, so the narrowed result set is the set of matches counted in this bucket.  For terms buckets the token is the key encoded as lowercase hex of its UTF-8 bytes, prefixed with &#x60;ǂǂ&#x60; (U+01C2 twice) and wrapped in double quotes, e.g. &#x60;\&quot;ǂǂ5361786f6e\&quot;&#x60; for the key &#x60;Saxon&#x60;. For range buckets the token is &#x60;range({from}, {to})&#x60; with the bounds of the matching &#x60;bucketAggregationRange&#x60;; an open lower bound is written as &#x60;min&#x60;, an open upper bound as &#x60;max&#x60; followed by &#x60;to&#x3D;\&quot;le\&quot;&#x60;, e.g. &#x60;range(min, 1980)&#x60;, &#x60;range(1980, 1990)&#x60; and &#x60;range(2010, max, to&#x3D;\&quot;le\&quot;)&#x60;. This is the same encoding MS Graph uses. Geohash buckets carry no token; narrow by location through the search query instead.  | [optional] [readonly] 
**LibreGraphSubAggregations** | Pointer to [**[]SearchAggregation**](SearchAggregation.md) | Nested aggregation results, one per sub-aggregation requested on the parent &#x60;aggregationOption&#x60;. Libregraph extension not present in MS Graph.  | [optional] 

## Methods

### NewSearchBucket

`func NewSearchBucket() *SearchBucket`

NewSearchBucket instantiates a new SearchBucket object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchBucketWithDefaults

`func NewSearchBucketWithDefaults() *SearchBucket`

NewSearchBucketWithDefaults instantiates a new SearchBucket object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *SearchBucket) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *SearchBucket) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *SearchBucket) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *SearchBucket) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetCount

`func (o *SearchBucket) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *SearchBucket) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *SearchBucket) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *SearchBucket) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetAggregationFilterToken

`func (o *SearchBucket) GetAggregationFilterToken() string`

GetAggregationFilterToken returns the AggregationFilterToken field if non-nil, zero value otherwise.

### GetAggregationFilterTokenOk

`func (o *SearchBucket) GetAggregationFilterTokenOk() (*string, bool)`

GetAggregationFilterTokenOk returns a tuple with the AggregationFilterToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationFilterToken

`func (o *SearchBucket) SetAggregationFilterToken(v string)`

SetAggregationFilterToken sets AggregationFilterToken field to given value.

### HasAggregationFilterToken

`func (o *SearchBucket) HasAggregationFilterToken() bool`

HasAggregationFilterToken returns a boolean if a field has been set.

### GetLibreGraphSubAggregations

`func (o *SearchBucket) GetLibreGraphSubAggregations() []SearchAggregation`

GetLibreGraphSubAggregations returns the LibreGraphSubAggregations field if non-nil, zero value otherwise.

### GetLibreGraphSubAggregationsOk

`func (o *SearchBucket) GetLibreGraphSubAggregationsOk() (*[]SearchAggregation, bool)`

GetLibreGraphSubAggregationsOk returns a tuple with the LibreGraphSubAggregations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibreGraphSubAggregations

`func (o *SearchBucket) SetLibreGraphSubAggregations(v []SearchAggregation)`

SetLibreGraphSubAggregations sets LibreGraphSubAggregations field to given value.

### HasLibreGraphSubAggregations

`func (o *SearchBucket) HasLibreGraphSubAggregations() bool`

HasLibreGraphSubAggregations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


