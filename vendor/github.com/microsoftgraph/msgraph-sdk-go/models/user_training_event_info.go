package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserTrainingEventInfo 
type UserTrainingEventInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Display name of the training.
    displayName *string
    // Latest status of the training assigned to the user. Possible values are: unknown, assigned, inProgress, completed, overdue, unknownFutureValue.
    latestTrainingStatus *TrainingStatus
    // The OdataType property
    odataType *string
    // Event details of the training when it was assigned to the user.
    trainingAssignedProperties UserTrainingContentEventInfoable
    // Event details of the training when it was completed by the user.
    trainingCompletedProperties UserTrainingContentEventInfoable
    // Event details of the training when it was updated/in-progress by the user.
    trainingUpdatedProperties UserTrainingContentEventInfoable
}
// NewUserTrainingEventInfo instantiates a new userTrainingEventInfo and sets the default values.
func NewUserTrainingEventInfo()(*UserTrainingEventInfo) {
    m := &UserTrainingEventInfo{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateUserTrainingEventInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserTrainingEventInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserTrainingEventInfo(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UserTrainingEventInfo) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDisplayName gets the displayName property value. Display name of the training.
func (m *UserTrainingEventInfo) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserTrainingEventInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["latestTrainingStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseTrainingStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLatestTrainingStatus(val.(*TrainingStatus))
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
    res["trainingAssignedProperties"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserTrainingContentEventInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTrainingAssignedProperties(val.(UserTrainingContentEventInfoable))
        }
        return nil
    }
    res["trainingCompletedProperties"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserTrainingContentEventInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTrainingCompletedProperties(val.(UserTrainingContentEventInfoable))
        }
        return nil
    }
    res["trainingUpdatedProperties"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserTrainingContentEventInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTrainingUpdatedProperties(val.(UserTrainingContentEventInfoable))
        }
        return nil
    }
    return res
}
// GetLatestTrainingStatus gets the latestTrainingStatus property value. Latest status of the training assigned to the user. Possible values are: unknown, assigned, inProgress, completed, overdue, unknownFutureValue.
func (m *UserTrainingEventInfo) GetLatestTrainingStatus()(*TrainingStatus) {
    return m.latestTrainingStatus
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *UserTrainingEventInfo) GetOdataType()(*string) {
    return m.odataType
}
// GetTrainingAssignedProperties gets the trainingAssignedProperties property value. Event details of the training when it was assigned to the user.
func (m *UserTrainingEventInfo) GetTrainingAssignedProperties()(UserTrainingContentEventInfoable) {
    return m.trainingAssignedProperties
}
// GetTrainingCompletedProperties gets the trainingCompletedProperties property value. Event details of the training when it was completed by the user.
func (m *UserTrainingEventInfo) GetTrainingCompletedProperties()(UserTrainingContentEventInfoable) {
    return m.trainingCompletedProperties
}
// GetTrainingUpdatedProperties gets the trainingUpdatedProperties property value. Event details of the training when it was updated/in-progress by the user.
func (m *UserTrainingEventInfo) GetTrainingUpdatedProperties()(UserTrainingContentEventInfoable) {
    return m.trainingUpdatedProperties
}
// Serialize serializes information the current object
func (m *UserTrainingEventInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetLatestTrainingStatus() != nil {
        cast := (*m.GetLatestTrainingStatus()).String()
        err := writer.WriteStringValue("latestTrainingStatus", &cast)
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
        err := writer.WriteObjectValue("trainingAssignedProperties", m.GetTrainingAssignedProperties())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("trainingCompletedProperties", m.GetTrainingCompletedProperties())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("trainingUpdatedProperties", m.GetTrainingUpdatedProperties())
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
func (m *UserTrainingEventInfo) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDisplayName sets the displayName property value. Display name of the training.
func (m *UserTrainingEventInfo) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetLatestTrainingStatus sets the latestTrainingStatus property value. Latest status of the training assigned to the user. Possible values are: unknown, assigned, inProgress, completed, overdue, unknownFutureValue.
func (m *UserTrainingEventInfo) SetLatestTrainingStatus(value *TrainingStatus)() {
    m.latestTrainingStatus = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *UserTrainingEventInfo) SetOdataType(value *string)() {
    m.odataType = value
}
// SetTrainingAssignedProperties sets the trainingAssignedProperties property value. Event details of the training when it was assigned to the user.
func (m *UserTrainingEventInfo) SetTrainingAssignedProperties(value UserTrainingContentEventInfoable)() {
    m.trainingAssignedProperties = value
}
// SetTrainingCompletedProperties sets the trainingCompletedProperties property value. Event details of the training when it was completed by the user.
func (m *UserTrainingEventInfo) SetTrainingCompletedProperties(value UserTrainingContentEventInfoable)() {
    m.trainingCompletedProperties = value
}
// SetTrainingUpdatedProperties sets the trainingUpdatedProperties property value. Event details of the training when it was updated/in-progress by the user.
func (m *UserTrainingEventInfo) SetTrainingUpdatedProperties(value UserTrainingContentEventInfoable)() {
    m.trainingUpdatedProperties = value
}
