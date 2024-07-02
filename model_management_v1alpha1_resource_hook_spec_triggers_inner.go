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

// checks if the ManagementV1alpha1ResourceHookSpecTriggersInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManagementV1alpha1ResourceHookSpecTriggersInner{}

// ManagementV1alpha1ResourceHookSpecTriggersInner struct for ManagementV1alpha1ResourceHookSpecTriggersInner
type ManagementV1alpha1ResourceHookSpecTriggersInner struct {
	// Value for the group of the resource. Use \"*\" for any.
	Group string `json:"group"`
	Scope *ManagementV1alpha1ResourceHookSpecTriggersInnerScope `json:"scope,omitempty"`
	// Value for the Kind of the resource. Use \"*\" for any.
	Kind string `json:"kind"`
	// Name of the resource. Use \"*\" for any.
	Name string `json:"name" validate:"regexp=^([\\\\*])|([a-z0-9]([-a-z0-9]*[a-z0-9])?(\\\\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*)"`
	Type []string `json:"type"`
}

type _ManagementV1alpha1ResourceHookSpecTriggersInner ManagementV1alpha1ResourceHookSpecTriggersInner

// NewManagementV1alpha1ResourceHookSpecTriggersInner instantiates a new ManagementV1alpha1ResourceHookSpecTriggersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagementV1alpha1ResourceHookSpecTriggersInner(group string, kind string, name string, type_ []string) *ManagementV1alpha1ResourceHookSpecTriggersInner {
	this := ManagementV1alpha1ResourceHookSpecTriggersInner{}
	this.Group = group
	this.Kind = kind
	this.Name = name
	this.Type = type_
	return &this
}

// NewManagementV1alpha1ResourceHookSpecTriggersInnerWithDefaults instantiates a new ManagementV1alpha1ResourceHookSpecTriggersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagementV1alpha1ResourceHookSpecTriggersInnerWithDefaults() *ManagementV1alpha1ResourceHookSpecTriggersInner {
	this := ManagementV1alpha1ResourceHookSpecTriggersInner{}
	return &this
}

// GetGroup returns the Group field value
func (o *ManagementV1alpha1ResourceHookSpecTriggersInner) GetGroup() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Group
}

// GetGroupOk returns a tuple with the Group field value
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1ResourceHookSpecTriggersInner) GetGroupOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Group, true
}

// SetGroup sets field value
func (o *ManagementV1alpha1ResourceHookSpecTriggersInner) SetGroup(v string) {
	o.Group = v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *ManagementV1alpha1ResourceHookSpecTriggersInner) GetScope() ManagementV1alpha1ResourceHookSpecTriggersInnerScope {
	if o == nil || IsNil(o.Scope) {
		var ret ManagementV1alpha1ResourceHookSpecTriggersInnerScope
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1ResourceHookSpecTriggersInner) GetScopeOk() (*ManagementV1alpha1ResourceHookSpecTriggersInnerScope, bool) {
	if o == nil || IsNil(o.Scope) {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *ManagementV1alpha1ResourceHookSpecTriggersInner) HasScope() bool {
	if o != nil && !IsNil(o.Scope) {
		return true
	}

	return false
}

// SetScope gets a reference to the given ManagementV1alpha1ResourceHookSpecTriggersInnerScope and assigns it to the Scope field.
func (o *ManagementV1alpha1ResourceHookSpecTriggersInner) SetScope(v ManagementV1alpha1ResourceHookSpecTriggersInnerScope) {
	o.Scope = &v
}

// GetKind returns the Kind field value
func (o *ManagementV1alpha1ResourceHookSpecTriggersInner) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1ResourceHookSpecTriggersInner) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *ManagementV1alpha1ResourceHookSpecTriggersInner) SetKind(v string) {
	o.Kind = v
}

// GetName returns the Name field value
func (o *ManagementV1alpha1ResourceHookSpecTriggersInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1ResourceHookSpecTriggersInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ManagementV1alpha1ResourceHookSpecTriggersInner) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *ManagementV1alpha1ResourceHookSpecTriggersInner) GetType() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1ResourceHookSpecTriggersInner) GetTypeOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type, true
}

// SetType sets field value
func (o *ManagementV1alpha1ResourceHookSpecTriggersInner) SetType(v []string) {
	o.Type = v
}

func (o ManagementV1alpha1ResourceHookSpecTriggersInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManagementV1alpha1ResourceHookSpecTriggersInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["group"] = o.Group
	if !IsNil(o.Scope) {
		toSerialize["scope"] = o.Scope
	}
	toSerialize["kind"] = o.Kind
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *ManagementV1alpha1ResourceHookSpecTriggersInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"group",
		"kind",
		"name",
		"type",
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

	varManagementV1alpha1ResourceHookSpecTriggersInner := _ManagementV1alpha1ResourceHookSpecTriggersInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varManagementV1alpha1ResourceHookSpecTriggersInner)

	if err != nil {
		return err
	}

	*o = ManagementV1alpha1ResourceHookSpecTriggersInner(varManagementV1alpha1ResourceHookSpecTriggersInner)

	return err
}

type NullableManagementV1alpha1ResourceHookSpecTriggersInner struct {
	value *ManagementV1alpha1ResourceHookSpecTriggersInner
	isSet bool
}

func (v NullableManagementV1alpha1ResourceHookSpecTriggersInner) Get() *ManagementV1alpha1ResourceHookSpecTriggersInner {
	return v.value
}

func (v *NullableManagementV1alpha1ResourceHookSpecTriggersInner) Set(val *ManagementV1alpha1ResourceHookSpecTriggersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableManagementV1alpha1ResourceHookSpecTriggersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableManagementV1alpha1ResourceHookSpecTriggersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagementV1alpha1ResourceHookSpecTriggersInner(val *ManagementV1alpha1ResourceHookSpecTriggersInner) *NullableManagementV1alpha1ResourceHookSpecTriggersInner {
	return &NullableManagementV1alpha1ResourceHookSpecTriggersInner{value: val, isSet: true}
}

func (v NullableManagementV1alpha1ResourceHookSpecTriggersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagementV1alpha1ResourceHookSpecTriggersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


