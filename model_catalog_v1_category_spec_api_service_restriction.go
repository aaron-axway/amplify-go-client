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

// checks if the CatalogV1CategorySpecAPIServiceRestriction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogV1CategorySpecAPIServiceRestriction{}

// CatalogV1CategorySpecAPIServiceRestriction struct for CatalogV1CategorySpecAPIServiceRestriction
type CatalogV1CategorySpecAPIServiceRestriction struct {
	Type string `json:"type"`
}

type _CatalogV1CategorySpecAPIServiceRestriction CatalogV1CategorySpecAPIServiceRestriction

// NewCatalogV1CategorySpecAPIServiceRestriction instantiates a new CatalogV1CategorySpecAPIServiceRestriction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogV1CategorySpecAPIServiceRestriction(type_ string) *CatalogV1CategorySpecAPIServiceRestriction {
	this := CatalogV1CategorySpecAPIServiceRestriction{}
	this.Type = type_
	return &this
}

// NewCatalogV1CategorySpecAPIServiceRestrictionWithDefaults instantiates a new CatalogV1CategorySpecAPIServiceRestriction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogV1CategorySpecAPIServiceRestrictionWithDefaults() *CatalogV1CategorySpecAPIServiceRestriction {
	this := CatalogV1CategorySpecAPIServiceRestriction{}
	return &this
}

// GetType returns the Type field value
func (o *CatalogV1CategorySpecAPIServiceRestriction) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CatalogV1CategorySpecAPIServiceRestriction) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CatalogV1CategorySpecAPIServiceRestriction) SetType(v string) {
	o.Type = v
}

func (o CatalogV1CategorySpecAPIServiceRestriction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogV1CategorySpecAPIServiceRestriction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *CatalogV1CategorySpecAPIServiceRestriction) UnmarshalJSON(data []byte) (err error) {
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

	varCatalogV1CategorySpecAPIServiceRestriction := _CatalogV1CategorySpecAPIServiceRestriction{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCatalogV1CategorySpecAPIServiceRestriction)

	if err != nil {
		return err
	}

	*o = CatalogV1CategorySpecAPIServiceRestriction(varCatalogV1CategorySpecAPIServiceRestriction)

	return err
}

type NullableCatalogV1CategorySpecAPIServiceRestriction struct {
	value *CatalogV1CategorySpecAPIServiceRestriction
	isSet bool
}

func (v NullableCatalogV1CategorySpecAPIServiceRestriction) Get() *CatalogV1CategorySpecAPIServiceRestriction {
	return v.value
}

func (v *NullableCatalogV1CategorySpecAPIServiceRestriction) Set(val *CatalogV1CategorySpecAPIServiceRestriction) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogV1CategorySpecAPIServiceRestriction) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogV1CategorySpecAPIServiceRestriction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogV1CategorySpecAPIServiceRestriction(val *CatalogV1CategorySpecAPIServiceRestriction) *NullableCatalogV1CategorySpecAPIServiceRestriction {
	return &NullableCatalogV1CategorySpecAPIServiceRestriction{value: val, isSet: true}
}

func (v NullableCatalogV1CategorySpecAPIServiceRestriction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogV1CategorySpecAPIServiceRestriction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

