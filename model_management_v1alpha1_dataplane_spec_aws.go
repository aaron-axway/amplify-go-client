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

// checks if the ManagementV1alpha1DataplaneSpecAWS type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManagementV1alpha1DataplaneSpecAWS{}

// ManagementV1alpha1DataplaneSpecAWS The configuration common to all AWS agents that use this dataplane
type ManagementV1alpha1DataplaneSpecAWS struct {
	Type string `json:"type"`
	// The ARN of the cloud watch log resource that AWS API Gateway will be configured to send API Access data to
	AccessLogARN *string `json:"accessLogARN,omitempty" validate:"regexp=^arn:aws[a-zA-Z-]*:logs:[a-zA-Z0-9\\\\-]*:\\\\d{12}:log-group:[a-zA-Z0-9_\\\\-\\/\\\\.#]{1,512}$"`
	// If true, the discovery agent will enable full transaction logging for discovered API stages
	FullTransactionLogging *bool `json:"fullTransactionLogging,omitempty"`
}

type _ManagementV1alpha1DataplaneSpecAWS ManagementV1alpha1DataplaneSpecAWS

// NewManagementV1alpha1DataplaneSpecAWS instantiates a new ManagementV1alpha1DataplaneSpecAWS object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagementV1alpha1DataplaneSpecAWS(type_ string) *ManagementV1alpha1DataplaneSpecAWS {
	this := ManagementV1alpha1DataplaneSpecAWS{}
	this.Type = type_
	return &this
}

// NewManagementV1alpha1DataplaneSpecAWSWithDefaults instantiates a new ManagementV1alpha1DataplaneSpecAWS object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagementV1alpha1DataplaneSpecAWSWithDefaults() *ManagementV1alpha1DataplaneSpecAWS {
	this := ManagementV1alpha1DataplaneSpecAWS{}
	return &this
}

// GetType returns the Type field value
func (o *ManagementV1alpha1DataplaneSpecAWS) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1DataplaneSpecAWS) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ManagementV1alpha1DataplaneSpecAWS) SetType(v string) {
	o.Type = v
}

// GetAccessLogARN returns the AccessLogARN field value if set, zero value otherwise.
func (o *ManagementV1alpha1DataplaneSpecAWS) GetAccessLogARN() string {
	if o == nil || IsNil(o.AccessLogARN) {
		var ret string
		return ret
	}
	return *o.AccessLogARN
}

// GetAccessLogARNOk returns a tuple with the AccessLogARN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1DataplaneSpecAWS) GetAccessLogARNOk() (*string, bool) {
	if o == nil || IsNil(o.AccessLogARN) {
		return nil, false
	}
	return o.AccessLogARN, true
}

// HasAccessLogARN returns a boolean if a field has been set.
func (o *ManagementV1alpha1DataplaneSpecAWS) HasAccessLogARN() bool {
	if o != nil && !IsNil(o.AccessLogARN) {
		return true
	}

	return false
}

// SetAccessLogARN gets a reference to the given string and assigns it to the AccessLogARN field.
func (o *ManagementV1alpha1DataplaneSpecAWS) SetAccessLogARN(v string) {
	o.AccessLogARN = &v
}

// GetFullTransactionLogging returns the FullTransactionLogging field value if set, zero value otherwise.
func (o *ManagementV1alpha1DataplaneSpecAWS) GetFullTransactionLogging() bool {
	if o == nil || IsNil(o.FullTransactionLogging) {
		var ret bool
		return ret
	}
	return *o.FullTransactionLogging
}

// GetFullTransactionLoggingOk returns a tuple with the FullTransactionLogging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1DataplaneSpecAWS) GetFullTransactionLoggingOk() (*bool, bool) {
	if o == nil || IsNil(o.FullTransactionLogging) {
		return nil, false
	}
	return o.FullTransactionLogging, true
}

// HasFullTransactionLogging returns a boolean if a field has been set.
func (o *ManagementV1alpha1DataplaneSpecAWS) HasFullTransactionLogging() bool {
	if o != nil && !IsNil(o.FullTransactionLogging) {
		return true
	}

	return false
}

// SetFullTransactionLogging gets a reference to the given bool and assigns it to the FullTransactionLogging field.
func (o *ManagementV1alpha1DataplaneSpecAWS) SetFullTransactionLogging(v bool) {
	o.FullTransactionLogging = &v
}

func (o ManagementV1alpha1DataplaneSpecAWS) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManagementV1alpha1DataplaneSpecAWS) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.AccessLogARN) {
		toSerialize["accessLogARN"] = o.AccessLogARN
	}
	if !IsNil(o.FullTransactionLogging) {
		toSerialize["fullTransactionLogging"] = o.FullTransactionLogging
	}
	return toSerialize, nil
}

func (o *ManagementV1alpha1DataplaneSpecAWS) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
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

	varManagementV1alpha1DataplaneSpecAWS := _ManagementV1alpha1DataplaneSpecAWS{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varManagementV1alpha1DataplaneSpecAWS)

	if err != nil {
		return err
	}

	*o = ManagementV1alpha1DataplaneSpecAWS(varManagementV1alpha1DataplaneSpecAWS)

	return err
}

type NullableManagementV1alpha1DataplaneSpecAWS struct {
	value *ManagementV1alpha1DataplaneSpecAWS
	isSet bool
}

func (v NullableManagementV1alpha1DataplaneSpecAWS) Get() *ManagementV1alpha1DataplaneSpecAWS {
	return v.value
}

func (v *NullableManagementV1alpha1DataplaneSpecAWS) Set(val *ManagementV1alpha1DataplaneSpecAWS) {
	v.value = val
	v.isSet = true
}

func (v NullableManagementV1alpha1DataplaneSpecAWS) IsSet() bool {
	return v.isSet
}

func (v *NullableManagementV1alpha1DataplaneSpecAWS) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagementV1alpha1DataplaneSpecAWS(val *ManagementV1alpha1DataplaneSpecAWS) *NullableManagementV1alpha1DataplaneSpecAWS {
	return &NullableManagementV1alpha1DataplaneSpecAWS{value: val, isSet: true}
}

func (v NullableManagementV1alpha1DataplaneSpecAWS) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagementV1alpha1DataplaneSpecAWS) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


