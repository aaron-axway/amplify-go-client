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
	"bytes"
	"fmt"
)

// checks if the CatalogV1alpha1AssetRequestApproval type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogV1alpha1AssetRequestApproval{}

// CatalogV1alpha1AssetRequestApproval struct for CatalogV1alpha1AssetRequestApproval
type CatalogV1alpha1AssetRequestApproval struct {
	State CatalogV1alpha1AssetRequestApprovalState `json:"state"`
}

type _CatalogV1alpha1AssetRequestApproval CatalogV1alpha1AssetRequestApproval

// NewCatalogV1alpha1AssetRequestApproval instantiates a new CatalogV1alpha1AssetRequestApproval object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogV1alpha1AssetRequestApproval(state CatalogV1alpha1AssetRequestApprovalState) *CatalogV1alpha1AssetRequestApproval {
	this := CatalogV1alpha1AssetRequestApproval{}
	this.State = state
	return &this
}

// NewCatalogV1alpha1AssetRequestApprovalWithDefaults instantiates a new CatalogV1alpha1AssetRequestApproval object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogV1alpha1AssetRequestApprovalWithDefaults() *CatalogV1alpha1AssetRequestApproval {
	this := CatalogV1alpha1AssetRequestApproval{}
	return &this
}

// GetState returns the State field value
func (o *CatalogV1alpha1AssetRequestApproval) GetState() CatalogV1alpha1AssetRequestApprovalState {
	if o == nil {
		var ret CatalogV1alpha1AssetRequestApprovalState
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1AssetRequestApproval) GetStateOk() (*CatalogV1alpha1AssetRequestApprovalState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *CatalogV1alpha1AssetRequestApproval) SetState(v CatalogV1alpha1AssetRequestApprovalState) {
	o.State = v
}

func (o CatalogV1alpha1AssetRequestApproval) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogV1alpha1AssetRequestApproval) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["state"] = o.State
	return toSerialize, nil
}

func (o *CatalogV1alpha1AssetRequestApproval) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"state",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCatalogV1alpha1AssetRequestApproval := _CatalogV1alpha1AssetRequestApproval{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCatalogV1alpha1AssetRequestApproval)

	if err != nil {
		return err
	}

	*o = CatalogV1alpha1AssetRequestApproval(varCatalogV1alpha1AssetRequestApproval)

	return err
}

type NullableCatalogV1alpha1AssetRequestApproval struct {
	value *CatalogV1alpha1AssetRequestApproval
	isSet bool
}

func (v NullableCatalogV1alpha1AssetRequestApproval) Get() *CatalogV1alpha1AssetRequestApproval {
	return v.value
}

func (v *NullableCatalogV1alpha1AssetRequestApproval) Set(val *CatalogV1alpha1AssetRequestApproval) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogV1alpha1AssetRequestApproval) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogV1alpha1AssetRequestApproval) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogV1alpha1AssetRequestApproval(val *CatalogV1alpha1AssetRequestApproval) *NullableCatalogV1alpha1AssetRequestApproval {
	return &NullableCatalogV1alpha1AssetRequestApproval{value: val, isSet: true}
}

func (v NullableCatalogV1alpha1AssetRequestApproval) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogV1alpha1AssetRequestApproval) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


