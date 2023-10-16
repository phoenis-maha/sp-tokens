# CreateRestrictedDataTokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetApplication** | Pointer to **string** | The application ID for the target application to which access is being delegated. | [optional] 
**RestrictedResources** | [**[]RestrictedResource**](RestrictedResource.md) | A list of restricted resources. Maximum: 50 | 

## Methods

### NewCreateRestrictedDataTokenRequest

`func NewCreateRestrictedDataTokenRequest(restrictedResources []RestrictedResource, ) *CreateRestrictedDataTokenRequest`

NewCreateRestrictedDataTokenRequest instantiates a new CreateRestrictedDataTokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRestrictedDataTokenRequestWithDefaults

`func NewCreateRestrictedDataTokenRequestWithDefaults() *CreateRestrictedDataTokenRequest`

NewCreateRestrictedDataTokenRequestWithDefaults instantiates a new CreateRestrictedDataTokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetApplication

`func (o *CreateRestrictedDataTokenRequest) GetTargetApplication() string`

GetTargetApplication returns the TargetApplication field if non-nil, zero value otherwise.

### GetTargetApplicationOk

`func (o *CreateRestrictedDataTokenRequest) GetTargetApplicationOk() (*string, bool)`

GetTargetApplicationOk returns a tuple with the TargetApplication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetApplication

`func (o *CreateRestrictedDataTokenRequest) SetTargetApplication(v string)`

SetTargetApplication sets TargetApplication field to given value.

### HasTargetApplication

`func (o *CreateRestrictedDataTokenRequest) HasTargetApplication() bool`

HasTargetApplication returns a boolean if a field has been set.

### GetRestrictedResources

`func (o *CreateRestrictedDataTokenRequest) GetRestrictedResources() []RestrictedResource`

GetRestrictedResources returns the RestrictedResources field if non-nil, zero value otherwise.

### GetRestrictedResourcesOk

`func (o *CreateRestrictedDataTokenRequest) GetRestrictedResourcesOk() (*[]RestrictedResource, bool)`

GetRestrictedResourcesOk returns a tuple with the RestrictedResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedResources

`func (o *CreateRestrictedDataTokenRequest) SetRestrictedResources(v []RestrictedResource)`

SetRestrictedResources sets RestrictedResources field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


