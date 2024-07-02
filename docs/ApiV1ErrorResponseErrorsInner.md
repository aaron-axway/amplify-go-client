# ApiV1ErrorResponseErrorsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **int32** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Detail** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to [**ApiV1ErrorResponseErrorsInnerMeta**](ApiV1ErrorResponseErrorsInnerMeta.md) |  | [optional] 

## Methods

### NewApiV1ErrorResponseErrorsInner

`func NewApiV1ErrorResponseErrorsInner() *ApiV1ErrorResponseErrorsInner`

NewApiV1ErrorResponseErrorsInner instantiates a new ApiV1ErrorResponseErrorsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV1ErrorResponseErrorsInnerWithDefaults

`func NewApiV1ErrorResponseErrorsInnerWithDefaults() *ApiV1ErrorResponseErrorsInner`

NewApiV1ErrorResponseErrorsInnerWithDefaults instantiates a new ApiV1ErrorResponseErrorsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ApiV1ErrorResponseErrorsInner) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApiV1ErrorResponseErrorsInner) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApiV1ErrorResponseErrorsInner) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApiV1ErrorResponseErrorsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTitle

`func (o *ApiV1ErrorResponseErrorsInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ApiV1ErrorResponseErrorsInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ApiV1ErrorResponseErrorsInner) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ApiV1ErrorResponseErrorsInner) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDetail

`func (o *ApiV1ErrorResponseErrorsInner) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *ApiV1ErrorResponseErrorsInner) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *ApiV1ErrorResponseErrorsInner) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *ApiV1ErrorResponseErrorsInner) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetMeta

`func (o *ApiV1ErrorResponseErrorsInner) GetMeta() ApiV1ErrorResponseErrorsInnerMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ApiV1ErrorResponseErrorsInner) GetMetaOk() (*ApiV1ErrorResponseErrorsInnerMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ApiV1ErrorResponseErrorsInner) SetMeta(v ApiV1ErrorResponseErrorsInnerMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ApiV1ErrorResponseErrorsInner) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


