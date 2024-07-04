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

// checks if the CatalogV1alpha1SubscriptionMarketplaceResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogV1alpha1SubscriptionMarketplaceResource{}

// CatalogV1alpha1SubscriptionMarketplaceResource The Marketplace Subscription resource details.
type CatalogV1alpha1SubscriptionMarketplaceResource struct {
	Metadata CatalogV1alpha1SubscriptionMarketplaceResourceMetadata `json:"metadata"`
	Owner *CatalogV1alpha1SubscriptionMarketplaceResourceOwner `json:"owner,omitempty"`
}

type _CatalogV1alpha1SubscriptionMarketplaceResource CatalogV1alpha1SubscriptionMarketplaceResource

// NewCatalogV1alpha1SubscriptionMarketplaceResource instantiates a new CatalogV1alpha1SubscriptionMarketplaceResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogV1alpha1SubscriptionMarketplaceResource(metadata CatalogV1alpha1SubscriptionMarketplaceResourceMetadata) *CatalogV1alpha1SubscriptionMarketplaceResource {
	this := CatalogV1alpha1SubscriptionMarketplaceResource{}
	this.Metadata = metadata
	return &this
}

// NewCatalogV1alpha1SubscriptionMarketplaceResourceWithDefaults instantiates a new CatalogV1alpha1SubscriptionMarketplaceResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogV1alpha1SubscriptionMarketplaceResourceWithDefaults() *CatalogV1alpha1SubscriptionMarketplaceResource {
	this := CatalogV1alpha1SubscriptionMarketplaceResource{}
	return &this
}

// GetMetadata returns the Metadata field value
func (o *CatalogV1alpha1SubscriptionMarketplaceResource) GetMetadata() CatalogV1alpha1SubscriptionMarketplaceResourceMetadata {
	if o == nil {
		var ret CatalogV1alpha1SubscriptionMarketplaceResourceMetadata
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1SubscriptionMarketplaceResource) GetMetadataOk() (*CatalogV1alpha1SubscriptionMarketplaceResourceMetadata, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *CatalogV1alpha1SubscriptionMarketplaceResource) SetMetadata(v CatalogV1alpha1SubscriptionMarketplaceResourceMetadata) {
	o.Metadata = v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *CatalogV1alpha1SubscriptionMarketplaceResource) GetOwner() CatalogV1alpha1SubscriptionMarketplaceResourceOwner {
	if o == nil || IsNil(o.Owner) {
		var ret CatalogV1alpha1SubscriptionMarketplaceResourceOwner
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogV1alpha1SubscriptionMarketplaceResource) GetOwnerOk() (*CatalogV1alpha1SubscriptionMarketplaceResourceOwner, bool) {
	if o == nil || IsNil(o.Owner) {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *CatalogV1alpha1SubscriptionMarketplaceResource) HasOwner() bool {
	if o != nil && !IsNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given CatalogV1alpha1SubscriptionMarketplaceResourceOwner and assigns it to the Owner field.
func (o *CatalogV1alpha1SubscriptionMarketplaceResource) SetOwner(v CatalogV1alpha1SubscriptionMarketplaceResourceOwner) {
	o.Owner = &v
}

func (o CatalogV1alpha1SubscriptionMarketplaceResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogV1alpha1SubscriptionMarketplaceResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["metadata"] = o.Metadata
	if !IsNil(o.Owner) {
		toSerialize["owner"] = o.Owner
	}
	return toSerialize, nil
}

func (o *CatalogV1alpha1SubscriptionMarketplaceResource) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"metadata",
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

	varCatalogV1alpha1SubscriptionMarketplaceResource := _CatalogV1alpha1SubscriptionMarketplaceResource{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCatalogV1alpha1SubscriptionMarketplaceResource)

	if err != nil {
		return err
	}

	*o = CatalogV1alpha1SubscriptionMarketplaceResource(varCatalogV1alpha1SubscriptionMarketplaceResource)

	return err
}

type NullableCatalogV1alpha1SubscriptionMarketplaceResource struct {
	value *CatalogV1alpha1SubscriptionMarketplaceResource
	isSet bool
}

func (v NullableCatalogV1alpha1SubscriptionMarketplaceResource) Get() *CatalogV1alpha1SubscriptionMarketplaceResource {
	return v.value
}

func (v *NullableCatalogV1alpha1SubscriptionMarketplaceResource) Set(val *CatalogV1alpha1SubscriptionMarketplaceResource) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogV1alpha1SubscriptionMarketplaceResource) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogV1alpha1SubscriptionMarketplaceResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogV1alpha1SubscriptionMarketplaceResource(val *CatalogV1alpha1SubscriptionMarketplaceResource) *NullableCatalogV1alpha1SubscriptionMarketplaceResource {
	return &NullableCatalogV1alpha1SubscriptionMarketplaceResource{value: val, isSet: true}
}

func (v NullableCatalogV1alpha1SubscriptionMarketplaceResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogV1alpha1SubscriptionMarketplaceResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


