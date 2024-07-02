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

// checks if the ManagementV1APISpecLintingJob type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManagementV1APISpecLintingJob{}

// ManagementV1APISpecLintingJob Definition of APISpecLintingJob for version v1 in group management.    Resource representing an Amplify Central API Specification Linting Job
type ManagementV1APISpecLintingJob struct {
	// Resource belongs to the management group. Cannot be updated. The server infers this from the endpoint the client submits the request to.
	Group *string `json:"group,omitempty" validate:"regexp=^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\\\\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*"`
	// Resource version is v1. The version defines the structure of the resource. Cannot be updated. The server infers this from the endpoint the client submits the request to.
	ApiVersion *string `json:"apiVersion,omitempty"`
	// Resource of kind APISpecLintingJob. Cannot be updated. The server infers this from the endpoint the client submits the request to.
	Kind *string `json:"kind,omitempty"`
	// Name of the APISpecLintingJob. APISpecLintingJob name is unique and cannot be updated. The characters allowed in names are: digits (0-9), lower case letters (a-z), -, and .
	Name *string `json:"name,omitempty" validate:"regexp=^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\\\\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*"`
	// The friendly name of the APISpecLintingJob.
	Title *string `json:"title,omitempty"`
	Metadata *ApiV1Metadata `json:"metadata,omitempty"`
	Owner *ApiV1Owner `json:"owner,omitempty"`
	// List of finalizers
	Finalizers []ApiV1Finalizer `json:"finalizers,omitempty"`
	// Custom attributes added to objects.
	Attributes *map[string]string `json:"attributes,omitempty"`
	// List of tags.
	Tags []string `json:"tags,omitempty"`
	Spec ManagementV1APISpecLintingJobSpec `json:"spec"`
	// Set true if this lint job and its result has been archived.
	Archived *bool `json:"archived,omitempty"`
	Result *ManagementV1APISpecLintingJobResult `json:"result,omitempty"`
	State *ManagementV1APISpecLintingJobState `json:"state,omitempty"`
	Languages *ApiV1Languages `json:"languages,omitempty"`
	LanguagesEnUs *ApiV1Language `json:"languages-en-us,omitempty"`
	LanguagesPtBr *ApiV1Language `json:"languages-pt-br,omitempty"`
	LanguagesDeDe *ApiV1Language `json:"languages-de-de,omitempty"`
	LanguagesFrFr *ApiV1Language `json:"languages-fr-fr,omitempty"`
	Embedded *ApiV1Embedded `json:"_embedded,omitempty"`
}

type _ManagementV1APISpecLintingJob ManagementV1APISpecLintingJob

// NewManagementV1APISpecLintingJob instantiates a new ManagementV1APISpecLintingJob object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagementV1APISpecLintingJob(spec ManagementV1APISpecLintingJobSpec) *ManagementV1APISpecLintingJob {
	this := ManagementV1APISpecLintingJob{}
	this.Spec = spec
	return &this
}

