package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementExchangeConnector entity which represents a connection to an Exchange environment.
type DeviceManagementExchangeConnector struct {
    Entity
    // The name of the server hosting the Exchange Connector.
    connectorServerName *string
    // An alias assigned to the Exchange server
    exchangeAlias *string
    // The type of Exchange Connector.
    exchangeConnectorType *DeviceManagementExchangeConnectorType
    // Exchange Organization to the Exchange server
    exchangeOrganization *string
    // Last sync time for the Exchange Connector
    lastSyncDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Email address used to configure the Service To Service Exchange Connector.
    primarySmtpAddress *string
    // The name of the Exchange server.
    serverName *string
    // The current status of the Exchange Connector.
    status *DeviceManagementExchangeConnectorStatus
    // The version of the ExchangeConnectorAgent
    version *string
}
// NewDeviceManagementExchangeConnector instantiates a new deviceManagementExchangeConnector and sets the default values.
func NewDeviceManagementExchangeConnector()(*DeviceManagementExchangeConnector) {
    m := &DeviceManagementExchangeConnector{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeviceManagementExchangeConnectorFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementExchangeConnectorFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementExchangeConnector(), nil
}
// GetConnectorServerName gets the connectorServerName property value. The name of the server hosting the Exchange Connector.
func (m *DeviceManagementExchangeConnector) GetConnectorServerName()(*string) {
    return m.connectorServerName
}
// GetExchangeAlias gets the exchangeAlias property value. An alias assigned to the Exchange server
func (m *DeviceManagementExchangeConnector) GetExchangeAlias()(*string) {
    return m.exchangeAlias
}
// GetExchangeConnectorType gets the exchangeConnectorType property value. The type of Exchange Connector.
func (m *DeviceManagementExchangeConnector) GetExchangeConnectorType()(*DeviceManagementExchangeConnectorType) {
    return m.exchangeConnectorType
}
// GetExchangeOrganization gets the exchangeOrganization property value. Exchange Organization to the Exchange server
func (m *DeviceManagementExchangeConnector) GetExchangeOrganization()(*string) {
    return m.exchangeOrganization
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementExchangeConnector) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["connectorServerName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConnectorServerName(val)
        }
        return nil
    }
    res["exchangeAlias"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExchangeAlias(val)
        }
        return nil
    }
    res["exchangeConnectorType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementExchangeConnectorType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExchangeConnectorType(val.(*DeviceManagementExchangeConnectorType))
        }
        return nil
    }
    res["exchangeOrganization"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExchangeOrganization(val)
        }
        return nil
    }
    res["lastSyncDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastSyncDateTime(val)
        }
        return nil
    }
    res["primarySmtpAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrimarySmtpAddress(val)
        }
        return nil
    }
    res["serverName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServerName(val)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementExchangeConnectorStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*DeviceManagementExchangeConnectorStatus))
        }
        return nil
    }
    res["version"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersion(val)
        }
        return nil
    }
    return res
}
// GetLastSyncDateTime gets the lastSyncDateTime property value. Last sync time for the Exchange Connector
func (m *DeviceManagementExchangeConnector) GetLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastSyncDateTime
}
// GetPrimarySmtpAddress gets the primarySmtpAddress property value. Email address used to configure the Service To Service Exchange Connector.
func (m *DeviceManagementExchangeConnector) GetPrimarySmtpAddress()(*string) {
    return m.primarySmtpAddress
}
// GetServerName gets the serverName property value. The name of the Exchange server.
func (m *DeviceManagementExchangeConnector) GetServerName()(*string) {
    return m.serverName
}
// GetStatus gets the status property value. The current status of the Exchange Connector.
func (m *DeviceManagementExchangeConnector) GetStatus()(*DeviceManagementExchangeConnectorStatus) {
    return m.status
}
// GetVersion gets the version property value. The version of the ExchangeConnectorAgent
func (m *DeviceManagementExchangeConnector) GetVersion()(*string) {
    return m.version
}
// Serialize serializes information the current object
func (m *DeviceManagementExchangeConnector) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("connectorServerName", m.GetConnectorServerName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("exchangeAlias", m.GetExchangeAlias())
        if err != nil {
            return err
        }
    }
    if m.GetExchangeConnectorType() != nil {
        cast := (*m.GetExchangeConnectorType()).String()
        err = writer.WriteStringValue("exchangeConnectorType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("exchangeOrganization", m.GetExchangeOrganization())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastSyncDateTime", m.GetLastSyncDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("primarySmtpAddress", m.GetPrimarySmtpAddress())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("serverName", m.GetServerName())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("version", m.GetVersion())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetConnectorServerName sets the connectorServerName property value. The name of the server hosting the Exchange Connector.
func (m *DeviceManagementExchangeConnector) SetConnectorServerName(value *string)() {
    m.connectorServerName = value
}
// SetExchangeAlias sets the exchangeAlias property value. An alias assigned to the Exchange server
func (m *DeviceManagementExchangeConnector) SetExchangeAlias(value *string)() {
    m.exchangeAlias = value
}
// SetExchangeConnectorType sets the exchangeConnectorType property value. The type of Exchange Connector.
func (m *DeviceManagementExchangeConnector) SetExchangeConnectorType(value *DeviceManagementExchangeConnectorType)() {
    m.exchangeConnectorType = value
}
// SetExchangeOrganization sets the exchangeOrganization property value. Exchange Organization to the Exchange server
func (m *DeviceManagementExchangeConnector) SetExchangeOrganization(value *string)() {
    m.exchangeOrganization = value
}
// SetLastSyncDateTime sets the lastSyncDateTime property value. Last sync time for the Exchange Connector
func (m *DeviceManagementExchangeConnector) SetLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastSyncDateTime = value
}
// SetPrimarySmtpAddress sets the primarySmtpAddress property value. Email address used to configure the Service To Service Exchange Connector.
func (m *DeviceManagementExchangeConnector) SetPrimarySmtpAddress(value *string)() {
    m.primarySmtpAddress = value
}
// SetServerName sets the serverName property value. The name of the Exchange server.
func (m *DeviceManagementExchangeConnector) SetServerName(value *string)() {
    m.serverName = value
}
// SetStatus sets the status property value. The current status of the Exchange Connector.
func (m *DeviceManagementExchangeConnector) SetStatus(value *DeviceManagementExchangeConnectorStatus)() {
    m.status = value
}
// SetVersion sets the version property value. The version of the ExchangeConnectorAgent
func (m *DeviceManagementExchangeConnector) SetVersion(value *string)() {
    m.version = value
}
