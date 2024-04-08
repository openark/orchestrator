package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AlertHistoryState 
type AlertHistoryState struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The appId property
    appId *string
    // The assignedTo property
    assignedTo *string
    // The comments property
    comments []string
    // The feedback property
    feedback *AlertFeedback
    // The OdataType property
    odataType *string
    // The status property
    status *AlertStatus
    // The updatedDateTime property
    updatedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The user property
    user *string
}
// NewAlertHistoryState instantiates a new alertHistoryState and sets the default values.
func NewAlertHistoryState()(*AlertHistoryState) {
    m := &AlertHistoryState{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAlertHistoryStateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAlertHistoryStateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAlertHistoryState(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AlertHistoryState) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAppId gets the appId property value. The appId property
func (m *AlertHistoryState) GetAppId()(*string) {
    return m.appId
}
// GetAssignedTo gets the assignedTo property value. The assignedTo property
func (m *AlertHistoryState) GetAssignedTo()(*string) {
    return m.assignedTo
}
// GetComments gets the comments property value. The comments property
func (m *AlertHistoryState) GetComments()([]string) {
    return m.comments
}
// GetFeedback gets the feedback property value. The feedback property
func (m *AlertHistoryState) GetFeedback()(*AlertFeedback) {
    return m.feedback
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AlertHistoryState) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["appId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppId(val)
        }
        return nil
    }
    res["assignedTo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignedTo(val)
        }
        return nil
    }
    res["comments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetComments(res)
        }
        return nil
    }
    res["feedback"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAlertFeedback)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFeedback(val.(*AlertFeedback))
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
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAlertStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*AlertStatus))
        }
        return nil
    }
    res["updatedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdatedDateTime(val)
        }
        return nil
    }
    res["user"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUser(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AlertHistoryState) GetOdataType()(*string) {
    return m.odataType
}
// GetStatus gets the status property value. The status property
func (m *AlertHistoryState) GetStatus()(*AlertStatus) {
    return m.status
}
// GetUpdatedDateTime gets the updatedDateTime property value. The updatedDateTime property
func (m *AlertHistoryState) GetUpdatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.updatedDateTime
}
// GetUser gets the user property value. The user property
func (m *AlertHistoryState) GetUser()(*string) {
    return m.user
}
// Serialize serializes information the current object
func (m *AlertHistoryState) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("appId", m.GetAppId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("assignedTo", m.GetAssignedTo())
        if err != nil {
            return err
        }
    }
    if m.GetComments() != nil {
        err := writer.WriteCollectionOfStringValues("comments", m.GetComments())
        if err != nil {
            return err
        }
    }
    if m.GetFeedback() != nil {
        cast := (*m.GetFeedback()).String()
        err := writer.WriteStringValue("feedback", &cast)
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
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err := writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("updatedDateTime", m.GetUpdatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("user", m.GetUser())
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
func (m *AlertHistoryState) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAppId sets the appId property value. The appId property
func (m *AlertHistoryState) SetAppId(value *string)() {
    m.appId = value
}
// SetAssignedTo sets the assignedTo property value. The assignedTo property
func (m *AlertHistoryState) SetAssignedTo(value *string)() {
    m.assignedTo = value
}
// SetComments sets the comments property value. The comments property
func (m *AlertHistoryState) SetComments(value []string)() {
    m.comments = value
}
// SetFeedback sets the feedback property value. The feedback property
func (m *AlertHistoryState) SetFeedback(value *AlertFeedback)() {
    m.feedback = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AlertHistoryState) SetOdataType(value *string)() {
    m.odataType = value
}
// SetStatus sets the status property value. The status property
func (m *AlertHistoryState) SetStatus(value *AlertStatus)() {
    m.status = value
}
// SetUpdatedDateTime sets the updatedDateTime property value. The updatedDateTime property
func (m *AlertHistoryState) SetUpdatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.updatedDateTime = value
}
// SetUser sets the user property value. The user property
func (m *AlertHistoryState) SetUser(value *string)() {
    m.user = value
}
