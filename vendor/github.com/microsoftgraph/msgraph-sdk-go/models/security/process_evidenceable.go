package security

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ProcessEvidenceable 
type ProcessEvidenceable interface {
    AlertEvidenceable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDetectionStatus()(*DetectionStatus)
    GetImageFile()(FileDetailsable)
    GetMdeDeviceId()(*string)
    GetParentProcessCreationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetParentProcessId()(*int64)
    GetParentProcessImageFile()(FileDetailsable)
    GetProcessCommandLine()(*string)
    GetProcessCreationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetProcessId()(*int64)
    GetUserAccount()(UserAccountable)
    SetDetectionStatus(value *DetectionStatus)()
    SetImageFile(value FileDetailsable)()
    SetMdeDeviceId(value *string)()
    SetParentProcessCreationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetParentProcessId(value *int64)()
    SetParentProcessImageFile(value FileDetailsable)()
    SetProcessCommandLine(value *string)()
    SetProcessCreationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetProcessId(value *int64)()
    SetUserAccount(value UserAccountable)()
}
