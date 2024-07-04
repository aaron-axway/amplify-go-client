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

// checks if the DefinitionsV1alpha1AccessControlListSpec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DefinitionsV1alpha1AccessControlListSpec{}

// DefinitionsV1alpha1AccessControlListSpec struct for DefinitionsV1alpha1AccessControlListSpec
type DefinitionsV1alpha1AccessControlListSpec struct {
	Rules []DefinitionsV1alpha1AccessControlListSpecRulesInner `json:"rules,omitempty"`
	Subjects []CatalogV1alpha1AccessControlListSpecSubjectsInner `json:"subjects,omitempty"`
}

// NewDefinitionsV1alpha1AccessControlListSpec instantiates a new DefinitionsV1alpha1AccessControlListSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDefinitionsV1alpha1AccessControlListSpec() *DefinitionsV1alpha1AccessControlListSpec {
	this := DefinitionsV1alpha1AccessControlListSpec{}
	return &this
}

// NewDefinitionsV1alpha1AccessControlListSpecWithDefaults instantiates a new DefinitionsV1alpha1AccessControlListSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDefinitionsV1alpha1AccessControlListSpecWithDefaults() *DefinitionsV1alpha1AccessControlListSpec {
	this := DefinitionsV1alpha1AccessControlListSpec{}
	return &this
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *DefinitionsV1alpha1AccessControlListSpec) GetRules() []DefinitionsV1alpha1AccessControlListSpecRulesInner {
	if o == nil || IsNil(o.Rules) {
		var ret []DefinitionsV1alpha1AccessControlListSpecRulesInner
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefinitionsV1alpha1AccessControlListSpec) GetRulesOk() ([]DefinitionsV1alpha1AccessControlListSpecRulesInner, bool) {
	if o == nil || IsNil(o.Rules) {
		return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *DefinitionsV1alpha1AccessControlListSpec) HasRules() bool {
	if o != nil && !IsNil(o.Rules) {
		return true
	}

	return false
}

// SetRules gets a reference to the given []DefinitionsV1alpha1AccessControlListSpecRulesInner and assigns it to the Rules field.
func (o *DefinitionsV1alpha1AccessControlListSpec) SetRules(v []DefinitionsV1alpha1AccessControlListSpecRulesInner) {
	o.Rules = v
}

// GetSubjects returns the Subjects field value if set, zero value otherwise.
func (o *DefinitionsV1alpha1AccessControlListSpec) GetSubjects() []CatalogV1alpha1AccessControlListSpecSubjectsInner {
	if o == nil || IsNil(o.Subjects) {
		var ret []CatalogV1alpha1AccessControlListSpecSubjectsInner
		return ret
	}
	return o.Subjects
}

// GetSubjectsOk returns a tuple with the Subjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefinitionsV1alpha1AccessControlListSpec) GetSubjectsOk() ([]CatalogV1alpha1AccessControlListSpecSubjectsInner, bool) {
	if o == nil || IsNil(o.Subjects) {
		return nil, false
	}
	return o.Subjects, true
}

// HasSubjects returns a boolean if a field has been set.
func (o *DefinitionsV1alpha1AccessControlListSpec) HasSubjects() bool {
	if o != nil && !IsNil(o.Subjects) {
		return true
	}

	return false
}

// SetSubjects gets a reference to the given []CatalogV1alpha1AccessControlListSpecSubjectsInner and assigns it to the Subjects field.
func (o *DefinitionsV1alpha1AccessControlListSpec) SetSubjects(v []CatalogV1alpha1AccessControlListSpecSubjectsInner) {
	o.Subjects = v
}

func (o DefinitionsV1alpha1AccessControlListSpec) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DefinitionsV1alpha1AccessControlListSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Rules) {
		toSerialize["rules"] = o.Rules
	}
	if !IsNil(o.Subjects) {
		toSerialize["subjects"] = o.Subjects
	}
	return toSerialize, nil
}

type NullableDefinitionsV1alpha1AccessControlListSpec struct {
	value *DefinitionsV1alpha1AccessControlListSpec
	isSet bool
}

func (v NullableDefinitionsV1alpha1AccessControlListSpec) Get() *DefinitionsV1alpha1AccessControlListSpec {
	return v.value
}

func (v *NullableDefinitionsV1alpha1AccessControlListSpec) Set(val *DefinitionsV1alpha1AccessControlListSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableDefinitionsV1alpha1AccessControlListSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableDefinitionsV1alpha1AccessControlListSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDefinitionsV1alpha1AccessControlListSpec(val *DefinitionsV1alpha1AccessControlListSpec) *NullableDefinitionsV1alpha1AccessControlListSpec {
	return &NullableDefinitionsV1alpha1AccessControlListSpec{value: val, isSet: true}
}

func (v NullableDefinitionsV1alpha1AccessControlListSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDefinitionsV1alpha1AccessControlListSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


