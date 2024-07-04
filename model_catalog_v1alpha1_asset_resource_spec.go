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

// checks if the CatalogV1alpha1AssetResourceSpec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogV1alpha1AssetResourceSpec{}

// CatalogV1alpha1AssetResourceSpec struct for CatalogV1alpha1AssetResourceSpec
type CatalogV1alpha1AssetResourceSpec struct {
	// The Stage this Asset Resource is deployed on.
	Stage *string `json:"stage,omitempty"`
	AssetRequestDefinition *string `json:"assetRequestDefinition,omitempty"`
	CredentialRequestDefinitions []string `json:"credentialRequestDefinitions,omitempty"`
	Type string `json:"type"`
	// content-type of the spec.
	ContentType *string `json:"contentType,omitempty" validate:"regexp=^[-\\\\w.]+\\/[-+\\\\w.]+$"`
	// The version of referenced resource.
	Version *string `json:"version,omitempty"`
	// Base64 encoded value of the api specification.
	Definition string `json:"definition"`
	// Resource availability
	Status string `json:"status"`
	// information to access the definition.
	AccessInfo []CatalogV1alpha1AssetResourceSpecAccessInfoInner `json:"accessInfo,omitempty"`
	SourceReleaseState *CatalogV1alpha1AssetResourceSpecSourceReleaseState `json:"sourceReleaseState,omitempty"`
}

type _CatalogV1alpha1AssetResourceSpec CatalogV1alpha1AssetResourceSpec

// NewCatalogV1alpha1AssetResourceSpec instantiates a new CatalogV1alpha1AssetResourceSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogV1alpha1AssetResourceSpec(type_ string, definition string, status string) *CatalogV1alpha1AssetResourceSpec {
	this := CatalogV1alpha1AssetResourceSpec{}
	this.Type = type_
	this.Definition = definition
	this.Status = status
	return &this
}

// NewCatalogV1alpha1AssetResourceSpecWithDefaults instantiates a new CatalogV1alpha1AssetResourceSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogV1alpha1AssetResourceSpecWithDefaults() *CatalogV1alpha1AssetResourceSpec {
	this := CatalogV1alpha1AssetResourceSpec{}
	return &this
}

// GetStage returns the Stage field value if set, zero value otherwise.
func (o *CatalogV1alpha1AssetResourceSpec) GetStage() string {
	if o == nil || IsNil(o.Stage) {
		var ret string
		return ret
	}
	return *o.Stage
}

// GetStageOk returns a tuple with the Stage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1AssetResourceSpec) GetStageOk() (*string, bool) {
	if o == nil || IsNil(o.Stage) {
		return nil, false
	}
	return o.Stage, true
}

// HasStage returns a boolean if a field has been set.
func (o *CatalogV1alpha1AssetResourceSpec) HasStage() bool {
	if o != nil && !IsNil(o.Stage) {
		return true
	}

	return false
}

// SetStage gets a reference to the given string and assigns it to the Stage field.
func (o *CatalogV1alpha1AssetResourceSpec) SetStage(v string) {
	o.Stage = &v
}

// GetAssetRequestDefinition returns the AssetRequestDefinition field value if set, zero value otherwise.
func (o *CatalogV1alpha1AssetResourceSpec) GetAssetRequestDefinition() string {
	if o == nil || IsNil(o.AssetRequestDefinition) {
		var ret string
		return ret
	}
	return *o.AssetRequestDefinition
}

// GetAssetRequestDefinitionOk returns a tuple with the AssetRequestDefinition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1AssetResourceSpec) GetAssetRequestDefinitionOk() (*string, bool) {
	if o == nil || IsNil(o.AssetRequestDefinition) {
		return nil, false
	}
	return o.AssetRequestDefinition, true
}

// HasAssetRequestDefinition returns a boolean if a field has been set.
func (o *CatalogV1alpha1AssetResourceSpec) HasAssetRequestDefinition() bool {
	if o != nil && !IsNil(o.AssetRequestDefinition) {
		return true
	}

	return false
}

// SetAssetRequestDefinition gets a reference to the given string and assigns it to the AssetRequestDefinition field.
func (o *CatalogV1alpha1AssetResourceSpec) SetAssetRequestDefinition(v string) {
	o.AssetRequestDefinition = &v
}

