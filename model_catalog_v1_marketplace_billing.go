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

// checks if the CatalogV1MarketplaceBilling type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogV1MarketplaceBilling{}

// CatalogV1MarketplaceBilling Billing options for the Marketplace.
type CatalogV1MarketplaceBilling struct {
	Payment *CatalogV1MarketplaceBillingPayment `json:"payment,omitempty"`
}

// NewCatalogV1MarketplaceBilling instantiates a new CatalogV1MarketplaceBilling object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogV1MarketplaceBilling() *CatalogV1MarketplaceBilling {
	this := CatalogV1MarketplaceBilling{}
	return &this
}

// NewCatalogV1MarketplaceBillingWithDefaults instantiates a new CatalogV1MarketplaceBilling object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogV1MarketplaceBillingWithDefaults() *CatalogV1MarketplaceBilling {
	this := CatalogV1MarketplaceBilling{}
	return &this
}

// GetPayment returns the Payment field value if set, zero value otherwise.
func (o *CatalogV1MarketplaceBilling) GetPayment() CatalogV1MarketplaceBillingPayment {
	if o == nil || IsNil(o.Payment) {
		var ret CatalogV1MarketplaceBillingPayment
		return ret
	}
	return *o.Payment
}

// GetPaymentOk returns a tuple with the Payment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1MarketplaceBilling) GetPaymentOk() (*CatalogV1MarketplaceBillingPayment, bool) {
	if o == nil || IsNil(o.Payment) {
		return nil, false
	}
	return o.Payment, true
}

// HasPayment returns a boolean if a field has been set.
func (o *CatalogV1MarketplaceBilling) HasPayment() bool {
	if o != nil && !IsNil(o.Payment) {
		return true
	}

	return false
}

// SetPayment gets a reference to the given CatalogV1MarketplaceBillingPayment and assigns it to the Payment field.
func (o *CatalogV1MarketplaceBilling) SetPayment(v CatalogV1MarketplaceBillingPayment) {
	o.Payment = &v
}

func (o CatalogV1MarketplaceBilling) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogV1MarketplaceBilling) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Payment) {
		toSerialize["payment"] = o.Payment
	}
	return toSerialize, nil
}

type NullableCatalogV1MarketplaceBilling struct {
	value *CatalogV1MarketplaceBilling
	isSet bool
}

func (v NullableCatalogV1MarketplaceBilling) Get() *CatalogV1MarketplaceBilling {
	return v.value
}

func (v *NullableCatalogV1MarketplaceBilling) Set(val *CatalogV1MarketplaceBilling) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogV1MarketplaceBilling) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogV1MarketplaceBilling) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogV1MarketplaceBilling(val *CatalogV1MarketplaceBilling) *NullableCatalogV1MarketplaceBilling {
	return &NullableCatalogV1MarketplaceBilling{value: val, isSet: true}
}

func (v NullableCatalogV1MarketplaceBilling) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogV1MarketplaceBilling) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


