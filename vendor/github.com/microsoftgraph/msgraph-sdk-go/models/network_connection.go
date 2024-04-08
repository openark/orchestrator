package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// NetworkConnection 
type NetworkConnection struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Name of the application managing the network connection (for example, Facebook or SMTP).
    applicationName *string
    // Destination IP address (of the network connection).
    destinationAddress *string
    // Destination domain portion of the destination URL. (for example 'www.contoso.com').
    destinationDomain *string
    // Location (by IP address mapping) associated with the destination of a network connection.
    destinationLocation *string
    // Destination port (of the network connection).
    destinationPort *string
    // Network connection URL/URI string - excluding parameters. (for example 'www.contoso.com/products/default.html')
    destinationUrl *string
    // Network connection direction. Possible values are: unknown, inbound, outbound.
    direction *ConnectionDirection
    // Date when the destination domain was registered. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    domainRegisteredDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The local DNS name resolution as it appears in the host's local DNS cache (for example, in case the 'hosts' file was tampered with).
    localDnsName *string
    // Network Address Translation destination IP address.
    natDestinationAddress *string
    // Network Address Translation destination port.
    natDestinationPort *string
    // Network Address Translation source IP address.
    natSourceAddress *string
    // Network Address Translation source port.
    natSourcePort *string
    // The OdataType property
    odataType *string
    // Network protocol. Possible values are: unknown, ip, icmp, igmp, ggp, ipv4, tcp, pup, udp, idp, ipv6, ipv6RoutingHeader, ipv6FragmentHeader, ipSecEncapsulatingSecurityPayload, ipSecAuthenticationHeader, icmpV6, ipv6NoNextHeader, ipv6DestinationOptions, nd, raw, ipx, spx, spxII.
    protocol *SecurityNetworkProtocol
    // Provider generated/calculated risk score of the network connection. Recommended value range of 0-1, which equates to a percentage.
    riskScore *string
    // Source (i.e. origin) IP address (of the network connection).
    sourceAddress *string
    // Location (by IP address mapping) associated with the source of a network connection.
    sourceLocation *string
    // Source (i.e. origin) IP port (of the network connection).
    sourcePort *string
    // Network connection status. Possible values are: unknown, attempted, succeeded, blocked, failed.
    status *ConnectionStatus
    // Parameters (suffix) of the destination URL.
    urlParameters *string
}
// NewNetworkConnection instantiates a new networkConnection and sets the default values.
func NewNetworkConnection()(*NetworkConnection) {
    m := &NetworkConnection{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateNetworkConnectionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateNetworkConnectionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewNetworkConnection(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *NetworkConnection) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetApplicationName gets the applicationName property value. Name of the application managing the network connection (for example, Facebook or SMTP).
func (m *NetworkConnection) GetApplicationName()(*string) {
    return m.applicationName
}
// GetDestinationAddress gets the destinationAddress property value. Destination IP address (of the network connection).
func (m *NetworkConnection) GetDestinationAddress()(*string) {
    return m.destinationAddress
}
// GetDestinationDomain gets the destinationDomain property value. Destination domain portion of the destination URL. (for example 'www.contoso.com').
func (m *NetworkConnection) GetDestinationDomain()(*string) {
    return m.destinationDomain
}
// GetDestinationLocation gets the destinationLocation property value. Location (by IP address mapping) associated with the destination of a network connection.
func (m *NetworkConnection) GetDestinationLocation()(*string) {
    return m.destinationLocation
}
// GetDestinationPort gets the destinationPort property value. Destination port (of the network connection).
func (m *NetworkConnection) GetDestinationPort()(*string) {
    return m.destinationPort
}
// GetDestinationUrl gets the destinationUrl property value. Network connection URL/URI string - excluding parameters. (for example 'www.contoso.com/products/default.html')
func (m *NetworkConnection) GetDestinationUrl()(*string) {
    return m.destinationUrl
}
// GetDirection gets the direction property value. Network connection direction. Possible values are: unknown, inbound, outbound.
func (m *NetworkConnection) GetDirection()(*ConnectionDirection) {
    return m.direction
}
// GetDomainRegisteredDateTime gets the domainRegisteredDateTime property value. Date when the destination domain was registered. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *NetworkConnection) GetDomainRegisteredDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.domainRegisteredDateTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *NetworkConnection) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["applicationName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationName(val)
        }
        return nil
    }
    res["destinationAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDestinationAddress(val)
        }
        return nil
    }
    res["destinationDomain"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDestinationDomain(val)
        }
        return nil
    }
    res["destinationLocation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDestinationLocation(val)
        }
        return nil
    }
    res["destinationPort"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDestinationPort(val)
        }
        return nil
    }
    res["destinationUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDestinationUrl(val)
        }
        return nil
    }
    res["direction"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseConnectionDirection)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDirection(val.(*ConnectionDirection))
        }
        return nil
    }
    res["domainRegisteredDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDomainRegisteredDateTime(val)
        }
        return nil
    }
    res["localDnsName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalDnsName(val)
        }
        return nil
    }
    res["natDestinationAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNatDestinationAddress(val)
        }
        return nil
    }
    res["natDestinationPort"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNatDestinationPort(val)
        }
        return nil
    }
    res["natSourceAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNatSourceAddress(val)
        }
        return nil
    }
    res["natSourcePort"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNatSourcePort(val)
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
    res["protocol"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSecurityNetworkProtocol)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProtocol(val.(*SecurityNetworkProtocol))
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
    res["sourceAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSourceAddress(val)
        }
        return nil
    }
    res["sourceLocation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSourceLocation(val)
        }
        return nil
    }
    res["sourcePort"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSourcePort(val)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseConnectionStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*ConnectionStatus))
        }
        return nil
    }
    res["urlParameters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUrlParameters(val)
        }
        return nil
    }
    return res
}
// GetLocalDnsName gets the localDnsName property value. The local DNS name resolution as it appears in the host's local DNS cache (for example, in case the 'hosts' file was tampered with).
func (m *NetworkConnection) GetLocalDnsName()(*string) {
    return m.localDnsName
}
// GetNatDestinationAddress gets the natDestinationAddress property value. Network Address Translation destination IP address.
func (m *NetworkConnection) GetNatDestinationAddress()(*string) {
    return m.natDestinationAddress
}
// GetNatDestinationPort gets the natDestinationPort property value. Network Address Translation destination port.
func (m *NetworkConnection) GetNatDestinationPort()(*string) {
    return m.natDestinationPort
}
// GetNatSourceAddress gets the natSourceAddress property value. Network Address Translation source IP address.
func (m *NetworkConnection) GetNatSourceAddress()(*string) {
    return m.natSourceAddress
}
// GetNatSourcePort gets the natSourcePort property value. Network Address Translation source port.
func (m *NetworkConnection) GetNatSourcePort()(*string) {
    return m.natSourcePort
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *NetworkConnection) GetOdataType()(*string) {
    return m.odataType
}
// GetProtocol gets the protocol property value. Network protocol. Possible values are: unknown, ip, icmp, igmp, ggp, ipv4, tcp, pup, udp, idp, ipv6, ipv6RoutingHeader, ipv6FragmentHeader, ipSecEncapsulatingSecurityPayload, ipSecAuthenticationHeader, icmpV6, ipv6NoNextHeader, ipv6DestinationOptions, nd, raw, ipx, spx, spxII.
func (m *NetworkConnection) GetProtocol()(*SecurityNetworkProtocol) {
    return m.protocol
}
// GetRiskScore gets the riskScore property value. Provider generated/calculated risk score of the network connection. Recommended value range of 0-1, which equates to a percentage.
func (m *NetworkConnection) GetRiskScore()(*string) {
    return m.riskScore
}
// GetSourceAddress gets the sourceAddress property value. Source (i.e. origin) IP address (of the network connection).
func (m *NetworkConnection) GetSourceAddress()(*string) {
    return m.sourceAddress
}
// GetSourceLocation gets the sourceLocation property value. Location (by IP address mapping) associated with the source of a network connection.
func (m *NetworkConnection) GetSourceLocation()(*string) {
    return m.sourceLocation
}
// GetSourcePort gets the sourcePort property value. Source (i.e. origin) IP port (of the network connection).
func (m *NetworkConnection) GetSourcePort()(*string) {
    return m.sourcePort
}
// GetStatus gets the status property value. Network connection status. Possible values are: unknown, attempted, succeeded, blocked, failed.
func (m *NetworkConnection) GetStatus()(*ConnectionStatus) {
    return m.status
}
// GetUrlParameters gets the urlParameters property value. Parameters (suffix) of the destination URL.
func (m *NetworkConnection) GetUrlParameters()(*string) {
    return m.urlParameters
}
// Serialize serializes information the current object
func (m *NetworkConnection) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("applicationName", m.GetApplicationName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("destinationAddress", m.GetDestinationAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("destinationDomain", m.GetDestinationDomain())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("destinationLocation", m.GetDestinationLocation())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("destinationPort", m.GetDestinationPort())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("destinationUrl", m.GetDestinationUrl())
        if err != nil {
            return err
        }
    }
    if m.GetDirection() != nil {
        cast := (*m.GetDirection()).String()
        err := writer.WriteStringValue("direction", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("domainRegisteredDateTime", m.GetDomainRegisteredDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("localDnsName", m.GetLocalDnsName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("natDestinationAddress", m.GetNatDestinationAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("natDestinationPort", m.GetNatDestinationPort())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("natSourceAddress", m.GetNatSourceAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("natSourcePort", m.GetNatSourcePort())
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
    if m.GetProtocol() != nil {
        cast := (*m.GetProtocol()).String()
        err := writer.WriteStringValue("protocol", &cast)
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
        err := writer.WriteStringValue("sourceAddress", m.GetSourceAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("sourceLocation", m.GetSourceLocation())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("sourcePort", m.GetSourcePort())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err := writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("urlParameters", m.GetUrlParameters())
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
func (m *NetworkConnection) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetApplicationName sets the applicationName property value. Name of the application managing the network connection (for example, Facebook or SMTP).
func (m *NetworkConnection) SetApplicationName(value *string)() {
    m.applicationName = value
}
// SetDestinationAddress sets the destinationAddress property value. Destination IP address (of the network connection).
func (m *NetworkConnection) SetDestinationAddress(value *string)() {
    m.destinationAddress = value
}
// SetDestinationDomain sets the destinationDomain property value. Destination domain portion of the destination URL. (for example 'www.contoso.com').
func (m *NetworkConnection) SetDestinationDomain(value *string)() {
    m.destinationDomain = value
}
// SetDestinationLocation sets the destinationLocation property value. Location (by IP address mapping) associated with the destination of a network connection.
func (m *NetworkConnection) SetDestinationLocation(value *string)() {
    m.destinationLocation = value
}
// SetDestinationPort sets the destinationPort property value. Destination port (of the network connection).
func (m *NetworkConnection) SetDestinationPort(value *string)() {
    m.destinationPort = value
}
// SetDestinationUrl sets the destinationUrl property value. Network connection URL/URI string - excluding parameters. (for example 'www.contoso.com/products/default.html')
func (m *NetworkConnection) SetDestinationUrl(value *string)() {
    m.destinationUrl = value
}
// SetDirection sets the direction property value. Network connection direction. Possible values are: unknown, inbound, outbound.
func (m *NetworkConnection) SetDirection(value *ConnectionDirection)() {
    m.direction = value
}
// SetDomainRegisteredDateTime sets the domainRegisteredDateTime property value. Date when the destination domain was registered. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *NetworkConnection) SetDomainRegisteredDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.domainRegisteredDateTime = value
}
// SetLocalDnsName sets the localDnsName property value. The local DNS name resolution as it appears in the host's local DNS cache (for example, in case the 'hosts' file was tampered with).
func (m *NetworkConnection) SetLocalDnsName(value *string)() {
    m.localDnsName = value
}
// SetNatDestinationAddress sets the natDestinationAddress property value. Network Address Translation destination IP address.
func (m *NetworkConnection) SetNatDestinationAddress(value *string)() {
    m.natDestinationAddress = value
}
// SetNatDestinationPort sets the natDestinationPort property value. Network Address Translation destination port.
func (m *NetworkConnection) SetNatDestinationPort(value *string)() {
    m.natDestinationPort = value
}
// SetNatSourceAddress sets the natSourceAddress property value. Network Address Translation source IP address.
func (m *NetworkConnection) SetNatSourceAddress(value *string)() {
    m.natSourceAddress = value
}
// SetNatSourcePort sets the natSourcePort property value. Network Address Translation source port.
func (m *NetworkConnection) SetNatSourcePort(value *string)() {
    m.natSourcePort = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *NetworkConnection) SetOdataType(value *string)() {
    m.odataType = value
}
// SetProtocol sets the protocol property value. Network protocol. Possible values are: unknown, ip, icmp, igmp, ggp, ipv4, tcp, pup, udp, idp, ipv6, ipv6RoutingHeader, ipv6FragmentHeader, ipSecEncapsulatingSecurityPayload, ipSecAuthenticationHeader, icmpV6, ipv6NoNextHeader, ipv6DestinationOptions, nd, raw, ipx, spx, spxII.
func (m *NetworkConnection) SetProtocol(value *SecurityNetworkProtocol)() {
    m.protocol = value
}
// SetRiskScore sets the riskScore property value. Provider generated/calculated risk score of the network connection. Recommended value range of 0-1, which equates to a percentage.
func (m *NetworkConnection) SetRiskScore(value *string)() {
    m.riskScore = value
}
// SetSourceAddress sets the sourceAddress property value. Source (i.e. origin) IP address (of the network connection).
func (m *NetworkConnection) SetSourceAddress(value *string)() {
    m.sourceAddress = value
}
// SetSourceLocation sets the sourceLocation property value. Location (by IP address mapping) associated with the source of a network connection.
func (m *NetworkConnection) SetSourceLocation(value *string)() {
    m.sourceLocation = value
}
// SetSourcePort sets the sourcePort property value. Source (i.e. origin) IP port (of the network connection).
func (m *NetworkConnection) SetSourcePort(value *string)() {
    m.sourcePort = value
}
// SetStatus sets the status property value. Network connection status. Possible values are: unknown, attempted, succeeded, blocked, failed.
func (m *NetworkConnection) SetStatus(value *ConnectionStatus)() {
    m.status = value
}
// SetUrlParameters sets the urlParameters property value. Parameters (suffix) of the destination URL.
func (m *NetworkConnection) SetUrlParameters(value *string)() {
    m.urlParameters = value
}
