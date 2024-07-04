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
	"bytes"
	"fmt"
)

// checks if the CatalogV1ProductReleaseSpec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogV1ProductReleaseSpec{}

// CatalogV1ProductReleaseSpec struct for CatalogV1ProductReleaseSpec
type CatalogV1ProductReleaseSpec struct {
	// Description of the product when the release was generated.
	Description *string `json:"description,omitempty"`
	// Version of the release.
	Version string `json:"version" validate:"regexp=^(0|([1-9][0-9]*))(\\\\.(0|([1-9][0-9]*))){2}$"`
	VersionProperties *CatalogV1ProductReleaseSpecVersionProperties `json:"versionProperties,omitempty"`
	Product string `json:"product"`
	Assets []CatalogV1ProductReleaseSpecAssetsInner `json:"assets,omitempty"`
	ReleaseTag string `json:"releaseTag"`
	State *string `json:"state,omitempty"`
	// list of categories for the released product.
	Categories []string `json:"categories,omitempty"`
}

type _CatalogV1ProductReleaseSpec CatalogV1ProductReleaseSpec

// NewCatalogV1ProductReleaseSpec instantiates a new CatalogV1ProductReleaseSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogV1ProductReleaseSpec(version string, product string, releaseTag string) *CatalogV1ProductReleaseSpec {
	this := CatalogV1ProductReleaseSpec{}
	this.Version = version
	this.Product = product
	this.ReleaseTag = releaseTag
	return &this
}

// NewCatalogV1ProductReleaseSpecWithDefaults instantiates a new CatalogV1ProductReleaseSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogV1ProductReleaseSpecWithDefaults() *CatalogV1ProductReleaseSpec {
	this := CatalogV1ProductReleaseSpec{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CatalogV1ProductReleaseSpec) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1ProductReleaseSpec) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CatalogV1ProductReleaseSpec) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CatalogV1ProductReleaseSpec) SetDescription(v string) {
	o.Description = &v
}

// GetVersion returns the Version field value
func (o *CatalogV1ProductReleaseSpec) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *CatalogV1ProductReleaseSpec) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *CatalogV1ProductReleaseSpec) SetVersion(v string) {
	o.Version = v
}

// GetVersionProperties returns the VersionProperties field value if set, zero value otherwise.
func (o *CatalogV1ProductReleaseSpec) GetVersionProperties() CatalogV1ProductReleaseSpecVersionProperties {
	if o == nil || IsNil(o.VersionProperties) {
		var ret CatalogV1ProductReleaseSpecVersionProperties
		return ret
	}
	return *o.VersionProperties
}

// GetVersionPropertiesOk returns a tuple with the VersionProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1ProductReleaseSpec) GetVersionPropertiesOk() (*CatalogV1ProductReleaseSpecVersionProperties, bool) {
	if o == nil || IsNil(o.VersionProperties) {
		return nil, false
	}
	return o.VersionProperties, true
}

// HasVersionProperties returns a boolean if a field has been set.
func (o *CatalogV1ProductReleaseSpec) HasVersionProperties() bool {
	if o != nil && !IsNil(o.VersionProperties) {
		return true
	}

	return false
}

// SetVersionProperties gets a reference to the given CatalogV1ProductReleaseSpecVersionProperties and assigns it to the VersionProperties field.
func (o *CatalogV1ProductReleaseSpec) SetVersionProperties(v CatalogV1ProductReleaseSpecVersionProperties) {
	o.VersionProperties = &v
}

// GetProduct returns the Product field value
func (o *CatalogV1ProductReleaseSpec) GetProduct() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Product
}

// GetProductOk returns a tuple with the Product field value
// and a boolean to check if the value has been set.
func (o *CatalogV1ProductReleaseSpec) GetProductOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Product, true
}

// SetProduct sets field value
func (o *CatalogV1ProductReleaseSpec) SetProduct(v string) {
	o.Product = v
}

// GetAssets returns the Assets field value if set, zero value otherwise.
func (o *CatalogV1ProductReleaseSpec) GetAssets() []CatalogV1ProductReleaseSpecAssetsInner {
	if o == nil || IsNil(o.Assets) {
		var ret []CatalogV1ProductReleaseSpecAssetsInner
		return ret
	}
	return o.Assets
}

