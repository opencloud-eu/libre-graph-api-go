# MetricDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | The reducer applied to the field values of all matches. Required.  &#x60;avg&#x60; is not a simple reducer (averages of averages are not averages): the backend carries &#x60;(sum, count)&#x60; internally and emits only the final value on the outermost merge.  | 

## Methods

### NewMetricDefinition

`func NewMetricDefinition(kind string, ) *MetricDefinition`

NewMetricDefinition instantiates a new MetricDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricDefinitionWithDefaults

`func NewMetricDefinitionWithDefaults() *MetricDefinition`

NewMetricDefinitionWithDefaults instantiates a new MetricDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *MetricDefinition) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *MetricDefinition) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *MetricDefinition) SetKind(v string)`

SetKind sets Kind field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


