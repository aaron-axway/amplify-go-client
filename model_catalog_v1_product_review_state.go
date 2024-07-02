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
)

// checks if the CatalogV1ProductReviewState type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogV1ProductReviewState{}

// CatalogV1ProductReviewState struct for CatalogV1ProductReviewState
type CatalogV1ProductReviewState struct {
	Name *string `json:"name,omitempty"`
	// Additional info on the state.
	Reason *string `json:"reason,omitempty"`
}

// NewCatalogV1ProductReviewState instantiates a new CatalogV1ProductReviewState object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogV1ProductReviewState() *CatalogV1ProductReviewState {
	this := CatalogV1ProductReviewState{}
	return &this
}

// NewCatalogV1ProductReviewStateWithDefaults instantiates a new CatalogV1ProductReviewState object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogV1ProductReviewStateWithDefaults() *CatalogV1ProductReviewState {
	this := CatalogV1ProductReviewState{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CatalogV1ProductReviewState) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1ProductReviewState) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CatalogV1ProductReviewState) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CatalogV1ProductReviewState) SetName(v string) {
	o.Name = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *CatalogV1ProductReviewState) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1ProductReviewState) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *CatalogV1ProductReviewState) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *CatalogV1ProductReviewState) SetReason(v string) {
	o.Reason = &v
}

func (o CatalogV1ProductReviewState) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogV1ProductReviewState) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	return toSerialize, nil
}

type NullableCatalogV1ProductReviewState struct {
	value *CatalogV1ProductReviewState
	isSet bool
}

func (v NullableCatalogV1ProductReviewState) Get() *CatalogV1ProductReviewState {
	return v.value
}

func (v *NullableCatalogV1ProductReviewState) Set(val *CatalogV1ProductReviewState) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogV1ProductReviewState) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogV1ProductReviewState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogV1ProductReviewState(val *CatalogV1ProductReviewState) *NullableCatalogV1ProductReviewState {
	return &NullableCatalogV1ProductReviewState{value: val, isSet: true}
}

func (v NullableCatalogV1ProductReviewState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogV1ProductReviewState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

