# \AccountApi

All URIs are relative to *https://localhost:7121*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiAccountRegisterPost**](AccountApi.md#ApiAccountRegisterPost) | **Post** /api/account/register | Register to KangDroid Cloud Service.



## ApiAccountRegisterPost

> ApiAccountRegisterPost(ctx).AccountRegisterRequest(accountRegisterRequest).Execute()

Register to KangDroid Cloud Service.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountRegisterRequest := *openapiclient.NewAccountRegisterRequest() // AccountRegisterRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountApi.ApiAccountRegisterPost(context.Background()).AccountRegisterRequest(accountRegisterRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.ApiAccountRegisterPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountRegisterPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountRegisterRequest** | [**AccountRegisterRequest**](AccountRegisterRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

