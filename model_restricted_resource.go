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

// checks if the RestrictedResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RestrictedResource{}

// RestrictedResource Model of a restricted resource.
type RestrictedResource struct {
	// The HTTP method in the restricted resource.
	Method string `json:"method"`
	// The path in the restricted resource. Here are some path examples: - ```/orders/v0/orders```. For getting an RDT for the getOrders operation of the Orders API. For bulk orders. - ```/orders/v0/orders/123-1234567-1234567```. For getting an RDT for the getOrder operation of the Orders API. For a specific order. - ```/orders/v0/orders/123-1234567-1234567/orderItems```. For getting an RDT for the getOrderItems operation of the Orders API. For the order items in a specific order. - ```/mfn/v0/shipments/FBA1234ABC5D```. For getting an RDT for the getShipment operation of the Shipping API. For a specific shipment. - ```/mfn/v0/shipments/{shipmentId}```. For getting an RDT for the getShipment operation of the Shipping API. For any of a selling partner's shipments that you specify when you call the getShipment operation.
	Path string `json:"path"`
	// Indicates the type of Personally Identifiable Information requested. This parameter is required only when getting an RDT for use with the getOrder, getOrders, or getOrderItems operation of the Orders API. For more information, see the [Tokens API Use Case Guide](doc:tokens-api-use-case-guide). Possible values include: - **buyerInfo**. On the order level this includes general identifying information about the buyer and tax-related information. On the order item level this includes gift wrap information and custom order information, if available. - **shippingAddress**. This includes information for fulfilling orders. - **buyerTaxInformation**. This includes information for issuing tax invoices.
	DataElements []string `json:"dataElements,omitempty"`
}

// NewRestrictedResource instantiates a new RestrictedResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRestrictedResource(method string, path string) *RestrictedResource {
	this := RestrictedResource{}
	this.Method = method
	this.Path = path
	return &this
}

// NewRestrictedResourceWithDefaults instantiates a new RestrictedResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRestrictedResourceWithDefaults() *RestrictedResource {
	this := RestrictedResource{}
	return &this
}

// GetMethod returns the Method field value
func (o *RestrictedResource) GetMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *RestrictedResource) GetMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Method, true
}

// SetMethod sets field value
func (o *RestrictedResource) SetMethod(v string) {
	o.Method = v
}

// GetPath returns the Path field value
func (o *RestrictedResource) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *RestrictedResource) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *RestrictedResource) SetPath(v string) {
	o.Path = v
}

// GetDataElements returns the DataElements field value if set, zero value otherwise.
func (o *RestrictedResource) GetDataElements() []string {
	if o == nil || IsNil(o.DataElements) {
		var ret []string
		return ret
	}
	return o.DataElements
}

// GetDataElementsOk returns a tuple with the DataElements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestrictedResource) GetDataElementsOk() ([]string, bool) {
	if o == nil || IsNil(o.DataElements) {
		return nil, false
	}
	return o.DataElements, true
}

// HasDataElements returns a boolean if a field has been set.
func (o *RestrictedResource) HasDataElements() bool {
	if o != nil && !IsNil(o.DataElements) {
		return true
	}

	return false
}

// SetDataElements gets a reference to the given []string and assigns it to the DataElements field.
func (o *RestrictedResource) SetDataElements(v []string) {
	o.DataElements = v
}

func (o RestrictedResource) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RestrictedResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["method"] = o.Method
	toSerialize["path"] = o.Path
	if !IsNil(o.DataElements) {
		toSerialize["dataElements"] = o.DataElements
	}
	return toSerialize, nil
}

type NullableRestrictedResource struct {
	value *RestrictedResource
	isSet bool
}

func (v NullableRestrictedResource) Get() *RestrictedResource {
	return v.value
}

func (v *NullableRestrictedResource) Set(val *RestrictedResource) {
	v.value = val
	v.isSet = true
}

func (v NullableRestrictedResource) IsSet() bool {
	return v.isSet
}

func (v *NullableRestrictedResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRestrictedResource(val *RestrictedResource) *NullableRestrictedResource {
	return &NullableRestrictedResource{value: val, isSet: true}
}

func (v NullableRestrictedResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRestrictedResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
