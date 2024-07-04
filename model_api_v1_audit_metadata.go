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
)

// checks if the ApiV1AuditMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiV1AuditMetadata{}

// ApiV1AuditMetadata Read-only metadata for the record.
type ApiV1AuditMetadata struct {
	// The creation time.
	CreateTimestamp *time.Time `json:"createTimestamp,omitempty"`
	// Id of the user that created the entity.
	CreateUserId *string `json:"createUserId,omitempty"`
	// The last modification time.
	ModifyTimestamp *time.Time `json:"modifyTimestamp,omitempty"`
	// Id of the user that last modified the entity.
	ModifyUserId *string `json:"modifyUserId,omitempty"`
}

// NewApiV1AuditMetadata instantiates a new ApiV1AuditMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiV1AuditMetadata() *ApiV1AuditMetadata {
	this := ApiV1AuditMetadata{}
	return &this
}

// NewApiV1AuditMetadataWithDefaults instantiates a new ApiV1AuditMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiV1AuditMetadataWithDefaults() *ApiV1AuditMetadata {
	this := ApiV1AuditMetadata{}
	return &this
}

// GetCreateTimestamp returns the CreateTimestamp field value if set, zero value otherwise.
func (o *ApiV1AuditMetadata) GetCreateTimestamp() time.Time {
	if o == nil || IsNil(o.CreateTimestamp) {
		var ret time.Time
		return ret
	}
	return *o.CreateTimestamp
}

// GetCreateTimestampOk returns a tuple with the CreateTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiV1AuditMetadata) GetCreateTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreateTimestamp) {
		return nil, false
	}
	return o.CreateTimestamp, true
}

// HasCreateTimestamp returns a boolean if a field has been set.
func (o *ApiV1AuditMetadata) HasCreateTimestamp() bool {
	if o != nil && !IsNil(o.CreateTimestamp) {
		return true
	}

	return false
}

// SetCreateTimestamp gets a reference to the given time.Time and assigns it to the CreateTimestamp field.
func (o *ApiV1AuditMetadata) SetCreateTimestamp(v time.Time) {
	o.CreateTimestamp = &v
}

// GetCreateUserId returns the CreateUserId field value if set, zero value otherwise.
func (o *ApiV1AuditMetadata) GetCreateUserId() string {
	if o == nil || IsNil(o.CreateUserId) {
		var ret string
		return ret
	}
	return *o.CreateUserId
}

// GetCreateUserIdOk returns a tuple with the CreateUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiV1AuditMetadata) GetCreateUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.CreateUserId) {
		return nil, false
	}
	return o.CreateUserId, true
}

// HasCreateUserId returns a boolean if a field has been set.
func (o *ApiV1AuditMetadata) HasCreateUserId() bool {
	if o != nil && !IsNil(o.CreateUserId) {
		return true
	}

	return false
}

// SetCreateUserId gets a reference to the given string and assigns it to the CreateUserId field.
func (o *ApiV1AuditMetadata) SetCreateUserId(v string) {
	o.CreateUserId = &v
}

// GetModifyTimestamp returns the ModifyTimestamp field value if set, zero value otherwise.
func (o *ApiV1AuditMetadata) GetModifyTimestamp() time.Time {
	if o == nil || IsNil(o.ModifyTimestamp) {
		var ret time.Time
		return ret
	}
	return *o.ModifyTimestamp
}

// GetModifyTimestampOk returns a tuple with the ModifyTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiV1AuditMetadata) GetModifyTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ModifyTimestamp) {
		return nil, false
	}
	return o.ModifyTimestamp, true
}

// HasModifyTimestamp returns a boolean if a field has been set.
func (o *ApiV1AuditMetadata) HasModifyTimestamp() bool {
	if o != nil && !IsNil(o.ModifyTimestamp) {
		return true
	}

	return false
}

// SetModifyTimestamp gets a reference to the given time.Time and assigns it to the ModifyTimestamp field.
func (o *ApiV1AuditMetadata) SetModifyTimestamp(v time.Time) {
	o.ModifyTimestamp = &v
}

// GetModifyUserId returns the ModifyUserId field value if set, zero value otherwise.
func (o *ApiV1AuditMetadata) GetModifyUserId() string {
	if o == nil || IsNil(o.ModifyUserId) {
		var ret string
		return ret
	}
	return *o.ModifyUserId
}

// GetModifyUserIdOk returns a tuple with the ModifyUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiV1AuditMetadata) GetModifyUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.ModifyUserId) {
		return nil, false
	}
	return o.ModifyUserId, true
}

// HasModifyUserId returns a boolean if a field has been set.
func (o *ApiV1AuditMetadata) HasModifyUserId() bool {
	if o != nil && !IsNil(o.ModifyUserId) {
		return true
	}

	return false
}

// SetModifyUserId gets a reference to the given string and assigns it to the ModifyUserId field.
func (o *ApiV1AuditMetadata) SetModifyUserId(v string) {
	o.ModifyUserId = &v
}

func (o ApiV1AuditMetadata) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiV1AuditMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreateTimestamp) {
		toSerialize["createTimestamp"] = o.CreateTimestamp
	}
	if !IsNil(o.CreateUserId) {
		toSerialize["createUserId"] = o.CreateUserId
	}
	if !IsNil(o.ModifyTimestamp) {
		toSerialize["modifyTimestamp"] = o.ModifyTimestamp
	}
	if !IsNil(o.ModifyUserId) {
		toSerialize["modifyUserId"] = o.ModifyUserId
	}
	return toSerialize, nil
}

type NullableApiV1AuditMetadata struct {
	value *ApiV1AuditMetadata
	isSet bool
}

func (v NullableApiV1AuditMetadata) Get() *ApiV1AuditMetadata {
	return v.value
}

func (v *NullableApiV1AuditMetadata) Set(val *ApiV1AuditMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableApiV1AuditMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableApiV1AuditMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiV1AuditMetadata(val *ApiV1AuditMetadata) *NullableApiV1AuditMetadata {
	return &NullableApiV1AuditMetadata{value: val, isSet: true}
}

func (v NullableApiV1AuditMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiV1AuditMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


