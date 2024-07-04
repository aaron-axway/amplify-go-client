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
	"gopkg.in/validator.v2"
	"fmt"
)

// CatalogV1alpha1SubscriptionJobSpecPostExecuteOnSuccessInnerAction - struct for CatalogV1alpha1SubscriptionJobSpecPostExecuteOnSuccessInnerAction
type CatalogV1alpha1SubscriptionJobSpecPostExecuteOnSuccessInnerAction struct {
	CatalogV1alpha1SubscriptionJobSpecApprovalStateChange *CatalogV1alpha1SubscriptionJobSpecApprovalStateChange
	CatalogV1alpha1SubscriptionJobSpecMigrateAssetRequests *CatalogV1alpha1SubscriptionJobSpecMigrateAssetRequests
}

// CatalogV1alpha1SubscriptionJobSpecApprovalStateChangeAsCatalogV1alpha1SubscriptionJobSpecPostExecuteOnSuccessInnerAction is a convenience function that returns CatalogV1alpha1SubscriptionJobSpecApprovalStateChange wrapped in CatalogV1alpha1SubscriptionJobSpecPostExecuteOnSuccessInnerAction
func CatalogV1alpha1SubscriptionJobSpecApprovalStateChangeAsCatalogV1alpha1SubscriptionJobSpecPostExecuteOnSuccessInnerAction(v *CatalogV1alpha1SubscriptionJobSpecApprovalStateChange) CatalogV1alpha1SubscriptionJobSpecPostExecuteOnSuccessInnerAction {
	return CatalogV1alpha1SubscriptionJobSpecPostExecuteOnSuccessInnerAction{
		CatalogV1alpha1SubscriptionJobSpecApprovalStateChange: v,
	}
}

