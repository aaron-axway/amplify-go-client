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
)

// checks if the ApiV1LanguagesMetadataAdditionalInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiV1LanguagesMetadataAdditionalInner{}

// ApiV1LanguagesMetadataAdditionalInner struct for ApiV1LanguagesMetadataAdditionalInner
type ApiV1LanguagesMetadataAdditionalInner struct {
	// Language code.
	Code *string `json:"code,omitempty"`
	// Status of the language. * \"complete\" language available for the resource with no missing properties or properties that need to be reviewed. * \"toReview\" language has properties that are missing or need to be reviewed. * \"undefined\" language was set initially and removed later. 
	Status *string `json:"status,omitempty"`
}

// NewApiV1LanguagesMetadataAdditionalInner instantiates a new ApiV1LanguagesMetadataAdditionalInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiV1LanguagesMetadataAdditionalInner() *ApiV1LanguagesMetadataAdditionalInner {
	this := ApiV1LanguagesMetadataAdditionalInner{}
	return &this
}

// NewApiV1LanguagesMetadataAdditionalInnerWithDefaults instantiates a new ApiV1LanguagesMetadataAdditionalInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiV1LanguagesMetadataAdditionalInnerWithDefaults() *ApiV1LanguagesMetadataAdditionalInner {
	this := ApiV1LanguagesMetadataAdditionalInner{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ApiV1LanguagesMetadataAdditionalInner) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiV1LanguagesMetadataAdditionalInner) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ApiV1LanguagesMetadataAdditionalInner) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ApiV1LanguagesMetadataAdditionalInner) SetCode(v string) {
	o.Code = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ApiV1LanguagesMetadataAdditionalInner) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiV1LanguagesMetadataAdditionalInner) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ApiV1LanguagesMetadataAdditionalInner) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ApiV1LanguagesMetadataAdditionalInner) SetStatus(v string) {
	o.Status = &v
}

func (o ApiV1LanguagesMetadataAdditionalInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiV1LanguagesMetadataAdditionalInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableApiV1LanguagesMetadataAdditionalInner struct {
	value *ApiV1LanguagesMetadataAdditionalInner
	isSet bool
}

func (v NullableApiV1LanguagesMetadataAdditionalInner) Get() *ApiV1LanguagesMetadataAdditionalInner {
	return v.value
}

func (v *NullableApiV1LanguagesMetadataAdditionalInner) Set(val *ApiV1LanguagesMetadataAdditionalInner) {
	v.value = val
	v.isSet = true
}

func (v NullableApiV1LanguagesMetadataAdditionalInner) IsSet() bool {
	return v.isSet
}

func (v *NullableApiV1LanguagesMetadataAdditionalInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiV1LanguagesMetadataAdditionalInner(val *ApiV1LanguagesMetadataAdditionalInner) *NullableApiV1LanguagesMetadataAdditionalInner {
	return &NullableApiV1LanguagesMetadataAdditionalInner{value: val, isSet: true}
}

func (v NullableApiV1LanguagesMetadataAdditionalInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiV1LanguagesMetadataAdditionalInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

