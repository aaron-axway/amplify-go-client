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

// checks if the CatalogV1alpha1ProductVisibilitySpecTeamRef type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogV1alpha1ProductVisibilitySpecTeamRef{}

// CatalogV1alpha1ProductVisibilitySpecTeamRef struct for CatalogV1alpha1ProductVisibilitySpecTeamRef
type CatalogV1alpha1ProductVisibilitySpecTeamRef struct {
	Type string `json:"type"`
	// ID of the subject
	Id string `json:"id"`
}

type _CatalogV1alpha1ProductVisibilitySpecTeamRef CatalogV1alpha1ProductVisibilitySpecTeamRef

// NewCatalogV1alpha1ProductVisibilitySpecTeamRef instantiates a new CatalogV1alpha1ProductVisibilitySpecTeamRef object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogV1alpha1ProductVisibilitySpecTeamRef(type_ string, id string) *CatalogV1alpha1ProductVisibilitySpecTeamRef {
	this := CatalogV1alpha1ProductVisibilitySpecTeamRef{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewCatalogV1alpha1ProductVisibilitySpecTeamRefWithDefaults instantiates a new CatalogV1alpha1ProductVisibilitySpecTeamRef object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogV1alpha1ProductVisibilitySpecTeamRefWithDefaults() *CatalogV1alpha1ProductVisibilitySpecTeamRef {
	this := CatalogV1alpha1ProductVisibilitySpecTeamRef{}
	return &this
}

// GetType returns the Type field value
func (o *CatalogV1alpha1ProductVisibilitySpecTeamRef) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1ProductVisibilitySpecTeamRef) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CatalogV1alpha1ProductVisibilitySpecTeamRef) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *CatalogV1alpha1ProductVisibilitySpecTeamRef) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1ProductVisibilitySpecTeamRef) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CatalogV1alpha1ProductVisibilitySpecTeamRef) SetId(v string) {
	o.Id = v
}

func (o CatalogV1alpha1ProductVisibilitySpecTeamRef) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogV1alpha1ProductVisibilitySpecTeamRef) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *CatalogV1alpha1ProductVisibilitySpecTeamRef) UnmarshalJSON(data []byte) (err error) {
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

	varCatalogV1alpha1ProductVisibilitySpecTeamRef := _CatalogV1alpha1ProductVisibilitySpecTeamRef{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCatalogV1alpha1ProductVisibilitySpecTeamRef)

	if err != nil {
		return err
	}

	*o = CatalogV1alpha1ProductVisibilitySpecTeamRef(varCatalogV1alpha1ProductVisibilitySpecTeamRef)

	return err
}

type NullableCatalogV1alpha1ProductVisibilitySpecTeamRef struct {
	value *CatalogV1alpha1ProductVisibilitySpecTeamRef
	isSet bool
}

func (v NullableCatalogV1alpha1ProductVisibilitySpecTeamRef) Get() *CatalogV1alpha1ProductVisibilitySpecTeamRef {
	return v.value
}

func (v *NullableCatalogV1alpha1ProductVisibilitySpecTeamRef) Set(val *CatalogV1alpha1ProductVisibilitySpecTeamRef) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogV1alpha1ProductVisibilitySpecTeamRef) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogV1alpha1ProductVisibilitySpecTeamRef) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogV1alpha1ProductVisibilitySpecTeamRef(val *CatalogV1alpha1ProductVisibilitySpecTeamRef) *NullableCatalogV1alpha1ProductVisibilitySpecTeamRef {
	return &NullableCatalogV1alpha1ProductVisibilitySpecTeamRef{value: val, isSet: true}
}

func (v NullableCatalogV1alpha1ProductVisibilitySpecTeamRef) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogV1alpha1ProductVisibilitySpecTeamRef) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


