# GuestLinkError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorType** | **string** | Machine-readable error identifier. | 
**Message** | **string** | Human-readable error message. | 
**PermissionId** | **string** | Permission (share) identifier related to the error, when known. | 

## Methods

### NewGuestLinkError

`func NewGuestLinkError(errorType string, message string, permissionId string, ) *GuestLinkError`

NewGuestLinkError instantiates a new GuestLinkError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGuestLinkErrorWithDefaults

`func NewGuestLinkErrorWithDefaults() *GuestLinkError`

NewGuestLinkErrorWithDefaults instantiates a new GuestLinkError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorType

`func (o *GuestLinkError) GetErrorType() string`

GetErrorType returns the ErrorType field if non-nil, zero value otherwise.

### GetErrorTypeOk

`func (o *GuestLinkError) GetErrorTypeOk() (*string, bool)`

GetErrorTypeOk returns a tuple with the ErrorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorType

`func (o *GuestLinkError) SetErrorType(v string)`

SetErrorType sets ErrorType field to given value.


### GetMessage

`func (o *GuestLinkError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GuestLinkError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GuestLinkError) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetPermissionId

`func (o *GuestLinkError) GetPermissionId() string`

GetPermissionId returns the PermissionId field if non-nil, zero value otherwise.

### GetPermissionIdOk

`func (o *GuestLinkError) GetPermissionIdOk() (*string, bool)`

GetPermissionIdOk returns a tuple with the PermissionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionId

`func (o *GuestLinkError) SetPermissionId(v string)`

SetPermissionId sets PermissionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


