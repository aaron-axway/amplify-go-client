# ManagementV1alpha1DataplaneSpecAWS

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**AccessLogARN** | Pointer to **string** | The ARN of the cloud watch log resource that AWS API Gateway will be configured to send API Access data to | [optional] 
**FullTransactionLogging** | Pointer to **bool** | If true, the discovery agent will enable full transaction logging for discovered API stages | [optional] 

## Methods

### NewManagementV1alpha1DataplaneSpecAWS

`func NewManagementV1alpha1DataplaneSpecAWS(type_ string, ) *ManagementV1alpha1DataplaneSpecAWS`

NewManagementV1alpha1DataplaneSpecAWS instantiates a new ManagementV1alpha1DataplaneSpecAWS object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementV1alpha1DataplaneSpecAWSWithDefaults

`func NewManagementV1alpha1DataplaneSpecAWSWithDefaults() *ManagementV1alpha1DataplaneSpecAWS`

NewManagementV1alpha1DataplaneSpecAWSWithDefaults instantiates a new ManagementV1alpha1DataplaneSpecAWS object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ManagementV1alpha1DataplaneSpecAWS) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ManagementV1alpha1DataplaneSpecAWS) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ManagementV1alpha1DataplaneSpecAWS) SetType(v string)`

SetType sets Type field to given value.


### GetAccessLogARN

`func (o *ManagementV1alpha1DataplaneSpecAWS) GetAccessLogARN() string`

GetAccessLogARN returns the AccessLogARN field if non-nil, zero value otherwise.

### GetAccessLogARNOk

`func (o *ManagementV1alpha1DataplaneSpecAWS) GetAccessLogARNOk() (*string, bool)`

GetAccessLogARNOk returns a tuple with the AccessLogARN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessLogARN

`func (o *ManagementV1alpha1DataplaneSpecAWS) SetAccessLogARN(v string)`

SetAccessLogARN sets AccessLogARN field to given value.

### HasAccessLogARN

`func (o *ManagementV1alpha1DataplaneSpecAWS) HasAccessLogARN() bool`

HasAccessLogARN returns a boolean if a field has been set.

### GetFullTransactionLogging

`func (o *ManagementV1alpha1DataplaneSpecAWS) GetFullTransactionLogging() bool`

GetFullTransactionLogging returns the FullTransactionLogging field if non-nil, zero value otherwise.

### GetFullTransactionLoggingOk

`func (o *ManagementV1alpha1DataplaneSpecAWS) GetFullTransactionLoggingOk() (*bool, bool)`

GetFullTransactionLoggingOk returns a tuple with the FullTransactionLogging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullTransactionLogging

`func (o *ManagementV1alpha1DataplaneSpecAWS) SetFullTransactionLogging(v bool)`

SetFullTransactionLogging sets FullTransactionLogging field to given value.

### HasFullTransactionLogging

`func (o *ManagementV1alpha1DataplaneSpecAWS) HasFullTransactionLogging() bool`

HasFullTransactionLogging returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


