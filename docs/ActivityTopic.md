# ActivityTopic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | **string** | How to read the value. Only &#x60;text&#x60; is supported at the time of writing. built from it.  | [default to "text"]
**Value** | **string** | With the source &#x60;text&#x60;, the id of the resource the activity happened on. | 

## Methods

### NewActivityTopic

`func NewActivityTopic(source string, value string, ) *ActivityTopic`

NewActivityTopic instantiates a new ActivityTopic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivityTopicWithDefaults

`func NewActivityTopicWithDefaults() *ActivityTopic`

NewActivityTopicWithDefaults instantiates a new ActivityTopic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *ActivityTopic) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ActivityTopic) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ActivityTopic) SetSource(v string)`

SetSource sets Source field to given value.


### GetValue

`func (o *ActivityTopic) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ActivityTopic) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ActivityTopic) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


