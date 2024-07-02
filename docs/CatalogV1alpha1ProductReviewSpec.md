# CatalogV1alpha1ProductReviewSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rating** | **int32** |  | 
**Comment** | Pointer to **string** |  | [optional] 

## Methods

### NewCatalogV1alpha1ProductReviewSpec

`func NewCatalogV1alpha1ProductReviewSpec(rating int32, ) *CatalogV1alpha1ProductReviewSpec`

NewCatalogV1alpha1ProductReviewSpec instantiates a new CatalogV1alpha1ProductReviewSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogV1alpha1ProductReviewSpecWithDefaults

`func NewCatalogV1alpha1ProductReviewSpecWithDefaults() *CatalogV1alpha1ProductReviewSpec`

NewCatalogV1alpha1ProductReviewSpecWithDefaults instantiates a new CatalogV1alpha1ProductReviewSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRating

`func (o *CatalogV1alpha1ProductReviewSpec) GetRating() int32`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *CatalogV1alpha1ProductReviewSpec) GetRatingOk() (*int32, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *CatalogV1alpha1ProductReviewSpec) SetRating(v int32)`

SetRating sets Rating field to given value.


### GetComment

`func (o *CatalogV1alpha1ProductReviewSpec) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *CatalogV1alpha1ProductReviewSpec) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *CatalogV1alpha1ProductReviewSpec) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *CatalogV1alpha1ProductReviewSpec) HasComment() bool`

HasComment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


