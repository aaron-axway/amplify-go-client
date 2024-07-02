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

// checks if the CatalogV1alpha1AuthorizationProfilePolicies type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogV1alpha1AuthorizationProfilePolicies{}

// CatalogV1alpha1AuthorizationProfilePolicies struct for CatalogV1alpha1AuthorizationProfilePolicies
type CatalogV1alpha1AuthorizationProfilePolicies struct {
	Credentials *CatalogV1alpha1AuthorizationProfilePoliciesCredentials `json:"credentials,omitempty"`
}

// NewCatalogV1alpha1AuthorizationProfilePolicies instantiates a new CatalogV1alpha1AuthorizationProfilePolicies object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogV1alpha1AuthorizationProfilePolicies() *CatalogV1alpha1AuthorizationProfilePolicies {
	this := CatalogV1alpha1AuthorizationProfilePolicies{}
	return &this
}

// NewCatalogV1alpha1AuthorizationProfilePoliciesWithDefaults instantiates a new CatalogV1alpha1AuthorizationProfilePolicies object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogV1alpha1AuthorizationProfilePoliciesWithDefaults() *CatalogV1alpha1AuthorizationProfilePolicies {
	this := CatalogV1alpha1AuthorizationProfilePolicies{}
	return &this
}

// GetCredentials returns the Credentials field value if set, zero value otherwise.
func (o *CatalogV1alpha1AuthorizationProfilePolicies) GetCredentials() CatalogV1alpha1AuthorizationProfilePoliciesCredentials {
	if o == nil || IsNil(o.Credentials) {
		var ret CatalogV1alpha1AuthorizationProfilePoliciesCredentials
		return ret
	}
	return *o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1AuthorizationProfilePolicies) GetCredentialsOk() (*CatalogV1alpha1AuthorizationProfilePoliciesCredentials, bool) {
	if o == nil || IsNil(o.Credentials) {
		return nil, false
	}
	return o.Credentials, true
}

// HasCredentials returns a boolean if a field has been set.
func (o *CatalogV1alpha1AuthorizationProfilePolicies) HasCredentials() bool {
	if o != nil && !IsNil(o.Credentials) {
		return true
	}

	return false
}

// SetCredentials gets a reference to the given CatalogV1alpha1AuthorizationProfilePoliciesCredentials and assigns it to the Credentials field.
func (o *CatalogV1alpha1AuthorizationProfilePolicies) SetCredentials(v CatalogV1alpha1AuthorizationProfilePoliciesCredentials) {
	o.Credentials = &v
}

func (o CatalogV1alpha1AuthorizationProfilePolicies) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogV1alpha1AuthorizationProfilePolicies) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Credentials) {
		toSerialize["credentials"] = o.Credentials
	}
	return toSerialize, nil
}

type NullableCatalogV1alpha1AuthorizationProfilePolicies struct {
	value *CatalogV1alpha1AuthorizationProfilePolicies
	isSet bool
}

func (v NullableCatalogV1alpha1AuthorizationProfilePolicies) Get() *CatalogV1alpha1AuthorizationProfilePolicies {
	return v.value
}

func (v *NullableCatalogV1alpha1AuthorizationProfilePolicies) Set(val *CatalogV1alpha1AuthorizationProfilePolicies) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogV1alpha1AuthorizationProfilePolicies) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogV1alpha1AuthorizationProfilePolicies) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogV1alpha1AuthorizationProfilePolicies(val *CatalogV1alpha1AuthorizationProfilePolicies) *NullableCatalogV1alpha1AuthorizationProfilePolicies {
	return &NullableCatalogV1alpha1AuthorizationProfilePolicies{value: val, isSet: true}
}

func (v NullableCatalogV1alpha1AuthorizationProfilePolicies) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogV1alpha1AuthorizationProfilePolicies) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


