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

// CatalogV1alpha1ConsumerProductVisibilitySpecSubjectsInner - struct for CatalogV1alpha1ConsumerProductVisibilitySpecSubjectsInner
type CatalogV1alpha1ConsumerProductVisibilitySpecSubjectsInner struct {
	CatalogV1alpha1ConsumerProductVisibilitySpecAuthenticatedRef *CatalogV1alpha1ConsumerProductVisibilitySpecAuthenticatedRef
	CatalogV1alpha1ConsumerProductVisibilitySpecOrgRef *CatalogV1alpha1ConsumerProductVisibilitySpecOrgRef
}

// CatalogV1alpha1ConsumerProductVisibilitySpecAuthenticatedRefAsCatalogV1alpha1ConsumerProductVisibilitySpecSubjectsInner is a convenience function that returns CatalogV1alpha1ConsumerProductVisibilitySpecAuthenticatedRef wrapped in CatalogV1alpha1ConsumerProductVisibilitySpecSubjectsInner
func CatalogV1alpha1ConsumerProductVisibilitySpecAuthenticatedRefAsCatalogV1alpha1ConsumerProductVisibilitySpecSubjectsInner(v *CatalogV1alpha1ConsumerProductVisibilitySpecAuthenticatedRef) CatalogV1alpha1ConsumerProductVisibilitySpecSubjectsInner {
	return CatalogV1alpha1ConsumerProductVisibilitySpecSubjectsInner{
		CatalogV1alpha1ConsumerProductVisibilitySpecAuthenticatedRef: v,
	}
}

// CatalogV1alpha1ConsumerProductVisibilitySpecOrgRefAsCatalogV1alpha1ConsumerProductVisibilitySpecSubjectsInner is a convenience function that returns CatalogV1alpha1ConsumerProductVisibilitySpecOrgRef wrapped in CatalogV1alpha1ConsumerProductVisibilitySpecSubjectsInner
func CatalogV1alpha1ConsumerProductVisibilitySpecOrgRefAsCatalogV1alpha1ConsumerProductVisibilitySpecSubjectsInner(v *CatalogV1alpha1ConsumerProductVisibilitySpecOrgRef) CatalogV1alpha1ConsumerProductVisibilitySpecSubjectsInner {
	return CatalogV1alpha1ConsumerProductVisibilitySpecSubjectsInner{
		CatalogV1alpha1ConsumerProductVisibilitySpecOrgRef: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *CatalogV1alpha1ConsumerProductVisibilitySpecSubjectsInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CatalogV1alpha1ConsumerProductVisibilitySpecAuthenticatedRef
	err = newStrictDecoder(data).Decode(&dst.CatalogV1alpha1ConsumerProductVisibilitySpecAuthenticatedRef)
	if err == nil {
		jsonCatalogV1alpha1ConsumerProductVisibilitySpecAuthenticatedRef, _ := json.Marshal(dst.CatalogV1alpha1ConsumerProductVisibilitySpecAuthenticatedRef)
		if string(jsonCatalogV1alpha1ConsumerProductVisibilitySpecAuthenticatedRef) == "{}" { // empty struct
			dst.CatalogV1alpha1ConsumerProductVisibilitySpecAuthenticatedRef = nil
		} else {
			if err = validator.Validate(dst.CatalogV1alpha1ConsumerProductVisibilitySpecAuthenticatedRef); err != nil {
				dst.CatalogV1alpha1ConsumerProductVisibilitySpecAuthenticatedRef = nil
			} else {
				match++
			}
		}
	} else {
		dst.CatalogV1alpha1ConsumerProductVisibilitySpecAuthenticatedRef = nil
	}

	// try to unmarshal data into CatalogV1alpha1ConsumerProductVisibilitySpecOrgRef
	err = newStrictDecoder(data).Decode(&dst.CatalogV1alpha1ConsumerProductVisibilitySpecOrgRef)
	if err == nil {
		jsonCatalogV1alpha1ConsumerProductVisibilitySpecOrgRef, _ := json.Marshal(dst.CatalogV1alpha1ConsumerProductVisibilitySpecOrgRef)
		if string(jsonCatalogV1alpha1ConsumerProductVisibilitySpecOrgRef) == "{}" { // empty struct
			dst.CatalogV1alpha1ConsumerProductVisibilitySpecOrgRef = nil
		} else {
			if err = validator.Validate(dst.CatalogV1alpha1ConsumerProductVisibilitySpecOrgRef); err != nil {
				dst.CatalogV1alpha1ConsumerProductVisibilitySpecOrgRef = nil
			} else {
				match++
			}
		}
	} else {
		dst.CatalogV1alpha1ConsumerProductVisibilitySpecOrgRef = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.CatalogV1alpha1ConsumerProductVisibilitySpecAuthenticatedRef = nil
		dst.CatalogV1alpha1ConsumerProductVisibilitySpecOrgRef = nil

		return fmt.Errorf("data matches more than one schema in oneOf(CatalogV1alpha1ConsumerProductVisibilitySpecSubjectsInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(CatalogV1alpha1ConsumerProductVisibilitySpecSubjectsInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CatalogV1alpha1ConsumerProductVisibilitySpecSubjectsInner) MarshalJSON() ([]byte, error) {
	if src.CatalogV1alpha1ConsumerProductVisibilitySpecAuthenticatedRef != nil {
		return json.Marshal(&src.CatalogV1alpha1ConsumerProductVisibilitySpecAuthenticatedRef)
	}

	if src.CatalogV1alpha1ConsumerProductVisibilitySpecOrgRef != nil {
		return json.Marshal(&src.CatalogV1alpha1ConsumerProductVisibilitySpecOrgRef)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CatalogV1alpha1ConsumerProductVisibilitySpecSubjectsInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.CatalogV1alpha1ConsumerProductVisibilitySpecAuthenticatedRef != nil {
		return obj.CatalogV1alpha1ConsumerProductVisibilitySpecAuthenticatedRef
	}

	if obj.CatalogV1alpha1ConsumerProductVisibilitySpecOrgRef != nil {
		return obj.CatalogV1alpha1ConsumerProductVisibilitySpecOrgRef
	}

	// all schemas are nil
	return nil
}

type NullableCatalogV1alpha1ConsumerProductVisibilitySpecSubjectsInner struct {
	value *CatalogV1alpha1ConsumerProductVisibilitySpecSubjectsInner
	isSet bool
}

func (v NullableCatalogV1alpha1ConsumerProductVisibilitySpecSubjectsInner) Get() *CatalogV1alpha1ConsumerProductVisibilitySpecSubjectsInner {
	return v.value
}

func (v *NullableCatalogV1alpha1ConsumerProductVisibilitySpecSubjectsInner) Set(val *CatalogV1alpha1ConsumerProductVisibilitySpecSubjectsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogV1alpha1ConsumerProductVisibilitySpecSubjectsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogV1alpha1ConsumerProductVisibilitySpecSubjectsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogV1alpha1ConsumerProductVisibilitySpecSubjectsInner(val *CatalogV1alpha1ConsumerProductVisibilitySpecSubjectsInner) *NullableCatalogV1alpha1ConsumerProductVisibilitySpecSubjectsInner {
	return &NullableCatalogV1alpha1ConsumerProductVisibilitySpecSubjectsInner{value: val, isSet: true}
}

func (v NullableCatalogV1alpha1ConsumerProductVisibilitySpecSubjectsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogV1alpha1ConsumerProductVisibilitySpecSubjectsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

