# ManagementV1alpha1DataplaneSecretSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dataplane** | **string** |  | 
**Data** | **string** | Key value pairs for accessing the dataplane. The value will be stored encrypted. | 

## Methods

### NewManagementV1alpha1DataplaneSecretSpec

`func NewManagementV1alpha1DataplaneSecretSpec(dataplane string, data string, ) *ManagementV1alpha1DataplaneSecretSpec`

NewManagementV1alpha1DataplaneSecretSpec instantiates a new ManagementV1alpha1DataplaneSecretSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementV1alpha1DataplaneSecretSpecWithDefaults

`func NewManagementV1alpha1DataplaneSecretSpecWithDefaults() *ManagementV1alpha1DataplaneSecretSpec`

NewManagementV1alpha1DataplaneSecretSpecWithDefaults instantiates a new ManagementV1alpha1DataplaneSecretSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataplane

`func (o *ManagementV1alpha1DataplaneSecretSpec) GetDataplane() string`

GetDataplane returns the Dataplane field if non-nil, zero value otherwise.

### GetDataplaneOk

`func (o *ManagementV1alpha1DataplaneSecretSpec) GetDataplaneOk() (*string, bool)`

GetDataplaneOk returns a tuple with the Dataplane field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataplane

`func (o *ManagementV1alpha1DataplaneSecretSpec) SetDataplane(v string)`

SetDataplane sets Dataplane field to given value.


### GetData

`func (o *ManagementV1alpha1DataplaneSecretSpec) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ManagementV1alpha1DataplaneSecretSpec) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ManagementV1alpha1DataplaneSecretSpec) SetData(v string)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


