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

// checks if the CatalogV1alpha1AssetRequestApprovalState type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogV1alpha1AssetRequestApprovalState{}

// CatalogV1alpha1AssetRequestApprovalState struct for CatalogV1alpha1AssetRequestApprovalState
type CatalogV1alpha1AssetRequestApprovalState struct {
	Name string `json:"name"`
	// Additional info on the state.
	Reason *string `json:"reason,omitempty"`
}

type _CatalogV1alpha1AssetRequestApprovalState CatalogV1alpha1AssetRequestApprovalState

// NewCatalogV1alpha1AssetRequestApprovalState instantiates a new CatalogV1alpha1AssetRequestApprovalState object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogV1alpha1AssetRequestApprovalState(name string) *CatalogV1alpha1AssetRequestApprovalState {
	this := CatalogV1alpha1AssetRequestApprovalState{}
	this.Name = name
	return &this
}

// NewCatalogV1alpha1AssetRequestApprovalStateWithDefaults instantiates a new CatalogV1alpha1AssetRequestApprovalState object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogV1alpha1AssetRequestApprovalStateWithDefaults() *CatalogV1alpha1AssetRequestApprovalState {
	this := CatalogV1alpha1AssetRequestApprovalState{}
	return &this
}

// GetName returns the Name field value
func (o *CatalogV1alpha1AssetRequestApprovalState) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1AssetRequestApprovalState) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CatalogV1alpha1AssetRequestApprovalState) SetName(v string) {
	o.Name = v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *CatalogV1alpha1AssetRequestApprovalState) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1AssetRequestApprovalState) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *CatalogV1alpha1AssetRequestApprovalState) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *CatalogV1alpha1AssetRequestApprovalState) SetReason(v string) {
	o.Reason = &v
}

func (o CatalogV1alpha1AssetRequestApprovalState) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogV1alpha1AssetRequestApprovalState) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	return toSerialize, nil
}

func (o *CatalogV1alpha1AssetRequestApprovalState) UnmarshalJSON(data []byte) (err error) {
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

	varCatalogV1alpha1AssetRequestApprovalState := _CatalogV1alpha1AssetRequestApprovalState{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCatalogV1alpha1AssetRequestApprovalState)

	if err != nil {
		return err
	}

	*o = CatalogV1alpha1AssetRequestApprovalState(varCatalogV1alpha1AssetRequestApprovalState)

	return err
}

type NullableCatalogV1alpha1AssetRequestApprovalState struct {
	value *CatalogV1alpha1AssetRequestApprovalState
	isSet bool
}

func (v NullableCatalogV1alpha1AssetRequestApprovalState) Get() *CatalogV1alpha1AssetRequestApprovalState {
	return v.value
}

func (v *NullableCatalogV1alpha1AssetRequestApprovalState) Set(val *CatalogV1alpha1AssetRequestApprovalState) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogV1alpha1AssetRequestApprovalState) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogV1alpha1AssetRequestApprovalState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogV1alpha1AssetRequestApprovalState(val *CatalogV1alpha1AssetRequestApprovalState) *NullableCatalogV1alpha1AssetRequestApprovalState {
	return &NullableCatalogV1alpha1AssetRequestApprovalState{value: val, isSet: true}
}

func (v NullableCatalogV1alpha1AssetRequestApprovalState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogV1alpha1AssetRequestApprovalState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

