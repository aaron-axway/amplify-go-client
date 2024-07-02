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

// checks if the CatalogV1QuotaSpecPerUnitPricingType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogV1QuotaSpecPerUnitPricingType{}

// CatalogV1QuotaSpecPerUnitPricingType struct for CatalogV1QuotaSpecPerUnitPricingType
type CatalogV1QuotaSpecPerUnitPricingType struct {
	Type string `json:"type"`
	// The cost per unit of usage.
	Cost float64 `json:"cost"`
}

type _CatalogV1QuotaSpecPerUnitPricingType CatalogV1QuotaSpecPerUnitPricingType

// NewCatalogV1QuotaSpecPerUnitPricingType instantiates a new CatalogV1QuotaSpecPerUnitPricingType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogV1QuotaSpecPerUnitPricingType(type_ string, cost float64) *CatalogV1QuotaSpecPerUnitPricingType {
	this := CatalogV1QuotaSpecPerUnitPricingType{}
	this.Type = type_
	this.Cost = cost
	return &this
}

// NewCatalogV1QuotaSpecPerUnitPricingTypeWithDefaults instantiates a new CatalogV1QuotaSpecPerUnitPricingType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogV1QuotaSpecPerUnitPricingTypeWithDefaults() *CatalogV1QuotaSpecPerUnitPricingType {
	this := CatalogV1QuotaSpecPerUnitPricingType{}
	return &this
}

// GetType returns the Type field value
func (o *CatalogV1QuotaSpecPerUnitPricingType) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CatalogV1QuotaSpecPerUnitPricingType) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CatalogV1QuotaSpecPerUnitPricingType) SetType(v string) {
	o.Type = v
}

// GetCost returns the Cost field value
func (o *CatalogV1QuotaSpecPerUnitPricingType) GetCost() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Cost
}

// GetCostOk returns a tuple with the Cost field value
// and a boolean to check if the value has been set.
func (o *CatalogV1QuotaSpecPerUnitPricingType) GetCostOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cost, true
}

// SetCost sets field value
func (o *CatalogV1QuotaSpecPerUnitPricingType) SetCost(v float64) {
	o.Cost = v
}

func (o CatalogV1QuotaSpecPerUnitPricingType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogV1QuotaSpecPerUnitPricingType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["cost"] = o.Cost
	return toSerialize, nil
}

func (o *CatalogV1QuotaSpecPerUnitPricingType) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"cost",
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

	varCatalogV1QuotaSpecPerUnitPricingType := _CatalogV1QuotaSpecPerUnitPricingType{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCatalogV1QuotaSpecPerUnitPricingType)

	if err != nil {
		return err
	}

	*o = CatalogV1QuotaSpecPerUnitPricingType(varCatalogV1QuotaSpecPerUnitPricingType)

	return err
}

type NullableCatalogV1QuotaSpecPerUnitPricingType struct {
	value *CatalogV1QuotaSpecPerUnitPricingType
	isSet bool
}

func (v NullableCatalogV1QuotaSpecPerUnitPricingType) Get() *CatalogV1QuotaSpecPerUnitPricingType {
	return v.value
}

func (v *NullableCatalogV1QuotaSpecPerUnitPricingType) Set(val *CatalogV1QuotaSpecPerUnitPricingType) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogV1QuotaSpecPerUnitPricingType) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogV1QuotaSpecPerUnitPricingType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogV1QuotaSpecPerUnitPricingType(val *CatalogV1QuotaSpecPerUnitPricingType) *NullableCatalogV1QuotaSpecPerUnitPricingType {
	return &NullableCatalogV1QuotaSpecPerUnitPricingType{value: val, isSet: true}
}

func (v NullableCatalogV1QuotaSpecPerUnitPricingType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogV1QuotaSpecPerUnitPricingType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


