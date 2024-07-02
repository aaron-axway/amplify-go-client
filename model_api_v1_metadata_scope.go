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

// checks if the ApiV1MetadataScope type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiV1MetadataScope{}

// ApiV1MetadataScope The scope where this resource was defined.
type ApiV1MetadataScope struct {
	// Internal id of the scope resource where the resource is defined.
	Id *string `json:"id,omitempty"`
	// The kind of the scope resource where the resource is defined.
	Kind *string `json:"kind,omitempty"`
	// The name of the scope where the resource is defined.
	Name *string `json:"name,omitempty" validate:"regexp=^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\\\\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*"`
	// The title of the scope where the resource is defined.
	Title *string `json:"title,omitempty"`
	// The URL representing the scope resource object.
	SelfLink *string `json:"selfLink,omitempty"`
	Owner *ApiV1Owner `json:"owner,omitempty"`
}

// NewApiV1MetadataScope instantiates a new ApiV1MetadataScope object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiV1MetadataScope() *ApiV1MetadataScope {
	this := ApiV1MetadataScope{}
	return &this
}

// NewApiV1MetadataScopeWithDefaults instantiates a new ApiV1MetadataScope object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiV1MetadataScopeWithDefaults() *ApiV1MetadataScope {
	this := ApiV1MetadataScope{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApiV1MetadataScope) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiV1MetadataScope) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApiV1MetadataScope) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ApiV1MetadataScope) SetId(v string) {
	o.Id = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *ApiV1MetadataScope) GetKind() string {
	if o == nil || IsNil(o.Kind) {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiV1MetadataScope) GetKindOk() (*string, bool) {
	if o == nil || IsNil(o.Kind) {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *ApiV1MetadataScope) HasKind() bool {
	if o != nil && !IsNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *ApiV1MetadataScope) SetKind(v string) {
	o.Kind = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ApiV1MetadataScope) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiV1MetadataScope) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ApiV1MetadataScope) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ApiV1MetadataScope) SetName(v string) {
	o.Name = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *ApiV1MetadataScope) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiV1MetadataScope) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ApiV1MetadataScope) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *ApiV1MetadataScope) SetTitle(v string) {
	o.Title = &v
}

// GetSelfLink returns the SelfLink field value if set, zero value otherwise.
func (o *ApiV1MetadataScope) GetSelfLink() string {
	if o == nil || IsNil(o.SelfLink) {
		var ret string
		return ret
	}
	return *o.SelfLink
}

// GetSelfLinkOk returns a tuple with the SelfLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiV1MetadataScope) GetSelfLinkOk() (*string, bool) {
	if o == nil || IsNil(o.SelfLink) {
		return nil, false
	}
	return o.SelfLink, true
}

// HasSelfLink returns a boolean if a field has been set.
func (o *ApiV1MetadataScope) HasSelfLink() bool {
	if o != nil && !IsNil(o.SelfLink) {
		return true
	}

	return false
}

// SetSelfLink gets a reference to the given string and assigns it to the SelfLink field.
func (o *ApiV1MetadataScope) SetSelfLink(v string) {
	o.SelfLink = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *ApiV1MetadataScope) GetOwner() ApiV1Owner {
	if o == nil || IsNil(o.Owner) {
		var ret ApiV1Owner
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiV1MetadataScope) GetOwnerOk() (*ApiV1Owner, bool) {
	if o == nil || IsNil(o.Owner) {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *ApiV1MetadataScope) HasOwner() bool {
	if o != nil && !IsNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given ApiV1Owner and assigns it to the Owner field.
func (o *ApiV1MetadataScope) SetOwner(v ApiV1Owner) {
	o.Owner = &v
}

func (o ApiV1MetadataScope) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiV1MetadataScope) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
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
	if !IsNil(o.SelfLink) {
		toSerialize["selfLink"] = o.SelfLink
	}
	if !IsNil(o.Owner) {
		toSerialize["owner"] = o.Owner
	}
	return toSerialize, nil
}

type NullableApiV1MetadataScope struct {
	value *ApiV1MetadataScope
	isSet bool
}

func (v NullableApiV1MetadataScope) Get() *ApiV1MetadataScope {
	return v.value
}

func (v *NullableApiV1MetadataScope) Set(val *ApiV1MetadataScope) {
	v.value = val
	v.isSet = true
}

func (v NullableApiV1MetadataScope) IsSet() bool {
	return v.isSet
}

func (v *NullableApiV1MetadataScope) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiV1MetadataScope(val *ApiV1MetadataScope) *NullableApiV1MetadataScope {
	return &NullableApiV1MetadataScope{value: val, isSet: true}
}

func (v NullableApiV1MetadataScope) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiV1MetadataScope) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


