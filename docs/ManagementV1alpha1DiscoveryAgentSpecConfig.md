# ManagementV1alpha1DiscoveryAgentSpecConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to **string** | Defines the tag based filter expression to be evaluated for discovering the API from the API Gateway. The filter value is a conditional expression that can use logical operators to compare two value. The conditional expression must have \&quot;tag\&quot; as the prefix/selector in the symbol name. For e.g. &#x60;&#x60;&#x60; tag.SOME_TAG &#x3D;&#x3D; \&quot;somevalue\&quot; &#x60;&#x60;&#x60; The expression can be a simple condition as shown above or compound condition in which more than one simple conditions are evaluated using logical operator. For e.g. &#x60;&#x60;&#x60; tag.SOME_TAG &#x3D;&#x3D; \&quot;somevalue\&quot; || tag.ANOTHER_TAG !&#x3D; \&quot;some_other_value\&quot; &#x60;&#x60;&#x60; In addition to logical expression, the filter can hold call based expressions. Below are the list of supported call expressions #### Exists Exists call can be made to evaluate if the tag name exists in the list of tags on API. This call expression can be used as unary expression For e.g. &#x60;&#x60;&#x60; tag.SOME_TAG.Exists() &#x60;&#x60;&#x60; #### Any Any call can be made in a simple expression to evaluate if the tag with any name has specified value or not in the list of tags on the API. For e.g. &#x60;&#x60;&#x60; tag.Any() &#x3D;&#x3D; \&quot;Tag with some value\&quot; || tag.Any() !&#x3D; \&quot;Tag with other value\&quot; &#x60;&#x60;&#x60; #### Contains Contains call can be made in a simple expression to evaluate if the the specified tag contains specified argument as value. This call expression requires string argument that will be used to perform lookup in tag value For e.g. &#x60;&#x60;&#x60; tag.Contains(\&quot;somevalue\&quot;) &#x60;&#x60;&#x60; #### MatchRegEx MatchRegEx call can be used for evaluating the specified tag value to match specified regular expression. This call expression requires a regular expression as the argument. For e.g. &#x60;&#x60;&#x60; tag.MatchRegEx(\&quot;(some){1}\&quot;) &#x60;&#x60;&#x60;  | [optional] 
**AdditionalTags** | Pointer to **[]string** | The list of tags to be added to the API service resource that the agent publishes to Amplify Central | [optional] 
**IgnoreTags** | Pointer to **[]string** | The list of tags to exclude from the API service resource that the agent publishes to Amplify Central | [optional] 
**OwningTeam** | Pointer to **string** | Name of the team that owns the catalog item created by agent. If not provided, the default team will be used. | [optional] 
**Owner** | Pointer to [**ManagementV1alpha1DiscoveryAgentSpecConfigOwner**](ManagementV1alpha1DiscoveryAgentSpecConfigOwner.md) |  | [optional] 

## Methods

### NewManagementV1alpha1DiscoveryAgentSpecConfig

`func NewManagementV1alpha1DiscoveryAgentSpecConfig() *ManagementV1alpha1DiscoveryAgentSpecConfig`

NewManagementV1alpha1DiscoveryAgentSpecConfig instantiates a new ManagementV1alpha1DiscoveryAgentSpecConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementV1alpha1DiscoveryAgentSpecConfigWithDefaults

`func NewManagementV1alpha1DiscoveryAgentSpecConfigWithDefaults() *ManagementV1alpha1DiscoveryAgentSpecConfig`

NewManagementV1alpha1DiscoveryAgentSpecConfigWithDefaults instantiates a new ManagementV1alpha1DiscoveryAgentSpecConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *ManagementV1alpha1DiscoveryAgentSpecConfig) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *ManagementV1alpha1DiscoveryAgentSpecConfig) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *ManagementV1alpha1DiscoveryAgentSpecConfig) SetFilter(v string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *ManagementV1alpha1DiscoveryAgentSpecConfig) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetAdditionalTags

`func (o *ManagementV1alpha1DiscoveryAgentSpecConfig) GetAdditionalTags() []string`

GetAdditionalTags returns the AdditionalTags field if non-nil, zero value otherwise.

### GetAdditionalTagsOk

`func (o *ManagementV1alpha1DiscoveryAgentSpecConfig) GetAdditionalTagsOk() (*[]string, bool)`

GetAdditionalTagsOk returns a tuple with the AdditionalTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalTags

`func (o *ManagementV1alpha1DiscoveryAgentSpecConfig) SetAdditionalTags(v []string)`

SetAdditionalTags sets AdditionalTags field to given value.

### HasAdditionalTags

`func (o *ManagementV1alpha1DiscoveryAgentSpecConfig) HasAdditionalTags() bool`

HasAdditionalTags returns a boolean if a field has been set.

### GetIgnoreTags

`func (o *ManagementV1alpha1DiscoveryAgentSpecConfig) GetIgnoreTags() []string`

GetIgnoreTags returns the IgnoreTags field if non-nil, zero value otherwise.

### GetIgnoreTagsOk

`func (o *ManagementV1alpha1DiscoveryAgentSpecConfig) GetIgnoreTagsOk() (*[]string, bool)`

GetIgnoreTagsOk returns a tuple with the IgnoreTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreTags

`func (o *ManagementV1alpha1DiscoveryAgentSpecConfig) SetIgnoreTags(v []string)`

SetIgnoreTags sets IgnoreTags field to given value.

### HasIgnoreTags

`func (o *ManagementV1alpha1DiscoveryAgentSpecConfig) HasIgnoreTags() bool`

HasIgnoreTags returns a boolean if a field has been set.

### GetOwningTeam

`func (o *ManagementV1alpha1DiscoveryAgentSpecConfig) GetOwningTeam() string`

GetOwningTeam returns the OwningTeam field if non-nil, zero value otherwise.

### GetOwningTeamOk

`func (o *ManagementV1alpha1DiscoveryAgentSpecConfig) GetOwningTeamOk() (*string, bool)`

GetOwningTeamOk returns a tuple with the OwningTeam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwningTeam

`func (o *ManagementV1alpha1DiscoveryAgentSpecConfig) SetOwningTeam(v string)`

SetOwningTeam sets OwningTeam field to given value.

### HasOwningTeam

`func (o *ManagementV1alpha1DiscoveryAgentSpecConfig) HasOwningTeam() bool`

HasOwningTeam returns a boolean if a field has been set.

### GetOwner

`func (o *ManagementV1alpha1DiscoveryAgentSpecConfig) GetOwner() ManagementV1alpha1DiscoveryAgentSpecConfigOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ManagementV1alpha1DiscoveryAgentSpecConfig) GetOwnerOk() (*ManagementV1alpha1DiscoveryAgentSpecConfigOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ManagementV1alpha1DiscoveryAgentSpecConfig) SetOwner(v ManagementV1alpha1DiscoveryAgentSpecConfigOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ManagementV1alpha1DiscoveryAgentSpecConfig) HasOwner() bool`

HasOwner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


