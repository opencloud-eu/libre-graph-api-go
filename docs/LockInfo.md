# LockInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LockType** | Pointer to **string** | The type of lock currently held on the file. OpenCloud currently only issues exclusive locks, same as MS Graph, even if it defines more. Read-only. | [optional] [readonly] 
**CreatedDateTime** | Pointer to **time.Time** | The date and time when the lock was created, in UTC. Read-only. | [optional] [readonly] 
**ExpirationDateTime** | Pointer to **time.Time** | The date and time when the lock expires, in UTC. Read-only. | [optional] [readonly] 
**Owners** | Pointer to [**[]Identity**](Identity.md) | The collection of users that currently hold the lock on the file. Read-only. | [optional] [readonly] 
**LibreGraphAppName** | Pointer to **string** | Name of the application holding the lock, for example an office application. Not part of MS Graph. Read-only. | [optional] [readonly] 

## Methods

### NewLockInfo

`func NewLockInfo() *LockInfo`

NewLockInfo instantiates a new LockInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLockInfoWithDefaults

`func NewLockInfoWithDefaults() *LockInfo`

NewLockInfoWithDefaults instantiates a new LockInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLockType

`func (o *LockInfo) GetLockType() string`

GetLockType returns the LockType field if non-nil, zero value otherwise.

### GetLockTypeOk

`func (o *LockInfo) GetLockTypeOk() (*string, bool)`

GetLockTypeOk returns a tuple with the LockType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockType

`func (o *LockInfo) SetLockType(v string)`

SetLockType sets LockType field to given value.

### HasLockType

`func (o *LockInfo) HasLockType() bool`

HasLockType returns a boolean if a field has been set.

### GetCreatedDateTime

`func (o *LockInfo) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *LockInfo) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *LockInfo) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *LockInfo) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### GetExpirationDateTime

`func (o *LockInfo) GetExpirationDateTime() time.Time`

GetExpirationDateTime returns the ExpirationDateTime field if non-nil, zero value otherwise.

### GetExpirationDateTimeOk

`func (o *LockInfo) GetExpirationDateTimeOk() (*time.Time, bool)`

GetExpirationDateTimeOk returns a tuple with the ExpirationDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDateTime

`func (o *LockInfo) SetExpirationDateTime(v time.Time)`

SetExpirationDateTime sets ExpirationDateTime field to given value.

### HasExpirationDateTime

`func (o *LockInfo) HasExpirationDateTime() bool`

HasExpirationDateTime returns a boolean if a field has been set.

### GetOwners

`func (o *LockInfo) GetOwners() []Identity`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *LockInfo) GetOwnersOk() (*[]Identity, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *LockInfo) SetOwners(v []Identity)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *LockInfo) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetLibreGraphAppName

`func (o *LockInfo) GetLibreGraphAppName() string`

GetLibreGraphAppName returns the LibreGraphAppName field if non-nil, zero value otherwise.

### GetLibreGraphAppNameOk

`func (o *LockInfo) GetLibreGraphAppNameOk() (*string, bool)`

GetLibreGraphAppNameOk returns a tuple with the LibreGraphAppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibreGraphAppName

`func (o *LockInfo) SetLibreGraphAppName(v string)`

SetLibreGraphAppName sets LibreGraphAppName field to given value.

### HasLibreGraphAppName

`func (o *LockInfo) HasLibreGraphAppName() bool`

HasLibreGraphAppName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


