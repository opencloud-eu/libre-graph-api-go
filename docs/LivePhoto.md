# LivePhoto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentId** | **string** | Identifier (UUID) shared by the still image and the paired video of one Live Photo. Read-only.  | [readonly] 
**StillImageTimeUs** | Pointer to **int64** | Time in microseconds within the paired video at which the still image was captured. The value comes from the video file only: it is the presentation time of the com.apple.quicktime.still-image-time timed metadata sample in the QuickTime movie, so it is only present on the video half. The still image carries no reliable equivalent (the Apple maker note tag 0x0017, historically documented as a derivable video index, does not encode a usable time on current iOS versions). If absent, readers should use a timestamp near the middle of the video track. Read-only.  | [optional] [readonly] 
**Auto** | Pointer to **bool** | True if the device decided automatically whether to capture the Live Photo video (com.apple.quicktime.live-photo.auto). Only present on the video half. Read-only.  | [optional] [readonly] 
**VitalityScore** | Pointer to **float64** | Score between 0 and 1 rating how interesting the motion clip is, used by readers to decide whether to autoplay it (com.apple.quicktime.live-photo.vitality-score). Only present on the video half. Read-only.  | [optional] [readonly] 
**VitalityScoringVersion** | Pointer to **int64** | Version of the algorithm that produced vitalityScore (com.apple.quicktime.live-photo.vitality-scoring-version). Only present on the video half. Read-only.  | [optional] [readonly] 

## Methods

### NewLivePhoto

`func NewLivePhoto(contentId string, ) *LivePhoto`

NewLivePhoto instantiates a new LivePhoto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLivePhotoWithDefaults

`func NewLivePhotoWithDefaults() *LivePhoto`

NewLivePhotoWithDefaults instantiates a new LivePhoto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentId

`func (o *LivePhoto) GetContentId() string`

GetContentId returns the ContentId field if non-nil, zero value otherwise.

### GetContentIdOk

`func (o *LivePhoto) GetContentIdOk() (*string, bool)`

GetContentIdOk returns a tuple with the ContentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentId

`func (o *LivePhoto) SetContentId(v string)`

SetContentId sets ContentId field to given value.


### GetStillImageTimeUs

`func (o *LivePhoto) GetStillImageTimeUs() int64`

GetStillImageTimeUs returns the StillImageTimeUs field if non-nil, zero value otherwise.

### GetStillImageTimeUsOk

`func (o *LivePhoto) GetStillImageTimeUsOk() (*int64, bool)`

GetStillImageTimeUsOk returns a tuple with the StillImageTimeUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStillImageTimeUs

`func (o *LivePhoto) SetStillImageTimeUs(v int64)`

SetStillImageTimeUs sets StillImageTimeUs field to given value.

### HasStillImageTimeUs

`func (o *LivePhoto) HasStillImageTimeUs() bool`

HasStillImageTimeUs returns a boolean if a field has been set.

### GetAuto

`func (o *LivePhoto) GetAuto() bool`

GetAuto returns the Auto field if non-nil, zero value otherwise.

### GetAutoOk

`func (o *LivePhoto) GetAutoOk() (*bool, bool)`

GetAutoOk returns a tuple with the Auto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuto

`func (o *LivePhoto) SetAuto(v bool)`

SetAuto sets Auto field to given value.

### HasAuto

`func (o *LivePhoto) HasAuto() bool`

HasAuto returns a boolean if a field has been set.

### GetVitalityScore

`func (o *LivePhoto) GetVitalityScore() float64`

GetVitalityScore returns the VitalityScore field if non-nil, zero value otherwise.

### GetVitalityScoreOk

`func (o *LivePhoto) GetVitalityScoreOk() (*float64, bool)`

GetVitalityScoreOk returns a tuple with the VitalityScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVitalityScore

`func (o *LivePhoto) SetVitalityScore(v float64)`

SetVitalityScore sets VitalityScore field to given value.

### HasVitalityScore

`func (o *LivePhoto) HasVitalityScore() bool`

HasVitalityScore returns a boolean if a field has been set.

### GetVitalityScoringVersion

`func (o *LivePhoto) GetVitalityScoringVersion() int64`

GetVitalityScoringVersion returns the VitalityScoringVersion field if non-nil, zero value otherwise.

### GetVitalityScoringVersionOk

`func (o *LivePhoto) GetVitalityScoringVersionOk() (*int64, bool)`

GetVitalityScoringVersionOk returns a tuple with the VitalityScoringVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVitalityScoringVersion

`func (o *LivePhoto) SetVitalityScoringVersion(v int64)`

SetVitalityScoringVersion sets VitalityScoringVersion field to given value.

### HasVitalityScoringVersion

`func (o *LivePhoto) HasVitalityScoringVersion() bool`

HasVitalityScoringVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


