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

// checks if the ManagementV1alpha1SpecDiscoverySpecNamespaceFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManagementV1alpha1SpecDiscoverySpecNamespaceFilter{}

// ManagementV1alpha1SpecDiscoverySpecNamespaceFilter a list of namespace names to follow. If not set, follows all namespaces.
type ManagementV1alpha1SpecDiscoverySpecNamespaceFilter struct {
	Names []string `json:"names,omitempty"`
}

// NewManagementV1alpha1SpecDiscoverySpecNamespaceFilter instantiates a new ManagementV1alpha1SpecDiscoverySpecNamespaceFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagementV1alpha1SpecDiscoverySpecNamespaceFilter() *ManagementV1alpha1SpecDiscoverySpecNamespaceFilter {
	this := ManagementV1alpha1SpecDiscoverySpecNamespaceFilter{}
	return &this
}

// NewManagementV1alpha1SpecDiscoverySpecNamespaceFilterWithDefaults instantiates a new ManagementV1alpha1SpecDiscoverySpecNamespaceFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagementV1alpha1SpecDiscoverySpecNamespaceFilterWithDefaults() *ManagementV1alpha1SpecDiscoverySpecNamespaceFilter {
	this := ManagementV1alpha1SpecDiscoverySpecNamespaceFilter{}
	return &this
}

// GetNames returns the Names field value if set, zero value otherwise.
func (o *ManagementV1alpha1SpecDiscoverySpecNamespaceFilter) GetNames() []string {
	if o == nil || IsNil(o.Names) {
		var ret []string
		return ret
	}
	return o.Names
}

// GetNamesOk returns a tuple with the Names field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1SpecDiscoverySpecNamespaceFilter) GetNamesOk() ([]string, bool) {
	if o == nil || IsNil(o.Names) {
		return nil, false
	}
	return o.Names, true
}

// HasNames returns a boolean if a field has been set.
func (o *ManagementV1alpha1SpecDiscoverySpecNamespaceFilter) HasNames() bool {
	if o != nil && !IsNil(o.Names) {
		return true
	}

	return false
}

// SetNames gets a reference to the given []string and assigns it to the Names field.
func (o *ManagementV1alpha1SpecDiscoverySpecNamespaceFilter) SetNames(v []string) {
	o.Names = v
}

func (o ManagementV1alpha1SpecDiscoverySpecNamespaceFilter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManagementV1alpha1SpecDiscoverySpecNamespaceFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Names) {
		toSerialize["names"] = o.Names
	}
	return toSerialize, nil
}

type NullableManagementV1alpha1SpecDiscoverySpecNamespaceFilter struct {
	value *ManagementV1alpha1SpecDiscoverySpecNamespaceFilter
	isSet bool
}

func (v NullableManagementV1alpha1SpecDiscoverySpecNamespaceFilter) Get() *ManagementV1alpha1SpecDiscoverySpecNamespaceFilter {
	return v.value
}

func (v *NullableManagementV1alpha1SpecDiscoverySpecNamespaceFilter) Set(val *ManagementV1alpha1SpecDiscoverySpecNamespaceFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableManagementV1alpha1SpecDiscoverySpecNamespaceFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableManagementV1alpha1SpecDiscoverySpecNamespaceFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagementV1alpha1SpecDiscoverySpecNamespaceFilter(val *ManagementV1alpha1SpecDiscoverySpecNamespaceFilter) *NullableManagementV1alpha1SpecDiscoverySpecNamespaceFilter {
	return &NullableManagementV1alpha1SpecDiscoverySpecNamespaceFilter{value: val, isSet: true}
}

func (v NullableManagementV1alpha1SpecDiscoverySpecNamespaceFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagementV1alpha1SpecDiscoverySpecNamespaceFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

