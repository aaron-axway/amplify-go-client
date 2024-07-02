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

// checks if the CatalogV1CategorySpec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogV1CategorySpec{}

// CatalogV1CategorySpec struct for CatalogV1CategorySpec
type CatalogV1CategorySpec struct {
	// Markdown representing the category description.
	Description *string `json:"description,omitempty"`
	// Defines a parent category reference. Write access needed on the parent category to allow referencing it.
	ParentCategory *string `json:"parentCategory,omitempty"`
	Restriction *CatalogV1CategorySpecRestriction `json:"restriction,omitempty"`
}

// NewCatalogV1CategorySpec instantiates a new CatalogV1CategorySpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogV1CategorySpec() *CatalogV1CategorySpec {
	this := CatalogV1CategorySpec{}
	return &this
}

// NewCatalogV1CategorySpecWithDefaults instantiates a new CatalogV1CategorySpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogV1CategorySpecWithDefaults() *CatalogV1CategorySpec {
	this := CatalogV1CategorySpec{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CatalogV1CategorySpec) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1CategorySpec) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CatalogV1CategorySpec) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CatalogV1CategorySpec) SetDescription(v string) {
	o.Description = &v
}

// GetParentCategory returns the ParentCategory field value if set, zero value otherwise.
func (o *CatalogV1CategorySpec) GetParentCategory() string {
	if o == nil || IsNil(o.ParentCategory) {
		var ret string
		return ret
	}
	return *o.ParentCategory
}

// GetParentCategoryOk returns a tuple with the ParentCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1CategorySpec) GetParentCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.ParentCategory) {
		return nil, false
	}
	return o.ParentCategory, true
}

// HasParentCategory returns a boolean if a field has been set.
func (o *CatalogV1CategorySpec) HasParentCategory() bool {
	if o != nil && !IsNil(o.ParentCategory) {
		return true
	}

	return false
}

// SetParentCategory gets a reference to the given string and assigns it to the ParentCategory field.
func (o *CatalogV1CategorySpec) SetParentCategory(v string) {
	o.ParentCategory = &v
}

// GetRestriction returns the Restriction field value if set, zero value otherwise.
func (o *CatalogV1CategorySpec) GetRestriction() CatalogV1CategorySpecRestriction {
	if o == nil || IsNil(o.Restriction) {
		var ret CatalogV1CategorySpecRestriction
		return ret
	}
	return *o.Restriction
}

// GetRestrictionOk returns a tuple with the Restriction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1CategorySpec) GetRestrictionOk() (*CatalogV1CategorySpecRestriction, bool) {
	if o == nil || IsNil(o.Restriction) {
		return nil, false
	}
	return o.Restriction, true
}

// HasRestriction returns a boolean if a field has been set.
func (o *CatalogV1CategorySpec) HasRestriction() bool {
	if o != nil && !IsNil(o.Restriction) {
		return true
	}

	return false
}

// SetRestriction gets a reference to the given CatalogV1CategorySpecRestriction and assigns it to the Restriction field.
func (o *CatalogV1CategorySpec) SetRestriction(v CatalogV1CategorySpecRestriction) {
	o.Restriction = &v
}

func (o CatalogV1CategorySpec) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogV1CategorySpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.ParentCategory) {
		toSerialize["parentCategory"] = o.ParentCategory
	}
	if !IsNil(o.Restriction) {
		toSerialize["restriction"] = o.Restriction
	}
	return toSerialize, nil
}

type NullableCatalogV1CategorySpec struct {
	value *CatalogV1CategorySpec
	isSet bool
}

func (v NullableCatalogV1CategorySpec) Get() *CatalogV1CategorySpec {
	return v.value
}

func (v *NullableCatalogV1CategorySpec) Set(val *CatalogV1CategorySpec) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogV1CategorySpec) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogV1CategorySpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogV1CategorySpec(val *CatalogV1CategorySpec) *NullableCatalogV1CategorySpec {
	return &NullableCatalogV1CategorySpec{value: val, isSet: true}
}

func (v NullableCatalogV1CategorySpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogV1CategorySpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


