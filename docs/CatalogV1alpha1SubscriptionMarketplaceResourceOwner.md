# CatalogV1alpha1SubscriptionMarketplaceResourceOwner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of the owner. | [optional] 
**Id** | Pointer to **string** | Id of the owner of the resource. | [optional] 
**Organization** | [**CatalogV1alpha1ApplicationMarketplaceResourceOwnerOrganization**](CatalogV1alpha1ApplicationMarketplaceResourceOwnerOrganization.md) |  | 

## Methods

### NewCatalogV1alpha1SubscriptionMarketplaceResourceOwner

`func NewCatalogV1alpha1SubscriptionMarketplaceResourceOwner(organization CatalogV1alpha1ApplicationMarketplaceResourceOwnerOrganization, ) *CatalogV1alpha1SubscriptionMarketplaceResourceOwner`

NewCatalogV1alpha1SubscriptionMarketplaceResourceOwner instantiates a new CatalogV1alpha1SubscriptionMarketplaceResourceOwner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogV1alpha1SubscriptionMarketplaceResourceOwnerWithDefaults

`func NewCatalogV1alpha1SubscriptionMarketplaceResourceOwnerWithDefaults() *CatalogV1alpha1SubscriptionMarketplaceResourceOwner`

NewCatalogV1alpha1SubscriptionMarketplaceResourceOwnerWithDefaults instantiates a new CatalogV1alpha1SubscriptionMarketplaceResourceOwner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CatalogV1alpha1SubscriptionMarketplaceResourceOwner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CatalogV1alpha1SubscriptionMarketplaceResourceOwner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CatalogV1alpha1SubscriptionMarketplaceResourceOwner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CatalogV1alpha1SubscriptionMarketplaceResourceOwner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *CatalogV1alpha1SubscriptionMarketplaceResourceOwner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CatalogV1alpha1SubscriptionMarketplaceResourceOwner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CatalogV1alpha1SubscriptionMarketplaceResourceOwner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CatalogV1alpha1SubscriptionMarketplaceResourceOwner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrganization

`func (o *CatalogV1alpha1SubscriptionMarketplaceResourceOwner) GetOrganization() CatalogV1alpha1ApplicationMarketplaceResourceOwnerOrganization`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *CatalogV1alpha1SubscriptionMarketplaceResourceOwner) GetOrganizationOk() (*CatalogV1alpha1ApplicationMarketplaceResourceOwnerOrganization, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *CatalogV1alpha1SubscriptionMarketplaceResourceOwner) SetOrganization(v CatalogV1alpha1ApplicationMarketplaceResourceOwnerOrganization)`

SetOrganization sets Organization field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


