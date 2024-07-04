# \CvrfAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CvrfIdGet**](CvrfAPI.md#CvrfIdGet) | **Get** /cvrf/{id} | Get security update details in CVRF format
[**CvrfIdGet_0**](CvrfAPI.md#CvrfIdGet_0) | **Get** /cvrf({id}) | Get security update details in CVRF format



## CvrfIdGet

> CvrfIdGet(ctx, id).Execute()

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
	r, err := apiClient.CvrfAPI.CvrfIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CvrfAPI.CvrfIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | CVRF document ID (yyyy-mmm) | 

### Other Parameters

Other parameters are passed through a pointer to a apiCvrfIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CvrfIdGet_0

> CvrfIdGet_0(ctx, id).Execute()

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
	r, err := apiClient.CvrfAPI.CvrfIdGet_0(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CvrfAPI.CvrfIdGet_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | CVRF document ID (yyyy-mmm) | 

### Other Parameters

Other parameters are passed through a pointer to a apiCvrfIdGet_1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

