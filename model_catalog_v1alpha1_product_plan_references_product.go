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

// checks if the CatalogV1alpha1ProductPlanReferencesProduct type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogV1alpha1ProductPlanReferencesProduct{}

// CatalogV1alpha1ProductPlanReferencesProduct struct for CatalogV1alpha1ProductPlanReferencesProduct
type CatalogV1alpha1ProductPlanReferencesProduct struct {
	// The latest active Product Release that corresponds to the Product referenced in the Plan.
	Release *string `json:"release,omitempty"`
}

// NewCatalogV1alpha1ProductPlanReferencesProduct instantiates a new CatalogV1alpha1ProductPlanReferencesProduct object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogV1alpha1ProductPlanReferencesProduct() *CatalogV1alpha1ProductPlanReferencesProduct {
	this := CatalogV1alpha1ProductPlanReferencesProduct{}
	return &this
}

// NewCatalogV1alpha1ProductPlanReferencesProductWithDefaults instantiates a new CatalogV1alpha1ProductPlanReferencesProduct object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogV1alpha1ProductPlanReferencesProductWithDefaults() *CatalogV1alpha1ProductPlanReferencesProduct {
	this := CatalogV1alpha1ProductPlanReferencesProduct{}
	return &this
}

// GetRelease returns the Release field value if set, zero value otherwise.
func (o *CatalogV1alpha1ProductPlanReferencesProduct) GetRelease() string {
	if o == nil || IsNil(o.Release) {
		var ret string
		return ret
	}
	return *o.Release
}

// GetReleaseOk returns a tuple with the Release field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1ProductPlanReferencesProduct) GetReleaseOk() (*string, bool) {
	if o == nil || IsNil(o.Release) {
		return nil, false
	}
	return o.Release, true
}

// HasRelease returns a boolean if a field has been set.
func (o *CatalogV1alpha1ProductPlanReferencesProduct) HasRelease() bool {
	if o != nil && !IsNil(o.Release) {
		return true
	}

	return false
}

// SetRelease gets a reference to the given string and assigns it to the Release field.
func (o *CatalogV1alpha1ProductPlanReferencesProduct) SetRelease(v string) {
	o.Release = &v
}

func (o CatalogV1alpha1ProductPlanReferencesProduct) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogV1alpha1ProductPlanReferencesProduct) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Release) {
		toSerialize["release"] = o.Release
	}
	return toSerialize, nil
}

type NullableCatalogV1alpha1ProductPlanReferencesProduct struct {
	value *CatalogV1alpha1ProductPlanReferencesProduct
	isSet bool
}

func (v NullableCatalogV1alpha1ProductPlanReferencesProduct) Get() *CatalogV1alpha1ProductPlanReferencesProduct {
	return v.value
}

func (v *NullableCatalogV1alpha1ProductPlanReferencesProduct) Set(val *CatalogV1alpha1ProductPlanReferencesProduct) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogV1alpha1ProductPlanReferencesProduct) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogV1alpha1ProductPlanReferencesProduct) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogV1alpha1ProductPlanReferencesProduct(val *CatalogV1alpha1ProductPlanReferencesProduct) *NullableCatalogV1alpha1ProductPlanReferencesProduct {
	return &NullableCatalogV1alpha1ProductPlanReferencesProduct{value: val, isSet: true}
}

func (v NullableCatalogV1alpha1ProductPlanReferencesProduct) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogV1alpha1ProductPlanReferencesProduct) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


