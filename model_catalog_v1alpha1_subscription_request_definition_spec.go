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

// checks if the CatalogV1alpha1SubscriptionRequestDefinitionSpec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogV1alpha1SubscriptionRequestDefinitionSpec{}

// CatalogV1alpha1SubscriptionRequestDefinitionSpec struct for CatalogV1alpha1SubscriptionRequestDefinitionSpec
type CatalogV1alpha1SubscriptionRequestDefinitionSpec struct {
	// JSON Schema draft \\#7 for defining the properties needed from a consumer to subscribe to a plan.
	Schema map[string]interface{} `json:"schema"`
}

type _CatalogV1alpha1SubscriptionRequestDefinitionSpec CatalogV1alpha1SubscriptionRequestDefinitionSpec

// NewCatalogV1alpha1SubscriptionRequestDefinitionSpec instantiates a new CatalogV1alpha1SubscriptionRequestDefinitionSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogV1alpha1SubscriptionRequestDefinitionSpec(schema map[string]interface{}) *CatalogV1alpha1SubscriptionRequestDefinitionSpec {
	this := CatalogV1alpha1SubscriptionRequestDefinitionSpec{}
	this.Schema = schema
	return &this
}

// NewCatalogV1alpha1SubscriptionRequestDefinitionSpecWithDefaults instantiates a new CatalogV1alpha1SubscriptionRequestDefinitionSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogV1alpha1SubscriptionRequestDefinitionSpecWithDefaults() *CatalogV1alpha1SubscriptionRequestDefinitionSpec {
	this := CatalogV1alpha1SubscriptionRequestDefinitionSpec{}
	return &this
}

// GetSchema returns the Schema field value
func (o *CatalogV1alpha1SubscriptionRequestDefinitionSpec) GetSchema() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1SubscriptionRequestDefinitionSpec) GetSchemaOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Schema, true
}

// SetSchema sets field value
func (o *CatalogV1alpha1SubscriptionRequestDefinitionSpec) SetSchema(v map[string]interface{}) {
	o.Schema = v
}

func (o CatalogV1alpha1SubscriptionRequestDefinitionSpec) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogV1alpha1SubscriptionRequestDefinitionSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schema"] = o.Schema
	return toSerialize, nil
}

func (o *CatalogV1alpha1SubscriptionRequestDefinitionSpec) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"schema",
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

	varCatalogV1alpha1SubscriptionRequestDefinitionSpec := _CatalogV1alpha1SubscriptionRequestDefinitionSpec{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCatalogV1alpha1SubscriptionRequestDefinitionSpec)

	if err != nil {
		return err
	}

	*o = CatalogV1alpha1SubscriptionRequestDefinitionSpec(varCatalogV1alpha1SubscriptionRequestDefinitionSpec)

	return err
}

type NullableCatalogV1alpha1SubscriptionRequestDefinitionSpec struct {
	value *CatalogV1alpha1SubscriptionRequestDefinitionSpec
	isSet bool
}

func (v NullableCatalogV1alpha1SubscriptionRequestDefinitionSpec) Get() *CatalogV1alpha1SubscriptionRequestDefinitionSpec {
	return v.value
}

func (v *NullableCatalogV1alpha1SubscriptionRequestDefinitionSpec) Set(val *CatalogV1alpha1SubscriptionRequestDefinitionSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogV1alpha1SubscriptionRequestDefinitionSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogV1alpha1SubscriptionRequestDefinitionSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogV1alpha1SubscriptionRequestDefinitionSpec(val *CatalogV1alpha1SubscriptionRequestDefinitionSpec) *NullableCatalogV1alpha1SubscriptionRequestDefinitionSpec {
	return &NullableCatalogV1alpha1SubscriptionRequestDefinitionSpec{value: val, isSet: true}
}

func (v NullableCatalogV1alpha1SubscriptionRequestDefinitionSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogV1alpha1SubscriptionRequestDefinitionSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


