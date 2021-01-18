# Threat1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ThreatType** | **string** | threat type: activity or content | 
**ThreatSubType** | **string** | threat subtype (depends on threat type) | 
**SourceIp** | Pointer to **string** | IP address sending unwanted traffic (required for activity threats) | [optional] 
**DestinationIp** | Pointer to **string** | IP address receiving unwanted traffic (required for activity threats) | [optional] 
**SourcePort** | Pointer to **string** | port sending unwanted traffic | [optional] 
**DestinationPort** | Pointer to **string** | port receiving unwanted traffic | [optional] 
**SourceUrl** | Pointer to **string** | URL hosting abusive content (required for content threats) | [optional] 
**DestinationUrl** | Pointer to **string** | URL receiving abusive content (required for content threats) | [optional] 
**Protocol** | Pointer to **string** | [internet protocol number](http://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml) of the protocol being used to send or serve abusive traffic or data | [optional] 
**ByteCount** | Pointer to **string** | total number of bytes received for reported traffic (activity threats only) | [optional] 
**PacketCount** | Pointer to **string** | total number of packets received for reported traffic (activity threats only) | [optional] 
**Date** | **string** | date of abuse | 
**Time** | **string** | time of abuse | 
**TimeZone** | **string** | reporter time zone | 
**Sample** | Pointer to **string** | sample data/artifacts relating to abuse | [optional] 
**SampleType** | Pointer to **string** | [Media Type](https://www.iana.org/assignments/media-types/media-types.xhtml) of sample (required if sample is provided) | [optional] 

## Methods

### NewThreat1

`func NewThreat1(threatType string, threatSubType string, date string, time string, timeZone string, ) *Threat1`

NewThreat1 instantiates a new Threat1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThreat1WithDefaults

`func NewThreat1WithDefaults() *Threat1`

NewThreat1WithDefaults instantiates a new Threat1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetThreatType

`func (o *Threat1) GetThreatType() string`

GetThreatType returns the ThreatType field if non-nil, zero value otherwise.

### GetThreatTypeOk

`func (o *Threat1) GetThreatTypeOk() (*string, bool)`

GetThreatTypeOk returns a tuple with the ThreatType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatType

`func (o *Threat1) SetThreatType(v string)`

SetThreatType sets ThreatType field to given value.


### GetThreatSubType

`func (o *Threat1) GetThreatSubType() string`

GetThreatSubType returns the ThreatSubType field if non-nil, zero value otherwise.

### GetThreatSubTypeOk

`func (o *Threat1) GetThreatSubTypeOk() (*string, bool)`

GetThreatSubTypeOk returns a tuple with the ThreatSubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatSubType

`func (o *Threat1) SetThreatSubType(v string)`

SetThreatSubType sets ThreatSubType field to given value.


### GetSourceIp

`func (o *Threat1) GetSourceIp() string`

GetSourceIp returns the SourceIp field if non-nil, zero value otherwise.

### GetSourceIpOk

`func (o *Threat1) GetSourceIpOk() (*string, bool)`

GetSourceIpOk returns a tuple with the SourceIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIp

`func (o *Threat1) SetSourceIp(v string)`

SetSourceIp sets SourceIp field to given value.

### HasSourceIp

`func (o *Threat1) HasSourceIp() bool`

HasSourceIp returns a boolean if a field has been set.

### GetDestinationIp

`func (o *Threat1) GetDestinationIp() string`

GetDestinationIp returns the DestinationIp field if non-nil, zero value otherwise.

### GetDestinationIpOk

`func (o *Threat1) GetDestinationIpOk() (*string, bool)`

GetDestinationIpOk returns a tuple with the DestinationIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationIp

`func (o *Threat1) SetDestinationIp(v string)`

SetDestinationIp sets DestinationIp field to given value.

### HasDestinationIp

`func (o *Threat1) HasDestinationIp() bool`

HasDestinationIp returns a boolean if a field has been set.

### GetSourcePort

`func (o *Threat1) GetSourcePort() string`

GetSourcePort returns the SourcePort field if non-nil, zero value otherwise.

### GetSourcePortOk

`func (o *Threat1) GetSourcePortOk() (*string, bool)`

GetSourcePortOk returns a tuple with the SourcePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePort

`func (o *Threat1) SetSourcePort(v string)`

SetSourcePort sets SourcePort field to given value.

### HasSourcePort

`func (o *Threat1) HasSourcePort() bool`

HasSourcePort returns a boolean if a field has been set.

### GetDestinationPort

`func (o *Threat1) GetDestinationPort() string`

GetDestinationPort returns the DestinationPort field if non-nil, zero value otherwise.

### GetDestinationPortOk

`func (o *Threat1) GetDestinationPortOk() (*string, bool)`

GetDestinationPortOk returns a tuple with the DestinationPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationPort

`func (o *Threat1) SetDestinationPort(v string)`

SetDestinationPort sets DestinationPort field to given value.

### HasDestinationPort

`func (o *Threat1) HasDestinationPort() bool`

HasDestinationPort returns a boolean if a field has been set.

### GetSourceUrl

`func (o *Threat1) GetSourceUrl() string`

GetSourceUrl returns the SourceUrl field if non-nil, zero value otherwise.

### GetSourceUrlOk

`func (o *Threat1) GetSourceUrlOk() (*string, bool)`

GetSourceUrlOk returns a tuple with the SourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUrl

`func (o *Threat1) SetSourceUrl(v string)`

SetSourceUrl sets SourceUrl field to given value.

### HasSourceUrl

`func (o *Threat1) HasSourceUrl() bool`

HasSourceUrl returns a boolean if a field has been set.

### GetDestinationUrl

`func (o *Threat1) GetDestinationUrl() string`

GetDestinationUrl returns the DestinationUrl field if non-nil, zero value otherwise.

### GetDestinationUrlOk

`func (o *Threat1) GetDestinationUrlOk() (*string, bool)`

GetDestinationUrlOk returns a tuple with the DestinationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationUrl

`func (o *Threat1) SetDestinationUrl(v string)`

SetDestinationUrl sets DestinationUrl field to given value.

### HasDestinationUrl

`func (o *Threat1) HasDestinationUrl() bool`

HasDestinationUrl returns a boolean if a field has been set.

### GetProtocol

`func (o *Threat1) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *Threat1) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *Threat1) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *Threat1) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetByteCount

`func (o *Threat1) GetByteCount() string`

GetByteCount returns the ByteCount field if non-nil, zero value otherwise.

### GetByteCountOk

`func (o *Threat1) GetByteCountOk() (*string, bool)`

GetByteCountOk returns a tuple with the ByteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByteCount

`func (o *Threat1) SetByteCount(v string)`

SetByteCount sets ByteCount field to given value.

### HasByteCount

`func (o *Threat1) HasByteCount() bool`

HasByteCount returns a boolean if a field has been set.

### GetPacketCount

`func (o *Threat1) GetPacketCount() string`

GetPacketCount returns the PacketCount field if non-nil, zero value otherwise.

### GetPacketCountOk

`func (o *Threat1) GetPacketCountOk() (*string, bool)`

GetPacketCountOk returns a tuple with the PacketCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPacketCount

`func (o *Threat1) SetPacketCount(v string)`

SetPacketCount sets PacketCount field to given value.

### HasPacketCount

`func (o *Threat1) HasPacketCount() bool`

HasPacketCount returns a boolean if a field has been set.

### GetDate

`func (o *Threat1) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *Threat1) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *Threat1) SetDate(v string)`

SetDate sets Date field to given value.


### GetTime

`func (o *Threat1) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *Threat1) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *Threat1) SetTime(v string)`

SetTime sets Time field to given value.


### GetTimeZone

`func (o *Threat1) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *Threat1) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *Threat1) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.


### GetSample

`func (o *Threat1) GetSample() string`

GetSample returns the Sample field if non-nil, zero value otherwise.

### GetSampleOk

`func (o *Threat1) GetSampleOk() (*string, bool)`

GetSampleOk returns a tuple with the Sample field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSample

`func (o *Threat1) SetSample(v string)`

SetSample sets Sample field to given value.

### HasSample

`func (o *Threat1) HasSample() bool`

HasSample returns a boolean if a field has been set.

### GetSampleType

`func (o *Threat1) GetSampleType() string`

GetSampleType returns the SampleType field if non-nil, zero value otherwise.

### GetSampleTypeOk

`func (o *Threat1) GetSampleTypeOk() (*string, bool)`

GetSampleTypeOk returns a tuple with the SampleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleType

`func (o *Threat1) SetSampleType(v string)`

SetSampleType sets SampleType field to given value.

### HasSampleType

`func (o *Threat1) HasSampleType() bool`

HasSampleType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


