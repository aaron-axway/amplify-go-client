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

// CatalogV1alpha1QuotaStatusReasonsInner - struct for CatalogV1alpha1QuotaStatusReasonsInner
type CatalogV1alpha1QuotaStatusReasonsInner struct {
	CatalogV1alpha1QuotaStatusError *CatalogV1alpha1QuotaStatusError
	CatalogV1alpha1QuotaStatusPending *CatalogV1alpha1QuotaStatusPending
	CatalogV1alpha1QuotaStatusSuccess *CatalogV1alpha1QuotaStatusSuccess
}

// CatalogV1alpha1QuotaStatusErrorAsCatalogV1alpha1QuotaStatusReasonsInner is a convenience function that returns CatalogV1alpha1QuotaStatusError wrapped in CatalogV1alpha1QuotaStatusReasonsInner
func CatalogV1alpha1QuotaStatusErrorAsCatalogV1alpha1QuotaStatusReasonsInner(v *CatalogV1alpha1QuotaStatusError) CatalogV1alpha1QuotaStatusReasonsInner {
	return CatalogV1alpha1QuotaStatusReasonsInner{
		CatalogV1alpha1QuotaStatusError: v,
	}
}

// CatalogV1alpha1QuotaStatusPendingAsCatalogV1alpha1QuotaStatusReasonsInner is a convenience function that returns CatalogV1alpha1QuotaStatusPending wrapped in CatalogV1alpha1QuotaStatusReasonsInner
func CatalogV1alpha1QuotaStatusPendingAsCatalogV1alpha1QuotaStatusReasonsInner(v *CatalogV1alpha1QuotaStatusPending) CatalogV1alpha1QuotaStatusReasonsInner {
	return CatalogV1alpha1QuotaStatusReasonsInner{
		CatalogV1alpha1QuotaStatusPending: v,
	}
}

// CatalogV1alpha1QuotaStatusSuccessAsCatalogV1alpha1QuotaStatusReasonsInner is a convenience function that returns CatalogV1alpha1QuotaStatusSuccess wrapped in CatalogV1alpha1QuotaStatusReasonsInner
func CatalogV1alpha1QuotaStatusSuccessAsCatalogV1alpha1QuotaStatusReasonsInner(v *CatalogV1alpha1QuotaStatusSuccess) CatalogV1alpha1QuotaStatusReasonsInner {
	return CatalogV1alpha1QuotaStatusReasonsInner{
		CatalogV1alpha1QuotaStatusSuccess: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *CatalogV1alpha1QuotaStatusReasonsInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CatalogV1alpha1QuotaStatusError
	err = newStrictDecoder(data).Decode(&dst.CatalogV1alpha1QuotaStatusError)
	if err == nil {
		jsonCatalogV1alpha1QuotaStatusError, _ := json.Marshal(dst.CatalogV1alpha1QuotaStatusError)
		if string(jsonCatalogV1alpha1QuotaStatusError) == "{}" { // empty struct
			dst.CatalogV1alpha1QuotaStatusError = nil
		} else {
			if err = validator.Validate(dst.CatalogV1alpha1QuotaStatusError); err != nil {
				dst.CatalogV1alpha1QuotaStatusError = nil
			} else {
				match++
			}
		}
	} else {
		dst.CatalogV1alpha1QuotaStatusError = nil
	}

	// try to unmarshal data into CatalogV1alpha1QuotaStatusPending
	err = newStrictDecoder(data).Decode(&dst.CatalogV1alpha1QuotaStatusPending)
	if err == nil {
		jsonCatalogV1alpha1QuotaStatusPending, _ := json.Marshal(dst.CatalogV1alpha1QuotaStatusPending)
		if string(jsonCatalogV1alpha1QuotaStatusPending) == "{}" { // empty struct
			dst.CatalogV1alpha1QuotaStatusPending = nil
		} else {
			if err = validator.Validate(dst.CatalogV1alpha1QuotaStatusPending); err != nil {
				dst.CatalogV1alpha1QuotaStatusPending = nil
			} else {
				match++
			}
		}
	} else {
		dst.CatalogV1alpha1QuotaStatusPending = nil
	}

	// try to unmarshal data into CatalogV1alpha1QuotaStatusSuccess
	err = newStrictDecoder(data).Decode(&dst.CatalogV1alpha1QuotaStatusSuccess)
	if err == nil {
		jsonCatalogV1alpha1QuotaStatusSuccess, _ := json.Marshal(dst.CatalogV1alpha1QuotaStatusSuccess)
		if string(jsonCatalogV1alpha1QuotaStatusSuccess) == "{}" { // empty struct
			dst.CatalogV1alpha1QuotaStatusSuccess = nil
		} else {
			if err = validator.Validate(dst.CatalogV1alpha1QuotaStatusSuccess); err != nil {
				dst.CatalogV1alpha1QuotaStatusSuccess = nil
			} else {
				match++
			}
		}
	} else {
		dst.CatalogV1alpha1QuotaStatusSuccess = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.CatalogV1alpha1QuotaStatusError = nil
		dst.CatalogV1alpha1QuotaStatusPending = nil
		dst.CatalogV1alpha1QuotaStatusSuccess = nil

		return fmt.Errorf("data matches more than one schema in oneOf(CatalogV1alpha1QuotaStatusReasonsInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(CatalogV1alpha1QuotaStatusReasonsInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CatalogV1alpha1QuotaStatusReasonsInner) MarshalJSON() ([]byte, error) {
	if src.CatalogV1alpha1QuotaStatusError != nil {
		return json.Marshal(&src.CatalogV1alpha1QuotaStatusError)
	}

	if src.CatalogV1alpha1QuotaStatusPending != nil {
		return json.Marshal(&src.CatalogV1alpha1QuotaStatusPending)
	}

	if src.CatalogV1alpha1QuotaStatusSuccess != nil {
		return json.Marshal(&src.CatalogV1alpha1QuotaStatusSuccess)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CatalogV1alpha1QuotaStatusReasonsInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.CatalogV1alpha1QuotaStatusError != nil {
		return obj.CatalogV1alpha1QuotaStatusError
	}

	if obj.CatalogV1alpha1QuotaStatusPending != nil {
		return obj.CatalogV1alpha1QuotaStatusPending
	}

	if obj.CatalogV1alpha1QuotaStatusSuccess != nil {
		return obj.CatalogV1alpha1QuotaStatusSuccess
	}

	// all schemas are nil
	return nil
}

type NullableCatalogV1alpha1QuotaStatusReasonsInner struct {
	value *CatalogV1alpha1QuotaStatusReasonsInner
	isSet bool
}

func (v NullableCatalogV1alpha1QuotaStatusReasonsInner) Get() *CatalogV1alpha1QuotaStatusReasonsInner {
	return v.value
}

func (v *NullableCatalogV1alpha1QuotaStatusReasonsInner) Set(val *CatalogV1alpha1QuotaStatusReasonsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogV1alpha1QuotaStatusReasonsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogV1alpha1QuotaStatusReasonsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogV1alpha1QuotaStatusReasonsInner(val *CatalogV1alpha1QuotaStatusReasonsInner) *NullableCatalogV1alpha1QuotaStatusReasonsInner {
	return &NullableCatalogV1alpha1QuotaStatusReasonsInner{value: val, isSet: true}
}

func (v NullableCatalogV1alpha1QuotaStatusReasonsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogV1alpha1QuotaStatusReasonsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

