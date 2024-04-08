package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserSimulationDetails 
type UserSimulationDetails struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Number of trainings assigned to a user in an attack simulation and training campaign.
    assignedTrainingsCount *int32
    // Number of trainings completed by a user in an attack simulation and training campaign.
    completedTrainingsCount *int32
    // Date and time of the compromising online action by a user in an attack simulation and training campaign.
    compromisedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Number of trainings in progress by a user in an attack simulation and training campaign.
    inProgressTrainingsCount *int32
    // Indicates whether a user was compromised in an attack simulation and training campaign.
    isCompromised *bool
    // The OdataType property
    odataType *string
    // Date and time when a user reported the delivered payload as phishing in the attack simulation and training campaign.
    reportedPhishDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // List of simulation events of a user in the attack simulation and training campaign.
    simulationEvents []UserSimulationEventInfoable
    // User in an attack simulation and training campaign.
    simulationUser AttackSimulationUserable
    // List of training events of a user in the attack simulation and training campaign.
    trainingEvents []UserTrainingEventInfoable
}
// NewUserSimulationDetails instantiates a new userSimulationDetails and sets the default values.
func NewUserSimulationDetails()(*UserSimulationDetails) {
    m := &UserSimulationDetails{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateUserSimulationDetailsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserSimulationDetailsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserSimulationDetails(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UserSimulationDetails) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAssignedTrainingsCount gets the assignedTrainingsCount property value. Number of trainings assigned to a user in an attack simulation and training campaign.
func (m *UserSimulationDetails) GetAssignedTrainingsCount()(*int32) {
    return m.assignedTrainingsCount
}
// GetCompletedTrainingsCount gets the completedTrainingsCount property value. Number of trainings completed by a user in an attack simulation and training campaign.
func (m *UserSimulationDetails) GetCompletedTrainingsCount()(*int32) {
    return m.completedTrainingsCount
}
// GetCompromisedDateTime gets the compromisedDateTime property value. Date and time of the compromising online action by a user in an attack simulation and training campaign.
func (m *UserSimulationDetails) GetCompromisedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.compromisedDateTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserSimulationDetails) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["assignedTrainingsCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignedTrainingsCount(val)
        }
        return nil
    }
    res["completedTrainingsCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompletedTrainingsCount(val)
        }
        return nil
    }
    res["compromisedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompromisedDateTime(val)
        }
        return nil
    }
    res["inProgressTrainingsCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInProgressTrainingsCount(val)
        }
        return nil
    }
    res["isCompromised"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsCompromised(val)
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
    res["reportedPhishDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReportedPhishDateTime(val)
        }
        return nil
    }
    res["simulationEvents"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserSimulationEventInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserSimulationEventInfoable, len(val))
            for i, v := range val {
                res[i] = v.(UserSimulationEventInfoable)
            }
            m.SetSimulationEvents(res)
        }
        return nil
    }
    res["simulationUser"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAttackSimulationUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSimulationUser(val.(AttackSimulationUserable))
        }
        return nil
    }
    res["trainingEvents"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserTrainingEventInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserTrainingEventInfoable, len(val))
            for i, v := range val {
                res[i] = v.(UserTrainingEventInfoable)
            }
            m.SetTrainingEvents(res)
        }
        return nil
    }
    return res
}
// GetInProgressTrainingsCount gets the inProgressTrainingsCount property value. Number of trainings in progress by a user in an attack simulation and training campaign.
func (m *UserSimulationDetails) GetInProgressTrainingsCount()(*int32) {
    return m.inProgressTrainingsCount
}
// GetIsCompromised gets the isCompromised property value. Indicates whether a user was compromised in an attack simulation and training campaign.
func (m *UserSimulationDetails) GetIsCompromised()(*bool) {
    return m.isCompromised
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *UserSimulationDetails) GetOdataType()(*string) {
    return m.odataType
}
// GetReportedPhishDateTime gets the reportedPhishDateTime property value. Date and time when a user reported the delivered payload as phishing in the attack simulation and training campaign.
func (m *UserSimulationDetails) GetReportedPhishDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.reportedPhishDateTime
}
// GetSimulationEvents gets the simulationEvents property value. List of simulation events of a user in the attack simulation and training campaign.
func (m *UserSimulationDetails) GetSimulationEvents()([]UserSimulationEventInfoable) {
    return m.simulationEvents
}
// GetSimulationUser gets the simulationUser property value. User in an attack simulation and training campaign.
func (m *UserSimulationDetails) GetSimulationUser()(AttackSimulationUserable) {
    return m.simulationUser
}
// GetTrainingEvents gets the trainingEvents property value. List of training events of a user in the attack simulation and training campaign.
func (m *UserSimulationDetails) GetTrainingEvents()([]UserTrainingEventInfoable) {
    return m.trainingEvents
}
// Serialize serializes information the current object
func (m *UserSimulationDetails) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("assignedTrainingsCount", m.GetAssignedTrainingsCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("completedTrainingsCount", m.GetCompletedTrainingsCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("compromisedDateTime", m.GetCompromisedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("inProgressTrainingsCount", m.GetInProgressTrainingsCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isCompromised", m.GetIsCompromised())
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
        err := writer.WriteTimeValue("reportedPhishDateTime", m.GetReportedPhishDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetSimulationEvents() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSimulationEvents()))
        for i, v := range m.GetSimulationEvents() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("simulationEvents", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("simulationUser", m.GetSimulationUser())
        if err != nil {
            return err
        }
    }
    if m.GetTrainingEvents() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTrainingEvents()))
        for i, v := range m.GetTrainingEvents() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("trainingEvents", cast)
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
func (m *UserSimulationDetails) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAssignedTrainingsCount sets the assignedTrainingsCount property value. Number of trainings assigned to a user in an attack simulation and training campaign.
func (m *UserSimulationDetails) SetAssignedTrainingsCount(value *int32)() {
    m.assignedTrainingsCount = value
}
// SetCompletedTrainingsCount sets the completedTrainingsCount property value. Number of trainings completed by a user in an attack simulation and training campaign.
func (m *UserSimulationDetails) SetCompletedTrainingsCount(value *int32)() {
    m.completedTrainingsCount = value
}
// SetCompromisedDateTime sets the compromisedDateTime property value. Date and time of the compromising online action by a user in an attack simulation and training campaign.
func (m *UserSimulationDetails) SetCompromisedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.compromisedDateTime = value
}
// SetInProgressTrainingsCount sets the inProgressTrainingsCount property value. Number of trainings in progress by a user in an attack simulation and training campaign.
func (m *UserSimulationDetails) SetInProgressTrainingsCount(value *int32)() {
    m.inProgressTrainingsCount = value
}
// SetIsCompromised sets the isCompromised property value. Indicates whether a user was compromised in an attack simulation and training campaign.
func (m *UserSimulationDetails) SetIsCompromised(value *bool)() {
    m.isCompromised = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *UserSimulationDetails) SetOdataType(value *string)() {
    m.odataType = value
}
// SetReportedPhishDateTime sets the reportedPhishDateTime property value. Date and time when a user reported the delivered payload as phishing in the attack simulation and training campaign.
func (m *UserSimulationDetails) SetReportedPhishDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.reportedPhishDateTime = value
}
// SetSimulationEvents sets the simulationEvents property value. List of simulation events of a user in the attack simulation and training campaign.
func (m *UserSimulationDetails) SetSimulationEvents(value []UserSimulationEventInfoable)() {
    m.simulationEvents = value
}
// SetSimulationUser sets the simulationUser property value. User in an attack simulation and training campaign.
func (m *UserSimulationDetails) SetSimulationUser(value AttackSimulationUserable)() {
    m.simulationUser = value
}
// SetTrainingEvents sets the trainingEvents property value. List of training events of a user in the attack simulation and training campaign.
func (m *UserSimulationDetails) SetTrainingEvents(value []UserTrainingEventInfoable)() {
    m.trainingEvents = value
}
