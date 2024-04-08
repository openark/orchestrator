package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MeetingParticipantInfo 
type MeetingParticipantInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Identity information of the participant.
    identity IdentitySetable
    // The OdataType property
    odataType *string
    // Specifies the participant's role in the meeting.
    role *OnlineMeetingRole
    // User principal name of the participant.
    upn *string
}
// NewMeetingParticipantInfo instantiates a new meetingParticipantInfo and sets the default values.
func NewMeetingParticipantInfo()(*MeetingParticipantInfo) {
    m := &MeetingParticipantInfo{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateMeetingParticipantInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMeetingParticipantInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMeetingParticipantInfo(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MeetingParticipantInfo) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MeetingParticipantInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["role"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseOnlineMeetingRole)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRole(val.(*OnlineMeetingRole))
        }
        return nil
    }
    res["upn"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpn(val)
        }
        return nil
    }
    return res
}
// GetIdentity gets the identity property value. Identity information of the participant.
func (m *MeetingParticipantInfo) GetIdentity()(IdentitySetable) {
    return m.identity
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *MeetingParticipantInfo) GetOdataType()(*string) {
    return m.odataType
}
// GetRole gets the role property value. Specifies the participant's role in the meeting.
func (m *MeetingParticipantInfo) GetRole()(*OnlineMeetingRole) {
    return m.role
}
// GetUpn gets the upn property value. User principal name of the participant.
func (m *MeetingParticipantInfo) GetUpn()(*string) {
    return m.upn
}
// Serialize serializes information the current object
func (m *MeetingParticipantInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("identity", m.GetIdentity())
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
    if m.GetRole() != nil {
        cast := (*m.GetRole()).String()
        err := writer.WriteStringValue("role", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("upn", m.GetUpn())
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
func (m *MeetingParticipantInfo) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetIdentity sets the identity property value. Identity information of the participant.
func (m *MeetingParticipantInfo) SetIdentity(value IdentitySetable)() {
    m.identity = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *MeetingParticipantInfo) SetOdataType(value *string)() {
    m.odataType = value
}
// SetRole sets the role property value. Specifies the participant's role in the meeting.
func (m *MeetingParticipantInfo) SetRole(value *OnlineMeetingRole)() {
    m.role = value
}
// SetUpn sets the upn property value. User principal name of the participant.
func (m *MeetingParticipantInfo) SetUpn(value *string)() {
    m.upn = value
}
