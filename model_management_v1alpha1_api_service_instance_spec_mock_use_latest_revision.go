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

// checks if the ManagementV1alpha1APIServiceInstanceSpecMockUseLatestRevision type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManagementV1alpha1APIServiceInstanceSpecMockUseLatestRevision{}

// ManagementV1alpha1APIServiceInstanceSpecMockUseLatestRevision Optional property used to automatically update the mock service with the newest APIServiceRevision. When not set, you must assign the \"spec.apiServiceRevision\" to revision you want to mock. 
type ManagementV1alpha1APIServiceInstanceSpecMockUseLatestRevision struct {
	ApiService string `json:"apiService"`
}

type _ManagementV1alpha1APIServiceInstanceSpecMockUseLatestRevision ManagementV1alpha1APIServiceInstanceSpecMockUseLatestRevision

// NewManagementV1alpha1APIServiceInstanceSpecMockUseLatestRevision instantiates a new ManagementV1alpha1APIServiceInstanceSpecMockUseLatestRevision object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagementV1alpha1APIServiceInstanceSpecMockUseLatestRevision(apiService string) *ManagementV1alpha1APIServiceInstanceSpecMockUseLatestRevision {
	this := ManagementV1alpha1APIServiceInstanceSpecMockUseLatestRevision{}
	this.ApiService = apiService
	return &this
}

// NewManagementV1alpha1APIServiceInstanceSpecMockUseLatestRevisionWithDefaults instantiates a new ManagementV1alpha1APIServiceInstanceSpecMockUseLatestRevision object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagementV1alpha1APIServiceInstanceSpecMockUseLatestRevisionWithDefaults() *ManagementV1alpha1APIServiceInstanceSpecMockUseLatestRevision {
	this := ManagementV1alpha1APIServiceInstanceSpecMockUseLatestRevision{}
	return &this
}

// GetApiService returns the ApiService field value
func (o *ManagementV1alpha1APIServiceInstanceSpecMockUseLatestRevision) GetApiService() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiService
}

// GetApiServiceOk returns a tuple with the ApiService field value
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1APIServiceInstanceSpecMockUseLatestRevision) GetApiServiceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiService, true
}

// SetApiService sets field value
func (o *ManagementV1alpha1APIServiceInstanceSpecMockUseLatestRevision) SetApiService(v string) {
	o.ApiService = v
}

func (o ManagementV1alpha1APIServiceInstanceSpecMockUseLatestRevision) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManagementV1alpha1APIServiceInstanceSpecMockUseLatestRevision) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["apiService"] = o.ApiService
	return toSerialize, nil
}

func (o *ManagementV1alpha1APIServiceInstanceSpecMockUseLatestRevision) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"apiService",
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

	varManagementV1alpha1APIServiceInstanceSpecMockUseLatestRevision := _ManagementV1alpha1APIServiceInstanceSpecMockUseLatestRevision{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varManagementV1alpha1APIServiceInstanceSpecMockUseLatestRevision)

	if err != nil {
		return err
	}

	*o = ManagementV1alpha1APIServiceInstanceSpecMockUseLatestRevision(varManagementV1alpha1APIServiceInstanceSpecMockUseLatestRevision)

	return err
}

type NullableManagementV1alpha1APIServiceInstanceSpecMockUseLatestRevision struct {
	value *ManagementV1alpha1APIServiceInstanceSpecMockUseLatestRevision
	isSet bool
}

func (v NullableManagementV1alpha1APIServiceInstanceSpecMockUseLatestRevision) Get() *ManagementV1alpha1APIServiceInstanceSpecMockUseLatestRevision {
	return v.value
}

func (v *NullableManagementV1alpha1APIServiceInstanceSpecMockUseLatestRevision) Set(val *ManagementV1alpha1APIServiceInstanceSpecMockUseLatestRevision) {
	v.value = val
	v.isSet = true
}

func (v NullableManagementV1alpha1APIServiceInstanceSpecMockUseLatestRevision) IsSet() bool {
	return v.isSet
}

func (v *NullableManagementV1alpha1APIServiceInstanceSpecMockUseLatestRevision) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagementV1alpha1APIServiceInstanceSpecMockUseLatestRevision(val *ManagementV1alpha1APIServiceInstanceSpecMockUseLatestRevision) *NullableManagementV1alpha1APIServiceInstanceSpecMockUseLatestRevision {
	return &NullableManagementV1alpha1APIServiceInstanceSpecMockUseLatestRevision{value: val, isSet: true}
}

func (v NullableManagementV1alpha1APIServiceInstanceSpecMockUseLatestRevision) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagementV1alpha1APIServiceInstanceSpecMockUseLatestRevision) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

