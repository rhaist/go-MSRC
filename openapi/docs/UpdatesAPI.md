# \UpdatesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Updates**](UpdatesAPI.md#Updates) | **Get** /Updates | Get all security update summaries
[**UpdatesCount**](UpdatesAPI.md#UpdatesCount) | **Get** /Updates/$count | Get all security update summaries
[**UpdatesGet**](UpdatesAPI.md#UpdatesGet) | **Get** /updates | Get all security update summaries
[**UpdatesKey**](UpdatesAPI.md#UpdatesKey) | **Get** /Updates({key}) | Get security update summaries by key
[**UpdatesKeyGet**](UpdatesAPI.md#UpdatesKeyGet) | **Get** /updates({key}) | Get security update summaries by key
[**UpdatesKeyGet_0**](UpdatesAPI.md#UpdatesKeyGet_0) | **Get** /updates/{key} | Get security update summaries by key
[**UpdatesKey_0**](UpdatesAPI.md#UpdatesKey_0) | **Get** /Updates/{key} | Get security update summaries by key



## Updates

> Updates(ctx).Options(options).Execute()

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
	options := *openapiclient.NewUpdateODataQueryOptions() // UpdateODataQueryOptions | OData query options (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UpdatesAPI.Updates(context.Background()).Options(options).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UpdatesAPI.Updates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **options** | [**UpdateODataQueryOptions**](UpdateODataQueryOptions.md) | OData query options | 

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


## UpdatesCount

> UpdatesCount(ctx).Options(options).Execute()

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
	options := *openapiclient.NewUpdateODataQueryOptions() // UpdateODataQueryOptions | OData query options (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UpdatesAPI.UpdatesCount(context.Background()).Options(options).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UpdatesAPI.UpdatesCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdatesCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **options** | [**UpdateODataQueryOptions**](UpdateODataQueryOptions.md) | OData query options | 

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


## UpdatesGet

> UpdatesGet(ctx).Options(options).Execute()

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
	options := *openapiclient.NewUpdateODataQueryOptions() // UpdateODataQueryOptions | OData query options (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UpdatesAPI.UpdatesGet(context.Background()).Options(options).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UpdatesAPI.UpdatesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdatesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **options** | [**UpdateODataQueryOptions**](UpdateODataQueryOptions.md) | OData query options | 

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


## UpdatesKey

> UpdatesKey(ctx).Key(key).Options(options).Execute()

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
	key := "key_example" // string | update ID (yyyy-mmm), vulnerability ID (CVE number), or year (yyyy) (optional)
	options := *openapiclient.NewUpdateODataQueryOptions() // UpdateODataQueryOptions | OData query options (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UpdatesAPI.UpdatesKey(context.Background()).Key(key).Options(options).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UpdatesAPI.UpdatesKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdatesKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **key** | **string** | update ID (yyyy-mmm), vulnerability ID (CVE number), or year (yyyy) | 
 **options** | [**UpdateODataQueryOptions**](UpdateODataQueryOptions.md) | OData query options | 

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


## UpdatesKeyGet

> UpdatesKeyGet(ctx).Key(key).Options(options).Execute()

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
	key := "key_example" // string | update ID (yyyy-mmm), vulnerability ID (CVE number), or year (yyyy) (optional)
	options := *openapiclient.NewUpdateODataQueryOptions() // UpdateODataQueryOptions | OData query options (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UpdatesAPI.UpdatesKeyGet(context.Background()).Key(key).Options(options).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UpdatesAPI.UpdatesKeyGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdatesKeyGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **key** | **string** | update ID (yyyy-mmm), vulnerability ID (CVE number), or year (yyyy) | 
 **options** | [**UpdateODataQueryOptions**](UpdateODataQueryOptions.md) | OData query options | 

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


## UpdatesKeyGet_0

> UpdatesKeyGet_0(ctx).Key(key).Options(options).Execute()

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
	key := "key_example" // string | update ID (yyyy-mmm), vulnerability ID (CVE number), or year (yyyy) (optional)
	options := *openapiclient.NewUpdateODataQueryOptions() // UpdateODataQueryOptions | OData query options (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UpdatesAPI.UpdatesKeyGet_0(context.Background()).Key(key).Options(options).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UpdatesAPI.UpdatesKeyGet_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdatesKeyGet_1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **key** | **string** | update ID (yyyy-mmm), vulnerability ID (CVE number), or year (yyyy) | 
 **options** | [**UpdateODataQueryOptions**](UpdateODataQueryOptions.md) | OData query options | 

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


## UpdatesKey_0

> UpdatesKey_0(ctx).Key(key).Options(options).Execute()

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
	key := "key_example" // string | update ID (yyyy-mmm), vulnerability ID (CVE number), or year (yyyy) (optional)
	options := *openapiclient.NewUpdateODataQueryOptions() // UpdateODataQueryOptions | OData query options (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UpdatesAPI.UpdatesKey_0(context.Background()).Key(key).Options(options).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UpdatesAPI.UpdatesKey_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdatesKey_2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **key** | **string** | update ID (yyyy-mmm), vulnerability ID (CVE number), or year (yyyy) | 
 **options** | [**UpdateODataQueryOptions**](UpdateODataQueryOptions.md) | OData query options | 

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

