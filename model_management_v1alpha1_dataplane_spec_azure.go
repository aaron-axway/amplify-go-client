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

// checks if the ManagementV1alpha1DataplaneSpecAzure type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManagementV1alpha1DataplaneSpecAzure{}

// ManagementV1alpha1DataplaneSpecAzure The configuration common to all Azure agents that use this dataplane
type ManagementV1alpha1DataplaneSpecAzure struct {
	Type *string `json:"type,omitempty"`
	// The tenantId is a Microsoft Entra ID entity that typically encompasses an organization
	TenantId string `json:"tenantId" validate:"regexp=^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$"`
	// The resource group holds related resources for any Azure solution
	ResourceGroup string `json:"resourceGroup"`
	// The subscriptionId is the ID given to the subscription tied to the tenant
	SubscriptionId string `json:"subscriptionId" validate:"regexp=^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$"`
	// The name of the azure API management
	ApimServiceName string `json:"apimServiceName"`
	// The event hub processes and stores events, data, or telemetry produced by distributed software or devices
	EventHubName *string `json:"eventHubName,omitempty"`
	// The event hub namespace is a management container for event hubs or topics
	EventHubNamespace *string `json:"eventHubNamespace,omitempty"`
	// Consumer groups enable consuming applications to each have a separate view of the event stream. They read the stream independently at their own pace and with their own offsets.
	EventHubConsumerGroup *string `json:"eventHubConsumerGroup,omitempty" validate:"regexp=^[a-zA-Z0-9$]+$"`
}

type _ManagementV1alpha1DataplaneSpecAzure ManagementV1alpha1DataplaneSpecAzure

// NewManagementV1alpha1DataplaneSpecAzure instantiates a new ManagementV1alpha1DataplaneSpecAzure object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagementV1alpha1DataplaneSpecAzure(tenantId string, resourceGroup string, subscriptionId string, apimServiceName string) *ManagementV1alpha1DataplaneSpecAzure {
	this := ManagementV1alpha1DataplaneSpecAzure{}
	this.TenantId = tenantId
	this.ResourceGroup = resourceGroup
	this.SubscriptionId = subscriptionId
	this.ApimServiceName = apimServiceName
	return &this
}

// NewManagementV1alpha1DataplaneSpecAzureWithDefaults instantiates a new ManagementV1alpha1DataplaneSpecAzure object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagementV1alpha1DataplaneSpecAzureWithDefaults() *ManagementV1alpha1DataplaneSpecAzure {
	this := ManagementV1alpha1DataplaneSpecAzure{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ManagementV1alpha1DataplaneSpecAzure) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1DataplaneSpecAzure) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ManagementV1alpha1DataplaneSpecAzure) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ManagementV1alpha1DataplaneSpecAzure) SetType(v string) {
	o.Type = &v
}

// GetTenantId returns the TenantId field value
func (o *ManagementV1alpha1DataplaneSpecAzure) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1DataplaneSpecAzure) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *ManagementV1alpha1DataplaneSpecAzure) SetTenantId(v string) {
	o.TenantId = v
}

// GetResourceGroup returns the ResourceGroup field value
func (o *ManagementV1alpha1DataplaneSpecAzure) GetResourceGroup() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceGroup
}

// GetResourceGroupOk returns a tuple with the ResourceGroup field value
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1DataplaneSpecAzure) GetResourceGroupOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceGroup, true
}

// SetResourceGroup sets field value
func (o *ManagementV1alpha1DataplaneSpecAzure) SetResourceGroup(v string) {
	o.ResourceGroup = v
}

// GetSubscriptionId returns the SubscriptionId field value
func (o *ManagementV1alpha1DataplaneSpecAzure) GetSubscriptionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1DataplaneSpecAzure) GetSubscriptionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubscriptionId, true
}

// SetSubscriptionId sets field value
func (o *ManagementV1alpha1DataplaneSpecAzure) SetSubscriptionId(v string) {
	o.SubscriptionId = v
}

// GetApimServiceName returns the ApimServiceName field value
func (o *ManagementV1alpha1DataplaneSpecAzure) GetApimServiceName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApimServiceName
}

// GetApimServiceNameOk returns a tuple with the ApimServiceName field value
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1DataplaneSpecAzure) GetApimServiceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApimServiceName, true
}

// SetApimServiceName sets field value
func (o *ManagementV1alpha1DataplaneSpecAzure) SetApimServiceName(v string) {
	o.ApimServiceName = v
}

// GetEventHubName returns the EventHubName field value if set, zero value otherwise.
func (o *ManagementV1alpha1DataplaneSpecAzure) GetEventHubName() string {
	if o == nil || IsNil(o.EventHubName) {
		var ret string
		return ret
	}
	return *o.EventHubName
}

// GetEventHubNameOk returns a tuple with the EventHubName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1DataplaneSpecAzure) GetEventHubNameOk() (*string, bool) {
	if o == nil || IsNil(o.EventHubName) {
		return nil, false
	}
	return o.EventHubName, true
}

// HasEventHubName returns a boolean if a field has been set.
func (o *ManagementV1alpha1DataplaneSpecAzure) HasEventHubName() bool {
	if o != nil && !IsNil(o.EventHubName) {
		return true
	}

	return false
}

