# ManagementV1alpha1DataplaneSpecApigeeMetricsFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilteredAPIs** | Pointer to **[]string** | The names of the APIs for which the metrics will be processed | [optional] 
**FilterMetrics** | Pointer to **bool** | The flag upon which is decided if the API metrics are filtered or not. Defaults to true | [optional] [default to true]

## Methods

### NewManagementV1alpha1DataplaneSpecApigeeMetricsFilter

`func NewManagementV1alpha1DataplaneSpecApigeeMetricsFilter() *ManagementV1alpha1DataplaneSpecApigeeMetricsFilter`

NewManagementV1alpha1DataplaneSpecApigeeMetricsFilter instantiates a new ManagementV1alpha1DataplaneSpecApigeeMetricsFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementV1alpha1DataplaneSpecApigeeMetricsFilterWithDefaults

`func NewManagementV1alpha1DataplaneSpecApigeeMetricsFilterWithDefaults() *ManagementV1alpha1DataplaneSpecApigeeMetricsFilter`

NewManagementV1alpha1DataplaneSpecApigeeMetricsFilterWithDefaults instantiates a new ManagementV1alpha1DataplaneSpecApigeeMetricsFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilteredAPIs

`func (o *ManagementV1alpha1DataplaneSpecApigeeMetricsFilter) GetFilteredAPIs() []string`

GetFilteredAPIs returns the FilteredAPIs field if non-nil, zero value otherwise.

### GetFilteredAPIsOk

`func (o *ManagementV1alpha1DataplaneSpecApigeeMetricsFilter) GetFilteredAPIsOk() (*[]string, bool)`

GetFilteredAPIsOk returns a tuple with the FilteredAPIs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilteredAPIs

`func (o *ManagementV1alpha1DataplaneSpecApigeeMetricsFilter) SetFilteredAPIs(v []string)`

SetFilteredAPIs sets FilteredAPIs field to given value.

### HasFilteredAPIs

`func (o *ManagementV1alpha1DataplaneSpecApigeeMetricsFilter) HasFilteredAPIs() bool`

HasFilteredAPIs returns a boolean if a field has been set.

### GetFilterMetrics

`func (o *ManagementV1alpha1DataplaneSpecApigeeMetricsFilter) GetFilterMetrics() bool`

GetFilterMetrics returns the FilterMetrics field if non-nil, zero value otherwise.

### GetFilterMetricsOk

`func (o *ManagementV1alpha1DataplaneSpecApigeeMetricsFilter) GetFilterMetricsOk() (*bool, bool)`

GetFilterMetricsOk returns a tuple with the FilterMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterMetrics

`func (o *ManagementV1alpha1DataplaneSpecApigeeMetricsFilter) SetFilterMetrics(v bool)`

SetFilterMetrics sets FilterMetrics field to given value.

### HasFilterMetrics

`func (o *ManagementV1alpha1DataplaneSpecApigeeMetricsFilter) HasFilterMetrics() bool`

HasFilterMetrics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


