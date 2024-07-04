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
	"bytes"
	"fmt"
)

// checks if the CatalogV1alpha1PublishedStageSpec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogV1alpha1PublishedStageSpec{}

// CatalogV1alpha1PublishedStageSpec struct for CatalogV1alpha1PublishedStageSpec
type CatalogV1alpha1PublishedStageSpec struct {
	Stage CatalogV1PublishedStageSpecStage `json:"stage"`
}

type _CatalogV1alpha1PublishedStageSpec CatalogV1alpha1PublishedStageSpec

// NewCatalogV1alpha1PublishedStageSpec instantiates a new CatalogV1alpha1PublishedStageSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogV1alpha1PublishedStageSpec(stage CatalogV1PublishedStageSpecStage) *CatalogV1alpha1PublishedStageSpec {
	this := CatalogV1alpha1PublishedStageSpec{}
	this.Stage = stage
	return &this
}

// NewCatalogV1alpha1PublishedStageSpecWithDefaults instantiates a new CatalogV1alpha1PublishedStageSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogV1alpha1PublishedStageSpecWithDefaults() *CatalogV1alpha1PublishedStageSpec {
	this := CatalogV1alpha1PublishedStageSpec{}
	return &this
}

// GetStage returns the Stage field value
func (o *CatalogV1alpha1PublishedStageSpec) GetStage() CatalogV1PublishedStageSpecStage {
	if o == nil {
		var ret CatalogV1PublishedStageSpecStage
		return ret
	}

	return o.Stage
}

// GetStageOk returns a tuple with the Stage field value
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1PublishedStageSpec) GetStageOk() (*CatalogV1PublishedStageSpecStage, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Stage, true
}

// SetStage sets field value
func (o *CatalogV1alpha1PublishedStageSpec) SetStage(v CatalogV1PublishedStageSpecStage) {
	o.Stage = v
}

func (o CatalogV1alpha1PublishedStageSpec) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogV1alpha1PublishedStageSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["stage"] = o.Stage
	return toSerialize, nil
}

func (o *CatalogV1alpha1PublishedStageSpec) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"stage",
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

	varCatalogV1alpha1PublishedStageSpec := _CatalogV1alpha1PublishedStageSpec{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCatalogV1alpha1PublishedStageSpec)

	if err != nil {
		return err
	}

	*o = CatalogV1alpha1PublishedStageSpec(varCatalogV1alpha1PublishedStageSpec)

	return err
}

type NullableCatalogV1alpha1PublishedStageSpec struct {
	value *CatalogV1alpha1PublishedStageSpec
	isSet bool
}

func (v NullableCatalogV1alpha1PublishedStageSpec) Get() *CatalogV1alpha1PublishedStageSpec {
	return v.value
}

func (v *NullableCatalogV1alpha1PublishedStageSpec) Set(val *CatalogV1alpha1PublishedStageSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogV1alpha1PublishedStageSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogV1alpha1PublishedStageSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogV1alpha1PublishedStageSpec(val *CatalogV1alpha1PublishedStageSpec) *NullableCatalogV1alpha1PublishedStageSpec {
	return &NullableCatalogV1alpha1PublishedStageSpec{value: val, isSet: true}
}

func (v NullableCatalogV1alpha1PublishedStageSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogV1alpha1PublishedStageSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


