package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DomainDnsRecord 
type DomainDnsRecord struct {
    Entity
    // If false, this record must be configured by the customer at the DNS host for Microsoft Online Services to operate correctly with the domain.
    isOptional *bool
    // Value used when configuring the name of the DNS record at the DNS host.
    label *string
    // Indicates what type of DNS record this entity represents.The value can be one of the following: CName, Mx, Srv, Txt.
    recordType *string
    // Microsoft Online Service or feature that has a dependency on this DNS record.Can be one of the following values: null, Email, Sharepoint, EmailInternalRelayOnly, OfficeCommunicationsOnline, SharePointDefaultDomain, FullRedelegation, SharePointPublic, OrgIdAuthentication, Yammer, Intune.
    supportedService *string
    // Value to use when configuring the time-to-live (ttl) property of the DNS record at the DNS host. Not nullable.
    ttl *int32
}
// NewDomainDnsRecord instantiates a new domainDnsRecord and sets the default values.
func NewDomainDnsRecord()(*DomainDnsRecord) {
    m := &DomainDnsRecord{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDomainDnsRecordFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDomainDnsRecordFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.domainDnsCnameRecord":
                        return NewDomainDnsCnameRecord(), nil
                    case "#microsoft.graph.domainDnsMxRecord":
                        return NewDomainDnsMxRecord(), nil
                    case "#microsoft.graph.domainDnsSrvRecord":
                        return NewDomainDnsSrvRecord(), nil
                    case "#microsoft.graph.domainDnsTxtRecord":
                        return NewDomainDnsTxtRecord(), nil
                    case "#microsoft.graph.domainDnsUnavailableRecord":
                        return NewDomainDnsUnavailableRecord(), nil
                }
            }
        }
    }
    return NewDomainDnsRecord(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DomainDnsRecord) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["isOptional"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsOptional(val)
        }
        return nil
    }
    res["label"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLabel(val)
        }
        return nil
    }
    res["recordType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecordType(val)
        }
        return nil
    }
    res["supportedService"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSupportedService(val)
        }
        return nil
    }
    res["ttl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTtl(val)
        }
        return nil
    }
    return res
}
// GetIsOptional gets the isOptional property value. If false, this record must be configured by the customer at the DNS host for Microsoft Online Services to operate correctly with the domain.
func (m *DomainDnsRecord) GetIsOptional()(*bool) {
    return m.isOptional
}
// GetLabel gets the label property value. Value used when configuring the name of the DNS record at the DNS host.
func (m *DomainDnsRecord) GetLabel()(*string) {
    return m.label
}
// GetRecordType gets the recordType property value. Indicates what type of DNS record this entity represents.The value can be one of the following: CName, Mx, Srv, Txt.
func (m *DomainDnsRecord) GetRecordType()(*string) {
    return m.recordType
}
// GetSupportedService gets the supportedService property value. Microsoft Online Service or feature that has a dependency on this DNS record.Can be one of the following values: null, Email, Sharepoint, EmailInternalRelayOnly, OfficeCommunicationsOnline, SharePointDefaultDomain, FullRedelegation, SharePointPublic, OrgIdAuthentication, Yammer, Intune.
func (m *DomainDnsRecord) GetSupportedService()(*string) {
    return m.supportedService
}
// GetTtl gets the ttl property value. Value to use when configuring the time-to-live (ttl) property of the DNS record at the DNS host. Not nullable.
func (m *DomainDnsRecord) GetTtl()(*int32) {
    return m.ttl
}
// Serialize serializes information the current object
func (m *DomainDnsRecord) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("isOptional", m.GetIsOptional())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("label", m.GetLabel())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("recordType", m.GetRecordType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("supportedService", m.GetSupportedService())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("ttl", m.GetTtl())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetIsOptional sets the isOptional property value. If false, this record must be configured by the customer at the DNS host for Microsoft Online Services to operate correctly with the domain.
func (m *DomainDnsRecord) SetIsOptional(value *bool)() {
    m.isOptional = value
}
// SetLabel sets the label property value. Value used when configuring the name of the DNS record at the DNS host.
func (m *DomainDnsRecord) SetLabel(value *string)() {
    m.label = value
}
// SetRecordType sets the recordType property value. Indicates what type of DNS record this entity represents.The value can be one of the following: CName, Mx, Srv, Txt.
func (m *DomainDnsRecord) SetRecordType(value *string)() {
    m.recordType = value
}
// SetSupportedService sets the supportedService property value. Microsoft Online Service or feature that has a dependency on this DNS record.Can be one of the following values: null, Email, Sharepoint, EmailInternalRelayOnly, OfficeCommunicationsOnline, SharePointDefaultDomain, FullRedelegation, SharePointPublic, OrgIdAuthentication, Yammer, Intune.
func (m *DomainDnsRecord) SetSupportedService(value *string)() {
    m.supportedService = value
}
// SetTtl sets the ttl property value. Value to use when configuring the time-to-live (ttl) property of the DNS record at the DNS host. Not nullable.
func (m *DomainDnsRecord) SetTtl(value *int32)() {
    m.ttl = value
}
