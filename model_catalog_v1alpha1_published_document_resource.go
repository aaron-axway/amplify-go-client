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

// checks if the CatalogV1alpha1PublishedDocumentResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogV1alpha1PublishedDocumentResource{}

// CatalogV1alpha1PublishedDocumentResource Definition of PublishedDocumentResource for version v1alpha1 in group catalog.    Resource representing a Published DocumentResource into a Marketplace
type CatalogV1alpha1PublishedDocumentResource struct {
	// Resource belongs to the catalog group. Cannot be updated. The server infers this from the endpoint the client submits the request to.
	Group *string `json:"group,omitempty" validate:"regexp=^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\\\\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*"`
	// Resource version is v1alpha1. The version defines the structure of the resource. Cannot be updated. The server infers this from the endpoint the client submits the request to.
	ApiVersion *string `json:"apiVersion,omitempty"`
	// Resource of kind PublishedDocumentResource. Cannot be updated. The server infers this from the endpoint the client submits the request to.
	Kind *string `json:"kind,omitempty"`
	// Name of the PublishedDocumentResource. PublishedDocumentResource name is unique and cannot be updated. The characters allowed in names are: digits (0-9), lower case letters (a-z), -, and .
	Name *string `json:"name,omitempty" validate:"regexp=^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\\\\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*"`
	// The friendly name of the PublishedDocumentResource.
	Title *string `json:"title,omitempty"`
	Metadata *ApiV1Metadata `json:"metadata,omitempty"`
	Owner *ApiV1Owner `json:"owner,omitempty"`
	// List of finalizers
	Finalizers []ApiV1Finalizer `json:"finalizers,omitempty"`
	// Custom attributes added to objects.
	Attributes *map[string]string `json:"attributes,omitempty"`
	// List of tags.
	Tags []string `json:"tags,omitempty"`
	Spec CatalogV1alpha1PublishedDocumentResourceSpec `json:"spec"`
	Languages *ApiV1Languages `json:"languages,omitempty"`
	LanguagesEnUs *ApiV1Language `json:"languages-en-us,omitempty"`
	LanguagesPtBr *ApiV1Language `json:"languages-pt-br,omitempty"`
	LanguagesDeDe *ApiV1Language `json:"languages-de-de,omitempty"`
	LanguagesFrFr *ApiV1Language `json:"languages-fr-fr,omitempty"`
	Embedded *ApiV1Embedded `json:"_embedded,omitempty"`
}

type _CatalogV1alpha1PublishedDocumentResource CatalogV1alpha1PublishedDocumentResource

// NewCatalogV1alpha1PublishedDocumentResource instantiates a new CatalogV1alpha1PublishedDocumentResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogV1alpha1PublishedDocumentResource(spec CatalogV1alpha1PublishedDocumentResourceSpec) *CatalogV1alpha1PublishedDocumentResource {
	this := CatalogV1alpha1PublishedDocumentResource{}
	this.Spec = spec
	return &this
}

