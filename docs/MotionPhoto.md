# MotionPhoto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **int32** | The file format version of the Motion Photo. Currently always 1. Read-only. | [optional] [readonly] 
**PresentationTimestampUs** | Pointer to **int64** | Presentation timestamp in microseconds of the video frame that corresponds to the still image. A value of -1 indicates unspecified. If absent, readers should use a timestamp near the middle of the video track. Read-only.  | [optional] [readonly] 
**VideoSize** | Pointer to **int64** | Size in bytes of the embedded video portion of the file. The video is appended at the end of the file, so clients can fetch it with a Range request: &#x60;Range: bytes&#x3D;&lt;fileSize - videoSize&gt;-&#x60;. Read-only.  | [optional] [readonly] 

## Methods

### NewMotionPhoto

`func NewMotionPhoto() *MotionPhoto`

NewMotionPhoto instantiates a new MotionPhoto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMotionPhotoWithDefaults

`func NewMotionPhotoWithDefaults() *MotionPhoto`

NewMotionPhotoWithDefaults instantiates a new MotionPhoto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *MotionPhoto) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MotionPhoto) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *MotionPhoto) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *MotionPhoto) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetPresentationTimestampUs

`func (o *MotionPhoto) GetPresentationTimestampUs() int64`

GetPresentationTimestampUs returns the PresentationTimestampUs field if non-nil, zero value otherwise.

### GetPresentationTimestampUsOk

`func (o *MotionPhoto) GetPresentationTimestampUsOk() (*int64, bool)`

GetPresentationTimestampUsOk returns a tuple with the PresentationTimestampUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresentationTimestampUs

`func (o *MotionPhoto) SetPresentationTimestampUs(v int64)`

SetPresentationTimestampUs sets PresentationTimestampUs field to given value.

### HasPresentationTimestampUs

`func (o *MotionPhoto) HasPresentationTimestampUs() bool`

HasPresentationTimestampUs returns a boolean if a field has been set.

### GetVideoSize

`func (o *MotionPhoto) GetVideoSize() int64`

GetVideoSize returns the VideoSize field if non-nil, zero value otherwise.

### GetVideoSizeOk

`func (o *MotionPhoto) GetVideoSizeOk() (*int64, bool)`

GetVideoSizeOk returns a tuple with the VideoSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoSize

`func (o *MotionPhoto) SetVideoSize(v int64)`

SetVideoSize sets VideoSize field to given value.

### HasVideoSize

`func (o *MotionPhoto) HasVideoSize() bool`

HasVideoSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


