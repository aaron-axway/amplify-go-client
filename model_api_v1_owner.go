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

// checks if the ApiV1Owner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiV1Owner{}

// ApiV1Owner Owner of the resource.
type ApiV1Owner struct {
	// The type of the owner. Defaults to team if not present.
	Type *string `json:"type,omitempty"`
	// Id of the owner of the resource.
	Id string `json:"id"`
}

type _ApiV1Owner ApiV1Owner

// NewApiV1Owner instantiates a new ApiV1Owner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiV1Owner(id string) *ApiV1Owner {
	this := ApiV1Owner{}
	var type_ string = "team"
	this.Type = &type_
	this.Id = id
	return &this
}

// NewApiV1OwnerWithDefaults instantiates a new ApiV1Owner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiV1OwnerWithDefaults() *ApiV1Owner {
	this := ApiV1Owner{}
	var type_ string = "team"
	this.Type = &type_
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ApiV1Owner) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiV1Owner) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ApiV1Owner) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ApiV1Owner) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value
func (o *ApiV1Owner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ApiV1Owner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ApiV1Owner) SetId(v string) {
	o.Id = v
}

func (o ApiV1Owner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiV1Owner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *ApiV1Owner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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

	varApiV1Owner := _ApiV1Owner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiV1Owner)

	if err != nil {
		return err
	}

	*o = ApiV1Owner(varApiV1Owner)

	return err
}

type NullableApiV1Owner struct {
	value *ApiV1Owner
	isSet bool
}

func (v NullableApiV1Owner) Get() *ApiV1Owner {
	return v.value
}

func (v *NullableApiV1Owner) Set(val *ApiV1Owner) {
	v.value = val
	v.isSet = true
}

func (v NullableApiV1Owner) IsSet() bool {
	return v.isSet
}

func (v *NullableApiV1Owner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiV1Owner(val *ApiV1Owner) *NullableApiV1Owner {
	return &NullableApiV1Owner{value: val, isSet: true}
}

func (v NullableApiV1Owner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiV1Owner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


