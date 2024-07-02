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

// CatalogV1alpha1ProductPlanJobSpecAction - struct for CatalogV1alpha1ProductPlanJobSpecAction
type CatalogV1alpha1ProductPlanJobSpecAction struct {
	CatalogV1alpha1ProductPlanJobSpecSubscriptionsArchive *CatalogV1alpha1ProductPlanJobSpecSubscriptionsArchive
	CatalogV1alpha1ProductPlanJobSpecSubscriptionsMigration *CatalogV1alpha1ProductPlanJobSpecSubscriptionsMigration
}

// CatalogV1alpha1ProductPlanJobSpecSubscriptionsArchiveAsCatalogV1alpha1ProductPlanJobSpecAction is a convenience function that returns CatalogV1alpha1ProductPlanJobSpecSubscriptionsArchive wrapped in CatalogV1alpha1ProductPlanJobSpecAction
func CatalogV1alpha1ProductPlanJobSpecSubscriptionsArchiveAsCatalogV1alpha1ProductPlanJobSpecAction(v *CatalogV1alpha1ProductPlanJobSpecSubscriptionsArchive) CatalogV1alpha1ProductPlanJobSpecAction {
	return CatalogV1alpha1ProductPlanJobSpecAction{
		CatalogV1alpha1ProductPlanJobSpecSubscriptionsArchive: v,
	}
}

// CatalogV1alpha1ProductPlanJobSpecSubscriptionsMigrationAsCatalogV1alpha1ProductPlanJobSpecAction is a convenience function that returns CatalogV1alpha1ProductPlanJobSpecSubscriptionsMigration wrapped in CatalogV1alpha1ProductPlanJobSpecAction
func CatalogV1alpha1ProductPlanJobSpecSubscriptionsMigrationAsCatalogV1alpha1ProductPlanJobSpecAction(v *CatalogV1alpha1ProductPlanJobSpecSubscriptionsMigration) CatalogV1alpha1ProductPlanJobSpecAction {
	return CatalogV1alpha1ProductPlanJobSpecAction{
		CatalogV1alpha1ProductPlanJobSpecSubscriptionsMigration: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *CatalogV1alpha1ProductPlanJobSpecAction) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CatalogV1alpha1ProductPlanJobSpecSubscriptionsArchive
	err = newStrictDecoder(data).Decode(&dst.CatalogV1alpha1ProductPlanJobSpecSubscriptionsArchive)
	if err == nil {
		jsonCatalogV1alpha1ProductPlanJobSpecSubscriptionsArchive, _ := json.Marshal(dst.CatalogV1alpha1ProductPlanJobSpecSubscriptionsArchive)
		if string(jsonCatalogV1alpha1ProductPlanJobSpecSubscriptionsArchive) == "{}" { // empty struct
			dst.CatalogV1alpha1ProductPlanJobSpecSubscriptionsArchive = nil
		} else {
			if err = validator.Validate(dst.CatalogV1alpha1ProductPlanJobSpecSubscriptionsArchive); err != nil {
				dst.CatalogV1alpha1ProductPlanJobSpecSubscriptionsArchive = nil
			} else {
				match++
			}
		}
	} else {
		dst.CatalogV1alpha1ProductPlanJobSpecSubscriptionsArchive = nil
	}

	// try to unmarshal data into CatalogV1alpha1ProductPlanJobSpecSubscriptionsMigration
	err = newStrictDecoder(data).Decode(&dst.CatalogV1alpha1ProductPlanJobSpecSubscriptionsMigration)
	if err == nil {
		jsonCatalogV1alpha1ProductPlanJobSpecSubscriptionsMigration, _ := json.Marshal(dst.CatalogV1alpha1ProductPlanJobSpecSubscriptionsMigration)
		if string(jsonCatalogV1alpha1ProductPlanJobSpecSubscriptionsMigration) == "{}" { // empty struct
			dst.CatalogV1alpha1ProductPlanJobSpecSubscriptionsMigration = nil
		} else {
			if err = validator.Validate(dst.CatalogV1alpha1ProductPlanJobSpecSubscriptionsMigration); err != nil {
				dst.CatalogV1alpha1ProductPlanJobSpecSubscriptionsMigration = nil
			} else {
				match++
			}
		}
	} else {
		dst.CatalogV1alpha1ProductPlanJobSpecSubscriptionsMigration = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.CatalogV1alpha1ProductPlanJobSpecSubscriptionsArchive = nil
		dst.CatalogV1alpha1ProductPlanJobSpecSubscriptionsMigration = nil

		return fmt.Errorf("data matches more than one schema in oneOf(CatalogV1alpha1ProductPlanJobSpecAction)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(CatalogV1alpha1ProductPlanJobSpecAction)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CatalogV1alpha1ProductPlanJobSpecAction) MarshalJSON() ([]byte, error) {
	if src.CatalogV1alpha1ProductPlanJobSpecSubscriptionsArchive != nil {
		return json.Marshal(&src.CatalogV1alpha1ProductPlanJobSpecSubscriptionsArchive)
	}

	if src.CatalogV1alpha1ProductPlanJobSpecSubscriptionsMigration != nil {
		return json.Marshal(&src.CatalogV1alpha1ProductPlanJobSpecSubscriptionsMigration)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CatalogV1alpha1ProductPlanJobSpecAction) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.CatalogV1alpha1ProductPlanJobSpecSubscriptionsArchive != nil {
		return obj.CatalogV1alpha1ProductPlanJobSpecSubscriptionsArchive
	}

	if obj.CatalogV1alpha1ProductPlanJobSpecSubscriptionsMigration != nil {
		return obj.CatalogV1alpha1ProductPlanJobSpecSubscriptionsMigration
	}

	// all schemas are nil
	return nil
}

type NullableCatalogV1alpha1ProductPlanJobSpecAction struct {
	value *CatalogV1alpha1ProductPlanJobSpecAction
	isSet bool
}

func (v NullableCatalogV1alpha1ProductPlanJobSpecAction) Get() *CatalogV1alpha1ProductPlanJobSpecAction {
	return v.value
}

func (v *NullableCatalogV1alpha1ProductPlanJobSpecAction) Set(val *CatalogV1alpha1ProductPlanJobSpecAction) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogV1alpha1ProductPlanJobSpecAction) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogV1alpha1ProductPlanJobSpecAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogV1alpha1ProductPlanJobSpecAction(val *CatalogV1alpha1ProductPlanJobSpecAction) *NullableCatalogV1alpha1ProductPlanJobSpecAction {
	return &NullableCatalogV1alpha1ProductPlanJobSpecAction{value: val, isSet: true}
}

func (v NullableCatalogV1alpha1ProductPlanJobSpecAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogV1alpha1ProductPlanJobSpecAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


