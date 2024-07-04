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

// checks if the ApiV1EmbeddedReferenceMetadataAccessRights type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiV1EmbeddedReferenceMetadataAccessRights{}

// ApiV1EmbeddedReferenceMetadataAccessRights struct for ApiV1EmbeddedReferenceMetadataAccessRights
type ApiV1EmbeddedReferenceMetadataAccessRights struct {
	// Defines if the current user can change the owner of the resource.
	CanChangeOwner *bool `json:"canChangeOwner,omitempty"`
	// Defines if the current user can delete the resource.
	CanDelete *bool `json:"canDelete,omitempty"`
	// Defines if the current user can update the resource.
	CanWrite *bool `json:"canWrite,omitempty"`
	// Defines if the current user can read the resource.
	CanRead *bool `json:"canRead,omitempty"`
}

// NewApiV1EmbeddedReferenceMetadataAccessRights instantiates a new ApiV1EmbeddedReferenceMetadataAccessRights object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiV1EmbeddedReferenceMetadataAccessRights() *ApiV1EmbeddedReferenceMetadataAccessRights {
	this := ApiV1EmbeddedReferenceMetadataAccessRights{}
	return &this
}

// NewApiV1EmbeddedReferenceMetadataAccessRightsWithDefaults instantiates a new ApiV1EmbeddedReferenceMetadataAccessRights object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiV1EmbeddedReferenceMetadataAccessRightsWithDefaults() *ApiV1EmbeddedReferenceMetadataAccessRights {
	this := ApiV1EmbeddedReferenceMetadataAccessRights{}
	return &this
}

// GetCanChangeOwner returns the CanChangeOwner field value if set, zero value otherwise.
func (o *ApiV1EmbeddedReferenceMetadataAccessRights) GetCanChangeOwner() bool {
	if o == nil || IsNil(o.CanChangeOwner) {
		var ret bool
		return ret
	}
	return *o.CanChangeOwner
}

// GetCanChangeOwnerOk returns a tuple with the CanChangeOwner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiV1EmbeddedReferenceMetadataAccessRights) GetCanChangeOwnerOk() (*bool, bool) {
	if o == nil || IsNil(o.CanChangeOwner) {
		return nil, false
	}
	return o.CanChangeOwner, true
}

// HasCanChangeOwner returns a boolean if a field has been set.
func (o *ApiV1EmbeddedReferenceMetadataAccessRights) HasCanChangeOwner() bool {
	if o != nil && !IsNil(o.CanChangeOwner) {
		return true
	}

	return false
}

// SetCanChangeOwner gets a reference to the given bool and assigns it to the CanChangeOwner field.
func (o *ApiV1EmbeddedReferenceMetadataAccessRights) SetCanChangeOwner(v bool) {
	o.CanChangeOwner = &v
}

// GetCanDelete returns the CanDelete field value if set, zero value otherwise.
func (o *ApiV1EmbeddedReferenceMetadataAccessRights) GetCanDelete() bool {
	if o == nil || IsNil(o.CanDelete) {
		var ret bool
		return ret
	}
	return *o.CanDelete
}

// GetCanDeleteOk returns a tuple with the CanDelete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiV1EmbeddedReferenceMetadataAccessRights) GetCanDeleteOk() (*bool, bool) {
	if o == nil || IsNil(o.CanDelete) {
		return nil, false
	}
	return o.CanDelete, true
}

// HasCanDelete returns a boolean if a field has been set.
func (o *ApiV1EmbeddedReferenceMetadataAccessRights) HasCanDelete() bool {
	if o != nil && !IsNil(o.CanDelete) {
		return true
	}

	return false
}

// SetCanDelete gets a reference to the given bool and assigns it to the CanDelete field.
func (o *ApiV1EmbeddedReferenceMetadataAccessRights) SetCanDelete(v bool) {
	o.CanDelete = &v
}

// GetCanWrite returns the CanWrite field value if set, zero value otherwise.
func (o *ApiV1EmbeddedReferenceMetadataAccessRights) GetCanWrite() bool {
	if o == nil || IsNil(o.CanWrite) {
		var ret bool
		return ret
	}
	return *o.CanWrite
}

// GetCanWriteOk returns a tuple with the CanWrite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiV1EmbeddedReferenceMetadataAccessRights) GetCanWriteOk() (*bool, bool) {
	if o == nil || IsNil(o.CanWrite) {
		return nil, false
	}
	return o.CanWrite, true
}

// HasCanWrite returns a boolean if a field has been set.
func (o *ApiV1EmbeddedReferenceMetadataAccessRights) HasCanWrite() bool {
	if o != nil && !IsNil(o.CanWrite) {
		return true
	}

	return false
}

// SetCanWrite gets a reference to the given bool and assigns it to the CanWrite field.
func (o *ApiV1EmbeddedReferenceMetadataAccessRights) SetCanWrite(v bool) {
	o.CanWrite = &v
}

// GetCanRead returns the CanRead field value if set, zero value otherwise.
func (o *ApiV1EmbeddedReferenceMetadataAccessRights) GetCanRead() bool {
	if o == nil || IsNil(o.CanRead) {
		var ret bool
		return ret
	}
	return *o.CanRead
}

// GetCanReadOk returns a tuple with the CanRead field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiV1EmbeddedReferenceMetadataAccessRights) GetCanReadOk() (*bool, bool) {
	if o == nil || IsNil(o.CanRead) {
		return nil, false
	}
	return o.CanRead, true
}

// HasCanRead returns a boolean if a field has been set.
func (o *ApiV1EmbeddedReferenceMetadataAccessRights) HasCanRead() bool {
	if o != nil && !IsNil(o.CanRead) {
		return true
	}

	return false
}

// SetCanRead gets a reference to the given bool and assigns it to the CanRead field.
func (o *ApiV1EmbeddedReferenceMetadataAccessRights) SetCanRead(v bool) {
	o.CanRead = &v
}

func (o ApiV1EmbeddedReferenceMetadataAccessRights) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiV1EmbeddedReferenceMetadataAccessRights) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CanChangeOwner) {
		toSerialize["canChangeOwner"] = o.CanChangeOwner
	}
	if !IsNil(o.CanDelete) {
		toSerialize["canDelete"] = o.CanDelete
	}
	if !IsNil(o.CanWrite) {
		toSerialize["canWrite"] = o.CanWrite
	}
	if !IsNil(o.CanRead) {
		toSerialize["canRead"] = o.CanRead
	}
	return toSerialize, nil
}

type NullableApiV1EmbeddedReferenceMetadataAccessRights struct {
	value *ApiV1EmbeddedReferenceMetadataAccessRights
	isSet bool
}

func (v NullableApiV1EmbeddedReferenceMetadataAccessRights) Get() *ApiV1EmbeddedReferenceMetadataAccessRights {
	return v.value
}

func (v *NullableApiV1EmbeddedReferenceMetadataAccessRights) Set(val *ApiV1EmbeddedReferenceMetadataAccessRights) {
	v.value = val
	v.isSet = true
}

func (v NullableApiV1EmbeddedReferenceMetadataAccessRights) IsSet() bool {
	return v.isSet
}

func (v *NullableApiV1EmbeddedReferenceMetadataAccessRights) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiV1EmbeddedReferenceMetadataAccessRights(val *ApiV1EmbeddedReferenceMetadataAccessRights) *NullableApiV1EmbeddedReferenceMetadataAccessRights {
	return &NullableApiV1EmbeddedReferenceMetadataAccessRights{value: val, isSet: true}
}

func (v NullableApiV1EmbeddedReferenceMetadataAccessRights) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiV1EmbeddedReferenceMetadataAccessRights) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


