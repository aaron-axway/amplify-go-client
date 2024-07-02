# ManagementV1alpha1BatchJobProgress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageNumber** | Pointer to **int32** | The last 1-based page number requested from API Server. | [optional] 
**PageSize** | Pointer to **int32** | The number of resources requested per page from API Server. | [optional] 
**QueryParams** | Pointer to **string** | The query parameters used in the request. | [optional] 
**ScopeName** | Pointer to **string** | Name of scoped resource the job fetches resources from. Undefined for unscoped searches. | [optional] 

## Methods

### NewManagementV1alpha1BatchJobProgress

`func NewManagementV1alpha1BatchJobProgress() *ManagementV1alpha1BatchJobProgress`

NewManagementV1alpha1BatchJobProgress instantiates a new ManagementV1alpha1BatchJobProgress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementV1alpha1BatchJobProgressWithDefaults

`func NewManagementV1alpha1BatchJobProgressWithDefaults() *ManagementV1alpha1BatchJobProgress`

NewManagementV1alpha1BatchJobProgressWithDefaults instantiates a new ManagementV1alpha1BatchJobProgress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageNumber

`func (o *ManagementV1alpha1BatchJobProgress) GetPageNumber() int32`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *ManagementV1alpha1BatchJobProgress) GetPageNumberOk() (*int32, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *ManagementV1alpha1BatchJobProgress) SetPageNumber(v int32)`

SetPageNumber sets PageNumber field to given value.

### HasPageNumber

`func (o *ManagementV1alpha1BatchJobProgress) HasPageNumber() bool`

HasPageNumber returns a boolean if a field has been set.

### GetPageSize

`func (o *ManagementV1alpha1BatchJobProgress) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ManagementV1alpha1BatchJobProgress) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ManagementV1alpha1BatchJobProgress) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *ManagementV1alpha1BatchJobProgress) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetQueryParams

`func (o *ManagementV1alpha1BatchJobProgress) GetQueryParams() string`

GetQueryParams returns the QueryParams field if non-nil, zero value otherwise.

### GetQueryParamsOk

`func (o *ManagementV1alpha1BatchJobProgress) GetQueryParamsOk() (*string, bool)`

GetQueryParamsOk returns a tuple with the QueryParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryParams

`func (o *ManagementV1alpha1BatchJobProgress) SetQueryParams(v string)`

SetQueryParams sets QueryParams field to given value.

### HasQueryParams

`func (o *ManagementV1alpha1BatchJobProgress) HasQueryParams() bool`

HasQueryParams returns a boolean if a field has been set.

### GetScopeName

`func (o *ManagementV1alpha1BatchJobProgress) GetScopeName() string`

GetScopeName returns the ScopeName field if non-nil, zero value otherwise.

### GetScopeNameOk

`func (o *ManagementV1alpha1BatchJobProgress) GetScopeNameOk() (*string, bool)`

GetScopeNameOk returns a tuple with the ScopeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeName

`func (o *ManagementV1alpha1BatchJobProgress) SetScopeName(v string)`

SetScopeName sets ScopeName field to given value.

### HasScopeName

`func (o *ManagementV1alpha1BatchJobProgress) HasScopeName() bool`

HasScopeName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


