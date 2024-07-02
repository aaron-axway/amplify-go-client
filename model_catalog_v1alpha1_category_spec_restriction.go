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

// CatalogV1alpha1CategorySpecRestriction - struct for CatalogV1alpha1CategorySpecRestriction
type CatalogV1alpha1CategorySpecRestriction struct {
	CatalogV1alpha1CategorySpecAPIServiceRestriction *CatalogV1alpha1CategorySpecAPIServiceRestriction
	CatalogV1alpha1CategorySpecAssetRestriction *CatalogV1alpha1CategorySpecAssetRestriction
	CatalogV1alpha1CategorySpecProductRestriction *CatalogV1alpha1CategorySpecProductRestriction
}

// CatalogV1alpha1CategorySpecAPIServiceRestrictionAsCatalogV1alpha1CategorySpecRestriction is a convenience function that returns CatalogV1alpha1CategorySpecAPIServiceRestriction wrapped in CatalogV1alpha1CategorySpecRestriction
func CatalogV1alpha1CategorySpecAPIServiceRestrictionAsCatalogV1alpha1CategorySpecRestriction(v *CatalogV1alpha1CategorySpecAPIServiceRestriction) CatalogV1alpha1CategorySpecRestriction {
	return CatalogV1alpha1CategorySpecRestriction{
		CatalogV1alpha1CategorySpecAPIServiceRestriction: v,
	}
}

// CatalogV1alpha1CategorySpecAssetRestrictionAsCatalogV1alpha1CategorySpecRestriction is a convenience function that returns CatalogV1alpha1CategorySpecAssetRestriction wrapped in CatalogV1alpha1CategorySpecRestriction
func CatalogV1alpha1CategorySpecAssetRestrictionAsCatalogV1alpha1CategorySpecRestriction(v *CatalogV1alpha1CategorySpecAssetRestriction) CatalogV1alpha1CategorySpecRestriction {
	return CatalogV1alpha1CategorySpecRestriction{
		CatalogV1alpha1CategorySpecAssetRestriction: v,
	}
}

