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

// checks if the CatalogV1QuotaSpecGraduatedPricingType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogV1QuotaSpecGraduatedPricingType{}

// CatalogV1QuotaSpecGraduatedPricingType struct for CatalogV1QuotaSpecGraduatedPricingType
type CatalogV1QuotaSpecGraduatedPricingType struct {
	Type string `json:"type"`
	Limit CatalogV1QuotaSpecGraduatedPricingTypeLimit `json:"limit"`
}

type _CatalogV1QuotaSpecGraduatedPricingType CatalogV1QuotaSpecGraduatedPricingType

// NewCatalogV1QuotaSpecGraduatedPricingType instantiates a new CatalogV1QuotaSpecGraduatedPricingType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogV1QuotaSpecGraduatedPricingType(type_ string, limit CatalogV1QuotaSpecGraduatedPricingTypeLimit) *CatalogV1QuotaSpecGraduatedPricingType {
	this := CatalogV1QuotaSpecGraduatedPricingType{}
	this.Type = type_
	this.Limit = limit
	return &this
}

// NewCatalogV1QuotaSpecGraduatedPricingTypeWithDefaults instantiates a new CatalogV1QuotaSpecGraduatedPricingType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogV1QuotaSpecGraduatedPricingTypeWithDefaults() *CatalogV1QuotaSpecGraduatedPricingType {
	this := CatalogV1QuotaSpecGraduatedPricingType{}
	return &this
}

// GetType returns the Type field value
func (o *CatalogV1QuotaSpecGraduatedPricingType) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CatalogV1QuotaSpecGraduatedPricingType) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CatalogV1QuotaSpecGraduatedPricingType) SetType(v string) {
	o.Type = v
}

// GetLimit returns the Limit field value
func (o *CatalogV1QuotaSpecGraduatedPricingType) GetLimit() CatalogV1QuotaSpecGraduatedPricingTypeLimit {
	if o == nil {
		var ret CatalogV1QuotaSpecGraduatedPricingTypeLimit
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *CatalogV1QuotaSpecGraduatedPricingType) GetLimitOk() (*CatalogV1QuotaSpecGraduatedPricingTypeLimit, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *CatalogV1QuotaSpecGraduatedPricingType) SetLimit(v CatalogV1QuotaSpecGraduatedPricingTypeLimit) {
	o.Limit = v
}

func (o CatalogV1QuotaSpecGraduatedPricingType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogV1QuotaSpecGraduatedPricingType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["limit"] = o.Limit
	return toSerialize, nil
}

func (o *CatalogV1QuotaSpecGraduatedPricingType) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"limit",
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

	varCatalogV1QuotaSpecGraduatedPricingType := _CatalogV1QuotaSpecGraduatedPricingType{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCatalogV1QuotaSpecGraduatedPricingType)

	if err != nil {
		return err
	}

	*o = CatalogV1QuotaSpecGraduatedPricingType(varCatalogV1QuotaSpecGraduatedPricingType)

	return err
}

type NullableCatalogV1QuotaSpecGraduatedPricingType struct {
	value *CatalogV1QuotaSpecGraduatedPricingType
	isSet bool
}

func (v NullableCatalogV1QuotaSpecGraduatedPricingType) Get() *CatalogV1QuotaSpecGraduatedPricingType {
	return v.value
}

func (v *NullableCatalogV1QuotaSpecGraduatedPricingType) Set(val *CatalogV1QuotaSpecGraduatedPricingType) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogV1QuotaSpecGraduatedPricingType) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogV1QuotaSpecGraduatedPricingType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogV1QuotaSpecGraduatedPricingType(val *CatalogV1QuotaSpecGraduatedPricingType) *NullableCatalogV1QuotaSpecGraduatedPricingType {
	return &NullableCatalogV1QuotaSpecGraduatedPricingType{value: val, isSet: true}
}

func (v NullableCatalogV1QuotaSpecGraduatedPricingType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogV1QuotaSpecGraduatedPricingType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


