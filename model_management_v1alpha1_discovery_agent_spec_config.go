/*
AMPLIFY Central API v0.347

APIs to manage AMPLIFY Central configuration resources.

API version: 0.347.0
Contact: support@axway.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ManagementV1alpha1DiscoveryAgentSpecConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManagementV1alpha1DiscoveryAgentSpecConfig{}

// ManagementV1alpha1DiscoveryAgentSpecConfig Represents the discovery agent config
type ManagementV1alpha1DiscoveryAgentSpecConfig struct {
	// Defines the tag based filter expression to be evaluated for discovering the API from the API Gateway. The filter value is a conditional expression that can use logical operators to compare two value. The conditional expression must have \"tag\" as the prefix/selector in the symbol name. For e.g. ``` tag.SOME_TAG == \"somevalue\" ``` The expression can be a simple condition as shown above or compound condition in which more than one simple conditions are evaluated using logical operator. For e.g. ``` tag.SOME_TAG == \"somevalue\" || tag.ANOTHER_TAG != \"some_other_value\" ``` In addition to logical expression, the filter can hold call based expressions. Below are the list of supported call expressions #### Exists Exists call can be made to evaluate if the tag name exists in the list of tags on API. This call expression can be used as unary expression For e.g. ``` tag.SOME_TAG.Exists() ``` #### Any Any call can be made in a simple expression to evaluate if the tag with any name has specified value or not in the list of tags on the API. For e.g. ``` tag.Any() == \"Tag with some value\" || tag.Any() != \"Tag with other value\" ``` #### Contains Contains call can be made in a simple expression to evaluate if the the specified tag contains specified argument as value. This call expression requires string argument that will be used to perform lookup in tag value For e.g. ``` tag.Contains(\"somevalue\") ``` #### MatchRegEx MatchRegEx call can be used for evaluating the specified tag value to match specified regular expression. This call expression requires a regular expression as the argument. For e.g. ``` tag.MatchRegEx(\"(some){1}\") ``` 
	Filter *string `json:"filter,omitempty"`
	// The list of tags to be added to the API service resource that the agent publishes to Amplify Central
	AdditionalTags []string `json:"additionalTags,omitempty"`
	// The list of tags to exclude from the API service resource that the agent publishes to Amplify Central
	IgnoreTags []string `json:"ignoreTags,omitempty"`
	// Name of the team that owns the catalog item created by agent. If not provided, the default team will be used.
	OwningTeam *string `json:"owningTeam,omitempty"`
	Owner *ManagementV1alpha1DiscoveryAgentSpecConfigOwner `json:"owner,omitempty"`
}

// NewManagementV1alpha1DiscoveryAgentSpecConfig instantiates a new ManagementV1alpha1DiscoveryAgentSpecConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagementV1alpha1DiscoveryAgentSpecConfig() *ManagementV1alpha1DiscoveryAgentSpecConfig {
	this := ManagementV1alpha1DiscoveryAgentSpecConfig{}
	return &this
}

// NewManagementV1alpha1DiscoveryAgentSpecConfigWithDefaults instantiates a new ManagementV1alpha1DiscoveryAgentSpecConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagementV1alpha1DiscoveryAgentSpecConfigWithDefaults() *ManagementV1alpha1DiscoveryAgentSpecConfig {
	this := ManagementV1alpha1DiscoveryAgentSpecConfig{}
	return &this
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *ManagementV1alpha1DiscoveryAgentSpecConfig) GetFilter() string {
	if o == nil || IsNil(o.Filter) {
		var ret string
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1DiscoveryAgentSpecConfig) GetFilterOk() (*string, bool) {
	if o == nil || IsNil(o.Filter) {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *ManagementV1alpha1DiscoveryAgentSpecConfig) HasFilter() bool {
	if o != nil && !IsNil(o.Filter) {
		return true
	}

	return false
}

// SetFilter gets a reference to the given string and assigns it to the Filter field.
func (o *ManagementV1alpha1DiscoveryAgentSpecConfig) SetFilter(v string) {
	o.Filter = &v
}

// GetAdditionalTags returns the AdditionalTags field value if set, zero value otherwise.
func (o *ManagementV1alpha1DiscoveryAgentSpecConfig) GetAdditionalTags() []string {
	if o == nil || IsNil(o.AdditionalTags) {
		var ret []string
		return ret
	}
	return o.AdditionalTags
}

// GetAdditionalTagsOk returns a tuple with the AdditionalTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1DiscoveryAgentSpecConfig) GetAdditionalTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.AdditionalTags) {
		return nil, false
	}
	return o.AdditionalTags, true
}

// HasAdditionalTags returns a boolean if a field has been set.
func (o *ManagementV1alpha1DiscoveryAgentSpecConfig) HasAdditionalTags() bool {
	if o != nil && !IsNil(o.AdditionalTags) {
		return true
	}

	return false
}

// SetAdditionalTags gets a reference to the given []string and assigns it to the AdditionalTags field.
func (o *ManagementV1alpha1DiscoveryAgentSpecConfig) SetAdditionalTags(v []string) {
	o.AdditionalTags = v
}

// GetIgnoreTags returns the IgnoreTags field value if set, zero value otherwise.
func (o *ManagementV1alpha1DiscoveryAgentSpecConfig) GetIgnoreTags() []string {
	if o == nil || IsNil(o.IgnoreTags) {
		var ret []string
		return ret
	}
	return o.IgnoreTags
}

// GetIgnoreTagsOk returns a tuple with the IgnoreTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1DiscoveryAgentSpecConfig) GetIgnoreTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.IgnoreTags) {
		return nil, false
	}
	return o.IgnoreTags, true
}

// HasIgnoreTags returns a boolean if a field has been set.
func (o *ManagementV1alpha1DiscoveryAgentSpecConfig) HasIgnoreTags() bool {
	if o != nil && !IsNil(o.IgnoreTags) {
		return true
	}

	return false
}

// SetIgnoreTags gets a reference to the given []string and assigns it to the IgnoreTags field.
func (o *ManagementV1alpha1DiscoveryAgentSpecConfig) SetIgnoreTags(v []string) {
	o.IgnoreTags = v
}

// GetOwningTeam returns the OwningTeam field value if set, zero value otherwise.
func (o *ManagementV1alpha1DiscoveryAgentSpecConfig) GetOwningTeam() string {
	if o == nil || IsNil(o.OwningTeam) {
		var ret string
		return ret
	}
	return *o.OwningTeam
}

// GetOwningTeamOk returns a tuple with the OwningTeam field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1DiscoveryAgentSpecConfig) GetOwningTeamOk() (*string, bool) {
	if o == nil || IsNil(o.OwningTeam) {
		return nil, false
	}
	return o.OwningTeam, true
}

// HasOwningTeam returns a boolean if a field has been set.
func (o *ManagementV1alpha1DiscoveryAgentSpecConfig) HasOwningTeam() bool {
	if o != nil && !IsNil(o.OwningTeam) {
		return true
	}

	return false
}

// SetOwningTeam gets a reference to the given string and assigns it to the OwningTeam field.
func (o *ManagementV1alpha1DiscoveryAgentSpecConfig) SetOwningTeam(v string) {
	o.OwningTeam = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *ManagementV1alpha1DiscoveryAgentSpecConfig) GetOwner() ManagementV1alpha1DiscoveryAgentSpecConfigOwner {
	if o == nil || IsNil(o.Owner) {
		var ret ManagementV1alpha1DiscoveryAgentSpecConfigOwner
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1DiscoveryAgentSpecConfig) GetOwnerOk() (*ManagementV1alpha1DiscoveryAgentSpecConfigOwner, bool) {
	if o == nil || IsNil(o.Owner) {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *ManagementV1alpha1DiscoveryAgentSpecConfig) HasOwner() bool {
	if o != nil && !IsNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given ManagementV1alpha1DiscoveryAgentSpecConfigOwner and assigns it to the Owner field.
func (o *ManagementV1alpha1DiscoveryAgentSpecConfig) SetOwner(v ManagementV1alpha1DiscoveryAgentSpecConfigOwner) {
	o.Owner = &v
}

func (o ManagementV1alpha1DiscoveryAgentSpecConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManagementV1alpha1DiscoveryAgentSpecConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Filter) {
		toSerialize["filter"] = o.Filter
	}
	if !IsNil(o.AdditionalTags) {
		toSerialize["additionalTags"] = o.AdditionalTags
	}
	if !IsNil(o.IgnoreTags) {
		toSerialize["ignoreTags"] = o.IgnoreTags
	}
	if !IsNil(o.OwningTeam) {
		toSerialize["owningTeam"] = o.OwningTeam
	}
	if !IsNil(o.Owner) {
		toSerialize["owner"] = o.Owner
	}
	return toSerialize, nil
}

type NullableManagementV1alpha1DiscoveryAgentSpecConfig struct {
	value *ManagementV1alpha1DiscoveryAgentSpecConfig
	isSet bool
}

func (v NullableManagementV1alpha1DiscoveryAgentSpecConfig) Get() *ManagementV1alpha1DiscoveryAgentSpecConfig {
	return v.value
}

func (v *NullableManagementV1alpha1DiscoveryAgentSpecConfig) Set(val *ManagementV1alpha1DiscoveryAgentSpecConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableManagementV1alpha1DiscoveryAgentSpecConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableManagementV1alpha1DiscoveryAgentSpecConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagementV1alpha1DiscoveryAgentSpecConfig(val *ManagementV1alpha1DiscoveryAgentSpecConfig) *NullableManagementV1alpha1DiscoveryAgentSpecConfig {
	return &NullableManagementV1alpha1DiscoveryAgentSpecConfig{value: val, isSet: true}
}

func (v NullableManagementV1alpha1DiscoveryAgentSpecConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagementV1alpha1DiscoveryAgentSpecConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


