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
	"time"
)

// checks if the ManagementV1alpha1APIServiceInstanceSourceRuntimeStatusResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManagementV1alpha1APIServiceInstanceSourceRuntimeStatusResult{}

// ManagementV1alpha1APIServiceInstanceSourceRuntimeStatusResult APIServiceInstance runtime results.
type ManagementV1alpha1APIServiceInstanceSourceRuntimeStatusResult struct {
	// Time when the update occurred.
	Timestamp *time.Time `json:"timestamp,omitempty"`
	// The average risk score in the runtime compliance result.
	RiskScore *float32 `json:"riskScore,omitempty"`
}

// NewManagementV1alpha1APIServiceInstanceSourceRuntimeStatusResult instantiates a new ManagementV1alpha1APIServiceInstanceSourceRuntimeStatusResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagementV1alpha1APIServiceInstanceSourceRuntimeStatusResult() *ManagementV1alpha1APIServiceInstanceSourceRuntimeStatusResult {
	this := ManagementV1alpha1APIServiceInstanceSourceRuntimeStatusResult{}
	return &this
}

// NewManagementV1alpha1APIServiceInstanceSourceRuntimeStatusResultWithDefaults instantiates a new ManagementV1alpha1APIServiceInstanceSourceRuntimeStatusResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagementV1alpha1APIServiceInstanceSourceRuntimeStatusResultWithDefaults() *ManagementV1alpha1APIServiceInstanceSourceRuntimeStatusResult {
	this := ManagementV1alpha1APIServiceInstanceSourceRuntimeStatusResult{}
	return &this
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *ManagementV1alpha1APIServiceInstanceSourceRuntimeStatusResult) GetTimestamp() time.Time {
	if o == nil || IsNil(o.Timestamp) {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1APIServiceInstanceSourceRuntimeStatusResult) GetTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *ManagementV1alpha1APIServiceInstanceSourceRuntimeStatusResult) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *ManagementV1alpha1APIServiceInstanceSourceRuntimeStatusResult) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

// GetRiskScore returns the RiskScore field value if set, zero value otherwise.
func (o *ManagementV1alpha1APIServiceInstanceSourceRuntimeStatusResult) GetRiskScore() float32 {
	if o == nil || IsNil(o.RiskScore) {
		var ret float32
		return ret
	}
	return *o.RiskScore
}

// GetRiskScoreOk returns a tuple with the RiskScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1APIServiceInstanceSourceRuntimeStatusResult) GetRiskScoreOk() (*float32, bool) {
	if o == nil || IsNil(o.RiskScore) {
		return nil, false
	}
	return o.RiskScore, true
}

// HasRiskScore returns a boolean if a field has been set.
func (o *ManagementV1alpha1APIServiceInstanceSourceRuntimeStatusResult) HasRiskScore() bool {
	if o != nil && !IsNil(o.RiskScore) {
		return true
	}

	return false
}

// SetRiskScore gets a reference to the given float32 and assigns it to the RiskScore field.
func (o *ManagementV1alpha1APIServiceInstanceSourceRuntimeStatusResult) SetRiskScore(v float32) {
	o.RiskScore = &v
}

func (o ManagementV1alpha1APIServiceInstanceSourceRuntimeStatusResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManagementV1alpha1APIServiceInstanceSourceRuntimeStatusResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	if !IsNil(o.RiskScore) {
		toSerialize["riskScore"] = o.RiskScore
	}
	return toSerialize, nil
}

type NullableManagementV1alpha1APIServiceInstanceSourceRuntimeStatusResult struct {
	value *ManagementV1alpha1APIServiceInstanceSourceRuntimeStatusResult
	isSet bool
}

func (v NullableManagementV1alpha1APIServiceInstanceSourceRuntimeStatusResult) Get() *ManagementV1alpha1APIServiceInstanceSourceRuntimeStatusResult {
	return v.value
}

func (v *NullableManagementV1alpha1APIServiceInstanceSourceRuntimeStatusResult) Set(val *ManagementV1alpha1APIServiceInstanceSourceRuntimeStatusResult) {
	v.value = val
	v.isSet = true
}

func (v NullableManagementV1alpha1APIServiceInstanceSourceRuntimeStatusResult) IsSet() bool {
	return v.isSet
}

func (v *NullableManagementV1alpha1APIServiceInstanceSourceRuntimeStatusResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagementV1alpha1APIServiceInstanceSourceRuntimeStatusResult(val *ManagementV1alpha1APIServiceInstanceSourceRuntimeStatusResult) *NullableManagementV1alpha1APIServiceInstanceSourceRuntimeStatusResult {
	return &NullableManagementV1alpha1APIServiceInstanceSourceRuntimeStatusResult{value: val, isSet: true}
}

func (v NullableManagementV1alpha1APIServiceInstanceSourceRuntimeStatusResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagementV1alpha1APIServiceInstanceSourceRuntimeStatusResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

