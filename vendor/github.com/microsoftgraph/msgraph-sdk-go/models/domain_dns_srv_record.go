package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DomainDnsSrvRecord 
type DomainDnsSrvRecord struct {
    DomainDnsRecord
    // Value to use when configuring the Target property of the SRV record at the DNS host.
    nameTarget *string
    // Value to use when configuring the port property of the SRV record at the DNS host.
    port *int32
    // Value to use when configuring the priority property of the SRV record at the DNS host.
    priority *int32
    // Value to use when configuring the protocol property of the SRV record at the DNS host.
    protocol *string
    // Value to use when configuring the service property of the SRV record at the DNS host.
    service *string
    // Value to use when configuring the weight property of the SRV record at the DNS host.
    weight *int32
}
// NewDomainDnsSrvRecord instantiates a new DomainDnsSrvRecord and sets the default values.
func NewDomainDnsSrvRecord()(*DomainDnsSrvRecord) {
    m := &DomainDnsSrvRecord{
        DomainDnsRecord: *NewDomainDnsRecord(),
    }
    return m
}
// CreateDomainDnsSrvRecordFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDomainDnsSrvRecordFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDomainDnsSrvRecord(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DomainDnsSrvRecord) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DomainDnsRecord.GetFieldDeserializers()
    res["nameTarget"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNameTarget(val)
        }
        return nil
    }
    res["port"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPort(val)
        }
        return nil
    }
    res["priority"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPriority(val)
        }
        return nil
    }
    res["protocol"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProtocol(val)
        }
        return nil
    }
    res["service"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetService(val)
        }
        return nil
    }
    res["weight"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWeight(val)
        }
        return nil
    }
    return res
}
// GetNameTarget gets the nameTarget property value. Value to use when configuring the Target property of the SRV record at the DNS host.
func (m *DomainDnsSrvRecord) GetNameTarget()(*string) {
    return m.nameTarget
}
// GetPort gets the port property value. Value to use when configuring the port property of the SRV record at the DNS host.
func (m *DomainDnsSrvRecord) GetPort()(*int32) {
    return m.port
}
// GetPriority gets the priority property value. Value to use when configuring the priority property of the SRV record at the DNS host.
func (m *DomainDnsSrvRecord) GetPriority()(*int32) {
    return m.priority
}
// GetProtocol gets the protocol property value. Value to use when configuring the protocol property of the SRV record at the DNS host.
func (m *DomainDnsSrvRecord) GetProtocol()(*string) {
    return m.protocol
}
// GetService gets the service property value. Value to use when configuring the service property of the SRV record at the DNS host.
func (m *DomainDnsSrvRecord) GetService()(*string) {
    return m.service
}
// GetWeight gets the weight property value. Value to use when configuring the weight property of the SRV record at the DNS host.
func (m *DomainDnsSrvRecord) GetWeight()(*int32) {
    return m.weight
}
// Serialize serializes information the current object
func (m *DomainDnsSrvRecord) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DomainDnsRecord.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("nameTarget", m.GetNameTarget())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("port", m.GetPort())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("priority", m.GetPriority())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("protocol", m.GetProtocol())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("service", m.GetService())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("weight", m.GetWeight())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetNameTarget sets the nameTarget property value. Value to use when configuring the Target property of the SRV record at the DNS host.
func (m *DomainDnsSrvRecord) SetNameTarget(value *string)() {
    m.nameTarget = value
}
// SetPort sets the port property value. Value to use when configuring the port property of the SRV record at the DNS host.
func (m *DomainDnsSrvRecord) SetPort(value *int32)() {
    m.port = value
}
// SetPriority sets the priority property value. Value to use when configuring the priority property of the SRV record at the DNS host.
func (m *DomainDnsSrvRecord) SetPriority(value *int32)() {
    m.priority = value
}
// SetProtocol sets the protocol property value. Value to use when configuring the protocol property of the SRV record at the DNS host.
func (m *DomainDnsSrvRecord) SetProtocol(value *string)() {
    m.protocol = value
}
// SetService sets the service property value. Value to use when configuring the service property of the SRV record at the DNS host.
func (m *DomainDnsSrvRecord) SetService(value *string)() {
    m.service = value
}
// SetWeight sets the weight property value. Value to use when configuring the weight property of the SRV record at the DNS host.
func (m *DomainDnsSrvRecord) SetWeight(value *int32)() {
    m.weight = value
}
