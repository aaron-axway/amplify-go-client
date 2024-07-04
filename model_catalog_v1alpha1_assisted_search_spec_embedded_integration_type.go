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

// checks if the CatalogV1alpha1AssistedSearchSpecEmbeddedIntegrationType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogV1alpha1AssistedSearchSpecEmbeddedIntegrationType{}

// CatalogV1alpha1AssistedSearchSpecEmbeddedIntegrationType Assisted Search is managed by Marketplace service.
type CatalogV1alpha1AssistedSearchSpecEmbeddedIntegrationType struct {
	Type string `json:"type"`
}

type _CatalogV1alpha1AssistedSearchSpecEmbeddedIntegrationType CatalogV1alpha1AssistedSearchSpecEmbeddedIntegrationType

// NewCatalogV1alpha1AssistedSearchSpecEmbeddedIntegrationType instantiates a new CatalogV1alpha1AssistedSearchSpecEmbeddedIntegrationType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogV1alpha1AssistedSearchSpecEmbeddedIntegrationType(type_ string) *CatalogV1alpha1AssistedSearchSpecEmbeddedIntegrationType {
	this := CatalogV1alpha1AssistedSearchSpecEmbeddedIntegrationType{}
	this.Type = type_
	return &this
}

// NewCatalogV1alpha1AssistedSearchSpecEmbeddedIntegrationTypeWithDefaults instantiates a new CatalogV1alpha1AssistedSearchSpecEmbeddedIntegrationType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogV1alpha1AssistedSearchSpecEmbeddedIntegrationTypeWithDefaults() *CatalogV1alpha1AssistedSearchSpecEmbeddedIntegrationType {
	this := CatalogV1alpha1AssistedSearchSpecEmbeddedIntegrationType{}
	return &this
}

// GetType returns the Type field value
func (o *CatalogV1alpha1AssistedSearchSpecEmbeddedIntegrationType) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1AssistedSearchSpecEmbeddedIntegrationType) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CatalogV1alpha1AssistedSearchSpecEmbeddedIntegrationType) SetType(v string) {
	o.Type = v
}

func (o CatalogV1alpha1AssistedSearchSpecEmbeddedIntegrationType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogV1alpha1AssistedSearchSpecEmbeddedIntegrationType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *CatalogV1alpha1AssistedSearchSpecEmbeddedIntegrationType) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
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

	varCatalogV1alpha1AssistedSearchSpecEmbeddedIntegrationType := _CatalogV1alpha1AssistedSearchSpecEmbeddedIntegrationType{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCatalogV1alpha1AssistedSearchSpecEmbeddedIntegrationType)

	if err != nil {
		return err
	}

	*o = CatalogV1alpha1AssistedSearchSpecEmbeddedIntegrationType(varCatalogV1alpha1AssistedSearchSpecEmbeddedIntegrationType)

	return err
}

type NullableCatalogV1alpha1AssistedSearchSpecEmbeddedIntegrationType struct {
	value *CatalogV1alpha1AssistedSearchSpecEmbeddedIntegrationType
	isSet bool
}

func (v NullableCatalogV1alpha1AssistedSearchSpecEmbeddedIntegrationType) Get() *CatalogV1alpha1AssistedSearchSpecEmbeddedIntegrationType {
	return v.value
}

func (v *NullableCatalogV1alpha1AssistedSearchSpecEmbeddedIntegrationType) Set(val *CatalogV1alpha1AssistedSearchSpecEmbeddedIntegrationType) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogV1alpha1AssistedSearchSpecEmbeddedIntegrationType) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogV1alpha1AssistedSearchSpecEmbeddedIntegrationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogV1alpha1AssistedSearchSpecEmbeddedIntegrationType(val *CatalogV1alpha1AssistedSearchSpecEmbeddedIntegrationType) *NullableCatalogV1alpha1AssistedSearchSpecEmbeddedIntegrationType {
	return &NullableCatalogV1alpha1AssistedSearchSpecEmbeddedIntegrationType{value: val, isSet: true}
}

func (v NullableCatalogV1alpha1AssistedSearchSpecEmbeddedIntegrationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogV1alpha1AssistedSearchSpecEmbeddedIntegrationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


