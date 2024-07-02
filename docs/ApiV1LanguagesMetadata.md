# ApiV1LanguagesMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Additional** | Pointer to [**[]ApiV1LanguagesMetadataAdditionalInner**](ApiV1LanguagesMetadataAdditionalInner.md) | Additional languages available for the resource instance. | [optional] 
**Audit** | Pointer to [**ApiV1AuditMetadata**](ApiV1AuditMetadata.md) |  | [optional] 

## Methods

### NewApiV1LanguagesMetadata

`func NewApiV1LanguagesMetadata() *ApiV1LanguagesMetadata`

NewApiV1LanguagesMetadata instantiates a new ApiV1LanguagesMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV1LanguagesMetadataWithDefaults

`func NewApiV1LanguagesMetadataWithDefaults() *ApiV1LanguagesMetadata`

NewApiV1LanguagesMetadataWithDefaults instantiates a new ApiV1LanguagesMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditional

`func (o *ApiV1LanguagesMetadata) GetAdditional() []ApiV1LanguagesMetadataAdditionalInner`

GetAdditional returns the Additional field if non-nil, zero value otherwise.

### GetAdditionalOk

`func (o *ApiV1LanguagesMetadata) GetAdditionalOk() (*[]ApiV1LanguagesMetadataAdditionalInner, bool)`

GetAdditionalOk returns a tuple with the Additional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditional

`func (o *ApiV1LanguagesMetadata) SetAdditional(v []ApiV1LanguagesMetadataAdditionalInner)`

SetAdditional sets Additional field to given value.

### HasAdditional

`func (o *ApiV1LanguagesMetadata) HasAdditional() bool`

HasAdditional returns a boolean if a field has been set.

### GetAudit

`func (o *ApiV1LanguagesMetadata) GetAudit() ApiV1AuditMetadata`

GetAudit returns the Audit field if non-nil, zero value otherwise.

### GetAuditOk

`func (o *ApiV1LanguagesMetadata) GetAuditOk() (*ApiV1AuditMetadata, bool)`

GetAuditOk returns a tuple with the Audit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudit

`func (o *ApiV1LanguagesMetadata) SetAudit(v ApiV1AuditMetadata)`

SetAudit sets Audit field to given value.

### HasAudit

`func (o *ApiV1LanguagesMetadata) HasAudit() bool`

HasAudit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


