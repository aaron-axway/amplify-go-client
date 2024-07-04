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

// checks if the DefinitionsV1alpha1CommandLineInterfaceSpecColumnsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DefinitionsV1alpha1CommandLineInterfaceSpecColumnsInner{}

// DefinitionsV1alpha1CommandLineInterfaceSpecColumnsInner struct for DefinitionsV1alpha1CommandLineInterfaceSpecColumnsInner
type DefinitionsV1alpha1CommandLineInterfaceSpecColumnsInner struct {
	// The name of the column for the resource.
	Name string `json:"name"`
	// The type of the column.
	Type string `json:"type"`
	// The JSONPath used to fetch data for the column starting from the root.
	JsonPath string `json:"jsonPath"`
	// The description of the column data.
	Description *string `json:"description,omitempty"`
}

type _DefinitionsV1alpha1CommandLineInterfaceSpecColumnsInner DefinitionsV1alpha1CommandLineInterfaceSpecColumnsInner

// NewDefinitionsV1alpha1CommandLineInterfaceSpecColumnsInner instantiates a new DefinitionsV1alpha1CommandLineInterfaceSpecColumnsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDefinitionsV1alpha1CommandLineInterfaceSpecColumnsInner(name string, type_ string, jsonPath string) *DefinitionsV1alpha1CommandLineInterfaceSpecColumnsInner {
	this := DefinitionsV1alpha1CommandLineInterfaceSpecColumnsInner{}
	this.Name = name
	this.Type = type_
	this.JsonPath = jsonPath
	return &this
}

// NewDefinitionsV1alpha1CommandLineInterfaceSpecColumnsInnerWithDefaults instantiates a new DefinitionsV1alpha1CommandLineInterfaceSpecColumnsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDefinitionsV1alpha1CommandLineInterfaceSpecColumnsInnerWithDefaults() *DefinitionsV1alpha1CommandLineInterfaceSpecColumnsInner {
	this := DefinitionsV1alpha1CommandLineInterfaceSpecColumnsInner{}
	return &this
}

// GetName returns the Name field value
func (o *DefinitionsV1alpha1CommandLineInterfaceSpecColumnsInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DefinitionsV1alpha1CommandLineInterfaceSpecColumnsInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DefinitionsV1alpha1CommandLineInterfaceSpecColumnsInner) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *DefinitionsV1alpha1CommandLineInterfaceSpecColumnsInner) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DefinitionsV1alpha1CommandLineInterfaceSpecColumnsInner) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DefinitionsV1alpha1CommandLineInterfaceSpecColumnsInner) SetType(v string) {
	o.Type = v
}

// GetJsonPath returns the JsonPath field value
func (o *DefinitionsV1alpha1CommandLineInterfaceSpecColumnsInner) GetJsonPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.JsonPath
}

// GetJsonPathOk returns a tuple with the JsonPath field value
// and a boolean to check if the value has been set.
func (o *DefinitionsV1alpha1CommandLineInterfaceSpecColumnsInner) GetJsonPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JsonPath, true
}

// SetJsonPath sets field value
func (o *DefinitionsV1alpha1CommandLineInterfaceSpecColumnsInner) SetJsonPath(v string) {
	o.JsonPath = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DefinitionsV1alpha1CommandLineInterfaceSpecColumnsInner) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefinitionsV1alpha1CommandLineInterfaceSpecColumnsInner) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DefinitionsV1alpha1CommandLineInterfaceSpecColumnsInner) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DefinitionsV1alpha1CommandLineInterfaceSpecColumnsInner) SetDescription(v string) {
	o.Description = &v
}

func (o DefinitionsV1alpha1CommandLineInterfaceSpecColumnsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DefinitionsV1alpha1CommandLineInterfaceSpecColumnsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	toSerialize["jsonPath"] = o.JsonPath
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

func (o *DefinitionsV1alpha1CommandLineInterfaceSpecColumnsInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"type",
		"jsonPath",
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

	varDefinitionsV1alpha1CommandLineInterfaceSpecColumnsInner := _DefinitionsV1alpha1CommandLineInterfaceSpecColumnsInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDefinitionsV1alpha1CommandLineInterfaceSpecColumnsInner)

	if err != nil {
		return err
	}

	*o = DefinitionsV1alpha1CommandLineInterfaceSpecColumnsInner(varDefinitionsV1alpha1CommandLineInterfaceSpecColumnsInner)

	return err
}

type NullableDefinitionsV1alpha1CommandLineInterfaceSpecColumnsInner struct {
	value *DefinitionsV1alpha1CommandLineInterfaceSpecColumnsInner
	isSet bool
}

func (v NullableDefinitionsV1alpha1CommandLineInterfaceSpecColumnsInner) Get() *DefinitionsV1alpha1CommandLineInterfaceSpecColumnsInner {
	return v.value
}

func (v *NullableDefinitionsV1alpha1CommandLineInterfaceSpecColumnsInner) Set(val *DefinitionsV1alpha1CommandLineInterfaceSpecColumnsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableDefinitionsV1alpha1CommandLineInterfaceSpecColumnsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableDefinitionsV1alpha1CommandLineInterfaceSpecColumnsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDefinitionsV1alpha1CommandLineInterfaceSpecColumnsInner(val *DefinitionsV1alpha1CommandLineInterfaceSpecColumnsInner) *NullableDefinitionsV1alpha1CommandLineInterfaceSpecColumnsInner {
	return &NullableDefinitionsV1alpha1CommandLineInterfaceSpecColumnsInner{value: val, isSet: true}
}

func (v NullableDefinitionsV1alpha1CommandLineInterfaceSpecColumnsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDefinitionsV1alpha1CommandLineInterfaceSpecColumnsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


