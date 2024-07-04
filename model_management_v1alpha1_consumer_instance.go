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

// checks if the ManagementV1alpha1ConsumerInstance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManagementV1alpha1ConsumerInstance{}

// ManagementV1alpha1ConsumerInstance Definition of ConsumerInstance for version v1alpha1 in group management.    Resource representing a Unified Catalog Item of type API
type ManagementV1alpha1ConsumerInstance struct {
	// Resource belongs to the management group. Cannot be updated. The server infers this from the endpoint the client submits the request to.
	Group *string `json:"group,omitempty" validate:"regexp=^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\\\\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*"`
	// Resource version is v1alpha1. The version defines the structure of the resource. Cannot be updated. The server infers this from the endpoint the client submits the request to.
	ApiVersion *string `json:"apiVersion,omitempty"`
	// Resource of kind ConsumerInstance. Cannot be updated. The server infers this from the endpoint the client submits the request to.
	Kind *string `json:"kind,omitempty"`
	// Name of the ConsumerInstance. ConsumerInstance name is unique and cannot be updated. The characters allowed in names are: digits (0-9), lower case letters (a-z), -, and .
	Name *string `json:"name,omitempty" validate:"regexp=^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\\\\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*"`
	// The friendly name of the ConsumerInstance.
	Title *string `json:"title,omitempty"`
	Metadata *ApiV1Metadata `json:"metadata,omitempty"`
	Owner *ApiV1Owner `json:"owner,omitempty"`
	// List of finalizers
	Finalizers []ApiV1Finalizer `json:"finalizers,omitempty"`
	// Custom attributes added to objects.
	Attributes *map[string]string `json:"attributes,omitempty"`
	// List of tags.
	Tags []string `json:"tags,omitempty"`
	Spec ManagementV1alpha1ConsumerInstanceSpec `json:"spec"`
	Status *ManagementV1alpha1ConsumerInstanceStatus `json:"status,omitempty"`
	References *ManagementV1alpha1ConsumerInstanceReferences `json:"references,omitempty"`
	Languages *ApiV1Languages `json:"languages,omitempty"`
	LanguagesEnUs *ApiV1Language `json:"languages-en-us,omitempty"`
	LanguagesPtBr *ApiV1Language `json:"languages-pt-br,omitempty"`
	LanguagesDeDe *ApiV1Language `json:"languages-de-de,omitempty"`
	LanguagesFrFr *ApiV1Language `json:"languages-fr-fr,omitempty"`
	Embedded *ApiV1Embedded `json:"_embedded,omitempty"`
}

type _ManagementV1alpha1ConsumerInstance ManagementV1alpha1ConsumerInstance

// NewManagementV1alpha1ConsumerInstance instantiates a new ManagementV1alpha1ConsumerInstance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagementV1alpha1ConsumerInstance(spec ManagementV1alpha1ConsumerInstanceSpec) *ManagementV1alpha1ConsumerInstance {
	this := ManagementV1alpha1ConsumerInstance{}
	this.Spec = spec
	return &this
}

