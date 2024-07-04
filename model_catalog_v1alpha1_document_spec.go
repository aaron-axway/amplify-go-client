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

// checks if the CatalogV1alpha1DocumentSpec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogV1alpha1DocumentSpec{}

// CatalogV1alpha1DocumentSpec struct for CatalogV1alpha1DocumentSpec
type CatalogV1alpha1DocumentSpec struct {
	// Document description.
	Description *string `json:"description,omitempty"`
	// Rank of document.
	Rank *float32 `json:"rank,omitempty"`
	Sections []CatalogV1alpha1DocumentSpecSectionsInner `json:"sections,omitempty"`
}

// NewCatalogV1alpha1DocumentSpec instantiates a new CatalogV1alpha1DocumentSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogV1alpha1DocumentSpec() *CatalogV1alpha1DocumentSpec {
	this := CatalogV1alpha1DocumentSpec{}
	var rank float32 = 0
	this.Rank = &rank
	return &this
}

// NewCatalogV1alpha1DocumentSpecWithDefaults instantiates a new CatalogV1alpha1DocumentSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogV1alpha1DocumentSpecWithDefaults() *CatalogV1alpha1DocumentSpec {
	this := CatalogV1alpha1DocumentSpec{}
	var rank float32 = 0
	this.Rank = &rank
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CatalogV1alpha1DocumentSpec) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1DocumentSpec) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CatalogV1alpha1DocumentSpec) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CatalogV1alpha1DocumentSpec) SetDescription(v string) {
	o.Description = &v
}

// GetRank returns the Rank field value if set, zero value otherwise.
func (o *CatalogV1alpha1DocumentSpec) GetRank() float32 {
	if o == nil || IsNil(o.Rank) {
		var ret float32
		return ret
	}
	return *o.Rank
}

// GetRankOk returns a tuple with the Rank field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1DocumentSpec) GetRankOk() (*float32, bool) {
	if o == nil || IsNil(o.Rank) {
		return nil, false
	}
	return o.Rank, true
}

// HasRank returns a boolean if a field has been set.
func (o *CatalogV1alpha1DocumentSpec) HasRank() bool {
	if o != nil && !IsNil(o.Rank) {
		return true
	}

	return false
}

// SetRank gets a reference to the given float32 and assigns it to the Rank field.
func (o *CatalogV1alpha1DocumentSpec) SetRank(v float32) {
	o.Rank = &v
}

// GetSections returns the Sections field value if set, zero value otherwise.
func (o *CatalogV1alpha1DocumentSpec) GetSections() []CatalogV1alpha1DocumentSpecSectionsInner {
	if o == nil || IsNil(o.Sections) {
		var ret []CatalogV1alpha1DocumentSpecSectionsInner
		return ret
	}
	return o.Sections
}

// GetSectionsOk returns a tuple with the Sections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1DocumentSpec) GetSectionsOk() ([]CatalogV1alpha1DocumentSpecSectionsInner, bool) {
	if o == nil || IsNil(o.Sections) {
		return nil, false
	}
	return o.Sections, true
}

// HasSections returns a boolean if a field has been set.
func (o *CatalogV1alpha1DocumentSpec) HasSections() bool {
	if o != nil && !IsNil(o.Sections) {
		return true
	}

	return false
}

// SetSections gets a reference to the given []CatalogV1alpha1DocumentSpecSectionsInner and assigns it to the Sections field.
func (o *CatalogV1alpha1DocumentSpec) SetSections(v []CatalogV1alpha1DocumentSpecSectionsInner) {
	o.Sections = v
}

func (o CatalogV1alpha1DocumentSpec) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogV1alpha1DocumentSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Rank) {
		toSerialize["rank"] = o.Rank
	}
	if !IsNil(o.Sections) {
		toSerialize["sections"] = o.Sections
	}
	return toSerialize, nil
}

type NullableCatalogV1alpha1DocumentSpec struct {
	value *CatalogV1alpha1DocumentSpec
	isSet bool
}

func (v NullableCatalogV1alpha1DocumentSpec) Get() *CatalogV1alpha1DocumentSpec {
	return v.value
}

func (v *NullableCatalogV1alpha1DocumentSpec) Set(val *CatalogV1alpha1DocumentSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogV1alpha1DocumentSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogV1alpha1DocumentSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogV1alpha1DocumentSpec(val *CatalogV1alpha1DocumentSpec) *NullableCatalogV1alpha1DocumentSpec {
	return &NullableCatalogV1alpha1DocumentSpec{value: val, isSet: true}
}

func (v NullableCatalogV1alpha1DocumentSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogV1alpha1DocumentSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


