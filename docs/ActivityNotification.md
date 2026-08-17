# ActivityNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Topic** | [**ActivityTopic**](ActivityTopic.md) |  | 
**ActivityType** | **string** | What happened. Only &#x60;mentioned&#x60; is supported at the time of writing.  | 
**TeamsAppId** | **string** | The app the activity happened in, as an id the server knows.  | 

## Methods

### NewActivityNotification

`func NewActivityNotification(topic ActivityTopic, activityType string, teamsAppId string, ) *ActivityNotification`

NewActivityNotification instantiates a new ActivityNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivityNotificationWithDefaults

`func NewActivityNotificationWithDefaults() *ActivityNotification`

NewActivityNotificationWithDefaults instantiates a new ActivityNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTopic

`func (o *ActivityNotification) GetTopic() ActivityTopic`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *ActivityNotification) GetTopicOk() (*ActivityTopic, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *ActivityNotification) SetTopic(v ActivityTopic)`

SetTopic sets Topic field to given value.


### GetActivityType

`func (o *ActivityNotification) GetActivityType() string`

GetActivityType returns the ActivityType field if non-nil, zero value otherwise.

### GetActivityTypeOk

`func (o *ActivityNotification) GetActivityTypeOk() (*string, bool)`

GetActivityTypeOk returns a tuple with the ActivityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityType

`func (o *ActivityNotification) SetActivityType(v string)`

SetActivityType sets ActivityType field to given value.


### GetTeamsAppId

`func (o *ActivityNotification) GetTeamsAppId() string`

GetTeamsAppId returns the TeamsAppId field if non-nil, zero value otherwise.

### GetTeamsAppIdOk

`func (o *ActivityNotification) GetTeamsAppIdOk() (*string, bool)`

GetTeamsAppIdOk returns a tuple with the TeamsAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamsAppId

`func (o *ActivityNotification) SetTeamsAppId(v string)`

SetTeamsAppId sets TeamsAppId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


