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

// checks if the ManagementV1alpha1AssetMappingTemplateSpec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManagementV1alpha1AssetMappingTemplateSpec{}

// ManagementV1alpha1AssetMappingTemplateSpec struct for ManagementV1alpha1AssetMappingTemplateSpec
type ManagementV1alpha1AssetMappingTemplateSpec struct {
	// A list of filters for the API Service resource on which the template applies.
	Filters []ManagementV1alpha1AssetMappingTemplateSpecFiltersInner `json:"filters,omitempty"`
}

// NewManagementV1alpha1AssetMappingTemplateSpec instantiates a new ManagementV1alpha1AssetMappingTemplateSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagementV1alpha1AssetMappingTemplateSpec() *ManagementV1alpha1AssetMappingTemplateSpec {
	this := ManagementV1alpha1AssetMappingTemplateSpec{}
	return &this
}

// NewManagementV1alpha1AssetMappingTemplateSpecWithDefaults instantiates a new ManagementV1alpha1AssetMappingTemplateSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagementV1alpha1AssetMappingTemplateSpecWithDefaults() *ManagementV1alpha1AssetMappingTemplateSpec {
	this := ManagementV1alpha1AssetMappingTemplateSpec{}
	return &this
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *ManagementV1alpha1AssetMappingTemplateSpec) GetFilters() []ManagementV1alpha1AssetMappingTemplateSpecFiltersInner {
	if o == nil || IsNil(o.Filters) {
		var ret []ManagementV1alpha1AssetMappingTemplateSpecFiltersInner
		return ret
	}
	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1AssetMappingTemplateSpec) GetFiltersOk() ([]ManagementV1alpha1AssetMappingTemplateSpecFiltersInner, bool) {
	if o == nil || IsNil(o.Filters) {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *ManagementV1alpha1AssetMappingTemplateSpec) HasFilters() bool {
	if o != nil && !IsNil(o.Filters) {
		return true
	}

	return false
}

// SetFilters gets a reference to the given []ManagementV1alpha1AssetMappingTemplateSpecFiltersInner and assigns it to the Filters field.
func (o *ManagementV1alpha1AssetMappingTemplateSpec) SetFilters(v []ManagementV1alpha1AssetMappingTemplateSpecFiltersInner) {
	o.Filters = v
}

func (o ManagementV1alpha1AssetMappingTemplateSpec) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManagementV1alpha1AssetMappingTemplateSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Filters) {
		toSerialize["filters"] = o.Filters
	}
	return toSerialize, nil
}

type NullableManagementV1alpha1AssetMappingTemplateSpec struct {
	value *ManagementV1alpha1AssetMappingTemplateSpec
	isSet bool
}

func (v NullableManagementV1alpha1AssetMappingTemplateSpec) Get() *ManagementV1alpha1AssetMappingTemplateSpec {
	return v.value
}

func (v *NullableManagementV1alpha1AssetMappingTemplateSpec) Set(val *ManagementV1alpha1AssetMappingTemplateSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableManagementV1alpha1AssetMappingTemplateSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableManagementV1alpha1AssetMappingTemplateSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagementV1alpha1AssetMappingTemplateSpec(val *ManagementV1alpha1AssetMappingTemplateSpec) *NullableManagementV1alpha1AssetMappingTemplateSpec {
	return &NullableManagementV1alpha1AssetMappingTemplateSpec{value: val, isSet: true}
}

func (v NullableManagementV1alpha1AssetMappingTemplateSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagementV1alpha1AssetMappingTemplateSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


