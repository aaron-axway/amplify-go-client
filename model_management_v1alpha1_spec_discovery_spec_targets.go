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

// checks if the ManagementV1alpha1SpecDiscoverySpecTargets type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManagementV1alpha1SpecDiscoverySpecTargets{}

// ManagementV1alpha1SpecDiscoverySpecTargets struct for ManagementV1alpha1SpecDiscoverySpecTargets
type ManagementV1alpha1SpecDiscoverySpecTargets struct {
	ExactPaths []ManagementV1alpha1SpecDiscoverySpecTargetsExactPathsInner `json:"exactPaths,omitempty"`
	FromAnnotation []ManagementV1alpha1SpecDiscoverySpecTargetsFromAnnotationInner `json:"fromAnnotation,omitempty"`
}

// NewManagementV1alpha1SpecDiscoverySpecTargets instantiates a new ManagementV1alpha1SpecDiscoverySpecTargets object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagementV1alpha1SpecDiscoverySpecTargets() *ManagementV1alpha1SpecDiscoverySpecTargets {
	this := ManagementV1alpha1SpecDiscoverySpecTargets{}
	return &this
}

// NewManagementV1alpha1SpecDiscoverySpecTargetsWithDefaults instantiates a new ManagementV1alpha1SpecDiscoverySpecTargets object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagementV1alpha1SpecDiscoverySpecTargetsWithDefaults() *ManagementV1alpha1SpecDiscoverySpecTargets {
	this := ManagementV1alpha1SpecDiscoverySpecTargets{}
	return &this
}

// GetExactPaths returns the ExactPaths field value if set, zero value otherwise.
func (o *ManagementV1alpha1SpecDiscoverySpecTargets) GetExactPaths() []ManagementV1alpha1SpecDiscoverySpecTargetsExactPathsInner {
	if o == nil || IsNil(o.ExactPaths) {
		var ret []ManagementV1alpha1SpecDiscoverySpecTargetsExactPathsInner
		return ret
	}
	return o.ExactPaths
}

// GetExactPathsOk returns a tuple with the ExactPaths field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1SpecDiscoverySpecTargets) GetExactPathsOk() ([]ManagementV1alpha1SpecDiscoverySpecTargetsExactPathsInner, bool) {
	if o == nil || IsNil(o.ExactPaths) {
		return nil, false
	}
	return o.ExactPaths, true
}

// HasExactPaths returns a boolean if a field has been set.
func (o *ManagementV1alpha1SpecDiscoverySpecTargets) HasExactPaths() bool {
	if o != nil && !IsNil(o.ExactPaths) {
		return true
	}

	return false
}

// SetExactPaths gets a reference to the given []ManagementV1alpha1SpecDiscoverySpecTargetsExactPathsInner and assigns it to the ExactPaths field.
func (o *ManagementV1alpha1SpecDiscoverySpecTargets) SetExactPaths(v []ManagementV1alpha1SpecDiscoverySpecTargetsExactPathsInner) {
	o.ExactPaths = v
}

// GetFromAnnotation returns the FromAnnotation field value if set, zero value otherwise.
func (o *ManagementV1alpha1SpecDiscoverySpecTargets) GetFromAnnotation() []ManagementV1alpha1SpecDiscoverySpecTargetsFromAnnotationInner {
	if o == nil || IsNil(o.FromAnnotation) {
		var ret []ManagementV1alpha1SpecDiscoverySpecTargetsFromAnnotationInner
		return ret
	}
	return o.FromAnnotation
}

// GetFromAnnotationOk returns a tuple with the FromAnnotation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1SpecDiscoverySpecTargets) GetFromAnnotationOk() ([]ManagementV1alpha1SpecDiscoverySpecTargetsFromAnnotationInner, bool) {
	if o == nil || IsNil(o.FromAnnotation) {
		return nil, false
	}
	return o.FromAnnotation, true
}

// HasFromAnnotation returns a boolean if a field has been set.
func (o *ManagementV1alpha1SpecDiscoverySpecTargets) HasFromAnnotation() bool {
	if o != nil && !IsNil(o.FromAnnotation) {
		return true
	}

	return false
}

// SetFromAnnotation gets a reference to the given []ManagementV1alpha1SpecDiscoverySpecTargetsFromAnnotationInner and assigns it to the FromAnnotation field.
func (o *ManagementV1alpha1SpecDiscoverySpecTargets) SetFromAnnotation(v []ManagementV1alpha1SpecDiscoverySpecTargetsFromAnnotationInner) {
	o.FromAnnotation = v
}

func (o ManagementV1alpha1SpecDiscoverySpecTargets) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManagementV1alpha1SpecDiscoverySpecTargets) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExactPaths) {
		toSerialize["exactPaths"] = o.ExactPaths
	}
	if !IsNil(o.FromAnnotation) {
		toSerialize["fromAnnotation"] = o.FromAnnotation
	}
	return toSerialize, nil
}

type NullableManagementV1alpha1SpecDiscoverySpecTargets struct {
	value *ManagementV1alpha1SpecDiscoverySpecTargets
	isSet bool
}

func (v NullableManagementV1alpha1SpecDiscoverySpecTargets) Get() *ManagementV1alpha1SpecDiscoverySpecTargets {
	return v.value
}

func (v *NullableManagementV1alpha1SpecDiscoverySpecTargets) Set(val *ManagementV1alpha1SpecDiscoverySpecTargets) {
	v.value = val
	v.isSet = true
}

func (v NullableManagementV1alpha1SpecDiscoverySpecTargets) IsSet() bool {
	return v.isSet
}

func (v *NullableManagementV1alpha1SpecDiscoverySpecTargets) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagementV1alpha1SpecDiscoverySpecTargets(val *ManagementV1alpha1SpecDiscoverySpecTargets) *NullableManagementV1alpha1SpecDiscoverySpecTargets {
	return &NullableManagementV1alpha1SpecDiscoverySpecTargets{value: val, isSet: true}
}

func (v NullableManagementV1alpha1SpecDiscoverySpecTargets) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagementV1alpha1SpecDiscoverySpecTargets) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


