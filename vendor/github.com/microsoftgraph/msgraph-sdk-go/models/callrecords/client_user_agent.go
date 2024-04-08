package callrecords

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ClientUserAgent 
type ClientUserAgent struct {
    UserAgent
    // The unique identifier of the Azure AD application used by this endpoint.
    azureADAppId *string
    // Immutable resource identifier of the Azure Communication Service associated with this endpoint based on Communication Services APIs.
    communicationServiceId *string
    // The platform property
    platform *ClientPlatform
    // The productFamily property
    productFamily *ProductFamily
}
// NewClientUserAgent instantiates a new ClientUserAgent and sets the default values.
func NewClientUserAgent()(*ClientUserAgent) {
    m := &ClientUserAgent{
        UserAgent: *NewUserAgent(),
    }
    odataTypeValue := "#microsoft.graph.callRecords.clientUserAgent"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateClientUserAgentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateClientUserAgentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewClientUserAgent(), nil
}
// GetAzureADAppId gets the azureADAppId property value. The unique identifier of the Azure AD application used by this endpoint.
func (m *ClientUserAgent) GetAzureADAppId()(*string) {
    return m.azureADAppId
}
// GetCommunicationServiceId gets the communicationServiceId property value. Immutable resource identifier of the Azure Communication Service associated with this endpoint based on Communication Services APIs.
func (m *ClientUserAgent) GetCommunicationServiceId()(*string) {
    return m.communicationServiceId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ClientUserAgent) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.UserAgent.GetFieldDeserializers()
    res["azureADAppId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAzureADAppId(val)
        }
        return nil
    }
    res["communicationServiceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCommunicationServiceId(val)
        }
        return nil
    }
    res["platform"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseClientPlatform)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPlatform(val.(*ClientPlatform))
        }
        return nil
    }
    res["productFamily"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseProductFamily)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProductFamily(val.(*ProductFamily))
        }
        return nil
    }
    return res
}
// GetPlatform gets the platform property value. The platform property
func (m *ClientUserAgent) GetPlatform()(*ClientPlatform) {
    return m.platform
}
// GetProductFamily gets the productFamily property value. The productFamily property
func (m *ClientUserAgent) GetProductFamily()(*ProductFamily) {
    return m.productFamily
}
// Serialize serializes information the current object
func (m *ClientUserAgent) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.UserAgent.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("azureADAppId", m.GetAzureADAppId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("communicationServiceId", m.GetCommunicationServiceId())
        if err != nil {
            return err
        }
    }
    if m.GetPlatform() != nil {
        cast := (*m.GetPlatform()).String()
        err = writer.WriteStringValue("platform", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetProductFamily() != nil {
        cast := (*m.GetProductFamily()).String()
        err = writer.WriteStringValue("productFamily", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAzureADAppId sets the azureADAppId property value. The unique identifier of the Azure AD application used by this endpoint.
func (m *ClientUserAgent) SetAzureADAppId(value *string)() {
    m.azureADAppId = value
}
// SetCommunicationServiceId sets the communicationServiceId property value. Immutable resource identifier of the Azure Communication Service associated with this endpoint based on Communication Services APIs.
func (m *ClientUserAgent) SetCommunicationServiceId(value *string)() {
    m.communicationServiceId = value
}
// SetPlatform sets the platform property value. The platform property
func (m *ClientUserAgent) SetPlatform(value *ClientPlatform)() {
    m.platform = value
}
// SetProductFamily sets the productFamily property value. The productFamily property
func (m *ClientUserAgent) SetProductFamily(value *ProductFamily)() {
    m.productFamily = value
}
