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
	"gopkg.in/validator.v2"
	"fmt"
)

// CatalogV1QuotaSpecGraduatedPricingTypeLimit - The tiered limits to set.
type CatalogV1QuotaSpecGraduatedPricingTypeLimit struct {
	CatalogV1QuotaSpecLimitTypeTiered *CatalogV1QuotaSpecLimitTypeTiered
}

// CatalogV1QuotaSpecLimitTypeTieredAsCatalogV1QuotaSpecGraduatedPricingTypeLimit is a convenience function that returns CatalogV1QuotaSpecLimitTypeTiered wrapped in CatalogV1QuotaSpecGraduatedPricingTypeLimit
func CatalogV1QuotaSpecLimitTypeTieredAsCatalogV1QuotaSpecGraduatedPricingTypeLimit(v *CatalogV1QuotaSpecLimitTypeTiered) CatalogV1QuotaSpecGraduatedPricingTypeLimit {
	return CatalogV1QuotaSpecGraduatedPricingTypeLimit{
		CatalogV1QuotaSpecLimitTypeTiered: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *CatalogV1QuotaSpecGraduatedPricingTypeLimit) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CatalogV1QuotaSpecLimitTypeTiered
	err = newStrictDecoder(data).Decode(&dst.CatalogV1QuotaSpecLimitTypeTiered)
	if err == nil {
		jsonCatalogV1QuotaSpecLimitTypeTiered, _ := json.Marshal(dst.CatalogV1QuotaSpecLimitTypeTiered)
		if string(jsonCatalogV1QuotaSpecLimitTypeTiered) == "{}" { // empty struct
			dst.CatalogV1QuotaSpecLimitTypeTiered = nil
		} else {
			if err = validator.Validate(dst.CatalogV1QuotaSpecLimitTypeTiered); err != nil {
				dst.CatalogV1QuotaSpecLimitTypeTiered = nil
			} else {
				match++
			}
		}
	} else {
		dst.CatalogV1QuotaSpecLimitTypeTiered = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.CatalogV1QuotaSpecLimitTypeTiered = nil

		return fmt.Errorf("data matches more than one schema in oneOf(CatalogV1QuotaSpecGraduatedPricingTypeLimit)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(CatalogV1QuotaSpecGraduatedPricingTypeLimit)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CatalogV1QuotaSpecGraduatedPricingTypeLimit) MarshalJSON() ([]byte, error) {
	if src.CatalogV1QuotaSpecLimitTypeTiered != nil {
		return json.Marshal(&src.CatalogV1QuotaSpecLimitTypeTiered)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CatalogV1QuotaSpecGraduatedPricingTypeLimit) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.CatalogV1QuotaSpecLimitTypeTiered != nil {
		return obj.CatalogV1QuotaSpecLimitTypeTiered
	}

	// all schemas are nil
	return nil
}

type NullableCatalogV1QuotaSpecGraduatedPricingTypeLimit struct {
	value *CatalogV1QuotaSpecGraduatedPricingTypeLimit
	isSet bool
}

func (v NullableCatalogV1QuotaSpecGraduatedPricingTypeLimit) Get() *CatalogV1QuotaSpecGraduatedPricingTypeLimit {
	return v.value
}

func (v *NullableCatalogV1QuotaSpecGraduatedPricingTypeLimit) Set(val *CatalogV1QuotaSpecGraduatedPricingTypeLimit) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogV1QuotaSpecGraduatedPricingTypeLimit) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogV1QuotaSpecGraduatedPricingTypeLimit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogV1QuotaSpecGraduatedPricingTypeLimit(val *CatalogV1QuotaSpecGraduatedPricingTypeLimit) *NullableCatalogV1QuotaSpecGraduatedPricingTypeLimit {
	return &NullableCatalogV1QuotaSpecGraduatedPricingTypeLimit{value: val, isSet: true}
}

func (v NullableCatalogV1QuotaSpecGraduatedPricingTypeLimit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogV1QuotaSpecGraduatedPricingTypeLimit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