// NewCatalogV1alpha1PublishedDocumentResourceWithDefaults instantiates a new CatalogV1alpha1PublishedDocumentResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogV1alpha1PublishedDocumentResourceWithDefaults() *CatalogV1alpha1PublishedDocumentResource {
	this := CatalogV1alpha1PublishedDocumentResource{}
	return &this
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *CatalogV1alpha1PublishedDocumentResource) GetGroup() string {
	if o == nil || IsNil(o.Group) {
		var ret string
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1PublishedDocumentResource) GetGroupOk() (*string, bool) {
	if o == nil || IsNil(o.Group) {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *CatalogV1alpha1PublishedDocumentResource) HasGroup() bool {
	if o != nil && !IsNil(o.Group) {
		return true
	}

	return false
}

// SetGroup gets a reference to the given string and assigns it to the Group field.
func (o *CatalogV1alpha1PublishedDocumentResource) SetGroup(v string) {
	o.Group = &v
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *CatalogV1alpha1PublishedDocumentResource) GetApiVersion() string {
	if o == nil || IsNil(o.ApiVersion) {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1PublishedDocumentResource) GetApiVersionOk() (*string, bool) {
	if o == nil || IsNil(o.ApiVersion) {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *CatalogV1alpha1PublishedDocumentResource) HasApiVersion() bool {
	if o != nil && !IsNil(o.ApiVersion) {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *CatalogV1alpha1PublishedDocumentResource) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *CatalogV1alpha1PublishedDocumentResource) GetKind() string {
	if o == nil || IsNil(o.Kind) {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1PublishedDocumentResource) GetKindOk() (*string, bool) {
	if o == nil || IsNil(o.Kind) {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *CatalogV1alpha1PublishedDocumentResource) HasKind() bool {
	if o != nil && !IsNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *CatalogV1alpha1PublishedDocumentResource) SetKind(v string) {
	o.Kind = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CatalogV1alpha1PublishedDocumentResource) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1PublishedDocumentResource) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CatalogV1alpha1PublishedDocumentResource) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CatalogV1alpha1PublishedDocumentResource) SetName(v string) {
	o.Name = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *CatalogV1alpha1PublishedDocumentResource) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1PublishedDocumentResource) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *CatalogV1alpha1PublishedDocumentResource) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *CatalogV1alpha1PublishedDocumentResource) SetTitle(v string) {
	o.Title = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *CatalogV1alpha1PublishedDocumentResource) GetMetadata() ApiV1Metadata {
	if o == nil || IsNil(o.Metadata) {
		var ret ApiV1Metadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1PublishedDocumentResource) GetMetadataOk() (*ApiV1Metadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *CatalogV1alpha1PublishedDocumentResource) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given ApiV1Metadata and assigns it to the Metadata field.
func (o *CatalogV1alpha1PublishedDocumentResource) SetMetadata(v ApiV1Metadata) {
	o.Metadata = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *CatalogV1alpha1PublishedDocumentResource) GetOwner() ApiV1Owner {
	if o == nil || IsNil(o.Owner) {
		var ret ApiV1Owner
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1PublishedDocumentResource) GetOwnerOk() (*ApiV1Owner, bool) {
	if o == nil || IsNil(o.Owner) {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *CatalogV1alpha1PublishedDocumentResource) HasOwner() bool {
	if o != nil && !IsNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given ApiV1Owner and assigns it to the Owner field.
func (o *CatalogV1alpha1PublishedDocumentResource) SetOwner(v ApiV1Owner) {
	o.Owner = &v
}

// GetFinalizers returns the Finalizers field value if set, zero value otherwise.
func (o *CatalogV1alpha1PublishedDocumentResource) GetFinalizers() []ApiV1Finalizer {
	if o == nil || IsNil(o.Finalizers) {
		var ret []ApiV1Finalizer
		return ret
	}
	return o.Finalizers
}

// GetFinalizersOk returns a tuple with the Finalizers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1PublishedDocumentResource) GetFinalizersOk() ([]ApiV1Finalizer, bool) {
	if o == nil || IsNil(o.Finalizers) {
		return nil, false
	}
	return o.Finalizers, true
}

// HasFinalizers returns a boolean if a field has been set.
func (o *CatalogV1alpha1PublishedDocumentResource) HasFinalizers() bool {
	if o != nil && !IsNil(o.Finalizers) {
		return true
	}

	return false
}

// SetFinalizers gets a reference to the given []ApiV1Finalizer and assigns it to the Finalizers field.
func (o *CatalogV1alpha1PublishedDocumentResource) SetFinalizers(v []ApiV1Finalizer) {
	o.Finalizers = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *CatalogV1alpha1PublishedDocumentResource) GetAttributes() map[string]string {
	if o == nil || IsNil(o.Attributes) {
		var ret map[string]string
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1PublishedDocumentResource) GetAttributesOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *CatalogV1alpha1PublishedDocumentResource) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]string and assigns it to the Attributes field.
func (o *CatalogV1alpha1PublishedDocumentResource) SetAttributes(v map[string]string) {
	o.Attributes = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *CatalogV1alpha1PublishedDocumentResource) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1PublishedDocumentResource) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *CatalogV1alpha1PublishedDocumentResource) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *CatalogV1alpha1PublishedDocumentResource) SetTags(v []string) {
	o.Tags = v
}

// GetSpec returns the Spec field value
func (o *CatalogV1alpha1PublishedDocumentResource) GetSpec() CatalogV1alpha1PublishedDocumentResourceSpec {
	if o == nil {
		var ret CatalogV1alpha1PublishedDocumentResourceSpec
		return ret
	}

	return o.Spec
}

// GetSpecOk returns a tuple with the Spec field value
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1PublishedDocumentResource) GetSpecOk() (*CatalogV1alpha1PublishedDocumentResourceSpec, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Spec, true
}

// SetSpec sets field value
func (o *CatalogV1alpha1PublishedDocumentResource) SetSpec(v CatalogV1alpha1PublishedDocumentResourceSpec) {
	o.Spec = v
}

// GetLanguages returns the Languages field value if set, zero value otherwise.
func (o *CatalogV1alpha1PublishedDocumentResource) GetLanguages() ApiV1Languages {
	if o == nil || IsNil(o.Languages) {
		var ret ApiV1Languages
		return ret
	}
	return *o.Languages
}

