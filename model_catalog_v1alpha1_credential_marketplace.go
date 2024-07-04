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

// checks if the CatalogV1alpha1CredentialMarketplace type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogV1alpha1CredentialMarketplace{}

// CatalogV1alpha1CredentialMarketplace Details about the marketplace Application.
type CatalogV1alpha1CredentialMarketplace struct {
	// The name of the Marketplace.
	Name string `json:"name"`
	Resource CatalogV1alpha1CredentialMarketplaceResource `json:"resource"`
}

type _CatalogV1alpha1CredentialMarketplace CatalogV1alpha1CredentialMarketplace

// NewCatalogV1alpha1CredentialMarketplace instantiates a new CatalogV1alpha1CredentialMarketplace object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogV1alpha1CredentialMarketplace(name string, resource CatalogV1alpha1CredentialMarketplaceResource) *CatalogV1alpha1CredentialMarketplace {
	this := CatalogV1alpha1CredentialMarketplace{}
	this.Name = name
	this.Resource = resource
	return &this
}

// NewCatalogV1alpha1CredentialMarketplaceWithDefaults instantiates a new CatalogV1alpha1CredentialMarketplace object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogV1alpha1CredentialMarketplaceWithDefaults() *CatalogV1alpha1CredentialMarketplace {
	this := CatalogV1alpha1CredentialMarketplace{}
	return &this
}

// GetName returns the Name field value
func (o *CatalogV1alpha1CredentialMarketplace) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1CredentialMarketplace) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CatalogV1alpha1CredentialMarketplace) SetName(v string) {
	o.Name = v
}

// GetResource returns the Resource field value
func (o *CatalogV1alpha1CredentialMarketplace) GetResource() CatalogV1alpha1CredentialMarketplaceResource {
	if o == nil {
		var ret CatalogV1alpha1CredentialMarketplaceResource
		return ret
	}

	return o.Resource
}

// GetResourceOk returns a tuple with the Resource field value
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1CredentialMarketplace) GetResourceOk() (*CatalogV1alpha1CredentialMarketplaceResource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Resource, true
}

// SetResource sets field value
func (o *CatalogV1alpha1CredentialMarketplace) SetResource(v CatalogV1alpha1CredentialMarketplaceResource) {
	o.Resource = v
}

func (o CatalogV1alpha1CredentialMarketplace) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogV1alpha1CredentialMarketplace) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["resource"] = o.Resource
	return toSerialize, nil
}

func (o *CatalogV1alpha1CredentialMarketplace) UnmarshalJSON(data []byte) (err error) {
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

	varCatalogV1alpha1CredentialMarketplace := _CatalogV1alpha1CredentialMarketplace{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCatalogV1alpha1CredentialMarketplace)

	if err != nil {
		return err
	}

	*o = CatalogV1alpha1CredentialMarketplace(varCatalogV1alpha1CredentialMarketplace)

	return err
}

type NullableCatalogV1alpha1CredentialMarketplace struct {
	value *CatalogV1alpha1CredentialMarketplace
	isSet bool
}

func (v NullableCatalogV1alpha1CredentialMarketplace) Get() *CatalogV1alpha1CredentialMarketplace {
	return v.value
}

func (v *NullableCatalogV1alpha1CredentialMarketplace) Set(val *CatalogV1alpha1CredentialMarketplace) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogV1alpha1CredentialMarketplace) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogV1alpha1CredentialMarketplace) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogV1alpha1CredentialMarketplace(val *CatalogV1alpha1CredentialMarketplace) *NullableCatalogV1alpha1CredentialMarketplace {
	return &NullableCatalogV1alpha1CredentialMarketplace{value: val, isSet: true}
}

func (v NullableCatalogV1alpha1CredentialMarketplace) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogV1alpha1CredentialMarketplace) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


