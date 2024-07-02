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

// checks if the CatalogV1alpha1AssetMappingStatusSourceResourceApiServiceInstance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogV1alpha1AssetMappingStatusSourceResourceApiServiceInstance{}

// CatalogV1alpha1AssetMappingStatusSourceResourceApiServiceInstance struct for CatalogV1alpha1AssetMappingStatusSourceResourceApiServiceInstance
type CatalogV1alpha1AssetMappingStatusSourceResourceApiServiceInstance struct {
	Ref *string `json:"ref,omitempty"`
	OperationType *CatalogV1alpha1AssetMappingStatusOperationType `json:"operationType,omitempty"`
}

// NewCatalogV1alpha1AssetMappingStatusSourceResourceApiServiceInstance instantiates a new CatalogV1alpha1AssetMappingStatusSourceResourceApiServiceInstance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogV1alpha1AssetMappingStatusSourceResourceApiServiceInstance() *CatalogV1alpha1AssetMappingStatusSourceResourceApiServiceInstance {
	this := CatalogV1alpha1AssetMappingStatusSourceResourceApiServiceInstance{}
	return &this
}

// NewCatalogV1alpha1AssetMappingStatusSourceResourceApiServiceInstanceWithDefaults instantiates a new CatalogV1alpha1AssetMappingStatusSourceResourceApiServiceInstance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogV1alpha1AssetMappingStatusSourceResourceApiServiceInstanceWithDefaults() *CatalogV1alpha1AssetMappingStatusSourceResourceApiServiceInstance {
	this := CatalogV1alpha1AssetMappingStatusSourceResourceApiServiceInstance{}
	return &this
}

// GetRef returns the Ref field value if set, zero value otherwise.
func (o *CatalogV1alpha1AssetMappingStatusSourceResourceApiServiceInstance) GetRef() string {
	if o == nil || IsNil(o.Ref) {
		var ret string
		return ret
	}
	return *o.Ref
}

// GetRefOk returns a tuple with the Ref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1AssetMappingStatusSourceResourceApiServiceInstance) GetRefOk() (*string, bool) {
	if o == nil || IsNil(o.Ref) {
		return nil, false
	}
	return o.Ref, true
}

// HasRef returns a boolean if a field has been set.
func (o *CatalogV1alpha1AssetMappingStatusSourceResourceApiServiceInstance) HasRef() bool {
	if o != nil && !IsNil(o.Ref) {
		return true
	}

	return false
}

// SetRef gets a reference to the given string and assigns it to the Ref field.
func (o *CatalogV1alpha1AssetMappingStatusSourceResourceApiServiceInstance) SetRef(v string) {
	o.Ref = &v
}

// GetOperationType returns the OperationType field value if set, zero value otherwise.
func (o *CatalogV1alpha1AssetMappingStatusSourceResourceApiServiceInstance) GetOperationType() CatalogV1alpha1AssetMappingStatusOperationType {
	if o == nil || IsNil(o.OperationType) {
		var ret CatalogV1alpha1AssetMappingStatusOperationType
		return ret
	}
	return *o.OperationType
}

// GetOperationTypeOk returns a tuple with the OperationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1AssetMappingStatusSourceResourceApiServiceInstance) GetOperationTypeOk() (*CatalogV1alpha1AssetMappingStatusOperationType, bool) {
	if o == nil || IsNil(o.OperationType) {
		return nil, false
	}
	return o.OperationType, true
}

// HasOperationType returns a boolean if a field has been set.
func (o *CatalogV1alpha1AssetMappingStatusSourceResourceApiServiceInstance) HasOperationType() bool {
	if o != nil && !IsNil(o.OperationType) {
		return true
	}

	return false
}

// SetOperationType gets a reference to the given CatalogV1alpha1AssetMappingStatusOperationType and assigns it to the OperationType field.
func (o *CatalogV1alpha1AssetMappingStatusSourceResourceApiServiceInstance) SetOperationType(v CatalogV1alpha1AssetMappingStatusOperationType) {
	o.OperationType = &v
}

func (o CatalogV1alpha1AssetMappingStatusSourceResourceApiServiceInstance) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogV1alpha1AssetMappingStatusSourceResourceApiServiceInstance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ref) {
		toSerialize["ref"] = o.Ref
	}
	if !IsNil(o.OperationType) {
		toSerialize["operationType"] = o.OperationType
	}
	return toSerialize, nil
}

type NullableCatalogV1alpha1AssetMappingStatusSourceResourceApiServiceInstance struct {
	value *CatalogV1alpha1AssetMappingStatusSourceResourceApiServiceInstance
	isSet bool
}

func (v NullableCatalogV1alpha1AssetMappingStatusSourceResourceApiServiceInstance) Get() *CatalogV1alpha1AssetMappingStatusSourceResourceApiServiceInstance {
	return v.value
}

func (v *NullableCatalogV1alpha1AssetMappingStatusSourceResourceApiServiceInstance) Set(val *CatalogV1alpha1AssetMappingStatusSourceResourceApiServiceInstance) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogV1alpha1AssetMappingStatusSourceResourceApiServiceInstance) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogV1alpha1AssetMappingStatusSourceResourceApiServiceInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogV1alpha1AssetMappingStatusSourceResourceApiServiceInstance(val *CatalogV1alpha1AssetMappingStatusSourceResourceApiServiceInstance) *NullableCatalogV1alpha1AssetMappingStatusSourceResourceApiServiceInstance {
	return &NullableCatalogV1alpha1AssetMappingStatusSourceResourceApiServiceInstance{value: val, isSet: true}
}

func (v NullableCatalogV1alpha1AssetMappingStatusSourceResourceApiServiceInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogV1alpha1AssetMappingStatusSourceResourceApiServiceInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


