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

// checks if the ManagementV1APISpecLintingJobResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManagementV1APISpecLintingJobResult{}

// ManagementV1APISpecLintingJobResult struct for ManagementV1APISpecLintingJobResult
type ManagementV1APISpecLintingJobResult struct {
	// The API Specification Linting Result details.
	Details []ManagementV1APISpecLintingJobResultDetailsInner `json:"details"`
	Summary ManagementV1APISpecLintingJobResultSummary `json:"summary"`
	// Reference to the APISpecLintingRuleset revision
	ApiSpecLintingRulesetRevision *int32 `json:"apiSpecLintingRulesetRevision,omitempty"`
	// Set the value to true if the linting result details count has reached the threshold
	DetailsThresholdExceeded *bool `json:"detailsThresholdExceeded,omitempty"`
}

type _ManagementV1APISpecLintingJobResult ManagementV1APISpecLintingJobResult

// NewManagementV1APISpecLintingJobResult instantiates a new ManagementV1APISpecLintingJobResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagementV1APISpecLintingJobResult(details []ManagementV1APISpecLintingJobResultDetailsInner, summary ManagementV1APISpecLintingJobResultSummary) *ManagementV1APISpecLintingJobResult {
	this := ManagementV1APISpecLintingJobResult{}
	this.Details = details
	this.Summary = summary
	return &this
}

// NewManagementV1APISpecLintingJobResultWithDefaults instantiates a new ManagementV1APISpecLintingJobResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagementV1APISpecLintingJobResultWithDefaults() *ManagementV1APISpecLintingJobResult {
	this := ManagementV1APISpecLintingJobResult{}
	return &this
}

// GetDetails returns the Details field value
func (o *ManagementV1APISpecLintingJobResult) GetDetails() []ManagementV1APISpecLintingJobResultDetailsInner {
	if o == nil {
		var ret []ManagementV1APISpecLintingJobResultDetailsInner
		return ret
	}

	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value
// and a boolean to check if the value has been set.
func (o *ManagementV1APISpecLintingJobResult) GetDetailsOk() ([]ManagementV1APISpecLintingJobResultDetailsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Details, true
}

// SetDetails sets field value
func (o *ManagementV1APISpecLintingJobResult) SetDetails(v []ManagementV1APISpecLintingJobResultDetailsInner) {
	o.Details = v
}

// GetSummary returns the Summary field value
func (o *ManagementV1APISpecLintingJobResult) GetSummary() ManagementV1APISpecLintingJobResultSummary {
	if o == nil {
		var ret ManagementV1APISpecLintingJobResultSummary
		return ret
	}

	return o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value
// and a boolean to check if the value has been set.
func (o *ManagementV1APISpecLintingJobResult) GetSummaryOk() (*ManagementV1APISpecLintingJobResultSummary, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Summary, true
}

// SetSummary sets field value
func (o *ManagementV1APISpecLintingJobResult) SetSummary(v ManagementV1APISpecLintingJobResultSummary) {
	o.Summary = v
}

// GetApiSpecLintingRulesetRevision returns the ApiSpecLintingRulesetRevision field value if set, zero value otherwise.
func (o *ManagementV1APISpecLintingJobResult) GetApiSpecLintingRulesetRevision() int32 {
	if o == nil || IsNil(o.ApiSpecLintingRulesetRevision) {
		var ret int32
		return ret
	}
	return *o.ApiSpecLintingRulesetRevision
}

// GetApiSpecLintingRulesetRevisionOk returns a tuple with the ApiSpecLintingRulesetRevision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1APISpecLintingJobResult) GetApiSpecLintingRulesetRevisionOk() (*int32, bool) {
	if o == nil || IsNil(o.ApiSpecLintingRulesetRevision) {
		return nil, false
	}
	return o.ApiSpecLintingRulesetRevision, true
}

// HasApiSpecLintingRulesetRevision returns a boolean if a field has been set.
func (o *ManagementV1APISpecLintingJobResult) HasApiSpecLintingRulesetRevision() bool {
	if o != nil && !IsNil(o.ApiSpecLintingRulesetRevision) {
		return true
	}

	return false
}

// SetApiSpecLintingRulesetRevision gets a reference to the given int32 and assigns it to the ApiSpecLintingRulesetRevision field.
func (o *ManagementV1APISpecLintingJobResult) SetApiSpecLintingRulesetRevision(v int32) {
	o.ApiSpecLintingRulesetRevision = &v
}

// GetDetailsThresholdExceeded returns the DetailsThresholdExceeded field value if set, zero value otherwise.
func (o *ManagementV1APISpecLintingJobResult) GetDetailsThresholdExceeded() bool {
	if o == nil || IsNil(o.DetailsThresholdExceeded) {
		var ret bool
		return ret
	}
	return *o.DetailsThresholdExceeded
}

// GetDetailsThresholdExceededOk returns a tuple with the DetailsThresholdExceeded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1APISpecLintingJobResult) GetDetailsThresholdExceededOk() (*bool, bool) {
	if o == nil || IsNil(o.DetailsThresholdExceeded) {
		return nil, false
	}
	return o.DetailsThresholdExceeded, true
}

// HasDetailsThresholdExceeded returns a boolean if a field has been set.
func (o *ManagementV1APISpecLintingJobResult) HasDetailsThresholdExceeded() bool {
	if o != nil && !IsNil(o.DetailsThresholdExceeded) {
		return true
	}

	return false
}

// SetDetailsThresholdExceeded gets a reference to the given bool and assigns it to the DetailsThresholdExceeded field.
func (o *ManagementV1APISpecLintingJobResult) SetDetailsThresholdExceeded(v bool) {
	o.DetailsThresholdExceeded = &v
}

func (o ManagementV1APISpecLintingJobResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManagementV1APISpecLintingJobResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["details"] = o.Details
	toSerialize["summary"] = o.Summary
	if !IsNil(o.ApiSpecLintingRulesetRevision) {
		toSerialize["apiSpecLintingRulesetRevision"] = o.ApiSpecLintingRulesetRevision
	}
	if !IsNil(o.DetailsThresholdExceeded) {
		toSerialize["detailsThresholdExceeded"] = o.DetailsThresholdExceeded
	}
	return toSerialize, nil
}

func (o *ManagementV1APISpecLintingJobResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"details",
		"summary",
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

	varManagementV1APISpecLintingJobResult := _ManagementV1APISpecLintingJobResult{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varManagementV1APISpecLintingJobResult)

	if err != nil {
		return err
	}

	*o = ManagementV1APISpecLintingJobResult(varManagementV1APISpecLintingJobResult)

	return err
}

type NullableManagementV1APISpecLintingJobResult struct {
	value *ManagementV1APISpecLintingJobResult
	isSet bool
}

func (v NullableManagementV1APISpecLintingJobResult) Get() *ManagementV1APISpecLintingJobResult {
	return v.value
}

func (v *NullableManagementV1APISpecLintingJobResult) Set(val *ManagementV1APISpecLintingJobResult) {
	v.value = val
	v.isSet = true
}

func (v NullableManagementV1APISpecLintingJobResult) IsSet() bool {
	return v.isSet
}

func (v *NullableManagementV1APISpecLintingJobResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagementV1APISpecLintingJobResult(val *ManagementV1APISpecLintingJobResult) *NullableManagementV1APISpecLintingJobResult {
	return &NullableManagementV1APISpecLintingJobResult{value: val, isSet: true}
}

func (v NullableManagementV1APISpecLintingJobResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagementV1APISpecLintingJobResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


