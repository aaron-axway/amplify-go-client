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

// checks if the ManagementV1alpha1ManagedApplicationReferences type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManagementV1alpha1ManagedApplicationReferences{}

// ManagementV1alpha1ManagedApplicationReferences struct for ManagementV1alpha1ManagedApplicationReferences
type ManagementV1alpha1ManagedApplicationReferences struct {
	// Reference to Application resource
	Application *string `json:"application,omitempty"`
}

// NewManagementV1alpha1ManagedApplicationReferences instantiates a new ManagementV1alpha1ManagedApplicationReferences object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagementV1alpha1ManagedApplicationReferences() *ManagementV1alpha1ManagedApplicationReferences {
	this := ManagementV1alpha1ManagedApplicationReferences{}
	return &this
}

// NewManagementV1alpha1ManagedApplicationReferencesWithDefaults instantiates a new ManagementV1alpha1ManagedApplicationReferences object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagementV1alpha1ManagedApplicationReferencesWithDefaults() *ManagementV1alpha1ManagedApplicationReferences {
	this := ManagementV1alpha1ManagedApplicationReferences{}
	return &this
}

// GetApplication returns the Application field value if set, zero value otherwise.
func (o *ManagementV1alpha1ManagedApplicationReferences) GetApplication() string {
	if o == nil || IsNil(o.Application) {
		var ret string
		return ret
	}
	return *o.Application
}

// GetApplicationOk returns a tuple with the Application field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1ManagedApplicationReferences) GetApplicationOk() (*string, bool) {
	if o == nil || IsNil(o.Application) {
		return nil, false
	}
	return o.Application, true
}

// HasApplication returns a boolean if a field has been set.
func (o *ManagementV1alpha1ManagedApplicationReferences) HasApplication() bool {
	if o != nil && !IsNil(o.Application) {
		return true
	}

	return false
}

// SetApplication gets a reference to the given string and assigns it to the Application field.
func (o *ManagementV1alpha1ManagedApplicationReferences) SetApplication(v string) {
	o.Application = &v
}

func (o ManagementV1alpha1ManagedApplicationReferences) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManagementV1alpha1ManagedApplicationReferences) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Application) {
		toSerialize["application"] = o.Application
	}
	return toSerialize, nil
}

type NullableManagementV1alpha1ManagedApplicationReferences struct {
	value *ManagementV1alpha1ManagedApplicationReferences
	isSet bool
}

func (v NullableManagementV1alpha1ManagedApplicationReferences) Get() *ManagementV1alpha1ManagedApplicationReferences {
	return v.value
}

func (v *NullableManagementV1alpha1ManagedApplicationReferences) Set(val *ManagementV1alpha1ManagedApplicationReferences) {
	v.value = val
	v.isSet = true
}

func (v NullableManagementV1alpha1ManagedApplicationReferences) IsSet() bool {
	return v.isSet
}

func (v *NullableManagementV1alpha1ManagedApplicationReferences) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagementV1alpha1ManagedApplicationReferences(val *ManagementV1alpha1ManagedApplicationReferences) *NullableManagementV1alpha1ManagedApplicationReferences {
	return &NullableManagementV1alpha1ManagedApplicationReferences{value: val, isSet: true}
}

func (v NullableManagementV1alpha1ManagedApplicationReferences) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagementV1alpha1ManagedApplicationReferences) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


