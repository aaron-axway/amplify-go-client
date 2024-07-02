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

// checks if the CatalogV1alpha1AssetMappingStatusOutputsInnerResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogV1alpha1AssetMappingStatusOutputsInnerResource{}

// CatalogV1alpha1AssetMappingStatusOutputsInnerResource The resources that were impacted with the trigger of asset mapping.
type CatalogV1alpha1AssetMappingStatusOutputsInnerResource struct {
	AssetResource *CatalogV1alpha1AssetMappingStatusOutputsInnerResourceAssetResource `json:"assetResource,omitempty"`
	Stage *CatalogV1alpha1AssetMappingStatusOutputsInnerResourceStage `json:"stage,omitempty"`
}

// NewCatalogV1alpha1AssetMappingStatusOutputsInnerResource instantiates a new CatalogV1alpha1AssetMappingStatusOutputsInnerResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogV1alpha1AssetMappingStatusOutputsInnerResource() *CatalogV1alpha1AssetMappingStatusOutputsInnerResource {
	this := CatalogV1alpha1AssetMappingStatusOutputsInnerResource{}
	return &this
}

// NewCatalogV1alpha1AssetMappingStatusOutputsInnerResourceWithDefaults instantiates a new CatalogV1alpha1AssetMappingStatusOutputsInnerResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogV1alpha1AssetMappingStatusOutputsInnerResourceWithDefaults() *CatalogV1alpha1AssetMappingStatusOutputsInnerResource {
	this := CatalogV1alpha1AssetMappingStatusOutputsInnerResource{}
	return &this
}

// GetAssetResource returns the AssetResource field value if set, zero value otherwise.
func (o *CatalogV1alpha1AssetMappingStatusOutputsInnerResource) GetAssetResource() CatalogV1alpha1AssetMappingStatusOutputsInnerResourceAssetResource {
	if o == nil || IsNil(o.AssetResource) {
		var ret CatalogV1alpha1AssetMappingStatusOutputsInnerResourceAssetResource
		return ret
	}
	return *o.AssetResource
}

// GetAssetResourceOk returns a tuple with the AssetResource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1AssetMappingStatusOutputsInnerResource) GetAssetResourceOk() (*CatalogV1alpha1AssetMappingStatusOutputsInnerResourceAssetResource, bool) {
	if o == nil || IsNil(o.AssetResource) {
		return nil, false
	}
	return o.AssetResource, true
}

// HasAssetResource returns a boolean if a field has been set.
func (o *CatalogV1alpha1AssetMappingStatusOutputsInnerResource) HasAssetResource() bool {
	if o != nil && !IsNil(o.AssetResource) {
		return true
	}

	return false
}

// SetAssetResource gets a reference to the given CatalogV1alpha1AssetMappingStatusOutputsInnerResourceAssetResource and assigns it to the AssetResource field.
func (o *CatalogV1alpha1AssetMappingStatusOutputsInnerResource) SetAssetResource(v CatalogV1alpha1AssetMappingStatusOutputsInnerResourceAssetResource) {
	o.AssetResource = &v
}

// GetStage returns the Stage field value if set, zero value otherwise.
func (o *CatalogV1alpha1AssetMappingStatusOutputsInnerResource) GetStage() CatalogV1alpha1AssetMappingStatusOutputsInnerResourceStage {
	if o == nil || IsNil(o.Stage) {
		var ret CatalogV1alpha1AssetMappingStatusOutputsInnerResourceStage
		return ret
	}
	return *o.Stage
}

// GetStageOk returns a tuple with the Stage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1AssetMappingStatusOutputsInnerResource) GetStageOk() (*CatalogV1alpha1AssetMappingStatusOutputsInnerResourceStage, bool) {
	if o == nil || IsNil(o.Stage) {
		return nil, false
	}
	return o.Stage, true
}

// HasStage returns a boolean if a field has been set.
func (o *CatalogV1alpha1AssetMappingStatusOutputsInnerResource) HasStage() bool {
	if o != nil && !IsNil(o.Stage) {
		return true
	}

	return false
}

// SetStage gets a reference to the given CatalogV1alpha1AssetMappingStatusOutputsInnerResourceStage and assigns it to the Stage field.
func (o *CatalogV1alpha1AssetMappingStatusOutputsInnerResource) SetStage(v CatalogV1alpha1AssetMappingStatusOutputsInnerResourceStage) {
	o.Stage = &v
}

func (o CatalogV1alpha1AssetMappingStatusOutputsInnerResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogV1alpha1AssetMappingStatusOutputsInnerResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AssetResource) {
		toSerialize["assetResource"] = o.AssetResource
	}
	if !IsNil(o.Stage) {
		toSerialize["stage"] = o.Stage
	}
	return toSerialize, nil
}

type NullableCatalogV1alpha1AssetMappingStatusOutputsInnerResource struct {
	value *CatalogV1alpha1AssetMappingStatusOutputsInnerResource
	isSet bool
}

func (v NullableCatalogV1alpha1AssetMappingStatusOutputsInnerResource) Get() *CatalogV1alpha1AssetMappingStatusOutputsInnerResource {
	return v.value
}

func (v *NullableCatalogV1alpha1AssetMappingStatusOutputsInnerResource) Set(val *CatalogV1alpha1AssetMappingStatusOutputsInnerResource) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogV1alpha1AssetMappingStatusOutputsInnerResource) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogV1alpha1AssetMappingStatusOutputsInnerResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogV1alpha1AssetMappingStatusOutputsInnerResource(val *CatalogV1alpha1AssetMappingStatusOutputsInnerResource) *NullableCatalogV1alpha1AssetMappingStatusOutputsInnerResource {
	return &NullableCatalogV1alpha1AssetMappingStatusOutputsInnerResource{value: val, isSet: true}
}

func (v NullableCatalogV1alpha1AssetMappingStatusOutputsInnerResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogV1alpha1AssetMappingStatusOutputsInnerResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


