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

// checks if the CatalogV1SupportContactSpecOfficeHours type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogV1SupportContactSpecOfficeHours{}

// CatalogV1SupportContactSpecOfficeHours struct for CatalogV1SupportContactSpecOfficeHours
type CatalogV1SupportContactSpecOfficeHours struct {
	// Long IANA Time Zone format. Examples: 'America/New_York' or 'Europe/Paris'
	Timezone string `json:"timezone" validate:"regexp=^[\\\\-A-Za-z_+0-9]{1,14}(\\/[\\\\-A-Za-z_+0-9]{1,14}){0,2}$"`
	Periods []CatalogV1SupportContactSpecOfficeHoursPeriodsInner `json:"periods"`
}

type _CatalogV1SupportContactSpecOfficeHours CatalogV1SupportContactSpecOfficeHours

// NewCatalogV1SupportContactSpecOfficeHours instantiates a new CatalogV1SupportContactSpecOfficeHours object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogV1SupportContactSpecOfficeHours(timezone string, periods []CatalogV1SupportContactSpecOfficeHoursPeriodsInner) *CatalogV1SupportContactSpecOfficeHours {
	this := CatalogV1SupportContactSpecOfficeHours{}
	this.Timezone = timezone
	this.Periods = periods
	return &this
}

// NewCatalogV1SupportContactSpecOfficeHoursWithDefaults instantiates a new CatalogV1SupportContactSpecOfficeHours object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogV1SupportContactSpecOfficeHoursWithDefaults() *CatalogV1SupportContactSpecOfficeHours {
	this := CatalogV1SupportContactSpecOfficeHours{}
	return &this
}

// GetTimezone returns the Timezone field value
func (o *CatalogV1SupportContactSpecOfficeHours) GetTimezone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value
// and a boolean to check if the value has been set.
func (o *CatalogV1SupportContactSpecOfficeHours) GetTimezoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timezone, true
}

// SetTimezone sets field value
func (o *CatalogV1SupportContactSpecOfficeHours) SetTimezone(v string) {
	o.Timezone = v
}

// GetPeriods returns the Periods field value
func (o *CatalogV1SupportContactSpecOfficeHours) GetPeriods() []CatalogV1SupportContactSpecOfficeHoursPeriodsInner {
	if o == nil {
		var ret []CatalogV1SupportContactSpecOfficeHoursPeriodsInner
		return ret
	}

	return o.Periods
}

// GetPeriodsOk returns a tuple with the Periods field value
// and a boolean to check if the value has been set.
func (o *CatalogV1SupportContactSpecOfficeHours) GetPeriodsOk() ([]CatalogV1SupportContactSpecOfficeHoursPeriodsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Periods, true
}

// SetPeriods sets field value
func (o *CatalogV1SupportContactSpecOfficeHours) SetPeriods(v []CatalogV1SupportContactSpecOfficeHoursPeriodsInner) {
	o.Periods = v
}

func (o CatalogV1SupportContactSpecOfficeHours) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogV1SupportContactSpecOfficeHours) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["timezone"] = o.Timezone
	toSerialize["periods"] = o.Periods
	return toSerialize, nil
}

func (o *CatalogV1SupportContactSpecOfficeHours) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"timezone",
		"periods",
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

	varCatalogV1SupportContactSpecOfficeHours := _CatalogV1SupportContactSpecOfficeHours{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCatalogV1SupportContactSpecOfficeHours)

	if err != nil {
		return err
	}

	*o = CatalogV1SupportContactSpecOfficeHours(varCatalogV1SupportContactSpecOfficeHours)

	return err
}

type NullableCatalogV1SupportContactSpecOfficeHours struct {
	value *CatalogV1SupportContactSpecOfficeHours
	isSet bool
}

func (v NullableCatalogV1SupportContactSpecOfficeHours) Get() *CatalogV1SupportContactSpecOfficeHours {
	return v.value
}

func (v *NullableCatalogV1SupportContactSpecOfficeHours) Set(val *CatalogV1SupportContactSpecOfficeHours) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogV1SupportContactSpecOfficeHours) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogV1SupportContactSpecOfficeHours) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogV1SupportContactSpecOfficeHours(val *CatalogV1SupportContactSpecOfficeHours) *NullableCatalogV1SupportContactSpecOfficeHours {
	return &NullableCatalogV1SupportContactSpecOfficeHours{value: val, isSet: true}
}

func (v NullableCatalogV1SupportContactSpecOfficeHours) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogV1SupportContactSpecOfficeHours) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


