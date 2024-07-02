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

// checks if the CatalogV1alpha1ProductReviewSpec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogV1alpha1ProductReviewSpec{}

// CatalogV1alpha1ProductReviewSpec struct for CatalogV1alpha1ProductReviewSpec
type CatalogV1alpha1ProductReviewSpec struct {
	Rating int32 `json:"rating"`
	Comment *string `json:"comment,omitempty"`
}

type _CatalogV1alpha1ProductReviewSpec CatalogV1alpha1ProductReviewSpec

// NewCatalogV1alpha1ProductReviewSpec instantiates a new CatalogV1alpha1ProductReviewSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogV1alpha1ProductReviewSpec(rating int32) *CatalogV1alpha1ProductReviewSpec {
	this := CatalogV1alpha1ProductReviewSpec{}
	this.Rating = rating
	return &this
}

// NewCatalogV1alpha1ProductReviewSpecWithDefaults instantiates a new CatalogV1alpha1ProductReviewSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogV1alpha1ProductReviewSpecWithDefaults() *CatalogV1alpha1ProductReviewSpec {
	this := CatalogV1alpha1ProductReviewSpec{}
	return &this
}

// GetRating returns the Rating field value
func (o *CatalogV1alpha1ProductReviewSpec) GetRating() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Rating
}

// GetRatingOk returns a tuple with the Rating field value
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1ProductReviewSpec) GetRatingOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rating, true
}

// SetRating sets field value
func (o *CatalogV1alpha1ProductReviewSpec) SetRating(v int32) {
	o.Rating = v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *CatalogV1alpha1ProductReviewSpec) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1ProductReviewSpec) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *CatalogV1alpha1ProductReviewSpec) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *CatalogV1alpha1ProductReviewSpec) SetComment(v string) {
	o.Comment = &v
}

func (o CatalogV1alpha1ProductReviewSpec) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogV1alpha1ProductReviewSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["rating"] = o.Rating
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	return toSerialize, nil
}

func (o *CatalogV1alpha1ProductReviewSpec) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"rating",
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

	varCatalogV1alpha1ProductReviewSpec := _CatalogV1alpha1ProductReviewSpec{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCatalogV1alpha1ProductReviewSpec)

	if err != nil {
		return err
	}

	*o = CatalogV1alpha1ProductReviewSpec(varCatalogV1alpha1ProductReviewSpec)

	return err
}

type NullableCatalogV1alpha1ProductReviewSpec struct {
	value *CatalogV1alpha1ProductReviewSpec
	isSet bool
}

func (v NullableCatalogV1alpha1ProductReviewSpec) Get() *CatalogV1alpha1ProductReviewSpec {
	return v.value
}

func (v *NullableCatalogV1alpha1ProductReviewSpec) Set(val *CatalogV1alpha1ProductReviewSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogV1alpha1ProductReviewSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogV1alpha1ProductReviewSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogV1alpha1ProductReviewSpec(val *CatalogV1alpha1ProductReviewSpec) *NullableCatalogV1alpha1ProductReviewSpec {
	return &NullableCatalogV1alpha1ProductReviewSpec{value: val, isSet: true}
}

func (v NullableCatalogV1alpha1ProductReviewSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogV1alpha1ProductReviewSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

