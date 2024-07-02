# ManagementV1alpha1DataplaneSpecGitHub

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**Name** | **string** | The repository name used for api specs discovery | 
**OwnerName** | **string** | The repository owner Name | 
**Filter** | Pointer to [**ManagementV1alpha1DataplaneSpecGitHubFilter**](ManagementV1alpha1DataplaneSpecGitHubFilter.md) |  | [optional] 

## Methods

### NewManagementV1alpha1DataplaneSpecGitHub

`func NewManagementV1alpha1DataplaneSpecGitHub(name string, ownerName string, ) *ManagementV1alpha1DataplaneSpecGitHub`

NewManagementV1alpha1DataplaneSpecGitHub instantiates a new ManagementV1alpha1DataplaneSpecGitHub object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementV1alpha1DataplaneSpecGitHubWithDefaults

`func NewManagementV1alpha1DataplaneSpecGitHubWithDefaults() *ManagementV1alpha1DataplaneSpecGitHub`

NewManagementV1alpha1DataplaneSpecGitHubWithDefaults instantiates a new ManagementV1alpha1DataplaneSpecGitHub object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ManagementV1alpha1DataplaneSpecGitHub) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ManagementV1alpha1DataplaneSpecGitHub) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ManagementV1alpha1DataplaneSpecGitHub) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ManagementV1alpha1DataplaneSpecGitHub) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *ManagementV1alpha1DataplaneSpecGitHub) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ManagementV1alpha1DataplaneSpecGitHub) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ManagementV1alpha1DataplaneSpecGitHub) SetName(v string)`

SetName sets Name field to given value.


### GetOwnerName

`func (o *ManagementV1alpha1DataplaneSpecGitHub) GetOwnerName() string`

GetOwnerName returns the OwnerName field if non-nil, zero value otherwise.

### GetOwnerNameOk

`func (o *ManagementV1alpha1DataplaneSpecGitHub) GetOwnerNameOk() (*string, bool)`

GetOwnerNameOk returns a tuple with the OwnerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerName

`func (o *ManagementV1alpha1DataplaneSpecGitHub) SetOwnerName(v string)`

SetOwnerName sets OwnerName field to given value.


### GetFilter

`func (o *ManagementV1alpha1DataplaneSpecGitHub) GetFilter() ManagementV1alpha1DataplaneSpecGitHubFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *ManagementV1alpha1DataplaneSpecGitHub) GetFilterOk() (*ManagementV1alpha1DataplaneSpecGitHubFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *ManagementV1alpha1DataplaneSpecGitHub) SetFilter(v ManagementV1alpha1DataplaneSpecGitHubFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *ManagementV1alpha1DataplaneSpecGitHub) HasFilter() bool`

HasFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


