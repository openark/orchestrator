package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PasswordResetResponse 
type PasswordResetResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The Azure AD-generated password.
    newPassword *string
    // The OdataType property
    odataType *string
}
// NewPasswordResetResponse instantiates a new passwordResetResponse and sets the default values.
func NewPasswordResetResponse()(*PasswordResetResponse) {
    m := &PasswordResetResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePasswordResetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePasswordResetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPasswordResetResponse(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PasswordResetResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PasswordResetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["newPassword"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNewPassword(val)
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
// GetNewPassword gets the newPassword property value. The Azure AD-generated password.
func (m *PasswordResetResponse) GetNewPassword()(*string) {
    return m.newPassword
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *PasswordResetResponse) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *PasswordResetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("newPassword", m.GetNewPassword())
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
func (m *PasswordResetResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetNewPassword sets the newPassword property value. The Azure AD-generated password.
func (m *PasswordResetResponse) SetNewPassword(value *string)() {
    m.newPassword = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *PasswordResetResponse) SetOdataType(value *string)() {
    m.odataType = value
}
