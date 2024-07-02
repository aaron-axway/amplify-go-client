# ApiV1Owner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of the owner. Defaults to team if not present. | [optional] [default to "team"]
**Id** | **string** | Id of the owner of the resource. | 

## Methods

### NewApiV1Owner

`func NewApiV1Owner(id string, ) *ApiV1Owner`

NewApiV1Owner instantiates a new ApiV1Owner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV1OwnerWithDefaults

`func NewApiV1OwnerWithDefaults() *ApiV1Owner`

NewApiV1OwnerWithDefaults instantiates a new ApiV1Owner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ApiV1Owner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiV1Owner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiV1Owner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ApiV1Owner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *ApiV1Owner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiV1Owner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiV1Owner) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


