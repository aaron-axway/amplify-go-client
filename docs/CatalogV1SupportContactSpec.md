# CatalogV1SupportContactSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | Email address of the Support Contact. | 
**PhoneNumber** | Pointer to **string** | String of the E.164 format of the phone number, e.g. +11234567899 | [optional] 
**OfficeHours** | Pointer to [**CatalogV1SupportContactSpecOfficeHours**](CatalogV1SupportContactSpecOfficeHours.md) |  | [optional] 
**AlternativeContacts** | Pointer to [**CatalogV1SupportContactSpecAlternativeContacts**](CatalogV1SupportContactSpecAlternativeContacts.md) |  | [optional] 

## Methods

### NewCatalogV1SupportContactSpec

`func NewCatalogV1SupportContactSpec(email string, ) *CatalogV1SupportContactSpec`

NewCatalogV1SupportContactSpec instantiates a new CatalogV1SupportContactSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogV1SupportContactSpecWithDefaults

`func NewCatalogV1SupportContactSpecWithDefaults() *CatalogV1SupportContactSpec`

NewCatalogV1SupportContactSpecWithDefaults instantiates a new CatalogV1SupportContactSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *CatalogV1SupportContactSpec) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CatalogV1SupportContactSpec) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CatalogV1SupportContactSpec) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetPhoneNumber

`func (o *CatalogV1SupportContactSpec) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *CatalogV1SupportContactSpec) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *CatalogV1SupportContactSpec) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *CatalogV1SupportContactSpec) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetOfficeHours

`func (o *CatalogV1SupportContactSpec) GetOfficeHours() CatalogV1SupportContactSpecOfficeHours`

GetOfficeHours returns the OfficeHours field if non-nil, zero value otherwise.

### GetOfficeHoursOk

`func (o *CatalogV1SupportContactSpec) GetOfficeHoursOk() (*CatalogV1SupportContactSpecOfficeHours, bool)`

GetOfficeHoursOk returns a tuple with the OfficeHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficeHours

`func (o *CatalogV1SupportContactSpec) SetOfficeHours(v CatalogV1SupportContactSpecOfficeHours)`

SetOfficeHours sets OfficeHours field to given value.

### HasOfficeHours

`func (o *CatalogV1SupportContactSpec) HasOfficeHours() bool`

HasOfficeHours returns a boolean if a field has been set.

### GetAlternativeContacts

`func (o *CatalogV1SupportContactSpec) GetAlternativeContacts() CatalogV1SupportContactSpecAlternativeContacts`

GetAlternativeContacts returns the AlternativeContacts field if non-nil, zero value otherwise.

### GetAlternativeContactsOk

`func (o *CatalogV1SupportContactSpec) GetAlternativeContactsOk() (*CatalogV1SupportContactSpecAlternativeContacts, bool)`

GetAlternativeContactsOk returns a tuple with the AlternativeContacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternativeContacts

`func (o *CatalogV1SupportContactSpec) SetAlternativeContacts(v CatalogV1SupportContactSpecAlternativeContacts)`

SetAlternativeContacts sets AlternativeContacts field to given value.

### HasAlternativeContacts

`func (o *CatalogV1SupportContactSpec) HasAlternativeContacts() bool`

HasAlternativeContacts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


