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
)

// checks if the CatalogV1ProductReleaseSpecVersionProperties type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogV1ProductReleaseSpecVersionProperties{}

// CatalogV1ProductReleaseSpecVersionProperties Defines version related properties.
type CatalogV1ProductReleaseSpecVersionProperties struct {
	// Label for the generated version.
	Label *string `json:"label,omitempty"`
	// Description of the version.
	Description *string `json:"description,omitempty"`
}

// NewCatalogV1ProductReleaseSpecVersionProperties instantiates a new CatalogV1ProductReleaseSpecVersionProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogV1ProductReleaseSpecVersionProperties() *CatalogV1ProductReleaseSpecVersionProperties {
	this := CatalogV1ProductReleaseSpecVersionProperties{}
	return &this
}

// NewCatalogV1ProductReleaseSpecVersionPropertiesWithDefaults instantiates a new CatalogV1ProductReleaseSpecVersionProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogV1ProductReleaseSpecVersionPropertiesWithDefaults() *CatalogV1ProductReleaseSpecVersionProperties {
	this := CatalogV1ProductReleaseSpecVersionProperties{}
	return &this
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *CatalogV1ProductReleaseSpecVersionProperties) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1ProductReleaseSpecVersionProperties) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *CatalogV1ProductReleaseSpecVersionProperties) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *CatalogV1ProductReleaseSpecVersionProperties) SetLabel(v string) {
	o.Label = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CatalogV1ProductReleaseSpecVersionProperties) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1ProductReleaseSpecVersionProperties) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CatalogV1ProductReleaseSpecVersionProperties) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CatalogV1ProductReleaseSpecVersionProperties) SetDescription(v string) {
	o.Description = &v
}

func (o CatalogV1ProductReleaseSpecVersionProperties) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogV1ProductReleaseSpecVersionProperties) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

type NullableCatalogV1ProductReleaseSpecVersionProperties struct {
	value *CatalogV1ProductReleaseSpecVersionProperties
	isSet bool
}

func (v NullableCatalogV1ProductReleaseSpecVersionProperties) Get() *CatalogV1ProductReleaseSpecVersionProperties {
	return v.value
}

func (v *NullableCatalogV1ProductReleaseSpecVersionProperties) Set(val *CatalogV1ProductReleaseSpecVersionProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogV1ProductReleaseSpecVersionProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogV1ProductReleaseSpecVersionProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogV1ProductReleaseSpecVersionProperties(val *CatalogV1ProductReleaseSpecVersionProperties) *NullableCatalogV1ProductReleaseSpecVersionProperties {
	return &NullableCatalogV1ProductReleaseSpecVersionProperties{value: val, isSet: true}
}

func (v NullableCatalogV1ProductReleaseSpecVersionProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogV1ProductReleaseSpecVersionProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


