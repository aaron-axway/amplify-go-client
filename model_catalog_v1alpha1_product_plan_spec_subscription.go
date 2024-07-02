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

// checks if the CatalogV1alpha1ProductPlanSpecSubscription type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogV1alpha1ProductPlanSpecSubscription{}

// CatalogV1alpha1ProductPlanSpecSubscription Defines Plan's subscription information
type CatalogV1alpha1ProductPlanSpecSubscription struct {
	// Defines properties required from a consumer to subscribe to the plan.
	Definition *string `json:"definition,omitempty"`
	Interval *CatalogV1alpha1ProductPlanSpecSubscriptionInterval `json:"interval,omitempty"`
	Renewal *string `json:"renewal,omitempty"`
	Approval *string `json:"approval,omitempty"`
	// Optional number of cycles after which the subscription will be archived. Cycles start once the subscription has been approved.
	Cycles *int32 `json:"cycles,omitempty"`
}

// NewCatalogV1alpha1ProductPlanSpecSubscription instantiates a new CatalogV1alpha1ProductPlanSpecSubscription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogV1alpha1ProductPlanSpecSubscription() *CatalogV1alpha1ProductPlanSpecSubscription {
	this := CatalogV1alpha1ProductPlanSpecSubscription{}
	return &this
}

// NewCatalogV1alpha1ProductPlanSpecSubscriptionWithDefaults instantiates a new CatalogV1alpha1ProductPlanSpecSubscription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogV1alpha1ProductPlanSpecSubscriptionWithDefaults() *CatalogV1alpha1ProductPlanSpecSubscription {
	this := CatalogV1alpha1ProductPlanSpecSubscription{}
	return &this
}

// GetDefinition returns the Definition field value if set, zero value otherwise.
func (o *CatalogV1alpha1ProductPlanSpecSubscription) GetDefinition() string {
	if o == nil || IsNil(o.Definition) {
		var ret string
		return ret
	}
	return *o.Definition
}

// GetDefinitionOk returns a tuple with the Definition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1ProductPlanSpecSubscription) GetDefinitionOk() (*string, bool) {
	if o == nil || IsNil(o.Definition) {
		return nil, false
	}
	return o.Definition, true
}

// HasDefinition returns a boolean if a field has been set.
func (o *CatalogV1alpha1ProductPlanSpecSubscription) HasDefinition() bool {
	if o != nil && !IsNil(o.Definition) {
		return true
	}

	return false
}

// SetDefinition gets a reference to the given string and assigns it to the Definition field.
func (o *CatalogV1alpha1ProductPlanSpecSubscription) SetDefinition(v string) {
	o.Definition = &v
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *CatalogV1alpha1ProductPlanSpecSubscription) GetInterval() CatalogV1alpha1ProductPlanSpecSubscriptionInterval {
	if o == nil || IsNil(o.Interval) {
		var ret CatalogV1alpha1ProductPlanSpecSubscriptionInterval
		return ret
	}
	return *o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1ProductPlanSpecSubscription) GetIntervalOk() (*CatalogV1alpha1ProductPlanSpecSubscriptionInterval, bool) {
	if o == nil || IsNil(o.Interval) {
		return nil, false
	}
	return o.Interval, true
}

// HasInterval returns a boolean if a field has been set.
func (o *CatalogV1alpha1ProductPlanSpecSubscription) HasInterval() bool {
	if o != nil && !IsNil(o.Interval) {
		return true
	}

	return false
}

// SetInterval gets a reference to the given CatalogV1alpha1ProductPlanSpecSubscriptionInterval and assigns it to the Interval field.
func (o *CatalogV1alpha1ProductPlanSpecSubscription) SetInterval(v CatalogV1alpha1ProductPlanSpecSubscriptionInterval) {
	o.Interval = &v
}

// GetRenewal returns the Renewal field value if set, zero value otherwise.
func (o *CatalogV1alpha1ProductPlanSpecSubscription) GetRenewal() string {
	if o == nil || IsNil(o.Renewal) {
		var ret string
		return ret
	}
	return *o.Renewal
}

