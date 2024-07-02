# CatalogV1alpha1SubscriptionInvoiceBillingPayment

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

### NewCatalogV1alpha1SubscriptionInvoiceBillingPayment

`func NewCatalogV1alpha1SubscriptionInvoiceBillingPayment(type_ string, ) *CatalogV1alpha1SubscriptionInvoiceBillingPayment`

NewCatalogV1alpha1SubscriptionInvoiceBillingPayment instantiates a new CatalogV1alpha1SubscriptionInvoiceBillingPayment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogV1alpha1SubscriptionInvoiceBillingPaymentWithDefaults

`func NewCatalogV1alpha1SubscriptionInvoiceBillingPaymentWithDefaults() *CatalogV1alpha1SubscriptionInvoiceBillingPayment`

NewCatalogV1alpha1SubscriptionInvoiceBillingPaymentWithDefaults instantiates a new CatalogV1alpha1SubscriptionInvoiceBillingPayment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPayment) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPayment) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPayment) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPayment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPayment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPayment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPayment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNumber

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPayment) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPayment) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPayment) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPayment) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetDueDate

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPayment) GetDueDate() time.Time`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPayment) GetDueDateOk() (*time.Time, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPayment) SetDueDate(v time.Time)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPayment) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetIssueDate

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPayment) GetIssueDate() time.Time`

GetIssueDate returns the IssueDate field if non-nil, zero value otherwise.

### GetIssueDateOk

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPayment) GetIssueDateOk() (*time.Time, bool)`

GetIssueDateOk returns a tuple with the IssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDate

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPayment) SetIssueDate(v time.Time)`

SetIssueDate sets IssueDate field to given value.

### HasIssueDate

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPayment) HasIssueDate() bool`

HasIssueDate returns a boolean if a field has been set.

### GetAmount

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPayment) GetAmount() CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeCustomAmount`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPayment) GetAmountOk() (*CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeCustomAmount, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPayment) SetAmount(v CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeCustomAmount)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPayment) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetLink

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPayment) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPayment) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPayment) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPayment) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetDocumentLink

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPayment) GetDocumentLink() string`

GetDocumentLink returns the DocumentLink field if non-nil, zero value otherwise.

### GetDocumentLinkOk

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPayment) GetDocumentLinkOk() (*string, bool)`

GetDocumentLinkOk returns a tuple with the DocumentLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentLink

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPayment) SetDocumentLink(v string)`

SetDocumentLink sets DocumentLink field to given value.

### HasDocumentLink

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPayment) HasDocumentLink() bool`

HasDocumentLink returns a boolean if a field has been set.

### GetCustomer

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPayment) GetCustomer() CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeCustomCustomer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPayment) GetCustomerOk() (*CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeCustomCustomer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPayment) SetCustomer(v CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeCustomCustomer)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPayment) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


