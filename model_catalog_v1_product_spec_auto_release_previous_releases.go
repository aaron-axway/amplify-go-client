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

// checks if the CatalogV1ProductSpecAutoReleasePreviousReleases type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogV1ProductSpecAutoReleasePreviousReleases{}

// CatalogV1ProductSpecAutoReleasePreviousReleases struct for CatalogV1ProductSpecAutoReleasePreviousReleases
type CatalogV1ProductSpecAutoReleasePreviousReleases struct {
	// Updates all prior non-archived releases to the desired state.
	UpdateState *string `json:"updateState,omitempty"`
}

// NewCatalogV1ProductSpecAutoReleasePreviousReleases instantiates a new CatalogV1ProductSpecAutoReleasePreviousReleases object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogV1ProductSpecAutoReleasePreviousReleases() *CatalogV1ProductSpecAutoReleasePreviousReleases {
	this := CatalogV1ProductSpecAutoReleasePreviousReleases{}
	return &this
}

// NewCatalogV1ProductSpecAutoReleasePreviousReleasesWithDefaults instantiates a new CatalogV1ProductSpecAutoReleasePreviousReleases object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogV1ProductSpecAutoReleasePreviousReleasesWithDefaults() *CatalogV1ProductSpecAutoReleasePreviousReleases {
	this := CatalogV1ProductSpecAutoReleasePreviousReleases{}
	return &this
}

// GetUpdateState returns the UpdateState field value if set, zero value otherwise.
func (o *CatalogV1ProductSpecAutoReleasePreviousReleases) GetUpdateState() string {
	if o == nil || IsNil(o.UpdateState) {
		var ret string
		return ret
	}
	return *o.UpdateState
}

// GetUpdateStateOk returns a tuple with the UpdateState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1ProductSpecAutoReleasePreviousReleases) GetUpdateStateOk() (*string, bool) {
	if o == nil || IsNil(o.UpdateState) {
		return nil, false
	}
	return o.UpdateState, true
}

// HasUpdateState returns a boolean if a field has been set.
func (o *CatalogV1ProductSpecAutoReleasePreviousReleases) HasUpdateState() bool {
	if o != nil && !IsNil(o.UpdateState) {
		return true
	}

	return false
}

// SetUpdateState gets a reference to the given string and assigns it to the UpdateState field.
func (o *CatalogV1ProductSpecAutoReleasePreviousReleases) SetUpdateState(v string) {
	o.UpdateState = &v
}

func (o CatalogV1ProductSpecAutoReleasePreviousReleases) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogV1ProductSpecAutoReleasePreviousReleases) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UpdateState) {
		toSerialize["updateState"] = o.UpdateState
	}
	return toSerialize, nil
}

type NullableCatalogV1ProductSpecAutoReleasePreviousReleases struct {
	value *CatalogV1ProductSpecAutoReleasePreviousReleases
	isSet bool
}

func (v NullableCatalogV1ProductSpecAutoReleasePreviousReleases) Get() *CatalogV1ProductSpecAutoReleasePreviousReleases {
	return v.value
}

func (v *NullableCatalogV1ProductSpecAutoReleasePreviousReleases) Set(val *CatalogV1ProductSpecAutoReleasePreviousReleases) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogV1ProductSpecAutoReleasePreviousReleases) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogV1ProductSpecAutoReleasePreviousReleases) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogV1ProductSpecAutoReleasePreviousReleases(val *CatalogV1ProductSpecAutoReleasePreviousReleases) *NullableCatalogV1ProductSpecAutoReleasePreviousReleases {
	return &NullableCatalogV1ProductSpecAutoReleasePreviousReleases{value: val, isSet: true}
}

func (v NullableCatalogV1ProductSpecAutoReleasePreviousReleases) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogV1ProductSpecAutoReleasePreviousReleases) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

