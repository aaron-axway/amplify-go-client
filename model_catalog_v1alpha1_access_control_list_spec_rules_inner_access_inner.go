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

// CatalogV1alpha1AccessControlListSpecRulesInnerAccessInner - struct for CatalogV1alpha1AccessControlListSpecRulesInnerAccessInner
type CatalogV1alpha1AccessControlListSpecRulesInnerAccessInner struct {
	CatalogV1alpha1AccessControlListSpecAccessLevelScope *CatalogV1alpha1AccessControlListSpecAccessLevelScope
	CatalogV1alpha1AccessControlListSpecAccessLevelScopedKind *CatalogV1alpha1AccessControlListSpecAccessLevelScopedKind
	CatalogV1alpha1AccessControlListSpecAccessLevelScopedResource *CatalogV1alpha1AccessControlListSpecAccessLevelScopedResource
}

// CatalogV1alpha1AccessControlListSpecAccessLevelScopeAsCatalogV1alpha1AccessControlListSpecRulesInnerAccessInner is a convenience function that returns CatalogV1alpha1AccessControlListSpecAccessLevelScope wrapped in CatalogV1alpha1AccessControlListSpecRulesInnerAccessInner
func CatalogV1alpha1AccessControlListSpecAccessLevelScopeAsCatalogV1alpha1AccessControlListSpecRulesInnerAccessInner(v *CatalogV1alpha1AccessControlListSpecAccessLevelScope) CatalogV1alpha1AccessControlListSpecRulesInnerAccessInner {
	return CatalogV1alpha1AccessControlListSpecRulesInnerAccessInner{
		CatalogV1alpha1AccessControlListSpecAccessLevelScope: v,
	}
}

// CatalogV1alpha1AccessControlListSpecAccessLevelScopedKindAsCatalogV1alpha1AccessControlListSpecRulesInnerAccessInner is a convenience function that returns CatalogV1alpha1AccessControlListSpecAccessLevelScopedKind wrapped in CatalogV1alpha1AccessControlListSpecRulesInnerAccessInner
func CatalogV1alpha1AccessControlListSpecAccessLevelScopedKindAsCatalogV1alpha1AccessControlListSpecRulesInnerAccessInner(v *CatalogV1alpha1AccessControlListSpecAccessLevelScopedKind) CatalogV1alpha1AccessControlListSpecRulesInnerAccessInner {
	return CatalogV1alpha1AccessControlListSpecRulesInnerAccessInner{
		CatalogV1alpha1AccessControlListSpecAccessLevelScopedKind: v,
	}
}