// GetCredentialRequestDefinitions returns the CredentialRequestDefinitions field value if set, zero value otherwise.
func (o *CatalogV1alpha1AssetResourceSpec) GetCredentialRequestDefinitions() []string {
	if o == nil || IsNil(o.CredentialRequestDefinitions) {
		var ret []string
		return ret
	}
	return o.CredentialRequestDefinitions
}

// GetCredentialRequestDefinitionsOk returns a tuple with the CredentialRequestDefinitions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1AssetResourceSpec) GetCredentialRequestDefinitionsOk() ([]string, bool) {
	if o == nil || IsNil(o.CredentialRequestDefinitions) {
		return nil, false
	}
	return o.CredentialRequestDefinitions, true
}

// HasCredentialRequestDefinitions returns a boolean if a field has been set.
func (o *CatalogV1alpha1AssetResourceSpec) HasCredentialRequestDefinitions() bool {
	if o != nil && !IsNil(o.CredentialRequestDefinitions) {
		return true
	}

	return false
}

// SetCredentialRequestDefinitions gets a reference to the given []string and assigns it to the CredentialRequestDefinitions field.
func (o *CatalogV1alpha1AssetResourceSpec) SetCredentialRequestDefinitions(v []string) {
	o.CredentialRequestDefinitions = v
}

// GetType returns the Type field value
func (o *CatalogV1alpha1AssetResourceSpec) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1AssetResourceSpec) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CatalogV1alpha1AssetResourceSpec) SetType(v string) {
	o.Type = v
}

// GetContentType returns the ContentType field value if set, zero value otherwise.
func (o *CatalogV1alpha1AssetResourceSpec) GetContentType() string {
	if o == nil || IsNil(o.ContentType) {
		var ret string
		return ret
	}
	return *o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1AssetResourceSpec) GetContentTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ContentType) {
		return nil, false
	}
	return o.ContentType, true
}

// HasContentType returns a boolean if a field has been set.
func (o *CatalogV1alpha1AssetResourceSpec) HasContentType() bool {
	if o != nil && !IsNil(o.ContentType) {
		return true
	}

	return false
}

// SetContentType gets a reference to the given string and assigns it to the ContentType field.
func (o *CatalogV1alpha1AssetResourceSpec) SetContentType(v string) {
	o.ContentType = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *CatalogV1alpha1AssetResourceSpec) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1AssetResourceSpec) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *CatalogV1alpha1AssetResourceSpec) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *CatalogV1alpha1AssetResourceSpec) SetVersion(v string) {
	o.Version = &v
}

// GetDefinition returns the Definition field value
func (o *CatalogV1alpha1AssetResourceSpec) GetDefinition() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Definition
}

// GetDefinitionOk returns a tuple with the Definition field value
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1AssetResourceSpec) GetDefinitionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Definition, true
}

// SetDefinition sets field value
func (o *CatalogV1alpha1AssetResourceSpec) SetDefinition(v string) {
	o.Definition = v
}

// GetStatus returns the Status field value
func (o *CatalogV1alpha1AssetResourceSpec) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1AssetResourceSpec) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *CatalogV1alpha1AssetResourceSpec) SetStatus(v string) {
	o.Status = v
}

// GetAccessInfo returns the AccessInfo field value if set, zero value otherwise.
func (o *CatalogV1alpha1AssetResourceSpec) GetAccessInfo() []CatalogV1alpha1AssetResourceSpecAccessInfoInner {
	if o == nil || IsNil(o.AccessInfo) {
		var ret []CatalogV1alpha1AssetResourceSpecAccessInfoInner
		return ret
	}
	return o.AccessInfo
}

// GetAccessInfoOk returns a tuple with the AccessInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1AssetResourceSpec) GetAccessInfoOk() ([]CatalogV1alpha1AssetResourceSpecAccessInfoInner, bool) {
	if o == nil || IsNil(o.AccessInfo) {
		return nil, false
	}
	return o.AccessInfo, true
}

// HasAccessInfo returns a boolean if a field has been set.
func (o *CatalogV1alpha1AssetResourceSpec) HasAccessInfo() bool {
	if o != nil && !IsNil(o.AccessInfo) {
		return true
	}

	return false
}

