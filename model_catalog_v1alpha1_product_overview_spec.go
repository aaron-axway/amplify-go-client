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

// checks if the CatalogV1alpha1ProductOverviewSpec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogV1alpha1ProductOverviewSpec{}

// CatalogV1alpha1ProductOverviewSpec struct for CatalogV1alpha1ProductOverviewSpec
type CatalogV1alpha1ProductOverviewSpec struct {
	// Defines all the documents and order for marketplace.
	Documents []CatalogV1ProductOverviewSpecDocumentsInner `json:"documents,omitempty"`
}

// NewCatalogV1alpha1ProductOverviewSpec instantiates a new CatalogV1alpha1ProductOverviewSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogV1alpha1ProductOverviewSpec() *CatalogV1alpha1ProductOverviewSpec {
	this := CatalogV1alpha1ProductOverviewSpec{}
	return &this
}

// NewCatalogV1alpha1ProductOverviewSpecWithDefaults instantiates a new CatalogV1alpha1ProductOverviewSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogV1alpha1ProductOverviewSpecWithDefaults() *CatalogV1alpha1ProductOverviewSpec {
	this := CatalogV1alpha1ProductOverviewSpec{}
	return &this
}

// GetDocuments returns the Documents field value if set, zero value otherwise.
func (o *CatalogV1alpha1ProductOverviewSpec) GetDocuments() []CatalogV1ProductOverviewSpecDocumentsInner {
	if o == nil || IsNil(o.Documents) {
		var ret []CatalogV1ProductOverviewSpecDocumentsInner
		return ret
	}
	return o.Documents
}

// GetDocumentsOk returns a tuple with the Documents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1ProductOverviewSpec) GetDocumentsOk() ([]CatalogV1ProductOverviewSpecDocumentsInner, bool) {
	if o == nil || IsNil(o.Documents) {
		return nil, false
	}
	return o.Documents, true
}

// HasDocuments returns a boolean if a field has been set.
func (o *CatalogV1alpha1ProductOverviewSpec) HasDocuments() bool {
	if o != nil && !IsNil(o.Documents) {
		return true
	}

	return false
}

// SetDocuments gets a reference to the given []CatalogV1ProductOverviewSpecDocumentsInner and assigns it to the Documents field.
func (o *CatalogV1alpha1ProductOverviewSpec) SetDocuments(v []CatalogV1ProductOverviewSpecDocumentsInner) {
	o.Documents = v
}

func (o CatalogV1alpha1ProductOverviewSpec) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogV1alpha1ProductOverviewSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Documents) {
		toSerialize["documents"] = o.Documents
	}
	return toSerialize, nil
}

type NullableCatalogV1alpha1ProductOverviewSpec struct {
	value *CatalogV1alpha1ProductOverviewSpec
	isSet bool
}

func (v NullableCatalogV1alpha1ProductOverviewSpec) Get() *CatalogV1alpha1ProductOverviewSpec {
	return v.value
}

func (v *NullableCatalogV1alpha1ProductOverviewSpec) Set(val *CatalogV1alpha1ProductOverviewSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogV1alpha1ProductOverviewSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogV1alpha1ProductOverviewSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogV1alpha1ProductOverviewSpec(val *CatalogV1alpha1ProductOverviewSpec) *NullableCatalogV1alpha1ProductOverviewSpec {
	return &NullableCatalogV1alpha1ProductOverviewSpec{value: val, isSet: true}
}

func (v NullableCatalogV1alpha1ProductOverviewSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogV1alpha1ProductOverviewSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


