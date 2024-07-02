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

// checks if the CatalogV1alpha1QuotaStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogV1alpha1QuotaStatus{}

// CatalogV1alpha1QuotaStatus struct for CatalogV1alpha1QuotaStatus
type CatalogV1alpha1QuotaStatus struct {
	// The current status level, indicating progress towards consistency.
	Level string `json:"level"`
	// Reasons for the generated status.
	Reasons []CatalogV1alpha1QuotaStatusReasonsInner `json:"reasons,omitempty"`
}

type _CatalogV1alpha1QuotaStatus CatalogV1alpha1QuotaStatus

// NewCatalogV1alpha1QuotaStatus instantiates a new CatalogV1alpha1QuotaStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogV1alpha1QuotaStatus(level string) *CatalogV1alpha1QuotaStatus {
	this := CatalogV1alpha1QuotaStatus{}
	this.Level = level
	return &this
}

// NewCatalogV1alpha1QuotaStatusWithDefaults instantiates a new CatalogV1alpha1QuotaStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogV1alpha1QuotaStatusWithDefaults() *CatalogV1alpha1QuotaStatus {
	this := CatalogV1alpha1QuotaStatus{}
	return &this
}

// GetLevel returns the Level field value
func (o *CatalogV1alpha1QuotaStatus) GetLevel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Level
}

// GetLevelOk returns a tuple with the Level field value
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1QuotaStatus) GetLevelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Level, true
}

// SetLevel sets field value
func (o *CatalogV1alpha1QuotaStatus) SetLevel(v string) {
	o.Level = v
}

// GetReasons returns the Reasons field value if set, zero value otherwise.
func (o *CatalogV1alpha1QuotaStatus) GetReasons() []CatalogV1alpha1QuotaStatusReasonsInner {
	if o == nil || IsNil(o.Reasons) {
		var ret []CatalogV1alpha1QuotaStatusReasonsInner
		return ret
	}
	return o.Reasons
}

// GetReasonsOk returns a tuple with the Reasons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1QuotaStatus) GetReasonsOk() ([]CatalogV1alpha1QuotaStatusReasonsInner, bool) {
	if o == nil || IsNil(o.Reasons) {
		return nil, false
	}
	return o.Reasons, true
}

// HasReasons returns a boolean if a field has been set.
func (o *CatalogV1alpha1QuotaStatus) HasReasons() bool {
	if o != nil && !IsNil(o.Reasons) {
		return true
	}

	return false
}

// SetReasons gets a reference to the given []CatalogV1alpha1QuotaStatusReasonsInner and assigns it to the Reasons field.
func (o *CatalogV1alpha1QuotaStatus) SetReasons(v []CatalogV1alpha1QuotaStatusReasonsInner) {
	o.Reasons = v
}

func (o CatalogV1alpha1QuotaStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogV1alpha1QuotaStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["level"] = o.Level
	if !IsNil(o.Reasons) {
		toSerialize["reasons"] = o.Reasons
	}
	return toSerialize, nil
}

func (o *CatalogV1alpha1QuotaStatus) UnmarshalJSON(data []byte) (err error) {
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

	varCatalogV1alpha1QuotaStatus := _CatalogV1alpha1QuotaStatus{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCatalogV1alpha1QuotaStatus)

	if err != nil {
		return err
	}

	*o = CatalogV1alpha1QuotaStatus(varCatalogV1alpha1QuotaStatus)

	return err
}

type NullableCatalogV1alpha1QuotaStatus struct {
	value *CatalogV1alpha1QuotaStatus
	isSet bool
}

func (v NullableCatalogV1alpha1QuotaStatus) Get() *CatalogV1alpha1QuotaStatus {
	return v.value
}

func (v *NullableCatalogV1alpha1QuotaStatus) Set(val *CatalogV1alpha1QuotaStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogV1alpha1QuotaStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogV1alpha1QuotaStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogV1alpha1QuotaStatus(val *CatalogV1alpha1QuotaStatus) *NullableCatalogV1alpha1QuotaStatus {
	return &NullableCatalogV1alpha1QuotaStatus{value: val, isSet: true}
}

func (v NullableCatalogV1alpha1QuotaStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogV1alpha1QuotaStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


