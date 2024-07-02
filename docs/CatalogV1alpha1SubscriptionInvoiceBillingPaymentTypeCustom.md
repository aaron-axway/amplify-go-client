# CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeCustom

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | Pointer to **string** | Invoice id from the billing system. | [optional] 
**Number** | Pointer to **string** | Custom Invoice number. | [optional] 
**DueDate** | Pointer to **time.Time** | Due date of the invoice. | [optional] 
**IssueDate** | Pointer to **time.Time** | Issue date of the invoice. | [optional] 
**Amount** | Pointer to [**CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeCustomAmount**](CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeCustomAmount.md) |  | [optional] 
**Link** | Pointer to **string** | Link where the payment can be done. | [optional] 
**DocumentLink** | Pointer to **string** | Link from where the invoice can be downloaded. | [optional] 
**Customer** | Pointer to [**CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeCustomCustomer**](CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeCustomCustomer.md) |  | [optional] 

## Methods

### NewCatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeCustom

`func NewCatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeCustom(type_ string, ) *CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeCustom`

NewCatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeCustom instantiates a new CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeCustom object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeCustomWithDefaults

`func NewCatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeCustomWithDefaults() *CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeCustom`

NewCatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeCustomWithDefaults instantiates a new CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeCustom object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeCustom) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeCustom) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeCustom) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeCustom) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeCustom) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeCustom) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeCustom) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNumber

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeCustom) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeCustom) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeCustom) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeCustom) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetDueDate

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeCustom) GetDueDate() time.Time`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeCustom) GetDueDateOk() (*time.Time, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeCustom) SetDueDate(v time.Time)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeCustom) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetIssueDate

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeCustom) GetIssueDate() time.Time`

GetIssueDate returns the IssueDate field if non-nil, zero value otherwise.

### GetIssueDateOk

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeCustom) GetIssueDateOk() (*time.Time, bool)`

GetIssueDateOk returns a tuple with the IssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDate

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeCustom) SetIssueDate(v time.Time)`

SetIssueDate sets IssueDate field to given value.

### HasIssueDate

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeCustom) HasIssueDate() bool`

HasIssueDate returns a boolean if a field has been set.

### GetAmount

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeCustom) GetAmount() CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeCustomAmount`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeCustom) GetAmountOk() (*CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeCustomAmount, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeCustom) SetAmount(v CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeCustomAmount)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeCustom) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetLink

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeCustom) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeCustom) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeCustom) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeCustom) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetDocumentLink

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeCustom) GetDocumentLink() string`

GetDocumentLink returns the DocumentLink field if non-nil, zero value otherwise.

### GetDocumentLinkOk

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeCustom) GetDocumentLinkOk() (*string, bool)`

GetDocumentLinkOk returns a tuple with the DocumentLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentLink

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeCustom) SetDocumentLink(v string)`

SetDocumentLink sets DocumentLink field to given value.

### HasDocumentLink

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeCustom) HasDocumentLink() bool`

HasDocumentLink returns a boolean if a field has been set.

### GetCustomer

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeCustom) GetCustomer() CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeCustomCustomer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeCustom) GetCustomerOk() (*CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeCustomCustomer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeCustom) SetCustomer(v CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeCustomCustomer)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeCustom) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


