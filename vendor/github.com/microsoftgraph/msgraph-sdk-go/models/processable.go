package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Processable 
type Processable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccountName()(*string)
    GetCommandLine()(*string)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetFileHash()(FileHashable)
    GetIntegrityLevel()(*ProcessIntegrityLevel)
    GetIsElevated()(*bool)
    GetName()(*string)
    GetOdataType()(*string)
    GetParentProcessCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetParentProcessId()(*int32)
    GetParentProcessName()(*string)
    GetPath()(*string)
    GetProcessId()(*int32)
    SetAccountName(value *string)()
    SetCommandLine(value *string)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetFileHash(value FileHashable)()
    SetIntegrityLevel(value *ProcessIntegrityLevel)()
    SetIsElevated(value *bool)()
    SetName(value *string)()
    SetOdataType(value *string)()
    SetParentProcessCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetParentProcessId(value *int32)()
    SetParentProcessName(value *string)()
    SetPath(value *string)()
    SetProcessId(value *int32)()
}
