package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Securityable 
type Securityable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAlerts()([]Alertable)
    GetAttackSimulation()(AttackSimulationRootable)
    GetSecureScoreControlProfiles()([]SecureScoreControlProfileable)
    GetSecureScores()([]SecureScoreable)
    SetAlerts(value []Alertable)()
    SetAttackSimulation(value AttackSimulationRootable)()
    SetSecureScoreControlProfiles(value []SecureScoreControlProfileable)()
    SetSecureScores(value []SecureScoreable)()
}
