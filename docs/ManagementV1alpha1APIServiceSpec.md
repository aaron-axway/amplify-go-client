# ManagementV1alpha1APIServiceSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of the api service. | [optional] 
**Categories** | Pointer to **[]string** |  | [optional] 
**Icon** | Pointer to [**ManagementV1alpha1APIServiceSpecIcon**](ManagementV1alpha1APIServiceSpecIcon.md) |  | [optional] 

## Methods

### NewManagementV1alpha1APIServiceSpec

`func NewManagementV1alpha1APIServiceSpec() *ManagementV1alpha1APIServiceSpec`

NewManagementV1alpha1APIServiceSpec instantiates a new ManagementV1alpha1APIServiceSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementV1alpha1APIServiceSpecWithDefaults

`func NewManagementV1alpha1APIServiceSpecWithDefaults() *ManagementV1alpha1APIServiceSpec`

NewManagementV1alpha1APIServiceSpecWithDefaults instantiates a new ManagementV1alpha1APIServiceSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ManagementV1alpha1APIServiceSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ManagementV1alpha1APIServiceSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ManagementV1alpha1APIServiceSpec) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ManagementV1alpha1APIServiceSpec) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCategories

`func (o *ManagementV1alpha1APIServiceSpec) GetCategories() []string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *ManagementV1alpha1APIServiceSpec) GetCategoriesOk() (*[]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *ManagementV1alpha1APIServiceSpec) SetCategories(v []string)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *ManagementV1alpha1APIServiceSpec) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetIcon

`func (o *ManagementV1alpha1APIServiceSpec) GetIcon() ManagementV1alpha1APIServiceSpecIcon`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *ManagementV1alpha1APIServiceSpec) GetIconOk() (*ManagementV1alpha1APIServiceSpecIcon, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *ManagementV1alpha1APIServiceSpec) SetIcon(v ManagementV1alpha1APIServiceSpecIcon)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *ManagementV1alpha1APIServiceSpec) HasIcon() bool`

HasIcon returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


