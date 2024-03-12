//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdns

// ARecord - An A record.
type ARecord struct {
	// The IPv4 address of this A record.
	IPv4Address *string
}

// AaaaRecord - An AAAA record.
type AaaaRecord struct {
	// The IPv6 address of this AAAA record.
	IPv6Address *string
}

// CaaRecord - A CAA record.
type CaaRecord struct {
	// The flags for this CAA record as an integer between 0 and 255.
	Flags *int32

	// The tag for this CAA record.
	Tag *string

	// The value for this CAA record.
	Value *string
}

// CnameRecord - A CNAME record.
type CnameRecord struct {
	// The canonical name for this CNAME record.
	Cname *string
}

// MxRecord - An MX record.
type MxRecord struct {
	// The domain name of the mail host for this MX record.
	Exchange *string

	// The preference value for this MX record.
	Preference *int32
}

// NsRecord - An NS record.
type NsRecord struct {
	// The name server name for this NS record.
	Nsdname *string
}

// PtrRecord - A PTR record.
type PtrRecord struct {
	// The PTR target domain name for this PTR record.
	Ptrdname *string
}

// RecordSet - Describes a DNS record set (a collection of DNS records with the same name and type).
type RecordSet struct {
	// The etag of the record set.
	Etag *string

	// The properties of the record set.
	Properties *RecordSetProperties

	// READ-ONLY; The ID of the record set.
	ID *string

	// READ-ONLY; The name of the record set.
	Name *string

	// READ-ONLY; The type of the record set.
	Type *string
}

// RecordSetListResult - The response to a record set List operation.
type RecordSetListResult struct {
	// Information about the record sets in the response.
	Value []*RecordSet

	// READ-ONLY; The continuation token for the next page of results.
	NextLink *string
}

// RecordSetProperties - Represents the properties of the records in the record set.
type RecordSetProperties struct {
	// The list of A records in the record set.
	ARecords []*ARecord

	// The list of AAAA records in the record set.
	AaaaRecords []*AaaaRecord

	// The list of CAA records in the record set.
	CaaRecords []*CaaRecord

	// The CNAME record in the record set.
	CnameRecord *CnameRecord

	// The metadata attached to the record set.
	Metadata map[string]*string

	// The list of MX records in the record set.
	MxRecords []*MxRecord

	// The list of NS records in the record set.
	NsRecords []*NsRecord

	// The list of PTR records in the record set.
	PtrRecords []*PtrRecord

	// The SOA record in the record set.
	SoaRecord *SoaRecord

	// The list of SRV records in the record set.
	SrvRecords []*SrvRecord

	// The TTL (time-to-live) of the records in the record set.
	TTL *int64

	// A reference to an azure resource from where the dns resource value is taken.
	TargetResource *SubResource

	// The list of TXT records in the record set.
	TxtRecords []*TxtRecord

	// READ-ONLY; Fully qualified domain name of the record set.
	Fqdn *string

	// READ-ONLY; provisioning State of the record set.
	ProvisioningState *string
}

// RecordSetUpdateParameters - Parameters supplied to update a record set.
type RecordSetUpdateParameters struct {
	// Specifies information about the record set being updated.
	RecordSet *RecordSet
}

// Resource - Common properties of an Azure Resource Manager resource
type Resource struct {
	// REQUIRED; Resource location.
	Location *string

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; Resource ID.
	ID *string

	// READ-ONLY; Resource name.
	Name *string

	// READ-ONLY; Resource type.
	Type *string
}

// ResourceReference - Represents a single Azure resource and its referencing DNS records.
type ResourceReference struct {
	// A list of dns Records
	DNSResources []*SubResource

	// A reference to an azure resource from where the dns resource value is taken.
	TargetResource *SubResource
}

// ResourceReferenceRequest - Represents the properties of the Dns Resource Reference Request.
type ResourceReferenceRequest struct {
	// The properties of the Resource Reference Request.
	Properties *ResourceReferenceRequestProperties
}

