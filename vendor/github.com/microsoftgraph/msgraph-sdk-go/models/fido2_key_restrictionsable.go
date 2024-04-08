package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Fido2KeyRestrictionsable 
type Fido2KeyRestrictionsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAaGuids()([]string)
    GetEnforcementType()(*Fido2RestrictionEnforcementType)
    GetIsEnforced()(*bool)
    GetOdataType()(*string)
    SetAaGuids(value []string)()
    SetEnforcementType(value *Fido2RestrictionEnforcementType)()
    SetIsEnforced(value *bool)()
    SetOdataType(value *string)()
}
