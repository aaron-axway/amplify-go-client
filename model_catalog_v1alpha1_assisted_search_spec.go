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

// checks if the CatalogV1alpha1AssistedSearchSpec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogV1alpha1AssistedSearchSpec{}

// CatalogV1alpha1AssistedSearchSpec struct for CatalogV1alpha1AssistedSearchSpec
type CatalogV1alpha1AssistedSearchSpec struct {
	Integration *CatalogV1alpha1AssistedSearchSpecIntegration `json:"integration,omitempty"`
}

// NewCatalogV1alpha1AssistedSearchSpec instantiates a new CatalogV1alpha1AssistedSearchSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogV1alpha1AssistedSearchSpec() *CatalogV1alpha1AssistedSearchSpec {
	this := CatalogV1alpha1AssistedSearchSpec{}
	return &this
}

// NewCatalogV1alpha1AssistedSearchSpecWithDefaults instantiates a new CatalogV1alpha1AssistedSearchSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogV1alpha1AssistedSearchSpecWithDefaults() *CatalogV1alpha1AssistedSearchSpec {
	this := CatalogV1alpha1AssistedSearchSpec{}
	return &this
}

// GetIntegration returns the Integration field value if set, zero value otherwise.
func (o *CatalogV1alpha1AssistedSearchSpec) GetIntegration() CatalogV1alpha1AssistedSearchSpecIntegration {
	if o == nil || IsNil(o.Integration) {
		var ret CatalogV1alpha1AssistedSearchSpecIntegration
		return ret
	}
	return *o.Integration
}

// GetIntegrationOk returns a tuple with the Integration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1AssistedSearchSpec) GetIntegrationOk() (*CatalogV1alpha1AssistedSearchSpecIntegration, bool) {
	if o == nil || IsNil(o.Integration) {
		return nil, false
	}
	return o.Integration, true
}

// HasIntegration returns a boolean if a field has been set.
func (o *CatalogV1alpha1AssistedSearchSpec) HasIntegration() bool {
	if o != nil && !IsNil(o.Integration) {
		return true
	}

	return false
}

// SetIntegration gets a reference to the given CatalogV1alpha1AssistedSearchSpecIntegration and assigns it to the Integration field.
func (o *CatalogV1alpha1AssistedSearchSpec) SetIntegration(v CatalogV1alpha1AssistedSearchSpecIntegration) {
	o.Integration = &v
}

func (o CatalogV1alpha1AssistedSearchSpec) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogV1alpha1AssistedSearchSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Integration) {
		toSerialize["integration"] = o.Integration
	}
	return toSerialize, nil
}

type NullableCatalogV1alpha1AssistedSearchSpec struct {
	value *CatalogV1alpha1AssistedSearchSpec
	isSet bool
}

func (v NullableCatalogV1alpha1AssistedSearchSpec) Get() *CatalogV1alpha1AssistedSearchSpec {
	return v.value
}

func (v *NullableCatalogV1alpha1AssistedSearchSpec) Set(val *CatalogV1alpha1AssistedSearchSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogV1alpha1AssistedSearchSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogV1alpha1AssistedSearchSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogV1alpha1AssistedSearchSpec(val *CatalogV1alpha1AssistedSearchSpec) *NullableCatalogV1alpha1AssistedSearchSpec {
	return &NullableCatalogV1alpha1AssistedSearchSpec{value: val, isSet: true}
}

func (v NullableCatalogV1alpha1AssistedSearchSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogV1alpha1AssistedSearchSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


