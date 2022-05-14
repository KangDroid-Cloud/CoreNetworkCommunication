# \AuthApi

All URIs are relative to *https://localhost:7121*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiAuthLoginPost**](AuthApi.md#ApiAuthLoginPost) | **Post** /api/auth/login | Login to KangDroid Cloud Service.



## ApiAuthLoginPost

> LoginResponse ApiAuthLoginPost(ctx).LoginRequest(loginRequest).Execute()

Login to KangDroid Cloud Service.

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
    loginRequest := *openapiclient.NewLoginRequest() // LoginRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthApi.ApiAuthLoginPost(context.Background()).LoginRequest(loginRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.ApiAuthLoginPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAuthLoginPost`: LoginResponse
    fmt.Fprintf(os.Stdout, "Response from `AuthApi.ApiAuthLoginPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAuthLoginPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **loginRequest** | [**LoginRequest**](LoginRequest.md) |  | 

### Return type

[**LoginResponse**](LoginResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

