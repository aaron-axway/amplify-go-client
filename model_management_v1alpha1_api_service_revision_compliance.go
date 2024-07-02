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

// checks if the ManagementV1alpha1APIServiceRevisionCompliance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManagementV1alpha1APIServiceRevisionCompliance{}

// ManagementV1alpha1APIServiceRevisionCompliance struct for ManagementV1alpha1APIServiceRevisionCompliance
type ManagementV1alpha1APIServiceRevisionCompliance struct {
	Design *ManagementV1alpha1APIServiceRevisionComplianceLintingStatus `json:"design,omitempty"`
	Security *ManagementV1alpha1APIServiceRevisionComplianceLintingStatus `json:"security,omitempty"`
}

// NewManagementV1alpha1APIServiceRevisionCompliance instantiates a new ManagementV1alpha1APIServiceRevisionCompliance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagementV1alpha1APIServiceRevisionCompliance() *ManagementV1alpha1APIServiceRevisionCompliance {
	this := ManagementV1alpha1APIServiceRevisionCompliance{}
	return &this
}

// NewManagementV1alpha1APIServiceRevisionComplianceWithDefaults instantiates a new ManagementV1alpha1APIServiceRevisionCompliance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagementV1alpha1APIServiceRevisionComplianceWithDefaults() *ManagementV1alpha1APIServiceRevisionCompliance {
	this := ManagementV1alpha1APIServiceRevisionCompliance{}
	return &this
}

// GetDesign returns the Design field value if set, zero value otherwise.
func (o *ManagementV1alpha1APIServiceRevisionCompliance) GetDesign() ManagementV1alpha1APIServiceRevisionComplianceLintingStatus {
	if o == nil || IsNil(o.Design) {
		var ret ManagementV1alpha1APIServiceRevisionComplianceLintingStatus
		return ret
	}
	return *o.Design
}

// GetDesignOk returns a tuple with the Design field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1APIServiceRevisionCompliance) GetDesignOk() (*ManagementV1alpha1APIServiceRevisionComplianceLintingStatus, bool) {
	if o == nil || IsNil(o.Design) {
		return nil, false
	}
	return o.Design, true
}

// HasDesign returns a boolean if a field has been set.
func (o *ManagementV1alpha1APIServiceRevisionCompliance) HasDesign() bool {
	if o != nil && !IsNil(o.Design) {
		return true
	}

	return false
}

// SetDesign gets a reference to the given ManagementV1alpha1APIServiceRevisionComplianceLintingStatus and assigns it to the Design field.
func (o *ManagementV1alpha1APIServiceRevisionCompliance) SetDesign(v ManagementV1alpha1APIServiceRevisionComplianceLintingStatus) {
	o.Design = &v
}

// GetSecurity returns the Security field value if set, zero value otherwise.
func (o *ManagementV1alpha1APIServiceRevisionCompliance) GetSecurity() ManagementV1alpha1APIServiceRevisionComplianceLintingStatus {
	if o == nil || IsNil(o.Security) {
		var ret ManagementV1alpha1APIServiceRevisionComplianceLintingStatus
		return ret
	}
	return *o.Security
}

// GetSecurityOk returns a tuple with the Security field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1APIServiceRevisionCompliance) GetSecurityOk() (*ManagementV1alpha1APIServiceRevisionComplianceLintingStatus, bool) {
	if o == nil || IsNil(o.Security) {
		return nil, false
	}
	return o.Security, true
}

// HasSecurity returns a boolean if a field has been set.
func (o *ManagementV1alpha1APIServiceRevisionCompliance) HasSecurity() bool {
	if o != nil && !IsNil(o.Security) {
		return true
	}

	return false
}

// SetSecurity gets a reference to the given ManagementV1alpha1APIServiceRevisionComplianceLintingStatus and assigns it to the Security field.
func (o *ManagementV1alpha1APIServiceRevisionCompliance) SetSecurity(v ManagementV1alpha1APIServiceRevisionComplianceLintingStatus) {
	o.Security = &v
}

func (o ManagementV1alpha1APIServiceRevisionCompliance) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManagementV1alpha1APIServiceRevisionCompliance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Design) {
		toSerialize["design"] = o.Design
	}
	if !IsNil(o.Security) {
		toSerialize["security"] = o.Security
	}
	return toSerialize, nil
}

type NullableManagementV1alpha1APIServiceRevisionCompliance struct {
	value *ManagementV1alpha1APIServiceRevisionCompliance
	isSet bool
}

func (v NullableManagementV1alpha1APIServiceRevisionCompliance) Get() *ManagementV1alpha1APIServiceRevisionCompliance {
	return v.value
}

func (v *NullableManagementV1alpha1APIServiceRevisionCompliance) Set(val *ManagementV1alpha1APIServiceRevisionCompliance) {
	v.value = val
	v.isSet = true
}

func (v NullableManagementV1alpha1APIServiceRevisionCompliance) IsSet() bool {
	return v.isSet
}

func (v *NullableManagementV1alpha1APIServiceRevisionCompliance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagementV1alpha1APIServiceRevisionCompliance(val *ManagementV1alpha1APIServiceRevisionCompliance) *NullableManagementV1alpha1APIServiceRevisionCompliance {
	return &NullableManagementV1alpha1APIServiceRevisionCompliance{value: val, isSet: true}
}

func (v NullableManagementV1alpha1APIServiceRevisionCompliance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagementV1alpha1APIServiceRevisionCompliance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


