# ApiV1JsonPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Op** | **string** | The patch operation to perform: * add - Adds or replaces a property with the given value. * replace - Replaces an existing property with the given value. Throws an error if property does not exist. * remove - Removes the property. Throws an error if the property does not exist. * x-build-object-tree - Adds object properties for each segment in the path if they don&#39;t already exist.  | 
**Path** | **string** |  | 
**Value** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewApiV1JsonPatch

`func NewApiV1JsonPatch(op string, path string, ) *ApiV1JsonPatch`

NewApiV1JsonPatch instantiates a new ApiV1JsonPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV1JsonPatchWithDefaults

`func NewApiV1JsonPatchWithDefaults() *ApiV1JsonPatch`

NewApiV1JsonPatchWithDefaults instantiates a new ApiV1JsonPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOp

`func (o *ApiV1JsonPatch) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *ApiV1JsonPatch) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *ApiV1JsonPatch) SetOp(v string)`

SetOp sets Op field to given value.


### GetPath

`func (o *ApiV1JsonPatch) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ApiV1JsonPatch) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ApiV1JsonPatch) SetPath(v string)`

SetPath sets Path field to given value.


### GetValue

`func (o *ApiV1JsonPatch) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ApiV1JsonPatch) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ApiV1JsonPatch) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *ApiV1JsonPatch) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


