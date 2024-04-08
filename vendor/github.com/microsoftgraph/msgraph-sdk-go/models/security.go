package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Security 
type Security struct {
    Entity
    // The alerts property
    alerts []Alertable
    // The attackSimulation property
    attackSimulation AttackSimulationRootable
    // The secureScoreControlProfiles property
    secureScoreControlProfiles []SecureScoreControlProfileable
    // The secureScores property
    secureScores []SecureScoreable
}
// NewSecurity instantiates a new Security and sets the default values.
func NewSecurity()(*Security) {
    m := &Security{
        Entity: *NewEntity(),
    }
    return m
}
// CreateSecurityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSecurityFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSecurity(), nil
}
// GetAlerts gets the alerts property value. The alerts property
func (m *Security) GetAlerts()([]Alertable) {
    return m.alerts
}
// GetAttackSimulation gets the attackSimulation property value. The attackSimulation property
func (m *Security) GetAttackSimulation()(AttackSimulationRootable) {
    return m.attackSimulation
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Security) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["alerts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAlertFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Alertable, len(val))
            for i, v := range val {
                res[i] = v.(Alertable)
            }
            m.SetAlerts(res)
        }
        return nil
    }
    res["attackSimulation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAttackSimulationRootFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAttackSimulation(val.(AttackSimulationRootable))
        }
        return nil
    }
    res["secureScoreControlProfiles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSecureScoreControlProfileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SecureScoreControlProfileable, len(val))
            for i, v := range val {
                res[i] = v.(SecureScoreControlProfileable)
            }
            m.SetSecureScoreControlProfiles(res)
        }
        return nil
    }
    res["secureScores"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSecureScoreFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SecureScoreable, len(val))
            for i, v := range val {
                res[i] = v.(SecureScoreable)
            }
            m.SetSecureScores(res)
        }
        return nil
    }
    return res
}
// GetSecureScoreControlProfiles gets the secureScoreControlProfiles property value. The secureScoreControlProfiles property
func (m *Security) GetSecureScoreControlProfiles()([]SecureScoreControlProfileable) {
    return m.secureScoreControlProfiles
}
// GetSecureScores gets the secureScores property value. The secureScores property
func (m *Security) GetSecureScores()([]SecureScoreable) {
    return m.secureScores
}
// Serialize serializes information the current object
func (m *Security) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAlerts() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAlerts()))
        for i, v := range m.GetAlerts() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("alerts", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("attackSimulation", m.GetAttackSimulation())
        if err != nil {
            return err
        }
    }
    if m.GetSecureScoreControlProfiles() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSecureScoreControlProfiles()))
        for i, v := range m.GetSecureScoreControlProfiles() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("secureScoreControlProfiles", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSecureScores() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSecureScores()))
        for i, v := range m.GetSecureScores() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("secureScores", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAlerts sets the alerts property value. The alerts property
func (m *Security) SetAlerts(value []Alertable)() {
    m.alerts = value
}
// SetAttackSimulation sets the attackSimulation property value. The attackSimulation property
func (m *Security) SetAttackSimulation(value AttackSimulationRootable)() {
    m.attackSimulation = value
}
// SetSecureScoreControlProfiles sets the secureScoreControlProfiles property value. The secureScoreControlProfiles property
func (m *Security) SetSecureScoreControlProfiles(value []SecureScoreControlProfileable)() {
    m.secureScoreControlProfiles = value
}
// SetSecureScores sets the secureScores property value. The secureScores property
func (m *Security) SetSecureScores(value []SecureScoreable)() {
    m.secureScores = value
}
