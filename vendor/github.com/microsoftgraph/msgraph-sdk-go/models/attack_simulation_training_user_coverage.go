package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AttackSimulationTrainingUserCoverage 
type AttackSimulationTrainingUserCoverage struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // User in an attack simulation and training campaign.
    attackSimulationUser AttackSimulationUserable
    // The OdataType property
    odataType *string
    // List of assigned trainings and their statuses for the user.
    userTrainings []UserTrainingStatusInfoable
}
// NewAttackSimulationTrainingUserCoverage instantiates a new attackSimulationTrainingUserCoverage and sets the default values.
func NewAttackSimulationTrainingUserCoverage()(*AttackSimulationTrainingUserCoverage) {
    m := &AttackSimulationTrainingUserCoverage{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAttackSimulationTrainingUserCoverageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAttackSimulationTrainingUserCoverageFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAttackSimulationTrainingUserCoverage(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AttackSimulationTrainingUserCoverage) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAttackSimulationUser gets the attackSimulationUser property value. User in an attack simulation and training campaign.
func (m *AttackSimulationTrainingUserCoverage) GetAttackSimulationUser()(AttackSimulationUserable) {
    return m.attackSimulationUser
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AttackSimulationTrainingUserCoverage) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["attackSimulationUser"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAttackSimulationUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAttackSimulationUser(val.(AttackSimulationUserable))
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
    res["userTrainings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserTrainingStatusInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserTrainingStatusInfoable, len(val))
            for i, v := range val {
                res[i] = v.(UserTrainingStatusInfoable)
            }
            m.SetUserTrainings(res)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AttackSimulationTrainingUserCoverage) GetOdataType()(*string) {
    return m.odataType
}
// GetUserTrainings gets the userTrainings property value. List of assigned trainings and their statuses for the user.
func (m *AttackSimulationTrainingUserCoverage) GetUserTrainings()([]UserTrainingStatusInfoable) {
    return m.userTrainings
}
// Serialize serializes information the current object
func (m *AttackSimulationTrainingUserCoverage) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("attackSimulationUser", m.GetAttackSimulationUser())
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
    if m.GetUserTrainings() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserTrainings()))
        for i, v := range m.GetUserTrainings() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("userTrainings", cast)
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
func (m *AttackSimulationTrainingUserCoverage) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAttackSimulationUser sets the attackSimulationUser property value. User in an attack simulation and training campaign.
func (m *AttackSimulationTrainingUserCoverage) SetAttackSimulationUser(value AttackSimulationUserable)() {
    m.attackSimulationUser = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AttackSimulationTrainingUserCoverage) SetOdataType(value *string)() {
    m.odataType = value
}
// SetUserTrainings sets the userTrainings property value. List of assigned trainings and their statuses for the user.
func (m *AttackSimulationTrainingUserCoverage) SetUserTrainings(value []UserTrainingStatusInfoable)() {
    m.userTrainings = value
}
