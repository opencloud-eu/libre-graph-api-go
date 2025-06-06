# UnifiedRolePermission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedResourceActions** | Pointer to **[]string** | Set of tasks that can be performed on a resource. Required.  The following is the schema for resource actions:  &#x60;&#x60;&#x60;    {Namespace}/{Entity}/{PropertySet}/{Action} &#x60;&#x60;&#x60;   For example: &#x60;libre.graph/applications/credentials/update&#x60;   * *{Namespace}* - The services that exposes the task. For example, all tasks in libre graph use the namespace &#x60;libre.graph&#x60;.  * *{Entity}* - The logical features or components exposed by the service in libre graph. For example, &#x60;applications&#x60;, &#x60;servicePrincipals&#x60;, or &#x60;groups&#x60;.  * *{PropertySet}* - Optional. The specific properties or aspects of the entity for which access is being granted.    For example, &#x60;libre.graph/applications/authentication/read&#x60; grants the ability to read the reply URL, logout URL,    and implicit flow property on the **application** object in libre graph. The following are reserved names for common property sets:    * &#x60;allProperties&#x60; - Designates all properties of the entity, including privileged properties.      Examples include &#x60;libre.graph/applications/allProperties/read&#x60; and &#x60;libre.graph/applications/allProperties/update&#x60;.    * &#x60;basic&#x60; - Designates common read properties but excludes privileged ones.      For example, &#x60;libre.graph/applications/basic/update&#x60; includes the ability to update standard properties like display name.    * &#x60;standard&#x60; - Designates common update properties but excludes privileged ones.      For example, &#x60;libre.graph/applications/standard/read&#x60;.  * *{Actions}* - The operations being granted. In most circumstances, permissions should be expressed in terms of CRUD operations or allTasks. Actions include:    * &#x60;create&#x60; - The ability to create a new instance of the entity.    * &#x60;read&#x60; - The ability to read a given property set (including allProperties).    * &#x60;update&#x60; - The ability to update a given property set (including allProperties).    * &#x60;delete&#x60; - The ability to delete a given entity.    * &#x60;allTasks&#x60; - Represents all CRUD operations (create, read, update, and delete).   Following the CS3 API we can represent the CS3 permissions by mapping them to driveItem properties or relations like this:  | [CS3 ResourcePermission](https://cs3org.github.io/cs3apis/#cs3.storage.provider.v1beta1.ResourcePermissions) | action | comment |  | ------------------------------------------------------------------------------------------------------------ | ------ | ------- |  | &#x60;stat&#x60; | &#x60;libre.graph/driveItem/basic/read&#x60; | &#x60;basic&#x60; because it does not include versions or trashed items |  | &#x60;get_quota&#x60; | &#x60;libre.graph/driveItem/quota/read&#x60; | read only the &#x60;quota&#x60; property |  | &#x60;get_path&#x60; | &#x60;libre.graph/driveItem/path/read&#x60; | read only the &#x60;path&#x60; property |  | &#x60;move&#x60; | &#x60;libre.graph/driveItem/path/update&#x60; | allows updating the &#x60;path&#x60; property of a CS3 resource |  | &#x60;delete&#x60; | &#x60;libre.graph/driveItem/standard/delete&#x60; | &#x60;standard&#x60; because deleting is a common update operation |  | &#x60;list_container&#x60; | &#x60;libre.graph/driveItem/children/read&#x60; | |  | &#x60;create_container&#x60; | &#x60;libre.graph/driveItem/children/create&#x60; | |  | &#x60;initiate_file_download&#x60; | &#x60;libre.graph/driveItem/content/read&#x60; | &#x60;content&#x60; is the property read when initiating a download |  | &#x60;initiate_file_upload&#x60; | &#x60;libre.graph/driveItem/upload/create&#x60; | &#x60;uploads&#x60; are a separate property. postprocessing creates the &#x60;content&#x60; |  | &#x60;add_grant&#x60; | &#x60;libre.graph/driveItem/permissions/create&#x60; | |  | &#x60;list_grant&#x60; | &#x60;libre.graph/driveItem/permissions/read&#x60; | |  | &#x60;update_grant&#x60; | &#x60;libre.graph/driveItem/permissions/update&#x60; | |  | &#x60;remove_grant&#x60; | &#x60;libre.graph/driveItem/permissions/delete&#x60; | |  | &#x60;deny_grant&#x60; | &#x60;libre.graph/driveItem/permissions/deny&#x60; | uses a non CRUD action &#x60;deny&#x60; |  | &#x60;list_file_versions&#x60; | &#x60;libre.graph/driveItem/versions/read&#x60; | &#x60;versions&#x60; is a &#x60;driveItemVersion&#x60; collection |  | &#x60;restore_file_version&#x60; | &#x60;libre.graph/driveItem/versions/update&#x60; | the only &#x60;update&#x60; action is restore |  | &#x60;list_recycle&#x60; | &#x60;libre.graph/driveItem/deleted/read&#x60; | reading a driveItem &#x60;deleted&#x60; property implies listing |  | &#x60;restore_recycle_item&#x60; | &#x60;libre.graph/driveItem/deleted/update&#x60; | the only &#x60;update&#x60; action is restore |  | &#x60;purge_recycle&#x60; | &#x60;libre.graph/driveItem/deleted/delete&#x60; | allows purging deleted &#x60;driveItems&#x60; |   Managing drives would be a different entity. A space manager role could be written as &#x60;libre.graph/drive/permission/allTasks&#x60;.  | [optional] 
**Condition** | Pointer to **string** | Optional constraints that must be met for the permission to be effective. Not supported for custom roles.  Conditions define constraints that must be met. For example, a requirement that target resource must have a certain property. The following are the supported conditions:  * Drive: &#x60;exists @Resource.Drive&#x60; - The target resource must be a drive/space * Folder: &#x60;exists @Resource.Folder&#x60; - The target resource must be a folder * File: &#x60;exists @Resource.File&#x60; - The target resource must be a file  The following is an example of a role permission with a condition that the target resource is a folder: &#x60;&#x60;&#x60;json   \&quot;rolePermissions\&quot;: [       {           \&quot;allowedResourceActions\&quot;: [               \&quot;libre.graph/applications/basic/update\&quot;,               \&quot;libre.graph/applications/credentials/update\&quot;           ],           \&quot;condition\&quot;:  \&quot;exists @Resource.File\&quot;       }   ] &#x60;&#x60;&#x60; Conditions aren&#39;t supported for custom roles.  | [optional] 

