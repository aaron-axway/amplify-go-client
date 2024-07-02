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

// checks if the ApiV1EmbeddedReferencedByResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiV1EmbeddedReferencedByResource{}

// ApiV1EmbeddedReferencedByResource struct for ApiV1EmbeddedReferencedByResource
type ApiV1EmbeddedReferencedByResource struct {
	// The kind of the resource.
	Kind *string `json:"kind,omitempty"`
	// The kind of the scope resource where the resource is defined.
	ScopeKind *string `json:"scopeKind,omitempty"`
	// Defines the group from which the resource belongs to.
	Group *string `json:"group,omitempty" validate:"regexp=^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\\\\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*"`
	// Number of resources.
	Count *int32 `json:"count,omitempty"`
	// The link that can be used to access those referencing resources.
	Link *string `json:"link,omitempty"`
}

// NewApiV1EmbeddedReferencedByResource instantiates a new ApiV1EmbeddedReferencedByResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiV1EmbeddedReferencedByResource() *ApiV1EmbeddedReferencedByResource {
	this := ApiV1EmbeddedReferencedByResource{}
	return &this
}

// NewApiV1EmbeddedReferencedByResourceWithDefaults instantiates a new ApiV1EmbeddedReferencedByResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiV1EmbeddedReferencedByResourceWithDefaults() *ApiV1EmbeddedReferencedByResource {
	this := ApiV1EmbeddedReferencedByResource{}
	return &this
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *ApiV1EmbeddedReferencedByResource) GetKind() string {
	if o == nil || IsNil(o.Kind) {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiV1EmbeddedReferencedByResource) GetKindOk() (*string, bool) {
	if o == nil || IsNil(o.Kind) {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *ApiV1EmbeddedReferencedByResource) HasKind() bool {
	if o != nil && !IsNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *ApiV1EmbeddedReferencedByResource) SetKind(v string) {
	o.Kind = &v
}

// GetScopeKind returns the ScopeKind field value if set, zero value otherwise.
func (o *ApiV1EmbeddedReferencedByResource) GetScopeKind() string {
	if o == nil || IsNil(o.ScopeKind) {
		var ret string
		return ret
	}
	return *o.ScopeKind
}

// GetScopeKindOk returns a tuple with the ScopeKind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiV1EmbeddedReferencedByResource) GetScopeKindOk() (*string, bool) {
	if o == nil || IsNil(o.ScopeKind) {
		return nil, false
	}
	return o.ScopeKind, true
}

// HasScopeKind returns a boolean if a field has been set.
func (o *ApiV1EmbeddedReferencedByResource) HasScopeKind() bool {
	if o != nil && !IsNil(o.ScopeKind) {
		return true
	}

	return false
}

// SetScopeKind gets a reference to the given string and assigns it to the ScopeKind field.
func (o *ApiV1EmbeddedReferencedByResource) SetScopeKind(v string) {
	o.ScopeKind = &v
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *ApiV1EmbeddedReferencedByResource) GetGroup() string {
	if o == nil || IsNil(o.Group) {
		var ret string
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiV1EmbeddedReferencedByResource) GetGroupOk() (*string, bool) {
	if o == nil || IsNil(o.Group) {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *ApiV1EmbeddedReferencedByResource) HasGroup() bool {
	if o != nil && !IsNil(o.Group) {
		return true
	}

	return false
}

// SetGroup gets a reference to the given string and assigns it to the Group field.
func (o *ApiV1EmbeddedReferencedByResource) SetGroup(v string) {
	o.Group = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *ApiV1EmbeddedReferencedByResource) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiV1EmbeddedReferencedByResource) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *ApiV1EmbeddedReferencedByResource) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *ApiV1EmbeddedReferencedByResource) SetCount(v int32) {
	o.Count = &v
}

// GetLink returns the Link field value if set, zero value otherwise.
func (o *ApiV1EmbeddedReferencedByResource) GetLink() string {
	if o == nil || IsNil(o.Link) {
		var ret string
		return ret
	}
	return *o.Link
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiV1EmbeddedReferencedByResource) GetLinkOk() (*string, bool) {
	if o == nil || IsNil(o.Link) {
		return nil, false
	}
	return o.Link, true
}

// HasLink returns a boolean if a field has been set.
func (o *ApiV1EmbeddedReferencedByResource) HasLink() bool {
	if o != nil && !IsNil(o.Link) {
		return true
	}

	return false
}

// SetLink gets a reference to the given string and assigns it to the Link field.
func (o *ApiV1EmbeddedReferencedByResource) SetLink(v string) {
	o.Link = &v
}

func (o ApiV1EmbeddedReferencedByResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiV1EmbeddedReferencedByResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Kind) {
		toSerialize["kind"] = o.Kind
	}
	if !IsNil(o.ScopeKind) {
		toSerialize["scopeKind"] = o.ScopeKind
	}
	if !IsNil(o.Group) {
		toSerialize["group"] = o.Group
	}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.Link) {
		toSerialize["link"] = o.Link
	}
	return toSerialize, nil
}

type NullableApiV1EmbeddedReferencedByResource struct {
	value *ApiV1EmbeddedReferencedByResource
	isSet bool
}

func (v NullableApiV1EmbeddedReferencedByResource) Get() *ApiV1EmbeddedReferencedByResource {
	return v.value
}

func (v *NullableApiV1EmbeddedReferencedByResource) Set(val *ApiV1EmbeddedReferencedByResource) {
	v.value = val
	v.isSet = true
}

func (v NullableApiV1EmbeddedReferencedByResource) IsSet() bool {
	return v.isSet
}

func (v *NullableApiV1EmbeddedReferencedByResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiV1EmbeddedReferencedByResource(val *ApiV1EmbeddedReferencedByResource) *NullableApiV1EmbeddedReferencedByResource {
	return &NullableApiV1EmbeddedReferencedByResource{value: val, isSet: true}
}

func (v NullableApiV1EmbeddedReferencedByResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiV1EmbeddedReferencedByResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

