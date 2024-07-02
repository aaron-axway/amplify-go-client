# ApiV1EmbeddedReferenceMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Internal id of the resource. | [optional] 
**AccessRights** | Pointer to [**ApiV1EmbeddedReferenceMetadataAccessRights**](ApiV1EmbeddedReferenceMetadataAccessRights.md) |  | [optional] 
**SelfLink** | Pointer to **string** | The URL representing the scope resource object. | [optional] 

## Methods

### NewApiV1EmbeddedReferenceMetadata

`func NewApiV1EmbeddedReferenceMetadata() *ApiV1EmbeddedReferenceMetadata`

NewApiV1EmbeddedReferenceMetadata instantiates a new ApiV1EmbeddedReferenceMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV1EmbeddedReferenceMetadataWithDefaults

`func NewApiV1EmbeddedReferenceMetadataWithDefaults() *ApiV1EmbeddedReferenceMetadata`

NewApiV1EmbeddedReferenceMetadataWithDefaults instantiates a new ApiV1EmbeddedReferenceMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApiV1EmbeddedReferenceMetadata) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiV1EmbeddedReferenceMetadata) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiV1EmbeddedReferenceMetadata) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApiV1EmbeddedReferenceMetadata) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccessRights

`func (o *ApiV1EmbeddedReferenceMetadata) GetAccessRights() ApiV1EmbeddedReferenceMetadataAccessRights`

GetAccessRights returns the AccessRights field if non-nil, zero value otherwise.

### GetAccessRightsOk

`func (o *ApiV1EmbeddedReferenceMetadata) GetAccessRightsOk() (*ApiV1EmbeddedReferenceMetadataAccessRights, bool)`

GetAccessRightsOk returns a tuple with the AccessRights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessRights

`func (o *ApiV1EmbeddedReferenceMetadata) SetAccessRights(v ApiV1EmbeddedReferenceMetadataAccessRights)`

SetAccessRights sets AccessRights field to given value.

### HasAccessRights

`func (o *ApiV1EmbeddedReferenceMetadata) HasAccessRights() bool`

HasAccessRights returns a boolean if a field has been set.

### GetSelfLink

`func (o *ApiV1EmbeddedReferenceMetadata) GetSelfLink() string`

GetSelfLink returns the SelfLink field if non-nil, zero value otherwise.

### GetSelfLinkOk

`func (o *ApiV1EmbeddedReferenceMetadata) GetSelfLinkOk() (*string, bool)`

GetSelfLinkOk returns a tuple with the SelfLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfLink

`func (o *ApiV1EmbeddedReferenceMetadata) SetSelfLink(v string)`

SetSelfLink sets SelfLink field to given value.

### HasSelfLink

`func (o *ApiV1EmbeddedReferenceMetadata) HasSelfLink() bool`

HasSelfLink returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


