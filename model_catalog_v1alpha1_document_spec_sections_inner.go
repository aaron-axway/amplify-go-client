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

// checks if the CatalogV1alpha1DocumentSpecSectionsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogV1alpha1DocumentSpecSectionsInner{}

// CatalogV1alpha1DocumentSpecSectionsInner struct for CatalogV1alpha1DocumentSpecSectionsInner
type CatalogV1alpha1DocumentSpecSectionsInner struct {
	// Title for the section.
	Title string `json:"title"`
	// Description for the section.
	Description *string `json:"description,omitempty"`
	// Section articles.
	Articles []CatalogV1alpha1DocumentSpecSectionsInnerArticlesInner `json:"articles,omitempty"`
}

type _CatalogV1alpha1DocumentSpecSectionsInner CatalogV1alpha1DocumentSpecSectionsInner

// NewCatalogV1alpha1DocumentSpecSectionsInner instantiates a new CatalogV1alpha1DocumentSpecSectionsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogV1alpha1DocumentSpecSectionsInner(title string) *CatalogV1alpha1DocumentSpecSectionsInner {
	this := CatalogV1alpha1DocumentSpecSectionsInner{}
	this.Title = title
	return &this
}

// NewCatalogV1alpha1DocumentSpecSectionsInnerWithDefaults instantiates a new CatalogV1alpha1DocumentSpecSectionsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogV1alpha1DocumentSpecSectionsInnerWithDefaults() *CatalogV1alpha1DocumentSpecSectionsInner {
	this := CatalogV1alpha1DocumentSpecSectionsInner{}
	return &this
}

// GetTitle returns the Title field value
func (o *CatalogV1alpha1DocumentSpecSectionsInner) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1DocumentSpecSectionsInner) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *CatalogV1alpha1DocumentSpecSectionsInner) SetTitle(v string) {
	o.Title = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CatalogV1alpha1DocumentSpecSectionsInner) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1DocumentSpecSectionsInner) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CatalogV1alpha1DocumentSpecSectionsInner) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CatalogV1alpha1DocumentSpecSectionsInner) SetDescription(v string) {
	o.Description = &v
}

// GetArticles returns the Articles field value if set, zero value otherwise.
func (o *CatalogV1alpha1DocumentSpecSectionsInner) GetArticles() []CatalogV1alpha1DocumentSpecSectionsInnerArticlesInner {
	if o == nil || IsNil(o.Articles) {
		var ret []CatalogV1alpha1DocumentSpecSectionsInnerArticlesInner
		return ret
	}
	return o.Articles
}

// GetArticlesOk returns a tuple with the Articles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1DocumentSpecSectionsInner) GetArticlesOk() ([]CatalogV1alpha1DocumentSpecSectionsInnerArticlesInner, bool) {
	if o == nil || IsNil(o.Articles) {
		return nil, false
	}
	return o.Articles, true
}

// HasArticles returns a boolean if a field has been set.
func (o *CatalogV1alpha1DocumentSpecSectionsInner) HasArticles() bool {
	if o != nil && !IsNil(o.Articles) {
		return true
	}

	return false
}

// SetArticles gets a reference to the given []CatalogV1alpha1DocumentSpecSectionsInnerArticlesInner and assigns it to the Articles field.
func (o *CatalogV1alpha1DocumentSpecSectionsInner) SetArticles(v []CatalogV1alpha1DocumentSpecSectionsInnerArticlesInner) {
	o.Articles = v
}

func (o CatalogV1alpha1DocumentSpecSectionsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogV1alpha1DocumentSpecSectionsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["title"] = o.Title
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Articles) {
		toSerialize["articles"] = o.Articles
	}
	return toSerialize, nil
}

func (o *CatalogV1alpha1DocumentSpecSectionsInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"title",
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

	varCatalogV1alpha1DocumentSpecSectionsInner := _CatalogV1alpha1DocumentSpecSectionsInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCatalogV1alpha1DocumentSpecSectionsInner)

	if err != nil {
		return err
	}

	*o = CatalogV1alpha1DocumentSpecSectionsInner(varCatalogV1alpha1DocumentSpecSectionsInner)

	return err
}

type NullableCatalogV1alpha1DocumentSpecSectionsInner struct {
	value *CatalogV1alpha1DocumentSpecSectionsInner
	isSet bool
}

func (v NullableCatalogV1alpha1DocumentSpecSectionsInner) Get() *CatalogV1alpha1DocumentSpecSectionsInner {
	return v.value
}

func (v *NullableCatalogV1alpha1DocumentSpecSectionsInner) Set(val *CatalogV1alpha1DocumentSpecSectionsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogV1alpha1DocumentSpecSectionsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogV1alpha1DocumentSpecSectionsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogV1alpha1DocumentSpecSectionsInner(val *CatalogV1alpha1DocumentSpecSectionsInner) *NullableCatalogV1alpha1DocumentSpecSectionsInner {
	return &NullableCatalogV1alpha1DocumentSpecSectionsInner{value: val, isSet: true}
}

func (v NullableCatalogV1alpha1DocumentSpecSectionsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogV1alpha1DocumentSpecSectionsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

