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

// checks if the ManagementV1alpha1APIServiceInstance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManagementV1alpha1APIServiceInstance{}

// ManagementV1alpha1APIServiceInstance Definition of APIServiceInstance for version v1alpha1 in group management.    Resource representing an Amplify Central API Service Instance
type ManagementV1alpha1APIServiceInstance struct {
	// Resource belongs to the management group. Cannot be updated. The server infers this from the endpoint the client submits the request to.
	Group *string `json:"group,omitempty" validate:"regexp=^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\\\\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*"`
	// Resource version is v1alpha1. The version defines the structure of the resource. Cannot be updated. The server infers this from the endpoint the client submits the request to.
	ApiVersion *string `json:"apiVersion,omitempty"`
	// Resource of kind APIServiceInstance. Cannot be updated. The server infers this from the endpoint the client submits the request to.
	Kind *string `json:"kind,omitempty"`
	// Name of the APIServiceInstance. APIServiceInstance name is unique and cannot be updated. The characters allowed in names are: digits (0-9), lower case letters (a-z), -, and .
	Name *string `json:"name,omitempty" validate:"regexp=^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\\\\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*"`
	// The friendly name of the APIServiceInstance.
	Title *string `json:"title,omitempty"`
	Metadata *ApiV1Metadata `json:"metadata,omitempty"`
	Owner *ApiV1Owner `json:"owner,omitempty"`
	// List of finalizers
	Finalizers []ApiV1Finalizer `json:"finalizers,omitempty"`
	// Custom attributes added to objects.
	Attributes *map[string]string `json:"attributes,omitempty"`
	// List of tags.
	Tags []string `json:"tags,omitempty"`
	Spec ManagementV1alpha1APIServiceInstanceSpec `json:"spec"`
	Lifecycle *ManagementV1alpha1APIServiceInstanceLifecycle `json:"lifecycle,omitempty"`
	References *ManagementV1alpha1APIServiceInstanceReferences `json:"references,omitempty"`
	Compliance *ManagementV1alpha1APIServiceInstanceCompliance `json:"compliance,omitempty"`
	Source *ManagementV1alpha1APIServiceInstanceSource `json:"source,omitempty"`
	Languages *ApiV1Languages `json:"languages,omitempty"`
	LanguagesEnUs *ApiV1Language `json:"languages-en-us,omitempty"`
	LanguagesPtBr *ApiV1Language `json:"languages-pt-br,omitempty"`
	LanguagesDeDe *ApiV1Language `json:"languages-de-de,omitempty"`
	LanguagesFrFr *ApiV1Language `json:"languages-fr-fr,omitempty"`
	Embedded *ApiV1Embedded `json:"_embedded,omitempty"`
}

type _ManagementV1alpha1APIServiceInstance ManagementV1alpha1APIServiceInstance

// NewManagementV1alpha1APIServiceInstance instantiates a new ManagementV1alpha1APIServiceInstance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagementV1alpha1APIServiceInstance(spec ManagementV1alpha1APIServiceInstanceSpec) *ManagementV1alpha1APIServiceInstance {
	this := ManagementV1alpha1APIServiceInstance{}
	this.Spec = spec
	return &this
}

