# ManagementV1alpha1ConsumerInstanceSpecAdditionalDataProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Defines the unstructured data type. Example &#39;APIBuilderConnector&#39; or &#39;SDK&#39;. | 
**ContentType** | **string** | Defines the Content Type of the data. | 
**Label** | Pointer to **string** | Short name for the unstructured data. | [optional] 
**FileName** | Pointer to **string** | File name used to be sent as part of the content disposition header for revision unstructured data download. | [optional] 
**Data** | **string** | Base64 encoded data for the Catalog Item revision content. | 

## Methods

### NewManagementV1alpha1ConsumerInstanceSpecAdditionalDataProperties

`func NewManagementV1alpha1ConsumerInstanceSpecAdditionalDataProperties(type_ string, contentType string, data string, ) *ManagementV1alpha1ConsumerInstanceSpecAdditionalDataProperties`

NewManagementV1alpha1ConsumerInstanceSpecAdditionalDataProperties instantiates a new ManagementV1alpha1ConsumerInstanceSpecAdditionalDataProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementV1alpha1ConsumerInstanceSpecAdditionalDataPropertiesWithDefaults

`func NewManagementV1alpha1ConsumerInstanceSpecAdditionalDataPropertiesWithDefaults() *ManagementV1alpha1ConsumerInstanceSpecAdditionalDataProperties`

NewManagementV1alpha1ConsumerInstanceSpecAdditionalDataPropertiesWithDefaults instantiates a new ManagementV1alpha1ConsumerInstanceSpecAdditionalDataProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ManagementV1alpha1ConsumerInstanceSpecAdditionalDataProperties) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ManagementV1alpha1ConsumerInstanceSpecAdditionalDataProperties) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ManagementV1alpha1ConsumerInstanceSpecAdditionalDataProperties) SetType(v string)`

SetType sets Type field to given value.


### GetContentType

`func (o *ManagementV1alpha1ConsumerInstanceSpecAdditionalDataProperties) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *ManagementV1alpha1ConsumerInstanceSpecAdditionalDataProperties) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *ManagementV1alpha1ConsumerInstanceSpecAdditionalDataProperties) SetContentType(v string)`

SetContentType sets ContentType field to given value.


### GetLabel

`func (o *ManagementV1alpha1ConsumerInstanceSpecAdditionalDataProperties) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ManagementV1alpha1ConsumerInstanceSpecAdditionalDataProperties) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ManagementV1alpha1ConsumerInstanceSpecAdditionalDataProperties) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ManagementV1alpha1ConsumerInstanceSpecAdditionalDataProperties) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetFileName

`func (o *ManagementV1alpha1ConsumerInstanceSpecAdditionalDataProperties) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *ManagementV1alpha1ConsumerInstanceSpecAdditionalDataProperties) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *ManagementV1alpha1ConsumerInstanceSpecAdditionalDataProperties) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *ManagementV1alpha1ConsumerInstanceSpecAdditionalDataProperties) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetData

`func (o *ManagementV1alpha1ConsumerInstanceSpecAdditionalDataProperties) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ManagementV1alpha1ConsumerInstanceSpecAdditionalDataProperties) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ManagementV1alpha1ConsumerInstanceSpecAdditionalDataProperties) SetData(v string)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


