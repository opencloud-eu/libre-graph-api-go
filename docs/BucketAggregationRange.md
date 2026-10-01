# BucketAggregationRange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | Pointer to **string** | Defines the lower bound from which to compute the aggregation. The value is always a string. Numeric bounds must be provided as their string representation (e.g. &#x60;\&quot;1980\&quot;&#x60;). Date bounds must use the &#x60;YYYY-MM-DDTHH:mm:ssZ&#x60; format. Optional if &#x60;to&#x60; is provided.  | [optional] 
**To** | Pointer to **string** | Defines the upper bound up to which to compute the aggregation. The value is always a string. Numeric bounds must be provided as their string representation (e.g. &#x60;\&quot;2000\&quot;&#x60;). Date bounds must use the &#x60;YYYY-MM-DDTHH:mm:ssZ&#x60; format. Optional if &#x60;from&#x60; is provided.  | [optional] 

## Methods

### NewBucketAggregationRange

`func NewBucketAggregationRange() *BucketAggregationRange`

NewBucketAggregationRange instantiates a new BucketAggregationRange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBucketAggregationRangeWithDefaults

`func NewBucketAggregationRangeWithDefaults() *BucketAggregationRange`

NewBucketAggregationRangeWithDefaults instantiates a new BucketAggregationRange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *BucketAggregationRange) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *BucketAggregationRange) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *BucketAggregationRange) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *BucketAggregationRange) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *BucketAggregationRange) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *BucketAggregationRange) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *BucketAggregationRange) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *BucketAggregationRange) HasTo() bool`

HasTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


