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

// checks if the CatalogV1alpha1AssistedSearchState type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogV1alpha1AssistedSearchState{}

// CatalogV1alpha1AssistedSearchState Current state of the integration. The integration can get to be disabled if resources quota is reached.
type CatalogV1alpha1AssistedSearchState struct {
	Name *string `json:"name,omitempty"`
	// Additional info on the state.
	Reason *string `json:"reason,omitempty"`
}

// NewCatalogV1alpha1AssistedSearchState instantiates a new CatalogV1alpha1AssistedSearchState object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogV1alpha1AssistedSearchState() *CatalogV1alpha1AssistedSearchState {
	this := CatalogV1alpha1AssistedSearchState{}
	return &this
}

// NewCatalogV1alpha1AssistedSearchStateWithDefaults instantiates a new CatalogV1alpha1AssistedSearchState object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogV1alpha1AssistedSearchStateWithDefaults() *CatalogV1alpha1AssistedSearchState {
	this := CatalogV1alpha1AssistedSearchState{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CatalogV1alpha1AssistedSearchState) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1AssistedSearchState) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CatalogV1alpha1AssistedSearchState) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CatalogV1alpha1AssistedSearchState) SetName(v string) {
	o.Name = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *CatalogV1alpha1AssistedSearchState) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1AssistedSearchState) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *CatalogV1alpha1AssistedSearchState) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *CatalogV1alpha1AssistedSearchState) SetReason(v string) {
	o.Reason = &v
}

func (o CatalogV1alpha1AssistedSearchState) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogV1alpha1AssistedSearchState) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	return toSerialize, nil
}

type NullableCatalogV1alpha1AssistedSearchState struct {
	value *CatalogV1alpha1AssistedSearchState
	isSet bool
}

func (v NullableCatalogV1alpha1AssistedSearchState) Get() *CatalogV1alpha1AssistedSearchState {
	return v.value
}

func (v *NullableCatalogV1alpha1AssistedSearchState) Set(val *CatalogV1alpha1AssistedSearchState) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogV1alpha1AssistedSearchState) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogV1alpha1AssistedSearchState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogV1alpha1AssistedSearchState(val *CatalogV1alpha1AssistedSearchState) *NullableCatalogV1alpha1AssistedSearchState {
	return &NullableCatalogV1alpha1AssistedSearchState{value: val, isSet: true}
}

func (v NullableCatalogV1alpha1AssistedSearchState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogV1alpha1AssistedSearchState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