// NewManagementV1alpha1APIServiceInstanceWithDefaults instantiates a new ManagementV1alpha1APIServiceInstance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagementV1alpha1APIServiceInstanceWithDefaults() *ManagementV1alpha1APIServiceInstance {
	this := ManagementV1alpha1APIServiceInstance{}
	return &this
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *ManagementV1alpha1APIServiceInstance) GetGroup() string {
	if o == nil || IsNil(o.Group) {
		var ret string
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1APIServiceInstance) GetGroupOk() (*string, bool) {
	if o == nil || IsNil(o.Group) {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *ManagementV1alpha1APIServiceInstance) HasGroup() bool {
	if o != nil && !IsNil(o.Group) {
		return true
	}

	return false
}

// SetGroup gets a reference to the given string and assigns it to the Group field.
func (o *ManagementV1alpha1APIServiceInstance) SetGroup(v string) {
	o.Group = &v
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *ManagementV1alpha1APIServiceInstance) GetApiVersion() string {
	if o == nil || IsNil(o.ApiVersion) {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1APIServiceInstance) GetApiVersionOk() (*string, bool) {
	if o == nil || IsNil(o.ApiVersion) {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *ManagementV1alpha1APIServiceInstance) HasApiVersion() bool {
	if o != nil && !IsNil(o.ApiVersion) {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *ManagementV1alpha1APIServiceInstance) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *ManagementV1alpha1APIServiceInstance) GetKind() string {
	if o == nil || IsNil(o.Kind) {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1APIServiceInstance) GetKindOk() (*string, bool) {
	if o == nil || IsNil(o.Kind) {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *ManagementV1alpha1APIServiceInstance) HasKind() bool {
	if o != nil && !IsNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *ManagementV1alpha1APIServiceInstance) SetKind(v string) {
	o.Kind = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ManagementV1alpha1APIServiceInstance) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1APIServiceInstance) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ManagementV1alpha1APIServiceInstance) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ManagementV1alpha1APIServiceInstance) SetName(v string) {
	o.Name = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *ManagementV1alpha1APIServiceInstance) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1APIServiceInstance) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ManagementV1alpha1APIServiceInstance) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *ManagementV1alpha1APIServiceInstance) SetTitle(v string) {
	o.Title = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *ManagementV1alpha1APIServiceInstance) GetMetadata() ApiV1Metadata {
	if o == nil || IsNil(o.Metadata) {
		var ret ApiV1Metadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1APIServiceInstance) GetMetadataOk() (*ApiV1Metadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *ManagementV1alpha1APIServiceInstance) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given ApiV1Metadata and assigns it to the Metadata field.
func (o *ManagementV1alpha1APIServiceInstance) SetMetadata(v ApiV1Metadata) {
	o.Metadata = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *ManagementV1alpha1APIServiceInstance) GetOwner() ApiV1Owner {
	if o == nil || IsNil(o.Owner) {
		var ret ApiV1Owner
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1APIServiceInstance) GetOwnerOk() (*ApiV1Owner, bool) {
	if o == nil || IsNil(o.Owner) {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *ManagementV1alpha1APIServiceInstance) HasOwner() bool {
	if o != nil && !IsNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given ApiV1Owner and assigns it to the Owner field.
func (o *ManagementV1alpha1APIServiceInstance) SetOwner(v ApiV1Owner) {
	o.Owner = &v
}

// GetFinalizers returns the Finalizers field value if set, zero value otherwise.
func (o *ManagementV1alpha1APIServiceInstance) GetFinalizers() []ApiV1Finalizer {
	if o == nil || IsNil(o.Finalizers) {
		var ret []ApiV1Finalizer
		return ret
	}
	return o.Finalizers
}

// GetFinalizersOk returns a tuple with the Finalizers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1APIServiceInstance) GetFinalizersOk() ([]ApiV1Finalizer, bool) {
	if o == nil || IsNil(o.Finalizers) {
		return nil, false
	}
	return o.Finalizers, true
}

// HasFinalizers returns a boolean if a field has been set.
func (o *ManagementV1alpha1APIServiceInstance) HasFinalizers() bool {
	if o != nil && !IsNil(o.Finalizers) {
		return true
	}

	return false
}

// SetFinalizers gets a reference to the given []ApiV1Finalizer and assigns it to the Finalizers field.
func (o *ManagementV1alpha1APIServiceInstance) SetFinalizers(v []ApiV1Finalizer) {
	o.Finalizers = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *ManagementV1alpha1APIServiceInstance) GetAttributes() map[string]string {
	if o == nil || IsNil(o.Attributes) {
		var ret map[string]string
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1APIServiceInstance) GetAttributesOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *ManagementV1alpha1APIServiceInstance) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]string and assigns it to the Attributes field.
func (o *ManagementV1alpha1APIServiceInstance) SetAttributes(v map[string]string) {
	o.Attributes = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ManagementV1alpha1APIServiceInstance) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1APIServiceInstance) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ManagementV1alpha1APIServiceInstance) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *ManagementV1alpha1APIServiceInstance) SetTags(v []string) {
	o.Tags = v
}

// GetSpec returns the Spec field value
func (o *ManagementV1alpha1APIServiceInstance) GetSpec() ManagementV1alpha1APIServiceInstanceSpec {
	if o == nil {
		var ret ManagementV1alpha1APIServiceInstanceSpec
		return ret
	}

	return o.Spec
}

// GetSpecOk returns a tuple with the Spec field value
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1APIServiceInstance) GetSpecOk() (*ManagementV1alpha1APIServiceInstanceSpec, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Spec, true
}

