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

// ManagementV1alpha1AccessControlListSpecRulesInnerAccessInner - struct for ManagementV1alpha1AccessControlListSpecRulesInnerAccessInner
type ManagementV1alpha1AccessControlListSpecRulesInnerAccessInner struct {
	ManagementV1alpha1AccessControlListSpecAccessLevelScope *ManagementV1alpha1AccessControlListSpecAccessLevelScope
	ManagementV1alpha1AccessControlListSpecAccessLevelScopedKind *ManagementV1alpha1AccessControlListSpecAccessLevelScopedKind
	ManagementV1alpha1AccessControlListSpecAccessLevelScopedResource *ManagementV1alpha1AccessControlListSpecAccessLevelScopedResource
}

// ManagementV1alpha1AccessControlListSpecAccessLevelScopeAsManagementV1alpha1AccessControlListSpecRulesInnerAccessInner is a convenience function that returns ManagementV1alpha1AccessControlListSpecAccessLevelScope wrapped in ManagementV1alpha1AccessControlListSpecRulesInnerAccessInner
func ManagementV1alpha1AccessControlListSpecAccessLevelScopeAsManagementV1alpha1AccessControlListSpecRulesInnerAccessInner(v *ManagementV1alpha1AccessControlListSpecAccessLevelScope) ManagementV1alpha1AccessControlListSpecRulesInnerAccessInner {
	return ManagementV1alpha1AccessControlListSpecRulesInnerAccessInner{
		ManagementV1alpha1AccessControlListSpecAccessLevelScope: v,
	}
}

// ManagementV1alpha1AccessControlListSpecAccessLevelScopedKindAsManagementV1alpha1AccessControlListSpecRulesInnerAccessInner is a convenience function that returns ManagementV1alpha1AccessControlListSpecAccessLevelScopedKind wrapped in ManagementV1alpha1AccessControlListSpecRulesInnerAccessInner
func ManagementV1alpha1AccessControlListSpecAccessLevelScopedKindAsManagementV1alpha1AccessControlListSpecRulesInnerAccessInner(v *ManagementV1alpha1AccessControlListSpecAccessLevelScopedKind) ManagementV1alpha1AccessControlListSpecRulesInnerAccessInner {
	return ManagementV1alpha1AccessControlListSpecRulesInnerAccessInner{
		ManagementV1alpha1AccessControlListSpecAccessLevelScopedKind: v,
	}
}

