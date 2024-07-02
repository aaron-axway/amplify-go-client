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

// checks if the CatalogV1alpha1CredentialMarketplaceResourceOwner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogV1alpha1CredentialMarketplaceResourceOwner{}

// CatalogV1alpha1CredentialMarketplaceResourceOwner Owner of the Credential.
type CatalogV1alpha1CredentialMarketplaceResourceOwner struct {
	// The type of the owner.
	Type *string `json:"type,omitempty"`
	// Id of the owner of the resource.
	Id *string `json:"id,omitempty"`
	Organization CatalogV1alpha1ApplicationMarketplaceResourceOwnerOrganization `json:"organization"`
}

type _CatalogV1alpha1CredentialMarketplaceResourceOwner CatalogV1alpha1CredentialMarketplaceResourceOwner

// NewCatalogV1alpha1CredentialMarketplaceResourceOwner instantiates a new CatalogV1alpha1CredentialMarketplaceResourceOwner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogV1alpha1CredentialMarketplaceResourceOwner(organization CatalogV1alpha1ApplicationMarketplaceResourceOwnerOrganization) *CatalogV1alpha1CredentialMarketplaceResourceOwner {
	this := CatalogV1alpha1CredentialMarketplaceResourceOwner{}
	this.Organization = organization
	return &this
}

// NewCatalogV1alpha1CredentialMarketplaceResourceOwnerWithDefaults instantiates a new CatalogV1alpha1CredentialMarketplaceResourceOwner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogV1alpha1CredentialMarketplaceResourceOwnerWithDefaults() *CatalogV1alpha1CredentialMarketplaceResourceOwner {
	this := CatalogV1alpha1CredentialMarketplaceResourceOwner{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CatalogV1alpha1CredentialMarketplaceResourceOwner) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1CredentialMarketplaceResourceOwner) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CatalogV1alpha1CredentialMarketplaceResourceOwner) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CatalogV1alpha1CredentialMarketplaceResourceOwner) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CatalogV1alpha1CredentialMarketplaceResourceOwner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1CredentialMarketplaceResourceOwner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CatalogV1alpha1CredentialMarketplaceResourceOwner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CatalogV1alpha1CredentialMarketplaceResourceOwner) SetId(v string) {
	o.Id = &v
}

// GetOrganization returns the Organization field value
func (o *CatalogV1alpha1CredentialMarketplaceResourceOwner) GetOrganization() CatalogV1alpha1ApplicationMarketplaceResourceOwnerOrganization {
	if o == nil {
		var ret CatalogV1alpha1ApplicationMarketplaceResourceOwnerOrganization
		return ret
	}

	return o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1CredentialMarketplaceResourceOwner) GetOrganizationOk() (*CatalogV1alpha1ApplicationMarketplaceResourceOwnerOrganization, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Organization, true
}

// SetOrganization sets field value
func (o *CatalogV1alpha1CredentialMarketplaceResourceOwner) SetOrganization(v CatalogV1alpha1ApplicationMarketplaceResourceOwnerOrganization) {
	o.Organization = v
}

func (o CatalogV1alpha1CredentialMarketplaceResourceOwner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogV1alpha1CredentialMarketplaceResourceOwner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["organization"] = o.Organization
	return toSerialize, nil
}

func (o *CatalogV1alpha1CredentialMarketplaceResourceOwner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"organization",
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

	varCatalogV1alpha1CredentialMarketplaceResourceOwner := _CatalogV1alpha1CredentialMarketplaceResourceOwner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCatalogV1alpha1CredentialMarketplaceResourceOwner)

	if err != nil {
		return err
	}

	*o = CatalogV1alpha1CredentialMarketplaceResourceOwner(varCatalogV1alpha1CredentialMarketplaceResourceOwner)

	return err
}

type NullableCatalogV1alpha1CredentialMarketplaceResourceOwner struct {
	value *CatalogV1alpha1CredentialMarketplaceResourceOwner
	isSet bool
}

func (v NullableCatalogV1alpha1CredentialMarketplaceResourceOwner) Get() *CatalogV1alpha1CredentialMarketplaceResourceOwner {
	return v.value
}

func (v *NullableCatalogV1alpha1CredentialMarketplaceResourceOwner) Set(val *CatalogV1alpha1CredentialMarketplaceResourceOwner) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogV1alpha1CredentialMarketplaceResourceOwner) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogV1alpha1CredentialMarketplaceResourceOwner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogV1alpha1CredentialMarketplaceResourceOwner(val *CatalogV1alpha1CredentialMarketplaceResourceOwner) *NullableCatalogV1alpha1CredentialMarketplaceResourceOwner {
	return &NullableCatalogV1alpha1CredentialMarketplaceResourceOwner{value: val, isSet: true}
}

func (v NullableCatalogV1alpha1CredentialMarketplaceResourceOwner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogV1alpha1CredentialMarketplaceResourceOwner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


