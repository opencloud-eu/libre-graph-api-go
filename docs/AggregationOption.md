# AggregationOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | **string** | Specifies the field in the schema of the specified entity type that the aggregation should be computed on. Required.  Examples: &#x60;audio.artist&#x60;, &#x60;audio.genre&#x60;, &#x60;audio.year&#x60;, &#x60;mimeType&#x60;.  | 
**Size** | Pointer to **int32** | The number of &#x60;searchBucket&#x60; resources to be returned. This is optional and only applies to terms aggregations. Combined with &#x60;bucketDefinition.sortBy&#x60; and &#x60;bucketDefinition.isDescending&#x60; to produce the top N results by count or key. When not specified, all buckets are returned.  | [optional] 
**BucketDefinition** | Pointer to [**BucketDefinition**](BucketDefinition.md) |  | [optional] 
**LibreGraphSubAggregations** | Pointer to [**[]AggregationOption**](AggregationOption.md) | Nested aggregations computed within each bucket of this aggregation. Libregraph extension not present in MS Graph.  Backends that don&#39;t support native composite aggregations (e.g. bleve) emulate them by walking the matched result set; OpenSearch translates them to native composite aggregations.  | [optional] 
**LibreGraphMetricDefinition** | Pointer to [**MetricDefinition**](MetricDefinition.md) |  | [optional] 

## Methods

### NewAggregationOption

`func NewAggregationOption(field string, ) *AggregationOption`

NewAggregationOption instantiates a new AggregationOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAggregationOptionWithDefaults

`func NewAggregationOptionWithDefaults() *AggregationOption`

NewAggregationOptionWithDefaults instantiates a new AggregationOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *AggregationOption) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *AggregationOption) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *AggregationOption) SetField(v string)`

SetField sets Field field to given value.


### GetSize

`func (o *AggregationOption) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *AggregationOption) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *AggregationOption) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *AggregationOption) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetBucketDefinition

`func (o *AggregationOption) GetBucketDefinition() BucketDefinition`

GetBucketDefinition returns the BucketDefinition field if non-nil, zero value otherwise.

### GetBucketDefinitionOk

`func (o *AggregationOption) GetBucketDefinitionOk() (*BucketDefinition, bool)`

GetBucketDefinitionOk returns a tuple with the BucketDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketDefinition

`func (o *AggregationOption) SetBucketDefinition(v BucketDefinition)`

SetBucketDefinition sets BucketDefinition field to given value.

### HasBucketDefinition

`func (o *AggregationOption) HasBucketDefinition() bool`

HasBucketDefinition returns a boolean if a field has been set.

### GetLibreGraphSubAggregations

`func (o *AggregationOption) GetLibreGraphSubAggregations() []AggregationOption`

GetLibreGraphSubAggregations returns the LibreGraphSubAggregations field if non-nil, zero value otherwise.

### GetLibreGraphSubAggregationsOk

`func (o *AggregationOption) GetLibreGraphSubAggregationsOk() (*[]AggregationOption, bool)`

GetLibreGraphSubAggregationsOk returns a tuple with the LibreGraphSubAggregations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibreGraphSubAggregations

`func (o *AggregationOption) SetLibreGraphSubAggregations(v []AggregationOption)`

SetLibreGraphSubAggregations sets LibreGraphSubAggregations field to given value.

### HasLibreGraphSubAggregations

`func (o *AggregationOption) HasLibreGraphSubAggregations() bool`

HasLibreGraphSubAggregations returns a boolean if a field has been set.

### GetLibreGraphMetricDefinition

`func (o *AggregationOption) GetLibreGraphMetricDefinition() MetricDefinition`

GetLibreGraphMetricDefinition returns the LibreGraphMetricDefinition field if non-nil, zero value otherwise.

### GetLibreGraphMetricDefinitionOk

`func (o *AggregationOption) GetLibreGraphMetricDefinitionOk() (*MetricDefinition, bool)`

GetLibreGraphMetricDefinitionOk returns a tuple with the LibreGraphMetricDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibreGraphMetricDefinition

`func (o *AggregationOption) SetLibreGraphMetricDefinition(v MetricDefinition)`

SetLibreGraphMetricDefinition sets LibreGraphMetricDefinition field to given value.

### HasLibreGraphMetricDefinition

`func (o *AggregationOption) HasLibreGraphMetricDefinition() bool`

HasLibreGraphMetricDefinition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


