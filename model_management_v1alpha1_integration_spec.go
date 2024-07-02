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

// checks if the ManagementV1alpha1IntegrationSpec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManagementV1alpha1IntegrationSpec{}

// ManagementV1alpha1IntegrationSpec struct for ManagementV1alpha1IntegrationSpec
type ManagementV1alpha1IntegrationSpec struct {
	Description *string `json:"description,omitempty"`
}

// NewManagementV1alpha1IntegrationSpec instantiates a new ManagementV1alpha1IntegrationSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagementV1alpha1IntegrationSpec() *ManagementV1alpha1IntegrationSpec {
	this := ManagementV1alpha1IntegrationSpec{}
	return &this
}

// NewManagementV1alpha1IntegrationSpecWithDefaults instantiates a new ManagementV1alpha1IntegrationSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagementV1alpha1IntegrationSpecWithDefaults() *ManagementV1alpha1IntegrationSpec {
	this := ManagementV1alpha1IntegrationSpec{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ManagementV1alpha1IntegrationSpec) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1IntegrationSpec) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ManagementV1alpha1IntegrationSpec) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ManagementV1alpha1IntegrationSpec) SetDescription(v string) {
	o.Description = &v
}

func (o ManagementV1alpha1IntegrationSpec) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManagementV1alpha1IntegrationSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

type NullableManagementV1alpha1IntegrationSpec struct {
	value *ManagementV1alpha1IntegrationSpec
	isSet bool
}

func (v NullableManagementV1alpha1IntegrationSpec) Get() *ManagementV1alpha1IntegrationSpec {
	return v.value
}

func (v *NullableManagementV1alpha1IntegrationSpec) Set(val *ManagementV1alpha1IntegrationSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableManagementV1alpha1IntegrationSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableManagementV1alpha1IntegrationSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagementV1alpha1IntegrationSpec(val *ManagementV1alpha1IntegrationSpec) *NullableManagementV1alpha1IntegrationSpec {
	return &NullableManagementV1alpha1IntegrationSpec{value: val, isSet: true}
}

func (v NullableManagementV1alpha1IntegrationSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagementV1alpha1IntegrationSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


