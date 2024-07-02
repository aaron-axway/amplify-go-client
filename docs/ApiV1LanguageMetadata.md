# ApiV1LanguageMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | Language code. | [optional] 
**Status** | Pointer to **string** | Status of the language. * \&quot;complete\&quot; language available for the resource with no missing properties or properties that need to be reviewed. * \&quot;toReview\&quot; language has properties that are missing or need to be reviewed. * \&quot;undefined\&quot; language is not defined on the resource.  | [optional] 
**Audit** | Pointer to [**ApiV1AuditMetadata**](ApiV1AuditMetadata.md) |  | [optional] 

## Methods

### NewApiV1LanguageMetadata

`func NewApiV1LanguageMetadata() *ApiV1LanguageMetadata`

NewApiV1LanguageMetadata instantiates a new ApiV1LanguageMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV1LanguageMetadataWithDefaults

`func NewApiV1LanguageMetadataWithDefaults() *ApiV1LanguageMetadata`

NewApiV1LanguageMetadataWithDefaults instantiates a new ApiV1LanguageMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ApiV1LanguageMetadata) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ApiV1LanguageMetadata) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ApiV1LanguageMetadata) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ApiV1LanguageMetadata) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetStatus

`func (o *ApiV1LanguageMetadata) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApiV1LanguageMetadata) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApiV1LanguageMetadata) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApiV1LanguageMetadata) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAudit

`func (o *ApiV1LanguageMetadata) GetAudit() ApiV1AuditMetadata`

GetAudit returns the Audit field if non-nil, zero value otherwise.

### GetAuditOk

`func (o *ApiV1LanguageMetadata) GetAuditOk() (*ApiV1AuditMetadata, bool)`

GetAuditOk returns a tuple with the Audit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudit

`func (o *ApiV1LanguageMetadata) SetAudit(v ApiV1AuditMetadata)`

SetAudit sets Audit field to given value.

### HasAudit

`func (o *ApiV1LanguageMetadata) HasAudit() bool`

HasAudit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


