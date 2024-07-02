# ApiV1AuditMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreateTimestamp** | Pointer to **time.Time** | The creation time. | [optional] 
**CreateUserId** | Pointer to **string** | Id of the user that created the entity. | [optional] 
**ModifyTimestamp** | Pointer to **time.Time** | The last modification time. | [optional] 
**ModifyUserId** | Pointer to **string** | Id of the user that last modified the entity. | [optional] 

## Methods

### NewApiV1AuditMetadata

`func NewApiV1AuditMetadata() *ApiV1AuditMetadata`

NewApiV1AuditMetadata instantiates a new ApiV1AuditMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV1AuditMetadataWithDefaults

`func NewApiV1AuditMetadataWithDefaults() *ApiV1AuditMetadata`

NewApiV1AuditMetadataWithDefaults instantiates a new ApiV1AuditMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreateTimestamp

`func (o *ApiV1AuditMetadata) GetCreateTimestamp() time.Time`

GetCreateTimestamp returns the CreateTimestamp field if non-nil, zero value otherwise.

### GetCreateTimestampOk

`func (o *ApiV1AuditMetadata) GetCreateTimestampOk() (*time.Time, bool)`

GetCreateTimestampOk returns a tuple with the CreateTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTimestamp

`func (o *ApiV1AuditMetadata) SetCreateTimestamp(v time.Time)`

SetCreateTimestamp sets CreateTimestamp field to given value.

### HasCreateTimestamp

`func (o *ApiV1AuditMetadata) HasCreateTimestamp() bool`

HasCreateTimestamp returns a boolean if a field has been set.

### GetCreateUserId

`func (o *ApiV1AuditMetadata) GetCreateUserId() string`

GetCreateUserId returns the CreateUserId field if non-nil, zero value otherwise.

### GetCreateUserIdOk

`func (o *ApiV1AuditMetadata) GetCreateUserIdOk() (*string, bool)`

GetCreateUserIdOk returns a tuple with the CreateUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateUserId

`func (o *ApiV1AuditMetadata) SetCreateUserId(v string)`

SetCreateUserId sets CreateUserId field to given value.

### HasCreateUserId

`func (o *ApiV1AuditMetadata) HasCreateUserId() bool`

HasCreateUserId returns a boolean if a field has been set.

### GetModifyTimestamp

`func (o *ApiV1AuditMetadata) GetModifyTimestamp() time.Time`

GetModifyTimestamp returns the ModifyTimestamp field if non-nil, zero value otherwise.

### GetModifyTimestampOk

`func (o *ApiV1AuditMetadata) GetModifyTimestampOk() (*time.Time, bool)`

GetModifyTimestampOk returns a tuple with the ModifyTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifyTimestamp

`func (o *ApiV1AuditMetadata) SetModifyTimestamp(v time.Time)`

SetModifyTimestamp sets ModifyTimestamp field to given value.

### HasModifyTimestamp

`func (o *ApiV1AuditMetadata) HasModifyTimestamp() bool`

HasModifyTimestamp returns a boolean if a field has been set.

### GetModifyUserId

`func (o *ApiV1AuditMetadata) GetModifyUserId() string`

GetModifyUserId returns the ModifyUserId field if non-nil, zero value otherwise.

### GetModifyUserIdOk

`func (o *ApiV1AuditMetadata) GetModifyUserIdOk() (*string, bool)`

GetModifyUserIdOk returns a tuple with the ModifyUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifyUserId

`func (o *ApiV1AuditMetadata) SetModifyUserId(v string)`

SetModifyUserId sets ModifyUserId field to given value.

### HasModifyUserId

`func (o *ApiV1AuditMetadata) HasModifyUserId() bool`

HasModifyUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


