# SearchHit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HitId** | Pointer to **string** | The internal identifier for the item. | [optional] [readonly] 
**Rank** | Pointer to **int32** | The rank or the order of the result. | [optional] [readonly] 
**Summary** | Pointer to **string** | A summary of the result, if a summary is available.  | [optional] [readonly] 
**Resource** | Pointer to [**DriveItem**](DriveItem.md) |  | [optional] 

## Methods

### NewSearchHit

`func NewSearchHit() *SearchHit`

NewSearchHit instantiates a new SearchHit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchHitWithDefaults

`func NewSearchHitWithDefaults() *SearchHit`

NewSearchHitWithDefaults instantiates a new SearchHit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHitId

`func (o *SearchHit) GetHitId() string`

GetHitId returns the HitId field if non-nil, zero value otherwise.

### GetHitIdOk

`func (o *SearchHit) GetHitIdOk() (*string, bool)`

GetHitIdOk returns a tuple with the HitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHitId

`func (o *SearchHit) SetHitId(v string)`

SetHitId sets HitId field to given value.

### HasHitId

`func (o *SearchHit) HasHitId() bool`

HasHitId returns a boolean if a field has been set.

### GetRank

`func (o *SearchHit) GetRank() int32`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *SearchHit) GetRankOk() (*int32, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRank

`func (o *SearchHit) SetRank(v int32)`

SetRank sets Rank field to given value.

### HasRank

`func (o *SearchHit) HasRank() bool`

HasRank returns a boolean if a field has been set.

### GetSummary

`func (o *SearchHit) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *SearchHit) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *SearchHit) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *SearchHit) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetResource

`func (o *SearchHit) GetResource() DriveItem`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *SearchHit) GetResourceOk() (*DriveItem, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *SearchHit) SetResource(v DriveItem)`

SetResource sets Resource field to given value.

### HasResource

`func (o *SearchHit) HasResource() bool`

HasResource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


