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

// checks if the ManagementV1alpha1SpecDiscoverySpec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManagementV1alpha1SpecDiscoverySpec{}

// ManagementV1alpha1SpecDiscoverySpec struct for ManagementV1alpha1SpecDiscoverySpec
type ManagementV1alpha1SpecDiscoverySpec struct {
	NamespaceFilter *ManagementV1alpha1SpecDiscoverySpecNamespaceFilter `json:"namespaceFilter,omitempty"`
	ResourceFilter *ManagementV1alpha1SpecDiscoverySpecResourceFilter `json:"resourceFilter,omitempty"`
	Targets *ManagementV1alpha1SpecDiscoverySpecTargets `json:"targets,omitempty"`
}

// NewManagementV1alpha1SpecDiscoverySpec instantiates a new ManagementV1alpha1SpecDiscoverySpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagementV1alpha1SpecDiscoverySpec() *ManagementV1alpha1SpecDiscoverySpec {
	this := ManagementV1alpha1SpecDiscoverySpec{}
	return &this
}

// NewManagementV1alpha1SpecDiscoverySpecWithDefaults instantiates a new ManagementV1alpha1SpecDiscoverySpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagementV1alpha1SpecDiscoverySpecWithDefaults() *ManagementV1alpha1SpecDiscoverySpec {
	this := ManagementV1alpha1SpecDiscoverySpec{}
	return &this
}

// GetNamespaceFilter returns the NamespaceFilter field value if set, zero value otherwise.
func (o *ManagementV1alpha1SpecDiscoverySpec) GetNamespaceFilter() ManagementV1alpha1SpecDiscoverySpecNamespaceFilter {
	if o == nil || IsNil(o.NamespaceFilter) {
		var ret ManagementV1alpha1SpecDiscoverySpecNamespaceFilter
		return ret
	}
	return *o.NamespaceFilter
}

// GetNamespaceFilterOk returns a tuple with the NamespaceFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1SpecDiscoverySpec) GetNamespaceFilterOk() (*ManagementV1alpha1SpecDiscoverySpecNamespaceFilter, bool) {
	if o == nil || IsNil(o.NamespaceFilter) {
		return nil, false
	}
	return o.NamespaceFilter, true
}

// HasNamespaceFilter returns a boolean if a field has been set.
func (o *ManagementV1alpha1SpecDiscoverySpec) HasNamespaceFilter() bool {
	if o != nil && !IsNil(o.NamespaceFilter) {
		return true
	}

	return false
}

// SetNamespaceFilter gets a reference to the given ManagementV1alpha1SpecDiscoverySpecNamespaceFilter and assigns it to the NamespaceFilter field.
func (o *ManagementV1alpha1SpecDiscoverySpec) SetNamespaceFilter(v ManagementV1alpha1SpecDiscoverySpecNamespaceFilter) {
	o.NamespaceFilter = &v
}

// GetResourceFilter returns the ResourceFilter field value if set, zero value otherwise.
func (o *ManagementV1alpha1SpecDiscoverySpec) GetResourceFilter() ManagementV1alpha1SpecDiscoverySpecResourceFilter {
	if o == nil || IsNil(o.ResourceFilter) {
		var ret ManagementV1alpha1SpecDiscoverySpecResourceFilter
		return ret
	}
	return *o.ResourceFilter
}

// GetResourceFilterOk returns a tuple with the ResourceFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1SpecDiscoverySpec) GetResourceFilterOk() (*ManagementV1alpha1SpecDiscoverySpecResourceFilter, bool) {
	if o == nil || IsNil(o.ResourceFilter) {
		return nil, false
	}
	return o.ResourceFilter, true
}

// HasResourceFilter returns a boolean if a field has been set.
func (o *ManagementV1alpha1SpecDiscoverySpec) HasResourceFilter() bool {
	if o != nil && !IsNil(o.ResourceFilter) {
		return true
	}

	return false
}

// SetResourceFilter gets a reference to the given ManagementV1alpha1SpecDiscoverySpecResourceFilter and assigns it to the ResourceFilter field.
func (o *ManagementV1alpha1SpecDiscoverySpec) SetResourceFilter(v ManagementV1alpha1SpecDiscoverySpecResourceFilter) {
	o.ResourceFilter = &v
}

// GetTargets returns the Targets field value if set, zero value otherwise.
func (o *ManagementV1alpha1SpecDiscoverySpec) GetTargets() ManagementV1alpha1SpecDiscoverySpecTargets {
	if o == nil || IsNil(o.Targets) {
		var ret ManagementV1alpha1SpecDiscoverySpecTargets
		return ret
	}
	return *o.Targets
}

// GetTargetsOk returns a tuple with the Targets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1SpecDiscoverySpec) GetTargetsOk() (*ManagementV1alpha1SpecDiscoverySpecTargets, bool) {
	if o == nil || IsNil(o.Targets) {
		return nil, false
	}
	return o.Targets, true
}

// HasTargets returns a boolean if a field has been set.
func (o *ManagementV1alpha1SpecDiscoverySpec) HasTargets() bool {
	if o != nil && !IsNil(o.Targets) {
		return true
	}

	return false
}

// SetTargets gets a reference to the given ManagementV1alpha1SpecDiscoverySpecTargets and assigns it to the Targets field.
func (o *ManagementV1alpha1SpecDiscoverySpec) SetTargets(v ManagementV1alpha1SpecDiscoverySpecTargets) {
	o.Targets = &v
}

func (o ManagementV1alpha1SpecDiscoverySpec) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManagementV1alpha1SpecDiscoverySpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NamespaceFilter) {
		toSerialize["namespaceFilter"] = o.NamespaceFilter
	}
	if !IsNil(o.ResourceFilter) {
		toSerialize["resourceFilter"] = o.ResourceFilter
	}
	if !IsNil(o.Targets) {
		toSerialize["targets"] = o.Targets
	}
	return toSerialize, nil
}

type NullableManagementV1alpha1SpecDiscoverySpec struct {
	value *ManagementV1alpha1SpecDiscoverySpec
	isSet bool
}

func (v NullableManagementV1alpha1SpecDiscoverySpec) Get() *ManagementV1alpha1SpecDiscoverySpec {
	return v.value
}

func (v *NullableManagementV1alpha1SpecDiscoverySpec) Set(val *ManagementV1alpha1SpecDiscoverySpec) {
	v.value = val
	v.isSet = true
}

func (v NullableManagementV1alpha1SpecDiscoverySpec) IsSet() bool {
	return v.isSet
}

func (v *NullableManagementV1alpha1SpecDiscoverySpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagementV1alpha1SpecDiscoverySpec(val *ManagementV1alpha1SpecDiscoverySpec) *NullableManagementV1alpha1SpecDiscoverySpec {
	return &NullableManagementV1alpha1SpecDiscoverySpec{value: val, isSet: true}
}

func (v NullableManagementV1alpha1SpecDiscoverySpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagementV1alpha1SpecDiscoverySpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