// GetRenewalOk returns a tuple with the Renewal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1ProductPlanSpecSubscription) GetRenewalOk() (*string, bool) {
	if o == nil || IsNil(o.Renewal) {
		return nil, false
	}
	return o.Renewal, true
}

// HasRenewal returns a boolean if a field has been set.
func (o *CatalogV1alpha1ProductPlanSpecSubscription) HasRenewal() bool {
	if o != nil && !IsNil(o.Renewal) {
		return true
	}

	return false
}

// SetRenewal gets a reference to the given string and assigns it to the Renewal field.
func (o *CatalogV1alpha1ProductPlanSpecSubscription) SetRenewal(v string) {
	o.Renewal = &v
}

// GetApproval returns the Approval field value if set, zero value otherwise.
func (o *CatalogV1alpha1ProductPlanSpecSubscription) GetApproval() string {
	if o == nil || IsNil(o.Approval) {
		var ret string
		return ret
	}
	return *o.Approval
}

// GetApprovalOk returns a tuple with the Approval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1ProductPlanSpecSubscription) GetApprovalOk() (*string, bool) {
	if o == nil || IsNil(o.Approval) {
		return nil, false
	}
	return o.Approval, true
}

// HasApproval returns a boolean if a field has been set.
func (o *CatalogV1alpha1ProductPlanSpecSubscription) HasApproval() bool {
	if o != nil && !IsNil(o.Approval) {
		return true
	}

	return false
}

// SetApproval gets a reference to the given string and assigns it to the Approval field.
func (o *CatalogV1alpha1ProductPlanSpecSubscription) SetApproval(v string) {
	o.Approval = &v
}

// GetCycles returns the Cycles field value if set, zero value otherwise.
func (o *CatalogV1alpha1ProductPlanSpecSubscription) GetCycles() int32 {
	if o == nil || IsNil(o.Cycles) {
		var ret int32
		return ret
	}
	return *o.Cycles
}

// GetCyclesOk returns a tuple with the Cycles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1ProductPlanSpecSubscription) GetCyclesOk() (*int32, bool) {
	if o == nil || IsNil(o.Cycles) {
		return nil, false
	}
	return o.Cycles, true
}

// HasCycles returns a boolean if a field has been set.
func (o *CatalogV1alpha1ProductPlanSpecSubscription) HasCycles() bool {
	if o != nil && !IsNil(o.Cycles) {
		return true
	}

	return false
}

// SetCycles gets a reference to the given int32 and assigns it to the Cycles field.
func (o *CatalogV1alpha1ProductPlanSpecSubscription) SetCycles(v int32) {
	o.Cycles = &v
}

func (o CatalogV1alpha1ProductPlanSpecSubscription) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogV1alpha1ProductPlanSpecSubscription) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Definition) {
		toSerialize["definition"] = o.Definition
	}
	if !IsNil(o.Interval) {
		toSerialize["interval"] = o.Interval
	}
	if !IsNil(o.Renewal) {
		toSerialize["renewal"] = o.Renewal
	}
	if !IsNil(o.Approval) {
		toSerialize["approval"] = o.Approval
	}
	if !IsNil(o.Cycles) {
		toSerialize["cycles"] = o.Cycles
	}
	return toSerialize, nil
}

type NullableCatalogV1alpha1ProductPlanSpecSubscription struct {
	value *CatalogV1alpha1ProductPlanSpecSubscription
	isSet bool
}

func (v NullableCatalogV1alpha1ProductPlanSpecSubscription) Get() *CatalogV1alpha1ProductPlanSpecSubscription {
	return v.value
}

func (v *NullableCatalogV1alpha1ProductPlanSpecSubscription) Set(val *CatalogV1alpha1ProductPlanSpecSubscription) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogV1alpha1ProductPlanSpecSubscription) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogV1alpha1ProductPlanSpecSubscription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogV1alpha1ProductPlanSpecSubscription(val *CatalogV1alpha1ProductPlanSpecSubscription) *NullableCatalogV1alpha1ProductPlanSpecSubscription {
	return &NullableCatalogV1alpha1ProductPlanSpecSubscription{value: val, isSet: true}
}

func (v NullableCatalogV1alpha1ProductPlanSpecSubscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogV1alpha1ProductPlanSpecSubscription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

