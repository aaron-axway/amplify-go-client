/*
AMPLIFY Central API v0.347

APIs to manage AMPLIFY Central configuration resources.

API version: 0.347.0
Contact: support@axway.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package amplify

import (
	"encoding/json"
)

// checks if the ManagementV1alpha1AssetMappingTemplateSpecFiltersInnerApiServiceInnerApiServiceRevisionApiServiceInstance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManagementV1alpha1AssetMappingTemplateSpecFiltersInnerApiServiceInnerApiServiceRevisionApiServiceInstance{}

// ManagementV1alpha1AssetMappingTemplateSpecFiltersInnerApiServiceInnerApiServiceRevisionApiServiceInstance struct for ManagementV1alpha1AssetMappingTemplateSpecFiltersInnerApiServiceInnerApiServiceRevisionApiServiceInstance
type ManagementV1alpha1AssetMappingTemplateSpecFiltersInnerApiServiceInnerApiServiceRevisionApiServiceInstance struct {
	// Attributes used to filter the APIServiceInstances for the API Service on which the template applies.
	Attributes *map[string]string `json:"attributes,omitempty"`
}

// NewManagementV1alpha1AssetMappingTemplateSpecFiltersInnerApiServiceInnerApiServiceRevisionApiServiceInstance instantiates a new ManagementV1alpha1AssetMappingTemplateSpecFiltersInnerApiServiceInnerApiServiceRevisionApiServiceInstance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagementV1alpha1AssetMappingTemplateSpecFiltersInnerApiServiceInnerApiServiceRevisionApiServiceInstance() *ManagementV1alpha1AssetMappingTemplateSpecFiltersInnerApiServiceInnerApiServiceRevisionApiServiceInstance {
	this := ManagementV1alpha1AssetMappingTemplateSpecFiltersInnerApiServiceInnerApiServiceRevisionApiServiceInstance{}
	return &this
}

// NewManagementV1alpha1AssetMappingTemplateSpecFiltersInnerApiServiceInnerApiServiceRevisionApiServiceInstanceWithDefaults instantiates a new ManagementV1alpha1AssetMappingTemplateSpecFiltersInnerApiServiceInnerApiServiceRevisionApiServiceInstance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagementV1alpha1AssetMappingTemplateSpecFiltersInnerApiServiceInnerApiServiceRevisionApiServiceInstanceWithDefaults() *ManagementV1alpha1AssetMappingTemplateSpecFiltersInnerApiServiceInnerApiServiceRevisionApiServiceInstance {
	this := ManagementV1alpha1AssetMappingTemplateSpecFiltersInnerApiServiceInnerApiServiceRevisionApiServiceInstance{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *ManagementV1alpha1AssetMappingTemplateSpecFiltersInnerApiServiceInnerApiServiceRevisionApiServiceInstance) GetAttributes() map[string]string {
	if o == nil || IsNil(o.Attributes) {
		var ret map[string]string
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1AssetMappingTemplateSpecFiltersInnerApiServiceInnerApiServiceRevisionApiServiceInstance) GetAttributesOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *ManagementV1alpha1AssetMappingTemplateSpecFiltersInnerApiServiceInnerApiServiceRevisionApiServiceInstance) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]string and assigns it to the Attributes field.
func (o *ManagementV1alpha1AssetMappingTemplateSpecFiltersInnerApiServiceInnerApiServiceRevisionApiServiceInstance) SetAttributes(v map[string]string) {
	o.Attributes = &v
}

func (o ManagementV1alpha1AssetMappingTemplateSpecFiltersInnerApiServiceInnerApiServiceRevisionApiServiceInstance) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManagementV1alpha1AssetMappingTemplateSpecFiltersInnerApiServiceInnerApiServiceRevisionApiServiceInstance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

type NullableManagementV1alpha1AssetMappingTemplateSpecFiltersInnerApiServiceInnerApiServiceRevisionApiServiceInstance struct {
	value *ManagementV1alpha1AssetMappingTemplateSpecFiltersInnerApiServiceInnerApiServiceRevisionApiServiceInstance
	isSet bool
}

func (v NullableManagementV1alpha1AssetMappingTemplateSpecFiltersInnerApiServiceInnerApiServiceRevisionApiServiceInstance) Get() *ManagementV1alpha1AssetMappingTemplateSpecFiltersInnerApiServiceInnerApiServiceRevisionApiServiceInstance {
	return v.value
}

func (v *NullableManagementV1alpha1AssetMappingTemplateSpecFiltersInnerApiServiceInnerApiServiceRevisionApiServiceInstance) Set(val *ManagementV1alpha1AssetMappingTemplateSpecFiltersInnerApiServiceInnerApiServiceRevisionApiServiceInstance) {
	v.value = val
	v.isSet = true
}

func (v NullableManagementV1alpha1AssetMappingTemplateSpecFiltersInnerApiServiceInnerApiServiceRevisionApiServiceInstance) IsSet() bool {
	return v.isSet
}

func (v *NullableManagementV1alpha1AssetMappingTemplateSpecFiltersInnerApiServiceInnerApiServiceRevisionApiServiceInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagementV1alpha1AssetMappingTemplateSpecFiltersInnerApiServiceInnerApiServiceRevisionApiServiceInstance(val *ManagementV1alpha1AssetMappingTemplateSpecFiltersInnerApiServiceInnerApiServiceRevisionApiServiceInstance) *NullableManagementV1alpha1AssetMappingTemplateSpecFiltersInnerApiServiceInnerApiServiceRevisionApiServiceInstance {
	return &NullableManagementV1alpha1AssetMappingTemplateSpecFiltersInnerApiServiceInnerApiServiceRevisionApiServiceInstance{value: val, isSet: true}
}

func (v NullableManagementV1alpha1AssetMappingTemplateSpecFiltersInnerApiServiceInnerApiServiceRevisionApiServiceInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagementV1alpha1AssetMappingTemplateSpecFiltersInnerApiServiceInnerApiServiceRevisionApiServiceInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


