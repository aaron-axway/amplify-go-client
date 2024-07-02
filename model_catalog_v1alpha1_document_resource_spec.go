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

// checks if the CatalogV1alpha1DocumentResourceSpec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogV1alpha1DocumentResourceSpec{}

// CatalogV1alpha1DocumentResourceSpec struct for CatalogV1alpha1DocumentResourceSpec
type CatalogV1alpha1DocumentResourceSpec struct {
	// Document description.
	Description *string `json:"description,omitempty"`
	// Version of the DocumentResource.
	Version string `json:"version"`
	Usage CatalogV1DocumentResourceSpecUsage `json:"usage"`
	Data CatalogV1alpha1DocumentResourceSpecData `json:"data"`
}

type _CatalogV1alpha1DocumentResourceSpec CatalogV1alpha1DocumentResourceSpec

// NewCatalogV1alpha1DocumentResourceSpec instantiates a new CatalogV1alpha1DocumentResourceSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogV1alpha1DocumentResourceSpec(version string, usage CatalogV1DocumentResourceSpecUsage, data CatalogV1alpha1DocumentResourceSpecData) *CatalogV1alpha1DocumentResourceSpec {
	this := CatalogV1alpha1DocumentResourceSpec{}
	this.Version = version
	this.Usage = usage
	this.Data = data
	return &this
}

// NewCatalogV1alpha1DocumentResourceSpecWithDefaults instantiates a new CatalogV1alpha1DocumentResourceSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogV1alpha1DocumentResourceSpecWithDefaults() *CatalogV1alpha1DocumentResourceSpec {
	this := CatalogV1alpha1DocumentResourceSpec{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CatalogV1alpha1DocumentResourceSpec) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1DocumentResourceSpec) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CatalogV1alpha1DocumentResourceSpec) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CatalogV1alpha1DocumentResourceSpec) SetDescription(v string) {
	o.Description = &v
}

// GetVersion returns the Version field value
func (o *CatalogV1alpha1DocumentResourceSpec) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1DocumentResourceSpec) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *CatalogV1alpha1DocumentResourceSpec) SetVersion(v string) {
	o.Version = v
}

// GetUsage returns the Usage field value
func (o *CatalogV1alpha1DocumentResourceSpec) GetUsage() CatalogV1DocumentResourceSpecUsage {
	if o == nil {
		var ret CatalogV1DocumentResourceSpecUsage
		return ret
	}

	return o.Usage
}

// GetUsageOk returns a tuple with the Usage field value
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1DocumentResourceSpec) GetUsageOk() (*CatalogV1DocumentResourceSpecUsage, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Usage, true
}

// SetUsage sets field value
func (o *CatalogV1alpha1DocumentResourceSpec) SetUsage(v CatalogV1DocumentResourceSpecUsage) {
	o.Usage = v
}

// GetData returns the Data field value
func (o *CatalogV1alpha1DocumentResourceSpec) GetData() CatalogV1alpha1DocumentResourceSpecData {
	if o == nil {
		var ret CatalogV1alpha1DocumentResourceSpecData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1DocumentResourceSpec) GetDataOk() (*CatalogV1alpha1DocumentResourceSpecData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CatalogV1alpha1DocumentResourceSpec) SetData(v CatalogV1alpha1DocumentResourceSpecData) {
	o.Data = v
}

func (o CatalogV1alpha1DocumentResourceSpec) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogV1alpha1DocumentResourceSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["version"] = o.Version
	toSerialize["usage"] = o.Usage
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *CatalogV1alpha1DocumentResourceSpec) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"version",
		"usage",
		"data",
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

	varCatalogV1alpha1DocumentResourceSpec := _CatalogV1alpha1DocumentResourceSpec{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCatalogV1alpha1DocumentResourceSpec)

	if err != nil {
		return err
	}

	*o = CatalogV1alpha1DocumentResourceSpec(varCatalogV1alpha1DocumentResourceSpec)

	return err
}

type NullableCatalogV1alpha1DocumentResourceSpec struct {
	value *CatalogV1alpha1DocumentResourceSpec
	isSet bool
}

func (v NullableCatalogV1alpha1DocumentResourceSpec) Get() *CatalogV1alpha1DocumentResourceSpec {
	return v.value
}

func (v *NullableCatalogV1alpha1DocumentResourceSpec) Set(val *CatalogV1alpha1DocumentResourceSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogV1alpha1DocumentResourceSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogV1alpha1DocumentResourceSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogV1alpha1DocumentResourceSpec(val *CatalogV1alpha1DocumentResourceSpec) *NullableCatalogV1alpha1DocumentResourceSpec {
	return &NullableCatalogV1alpha1DocumentResourceSpec{value: val, isSet: true}
}

func (v NullableCatalogV1alpha1DocumentResourceSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogV1alpha1DocumentResourceSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


