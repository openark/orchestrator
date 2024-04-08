package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DomainDnsUnavailableRecord 
type DomainDnsUnavailableRecord struct {
    DomainDnsRecord
    // Provides the reason why the DomainDnsUnavailableRecord entity is returned.
    description *string
}
// NewDomainDnsUnavailableRecord instantiates a new DomainDnsUnavailableRecord and sets the default values.
func NewDomainDnsUnavailableRecord()(*DomainDnsUnavailableRecord) {
    m := &DomainDnsUnavailableRecord{
        DomainDnsRecord: *NewDomainDnsRecord(),
    }
    return m
}
// CreateDomainDnsUnavailableRecordFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDomainDnsUnavailableRecordFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDomainDnsUnavailableRecord(), nil
}
// GetDescription gets the description property value. Provides the reason why the DomainDnsUnavailableRecord entity is returned.
func (m *DomainDnsUnavailableRecord) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DomainDnsUnavailableRecord) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DomainDnsRecord.GetFieldDeserializers()
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *DomainDnsUnavailableRecord) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DomainDnsRecord.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDescription sets the description property value. Provides the reason why the DomainDnsUnavailableRecord entity is returned.
func (m *DomainDnsUnavailableRecord) SetDescription(value *string)() {
    m.description = value
}