// CatalogV1alpha1CategorySpecProductRestrictionAsCatalogV1alpha1CategorySpecRestriction is a convenience function that returns CatalogV1alpha1CategorySpecProductRestriction wrapped in CatalogV1alpha1CategorySpecRestriction
func CatalogV1alpha1CategorySpecProductRestrictionAsCatalogV1alpha1CategorySpecRestriction(v *CatalogV1alpha1CategorySpecProductRestriction) CatalogV1alpha1CategorySpecRestriction {
	return CatalogV1alpha1CategorySpecRestriction{
		CatalogV1alpha1CategorySpecProductRestriction: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *CatalogV1alpha1CategorySpecRestriction) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CatalogV1alpha1CategorySpecAPIServiceRestriction
	err = newStrictDecoder(data).Decode(&dst.CatalogV1alpha1CategorySpecAPIServiceRestriction)
	if err == nil {
		jsonCatalogV1alpha1CategorySpecAPIServiceRestriction, _ := json.Marshal(dst.CatalogV1alpha1CategorySpecAPIServiceRestriction)
		if string(jsonCatalogV1alpha1CategorySpecAPIServiceRestriction) == "{}" { // empty struct
			dst.CatalogV1alpha1CategorySpecAPIServiceRestriction = nil
		} else {
			if err = validator.Validate(dst.CatalogV1alpha1CategorySpecAPIServiceRestriction); err != nil {
				dst.CatalogV1alpha1CategorySpecAPIServiceRestriction = nil
			} else {
				match++
			}
		}
	} else {
		dst.CatalogV1alpha1CategorySpecAPIServiceRestriction = nil
	}

	// try to unmarshal data into CatalogV1alpha1CategorySpecAssetRestriction
	err = newStrictDecoder(data).Decode(&dst.CatalogV1alpha1CategorySpecAssetRestriction)
	if err == nil {
		jsonCatalogV1alpha1CategorySpecAssetRestriction, _ := json.Marshal(dst.CatalogV1alpha1CategorySpecAssetRestriction)
		if string(jsonCatalogV1alpha1CategorySpecAssetRestriction) == "{}" { // empty struct
			dst.CatalogV1alpha1CategorySpecAssetRestriction = nil
		} else {
			if err = validator.Validate(dst.CatalogV1alpha1CategorySpecAssetRestriction); err != nil {
				dst.CatalogV1alpha1CategorySpecAssetRestriction = nil
			} else {
				match++
			}
		}
	} else {
		dst.CatalogV1alpha1CategorySpecAssetRestriction = nil
	}

	// try to unmarshal data into CatalogV1alpha1CategorySpecProductRestriction
	err = newStrictDecoder(data).Decode(&dst.CatalogV1alpha1CategorySpecProductRestriction)
	if err == nil {
		jsonCatalogV1alpha1CategorySpecProductRestriction, _ := json.Marshal(dst.CatalogV1alpha1CategorySpecProductRestriction)
		if string(jsonCatalogV1alpha1CategorySpecProductRestriction) == "{}" { // empty struct
			dst.CatalogV1alpha1CategorySpecProductRestriction = nil
		} else {
			if err = validator.Validate(dst.CatalogV1alpha1CategorySpecProductRestriction); err != nil {
				dst.CatalogV1alpha1CategorySpecProductRestriction = nil
			} else {
				match++
			}
		}
	} else {
		dst.CatalogV1alpha1CategorySpecProductRestriction = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.CatalogV1alpha1CategorySpecAPIServiceRestriction = nil
		dst.CatalogV1alpha1CategorySpecAssetRestriction = nil
		dst.CatalogV1alpha1CategorySpecProductRestriction = nil

		return fmt.Errorf("data matches more than one schema in oneOf(CatalogV1alpha1CategorySpecRestriction)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(CatalogV1alpha1CategorySpecRestriction)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CatalogV1alpha1CategorySpecRestriction) MarshalJSON() ([]byte, error) {
	if src.CatalogV1alpha1CategorySpecAPIServiceRestriction != nil {
		return json.Marshal(&src.CatalogV1alpha1CategorySpecAPIServiceRestriction)
	}

	if src.CatalogV1alpha1CategorySpecAssetRestriction != nil {
		return json.Marshal(&src.CatalogV1alpha1CategorySpecAssetRestriction)
	}

	if src.CatalogV1alpha1CategorySpecProductRestriction != nil {
		return json.Marshal(&src.CatalogV1alpha1CategorySpecProductRestriction)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CatalogV1alpha1CategorySpecRestriction) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.CatalogV1alpha1CategorySpecAPIServiceRestriction != nil {
		return obj.CatalogV1alpha1CategorySpecAPIServiceRestriction
	}

	if obj.CatalogV1alpha1CategorySpecAssetRestriction != nil {
		return obj.CatalogV1alpha1CategorySpecAssetRestriction
	}

	if obj.CatalogV1alpha1CategorySpecProductRestriction != nil {
		return obj.CatalogV1alpha1CategorySpecProductRestriction
	}

	// all schemas are nil
	return nil
}

type NullableCatalogV1alpha1CategorySpecRestriction struct {
	value *CatalogV1alpha1CategorySpecRestriction
	isSet bool
}

func (v NullableCatalogV1alpha1CategorySpecRestriction) Get() *CatalogV1alpha1CategorySpecRestriction {
	return v.value
}

func (v *NullableCatalogV1alpha1CategorySpecRestriction) Set(val *CatalogV1alpha1CategorySpecRestriction) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogV1alpha1CategorySpecRestriction) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogV1alpha1CategorySpecRestriction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogV1alpha1CategorySpecRestriction(val *CatalogV1alpha1CategorySpecRestriction) *NullableCatalogV1alpha1CategorySpecRestriction {
	return &NullableCatalogV1alpha1CategorySpecRestriction{value: val, isSet: true}
}

func (v NullableCatalogV1alpha1CategorySpecRestriction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogV1alpha1CategorySpecRestriction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


