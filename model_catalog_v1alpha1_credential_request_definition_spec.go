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

// checks if the CatalogV1alpha1CredentialRequestDefinitionSpec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogV1alpha1CredentialRequestDefinitionSpec{}

// CatalogV1alpha1CredentialRequestDefinitionSpec struct for CatalogV1alpha1CredentialRequestDefinitionSpec
type CatalogV1alpha1CredentialRequestDefinitionSpec struct {
	// JSON Schema draft \\#7 for describing the fields needed to provision Credentials of that type.
	Schema map[string]interface{} `json:"schema"`
	Provision ManagementV1alpha1CredentialRequestDefinitionSpecProvision `json:"provision"`
}

type _CatalogV1alpha1CredentialRequestDefinitionSpec CatalogV1alpha1CredentialRequestDefinitionSpec

// NewCatalogV1alpha1CredentialRequestDefinitionSpec instantiates a new CatalogV1alpha1CredentialRequestDefinitionSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogV1alpha1CredentialRequestDefinitionSpec(schema map[string]interface{}, provision ManagementV1alpha1CredentialRequestDefinitionSpecProvision) *CatalogV1alpha1CredentialRequestDefinitionSpec {
	this := CatalogV1alpha1CredentialRequestDefinitionSpec{}
	this.Schema = schema
	this.Provision = provision
	return &this
}

// NewCatalogV1alpha1CredentialRequestDefinitionSpecWithDefaults instantiates a new CatalogV1alpha1CredentialRequestDefinitionSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogV1alpha1CredentialRequestDefinitionSpecWithDefaults() *CatalogV1alpha1CredentialRequestDefinitionSpec {
	this := CatalogV1alpha1CredentialRequestDefinitionSpec{}
	return &this
}

// GetSchema returns the Schema field value
func (o *CatalogV1alpha1CredentialRequestDefinitionSpec) GetSchema() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1CredentialRequestDefinitionSpec) GetSchemaOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Schema, true
}

// SetSchema sets field value
func (o *CatalogV1alpha1CredentialRequestDefinitionSpec) SetSchema(v map[string]interface{}) {
	o.Schema = v
}

// GetProvision returns the Provision field value
func (o *CatalogV1alpha1CredentialRequestDefinitionSpec) GetProvision() ManagementV1alpha1CredentialRequestDefinitionSpecProvision {
	if o == nil {
		var ret ManagementV1alpha1CredentialRequestDefinitionSpecProvision
		return ret
	}

	return o.Provision
}

// GetProvisionOk returns a tuple with the Provision field value
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1CredentialRequestDefinitionSpec) GetProvisionOk() (*ManagementV1alpha1CredentialRequestDefinitionSpecProvision, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Provision, true
}

// SetProvision sets field value
func (o *CatalogV1alpha1CredentialRequestDefinitionSpec) SetProvision(v ManagementV1alpha1CredentialRequestDefinitionSpecProvision) {
	o.Provision = v
}

func (o CatalogV1alpha1CredentialRequestDefinitionSpec) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogV1alpha1CredentialRequestDefinitionSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schema"] = o.Schema
	toSerialize["provision"] = o.Provision
	return toSerialize, nil
}

func (o *CatalogV1alpha1CredentialRequestDefinitionSpec) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"schema",
		"provision",
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

	varCatalogV1alpha1CredentialRequestDefinitionSpec := _CatalogV1alpha1CredentialRequestDefinitionSpec{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCatalogV1alpha1CredentialRequestDefinitionSpec)

	if err != nil {
		return err
	}

	*o = CatalogV1alpha1CredentialRequestDefinitionSpec(varCatalogV1alpha1CredentialRequestDefinitionSpec)

	return err
}

type NullableCatalogV1alpha1CredentialRequestDefinitionSpec struct {
	value *CatalogV1alpha1CredentialRequestDefinitionSpec
	isSet bool
}

func (v NullableCatalogV1alpha1CredentialRequestDefinitionSpec) Get() *CatalogV1alpha1CredentialRequestDefinitionSpec {
	return v.value
}

func (v *NullableCatalogV1alpha1CredentialRequestDefinitionSpec) Set(val *CatalogV1alpha1CredentialRequestDefinitionSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogV1alpha1CredentialRequestDefinitionSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogV1alpha1CredentialRequestDefinitionSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogV1alpha1CredentialRequestDefinitionSpec(val *CatalogV1alpha1CredentialRequestDefinitionSpec) *NullableCatalogV1alpha1CredentialRequestDefinitionSpec {
	return &NullableCatalogV1alpha1CredentialRequestDefinitionSpec{value: val, isSet: true}
}

func (v NullableCatalogV1alpha1CredentialRequestDefinitionSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogV1alpha1CredentialRequestDefinitionSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

