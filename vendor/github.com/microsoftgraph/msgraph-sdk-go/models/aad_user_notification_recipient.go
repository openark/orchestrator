package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AadUserNotificationRecipient 
type AadUserNotificationRecipient struct {
    TeamworkNotificationRecipient
    // Azure AD user identifier. Use the List users method to get this ID.
    userId *string
}
// NewAadUserNotificationRecipient instantiates a new AadUserNotificationRecipient and sets the default values.
func NewAadUserNotificationRecipient()(*AadUserNotificationRecipient) {
    m := &AadUserNotificationRecipient{
        TeamworkNotificationRecipient: *NewTeamworkNotificationRecipient(),
    }
    odataTypeValue := "#microsoft.graph.aadUserNotificationRecipient"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateAadUserNotificationRecipientFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAadUserNotificationRecipientFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAadUserNotificationRecipient(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AadUserNotificationRecipient) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.TeamworkNotificationRecipient.GetFieldDeserializers()
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
// GetUserId gets the userId property value. Azure AD user identifier. Use the List users method to get this ID.
func (m *AadUserNotificationRecipient) GetUserId()(*string) {
    return m.userId
}
// Serialize serializes information the current object
func (m *AadUserNotificationRecipient) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.TeamworkNotificationRecipient.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("userId", m.GetUserId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetUserId sets the userId property value. Azure AD user identifier. Use the List users method to get this ID.
func (m *AadUserNotificationRecipient) SetUserId(value *string)() {
    m.userId = value
}
