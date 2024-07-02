# ManagementV1alpha1TraceabilityAgentSpecConfigRedaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | Pointer to **[]string** | The regular expressions for URL paths that, when matched, will be saved for Business and Consumer Insights | [optional] 
**QueryArgument** | Pointer to [**ManagementV1alpha1TraceabilityAgentSpecConfigRedactionQueryArgument**](ManagementV1alpha1TraceabilityAgentSpecConfigRedactionQueryArgument.md) |  | [optional] 
**RequestHeaders** | Pointer to [**ManagementV1alpha1TraceabilityAgentSpecConfigRedactionRequestHeaders**](ManagementV1alpha1TraceabilityAgentSpecConfigRedactionRequestHeaders.md) |  | [optional] 
**ResponseHeaders** | Pointer to [**ManagementV1alpha1TraceabilityAgentSpecConfigRedactionResponseHeaders**](ManagementV1alpha1TraceabilityAgentSpecConfigRedactionResponseHeaders.md) |  | [optional] 
**MaskingCharacter** | Pointer to **string** | The character(s) to apply to values that are sanitized before saving for Business and Consumer Insights. Max length of 5 Default is &#39;{*}&#39; | [optional] 

## Methods

### NewManagementV1alpha1TraceabilityAgentSpecConfigRedaction

`func NewManagementV1alpha1TraceabilityAgentSpecConfigRedaction() *ManagementV1alpha1TraceabilityAgentSpecConfigRedaction`

NewManagementV1alpha1TraceabilityAgentSpecConfigRedaction instantiates a new ManagementV1alpha1TraceabilityAgentSpecConfigRedaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementV1alpha1TraceabilityAgentSpecConfigRedactionWithDefaults

`func NewManagementV1alpha1TraceabilityAgentSpecConfigRedactionWithDefaults() *ManagementV1alpha1TraceabilityAgentSpecConfigRedaction`

NewManagementV1alpha1TraceabilityAgentSpecConfigRedactionWithDefaults instantiates a new ManagementV1alpha1TraceabilityAgentSpecConfigRedaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *ManagementV1alpha1TraceabilityAgentSpecConfigRedaction) GetPath() []string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ManagementV1alpha1TraceabilityAgentSpecConfigRedaction) GetPathOk() (*[]string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ManagementV1alpha1TraceabilityAgentSpecConfigRedaction) SetPath(v []string)`

SetPath sets Path field to given value.

### HasPath

`func (o *ManagementV1alpha1TraceabilityAgentSpecConfigRedaction) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetQueryArgument

`func (o *ManagementV1alpha1TraceabilityAgentSpecConfigRedaction) GetQueryArgument() ManagementV1alpha1TraceabilityAgentSpecConfigRedactionQueryArgument`

GetQueryArgument returns the QueryArgument field if non-nil, zero value otherwise.

### GetQueryArgumentOk

`func (o *ManagementV1alpha1TraceabilityAgentSpecConfigRedaction) GetQueryArgumentOk() (*ManagementV1alpha1TraceabilityAgentSpecConfigRedactionQueryArgument, bool)`

GetQueryArgumentOk returns a tuple with the QueryArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryArgument

`func (o *ManagementV1alpha1TraceabilityAgentSpecConfigRedaction) SetQueryArgument(v ManagementV1alpha1TraceabilityAgentSpecConfigRedactionQueryArgument)`

SetQueryArgument sets QueryArgument field to given value.

### HasQueryArgument

`func (o *ManagementV1alpha1TraceabilityAgentSpecConfigRedaction) HasQueryArgument() bool`

HasQueryArgument returns a boolean if a field has been set.

### GetRequestHeaders

`func (o *ManagementV1alpha1TraceabilityAgentSpecConfigRedaction) GetRequestHeaders() ManagementV1alpha1TraceabilityAgentSpecConfigRedactionRequestHeaders`

GetRequestHeaders returns the RequestHeaders field if non-nil, zero value otherwise.

### GetRequestHeadersOk

`func (o *ManagementV1alpha1TraceabilityAgentSpecConfigRedaction) GetRequestHeadersOk() (*ManagementV1alpha1TraceabilityAgentSpecConfigRedactionRequestHeaders, bool)`

GetRequestHeadersOk returns a tuple with the RequestHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestHeaders

`func (o *ManagementV1alpha1TraceabilityAgentSpecConfigRedaction) SetRequestHeaders(v ManagementV1alpha1TraceabilityAgentSpecConfigRedactionRequestHeaders)`

SetRequestHeaders sets RequestHeaders field to given value.

### HasRequestHeaders

`func (o *ManagementV1alpha1TraceabilityAgentSpecConfigRedaction) HasRequestHeaders() bool`

HasRequestHeaders returns a boolean if a field has been set.

### GetResponseHeaders

`func (o *ManagementV1alpha1TraceabilityAgentSpecConfigRedaction) GetResponseHeaders() ManagementV1alpha1TraceabilityAgentSpecConfigRedactionResponseHeaders`

GetResponseHeaders returns the ResponseHeaders field if non-nil, zero value otherwise.

### GetResponseHeadersOk

`func (o *ManagementV1alpha1TraceabilityAgentSpecConfigRedaction) GetResponseHeadersOk() (*ManagementV1alpha1TraceabilityAgentSpecConfigRedactionResponseHeaders, bool)`

GetResponseHeadersOk returns a tuple with the ResponseHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseHeaders

`func (o *ManagementV1alpha1TraceabilityAgentSpecConfigRedaction) SetResponseHeaders(v ManagementV1alpha1TraceabilityAgentSpecConfigRedactionResponseHeaders)`

SetResponseHeaders sets ResponseHeaders field to given value.

### HasResponseHeaders

`func (o *ManagementV1alpha1TraceabilityAgentSpecConfigRedaction) HasResponseHeaders() bool`

HasResponseHeaders returns a boolean if a field has been set.

### GetMaskingCharacter

`func (o *ManagementV1alpha1TraceabilityAgentSpecConfigRedaction) GetMaskingCharacter() string`

GetMaskingCharacter returns the MaskingCharacter field if non-nil, zero value otherwise.

### GetMaskingCharacterOk

`func (o *ManagementV1alpha1TraceabilityAgentSpecConfigRedaction) GetMaskingCharacterOk() (*string, bool)`

GetMaskingCharacterOk returns a tuple with the MaskingCharacter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaskingCharacter

`func (o *ManagementV1alpha1TraceabilityAgentSpecConfigRedaction) SetMaskingCharacter(v string)`

SetMaskingCharacter sets MaskingCharacter field to given value.

### HasMaskingCharacter

`func (o *ManagementV1alpha1TraceabilityAgentSpecConfigRedaction) HasMaskingCharacter() bool`

HasMaskingCharacter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


