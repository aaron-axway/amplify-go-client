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

// checks if the ManagementV1alpha1APISpecSpec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManagementV1alpha1APISpecSpec{}

// ManagementV1alpha1APISpecSpec struct for ManagementV1alpha1APISpecSpec
type ManagementV1alpha1APISpecSpec struct {
	SpecDiscoveryRef *string `json:"specDiscoveryRef,omitempty"`
	Definition *ManagementV1alpha1APISpecSpecDefinition `json:"definition,omitempty"`
	Endpoints []ManagementV1alpha1APISpecSpecEndpointsInner `json:"endpoints,omitempty"`
}

// NewManagementV1alpha1APISpecSpec instantiates a new ManagementV1alpha1APISpecSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagementV1alpha1APISpecSpec() *ManagementV1alpha1APISpecSpec {
	this := ManagementV1alpha1APISpecSpec{}
	return &this
}

// NewManagementV1alpha1APISpecSpecWithDefaults instantiates a new ManagementV1alpha1APISpecSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagementV1alpha1APISpecSpecWithDefaults() *ManagementV1alpha1APISpecSpec {
	this := ManagementV1alpha1APISpecSpec{}
	return &this
}

// GetSpecDiscoveryRef returns the SpecDiscoveryRef field value if set, zero value otherwise.
func (o *ManagementV1alpha1APISpecSpec) GetSpecDiscoveryRef() string {
	if o == nil || IsNil(o.SpecDiscoveryRef) {
		var ret string
		return ret
	}
	return *o.SpecDiscoveryRef
}

// GetSpecDiscoveryRefOk returns a tuple with the SpecDiscoveryRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1APISpecSpec) GetSpecDiscoveryRefOk() (*string, bool) {
	if o == nil || IsNil(o.SpecDiscoveryRef) {
		return nil, false
	}
	return o.SpecDiscoveryRef, true
}

// HasSpecDiscoveryRef returns a boolean if a field has been set.
func (o *ManagementV1alpha1APISpecSpec) HasSpecDiscoveryRef() bool {
	if o != nil && !IsNil(o.SpecDiscoveryRef) {
		return true
	}

	return false
}

// SetSpecDiscoveryRef gets a reference to the given string and assigns it to the SpecDiscoveryRef field.
func (o *ManagementV1alpha1APISpecSpec) SetSpecDiscoveryRef(v string) {
	o.SpecDiscoveryRef = &v
}

// GetDefinition returns the Definition field value if set, zero value otherwise.
func (o *ManagementV1alpha1APISpecSpec) GetDefinition() ManagementV1alpha1APISpecSpecDefinition {
	if o == nil || IsNil(o.Definition) {
		var ret ManagementV1alpha1APISpecSpecDefinition
		return ret
	}
	return *o.Definition
}

// GetDefinitionOk returns a tuple with the Definition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1APISpecSpec) GetDefinitionOk() (*ManagementV1alpha1APISpecSpecDefinition, bool) {
	if o == nil || IsNil(o.Definition) {
		return nil, false
	}
	return o.Definition, true
}

// HasDefinition returns a boolean if a field has been set.
func (o *ManagementV1alpha1APISpecSpec) HasDefinition() bool {
	if o != nil && !IsNil(o.Definition) {
		return true
	}

	return false
}

// SetDefinition gets a reference to the given ManagementV1alpha1APISpecSpecDefinition and assigns it to the Definition field.
func (o *ManagementV1alpha1APISpecSpec) SetDefinition(v ManagementV1alpha1APISpecSpecDefinition) {
	o.Definition = &v
}

// GetEndpoints returns the Endpoints field value if set, zero value otherwise.
func (o *ManagementV1alpha1APISpecSpec) GetEndpoints() []ManagementV1alpha1APISpecSpecEndpointsInner {
	if o == nil || IsNil(o.Endpoints) {
		var ret []ManagementV1alpha1APISpecSpecEndpointsInner
		return ret
	}
	return o.Endpoints
}

// GetEndpointsOk returns a tuple with the Endpoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1APISpecSpec) GetEndpointsOk() ([]ManagementV1alpha1APISpecSpecEndpointsInner, bool) {
	if o == nil || IsNil(o.Endpoints) {
		return nil, false
	}
	return o.Endpoints, true
}

// HasEndpoints returns a boolean if a field has been set.
func (o *ManagementV1alpha1APISpecSpec) HasEndpoints() bool {
	if o != nil && !IsNil(o.Endpoints) {
		return true
	}

	return false
}

// SetEndpoints gets a reference to the given []ManagementV1alpha1APISpecSpecEndpointsInner and assigns it to the Endpoints field.
func (o *ManagementV1alpha1APISpecSpec) SetEndpoints(v []ManagementV1alpha1APISpecSpecEndpointsInner) {
	o.Endpoints = v
}

func (o ManagementV1alpha1APISpecSpec) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManagementV1alpha1APISpecSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SpecDiscoveryRef) {
		toSerialize["specDiscoveryRef"] = o.SpecDiscoveryRef
	}
	if !IsNil(o.Definition) {
		toSerialize["definition"] = o.Definition
	}
	if !IsNil(o.Endpoints) {
		toSerialize["endpoints"] = o.Endpoints
	}
	return toSerialize, nil
}

type NullableManagementV1alpha1APISpecSpec struct {
	value *ManagementV1alpha1APISpecSpec
	isSet bool
}

func (v NullableManagementV1alpha1APISpecSpec) Get() *ManagementV1alpha1APISpecSpec {
	return v.value
}

func (v *NullableManagementV1alpha1APISpecSpec) Set(val *ManagementV1alpha1APISpecSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableManagementV1alpha1APISpecSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableManagementV1alpha1APISpecSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagementV1alpha1APISpecSpec(val *ManagementV1alpha1APISpecSpec) *NullableManagementV1alpha1APISpecSpec {
	return &NullableManagementV1alpha1APISpecSpec{value: val, isSet: true}
}

func (v NullableManagementV1alpha1APISpecSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagementV1alpha1APISpecSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


