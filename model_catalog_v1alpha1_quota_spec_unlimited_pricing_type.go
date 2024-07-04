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
	"bytes"
	"fmt"
)

// checks if the CatalogV1alpha1QuotaSpecUnlimitedPricingType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogV1alpha1QuotaSpecUnlimitedPricingType{}

// CatalogV1alpha1QuotaSpecUnlimitedPricingType struct for CatalogV1alpha1QuotaSpecUnlimitedPricingType
type CatalogV1alpha1QuotaSpecUnlimitedPricingType struct {
	Type string `json:"type"`
}

type _CatalogV1alpha1QuotaSpecUnlimitedPricingType CatalogV1alpha1QuotaSpecUnlimitedPricingType

// NewCatalogV1alpha1QuotaSpecUnlimitedPricingType instantiates a new CatalogV1alpha1QuotaSpecUnlimitedPricingType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogV1alpha1QuotaSpecUnlimitedPricingType(type_ string) *CatalogV1alpha1QuotaSpecUnlimitedPricingType {
	this := CatalogV1alpha1QuotaSpecUnlimitedPricingType{}
	this.Type = type_
	return &this
}

// NewCatalogV1alpha1QuotaSpecUnlimitedPricingTypeWithDefaults instantiates a new CatalogV1alpha1QuotaSpecUnlimitedPricingType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogV1alpha1QuotaSpecUnlimitedPricingTypeWithDefaults() *CatalogV1alpha1QuotaSpecUnlimitedPricingType {
	this := CatalogV1alpha1QuotaSpecUnlimitedPricingType{}
	return &this
}

// GetType returns the Type field value
func (o *CatalogV1alpha1QuotaSpecUnlimitedPricingType) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1QuotaSpecUnlimitedPricingType) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CatalogV1alpha1QuotaSpecUnlimitedPricingType) SetType(v string) {
	o.Type = v
}

func (o CatalogV1alpha1QuotaSpecUnlimitedPricingType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogV1alpha1QuotaSpecUnlimitedPricingType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *CatalogV1alpha1QuotaSpecUnlimitedPricingType) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
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

	varCatalogV1alpha1QuotaSpecUnlimitedPricingType := _CatalogV1alpha1QuotaSpecUnlimitedPricingType{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCatalogV1alpha1QuotaSpecUnlimitedPricingType)

	if err != nil {
		return err
	}

	*o = CatalogV1alpha1QuotaSpecUnlimitedPricingType(varCatalogV1alpha1QuotaSpecUnlimitedPricingType)

	return err
}

type NullableCatalogV1alpha1QuotaSpecUnlimitedPricingType struct {
	value *CatalogV1alpha1QuotaSpecUnlimitedPricingType
	isSet bool
}

func (v NullableCatalogV1alpha1QuotaSpecUnlimitedPricingType) Get() *CatalogV1alpha1QuotaSpecUnlimitedPricingType {
	return v.value
}

func (v *NullableCatalogV1alpha1QuotaSpecUnlimitedPricingType) Set(val *CatalogV1alpha1QuotaSpecUnlimitedPricingType) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogV1alpha1QuotaSpecUnlimitedPricingType) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogV1alpha1QuotaSpecUnlimitedPricingType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogV1alpha1QuotaSpecUnlimitedPricingType(val *CatalogV1alpha1QuotaSpecUnlimitedPricingType) *NullableCatalogV1alpha1QuotaSpecUnlimitedPricingType {
	return &NullableCatalogV1alpha1QuotaSpecUnlimitedPricingType{value: val, isSet: true}
}

func (v NullableCatalogV1alpha1QuotaSpecUnlimitedPricingType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogV1alpha1QuotaSpecUnlimitedPricingType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


