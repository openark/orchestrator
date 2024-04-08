package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IdentityProtectionRootable 
type IdentityProtectionRootable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetOdataType()(*string)
    GetRiskDetections()([]RiskDetectionable)
    GetRiskyServicePrincipals()([]RiskyServicePrincipalable)
    GetRiskyUsers()([]RiskyUserable)
    GetServicePrincipalRiskDetections()([]ServicePrincipalRiskDetectionable)
    SetOdataType(value *string)()
    SetRiskDetections(value []RiskDetectionable)()
    SetRiskyServicePrincipals(value []RiskyServicePrincipalable)()
    SetRiskyUsers(value []RiskyUserable)()
    SetServicePrincipalRiskDetections(value []ServicePrincipalRiskDetectionable)()
}
