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
)

// checks if the CatalogV1ProductReferencesMarketplacesInnerRatings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogV1ProductReferencesMarketplacesInnerRatings{}

// CatalogV1ProductReferencesMarketplacesInnerRatings struct for CatalogV1ProductReferencesMarketplacesInnerRatings
type CatalogV1ProductReferencesMarketplacesInnerRatings struct {
	Total *int32 `json:"total,omitempty"`
	Average *float32 `json:"average,omitempty"`
	// The ratings distribution per value.
	Distribution []CatalogV1ProductReferencesMarketplacesInnerRatingsDistributionInner `json:"distribution,omitempty"`
}

// NewCatalogV1ProductReferencesMarketplacesInnerRatings instantiates a new CatalogV1ProductReferencesMarketplacesInnerRatings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogV1ProductReferencesMarketplacesInnerRatings() *CatalogV1ProductReferencesMarketplacesInnerRatings {
	this := CatalogV1ProductReferencesMarketplacesInnerRatings{}
	return &this
}

// NewCatalogV1ProductReferencesMarketplacesInnerRatingsWithDefaults instantiates a new CatalogV1ProductReferencesMarketplacesInnerRatings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogV1ProductReferencesMarketplacesInnerRatingsWithDefaults() *CatalogV1ProductReferencesMarketplacesInnerRatings {
	this := CatalogV1ProductReferencesMarketplacesInnerRatings{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *CatalogV1ProductReferencesMarketplacesInnerRatings) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1ProductReferencesMarketplacesInnerRatings) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *CatalogV1ProductReferencesMarketplacesInnerRatings) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *CatalogV1ProductReferencesMarketplacesInnerRatings) SetTotal(v int32) {
	o.Total = &v
}

// GetAverage returns the Average field value if set, zero value otherwise.
func (o *CatalogV1ProductReferencesMarketplacesInnerRatings) GetAverage() float32 {
	if o == nil || IsNil(o.Average) {
		var ret float32
		return ret
	}
	return *o.Average
}

// GetAverageOk returns a tuple with the Average field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1ProductReferencesMarketplacesInnerRatings) GetAverageOk() (*float32, bool) {
	if o == nil || IsNil(o.Average) {
		return nil, false
	}
	return o.Average, true
}

// HasAverage returns a boolean if a field has been set.
func (o *CatalogV1ProductReferencesMarketplacesInnerRatings) HasAverage() bool {
	if o != nil && !IsNil(o.Average) {
		return true
	}

	return false
}

// SetAverage gets a reference to the given float32 and assigns it to the Average field.
func (o *CatalogV1ProductReferencesMarketplacesInnerRatings) SetAverage(v float32) {
	o.Average = &v
}

// GetDistribution returns the Distribution field value if set, zero value otherwise.
func (o *CatalogV1ProductReferencesMarketplacesInnerRatings) GetDistribution() []CatalogV1ProductReferencesMarketplacesInnerRatingsDistributionInner {
	if o == nil || IsNil(o.Distribution) {
		var ret []CatalogV1ProductReferencesMarketplacesInnerRatingsDistributionInner
		return ret
	}
	return o.Distribution
}

// GetDistributionOk returns a tuple with the Distribution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1ProductReferencesMarketplacesInnerRatings) GetDistributionOk() ([]CatalogV1ProductReferencesMarketplacesInnerRatingsDistributionInner, bool) {
	if o == nil || IsNil(o.Distribution) {
		return nil, false
	}
	return o.Distribution, true
}

// HasDistribution returns a boolean if a field has been set.
func (o *CatalogV1ProductReferencesMarketplacesInnerRatings) HasDistribution() bool {
	if o != nil && !IsNil(o.Distribution) {
		return true
	}

	return false
}

// SetDistribution gets a reference to the given []CatalogV1ProductReferencesMarketplacesInnerRatingsDistributionInner and assigns it to the Distribution field.
func (o *CatalogV1ProductReferencesMarketplacesInnerRatings) SetDistribution(v []CatalogV1ProductReferencesMarketplacesInnerRatingsDistributionInner) {
	o.Distribution = v
}

func (o CatalogV1ProductReferencesMarketplacesInnerRatings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogV1ProductReferencesMarketplacesInnerRatings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !IsNil(o.Average) {
		toSerialize["average"] = o.Average
	}
	if !IsNil(o.Distribution) {
		toSerialize["distribution"] = o.Distribution
	}
	return toSerialize, nil
}

type NullableCatalogV1ProductReferencesMarketplacesInnerRatings struct {
	value *CatalogV1ProductReferencesMarketplacesInnerRatings
	isSet bool
}

func (v NullableCatalogV1ProductReferencesMarketplacesInnerRatings) Get() *CatalogV1ProductReferencesMarketplacesInnerRatings {
	return v.value
}

func (v *NullableCatalogV1ProductReferencesMarketplacesInnerRatings) Set(val *CatalogV1ProductReferencesMarketplacesInnerRatings) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogV1ProductReferencesMarketplacesInnerRatings) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogV1ProductReferencesMarketplacesInnerRatings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogV1ProductReferencesMarketplacesInnerRatings(val *CatalogV1ProductReferencesMarketplacesInnerRatings) *NullableCatalogV1ProductReferencesMarketplacesInnerRatings {
	return &NullableCatalogV1ProductReferencesMarketplacesInnerRatings{value: val, isSet: true}
}

func (v NullableCatalogV1ProductReferencesMarketplacesInnerRatings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogV1ProductReferencesMarketplacesInnerRatings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


