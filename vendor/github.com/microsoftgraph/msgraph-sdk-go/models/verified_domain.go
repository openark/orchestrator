package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// VerifiedDomain 
type VerifiedDomain struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // For example, Email, OfficeCommunicationsOnline.
    capabilities *string
    // true if this is the default domain associated with the tenant; otherwise, false.
    isDefault *bool
    // true if this is the initial domain associated with the tenant; otherwise, false.
    isInitial *bool
    // The domain name; for example, contoso.onmicrosoft.com.
    name *string
    // The OdataType property
    odataType *string
    // For example, Managed.
    typeEscaped *string
}
// NewVerifiedDomain instantiates a new verifiedDomain and sets the default values.
func NewVerifiedDomain()(*VerifiedDomain) {
    m := &VerifiedDomain{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateVerifiedDomainFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateVerifiedDomainFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVerifiedDomain(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *VerifiedDomain) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCapabilities gets the capabilities property value. For example, Email, OfficeCommunicationsOnline.
func (m *VerifiedDomain) GetCapabilities()(*string) {
    return m.capabilities
}
// GetFieldDeserializers the deserialization information for the current model
func (m *VerifiedDomain) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["capabilities"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCapabilities(val)
        }
        return nil
    }
    res["isDefault"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDefault(val)
        }
        return nil
    }
    res["isInitial"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsInitial(val)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
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
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val)
        }
        return nil
    }
    return res
}
// GetIsDefault gets the isDefault property value. true if this is the default domain associated with the tenant; otherwise, false.
func (m *VerifiedDomain) GetIsDefault()(*bool) {
    return m.isDefault
}
// GetIsInitial gets the isInitial property value. true if this is the initial domain associated with the tenant; otherwise, false.
func (m *VerifiedDomain) GetIsInitial()(*bool) {
    return m.isInitial
}
// GetName gets the name property value. The domain name; for example, contoso.onmicrosoft.com.
func (m *VerifiedDomain) GetName()(*string) {
    return m.name
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *VerifiedDomain) GetOdataType()(*string) {
    return m.odataType
}
// GetType gets the type property value. For example, Managed.
func (m *VerifiedDomain) GetType()(*string) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *VerifiedDomain) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("capabilities", m.GetCapabilities())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isDefault", m.GetIsDefault())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isInitial", m.GetIsInitial())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
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
        err := writer.WriteStringValue("type", m.GetType())
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
func (m *VerifiedDomain) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCapabilities sets the capabilities property value. For example, Email, OfficeCommunicationsOnline.
func (m *VerifiedDomain) SetCapabilities(value *string)() {
    m.capabilities = value
}
// SetIsDefault sets the isDefault property value. true if this is the default domain associated with the tenant; otherwise, false.
func (m *VerifiedDomain) SetIsDefault(value *bool)() {
    m.isDefault = value
}
// SetIsInitial sets the isInitial property value. true if this is the initial domain associated with the tenant; otherwise, false.
func (m *VerifiedDomain) SetIsInitial(value *bool)() {
    m.isInitial = value
}
// SetName sets the name property value. The domain name; for example, contoso.onmicrosoft.com.
func (m *VerifiedDomain) SetName(value *string)() {
    m.name = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *VerifiedDomain) SetOdataType(value *string)() {
    m.odataType = value
}
// SetType sets the type property value. For example, Managed.
func (m *VerifiedDomain) SetType(value *string)() {
    m.typeEscaped = value
}
