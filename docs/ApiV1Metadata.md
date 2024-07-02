# ApiV1Metadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Internal id of the resource. | [optional] 
**Audit** | Pointer to [**ApiV1AuditMetadata**](ApiV1AuditMetadata.md) |  | [optional] 
**Scope** | Pointer to [**ApiV1MetadataScope**](ApiV1MetadataScope.md) |  | [optional] 
**Acls** | Pointer to [**[]ApiV1Acl**](ApiV1Acl.md) | Access Control Lists for this resource. | [optional] 
**ResourceVersion** | Pointer to **string** | Internal version of this object that can be used by clients to determine when objects have changed. | [optional] 
**References** | Pointer to [**[]ApiV1Reference**](ApiV1Reference.md) |  | [optional] 
**DeletedReferences** | Pointer to [**[]ApiV1Reference**](ApiV1Reference.md) |  | [optional] 
**State** | Pointer to **string** | Defines the state of the resource. If present, indicates if the resource is in Deleting state. | [optional] 
**SelfLink** | Pointer to **string** | The URL representing this resource object. | [optional] 

## Methods

### NewApiV1Metadata

`func NewApiV1Metadata() *ApiV1Metadata`

NewApiV1Metadata instantiates a new ApiV1Metadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV1MetadataWithDefaults

`func NewApiV1MetadataWithDefaults() *ApiV1Metadata`

NewApiV1MetadataWithDefaults instantiates a new ApiV1Metadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApiV1Metadata) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiV1Metadata) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiV1Metadata) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApiV1Metadata) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAudit

`func (o *ApiV1Metadata) GetAudit() ApiV1AuditMetadata`

GetAudit returns the Audit field if non-nil, zero value otherwise.

### GetAuditOk

`func (o *ApiV1Metadata) GetAuditOk() (*ApiV1AuditMetadata, bool)`

GetAuditOk returns a tuple with the Audit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudit

`func (o *ApiV1Metadata) SetAudit(v ApiV1AuditMetadata)`

SetAudit sets Audit field to given value.

### HasAudit

`func (o *ApiV1Metadata) HasAudit() bool`

HasAudit returns a boolean if a field has been set.

### GetScope

`func (o *ApiV1Metadata) GetScope() ApiV1MetadataScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *ApiV1Metadata) GetScopeOk() (*ApiV1MetadataScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *ApiV1Metadata) SetScope(v ApiV1MetadataScope)`

SetScope sets Scope field to given value.

### HasScope

`func (o *ApiV1Metadata) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetAcls

`func (o *ApiV1Metadata) GetAcls() []ApiV1Acl`

GetAcls returns the Acls field if non-nil, zero value otherwise.

### GetAclsOk

`func (o *ApiV1Metadata) GetAclsOk() (*[]ApiV1Acl, bool)`

GetAclsOk returns a tuple with the Acls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcls

`func (o *ApiV1Metadata) SetAcls(v []ApiV1Acl)`

SetAcls sets Acls field to given value.

### HasAcls

`func (o *ApiV1Metadata) HasAcls() bool`

HasAcls returns a boolean if a field has been set.

### GetResourceVersion

`func (o *ApiV1Metadata) GetResourceVersion() string`

GetResourceVersion returns the ResourceVersion field if non-nil, zero value otherwise.

### GetResourceVersionOk

`func (o *ApiV1Metadata) GetResourceVersionOk() (*string, bool)`

GetResourceVersionOk returns a tuple with the ResourceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceVersion

`func (o *ApiV1Metadata) SetResourceVersion(v string)`

SetResourceVersion sets ResourceVersion field to given value.

### HasResourceVersion

`func (o *ApiV1Metadata) HasResourceVersion() bool`

HasResourceVersion returns a boolean if a field has been set.

### GetReferences

`func (o *ApiV1Metadata) GetReferences() []ApiV1Reference`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *ApiV1Metadata) GetReferencesOk() (*[]ApiV1Reference, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *ApiV1Metadata) SetReferences(v []ApiV1Reference)`

SetReferences sets References field to given value.

### HasReferences

`func (o *ApiV1Metadata) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetDeletedReferences

`func (o *ApiV1Metadata) GetDeletedReferences() []ApiV1Reference`

GetDeletedReferences returns the DeletedReferences field if non-nil, zero value otherwise.

### GetDeletedReferencesOk

`func (o *ApiV1Metadata) GetDeletedReferencesOk() (*[]ApiV1Reference, bool)`

GetDeletedReferencesOk returns a tuple with the DeletedReferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedReferences

`func (o *ApiV1Metadata) SetDeletedReferences(v []ApiV1Reference)`

SetDeletedReferences sets DeletedReferences field to given value.

### HasDeletedReferences

`func (o *ApiV1Metadata) HasDeletedReferences() bool`

HasDeletedReferences returns a boolean if a field has been set.

### GetState

`func (o *ApiV1Metadata) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ApiV1Metadata) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ApiV1Metadata) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ApiV1Metadata) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSelfLink

`func (o *ApiV1Metadata) GetSelfLink() string`

GetSelfLink returns the SelfLink field if non-nil, zero value otherwise.

### GetSelfLinkOk

`func (o *ApiV1Metadata) GetSelfLinkOk() (*string, bool)`

GetSelfLinkOk returns a tuple with the SelfLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfLink

`func (o *ApiV1Metadata) SetSelfLink(v string)`

SetSelfLink sets SelfLink field to given value.

### HasSelfLink

`func (o *ApiV1Metadata) HasSelfLink() bool`

HasSelfLink returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


