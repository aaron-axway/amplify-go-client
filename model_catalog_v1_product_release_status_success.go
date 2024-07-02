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

// checks if the CatalogV1ProductReleaseStatusSuccess type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogV1ProductReleaseStatusSuccess{}

// CatalogV1ProductReleaseStatusSuccess struct for CatalogV1ProductReleaseStatusSuccess
type CatalogV1ProductReleaseStatusSuccess struct {
	Type string `json:"type"`
	// Time when the change occured.
	Timestamp time.Time `json:"timestamp"`
	// message of the result
	Detail string `json:"detail"`
	Meta *CatalogV1ProductStatusReasonsInnerMeta `json:"meta,omitempty"`
}

type _CatalogV1ProductReleaseStatusSuccess CatalogV1ProductReleaseStatusSuccess

// NewCatalogV1ProductReleaseStatusSuccess instantiates a new CatalogV1ProductReleaseStatusSuccess object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogV1ProductReleaseStatusSuccess(type_ string, timestamp time.Time, detail string) *CatalogV1ProductReleaseStatusSuccess {
	this := CatalogV1ProductReleaseStatusSuccess{}
	this.Type = type_
	this.Timestamp = timestamp
	this.Detail = detail
	return &this
}

// NewCatalogV1ProductReleaseStatusSuccessWithDefaults instantiates a new CatalogV1ProductReleaseStatusSuccess object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogV1ProductReleaseStatusSuccessWithDefaults() *CatalogV1ProductReleaseStatusSuccess {
	this := CatalogV1ProductReleaseStatusSuccess{}
	return &this
}

// GetType returns the Type field value
func (o *CatalogV1ProductReleaseStatusSuccess) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CatalogV1ProductReleaseStatusSuccess) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CatalogV1ProductReleaseStatusSuccess) SetType(v string) {
	o.Type = v
}

// GetTimestamp returns the Timestamp field value
func (o *CatalogV1ProductReleaseStatusSuccess) GetTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *CatalogV1ProductReleaseStatusSuccess) GetTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *CatalogV1ProductReleaseStatusSuccess) SetTimestamp(v time.Time) {
	o.Timestamp = v
}

// GetDetail returns the Detail field value
func (o *CatalogV1ProductReleaseStatusSuccess) GetDetail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Detail
}

// GetDetailOk returns a tuple with the Detail field value
// and a boolean to check if the value has been set.
func (o *CatalogV1ProductReleaseStatusSuccess) GetDetailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Detail, true
}

// SetDetail sets field value
func (o *CatalogV1ProductReleaseStatusSuccess) SetDetail(v string) {
	o.Detail = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *CatalogV1ProductReleaseStatusSuccess) GetMeta() CatalogV1ProductStatusReasonsInnerMeta {
	if o == nil || IsNil(o.Meta) {
		var ret CatalogV1ProductStatusReasonsInnerMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1ProductReleaseStatusSuccess) GetMetaOk() (*CatalogV1ProductStatusReasonsInnerMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *CatalogV1ProductReleaseStatusSuccess) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given CatalogV1ProductStatusReasonsInnerMeta and assigns it to the Meta field.
func (o *CatalogV1ProductReleaseStatusSuccess) SetMeta(v CatalogV1ProductStatusReasonsInnerMeta) {
	o.Meta = &v
}

func (o CatalogV1ProductReleaseStatusSuccess) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogV1ProductReleaseStatusSuccess) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["timestamp"] = o.Timestamp
	toSerialize["detail"] = o.Detail
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return toSerialize, nil
}

func (o *CatalogV1ProductReleaseStatusSuccess) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"timestamp",
		"detail",
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

	varCatalogV1ProductReleaseStatusSuccess := _CatalogV1ProductReleaseStatusSuccess{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCatalogV1ProductReleaseStatusSuccess)

	if err != nil {
		return err
	}

	*o = CatalogV1ProductReleaseStatusSuccess(varCatalogV1ProductReleaseStatusSuccess)

	return err
}

type NullableCatalogV1ProductReleaseStatusSuccess struct {
	value *CatalogV1ProductReleaseStatusSuccess
	isSet bool
}

func (v NullableCatalogV1ProductReleaseStatusSuccess) Get() *CatalogV1ProductReleaseStatusSuccess {
	return v.value
}

func (v *NullableCatalogV1ProductReleaseStatusSuccess) Set(val *CatalogV1ProductReleaseStatusSuccess) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogV1ProductReleaseStatusSuccess) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogV1ProductReleaseStatusSuccess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogV1ProductReleaseStatusSuccess(val *CatalogV1ProductReleaseStatusSuccess) *NullableCatalogV1ProductReleaseStatusSuccess {
	return &NullableCatalogV1ProductReleaseStatusSuccess{value: val, isSet: true}
}

func (v NullableCatalogV1ProductReleaseStatusSuccess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogV1ProductReleaseStatusSuccess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

