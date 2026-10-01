# SearchQuery200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to [**[]SearchResponse**](SearchResponse.md) | A collection of search response objects, one per request. | [optional] 

## Methods

### NewSearchQuery200Response

`func NewSearchQuery200Response() *SearchQuery200Response`

NewSearchQuery200Response instantiates a new SearchQuery200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchQuery200ResponseWithDefaults

`func NewSearchQuery200ResponseWithDefaults() *SearchQuery200Response`

NewSearchQuery200ResponseWithDefaults instantiates a new SearchQuery200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *SearchQuery200Response) GetValue() []SearchResponse`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SearchQuery200Response) GetValueOk() (*[]SearchResponse, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SearchQuery200Response) SetValue(v []SearchResponse)`

SetValue sets Value field to given value.

### HasValue

`func (o *SearchQuery200Response) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


