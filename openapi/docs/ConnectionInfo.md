# ConnectionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**RemoteIpAddress** | Pointer to [**IPAddress**](IPAddress.md) |  | [optional] 
**RemotePort** | Pointer to **int32** |  | [optional] 
**LocalIpAddress** | Pointer to [**IPAddress**](IPAddress.md) |  | [optional] 
**LocalPort** | Pointer to **int32** |  | [optional] 
**ClientCertificate** | Pointer to [**X509Certificate2**](X509Certificate2.md) |  | [optional] 

## Methods

### NewConnectionInfo

`func NewConnectionInfo() *ConnectionInfo`

NewConnectionInfo instantiates a new ConnectionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionInfoWithDefaults

`func NewConnectionInfoWithDefaults() *ConnectionInfo`

NewConnectionInfoWithDefaults instantiates a new ConnectionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConnectionInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConnectionInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConnectionInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ConnectionInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ConnectionInfo) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ConnectionInfo) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetRemoteIpAddress

`func (o *ConnectionInfo) GetRemoteIpAddress() IPAddress`

GetRemoteIpAddress returns the RemoteIpAddress field if non-nil, zero value otherwise.

### GetRemoteIpAddressOk

`func (o *ConnectionInfo) GetRemoteIpAddressOk() (*IPAddress, bool)`

GetRemoteIpAddressOk returns a tuple with the RemoteIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteIpAddress

`func (o *ConnectionInfo) SetRemoteIpAddress(v IPAddress)`

SetRemoteIpAddress sets RemoteIpAddress field to given value.

### HasRemoteIpAddress

`func (o *ConnectionInfo) HasRemoteIpAddress() bool`

HasRemoteIpAddress returns a boolean if a field has been set.

### GetRemotePort

`func (o *ConnectionInfo) GetRemotePort() int32`

GetRemotePort returns the RemotePort field if non-nil, zero value otherwise.

### GetRemotePortOk

`func (o *ConnectionInfo) GetRemotePortOk() (*int32, bool)`

GetRemotePortOk returns a tuple with the RemotePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotePort

`func (o *ConnectionInfo) SetRemotePort(v int32)`

SetRemotePort sets RemotePort field to given value.

### HasRemotePort

`func (o *ConnectionInfo) HasRemotePort() bool`

HasRemotePort returns a boolean if a field has been set.

### GetLocalIpAddress

`func (o *ConnectionInfo) GetLocalIpAddress() IPAddress`

GetLocalIpAddress returns the LocalIpAddress field if non-nil, zero value otherwise.

### GetLocalIpAddressOk

`func (o *ConnectionInfo) GetLocalIpAddressOk() (*IPAddress, bool)`

GetLocalIpAddressOk returns a tuple with the LocalIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalIpAddress

`func (o *ConnectionInfo) SetLocalIpAddress(v IPAddress)`

SetLocalIpAddress sets LocalIpAddress field to given value.

### HasLocalIpAddress

`func (o *ConnectionInfo) HasLocalIpAddress() bool`

HasLocalIpAddress returns a boolean if a field has been set.

### GetLocalPort

`func (o *ConnectionInfo) GetLocalPort() int32`

GetLocalPort returns the LocalPort field if non-nil, zero value otherwise.

### GetLocalPortOk

`func (o *ConnectionInfo) GetLocalPortOk() (*int32, bool)`

GetLocalPortOk returns a tuple with the LocalPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalPort

`func (o *ConnectionInfo) SetLocalPort(v int32)`

SetLocalPort sets LocalPort field to given value.

### HasLocalPort

`func (o *ConnectionInfo) HasLocalPort() bool`

HasLocalPort returns a boolean if a field has been set.

### GetClientCertificate

`func (o *ConnectionInfo) GetClientCertificate() X509Certificate2`

GetClientCertificate returns the ClientCertificate field if non-nil, zero value otherwise.

### GetClientCertificateOk

`func (o *ConnectionInfo) GetClientCertificateOk() (*X509Certificate2, bool)`

GetClientCertificateOk returns a tuple with the ClientCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertificate

`func (o *ConnectionInfo) SetClientCertificate(v X509Certificate2)`

SetClientCertificate sets ClientCertificate field to given value.

### HasClientCertificate

`func (o *ConnectionInfo) HasClientCertificate() bool`

HasClientCertificate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


