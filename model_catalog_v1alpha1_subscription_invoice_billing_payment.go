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
	"gopkg.in/validator.v2"
	"fmt"
)

// CatalogV1alpha1SubscriptionInvoiceBillingPayment - struct for CatalogV1alpha1SubscriptionInvoiceBillingPayment
type CatalogV1alpha1SubscriptionInvoiceBillingPayment struct {
	CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeCustom *CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeCustom
	CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeStripe *CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeStripe
}

// CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeCustomAsCatalogV1alpha1SubscriptionInvoiceBillingPayment is a convenience function that returns CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeCustom wrapped in CatalogV1alpha1SubscriptionInvoiceBillingPayment
func CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeCustomAsCatalogV1alpha1SubscriptionInvoiceBillingPayment(v *CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeCustom) CatalogV1alpha1SubscriptionInvoiceBillingPayment {
	return CatalogV1alpha1SubscriptionInvoiceBillingPayment{
		CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeCustom: v,
	}
}

// CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeStripeAsCatalogV1alpha1SubscriptionInvoiceBillingPayment is a convenience function that returns CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeStripe wrapped in CatalogV1alpha1SubscriptionInvoiceBillingPayment
func CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeStripeAsCatalogV1alpha1SubscriptionInvoiceBillingPayment(v *CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeStripe) CatalogV1alpha1SubscriptionInvoiceBillingPayment {
	return CatalogV1alpha1SubscriptionInvoiceBillingPayment{
		CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeStripe: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *CatalogV1alpha1SubscriptionInvoiceBillingPayment) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeCustom
	err = newStrictDecoder(data).Decode(&dst.CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeCustom)
	if err == nil {
		jsonCatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeCustom, _ := json.Marshal(dst.CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeCustom)
		if string(jsonCatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeCustom) == "{}" { // empty struct
			dst.CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeCustom = nil
		} else {
			if err = validator.Validate(dst.CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeCustom); err != nil {
				dst.CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeCustom = nil
			} else {
				match++
			}
		}
	} else {
		dst.CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeCustom = nil
	}

	// try to unmarshal data into CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeStripe
	err = newStrictDecoder(data).Decode(&dst.CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeStripe)
	if err == nil {
		jsonCatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeStripe, _ := json.Marshal(dst.CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeStripe)
		if string(jsonCatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeStripe) == "{}" { // empty struct
			dst.CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeStripe = nil
		} else {
			if err = validator.Validate(dst.CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeStripe); err != nil {
				dst.CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeStripe = nil
			} else {
				match++
			}
		}
	} else {
		dst.CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeStripe = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeCustom = nil
		dst.CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeStripe = nil

		return fmt.Errorf("data matches more than one schema in oneOf(CatalogV1alpha1SubscriptionInvoiceBillingPayment)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(CatalogV1alpha1SubscriptionInvoiceBillingPayment)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CatalogV1alpha1SubscriptionInvoiceBillingPayment) MarshalJSON() ([]byte, error) {
	if src.CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeCustom != nil {
		return json.Marshal(&src.CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeCustom)
	}

	if src.CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeStripe != nil {
		return json.Marshal(&src.CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeStripe)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CatalogV1alpha1SubscriptionInvoiceBillingPayment) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeCustom != nil {
		return obj.CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeCustom
	}

	if obj.CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeStripe != nil {
		return obj.CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeStripe
	}

	// all schemas are nil
	return nil
}

type NullableCatalogV1alpha1SubscriptionInvoiceBillingPayment struct {
	value *CatalogV1alpha1SubscriptionInvoiceBillingPayment
	isSet bool
}

func (v NullableCatalogV1alpha1SubscriptionInvoiceBillingPayment) Get() *CatalogV1alpha1SubscriptionInvoiceBillingPayment {
	return v.value
}

func (v *NullableCatalogV1alpha1SubscriptionInvoiceBillingPayment) Set(val *CatalogV1alpha1SubscriptionInvoiceBillingPayment) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogV1alpha1SubscriptionInvoiceBillingPayment) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogV1alpha1SubscriptionInvoiceBillingPayment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogV1alpha1SubscriptionInvoiceBillingPayment(val *CatalogV1alpha1SubscriptionInvoiceBillingPayment) *NullableCatalogV1alpha1SubscriptionInvoiceBillingPayment {
	return &NullableCatalogV1alpha1SubscriptionInvoiceBillingPayment{value: val, isSet: true}
}

func (v NullableCatalogV1alpha1SubscriptionInvoiceBillingPayment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogV1alpha1SubscriptionInvoiceBillingPayment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

