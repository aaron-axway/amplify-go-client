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
	"gopkg.in/validator.v2"
	"fmt"
)

// CatalogV1QuotaSpecPricing - struct for CatalogV1QuotaSpecPricing
type CatalogV1QuotaSpecPricing struct {
	CatalogV1QuotaSpecFixedPricingType *CatalogV1QuotaSpecFixedPricingType
	CatalogV1QuotaSpecGraduatedPricingType *CatalogV1QuotaSpecGraduatedPricingType
	CatalogV1QuotaSpecPerUnitPricingType *CatalogV1QuotaSpecPerUnitPricingType
	CatalogV1QuotaSpecUnlimitedPricingType *CatalogV1QuotaSpecUnlimitedPricingType
	CatalogV1QuotaSpecVolumePricingType *CatalogV1QuotaSpecVolumePricingType
}

// CatalogV1QuotaSpecFixedPricingTypeAsCatalogV1QuotaSpecPricing is a convenience function that returns CatalogV1QuotaSpecFixedPricingType wrapped in CatalogV1QuotaSpecPricing
func CatalogV1QuotaSpecFixedPricingTypeAsCatalogV1QuotaSpecPricing(v *CatalogV1QuotaSpecFixedPricingType) CatalogV1QuotaSpecPricing {
	return CatalogV1QuotaSpecPricing{
		CatalogV1QuotaSpecFixedPricingType: v,
	}
}

// CatalogV1QuotaSpecGraduatedPricingTypeAsCatalogV1QuotaSpecPricing is a convenience function that returns CatalogV1QuotaSpecGraduatedPricingType wrapped in CatalogV1QuotaSpecPricing
func CatalogV1QuotaSpecGraduatedPricingTypeAsCatalogV1QuotaSpecPricing(v *CatalogV1QuotaSpecGraduatedPricingType) CatalogV1QuotaSpecPricing {
	return CatalogV1QuotaSpecPricing{
		CatalogV1QuotaSpecGraduatedPricingType: v,
	}
}

// CatalogV1QuotaSpecPerUnitPricingTypeAsCatalogV1QuotaSpecPricing is a convenience function that returns CatalogV1QuotaSpecPerUnitPricingType wrapped in CatalogV1QuotaSpecPricing
func CatalogV1QuotaSpecPerUnitPricingTypeAsCatalogV1QuotaSpecPricing(v *CatalogV1QuotaSpecPerUnitPricingType) CatalogV1QuotaSpecPricing {
	return CatalogV1QuotaSpecPricing{
		CatalogV1QuotaSpecPerUnitPricingType: v,
	}
}

// CatalogV1QuotaSpecUnlimitedPricingTypeAsCatalogV1QuotaSpecPricing is a convenience function that returns CatalogV1QuotaSpecUnlimitedPricingType wrapped in CatalogV1QuotaSpecPricing
func CatalogV1QuotaSpecUnlimitedPricingTypeAsCatalogV1QuotaSpecPricing(v *CatalogV1QuotaSpecUnlimitedPricingType) CatalogV1QuotaSpecPricing {
	return CatalogV1QuotaSpecPricing{
		CatalogV1QuotaSpecUnlimitedPricingType: v,
	}
}