// ManagementV1alpha1AccessControlListSpecAccessLevelScopedResourceAsManagementV1alpha1AccessControlListSpecRulesInnerAccessInner is a convenience function that returns ManagementV1alpha1AccessControlListSpecAccessLevelScopedResource wrapped in ManagementV1alpha1AccessControlListSpecRulesInnerAccessInner
func ManagementV1alpha1AccessControlListSpecAccessLevelScopedResourceAsManagementV1alpha1AccessControlListSpecRulesInnerAccessInner(v *ManagementV1alpha1AccessControlListSpecAccessLevelScopedResource) ManagementV1alpha1AccessControlListSpecRulesInnerAccessInner {
	return ManagementV1alpha1AccessControlListSpecRulesInnerAccessInner{
		ManagementV1alpha1AccessControlListSpecAccessLevelScopedResource: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ManagementV1alpha1AccessControlListSpecRulesInnerAccessInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ManagementV1alpha1AccessControlListSpecAccessLevelScope
	err = newStrictDecoder(data).Decode(&dst.ManagementV1alpha1AccessControlListSpecAccessLevelScope)
	if err == nil {
		jsonManagementV1alpha1AccessControlListSpecAccessLevelScope, _ := json.Marshal(dst.ManagementV1alpha1AccessControlListSpecAccessLevelScope)
		if string(jsonManagementV1alpha1AccessControlListSpecAccessLevelScope) == "{}" { // empty struct
			dst.ManagementV1alpha1AccessControlListSpecAccessLevelScope = nil
		} else {
			if err = validator.Validate(dst.ManagementV1alpha1AccessControlListSpecAccessLevelScope); err != nil {
				dst.ManagementV1alpha1AccessControlListSpecAccessLevelScope = nil
			} else {
				match++
			}
		}
	} else {
		dst.ManagementV1alpha1AccessControlListSpecAccessLevelScope = nil
	}

	// try to unmarshal data into ManagementV1alpha1AccessControlListSpecAccessLevelScopedKind
	err = newStrictDecoder(data).Decode(&dst.ManagementV1alpha1AccessControlListSpecAccessLevelScopedKind)
	if err == nil {
		jsonManagementV1alpha1AccessControlListSpecAccessLevelScopedKind, _ := json.Marshal(dst.ManagementV1alpha1AccessControlListSpecAccessLevelScopedKind)
		if string(jsonManagementV1alpha1AccessControlListSpecAccessLevelScopedKind) == "{}" { // empty struct
			dst.ManagementV1alpha1AccessControlListSpecAccessLevelScopedKind = nil
		} else {
			if err = validator.Validate(dst.ManagementV1alpha1AccessControlListSpecAccessLevelScopedKind); err != nil {
				dst.ManagementV1alpha1AccessControlListSpecAccessLevelScopedKind = nil
			} else {
				match++
			}
		}
	} else {
		dst.ManagementV1alpha1AccessControlListSpecAccessLevelScopedKind = nil
	}

	// try to unmarshal data into ManagementV1alpha1AccessControlListSpecAccessLevelScopedResource
	err = newStrictDecoder(data).Decode(&dst.ManagementV1alpha1AccessControlListSpecAccessLevelScopedResource)
	if err == nil {
		jsonManagementV1alpha1AccessControlListSpecAccessLevelScopedResource, _ := json.Marshal(dst.ManagementV1alpha1AccessControlListSpecAccessLevelScopedResource)
		if string(jsonManagementV1alpha1AccessControlListSpecAccessLevelScopedResource) == "{}" { // empty struct
			dst.ManagementV1alpha1AccessControlListSpecAccessLevelScopedResource = nil
		} else {
			if err = validator.Validate(dst.ManagementV1alpha1AccessControlListSpecAccessLevelScopedResource); err != nil {
				dst.ManagementV1alpha1AccessControlListSpecAccessLevelScopedResource = nil
			} else {
				match++
			}
		}
	} else {
		dst.ManagementV1alpha1AccessControlListSpecAccessLevelScopedResource = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ManagementV1alpha1AccessControlListSpecAccessLevelScope = nil
		dst.ManagementV1alpha1AccessControlListSpecAccessLevelScopedKind = nil
		dst.ManagementV1alpha1AccessControlListSpecAccessLevelScopedResource = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ManagementV1alpha1AccessControlListSpecRulesInnerAccessInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ManagementV1alpha1AccessControlListSpecRulesInnerAccessInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ManagementV1alpha1AccessControlListSpecRulesInnerAccessInner) MarshalJSON() ([]byte, error) {
	if src.ManagementV1alpha1AccessControlListSpecAccessLevelScope != nil {
		return json.Marshal(&src.ManagementV1alpha1AccessControlListSpecAccessLevelScope)
	}

	if src.ManagementV1alpha1AccessControlListSpecAccessLevelScopedKind != nil {
		return json.Marshal(&src.ManagementV1alpha1AccessControlListSpecAccessLevelScopedKind)
	}

	if src.ManagementV1alpha1AccessControlListSpecAccessLevelScopedResource != nil {
		return json.Marshal(&src.ManagementV1alpha1AccessControlListSpecAccessLevelScopedResource)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ManagementV1alpha1AccessControlListSpecRulesInnerAccessInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ManagementV1alpha1AccessControlListSpecAccessLevelScope != nil {
		return obj.ManagementV1alpha1AccessControlListSpecAccessLevelScope
	}

	if obj.ManagementV1alpha1AccessControlListSpecAccessLevelScopedKind != nil {
		return obj.ManagementV1alpha1AccessControlListSpecAccessLevelScopedKind
	}

	if obj.ManagementV1alpha1AccessControlListSpecAccessLevelScopedResource != nil {
		return obj.ManagementV1alpha1AccessControlListSpecAccessLevelScopedResource
	}

	// all schemas are nil
	return nil
}

type NullableManagementV1alpha1AccessControlListSpecRulesInnerAccessInner struct {
	value *ManagementV1alpha1AccessControlListSpecRulesInnerAccessInner
	isSet bool
}

func (v NullableManagementV1alpha1AccessControlListSpecRulesInnerAccessInner) Get() *ManagementV1alpha1AccessControlListSpecRulesInnerAccessInner {
	return v.value
}

func (v *NullableManagementV1alpha1AccessControlListSpecRulesInnerAccessInner) Set(val *ManagementV1alpha1AccessControlListSpecRulesInnerAccessInner) {
	v.value = val
	v.isSet = true
}

func (v NullableManagementV1alpha1AccessControlListSpecRulesInnerAccessInner) IsSet() bool {
	return v.isSet
}

func (v *NullableManagementV1alpha1AccessControlListSpecRulesInnerAccessInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagementV1alpha1AccessControlListSpecRulesInnerAccessInner(val *ManagementV1alpha1AccessControlListSpecRulesInnerAccessInner) *NullableManagementV1alpha1AccessControlListSpecRulesInnerAccessInner {
	return &NullableManagementV1alpha1AccessControlListSpecRulesInnerAccessInner{value: val, isSet: true}
}

func (v NullableManagementV1alpha1AccessControlListSpecRulesInnerAccessInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagementV1alpha1AccessControlListSpecRulesInnerAccessInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


