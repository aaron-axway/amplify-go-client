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

// checks if the CatalogV1ProductReferencesAssetsInnerRelease type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogV1ProductReferencesAssetsInnerRelease{}

// CatalogV1ProductReferencesAssetsInnerRelease struct for CatalogV1ProductReferencesAssetsInnerRelease
type CatalogV1ProductReferencesAssetsInnerRelease struct {
	// The latest AssetRelease computed based on the provided Asset filters.
	Name *string `json:"name,omitempty"`
	// The AssetRelease version that the Product currently points to.
	Version *string `json:"version,omitempty"`
	// The AssetRelease state.
	State *string `json:"state,omitempty"`
	// APIService references belonging to referenced AssetReleases.
	ApiServices []string `json:"apiServices,omitempty"`
}

// NewCatalogV1ProductReferencesAssetsInnerRelease instantiates a new CatalogV1ProductReferencesAssetsInnerRelease object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogV1ProductReferencesAssetsInnerRelease() *CatalogV1ProductReferencesAssetsInnerRelease {
	this := CatalogV1ProductReferencesAssetsInnerRelease{}
	return &this
}

// NewCatalogV1ProductReferencesAssetsInnerReleaseWithDefaults instantiates a new CatalogV1ProductReferencesAssetsInnerRelease object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogV1ProductReferencesAssetsInnerReleaseWithDefaults() *CatalogV1ProductReferencesAssetsInnerRelease {
	this := CatalogV1ProductReferencesAssetsInnerRelease{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CatalogV1ProductReferencesAssetsInnerRelease) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1ProductReferencesAssetsInnerRelease) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CatalogV1ProductReferencesAssetsInnerRelease) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CatalogV1ProductReferencesAssetsInnerRelease) SetName(v string) {
	o.Name = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *CatalogV1ProductReferencesAssetsInnerRelease) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1ProductReferencesAssetsInnerRelease) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *CatalogV1ProductReferencesAssetsInnerRelease) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *CatalogV1ProductReferencesAssetsInnerRelease) SetVersion(v string) {
	o.Version = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *CatalogV1ProductReferencesAssetsInnerRelease) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1ProductReferencesAssetsInnerRelease) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *CatalogV1ProductReferencesAssetsInnerRelease) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *CatalogV1ProductReferencesAssetsInnerRelease) SetState(v string) {
	o.State = &v
}

// GetApiServices returns the ApiServices field value if set, zero value otherwise.
func (o *CatalogV1ProductReferencesAssetsInnerRelease) GetApiServices() []string {
	if o == nil || IsNil(o.ApiServices) {
		var ret []string
		return ret
	}
	return o.ApiServices
}

// GetApiServicesOk returns a tuple with the ApiServices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1ProductReferencesAssetsInnerRelease) GetApiServicesOk() ([]string, bool) {
	if o == nil || IsNil(o.ApiServices) {
		return nil, false
	}
	return o.ApiServices, true
}

// HasApiServices returns a boolean if a field has been set.
func (o *CatalogV1ProductReferencesAssetsInnerRelease) HasApiServices() bool {
	if o != nil && !IsNil(o.ApiServices) {
		return true
	}

	return false
}

// SetApiServices gets a reference to the given []string and assigns it to the ApiServices field.
func (o *CatalogV1ProductReferencesAssetsInnerRelease) SetApiServices(v []string) {
	o.ApiServices = v
}

func (o CatalogV1ProductReferencesAssetsInnerRelease) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogV1ProductReferencesAssetsInnerRelease) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.ApiServices) {
		toSerialize["apiServices"] = o.ApiServices
	}
	return toSerialize, nil
}

type NullableCatalogV1ProductReferencesAssetsInnerRelease struct {
	value *CatalogV1ProductReferencesAssetsInnerRelease
	isSet bool
}

func (v NullableCatalogV1ProductReferencesAssetsInnerRelease) Get() *CatalogV1ProductReferencesAssetsInnerRelease {
	return v.value
}

func (v *NullableCatalogV1ProductReferencesAssetsInnerRelease) Set(val *CatalogV1ProductReferencesAssetsInnerRelease) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogV1ProductReferencesAssetsInnerRelease) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogV1ProductReferencesAssetsInnerRelease) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogV1ProductReferencesAssetsInnerRelease(val *CatalogV1ProductReferencesAssetsInnerRelease) *NullableCatalogV1ProductReferencesAssetsInnerRelease {
	return &NullableCatalogV1ProductReferencesAssetsInnerRelease{value: val, isSet: true}
}

func (v NullableCatalogV1ProductReferencesAssetsInnerRelease) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogV1ProductReferencesAssetsInnerRelease) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


