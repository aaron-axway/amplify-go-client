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
	"bytes"
	"fmt"
)

// checks if the CatalogV1alpha1StageVisibilitySpec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogV1alpha1StageVisibilitySpec{}

// CatalogV1alpha1StageVisibilitySpec struct for CatalogV1alpha1StageVisibilitySpec
type CatalogV1alpha1StageVisibilitySpec struct {
	// Defines where the visibility settings apply.
	Stages []CatalogV1alpha1StageVisibilitySpecStagesInner `json:"stages"`
	// Determines if the list of subjects should be excluded from the stage visibility.
	Exclude *bool `json:"exclude,omitempty"`
	Subjects []CatalogV1alpha1StageVisibilitySpecSubjectsInner `json:"subjects,omitempty"`
}

type _CatalogV1alpha1StageVisibilitySpec CatalogV1alpha1StageVisibilitySpec

// NewCatalogV1alpha1StageVisibilitySpec instantiates a new CatalogV1alpha1StageVisibilitySpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogV1alpha1StageVisibilitySpec(stages []CatalogV1alpha1StageVisibilitySpecStagesInner) *CatalogV1alpha1StageVisibilitySpec {
	this := CatalogV1alpha1StageVisibilitySpec{}
	this.Stages = stages
	return &this
}

// NewCatalogV1alpha1StageVisibilitySpecWithDefaults instantiates a new CatalogV1alpha1StageVisibilitySpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogV1alpha1StageVisibilitySpecWithDefaults() *CatalogV1alpha1StageVisibilitySpec {
	this := CatalogV1alpha1StageVisibilitySpec{}
	return &this
}

// GetStages returns the Stages field value
func (o *CatalogV1alpha1StageVisibilitySpec) GetStages() []CatalogV1alpha1StageVisibilitySpecStagesInner {
	if o == nil {
		var ret []CatalogV1alpha1StageVisibilitySpecStagesInner
		return ret
	}

	return o.Stages
}

// GetStagesOk returns a tuple with the Stages field value
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1StageVisibilitySpec) GetStagesOk() ([]CatalogV1alpha1StageVisibilitySpecStagesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Stages, true
}

// SetStages sets field value
func (o *CatalogV1alpha1StageVisibilitySpec) SetStages(v []CatalogV1alpha1StageVisibilitySpecStagesInner) {
	o.Stages = v
}

// GetExclude returns the Exclude field value if set, zero value otherwise.
func (o *CatalogV1alpha1StageVisibilitySpec) GetExclude() bool {
	if o == nil || IsNil(o.Exclude) {
		var ret bool
		return ret
	}
	return *o.Exclude
}

// GetExcludeOk returns a tuple with the Exclude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1StageVisibilitySpec) GetExcludeOk() (*bool, bool) {
	if o == nil || IsNil(o.Exclude) {
		return nil, false
	}
	return o.Exclude, true
}

// HasExclude returns a boolean if a field has been set.
func (o *CatalogV1alpha1StageVisibilitySpec) HasExclude() bool {
	if o != nil && !IsNil(o.Exclude) {
		return true
	}

	return false
}

// SetExclude gets a reference to the given bool and assigns it to the Exclude field.
func (o *CatalogV1alpha1StageVisibilitySpec) SetExclude(v bool) {
	o.Exclude = &v
}

// GetSubjects returns the Subjects field value if set, zero value otherwise.
func (o *CatalogV1alpha1StageVisibilitySpec) GetSubjects() []CatalogV1alpha1StageVisibilitySpecSubjectsInner {
	if o == nil || IsNil(o.Subjects) {
		var ret []CatalogV1alpha1StageVisibilitySpecSubjectsInner
		return ret
	}
	return o.Subjects
}

// GetSubjectsOk returns a tuple with the Subjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1StageVisibilitySpec) GetSubjectsOk() ([]CatalogV1alpha1StageVisibilitySpecSubjectsInner, bool) {
	if o == nil || IsNil(o.Subjects) {
		return nil, false
	}
	return o.Subjects, true
}

// HasSubjects returns a boolean if a field has been set.
func (o *CatalogV1alpha1StageVisibilitySpec) HasSubjects() bool {
	if o != nil && !IsNil(o.Subjects) {
		return true
	}

	return false
}

// SetSubjects gets a reference to the given []CatalogV1alpha1StageVisibilitySpecSubjectsInner and assigns it to the Subjects field.
func (o *CatalogV1alpha1StageVisibilitySpec) SetSubjects(v []CatalogV1alpha1StageVisibilitySpecSubjectsInner) {
	o.Subjects = v
}

func (o CatalogV1alpha1StageVisibilitySpec) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogV1alpha1StageVisibilitySpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["stages"] = o.Stages
	if !IsNil(o.Exclude) {
		toSerialize["exclude"] = o.Exclude
	}
	if !IsNil(o.Subjects) {
		toSerialize["subjects"] = o.Subjects
	}
	return toSerialize, nil
}

func (o *CatalogV1alpha1StageVisibilitySpec) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"stages",
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

	varCatalogV1alpha1StageVisibilitySpec := _CatalogV1alpha1StageVisibilitySpec{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCatalogV1alpha1StageVisibilitySpec)

	if err != nil {
		return err
	}

	*o = CatalogV1alpha1StageVisibilitySpec(varCatalogV1alpha1StageVisibilitySpec)

	return err
}

type NullableCatalogV1alpha1StageVisibilitySpec struct {
	value *CatalogV1alpha1StageVisibilitySpec
	isSet bool
}

func (v NullableCatalogV1alpha1StageVisibilitySpec) Get() *CatalogV1alpha1StageVisibilitySpec {
	return v.value
}

func (v *NullableCatalogV1alpha1StageVisibilitySpec) Set(val *CatalogV1alpha1StageVisibilitySpec) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogV1alpha1StageVisibilitySpec) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogV1alpha1StageVisibilitySpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogV1alpha1StageVisibilitySpec(val *CatalogV1alpha1StageVisibilitySpec) *NullableCatalogV1alpha1StageVisibilitySpec {
	return &NullableCatalogV1alpha1StageVisibilitySpec{value: val, isSet: true}
}

func (v NullableCatalogV1alpha1StageVisibilitySpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogV1alpha1StageVisibilitySpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


