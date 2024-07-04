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
	"bytes"
	"fmt"
)

// checks if the CatalogV1alpha1AssetMappingSpecInputs type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogV1alpha1AssetMappingSpecInputs{}

// CatalogV1alpha1AssetMappingSpecInputs The list of the inputs provided to the template.
type CatalogV1alpha1AssetMappingSpecInputs struct {
	ApiService string `json:"apiService"`
	ApiServiceRevision *string `json:"apiServiceRevision,omitempty"`
	ApiServiceInstance *string `json:"apiServiceInstance,omitempty"`
	// This property is deprecated and will be ignored.
	// Deprecated
	Stage *string `json:"stage,omitempty"`
	// list of categories for the asset.
	Categories []string `json:"categories,omitempty"`
	// title for generated asset resource.
	AssetResourceTitle *string `json:"assetResourceTitle,omitempty"`
}

type _CatalogV1alpha1AssetMappingSpecInputs CatalogV1alpha1AssetMappingSpecInputs

// NewCatalogV1alpha1AssetMappingSpecInputs instantiates a new CatalogV1alpha1AssetMappingSpecInputs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogV1alpha1AssetMappingSpecInputs(apiService string) *CatalogV1alpha1AssetMappingSpecInputs {
	this := CatalogV1alpha1AssetMappingSpecInputs{}
	this.ApiService = apiService
	return &this
}

// NewCatalogV1alpha1AssetMappingSpecInputsWithDefaults instantiates a new CatalogV1alpha1AssetMappingSpecInputs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogV1alpha1AssetMappingSpecInputsWithDefaults() *CatalogV1alpha1AssetMappingSpecInputs {
	this := CatalogV1alpha1AssetMappingSpecInputs{}
	return &this
}

// GetApiService returns the ApiService field value
func (o *CatalogV1alpha1AssetMappingSpecInputs) GetApiService() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiService
}

// GetApiServiceOk returns a tuple with the ApiService field value
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1AssetMappingSpecInputs) GetApiServiceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiService, true
}

// SetApiService sets field value
func (o *CatalogV1alpha1AssetMappingSpecInputs) SetApiService(v string) {
	o.ApiService = v
}

// GetApiServiceRevision returns the ApiServiceRevision field value if set, zero value otherwise.
func (o *CatalogV1alpha1AssetMappingSpecInputs) GetApiServiceRevision() string {
	if o == nil || IsNil(o.ApiServiceRevision) {
		var ret string
		return ret
	}
	return *o.ApiServiceRevision
}

// GetApiServiceRevisionOk returns a tuple with the ApiServiceRevision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1AssetMappingSpecInputs) GetApiServiceRevisionOk() (*string, bool) {
	if o == nil || IsNil(o.ApiServiceRevision) {
		return nil, false
	}
	return o.ApiServiceRevision, true
}

// HasApiServiceRevision returns a boolean if a field has been set.
func (o *CatalogV1alpha1AssetMappingSpecInputs) HasApiServiceRevision() bool {
	if o != nil && !IsNil(o.ApiServiceRevision) {
		return true
	}

	return false
}

// SetApiServiceRevision gets a reference to the given string and assigns it to the ApiServiceRevision field.
func (o *CatalogV1alpha1AssetMappingSpecInputs) SetApiServiceRevision(v string) {
	o.ApiServiceRevision = &v
}

// GetApiServiceInstance returns the ApiServiceInstance field value if set, zero value otherwise.
func (o *CatalogV1alpha1AssetMappingSpecInputs) GetApiServiceInstance() string {
	if o == nil || IsNil(o.ApiServiceInstance) {
		var ret string
		return ret
	}
	return *o.ApiServiceInstance
}

// GetApiServiceInstanceOk returns a tuple with the ApiServiceInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1AssetMappingSpecInputs) GetApiServiceInstanceOk() (*string, bool) {
	if o == nil || IsNil(o.ApiServiceInstance) {
		return nil, false
	}
	return o.ApiServiceInstance, true
}

// HasApiServiceInstance returns a boolean if a field has been set.
func (o *CatalogV1alpha1AssetMappingSpecInputs) HasApiServiceInstance() bool {
	if o != nil && !IsNil(o.ApiServiceInstance) {
		return true
	}

	return false
}

// SetApiServiceInstance gets a reference to the given string and assigns it to the ApiServiceInstance field.
func (o *CatalogV1alpha1AssetMappingSpecInputs) SetApiServiceInstance(v string) {
	o.ApiServiceInstance = &v
}

// GetStage returns the Stage field value if set, zero value otherwise.
// Deprecated
func (o *CatalogV1alpha1AssetMappingSpecInputs) GetStage() string {
	if o == nil || IsNil(o.Stage) {
		var ret string
		return ret
	}
	return *o.Stage
}

// GetStageOk returns a tuple with the Stage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *CatalogV1alpha1AssetMappingSpecInputs) GetStageOk() (*string, bool) {
	if o == nil || IsNil(o.Stage) {
		return nil, false
	}
	return o.Stage, true
}

