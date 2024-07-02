# ApiV1Language

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**ApiV1LanguageMetadata**](ApiV1LanguageMetadata.md) |  | [optional] 
**Values** | Pointer to [**[]ApiV1LanguageValuesInner**](ApiV1LanguageValuesInner.md) | Defines the language specific values set on the resource. | [optional] 

## Methods

### NewApiV1Language

`func NewApiV1Language() *ApiV1Language`

NewApiV1Language instantiates a new ApiV1Language object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV1LanguageWithDefaults

`func NewApiV1LanguageWithDefaults() *ApiV1Language`

NewApiV1LanguageWithDefaults instantiates a new ApiV1Language object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *ApiV1Language) GetMetadata() ApiV1LanguageMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ApiV1Language) GetMetadataOk() (*ApiV1LanguageMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ApiV1Language) SetMetadata(v ApiV1LanguageMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ApiV1Language) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetValues

`func (o *ApiV1Language) GetValues() []ApiV1LanguageValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *ApiV1Language) GetValuesOk() (*[]ApiV1LanguageValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *ApiV1Language) SetValues(v []ApiV1LanguageValuesInner)`

SetValues sets Values field to given value.

### HasValues

`func (o *ApiV1Language) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


