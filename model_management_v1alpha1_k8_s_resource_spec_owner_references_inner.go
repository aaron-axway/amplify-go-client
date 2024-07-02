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

// checks if the ManagementV1alpha1K8SResourceSpecOwnerReferencesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManagementV1alpha1K8SResourceSpecOwnerReferencesInner{}

// ManagementV1alpha1K8SResourceSpecOwnerReferencesInner struct for ManagementV1alpha1K8SResourceSpecOwnerReferencesInner
type ManagementV1alpha1K8SResourceSpecOwnerReferencesInner struct {
	Kind *string `json:"kind,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewManagementV1alpha1K8SResourceSpecOwnerReferencesInner instantiates a new ManagementV1alpha1K8SResourceSpecOwnerReferencesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagementV1alpha1K8SResourceSpecOwnerReferencesInner() *ManagementV1alpha1K8SResourceSpecOwnerReferencesInner {
	this := ManagementV1alpha1K8SResourceSpecOwnerReferencesInner{}
	return &this
}

// NewManagementV1alpha1K8SResourceSpecOwnerReferencesInnerWithDefaults instantiates a new ManagementV1alpha1K8SResourceSpecOwnerReferencesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagementV1alpha1K8SResourceSpecOwnerReferencesInnerWithDefaults() *ManagementV1alpha1K8SResourceSpecOwnerReferencesInner {
	this := ManagementV1alpha1K8SResourceSpecOwnerReferencesInner{}
	return &this
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *ManagementV1alpha1K8SResourceSpecOwnerReferencesInner) GetKind() string {
	if o == nil || IsNil(o.Kind) {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1K8SResourceSpecOwnerReferencesInner) GetKindOk() (*string, bool) {
	if o == nil || IsNil(o.Kind) {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *ManagementV1alpha1K8SResourceSpecOwnerReferencesInner) HasKind() bool {
	if o != nil && !IsNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *ManagementV1alpha1K8SResourceSpecOwnerReferencesInner) SetKind(v string) {
	o.Kind = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ManagementV1alpha1K8SResourceSpecOwnerReferencesInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1K8SResourceSpecOwnerReferencesInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ManagementV1alpha1K8SResourceSpecOwnerReferencesInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ManagementV1alpha1K8SResourceSpecOwnerReferencesInner) SetName(v string) {
	o.Name = &v
}

func (o ManagementV1alpha1K8SResourceSpecOwnerReferencesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManagementV1alpha1K8SResourceSpecOwnerReferencesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Kind) {
		toSerialize["kind"] = o.Kind
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableManagementV1alpha1K8SResourceSpecOwnerReferencesInner struct {
	value *ManagementV1alpha1K8SResourceSpecOwnerReferencesInner
	isSet bool
}

func (v NullableManagementV1alpha1K8SResourceSpecOwnerReferencesInner) Get() *ManagementV1alpha1K8SResourceSpecOwnerReferencesInner {
	return v.value
}

func (v *NullableManagementV1alpha1K8SResourceSpecOwnerReferencesInner) Set(val *ManagementV1alpha1K8SResourceSpecOwnerReferencesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableManagementV1alpha1K8SResourceSpecOwnerReferencesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableManagementV1alpha1K8SResourceSpecOwnerReferencesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagementV1alpha1K8SResourceSpecOwnerReferencesInner(val *ManagementV1alpha1K8SResourceSpecOwnerReferencesInner) *NullableManagementV1alpha1K8SResourceSpecOwnerReferencesInner {
	return &NullableManagementV1alpha1K8SResourceSpecOwnerReferencesInner{value: val, isSet: true}
}

func (v NullableManagementV1alpha1K8SResourceSpecOwnerReferencesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagementV1alpha1K8SResourceSpecOwnerReferencesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