// GetLanguagesOk returns a tuple with the Languages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1PublishedDocumentResource) GetLanguagesOk() (*ApiV1Languages, bool) {
	if o == nil || IsNil(o.Languages) {
		return nil, false
	}
	return o.Languages, true
}

// HasLanguages returns a boolean if a field has been set.
func (o *CatalogV1alpha1PublishedDocumentResource) HasLanguages() bool {
	if o != nil && !IsNil(o.Languages) {
		return true
	}

	return false
}

// SetLanguages gets a reference to the given ApiV1Languages and assigns it to the Languages field.
func (o *CatalogV1alpha1PublishedDocumentResource) SetLanguages(v ApiV1Languages) {
	o.Languages = &v
}

// GetLanguagesEnUs returns the LanguagesEnUs field value if set, zero value otherwise.
func (o *CatalogV1alpha1PublishedDocumentResource) GetLanguagesEnUs() ApiV1Language {
	if o == nil || IsNil(o.LanguagesEnUs) {
		var ret ApiV1Language
		return ret
	}
	return *o.LanguagesEnUs
}

// GetLanguagesEnUsOk returns a tuple with the LanguagesEnUs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1PublishedDocumentResource) GetLanguagesEnUsOk() (*ApiV1Language, bool) {
	if o == nil || IsNil(o.LanguagesEnUs) {
		return nil, false
	}
	return o.LanguagesEnUs, true
}

// HasLanguagesEnUs returns a boolean if a field has been set.
func (o *CatalogV1alpha1PublishedDocumentResource) HasLanguagesEnUs() bool {
	if o != nil && !IsNil(o.LanguagesEnUs) {
		return true
	}

	return false
}

// SetLanguagesEnUs gets a reference to the given ApiV1Language and assigns it to the LanguagesEnUs field.
func (o *CatalogV1alpha1PublishedDocumentResource) SetLanguagesEnUs(v ApiV1Language) {
	o.LanguagesEnUs = &v
}

// GetLanguagesPtBr returns the LanguagesPtBr field value if set, zero value otherwise.
func (o *CatalogV1alpha1PublishedDocumentResource) GetLanguagesPtBr() ApiV1Language {
	if o == nil || IsNil(o.LanguagesPtBr) {
		var ret ApiV1Language
		return ret
	}
	return *o.LanguagesPtBr
}

// GetLanguagesPtBrOk returns a tuple with the LanguagesPtBr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1PublishedDocumentResource) GetLanguagesPtBrOk() (*ApiV1Language, bool) {
	if o == nil || IsNil(o.LanguagesPtBr) {
		return nil, false
	}
	return o.LanguagesPtBr, true
}

// HasLanguagesPtBr returns a boolean if a field has been set.
func (o *CatalogV1alpha1PublishedDocumentResource) HasLanguagesPtBr() bool {
	if o != nil && !IsNil(o.LanguagesPtBr) {
		return true
	}

	return false
}

// SetLanguagesPtBr gets a reference to the given ApiV1Language and assigns it to the LanguagesPtBr field.
func (o *CatalogV1alpha1PublishedDocumentResource) SetLanguagesPtBr(v ApiV1Language) {
	o.LanguagesPtBr = &v
}

// GetLanguagesDeDe returns the LanguagesDeDe field value if set, zero value otherwise.
func (o *CatalogV1alpha1PublishedDocumentResource) GetLanguagesDeDe() ApiV1Language {
	if o == nil || IsNil(o.LanguagesDeDe) {
		var ret ApiV1Language
		return ret
	}
	return *o.LanguagesDeDe
}

// GetLanguagesDeDeOk returns a tuple with the LanguagesDeDe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1PublishedDocumentResource) GetLanguagesDeDeOk() (*ApiV1Language, bool) {
	if o == nil || IsNil(o.LanguagesDeDe) {
		return nil, false
	}
	return o.LanguagesDeDe, true
}

// HasLanguagesDeDe returns a boolean if a field has been set.
func (o *CatalogV1alpha1PublishedDocumentResource) HasLanguagesDeDe() bool {
	if o != nil && !IsNil(o.LanguagesDeDe) {
		return true
	}

	return false
}

// SetLanguagesDeDe gets a reference to the given ApiV1Language and assigns it to the LanguagesDeDe field.
func (o *CatalogV1alpha1PublishedDocumentResource) SetLanguagesDeDe(v ApiV1Language) {
	o.LanguagesDeDe = &v
}

// GetLanguagesFrFr returns the LanguagesFrFr field value if set, zero value otherwise.
func (o *CatalogV1alpha1PublishedDocumentResource) GetLanguagesFrFr() ApiV1Language {
	if o == nil || IsNil(o.LanguagesFrFr) {
		var ret ApiV1Language
		return ret
	}
	return *o.LanguagesFrFr
}

