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

// checks if the ApiV1LanguagesResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiV1LanguagesResource{}

// ApiV1LanguagesResource Provides information about the default language in which the resource was defined
type ApiV1LanguagesResource struct {
	// Language code.
	Code string `json:"code"`
}

type _ApiV1LanguagesResource ApiV1LanguagesResource

// NewApiV1LanguagesResource instantiates a new ApiV1LanguagesResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiV1LanguagesResource(code string) *ApiV1LanguagesResource {
	this := ApiV1LanguagesResource{}
	this.Code = code
	return &this
}

// NewApiV1LanguagesResourceWithDefaults instantiates a new ApiV1LanguagesResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiV1LanguagesResourceWithDefaults() *ApiV1LanguagesResource {
	this := ApiV1LanguagesResource{}
	return &this
}

// GetCode returns the Code field value
func (o *ApiV1LanguagesResource) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *ApiV1LanguagesResource) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *ApiV1LanguagesResource) SetCode(v string) {
	o.Code = v
}

func (o ApiV1LanguagesResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiV1LanguagesResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	return toSerialize, nil
}

func (o *ApiV1LanguagesResource) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"code",
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

	varApiV1LanguagesResource := _ApiV1LanguagesResource{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiV1LanguagesResource)

	if err != nil {
		return err
	}

	*o = ApiV1LanguagesResource(varApiV1LanguagesResource)

	return err
}

type NullableApiV1LanguagesResource struct {
	value *ApiV1LanguagesResource
	isSet bool
}

func (v NullableApiV1LanguagesResource) Get() *ApiV1LanguagesResource {
	return v.value
}

func (v *NullableApiV1LanguagesResource) Set(val *ApiV1LanguagesResource) {
	v.value = val
	v.isSet = true
}

func (v NullableApiV1LanguagesResource) IsSet() bool {
	return v.isSet
}

func (v *NullableApiV1LanguagesResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiV1LanguagesResource(val *ApiV1LanguagesResource) *NullableApiV1LanguagesResource {
	return &NullableApiV1LanguagesResource{value: val, isSet: true}
}

func (v NullableApiV1LanguagesResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiV1LanguagesResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


