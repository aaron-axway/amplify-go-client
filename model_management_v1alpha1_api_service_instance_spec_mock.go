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
	"bytes"
	"fmt"
)

// checks if the ManagementV1alpha1APIServiceInstanceSpecMock type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManagementV1alpha1APIServiceInstanceSpecMock{}

// ManagementV1alpha1APIServiceInstanceSpecMock Sets up the referenced API to be mocked by an Axway server. Can only be set upon creation. Requires an \"API Mocking\" entitlement. 
type ManagementV1alpha1APIServiceInstanceSpecMock struct {
	// Assigned to the mock server's URL base path. Must be unique for the organization.
	Name string `json:"name" validate:"regexp=^([\\\\*])|([a-z0-9]([-a-z0-9]*[a-z0-9])?(\\\\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*)"`
	UseLatestRevision *ManagementV1alpha1APIServiceInstanceSpecMockUseLatestRevision `json:"useLatestRevision,omitempty"`
}

type _ManagementV1alpha1APIServiceInstanceSpecMock ManagementV1alpha1APIServiceInstanceSpecMock

// NewManagementV1alpha1APIServiceInstanceSpecMock instantiates a new ManagementV1alpha1APIServiceInstanceSpecMock object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagementV1alpha1APIServiceInstanceSpecMock(name string) *ManagementV1alpha1APIServiceInstanceSpecMock {
	this := ManagementV1alpha1APIServiceInstanceSpecMock{}
	this.Name = name
	return &this
}

// NewManagementV1alpha1APIServiceInstanceSpecMockWithDefaults instantiates a new ManagementV1alpha1APIServiceInstanceSpecMock object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagementV1alpha1APIServiceInstanceSpecMockWithDefaults() *ManagementV1alpha1APIServiceInstanceSpecMock {
	this := ManagementV1alpha1APIServiceInstanceSpecMock{}
	return &this
}

// GetName returns the Name field value
func (o *ManagementV1alpha1APIServiceInstanceSpecMock) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1APIServiceInstanceSpecMock) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ManagementV1alpha1APIServiceInstanceSpecMock) SetName(v string) {
	o.Name = v
}

// GetUseLatestRevision returns the UseLatestRevision field value if set, zero value otherwise.
func (o *ManagementV1alpha1APIServiceInstanceSpecMock) GetUseLatestRevision() ManagementV1alpha1APIServiceInstanceSpecMockUseLatestRevision {
	if o == nil || IsNil(o.UseLatestRevision) {
		var ret ManagementV1alpha1APIServiceInstanceSpecMockUseLatestRevision
		return ret
	}
	return *o.UseLatestRevision
}

// GetUseLatestRevisionOk returns a tuple with the UseLatestRevision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1APIServiceInstanceSpecMock) GetUseLatestRevisionOk() (*ManagementV1alpha1APIServiceInstanceSpecMockUseLatestRevision, bool) {
	if o == nil || IsNil(o.UseLatestRevision) {
		return nil, false
	}
	return o.UseLatestRevision, true
}

// HasUseLatestRevision returns a boolean if a field has been set.
func (o *ManagementV1alpha1APIServiceInstanceSpecMock) HasUseLatestRevision() bool {
	if o != nil && !IsNil(o.UseLatestRevision) {
		return true
	}

	return false
}

// SetUseLatestRevision gets a reference to the given ManagementV1alpha1APIServiceInstanceSpecMockUseLatestRevision and assigns it to the UseLatestRevision field.
func (o *ManagementV1alpha1APIServiceInstanceSpecMock) SetUseLatestRevision(v ManagementV1alpha1APIServiceInstanceSpecMockUseLatestRevision) {
	o.UseLatestRevision = &v
}

func (o ManagementV1alpha1APIServiceInstanceSpecMock) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManagementV1alpha1APIServiceInstanceSpecMock) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.UseLatestRevision) {
		toSerialize["useLatestRevision"] = o.UseLatestRevision
	}
	return toSerialize, nil
}

func (o *ManagementV1alpha1APIServiceInstanceSpecMock) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varManagementV1alpha1APIServiceInstanceSpecMock := _ManagementV1alpha1APIServiceInstanceSpecMock{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varManagementV1alpha1APIServiceInstanceSpecMock)

	if err != nil {
		return err
	}

	*o = ManagementV1alpha1APIServiceInstanceSpecMock(varManagementV1alpha1APIServiceInstanceSpecMock)

	return err
}

type NullableManagementV1alpha1APIServiceInstanceSpecMock struct {
	value *ManagementV1alpha1APIServiceInstanceSpecMock
	isSet bool
}

func (v NullableManagementV1alpha1APIServiceInstanceSpecMock) Get() *ManagementV1alpha1APIServiceInstanceSpecMock {
	return v.value
}

func (v *NullableManagementV1alpha1APIServiceInstanceSpecMock) Set(val *ManagementV1alpha1APIServiceInstanceSpecMock) {
	v.value = val
	v.isSet = true
}

func (v NullableManagementV1alpha1APIServiceInstanceSpecMock) IsSet() bool {
	return v.isSet
}

func (v *NullableManagementV1alpha1APIServiceInstanceSpecMock) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagementV1alpha1APIServiceInstanceSpecMock(val *ManagementV1alpha1APIServiceInstanceSpecMock) *NullableManagementV1alpha1APIServiceInstanceSpecMock {
	return &NullableManagementV1alpha1APIServiceInstanceSpecMock{value: val, isSet: true}
}

func (v NullableManagementV1alpha1APIServiceInstanceSpecMock) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagementV1alpha1APIServiceInstanceSpecMock) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