// NewManagementV1APISpecLintingJobWithDefaults instantiates a new ManagementV1APISpecLintingJob object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagementV1APISpecLintingJobWithDefaults() *ManagementV1APISpecLintingJob {
	this := ManagementV1APISpecLintingJob{}
	return &this
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *ManagementV1APISpecLintingJob) GetGroup() string {
	if o == nil || IsNil(o.Group) {
		var ret string
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1APISpecLintingJob) GetGroupOk() (*string, bool) {
	if o == nil || IsNil(o.Group) {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *ManagementV1APISpecLintingJob) HasGroup() bool {
	if o != nil && !IsNil(o.Group) {
		return true
	}

	return false
}

// SetGroup gets a reference to the given string and assigns it to the Group field.
func (o *ManagementV1APISpecLintingJob) SetGroup(v string) {
	o.Group = &v
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *ManagementV1APISpecLintingJob) GetApiVersion() string {
	if o == nil || IsNil(o.ApiVersion) {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1APISpecLintingJob) GetApiVersionOk() (*string, bool) {
	if o == nil || IsNil(o.ApiVersion) {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *ManagementV1APISpecLintingJob) HasApiVersion() bool {
	if o != nil && !IsNil(o.ApiVersion) {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *ManagementV1APISpecLintingJob) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *ManagementV1APISpecLintingJob) GetKind() string {
	if o == nil || IsNil(o.Kind) {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1APISpecLintingJob) GetKindOk() (*string, bool) {
	if o == nil || IsNil(o.Kind) {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *ManagementV1APISpecLintingJob) HasKind() bool {
	if o != nil && !IsNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *ManagementV1APISpecLintingJob) SetKind(v string) {
	o.Kind = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ManagementV1APISpecLintingJob) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1APISpecLintingJob) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ManagementV1APISpecLintingJob) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ManagementV1APISpecLintingJob) SetName(v string) {
	o.Name = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *ManagementV1APISpecLintingJob) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1APISpecLintingJob) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ManagementV1APISpecLintingJob) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *ManagementV1APISpecLintingJob) SetTitle(v string) {
	o.Title = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *ManagementV1APISpecLintingJob) GetMetadata() ApiV1Metadata {
	if o == nil || IsNil(o.Metadata) {
		var ret ApiV1Metadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1APISpecLintingJob) GetMetadataOk() (*ApiV1Metadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *ManagementV1APISpecLintingJob) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given ApiV1Metadata and assigns it to the Metadata field.
func (o *ManagementV1APISpecLintingJob) SetMetadata(v ApiV1Metadata) {
	o.Metadata = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *ManagementV1APISpecLintingJob) GetOwner() ApiV1Owner {
	if o == nil || IsNil(o.Owner) {
		var ret ApiV1Owner
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1APISpecLintingJob) GetOwnerOk() (*ApiV1Owner, bool) {
	if o == nil || IsNil(o.Owner) {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *ManagementV1APISpecLintingJob) HasOwner() bool {
	if o != nil && !IsNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given ApiV1Owner and assigns it to the Owner field.
func (o *ManagementV1APISpecLintingJob) SetOwner(v ApiV1Owner) {
	o.Owner = &v
}

// GetFinalizers returns the Finalizers field value if set, zero value otherwise.
func (o *ManagementV1APISpecLintingJob) GetFinalizers() []ApiV1Finalizer {
	if o == nil || IsNil(o.Finalizers) {
		var ret []ApiV1Finalizer
		return ret
	}
	return o.Finalizers
}

// GetFinalizersOk returns a tuple with the Finalizers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1APISpecLintingJob) GetFinalizersOk() ([]ApiV1Finalizer, bool) {
	if o == nil || IsNil(o.Finalizers) {
		return nil, false
	}
	return o.Finalizers, true
}

// HasFinalizers returns a boolean if a field has been set.
func (o *ManagementV1APISpecLintingJob) HasFinalizers() bool {
	if o != nil && !IsNil(o.Finalizers) {
		return true
	}

	return false
}

// SetFinalizers gets a reference to the given []ApiV1Finalizer and assigns it to the Finalizers field.
func (o *ManagementV1APISpecLintingJob) SetFinalizers(v []ApiV1Finalizer) {
	o.Finalizers = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *ManagementV1APISpecLintingJob) GetAttributes() map[string]string {
	if o == nil || IsNil(o.Attributes) {
		var ret map[string]string
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1APISpecLintingJob) GetAttributesOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *ManagementV1APISpecLintingJob) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]string and assigns it to the Attributes field.
func (o *ManagementV1APISpecLintingJob) SetAttributes(v map[string]string) {
	o.Attributes = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ManagementV1APISpecLintingJob) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1APISpecLintingJob) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ManagementV1APISpecLintingJob) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *ManagementV1APISpecLintingJob) SetTags(v []string) {
	o.Tags = v
}

// GetSpec returns the Spec field value
func (o *ManagementV1APISpecLintingJob) GetSpec() ManagementV1APISpecLintingJobSpec {
	if o == nil {
		var ret ManagementV1APISpecLintingJobSpec
		return ret
	}

	return o.Spec
}

// GetSpecOk returns a tuple with the Spec field value
// and a boolean to check if the value has been set.
func (o *ManagementV1APISpecLintingJob) GetSpecOk() (*ManagementV1APISpecLintingJobSpec, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Spec, true
}

// SetSpec sets field value
func (o *ManagementV1APISpecLintingJob) SetSpec(v ManagementV1APISpecLintingJobSpec) {
	o.Spec = v
}

// GetArchived returns the Archived field value if set, zero value otherwise.
func (o *ManagementV1APISpecLintingJob) GetArchived() bool {
	if o == nil || IsNil(o.Archived) {
		var ret bool
		return ret
	}
	return *o.Archived
}

// GetArchivedOk returns a tuple with the Archived field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1APISpecLintingJob) GetArchivedOk() (*bool, bool) {
	if o == nil || IsNil(o.Archived) {
		return nil, false
	}
	return o.Archived, true
}

// HasArchived returns a boolean if a field has been set.
func (o *ManagementV1APISpecLintingJob) HasArchived() bool {
	if o != nil && !IsNil(o.Archived) {
		return true
	}

	return false
}