// NewManagementV1alpha1ConsumerInstanceWithDefaults instantiates a new ManagementV1alpha1ConsumerInstance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagementV1alpha1ConsumerInstanceWithDefaults() *ManagementV1alpha1ConsumerInstance {
	this := ManagementV1alpha1ConsumerInstance{}
	return &this
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *ManagementV1alpha1ConsumerInstance) GetGroup() string {
	if o == nil || IsNil(o.Group) {
		var ret string
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1ConsumerInstance) GetGroupOk() (*string, bool) {
	if o == nil || IsNil(o.Group) {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *ManagementV1alpha1ConsumerInstance) HasGroup() bool {
	if o != nil && !IsNil(o.Group) {
		return true
	}

	return false
}

// SetGroup gets a reference to the given string and assigns it to the Group field.
func (o *ManagementV1alpha1ConsumerInstance) SetGroup(v string) {
	o.Group = &v
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *ManagementV1alpha1ConsumerInstance) GetApiVersion() string {
	if o == nil || IsNil(o.ApiVersion) {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1ConsumerInstance) GetApiVersionOk() (*string, bool) {
	if o == nil || IsNil(o.ApiVersion) {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *ManagementV1alpha1ConsumerInstance) HasApiVersion() bool {
	if o != nil && !IsNil(o.ApiVersion) {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *ManagementV1alpha1ConsumerInstance) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *ManagementV1alpha1ConsumerInstance) GetKind() string {
	if o == nil || IsNil(o.Kind) {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1ConsumerInstance) GetKindOk() (*string, bool) {
	if o == nil || IsNil(o.Kind) {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *ManagementV1alpha1ConsumerInstance) HasKind() bool {
	if o != nil && !IsNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *ManagementV1alpha1ConsumerInstance) SetKind(v string) {
	o.Kind = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ManagementV1alpha1ConsumerInstance) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1ConsumerInstance) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ManagementV1alpha1ConsumerInstance) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ManagementV1alpha1ConsumerInstance) SetName(v string) {
	o.Name = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *ManagementV1alpha1ConsumerInstance) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1ConsumerInstance) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ManagementV1alpha1ConsumerInstance) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *ManagementV1alpha1ConsumerInstance) SetTitle(v string) {
	o.Title = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *ManagementV1alpha1ConsumerInstance) GetMetadata() ApiV1Metadata {
	if o == nil || IsNil(o.Metadata) {
		var ret ApiV1Metadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1ConsumerInstance) GetMetadataOk() (*ApiV1Metadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *ManagementV1alpha1ConsumerInstance) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given ApiV1Metadata and assigns it to the Metadata field.
func (o *ManagementV1alpha1ConsumerInstance) SetMetadata(v ApiV1Metadata) {
	o.Metadata = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *ManagementV1alpha1ConsumerInstance) GetOwner() ApiV1Owner {
	if o == nil || IsNil(o.Owner) {
		var ret ApiV1Owner
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1ConsumerInstance) GetOwnerOk() (*ApiV1Owner, bool) {
	if o == nil || IsNil(o.Owner) {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *ManagementV1alpha1ConsumerInstance) HasOwner() bool {
	if o != nil && !IsNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given ApiV1Owner and assigns it to the Owner field.
func (o *ManagementV1alpha1ConsumerInstance) SetOwner(v ApiV1Owner) {
	o.Owner = &v
}

// GetFinalizers returns the Finalizers field value if set, zero value otherwise.
func (o *ManagementV1alpha1ConsumerInstance) GetFinalizers() []ApiV1Finalizer {
	if o == nil || IsNil(o.Finalizers) {
		var ret []ApiV1Finalizer
		return ret
	}
	return o.Finalizers
}

// GetFinalizersOk returns a tuple with the Finalizers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1ConsumerInstance) GetFinalizersOk() ([]ApiV1Finalizer, bool) {
	if o == nil || IsNil(o.Finalizers) {
		return nil, false
	}
	return o.Finalizers, true
}

// HasFinalizers returns a boolean if a field has been set.
func (o *ManagementV1alpha1ConsumerInstance) HasFinalizers() bool {
	if o != nil && !IsNil(o.Finalizers) {
		return true
	}

	return false
}

// SetFinalizers gets a reference to the given []ApiV1Finalizer and assigns it to the Finalizers field.
func (o *ManagementV1alpha1ConsumerInstance) SetFinalizers(v []ApiV1Finalizer) {
	o.Finalizers = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *ManagementV1alpha1ConsumerInstance) GetAttributes() map[string]string {
	if o == nil || IsNil(o.Attributes) {
		var ret map[string]string
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1ConsumerInstance) GetAttributesOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *ManagementV1alpha1ConsumerInstance) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]string and assigns it to the Attributes field.
func (o *ManagementV1alpha1ConsumerInstance) SetAttributes(v map[string]string) {
	o.Attributes = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ManagementV1alpha1ConsumerInstance) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1ConsumerInstance) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ManagementV1alpha1ConsumerInstance) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *ManagementV1alpha1ConsumerInstance) SetTags(v []string) {
	o.Tags = v
}

// GetSpec returns the Spec field value
func (o *ManagementV1alpha1ConsumerInstance) GetSpec() ManagementV1alpha1ConsumerInstanceSpec {
	if o == nil {
		var ret ManagementV1alpha1ConsumerInstanceSpec
		return ret
	}

	return o.Spec
}

// GetSpecOk returns a tuple with the Spec field value
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1ConsumerInstance) GetSpecOk() (*ManagementV1alpha1ConsumerInstanceSpec, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Spec, true
}

// SetSpec sets field value
func (o *ManagementV1alpha1ConsumerInstance) SetSpec(v ManagementV1alpha1ConsumerInstanceSpec) {
	o.Spec = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ManagementV1alpha1ConsumerInstance) GetStatus() ManagementV1alpha1ConsumerInstanceStatus {
	if o == nil || IsNil(o.Status) {
		var ret ManagementV1alpha1ConsumerInstanceStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1ConsumerInstance) GetStatusOk() (*ManagementV1alpha1ConsumerInstanceStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ManagementV1alpha1ConsumerInstance) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given ManagementV1alpha1ConsumerInstanceStatus and assigns it to the Status field.
func (o *ManagementV1alpha1ConsumerInstance) SetStatus(v ManagementV1alpha1ConsumerInstanceStatus) {
	o.Status = &v
}

// GetReferences returns the References field value if set, zero value otherwise.
func (o *ManagementV1alpha1ConsumerInstance) GetReferences() ManagementV1alpha1ConsumerInstanceReferences {
	if o == nil || IsNil(o.References) {
		var ret ManagementV1alpha1ConsumerInstanceReferences
		return ret
	}
	return *o.References
}

// GetReferencesOk returns a tuple with the References field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1ConsumerInstance) GetReferencesOk() (*ManagementV1alpha1ConsumerInstanceReferences, bool) {
	if o == nil || IsNil(o.References) {
		return nil, false
	}
	return o.References, true
}

