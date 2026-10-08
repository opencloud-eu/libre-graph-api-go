# GuestLinkVerifyTokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | **string** | One-time guest link token received from the guest link. | 

## Methods

### NewGuestLinkVerifyTokenRequest

`func NewGuestLinkVerifyTokenRequest(token string, ) *GuestLinkVerifyTokenRequest`

NewGuestLinkVerifyTokenRequest instantiates a new GuestLinkVerifyTokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGuestLinkVerifyTokenRequestWithDefaults

`func NewGuestLinkVerifyTokenRequestWithDefaults() *GuestLinkVerifyTokenRequest`

NewGuestLinkVerifyTokenRequestWithDefaults instantiates a new GuestLinkVerifyTokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *GuestLinkVerifyTokenRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GuestLinkVerifyTokenRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GuestLinkVerifyTokenRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


