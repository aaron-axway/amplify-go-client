# CatalogV1alpha1WebhookSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Auth** | Pointer to [**ManagementV1alpha1WebhookSpecAuth**](ManagementV1alpha1WebhookSpecAuth.md) |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Url** | **string** |  | 
**Headers** | Pointer to **map[string]string** | A list of headers that will be sent as par of the http call to the webhook endpoint. | [optional] 

## Methods

### NewCatalogV1alpha1WebhookSpec

`func NewCatalogV1alpha1WebhookSpec(url string, ) *CatalogV1alpha1WebhookSpec`

NewCatalogV1alpha1WebhookSpec instantiates a new CatalogV1alpha1WebhookSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogV1alpha1WebhookSpecWithDefaults

`func NewCatalogV1alpha1WebhookSpecWithDefaults() *CatalogV1alpha1WebhookSpec`

NewCatalogV1alpha1WebhookSpecWithDefaults instantiates a new CatalogV1alpha1WebhookSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuth

`func (o *CatalogV1alpha1WebhookSpec) GetAuth() ManagementV1alpha1WebhookSpecAuth`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *CatalogV1alpha1WebhookSpec) GetAuthOk() (*ManagementV1alpha1WebhookSpecAuth, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *CatalogV1alpha1WebhookSpec) SetAuth(v ManagementV1alpha1WebhookSpecAuth)`

SetAuth sets Auth field to given value.

### HasAuth

`func (o *CatalogV1alpha1WebhookSpec) HasAuth() bool`

HasAuth returns a boolean if a field has been set.

### GetEnabled

`func (o *CatalogV1alpha1WebhookSpec) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CatalogV1alpha1WebhookSpec) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CatalogV1alpha1WebhookSpec) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CatalogV1alpha1WebhookSpec) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetUrl

`func (o *CatalogV1alpha1WebhookSpec) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CatalogV1alpha1WebhookSpec) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CatalogV1alpha1WebhookSpec) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetHeaders

`func (o *CatalogV1alpha1WebhookSpec) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *CatalogV1alpha1WebhookSpec) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *CatalogV1alpha1WebhookSpec) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *CatalogV1alpha1WebhookSpec) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


