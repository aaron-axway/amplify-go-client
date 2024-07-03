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

// CatalogV1alpha1AssetState the model 'CatalogV1alpha1AssetState'
type CatalogV1alpha1AssetState string

// List of catalog.v1alpha1.AssetState
const (
	DRAFTAssetState CatalogV1alpha1AssetState = "draft"
	ACTIVEAssetState CatalogV1alpha1AssetState = "active"
	DEPRECATEDAssetState CatalogV1alpha1AssetState = "deprecated"
	ARCHIVEDAssetState CatalogV1alpha1AssetState = "archived"
)

// All allowed values of CatalogV1alpha1AssetState enum
var AllowedCatalogV1alpha1AssetStateEnumValues = []CatalogV1alpha1AssetState{
	"draft",
	"active",
	"deprecated",
	"archived",
}

func (v *CatalogV1alpha1AssetState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CatalogV1alpha1AssetState(value)
	for _, existing := range AllowedCatalogV1alpha1AssetStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CatalogV1alpha1AssetState", value)
}

// NewCatalogV1alpha1AssetStateFromValue returns a pointer to a valid CatalogV1alpha1AssetState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCatalogV1alpha1AssetStateFromValue(v string) (*CatalogV1alpha1AssetState, error) {
	ev := CatalogV1alpha1AssetState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CatalogV1alpha1AssetState: valid values are %v", v, AllowedCatalogV1alpha1AssetStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CatalogV1alpha1AssetState) IsValid() bool {
	for _, existing := range AllowedCatalogV1alpha1AssetStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to catalog.v1alpha1.AssetState value
func (v CatalogV1alpha1AssetState) Ptr() *CatalogV1alpha1AssetState {
	return &v
}

type NullableCatalogV1alpha1AssetState struct {
	value *CatalogV1alpha1AssetState
	isSet bool
}

func (v NullableCatalogV1alpha1AssetState) Get() *CatalogV1alpha1AssetState {
	return v.value
}

func (v *NullableCatalogV1alpha1AssetState) Set(val *CatalogV1alpha1AssetState) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogV1alpha1AssetState) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogV1alpha1AssetState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogV1alpha1AssetState(val *CatalogV1alpha1AssetState) *NullableCatalogV1alpha1AssetState {
	return &NullableCatalogV1alpha1AssetState{value: val, isSet: true}
}

func (v NullableCatalogV1alpha1AssetState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogV1alpha1AssetState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

