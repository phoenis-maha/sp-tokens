# CreateRestrictedDataTokenResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RestrictedDataToken** | Pointer to **string** | A Restricted Data Token (RDT). This is a short-lived access token that authorizes calls to restricted operations. Pass this value with the x-amz-access-token header when making subsequent calls to these restricted resources. | [optional] 
**ExpiresIn** | Pointer to **int32** | The lifetime of the Restricted Data Token, in seconds. | [optional] 

## Methods

### NewCreateRestrictedDataTokenResponse

`func NewCreateRestrictedDataTokenResponse() *CreateRestrictedDataTokenResponse`

NewCreateRestrictedDataTokenResponse instantiates a new CreateRestrictedDataTokenResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRestrictedDataTokenResponseWithDefaults

`func NewCreateRestrictedDataTokenResponseWithDefaults() *CreateRestrictedDataTokenResponse`

NewCreateRestrictedDataTokenResponseWithDefaults instantiates a new CreateRestrictedDataTokenResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRestrictedDataToken

`func (o *CreateRestrictedDataTokenResponse) GetRestrictedDataToken() string`

GetRestrictedDataToken returns the RestrictedDataToken field if non-nil, zero value otherwise.

### GetRestrictedDataTokenOk

`func (o *CreateRestrictedDataTokenResponse) GetRestrictedDataTokenOk() (*string, bool)`

GetRestrictedDataTokenOk returns a tuple with the RestrictedDataToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedDataToken

`func (o *CreateRestrictedDataTokenResponse) SetRestrictedDataToken(v string)`

SetRestrictedDataToken sets RestrictedDataToken field to given value.

### HasRestrictedDataToken

`func (o *CreateRestrictedDataTokenResponse) HasRestrictedDataToken() bool`

HasRestrictedDataToken returns a boolean if a field has been set.

### GetExpiresIn

`func (o *CreateRestrictedDataTokenResponse) GetExpiresIn() int32`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *CreateRestrictedDataTokenResponse) GetExpiresInOk() (*int32, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *CreateRestrictedDataTokenResponse) SetExpiresIn(v int32)`

SetExpiresIn sets ExpiresIn field to given value.

### HasExpiresIn

`func (o *CreateRestrictedDataTokenResponse) HasExpiresIn() bool`

HasExpiresIn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


