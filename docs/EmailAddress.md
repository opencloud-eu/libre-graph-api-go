# EmailAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | The email address. | [optional] 
**Name** | Pointer to **string** | The name associated with the email address. | [optional] 

## Methods

### NewEmailAddress

`func NewEmailAddress() *EmailAddress`

NewEmailAddress instantiates a new EmailAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailAddressWithDefaults

`func NewEmailAddressWithDefaults() *EmailAddress`

NewEmailAddressWithDefaults instantiates a new EmailAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *EmailAddress) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *EmailAddress) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *EmailAddress) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *EmailAddress) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetName

`func (o *EmailAddress) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EmailAddress) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EmailAddress) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EmailAddress) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


