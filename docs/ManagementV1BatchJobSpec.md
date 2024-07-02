# ManagementV1BatchJobSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | Unique string ID assigned by a controller indicating what action this job is performing. | 
**References** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewManagementV1BatchJobSpec

`func NewManagementV1BatchJobSpec(action string, ) *ManagementV1BatchJobSpec`

NewManagementV1BatchJobSpec instantiates a new ManagementV1BatchJobSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementV1BatchJobSpecWithDefaults

`func NewManagementV1BatchJobSpecWithDefaults() *ManagementV1BatchJobSpec`

NewManagementV1BatchJobSpecWithDefaults instantiates a new ManagementV1BatchJobSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *ManagementV1BatchJobSpec) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ManagementV1BatchJobSpec) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ManagementV1BatchJobSpec) SetAction(v string)`

SetAction sets Action field to given value.


### GetReferences

`func (o *ManagementV1BatchJobSpec) GetReferences() []map[string]interface{}`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *ManagementV1BatchJobSpec) GetReferencesOk() (*[]map[string]interface{}, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *ManagementV1BatchJobSpec) SetReferences(v []map[string]interface{})`

SetReferences sets References field to given value.

### HasReferences

`func (o *ManagementV1BatchJobSpec) HasReferences() bool`

HasReferences returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


