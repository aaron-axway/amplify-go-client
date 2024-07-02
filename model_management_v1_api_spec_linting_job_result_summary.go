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

// checks if the ManagementV1APISpecLintingJobResultSummary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManagementV1APISpecLintingJobResultSummary{}

// ManagementV1APISpecLintingJobResultSummary Summary of the linting results.
type ManagementV1APISpecLintingJobResultSummary struct {
	// The total number of errors in the linting result.
	ErrorCount *int32 `json:"errorCount,omitempty"`
	// The total number of warnings in the linting result.
	WarningCount *int32 `json:"warningCount,omitempty"`
	// The total number of hints in the linting result.
	HintCount *int32 `json:"hintCount,omitempty"`
	// The total number of infos in the linting result.
	InfoCount *int32 `json:"infoCount,omitempty"`
	// Grade result from the results summary.
	Grade *string `json:"grade,omitempty"`
}

// NewManagementV1APISpecLintingJobResultSummary instantiates a new ManagementV1APISpecLintingJobResultSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagementV1APISpecLintingJobResultSummary() *ManagementV1APISpecLintingJobResultSummary {
	this := ManagementV1APISpecLintingJobResultSummary{}
	return &this
}

// NewManagementV1APISpecLintingJobResultSummaryWithDefaults instantiates a new ManagementV1APISpecLintingJobResultSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagementV1APISpecLintingJobResultSummaryWithDefaults() *ManagementV1APISpecLintingJobResultSummary {
	this := ManagementV1APISpecLintingJobResultSummary{}
	return &this
}

// GetErrorCount returns the ErrorCount field value if set, zero value otherwise.
func (o *ManagementV1APISpecLintingJobResultSummary) GetErrorCount() int32 {
	if o == nil || IsNil(o.ErrorCount) {
		var ret int32
		return ret
	}
	return *o.ErrorCount
}

// GetErrorCountOk returns a tuple with the ErrorCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1APISpecLintingJobResultSummary) GetErrorCountOk() (*int32, bool) {
	if o == nil || IsNil(o.ErrorCount) {
		return nil, false
	}
	return o.ErrorCount, true
}

// HasErrorCount returns a boolean if a field has been set.
func (o *ManagementV1APISpecLintingJobResultSummary) HasErrorCount() bool {
	if o != nil && !IsNil(o.ErrorCount) {
		return true
	}

	return false
}

// SetErrorCount gets a reference to the given int32 and assigns it to the ErrorCount field.
func (o *ManagementV1APISpecLintingJobResultSummary) SetErrorCount(v int32) {
	o.ErrorCount = &v
}

// GetWarningCount returns the WarningCount field value if set, zero value otherwise.
func (o *ManagementV1APISpecLintingJobResultSummary) GetWarningCount() int32 {
	if o == nil || IsNil(o.WarningCount) {
		var ret int32
		return ret
	}
	return *o.WarningCount
}

// GetWarningCountOk returns a tuple with the WarningCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1APISpecLintingJobResultSummary) GetWarningCountOk() (*int32, bool) {
	if o == nil || IsNil(o.WarningCount) {
		return nil, false
	}
	return o.WarningCount, true
}

// HasWarningCount returns a boolean if a field has been set.
func (o *ManagementV1APISpecLintingJobResultSummary) HasWarningCount() bool {
	if o != nil && !IsNil(o.WarningCount) {
		return true
	}

	return false
}

// SetWarningCount gets a reference to the given int32 and assigns it to the WarningCount field.
func (o *ManagementV1APISpecLintingJobResultSummary) SetWarningCount(v int32) {
	o.WarningCount = &v
}

// GetHintCount returns the HintCount field value if set, zero value otherwise.
func (o *ManagementV1APISpecLintingJobResultSummary) GetHintCount() int32 {
	if o == nil || IsNil(o.HintCount) {
		var ret int32
		return ret
	}
	return *o.HintCount
}

// GetHintCountOk returns a tuple with the HintCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1APISpecLintingJobResultSummary) GetHintCountOk() (*int32, bool) {
	if o == nil || IsNil(o.HintCount) {
		return nil, false
	}
	return o.HintCount, true
}

