# ManagementV1alpha1DataplaneSpecGitHubFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Paths** | Pointer to **[]string** | The paths within the repository that the agent will gather files for looking for specs | [optional] 
**Pattern** | Pointer to **[]string** | The regular expressions that when a file name passes any of the patterns, a spec file will be used to create an api service | [optional] 
**Branch** | Pointer to **string** | The GitHub repository branch that&#39;ll be used for the discovery process | [optional] 

## Methods

### NewManagementV1alpha1DataplaneSpecGitHubFilter

`func NewManagementV1alpha1DataplaneSpecGitHubFilter() *ManagementV1alpha1DataplaneSpecGitHubFilter`

NewManagementV1alpha1DataplaneSpecGitHubFilter instantiates a new ManagementV1alpha1DataplaneSpecGitHubFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementV1alpha1DataplaneSpecGitHubFilterWithDefaults

`func NewManagementV1alpha1DataplaneSpecGitHubFilterWithDefaults() *ManagementV1alpha1DataplaneSpecGitHubFilter`

NewManagementV1alpha1DataplaneSpecGitHubFilterWithDefaults instantiates a new ManagementV1alpha1DataplaneSpecGitHubFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaths

`func (o *ManagementV1alpha1DataplaneSpecGitHubFilter) GetPaths() []string`

GetPaths returns the Paths field if non-nil, zero value otherwise.

### GetPathsOk

`func (o *ManagementV1alpha1DataplaneSpecGitHubFilter) GetPathsOk() (*[]string, bool)`

GetPathsOk returns a tuple with the Paths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaths

`func (o *ManagementV1alpha1DataplaneSpecGitHubFilter) SetPaths(v []string)`

SetPaths sets Paths field to given value.

### HasPaths

`func (o *ManagementV1alpha1DataplaneSpecGitHubFilter) HasPaths() bool`

HasPaths returns a boolean if a field has been set.

### GetPattern

`func (o *ManagementV1alpha1DataplaneSpecGitHubFilter) GetPattern() []string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *ManagementV1alpha1DataplaneSpecGitHubFilter) GetPatternOk() (*[]string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *ManagementV1alpha1DataplaneSpecGitHubFilter) SetPattern(v []string)`

SetPattern sets Pattern field to given value.

### HasPattern

`func (o *ManagementV1alpha1DataplaneSpecGitHubFilter) HasPattern() bool`

HasPattern returns a boolean if a field has been set.

### GetBranch

`func (o *ManagementV1alpha1DataplaneSpecGitHubFilter) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *ManagementV1alpha1DataplaneSpecGitHubFilter) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *ManagementV1alpha1DataplaneSpecGitHubFilter) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *ManagementV1alpha1DataplaneSpecGitHubFilter) HasBranch() bool`

HasBranch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


