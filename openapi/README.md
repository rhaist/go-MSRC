# Go API client for openapi

This RESTful API can be used to engage the Microsoft Security Response Center (MSRC) in the following ways:

- Get security update summaries and details using the [Common Vulnerability Reporting Framework](https://www.icasi.org/cvrf) (CVRF).

- Report suspected cyberattacks or abuse originating from Microsoft Online Services.

- Notify Microsoft of any planned penetration tests against your Azure assets.

**Sample client code** is available on the Microsoft Security [Updates](https://github.com/microsoft/MSRC-Microsoft-Security-Updates-API) and [Engage](https://github.com/Microsoft/MSRC-Microsoft-Engage-API) 
Github repositories.

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: v1
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import sw "./openapi"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identifield by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), sw.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), sw.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.msrc.microsoft.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*EngageTheMicrosoftSecurityResponseCenterMSRCApi* | [**CarsPostCarsReport**](docs/EngageTheMicrosoftSecurityResponseCenterMSRCApi.md#carspostcarsreport) | **Post** /engage/cars | Submit an abuse report
*GetSecurityUpdatesApi* | [**CvrfGetCvrf**](docs/GetSecurityUpdatesApi.md#cvrfgetcvrf) | **Get** /cvrf/{id} | Get security update details in CVRF format
*GetSecurityUpdatesApi* | [**UpdatesGetUpdates**](docs/GetSecurityUpdatesApi.md#updatesgetupdates) | **Get** /updates | Get all security update summaries
*GetSecurityUpdatesApi* | [**UpdatesGetUpdatesByKey**](docs/GetSecurityUpdatesApi.md#updatesgetupdatesbykey) | **Get** /Updates(&#39;{key}&#39;) | Get security update summaries by key


## Documentation For Models

 - [AffectedFile](docs/AffectedFile.md)
 - [BranchType](docs/BranchType.md)
 - [CarsDto](docs/CarsDto.md)
 - [Cvrfdoc](docs/Cvrfdoc.md)
 - [CvrfdocAcknowledgment](docs/CvrfdocAcknowledgment.md)
 - [CvrfdocAcknowledgmentDescription](docs/CvrfdocAcknowledgmentDescription.md)
 - [CvrfdocAcknowledgmentName](docs/CvrfdocAcknowledgmentName.md)
 - [CvrfdocAcknowledgmentOrganization](docs/CvrfdocAcknowledgmentOrganization.md)
 - [CvrfdocAggregateSeverity](docs/CvrfdocAggregateSeverity.md)
 - [CvrfdocDocumentDistribution](docs/CvrfdocDocumentDistribution.md)
 - [CvrfdocDocumentPublisher](docs/CvrfdocDocumentPublisher.md)
 - [CvrfdocDocumentPublisherContactDetails](docs/CvrfdocDocumentPublisherContactDetails.md)
 - [CvrfdocDocumentPublisherIssuingAuthority](docs/CvrfdocDocumentPublisherIssuingAuthority.md)
 - [CvrfdocDocumentTitle](docs/CvrfdocDocumentTitle.md)
 - [CvrfdocDocumentTracking](docs/CvrfdocDocumentTracking.md)
 - [CvrfdocDocumentTrackingGenerator](docs/CvrfdocDocumentTrackingGenerator.md)
 - [CvrfdocDocumentTrackingGeneratorEngine](docs/CvrfdocDocumentTrackingGeneratorEngine.md)
 - [CvrfdocDocumentTrackingIdentification](docs/CvrfdocDocumentTrackingIdentification.md)
 - [CvrfdocDocumentTrackingIdentificationAlias](docs/CvrfdocDocumentTrackingIdentificationAlias.md)
 - [CvrfdocDocumentTrackingIdentificationID](docs/CvrfdocDocumentTrackingIdentificationID.md)
 - [CvrfdocDocumentTrackingRevision](docs/CvrfdocDocumentTrackingRevision.md)
 - [CvrfdocDocumentTrackingRevisionDescription](docs/CvrfdocDocumentTrackingRevisionDescription.md)
 - [CvrfdocDocumentType](docs/CvrfdocDocumentType.md)
 - [CvrfdocNote](docs/CvrfdocNote.md)
 - [CvrfdocReference](docs/CvrfdocReference.md)
 - [CvrfdocReferenceDescription](docs/CvrfdocReferenceDescription.md)
 - [FullProductName](docs/FullProductName.md)
 - [InlineResponse200](docs/InlineResponse200.md)
 - [ProductTree](docs/ProductTree.md)
 - [ProductTreeGroup](docs/ProductTreeGroup.md)
 - [ProductTreeGroupDescription](docs/ProductTreeGroupDescription.md)
 - [ProductTreeRelationship](docs/ProductTreeRelationship.md)
 - [Report](docs/Report.md)
 - [Report1](docs/Report1.md)
 - [ReporterInfo](docs/ReporterInfo.md)
 - [RequesterInfo](docs/RequesterInfo.md)
 - [SubscriptionInfo](docs/SubscriptionInfo.md)
 - [TestInfo](docs/TestInfo.md)
 - [Threat](docs/Threat.md)
 - [Threat1](docs/Threat1.md)
 - [Update](docs/Update.md)
 - [Vulnerability](docs/Vulnerability.md)
 - [VulnerabilityAcknowledgment](docs/VulnerabilityAcknowledgment.md)
 - [VulnerabilityAcknowledgmentDescription](docs/VulnerabilityAcknowledgmentDescription.md)
 - [VulnerabilityAcknowledgmentName](docs/VulnerabilityAcknowledgmentName.md)
 - [VulnerabilityAcknowledgmentOrganization](docs/VulnerabilityAcknowledgmentOrganization.md)
 - [VulnerabilityCWE](docs/VulnerabilityCWE.md)
 - [VulnerabilityID](docs/VulnerabilityID.md)
 - [VulnerabilityInvolvement](docs/VulnerabilityInvolvement.md)
 - [VulnerabilityInvolvementDescription](docs/VulnerabilityInvolvementDescription.md)
 - [VulnerabilityNote](docs/VulnerabilityNote.md)
 - [VulnerabilityReference](docs/VulnerabilityReference.md)
 - [VulnerabilityReferenceDescription](docs/VulnerabilityReferenceDescription.md)
 - [VulnerabilityRemediation](docs/VulnerabilityRemediation.md)
 - [VulnerabilityRemediationDescription](docs/VulnerabilityRemediationDescription.md)
 - [VulnerabilityRemediationEntitlement](docs/VulnerabilityRemediationEntitlement.md)
 - [VulnerabilityRemediationRestartRequired](docs/VulnerabilityRemediationRestartRequired.md)
 - [VulnerabilityScoreSet](docs/VulnerabilityScoreSet.md)
 - [VulnerabilityStatus](docs/VulnerabilityStatus.md)
 - [VulnerabilityThreat](docs/VulnerabilityThreat.md)
 - [VulnerabilityThreatDescription](docs/VulnerabilityThreatDescription.md)
 - [VulnerabilityTitle](docs/VulnerabilityTitle.md)


## Documentation For Authorization



### apiKeyHeader

- **Type**: API key
- **API key parameter name**: api-key
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: api-key and passed in as the auth context for each request.


### oauth2


- **Type**: OAuth
- **Flow**: accessCode
- **Authorization URL**: https://login.windows.net/72f988bf-86f1-41af-91ab-2d7cd011db47/oauth2/authorize
- **Scopes**: N/A

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.

```golang
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, sw.ContextOAuth2, tokenSource)
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author


