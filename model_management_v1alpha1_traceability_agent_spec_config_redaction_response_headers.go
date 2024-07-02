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

// checks if the ManagementV1alpha1TraceabilityAgentSpecConfigRedactionResponseHeaders type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManagementV1alpha1TraceabilityAgentSpecConfigRedactionResponseHeaders{}

// ManagementV1alpha1TraceabilityAgentSpecConfigRedactionResponseHeaders The allowed and sanitization setup of response headers in transactional data
type ManagementV1alpha1TraceabilityAgentSpecConfigRedactionResponseHeaders struct {
	// The regular expressions for response headers that, when matched, will be saved for Business and Consumer Insights
	Show []string `json:"show,omitempty"`
	// The regular expressions for response headers keys that, when matched, will sanitize based off the value regular expression before saving for Business and Consumer Insights
	Sanitize []ManagementV1alpha1TraceabilityAgentSpecConfigRedactionQueryArgumentSanitizeInner `json:"sanitize,omitempty"`
}

// NewManagementV1alpha1TraceabilityAgentSpecConfigRedactionResponseHeaders instantiates a new ManagementV1alpha1TraceabilityAgentSpecConfigRedactionResponseHeaders object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagementV1alpha1TraceabilityAgentSpecConfigRedactionResponseHeaders() *ManagementV1alpha1TraceabilityAgentSpecConfigRedactionResponseHeaders {
	this := ManagementV1alpha1TraceabilityAgentSpecConfigRedactionResponseHeaders{}
	return &this
}

// NewManagementV1alpha1TraceabilityAgentSpecConfigRedactionResponseHeadersWithDefaults instantiates a new ManagementV1alpha1TraceabilityAgentSpecConfigRedactionResponseHeaders object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagementV1alpha1TraceabilityAgentSpecConfigRedactionResponseHeadersWithDefaults() *ManagementV1alpha1TraceabilityAgentSpecConfigRedactionResponseHeaders {
	this := ManagementV1alpha1TraceabilityAgentSpecConfigRedactionResponseHeaders{}
	return &this
}

// GetShow returns the Show field value if set, zero value otherwise.
func (o *ManagementV1alpha1TraceabilityAgentSpecConfigRedactionResponseHeaders) GetShow() []string {
	if o == nil || IsNil(o.Show) {
		var ret []string
		return ret
	}
	return o.Show
}

// GetShowOk returns a tuple with the Show field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1TraceabilityAgentSpecConfigRedactionResponseHeaders) GetShowOk() ([]string, bool) {
	if o == nil || IsNil(o.Show) {
		return nil, false
	}
	return o.Show, true
}

// HasShow returns a boolean if a field has been set.
func (o *ManagementV1alpha1TraceabilityAgentSpecConfigRedactionResponseHeaders) HasShow() bool {
	if o != nil && !IsNil(o.Show) {
		return true
	}

	return false
}

// SetShow gets a reference to the given []string and assigns it to the Show field.
func (o *ManagementV1alpha1TraceabilityAgentSpecConfigRedactionResponseHeaders) SetShow(v []string) {
	o.Show = v
}

// GetSanitize returns the Sanitize field value if set, zero value otherwise.
func (o *ManagementV1alpha1TraceabilityAgentSpecConfigRedactionResponseHeaders) GetSanitize() []ManagementV1alpha1TraceabilityAgentSpecConfigRedactionQueryArgumentSanitizeInner {
	if o == nil || IsNil(o.Sanitize) {
		var ret []ManagementV1alpha1TraceabilityAgentSpecConfigRedactionQueryArgumentSanitizeInner
		return ret
	}
	return o.Sanitize
}

// GetSanitizeOk returns a tuple with the Sanitize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1TraceabilityAgentSpecConfigRedactionResponseHeaders) GetSanitizeOk() ([]ManagementV1alpha1TraceabilityAgentSpecConfigRedactionQueryArgumentSanitizeInner, bool) {
	if o == nil || IsNil(o.Sanitize) {
		return nil, false
	}
	return o.Sanitize, true
}

// HasSanitize returns a boolean if a field has been set.
func (o *ManagementV1alpha1TraceabilityAgentSpecConfigRedactionResponseHeaders) HasSanitize() bool {
	if o != nil && !IsNil(o.Sanitize) {
		return true
	}

	return false
}

// SetSanitize gets a reference to the given []ManagementV1alpha1TraceabilityAgentSpecConfigRedactionQueryArgumentSanitizeInner and assigns it to the Sanitize field.
func (o *ManagementV1alpha1TraceabilityAgentSpecConfigRedactionResponseHeaders) SetSanitize(v []ManagementV1alpha1TraceabilityAgentSpecConfigRedactionQueryArgumentSanitizeInner) {
	o.Sanitize = v
}

func (o ManagementV1alpha1TraceabilityAgentSpecConfigRedactionResponseHeaders) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManagementV1alpha1TraceabilityAgentSpecConfigRedactionResponseHeaders) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Show) {
		toSerialize["show"] = o.Show
	}
	if !IsNil(o.Sanitize) {
		toSerialize["sanitize"] = o.Sanitize
	}
	return toSerialize, nil
}

type NullableManagementV1alpha1TraceabilityAgentSpecConfigRedactionResponseHeaders struct {
	value *ManagementV1alpha1TraceabilityAgentSpecConfigRedactionResponseHeaders
	isSet bool
}

func (v NullableManagementV1alpha1TraceabilityAgentSpecConfigRedactionResponseHeaders) Get() *ManagementV1alpha1TraceabilityAgentSpecConfigRedactionResponseHeaders {
	return v.value
}

func (v *NullableManagementV1alpha1TraceabilityAgentSpecConfigRedactionResponseHeaders) Set(val *ManagementV1alpha1TraceabilityAgentSpecConfigRedactionResponseHeaders) {
	v.value = val
	v.isSet = true
}

func (v NullableManagementV1alpha1TraceabilityAgentSpecConfigRedactionResponseHeaders) IsSet() bool {
	return v.isSet
}

func (v *NullableManagementV1alpha1TraceabilityAgentSpecConfigRedactionResponseHeaders) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagementV1alpha1TraceabilityAgentSpecConfigRedactionResponseHeaders(val *ManagementV1alpha1TraceabilityAgentSpecConfigRedactionResponseHeaders) *NullableManagementV1alpha1TraceabilityAgentSpecConfigRedactionResponseHeaders {
	return &NullableManagementV1alpha1TraceabilityAgentSpecConfigRedactionResponseHeaders{value: val, isSet: true}
}

func (v NullableManagementV1alpha1TraceabilityAgentSpecConfigRedactionResponseHeaders) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagementV1alpha1TraceabilityAgentSpecConfigRedactionResponseHeaders) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

