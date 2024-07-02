# CatalogV1alpha1SubscriptionSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **map[string]interface{}** | Data matching the subscription definition schema set on the subscribed plan. | [optional] 
**Schema** | Pointer to **map[string]interface{}** | Schema snapshot from the subscription definition when the subscription got created. | [optional] 
**Product** | **string** | Reference to Product resource | 
**Plan** | [**CatalogV1alpha1SubscriptionSpecPlan**](CatalogV1alpha1SubscriptionSpecPlan.md) |  | 

## Methods

### NewCatalogV1alpha1SubscriptionSpec

`func NewCatalogV1alpha1SubscriptionSpec(product string, plan CatalogV1alpha1SubscriptionSpecPlan, ) *CatalogV1alpha1SubscriptionSpec`

NewCatalogV1alpha1SubscriptionSpec instantiates a new CatalogV1alpha1SubscriptionSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogV1alpha1SubscriptionSpecWithDefaults

`func NewCatalogV1alpha1SubscriptionSpecWithDefaults() *CatalogV1alpha1SubscriptionSpec`

NewCatalogV1alpha1SubscriptionSpecWithDefaults instantiates a new CatalogV1alpha1SubscriptionSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CatalogV1alpha1SubscriptionSpec) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CatalogV1alpha1SubscriptionSpec) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CatalogV1alpha1SubscriptionSpec) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *CatalogV1alpha1SubscriptionSpec) HasData() bool`

HasData returns a boolean if a field has been set.

### GetSchema

`func (o *CatalogV1alpha1SubscriptionSpec) GetSchema() map[string]interface{}`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *CatalogV1alpha1SubscriptionSpec) GetSchemaOk() (*map[string]interface{}, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *CatalogV1alpha1SubscriptionSpec) SetSchema(v map[string]interface{})`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *CatalogV1alpha1SubscriptionSpec) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetProduct

`func (o *CatalogV1alpha1SubscriptionSpec) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *CatalogV1alpha1SubscriptionSpec) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *CatalogV1alpha1SubscriptionSpec) SetProduct(v string)`

SetProduct sets Product field to given value.


### GetPlan

`func (o *CatalogV1alpha1SubscriptionSpec) GetPlan() CatalogV1alpha1SubscriptionSpecPlan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *CatalogV1alpha1SubscriptionSpec) GetPlanOk() (*CatalogV1alpha1SubscriptionSpecPlan, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *CatalogV1alpha1SubscriptionSpec) SetPlan(v CatalogV1alpha1SubscriptionSpecPlan)`

SetPlan sets Plan field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