// SetEventHubName gets a reference to the given string and assigns it to the EventHubName field.
func (o *ManagementV1alpha1DataplaneSpecAzure) SetEventHubName(v string) {
	o.EventHubName = &v
}

// GetEventHubNamespace returns the EventHubNamespace field value if set, zero value otherwise.
func (o *ManagementV1alpha1DataplaneSpecAzure) GetEventHubNamespace() string {
	if o == nil || IsNil(o.EventHubNamespace) {
		var ret string
		return ret
	}
	return *o.EventHubNamespace
}

// GetEventHubNamespaceOk returns a tuple with the EventHubNamespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1DataplaneSpecAzure) GetEventHubNamespaceOk() (*string, bool) {
	if o == nil || IsNil(o.EventHubNamespace) {
		return nil, false
	}
	return o.EventHubNamespace, true
}

// HasEventHubNamespace returns a boolean if a field has been set.
func (o *ManagementV1alpha1DataplaneSpecAzure) HasEventHubNamespace() bool {
	if o != nil && !IsNil(o.EventHubNamespace) {
		return true
	}

	return false
}

// SetEventHubNamespace gets a reference to the given string and assigns it to the EventHubNamespace field.
func (o *ManagementV1alpha1DataplaneSpecAzure) SetEventHubNamespace(v string) {
	o.EventHubNamespace = &v
}

// GetEventHubConsumerGroup returns the EventHubConsumerGroup field value if set, zero value otherwise.
func (o *ManagementV1alpha1DataplaneSpecAzure) GetEventHubConsumerGroup() string {
	if o == nil || IsNil(o.EventHubConsumerGroup) {
		var ret string
		return ret
	}
	return *o.EventHubConsumerGroup
}

// GetEventHubConsumerGroupOk returns a tuple with the EventHubConsumerGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1DataplaneSpecAzure) GetEventHubConsumerGroupOk() (*string, bool) {
	if o == nil || IsNil(o.EventHubConsumerGroup) {
		return nil, false
	}
	return o.EventHubConsumerGroup, true
}

// HasEventHubConsumerGroup returns a boolean if a field has been set.
func (o *ManagementV1alpha1DataplaneSpecAzure) HasEventHubConsumerGroup() bool {
	if o != nil && !IsNil(o.EventHubConsumerGroup) {
		return true
	}

	return false
}

// SetEventHubConsumerGroup gets a reference to the given string and assigns it to the EventHubConsumerGroup field.
func (o *ManagementV1alpha1DataplaneSpecAzure) SetEventHubConsumerGroup(v string) {
	o.EventHubConsumerGroup = &v
}

func (o ManagementV1alpha1DataplaneSpecAzure) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManagementV1alpha1DataplaneSpecAzure) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	toSerialize["tenantId"] = o.TenantId
	toSerialize["resourceGroup"] = o.ResourceGroup
	toSerialize["subscriptionId"] = o.SubscriptionId
	toSerialize["apimServiceName"] = o.ApimServiceName
	if !IsNil(o.EventHubName) {
		toSerialize["eventHubName"] = o.EventHubName
	}
	if !IsNil(o.EventHubNamespace) {
		toSerialize["eventHubNamespace"] = o.EventHubNamespace
	}
	if !IsNil(o.EventHubConsumerGroup) {
		toSerialize["eventHubConsumerGroup"] = o.EventHubConsumerGroup
	}
	return toSerialize, nil
}

func (o *ManagementV1alpha1DataplaneSpecAzure) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tenantId",
		"resourceGroup",
		"subscriptionId",
		"apimServiceName",
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

	varManagementV1alpha1DataplaneSpecAzure := _ManagementV1alpha1DataplaneSpecAzure{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varManagementV1alpha1DataplaneSpecAzure)

	if err != nil {
		return err
	}

	*o = ManagementV1alpha1DataplaneSpecAzure(varManagementV1alpha1DataplaneSpecAzure)

	return err
}

type NullableManagementV1alpha1DataplaneSpecAzure struct {
	value *ManagementV1alpha1DataplaneSpecAzure
	isSet bool
}

func (v NullableManagementV1alpha1DataplaneSpecAzure) Get() *ManagementV1alpha1DataplaneSpecAzure {
	return v.value
}

func (v *NullableManagementV1alpha1DataplaneSpecAzure) Set(val *ManagementV1alpha1DataplaneSpecAzure) {
	v.value = val
	v.isSet = true
}

func (v NullableManagementV1alpha1DataplaneSpecAzure) IsSet() bool {
	return v.isSet
}

func (v *NullableManagementV1alpha1DataplaneSpecAzure) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagementV1alpha1DataplaneSpecAzure(val *ManagementV1alpha1DataplaneSpecAzure) *NullableManagementV1alpha1DataplaneSpecAzure {
	return &NullableManagementV1alpha1DataplaneSpecAzure{value: val, isSet: true}
}

func (v NullableManagementV1alpha1DataplaneSpecAzure) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagementV1alpha1DataplaneSpecAzure) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


