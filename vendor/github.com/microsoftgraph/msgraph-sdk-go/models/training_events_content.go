package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TrainingEventsContent 
type TrainingEventsContent struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // List of assigned trainings and their information in an attack simulation and training campaign.
    assignedTrainingsInfos []AssignedTrainingInfoable
    // The OdataType property
    odataType *string
    // Number of users who were assigned trainings in an attack simulation and training campaign.
    trainingsAssignedUserCount *int32
}
// NewTrainingEventsContent instantiates a new trainingEventsContent and sets the default values.
func NewTrainingEventsContent()(*TrainingEventsContent) {
    m := &TrainingEventsContent{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTrainingEventsContentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTrainingEventsContentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTrainingEventsContent(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TrainingEventsContent) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAssignedTrainingsInfos gets the assignedTrainingsInfos property value. List of assigned trainings and their information in an attack simulation and training campaign.
func (m *TrainingEventsContent) GetAssignedTrainingsInfos()([]AssignedTrainingInfoable) {
    return m.assignedTrainingsInfos
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TrainingEventsContent) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["assignedTrainingsInfos"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAssignedTrainingInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AssignedTrainingInfoable, len(val))
            for i, v := range val {
                res[i] = v.(AssignedTrainingInfoable)
            }
            m.SetAssignedTrainingsInfos(res)
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
    res["trainingsAssignedUserCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTrainingsAssignedUserCount(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *TrainingEventsContent) GetOdataType()(*string) {
    return m.odataType
}
// GetTrainingsAssignedUserCount gets the trainingsAssignedUserCount property value. Number of users who were assigned trainings in an attack simulation and training campaign.
func (m *TrainingEventsContent) GetTrainingsAssignedUserCount()(*int32) {
    return m.trainingsAssignedUserCount
}
// Serialize serializes information the current object
func (m *TrainingEventsContent) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAssignedTrainingsInfos() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAssignedTrainingsInfos()))
        for i, v := range m.GetAssignedTrainingsInfos() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("assignedTrainingsInfos", cast)
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
        err := writer.WriteInt32Value("trainingsAssignedUserCount", m.GetTrainingsAssignedUserCount())
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
func (m *TrainingEventsContent) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAssignedTrainingsInfos sets the assignedTrainingsInfos property value. List of assigned trainings and their information in an attack simulation and training campaign.
func (m *TrainingEventsContent) SetAssignedTrainingsInfos(value []AssignedTrainingInfoable)() {
    m.assignedTrainingsInfos = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *TrainingEventsContent) SetOdataType(value *string)() {
    m.odataType = value
}
// SetTrainingsAssignedUserCount sets the trainingsAssignedUserCount property value. Number of users who were assigned trainings in an attack simulation and training campaign.
func (m *TrainingEventsContent) SetTrainingsAssignedUserCount(value *int32)() {
    m.trainingsAssignedUserCount = value
}