// SetSpec sets field value
func (o *ManagementV1alpha1APIServiceInstance) SetSpec(v ManagementV1alpha1APIServiceInstanceSpec) {
	o.Spec = v
}

// GetLifecycle returns the Lifecycle field value if set, zero value otherwise.
func (o *ManagementV1alpha1APIServiceInstance) GetLifecycle() ManagementV1alpha1APIServiceInstanceLifecycle {
	if o == nil || IsNil(o.Lifecycle) {
		var ret ManagementV1alpha1APIServiceInstanceLifecycle
		return ret
	}
	return *o.Lifecycle
}

// GetLifecycleOk returns a tuple with the Lifecycle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1APIServiceInstance) GetLifecycleOk() (*ManagementV1alpha1APIServiceInstanceLifecycle, bool) {
	if o == nil || IsNil(o.Lifecycle) {
		return nil, false
	}
	return o.Lifecycle, true
}

// HasLifecycle returns a boolean if a field has been set.
func (o *ManagementV1alpha1APIServiceInstance) HasLifecycle() bool {
	if o != nil && !IsNil(o.Lifecycle) {
		return true
	}

	return false
}

// SetLifecycle gets a reference to the given ManagementV1alpha1APIServiceInstanceLifecycle and assigns it to the Lifecycle field.
func (o *ManagementV1alpha1APIServiceInstance) SetLifecycle(v ManagementV1alpha1APIServiceInstanceLifecycle) {
	o.Lifecycle = &v
}

// GetReferences returns the References field value if set, zero value otherwise.
func (o *ManagementV1alpha1APIServiceInstance) GetReferences() ManagementV1alpha1APIServiceInstanceReferences {
	if o == nil || IsNil(o.References) {
		var ret ManagementV1alpha1APIServiceInstanceReferences
		return ret
	}
	return *o.References
}

// GetReferencesOk returns a tuple with the References field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1APIServiceInstance) GetReferencesOk() (*ManagementV1alpha1APIServiceInstanceReferences, bool) {
	if o == nil || IsNil(o.References) {
		return nil, false
	}
	return o.References, true
}

// HasReferences returns a boolean if a field has been set.
func (o *ManagementV1alpha1APIServiceInstance) HasReferences() bool {
	if o != nil && !IsNil(o.References) {
		return true
	}

	return false
}

// SetReferences gets a reference to the given ManagementV1alpha1APIServiceInstanceReferences and assigns it to the References field.
func (o *ManagementV1alpha1APIServiceInstance) SetReferences(v ManagementV1alpha1APIServiceInstanceReferences) {
	o.References = &v
}

// GetCompliance returns the Compliance field value if set, zero value otherwise.
func (o *ManagementV1alpha1APIServiceInstance) GetCompliance() ManagementV1alpha1APIServiceInstanceCompliance {
	if o == nil || IsNil(o.Compliance) {
		var ret ManagementV1alpha1APIServiceInstanceCompliance
		return ret
	}
	return *o.Compliance
}

// GetComplianceOk returns a tuple with the Compliance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1APIServiceInstance) GetComplianceOk() (*ManagementV1alpha1APIServiceInstanceCompliance, bool) {
	if o == nil || IsNil(o.Compliance) {
		return nil, false
	}
	return o.Compliance, true
}

// HasCompliance returns a boolean if a field has been set.
func (o *ManagementV1alpha1APIServiceInstance) HasCompliance() bool {
	if o != nil && !IsNil(o.Compliance) {
		return true
	}

	return false
}

// SetCompliance gets a reference to the given ManagementV1alpha1APIServiceInstanceCompliance and assigns it to the Compliance field.
func (o *ManagementV1alpha1APIServiceInstance) SetCompliance(v ManagementV1alpha1APIServiceInstanceCompliance) {
	o.Compliance = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *ManagementV1alpha1APIServiceInstance) GetSource() ManagementV1alpha1APIServiceInstanceSource {
	if o == nil || IsNil(o.Source) {
		var ret ManagementV1alpha1APIServiceInstanceSource
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1APIServiceInstance) GetSourceOk() (*ManagementV1alpha1APIServiceInstanceSource, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *ManagementV1alpha1APIServiceInstance) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given ManagementV1alpha1APIServiceInstanceSource and assigns it to the Source field.
func (o *ManagementV1alpha1APIServiceInstance) SetSource(v ManagementV1alpha1APIServiceInstanceSource) {
	o.Source = &v
}

// GetLanguages returns the Languages field value if set, zero value otherwise.
func (o *ManagementV1alpha1APIServiceInstance) GetLanguages() ApiV1Languages {
	if o == nil || IsNil(o.Languages) {
		var ret ApiV1Languages
		return ret
	}
	return *o.Languages
}

// GetLanguagesOk returns a tuple with the Languages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1APIServiceInstance) GetLanguagesOk() (*ApiV1Languages, bool) {
	if o == nil || IsNil(o.Languages) {
		return nil, false
	}
	return o.Languages, true
}

