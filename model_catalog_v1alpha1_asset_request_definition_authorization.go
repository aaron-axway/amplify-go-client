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

// checks if the CatalogV1alpha1AssetRequestDefinitionAuthorization type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogV1alpha1AssetRequestDefinitionAuthorization{}

// CatalogV1alpha1AssetRequestDefinitionAuthorization struct for CatalogV1alpha1AssetRequestDefinitionAuthorization
type CatalogV1alpha1AssetRequestDefinitionAuthorization struct {
	Approval string `json:"approval"`
}

type _CatalogV1alpha1AssetRequestDefinitionAuthorization CatalogV1alpha1AssetRequestDefinitionAuthorization

// NewCatalogV1alpha1AssetRequestDefinitionAuthorization instantiates a new CatalogV1alpha1AssetRequestDefinitionAuthorization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogV1alpha1AssetRequestDefinitionAuthorization(approval string) *CatalogV1alpha1AssetRequestDefinitionAuthorization {
	this := CatalogV1alpha1AssetRequestDefinitionAuthorization{}
	this.Approval = approval
	return &this
}

// NewCatalogV1alpha1AssetRequestDefinitionAuthorizationWithDefaults instantiates a new CatalogV1alpha1AssetRequestDefinitionAuthorization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogV1alpha1AssetRequestDefinitionAuthorizationWithDefaults() *CatalogV1alpha1AssetRequestDefinitionAuthorization {
	this := CatalogV1alpha1AssetRequestDefinitionAuthorization{}
	return &this
}

// GetApproval returns the Approval field value
func (o *CatalogV1alpha1AssetRequestDefinitionAuthorization) GetApproval() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Approval
}

// GetApprovalOk returns a tuple with the Approval field value
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1AssetRequestDefinitionAuthorization) GetApprovalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Approval, true
}

// SetApproval sets field value
func (o *CatalogV1alpha1AssetRequestDefinitionAuthorization) SetApproval(v string) {
	o.Approval = v
}

func (o CatalogV1alpha1AssetRequestDefinitionAuthorization) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogV1alpha1AssetRequestDefinitionAuthorization) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["approval"] = o.Approval
	return toSerialize, nil
}

func (o *CatalogV1alpha1AssetRequestDefinitionAuthorization) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"approval",
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

	varCatalogV1alpha1AssetRequestDefinitionAuthorization := _CatalogV1alpha1AssetRequestDefinitionAuthorization{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCatalogV1alpha1AssetRequestDefinitionAuthorization)

	if err != nil {
		return err
	}

	*o = CatalogV1alpha1AssetRequestDefinitionAuthorization(varCatalogV1alpha1AssetRequestDefinitionAuthorization)

	return err
}

type NullableCatalogV1alpha1AssetRequestDefinitionAuthorization struct {
	value *CatalogV1alpha1AssetRequestDefinitionAuthorization
	isSet bool
}

func (v NullableCatalogV1alpha1AssetRequestDefinitionAuthorization) Get() *CatalogV1alpha1AssetRequestDefinitionAuthorization {
	return v.value
}

func (v *NullableCatalogV1alpha1AssetRequestDefinitionAuthorization) Set(val *CatalogV1alpha1AssetRequestDefinitionAuthorization) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogV1alpha1AssetRequestDefinitionAuthorization) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogV1alpha1AssetRequestDefinitionAuthorization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogV1alpha1AssetRequestDefinitionAuthorization(val *CatalogV1alpha1AssetRequestDefinitionAuthorization) *NullableCatalogV1alpha1AssetRequestDefinitionAuthorization {
	return &NullableCatalogV1alpha1AssetRequestDefinitionAuthorization{value: val, isSet: true}
}

func (v NullableCatalogV1alpha1AssetRequestDefinitionAuthorization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogV1alpha1AssetRequestDefinitionAuthorization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

