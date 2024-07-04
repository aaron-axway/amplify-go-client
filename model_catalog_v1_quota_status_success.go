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
	"time"
	"bytes"
	"fmt"
)

// checks if the CatalogV1QuotaStatusSuccess type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogV1QuotaStatusSuccess{}

// CatalogV1QuotaStatusSuccess struct for CatalogV1QuotaStatusSuccess
type CatalogV1QuotaStatusSuccess struct {
	Type string `json:"type"`
	// Time when the change occured.
	Timestamp time.Time `json:"timestamp"`
	// Details of the result.
	Detail string `json:"detail"`
	Meta map[string]interface{} `json:"meta,omitempty"`
}

type _CatalogV1QuotaStatusSuccess CatalogV1QuotaStatusSuccess

// NewCatalogV1QuotaStatusSuccess instantiates a new CatalogV1QuotaStatusSuccess object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogV1QuotaStatusSuccess(type_ string, timestamp time.Time, detail string) *CatalogV1QuotaStatusSuccess {
	this := CatalogV1QuotaStatusSuccess{}
	this.Type = type_
	this.Timestamp = timestamp
	this.Detail = detail
	return &this
}

// NewCatalogV1QuotaStatusSuccessWithDefaults instantiates a new CatalogV1QuotaStatusSuccess object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogV1QuotaStatusSuccessWithDefaults() *CatalogV1QuotaStatusSuccess {
	this := CatalogV1QuotaStatusSuccess{}
	return &this
}

// GetType returns the Type field value
func (o *CatalogV1QuotaStatusSuccess) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CatalogV1QuotaStatusSuccess) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CatalogV1QuotaStatusSuccess) SetType(v string) {
	o.Type = v
}

// GetTimestamp returns the Timestamp field value
func (o *CatalogV1QuotaStatusSuccess) GetTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *CatalogV1QuotaStatusSuccess) GetTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *CatalogV1QuotaStatusSuccess) SetTimestamp(v time.Time) {
	o.Timestamp = v
}

// GetDetail returns the Detail field value
func (o *CatalogV1QuotaStatusSuccess) GetDetail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Detail
}

// GetDetailOk returns a tuple with the Detail field value
// and a boolean to check if the value has been set.
func (o *CatalogV1QuotaStatusSuccess) GetDetailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Detail, true
}

// SetDetail sets field value
func (o *CatalogV1QuotaStatusSuccess) SetDetail(v string) {
	o.Detail = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *CatalogV1QuotaStatusSuccess) GetMeta() map[string]interface{} {
	if o == nil || IsNil(o.Meta) {
		var ret map[string]interface{}
		return ret
	}
	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1QuotaStatusSuccess) GetMetaOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Meta) {
		return map[string]interface{}{}, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *CatalogV1QuotaStatusSuccess) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given map[string]interface{} and assigns it to the Meta field.
func (o *CatalogV1QuotaStatusSuccess) SetMeta(v map[string]interface{}) {
	o.Meta = v
}

func (o CatalogV1QuotaStatusSuccess) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogV1QuotaStatusSuccess) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["timestamp"] = o.Timestamp
	toSerialize["detail"] = o.Detail
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return toSerialize, nil
}

func (o *CatalogV1QuotaStatusSuccess) UnmarshalJSON(data []byte) (err error) {
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

	varCatalogV1QuotaStatusSuccess := _CatalogV1QuotaStatusSuccess{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCatalogV1QuotaStatusSuccess)

	if err != nil {
		return err
	}

	*o = CatalogV1QuotaStatusSuccess(varCatalogV1QuotaStatusSuccess)

	return err
}

type NullableCatalogV1QuotaStatusSuccess struct {
	value *CatalogV1QuotaStatusSuccess
	isSet bool
}

func (v NullableCatalogV1QuotaStatusSuccess) Get() *CatalogV1QuotaStatusSuccess {
	return v.value
}

func (v *NullableCatalogV1QuotaStatusSuccess) Set(val *CatalogV1QuotaStatusSuccess) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogV1QuotaStatusSuccess) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogV1QuotaStatusSuccess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogV1QuotaStatusSuccess(val *CatalogV1QuotaStatusSuccess) *NullableCatalogV1QuotaStatusSuccess {
	return &NullableCatalogV1QuotaStatusSuccess{value: val, isSet: true}
}

func (v NullableCatalogV1QuotaStatusSuccess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogV1QuotaStatusSuccess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


