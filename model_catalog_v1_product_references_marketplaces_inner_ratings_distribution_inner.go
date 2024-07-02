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

// checks if the CatalogV1ProductReferencesMarketplacesInnerRatingsDistributionInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogV1ProductReferencesMarketplacesInnerRatingsDistributionInner{}

// CatalogV1ProductReferencesMarketplacesInnerRatingsDistributionInner struct for CatalogV1ProductReferencesMarketplacesInnerRatingsDistributionInner
type CatalogV1ProductReferencesMarketplacesInnerRatingsDistributionInner struct {
	Rating *int32 `json:"rating,omitempty"`
	Total *int32 `json:"total,omitempty"`
}

// NewCatalogV1ProductReferencesMarketplacesInnerRatingsDistributionInner instantiates a new CatalogV1ProductReferencesMarketplacesInnerRatingsDistributionInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogV1ProductReferencesMarketplacesInnerRatingsDistributionInner() *CatalogV1ProductReferencesMarketplacesInnerRatingsDistributionInner {
	this := CatalogV1ProductReferencesMarketplacesInnerRatingsDistributionInner{}
	return &this
}

// NewCatalogV1ProductReferencesMarketplacesInnerRatingsDistributionInnerWithDefaults instantiates a new CatalogV1ProductReferencesMarketplacesInnerRatingsDistributionInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogV1ProductReferencesMarketplacesInnerRatingsDistributionInnerWithDefaults() *CatalogV1ProductReferencesMarketplacesInnerRatingsDistributionInner {
	this := CatalogV1ProductReferencesMarketplacesInnerRatingsDistributionInner{}
	return &this
}

// GetRating returns the Rating field value if set, zero value otherwise.
func (o *CatalogV1ProductReferencesMarketplacesInnerRatingsDistributionInner) GetRating() int32 {
	if o == nil || IsNil(o.Rating) {
		var ret int32
		return ret
	}
	return *o.Rating
}

// GetRatingOk returns a tuple with the Rating field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1ProductReferencesMarketplacesInnerRatingsDistributionInner) GetRatingOk() (*int32, bool) {
	if o == nil || IsNil(o.Rating) {
		return nil, false
	}
	return o.Rating, true
}

// HasRating returns a boolean if a field has been set.
func (o *CatalogV1ProductReferencesMarketplacesInnerRatingsDistributionInner) HasRating() bool {
	if o != nil && !IsNil(o.Rating) {
		return true
	}

	return false
}

// SetRating gets a reference to the given int32 and assigns it to the Rating field.
func (o *CatalogV1ProductReferencesMarketplacesInnerRatingsDistributionInner) SetRating(v int32) {
	o.Rating = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *CatalogV1ProductReferencesMarketplacesInnerRatingsDistributionInner) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1ProductReferencesMarketplacesInnerRatingsDistributionInner) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *CatalogV1ProductReferencesMarketplacesInnerRatingsDistributionInner) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *CatalogV1ProductReferencesMarketplacesInnerRatingsDistributionInner) SetTotal(v int32) {
	o.Total = &v
}

func (o CatalogV1ProductReferencesMarketplacesInnerRatingsDistributionInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogV1ProductReferencesMarketplacesInnerRatingsDistributionInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Rating) {
		toSerialize["rating"] = o.Rating
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	return toSerialize, nil
}

type NullableCatalogV1ProductReferencesMarketplacesInnerRatingsDistributionInner struct {
	value *CatalogV1ProductReferencesMarketplacesInnerRatingsDistributionInner
	isSet bool
}

func (v NullableCatalogV1ProductReferencesMarketplacesInnerRatingsDistributionInner) Get() *CatalogV1ProductReferencesMarketplacesInnerRatingsDistributionInner {
	return v.value
}

func (v *NullableCatalogV1ProductReferencesMarketplacesInnerRatingsDistributionInner) Set(val *CatalogV1ProductReferencesMarketplacesInnerRatingsDistributionInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogV1ProductReferencesMarketplacesInnerRatingsDistributionInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogV1ProductReferencesMarketplacesInnerRatingsDistributionInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogV1ProductReferencesMarketplacesInnerRatingsDistributionInner(val *CatalogV1ProductReferencesMarketplacesInnerRatingsDistributionInner) *NullableCatalogV1ProductReferencesMarketplacesInnerRatingsDistributionInner {
	return &NullableCatalogV1ProductReferencesMarketplacesInnerRatingsDistributionInner{value: val, isSet: true}
}

func (v NullableCatalogV1ProductReferencesMarketplacesInnerRatingsDistributionInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogV1ProductReferencesMarketplacesInnerRatingsDistributionInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

