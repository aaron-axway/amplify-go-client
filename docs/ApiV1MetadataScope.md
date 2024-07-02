# ApiV1MetadataScope

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Internal id of the scope resource where the resource is defined. | [optional] 
**Kind** | Pointer to **string** | The kind of the scope resource where the resource is defined. | [optional] 
**Name** | Pointer to **string** | The name of the scope where the resource is defined. | [optional] 
**Title** | Pointer to **string** | The title of the scope where the resource is defined. | [optional] 
**SelfLink** | Pointer to **string** | The URL representing the scope resource object. | [optional] 
**Owner** | Pointer to [**ApiV1Owner**](ApiV1Owner.md) |  | [optional] 

## Methods

### NewApiV1MetadataScope

`func NewApiV1MetadataScope() *ApiV1MetadataScope`

NewApiV1MetadataScope instantiates a new ApiV1MetadataScope object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV1MetadataScopeWithDefaults

`func NewApiV1MetadataScopeWithDefaults() *ApiV1MetadataScope`

NewApiV1MetadataScopeWithDefaults instantiates a new ApiV1MetadataScope object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApiV1MetadataScope) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiV1MetadataScope) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiV1MetadataScope) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApiV1MetadataScope) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *ApiV1MetadataScope) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ApiV1MetadataScope) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ApiV1MetadataScope) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ApiV1MetadataScope) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetName

`func (o *ApiV1MetadataScope) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiV1MetadataScope) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiV1MetadataScope) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiV1MetadataScope) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTitle

`func (o *ApiV1MetadataScope) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ApiV1MetadataScope) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ApiV1MetadataScope) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ApiV1MetadataScope) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetSelfLink

`func (o *ApiV1MetadataScope) GetSelfLink() string`

GetSelfLink returns the SelfLink field if non-nil, zero value otherwise.

### GetSelfLinkOk

`func (o *ApiV1MetadataScope) GetSelfLinkOk() (*string, bool)`

GetSelfLinkOk returns a tuple with the SelfLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfLink

`func (o *ApiV1MetadataScope) SetSelfLink(v string)`

SetSelfLink sets SelfLink field to given value.

### HasSelfLink

`func (o *ApiV1MetadataScope) HasSelfLink() bool`

HasSelfLink returns a boolean if a field has been set.

### GetOwner

`func (o *ApiV1MetadataScope) GetOwner() ApiV1Owner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ApiV1MetadataScope) GetOwnerOk() (*ApiV1Owner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ApiV1MetadataScope) SetOwner(v ApiV1Owner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ApiV1MetadataScope) HasOwner() bool`

HasOwner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


