package callrecords

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// CallRecordable 
type CallRecordable interface {
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetJoinWebUrl()(*string)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetModalities()([]Modality)
    GetOrganizer()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.IdentitySetable)
    GetParticipants()([]iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.IdentitySetable)
    GetSessions()([]Sessionable)
    GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetType()(*CallType)
    GetVersion()(*int64)
    SetEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetJoinWebUrl(value *string)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetModalities(value []Modality)()
    SetOrganizer(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.IdentitySetable)()
    SetParticipants(value []iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.IdentitySetable)()
    SetSessions(value []Sessionable)()
    SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetType(value *CallType)()
    SetVersion(value *int64)()
}
