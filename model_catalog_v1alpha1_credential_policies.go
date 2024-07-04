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

// checks if the CatalogV1alpha1CredentialPolicies type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogV1alpha1CredentialPolicies{}

// CatalogV1alpha1CredentialPolicies struct for CatalogV1alpha1CredentialPolicies
type CatalogV1alpha1CredentialPolicies struct {
	Expiry *CatalogV1alpha1CredentialPoliciesExpiry `json:"expiry,omitempty"`
}

// NewCatalogV1alpha1CredentialPolicies instantiates a new CatalogV1alpha1CredentialPolicies object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogV1alpha1CredentialPolicies() *CatalogV1alpha1CredentialPolicies {
	this := CatalogV1alpha1CredentialPolicies{}
	return &this
}

// NewCatalogV1alpha1CredentialPoliciesWithDefaults instantiates a new CatalogV1alpha1CredentialPolicies object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogV1alpha1CredentialPoliciesWithDefaults() *CatalogV1alpha1CredentialPolicies {
	this := CatalogV1alpha1CredentialPolicies{}
	return &this
}

// GetExpiry returns the Expiry field value if set, zero value otherwise.
func (o *CatalogV1alpha1CredentialPolicies) GetExpiry() CatalogV1alpha1CredentialPoliciesExpiry {
	if o == nil || IsNil(o.Expiry) {
		var ret CatalogV1alpha1CredentialPoliciesExpiry
		return ret
	}
	return *o.Expiry
}

// GetExpiryOk returns a tuple with the Expiry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1CredentialPolicies) GetExpiryOk() (*CatalogV1alpha1CredentialPoliciesExpiry, bool) {
	if o == nil || IsNil(o.Expiry) {
		return nil, false
	}
	return o.Expiry, true
}

// HasExpiry returns a boolean if a field has been set.
func (o *CatalogV1alpha1CredentialPolicies) HasExpiry() bool {
	if o != nil && !IsNil(o.Expiry) {
		return true
	}

	return false
}

// SetExpiry gets a reference to the given CatalogV1alpha1CredentialPoliciesExpiry and assigns it to the Expiry field.
func (o *CatalogV1alpha1CredentialPolicies) SetExpiry(v CatalogV1alpha1CredentialPoliciesExpiry) {
	o.Expiry = &v
}

func (o CatalogV1alpha1CredentialPolicies) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogV1alpha1CredentialPolicies) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Expiry) {
		toSerialize["expiry"] = o.Expiry
	}
	return toSerialize, nil
}

type NullableCatalogV1alpha1CredentialPolicies struct {
	value *CatalogV1alpha1CredentialPolicies
	isSet bool
}

func (v NullableCatalogV1alpha1CredentialPolicies) Get() *CatalogV1alpha1CredentialPolicies {
	return v.value
}

func (v *NullableCatalogV1alpha1CredentialPolicies) Set(val *CatalogV1alpha1CredentialPolicies) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogV1alpha1CredentialPolicies) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogV1alpha1CredentialPolicies) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogV1alpha1CredentialPolicies(val *CatalogV1alpha1CredentialPolicies) *NullableCatalogV1alpha1CredentialPolicies {
	return &NullableCatalogV1alpha1CredentialPolicies{value: val, isSet: true}
}

func (v NullableCatalogV1alpha1CredentialPolicies) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogV1alpha1CredentialPolicies) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


