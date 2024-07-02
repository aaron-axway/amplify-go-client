# CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeStripe

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | Pointer to **string** | Stripe Invoice id. | [optional] 
**Number** | Pointer to **string** | Stripe Invoice number. | [optional] 
**DueDate** | Pointer to **time.Time** | Due date of the invoice. | [optional] 
**IssueDate** | Pointer to **time.Time** | Issue date of the invoice. | [optional] 
**Amount** | Pointer to [**CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeStripeAmount**](CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeStripeAmount.md) |  | [optional] 
**Link** | Pointer to **string** | Link where the payment can be done. | [optional] 
**DocumentLink** | Pointer to **string** | Link from where the invoice can be downloaded. | [optional] 
**Customer** | Pointer to [**CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeStripeCustomer**](CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeStripeCustomer.md) |  | [optional] 

## Methods

### NewCatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeStripe

`func NewCatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeStripe(type_ string, ) *CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeStripe`

NewCatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeStripe instantiates a new CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeStripe object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeStripeWithDefaults

`func NewCatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeStripeWithDefaults() *CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeStripe`

NewCatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeStripeWithDefaults instantiates a new CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeStripe object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeStripe) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeStripe) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeStripe) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeStripe) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeStripe) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeStripe) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeStripe) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNumber

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeStripe) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeStripe) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeStripe) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeStripe) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetDueDate

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeStripe) GetDueDate() time.Time`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeStripe) GetDueDateOk() (*time.Time, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeStripe) SetDueDate(v time.Time)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeStripe) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetIssueDate

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeStripe) GetIssueDate() time.Time`

GetIssueDate returns the IssueDate field if non-nil, zero value otherwise.

### GetIssueDateOk

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeStripe) GetIssueDateOk() (*time.Time, bool)`

GetIssueDateOk returns a tuple with the IssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDate

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeStripe) SetIssueDate(v time.Time)`

SetIssueDate sets IssueDate field to given value.

### HasIssueDate

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeStripe) HasIssueDate() bool`

HasIssueDate returns a boolean if a field has been set.

### GetAmount

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeStripe) GetAmount() CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeStripeAmount`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeStripe) GetAmountOk() (*CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeStripeAmount, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeStripe) SetAmount(v CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeStripeAmount)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeStripe) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetLink

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeStripe) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeStripe) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeStripe) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeStripe) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetDocumentLink

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeStripe) GetDocumentLink() string`

GetDocumentLink returns the DocumentLink field if non-nil, zero value otherwise.

### GetDocumentLinkOk

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeStripe) GetDocumentLinkOk() (*string, bool)`

GetDocumentLinkOk returns a tuple with the DocumentLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentLink

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeStripe) SetDocumentLink(v string)`

SetDocumentLink sets DocumentLink field to given value.

### HasDocumentLink

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeStripe) HasDocumentLink() bool`

HasDocumentLink returns a boolean if a field has been set.

### GetCustomer

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeStripe) GetCustomer() CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeStripeCustomer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeStripe) GetCustomerOk() (*CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeStripeCustomer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeStripe) SetCustomer(v CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeStripeCustomer)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeStripe) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


