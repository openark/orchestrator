package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ConditionalAccessSessionControlsable 
type ConditionalAccessSessionControlsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApplicationEnforcedRestrictions()(ApplicationEnforcedRestrictionsSessionControlable)
    GetCloudAppSecurity()(CloudAppSecuritySessionControlable)
    GetDisableResilienceDefaults()(*bool)
    GetOdataType()(*string)
    GetPersistentBrowser()(PersistentBrowserSessionControlable)
    GetSignInFrequency()(SignInFrequencySessionControlable)
    SetApplicationEnforcedRestrictions(value ApplicationEnforcedRestrictionsSessionControlable)()
    SetCloudAppSecurity(value CloudAppSecuritySessionControlable)()
    SetDisableResilienceDefaults(value *bool)()
    SetOdataType(value *string)()
    SetPersistentBrowser(value PersistentBrowserSessionControlable)()
    SetSignInFrequency(value SignInFrequencySessionControlable)()
}
