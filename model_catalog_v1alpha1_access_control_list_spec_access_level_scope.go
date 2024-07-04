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

// checks if the CatalogV1alpha1AccessControlListSpecAccessLevelScope type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogV1alpha1AccessControlListSpecAccessLevelScope{}

// CatalogV1alpha1AccessControlListSpecAccessLevelScope struct for CatalogV1alpha1AccessControlListSpecAccessLevelScope
type CatalogV1alpha1AccessControlListSpecAccessLevelScope struct {
	// Resource level at which access is being granted.
	Level *string `json:"level,omitempty"`
	// Set true to allow users to create scoped resources.
	AllowCreateScoped *bool `json:"allowCreateScoped,omitempty"`
	// Set true to allow users to delete the unscoped resource.
	AllowDelete *bool `json:"allowDelete,omitempty"`
	// Set true to allow users to update the unscoped resource.
	AllowWrite *bool `json:"allowWrite,omitempty"`
}

// NewCatalogV1alpha1AccessControlListSpecAccessLevelScope instantiates a new CatalogV1alpha1AccessControlListSpecAccessLevelScope object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogV1alpha1AccessControlListSpecAccessLevelScope() *CatalogV1alpha1AccessControlListSpecAccessLevelScope {
	this := CatalogV1alpha1AccessControlListSpecAccessLevelScope{}
	return &this
}

// NewCatalogV1alpha1AccessControlListSpecAccessLevelScopeWithDefaults instantiates a new CatalogV1alpha1AccessControlListSpecAccessLevelScope object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogV1alpha1AccessControlListSpecAccessLevelScopeWithDefaults() *CatalogV1alpha1AccessControlListSpecAccessLevelScope {
	this := CatalogV1alpha1AccessControlListSpecAccessLevelScope{}
	return &this
}

// GetLevel returns the Level field value if set, zero value otherwise.
func (o *CatalogV1alpha1AccessControlListSpecAccessLevelScope) GetLevel() string {
	if o == nil || IsNil(o.Level) {
		var ret string
		return ret
	}
	return *o.Level
}

// GetLevelOk returns a tuple with the Level field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1AccessControlListSpecAccessLevelScope) GetLevelOk() (*string, bool) {
	if o == nil || IsNil(o.Level) {
		return nil, false
	}
	return o.Level, true
}

// HasLevel returns a boolean if a field has been set.
func (o *CatalogV1alpha1AccessControlListSpecAccessLevelScope) HasLevel() bool {
	if o != nil && !IsNil(o.Level) {
		return true
	}

	return false
}

// SetLevel gets a reference to the given string and assigns it to the Level field.
func (o *CatalogV1alpha1AccessControlListSpecAccessLevelScope) SetLevel(v string) {
	o.Level = &v
}

// GetAllowCreateScoped returns the AllowCreateScoped field value if set, zero value otherwise.
func (o *CatalogV1alpha1AccessControlListSpecAccessLevelScope) GetAllowCreateScoped() bool {
	if o == nil || IsNil(o.AllowCreateScoped) {
		var ret bool
		return ret
	}
	return *o.AllowCreateScoped
}

// GetAllowCreateScopedOk returns a tuple with the AllowCreateScoped field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1AccessControlListSpecAccessLevelScope) GetAllowCreateScopedOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowCreateScoped) {
		return nil, false
	}
	return o.AllowCreateScoped, true
}

// HasAllowCreateScoped returns a boolean if a field has been set.
func (o *CatalogV1alpha1AccessControlListSpecAccessLevelScope) HasAllowCreateScoped() bool {
	if o != nil && !IsNil(o.AllowCreateScoped) {
		return true
	}

	return false
}

// SetAllowCreateScoped gets a reference to the given bool and assigns it to the AllowCreateScoped field.
func (o *CatalogV1alpha1AccessControlListSpecAccessLevelScope) SetAllowCreateScoped(v bool) {
	o.AllowCreateScoped = &v
}

