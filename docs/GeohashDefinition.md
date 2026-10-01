# GeohashDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Precision** | **int32** | The geohash length of the returned cells (1-12); higher means finer cells. Required.  | 

## Methods

### NewGeohashDefinition

`func NewGeohashDefinition(precision int32, ) *GeohashDefinition`

NewGeohashDefinition instantiates a new GeohashDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeohashDefinitionWithDefaults

`func NewGeohashDefinitionWithDefaults() *GeohashDefinition`

NewGeohashDefinitionWithDefaults instantiates a new GeohashDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrecision

`func (o *GeohashDefinition) GetPrecision() int32`

GetPrecision returns the Precision field if non-nil, zero value otherwise.

### GetPrecisionOk

`func (o *GeohashDefinition) GetPrecisionOk() (*int32, bool)`

GetPrecisionOk returns a tuple with the Precision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecision

`func (o *GeohashDefinition) SetPrecision(v int32)`

SetPrecision sets Precision field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


