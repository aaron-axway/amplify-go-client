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

// checks if the CatalogV1alpha1StageVisibilitySpecStagesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogV1alpha1StageVisibilitySpecStagesInner{}

// CatalogV1alpha1StageVisibilitySpecStagesInner struct for CatalogV1alpha1StageVisibilitySpecStagesInner
type CatalogV1alpha1StageVisibilitySpecStagesInner struct {
	Name string `json:"name" validate:"regexp=^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\\\\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*"`
}

type _CatalogV1alpha1StageVisibilitySpecStagesInner CatalogV1alpha1StageVisibilitySpecStagesInner

// NewCatalogV1alpha1StageVisibilitySpecStagesInner instantiates a new CatalogV1alpha1StageVisibilitySpecStagesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogV1alpha1StageVisibilitySpecStagesInner(name string) *CatalogV1alpha1StageVisibilitySpecStagesInner {
	this := CatalogV1alpha1StageVisibilitySpecStagesInner{}
	this.Name = name
	return &this
}

// NewCatalogV1alpha1StageVisibilitySpecStagesInnerWithDefaults instantiates a new CatalogV1alpha1StageVisibilitySpecStagesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogV1alpha1StageVisibilitySpecStagesInnerWithDefaults() *CatalogV1alpha1StageVisibilitySpecStagesInner {
	this := CatalogV1alpha1StageVisibilitySpecStagesInner{}
	return &this
}

// GetName returns the Name field value
func (o *CatalogV1alpha1StageVisibilitySpecStagesInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1StageVisibilitySpecStagesInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CatalogV1alpha1StageVisibilitySpecStagesInner) SetName(v string) {
	o.Name = v
}

func (o CatalogV1alpha1StageVisibilitySpecStagesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogV1alpha1StageVisibilitySpecStagesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

func (o *CatalogV1alpha1StageVisibilitySpecStagesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varCatalogV1alpha1StageVisibilitySpecStagesInner := _CatalogV1alpha1StageVisibilitySpecStagesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCatalogV1alpha1StageVisibilitySpecStagesInner)

	if err != nil {
		return err
	}

	*o = CatalogV1alpha1StageVisibilitySpecStagesInner(varCatalogV1alpha1StageVisibilitySpecStagesInner)

	return err
}

type NullableCatalogV1alpha1StageVisibilitySpecStagesInner struct {
	value *CatalogV1alpha1StageVisibilitySpecStagesInner
	isSet bool
}

func (v NullableCatalogV1alpha1StageVisibilitySpecStagesInner) Get() *CatalogV1alpha1StageVisibilitySpecStagesInner {
	return v.value
}

func (v *NullableCatalogV1alpha1StageVisibilitySpecStagesInner) Set(val *CatalogV1alpha1StageVisibilitySpecStagesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogV1alpha1StageVisibilitySpecStagesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogV1alpha1StageVisibilitySpecStagesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogV1alpha1StageVisibilitySpecStagesInner(val *CatalogV1alpha1StageVisibilitySpecStagesInner) *NullableCatalogV1alpha1StageVisibilitySpecStagesInner {
	return &NullableCatalogV1alpha1StageVisibilitySpecStagesInner{value: val, isSet: true}
}

func (v NullableCatalogV1alpha1StageVisibilitySpecStagesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogV1alpha1StageVisibilitySpecStagesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

