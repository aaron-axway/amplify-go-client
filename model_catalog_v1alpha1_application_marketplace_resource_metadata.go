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

// checks if the CatalogV1alpha1ApplicationMarketplaceResourceMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogV1alpha1ApplicationMarketplaceResourceMetadata{}

// CatalogV1alpha1ApplicationMarketplaceResourceMetadata Marketplace Application metadata.
type CatalogV1alpha1ApplicationMarketplaceResourceMetadata struct {
	// Id of the user that created the entity.
	CreateUserId *string `json:"createUserId,omitempty"`
	// Id of the user that created the entity.
	ModifyUserId *string `json:"modifyUserId,omitempty"`
}

// NewCatalogV1alpha1ApplicationMarketplaceResourceMetadata instantiates a new CatalogV1alpha1ApplicationMarketplaceResourceMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogV1alpha1ApplicationMarketplaceResourceMetadata() *CatalogV1alpha1ApplicationMarketplaceResourceMetadata {
	this := CatalogV1alpha1ApplicationMarketplaceResourceMetadata{}
	return &this
}

// NewCatalogV1alpha1ApplicationMarketplaceResourceMetadataWithDefaults instantiates a new CatalogV1alpha1ApplicationMarketplaceResourceMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogV1alpha1ApplicationMarketplaceResourceMetadataWithDefaults() *CatalogV1alpha1ApplicationMarketplaceResourceMetadata {
	this := CatalogV1alpha1ApplicationMarketplaceResourceMetadata{}
	return &this
}

// GetCreateUserId returns the CreateUserId field value if set, zero value otherwise.
func (o *CatalogV1alpha1ApplicationMarketplaceResourceMetadata) GetCreateUserId() string {
	if o == nil || IsNil(o.CreateUserId) {
		var ret string
		return ret
	}
	return *o.CreateUserId
}

// GetCreateUserIdOk returns a tuple with the CreateUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1ApplicationMarketplaceResourceMetadata) GetCreateUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.CreateUserId) {
		return nil, false
	}
	return o.CreateUserId, true
}

// HasCreateUserId returns a boolean if a field has been set.
func (o *CatalogV1alpha1ApplicationMarketplaceResourceMetadata) HasCreateUserId() bool {
	if o != nil && !IsNil(o.CreateUserId) {
		return true
	}

	return false
}

// SetCreateUserId gets a reference to the given string and assigns it to the CreateUserId field.
func (o *CatalogV1alpha1ApplicationMarketplaceResourceMetadata) SetCreateUserId(v string) {
	o.CreateUserId = &v
}

// GetModifyUserId returns the ModifyUserId field value if set, zero value otherwise.
func (o *CatalogV1alpha1ApplicationMarketplaceResourceMetadata) GetModifyUserId() string {
	if o == nil || IsNil(o.ModifyUserId) {
		var ret string
		return ret
	}
	return *o.ModifyUserId
}

// GetModifyUserIdOk returns a tuple with the ModifyUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1ApplicationMarketplaceResourceMetadata) GetModifyUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.ModifyUserId) {
		return nil, false
	}
	return o.ModifyUserId, true
}

// HasModifyUserId returns a boolean if a field has been set.
func (o *CatalogV1alpha1ApplicationMarketplaceResourceMetadata) HasModifyUserId() bool {
	if o != nil && !IsNil(o.ModifyUserId) {
		return true
	}

	return false
}

// SetModifyUserId gets a reference to the given string and assigns it to the ModifyUserId field.
func (o *CatalogV1alpha1ApplicationMarketplaceResourceMetadata) SetModifyUserId(v string) {
	o.ModifyUserId = &v
}

func (o CatalogV1alpha1ApplicationMarketplaceResourceMetadata) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogV1alpha1ApplicationMarketplaceResourceMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreateUserId) {
		toSerialize["createUserId"] = o.CreateUserId
	}
	if !IsNil(o.ModifyUserId) {
		toSerialize["modifyUserId"] = o.ModifyUserId
	}
	return toSerialize, nil
}

type NullableCatalogV1alpha1ApplicationMarketplaceResourceMetadata struct {
	value *CatalogV1alpha1ApplicationMarketplaceResourceMetadata
	isSet bool
}

func (v NullableCatalogV1alpha1ApplicationMarketplaceResourceMetadata) Get() *CatalogV1alpha1ApplicationMarketplaceResourceMetadata {
	return v.value
}

func (v *NullableCatalogV1alpha1ApplicationMarketplaceResourceMetadata) Set(val *CatalogV1alpha1ApplicationMarketplaceResourceMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogV1alpha1ApplicationMarketplaceResourceMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogV1alpha1ApplicationMarketplaceResourceMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogV1alpha1ApplicationMarketplaceResourceMetadata(val *CatalogV1alpha1ApplicationMarketplaceResourceMetadata) *NullableCatalogV1alpha1ApplicationMarketplaceResourceMetadata {
	return &NullableCatalogV1alpha1ApplicationMarketplaceResourceMetadata{value: val, isSet: true}
}

func (v NullableCatalogV1alpha1ApplicationMarketplaceResourceMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogV1alpha1ApplicationMarketplaceResourceMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


