# PendingOperationsPendingContentUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QueuedDateTime** | Pointer to **time.Time** | Time the operation was queued. May be absent. Read-only. | [optional] [readonly] 

## Methods

### NewPendingOperationsPendingContentUpdate

`func NewPendingOperationsPendingContentUpdate() *PendingOperationsPendingContentUpdate`

NewPendingOperationsPendingContentUpdate instantiates a new PendingOperationsPendingContentUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPendingOperationsPendingContentUpdateWithDefaults

`func NewPendingOperationsPendingContentUpdateWithDefaults() *PendingOperationsPendingContentUpdate`

NewPendingOperationsPendingContentUpdateWithDefaults instantiates a new PendingOperationsPendingContentUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueuedDateTime

`func (o *PendingOperationsPendingContentUpdate) GetQueuedDateTime() time.Time`

GetQueuedDateTime returns the QueuedDateTime field if non-nil, zero value otherwise.

### GetQueuedDateTimeOk

`func (o *PendingOperationsPendingContentUpdate) GetQueuedDateTimeOk() (*time.Time, bool)`

GetQueuedDateTimeOk returns a tuple with the QueuedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueuedDateTime

`func (o *PendingOperationsPendingContentUpdate) SetQueuedDateTime(v time.Time)`

SetQueuedDateTime sets QueuedDateTime field to given value.

### HasQueuedDateTime

`func (o *PendingOperationsPendingContentUpdate) HasQueuedDateTime() bool`

HasQueuedDateTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


