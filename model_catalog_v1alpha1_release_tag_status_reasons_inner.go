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

// CatalogV1alpha1ReleaseTagStatusReasonsInner - struct for CatalogV1alpha1ReleaseTagStatusReasonsInner
type CatalogV1alpha1ReleaseTagStatusReasonsInner struct {
	CatalogV1alpha1ReleaseTagStatusError *CatalogV1alpha1ReleaseTagStatusError
	CatalogV1alpha1ReleaseTagStatusPending *CatalogV1alpha1ReleaseTagStatusPending
	CatalogV1alpha1ReleaseTagStatusSuccess *CatalogV1alpha1ReleaseTagStatusSuccess
}

// CatalogV1alpha1ReleaseTagStatusErrorAsCatalogV1alpha1ReleaseTagStatusReasonsInner is a convenience function that returns CatalogV1alpha1ReleaseTagStatusError wrapped in CatalogV1alpha1ReleaseTagStatusReasonsInner
func CatalogV1alpha1ReleaseTagStatusErrorAsCatalogV1alpha1ReleaseTagStatusReasonsInner(v *CatalogV1alpha1ReleaseTagStatusError) CatalogV1alpha1ReleaseTagStatusReasonsInner {
	return CatalogV1alpha1ReleaseTagStatusReasonsInner{
		CatalogV1alpha1ReleaseTagStatusError: v,
	}
}

// CatalogV1alpha1ReleaseTagStatusPendingAsCatalogV1alpha1ReleaseTagStatusReasonsInner is a convenience function that returns CatalogV1alpha1ReleaseTagStatusPending wrapped in CatalogV1alpha1ReleaseTagStatusReasonsInner
func CatalogV1alpha1ReleaseTagStatusPendingAsCatalogV1alpha1ReleaseTagStatusReasonsInner(v *CatalogV1alpha1ReleaseTagStatusPending) CatalogV1alpha1ReleaseTagStatusReasonsInner {
	return CatalogV1alpha1ReleaseTagStatusReasonsInner{
		CatalogV1alpha1ReleaseTagStatusPending: v,
	}
}

