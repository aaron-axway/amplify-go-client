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
	"time"
	"bytes"
	"fmt"
)

// checks if the CatalogV1ProductStatusReasonsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogV1ProductStatusReasonsInner{}

// CatalogV1ProductStatusReasonsInner struct for CatalogV1ProductStatusReasonsInner
type CatalogV1ProductStatusReasonsInner struct {
	Type string `json:"type"`
	// Details of the error.
	Detail string `json:"detail"`
	// Time when the update occurred.
	Timestamp time.Time `json:"timestamp"`
	Meta *CatalogV1ProductStatusReasonsInnerMeta `json:"meta,omitempty"`
}

type _CatalogV1ProductStatusReasonsInner CatalogV1ProductStatusReasonsInner

// NewCatalogV1ProductStatusReasonsInner instantiates a new CatalogV1ProductStatusReasonsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogV1ProductStatusReasonsInner(type_ string, detail string, timestamp time.Time) *CatalogV1ProductStatusReasonsInner {
	this := CatalogV1ProductStatusReasonsInner{}
	this.Type = type_
	this.Detail = detail
	this.Timestamp = timestamp
	return &this
}

// NewCatalogV1ProductStatusReasonsInnerWithDefaults instantiates a new CatalogV1ProductStatusReasonsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogV1ProductStatusReasonsInnerWithDefaults() *CatalogV1ProductStatusReasonsInner {
	this := CatalogV1ProductStatusReasonsInner{}
	return &this
}

// GetType returns the Type field value
func (o *CatalogV1ProductStatusReasonsInner) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CatalogV1ProductStatusReasonsInner) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CatalogV1ProductStatusReasonsInner) SetType(v string) {
	o.Type = v
}

// GetDetail returns the Detail field value
func (o *CatalogV1ProductStatusReasonsInner) GetDetail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Detail
}

// GetDetailOk returns a tuple with the Detail field value
// and a boolean to check if the value has been set.
func (o *CatalogV1ProductStatusReasonsInner) GetDetailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Detail, true
}

// SetDetail sets field value
func (o *CatalogV1ProductStatusReasonsInner) SetDetail(v string) {
	o.Detail = v
}

// GetTimestamp returns the Timestamp field value
func (o *CatalogV1ProductStatusReasonsInner) GetTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *CatalogV1ProductStatusReasonsInner) GetTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *CatalogV1ProductStatusReasonsInner) SetTimestamp(v time.Time) {
	o.Timestamp = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *CatalogV1ProductStatusReasonsInner) GetMeta() CatalogV1ProductStatusReasonsInnerMeta {
	if o == nil || IsNil(o.Meta) {
		var ret CatalogV1ProductStatusReasonsInnerMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1ProductStatusReasonsInner) GetMetaOk() (*CatalogV1ProductStatusReasonsInnerMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *CatalogV1ProductStatusReasonsInner) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given CatalogV1ProductStatusReasonsInnerMeta and assigns it to the Meta field.
func (o *CatalogV1ProductStatusReasonsInner) SetMeta(v CatalogV1ProductStatusReasonsInnerMeta) {
	o.Meta = &v
}

func (o CatalogV1ProductStatusReasonsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogV1ProductStatusReasonsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["detail"] = o.Detail
	toSerialize["timestamp"] = o.Timestamp
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return toSerialize, nil
}

func (o *CatalogV1ProductStatusReasonsInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"detail",
		"timestamp",
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

	varCatalogV1ProductStatusReasonsInner := _CatalogV1ProductStatusReasonsInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCatalogV1ProductStatusReasonsInner)

	if err != nil {
		return err
	}

	*o = CatalogV1ProductStatusReasonsInner(varCatalogV1ProductStatusReasonsInner)

	return err
}

type NullableCatalogV1ProductStatusReasonsInner struct {
	value *CatalogV1ProductStatusReasonsInner
	isSet bool
}

func (v NullableCatalogV1ProductStatusReasonsInner) Get() *CatalogV1ProductStatusReasonsInner {
	return v.value
}

func (v *NullableCatalogV1ProductStatusReasonsInner) Set(val *CatalogV1ProductStatusReasonsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogV1ProductStatusReasonsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogV1ProductStatusReasonsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogV1ProductStatusReasonsInner(val *CatalogV1ProductStatusReasonsInner) *NullableCatalogV1ProductStatusReasonsInner {
	return &NullableCatalogV1ProductStatusReasonsInner{value: val, isSet: true}
}

func (v NullableCatalogV1ProductStatusReasonsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogV1ProductStatusReasonsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


