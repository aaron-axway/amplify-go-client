# ApiV1LanguageValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | Pointer to **string** | The json path for where the value is defined | [optional] 
**Value** | Pointer to **map[string]interface{}** | The value | [optional] 
**Status** | Pointer to **string** | Status of the value for the path. * \&quot;complete\&quot; value provided and does not need to be reviewed. * \&quot;toReview\&quot; value provided, but it needs to be reviewed. * \&quot;undefined\&quot; value not provided.  | [optional] 

## Methods

### NewApiV1LanguageValuesInner

`func NewApiV1LanguageValuesInner() *ApiV1LanguageValuesInner`

NewApiV1LanguageValuesInner instantiates a new ApiV1LanguageValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV1LanguageValuesInnerWithDefaults

`func NewApiV1LanguageValuesInnerWithDefaults() *ApiV1LanguageValuesInner`

NewApiV1LanguageValuesInnerWithDefaults instantiates a new ApiV1LanguageValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *ApiV1LanguageValuesInner) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ApiV1LanguageValuesInner) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ApiV1LanguageValuesInner) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *ApiV1LanguageValuesInner) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetValue

`func (o *ApiV1LanguageValuesInner) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ApiV1LanguageValuesInner) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ApiV1LanguageValuesInner) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *ApiV1LanguageValuesInner) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetStatus

`func (o *ApiV1LanguageValuesInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApiV1LanguageValuesInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApiV1LanguageValuesInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApiV1LanguageValuesInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


