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

// checks if the CatalogV1alpha1AssetStatusReasonsInnerMeta type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogV1alpha1AssetStatusReasonsInnerMeta{}

// CatalogV1alpha1AssetStatusReasonsInnerMeta struct for CatalogV1alpha1AssetStatusReasonsInnerMeta
type CatalogV1alpha1AssetStatusReasonsInnerMeta struct {
	AssetResource *string `json:"assetResource,omitempty"`
}

// NewCatalogV1alpha1AssetStatusReasonsInnerMeta instantiates a new CatalogV1alpha1AssetStatusReasonsInnerMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogV1alpha1AssetStatusReasonsInnerMeta() *CatalogV1alpha1AssetStatusReasonsInnerMeta {
	this := CatalogV1alpha1AssetStatusReasonsInnerMeta{}
	return &this
}

// NewCatalogV1alpha1AssetStatusReasonsInnerMetaWithDefaults instantiates a new CatalogV1alpha1AssetStatusReasonsInnerMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogV1alpha1AssetStatusReasonsInnerMetaWithDefaults() *CatalogV1alpha1AssetStatusReasonsInnerMeta {
	this := CatalogV1alpha1AssetStatusReasonsInnerMeta{}
	return &this
}

// GetAssetResource returns the AssetResource field value if set, zero value otherwise.
func (o *CatalogV1alpha1AssetStatusReasonsInnerMeta) GetAssetResource() string {
	if o == nil || IsNil(o.AssetResource) {
		var ret string
		return ret
	}
	return *o.AssetResource
}

// GetAssetResourceOk returns a tuple with the AssetResource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1AssetStatusReasonsInnerMeta) GetAssetResourceOk() (*string, bool) {
	if o == nil || IsNil(o.AssetResource) {
		return nil, false
	}
	return o.AssetResource, true
}

// HasAssetResource returns a boolean if a field has been set.
func (o *CatalogV1alpha1AssetStatusReasonsInnerMeta) HasAssetResource() bool {
	if o != nil && !IsNil(o.AssetResource) {
		return true
	}

	return false
}

// SetAssetResource gets a reference to the given string and assigns it to the AssetResource field.
func (o *CatalogV1alpha1AssetStatusReasonsInnerMeta) SetAssetResource(v string) {
	o.AssetResource = &v
}

func (o CatalogV1alpha1AssetStatusReasonsInnerMeta) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogV1alpha1AssetStatusReasonsInnerMeta) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AssetResource) {
		toSerialize["assetResource"] = o.AssetResource
	}
	return toSerialize, nil
}

type NullableCatalogV1alpha1AssetStatusReasonsInnerMeta struct {
	value *CatalogV1alpha1AssetStatusReasonsInnerMeta
	isSet bool
}

func (v NullableCatalogV1alpha1AssetStatusReasonsInnerMeta) Get() *CatalogV1alpha1AssetStatusReasonsInnerMeta {
	return v.value
}

func (v *NullableCatalogV1alpha1AssetStatusReasonsInnerMeta) Set(val *CatalogV1alpha1AssetStatusReasonsInnerMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogV1alpha1AssetStatusReasonsInnerMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogV1alpha1AssetStatusReasonsInnerMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogV1alpha1AssetStatusReasonsInnerMeta(val *CatalogV1alpha1AssetStatusReasonsInnerMeta) *NullableCatalogV1alpha1AssetStatusReasonsInnerMeta {
	return &NullableCatalogV1alpha1AssetStatusReasonsInnerMeta{value: val, isSet: true}
}

func (v NullableCatalogV1alpha1AssetStatusReasonsInnerMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogV1alpha1AssetStatusReasonsInnerMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


