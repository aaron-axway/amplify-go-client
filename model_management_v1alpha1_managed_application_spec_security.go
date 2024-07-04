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

// checks if the ManagementV1alpha1ManagedApplicationSpecSecurity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManagementV1alpha1ManagedApplicationSpecSecurity{}

// ManagementV1alpha1ManagedApplicationSpecSecurity struct for ManagementV1alpha1ManagedApplicationSpecSecurity
type ManagementV1alpha1ManagedApplicationSpecSecurity struct {
	// public key to be used to encrypt the credentials linked to this Managed Application.
	EncryptionKey string `json:"encryptionKey"`
	EncryptionAlgorithm *string `json:"encryptionAlgorithm,omitempty"`
	EncryptionHash *string `json:"encryptionHash,omitempty"`
}

type _ManagementV1alpha1ManagedApplicationSpecSecurity ManagementV1alpha1ManagedApplicationSpecSecurity

// NewManagementV1alpha1ManagedApplicationSpecSecurity instantiates a new ManagementV1alpha1ManagedApplicationSpecSecurity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagementV1alpha1ManagedApplicationSpecSecurity(encryptionKey string) *ManagementV1alpha1ManagedApplicationSpecSecurity {
	this := ManagementV1alpha1ManagedApplicationSpecSecurity{}
	this.EncryptionKey = encryptionKey
	return &this
}

// NewManagementV1alpha1ManagedApplicationSpecSecurityWithDefaults instantiates a new ManagementV1alpha1ManagedApplicationSpecSecurity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagementV1alpha1ManagedApplicationSpecSecurityWithDefaults() *ManagementV1alpha1ManagedApplicationSpecSecurity {
	this := ManagementV1alpha1ManagedApplicationSpecSecurity{}
	return &this
}

// GetEncryptionKey returns the EncryptionKey field value
func (o *ManagementV1alpha1ManagedApplicationSpecSecurity) GetEncryptionKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EncryptionKey
}

// GetEncryptionKeyOk returns a tuple with the EncryptionKey field value
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1ManagedApplicationSpecSecurity) GetEncryptionKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EncryptionKey, true
}

// SetEncryptionKey sets field value
func (o *ManagementV1alpha1ManagedApplicationSpecSecurity) SetEncryptionKey(v string) {
	o.EncryptionKey = v
}

// GetEncryptionAlgorithm returns the EncryptionAlgorithm field value if set, zero value otherwise.
func (o *ManagementV1alpha1ManagedApplicationSpecSecurity) GetEncryptionAlgorithm() string {
	if o == nil || IsNil(o.EncryptionAlgorithm) {
		var ret string
		return ret
	}
	return *o.EncryptionAlgorithm
}

// GetEncryptionAlgorithmOk returns a tuple with the EncryptionAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1ManagedApplicationSpecSecurity) GetEncryptionAlgorithmOk() (*string, bool) {
	if o == nil || IsNil(o.EncryptionAlgorithm) {
		return nil, false
	}
	return o.EncryptionAlgorithm, true
}

// HasEncryptionAlgorithm returns a boolean if a field has been set.
func (o *ManagementV1alpha1ManagedApplicationSpecSecurity) HasEncryptionAlgorithm() bool {
	if o != nil && !IsNil(o.EncryptionAlgorithm) {
		return true
	}

	return false
}

// SetEncryptionAlgorithm gets a reference to the given string and assigns it to the EncryptionAlgorithm field.
func (o *ManagementV1alpha1ManagedApplicationSpecSecurity) SetEncryptionAlgorithm(v string) {
	o.EncryptionAlgorithm = &v
}

// GetEncryptionHash returns the EncryptionHash field value if set, zero value otherwise.
func (o *ManagementV1alpha1ManagedApplicationSpecSecurity) GetEncryptionHash() string {
	if o == nil || IsNil(o.EncryptionHash) {
		var ret string
		return ret
	}
	return *o.EncryptionHash
}

// GetEncryptionHashOk returns a tuple with the EncryptionHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1ManagedApplicationSpecSecurity) GetEncryptionHashOk() (*string, bool) {
	if o == nil || IsNil(o.EncryptionHash) {
		return nil, false
	}
	return o.EncryptionHash, true
}

// HasEncryptionHash returns a boolean if a field has been set.
func (o *ManagementV1alpha1ManagedApplicationSpecSecurity) HasEncryptionHash() bool {
	if o != nil && !IsNil(o.EncryptionHash) {
		return true
	}

	return false
}

// SetEncryptionHash gets a reference to the given string and assigns it to the EncryptionHash field.
func (o *ManagementV1alpha1ManagedApplicationSpecSecurity) SetEncryptionHash(v string) {
	o.EncryptionHash = &v
}

func (o ManagementV1alpha1ManagedApplicationSpecSecurity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManagementV1alpha1ManagedApplicationSpecSecurity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["encryptionKey"] = o.EncryptionKey
	if !IsNil(o.EncryptionAlgorithm) {
		toSerialize["encryptionAlgorithm"] = o.EncryptionAlgorithm
	}
	if !IsNil(o.EncryptionHash) {
		toSerialize["encryptionHash"] = o.EncryptionHash
	}
	return toSerialize, nil
}

func (o *ManagementV1alpha1ManagedApplicationSpecSecurity) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"encryptionKey",
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

	varManagementV1alpha1ManagedApplicationSpecSecurity := _ManagementV1alpha1ManagedApplicationSpecSecurity{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varManagementV1alpha1ManagedApplicationSpecSecurity)

	if err != nil {
		return err
	}

	*o = ManagementV1alpha1ManagedApplicationSpecSecurity(varManagementV1alpha1ManagedApplicationSpecSecurity)

	return err
}

type NullableManagementV1alpha1ManagedApplicationSpecSecurity struct {
	value *ManagementV1alpha1ManagedApplicationSpecSecurity
	isSet bool
}

func (v NullableManagementV1alpha1ManagedApplicationSpecSecurity) Get() *ManagementV1alpha1ManagedApplicationSpecSecurity {
	return v.value
}

func (v *NullableManagementV1alpha1ManagedApplicationSpecSecurity) Set(val *ManagementV1alpha1ManagedApplicationSpecSecurity) {
	v.value = val
	v.isSet = true
}

func (v NullableManagementV1alpha1ManagedApplicationSpecSecurity) IsSet() bool {
	return v.isSet
}

func (v *NullableManagementV1alpha1ManagedApplicationSpecSecurity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagementV1alpha1ManagedApplicationSpecSecurity(val *ManagementV1alpha1ManagedApplicationSpecSecurity) *NullableManagementV1alpha1ManagedApplicationSpecSecurity {
	return &NullableManagementV1alpha1ManagedApplicationSpecSecurity{value: val, isSet: true}
}

func (v NullableManagementV1alpha1ManagedApplicationSpecSecurity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagementV1alpha1ManagedApplicationSpecSecurity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