// HasReferences returns a boolean if a field has been set.
func (o *ManagementV1alpha1ConsumerInstance) HasReferences() bool {
	if o != nil && !IsNil(o.References) {
		return true
	}

	return false
}

// SetReferences gets a reference to the given ManagementV1alpha1ConsumerInstanceReferences and assigns it to the References field.
func (o *ManagementV1alpha1ConsumerInstance) SetReferences(v ManagementV1alpha1ConsumerInstanceReferences) {
	o.References = &v
}

// GetLanguages returns the Languages field value if set, zero value otherwise.
func (o *ManagementV1alpha1ConsumerInstance) GetLanguages() ApiV1Languages {
	if o == nil || IsNil(o.Languages) {
		var ret ApiV1Languages
		return ret
	}
	return *o.Languages
}

// GetLanguagesOk returns a tuple with the Languages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1ConsumerInstance) GetLanguagesOk() (*ApiV1Languages, bool) {
	if o == nil || IsNil(o.Languages) {
		return nil, false
	}
	return o.Languages, true
}

// HasLanguages returns a boolean if a field has been set.
func (o *ManagementV1alpha1ConsumerInstance) HasLanguages() bool {
	if o != nil && !IsNil(o.Languages) {
		return true
	}

	return false
}

// SetLanguages gets a reference to the given ApiV1Languages and assigns it to the Languages field.
func (o *ManagementV1alpha1ConsumerInstance) SetLanguages(v ApiV1Languages) {
	o.Languages = &v
}

// GetLanguagesEnUs returns the LanguagesEnUs field value if set, zero value otherwise.
func (o *ManagementV1alpha1ConsumerInstance) GetLanguagesEnUs() ApiV1Language {
	if o == nil || IsNil(o.LanguagesEnUs) {
		var ret ApiV1Language
		return ret
	}
	return *o.LanguagesEnUs
}

// GetLanguagesEnUsOk returns a tuple with the LanguagesEnUs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1ConsumerInstance) GetLanguagesEnUsOk() (*ApiV1Language, bool) {
	if o == nil || IsNil(o.LanguagesEnUs) {
		return nil, false
	}
	return o.LanguagesEnUs, true
}

// HasLanguagesEnUs returns a boolean if a field has been set.
func (o *ManagementV1alpha1ConsumerInstance) HasLanguagesEnUs() bool {
	if o != nil && !IsNil(o.LanguagesEnUs) {
		return true
	}

	return false
}

// SetLanguagesEnUs gets a reference to the given ApiV1Language and assigns it to the LanguagesEnUs field.
func (o *ManagementV1alpha1ConsumerInstance) SetLanguagesEnUs(v ApiV1Language) {
	o.LanguagesEnUs = &v
}

// GetLanguagesPtBr returns the LanguagesPtBr field value if set, zero value otherwise.
func (o *ManagementV1alpha1ConsumerInstance) GetLanguagesPtBr() ApiV1Language {
	if o == nil || IsNil(o.LanguagesPtBr) {
		var ret ApiV1Language
		return ret
	}
	return *o.LanguagesPtBr
}

// GetLanguagesPtBrOk returns a tuple with the LanguagesPtBr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1ConsumerInstance) GetLanguagesPtBrOk() (*ApiV1Language, bool) {
	if o == nil || IsNil(o.LanguagesPtBr) {
		return nil, false
	}
	return o.LanguagesPtBr, true
}

// HasLanguagesPtBr returns a boolean if a field has been set.
func (o *ManagementV1alpha1ConsumerInstance) HasLanguagesPtBr() bool {
	if o != nil && !IsNil(o.LanguagesPtBr) {
		return true
	}

	return false
}

// SetLanguagesPtBr gets a reference to the given ApiV1Language and assigns it to the LanguagesPtBr field.
func (o *ManagementV1alpha1ConsumerInstance) SetLanguagesPtBr(v ApiV1Language) {
	o.LanguagesPtBr = &v
}

// GetLanguagesDeDe returns the LanguagesDeDe field value if set, zero value otherwise.
func (o *ManagementV1alpha1ConsumerInstance) GetLanguagesDeDe() ApiV1Language {
	if o == nil || IsNil(o.LanguagesDeDe) {
		var ret ApiV1Language
		return ret
	}
	return *o.LanguagesDeDe
}

