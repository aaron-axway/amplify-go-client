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

// checks if the ManagementV1alpha1AssetMappingStatusOutputsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManagementV1alpha1AssetMappingStatusOutputsInner{}

// ManagementV1alpha1AssetMappingStatusOutputsInner struct for ManagementV1alpha1AssetMappingStatusOutputsInner
type ManagementV1alpha1AssetMappingStatusOutputsInner struct {
	Resource *ManagementV1alpha1AssetMappingStatusOutputsInnerResource `json:"resource,omitempty"`
}

// NewManagementV1alpha1AssetMappingStatusOutputsInner instantiates a new ManagementV1alpha1AssetMappingStatusOutputsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagementV1alpha1AssetMappingStatusOutputsInner() *ManagementV1alpha1AssetMappingStatusOutputsInner {
	this := ManagementV1alpha1AssetMappingStatusOutputsInner{}
	return &this
}

// NewManagementV1alpha1AssetMappingStatusOutputsInnerWithDefaults instantiates a new ManagementV1alpha1AssetMappingStatusOutputsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagementV1alpha1AssetMappingStatusOutputsInnerWithDefaults() *ManagementV1alpha1AssetMappingStatusOutputsInner {
	this := ManagementV1alpha1AssetMappingStatusOutputsInner{}
	return &this
}

// GetResource returns the Resource field value if set, zero value otherwise.
func (o *ManagementV1alpha1AssetMappingStatusOutputsInner) GetResource() ManagementV1alpha1AssetMappingStatusOutputsInnerResource {
	if o == nil || IsNil(o.Resource) {
		var ret ManagementV1alpha1AssetMappingStatusOutputsInnerResource
		return ret
	}
	return *o.Resource
}

// GetResourceOk returns a tuple with the Resource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1AssetMappingStatusOutputsInner) GetResourceOk() (*ManagementV1alpha1AssetMappingStatusOutputsInnerResource, bool) {
	if o == nil || IsNil(o.Resource) {
		return nil, false
	}
	return o.Resource, true
}

// HasResource returns a boolean if a field has been set.
func (o *ManagementV1alpha1AssetMappingStatusOutputsInner) HasResource() bool {
	if o != nil && !IsNil(o.Resource) {
		return true
	}

	return false
}

// SetResource gets a reference to the given ManagementV1alpha1AssetMappingStatusOutputsInnerResource and assigns it to the Resource field.
func (o *ManagementV1alpha1AssetMappingStatusOutputsInner) SetResource(v ManagementV1alpha1AssetMappingStatusOutputsInnerResource) {
	o.Resource = &v
}

func (o ManagementV1alpha1AssetMappingStatusOutputsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManagementV1alpha1AssetMappingStatusOutputsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Resource) {
		toSerialize["resource"] = o.Resource
	}
	return toSerialize, nil
}

type NullableManagementV1alpha1AssetMappingStatusOutputsInner struct {
	value *ManagementV1alpha1AssetMappingStatusOutputsInner
	isSet bool
}

func (v NullableManagementV1alpha1AssetMappingStatusOutputsInner) Get() *ManagementV1alpha1AssetMappingStatusOutputsInner {
	return v.value
}

func (v *NullableManagementV1alpha1AssetMappingStatusOutputsInner) Set(val *ManagementV1alpha1AssetMappingStatusOutputsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableManagementV1alpha1AssetMappingStatusOutputsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableManagementV1alpha1AssetMappingStatusOutputsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagementV1alpha1AssetMappingStatusOutputsInner(val *ManagementV1alpha1AssetMappingStatusOutputsInner) *NullableManagementV1alpha1AssetMappingStatusOutputsInner {
	return &NullableManagementV1alpha1AssetMappingStatusOutputsInner{value: val, isSet: true}
}

func (v NullableManagementV1alpha1AssetMappingStatusOutputsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagementV1alpha1AssetMappingStatusOutputsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