// SetAccessInfo gets a reference to the given []CatalogV1alpha1AssetResourceSpecAccessInfoInner and assigns it to the AccessInfo field.
func (o *CatalogV1alpha1AssetResourceSpec) SetAccessInfo(v []CatalogV1alpha1AssetResourceSpecAccessInfoInner) {
	o.AccessInfo = v
}

// GetSourceReleaseState returns the SourceReleaseState field value if set, zero value otherwise.
func (o *CatalogV1alpha1AssetResourceSpec) GetSourceReleaseState() CatalogV1alpha1AssetResourceSpecSourceReleaseState {
	if o == nil || IsNil(o.SourceReleaseState) {
		var ret CatalogV1alpha1AssetResourceSpecSourceReleaseState
		return ret
	}
	return *o.SourceReleaseState
}

// GetSourceReleaseStateOk returns a tuple with the SourceReleaseState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1AssetResourceSpec) GetSourceReleaseStateOk() (*CatalogV1alpha1AssetResourceSpecSourceReleaseState, bool) {
	if o == nil || IsNil(o.SourceReleaseState) {
		return nil, false
	}
	return o.SourceReleaseState, true
}

// HasSourceReleaseState returns a boolean if a field has been set.
func (o *CatalogV1alpha1AssetResourceSpec) HasSourceReleaseState() bool {
	if o != nil && !IsNil(o.SourceReleaseState) {
		return true
	}

	return false
}

// SetSourceReleaseState gets a reference to the given CatalogV1alpha1AssetResourceSpecSourceReleaseState and assigns it to the SourceReleaseState field.
func (o *CatalogV1alpha1AssetResourceSpec) SetSourceReleaseState(v CatalogV1alpha1AssetResourceSpecSourceReleaseState) {
	o.SourceReleaseState = &v
}

func (o CatalogV1alpha1AssetResourceSpec) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogV1alpha1AssetResourceSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Stage) {
		toSerialize["stage"] = o.Stage
	}
	if !IsNil(o.AssetRequestDefinition) {
		toSerialize["assetRequestDefinition"] = o.AssetRequestDefinition
	}
	if !IsNil(o.CredentialRequestDefinitions) {
		toSerialize["credentialRequestDefinitions"] = o.CredentialRequestDefinitions
	}
	toSerialize["type"] = o.Type
	if !IsNil(o.ContentType) {
		toSerialize["contentType"] = o.ContentType
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	toSerialize["definition"] = o.Definition
	toSerialize["status"] = o.Status
	if !IsNil(o.AccessInfo) {
		toSerialize["accessInfo"] = o.AccessInfo
	}
	if !IsNil(o.SourceReleaseState) {
		toSerialize["sourceReleaseState"] = o.SourceReleaseState
	}
	return toSerialize, nil
}

func (o *CatalogV1alpha1AssetResourceSpec) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"definition",
		"status",
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

	varCatalogV1alpha1AssetResourceSpec := _CatalogV1alpha1AssetResourceSpec{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCatalogV1alpha1AssetResourceSpec)

	if err != nil {
		return err
	}

	*o = CatalogV1alpha1AssetResourceSpec(varCatalogV1alpha1AssetResourceSpec)

	return err
}

type NullableCatalogV1alpha1AssetResourceSpec struct {
	value *CatalogV1alpha1AssetResourceSpec
	isSet bool
}

func (v NullableCatalogV1alpha1AssetResourceSpec) Get() *CatalogV1alpha1AssetResourceSpec {
	return v.value
}

func (v *NullableCatalogV1alpha1AssetResourceSpec) Set(val *CatalogV1alpha1AssetResourceSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogV1alpha1AssetResourceSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogV1alpha1AssetResourceSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogV1alpha1AssetResourceSpec(val *CatalogV1alpha1AssetResourceSpec) *NullableCatalogV1alpha1AssetResourceSpec {
	return &NullableCatalogV1alpha1AssetResourceSpec{value: val, isSet: true}
}

func (v NullableCatalogV1alpha1AssetResourceSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogV1alpha1AssetResourceSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


