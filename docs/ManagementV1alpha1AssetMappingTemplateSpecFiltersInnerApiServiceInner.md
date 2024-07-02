# ManagementV1alpha1AssetMappingTemplateSpecFiltersInnerApiServiceInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the API Service. | [optional] 
**Attributes** | Pointer to **map[string]string** | Attributes used to filter the APIServiceRevisions for the API Service on which the template applies. | [optional] 
**ApiServiceRevision** | Pointer to [**ManagementV1alpha1AssetMappingTemplateSpecFiltersInnerApiServiceInnerApiServiceRevision**](ManagementV1alpha1AssetMappingTemplateSpecFiltersInnerApiServiceInnerApiServiceRevision.md) |  | [optional] 

## Methods

### NewManagementV1alpha1AssetMappingTemplateSpecFiltersInnerApiServiceInner

`func NewManagementV1alpha1AssetMappingTemplateSpecFiltersInnerApiServiceInner() *ManagementV1alpha1AssetMappingTemplateSpecFiltersInnerApiServiceInner`

NewManagementV1alpha1AssetMappingTemplateSpecFiltersInnerApiServiceInner instantiates a new ManagementV1alpha1AssetMappingTemplateSpecFiltersInnerApiServiceInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementV1alpha1AssetMappingTemplateSpecFiltersInnerApiServiceInnerWithDefaults

`func NewManagementV1alpha1AssetMappingTemplateSpecFiltersInnerApiServiceInnerWithDefaults() *ManagementV1alpha1AssetMappingTemplateSpecFiltersInnerApiServiceInner`

NewManagementV1alpha1AssetMappingTemplateSpecFiltersInnerApiServiceInnerWithDefaults instantiates a new ManagementV1alpha1AssetMappingTemplateSpecFiltersInnerApiServiceInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ManagementV1alpha1AssetMappingTemplateSpecFiltersInnerApiServiceInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ManagementV1alpha1AssetMappingTemplateSpecFiltersInnerApiServiceInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ManagementV1alpha1AssetMappingTemplateSpecFiltersInnerApiServiceInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ManagementV1alpha1AssetMappingTemplateSpecFiltersInnerApiServiceInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAttributes

`func (o *ManagementV1alpha1AssetMappingTemplateSpecFiltersInnerApiServiceInner) GetAttributes() map[string]string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ManagementV1alpha1AssetMappingTemplateSpecFiltersInnerApiServiceInner) GetAttributesOk() (*map[string]string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ManagementV1alpha1AssetMappingTemplateSpecFiltersInnerApiServiceInner) SetAttributes(v map[string]string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *ManagementV1alpha1AssetMappingTemplateSpecFiltersInnerApiServiceInner) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetApiServiceRevision

`func (o *ManagementV1alpha1AssetMappingTemplateSpecFiltersInnerApiServiceInner) GetApiServiceRevision() ManagementV1alpha1AssetMappingTemplateSpecFiltersInnerApiServiceInnerApiServiceRevision`

GetApiServiceRevision returns the ApiServiceRevision field if non-nil, zero value otherwise.

### GetApiServiceRevisionOk

`func (o *ManagementV1alpha1AssetMappingTemplateSpecFiltersInnerApiServiceInner) GetApiServiceRevisionOk() (*ManagementV1alpha1AssetMappingTemplateSpecFiltersInnerApiServiceInnerApiServiceRevision, bool)`

GetApiServiceRevisionOk returns a tuple with the ApiServiceRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiServiceRevision

`func (o *ManagementV1alpha1AssetMappingTemplateSpecFiltersInnerApiServiceInner) SetApiServiceRevision(v ManagementV1alpha1AssetMappingTemplateSpecFiltersInnerApiServiceInnerApiServiceRevision)`

SetApiServiceRevision sets ApiServiceRevision field to given value.

### HasApiServiceRevision

`func (o *ManagementV1alpha1AssetMappingTemplateSpecFiltersInnerApiServiceInner) HasApiServiceRevision() bool`

HasApiServiceRevision returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


