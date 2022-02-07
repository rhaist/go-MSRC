# \GetSecurityUpdatesApi

All URIs are relative to *https://api.msrc.microsoft.com/cvrf/v2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CvrfGetCvrf**](GetSecurityUpdatesApi.md#CvrfGetCvrf) | **Get** /cvrf/{id} | Get security update details in CVRF format
[**UpdatesGetUpdates**](GetSecurityUpdatesApi.md#UpdatesGetUpdates) | **Get** /updates | Get all security update summaries
[**UpdatesGetUpdatesByKey**](GetSecurityUpdatesApi.md#UpdatesGetUpdatesByKey) | **Get** /Updates(&#39;{key}&#39;) | Get security update summaries by key



## CvrfGetCvrf

> Cvrfdoc CvrfGetCvrf(ctx, id).Execute()

Get security update details in CVRF format



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
    id := "id_example" // string | CVRF document ID (yyyy-mmm)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GetSecurityUpdatesApi.CvrfGetCvrf(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GetSecurityUpdatesApi.CvrfGetCvrf``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CvrfGetCvrf`: Cvrfdoc
    fmt.Fprintf(os.Stdout, "Response from `GetSecurityUpdatesApi.CvrfGetCvrf`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | CVRF document ID (yyyy-mmm) | 

### Other Parameters

Other parameters are passed through a pointer to a apiCvrfGetCvrfRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Cvrfdoc**](Cvrfdoc.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatesGetUpdates

> InlineResponse200 UpdatesGetUpdates(ctx).Execute()

Get all security update summaries



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GetSecurityUpdatesApi.UpdatesGetUpdates(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GetSecurityUpdatesApi.UpdatesGetUpdates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatesGetUpdates`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `GetSecurityUpdatesApi.UpdatesGetUpdates`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatesGetUpdatesRequest struct via the builder pattern


### Return type

[**InlineResponse200**](InlineResponse200.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatesGetUpdatesByKey

> InlineResponse200 UpdatesGetUpdatesByKey(ctx, key).Execute()

Get security update summaries by key



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
    key := "key_example" // string | update ID (yyyy-mmm), vulnerability ID (CVE number), or year (yyyy)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GetSecurityUpdatesApi.UpdatesGetUpdatesByKey(context.Background(), key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GetSecurityUpdatesApi.UpdatesGetUpdatesByKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatesGetUpdatesByKey`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `GetSecurityUpdatesApi.UpdatesGetUpdatesByKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | update ID (yyyy-mmm), vulnerability ID (CVE number), or year (yyyy) | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatesGetUpdatesByKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse200**](InlineResponse200.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

