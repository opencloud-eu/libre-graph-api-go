# SearchAggregation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | Pointer to **string** | Defines the field in the request on which the aggregation was computed.  | [optional] 
**Buckets** | Pointer to [**[]SearchBucket**](SearchBucket.md) | Defines the computed buckets for this aggregation. For bucket aggregations they are sorted according to the &#x60;sortBy&#x60; and &#x60;isDescending&#x60; specified in the &#x60;bucketDefinition&#x60; of the corresponding &#x60;aggregationOption&#x60;; for geohash aggregations they are ordered by &#x60;count&#x60;, descending.  | [optional] 
**LibreGraphMetric** | Pointer to [**SearchMetric**](SearchMetric.md) |  | [optional] 

## Methods

### NewSearchAggregation

`func NewSearchAggregation() *SearchAggregation`

NewSearchAggregation instantiates a new SearchAggregation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchAggregationWithDefaults

`func NewSearchAggregationWithDefaults() *SearchAggregation`

NewSearchAggregationWithDefaults instantiates a new SearchAggregation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *SearchAggregation) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *SearchAggregation) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *SearchAggregation) SetField(v string)`

SetField sets Field field to given value.

### HasField

`func (o *SearchAggregation) HasField() bool`

HasField returns a boolean if a field has been set.

### GetBuckets

`func (o *SearchAggregation) GetBuckets() []SearchBucket`

GetBuckets returns the Buckets field if non-nil, zero value otherwise.

### GetBucketsOk

`func (o *SearchAggregation) GetBucketsOk() (*[]SearchBucket, bool)`

GetBucketsOk returns a tuple with the Buckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuckets

`func (o *SearchAggregation) SetBuckets(v []SearchBucket)`

SetBuckets sets Buckets field to given value.

### HasBuckets

`func (o *SearchAggregation) HasBuckets() bool`

HasBuckets returns a boolean if a field has been set.

### GetLibreGraphMetric

`func (o *SearchAggregation) GetLibreGraphMetric() SearchMetric`

GetLibreGraphMetric returns the LibreGraphMetric field if non-nil, zero value otherwise.

### GetLibreGraphMetricOk

`func (o *SearchAggregation) GetLibreGraphMetricOk() (*SearchMetric, bool)`

GetLibreGraphMetricOk returns a tuple with the LibreGraphMetric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibreGraphMetric

`func (o *SearchAggregation) SetLibreGraphMetric(v SearchMetric)`

SetLibreGraphMetric sets LibreGraphMetric field to given value.

### HasLibreGraphMetric

`func (o *SearchAggregation) HasLibreGraphMetric() bool`

HasLibreGraphMetric returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


