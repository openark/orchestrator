package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// HostSecurityState 
type HostSecurityState struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Host FQDN (Fully Qualified Domain Name) (for example, machine.company.com).
    fqdn *string
    // The isAzureAdJoined property
    isAzureAdJoined *bool
    // The isAzureAdRegistered property
    isAzureAdRegistered *bool
    // True if the host is domain joined to an on-premises Active Directory domain.
    isHybridAzureDomainJoined *bool
    // The local host name, without the DNS domain name.
    netBiosName *string
    // The OdataType property
    odataType *string
    // Host Operating System. (For example, Windows10, MacOS, RHEL, etc.).
    os *string
    // Private (not routable) IPv4 or IPv6 address (see RFC 1918) at the time of the alert.
    privateIpAddress *string
    // Publicly routable IPv4 or IPv6 address (see RFC 1918) at time of the alert.
    publicIpAddress *string
    // Provider-generated/calculated risk score of the host.  Recommended value range of 0-1, which equates to a percentage.
    riskScore *string
}
// NewHostSecurityState instantiates a new hostSecurityState and sets the default values.
func NewHostSecurityState()(*HostSecurityState) {
    m := &HostSecurityState{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateHostSecurityStateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateHostSecurityStateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewHostSecurityState(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *HostSecurityState) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *HostSecurityState) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["fqdn"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFqdn(val)
        }
        return nil
    }
    res["isAzureAdJoined"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsAzureAdJoined(val)
        }
        return nil
    }
    res["isAzureAdRegistered"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsAzureAdRegistered(val)
        }
        return nil
    }
    res["isHybridAzureDomainJoined"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsHybridAzureDomainJoined(val)
        }
        return nil
    }
    res["netBiosName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetBiosName(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["os"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOs(val)
        }
        return nil
    }
    res["privateIpAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrivateIpAddress(val)
        }
        return nil
    }
    res["publicIpAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublicIpAddress(val)
        }
        return nil
    }
    res["riskScore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRiskScore(val)
        }
        return nil
    }
    return res
}
// GetFqdn gets the fqdn property value. Host FQDN (Fully Qualified Domain Name) (for example, machine.company.com).
func (m *HostSecurityState) GetFqdn()(*string) {
    return m.fqdn
}
// GetIsAzureAdJoined gets the isAzureAdJoined property value. The isAzureAdJoined property
func (m *HostSecurityState) GetIsAzureAdJoined()(*bool) {
    return m.isAzureAdJoined
}
// GetIsAzureAdRegistered gets the isAzureAdRegistered property value. The isAzureAdRegistered property
func (m *HostSecurityState) GetIsAzureAdRegistered()(*bool) {
    return m.isAzureAdRegistered
}
// GetIsHybridAzureDomainJoined gets the isHybridAzureDomainJoined property value. True if the host is domain joined to an on-premises Active Directory domain.
func (m *HostSecurityState) GetIsHybridAzureDomainJoined()(*bool) {
    return m.isHybridAzureDomainJoined
}
// GetNetBiosName gets the netBiosName property value. The local host name, without the DNS domain name.
func (m *HostSecurityState) GetNetBiosName()(*string) {
    return m.netBiosName
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *HostSecurityState) GetOdataType()(*string) {
    return m.odataType
}
// GetOs gets the os property value. Host Operating System. (For example, Windows10, MacOS, RHEL, etc.).
func (m *HostSecurityState) GetOs()(*string) {
    return m.os
}
// GetPrivateIpAddress gets the privateIpAddress property value. Private (not routable) IPv4 or IPv6 address (see RFC 1918) at the time of the alert.
func (m *HostSecurityState) GetPrivateIpAddress()(*string) {
    return m.privateIpAddress
}
// GetPublicIpAddress gets the publicIpAddress property value. Publicly routable IPv4 or IPv6 address (see RFC 1918) at time of the alert.
func (m *HostSecurityState) GetPublicIpAddress()(*string) {
    return m.publicIpAddress
}
// GetRiskScore gets the riskScore property value. Provider-generated/calculated risk score of the host.  Recommended value range of 0-1, which equates to a percentage.
func (m *HostSecurityState) GetRiskScore()(*string) {
    return m.riskScore
}
// Serialize serializes information the current object
func (m *HostSecurityState) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("fqdn", m.GetFqdn())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isAzureAdJoined", m.GetIsAzureAdJoined())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isAzureAdRegistered", m.GetIsAzureAdRegistered())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isHybridAzureDomainJoined", m.GetIsHybridAzureDomainJoined())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("netBiosName", m.GetNetBiosName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("os", m.GetOs())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("privateIpAddress", m.GetPrivateIpAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("publicIpAddress", m.GetPublicIpAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("riskScore", m.GetRiskScore())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *HostSecurityState) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetFqdn sets the fqdn property value. Host FQDN (Fully Qualified Domain Name) (for example, machine.company.com).
func (m *HostSecurityState) SetFqdn(value *string)() {
    m.fqdn = value
}
// SetIsAzureAdJoined sets the isAzureAdJoined property value. The isAzureAdJoined property
func (m *HostSecurityState) SetIsAzureAdJoined(value *bool)() {
    m.isAzureAdJoined = value
}
// SetIsAzureAdRegistered sets the isAzureAdRegistered property value. The isAzureAdRegistered property
func (m *HostSecurityState) SetIsAzureAdRegistered(value *bool)() {
    m.isAzureAdRegistered = value
}
// SetIsHybridAzureDomainJoined sets the isHybridAzureDomainJoined property value. True if the host is domain joined to an on-premises Active Directory domain.
func (m *HostSecurityState) SetIsHybridAzureDomainJoined(value *bool)() {
    m.isHybridAzureDomainJoined = value
}
// SetNetBiosName sets the netBiosName property value. The local host name, without the DNS domain name.
func (m *HostSecurityState) SetNetBiosName(value *string)() {
    m.netBiosName = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *HostSecurityState) SetOdataType(value *string)() {
    m.odataType = value
}
// SetOs sets the os property value. Host Operating System. (For example, Windows10, MacOS, RHEL, etc.).
func (m *HostSecurityState) SetOs(value *string)() {
    m.os = value
}
// SetPrivateIpAddress sets the privateIpAddress property value. Private (not routable) IPv4 or IPv6 address (see RFC 1918) at the time of the alert.
func (m *HostSecurityState) SetPrivateIpAddress(value *string)() {
    m.privateIpAddress = value
}
// SetPublicIpAddress sets the publicIpAddress property value. Publicly routable IPv4 or IPv6 address (see RFC 1918) at time of the alert.
func (m *HostSecurityState) SetPublicIpAddress(value *string)() {
    m.publicIpAddress = value
}
// SetRiskScore sets the riskScore property value. Provider-generated/calculated risk score of the host.  Recommended value range of 0-1, which equates to a percentage.
func (m *HostSecurityState) SetRiskScore(value *string)() {
    m.riskScore = value
}
