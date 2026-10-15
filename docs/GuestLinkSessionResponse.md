# GuestLinkSessionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PermissionId** | **string** | Identifier of the share (permission) the guest was invited to. | 
**ExpirationDateTime** | **time.Time** | The date and time when the guest session expires, in UTC. After that the guest has to renew the guest link. | 

## Methods

### NewGuestLinkSessionResponse

`func NewGuestLinkSessionResponse(permissionId string, expirationDateTime time.Time, ) *GuestLinkSessionResponse`

NewGuestLinkSessionResponse instantiates a new GuestLinkSessionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGuestLinkSessionResponseWithDefaults

`func NewGuestLinkSessionResponseWithDefaults() *GuestLinkSessionResponse`

NewGuestLinkSessionResponseWithDefaults instantiates a new GuestLinkSessionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermissionId

`func (o *GuestLinkSessionResponse) GetPermissionId() string`

GetPermissionId returns the PermissionId field if non-nil, zero value otherwise.

### GetPermissionIdOk

`func (o *GuestLinkSessionResponse) GetPermissionIdOk() (*string, bool)`

GetPermissionIdOk returns a tuple with the PermissionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionId

`func (o *GuestLinkSessionResponse) SetPermissionId(v string)`

SetPermissionId sets PermissionId field to given value.


### GetExpirationDateTime

`func (o *GuestLinkSessionResponse) GetExpirationDateTime() time.Time`

GetExpirationDateTime returns the ExpirationDateTime field if non-nil, zero value otherwise.

### GetExpirationDateTimeOk

`func (o *GuestLinkSessionResponse) GetExpirationDateTimeOk() (*time.Time, bool)`

GetExpirationDateTimeOk returns a tuple with the ExpirationDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDateTime

`func (o *GuestLinkSessionResponse) SetExpirationDateTime(v time.Time)`

SetExpirationDateTime sets ExpirationDateTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


