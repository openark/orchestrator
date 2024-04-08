package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AttackSimulationSimulationUserCoverage 
type AttackSimulationSimulationUserCoverage struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // User in an attack simulation and training campaign.
    attackSimulationUser AttackSimulationUserable
    // Number of link clicks in the received payloads by the user in attack simulation and training campaigns.
    clickCount *int32
    // Number of compromising actions by the user in attack simulation and training campaigns.
    compromisedCount *int32
    // Date and time of the latest attack simulation and training campaign that the user was included in.
    latestSimulationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The OdataType property
    odataType *string
    // Number of attack simulation and training campaigns that the user was included in.
    simulationCount *int32
}
// NewAttackSimulationSimulationUserCoverage instantiates a new attackSimulationSimulationUserCoverage and sets the default values.
func NewAttackSimulationSimulationUserCoverage()(*AttackSimulationSimulationUserCoverage) {
    m := &AttackSimulationSimulationUserCoverage{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAttackSimulationSimulationUserCoverageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAttackSimulationSimulationUserCoverageFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAttackSimulationSimulationUserCoverage(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AttackSimulationSimulationUserCoverage) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAttackSimulationUser gets the attackSimulationUser property value. User in an attack simulation and training campaign.
func (m *AttackSimulationSimulationUserCoverage) GetAttackSimulationUser()(AttackSimulationUserable) {
    return m.attackSimulationUser
}
// GetClickCount gets the clickCount property value. Number of link clicks in the received payloads by the user in attack simulation and training campaigns.
func (m *AttackSimulationSimulationUserCoverage) GetClickCount()(*int32) {
    return m.clickCount
}
// GetCompromisedCount gets the compromisedCount property value. Number of compromising actions by the user in attack simulation and training campaigns.
func (m *AttackSimulationSimulationUserCoverage) GetCompromisedCount()(*int32) {
    return m.compromisedCount
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AttackSimulationSimulationUserCoverage) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["clickCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClickCount(val)
        }
        return nil
    }
    res["compromisedCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompromisedCount(val)
        }
        return nil
    }
    res["latestSimulationDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLatestSimulationDateTime(val)
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
    res["simulationCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSimulationCount(val)
        }
        return nil
    }
    return res
}
// GetLatestSimulationDateTime gets the latestSimulationDateTime property value. Date and time of the latest attack simulation and training campaign that the user was included in.
func (m *AttackSimulationSimulationUserCoverage) GetLatestSimulationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.latestSimulationDateTime
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AttackSimulationSimulationUserCoverage) GetOdataType()(*string) {
    return m.odataType
}
// GetSimulationCount gets the simulationCount property value. Number of attack simulation and training campaigns that the user was included in.
func (m *AttackSimulationSimulationUserCoverage) GetSimulationCount()(*int32) {
    return m.simulationCount
}
// Serialize serializes information the current object
func (m *AttackSimulationSimulationUserCoverage) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("attackSimulationUser", m.GetAttackSimulationUser())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("clickCount", m.GetClickCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("compromisedCount", m.GetCompromisedCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("latestSimulationDateTime", m.GetLatestSimulationDateTime())
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
        err := writer.WriteInt32Value("simulationCount", m.GetSimulationCount())
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
func (m *AttackSimulationSimulationUserCoverage) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAttackSimulationUser sets the attackSimulationUser property value. User in an attack simulation and training campaign.
func (m *AttackSimulationSimulationUserCoverage) SetAttackSimulationUser(value AttackSimulationUserable)() {
    m.attackSimulationUser = value
}
// SetClickCount sets the clickCount property value. Number of link clicks in the received payloads by the user in attack simulation and training campaigns.
func (m *AttackSimulationSimulationUserCoverage) SetClickCount(value *int32)() {
    m.clickCount = value
}
// SetCompromisedCount sets the compromisedCount property value. Number of compromising actions by the user in attack simulation and training campaigns.
func (m *AttackSimulationSimulationUserCoverage) SetCompromisedCount(value *int32)() {
    m.compromisedCount = value
}
// SetLatestSimulationDateTime sets the latestSimulationDateTime property value. Date and time of the latest attack simulation and training campaign that the user was included in.
func (m *AttackSimulationSimulationUserCoverage) SetLatestSimulationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.latestSimulationDateTime = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AttackSimulationSimulationUserCoverage) SetOdataType(value *string)() {
    m.odataType = value
}
// SetSimulationCount sets the simulationCount property value. Number of attack simulation and training campaigns that the user was included in.
func (m *AttackSimulationSimulationUserCoverage) SetSimulationCount(value *int32)() {
    m.simulationCount = value
}
