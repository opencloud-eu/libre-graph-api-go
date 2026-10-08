# GuestLinkRenewRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PermissionId** | **string** | Identifier of the share (permission) the guest was invited to. | 
**Token** | Pointer to **string** | Previous guest link token. Optional when the guest session cookie is sent instead. | [optional] 

## Methods

### NewGuestLinkRenewRequest

`func NewGuestLinkRenewRequest(permissionId string, ) *GuestLinkRenewRequest`

NewGuestLinkRenewRequest instantiates a new GuestLinkRenewRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGuestLinkRenewRequestWithDefaults

`func NewGuestLinkRenewRequestWithDefaults() *GuestLinkRenewRequest`

NewGuestLinkRenewRequestWithDefaults instantiates a new GuestLinkRenewRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermissionId

`func (o *GuestLinkRenewRequest) GetPermissionId() string`

GetPermissionId returns the PermissionId field if non-nil, zero value otherwise.

### GetPermissionIdOk

`func (o *GuestLinkRenewRequest) GetPermissionIdOk() (*string, bool)`

GetPermissionIdOk returns a tuple with the PermissionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionId

`func (o *GuestLinkRenewRequest) SetPermissionId(v string)`

SetPermissionId sets PermissionId field to given value.


### GetToken

`func (o *GuestLinkRenewRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GuestLinkRenewRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GuestLinkRenewRequest) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GuestLinkRenewRequest) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


