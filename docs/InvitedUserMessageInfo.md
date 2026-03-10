# InvitedUserMessageInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CcRecipients** | Pointer to [**[]Recipient**](Recipient.md) | Additional recipients who will receive a copy of the invitation message. | [optional] 
**CustomizedMessageBody** | Pointer to **string** | The customized message body that will be included in the invitation message. | [optional] 
**MessageLanguage** | Pointer to **string** | The language of the invitation message. | [optional] 

## Methods

### NewInvitedUserMessageInfo

`func NewInvitedUserMessageInfo() *InvitedUserMessageInfo`

NewInvitedUserMessageInfo instantiates a new InvitedUserMessageInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvitedUserMessageInfoWithDefaults

`func NewInvitedUserMessageInfoWithDefaults() *InvitedUserMessageInfo`

NewInvitedUserMessageInfoWithDefaults instantiates a new InvitedUserMessageInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCcRecipients

`func (o *InvitedUserMessageInfo) GetCcRecipients() []Recipient`

GetCcRecipients returns the CcRecipients field if non-nil, zero value otherwise.

### GetCcRecipientsOk

`func (o *InvitedUserMessageInfo) GetCcRecipientsOk() (*[]Recipient, bool)`

GetCcRecipientsOk returns a tuple with the CcRecipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcRecipients

`func (o *InvitedUserMessageInfo) SetCcRecipients(v []Recipient)`

SetCcRecipients sets CcRecipients field to given value.

### HasCcRecipients

`func (o *InvitedUserMessageInfo) HasCcRecipients() bool`

HasCcRecipients returns a boolean if a field has been set.

### GetCustomizedMessageBody

`func (o *InvitedUserMessageInfo) GetCustomizedMessageBody() string`

GetCustomizedMessageBody returns the CustomizedMessageBody field if non-nil, zero value otherwise.

### GetCustomizedMessageBodyOk

`func (o *InvitedUserMessageInfo) GetCustomizedMessageBodyOk() (*string, bool)`

GetCustomizedMessageBodyOk returns a tuple with the CustomizedMessageBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomizedMessageBody

`func (o *InvitedUserMessageInfo) SetCustomizedMessageBody(v string)`

SetCustomizedMessageBody sets CustomizedMessageBody field to given value.

### HasCustomizedMessageBody

`func (o *InvitedUserMessageInfo) HasCustomizedMessageBody() bool`

HasCustomizedMessageBody returns a boolean if a field has been set.

### GetMessageLanguage

`func (o *InvitedUserMessageInfo) GetMessageLanguage() string`

GetMessageLanguage returns the MessageLanguage field if non-nil, zero value otherwise.

### GetMessageLanguageOk

`func (o *InvitedUserMessageInfo) GetMessageLanguageOk() (*string, bool)`

GetMessageLanguageOk returns a tuple with the MessageLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageLanguage

`func (o *InvitedUserMessageInfo) SetMessageLanguage(v string)`

SetMessageLanguage sets MessageLanguage field to given value.

### HasMessageLanguage

`func (o *InvitedUserMessageInfo) HasMessageLanguage() bool`

HasMessageLanguage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