// ResourceReferenceRequestProperties - Represents the properties of the Dns Resource Reference Request.
type ResourceReferenceRequestProperties struct {
	// A list of references to azure resources for which referencing dns records need to be queried.
	TargetResources []*SubResource
}

// ResourceReferenceResult - Represents the properties of the Dns Resource Reference Result.
type ResourceReferenceResult struct {
	// The result of dns resource reference request. Returns a list of dns resource references for each of the azure resource
	// in the request.
	Properties *ResourceReferenceResultProperties
}

// ResourceReferenceResultProperties - The result of dns resource reference request. Returns a list of dns resource references
// for each of the azure resource in the request.
type ResourceReferenceResultProperties struct {
	// The result of dns resource reference request. A list of dns resource references for each of the azure resource in the request
	DNSResourceReferences []*ResourceReference
}

// SoaRecord - An SOA record.
type SoaRecord struct {
	// The email contact for this SOA record.
	Email *string

	// The expire time for this SOA record.
	ExpireTime *int64

	// The domain name of the authoritative name server for this SOA record.
	Host *string

	// The minimum value for this SOA record. By convention this is used to determine the negative caching duration.
	MinimumTTL *int64

	// The refresh value for this SOA record.
	RefreshTime *int64

	// The retry time for this SOA record.
	RetryTime *int64

	// The serial number for this SOA record.
	SerialNumber *int64
}

// SrvRecord - An SRV record.
type SrvRecord struct {
	// The port value for this SRV record.
	Port *int32

	// The priority value for this SRV record.
	Priority *int32

	// The target domain name for this SRV record.
	Target *string

	// The weight value for this SRV record.
	Weight *int32
}

// SubResource - A reference to a another resource
type SubResource struct {
	// Resource Id.
	ID *string
}

// TxtRecord - A TXT record.
type TxtRecord struct {
	// The text value of this TXT record.
	Value []*string
}

// Zone - Describes a DNS zone.
type Zone struct {
	// REQUIRED; Resource location.
	Location *string

	// The etag of the zone.
	Etag *string

	// The properties of the zone.
	Properties *ZoneProperties

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; Resource ID.
	ID *string

	// READ-ONLY; Resource name.
	Name *string

	// READ-ONLY; Resource type.
	Type *string
}

// ZoneListResult - The response to a Zone List or ListAll operation.
type ZoneListResult struct {
	// Information about the DNS zones.
	Value []*Zone

	// READ-ONLY; The continuation token for the next page of results.
	NextLink *string
}

// ZoneProperties - Represents the properties of the zone.
type ZoneProperties struct {
	// A list of references to virtual networks that register hostnames in this DNS zone. This is a only when ZoneType is Private.
	RegistrationVirtualNetworks []*SubResource

	// A list of references to virtual networks that resolve records in this DNS zone. This is a only when ZoneType is Private.
	ResolutionVirtualNetworks []*SubResource

	// The type of this DNS zone (Public or Private).
	ZoneType *ZoneType

	// READ-ONLY; The maximum number of record sets that can be created in this DNS zone. This is a read-only property and any
	// attempt to set this value will be ignored.
	MaxNumberOfRecordSets *int64

	// READ-ONLY; The maximum number of records per record set that can be created in this DNS zone. This is a read-only property
	// and any attempt to set this value will be ignored.
	MaxNumberOfRecordsPerRecordSet *int64

	// READ-ONLY; The name servers for this DNS zone. This is a read-only property and any attempt to set this value will be ignored.
	NameServers []*string

	// READ-ONLY; The current number of record sets in this DNS zone. This is a read-only property and any attempt to set this
	// value will be ignored.
	NumberOfRecordSets *int64
}

// ZoneUpdate - Describes a request to update a DNS zone.
type ZoneUpdate struct {
	// Resource tags.
	Tags map[string]*string
}