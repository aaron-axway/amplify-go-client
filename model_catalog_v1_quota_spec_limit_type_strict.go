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

// checks if the CatalogV1QuotaSpecLimitTypeStrict type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogV1QuotaSpecLimitTypeStrict{}

// CatalogV1QuotaSpecLimitTypeStrict struct for CatalogV1QuotaSpecLimitTypeStrict
type CatalogV1QuotaSpecLimitTypeStrict struct {
	Type string `json:"type"`
	// The limit of the unit that is provided.
	Value int32 `json:"value"`
}

type _CatalogV1QuotaSpecLimitTypeStrict CatalogV1QuotaSpecLimitTypeStrict

// NewCatalogV1QuotaSpecLimitTypeStrict instantiates a new CatalogV1QuotaSpecLimitTypeStrict object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogV1QuotaSpecLimitTypeStrict(type_ string, value int32) *CatalogV1QuotaSpecLimitTypeStrict {
	this := CatalogV1QuotaSpecLimitTypeStrict{}
	this.Type = type_
	this.Value = value
	return &this
}

// NewCatalogV1QuotaSpecLimitTypeStrictWithDefaults instantiates a new CatalogV1QuotaSpecLimitTypeStrict object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogV1QuotaSpecLimitTypeStrictWithDefaults() *CatalogV1QuotaSpecLimitTypeStrict {
	this := CatalogV1QuotaSpecLimitTypeStrict{}
	return &this
}

// GetType returns the Type field value
func (o *CatalogV1QuotaSpecLimitTypeStrict) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CatalogV1QuotaSpecLimitTypeStrict) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CatalogV1QuotaSpecLimitTypeStrict) SetType(v string) {
	o.Type = v
}

// GetValue returns the Value field value
func (o *CatalogV1QuotaSpecLimitTypeStrict) GetValue() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *CatalogV1QuotaSpecLimitTypeStrict) GetValueOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *CatalogV1QuotaSpecLimitTypeStrict) SetValue(v int32) {
	o.Value = v
}

func (o CatalogV1QuotaSpecLimitTypeStrict) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogV1QuotaSpecLimitTypeStrict) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

func (o *CatalogV1QuotaSpecLimitTypeStrict) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"value",
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

	varCatalogV1QuotaSpecLimitTypeStrict := _CatalogV1QuotaSpecLimitTypeStrict{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCatalogV1QuotaSpecLimitTypeStrict)

	if err != nil {
		return err
	}

	*o = CatalogV1QuotaSpecLimitTypeStrict(varCatalogV1QuotaSpecLimitTypeStrict)

	return err
}

type NullableCatalogV1QuotaSpecLimitTypeStrict struct {
	value *CatalogV1QuotaSpecLimitTypeStrict
	isSet bool
}

func (v NullableCatalogV1QuotaSpecLimitTypeStrict) Get() *CatalogV1QuotaSpecLimitTypeStrict {
	return v.value
}

func (v *NullableCatalogV1QuotaSpecLimitTypeStrict) Set(val *CatalogV1QuotaSpecLimitTypeStrict) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogV1QuotaSpecLimitTypeStrict) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogV1QuotaSpecLimitTypeStrict) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogV1QuotaSpecLimitTypeStrict(val *CatalogV1QuotaSpecLimitTypeStrict) *NullableCatalogV1QuotaSpecLimitTypeStrict {
	return &NullableCatalogV1QuotaSpecLimitTypeStrict{value: val, isSet: true}
}

func (v NullableCatalogV1QuotaSpecLimitTypeStrict) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogV1QuotaSpecLimitTypeStrict) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