// HasLanguages returns a boolean if a field has been set.
func (o *ManagementV1alpha1APIServiceInstance) HasLanguages() bool {
	if o != nil && !IsNil(o.Languages) {
		return true
	}

	return false
}

// SetLanguages gets a reference to the given ApiV1Languages and assigns it to the Languages field.
func (o *ManagementV1alpha1APIServiceInstance) SetLanguages(v ApiV1Languages) {
	o.Languages = &v
}

// GetLanguagesEnUs returns the LanguagesEnUs field value if set, zero value otherwise.
func (o *ManagementV1alpha1APIServiceInstance) GetLanguagesEnUs() ApiV1Language {
	if o == nil || IsNil(o.LanguagesEnUs) {
		var ret ApiV1Language
		return ret
	}
	return *o.LanguagesEnUs
}

// GetLanguagesEnUsOk returns a tuple with the LanguagesEnUs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1APIServiceInstance) GetLanguagesEnUsOk() (*ApiV1Language, bool) {
	if o == nil || IsNil(o.LanguagesEnUs) {
		return nil, false
	}
	return o.LanguagesEnUs, true
}

// HasLanguagesEnUs returns a boolean if a field has been set.
func (o *ManagementV1alpha1APIServiceInstance) HasLanguagesEnUs() bool {
	if o != nil && !IsNil(o.LanguagesEnUs) {
		return true
	}

	return false
}

// SetLanguagesEnUs gets a reference to the given ApiV1Language and assigns it to the LanguagesEnUs field.
func (o *ManagementV1alpha1APIServiceInstance) SetLanguagesEnUs(v ApiV1Language) {
	o.LanguagesEnUs = &v
}

// GetLanguagesPtBr returns the LanguagesPtBr field value if set, zero value otherwise.
func (o *ManagementV1alpha1APIServiceInstance) GetLanguagesPtBr() ApiV1Language {
	if o == nil || IsNil(o.LanguagesPtBr) {
		var ret ApiV1Language
		return ret
	}
	return *o.LanguagesPtBr
}

// GetLanguagesPtBrOk returns a tuple with the LanguagesPtBr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1APIServiceInstance) GetLanguagesPtBrOk() (*ApiV1Language, bool) {
	if o == nil || IsNil(o.LanguagesPtBr) {
		return nil, false
	}
	return o.LanguagesPtBr, true
}

// HasLanguagesPtBr returns a boolean if a field has been set.
func (o *ManagementV1alpha1APIServiceInstance) HasLanguagesPtBr() bool {
	if o != nil && !IsNil(o.LanguagesPtBr) {
		return true
	}

	return false
}

// SetLanguagesPtBr gets a reference to the given ApiV1Language and assigns it to the LanguagesPtBr field.
func (o *ManagementV1alpha1APIServiceInstance) SetLanguagesPtBr(v ApiV1Language) {
	o.LanguagesPtBr = &v
}

// GetLanguagesDeDe returns the LanguagesDeDe field value if set, zero value otherwise.
func (o *ManagementV1alpha1APIServiceInstance) GetLanguagesDeDe() ApiV1Language {
	if o == nil || IsNil(o.LanguagesDeDe) {
		var ret ApiV1Language
		return ret
	}
	return *o.LanguagesDeDe
}

// GetLanguagesDeDeOk returns a tuple with the LanguagesDeDe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1APIServiceInstance) GetLanguagesDeDeOk() (*ApiV1Language, bool) {
	if o == nil || IsNil(o.LanguagesDeDe) {
		return nil, false
	}
	return o.LanguagesDeDe, true
}

// HasLanguagesDeDe returns a boolean if a field has been set.
func (o *ManagementV1alpha1APIServiceInstance) HasLanguagesDeDe() bool {
	if o != nil && !IsNil(o.LanguagesDeDe) {
		return true
	}

	return false
}

