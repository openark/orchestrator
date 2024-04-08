package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ParticipantInfo 
type ParticipantInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The ISO 3166-1 Alpha-2 country code of the participant's best estimated physical location at the start of the call. Read-only.
    countryCode *string
    // The type of endpoint the participant is using. Possible values are: default, skypeForBusiness, or skypeForBusinessVoipPhone. Read-only.
    endpointType *EndpointType
    // The identity property
    identity IdentitySetable
    // The language culture string. Read-only.
    languageId *string
    // The OdataType property
    odataType *string
    // The participant ID of the participant. Read-only.
    participantId *string
    // The home region of the participant. This can be a country, a continent, or a larger geographic region. This does not change based on the participant's current physical location. Read-only.
    region *string
}
// NewParticipantInfo instantiates a new participantInfo and sets the default values.
func NewParticipantInfo()(*ParticipantInfo) {
    m := &ParticipantInfo{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateParticipantInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateParticipantInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewParticipantInfo(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ParticipantInfo) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCountryCode gets the countryCode property value. The ISO 3166-1 Alpha-2 country code of the participant's best estimated physical location at the start of the call. Read-only.
func (m *ParticipantInfo) GetCountryCode()(*string) {
    return m.countryCode
}
// GetEndpointType gets the endpointType property value. The type of endpoint the participant is using. Possible values are: default, skypeForBusiness, or skypeForBusinessVoipPhone. Read-only.
func (m *ParticipantInfo) GetEndpointType()(*EndpointType) {
    return m.endpointType
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ParticipantInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["countryCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCountryCode(val)
        }
        return nil
    }
    res["endpointType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEndpointType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndpointType(val.(*EndpointType))
        }
        return nil
    }
    res["identity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentity(val.(IdentitySetable))
        }
        return nil
    }
    res["languageId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLanguageId(val)
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
    res["participantId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParticipantId(val)
        }
        return nil
    }
    res["region"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRegion(val)
        }
        return nil
    }
    return res
}
// GetIdentity gets the identity property value. The identity property
func (m *ParticipantInfo) GetIdentity()(IdentitySetable) {
    return m.identity
}
// GetLanguageId gets the languageId property value. The language culture string. Read-only.
func (m *ParticipantInfo) GetLanguageId()(*string) {
    return m.languageId
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ParticipantInfo) GetOdataType()(*string) {
    return m.odataType
}
// GetParticipantId gets the participantId property value. The participant ID of the participant. Read-only.
func (m *ParticipantInfo) GetParticipantId()(*string) {
    return m.participantId
}
// GetRegion gets the region property value. The home region of the participant. This can be a country, a continent, or a larger geographic region. This does not change based on the participant's current physical location. Read-only.
func (m *ParticipantInfo) GetRegion()(*string) {
    return m.region
}
// Serialize serializes information the current object
func (m *ParticipantInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("countryCode", m.GetCountryCode())
        if err != nil {
            return err
        }
    }
    if m.GetEndpointType() != nil {
        cast := (*m.GetEndpointType()).String()
        err := writer.WriteStringValue("endpointType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("identity", m.GetIdentity())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("languageId", m.GetLanguageId())
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
        err := writer.WriteStringValue("participantId", m.GetParticipantId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("region", m.GetRegion())
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
func (m *ParticipantInfo) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCountryCode sets the countryCode property value. The ISO 3166-1 Alpha-2 country code of the participant's best estimated physical location at the start of the call. Read-only.
func (m *ParticipantInfo) SetCountryCode(value *string)() {
    m.countryCode = value
}
// SetEndpointType sets the endpointType property value. The type of endpoint the participant is using. Possible values are: default, skypeForBusiness, or skypeForBusinessVoipPhone. Read-only.
func (m *ParticipantInfo) SetEndpointType(value *EndpointType)() {
    m.endpointType = value
}
// SetIdentity sets the identity property value. The identity property
func (m *ParticipantInfo) SetIdentity(value IdentitySetable)() {
    m.identity = value
}
// SetLanguageId sets the languageId property value. The language culture string. Read-only.
func (m *ParticipantInfo) SetLanguageId(value *string)() {
    m.languageId = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ParticipantInfo) SetOdataType(value *string)() {
    m.odataType = value
}
// SetParticipantId sets the participantId property value. The participant ID of the participant. Read-only.
func (m *ParticipantInfo) SetParticipantId(value *string)() {
    m.participantId = value
}
// SetRegion sets the region property value. The home region of the participant. This can be a country, a continent, or a larger geographic region. This does not change based on the participant's current physical location. Read-only.
func (m *ParticipantInfo) SetRegion(value *string)() {
    m.region = value
}