// GetAllowDelete returns the AllowDelete field value if set, zero value otherwise.
func (o *CatalogV1alpha1AccessControlListSpecAccessLevelScope) GetAllowDelete() bool {
	if o == nil || IsNil(o.AllowDelete) {
		var ret bool
		return ret
	}
	return *o.AllowDelete
}

// GetAllowDeleteOk returns a tuple with the AllowDelete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1AccessControlListSpecAccessLevelScope) GetAllowDeleteOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowDelete) {
		return nil, false
	}
	return o.AllowDelete, true
}

// HasAllowDelete returns a boolean if a field has been set.
func (o *CatalogV1alpha1AccessControlListSpecAccessLevelScope) HasAllowDelete() bool {
	if o != nil && !IsNil(o.AllowDelete) {
		return true
	}

	return false
}

// SetAllowDelete gets a reference to the given bool and assigns it to the AllowDelete field.
func (o *CatalogV1alpha1AccessControlListSpecAccessLevelScope) SetAllowDelete(v bool) {
	o.AllowDelete = &v
}

// GetAllowWrite returns the AllowWrite field value if set, zero value otherwise.
func (o *CatalogV1alpha1AccessControlListSpecAccessLevelScope) GetAllowWrite() bool {
	if o == nil || IsNil(o.AllowWrite) {
		var ret bool
		return ret
	}
	return *o.AllowWrite
}

// GetAllowWriteOk returns a tuple with the AllowWrite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1AccessControlListSpecAccessLevelScope) GetAllowWriteOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowWrite) {
		return nil, false
	}
	return o.AllowWrite, true
}

// HasAllowWrite returns a boolean if a field has been set.
func (o *CatalogV1alpha1AccessControlListSpecAccessLevelScope) HasAllowWrite() bool {
	if o != nil && !IsNil(o.AllowWrite) {
		return true
	}

	return false
}

// SetAllowWrite gets a reference to the given bool and assigns it to the AllowWrite field.
func (o *CatalogV1alpha1AccessControlListSpecAccessLevelScope) SetAllowWrite(v bool) {
	o.AllowWrite = &v
}

func (o CatalogV1alpha1AccessControlListSpecAccessLevelScope) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogV1alpha1AccessControlListSpecAccessLevelScope) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Level) {
		toSerialize["level"] = o.Level
	}
	if !IsNil(o.AllowCreateScoped) {
		toSerialize["allowCreateScoped"] = o.AllowCreateScoped
	}
	if !IsNil(o.AllowDelete) {
		toSerialize["allowDelete"] = o.AllowDelete
	}
	if !IsNil(o.AllowWrite) {
		toSerialize["allowWrite"] = o.AllowWrite
	}
	return toSerialize, nil
}

type NullableCatalogV1alpha1AccessControlListSpecAccessLevelScope struct {
	value *CatalogV1alpha1AccessControlListSpecAccessLevelScope
	isSet bool
}

func (v NullableCatalogV1alpha1AccessControlListSpecAccessLevelScope) Get() *CatalogV1alpha1AccessControlListSpecAccessLevelScope {
	return v.value
}

func (v *NullableCatalogV1alpha1AccessControlListSpecAccessLevelScope) Set(val *CatalogV1alpha1AccessControlListSpecAccessLevelScope) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogV1alpha1AccessControlListSpecAccessLevelScope) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogV1alpha1AccessControlListSpecAccessLevelScope) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogV1alpha1AccessControlListSpecAccessLevelScope(val *CatalogV1alpha1AccessControlListSpecAccessLevelScope) *NullableCatalogV1alpha1AccessControlListSpecAccessLevelScope {
	return &NullableCatalogV1alpha1AccessControlListSpecAccessLevelScope{value: val, isSet: true}
}

func (v NullableCatalogV1alpha1AccessControlListSpecAccessLevelScope) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogV1alpha1AccessControlListSpecAccessLevelScope) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


