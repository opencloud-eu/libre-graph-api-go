# BucketDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SortBy** | **string** | The possible values are &#x60;count&#x60; to sort by the number of matches in the aggregation, &#x60;keyAsString&#x60; to sort alphabetically based on the key in the aggregation, and &#x60;keyAsNumber&#x60; to sort numerically based on the key in the aggregation. Required.  | 
**IsDescending** | Pointer to **bool** | Set to &#x60;true&#x60; to specify the sort order as descending. Optional, defaults to &#x60;false&#x60; (ascending).  | [optional] [default to false]
**MinimumCount** | Pointer to **int32** | The minimum number of items that should be present in the aggregation for the bucket to be returned in the response. Optional, default is 0.  | [optional] [default to 0]
**Ranges** | Pointer to [**[]BucketAggregationRange**](BucketAggregationRange.md) | Specifies the manual ranges to compute the aggregation buckets. This is only valid for non-string facets of date or numeric type. Optional. Follows the [MS Graph bucketAggregationRange](https://learn.microsoft.com/en-us/graph/api/resources/bucketaggregationrange) resource type.  | [optional] 

## Methods

### NewBucketDefinition

`func NewBucketDefinition(sortBy string, ) *BucketDefinition`

NewBucketDefinition instantiates a new BucketDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBucketDefinitionWithDefaults

`func NewBucketDefinitionWithDefaults() *BucketDefinition`

NewBucketDefinitionWithDefaults instantiates a new BucketDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSortBy

`func (o *BucketDefinition) GetSortBy() string`

GetSortBy returns the SortBy field if non-nil, zero value otherwise.

### GetSortByOk

`func (o *BucketDefinition) GetSortByOk() (*string, bool)`

GetSortByOk returns a tuple with the SortBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortBy

`func (o *BucketDefinition) SetSortBy(v string)`

SetSortBy sets SortBy field to given value.


### GetIsDescending

`func (o *BucketDefinition) GetIsDescending() bool`

GetIsDescending returns the IsDescending field if non-nil, zero value otherwise.

### GetIsDescendingOk

`func (o *BucketDefinition) GetIsDescendingOk() (*bool, bool)`

GetIsDescendingOk returns a tuple with the IsDescending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDescending

`func (o *BucketDefinition) SetIsDescending(v bool)`

SetIsDescending sets IsDescending field to given value.

### HasIsDescending

`func (o *BucketDefinition) HasIsDescending() bool`

HasIsDescending returns a boolean if a field has been set.

### GetMinimumCount

`func (o *BucketDefinition) GetMinimumCount() int32`

GetMinimumCount returns the MinimumCount field if non-nil, zero value otherwise.

### GetMinimumCountOk

`func (o *BucketDefinition) GetMinimumCountOk() (*int32, bool)`

GetMinimumCountOk returns a tuple with the MinimumCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumCount

`func (o *BucketDefinition) SetMinimumCount(v int32)`

SetMinimumCount sets MinimumCount field to given value.

### HasMinimumCount

`func (o *BucketDefinition) HasMinimumCount() bool`

HasMinimumCount returns a boolean if a field has been set.

### GetRanges

`func (o *BucketDefinition) GetRanges() []BucketAggregationRange`

GetRanges returns the Ranges field if non-nil, zero value otherwise.

### GetRangesOk

`func (o *BucketDefinition) GetRangesOk() (*[]BucketAggregationRange, bool)`

GetRangesOk returns a tuple with the Ranges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRanges

`func (o *BucketDefinition) SetRanges(v []BucketAggregationRange)`

SetRanges sets Ranges field to given value.

### HasRanges

`func (o *BucketDefinition) HasRanges() bool`

HasRanges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


