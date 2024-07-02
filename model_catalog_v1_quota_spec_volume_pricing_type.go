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

// checks if the CatalogV1QuotaSpecVolumePricingType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogV1QuotaSpecVolumePricingType{}

// CatalogV1QuotaSpecVolumePricingType struct for CatalogV1QuotaSpecVolumePricingType
type CatalogV1QuotaSpecVolumePricingType struct {
	Type string `json:"type"`
	Limit CatalogV1QuotaSpecVolumePricingTypeLimit `json:"limit"`
}

type _CatalogV1QuotaSpecVolumePricingType CatalogV1QuotaSpecVolumePricingType

// NewCatalogV1QuotaSpecVolumePricingType instantiates a new CatalogV1QuotaSpecVolumePricingType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogV1QuotaSpecVolumePricingType(type_ string, limit CatalogV1QuotaSpecVolumePricingTypeLimit) *CatalogV1QuotaSpecVolumePricingType {
	this := CatalogV1QuotaSpecVolumePricingType{}
	this.Type = type_
	this.Limit = limit
	return &this
}

// NewCatalogV1QuotaSpecVolumePricingTypeWithDefaults instantiates a new CatalogV1QuotaSpecVolumePricingType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogV1QuotaSpecVolumePricingTypeWithDefaults() *CatalogV1QuotaSpecVolumePricingType {
	this := CatalogV1QuotaSpecVolumePricingType{}
	return &this
}

// GetType returns the Type field value
func (o *CatalogV1QuotaSpecVolumePricingType) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CatalogV1QuotaSpecVolumePricingType) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CatalogV1QuotaSpecVolumePricingType) SetType(v string) {
	o.Type = v
}

// GetLimit returns the Limit field value
func (o *CatalogV1QuotaSpecVolumePricingType) GetLimit() CatalogV1QuotaSpecVolumePricingTypeLimit {
	if o == nil {
		var ret CatalogV1QuotaSpecVolumePricingTypeLimit
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *CatalogV1QuotaSpecVolumePricingType) GetLimitOk() (*CatalogV1QuotaSpecVolumePricingTypeLimit, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *CatalogV1QuotaSpecVolumePricingType) SetLimit(v CatalogV1QuotaSpecVolumePricingTypeLimit) {
	o.Limit = v
}

func (o CatalogV1QuotaSpecVolumePricingType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogV1QuotaSpecVolumePricingType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["limit"] = o.Limit
	return toSerialize, nil
}

func (o *CatalogV1QuotaSpecVolumePricingType) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"limit",
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

	varCatalogV1QuotaSpecVolumePricingType := _CatalogV1QuotaSpecVolumePricingType{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCatalogV1QuotaSpecVolumePricingType)

	if err != nil {
		return err
	}

	*o = CatalogV1QuotaSpecVolumePricingType(varCatalogV1QuotaSpecVolumePricingType)

	return err
}

type NullableCatalogV1QuotaSpecVolumePricingType struct {
	value *CatalogV1QuotaSpecVolumePricingType
	isSet bool
}

func (v NullableCatalogV1QuotaSpecVolumePricingType) Get() *CatalogV1QuotaSpecVolumePricingType {
	return v.value
}

func (v *NullableCatalogV1QuotaSpecVolumePricingType) Set(val *CatalogV1QuotaSpecVolumePricingType) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogV1QuotaSpecVolumePricingType) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogV1QuotaSpecVolumePricingType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogV1QuotaSpecVolumePricingType(val *CatalogV1QuotaSpecVolumePricingType) *NullableCatalogV1QuotaSpecVolumePricingType {
	return &NullableCatalogV1QuotaSpecVolumePricingType{value: val, isSet: true}
}

func (v NullableCatalogV1QuotaSpecVolumePricingType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogV1QuotaSpecVolumePricingType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


