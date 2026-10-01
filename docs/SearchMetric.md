# SearchMetric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** | Echoes the &#x60;kind&#x60; of the corresponding &#x60;metricDefinition&#x60;, allowing consumers (and the search service&#39;s cross-space merge layer) to pick the right reducer when combining results.  | [optional] 
**Value** | Pointer to **float64** | The scalar result of the metric. | [optional] 

## Methods

### NewSearchMetric

`func NewSearchMetric() *SearchMetric`

NewSearchMetric instantiates a new SearchMetric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchMetricWithDefaults

`func NewSearchMetricWithDefaults() *SearchMetric`

NewSearchMetricWithDefaults instantiates a new SearchMetric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *SearchMetric) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *SearchMetric) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *SearchMetric) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *SearchMetric) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetValue

`func (o *SearchMetric) GetValue() float64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SearchMetric) GetValueOk() (*float64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SearchMetric) SetValue(v float64)`

SetValue sets Value field to given value.

### HasValue

`func (o *SearchMetric) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


