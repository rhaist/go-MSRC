# \GetSecurityUpdatesAPI

All URIs are relative to *https://api.msrc.microsoft.com/cvrf/v2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CvrfGetCvrf**](GetSecurityUpdatesAPI.md#CvrfGetCvrf) | **Get** /cvrf/{id} | Get security update details in CVRF format
[**UpdatesGetUpdates**](GetSecurityUpdatesAPI.md#UpdatesGetUpdates) | **Get** /updates | Get all security update summaries
[**UpdatesGetUpdatesByKey**](GetSecurityUpdatesAPI.md#UpdatesGetUpdatesByKey) | **Get** /Updates(&#39;{key}&#39;) | Get security update summaries by key



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
	openapiclient "github.com/rhaist/go-MSRC/openapi"
)

func main() {
	id := "id_example" // string | CVRF document ID (yyyy-mmm)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GetSecurityUpdatesAPI.CvrfGetCvrf(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GetSecurityUpdatesAPI.CvrfGetCvrf``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CvrfGetCvrf`: Cvrfdoc
	fmt.Fprintf(os.Stdout, "Response from `GetSecurityUpdatesAPI.CvrfGetCvrf`: %v\n", resp)
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

> UpdatesGetUpdates200Response UpdatesGetUpdates(ctx).Execute()

Get all security update summaries



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/rhaist/go-MSRC/openapi"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GetSecurityUpdatesAPI.UpdatesGetUpdates(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GetSecurityUpdatesAPI.UpdatesGetUpdates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatesGetUpdates`: UpdatesGetUpdates200Response
	fmt.Fprintf(os.Stdout, "Response from `GetSecurityUpdatesAPI.UpdatesGetUpdates`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatesGetUpdatesRequest struct via the builder pattern


### Return type

[**UpdatesGetUpdates200Response**](UpdatesGetUpdates200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatesGetUpdatesByKey

> UpdatesGetUpdates200Response UpdatesGetUpdatesByKey(ctx, key).Execute()

Get security update summaries by key



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/rhaist/go-MSRC/openapi"
)

func main() {
	key := "key_example" // string | update ID (yyyy-mmm), vulnerability ID (CVE number), or year (yyyy)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GetSecurityUpdatesAPI.UpdatesGetUpdatesByKey(context.Background(), key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GetSecurityUpdatesAPI.UpdatesGetUpdatesByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatesGetUpdatesByKey`: UpdatesGetUpdates200Response
	fmt.Fprintf(os.Stdout, "Response from `GetSecurityUpdatesAPI.UpdatesGetUpdatesByKey`: %v\n", resp)
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

[**UpdatesGetUpdates200Response**](UpdatesGetUpdates200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

