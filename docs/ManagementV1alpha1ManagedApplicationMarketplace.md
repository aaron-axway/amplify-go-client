# ManagementV1alpha1ManagedApplicationMarketplace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the Marketplace. | 
**Resource** | [**CatalogV1alpha1ApplicationMarketplaceResource**](CatalogV1alpha1ApplicationMarketplaceResource.md) |  | 

## Methods

### NewManagementV1alpha1ManagedApplicationMarketplace

`func NewManagementV1alpha1ManagedApplicationMarketplace(name string, resource CatalogV1alpha1ApplicationMarketplaceResource, ) *ManagementV1alpha1ManagedApplicationMarketplace`

NewManagementV1alpha1ManagedApplicationMarketplace instantiates a new ManagementV1alpha1ManagedApplicationMarketplace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementV1alpha1ManagedApplicationMarketplaceWithDefaults

`func NewManagementV1alpha1ManagedApplicationMarketplaceWithDefaults() *ManagementV1alpha1ManagedApplicationMarketplace`

NewManagementV1alpha1ManagedApplicationMarketplaceWithDefaults instantiates a new ManagementV1alpha1ManagedApplicationMarketplace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ManagementV1alpha1ManagedApplicationMarketplace) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ManagementV1alpha1ManagedApplicationMarketplace) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ManagementV1alpha1ManagedApplicationMarketplace) SetName(v string)`

SetName sets Name field to given value.


### GetResource

`func (o *ManagementV1alpha1ManagedApplicationMarketplace) GetResource() CatalogV1alpha1ApplicationMarketplaceResource`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *ManagementV1alpha1ManagedApplicationMarketplace) GetResourceOk() (*CatalogV1alpha1ApplicationMarketplaceResource, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *ManagementV1alpha1ManagedApplicationMarketplace) SetResource(v CatalogV1alpha1ApplicationMarketplaceResource)`

SetResource sets Resource field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