// GetAssetsOk returns a tuple with the Assets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1ProductReleaseSpec) GetAssetsOk() ([]CatalogV1ProductReleaseSpecAssetsInner, bool) {
	if o == nil || IsNil(o.Assets) {
		return nil, false
	}
	return o.Assets, true
}

// HasAssets returns a boolean if a field has been set.
func (o *CatalogV1ProductReleaseSpec) HasAssets() bool {
	if o != nil && !IsNil(o.Assets) {
		return true
	}

	return false
}

// SetAssets gets a reference to the given []CatalogV1ProductReleaseSpecAssetsInner and assigns it to the Assets field.
func (o *CatalogV1ProductReleaseSpec) SetAssets(v []CatalogV1ProductReleaseSpecAssetsInner) {
	o.Assets = v
}

// GetReleaseTag returns the ReleaseTag field value
func (o *CatalogV1ProductReleaseSpec) GetReleaseTag() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReleaseTag
}

// GetReleaseTagOk returns a tuple with the ReleaseTag field value
// and a boolean to check if the value has been set.
func (o *CatalogV1ProductReleaseSpec) GetReleaseTagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReleaseTag, true
}

// SetReleaseTag sets field value
func (o *CatalogV1ProductReleaseSpec) SetReleaseTag(v string) {
	o.ReleaseTag = v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *CatalogV1ProductReleaseSpec) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1ProductReleaseSpec) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *CatalogV1ProductReleaseSpec) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *CatalogV1ProductReleaseSpec) SetState(v string) {
	o.State = &v
}

// GetCategories returns the Categories field value if set, zero value otherwise.
func (o *CatalogV1ProductReleaseSpec) GetCategories() []string {
	if o == nil || IsNil(o.Categories) {
		var ret []string
		return ret
	}
	return o.Categories
}

// GetCategoriesOk returns a tuple with the Categories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1ProductReleaseSpec) GetCategoriesOk() ([]string, bool) {
	if o == nil || IsNil(o.Categories) {
		return nil, false
	}
	return o.Categories, true
}

// HasCategories returns a boolean if a field has been set.
func (o *CatalogV1ProductReleaseSpec) HasCategories() bool {
	if o != nil && !IsNil(o.Categories) {
		return true
	}

	return false
}

// SetCategories gets a reference to the given []string and assigns it to the Categories field.
func (o *CatalogV1ProductReleaseSpec) SetCategories(v []string) {
	o.Categories = v
}

func (o CatalogV1ProductReleaseSpec) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogV1ProductReleaseSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["version"] = o.Version
	if !IsNil(o.VersionProperties) {
		toSerialize["versionProperties"] = o.VersionProperties
	}
	toSerialize["product"] = o.Product
	if !IsNil(o.Assets) {
		toSerialize["assets"] = o.Assets
	}
	toSerialize["releaseTag"] = o.ReleaseTag
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.Categories) {
		toSerialize["categories"] = o.Categories
	}
	return toSerialize, nil
}

func (o *CatalogV1ProductReleaseSpec) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"version",
		"product",
		"releaseTag",
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

	varCatalogV1ProductReleaseSpec := _CatalogV1ProductReleaseSpec{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCatalogV1ProductReleaseSpec)

	if err != nil {
		return err
	}

	*o = CatalogV1ProductReleaseSpec(varCatalogV1ProductReleaseSpec)

	return err
}

type NullableCatalogV1ProductReleaseSpec struct {
	value *CatalogV1ProductReleaseSpec
	isSet bool
}

func (v NullableCatalogV1ProductReleaseSpec) Get() *CatalogV1ProductReleaseSpec {
	return v.value
}

func (v *NullableCatalogV1ProductReleaseSpec) Set(val *CatalogV1ProductReleaseSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogV1ProductReleaseSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogV1ProductReleaseSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogV1ProductReleaseSpec(val *CatalogV1ProductReleaseSpec) *NullableCatalogV1ProductReleaseSpec {
	return &NullableCatalogV1ProductReleaseSpec{value: val, isSet: true}
}

func (v NullableCatalogV1ProductReleaseSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogV1ProductReleaseSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


