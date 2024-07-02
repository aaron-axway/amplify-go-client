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

// checks if the CatalogV1alpha1PublishedProductReferences type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogV1alpha1PublishedProductReferences{}

// CatalogV1alpha1PublishedProductReferences struct for CatalogV1alpha1PublishedProductReferences
type CatalogV1alpha1PublishedProductReferences struct {
	Release CatalogV1alpha1PublishedProductReferencesRelease `json:"release"`
}

type _CatalogV1alpha1PublishedProductReferences CatalogV1alpha1PublishedProductReferences

// NewCatalogV1alpha1PublishedProductReferences instantiates a new CatalogV1alpha1PublishedProductReferences object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogV1alpha1PublishedProductReferences(release CatalogV1alpha1PublishedProductReferencesRelease) *CatalogV1alpha1PublishedProductReferences {
	this := CatalogV1alpha1PublishedProductReferences{}
	this.Release = release
	return &this
}

// NewCatalogV1alpha1PublishedProductReferencesWithDefaults instantiates a new CatalogV1alpha1PublishedProductReferences object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogV1alpha1PublishedProductReferencesWithDefaults() *CatalogV1alpha1PublishedProductReferences {
	this := CatalogV1alpha1PublishedProductReferences{}
	return &this
}

// GetRelease returns the Release field value
func (o *CatalogV1alpha1PublishedProductReferences) GetRelease() CatalogV1alpha1PublishedProductReferencesRelease {
	if o == nil {
		var ret CatalogV1alpha1PublishedProductReferencesRelease
		return ret
	}

	return o.Release
}

// GetReleaseOk returns a tuple with the Release field value
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1PublishedProductReferences) GetReleaseOk() (*CatalogV1alpha1PublishedProductReferencesRelease, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Release, true
}

// SetRelease sets field value
func (o *CatalogV1alpha1PublishedProductReferences) SetRelease(v CatalogV1alpha1PublishedProductReferencesRelease) {
	o.Release = v
}

func (o CatalogV1alpha1PublishedProductReferences) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogV1alpha1PublishedProductReferences) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["release"] = o.Release
	return toSerialize, nil
}

func (o *CatalogV1alpha1PublishedProductReferences) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"release",
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

	varCatalogV1alpha1PublishedProductReferences := _CatalogV1alpha1PublishedProductReferences{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCatalogV1alpha1PublishedProductReferences)

	if err != nil {
		return err
	}

	*o = CatalogV1alpha1PublishedProductReferences(varCatalogV1alpha1PublishedProductReferences)

	return err
}

type NullableCatalogV1alpha1PublishedProductReferences struct {
	value *CatalogV1alpha1PublishedProductReferences
	isSet bool
}

func (v NullableCatalogV1alpha1PublishedProductReferences) Get() *CatalogV1alpha1PublishedProductReferences {
	return v.value
}

func (v *NullableCatalogV1alpha1PublishedProductReferences) Set(val *CatalogV1alpha1PublishedProductReferences) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogV1alpha1PublishedProductReferences) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogV1alpha1PublishedProductReferences) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogV1alpha1PublishedProductReferences(val *CatalogV1alpha1PublishedProductReferences) *NullableCatalogV1alpha1PublishedProductReferences {
	return &NullableCatalogV1alpha1PublishedProductReferences{value: val, isSet: true}
}

func (v NullableCatalogV1alpha1PublishedProductReferences) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogV1alpha1PublishedProductReferences) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


