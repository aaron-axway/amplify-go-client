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

// checks if the CatalogV1alpha1AssetRequestMarketplace type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogV1alpha1AssetRequestMarketplace{}

// CatalogV1alpha1AssetRequestMarketplace Details about the marketplace Access Request.
type CatalogV1alpha1AssetRequestMarketplace struct {
	// The name of the Marketplace.
	Name string `json:"name"`
	Resource CatalogV1alpha1AssetRequestMarketplaceResource `json:"resource"`
}

type _CatalogV1alpha1AssetRequestMarketplace CatalogV1alpha1AssetRequestMarketplace

// NewCatalogV1alpha1AssetRequestMarketplace instantiates a new CatalogV1alpha1AssetRequestMarketplace object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogV1alpha1AssetRequestMarketplace(name string, resource CatalogV1alpha1AssetRequestMarketplaceResource) *CatalogV1alpha1AssetRequestMarketplace {
	this := CatalogV1alpha1AssetRequestMarketplace{}
	this.Name = name
	this.Resource = resource
	return &this
}

// NewCatalogV1alpha1AssetRequestMarketplaceWithDefaults instantiates a new CatalogV1alpha1AssetRequestMarketplace object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogV1alpha1AssetRequestMarketplaceWithDefaults() *CatalogV1alpha1AssetRequestMarketplace {
	this := CatalogV1alpha1AssetRequestMarketplace{}
	return &this
}

// GetName returns the Name field value
func (o *CatalogV1alpha1AssetRequestMarketplace) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1AssetRequestMarketplace) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CatalogV1alpha1AssetRequestMarketplace) SetName(v string) {
	o.Name = v
}

// GetResource returns the Resource field value
func (o *CatalogV1alpha1AssetRequestMarketplace) GetResource() CatalogV1alpha1AssetRequestMarketplaceResource {
	if o == nil {
		var ret CatalogV1alpha1AssetRequestMarketplaceResource
		return ret
	}

	return o.Resource
}

// GetResourceOk returns a tuple with the Resource field value
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1AssetRequestMarketplace) GetResourceOk() (*CatalogV1alpha1AssetRequestMarketplaceResource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Resource, true
}

// SetResource sets field value
func (o *CatalogV1alpha1AssetRequestMarketplace) SetResource(v CatalogV1alpha1AssetRequestMarketplaceResource) {
	o.Resource = v
}

func (o CatalogV1alpha1AssetRequestMarketplace) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogV1alpha1AssetRequestMarketplace) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["resource"] = o.Resource
	return toSerialize, nil
}

func (o *CatalogV1alpha1AssetRequestMarketplace) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"resource",
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

	varCatalogV1alpha1AssetRequestMarketplace := _CatalogV1alpha1AssetRequestMarketplace{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCatalogV1alpha1AssetRequestMarketplace)

	if err != nil {
		return err
	}

	*o = CatalogV1alpha1AssetRequestMarketplace(varCatalogV1alpha1AssetRequestMarketplace)

	return err
}

type NullableCatalogV1alpha1AssetRequestMarketplace struct {
	value *CatalogV1alpha1AssetRequestMarketplace
	isSet bool
}

func (v NullableCatalogV1alpha1AssetRequestMarketplace) Get() *CatalogV1alpha1AssetRequestMarketplace {
	return v.value
}

func (v *NullableCatalogV1alpha1AssetRequestMarketplace) Set(val *CatalogV1alpha1AssetRequestMarketplace) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogV1alpha1AssetRequestMarketplace) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogV1alpha1AssetRequestMarketplace) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogV1alpha1AssetRequestMarketplace(val *CatalogV1alpha1AssetRequestMarketplace) *NullableCatalogV1alpha1AssetRequestMarketplace {
	return &NullableCatalogV1alpha1AssetRequestMarketplace{value: val, isSet: true}
}

func (v NullableCatalogV1alpha1AssetRequestMarketplace) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogV1alpha1AssetRequestMarketplace) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

