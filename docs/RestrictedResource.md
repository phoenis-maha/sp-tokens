# RestrictedResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Method** | **string** | The HTTP method in the restricted resource. | 
**Path** | **string** | The path in the restricted resource. Here are some path examples: - &#x60;&#x60;&#x60;/orders/v0/orders&#x60;&#x60;&#x60;. For getting an RDT for the getOrders operation of the Orders API. For bulk orders. - &#x60;&#x60;&#x60;/orders/v0/orders/123-1234567-1234567&#x60;&#x60;&#x60;. For getting an RDT for the getOrder operation of the Orders API. For a specific order. - &#x60;&#x60;&#x60;/orders/v0/orders/123-1234567-1234567/orderItems&#x60;&#x60;&#x60;. For getting an RDT for the getOrderItems operation of the Orders API. For the order items in a specific order. - &#x60;&#x60;&#x60;/mfn/v0/shipments/FBA1234ABC5D&#x60;&#x60;&#x60;. For getting an RDT for the getShipment operation of the Shipping API. For a specific shipment. - &#x60;&#x60;&#x60;/mfn/v0/shipments/{shipmentId}&#x60;&#x60;&#x60;. For getting an RDT for the getShipment operation of the Shipping API. For any of a selling partner&#39;s shipments that you specify when you call the getShipment operation. | 
**DataElements** | Pointer to **[]string** | Indicates the type of Personally Identifiable Information requested. This parameter is required only when getting an RDT for use with the getOrder, getOrders, or getOrderItems operation of the Orders API. For more information, see the [Tokens API Use Case Guide](doc:tokens-api-use-case-guide). Possible values include: - **buyerInfo**. On the order level this includes general identifying information about the buyer and tax-related information. On the order item level this includes gift wrap information and custom order information, if available. - **shippingAddress**. This includes information for fulfilling orders. - **buyerTaxInformation**. This includes information for issuing tax invoices. | [optional] 

## Methods

### NewRestrictedResource

`func NewRestrictedResource(method string, path string, ) *RestrictedResource`

NewRestrictedResource instantiates a new RestrictedResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestrictedResourceWithDefaults

`func NewRestrictedResourceWithDefaults() *RestrictedResource`

NewRestrictedResourceWithDefaults instantiates a new RestrictedResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethod

`func (o *RestrictedResource) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *RestrictedResource) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *RestrictedResource) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetPath

`func (o *RestrictedResource) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *RestrictedResource) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *RestrictedResource) SetPath(v string)`

SetPath sets Path field to given value.


### GetDataElements

`func (o *RestrictedResource) GetDataElements() []string`

GetDataElements returns the DataElements field if non-nil, zero value otherwise.

### GetDataElementsOk

`func (o *RestrictedResource) GetDataElementsOk() (*[]string, bool)`

GetDataElementsOk returns a tuple with the DataElements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataElements

`func (o *RestrictedResource) SetDataElements(v []string)`

SetDataElements sets DataElements field to given value.

### HasDataElements

`func (o *RestrictedResource) HasDataElements() bool`

HasDataElements returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


