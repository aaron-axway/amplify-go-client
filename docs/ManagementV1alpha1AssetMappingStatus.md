# ManagementV1alpha1AssetMappingStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Level** | Pointer to **string** |  | [optional] 
**Source** | Pointer to [**ManagementV1alpha1AssetMappingStatusSource**](ManagementV1alpha1AssetMappingStatusSource.md) |  | [optional] 
**Outputs** | Pointer to [**[]ManagementV1alpha1AssetMappingStatusOutputsInner**](ManagementV1alpha1AssetMappingStatusOutputsInner.md) | Generated catalog resources. | [optional] 

## Methods

### NewManagementV1alpha1AssetMappingStatus

`func NewManagementV1alpha1AssetMappingStatus() *ManagementV1alpha1AssetMappingStatus`

NewManagementV1alpha1AssetMappingStatus instantiates a new ManagementV1alpha1AssetMappingStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementV1alpha1AssetMappingStatusWithDefaults

`func NewManagementV1alpha1AssetMappingStatusWithDefaults() *ManagementV1alpha1AssetMappingStatus`

NewManagementV1alpha1AssetMappingStatusWithDefaults instantiates a new ManagementV1alpha1AssetMappingStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLevel

`func (o *ManagementV1alpha1AssetMappingStatus) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *ManagementV1alpha1AssetMappingStatus) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *ManagementV1alpha1AssetMappingStatus) SetLevel(v string)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *ManagementV1alpha1AssetMappingStatus) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetSource

`func (o *ManagementV1alpha1AssetMappingStatus) GetSource() ManagementV1alpha1AssetMappingStatusSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ManagementV1alpha1AssetMappingStatus) GetSourceOk() (*ManagementV1alpha1AssetMappingStatusSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ManagementV1alpha1AssetMappingStatus) SetSource(v ManagementV1alpha1AssetMappingStatusSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *ManagementV1alpha1AssetMappingStatus) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetOutputs

`func (o *ManagementV1alpha1AssetMappingStatus) GetOutputs() []ManagementV1alpha1AssetMappingStatusOutputsInner`

GetOutputs returns the Outputs field if non-nil, zero value otherwise.

### GetOutputsOk

`func (o *ManagementV1alpha1AssetMappingStatus) GetOutputsOk() (*[]ManagementV1alpha1AssetMappingStatusOutputsInner, bool)`

GetOutputsOk returns a tuple with the Outputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputs

`func (o *ManagementV1alpha1AssetMappingStatus) SetOutputs(v []ManagementV1alpha1AssetMappingStatusOutputsInner)`

SetOutputs sets Outputs field to given value.

### HasOutputs

`func (o *ManagementV1alpha1AssetMappingStatus) HasOutputs() bool`

HasOutputs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


