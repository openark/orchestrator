package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// FollowupFlag 
type FollowupFlag struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The date and time that the follow-up was finished.
    completedDateTime DateTimeTimeZoneable
    // The date and time that the follow up is to be finished. Note: To set the due date, you must also specify the startDateTime; otherwise, you will get a 400 Bad Request response.
    dueDateTime DateTimeTimeZoneable
    // The status for follow-up for an item. Possible values are notFlagged, complete, and flagged.
    flagStatus *FollowupFlagStatus
    // The OdataType property
    odataType *string
    // The date and time that the follow-up is to begin.
    startDateTime DateTimeTimeZoneable
}
// NewFollowupFlag instantiates a new followupFlag and sets the default values.
func NewFollowupFlag()(*FollowupFlag) {
    m := &FollowupFlag{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateFollowupFlagFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateFollowupFlagFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFollowupFlag(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *FollowupFlag) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCompletedDateTime gets the completedDateTime property value. The date and time that the follow-up was finished.
func (m *FollowupFlag) GetCompletedDateTime()(DateTimeTimeZoneable) {
    return m.completedDateTime
}
// GetDueDateTime gets the dueDateTime property value. The date and time that the follow up is to be finished. Note: To set the due date, you must also specify the startDateTime; otherwise, you will get a 400 Bad Request response.
func (m *FollowupFlag) GetDueDateTime()(DateTimeTimeZoneable) {
    return m.dueDateTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *FollowupFlag) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["completedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDateTimeTimeZoneFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompletedDateTime(val.(DateTimeTimeZoneable))
        }
        return nil
    }
    res["dueDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDateTimeTimeZoneFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDueDateTime(val.(DateTimeTimeZoneable))
        }
        return nil
    }
    res["flagStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseFollowupFlagStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFlagStatus(val.(*FollowupFlagStatus))
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
    res["startDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDateTimeTimeZoneFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartDateTime(val.(DateTimeTimeZoneable))
        }
        return nil
    }
    return res
}
// GetFlagStatus gets the flagStatus property value. The status for follow-up for an item. Possible values are notFlagged, complete, and flagged.
func (m *FollowupFlag) GetFlagStatus()(*FollowupFlagStatus) {
    return m.flagStatus
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *FollowupFlag) GetOdataType()(*string) {
    return m.odataType
}
// GetStartDateTime gets the startDateTime property value. The date and time that the follow-up is to begin.
func (m *FollowupFlag) GetStartDateTime()(DateTimeTimeZoneable) {
    return m.startDateTime
}
// Serialize serializes information the current object
func (m *FollowupFlag) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("completedDateTime", m.GetCompletedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("dueDateTime", m.GetDueDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetFlagStatus() != nil {
        cast := (*m.GetFlagStatus()).String()
        err := writer.WriteStringValue("flagStatus", &cast)
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
        err := writer.WriteObjectValue("startDateTime", m.GetStartDateTime())
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
func (m *FollowupFlag) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCompletedDateTime sets the completedDateTime property value. The date and time that the follow-up was finished.
func (m *FollowupFlag) SetCompletedDateTime(value DateTimeTimeZoneable)() {
    m.completedDateTime = value
}
// SetDueDateTime sets the dueDateTime property value. The date and time that the follow up is to be finished. Note: To set the due date, you must also specify the startDateTime; otherwise, you will get a 400 Bad Request response.
func (m *FollowupFlag) SetDueDateTime(value DateTimeTimeZoneable)() {
    m.dueDateTime = value
}
// SetFlagStatus sets the flagStatus property value. The status for follow-up for an item. Possible values are notFlagged, complete, and flagged.
func (m *FollowupFlag) SetFlagStatus(value *FollowupFlagStatus)() {
    m.flagStatus = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *FollowupFlag) SetOdataType(value *string)() {
    m.odataType = value
}
// SetStartDateTime sets the startDateTime property value. The date and time that the follow-up is to begin.
func (m *FollowupFlag) SetStartDateTime(value DateTimeTimeZoneable)() {
    m.startDateTime = value
}
