package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Windows10SecureAssessmentConfigurationable 
type Windows10SecureAssessmentConfigurationable interface {
    DeviceConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowPrinting()(*bool)
    GetAllowScreenCapture()(*bool)
    GetAllowTextSuggestion()(*bool)
    GetConfigurationAccount()(*string)
    GetLaunchUri()(*string)
    SetAllowPrinting(value *bool)()
    SetAllowScreenCapture(value *bool)()
    SetAllowTextSuggestion(value *bool)()
    SetConfigurationAccount(value *string)()
    SetLaunchUri(value *string)()
}