// GetLanguagesFrFrOk returns a tuple with the LanguagesFrFr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1PublishedDocumentResource) GetLanguagesFrFrOk() (*ApiV1Language, bool) {
	if o == nil || IsNil(o.LanguagesFrFr) {
		return nil, false
	}
	return o.LanguagesFrFr, true
}

// HasLanguagesFrFr returns a boolean if a field has been set.
func (o *CatalogV1alpha1PublishedDocumentResource) HasLanguagesFrFr() bool {
	if o != nil && !IsNil(o.LanguagesFrFr) {
		return true
	}

	return false
}

// SetLanguagesFrFr gets a reference to the given ApiV1Language and assigns it to the LanguagesFrFr field.
func (o *CatalogV1alpha1PublishedDocumentResource) SetLanguagesFrFr(v ApiV1Language) {
	o.LanguagesFrFr = &v
}

// GetEmbedded returns the Embedded field value if set, zero value otherwise.
func (o *CatalogV1alpha1PublishedDocumentResource) GetEmbedded() ApiV1Embedded {
	if o == nil || IsNil(o.Embedded) {
		var ret ApiV1Embedded
		return ret
	}
	return *o.Embedded
}

// GetEmbeddedOk returns a tuple with the Embedded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1PublishedDocumentResource) GetEmbeddedOk() (*ApiV1Embedded, bool) {
	if o == nil || IsNil(o.Embedded) {
		return nil, false
	}
	return o.Embedded, true
}

// HasEmbedded returns a boolean if a field has been set.
func (o *CatalogV1alpha1PublishedDocumentResource) HasEmbedded() bool {
	if o != nil && !IsNil(o.Embedded) {
		return true
	}

	return false
}

// SetEmbedded gets a reference to the given ApiV1Embedded and assigns it to the Embedded field.
func (o *CatalogV1alpha1PublishedDocumentResource) SetEmbedded(v ApiV1Embedded) {
	o.Embedded = &v
}

func (o CatalogV1alpha1PublishedDocumentResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogV1alpha1PublishedDocumentResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Group) {
		toSerialize["group"] = o.Group
	}
	if !IsNil(o.ApiVersion) {
		toSerialize["apiVersion"] = o.ApiVersion
	}
	if !IsNil(o.Kind) {
		toSerialize["kind"] = o.Kind
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.Owner) {
		toSerialize["owner"] = o.Owner
	}
	if !IsNil(o.Finalizers) {
		toSerialize["finalizers"] = o.Finalizers
	}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	toSerialize["spec"] = o.Spec
	if !IsNil(o.Languages) {
		toSerialize["languages"] = o.Languages
	}
	if !IsNil(o.LanguagesEnUs) {
		toSerialize["languages-en-us"] = o.LanguagesEnUs
	}
	if !IsNil(o.LanguagesPtBr) {
		toSerialize["languages-pt-br"] = o.LanguagesPtBr
	}
	if !IsNil(o.LanguagesDeDe) {
		toSerialize["languages-de-de"] = o.LanguagesDeDe
	}
	if !IsNil(o.LanguagesFrFr) {
		toSerialize["languages-fr-fr"] = o.LanguagesFrFr
	}
	if !IsNil(o.Embedded) {
		toSerialize["_embedded"] = o.Embedded
	}
	return toSerialize, nil
}

func (o *CatalogV1alpha1PublishedDocumentResource) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"spec",
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

	varCatalogV1alpha1PublishedDocumentResource := _CatalogV1alpha1PublishedDocumentResource{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCatalogV1alpha1PublishedDocumentResource)

	if err != nil {
		return err
	}

	*o = CatalogV1alpha1PublishedDocumentResource(varCatalogV1alpha1PublishedDocumentResource)

	return err
}

type NullableCatalogV1alpha1PublishedDocumentResource struct {
	value *CatalogV1alpha1PublishedDocumentResource
	isSet bool
}

func (v NullableCatalogV1alpha1PublishedDocumentResource) Get() *CatalogV1alpha1PublishedDocumentResource {
	return v.value
}

func (v *NullableCatalogV1alpha1PublishedDocumentResource) Set(val *CatalogV1alpha1PublishedDocumentResource) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogV1alpha1PublishedDocumentResource) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogV1alpha1PublishedDocumentResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogV1alpha1PublishedDocumentResource(val *CatalogV1alpha1PublishedDocumentResource) *NullableCatalogV1alpha1PublishedDocumentResource {
	return &NullableCatalogV1alpha1PublishedDocumentResource{value: val, isSet: true}
}

func (v NullableCatalogV1alpha1PublishedDocumentResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogV1alpha1PublishedDocumentResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


