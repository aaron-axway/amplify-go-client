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

// checks if the CatalogV1alpha1AssetLatestrelease type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogV1alpha1AssetLatestrelease{}

// CatalogV1alpha1AssetLatestrelease Provides newest non-archived release and version. Will be unassigned if no releases exist.
type CatalogV1alpha1AssetLatestrelease struct {
	Name *string `json:"name,omitempty"`
	Version *string `json:"version,omitempty"`
}

// NewCatalogV1alpha1AssetLatestrelease instantiates a new CatalogV1alpha1AssetLatestrelease object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogV1alpha1AssetLatestrelease() *CatalogV1alpha1AssetLatestrelease {
	this := CatalogV1alpha1AssetLatestrelease{}
	return &this
}

// NewCatalogV1alpha1AssetLatestreleaseWithDefaults instantiates a new CatalogV1alpha1AssetLatestrelease object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogV1alpha1AssetLatestreleaseWithDefaults() *CatalogV1alpha1AssetLatestrelease {
	this := CatalogV1alpha1AssetLatestrelease{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CatalogV1alpha1AssetLatestrelease) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1AssetLatestrelease) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CatalogV1alpha1AssetLatestrelease) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CatalogV1alpha1AssetLatestrelease) SetName(v string) {
	o.Name = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *CatalogV1alpha1AssetLatestrelease) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1AssetLatestrelease) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *CatalogV1alpha1AssetLatestrelease) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *CatalogV1alpha1AssetLatestrelease) SetVersion(v string) {
	o.Version = &v
}

func (o CatalogV1alpha1AssetLatestrelease) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogV1alpha1AssetLatestrelease) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableCatalogV1alpha1AssetLatestrelease struct {
	value *CatalogV1alpha1AssetLatestrelease
	isSet bool
}

func (v NullableCatalogV1alpha1AssetLatestrelease) Get() *CatalogV1alpha1AssetLatestrelease {
	return v.value
}

func (v *NullableCatalogV1alpha1AssetLatestrelease) Set(val *CatalogV1alpha1AssetLatestrelease) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogV1alpha1AssetLatestrelease) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogV1alpha1AssetLatestrelease) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogV1alpha1AssetLatestrelease(val *CatalogV1alpha1AssetLatestrelease) *NullableCatalogV1alpha1AssetLatestrelease {
	return &NullableCatalogV1alpha1AssetLatestrelease{value: val, isSet: true}
}

func (v NullableCatalogV1alpha1AssetLatestrelease) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogV1alpha1AssetLatestrelease) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


