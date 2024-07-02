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

// checks if the ManagementV1alpha1AssetMappingStatusOutputsInnerResourceAssetResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManagementV1alpha1AssetMappingStatusOutputsInnerResourceAssetResource{}

// ManagementV1alpha1AssetMappingStatusOutputsInnerResourceAssetResource struct for ManagementV1alpha1AssetMappingStatusOutputsInnerResourceAssetResource
type ManagementV1alpha1AssetMappingStatusOutputsInnerResourceAssetResource struct {
	Ref *string `json:"ref,omitempty"`
	OperationType *ManagementV1alpha1AssetMappingStatusOperationType `json:"operationType,omitempty"`
}

// NewManagementV1alpha1AssetMappingStatusOutputsInnerResourceAssetResource instantiates a new ManagementV1alpha1AssetMappingStatusOutputsInnerResourceAssetResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagementV1alpha1AssetMappingStatusOutputsInnerResourceAssetResource() *ManagementV1alpha1AssetMappingStatusOutputsInnerResourceAssetResource {
	this := ManagementV1alpha1AssetMappingStatusOutputsInnerResourceAssetResource{}
	return &this
}

// NewManagementV1alpha1AssetMappingStatusOutputsInnerResourceAssetResourceWithDefaults instantiates a new ManagementV1alpha1AssetMappingStatusOutputsInnerResourceAssetResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagementV1alpha1AssetMappingStatusOutputsInnerResourceAssetResourceWithDefaults() *ManagementV1alpha1AssetMappingStatusOutputsInnerResourceAssetResource {
	this := ManagementV1alpha1AssetMappingStatusOutputsInnerResourceAssetResource{}
	return &this
}

// GetRef returns the Ref field value if set, zero value otherwise.
func (o *ManagementV1alpha1AssetMappingStatusOutputsInnerResourceAssetResource) GetRef() string {
	if o == nil || IsNil(o.Ref) {
		var ret string
		return ret
	}
	return *o.Ref
}

// GetRefOk returns a tuple with the Ref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1AssetMappingStatusOutputsInnerResourceAssetResource) GetRefOk() (*string, bool) {
	if o == nil || IsNil(o.Ref) {
		return nil, false
	}
	return o.Ref, true
}

// HasRef returns a boolean if a field has been set.
func (o *ManagementV1alpha1AssetMappingStatusOutputsInnerResourceAssetResource) HasRef() bool {
	if o != nil && !IsNil(o.Ref) {
		return true
	}

	return false
}

// SetRef gets a reference to the given string and assigns it to the Ref field.
func (o *ManagementV1alpha1AssetMappingStatusOutputsInnerResourceAssetResource) SetRef(v string) {
	o.Ref = &v
}

// GetOperationType returns the OperationType field value if set, zero value otherwise.
func (o *ManagementV1alpha1AssetMappingStatusOutputsInnerResourceAssetResource) GetOperationType() ManagementV1alpha1AssetMappingStatusOperationType {
	if o == nil || IsNil(o.OperationType) {
		var ret ManagementV1alpha1AssetMappingStatusOperationType
		return ret
	}
	return *o.OperationType
}

// GetOperationTypeOk returns a tuple with the OperationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1AssetMappingStatusOutputsInnerResourceAssetResource) GetOperationTypeOk() (*ManagementV1alpha1AssetMappingStatusOperationType, bool) {
	if o == nil || IsNil(o.OperationType) {
		return nil, false
	}
	return o.OperationType, true
}

// HasOperationType returns a boolean if a field has been set.
func (o *ManagementV1alpha1AssetMappingStatusOutputsInnerResourceAssetResource) HasOperationType() bool {
	if o != nil && !IsNil(o.OperationType) {
		return true
	}

	return false
}

// SetOperationType gets a reference to the given ManagementV1alpha1AssetMappingStatusOperationType and assigns it to the OperationType field.
func (o *ManagementV1alpha1AssetMappingStatusOutputsInnerResourceAssetResource) SetOperationType(v ManagementV1alpha1AssetMappingStatusOperationType) {
	o.OperationType = &v
}

func (o ManagementV1alpha1AssetMappingStatusOutputsInnerResourceAssetResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManagementV1alpha1AssetMappingStatusOutputsInnerResourceAssetResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ref) {
		toSerialize["ref"] = o.Ref
	}
	if !IsNil(o.OperationType) {
		toSerialize["operationType"] = o.OperationType
	}
	return toSerialize, nil
}

type NullableManagementV1alpha1AssetMappingStatusOutputsInnerResourceAssetResource struct {
	value *ManagementV1alpha1AssetMappingStatusOutputsInnerResourceAssetResource
	isSet bool
}

func (v NullableManagementV1alpha1AssetMappingStatusOutputsInnerResourceAssetResource) Get() *ManagementV1alpha1AssetMappingStatusOutputsInnerResourceAssetResource {
	return v.value
}

func (v *NullableManagementV1alpha1AssetMappingStatusOutputsInnerResourceAssetResource) Set(val *ManagementV1alpha1AssetMappingStatusOutputsInnerResourceAssetResource) {
	v.value = val
	v.isSet = true
}

func (v NullableManagementV1alpha1AssetMappingStatusOutputsInnerResourceAssetResource) IsSet() bool {
	return v.isSet
}

func (v *NullableManagementV1alpha1AssetMappingStatusOutputsInnerResourceAssetResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagementV1alpha1AssetMappingStatusOutputsInnerResourceAssetResource(val *ManagementV1alpha1AssetMappingStatusOutputsInnerResourceAssetResource) *NullableManagementV1alpha1AssetMappingStatusOutputsInnerResourceAssetResource {
	return &NullableManagementV1alpha1AssetMappingStatusOutputsInnerResourceAssetResource{value: val, isSet: true}
}

func (v NullableManagementV1alpha1AssetMappingStatusOutputsInnerResourceAssetResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagementV1alpha1AssetMappingStatusOutputsInnerResourceAssetResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

