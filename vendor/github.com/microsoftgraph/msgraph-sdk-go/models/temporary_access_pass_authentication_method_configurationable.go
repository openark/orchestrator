package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TemporaryAccessPassAuthenticationMethodConfigurationable 
type TemporaryAccessPassAuthenticationMethodConfigurationable interface {
    AuthenticationMethodConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDefaultLength()(*int32)
    GetDefaultLifetimeInMinutes()(*int32)
    GetIncludeTargets()([]AuthenticationMethodTargetable)
    GetIsUsableOnce()(*bool)
    GetMaximumLifetimeInMinutes()(*int32)
    GetMinimumLifetimeInMinutes()(*int32)
    SetDefaultLength(value *int32)()
    SetDefaultLifetimeInMinutes(value *int32)()
    SetIncludeTargets(value []AuthenticationMethodTargetable)()
    SetIsUsableOnce(value *bool)()
    SetMaximumLifetimeInMinutes(value *int32)()
    SetMinimumLifetimeInMinutes(value *int32)()
}
