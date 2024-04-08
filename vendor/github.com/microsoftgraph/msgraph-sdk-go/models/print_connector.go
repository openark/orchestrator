package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PrintConnector 
type PrintConnector struct {
    Entity
    // The connector's version.
    appVersion *string
    // The name of the connector.
    displayName *string
    // The connector machine's hostname.
    fullyQualifiedDomainName *string
    // The physical and/or organizational location of the connector.
    location PrinterLocationable
    // The connector machine's operating system version.
    operatingSystem *string
    // The DateTimeOffset when the connector was registered.
    registeredDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// NewPrintConnector instantiates a new printConnector and sets the default values.
func NewPrintConnector()(*PrintConnector) {
    m := &PrintConnector{
        Entity: *NewEntity(),
    }
    return m
}
// CreatePrintConnectorFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePrintConnectorFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPrintConnector(), nil
}
// GetAppVersion gets the appVersion property value. The connector's version.
func (m *PrintConnector) GetAppVersion()(*string) {
    return m.appVersion
}
// GetDisplayName gets the displayName property value. The name of the connector.
func (m *PrintConnector) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PrintConnector) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["appVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppVersion(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["fullyQualifiedDomainName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFullyQualifiedDomainName(val)
        }
        return nil
    }
    res["location"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePrinterLocationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocation(val.(PrinterLocationable))
        }
        return nil
    }
    res["operatingSystem"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOperatingSystem(val)
        }
        return nil
    }
    res["registeredDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRegisteredDateTime(val)
        }
        return nil
    }
    return res
}
// GetFullyQualifiedDomainName gets the fullyQualifiedDomainName property value. The connector machine's hostname.
func (m *PrintConnector) GetFullyQualifiedDomainName()(*string) {
    return m.fullyQualifiedDomainName
}
// GetLocation gets the location property value. The physical and/or organizational location of the connector.
func (m *PrintConnector) GetLocation()(PrinterLocationable) {
    return m.location
}
// GetOperatingSystem gets the operatingSystem property value. The connector machine's operating system version.
func (m *PrintConnector) GetOperatingSystem()(*string) {
    return m.operatingSystem
}
// GetRegisteredDateTime gets the registeredDateTime property value. The DateTimeOffset when the connector was registered.
func (m *PrintConnector) GetRegisteredDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.registeredDateTime
}
// Serialize serializes information the current object
func (m *PrintConnector) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("appVersion", m.GetAppVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("fullyQualifiedDomainName", m.GetFullyQualifiedDomainName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("location", m.GetLocation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("operatingSystem", m.GetOperatingSystem())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("registeredDateTime", m.GetRegisteredDateTime())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAppVersion sets the appVersion property value. The connector's version.
func (m *PrintConnector) SetAppVersion(value *string)() {
    m.appVersion = value
}
// SetDisplayName sets the displayName property value. The name of the connector.
func (m *PrintConnector) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetFullyQualifiedDomainName sets the fullyQualifiedDomainName property value. The connector machine's hostname.
func (m *PrintConnector) SetFullyQualifiedDomainName(value *string)() {
    m.fullyQualifiedDomainName = value
}
// SetLocation sets the location property value. The physical and/or organizational location of the connector.
func (m *PrintConnector) SetLocation(value PrinterLocationable)() {
    m.location = value
}
// SetOperatingSystem sets the operatingSystem property value. The connector machine's operating system version.
func (m *PrintConnector) SetOperatingSystem(value *string)() {
    m.operatingSystem = value
}
// SetRegisteredDateTime sets the registeredDateTime property value. The DateTimeOffset when the connector was registered.
func (m *PrintConnector) SetRegisteredDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.registeredDateTime = value
}
