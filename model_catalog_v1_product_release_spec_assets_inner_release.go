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

// checks if the CatalogV1ProductReleaseSpecAssetsInnerRelease type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogV1ProductReleaseSpecAssetsInnerRelease{}

// CatalogV1ProductReleaseSpecAssetsInnerRelease struct for CatalogV1ProductReleaseSpecAssetsInnerRelease
type CatalogV1ProductReleaseSpecAssetsInnerRelease struct {
	Name *string `json:"name,omitempty"`
	// The AssetRelease version.
	Version *string `json:"version,omitempty"`
	// The AssetRelease state.
	State *string `json:"state,omitempty"`
}

// NewCatalogV1ProductReleaseSpecAssetsInnerRelease instantiates a new CatalogV1ProductReleaseSpecAssetsInnerRelease object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogV1ProductReleaseSpecAssetsInnerRelease() *CatalogV1ProductReleaseSpecAssetsInnerRelease {
	this := CatalogV1ProductReleaseSpecAssetsInnerRelease{}
	return &this
}

// NewCatalogV1ProductReleaseSpecAssetsInnerReleaseWithDefaults instantiates a new CatalogV1ProductReleaseSpecAssetsInnerRelease object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogV1ProductReleaseSpecAssetsInnerReleaseWithDefaults() *CatalogV1ProductReleaseSpecAssetsInnerRelease {
	this := CatalogV1ProductReleaseSpecAssetsInnerRelease{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CatalogV1ProductReleaseSpecAssetsInnerRelease) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1ProductReleaseSpecAssetsInnerRelease) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CatalogV1ProductReleaseSpecAssetsInnerRelease) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CatalogV1ProductReleaseSpecAssetsInnerRelease) SetName(v string) {
	o.Name = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *CatalogV1ProductReleaseSpecAssetsInnerRelease) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1ProductReleaseSpecAssetsInnerRelease) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *CatalogV1ProductReleaseSpecAssetsInnerRelease) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *CatalogV1ProductReleaseSpecAssetsInnerRelease) SetVersion(v string) {
	o.Version = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *CatalogV1ProductReleaseSpecAssetsInnerRelease) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1ProductReleaseSpecAssetsInnerRelease) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *CatalogV1ProductReleaseSpecAssetsInnerRelease) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *CatalogV1ProductReleaseSpecAssetsInnerRelease) SetState(v string) {
	o.State = &v
}

func (o CatalogV1ProductReleaseSpecAssetsInnerRelease) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogV1ProductReleaseSpecAssetsInnerRelease) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	return toSerialize, nil
}

type NullableCatalogV1ProductReleaseSpecAssetsInnerRelease struct {
	value *CatalogV1ProductReleaseSpecAssetsInnerRelease
	isSet bool
}

func (v NullableCatalogV1ProductReleaseSpecAssetsInnerRelease) Get() *CatalogV1ProductReleaseSpecAssetsInnerRelease {
	return v.value
}

func (v *NullableCatalogV1ProductReleaseSpecAssetsInnerRelease) Set(val *CatalogV1ProductReleaseSpecAssetsInnerRelease) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogV1ProductReleaseSpecAssetsInnerRelease) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogV1ProductReleaseSpecAssetsInnerRelease) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogV1ProductReleaseSpecAssetsInnerRelease(val *CatalogV1ProductReleaseSpecAssetsInnerRelease) *NullableCatalogV1ProductReleaseSpecAssetsInnerRelease {
	return &NullableCatalogV1ProductReleaseSpecAssetsInnerRelease{value: val, isSet: true}
}

func (v NullableCatalogV1ProductReleaseSpecAssetsInnerRelease) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogV1ProductReleaseSpecAssetsInnerRelease) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


