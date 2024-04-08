package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ImplicitGrantSettings 
type ImplicitGrantSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Specifies whether this web application can request an access token using the OAuth 2.0 implicit flow.
    enableAccessTokenIssuance *bool
    // Specifies whether this web application can request an ID token using the OAuth 2.0 implicit flow.
    enableIdTokenIssuance *bool
    // The OdataType property
    odataType *string
}
// NewImplicitGrantSettings instantiates a new implicitGrantSettings and sets the default values.
func NewImplicitGrantSettings()(*ImplicitGrantSettings) {
    m := &ImplicitGrantSettings{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateImplicitGrantSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateImplicitGrantSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewImplicitGrantSettings(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ImplicitGrantSettings) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetEnableAccessTokenIssuance gets the enableAccessTokenIssuance property value. Specifies whether this web application can request an access token using the OAuth 2.0 implicit flow.
func (m *ImplicitGrantSettings) GetEnableAccessTokenIssuance()(*bool) {
    return m.enableAccessTokenIssuance
}
// GetEnableIdTokenIssuance gets the enableIdTokenIssuance property value. Specifies whether this web application can request an ID token using the OAuth 2.0 implicit flow.
func (m *ImplicitGrantSettings) GetEnableIdTokenIssuance()(*bool) {
    return m.enableIdTokenIssuance
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ImplicitGrantSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["enableAccessTokenIssuance"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableAccessTokenIssuance(val)
        }
        return nil
    }
    res["enableIdTokenIssuance"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableIdTokenIssuance(val)
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
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ImplicitGrantSettings) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *ImplicitGrantSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("enableAccessTokenIssuance", m.GetEnableAccessTokenIssuance())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("enableIdTokenIssuance", m.GetEnableIdTokenIssuance())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ImplicitGrantSettings) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetEnableAccessTokenIssuance sets the enableAccessTokenIssuance property value. Specifies whether this web application can request an access token using the OAuth 2.0 implicit flow.
func (m *ImplicitGrantSettings) SetEnableAccessTokenIssuance(value *bool)() {
    m.enableAccessTokenIssuance = value
}
// SetEnableIdTokenIssuance sets the enableIdTokenIssuance property value. Specifies whether this web application can request an ID token using the OAuth 2.0 implicit flow.
func (m *ImplicitGrantSettings) SetEnableIdTokenIssuance(value *bool)() {
    m.enableIdTokenIssuance = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ImplicitGrantSettings) SetOdataType(value *string)() {
    m.odataType = value
}