// GetLanguagesDeDeOk returns a tuple with the LanguagesDeDe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1ConsumerInstance) GetLanguagesDeDeOk() (*ApiV1Language, bool) {
	if o == nil || IsNil(o.LanguagesDeDe) {
		return nil, false
	}
	return o.LanguagesDeDe, true
}

// HasLanguagesDeDe returns a boolean if a field has been set.
func (o *ManagementV1alpha1ConsumerInstance) HasLanguagesDeDe() bool {
	if o != nil && !IsNil(o.LanguagesDeDe) {
		return true
	}

	return false
}

// SetLanguagesDeDe gets a reference to the given ApiV1Language and assigns it to the LanguagesDeDe field.
func (o *ManagementV1alpha1ConsumerInstance) SetLanguagesDeDe(v ApiV1Language) {
	o.LanguagesDeDe = &v
}

// GetLanguagesFrFr returns the LanguagesFrFr field value if set, zero value otherwise.
func (o *ManagementV1alpha1ConsumerInstance) GetLanguagesFrFr() ApiV1Language {
	if o == nil || IsNil(o.LanguagesFrFr) {
		var ret ApiV1Language
		return ret
	}
	return *o.LanguagesFrFr
}

// GetLanguagesFrFrOk returns a tuple with the LanguagesFrFr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1ConsumerInstance) GetLanguagesFrFrOk() (*ApiV1Language, bool) {
	if o == nil || IsNil(o.LanguagesFrFr) {
		return nil, false
	}
	return o.LanguagesFrFr, true
}

// HasLanguagesFrFr returns a boolean if a field has been set.
func (o *ManagementV1alpha1ConsumerInstance) HasLanguagesFrFr() bool {
	if o != nil && !IsNil(o.LanguagesFrFr) {
		return true
	}

	return false
}

// SetLanguagesFrFr gets a reference to the given ApiV1Language and assigns it to the LanguagesFrFr field.
func (o *ManagementV1alpha1ConsumerInstance) SetLanguagesFrFr(v ApiV1Language) {
	o.LanguagesFrFr = &v
}

// GetEmbedded returns the Embedded field value if set, zero value otherwise.
func (o *ManagementV1alpha1ConsumerInstance) GetEmbedded() ApiV1Embedded {
	if o == nil || IsNil(o.Embedded) {
		var ret ApiV1Embedded
		return ret
	}
	return *o.Embedded
}

// GetEmbeddedOk returns a tuple with the Embedded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1ConsumerInstance) GetEmbeddedOk() (*ApiV1Embedded, bool) {
	if o == nil || IsNil(o.Embedded) {
		return nil, false
	}
	return o.Embedded, true
}

// HasEmbedded returns a boolean if a field has been set.
func (o *ManagementV1alpha1ConsumerInstance) HasEmbedded() bool {
	if o != nil && !IsNil(o.Embedded) {
		return true
	}

	return false
}

// SetEmbedded gets a reference to the given ApiV1Embedded and assigns it to the Embedded field.
func (o *ManagementV1alpha1ConsumerInstance) SetEmbedded(v ApiV1Embedded) {
	o.Embedded = &v
}

func (o ManagementV1alpha1ConsumerInstance) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManagementV1alpha1ConsumerInstance) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.References) {
		toSerialize["references"] = o.References
	}
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

func (o *ManagementV1alpha1ConsumerInstance) UnmarshalJSON(data []byte) (err error) {
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

	varManagementV1alpha1ConsumerInstance := _ManagementV1alpha1ConsumerInstance{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varManagementV1alpha1ConsumerInstance)

	if err != nil {
		return err
	}

	*o = ManagementV1alpha1ConsumerInstance(varManagementV1alpha1ConsumerInstance)

	return err
}

type NullableManagementV1alpha1ConsumerInstance struct {
	value *ManagementV1alpha1ConsumerInstance
	isSet bool
}

func (v NullableManagementV1alpha1ConsumerInstance) Get() *ManagementV1alpha1ConsumerInstance {
	return v.value
}

func (v *NullableManagementV1alpha1ConsumerInstance) Set(val *ManagementV1alpha1ConsumerInstance) {
	v.value = val
	v.isSet = true
}

func (v NullableManagementV1alpha1ConsumerInstance) IsSet() bool {
	return v.isSet
}

func (v *NullableManagementV1alpha1ConsumerInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagementV1alpha1ConsumerInstance(val *ManagementV1alpha1ConsumerInstance) *NullableManagementV1alpha1ConsumerInstance {
	return &NullableManagementV1alpha1ConsumerInstance{value: val, isSet: true}
}

func (v NullableManagementV1alpha1ConsumerInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagementV1alpha1ConsumerInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