// CatalogV1alpha1ReleaseTagStatusSuccessAsCatalogV1alpha1ReleaseTagStatusReasonsInner is a convenience function that returns CatalogV1alpha1ReleaseTagStatusSuccess wrapped in CatalogV1alpha1ReleaseTagStatusReasonsInner
func CatalogV1alpha1ReleaseTagStatusSuccessAsCatalogV1alpha1ReleaseTagStatusReasonsInner(v *CatalogV1alpha1ReleaseTagStatusSuccess) CatalogV1alpha1ReleaseTagStatusReasonsInner {
	return CatalogV1alpha1ReleaseTagStatusReasonsInner{
		CatalogV1alpha1ReleaseTagStatusSuccess: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *CatalogV1alpha1ReleaseTagStatusReasonsInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CatalogV1alpha1ReleaseTagStatusError
	err = newStrictDecoder(data).Decode(&dst.CatalogV1alpha1ReleaseTagStatusError)
	if err == nil {
		jsonCatalogV1alpha1ReleaseTagStatusError, _ := json.Marshal(dst.CatalogV1alpha1ReleaseTagStatusError)
		if string(jsonCatalogV1alpha1ReleaseTagStatusError) == "{}" { // empty struct
			dst.CatalogV1alpha1ReleaseTagStatusError = nil
		} else {
			if err = validator.Validate(dst.CatalogV1alpha1ReleaseTagStatusError); err != nil {
				dst.CatalogV1alpha1ReleaseTagStatusError = nil
			} else {
				match++
			}
		}
	} else {
		dst.CatalogV1alpha1ReleaseTagStatusError = nil
	}

	// try to unmarshal data into CatalogV1alpha1ReleaseTagStatusPending
	err = newStrictDecoder(data).Decode(&dst.CatalogV1alpha1ReleaseTagStatusPending)
	if err == nil {
		jsonCatalogV1alpha1ReleaseTagStatusPending, _ := json.Marshal(dst.CatalogV1alpha1ReleaseTagStatusPending)
		if string(jsonCatalogV1alpha1ReleaseTagStatusPending) == "{}" { // empty struct
			dst.CatalogV1alpha1ReleaseTagStatusPending = nil
		} else {
			if err = validator.Validate(dst.CatalogV1alpha1ReleaseTagStatusPending); err != nil {
				dst.CatalogV1alpha1ReleaseTagStatusPending = nil
			} else {
				match++
			}
		}
	} else {
		dst.CatalogV1alpha1ReleaseTagStatusPending = nil
	}

	// try to unmarshal data into CatalogV1alpha1ReleaseTagStatusSuccess
	err = newStrictDecoder(data).Decode(&dst.CatalogV1alpha1ReleaseTagStatusSuccess)
	if err == nil {
		jsonCatalogV1alpha1ReleaseTagStatusSuccess, _ := json.Marshal(dst.CatalogV1alpha1ReleaseTagStatusSuccess)
		if string(jsonCatalogV1alpha1ReleaseTagStatusSuccess) == "{}" { // empty struct
			dst.CatalogV1alpha1ReleaseTagStatusSuccess = nil
		} else {
			if err = validator.Validate(dst.CatalogV1alpha1ReleaseTagStatusSuccess); err != nil {
				dst.CatalogV1alpha1ReleaseTagStatusSuccess = nil
			} else {
				match++
			}
		}
	} else {
		dst.CatalogV1alpha1ReleaseTagStatusSuccess = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.CatalogV1alpha1ReleaseTagStatusError = nil
		dst.CatalogV1alpha1ReleaseTagStatusPending = nil
		dst.CatalogV1alpha1ReleaseTagStatusSuccess = nil

		return fmt.Errorf("data matches more than one schema in oneOf(CatalogV1alpha1ReleaseTagStatusReasonsInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(CatalogV1alpha1ReleaseTagStatusReasonsInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CatalogV1alpha1ReleaseTagStatusReasonsInner) MarshalJSON() ([]byte, error) {
	if src.CatalogV1alpha1ReleaseTagStatusError != nil {
		return json.Marshal(&src.CatalogV1alpha1ReleaseTagStatusError)
	}

	if src.CatalogV1alpha1ReleaseTagStatusPending != nil {
		return json.Marshal(&src.CatalogV1alpha1ReleaseTagStatusPending)
	}

	if src.CatalogV1alpha1ReleaseTagStatusSuccess != nil {
		return json.Marshal(&src.CatalogV1alpha1ReleaseTagStatusSuccess)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CatalogV1alpha1ReleaseTagStatusReasonsInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.CatalogV1alpha1ReleaseTagStatusError != nil {
		return obj.CatalogV1alpha1ReleaseTagStatusError
	}

	if obj.CatalogV1alpha1ReleaseTagStatusPending != nil {
		return obj.CatalogV1alpha1ReleaseTagStatusPending
	}

	if obj.CatalogV1alpha1ReleaseTagStatusSuccess != nil {
		return obj.CatalogV1alpha1ReleaseTagStatusSuccess
	}

	// all schemas are nil
	return nil
}

type NullableCatalogV1alpha1ReleaseTagStatusReasonsInner struct {
	value *CatalogV1alpha1ReleaseTagStatusReasonsInner
	isSet bool
}

func (v NullableCatalogV1alpha1ReleaseTagStatusReasonsInner) Get() *CatalogV1alpha1ReleaseTagStatusReasonsInner {
	return v.value
}

func (v *NullableCatalogV1alpha1ReleaseTagStatusReasonsInner) Set(val *CatalogV1alpha1ReleaseTagStatusReasonsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogV1alpha1ReleaseTagStatusReasonsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogV1alpha1ReleaseTagStatusReasonsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogV1alpha1ReleaseTagStatusReasonsInner(val *CatalogV1alpha1ReleaseTagStatusReasonsInner) *NullableCatalogV1alpha1ReleaseTagStatusReasonsInner {
	return &NullableCatalogV1alpha1ReleaseTagStatusReasonsInner{value: val, isSet: true}
}

func (v NullableCatalogV1alpha1ReleaseTagStatusReasonsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogV1alpha1ReleaseTagStatusReasonsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


