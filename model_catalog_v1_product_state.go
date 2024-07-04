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
	"fmt"
)

// CatalogV1ProductState the model 'CatalogV1ProductState'
type CatalogV1ProductState string

// List of catalog.v1.ProductState
const (
	DRAFT CatalogV1ProductState = "draft"
	ACTIVE CatalogV1ProductState = "active"
	DEPRECATED CatalogV1ProductState = "deprecated"
	ARCHIVED CatalogV1ProductState = "archived"
)

// All allowed values of CatalogV1ProductState enum
var AllowedCatalogV1ProductStateEnumValues = []CatalogV1ProductState{
	"draft",
	"active",
	"deprecated",
	"archived",
}

func (v *CatalogV1ProductState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CatalogV1ProductState(value)
	for _, existing := range AllowedCatalogV1ProductStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CatalogV1ProductState", value)
}

// NewCatalogV1ProductStateFromValue returns a pointer to a valid CatalogV1ProductState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCatalogV1ProductStateFromValue(v string) (*CatalogV1ProductState, error) {
	ev := CatalogV1ProductState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CatalogV1ProductState: valid values are %v", v, AllowedCatalogV1ProductStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CatalogV1ProductState) IsValid() bool {
	for _, existing := range AllowedCatalogV1ProductStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to catalog.v1.ProductState value
func (v CatalogV1ProductState) Ptr() *CatalogV1ProductState {
	return &v
}

type NullableCatalogV1ProductState struct {
	value *CatalogV1ProductState
	isSet bool
}

func (v NullableCatalogV1ProductState) Get() *CatalogV1ProductState {
	return v.value
}

func (v *NullableCatalogV1ProductState) Set(val *CatalogV1ProductState) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogV1ProductState) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogV1ProductState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogV1ProductState(val *CatalogV1ProductState) *NullableCatalogV1ProductState {
	return &NullableCatalogV1ProductState{value: val, isSet: true}
}

func (v NullableCatalogV1ProductState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogV1ProductState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

