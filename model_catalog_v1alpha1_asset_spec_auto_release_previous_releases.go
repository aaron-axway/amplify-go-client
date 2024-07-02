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

// checks if the CatalogV1alpha1AssetSpecAutoReleasePreviousReleases type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogV1alpha1AssetSpecAutoReleasePreviousReleases{}

// CatalogV1alpha1AssetSpecAutoReleasePreviousReleases struct for CatalogV1alpha1AssetSpecAutoReleasePreviousReleases
type CatalogV1alpha1AssetSpecAutoReleasePreviousReleases struct {
	// Updates all prior non-archived releases to the desired state.
	UpdateState *string `json:"updateState,omitempty"`
}

// NewCatalogV1alpha1AssetSpecAutoReleasePreviousReleases instantiates a new CatalogV1alpha1AssetSpecAutoReleasePreviousReleases object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogV1alpha1AssetSpecAutoReleasePreviousReleases() *CatalogV1alpha1AssetSpecAutoReleasePreviousReleases {
	this := CatalogV1alpha1AssetSpecAutoReleasePreviousReleases{}
	return &this
}

// NewCatalogV1alpha1AssetSpecAutoReleasePreviousReleasesWithDefaults instantiates a new CatalogV1alpha1AssetSpecAutoReleasePreviousReleases object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogV1alpha1AssetSpecAutoReleasePreviousReleasesWithDefaults() *CatalogV1alpha1AssetSpecAutoReleasePreviousReleases {
	this := CatalogV1alpha1AssetSpecAutoReleasePreviousReleases{}
	return &this
}

// GetUpdateState returns the UpdateState field value if set, zero value otherwise.
func (o *CatalogV1alpha1AssetSpecAutoReleasePreviousReleases) GetUpdateState() string {
	if o == nil || IsNil(o.UpdateState) {
		var ret string
		return ret
	}
	return *o.UpdateState
}

// GetUpdateStateOk returns a tuple with the UpdateState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1AssetSpecAutoReleasePreviousReleases) GetUpdateStateOk() (*string, bool) {
	if o == nil || IsNil(o.UpdateState) {
		return nil, false
	}
	return o.UpdateState, true
}

// HasUpdateState returns a boolean if a field has been set.
func (o *CatalogV1alpha1AssetSpecAutoReleasePreviousReleases) HasUpdateState() bool {
	if o != nil && !IsNil(o.UpdateState) {
		return true
	}

	return false
}

// SetUpdateState gets a reference to the given string and assigns it to the UpdateState field.
func (o *CatalogV1alpha1AssetSpecAutoReleasePreviousReleases) SetUpdateState(v string) {
	o.UpdateState = &v
}

func (o CatalogV1alpha1AssetSpecAutoReleasePreviousReleases) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogV1alpha1AssetSpecAutoReleasePreviousReleases) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UpdateState) {
		toSerialize["updateState"] = o.UpdateState
	}
	return toSerialize, nil
}

type NullableCatalogV1alpha1AssetSpecAutoReleasePreviousReleases struct {
	value *CatalogV1alpha1AssetSpecAutoReleasePreviousReleases
	isSet bool
}

func (v NullableCatalogV1alpha1AssetSpecAutoReleasePreviousReleases) Get() *CatalogV1alpha1AssetSpecAutoReleasePreviousReleases {
	return v.value
}

func (v *NullableCatalogV1alpha1AssetSpecAutoReleasePreviousReleases) Set(val *CatalogV1alpha1AssetSpecAutoReleasePreviousReleases) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogV1alpha1AssetSpecAutoReleasePreviousReleases) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogV1alpha1AssetSpecAutoReleasePreviousReleases) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogV1alpha1AssetSpecAutoReleasePreviousReleases(val *CatalogV1alpha1AssetSpecAutoReleasePreviousReleases) *NullableCatalogV1alpha1AssetSpecAutoReleasePreviousReleases {
	return &NullableCatalogV1alpha1AssetSpecAutoReleasePreviousReleases{value: val, isSet: true}
}

func (v NullableCatalogV1alpha1AssetSpecAutoReleasePreviousReleases) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogV1alpha1AssetSpecAutoReleasePreviousReleases) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


