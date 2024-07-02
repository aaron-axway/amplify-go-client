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

// checks if the CatalogV1alpha1ConsumerProductVisibilitySpecOrgRef type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogV1alpha1ConsumerProductVisibilitySpecOrgRef{}

// CatalogV1alpha1ConsumerProductVisibilitySpecOrgRef struct for CatalogV1alpha1ConsumerProductVisibilitySpecOrgRef
type CatalogV1alpha1ConsumerProductVisibilitySpecOrgRef struct {
	Type string `json:"type"`
	// ID of the consumer organization.
	Id string `json:"id"`
}

type _CatalogV1alpha1ConsumerProductVisibilitySpecOrgRef CatalogV1alpha1ConsumerProductVisibilitySpecOrgRef

// NewCatalogV1alpha1ConsumerProductVisibilitySpecOrgRef instantiates a new CatalogV1alpha1ConsumerProductVisibilitySpecOrgRef object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogV1alpha1ConsumerProductVisibilitySpecOrgRef(type_ string, id string) *CatalogV1alpha1ConsumerProductVisibilitySpecOrgRef {
	this := CatalogV1alpha1ConsumerProductVisibilitySpecOrgRef{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewCatalogV1alpha1ConsumerProductVisibilitySpecOrgRefWithDefaults instantiates a new CatalogV1alpha1ConsumerProductVisibilitySpecOrgRef object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogV1alpha1ConsumerProductVisibilitySpecOrgRefWithDefaults() *CatalogV1alpha1ConsumerProductVisibilitySpecOrgRef {
	this := CatalogV1alpha1ConsumerProductVisibilitySpecOrgRef{}
	return &this
}

// GetType returns the Type field value
func (o *CatalogV1alpha1ConsumerProductVisibilitySpecOrgRef) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1ConsumerProductVisibilitySpecOrgRef) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CatalogV1alpha1ConsumerProductVisibilitySpecOrgRef) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *CatalogV1alpha1ConsumerProductVisibilitySpecOrgRef) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1ConsumerProductVisibilitySpecOrgRef) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CatalogV1alpha1ConsumerProductVisibilitySpecOrgRef) SetId(v string) {
	o.Id = v
}

func (o CatalogV1alpha1ConsumerProductVisibilitySpecOrgRef) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogV1alpha1ConsumerProductVisibilitySpecOrgRef) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *CatalogV1alpha1ConsumerProductVisibilitySpecOrgRef) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
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

	varCatalogV1alpha1ConsumerProductVisibilitySpecOrgRef := _CatalogV1alpha1ConsumerProductVisibilitySpecOrgRef{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCatalogV1alpha1ConsumerProductVisibilitySpecOrgRef)

	if err != nil {
		return err
	}

	*o = CatalogV1alpha1ConsumerProductVisibilitySpecOrgRef(varCatalogV1alpha1ConsumerProductVisibilitySpecOrgRef)

	return err
}

type NullableCatalogV1alpha1ConsumerProductVisibilitySpecOrgRef struct {
	value *CatalogV1alpha1ConsumerProductVisibilitySpecOrgRef
	isSet bool
}

func (v NullableCatalogV1alpha1ConsumerProductVisibilitySpecOrgRef) Get() *CatalogV1alpha1ConsumerProductVisibilitySpecOrgRef {
	return v.value
}

func (v *NullableCatalogV1alpha1ConsumerProductVisibilitySpecOrgRef) Set(val *CatalogV1alpha1ConsumerProductVisibilitySpecOrgRef) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogV1alpha1ConsumerProductVisibilitySpecOrgRef) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogV1alpha1ConsumerProductVisibilitySpecOrgRef) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogV1alpha1ConsumerProductVisibilitySpecOrgRef(val *CatalogV1alpha1ConsumerProductVisibilitySpecOrgRef) *NullableCatalogV1alpha1ConsumerProductVisibilitySpecOrgRef {
	return &NullableCatalogV1alpha1ConsumerProductVisibilitySpecOrgRef{value: val, isSet: true}
}

func (v NullableCatalogV1alpha1ConsumerProductVisibilitySpecOrgRef) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogV1alpha1ConsumerProductVisibilitySpecOrgRef) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


