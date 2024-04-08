package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UriClickSecurityStateable 
type UriClickSecurityStateable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetClickAction()(*string)
    GetClickDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetId()(*string)
    GetOdataType()(*string)
    GetSourceId()(*string)
    GetUriDomain()(*string)
    GetVerdict()(*string)
    SetClickAction(value *string)()
    SetClickDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetId(value *string)()
    SetOdataType(value *string)()
    SetSourceId(value *string)()
    SetUriDomain(value *string)()
    SetVerdict(value *string)()
}
