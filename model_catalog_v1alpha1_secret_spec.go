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

// checks if the CatalogV1alpha1SecretSpec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogV1alpha1SecretSpec{}

// CatalogV1alpha1SecretSpec struct for CatalogV1alpha1SecretSpec
type CatalogV1alpha1SecretSpec struct {
	// Key value pairs. The value will be stored encrypted.
	Data *map[string]string `json:"data,omitempty"`
}

// NewCatalogV1alpha1SecretSpec instantiates a new CatalogV1alpha1SecretSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogV1alpha1SecretSpec() *CatalogV1alpha1SecretSpec {
	this := CatalogV1alpha1SecretSpec{}
	return &this
}

// NewCatalogV1alpha1SecretSpecWithDefaults instantiates a new CatalogV1alpha1SecretSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogV1alpha1SecretSpecWithDefaults() *CatalogV1alpha1SecretSpec {
	this := CatalogV1alpha1SecretSpec{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CatalogV1alpha1SecretSpec) GetData() map[string]string {
	if o == nil || IsNil(o.Data) {
		var ret map[string]string
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1SecretSpec) GetDataOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CatalogV1alpha1SecretSpec) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given map[string]string and assigns it to the Data field.
func (o *CatalogV1alpha1SecretSpec) SetData(v map[string]string) {
	o.Data = &v
}

func (o CatalogV1alpha1SecretSpec) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogV1alpha1SecretSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableCatalogV1alpha1SecretSpec struct {
	value *CatalogV1alpha1SecretSpec
	isSet bool
}

func (v NullableCatalogV1alpha1SecretSpec) Get() *CatalogV1alpha1SecretSpec {
	return v.value
}

func (v *NullableCatalogV1alpha1SecretSpec) Set(val *CatalogV1alpha1SecretSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogV1alpha1SecretSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogV1alpha1SecretSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogV1alpha1SecretSpec(val *CatalogV1alpha1SecretSpec) *NullableCatalogV1alpha1SecretSpec {
	return &NullableCatalogV1alpha1SecretSpec{value: val, isSet: true}
}

func (v NullableCatalogV1alpha1SecretSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogV1alpha1SecretSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