// SetArchived gets a reference to the given bool and assigns it to the Archived field.
func (o *ManagementV1APISpecLintingJob) SetArchived(v bool) {
	o.Archived = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *ManagementV1APISpecLintingJob) GetResult() ManagementV1APISpecLintingJobResult {
	if o == nil || IsNil(o.Result) {
		var ret ManagementV1APISpecLintingJobResult
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1APISpecLintingJob) GetResultOk() (*ManagementV1APISpecLintingJobResult, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *ManagementV1APISpecLintingJob) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given ManagementV1APISpecLintingJobResult and assigns it to the Result field.
func (o *ManagementV1APISpecLintingJob) SetResult(v ManagementV1APISpecLintingJobResult) {
	o.Result = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *ManagementV1APISpecLintingJob) GetState() ManagementV1APISpecLintingJobState {
	if o == nil || IsNil(o.State) {
		var ret ManagementV1APISpecLintingJobState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1APISpecLintingJob) GetStateOk() (*ManagementV1APISpecLintingJobState, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *ManagementV1APISpecLintingJob) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given ManagementV1APISpecLintingJobState and assigns it to the State field.
func (o *ManagementV1APISpecLintingJob) SetState(v ManagementV1APISpecLintingJobState) {
	o.State = &v
}

// GetLanguages returns the Languages field value if set, zero value otherwise.
func (o *ManagementV1APISpecLintingJob) GetLanguages() ApiV1Languages {
	if o == nil || IsNil(o.Languages) {
		var ret ApiV1Languages
		return ret
	}
	return *o.Languages
}

// GetLanguagesOk returns a tuple with the Languages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1APISpecLintingJob) GetLanguagesOk() (*ApiV1Languages, bool) {
	if o == nil || IsNil(o.Languages) {
		return nil, false
	}
	return o.Languages, true
}

// HasLanguages returns a boolean if a field has been set.
func (o *ManagementV1APISpecLintingJob) HasLanguages() bool {
	if o != nil && !IsNil(o.Languages) {
		return true
	}

	return false
}

// SetLanguages gets a reference to the given ApiV1Languages and assigns it to the Languages field.
func (o *ManagementV1APISpecLintingJob) SetLanguages(v ApiV1Languages) {
	o.Languages = &v
}

// GetLanguagesEnUs returns the LanguagesEnUs field value if set, zero value otherwise.
func (o *ManagementV1APISpecLintingJob) GetLanguagesEnUs() ApiV1Language {
	if o == nil || IsNil(o.LanguagesEnUs) {
		var ret ApiV1Language
		return ret
	}
	return *o.LanguagesEnUs
}

// GetLanguagesEnUsOk returns a tuple with the LanguagesEnUs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1APISpecLintingJob) GetLanguagesEnUsOk() (*ApiV1Language, bool) {
	if o == nil || IsNil(o.LanguagesEnUs) {
		return nil, false
	}
	return o.LanguagesEnUs, true
}

// HasLanguagesEnUs returns a boolean if a field has been set.
func (o *ManagementV1APISpecLintingJob) HasLanguagesEnUs() bool {
	if o != nil && !IsNil(o.LanguagesEnUs) {
		return true
	}

	return false
}

// SetLanguagesEnUs gets a reference to the given ApiV1Language and assigns it to the LanguagesEnUs field.
func (o *ManagementV1APISpecLintingJob) SetLanguagesEnUs(v ApiV1Language) {
	o.LanguagesEnUs = &v
}

// GetLanguagesPtBr returns the LanguagesPtBr field value if set, zero value otherwise.
func (o *ManagementV1APISpecLintingJob) GetLanguagesPtBr() ApiV1Language {
	if o == nil || IsNil(o.LanguagesPtBr) {
		var ret ApiV1Language
		return ret
	}
	return *o.LanguagesPtBr
}

// GetLanguagesPtBrOk returns a tuple with the LanguagesPtBr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1APISpecLintingJob) GetLanguagesPtBrOk() (*ApiV1Language, bool) {
	if o == nil || IsNil(o.LanguagesPtBr) {
		return nil, false
	}
	return o.LanguagesPtBr, true
}

// HasLanguagesPtBr returns a boolean if a field has been set.
func (o *ManagementV1APISpecLintingJob) HasLanguagesPtBr() bool {
	if o != nil && !IsNil(o.LanguagesPtBr) {
		return true
	}

	return false
}

// SetLanguagesPtBr gets a reference to the given ApiV1Language and assigns it to the LanguagesPtBr field.
func (o *ManagementV1APISpecLintingJob) SetLanguagesPtBr(v ApiV1Language) {
	o.LanguagesPtBr = &v
}

// GetLanguagesDeDe returns the LanguagesDeDe field value if set, zero value otherwise.
func (o *ManagementV1APISpecLintingJob) GetLanguagesDeDe() ApiV1Language {
	if o == nil || IsNil(o.LanguagesDeDe) {
		var ret ApiV1Language
		return ret
	}
	return *o.LanguagesDeDe
}

