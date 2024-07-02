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

// checks if the CatalogV1alpha1AccessControlListSpecSubjectsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogV1alpha1AccessControlListSpecSubjectsInner{}

// CatalogV1alpha1AccessControlListSpecSubjectsInner struct for CatalogV1alpha1AccessControlListSpecSubjectsInner
type CatalogV1alpha1AccessControlListSpecSubjectsInner struct {
	// Type of the subject
	Type *string `json:"type,omitempty"`
	// ID of the subject
	Id *string `json:"id,omitempty"`
}

// NewCatalogV1alpha1AccessControlListSpecSubjectsInner instantiates a new CatalogV1alpha1AccessControlListSpecSubjectsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogV1alpha1AccessControlListSpecSubjectsInner() *CatalogV1alpha1AccessControlListSpecSubjectsInner {
	this := CatalogV1alpha1AccessControlListSpecSubjectsInner{}
	return &this
}

// NewCatalogV1alpha1AccessControlListSpecSubjectsInnerWithDefaults instantiates a new CatalogV1alpha1AccessControlListSpecSubjectsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogV1alpha1AccessControlListSpecSubjectsInnerWithDefaults() *CatalogV1alpha1AccessControlListSpecSubjectsInner {
	this := CatalogV1alpha1AccessControlListSpecSubjectsInner{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CatalogV1alpha1AccessControlListSpecSubjectsInner) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1AccessControlListSpecSubjectsInner) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CatalogV1alpha1AccessControlListSpecSubjectsInner) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CatalogV1alpha1AccessControlListSpecSubjectsInner) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CatalogV1alpha1AccessControlListSpecSubjectsInner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1AccessControlListSpecSubjectsInner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CatalogV1alpha1AccessControlListSpecSubjectsInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CatalogV1alpha1AccessControlListSpecSubjectsInner) SetId(v string) {
	o.Id = &v
}

func (o CatalogV1alpha1AccessControlListSpecSubjectsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogV1alpha1AccessControlListSpecSubjectsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableCatalogV1alpha1AccessControlListSpecSubjectsInner struct {
	value *CatalogV1alpha1AccessControlListSpecSubjectsInner
	isSet bool
}

func (v NullableCatalogV1alpha1AccessControlListSpecSubjectsInner) Get() *CatalogV1alpha1AccessControlListSpecSubjectsInner {
	return v.value
}

func (v *NullableCatalogV1alpha1AccessControlListSpecSubjectsInner) Set(val *CatalogV1alpha1AccessControlListSpecSubjectsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogV1alpha1AccessControlListSpecSubjectsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogV1alpha1AccessControlListSpecSubjectsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogV1alpha1AccessControlListSpecSubjectsInner(val *CatalogV1alpha1AccessControlListSpecSubjectsInner) *NullableCatalogV1alpha1AccessControlListSpecSubjectsInner {
	return &NullableCatalogV1alpha1AccessControlListSpecSubjectsInner{value: val, isSet: true}
}

func (v NullableCatalogV1alpha1AccessControlListSpecSubjectsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogV1alpha1AccessControlListSpecSubjectsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

