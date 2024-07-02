# ApiV1ResourceInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | Pointer to **string** | Defines the group from which the resource belongs to. The server infers this from the endpoint the client submits the request to. | [optional] 
**ApiVersion** | Pointer to **string** | Defines the structure of the resource. The server infers this from the endpoint the client submits the request to. | [optional] 
**Kind** | Pointer to **string** | Defines the kind of the resource. The server infers this from the endpoint the client submits the request to. | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to [**ApiV1Metadata**](ApiV1Metadata.md) |  | [optional] 
**Attributes** | Pointer to **map[string]string** | Custom attributes added to objects. | [optional] 
**Finalizers** | Pointer to [**[]ApiV1Finalizer**](ApiV1Finalizer.md) |  | [optional] 
**Owner** | Pointer to [**ApiV1Owner**](ApiV1Owner.md) |  | [optional] 
**Tags** | Pointer to **[]string** | List of tags. | [optional] 
**Spec** | **map[string]interface{}** | Resource instance specs. | 
**Embedded** | Pointer to [**ApiV1Embedded**](ApiV1Embedded.md) |  | [optional] 

## Methods

### NewApiV1ResourceInstance

`func NewApiV1ResourceInstance(spec map[string]interface{}, ) *ApiV1ResourceInstance`

NewApiV1ResourceInstance instantiates a new ApiV1ResourceInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV1ResourceInstanceWithDefaults

`func NewApiV1ResourceInstanceWithDefaults() *ApiV1ResourceInstance`

NewApiV1ResourceInstanceWithDefaults instantiates a new ApiV1ResourceInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *ApiV1ResourceInstance) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *ApiV1ResourceInstance) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *ApiV1ResourceInstance) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *ApiV1ResourceInstance) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetApiVersion

`func (o *ApiV1ResourceInstance) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ApiV1ResourceInstance) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ApiV1ResourceInstance) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *ApiV1ResourceInstance) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *ApiV1ResourceInstance) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ApiV1ResourceInstance) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ApiV1ResourceInstance) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ApiV1ResourceInstance) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetName

`func (o *ApiV1ResourceInstance) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiV1ResourceInstance) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiV1ResourceInstance) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiV1ResourceInstance) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTitle

`func (o *ApiV1ResourceInstance) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ApiV1ResourceInstance) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ApiV1ResourceInstance) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ApiV1ResourceInstance) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetMetadata

`func (o *ApiV1ResourceInstance) GetMetadata() ApiV1Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ApiV1ResourceInstance) GetMetadataOk() (*ApiV1Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ApiV1ResourceInstance) SetMetadata(v ApiV1Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ApiV1ResourceInstance) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetAttributes

`func (o *ApiV1ResourceInstance) GetAttributes() map[string]string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ApiV1ResourceInstance) GetAttributesOk() (*map[string]string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ApiV1ResourceInstance) SetAttributes(v map[string]string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *ApiV1ResourceInstance) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetFinalizers

`func (o *ApiV1ResourceInstance) GetFinalizers() []ApiV1Finalizer`

GetFinalizers returns the Finalizers field if non-nil, zero value otherwise.

### GetFinalizersOk

`func (o *ApiV1ResourceInstance) GetFinalizersOk() (*[]ApiV1Finalizer, bool)`

GetFinalizersOk returns a tuple with the Finalizers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalizers

`func (o *ApiV1ResourceInstance) SetFinalizers(v []ApiV1Finalizer)`

SetFinalizers sets Finalizers field to given value.

### HasFinalizers

`func (o *ApiV1ResourceInstance) HasFinalizers() bool`

HasFinalizers returns a boolean if a field has been set.

### GetOwner

`func (o *ApiV1ResourceInstance) GetOwner() ApiV1Owner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ApiV1ResourceInstance) GetOwnerOk() (*ApiV1Owner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ApiV1ResourceInstance) SetOwner(v ApiV1Owner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ApiV1ResourceInstance) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetTags

`func (o *ApiV1ResourceInstance) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ApiV1ResourceInstance) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ApiV1ResourceInstance) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ApiV1ResourceInstance) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetSpec

`func (o *ApiV1ResourceInstance) GetSpec() map[string]interface{}`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *ApiV1ResourceInstance) GetSpecOk() (*map[string]interface{}, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *ApiV1ResourceInstance) SetSpec(v map[string]interface{})`

SetSpec sets Spec field to given value.


### GetEmbedded

`func (o *ApiV1ResourceInstance) GetEmbedded() ApiV1Embedded`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *ApiV1ResourceInstance) GetEmbeddedOk() (*ApiV1Embedded, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *ApiV1ResourceInstance) SetEmbedded(v ApiV1Embedded)`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *ApiV1ResourceInstance) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


