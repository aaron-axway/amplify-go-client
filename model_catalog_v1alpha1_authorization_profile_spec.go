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

// checks if the CatalogV1alpha1AuthorizationProfileSpec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogV1alpha1AuthorizationProfileSpec{}

// CatalogV1alpha1AuthorizationProfileSpec struct for CatalogV1alpha1AuthorizationProfileSpec
type CatalogV1alpha1AuthorizationProfileSpec struct {
	// Description of AuthorizationProfile that can be used to access Asset.
	Description *string `json:"description,omitempty"`
}

// NewCatalogV1alpha1AuthorizationProfileSpec instantiates a new CatalogV1alpha1AuthorizationProfileSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogV1alpha1AuthorizationProfileSpec() *CatalogV1alpha1AuthorizationProfileSpec {
	this := CatalogV1alpha1AuthorizationProfileSpec{}
	return &this
}

// NewCatalogV1alpha1AuthorizationProfileSpecWithDefaults instantiates a new CatalogV1alpha1AuthorizationProfileSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogV1alpha1AuthorizationProfileSpecWithDefaults() *CatalogV1alpha1AuthorizationProfileSpec {
	this := CatalogV1alpha1AuthorizationProfileSpec{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CatalogV1alpha1AuthorizationProfileSpec) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1AuthorizationProfileSpec) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CatalogV1alpha1AuthorizationProfileSpec) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CatalogV1alpha1AuthorizationProfileSpec) SetDescription(v string) {
	o.Description = &v
}

func (o CatalogV1alpha1AuthorizationProfileSpec) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogV1alpha1AuthorizationProfileSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

type NullableCatalogV1alpha1AuthorizationProfileSpec struct {
	value *CatalogV1alpha1AuthorizationProfileSpec
	isSet bool
}

func (v NullableCatalogV1alpha1AuthorizationProfileSpec) Get() *CatalogV1alpha1AuthorizationProfileSpec {
	return v.value
}

func (v *NullableCatalogV1alpha1AuthorizationProfileSpec) Set(val *CatalogV1alpha1AuthorizationProfileSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogV1alpha1AuthorizationProfileSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogV1alpha1AuthorizationProfileSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogV1alpha1AuthorizationProfileSpec(val *CatalogV1alpha1AuthorizationProfileSpec) *NullableCatalogV1alpha1AuthorizationProfileSpec {
	return &NullableCatalogV1alpha1AuthorizationProfileSpec{value: val, isSet: true}
}

func (v NullableCatalogV1alpha1AuthorizationProfileSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogV1alpha1AuthorizationProfileSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

