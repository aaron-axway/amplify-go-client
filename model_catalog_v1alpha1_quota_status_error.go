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
	"fmt"
)

// checks if the CatalogV1alpha1QuotaStatusError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogV1alpha1QuotaStatusError{}

// CatalogV1alpha1QuotaStatusError struct for CatalogV1alpha1QuotaStatusError
type CatalogV1alpha1QuotaStatusError struct {
	Type string `json:"type"`
	// Details of the error.
	Detail string `json:"detail"`
	// Time when the update occurred.
	Timestamp time.Time `json:"timestamp"`
	Meta map[string]interface{} `json:"meta,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CatalogV1alpha1QuotaStatusError CatalogV1alpha1QuotaStatusError

// NewCatalogV1alpha1QuotaStatusError instantiates a new CatalogV1alpha1QuotaStatusError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogV1alpha1QuotaStatusError(type_ string, detail string, timestamp time.Time) *CatalogV1alpha1QuotaStatusError {
	this := CatalogV1alpha1QuotaStatusError{}
	this.Type = type_
	this.Detail = detail
	this.Timestamp = timestamp
	return &this
}

// NewCatalogV1alpha1QuotaStatusErrorWithDefaults instantiates a new CatalogV1alpha1QuotaStatusError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogV1alpha1QuotaStatusErrorWithDefaults() *CatalogV1alpha1QuotaStatusError {
	this := CatalogV1alpha1QuotaStatusError{}
	return &this
}

// GetType returns the Type field value
func (o *CatalogV1alpha1QuotaStatusError) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1QuotaStatusError) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CatalogV1alpha1QuotaStatusError) SetType(v string) {
	o.Type = v
}

// GetDetail returns the Detail field value
func (o *CatalogV1alpha1QuotaStatusError) GetDetail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Detail
}

// GetDetailOk returns a tuple with the Detail field value
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1QuotaStatusError) GetDetailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Detail, true
}

// SetDetail sets field value
func (o *CatalogV1alpha1QuotaStatusError) SetDetail(v string) {
	o.Detail = v
}

// GetTimestamp returns the Timestamp field value
func (o *CatalogV1alpha1QuotaStatusError) GetTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1QuotaStatusError) GetTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *CatalogV1alpha1QuotaStatusError) SetTimestamp(v time.Time) {
	o.Timestamp = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *CatalogV1alpha1QuotaStatusError) GetMeta() map[string]interface{} {
	if o == nil || IsNil(o.Meta) {
		var ret map[string]interface{}
		return ret
	}
	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1QuotaStatusError) GetMetaOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Meta) {
		return map[string]interface{}{}, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *CatalogV1alpha1QuotaStatusError) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given map[string]interface{} and assigns it to the Meta field.
func (o *CatalogV1alpha1QuotaStatusError) SetMeta(v map[string]interface{}) {
	o.Meta = v
}

func (o CatalogV1alpha1QuotaStatusError) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogV1alpha1QuotaStatusError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["detail"] = o.Detail
	toSerialize["timestamp"] = o.Timestamp
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CatalogV1alpha1QuotaStatusError) UnmarshalJSON(data []byte) (err error) {
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

	varCatalogV1alpha1QuotaStatusError := _CatalogV1alpha1QuotaStatusError{}

	err = json.Unmarshal(data, &varCatalogV1alpha1QuotaStatusError)

	if err != nil {
		return err
	}

	*o = CatalogV1alpha1QuotaStatusError(varCatalogV1alpha1QuotaStatusError)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "detail")
		delete(additionalProperties, "timestamp")
		delete(additionalProperties, "meta")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCatalogV1alpha1QuotaStatusError struct {
	value *CatalogV1alpha1QuotaStatusError
	isSet bool
}

func (v NullableCatalogV1alpha1QuotaStatusError) Get() *CatalogV1alpha1QuotaStatusError {
	return v.value
}

func (v *NullableCatalogV1alpha1QuotaStatusError) Set(val *CatalogV1alpha1QuotaStatusError) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogV1alpha1QuotaStatusError) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogV1alpha1QuotaStatusError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogV1alpha1QuotaStatusError(val *CatalogV1alpha1QuotaStatusError) *NullableCatalogV1alpha1QuotaStatusError {
	return &NullableCatalogV1alpha1QuotaStatusError{value: val, isSet: true}
}

func (v NullableCatalogV1alpha1QuotaStatusError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogV1alpha1QuotaStatusError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


