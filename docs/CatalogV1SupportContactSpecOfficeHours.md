# CatalogV1SupportContactSpecOfficeHours

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timezone** | **string** | Long IANA Time Zone format. Examples: &#39;America/New_York&#39; or &#39;Europe/Paris&#39; | 
**Periods** | [**[]CatalogV1SupportContactSpecOfficeHoursPeriodsInner**](CatalogV1SupportContactSpecOfficeHoursPeriodsInner.md) |  | 

## Methods

### NewCatalogV1SupportContactSpecOfficeHours

`func NewCatalogV1SupportContactSpecOfficeHours(timezone string, periods []CatalogV1SupportContactSpecOfficeHoursPeriodsInner, ) *CatalogV1SupportContactSpecOfficeHours`

NewCatalogV1SupportContactSpecOfficeHours instantiates a new CatalogV1SupportContactSpecOfficeHours object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogV1SupportContactSpecOfficeHoursWithDefaults

`func NewCatalogV1SupportContactSpecOfficeHoursWithDefaults() *CatalogV1SupportContactSpecOfficeHours`

NewCatalogV1SupportContactSpecOfficeHoursWithDefaults instantiates a new CatalogV1SupportContactSpecOfficeHours object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimezone

`func (o *CatalogV1SupportContactSpecOfficeHours) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *CatalogV1SupportContactSpecOfficeHours) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *CatalogV1SupportContactSpecOfficeHours) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.


### GetPeriods

`func (o *CatalogV1SupportContactSpecOfficeHours) GetPeriods() []CatalogV1SupportContactSpecOfficeHoursPeriodsInner`

GetPeriods returns the Periods field if non-nil, zero value otherwise.

### GetPeriodsOk

`func (o *CatalogV1SupportContactSpecOfficeHours) GetPeriodsOk() (*[]CatalogV1SupportContactSpecOfficeHoursPeriodsInner, bool)`

GetPeriodsOk returns a tuple with the Periods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriods

`func (o *CatalogV1SupportContactSpecOfficeHours) SetPeriods(v []CatalogV1SupportContactSpecOfficeHoursPeriodsInner)`

SetPeriods sets Periods field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


