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

// checks if the CatalogV1alpha1SubscriptionInvoiceStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogV1alpha1SubscriptionInvoiceStatus{}

// CatalogV1alpha1SubscriptionInvoiceStatus struct for CatalogV1alpha1SubscriptionInvoiceStatus
type CatalogV1alpha1SubscriptionInvoiceStatus struct {
	// The current status level, indicating progress towards consistency.
	Level string `json:"level"`
	// Reasons for the generated status.
	Reasons []CatalogV1alpha1SubscriptionStatusReasonsInner `json:"reasons,omitempty"`
}

type _CatalogV1alpha1SubscriptionInvoiceStatus CatalogV1alpha1SubscriptionInvoiceStatus

// NewCatalogV1alpha1SubscriptionInvoiceStatus instantiates a new CatalogV1alpha1SubscriptionInvoiceStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogV1alpha1SubscriptionInvoiceStatus(level string) *CatalogV1alpha1SubscriptionInvoiceStatus {
	this := CatalogV1alpha1SubscriptionInvoiceStatus{}
	this.Level = level
	return &this
}

// NewCatalogV1alpha1SubscriptionInvoiceStatusWithDefaults instantiates a new CatalogV1alpha1SubscriptionInvoiceStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogV1alpha1SubscriptionInvoiceStatusWithDefaults() *CatalogV1alpha1SubscriptionInvoiceStatus {
	this := CatalogV1alpha1SubscriptionInvoiceStatus{}
	return &this
}

// GetLevel returns the Level field value
func (o *CatalogV1alpha1SubscriptionInvoiceStatus) GetLevel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Level
}

// GetLevelOk returns a tuple with the Level field value
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1SubscriptionInvoiceStatus) GetLevelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Level, true
}

// SetLevel sets field value
func (o *CatalogV1alpha1SubscriptionInvoiceStatus) SetLevel(v string) {
	o.Level = v
}

// GetReasons returns the Reasons field value if set, zero value otherwise.
func (o *CatalogV1alpha1SubscriptionInvoiceStatus) GetReasons() []CatalogV1alpha1SubscriptionStatusReasonsInner {
	if o == nil || IsNil(o.Reasons) {
		var ret []CatalogV1alpha1SubscriptionStatusReasonsInner
		return ret
	}
	return o.Reasons
}

// GetReasonsOk returns a tuple with the Reasons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1SubscriptionInvoiceStatus) GetReasonsOk() ([]CatalogV1alpha1SubscriptionStatusReasonsInner, bool) {
	if o == nil || IsNil(o.Reasons) {
		return nil, false
	}
	return o.Reasons, true
}

// HasReasons returns a boolean if a field has been set.
func (o *CatalogV1alpha1SubscriptionInvoiceStatus) HasReasons() bool {
	if o != nil && !IsNil(o.Reasons) {
		return true
	}

	return false
}

// SetReasons gets a reference to the given []CatalogV1alpha1SubscriptionStatusReasonsInner and assigns it to the Reasons field.
func (o *CatalogV1alpha1SubscriptionInvoiceStatus) SetReasons(v []CatalogV1alpha1SubscriptionStatusReasonsInner) {
	o.Reasons = v
}

func (o CatalogV1alpha1SubscriptionInvoiceStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogV1alpha1SubscriptionInvoiceStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["level"] = o.Level
	if !IsNil(o.Reasons) {
		toSerialize["reasons"] = o.Reasons
	}
	return toSerialize, nil
}

func (o *CatalogV1alpha1SubscriptionInvoiceStatus) UnmarshalJSON(data []byte) (err error) {
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

	varCatalogV1alpha1SubscriptionInvoiceStatus := _CatalogV1alpha1SubscriptionInvoiceStatus{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCatalogV1alpha1SubscriptionInvoiceStatus)

	if err != nil {
		return err
	}

	*o = CatalogV1alpha1SubscriptionInvoiceStatus(varCatalogV1alpha1SubscriptionInvoiceStatus)

	return err
}

type NullableCatalogV1alpha1SubscriptionInvoiceStatus struct {
	value *CatalogV1alpha1SubscriptionInvoiceStatus
	isSet bool
}

func (v NullableCatalogV1alpha1SubscriptionInvoiceStatus) Get() *CatalogV1alpha1SubscriptionInvoiceStatus {
	return v.value
}

func (v *NullableCatalogV1alpha1SubscriptionInvoiceStatus) Set(val *CatalogV1alpha1SubscriptionInvoiceStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogV1alpha1SubscriptionInvoiceStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogV1alpha1SubscriptionInvoiceStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogV1alpha1SubscriptionInvoiceStatus(val *CatalogV1alpha1SubscriptionInvoiceStatus) *NullableCatalogV1alpha1SubscriptionInvoiceStatus {
	return &NullableCatalogV1alpha1SubscriptionInvoiceStatus{value: val, isSet: true}
}

func (v NullableCatalogV1alpha1SubscriptionInvoiceStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogV1alpha1SubscriptionInvoiceStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


