# ApiV1EmbeddedMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReferencedByResources** | Pointer to [**[]ApiV1EmbeddedReferencedByResource**](ApiV1EmbeddedReferencedByResource.md) | Provides details on resources referencing the current resource. Included if embed parameter contains referencedByResources. | [optional] 
**ScopedResources** | Pointer to [**[]ApiV1EmbeddedScopedResources**](ApiV1EmbeddedScopedResources.md) | Provides information about resources scoped to the current resource. Included if embed parameter contains metadata.scopedResources. | [optional] 
**References** | Pointer to [**[]ApiV1EmbeddedReference**](ApiV1EmbeddedReference.md) | Provides additional information about the resources being referenced by the current resource. Included if embed parameter contains metadata.references. | [optional] 

## Methods

### NewApiV1EmbeddedMetadata

`func NewApiV1EmbeddedMetadata() *ApiV1EmbeddedMetadata`

NewApiV1EmbeddedMetadata instantiates a new ApiV1EmbeddedMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV1EmbeddedMetadataWithDefaults

`func NewApiV1EmbeddedMetadataWithDefaults() *ApiV1EmbeddedMetadata`

NewApiV1EmbeddedMetadataWithDefaults instantiates a new ApiV1EmbeddedMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReferencedByResources

`func (o *ApiV1EmbeddedMetadata) GetReferencedByResources() []ApiV1EmbeddedReferencedByResource`

GetReferencedByResources returns the ReferencedByResources field if non-nil, zero value otherwise.

### GetReferencedByResourcesOk

`func (o *ApiV1EmbeddedMetadata) GetReferencedByResourcesOk() (*[]ApiV1EmbeddedReferencedByResource, bool)`

GetReferencedByResourcesOk returns a tuple with the ReferencedByResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencedByResources

`func (o *ApiV1EmbeddedMetadata) SetReferencedByResources(v []ApiV1EmbeddedReferencedByResource)`

SetReferencedByResources sets ReferencedByResources field to given value.

### HasReferencedByResources

`func (o *ApiV1EmbeddedMetadata) HasReferencedByResources() bool`

HasReferencedByResources returns a boolean if a field has been set.

### GetScopedResources

`func (o *ApiV1EmbeddedMetadata) GetScopedResources() []ApiV1EmbeddedScopedResources`

GetScopedResources returns the ScopedResources field if non-nil, zero value otherwise.

### GetScopedResourcesOk

`func (o *ApiV1EmbeddedMetadata) GetScopedResourcesOk() (*[]ApiV1EmbeddedScopedResources, bool)`

GetScopedResourcesOk returns a tuple with the ScopedResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopedResources

`func (o *ApiV1EmbeddedMetadata) SetScopedResources(v []ApiV1EmbeddedScopedResources)`

SetScopedResources sets ScopedResources field to given value.

### HasScopedResources

`func (o *ApiV1EmbeddedMetadata) HasScopedResources() bool`

HasScopedResources returns a boolean if a field has been set.

### GetReferences

`func (o *ApiV1EmbeddedMetadata) GetReferences() []ApiV1EmbeddedReference`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *ApiV1EmbeddedMetadata) GetReferencesOk() (*[]ApiV1EmbeddedReference, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *ApiV1EmbeddedMetadata) SetReferences(v []ApiV1EmbeddedReference)`

SetReferences sets References field to given value.

### HasReferences

`func (o *ApiV1EmbeddedMetadata) HasReferences() bool`

HasReferences returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