// CatalogV1QuotaSpecVolumePricingTypeAsCatalogV1QuotaSpecPricing is a convenience function that returns CatalogV1QuotaSpecVolumePricingType wrapped in CatalogV1QuotaSpecPricing
func CatalogV1QuotaSpecVolumePricingTypeAsCatalogV1QuotaSpecPricing(v *CatalogV1QuotaSpecVolumePricingType) CatalogV1QuotaSpecPricing {
	return CatalogV1QuotaSpecPricing{
		CatalogV1QuotaSpecVolumePricingType: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *CatalogV1QuotaSpecPricing) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CatalogV1QuotaSpecFixedPricingType
	err = newStrictDecoder(data).Decode(&dst.CatalogV1QuotaSpecFixedPricingType)
	if err == nil {
		jsonCatalogV1QuotaSpecFixedPricingType, _ := json.Marshal(dst.CatalogV1QuotaSpecFixedPricingType)
		if string(jsonCatalogV1QuotaSpecFixedPricingType) == "{}" { // empty struct
			dst.CatalogV1QuotaSpecFixedPricingType = nil
		} else {
			if err = validator.Validate(dst.CatalogV1QuotaSpecFixedPricingType); err != nil {
				dst.CatalogV1QuotaSpecFixedPricingType = nil
			} else {
				match++
			}
		}
	} else {
		dst.CatalogV1QuotaSpecFixedPricingType = nil
	}

	// try to unmarshal data into CatalogV1QuotaSpecGraduatedPricingType
	err = newStrictDecoder(data).Decode(&dst.CatalogV1QuotaSpecGraduatedPricingType)
	if err == nil {
		jsonCatalogV1QuotaSpecGraduatedPricingType, _ := json.Marshal(dst.CatalogV1QuotaSpecGraduatedPricingType)
		if string(jsonCatalogV1QuotaSpecGraduatedPricingType) == "{}" { // empty struct
			dst.CatalogV1QuotaSpecGraduatedPricingType = nil
		} else {
			if err = validator.Validate(dst.CatalogV1QuotaSpecGraduatedPricingType); err != nil {
				dst.CatalogV1QuotaSpecGraduatedPricingType = nil
			} else {
				match++
			}
		}
	} else {
		dst.CatalogV1QuotaSpecGraduatedPricingType = nil
	}

	// try to unmarshal data into CatalogV1QuotaSpecPerUnitPricingType
	err = newStrictDecoder(data).Decode(&dst.CatalogV1QuotaSpecPerUnitPricingType)
	if err == nil {
		jsonCatalogV1QuotaSpecPerUnitPricingType, _ := json.Marshal(dst.CatalogV1QuotaSpecPerUnitPricingType)
		if string(jsonCatalogV1QuotaSpecPerUnitPricingType) == "{}" { // empty struct
			dst.CatalogV1QuotaSpecPerUnitPricingType = nil
		} else {
			if err = validator.Validate(dst.CatalogV1QuotaSpecPerUnitPricingType); err != nil {
				dst.CatalogV1QuotaSpecPerUnitPricingType = nil
			} else {
				match++
			}
		}
	} else {
		dst.CatalogV1QuotaSpecPerUnitPricingType = nil
	}

	// try to unmarshal data into CatalogV1QuotaSpecUnlimitedPricingType
	err = newStrictDecoder(data).Decode(&dst.CatalogV1QuotaSpecUnlimitedPricingType)
	if err == nil {
		jsonCatalogV1QuotaSpecUnlimitedPricingType, _ := json.Marshal(dst.CatalogV1QuotaSpecUnlimitedPricingType)
		if string(jsonCatalogV1QuotaSpecUnlimitedPricingType) == "{}" { // empty struct
			dst.CatalogV1QuotaSpecUnlimitedPricingType = nil
		} else {
			if err = validator.Validate(dst.CatalogV1QuotaSpecUnlimitedPricingType); err != nil {
				dst.CatalogV1QuotaSpecUnlimitedPricingType = nil
			} else {
				match++
			}
		}
	} else {
		dst.CatalogV1QuotaSpecUnlimitedPricingType = nil
	}

	// try to unmarshal data into CatalogV1QuotaSpecVolumePricingType
	err = newStrictDecoder(data).Decode(&dst.CatalogV1QuotaSpecVolumePricingType)
	if err == nil {
		jsonCatalogV1QuotaSpecVolumePricingType, _ := json.Marshal(dst.CatalogV1QuotaSpecVolumePricingType)
		if string(jsonCatalogV1QuotaSpecVolumePricingType) == "{}" { // empty struct
			dst.CatalogV1QuotaSpecVolumePricingType = nil
		} else {
			if err = validator.Validate(dst.CatalogV1QuotaSpecVolumePricingType); err != nil {
				dst.CatalogV1QuotaSpecVolumePricingType = nil
			} else {
				match++
			}
		}
	} else {
		dst.CatalogV1QuotaSpecVolumePricingType = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.CatalogV1QuotaSpecFixedPricingType = nil
		dst.CatalogV1QuotaSpecGraduatedPricingType = nil
		dst.CatalogV1QuotaSpecPerUnitPricingType = nil
		dst.CatalogV1QuotaSpecUnlimitedPricingType = nil
		dst.CatalogV1QuotaSpecVolumePricingType = nil

		return fmt.Errorf("data matches more than one schema in oneOf(CatalogV1QuotaSpecPricing)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(CatalogV1QuotaSpecPricing)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CatalogV1QuotaSpecPricing) MarshalJSON() ([]byte, error) {
	if src.CatalogV1QuotaSpecFixedPricingType != nil {
		return json.Marshal(&src.CatalogV1QuotaSpecFixedPricingType)
	}

	if src.CatalogV1QuotaSpecGraduatedPricingType != nil {
		return json.Marshal(&src.CatalogV1QuotaSpecGraduatedPricingType)
	}

	if src.CatalogV1QuotaSpecPerUnitPricingType != nil {
		return json.Marshal(&src.CatalogV1QuotaSpecPerUnitPricingType)
	}

	if src.CatalogV1QuotaSpecUnlimitedPricingType != nil {
		return json.Marshal(&src.CatalogV1QuotaSpecUnlimitedPricingType)
	}

	if src.CatalogV1QuotaSpecVolumePricingType != nil {
		return json.Marshal(&src.CatalogV1QuotaSpecVolumePricingType)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CatalogV1QuotaSpecPricing) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.CatalogV1QuotaSpecFixedPricingType != nil {
		return obj.CatalogV1QuotaSpecFixedPricingType
	}

	if obj.CatalogV1QuotaSpecGraduatedPricingType != nil {
		return obj.CatalogV1QuotaSpecGraduatedPricingType
	}

	if obj.CatalogV1QuotaSpecPerUnitPricingType != nil {
		return obj.CatalogV1QuotaSpecPerUnitPricingType
	}

	if obj.CatalogV1QuotaSpecUnlimitedPricingType != nil {
		return obj.CatalogV1QuotaSpecUnlimitedPricingType
	}

	if obj.CatalogV1QuotaSpecVolumePricingType != nil {
		return obj.CatalogV1QuotaSpecVolumePricingType
	}

	// all schemas are nil
	return nil
}

type NullableCatalogV1QuotaSpecPricing struct {
	value *CatalogV1QuotaSpecPricing
	isSet bool
}

func (v NullableCatalogV1QuotaSpecPricing) Get() *CatalogV1QuotaSpecPricing {
	return v.value
}

func (v *NullableCatalogV1QuotaSpecPricing) Set(val *CatalogV1QuotaSpecPricing) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogV1QuotaSpecPricing) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogV1QuotaSpecPricing) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogV1QuotaSpecPricing(val *CatalogV1QuotaSpecPricing) *NullableCatalogV1QuotaSpecPricing {
	return &NullableCatalogV1QuotaSpecPricing{value: val, isSet: true}
}

func (v NullableCatalogV1QuotaSpecPricing) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogV1QuotaSpecPricing) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