// SetLanguagesDeDe gets a reference to the given ApiV1Language and assigns it to the LanguagesDeDe field.
func (o *ManagementV1alpha1APIServiceInstance) SetLanguagesDeDe(v ApiV1Language) {
	o.LanguagesDeDe = &v
}

// GetLanguagesFrFr returns the LanguagesFrFr field value if set, zero value otherwise.
func (o *ManagementV1alpha1APIServiceInstance) GetLanguagesFrFr() ApiV1Language {
	if o == nil || IsNil(o.LanguagesFrFr) {
		var ret ApiV1Language
		return ret
	}
	return *o.LanguagesFrFr
}

// GetLanguagesFrFrOk returns a tuple with the LanguagesFrFr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1APIServiceInstance) GetLanguagesFrFrOk() (*ApiV1Language, bool) {
	if o == nil || IsNil(o.LanguagesFrFr) {
		return nil, false
	}
	return o.LanguagesFrFr, true
}

// HasLanguagesFrFr returns a boolean if a field has been set.
func (o *ManagementV1alpha1APIServiceInstance) HasLanguagesFrFr() bool {
	if o != nil && !IsNil(o.LanguagesFrFr) {
		return true
	}

	return false
}

// SetLanguagesFrFr gets a reference to the given ApiV1Language and assigns it to the LanguagesFrFr field.
func (o *ManagementV1alpha1APIServiceInstance) SetLanguagesFrFr(v ApiV1Language) {
	o.LanguagesFrFr = &v
}

// GetEmbedded returns the Embedded field value if set, zero value otherwise.
func (o *ManagementV1alpha1APIServiceInstance) GetEmbedded() ApiV1Embedded {
	if o == nil || IsNil(o.Embedded) {
		var ret ApiV1Embedded
		return ret
	}
	return *o.Embedded
}

// GetEmbeddedOk returns a tuple with the Embedded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1APIServiceInstance) GetEmbeddedOk() (*ApiV1Embedded, bool) {
	if o == nil || IsNil(o.Embedded) {
		return nil, false
	}
	return o.Embedded, true
}

// HasEmbedded returns a boolean if a field has been set.
func (o *ManagementV1alpha1APIServiceInstance) HasEmbedded() bool {
	if o != nil && !IsNil(o.Embedded) {
		return true
	}

	return false
}

// SetEmbedded gets a reference to the given ApiV1Embedded and assigns it to the Embedded field.
func (o *ManagementV1alpha1APIServiceInstance) SetEmbedded(v ApiV1Embedded) {
	o.Embedded = &v
}

func (o ManagementV1alpha1APIServiceInstance) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManagementV1alpha1APIServiceInstance) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Lifecycle) {
		toSerialize["lifecycle"] = o.Lifecycle
	}
	if !IsNil(o.References) {
		toSerialize["references"] = o.References
	}
	if !IsNil(o.Compliance) {
		toSerialize["compliance"] = o.Compliance
	}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
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

func (o *ManagementV1alpha1APIServiceInstance) UnmarshalJSON(data []byte) (err error) {
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

	varManagementV1alpha1APIServiceInstance := _ManagementV1alpha1APIServiceInstance{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varManagementV1alpha1APIServiceInstance)

	if err != nil {
		return err
	}

	*o = ManagementV1alpha1APIServiceInstance(varManagementV1alpha1APIServiceInstance)

	return err
}

type NullableManagementV1alpha1APIServiceInstance struct {
	value *ManagementV1alpha1APIServiceInstance
	isSet bool
}

func (v NullableManagementV1alpha1APIServiceInstance) Get() *ManagementV1alpha1APIServiceInstance {
	return v.value
}

func (v *NullableManagementV1alpha1APIServiceInstance) Set(val *ManagementV1alpha1APIServiceInstance) {
	v.value = val
	v.isSet = true
}

func (v NullableManagementV1alpha1APIServiceInstance) IsSet() bool {
	return v.isSet
}

func (v *NullableManagementV1alpha1APIServiceInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagementV1alpha1APIServiceInstance(val *ManagementV1alpha1APIServiceInstance) *NullableManagementV1alpha1APIServiceInstance {
	return &NullableManagementV1alpha1APIServiceInstance{value: val, isSet: true}
}

func (v NullableManagementV1alpha1APIServiceInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagementV1alpha1APIServiceInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


