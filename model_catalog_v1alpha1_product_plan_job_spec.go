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

// checks if the CatalogV1alpha1ProductPlanJobSpec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogV1alpha1ProductPlanJobSpec{}

// CatalogV1alpha1ProductPlanJobSpec struct for CatalogV1alpha1ProductPlanJobSpec
type CatalogV1alpha1ProductPlanJobSpec struct {
	Action CatalogV1alpha1ProductPlanJobSpecAction `json:"action"`
	When *CatalogV1alpha1SubscriptionJobSpecWhen `json:"when,omitempty"`
}

type _CatalogV1alpha1ProductPlanJobSpec CatalogV1alpha1ProductPlanJobSpec

// NewCatalogV1alpha1ProductPlanJobSpec instantiates a new CatalogV1alpha1ProductPlanJobSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogV1alpha1ProductPlanJobSpec(action CatalogV1alpha1ProductPlanJobSpecAction) *CatalogV1alpha1ProductPlanJobSpec {
	this := CatalogV1alpha1ProductPlanJobSpec{}
	this.Action = action
	return &this
}

// NewCatalogV1alpha1ProductPlanJobSpecWithDefaults instantiates a new CatalogV1alpha1ProductPlanJobSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogV1alpha1ProductPlanJobSpecWithDefaults() *CatalogV1alpha1ProductPlanJobSpec {
	this := CatalogV1alpha1ProductPlanJobSpec{}
	return &this
}

// GetAction returns the Action field value
func (o *CatalogV1alpha1ProductPlanJobSpec) GetAction() CatalogV1alpha1ProductPlanJobSpecAction {
	if o == nil {
		var ret CatalogV1alpha1ProductPlanJobSpecAction
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1ProductPlanJobSpec) GetActionOk() (*CatalogV1alpha1ProductPlanJobSpecAction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *CatalogV1alpha1ProductPlanJobSpec) SetAction(v CatalogV1alpha1ProductPlanJobSpecAction) {
	o.Action = v
}

// GetWhen returns the When field value if set, zero value otherwise.
func (o *CatalogV1alpha1ProductPlanJobSpec) GetWhen() CatalogV1alpha1SubscriptionJobSpecWhen {
	if o == nil || IsNil(o.When) {
		var ret CatalogV1alpha1SubscriptionJobSpecWhen
		return ret
	}
	return *o.When
}

// GetWhenOk returns a tuple with the When field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1ProductPlanJobSpec) GetWhenOk() (*CatalogV1alpha1SubscriptionJobSpecWhen, bool) {
	if o == nil || IsNil(o.When) {
		return nil, false
	}
	return o.When, true
}

// HasWhen returns a boolean if a field has been set.
func (o *CatalogV1alpha1ProductPlanJobSpec) HasWhen() bool {
	if o != nil && !IsNil(o.When) {
		return true
	}

	return false
}

// SetWhen gets a reference to the given CatalogV1alpha1SubscriptionJobSpecWhen and assigns it to the When field.
func (o *CatalogV1alpha1ProductPlanJobSpec) SetWhen(v CatalogV1alpha1SubscriptionJobSpecWhen) {
	o.When = &v
}

func (o CatalogV1alpha1ProductPlanJobSpec) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogV1alpha1ProductPlanJobSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["action"] = o.Action
	if !IsNil(o.When) {
		toSerialize["when"] = o.When
	}
	return toSerialize, nil
}

func (o *CatalogV1alpha1ProductPlanJobSpec) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"action",
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

	varCatalogV1alpha1ProductPlanJobSpec := _CatalogV1alpha1ProductPlanJobSpec{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCatalogV1alpha1ProductPlanJobSpec)

	if err != nil {
		return err
	}

	*o = CatalogV1alpha1ProductPlanJobSpec(varCatalogV1alpha1ProductPlanJobSpec)

	return err
}

type NullableCatalogV1alpha1ProductPlanJobSpec struct {
	value *CatalogV1alpha1ProductPlanJobSpec
	isSet bool
}

func (v NullableCatalogV1alpha1ProductPlanJobSpec) Get() *CatalogV1alpha1ProductPlanJobSpec {
	return v.value
}

func (v *NullableCatalogV1alpha1ProductPlanJobSpec) Set(val *CatalogV1alpha1ProductPlanJobSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogV1alpha1ProductPlanJobSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogV1alpha1ProductPlanJobSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogV1alpha1ProductPlanJobSpec(val *CatalogV1alpha1ProductPlanJobSpec) *NullableCatalogV1alpha1ProductPlanJobSpec {
	return &NullableCatalogV1alpha1ProductPlanJobSpec{value: val, isSet: true}
}

func (v NullableCatalogV1alpha1ProductPlanJobSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogV1alpha1ProductPlanJobSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


