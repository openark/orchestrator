package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AttackSimulationUser 
type AttackSimulationUser struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Display name of the user.
    displayName *string
    // Email address of the user.
    email *string
    // The OdataType property
    odataType *string
    // This is the id property value of the user resource that represents the user in the Azure Active Directory tenant.
    userId *string
}
// NewAttackSimulationUser instantiates a new attackSimulationUser and sets the default values.
func NewAttackSimulationUser()(*AttackSimulationUser) {
    m := &AttackSimulationUser{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAttackSimulationUserFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAttackSimulationUserFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAttackSimulationUser(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AttackSimulationUser) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDisplayName gets the displayName property value. Display name of the user.
func (m *AttackSimulationUser) GetDisplayName()(*string) {
    return m.displayName
}
// GetEmail gets the email property value. Email address of the user.
func (m *AttackSimulationUser) GetEmail()(*string) {
    return m.email
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AttackSimulationUser) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["email"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmail(val)
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
    res["userId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserId(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AttackSimulationUser) GetOdataType()(*string) {
    return m.odataType
}
// GetUserId gets the userId property value. This is the id property value of the user resource that represents the user in the Azure Active Directory tenant.
func (m *AttackSimulationUser) GetUserId()(*string) {
    return m.userId
}
// Serialize serializes information the current object
func (m *AttackSimulationUser) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("email", m.GetEmail())
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
        err := writer.WriteStringValue("userId", m.GetUserId())
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
func (m *AttackSimulationUser) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDisplayName sets the displayName property value. Display name of the user.
func (m *AttackSimulationUser) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetEmail sets the email property value. Email address of the user.
func (m *AttackSimulationUser) SetEmail(value *string)() {
    m.email = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AttackSimulationUser) SetOdataType(value *string)() {
    m.odataType = value
}
// SetUserId sets the userId property value. This is the id property value of the user resource that represents the user in the Azure Active Directory tenant.
func (m *AttackSimulationUser) SetUserId(value *string)() {
    m.userId = value
}
