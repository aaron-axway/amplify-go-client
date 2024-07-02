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

// checks if the CatalogV1alpha1ApplicationMarketplace type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogV1alpha1ApplicationMarketplace{}

// CatalogV1alpha1ApplicationMarketplace Details about marketplace in Application.
type CatalogV1alpha1ApplicationMarketplace struct {
	// The name of the Marketplace.
	Name string `json:"name"`
	Resource CatalogV1alpha1ApplicationMarketplaceResource `json:"resource"`
}

type _CatalogV1alpha1ApplicationMarketplace CatalogV1alpha1ApplicationMarketplace

// NewCatalogV1alpha1ApplicationMarketplace instantiates a new CatalogV1alpha1ApplicationMarketplace object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogV1alpha1ApplicationMarketplace(name string, resource CatalogV1alpha1ApplicationMarketplaceResource) *CatalogV1alpha1ApplicationMarketplace {
	this := CatalogV1alpha1ApplicationMarketplace{}
	this.Name = name
	this.Resource = resource
	return &this
}

// NewCatalogV1alpha1ApplicationMarketplaceWithDefaults instantiates a new CatalogV1alpha1ApplicationMarketplace object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogV1alpha1ApplicationMarketplaceWithDefaults() *CatalogV1alpha1ApplicationMarketplace {
	this := CatalogV1alpha1ApplicationMarketplace{}
	return &this
}

// GetName returns the Name field value
func (o *CatalogV1alpha1ApplicationMarketplace) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1ApplicationMarketplace) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CatalogV1alpha1ApplicationMarketplace) SetName(v string) {
	o.Name = v
}

// GetResource returns the Resource field value
func (o *CatalogV1alpha1ApplicationMarketplace) GetResource() CatalogV1alpha1ApplicationMarketplaceResource {
	if o == nil {
		var ret CatalogV1alpha1ApplicationMarketplaceResource
		return ret
	}

	return o.Resource
}

// GetResourceOk returns a tuple with the Resource field value
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1ApplicationMarketplace) GetResourceOk() (*CatalogV1alpha1ApplicationMarketplaceResource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Resource, true
}

// SetResource sets field value
func (o *CatalogV1alpha1ApplicationMarketplace) SetResource(v CatalogV1alpha1ApplicationMarketplaceResource) {
	o.Resource = v
}

func (o CatalogV1alpha1ApplicationMarketplace) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogV1alpha1ApplicationMarketplace) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["resource"] = o.Resource
	return toSerialize, nil
}

func (o *CatalogV1alpha1ApplicationMarketplace) UnmarshalJSON(data []byte) (err error) {
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

	varCatalogV1alpha1ApplicationMarketplace := _CatalogV1alpha1ApplicationMarketplace{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCatalogV1alpha1ApplicationMarketplace)

	if err != nil {
		return err
	}

	*o = CatalogV1alpha1ApplicationMarketplace(varCatalogV1alpha1ApplicationMarketplace)

	return err
}

type NullableCatalogV1alpha1ApplicationMarketplace struct {
	value *CatalogV1alpha1ApplicationMarketplace
	isSet bool
}

func (v NullableCatalogV1alpha1ApplicationMarketplace) Get() *CatalogV1alpha1ApplicationMarketplace {
	return v.value
}

func (v *NullableCatalogV1alpha1ApplicationMarketplace) Set(val *CatalogV1alpha1ApplicationMarketplace) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogV1alpha1ApplicationMarketplace) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogV1alpha1ApplicationMarketplace) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogV1alpha1ApplicationMarketplace(val *CatalogV1alpha1ApplicationMarketplace) *NullableCatalogV1alpha1ApplicationMarketplace {
	return &NullableCatalogV1alpha1ApplicationMarketplace{value: val, isSet: true}
}

func (v NullableCatalogV1alpha1ApplicationMarketplace) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogV1alpha1ApplicationMarketplace) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