// HasStage returns a boolean if a field has been set.
func (o *CatalogV1alpha1AssetMappingSpecInputs) HasStage() bool {
	if o != nil && !IsNil(o.Stage) {
		return true
	}

	return false
}

// SetStage gets a reference to the given string and assigns it to the Stage field.
// Deprecated
func (o *CatalogV1alpha1AssetMappingSpecInputs) SetStage(v string) {
	o.Stage = &v
}

// GetCategories returns the Categories field value if set, zero value otherwise.
func (o *CatalogV1alpha1AssetMappingSpecInputs) GetCategories() []string {
	if o == nil || IsNil(o.Categories) {
		var ret []string
		return ret
	}
	return o.Categories
}

// GetCategoriesOk returns a tuple with the Categories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1AssetMappingSpecInputs) GetCategoriesOk() ([]string, bool) {
	if o == nil || IsNil(o.Categories) {
		return nil, false
	}
	return o.Categories, true
}

// HasCategories returns a boolean if a field has been set.
func (o *CatalogV1alpha1AssetMappingSpecInputs) HasCategories() bool {
	if o != nil && !IsNil(o.Categories) {
		return true
	}

	return false
}

// SetCategories gets a reference to the given []string and assigns it to the Categories field.
func (o *CatalogV1alpha1AssetMappingSpecInputs) SetCategories(v []string) {
	o.Categories = v
}

// GetAssetResourceTitle returns the AssetResourceTitle field value if set, zero value otherwise.
func (o *CatalogV1alpha1AssetMappingSpecInputs) GetAssetResourceTitle() string {
	if o == nil || IsNil(o.AssetResourceTitle) {
		var ret string
		return ret
	}
	return *o.AssetResourceTitle
}

// GetAssetResourceTitleOk returns a tuple with the AssetResourceTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1AssetMappingSpecInputs) GetAssetResourceTitleOk() (*string, bool) {
	if o == nil || IsNil(o.AssetResourceTitle) {
		return nil, false
	}
	return o.AssetResourceTitle, true
}

// HasAssetResourceTitle returns a boolean if a field has been set.
func (o *CatalogV1alpha1AssetMappingSpecInputs) HasAssetResourceTitle() bool {
	if o != nil && !IsNil(o.AssetResourceTitle) {
		return true
	}

	return false
}

// SetAssetResourceTitle gets a reference to the given string and assigns it to the AssetResourceTitle field.
func (o *CatalogV1alpha1AssetMappingSpecInputs) SetAssetResourceTitle(v string) {
	o.AssetResourceTitle = &v
}

func (o CatalogV1alpha1AssetMappingSpecInputs) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogV1alpha1AssetMappingSpecInputs) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["apiService"] = o.ApiService
	if !IsNil(o.ApiServiceRevision) {
		toSerialize["apiServiceRevision"] = o.ApiServiceRevision
	}
	if !IsNil(o.ApiServiceInstance) {
		toSerialize["apiServiceInstance"] = o.ApiServiceInstance
	}
	if !IsNil(o.Stage) {
		toSerialize["stage"] = o.Stage
	}
	if !IsNil(o.Categories) {
		toSerialize["categories"] = o.Categories
	}
	if !IsNil(o.AssetResourceTitle) {
		toSerialize["assetResourceTitle"] = o.AssetResourceTitle
	}
	return toSerialize, nil
}

func (o *CatalogV1alpha1AssetMappingSpecInputs) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"apiService",
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

	varCatalogV1alpha1AssetMappingSpecInputs := _CatalogV1alpha1AssetMappingSpecInputs{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCatalogV1alpha1AssetMappingSpecInputs)

	if err != nil {
		return err
	}

	*o = CatalogV1alpha1AssetMappingSpecInputs(varCatalogV1alpha1AssetMappingSpecInputs)

	return err
}

type NullableCatalogV1alpha1AssetMappingSpecInputs struct {
	value *CatalogV1alpha1AssetMappingSpecInputs
	isSet bool
}

func (v NullableCatalogV1alpha1AssetMappingSpecInputs) Get() *CatalogV1alpha1AssetMappingSpecInputs {
	return v.value
}

func (v *NullableCatalogV1alpha1AssetMappingSpecInputs) Set(val *CatalogV1alpha1AssetMappingSpecInputs) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogV1alpha1AssetMappingSpecInputs) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogV1alpha1AssetMappingSpecInputs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogV1alpha1AssetMappingSpecInputs(val *CatalogV1alpha1AssetMappingSpecInputs) *NullableCatalogV1alpha1AssetMappingSpecInputs {
	return &NullableCatalogV1alpha1AssetMappingSpecInputs{value: val, isSet: true}
}

func (v NullableCatalogV1alpha1AssetMappingSpecInputs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogV1alpha1AssetMappingSpecInputs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


