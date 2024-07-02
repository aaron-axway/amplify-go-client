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
	"bytes"
	"fmt"
)

// checks if the CatalogV1CategorySpecProductRestrictionMarketplace type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogV1CategorySpecProductRestrictionMarketplace{}

// CatalogV1CategorySpecProductRestrictionMarketplace Defines properties to be applied to all marketplaces.
type CatalogV1CategorySpecProductRestrictionMarketplace struct {
	// Property not used anymore. CategoryVisibility resource to be used to control the featured property of the Category in a specific Marketplace.
	// Deprecated
	Featured *bool `json:"featured,omitempty"`
	// Defines if the Category is visible in the marketplace.
	Visible bool `json:"visible"`
}

type _CatalogV1CategorySpecProductRestrictionMarketplace CatalogV1CategorySpecProductRestrictionMarketplace

// NewCatalogV1CategorySpecProductRestrictionMarketplace instantiates a new CatalogV1CategorySpecProductRestrictionMarketplace object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogV1CategorySpecProductRestrictionMarketplace(visible bool) *CatalogV1CategorySpecProductRestrictionMarketplace {
	this := CatalogV1CategorySpecProductRestrictionMarketplace{}
	this.Visible = visible
	return &this
}

// NewCatalogV1CategorySpecProductRestrictionMarketplaceWithDefaults instantiates a new CatalogV1CategorySpecProductRestrictionMarketplace object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogV1CategorySpecProductRestrictionMarketplaceWithDefaults() *CatalogV1CategorySpecProductRestrictionMarketplace {
	this := CatalogV1CategorySpecProductRestrictionMarketplace{}
	return &this
}

// GetFeatured returns the Featured field value if set, zero value otherwise.
// Deprecated
func (o *CatalogV1CategorySpecProductRestrictionMarketplace) GetFeatured() bool {
	if o == nil || IsNil(o.Featured) {
		var ret bool
		return ret
	}
	return *o.Featured
}

// GetFeaturedOk returns a tuple with the Featured field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *CatalogV1CategorySpecProductRestrictionMarketplace) GetFeaturedOk() (*bool, bool) {
	if o == nil || IsNil(o.Featured) {
		return nil, false
	}
	return o.Featured, true
}

// HasFeatured returns a boolean if a field has been set.
func (o *CatalogV1CategorySpecProductRestrictionMarketplace) HasFeatured() bool {
	if o != nil && !IsNil(o.Featured) {
		return true
	}

	return false
}

// SetFeatured gets a reference to the given bool and assigns it to the Featured field.
// Deprecated
func (o *CatalogV1CategorySpecProductRestrictionMarketplace) SetFeatured(v bool) {
	o.Featured = &v
}

// GetVisible returns the Visible field value
func (o *CatalogV1CategorySpecProductRestrictionMarketplace) GetVisible() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Visible
}

// GetVisibleOk returns a tuple with the Visible field value
// and a boolean to check if the value has been set.
func (o *CatalogV1CategorySpecProductRestrictionMarketplace) GetVisibleOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Visible, true
}

// SetVisible sets field value
func (o *CatalogV1CategorySpecProductRestrictionMarketplace) SetVisible(v bool) {
	o.Visible = v
}

func (o CatalogV1CategorySpecProductRestrictionMarketplace) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogV1CategorySpecProductRestrictionMarketplace) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Featured) {
		toSerialize["featured"] = o.Featured
	}
	toSerialize["visible"] = o.Visible
	return toSerialize, nil
}

func (o *CatalogV1CategorySpecProductRestrictionMarketplace) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"visible",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCatalogV1CategorySpecProductRestrictionMarketplace := _CatalogV1CategorySpecProductRestrictionMarketplace{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCatalogV1CategorySpecProductRestrictionMarketplace)

	if err != nil {
		return err
	}

	*o = CatalogV1CategorySpecProductRestrictionMarketplace(varCatalogV1CategorySpecProductRestrictionMarketplace)

	return err
}

type NullableCatalogV1CategorySpecProductRestrictionMarketplace struct {
	value *CatalogV1CategorySpecProductRestrictionMarketplace
	isSet bool
}

func (v NullableCatalogV1CategorySpecProductRestrictionMarketplace) Get() *CatalogV1CategorySpecProductRestrictionMarketplace {
	return v.value
}

func (v *NullableCatalogV1CategorySpecProductRestrictionMarketplace) Set(val *CatalogV1CategorySpecProductRestrictionMarketplace) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogV1CategorySpecProductRestrictionMarketplace) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogV1CategorySpecProductRestrictionMarketplace) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogV1CategorySpecProductRestrictionMarketplace(val *CatalogV1CategorySpecProductRestrictionMarketplace) *NullableCatalogV1CategorySpecProductRestrictionMarketplace {
	return &NullableCatalogV1CategorySpecProductRestrictionMarketplace{value: val, isSet: true}
}

func (v NullableCatalogV1CategorySpecProductRestrictionMarketplace) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogV1CategorySpecProductRestrictionMarketplace) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

