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

// checks if the CatalogV1DocumentStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogV1DocumentStatus{}

// CatalogV1DocumentStatus struct for CatalogV1DocumentStatus
type CatalogV1DocumentStatus struct {
	// The current status level, indicating progress towards consistency.
	Level string `json:"level"`
	// Reasons for the generated status.
	Reasons []CatalogV1DocumentStatusReasonsInner `json:"reasons,omitempty"`
}

type _CatalogV1DocumentStatus CatalogV1DocumentStatus

// NewCatalogV1DocumentStatus instantiates a new CatalogV1DocumentStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogV1DocumentStatus(level string) *CatalogV1DocumentStatus {
	this := CatalogV1DocumentStatus{}
	this.Level = level
	return &this
}

// NewCatalogV1DocumentStatusWithDefaults instantiates a new CatalogV1DocumentStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogV1DocumentStatusWithDefaults() *CatalogV1DocumentStatus {
	this := CatalogV1DocumentStatus{}
	return &this
}

// GetLevel returns the Level field value
func (o *CatalogV1DocumentStatus) GetLevel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Level
}

// GetLevelOk returns a tuple with the Level field value
// and a boolean to check if the value has been set.
func (o *CatalogV1DocumentStatus) GetLevelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Level, true
}

// SetLevel sets field value
func (o *CatalogV1DocumentStatus) SetLevel(v string) {
	o.Level = v
}

// GetReasons returns the Reasons field value if set, zero value otherwise.
func (o *CatalogV1DocumentStatus) GetReasons() []CatalogV1DocumentStatusReasonsInner {
	if o == nil || IsNil(o.Reasons) {
		var ret []CatalogV1DocumentStatusReasonsInner
		return ret
	}
	return o.Reasons
}

// GetReasonsOk returns a tuple with the Reasons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1DocumentStatus) GetReasonsOk() ([]CatalogV1DocumentStatusReasonsInner, bool) {
	if o == nil || IsNil(o.Reasons) {
		return nil, false
	}
	return o.Reasons, true
}

// HasReasons returns a boolean if a field has been set.
func (o *CatalogV1DocumentStatus) HasReasons() bool {
	if o != nil && !IsNil(o.Reasons) {
		return true
	}

	return false
}

// SetReasons gets a reference to the given []CatalogV1DocumentStatusReasonsInner and assigns it to the Reasons field.
func (o *CatalogV1DocumentStatus) SetReasons(v []CatalogV1DocumentStatusReasonsInner) {
	o.Reasons = v
}

func (o CatalogV1DocumentStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogV1DocumentStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["level"] = o.Level
	if !IsNil(o.Reasons) {
		toSerialize["reasons"] = o.Reasons
	}
	return toSerialize, nil
}

func (o *CatalogV1DocumentStatus) UnmarshalJSON(data []byte) (err error) {
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

	varCatalogV1DocumentStatus := _CatalogV1DocumentStatus{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCatalogV1DocumentStatus)

	if err != nil {
		return err
	}

	*o = CatalogV1DocumentStatus(varCatalogV1DocumentStatus)

	return err
}

type NullableCatalogV1DocumentStatus struct {
	value *CatalogV1DocumentStatus
	isSet bool
}

func (v NullableCatalogV1DocumentStatus) Get() *CatalogV1DocumentStatus {
	return v.value
}

func (v *NullableCatalogV1DocumentStatus) Set(val *CatalogV1DocumentStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogV1DocumentStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogV1DocumentStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogV1DocumentStatus(val *CatalogV1DocumentStatus) *NullableCatalogV1DocumentStatus {
	return &NullableCatalogV1DocumentStatus{value: val, isSet: true}
}

func (v NullableCatalogV1DocumentStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogV1DocumentStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


