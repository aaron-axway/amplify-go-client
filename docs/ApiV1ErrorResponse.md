# ApiV1ErrorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to [**[]ApiV1ErrorResponseErrorsInner**](ApiV1ErrorResponseErrorsInner.md) |  | [optional] 

## Methods

### NewApiV1ErrorResponse

`func NewApiV1ErrorResponse() *ApiV1ErrorResponse`

NewApiV1ErrorResponse instantiates a new ApiV1ErrorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV1ErrorResponseWithDefaults

`func NewApiV1ErrorResponseWithDefaults() *ApiV1ErrorResponse`

NewApiV1ErrorResponseWithDefaults instantiates a new ApiV1ErrorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *ApiV1ErrorResponse) GetErrors() []ApiV1ErrorResponseErrorsInner`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ApiV1ErrorResponse) GetErrorsOk() (*[]ApiV1ErrorResponseErrorsInner, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ApiV1ErrorResponse) SetErrors(v []ApiV1ErrorResponseErrorsInner)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *ApiV1ErrorResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


