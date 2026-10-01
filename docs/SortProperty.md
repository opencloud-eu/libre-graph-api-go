# SortProperty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the property to sort the search results by. Required.  Sortable are the scalar search fields of the search hit&#39;s resource: &#x60;name&#x60;, &#x60;size&#x60;, &#x60;lastModifiedDateTime&#x60;, &#x60;mimeType&#x60; and the scalar facet properties such as &#x60;photo.takenDateTime&#x60;, &#x60;photo.iso&#x60;, &#x60;audio.artist&#x60;, &#x60;audio.year&#x60; or &#x60;image.width&#x60;. Strings sort lexicographically, numbers and dates by value. Multivalued properties (e.g. &#x60;@libre.graph.tags&#x60;) and unknown properties are rejected with &#x60;invalidRequest&#x60;.  | 
**IsDescending** | Pointer to **bool** | Set to &#x60;true&#x60; to specify the sort order as descending. Optional, defaults to &#x60;false&#x60; (ascending).  | [optional] [default to false]

## Methods

### NewSortProperty

`func NewSortProperty(name string, ) *SortProperty`

NewSortProperty instantiates a new SortProperty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSortPropertyWithDefaults

`func NewSortPropertyWithDefaults() *SortProperty`

NewSortPropertyWithDefaults instantiates a new SortProperty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SortProperty) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SortProperty) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SortProperty) SetName(v string)`

SetName sets Name field to given value.


### GetIsDescending

`func (o *SortProperty) GetIsDescending() bool`

GetIsDescending returns the IsDescending field if non-nil, zero value otherwise.

### GetIsDescendingOk

`func (o *SortProperty) GetIsDescendingOk() (*bool, bool)`

GetIsDescendingOk returns a tuple with the IsDescending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDescending

`func (o *SortProperty) SetIsDescending(v bool)`

SetIsDescending sets IsDescending field to given value.

### HasIsDescending

`func (o *SortProperty) HasIsDescending() bool`

HasIsDescending returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


