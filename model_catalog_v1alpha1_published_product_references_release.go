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

// checks if the CatalogV1alpha1PublishedProductReferencesRelease type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogV1alpha1PublishedProductReferencesRelease{}

// CatalogV1alpha1PublishedProductReferencesRelease The Product release that was published to the marketplace
type CatalogV1alpha1PublishedProductReferencesRelease struct {
	// Reference to product release
	Name *string `json:"name,omitempty"`
	// Product release version
	Version *string `json:"version,omitempty"`
	// Product release state
	State *string `json:"state,omitempty"`
}

// NewCatalogV1alpha1PublishedProductReferencesRelease instantiates a new CatalogV1alpha1PublishedProductReferencesRelease object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogV1alpha1PublishedProductReferencesRelease() *CatalogV1alpha1PublishedProductReferencesRelease {
	this := CatalogV1alpha1PublishedProductReferencesRelease{}
	return &this
}

// NewCatalogV1alpha1PublishedProductReferencesReleaseWithDefaults instantiates a new CatalogV1alpha1PublishedProductReferencesRelease object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogV1alpha1PublishedProductReferencesReleaseWithDefaults() *CatalogV1alpha1PublishedProductReferencesRelease {
	this := CatalogV1alpha1PublishedProductReferencesRelease{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CatalogV1alpha1PublishedProductReferencesRelease) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1PublishedProductReferencesRelease) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CatalogV1alpha1PublishedProductReferencesRelease) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CatalogV1alpha1PublishedProductReferencesRelease) SetName(v string) {
	o.Name = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *CatalogV1alpha1PublishedProductReferencesRelease) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1PublishedProductReferencesRelease) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *CatalogV1alpha1PublishedProductReferencesRelease) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *CatalogV1alpha1PublishedProductReferencesRelease) SetVersion(v string) {
	o.Version = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *CatalogV1alpha1PublishedProductReferencesRelease) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1PublishedProductReferencesRelease) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *CatalogV1alpha1PublishedProductReferencesRelease) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *CatalogV1alpha1PublishedProductReferencesRelease) SetState(v string) {
	o.State = &v
}

func (o CatalogV1alpha1PublishedProductReferencesRelease) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogV1alpha1PublishedProductReferencesRelease) ToMap() (map[string]interface{}, error) {
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

type NullableCatalogV1alpha1PublishedProductReferencesRelease struct {
	value *CatalogV1alpha1PublishedProductReferencesRelease
	isSet bool
}

func (v NullableCatalogV1alpha1PublishedProductReferencesRelease) Get() *CatalogV1alpha1PublishedProductReferencesRelease {
	return v.value
}

func (v *NullableCatalogV1alpha1PublishedProductReferencesRelease) Set(val *CatalogV1alpha1PublishedProductReferencesRelease) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogV1alpha1PublishedProductReferencesRelease) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogV1alpha1PublishedProductReferencesRelease) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogV1alpha1PublishedProductReferencesRelease(val *CatalogV1alpha1PublishedProductReferencesRelease) *NullableCatalogV1alpha1PublishedProductReferencesRelease {
	return &NullableCatalogV1alpha1PublishedProductReferencesRelease{value: val, isSet: true}
}

func (v NullableCatalogV1alpha1PublishedProductReferencesRelease) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogV1alpha1PublishedProductReferencesRelease) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


