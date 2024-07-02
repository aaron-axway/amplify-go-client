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

// checks if the ManagementV1alpha1APIServiceInstanceSpecEndpointInnerRouting type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManagementV1alpha1APIServiceInstanceSpecEndpointInnerRouting{}

// ManagementV1alpha1APIServiceInstanceSpecEndpointInnerRouting struct for ManagementV1alpha1APIServiceInstanceSpecEndpointInnerRouting
type ManagementV1alpha1APIServiceInstanceSpecEndpointInnerRouting struct {
	// The base path to the API. Example: '/api'.
	BasePath *string `json:"basePath,omitempty" validate:"regexp=^\\/"`
	// Specify any additional routing details needed
	Details map[string]interface{} `json:"details,omitempty"`
}

// NewManagementV1alpha1APIServiceInstanceSpecEndpointInnerRouting instantiates a new ManagementV1alpha1APIServiceInstanceSpecEndpointInnerRouting object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagementV1alpha1APIServiceInstanceSpecEndpointInnerRouting() *ManagementV1alpha1APIServiceInstanceSpecEndpointInnerRouting {
	this := ManagementV1alpha1APIServiceInstanceSpecEndpointInnerRouting{}
	return &this
}

// NewManagementV1alpha1APIServiceInstanceSpecEndpointInnerRoutingWithDefaults instantiates a new ManagementV1alpha1APIServiceInstanceSpecEndpointInnerRouting object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagementV1alpha1APIServiceInstanceSpecEndpointInnerRoutingWithDefaults() *ManagementV1alpha1APIServiceInstanceSpecEndpointInnerRouting {
	this := ManagementV1alpha1APIServiceInstanceSpecEndpointInnerRouting{}
	return &this
}

// GetBasePath returns the BasePath field value if set, zero value otherwise.
func (o *ManagementV1alpha1APIServiceInstanceSpecEndpointInnerRouting) GetBasePath() string {
	if o == nil || IsNil(o.BasePath) {
		var ret string
		return ret
	}
	return *o.BasePath
}

// GetBasePathOk returns a tuple with the BasePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1APIServiceInstanceSpecEndpointInnerRouting) GetBasePathOk() (*string, bool) {
	if o == nil || IsNil(o.BasePath) {
		return nil, false
	}
	return o.BasePath, true
}

// HasBasePath returns a boolean if a field has been set.
func (o *ManagementV1alpha1APIServiceInstanceSpecEndpointInnerRouting) HasBasePath() bool {
	if o != nil && !IsNil(o.BasePath) {
		return true
	}

	return false
}

// SetBasePath gets a reference to the given string and assigns it to the BasePath field.
func (o *ManagementV1alpha1APIServiceInstanceSpecEndpointInnerRouting) SetBasePath(v string) {
	o.BasePath = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *ManagementV1alpha1APIServiceInstanceSpecEndpointInnerRouting) GetDetails() map[string]interface{} {
	if o == nil || IsNil(o.Details) {
		var ret map[string]interface{}
		return ret
	}
	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1APIServiceInstanceSpecEndpointInnerRouting) GetDetailsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Details) {
		return map[string]interface{}{}, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *ManagementV1alpha1APIServiceInstanceSpecEndpointInnerRouting) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given map[string]interface{} and assigns it to the Details field.
func (o *ManagementV1alpha1APIServiceInstanceSpecEndpointInnerRouting) SetDetails(v map[string]interface{}) {
	o.Details = v
}

func (o ManagementV1alpha1APIServiceInstanceSpecEndpointInnerRouting) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManagementV1alpha1APIServiceInstanceSpecEndpointInnerRouting) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BasePath) {
		toSerialize["basePath"] = o.BasePath
	}
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	return toSerialize, nil
}

type NullableManagementV1alpha1APIServiceInstanceSpecEndpointInnerRouting struct {
	value *ManagementV1alpha1APIServiceInstanceSpecEndpointInnerRouting
	isSet bool
}

func (v NullableManagementV1alpha1APIServiceInstanceSpecEndpointInnerRouting) Get() *ManagementV1alpha1APIServiceInstanceSpecEndpointInnerRouting {
	return v.value
}

func (v *NullableManagementV1alpha1APIServiceInstanceSpecEndpointInnerRouting) Set(val *ManagementV1alpha1APIServiceInstanceSpecEndpointInnerRouting) {
	v.value = val
	v.isSet = true
}

func (v NullableManagementV1alpha1APIServiceInstanceSpecEndpointInnerRouting) IsSet() bool {
	return v.isSet
}

func (v *NullableManagementV1alpha1APIServiceInstanceSpecEndpointInnerRouting) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagementV1alpha1APIServiceInstanceSpecEndpointInnerRouting(val *ManagementV1alpha1APIServiceInstanceSpecEndpointInnerRouting) *NullableManagementV1alpha1APIServiceInstanceSpecEndpointInnerRouting {
	return &NullableManagementV1alpha1APIServiceInstanceSpecEndpointInnerRouting{value: val, isSet: true}
}

func (v NullableManagementV1alpha1APIServiceInstanceSpecEndpointInnerRouting) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagementV1alpha1APIServiceInstanceSpecEndpointInnerRouting) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


