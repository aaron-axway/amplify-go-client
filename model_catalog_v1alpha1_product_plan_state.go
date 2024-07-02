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
	"fmt"
)

// CatalogV1alpha1ProductPlanState the model 'CatalogV1alpha1ProductPlanState'
type CatalogV1alpha1ProductPlanState string

// List of catalog.v1alpha1.ProductPlanState
const (
	DRAFT CatalogV1alpha1ProductPlanState = "draft"
	ACTIVE CatalogV1alpha1ProductPlanState = "active"
	DEPRECATED CatalogV1alpha1ProductPlanState = "deprecated"
	ARCHIVED CatalogV1alpha1ProductPlanState = "archived"
)

// All allowed values of CatalogV1alpha1ProductPlanState enum
var AllowedCatalogV1alpha1ProductPlanStateEnumValues = []CatalogV1alpha1ProductPlanState{
	"draft",
	"active",
	"deprecated",
	"archived",
}

func (v *CatalogV1alpha1ProductPlanState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CatalogV1alpha1ProductPlanState(value)
	for _, existing := range AllowedCatalogV1alpha1ProductPlanStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CatalogV1alpha1ProductPlanState", value)
}

// NewCatalogV1alpha1ProductPlanStateFromValue returns a pointer to a valid CatalogV1alpha1ProductPlanState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCatalogV1alpha1ProductPlanStateFromValue(v string) (*CatalogV1alpha1ProductPlanState, error) {
	ev := CatalogV1alpha1ProductPlanState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CatalogV1alpha1ProductPlanState: valid values are %v", v, AllowedCatalogV1alpha1ProductPlanStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CatalogV1alpha1ProductPlanState) IsValid() bool {
	for _, existing := range AllowedCatalogV1alpha1ProductPlanStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to catalog.v1alpha1.ProductPlanState value
func (v CatalogV1alpha1ProductPlanState) Ptr() *CatalogV1alpha1ProductPlanState {
	return &v
}

type NullableCatalogV1alpha1ProductPlanState struct {
	value *CatalogV1alpha1ProductPlanState
	isSet bool
}

func (v NullableCatalogV1alpha1ProductPlanState) Get() *CatalogV1alpha1ProductPlanState {
	return v.value
}

func (v *NullableCatalogV1alpha1ProductPlanState) Set(val *CatalogV1alpha1ProductPlanState) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogV1alpha1ProductPlanState) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogV1alpha1ProductPlanState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogV1alpha1ProductPlanState(val *CatalogV1alpha1ProductPlanState) *NullableCatalogV1alpha1ProductPlanState {
	return &NullableCatalogV1alpha1ProductPlanState{value: val, isSet: true}
}

func (v NullableCatalogV1alpha1ProductPlanState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogV1alpha1ProductPlanState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