// GetLanguagesDeDeOk returns a tuple with the LanguagesDeDe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1APISpecLintingJob) GetLanguagesDeDeOk() (*ApiV1Language, bool) {
	if o == nil || IsNil(o.LanguagesDeDe) {
		return nil, false
	}
	return o.LanguagesDeDe, true
}

// HasLanguagesDeDe returns a boolean if a field has been set.
func (o *ManagementV1APISpecLintingJob) HasLanguagesDeDe() bool {
	if o != nil && !IsNil(o.LanguagesDeDe) {
		return true
	}

	return false
}

// SetLanguagesDeDe gets a reference to the given ApiV1Language and assigns it to the LanguagesDeDe field.
func (o *ManagementV1APISpecLintingJob) SetLanguagesDeDe(v ApiV1Language) {
	o.LanguagesDeDe = &v
}

// GetLanguagesFrFr returns the LanguagesFrFr field value if set, zero value otherwise.
func (o *ManagementV1APISpecLintingJob) GetLanguagesFrFr() ApiV1Language {
	if o == nil || IsNil(o.LanguagesFrFr) {
		var ret ApiV1Language
		return ret
	}
	return *o.LanguagesFrFr
}

// GetLanguagesFrFrOk returns a tuple with the LanguagesFrFr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1APISpecLintingJob) GetLanguagesFrFrOk() (*ApiV1Language, bool) {
	if o == nil || IsNil(o.LanguagesFrFr) {
		return nil, false
	}
	return o.LanguagesFrFr, true
}

// HasLanguagesFrFr returns a boolean if a field has been set.
func (o *ManagementV1APISpecLintingJob) HasLanguagesFrFr() bool {
	if o != nil && !IsNil(o.LanguagesFrFr) {
		return true
	}

	return false
}

// SetLanguagesFrFr gets a reference to the given ApiV1Language and assigns it to the LanguagesFrFr field.
func (o *ManagementV1APISpecLintingJob) SetLanguagesFrFr(v ApiV1Language) {
	o.LanguagesFrFr = &v
}

// GetEmbedded returns the Embedded field value if set, zero value otherwise.
func (o *ManagementV1APISpecLintingJob) GetEmbedded() ApiV1Embedded {
	if o == nil || IsNil(o.Embedded) {
		var ret ApiV1Embedded
		return ret
	}
	return *o.Embedded
}

// GetEmbeddedOk returns a tuple with the Embedded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1APISpecLintingJob) GetEmbeddedOk() (*ApiV1Embedded, bool) {
	if o == nil || IsNil(o.Embedded) {
		return nil, false
	}
	return o.Embedded, true
}

// HasEmbedded returns a boolean if a field has been set.
func (o *ManagementV1APISpecLintingJob) HasEmbedded() bool {
	if o != nil && !IsNil(o.Embedded) {
		return true
	}

	return false
}

// SetEmbedded gets a reference to the given ApiV1Embedded and assigns it to the Embedded field.
func (o *ManagementV1APISpecLintingJob) SetEmbedded(v ApiV1Embedded) {
	o.Embedded = &v
}

func (o ManagementV1APISpecLintingJob) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManagementV1APISpecLintingJob) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Archived) {
		toSerialize["archived"] = o.Archived
	}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
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

func (o *ManagementV1APISpecLintingJob) UnmarshalJSON(data []byte) (err error) {
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

	varManagementV1APISpecLintingJob := _ManagementV1APISpecLintingJob{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varManagementV1APISpecLintingJob)

	if err != nil {
		return err
	}

	*o = ManagementV1APISpecLintingJob(varManagementV1APISpecLintingJob)

	return err
}

type NullableManagementV1APISpecLintingJob struct {
	value *ManagementV1APISpecLintingJob
	isSet bool
}

func (v NullableManagementV1APISpecLintingJob) Get() *ManagementV1APISpecLintingJob {
	return v.value
}

func (v *NullableManagementV1APISpecLintingJob) Set(val *ManagementV1APISpecLintingJob) {
	v.value = val
	v.isSet = true
}

func (v NullableManagementV1APISpecLintingJob) IsSet() bool {
	return v.isSet
}

func (v *NullableManagementV1APISpecLintingJob) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagementV1APISpecLintingJob(val *ManagementV1APISpecLintingJob) *NullableManagementV1APISpecLintingJob {
	return &NullableManagementV1APISpecLintingJob{value: val, isSet: true}
}

func (v NullableManagementV1APISpecLintingJob) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagementV1APISpecLintingJob) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


