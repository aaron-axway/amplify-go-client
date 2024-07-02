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

// checks if the ManagementV1alpha1SpecDiscoverySpecTargetsFromAnnotationInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManagementV1alpha1SpecDiscoverySpecTargetsFromAnnotationInner{}

// ManagementV1alpha1SpecDiscoverySpecTargetsFromAnnotationInner struct for ManagementV1alpha1SpecDiscoverySpecTargetsFromAnnotationInner
type ManagementV1alpha1SpecDiscoverySpecTargetsFromAnnotationInner struct {
	// annotation template. golang template that's fed the following object for each declared port of the pod: {\"name\", \"number\"}
	Template *string `json:"template,omitempty"`
}

// NewManagementV1alpha1SpecDiscoverySpecTargetsFromAnnotationInner instantiates a new ManagementV1alpha1SpecDiscoverySpecTargetsFromAnnotationInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagementV1alpha1SpecDiscoverySpecTargetsFromAnnotationInner() *ManagementV1alpha1SpecDiscoverySpecTargetsFromAnnotationInner {
	this := ManagementV1alpha1SpecDiscoverySpecTargetsFromAnnotationInner{}
	return &this
}

// NewManagementV1alpha1SpecDiscoverySpecTargetsFromAnnotationInnerWithDefaults instantiates a new ManagementV1alpha1SpecDiscoverySpecTargetsFromAnnotationInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagementV1alpha1SpecDiscoverySpecTargetsFromAnnotationInnerWithDefaults() *ManagementV1alpha1SpecDiscoverySpecTargetsFromAnnotationInner {
	this := ManagementV1alpha1SpecDiscoverySpecTargetsFromAnnotationInner{}
	return &this
}

// GetTemplate returns the Template field value if set, zero value otherwise.
func (o *ManagementV1alpha1SpecDiscoverySpecTargetsFromAnnotationInner) GetTemplate() string {
	if o == nil || IsNil(o.Template) {
		var ret string
		return ret
	}
	return *o.Template
}

// GetTemplateOk returns a tuple with the Template field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1SpecDiscoverySpecTargetsFromAnnotationInner) GetTemplateOk() (*string, bool) {
	if o == nil || IsNil(o.Template) {
		return nil, false
	}
	return o.Template, true
}

// HasTemplate returns a boolean if a field has been set.
func (o *ManagementV1alpha1SpecDiscoverySpecTargetsFromAnnotationInner) HasTemplate() bool {
	if o != nil && !IsNil(o.Template) {
		return true
	}

	return false
}

// SetTemplate gets a reference to the given string and assigns it to the Template field.
func (o *ManagementV1alpha1SpecDiscoverySpecTargetsFromAnnotationInner) SetTemplate(v string) {
	o.Template = &v
}

func (o ManagementV1alpha1SpecDiscoverySpecTargetsFromAnnotationInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManagementV1alpha1SpecDiscoverySpecTargetsFromAnnotationInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Template) {
		toSerialize["template"] = o.Template
	}
	return toSerialize, nil
}

type NullableManagementV1alpha1SpecDiscoverySpecTargetsFromAnnotationInner struct {
	value *ManagementV1alpha1SpecDiscoverySpecTargetsFromAnnotationInner
	isSet bool
}

func (v NullableManagementV1alpha1SpecDiscoverySpecTargetsFromAnnotationInner) Get() *ManagementV1alpha1SpecDiscoverySpecTargetsFromAnnotationInner {
	return v.value
}

func (v *NullableManagementV1alpha1SpecDiscoverySpecTargetsFromAnnotationInner) Set(val *ManagementV1alpha1SpecDiscoverySpecTargetsFromAnnotationInner) {
	v.value = val
	v.isSet = true
}

func (v NullableManagementV1alpha1SpecDiscoverySpecTargetsFromAnnotationInner) IsSet() bool {
	return v.isSet
}

func (v *NullableManagementV1alpha1SpecDiscoverySpecTargetsFromAnnotationInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagementV1alpha1SpecDiscoverySpecTargetsFromAnnotationInner(val *ManagementV1alpha1SpecDiscoverySpecTargetsFromAnnotationInner) *NullableManagementV1alpha1SpecDiscoverySpecTargetsFromAnnotationInner {
	return &NullableManagementV1alpha1SpecDiscoverySpecTargetsFromAnnotationInner{value: val, isSet: true}
}

func (v NullableManagementV1alpha1SpecDiscoverySpecTargetsFromAnnotationInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagementV1alpha1SpecDiscoverySpecTargetsFromAnnotationInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


