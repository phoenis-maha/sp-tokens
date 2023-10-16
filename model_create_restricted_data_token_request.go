/*
Selling Partner API for Tokens

The Selling Partner API for Tokens provides a secure way to access a customer's PII (Personally Identifiable Information). You can call the Tokens API to get a Restricted Data Token (RDT) for one or more restricted resources that you specify. The RDT authorizes subsequent calls to restricted operations that correspond to the restricted resources that you specified.  For more information, see the [Tokens API Use Case Guide](doc:tokens-api-use-case-guide).

API version: 2021-03-01
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the CreateRestrictedDataTokenRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateRestrictedDataTokenRequest{}

// CreateRestrictedDataTokenRequest The request schema for the createRestrictedDataToken operation.
type CreateRestrictedDataTokenRequest struct {
	// The application ID for the target application to which access is being delegated.
	TargetApplication *string `json:"targetApplication,omitempty"`
	// A list of restricted resources. Maximum: 50
	RestrictedResources []RestrictedResource `json:"restrictedResources"`
}

// NewCreateRestrictedDataTokenRequest instantiates a new CreateRestrictedDataTokenRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateRestrictedDataTokenRequest(restrictedResources []RestrictedResource) *CreateRestrictedDataTokenRequest {
	this := CreateRestrictedDataTokenRequest{}
	this.RestrictedResources = restrictedResources
	return &this
}

// NewCreateRestrictedDataTokenRequestWithDefaults instantiates a new CreateRestrictedDataTokenRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateRestrictedDataTokenRequestWithDefaults() *CreateRestrictedDataTokenRequest {
	this := CreateRestrictedDataTokenRequest{}
	return &this
}

// GetTargetApplication returns the TargetApplication field value if set, zero value otherwise.
func (o *CreateRestrictedDataTokenRequest) GetTargetApplication() string {
	if o == nil || IsNil(o.TargetApplication) {
		var ret string
		return ret
	}
	return *o.TargetApplication
}

// GetTargetApplicationOk returns a tuple with the TargetApplication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRestrictedDataTokenRequest) GetTargetApplicationOk() (*string, bool) {
	if o == nil || IsNil(o.TargetApplication) {
		return nil, false
	}
	return o.TargetApplication, true
}

// HasTargetApplication returns a boolean if a field has been set.
func (o *CreateRestrictedDataTokenRequest) HasTargetApplication() bool {
	if o != nil && !IsNil(o.TargetApplication) {
		return true
	}

	return false
}

// SetTargetApplication gets a reference to the given string and assigns it to the TargetApplication field.
func (o *CreateRestrictedDataTokenRequest) SetTargetApplication(v string) {
	o.TargetApplication = &v
}

// GetRestrictedResources returns the RestrictedResources field value
func (o *CreateRestrictedDataTokenRequest) GetRestrictedResources() []RestrictedResource {
	if o == nil {
		var ret []RestrictedResource
		return ret
	}

	return o.RestrictedResources
}

// GetRestrictedResourcesOk returns a tuple with the RestrictedResources field value
// and a boolean to check if the value has been set.
func (o *CreateRestrictedDataTokenRequest) GetRestrictedResourcesOk() ([]RestrictedResource, bool) {
	if o == nil {
		return nil, false
	}
	return o.RestrictedResources, true
}

// SetRestrictedResources sets field value
func (o *CreateRestrictedDataTokenRequest) SetRestrictedResources(v []RestrictedResource) {
	o.RestrictedResources = v
}

func (o CreateRestrictedDataTokenRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateRestrictedDataTokenRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TargetApplication) {
		toSerialize["targetApplication"] = o.TargetApplication
	}
	toSerialize["restrictedResources"] = o.RestrictedResources
	return toSerialize, nil
}

type NullableCreateRestrictedDataTokenRequest struct {
	value *CreateRestrictedDataTokenRequest
	isSet bool
}

func (v NullableCreateRestrictedDataTokenRequest) Get() *CreateRestrictedDataTokenRequest {
	return v.value
}

func (v *NullableCreateRestrictedDataTokenRequest) Set(val *CreateRestrictedDataTokenRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateRestrictedDataTokenRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateRestrictedDataTokenRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateRestrictedDataTokenRequest(val *CreateRestrictedDataTokenRequest) *NullableCreateRestrictedDataTokenRequest {
	return &NullableCreateRestrictedDataTokenRequest{value: val, isSet: true}
}

func (v NullableCreateRestrictedDataTokenRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateRestrictedDataTokenRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}