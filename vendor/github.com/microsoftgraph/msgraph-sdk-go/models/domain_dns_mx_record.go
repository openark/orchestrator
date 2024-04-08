package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DomainDnsMxRecord 
type DomainDnsMxRecord struct {
    DomainDnsRecord
    // Value used when configuring the answer/destination/value of the MX record at the DNS host.
    mailExchange *string
    // Value used when configuring the Preference/Priority property of the MX record at the DNS host.
    preference *int32
}
// NewDomainDnsMxRecord instantiates a new DomainDnsMxRecord and sets the default values.
func NewDomainDnsMxRecord()(*DomainDnsMxRecord) {
    m := &DomainDnsMxRecord{
        DomainDnsRecord: *NewDomainDnsRecord(),
    }
    return m
}
// CreateDomainDnsMxRecordFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDomainDnsMxRecordFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDomainDnsMxRecord(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DomainDnsMxRecord) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DomainDnsRecord.GetFieldDeserializers()
    res["mailExchange"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMailExchange(val)
        }
        return nil
    }
    res["preference"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreference(val)
        }
        return nil
    }
    return res
}
// GetMailExchange gets the mailExchange property value. Value used when configuring the answer/destination/value of the MX record at the DNS host.
func (m *DomainDnsMxRecord) GetMailExchange()(*string) {
    return m.mailExchange
}
// GetPreference gets the preference property value. Value used when configuring the Preference/Priority property of the MX record at the DNS host.
func (m *DomainDnsMxRecord) GetPreference()(*int32) {
    return m.preference
}
// Serialize serializes information the current object
func (m *DomainDnsMxRecord) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DomainDnsRecord.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("mailExchange", m.GetMailExchange())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("preference", m.GetPreference())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetMailExchange sets the mailExchange property value. Value used when configuring the answer/destination/value of the MX record at the DNS host.
func (m *DomainDnsMxRecord) SetMailExchange(value *string)() {
    m.mailExchange = value
}
// SetPreference sets the preference property value. Value used when configuring the Preference/Priority property of the MX record at the DNS host.
func (m *DomainDnsMxRecord) SetPreference(value *int32)() {
    m.preference = value
}
