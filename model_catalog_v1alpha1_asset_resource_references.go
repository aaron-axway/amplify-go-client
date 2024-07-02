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

// checks if the CatalogV1alpha1AssetResourceReferences type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogV1alpha1AssetResourceReferences{}

// CatalogV1alpha1AssetResourceReferences struct for CatalogV1alpha1AssetResourceReferences
type CatalogV1alpha1AssetResourceReferences struct {
	// Reference to API Service Revision resource
	ApiServiceRevision *string `json:"apiServiceRevision,omitempty"`
	// Reference to API Service Instance resource
	ApiServiceInstance *string `json:"apiServiceInstance,omitempty"`
}

// NewCatalogV1alpha1AssetResourceReferences instantiates a new CatalogV1alpha1AssetResourceReferences object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogV1alpha1AssetResourceReferences() *CatalogV1alpha1AssetResourceReferences {
	this := CatalogV1alpha1AssetResourceReferences{}
	return &this
}

// NewCatalogV1alpha1AssetResourceReferencesWithDefaults instantiates a new CatalogV1alpha1AssetResourceReferences object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogV1alpha1AssetResourceReferencesWithDefaults() *CatalogV1alpha1AssetResourceReferences {
	this := CatalogV1alpha1AssetResourceReferences{}
	return &this
}

// GetApiServiceRevision returns the ApiServiceRevision field value if set, zero value otherwise.
func (o *CatalogV1alpha1AssetResourceReferences) GetApiServiceRevision() string {
	if o == nil || IsNil(o.ApiServiceRevision) {
		var ret string
		return ret
	}
	return *o.ApiServiceRevision
}

// GetApiServiceRevisionOk returns a tuple with the ApiServiceRevision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1AssetResourceReferences) GetApiServiceRevisionOk() (*string, bool) {
	if o == nil || IsNil(o.ApiServiceRevision) {
		return nil, false
	}
	return o.ApiServiceRevision, true
}

// HasApiServiceRevision returns a boolean if a field has been set.
func (o *CatalogV1alpha1AssetResourceReferences) HasApiServiceRevision() bool {
	if o != nil && !IsNil(o.ApiServiceRevision) {
		return true
	}

	return false
}

// SetApiServiceRevision gets a reference to the given string and assigns it to the ApiServiceRevision field.
func (o *CatalogV1alpha1AssetResourceReferences) SetApiServiceRevision(v string) {
	o.ApiServiceRevision = &v
}

// GetApiServiceInstance returns the ApiServiceInstance field value if set, zero value otherwise.
func (o *CatalogV1alpha1AssetResourceReferences) GetApiServiceInstance() string {
	if o == nil || IsNil(o.ApiServiceInstance) {
		var ret string
		return ret
	}
	return *o.ApiServiceInstance
}

// GetApiServiceInstanceOk returns a tuple with the ApiServiceInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1AssetResourceReferences) GetApiServiceInstanceOk() (*string, bool) {
	if o == nil || IsNil(o.ApiServiceInstance) {
		return nil, false
	}
	return o.ApiServiceInstance, true
}

// HasApiServiceInstance returns a boolean if a field has been set.
func (o *CatalogV1alpha1AssetResourceReferences) HasApiServiceInstance() bool {
	if o != nil && !IsNil(o.ApiServiceInstance) {
		return true
	}

	return false
}

// SetApiServiceInstance gets a reference to the given string and assigns it to the ApiServiceInstance field.
func (o *CatalogV1alpha1AssetResourceReferences) SetApiServiceInstance(v string) {
	o.ApiServiceInstance = &v
}

func (o CatalogV1alpha1AssetResourceReferences) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogV1alpha1AssetResourceReferences) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApiServiceRevision) {
		toSerialize["apiServiceRevision"] = o.ApiServiceRevision
	}
	if !IsNil(o.ApiServiceInstance) {
		toSerialize["apiServiceInstance"] = o.ApiServiceInstance
	}
	return toSerialize, nil
}

type NullableCatalogV1alpha1AssetResourceReferences struct {
	value *CatalogV1alpha1AssetResourceReferences
	isSet bool
}

func (v NullableCatalogV1alpha1AssetResourceReferences) Get() *CatalogV1alpha1AssetResourceReferences {
	return v.value
}

func (v *NullableCatalogV1alpha1AssetResourceReferences) Set(val *CatalogV1alpha1AssetResourceReferences) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogV1alpha1AssetResourceReferences) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogV1alpha1AssetResourceReferences) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogV1alpha1AssetResourceReferences(val *CatalogV1alpha1AssetResourceReferences) *NullableCatalogV1alpha1AssetResourceReferences {
	return &NullableCatalogV1alpha1AssetResourceReferences{value: val, isSet: true}
}

func (v NullableCatalogV1alpha1AssetResourceReferences) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogV1alpha1AssetResourceReferences) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


