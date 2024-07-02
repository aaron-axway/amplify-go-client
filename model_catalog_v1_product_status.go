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

// checks if the CatalogV1ProductStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogV1ProductStatus{}

// CatalogV1ProductStatus struct for CatalogV1ProductStatus
type CatalogV1ProductStatus struct {
	// The current status level, indicating progress towards consistency.
	Level string `json:"level"`
	// Reasons for the generated status.
	Reasons []CatalogV1ProductStatusReasonsInner `json:"reasons,omitempty"`
}

type _CatalogV1ProductStatus CatalogV1ProductStatus

// NewCatalogV1ProductStatus instantiates a new CatalogV1ProductStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogV1ProductStatus(level string) *CatalogV1ProductStatus {
	this := CatalogV1ProductStatus{}
	this.Level = level
	return &this
}

// NewCatalogV1ProductStatusWithDefaults instantiates a new CatalogV1ProductStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogV1ProductStatusWithDefaults() *CatalogV1ProductStatus {
	this := CatalogV1ProductStatus{}
	return &this
}

// GetLevel returns the Level field value
func (o *CatalogV1ProductStatus) GetLevel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Level
}

// GetLevelOk returns a tuple with the Level field value
// and a boolean to check if the value has been set.
func (o *CatalogV1ProductStatus) GetLevelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Level, true
}

// SetLevel sets field value
func (o *CatalogV1ProductStatus) SetLevel(v string) {
	o.Level = v
}

// GetReasons returns the Reasons field value if set, zero value otherwise.
func (o *CatalogV1ProductStatus) GetReasons() []CatalogV1ProductStatusReasonsInner {
	if o == nil || IsNil(o.Reasons) {
		var ret []CatalogV1ProductStatusReasonsInner
		return ret
	}
	return o.Reasons
}

// GetReasonsOk returns a tuple with the Reasons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1ProductStatus) GetReasonsOk() ([]CatalogV1ProductStatusReasonsInner, bool) {
	if o == nil || IsNil(o.Reasons) {
		return nil, false
	}
	return o.Reasons, true
}

// HasReasons returns a boolean if a field has been set.
func (o *CatalogV1ProductStatus) HasReasons() bool {
	if o != nil && !IsNil(o.Reasons) {
		return true
	}

	return false
}

// SetReasons gets a reference to the given []CatalogV1ProductStatusReasonsInner and assigns it to the Reasons field.
func (o *CatalogV1ProductStatus) SetReasons(v []CatalogV1ProductStatusReasonsInner) {
	o.Reasons = v
}

func (o CatalogV1ProductStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogV1ProductStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["level"] = o.Level
	if !IsNil(o.Reasons) {
		toSerialize["reasons"] = o.Reasons
	}
	return toSerialize, nil
}

func (o *CatalogV1ProductStatus) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"level",
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

	varCatalogV1ProductStatus := _CatalogV1ProductStatus{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCatalogV1ProductStatus)

	if err != nil {
		return err
	}

	*o = CatalogV1ProductStatus(varCatalogV1ProductStatus)

	return err
}

type NullableCatalogV1ProductStatus struct {
	value *CatalogV1ProductStatus
	isSet bool
}

func (v NullableCatalogV1ProductStatus) Get() *CatalogV1ProductStatus {
	return v.value
}

func (v *NullableCatalogV1ProductStatus) Set(val *CatalogV1ProductStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogV1ProductStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogV1ProductStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogV1ProductStatus(val *CatalogV1ProductStatus) *NullableCatalogV1ProductStatus {
	return &NullableCatalogV1ProductStatus{value: val, isSet: true}
}

func (v NullableCatalogV1ProductStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogV1ProductStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


