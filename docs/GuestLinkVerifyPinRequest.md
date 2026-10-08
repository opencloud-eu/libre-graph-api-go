# GuestLinkVerifyPinRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pin** | **string** | One-time PIN received from the renewed guest link. | 
**PermissionId** | **string** | Identifier of the share (permission) the guest was invited to. | 

## Methods

### NewGuestLinkVerifyPinRequest

`func NewGuestLinkVerifyPinRequest(pin string, permissionId string, ) *GuestLinkVerifyPinRequest`

NewGuestLinkVerifyPinRequest instantiates a new GuestLinkVerifyPinRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGuestLinkVerifyPinRequestWithDefaults

`func NewGuestLinkVerifyPinRequestWithDefaults() *GuestLinkVerifyPinRequest`

NewGuestLinkVerifyPinRequestWithDefaults instantiates a new GuestLinkVerifyPinRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPin

`func (o *GuestLinkVerifyPinRequest) GetPin() string`

GetPin returns the Pin field if non-nil, zero value otherwise.

### GetPinOk

`func (o *GuestLinkVerifyPinRequest) GetPinOk() (*string, bool)`

GetPinOk returns a tuple with the Pin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPin

`func (o *GuestLinkVerifyPinRequest) SetPin(v string)`

SetPin sets Pin field to given value.


### GetPermissionId

`func (o *GuestLinkVerifyPinRequest) GetPermissionId() string`

GetPermissionId returns the PermissionId field if non-nil, zero value otherwise.

### GetPermissionIdOk

`func (o *GuestLinkVerifyPinRequest) GetPermissionIdOk() (*string, bool)`

GetPermissionIdOk returns a tuple with the PermissionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionId

`func (o *GuestLinkVerifyPinRequest) SetPermissionId(v string)`

SetPermissionId sets PermissionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


