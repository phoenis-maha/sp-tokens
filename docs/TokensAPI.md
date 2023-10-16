# \TokensAPI

All URIs are relative to *https://sellingpartnerapi-na.amazon.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRestrictedDataToken**](TokensAPI.md#CreateRestrictedDataToken) | **Post** /tokens/2021-03-01/restrictedDataToken | 



## CreateRestrictedDataToken

> CreateRestrictedDataTokenResponse CreateRestrictedDataToken(ctx).Body(body).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    body := *openapiclient.NewCreateRestrictedDataTokenRequest([]openapiclient.RestrictedResource{*openapiclient.NewRestrictedResource("Method_example", "Path_example")}) // CreateRestrictedDataTokenRequest | The restricted data token request details.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TokensAPI.CreateRestrictedDataToken(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokensAPI.CreateRestrictedDataToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRestrictedDataToken`: CreateRestrictedDataTokenResponse
    fmt.Fprintf(os.Stdout, "Response from `TokensAPI.CreateRestrictedDataToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRestrictedDataTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateRestrictedDataTokenRequest**](CreateRestrictedDataTokenRequest.md) | The restricted data token request details. | 

### Return type

[**CreateRestrictedDataTokenResponse**](CreateRestrictedDataTokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