## Methods

### NewUnifiedRolePermission

`func NewUnifiedRolePermission() *UnifiedRolePermission`

NewUnifiedRolePermission instantiates a new UnifiedRolePermission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnifiedRolePermissionWithDefaults

`func NewUnifiedRolePermissionWithDefaults() *UnifiedRolePermission`

NewUnifiedRolePermissionWithDefaults instantiates a new UnifiedRolePermission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedResourceActions

`func (o *UnifiedRolePermission) GetAllowedResourceActions() []string`

GetAllowedResourceActions returns the AllowedResourceActions field if non-nil, zero value otherwise.

### GetAllowedResourceActionsOk

`func (o *UnifiedRolePermission) GetAllowedResourceActionsOk() (*[]string, bool)`

GetAllowedResourceActionsOk returns a tuple with the AllowedResourceActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedResourceActions

`func (o *UnifiedRolePermission) SetAllowedResourceActions(v []string)`

SetAllowedResourceActions sets AllowedResourceActions field to given value.

### HasAllowedResourceActions

`func (o *UnifiedRolePermission) HasAllowedResourceActions() bool`

HasAllowedResourceActions returns a boolean if a field has been set.

### GetCondition

`func (o *UnifiedRolePermission) GetCondition() string`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *UnifiedRolePermission) GetConditionOk() (*string, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *UnifiedRolePermission) SetCondition(v string)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *UnifiedRolePermission) HasCondition() bool`

HasCondition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


