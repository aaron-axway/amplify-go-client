# ManagementV1alpha1SecretSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **map[string]string** | Key value pairs. The value will be stored encrypted. | [optional] 

## Methods

### NewManagementV1alpha1SecretSpec

`func NewManagementV1alpha1SecretSpec() *ManagementV1alpha1SecretSpec`

NewManagementV1alpha1SecretSpec instantiates a new ManagementV1alpha1SecretSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementV1alpha1SecretSpecWithDefaults

`func NewManagementV1alpha1SecretSpecWithDefaults() *ManagementV1alpha1SecretSpec`

NewManagementV1alpha1SecretSpecWithDefaults instantiates a new ManagementV1alpha1SecretSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ManagementV1alpha1SecretSpec) GetData() map[string]string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ManagementV1alpha1SecretSpec) GetDataOk() (*map[string]string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ManagementV1alpha1SecretSpec) SetData(v map[string]string)`

SetData sets Data field to given value.

### HasData

`func (o *ManagementV1alpha1SecretSpec) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


