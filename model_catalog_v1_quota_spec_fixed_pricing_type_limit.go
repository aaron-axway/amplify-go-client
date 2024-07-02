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

// CatalogV1QuotaSpecFixedPricingTypeLimit - struct for CatalogV1QuotaSpecFixedPricingTypeLimit
type CatalogV1QuotaSpecFixedPricingTypeLimit struct {
	CatalogV1QuotaSpecLimitTypeLoose *CatalogV1QuotaSpecLimitTypeLoose
	CatalogV1QuotaSpecLimitTypeStrict *CatalogV1QuotaSpecLimitTypeStrict
}

// CatalogV1QuotaSpecLimitTypeLooseAsCatalogV1QuotaSpecFixedPricingTypeLimit is a convenience function that returns CatalogV1QuotaSpecLimitTypeLoose wrapped in CatalogV1QuotaSpecFixedPricingTypeLimit
func CatalogV1QuotaSpecLimitTypeLooseAsCatalogV1QuotaSpecFixedPricingTypeLimit(v *CatalogV1QuotaSpecLimitTypeLoose) CatalogV1QuotaSpecFixedPricingTypeLimit {
	return CatalogV1QuotaSpecFixedPricingTypeLimit{
		CatalogV1QuotaSpecLimitTypeLoose: v,
	}
}

// CatalogV1QuotaSpecLimitTypeStrictAsCatalogV1QuotaSpecFixedPricingTypeLimit is a convenience function that returns CatalogV1QuotaSpecLimitTypeStrict wrapped in CatalogV1QuotaSpecFixedPricingTypeLimit
func CatalogV1QuotaSpecLimitTypeStrictAsCatalogV1QuotaSpecFixedPricingTypeLimit(v *CatalogV1QuotaSpecLimitTypeStrict) CatalogV1QuotaSpecFixedPricingTypeLimit {
	return CatalogV1QuotaSpecFixedPricingTypeLimit{
		CatalogV1QuotaSpecLimitTypeStrict: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *CatalogV1QuotaSpecFixedPricingTypeLimit) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CatalogV1QuotaSpecLimitTypeLoose
	err = newStrictDecoder(data).Decode(&dst.CatalogV1QuotaSpecLimitTypeLoose)
	if err == nil {
		jsonCatalogV1QuotaSpecLimitTypeLoose, _ := json.Marshal(dst.CatalogV1QuotaSpecLimitTypeLoose)
		if string(jsonCatalogV1QuotaSpecLimitTypeLoose) == "{}" { // empty struct
			dst.CatalogV1QuotaSpecLimitTypeLoose = nil
		} else {
			if err = validator.Validate(dst.CatalogV1QuotaSpecLimitTypeLoose); err != nil {
				dst.CatalogV1QuotaSpecLimitTypeLoose = nil
			} else {
				match++
			}
		}
	} else {
		dst.CatalogV1QuotaSpecLimitTypeLoose = nil
	}

	// try to unmarshal data into CatalogV1QuotaSpecLimitTypeStrict
	err = newStrictDecoder(data).Decode(&dst.CatalogV1QuotaSpecLimitTypeStrict)
	if err == nil {
		jsonCatalogV1QuotaSpecLimitTypeStrict, _ := json.Marshal(dst.CatalogV1QuotaSpecLimitTypeStrict)
		if string(jsonCatalogV1QuotaSpecLimitTypeStrict) == "{}" { // empty struct
			dst.CatalogV1QuotaSpecLimitTypeStrict = nil
		} else {
			if err = validator.Validate(dst.CatalogV1QuotaSpecLimitTypeStrict); err != nil {
				dst.CatalogV1QuotaSpecLimitTypeStrict = nil
			} else {
				match++
			}
		}
	} else {
		dst.CatalogV1QuotaSpecLimitTypeStrict = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.CatalogV1QuotaSpecLimitTypeLoose = nil
		dst.CatalogV1QuotaSpecLimitTypeStrict = nil

		return fmt.Errorf("data matches more than one schema in oneOf(CatalogV1QuotaSpecFixedPricingTypeLimit)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(CatalogV1QuotaSpecFixedPricingTypeLimit)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CatalogV1QuotaSpecFixedPricingTypeLimit) MarshalJSON() ([]byte, error) {
	if src.CatalogV1QuotaSpecLimitTypeLoose != nil {
		return json.Marshal(&src.CatalogV1QuotaSpecLimitTypeLoose)
	}

	if src.CatalogV1QuotaSpecLimitTypeStrict != nil {
		return json.Marshal(&src.CatalogV1QuotaSpecLimitTypeStrict)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CatalogV1QuotaSpecFixedPricingTypeLimit) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.CatalogV1QuotaSpecLimitTypeLoose != nil {
		return obj.CatalogV1QuotaSpecLimitTypeLoose
	}

	if obj.CatalogV1QuotaSpecLimitTypeStrict != nil {
		return obj.CatalogV1QuotaSpecLimitTypeStrict
	}

	// all schemas are nil
	return nil
}

type NullableCatalogV1QuotaSpecFixedPricingTypeLimit struct {
	value *CatalogV1QuotaSpecFixedPricingTypeLimit
	isSet bool
}

func (v NullableCatalogV1QuotaSpecFixedPricingTypeLimit) Get() *CatalogV1QuotaSpecFixedPricingTypeLimit {
	return v.value
}

func (v *NullableCatalogV1QuotaSpecFixedPricingTypeLimit) Set(val *CatalogV1QuotaSpecFixedPricingTypeLimit) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogV1QuotaSpecFixedPricingTypeLimit) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogV1QuotaSpecFixedPricingTypeLimit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogV1QuotaSpecFixedPricingTypeLimit(val *CatalogV1QuotaSpecFixedPricingTypeLimit) *NullableCatalogV1QuotaSpecFixedPricingTypeLimit {
	return &NullableCatalogV1QuotaSpecFixedPricingTypeLimit{value: val, isSet: true}
}

func (v NullableCatalogV1QuotaSpecFixedPricingTypeLimit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogV1QuotaSpecFixedPricingTypeLimit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