// HasHintCount returns a boolean if a field has been set.
func (o *ManagementV1APISpecLintingJobResultSummary) HasHintCount() bool {
	if o != nil && !IsNil(o.HintCount) {
		return true
	}

	return false
}

// SetHintCount gets a reference to the given int32 and assigns it to the HintCount field.
func (o *ManagementV1APISpecLintingJobResultSummary) SetHintCount(v int32) {
	o.HintCount = &v
}

// GetInfoCount returns the InfoCount field value if set, zero value otherwise.
func (o *ManagementV1APISpecLintingJobResultSummary) GetInfoCount() int32 {
	if o == nil || IsNil(o.InfoCount) {
		var ret int32
		return ret
	}
	return *o.InfoCount
}

// GetInfoCountOk returns a tuple with the InfoCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1APISpecLintingJobResultSummary) GetInfoCountOk() (*int32, bool) {
	if o == nil || IsNil(o.InfoCount) {
		return nil, false
	}
	return o.InfoCount, true
}

// HasInfoCount returns a boolean if a field has been set.
func (o *ManagementV1APISpecLintingJobResultSummary) HasInfoCount() bool {
	if o != nil && !IsNil(o.InfoCount) {
		return true
	}

	return false
}

// SetInfoCount gets a reference to the given int32 and assigns it to the InfoCount field.
func (o *ManagementV1APISpecLintingJobResultSummary) SetInfoCount(v int32) {
	o.InfoCount = &v
}

// GetGrade returns the Grade field value if set, zero value otherwise.
func (o *ManagementV1APISpecLintingJobResultSummary) GetGrade() string {
	if o == nil || IsNil(o.Grade) {
		var ret string
		return ret
	}
	return *o.Grade
}

// GetGradeOk returns a tuple with the Grade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1APISpecLintingJobResultSummary) GetGradeOk() (*string, bool) {
	if o == nil || IsNil(o.Grade) {
		return nil, false
	}
	return o.Grade, true
}

// HasGrade returns a boolean if a field has been set.
func (o *ManagementV1APISpecLintingJobResultSummary) HasGrade() bool {
	if o != nil && !IsNil(o.Grade) {
		return true
	}

	return false
}

// SetGrade gets a reference to the given string and assigns it to the Grade field.
func (o *ManagementV1APISpecLintingJobResultSummary) SetGrade(v string) {
	o.Grade = &v
}

func (o ManagementV1APISpecLintingJobResultSummary) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManagementV1APISpecLintingJobResultSummary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ErrorCount) {
		toSerialize["errorCount"] = o.ErrorCount
	}
	if !IsNil(o.WarningCount) {
		toSerialize["warningCount"] = o.WarningCount
	}
	if !IsNil(o.HintCount) {
		toSerialize["hintCount"] = o.HintCount
	}
	if !IsNil(o.InfoCount) {
		toSerialize["infoCount"] = o.InfoCount
	}
	if !IsNil(o.Grade) {
		toSerialize["grade"] = o.Grade
	}
	return toSerialize, nil
}

type NullableManagementV1APISpecLintingJobResultSummary struct {
	value *ManagementV1APISpecLintingJobResultSummary
	isSet bool
}

func (v NullableManagementV1APISpecLintingJobResultSummary) Get() *ManagementV1APISpecLintingJobResultSummary {
	return v.value
}

func (v *NullableManagementV1APISpecLintingJobResultSummary) Set(val *ManagementV1APISpecLintingJobResultSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableManagementV1APISpecLintingJobResultSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableManagementV1APISpecLintingJobResultSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagementV1APISpecLintingJobResultSummary(val *ManagementV1APISpecLintingJobResultSummary) *NullableManagementV1APISpecLintingJobResultSummary {
	return &NullableManagementV1APISpecLintingJobResultSummary{value: val, isSet: true}
}

func (v NullableManagementV1APISpecLintingJobResultSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagementV1APISpecLintingJobResultSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