// CatalogV1alpha1AccessControlListSpecAccessLevelScopedResourceAsCatalogV1alpha1AccessControlListSpecRulesInnerAccessInner is a convenience function that returns CatalogV1alpha1AccessControlListSpecAccessLevelScopedResource wrapped in CatalogV1alpha1AccessControlListSpecRulesInnerAccessInner
func CatalogV1alpha1AccessControlListSpecAccessLevelScopedResourceAsCatalogV1alpha1AccessControlListSpecRulesInnerAccessInner(v *CatalogV1alpha1AccessControlListSpecAccessLevelScopedResource) CatalogV1alpha1AccessControlListSpecRulesInnerAccessInner {
	return CatalogV1alpha1AccessControlListSpecRulesInnerAccessInner{
		CatalogV1alpha1AccessControlListSpecAccessLevelScopedResource: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *CatalogV1alpha1AccessControlListSpecRulesInnerAccessInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CatalogV1alpha1AccessControlListSpecAccessLevelScope
	err = newStrictDecoder(data).Decode(&dst.CatalogV1alpha1AccessControlListSpecAccessLevelScope)
	if err == nil {
		jsonCatalogV1alpha1AccessControlListSpecAccessLevelScope, _ := json.Marshal(dst.CatalogV1alpha1AccessControlListSpecAccessLevelScope)
		if string(jsonCatalogV1alpha1AccessControlListSpecAccessLevelScope) == "{}" { // empty struct
			dst.CatalogV1alpha1AccessControlListSpecAccessLevelScope = nil
		} else {
			if err = validator.Validate(dst.CatalogV1alpha1AccessControlListSpecAccessLevelScope); err != nil {
				dst.CatalogV1alpha1AccessControlListSpecAccessLevelScope = nil
			} else {
				match++
			}
		}
	} else {
		dst.CatalogV1alpha1AccessControlListSpecAccessLevelScope = nil
	}

	// try to unmarshal data into CatalogV1alpha1AccessControlListSpecAccessLevelScopedKind
	err = newStrictDecoder(data).Decode(&dst.CatalogV1alpha1AccessControlListSpecAccessLevelScopedKind)
	if err == nil {
		jsonCatalogV1alpha1AccessControlListSpecAccessLevelScopedKind, _ := json.Marshal(dst.CatalogV1alpha1AccessControlListSpecAccessLevelScopedKind)
		if string(jsonCatalogV1alpha1AccessControlListSpecAccessLevelScopedKind) == "{}" { // empty struct
			dst.CatalogV1alpha1AccessControlListSpecAccessLevelScopedKind = nil
		} else {
			if err = validator.Validate(dst.CatalogV1alpha1AccessControlListSpecAccessLevelScopedKind); err != nil {
				dst.CatalogV1alpha1AccessControlListSpecAccessLevelScopedKind = nil
			} else {
				match++
			}
		}
	} else {
		dst.CatalogV1alpha1AccessControlListSpecAccessLevelScopedKind = nil
	}

	// try to unmarshal data into CatalogV1alpha1AccessControlListSpecAccessLevelScopedResource
	err = newStrictDecoder(data).Decode(&dst.CatalogV1alpha1AccessControlListSpecAccessLevelScopedResource)
	if err == nil {
		jsonCatalogV1alpha1AccessControlListSpecAccessLevelScopedResource, _ := json.Marshal(dst.CatalogV1alpha1AccessControlListSpecAccessLevelScopedResource)
		if string(jsonCatalogV1alpha1AccessControlListSpecAccessLevelScopedResource) == "{}" { // empty struct
			dst.CatalogV1alpha1AccessControlListSpecAccessLevelScopedResource = nil
		} else {
			if err = validator.Validate(dst.CatalogV1alpha1AccessControlListSpecAccessLevelScopedResource); err != nil {
				dst.CatalogV1alpha1AccessControlListSpecAccessLevelScopedResource = nil
			} else {
				match++
			}
		}
	} else {
		dst.CatalogV1alpha1AccessControlListSpecAccessLevelScopedResource = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.CatalogV1alpha1AccessControlListSpecAccessLevelScope = nil
		dst.CatalogV1alpha1AccessControlListSpecAccessLevelScopedKind = nil
		dst.CatalogV1alpha1AccessControlListSpecAccessLevelScopedResource = nil

		return fmt.Errorf("data matches more than one schema in oneOf(CatalogV1alpha1AccessControlListSpecRulesInnerAccessInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(CatalogV1alpha1AccessControlListSpecRulesInnerAccessInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CatalogV1alpha1AccessControlListSpecRulesInnerAccessInner) MarshalJSON() ([]byte, error) {
	if src.CatalogV1alpha1AccessControlListSpecAccessLevelScope != nil {
		return json.Marshal(&src.CatalogV1alpha1AccessControlListSpecAccessLevelScope)
	}

	if src.CatalogV1alpha1AccessControlListSpecAccessLevelScopedKind != nil {
		return json.Marshal(&src.CatalogV1alpha1AccessControlListSpecAccessLevelScopedKind)
	}

	if src.CatalogV1alpha1AccessControlListSpecAccessLevelScopedResource != nil {
		return json.Marshal(&src.CatalogV1alpha1AccessControlListSpecAccessLevelScopedResource)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CatalogV1alpha1AccessControlListSpecRulesInnerAccessInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.CatalogV1alpha1AccessControlListSpecAccessLevelScope != nil {
		return obj.CatalogV1alpha1AccessControlListSpecAccessLevelScope
	}

	if obj.CatalogV1alpha1AccessControlListSpecAccessLevelScopedKind != nil {
		return obj.CatalogV1alpha1AccessControlListSpecAccessLevelScopedKind
	}

	if obj.CatalogV1alpha1AccessControlListSpecAccessLevelScopedResource != nil {
		return obj.CatalogV1alpha1AccessControlListSpecAccessLevelScopedResource
	}

	// all schemas are nil
	return nil
}

type NullableCatalogV1alpha1AccessControlListSpecRulesInnerAccessInner struct {
	value *CatalogV1alpha1AccessControlListSpecRulesInnerAccessInner
	isSet bool
}

func (v NullableCatalogV1alpha1AccessControlListSpecRulesInnerAccessInner) Get() *CatalogV1alpha1AccessControlListSpecRulesInnerAccessInner {
	return v.value
}

func (v *NullableCatalogV1alpha1AccessControlListSpecRulesInnerAccessInner) Set(val *CatalogV1alpha1AccessControlListSpecRulesInnerAccessInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogV1alpha1AccessControlListSpecRulesInnerAccessInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogV1alpha1AccessControlListSpecRulesInnerAccessInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogV1alpha1AccessControlListSpecRulesInnerAccessInner(val *CatalogV1alpha1AccessControlListSpecRulesInnerAccessInner) *NullableCatalogV1alpha1AccessControlListSpecRulesInnerAccessInner {
	return &NullableCatalogV1alpha1AccessControlListSpecRulesInnerAccessInner{value: val, isSet: true}
}

func (v NullableCatalogV1alpha1AccessControlListSpecRulesInnerAccessInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogV1alpha1AccessControlListSpecRulesInnerAccessInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


