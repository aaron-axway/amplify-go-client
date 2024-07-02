# ManagementV1alpha1TraceabilityAgentSpecConfigRedactionResponseHeaders

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Show** | Pointer to **[]string** | The regular expressions for response headers that, when matched, will be saved for Business and Consumer Insights | [optional] 
**Sanitize** | Pointer to [**[]ManagementV1alpha1TraceabilityAgentSpecConfigRedactionQueryArgumentSanitizeInner**](ManagementV1alpha1TraceabilityAgentSpecConfigRedactionQueryArgumentSanitizeInner.md) | The regular expressions for response headers keys that, when matched, will sanitize based off the value regular expression before saving for Business and Consumer Insights | [optional] 

## Methods

### NewManagementV1alpha1TraceabilityAgentSpecConfigRedactionResponseHeaders

`func NewManagementV1alpha1TraceabilityAgentSpecConfigRedactionResponseHeaders() *ManagementV1alpha1TraceabilityAgentSpecConfigRedactionResponseHeaders`

NewManagementV1alpha1TraceabilityAgentSpecConfigRedactionResponseHeaders instantiates a new ManagementV1alpha1TraceabilityAgentSpecConfigRedactionResponseHeaders object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementV1alpha1TraceabilityAgentSpecConfigRedactionResponseHeadersWithDefaults

`func NewManagementV1alpha1TraceabilityAgentSpecConfigRedactionResponseHeadersWithDefaults() *ManagementV1alpha1TraceabilityAgentSpecConfigRedactionResponseHeaders`

NewManagementV1alpha1TraceabilityAgentSpecConfigRedactionResponseHeadersWithDefaults instantiates a new ManagementV1alpha1TraceabilityAgentSpecConfigRedactionResponseHeaders object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShow

`func (o *ManagementV1alpha1TraceabilityAgentSpecConfigRedactionResponseHeaders) GetShow() []string`

GetShow returns the Show field if non-nil, zero value otherwise.

### GetShowOk

`func (o *ManagementV1alpha1TraceabilityAgentSpecConfigRedactionResponseHeaders) GetShowOk() (*[]string, bool)`

GetShowOk returns a tuple with the Show field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShow

`func (o *ManagementV1alpha1TraceabilityAgentSpecConfigRedactionResponseHeaders) SetShow(v []string)`

SetShow sets Show field to given value.

### HasShow

`func (o *ManagementV1alpha1TraceabilityAgentSpecConfigRedactionResponseHeaders) HasShow() bool`

HasShow returns a boolean if a field has been set.

### GetSanitize

`func (o *ManagementV1alpha1TraceabilityAgentSpecConfigRedactionResponseHeaders) GetSanitize() []ManagementV1alpha1TraceabilityAgentSpecConfigRedactionQueryArgumentSanitizeInner`

GetSanitize returns the Sanitize field if non-nil, zero value otherwise.

### GetSanitizeOk

`func (o *ManagementV1alpha1TraceabilityAgentSpecConfigRedactionResponseHeaders) GetSanitizeOk() (*[]ManagementV1alpha1TraceabilityAgentSpecConfigRedactionQueryArgumentSanitizeInner, bool)`

GetSanitizeOk returns a tuple with the Sanitize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSanitize

`func (o *ManagementV1alpha1TraceabilityAgentSpecConfigRedactionResponseHeaders) SetSanitize(v []ManagementV1alpha1TraceabilityAgentSpecConfigRedactionQueryArgumentSanitizeInner)`

SetSanitize sets Sanitize field to given value.

### HasSanitize

`func (o *ManagementV1alpha1TraceabilityAgentSpecConfigRedactionResponseHeaders) HasSanitize() bool`

HasSanitize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


