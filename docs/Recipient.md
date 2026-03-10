# Recipient

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailAddress** | Pointer to [**EmailAddress**](EmailAddress.md) |  | [optional] 

## Methods

### NewRecipient

`func NewRecipient() *Recipient`

NewRecipient instantiates a new Recipient object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecipientWithDefaults

`func NewRecipientWithDefaults() *Recipient`

NewRecipientWithDefaults instantiates a new Recipient object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmailAddress

`func (o *Recipient) GetEmailAddress() EmailAddress`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *Recipient) GetEmailAddressOk() (*EmailAddress, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *Recipient) SetEmailAddress(v EmailAddress)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *Recipient) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