// CatalogV1alpha1SubscriptionJobSpecMigrateAssetRequestsAsCatalogV1alpha1SubscriptionJobSpecPostExecuteOnSuccessInnerAction is a convenience function that returns CatalogV1alpha1SubscriptionJobSpecMigrateAssetRequests wrapped in CatalogV1alpha1SubscriptionJobSpecPostExecuteOnSuccessInnerAction
func CatalogV1alpha1SubscriptionJobSpecMigrateAssetRequestsAsCatalogV1alpha1SubscriptionJobSpecPostExecuteOnSuccessInnerAction(v *CatalogV1alpha1SubscriptionJobSpecMigrateAssetRequests) CatalogV1alpha1SubscriptionJobSpecPostExecuteOnSuccessInnerAction {
	return CatalogV1alpha1SubscriptionJobSpecPostExecuteOnSuccessInnerAction{
		CatalogV1alpha1SubscriptionJobSpecMigrateAssetRequests: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *CatalogV1alpha1SubscriptionJobSpecPostExecuteOnSuccessInnerAction) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CatalogV1alpha1SubscriptionJobSpecApprovalStateChange
	err = newStrictDecoder(data).Decode(&dst.CatalogV1alpha1SubscriptionJobSpecApprovalStateChange)
	if err == nil {
		jsonCatalogV1alpha1SubscriptionJobSpecApprovalStateChange, _ := json.Marshal(dst.CatalogV1alpha1SubscriptionJobSpecApprovalStateChange)
		if string(jsonCatalogV1alpha1SubscriptionJobSpecApprovalStateChange) == "{}" { // empty struct
			dst.CatalogV1alpha1SubscriptionJobSpecApprovalStateChange = nil
		} else {
			if err = validator.Validate(dst.CatalogV1alpha1SubscriptionJobSpecApprovalStateChange); err != nil {
				dst.CatalogV1alpha1SubscriptionJobSpecApprovalStateChange = nil
			} else {
				match++
			}
		}
	} else {
		dst.CatalogV1alpha1SubscriptionJobSpecApprovalStateChange = nil
	}

	// try to unmarshal data into CatalogV1alpha1SubscriptionJobSpecMigrateAssetRequests
	err = newStrictDecoder(data).Decode(&dst.CatalogV1alpha1SubscriptionJobSpecMigrateAssetRequests)
	if err == nil {
		jsonCatalogV1alpha1SubscriptionJobSpecMigrateAssetRequests, _ := json.Marshal(dst.CatalogV1alpha1SubscriptionJobSpecMigrateAssetRequests)
		if string(jsonCatalogV1alpha1SubscriptionJobSpecMigrateAssetRequests) == "{}" { // empty struct
			dst.CatalogV1alpha1SubscriptionJobSpecMigrateAssetRequests = nil
		} else {
			if err = validator.Validate(dst.CatalogV1alpha1SubscriptionJobSpecMigrateAssetRequests); err != nil {
				dst.CatalogV1alpha1SubscriptionJobSpecMigrateAssetRequests = nil
			} else {
				match++
			}
		}
	} else {
		dst.CatalogV1alpha1SubscriptionJobSpecMigrateAssetRequests = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.CatalogV1alpha1SubscriptionJobSpecApprovalStateChange = nil
		dst.CatalogV1alpha1SubscriptionJobSpecMigrateAssetRequests = nil

		return fmt.Errorf("data matches more than one schema in oneOf(CatalogV1alpha1SubscriptionJobSpecPostExecuteOnSuccessInnerAction)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(CatalogV1alpha1SubscriptionJobSpecPostExecuteOnSuccessInnerAction)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CatalogV1alpha1SubscriptionJobSpecPostExecuteOnSuccessInnerAction) MarshalJSON() ([]byte, error) {
	if src.CatalogV1alpha1SubscriptionJobSpecApprovalStateChange != nil {
		return json.Marshal(&src.CatalogV1alpha1SubscriptionJobSpecApprovalStateChange)
	}

	if src.CatalogV1alpha1SubscriptionJobSpecMigrateAssetRequests != nil {
		return json.Marshal(&src.CatalogV1alpha1SubscriptionJobSpecMigrateAssetRequests)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CatalogV1alpha1SubscriptionJobSpecPostExecuteOnSuccessInnerAction) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.CatalogV1alpha1SubscriptionJobSpecApprovalStateChange != nil {
		return obj.CatalogV1alpha1SubscriptionJobSpecApprovalStateChange
	}

	if obj.CatalogV1alpha1SubscriptionJobSpecMigrateAssetRequests != nil {
		return obj.CatalogV1alpha1SubscriptionJobSpecMigrateAssetRequests
	}

	// all schemas are nil
	return nil
}

type NullableCatalogV1alpha1SubscriptionJobSpecPostExecuteOnSuccessInnerAction struct {
	value *CatalogV1alpha1SubscriptionJobSpecPostExecuteOnSuccessInnerAction
	isSet bool
}

func (v NullableCatalogV1alpha1SubscriptionJobSpecPostExecuteOnSuccessInnerAction) Get() *CatalogV1alpha1SubscriptionJobSpecPostExecuteOnSuccessInnerAction {
	return v.value
}

func (v *NullableCatalogV1alpha1SubscriptionJobSpecPostExecuteOnSuccessInnerAction) Set(val *CatalogV1alpha1SubscriptionJobSpecPostExecuteOnSuccessInnerAction) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogV1alpha1SubscriptionJobSpecPostExecuteOnSuccessInnerAction) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogV1alpha1SubscriptionJobSpecPostExecuteOnSuccessInnerAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogV1alpha1SubscriptionJobSpecPostExecuteOnSuccessInnerAction(val *CatalogV1alpha1SubscriptionJobSpecPostExecuteOnSuccessInnerAction) *NullableCatalogV1alpha1SubscriptionJobSpecPostExecuteOnSuccessInnerAction {
	return &NullableCatalogV1alpha1SubscriptionJobSpecPostExecuteOnSuccessInnerAction{value: val, isSet: true}
}

func (v NullableCatalogV1alpha1SubscriptionJobSpecPostExecuteOnSuccessInnerAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogV1alpha1SubscriptionJobSpecPostExecuteOnSuccessInnerAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


