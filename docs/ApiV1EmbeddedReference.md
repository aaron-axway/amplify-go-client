# ApiV1EmbeddedReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | Pointer to **string** | Defines the group from which the resource belongs to. | [optional] 
**ApiVersion** | Pointer to **string** | Defines the api version of the resource. | [optional] 
**Kind** | Pointer to **string** | Defines the kind of the resource. The server infers this from the endpoint the client submits the request to. | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to [**ApiV1EmbeddedReferenceMetadata**](ApiV1EmbeddedReferenceMetadata.md) |  | [optional] 

## Methods

### NewApiV1EmbeddedReference

`func NewApiV1EmbeddedReference() *ApiV1EmbeddedReference`

NewApiV1EmbeddedReference instantiates a new ApiV1EmbeddedReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV1EmbeddedReferenceWithDefaults

`func NewApiV1EmbeddedReferenceWithDefaults() *ApiV1EmbeddedReference`

NewApiV1EmbeddedReferenceWithDefaults instantiates a new ApiV1EmbeddedReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *ApiV1EmbeddedReference) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *ApiV1EmbeddedReference) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *ApiV1EmbeddedReference) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *ApiV1EmbeddedReference) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetApiVersion

`func (o *ApiV1EmbeddedReference) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ApiV1EmbeddedReference) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ApiV1EmbeddedReference) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *ApiV1EmbeddedReference) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *ApiV1EmbeddedReference) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ApiV1EmbeddedReference) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ApiV1EmbeddedReference) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ApiV1EmbeddedReference) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetName

`func (o *ApiV1EmbeddedReference) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiV1EmbeddedReference) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiV1EmbeddedReference) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiV1EmbeddedReference) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTitle

`func (o *ApiV1EmbeddedReference) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ApiV1EmbeddedReference) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ApiV1EmbeddedReference) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ApiV1EmbeddedReference) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetMetadata

`func (o *ApiV1EmbeddedReference) GetMetadata() ApiV1EmbeddedReferenceMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ApiV1EmbeddedReference) GetMetadataOk() (*ApiV1EmbeddedReferenceMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ApiV1EmbeddedReference) SetMetadata(v ApiV1EmbeddedReferenceMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ApiV1EmbeddedReference) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


