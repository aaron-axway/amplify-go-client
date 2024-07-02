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

// checks if the CatalogV1alpha1SubscriptionBilling type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogV1alpha1SubscriptionBilling{}

// CatalogV1alpha1SubscriptionBilling Details about the Subscription billing.
type CatalogV1alpha1SubscriptionBilling struct {
	Payment *CatalogV1alpha1SubscriptionBillingPayment `json:"payment,omitempty"`
}

// NewCatalogV1alpha1SubscriptionBilling instantiates a new CatalogV1alpha1SubscriptionBilling object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogV1alpha1SubscriptionBilling() *CatalogV1alpha1SubscriptionBilling {
	this := CatalogV1alpha1SubscriptionBilling{}
	return &this
}

// NewCatalogV1alpha1SubscriptionBillingWithDefaults instantiates a new CatalogV1alpha1SubscriptionBilling object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogV1alpha1SubscriptionBillingWithDefaults() *CatalogV1alpha1SubscriptionBilling {
	this := CatalogV1alpha1SubscriptionBilling{}
	return &this
}

// GetPayment returns the Payment field value if set, zero value otherwise.
func (o *CatalogV1alpha1SubscriptionBilling) GetPayment() CatalogV1alpha1SubscriptionBillingPayment {
	if o == nil || IsNil(o.Payment) {
		var ret CatalogV1alpha1SubscriptionBillingPayment
		return ret
	}
	return *o.Payment
}

// GetPaymentOk returns a tuple with the Payment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1SubscriptionBilling) GetPaymentOk() (*CatalogV1alpha1SubscriptionBillingPayment, bool) {
	if o == nil || IsNil(o.Payment) {
		return nil, false
	}
	return o.Payment, true
}

// HasPayment returns a boolean if a field has been set.
func (o *CatalogV1alpha1SubscriptionBilling) HasPayment() bool {
	if o != nil && !IsNil(o.Payment) {
		return true
	}

	return false
}

// SetPayment gets a reference to the given CatalogV1alpha1SubscriptionBillingPayment and assigns it to the Payment field.
func (o *CatalogV1alpha1SubscriptionBilling) SetPayment(v CatalogV1alpha1SubscriptionBillingPayment) {
	o.Payment = &v
}

func (o CatalogV1alpha1SubscriptionBilling) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogV1alpha1SubscriptionBilling) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Payment) {
		toSerialize["payment"] = o.Payment
	}
	return toSerialize, nil
}

type NullableCatalogV1alpha1SubscriptionBilling struct {
	value *CatalogV1alpha1SubscriptionBilling
	isSet bool
}

func (v NullableCatalogV1alpha1SubscriptionBilling) Get() *CatalogV1alpha1SubscriptionBilling {
	return v.value
}

func (v *NullableCatalogV1alpha1SubscriptionBilling) Set(val *CatalogV1alpha1SubscriptionBilling) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogV1alpha1SubscriptionBilling) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogV1alpha1SubscriptionBilling) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogV1alpha1SubscriptionBilling(val *CatalogV1alpha1SubscriptionBilling) *NullableCatalogV1alpha1SubscriptionBilling {
	return &NullableCatalogV1alpha1SubscriptionBilling{value: val, isSet: true}
}

func (v NullableCatalogV1alpha1SubscriptionBilling) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogV1alpha1SubscriptionBilling) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


