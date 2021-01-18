# \EngageTheMicrosoftSecurityResponseCenterMSRCApi

All URIs are relative to *https://api.msrc.microsoft.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CarsPostCarsReport**](EngageTheMicrosoftSecurityResponseCenterMSRCApi.md#CarsPostCarsReport) | **Post** /engage/cars | Submit an abuse report



## CarsPostCarsReport

> CarsPostCarsReport(ctx).CarsDto(carsDto).ApiVersion(apiVersion).Execute()

Submit an abuse report



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
    carsDto := *openapiclient.NewCarsDto(false, *openapiclient.NewReporterInfo("ReporterEmail_example", "ReporterName_example", "DiscloseEmail_example", "ReporterNotes_example"), []openapiclient.Report{*openapiclient.NewReport(*openapiclient.NewReport1("ReportNotes_example", []openapiclient.Threat{*openapiclient.NewThreat(*openapiclient.NewThreat1("ThreatType_example", "ThreatSubType_example", "Date_example", "Time_example", "TimeZone_example"))}))}) // CarsDto | CARS-formatted abuse report
    apiVersion := "apiVersion_example" // string | API version (yyyy) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EngageTheMicrosoftSecurityResponseCenterMSRCApi.CarsPostCarsReport(context.Background()).CarsDto(carsDto).ApiVersion(apiVersion).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `EngageTheMicrosoftSecurityResponseCenterMSRCApi.CarsPostCarsReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCarsPostCarsReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **carsDto** | [**CarsDto**](CarsDto.md) | CARS-formatted abuse report | 
 **apiVersion** | **string** | API version (yyyy) | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

