# CatalogV1alpha1AssetRequestSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | **map[string]interface{}** |  | 
**AssetResource** | **string** |  | 
**Subscription** | Pointer to **string** | reference to the Subscription to be used to access the Asset Resource. | [optional] 
**AssetRequest** | Pointer to **string** | The AssetRequest from which this resource is being migrated from. Reference must be in the same Application. | [optional] 

## Methods

### NewCatalogV1alpha1AssetRequestSpec

`func NewCatalogV1alpha1AssetRequestSpec(data map[string]interface{}, assetResource string, ) *CatalogV1alpha1AssetRequestSpec`

NewCatalogV1alpha1AssetRequestSpec instantiates a new CatalogV1alpha1AssetRequestSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogV1alpha1AssetRequestSpecWithDefaults

`func NewCatalogV1alpha1AssetRequestSpecWithDefaults() *CatalogV1alpha1AssetRequestSpec`

NewCatalogV1alpha1AssetRequestSpecWithDefaults instantiates a new CatalogV1alpha1AssetRequestSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CatalogV1alpha1AssetRequestSpec) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CatalogV1alpha1AssetRequestSpec) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CatalogV1alpha1AssetRequestSpec) SetData(v map[string]interface{})`

SetData sets Data field to given value.


### GetAssetResource

`func (o *CatalogV1alpha1AssetRequestSpec) GetAssetResource() string`

GetAssetResource returns the AssetResource field if non-nil, zero value otherwise.

### GetAssetResourceOk

`func (o *CatalogV1alpha1AssetRequestSpec) GetAssetResourceOk() (*string, bool)`

GetAssetResourceOk returns a tuple with the AssetResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetResource

`func (o *CatalogV1alpha1AssetRequestSpec) SetAssetResource(v string)`

SetAssetResource sets AssetResource field to given value.


### GetSubscription

`func (o *CatalogV1alpha1AssetRequestSpec) GetSubscription() string`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *CatalogV1alpha1AssetRequestSpec) GetSubscriptionOk() (*string, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *CatalogV1alpha1AssetRequestSpec) SetSubscription(v string)`

SetSubscription sets Subscription field to given value.

### HasSubscription

`func (o *CatalogV1alpha1AssetRequestSpec) HasSubscription() bool`

HasSubscription returns a boolean if a field has been set.

### GetAssetRequest

`func (o *CatalogV1alpha1AssetRequestSpec) GetAssetRequest() string`

GetAssetRequest returns the AssetRequest field if non-nil, zero value otherwise.

### GetAssetRequestOk

`func (o *CatalogV1alpha1AssetRequestSpec) GetAssetRequestOk() (*string, bool)`

GetAssetRequestOk returns a tuple with the AssetRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetRequest

`func (o *CatalogV1alpha1AssetRequestSpec) SetAssetRequest(v string)`

SetAssetRequest sets AssetRequest field to given value.

### HasAssetRequest

`func (o *CatalogV1alpha1AssetRequestSpec) HasAssetRequest() bool`

HasAssetRequest returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


