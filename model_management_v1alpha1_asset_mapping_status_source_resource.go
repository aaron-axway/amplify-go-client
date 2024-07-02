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

// checks if the ManagementV1alpha1AssetMappingStatusSourceResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManagementV1alpha1AssetMappingStatusSourceResource{}

// ManagementV1alpha1AssetMappingStatusSourceResource The resource that triggered the Asset Mapping.
type ManagementV1alpha1AssetMappingStatusSourceResource struct {
	ApiService *ManagementV1alpha1AssetMappingStatusSourceResourceApiService `json:"apiService,omitempty"`
	ApiServiceRevision *ManagementV1alpha1AssetMappingStatusSourceResourceApiServiceRevision `json:"apiServiceRevision,omitempty"`
	ApiServiceInstance *ManagementV1alpha1AssetMappingStatusSourceResourceApiServiceInstance `json:"apiServiceInstance,omitempty"`
	AssetMappingTemplate *ManagementV1alpha1AssetMappingStatusSourceResourceAssetMappingTemplate `json:"assetMappingTemplate,omitempty"`
}

// NewManagementV1alpha1AssetMappingStatusSourceResource instantiates a new ManagementV1alpha1AssetMappingStatusSourceResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagementV1alpha1AssetMappingStatusSourceResource() *ManagementV1alpha1AssetMappingStatusSourceResource {
	this := ManagementV1alpha1AssetMappingStatusSourceResource{}
	return &this
}

// NewManagementV1alpha1AssetMappingStatusSourceResourceWithDefaults instantiates a new ManagementV1alpha1AssetMappingStatusSourceResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagementV1alpha1AssetMappingStatusSourceResourceWithDefaults() *ManagementV1alpha1AssetMappingStatusSourceResource {
	this := ManagementV1alpha1AssetMappingStatusSourceResource{}
	return &this
}

// GetApiService returns the ApiService field value if set, zero value otherwise.
func (o *ManagementV1alpha1AssetMappingStatusSourceResource) GetApiService() ManagementV1alpha1AssetMappingStatusSourceResourceApiService {
	if o == nil || IsNil(o.ApiService) {
		var ret ManagementV1alpha1AssetMappingStatusSourceResourceApiService
		return ret
	}
	return *o.ApiService
}

// GetApiServiceOk returns a tuple with the ApiService field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1AssetMappingStatusSourceResource) GetApiServiceOk() (*ManagementV1alpha1AssetMappingStatusSourceResourceApiService, bool) {
	if o == nil || IsNil(o.ApiService) {
		return nil, false
	}
	return o.ApiService, true
}

// HasApiService returns a boolean if a field has been set.
func (o *ManagementV1alpha1AssetMappingStatusSourceResource) HasApiService() bool {
	if o != nil && !IsNil(o.ApiService) {
		return true
	}

	return false
}

// SetApiService gets a reference to the given ManagementV1alpha1AssetMappingStatusSourceResourceApiService and assigns it to the ApiService field.
func (o *ManagementV1alpha1AssetMappingStatusSourceResource) SetApiService(v ManagementV1alpha1AssetMappingStatusSourceResourceApiService) {
	o.ApiService = &v
}

// GetApiServiceRevision returns the ApiServiceRevision field value if set, zero value otherwise.
func (o *ManagementV1alpha1AssetMappingStatusSourceResource) GetApiServiceRevision() ManagementV1alpha1AssetMappingStatusSourceResourceApiServiceRevision {
	if o == nil || IsNil(o.ApiServiceRevision) {
		var ret ManagementV1alpha1AssetMappingStatusSourceResourceApiServiceRevision
		return ret
	}
	return *o.ApiServiceRevision
}

// GetApiServiceRevisionOk returns a tuple with the ApiServiceRevision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1AssetMappingStatusSourceResource) GetApiServiceRevisionOk() (*ManagementV1alpha1AssetMappingStatusSourceResourceApiServiceRevision, bool) {
	if o == nil || IsNil(o.ApiServiceRevision) {
		return nil, false
	}
	return o.ApiServiceRevision, true
}

// HasApiServiceRevision returns a boolean if a field has been set.
func (o *ManagementV1alpha1AssetMappingStatusSourceResource) HasApiServiceRevision() bool {
	if o != nil && !IsNil(o.ApiServiceRevision) {
		return true
	}

	return false
}

// SetApiServiceRevision gets a reference to the given ManagementV1alpha1AssetMappingStatusSourceResourceApiServiceRevision and assigns it to the ApiServiceRevision field.
func (o *ManagementV1alpha1AssetMappingStatusSourceResource) SetApiServiceRevision(v ManagementV1alpha1AssetMappingStatusSourceResourceApiServiceRevision) {
	o.ApiServiceRevision = &v
}

