package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DomainDnsCnameRecord 
type DomainDnsCnameRecord struct {
    DomainDnsRecord
    // The canonical name of the CNAME record. Used to configure the CNAME record at the DNS host.
    canonicalName *string
}
// NewDomainDnsCnameRecord instantiates a new DomainDnsCnameRecord and sets the default values.
func NewDomainDnsCnameRecord()(*DomainDnsCnameRecord) {
    m := &DomainDnsCnameRecord{
        DomainDnsRecord: *NewDomainDnsRecord(),
    }
    return m
}
// CreateDomainDnsCnameRecordFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDomainDnsCnameRecordFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDomainDnsCnameRecord(), nil
}
// GetCanonicalName gets the canonicalName property value. The canonical name of the CNAME record. Used to configure the CNAME record at the DNS host.
func (m *DomainDnsCnameRecord) GetCanonicalName()(*string) {
    return m.canonicalName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DomainDnsCnameRecord) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DomainDnsRecord.GetFieldDeserializers()
    res["canonicalName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCanonicalName(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *DomainDnsCnameRecord) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DomainDnsRecord.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("canonicalName", m.GetCanonicalName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCanonicalName sets the canonicalName property value. The canonical name of the CNAME record. Used to configure the CNAME record at the DNS host.
func (m *DomainDnsCnameRecord) SetCanonicalName(value *string)() {
    m.canonicalName = value
}
