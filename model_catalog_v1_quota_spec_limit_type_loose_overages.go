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
)

// checks if the CatalogV1QuotaSpecLimitTypeLooseOverages type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogV1QuotaSpecLimitTypeLooseOverages{}

// CatalogV1QuotaSpecLimitTypeLooseOverages struct for CatalogV1QuotaSpecLimitTypeLooseOverages
type CatalogV1QuotaSpecLimitTypeLooseOverages struct {
	// The overage price per unit.
	Value *float64 `json:"value,omitempty"`
}

// NewCatalogV1QuotaSpecLimitTypeLooseOverages instantiates a new CatalogV1QuotaSpecLimitTypeLooseOverages object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogV1QuotaSpecLimitTypeLooseOverages() *CatalogV1QuotaSpecLimitTypeLooseOverages {
	this := CatalogV1QuotaSpecLimitTypeLooseOverages{}
	return &this
}

// NewCatalogV1QuotaSpecLimitTypeLooseOveragesWithDefaults instantiates a new CatalogV1QuotaSpecLimitTypeLooseOverages object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogV1QuotaSpecLimitTypeLooseOveragesWithDefaults() *CatalogV1QuotaSpecLimitTypeLooseOverages {
	this := CatalogV1QuotaSpecLimitTypeLooseOverages{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CatalogV1QuotaSpecLimitTypeLooseOverages) GetValue() float64 {
	if o == nil || IsNil(o.Value) {
		var ret float64
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1QuotaSpecLimitTypeLooseOverages) GetValueOk() (*float64, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CatalogV1QuotaSpecLimitTypeLooseOverages) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given float64 and assigns it to the Value field.
func (o *CatalogV1QuotaSpecLimitTypeLooseOverages) SetValue(v float64) {
	o.Value = &v
}

func (o CatalogV1QuotaSpecLimitTypeLooseOverages) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogV1QuotaSpecLimitTypeLooseOverages) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableCatalogV1QuotaSpecLimitTypeLooseOverages struct {
	value *CatalogV1QuotaSpecLimitTypeLooseOverages
	isSet bool
}

func (v NullableCatalogV1QuotaSpecLimitTypeLooseOverages) Get() *CatalogV1QuotaSpecLimitTypeLooseOverages {
	return v.value
}

func (v *NullableCatalogV1QuotaSpecLimitTypeLooseOverages) Set(val *CatalogV1QuotaSpecLimitTypeLooseOverages) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogV1QuotaSpecLimitTypeLooseOverages) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogV1QuotaSpecLimitTypeLooseOverages) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogV1QuotaSpecLimitTypeLooseOverages(val *CatalogV1QuotaSpecLimitTypeLooseOverages) *NullableCatalogV1QuotaSpecLimitTypeLooseOverages {
	return &NullableCatalogV1QuotaSpecLimitTypeLooseOverages{value: val, isSet: true}
}

func (v NullableCatalogV1QuotaSpecLimitTypeLooseOverages) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogV1QuotaSpecLimitTypeLooseOverages) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

