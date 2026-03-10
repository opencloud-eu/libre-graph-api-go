# Invitation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InvitedUserDisplayName** | Pointer to **string** | The display name of the user being invited. | [optional] 
**InvitedUserEmailAddress** | Pointer to **string** | The email address of the user being invited. Required. | [optional] 
**InvitedUserMessageInfo** | Pointer to [**InvitedUserMessageInfo**](InvitedUserMessageInfo.md) |  | [optional] 
**SendInvitationMessage** | Pointer to **bool** | Indicates whether an invitation message should be sent to the user. | [optional] 
**InviteRedirectUrl** | Pointer to **string** | The URL to which the user is redirected after accepting the invitation. Required. | [optional] 
**InviteRedeemUrl** | Pointer to **string** | The URL that the user can use to redeem the invitation. Read-only. | [optional] [readonly] 
**Status** | Pointer to **string** | The status of the invitation. Read-only. | [optional] [readonly] 
**InvitedUser** | Pointer to [**User**](User.md) |  | [optional] 
**InvitedUserType** | Pointer to **string** | The type of user being invited. | [optional] 

## Methods

### NewInvitation

`func NewInvitation() *Invitation`

NewInvitation instantiates a new Invitation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvitationWithDefaults

`func NewInvitationWithDefaults() *Invitation`

NewInvitationWithDefaults instantiates a new Invitation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvitedUserDisplayName

`func (o *Invitation) GetInvitedUserDisplayName() string`

GetInvitedUserDisplayName returns the InvitedUserDisplayName field if non-nil, zero value otherwise.

### GetInvitedUserDisplayNameOk

`func (o *Invitation) GetInvitedUserDisplayNameOk() (*string, bool)`

GetInvitedUserDisplayNameOk returns a tuple with the InvitedUserDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitedUserDisplayName

`func (o *Invitation) SetInvitedUserDisplayName(v string)`

SetInvitedUserDisplayName sets InvitedUserDisplayName field to given value.

### HasInvitedUserDisplayName

`func (o *Invitation) HasInvitedUserDisplayName() bool`

HasInvitedUserDisplayName returns a boolean if a field has been set.

### GetInvitedUserEmailAddress

`func (o *Invitation) GetInvitedUserEmailAddress() string`

GetInvitedUserEmailAddress returns the InvitedUserEmailAddress field if non-nil, zero value otherwise.

### GetInvitedUserEmailAddressOk

`func (o *Invitation) GetInvitedUserEmailAddressOk() (*string, bool)`

GetInvitedUserEmailAddressOk returns a tuple with the InvitedUserEmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitedUserEmailAddress

`func (o *Invitation) SetInvitedUserEmailAddress(v string)`

SetInvitedUserEmailAddress sets InvitedUserEmailAddress field to given value.

### HasInvitedUserEmailAddress

`func (o *Invitation) HasInvitedUserEmailAddress() bool`

HasInvitedUserEmailAddress returns a boolean if a field has been set.

### GetInvitedUserMessageInfo

`func (o *Invitation) GetInvitedUserMessageInfo() InvitedUserMessageInfo`

GetInvitedUserMessageInfo returns the InvitedUserMessageInfo field if non-nil, zero value otherwise.

### GetInvitedUserMessageInfoOk

`func (o *Invitation) GetInvitedUserMessageInfoOk() (*InvitedUserMessageInfo, bool)`

GetInvitedUserMessageInfoOk returns a tuple with the InvitedUserMessageInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitedUserMessageInfo

`func (o *Invitation) SetInvitedUserMessageInfo(v InvitedUserMessageInfo)`

SetInvitedUserMessageInfo sets InvitedUserMessageInfo field to given value.

### HasInvitedUserMessageInfo

`func (o *Invitation) HasInvitedUserMessageInfo() bool`

HasInvitedUserMessageInfo returns a boolean if a field has been set.

### GetSendInvitationMessage

`func (o *Invitation) GetSendInvitationMessage() bool`

GetSendInvitationMessage returns the SendInvitationMessage field if non-nil, zero value otherwise.

### GetSendInvitationMessageOk

`func (o *Invitation) GetSendInvitationMessageOk() (*bool, bool)`

GetSendInvitationMessageOk returns a tuple with the SendInvitationMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendInvitationMessage

`func (o *Invitation) SetSendInvitationMessage(v bool)`

SetSendInvitationMessage sets SendInvitationMessage field to given value.

### HasSendInvitationMessage

`func (o *Invitation) HasSendInvitationMessage() bool`

HasSendInvitationMessage returns a boolean if a field has been set.

### GetInviteRedirectUrl

`func (o *Invitation) GetInviteRedirectUrl() string`

GetInviteRedirectUrl returns the InviteRedirectUrl field if non-nil, zero value otherwise.

### GetInviteRedirectUrlOk

`func (o *Invitation) GetInviteRedirectUrlOk() (*string, bool)`

GetInviteRedirectUrlOk returns a tuple with the InviteRedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviteRedirectUrl

`func (o *Invitation) SetInviteRedirectUrl(v string)`

SetInviteRedirectUrl sets InviteRedirectUrl field to given value.

### HasInviteRedirectUrl

`func (o *Invitation) HasInviteRedirectUrl() bool`

HasInviteRedirectUrl returns a boolean if a field has been set.

### GetInviteRedeemUrl

`func (o *Invitation) GetInviteRedeemUrl() string`

GetInviteRedeemUrl returns the InviteRedeemUrl field if non-nil, zero value otherwise.

### GetInviteRedeemUrlOk

`func (o *Invitation) GetInviteRedeemUrlOk() (*string, bool)`

GetInviteRedeemUrlOk returns a tuple with the InviteRedeemUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviteRedeemUrl

`func (o *Invitation) SetInviteRedeemUrl(v string)`

SetInviteRedeemUrl sets InviteRedeemUrl field to given value.

### HasInviteRedeemUrl

`func (o *Invitation) HasInviteRedeemUrl() bool`

HasInviteRedeemUrl returns a boolean if a field has been set.

### GetStatus

`func (o *Invitation) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Invitation) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Invitation) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Invitation) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetInvitedUser

`func (o *Invitation) GetInvitedUser() User`

GetInvitedUser returns the InvitedUser field if non-nil, zero value otherwise.

### GetInvitedUserOk

`func (o *Invitation) GetInvitedUserOk() (*User, bool)`

GetInvitedUserOk returns a tuple with the InvitedUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitedUser

`func (o *Invitation) SetInvitedUser(v User)`

SetInvitedUser sets InvitedUser field to given value.

### HasInvitedUser

`func (o *Invitation) HasInvitedUser() bool`

HasInvitedUser returns a boolean if a field has been set.

### GetInvitedUserType

`func (o *Invitation) GetInvitedUserType() string`

GetInvitedUserType returns the InvitedUserType field if non-nil, zero value otherwise.

### GetInvitedUserTypeOk

`func (o *Invitation) GetInvitedUserTypeOk() (*string, bool)`

GetInvitedUserTypeOk returns a tuple with the InvitedUserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitedUserType

`func (o *Invitation) SetInvitedUserType(v string)`

SetInvitedUserType sets InvitedUserType field to given value.

### HasInvitedUserType

`func (o *Invitation) HasInvitedUserType() bool`

HasInvitedUserType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