// GetApiServiceInstance returns the ApiServiceInstance field value if set, zero value otherwise.
func (o *ManagementV1alpha1AssetMappingStatusSourceResource) GetApiServiceInstance() ManagementV1alpha1AssetMappingStatusSourceResourceApiServiceInstance {
	if o == nil || IsNil(o.ApiServiceInstance) {
		var ret ManagementV1alpha1AssetMappingStatusSourceResourceApiServiceInstance
		return ret
	}
	return *o.ApiServiceInstance
}

// GetApiServiceInstanceOk returns a tuple with the ApiServiceInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1AssetMappingStatusSourceResource) GetApiServiceInstanceOk() (*ManagementV1alpha1AssetMappingStatusSourceResourceApiServiceInstance, bool) {
	if o == nil || IsNil(o.ApiServiceInstance) {
		return nil, false
	}
	return o.ApiServiceInstance, true
}

// HasApiServiceInstance returns a boolean if a field has been set.
func (o *ManagementV1alpha1AssetMappingStatusSourceResource) HasApiServiceInstance() bool {
	if o != nil && !IsNil(o.ApiServiceInstance) {
		return true
	}

	return false
}

// SetApiServiceInstance gets a reference to the given ManagementV1alpha1AssetMappingStatusSourceResourceApiServiceInstance and assigns it to the ApiServiceInstance field.
func (o *ManagementV1alpha1AssetMappingStatusSourceResource) SetApiServiceInstance(v ManagementV1alpha1AssetMappingStatusSourceResourceApiServiceInstance) {
	o.ApiServiceInstance = &v
}

// GetAssetMappingTemplate returns the AssetMappingTemplate field value if set, zero value otherwise.
func (o *ManagementV1alpha1AssetMappingStatusSourceResource) GetAssetMappingTemplate() ManagementV1alpha1AssetMappingStatusSourceResourceAssetMappingTemplate {
	if o == nil || IsNil(o.AssetMappingTemplate) {
		var ret ManagementV1alpha1AssetMappingStatusSourceResourceAssetMappingTemplate
		return ret
	}
	return *o.AssetMappingTemplate
}

// GetAssetMappingTemplateOk returns a tuple with the AssetMappingTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1AssetMappingStatusSourceResource) GetAssetMappingTemplateOk() (*ManagementV1alpha1AssetMappingStatusSourceResourceAssetMappingTemplate, bool) {
	if o == nil || IsNil(o.AssetMappingTemplate) {
		return nil, false
	}
	return o.AssetMappingTemplate, true
}

// HasAssetMappingTemplate returns a boolean if a field has been set.
func (o *ManagementV1alpha1AssetMappingStatusSourceResource) HasAssetMappingTemplate() bool {
	if o != nil && !IsNil(o.AssetMappingTemplate) {
		return true
	}

	return false
}

// SetAssetMappingTemplate gets a reference to the given ManagementV1alpha1AssetMappingStatusSourceResourceAssetMappingTemplate and assigns it to the AssetMappingTemplate field.
func (o *ManagementV1alpha1AssetMappingStatusSourceResource) SetAssetMappingTemplate(v ManagementV1alpha1AssetMappingStatusSourceResourceAssetMappingTemplate) {
	o.AssetMappingTemplate = &v
}

func (o ManagementV1alpha1AssetMappingStatusSourceResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManagementV1alpha1AssetMappingStatusSourceResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApiService) {
		toSerialize["apiService"] = o.ApiService
	}
	if !IsNil(o.ApiServiceRevision) {
		toSerialize["apiServiceRevision"] = o.ApiServiceRevision
	}
	if !IsNil(o.ApiServiceInstance) {
		toSerialize["apiServiceInstance"] = o.ApiServiceInstance
	}
	if !IsNil(o.AssetMappingTemplate) {
		toSerialize["assetMappingTemplate"] = o.AssetMappingTemplate
	}
	return toSerialize, nil
}

type NullableManagementV1alpha1AssetMappingStatusSourceResource struct {
	value *ManagementV1alpha1AssetMappingStatusSourceResource
	isSet bool
}

func (v NullableManagementV1alpha1AssetMappingStatusSourceResource) Get() *ManagementV1alpha1AssetMappingStatusSourceResource {
	return v.value
}

func (v *NullableManagementV1alpha1AssetMappingStatusSourceResource) Set(val *ManagementV1alpha1AssetMappingStatusSourceResource) {
	v.value = val
	v.isSet = true
}

func (v NullableManagementV1alpha1AssetMappingStatusSourceResource) IsSet() bool {
	return v.isSet
}

func (v *NullableManagementV1alpha1AssetMappingStatusSourceResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagementV1alpha1AssetMappingStatusSourceResource(val *ManagementV1alpha1AssetMappingStatusSourceResource) *NullableManagementV1alpha1AssetMappingStatusSourceResource {
	return &NullableManagementV1alpha1AssetMappingStatusSourceResource{value: val, isSet: true}
}

func (v NullableManagementV1alpha1AssetMappingStatusSourceResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagementV1alpha1AssetMappingStatusSourceResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


