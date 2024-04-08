package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CopyNotebookModelable 
type CopyNotebookModelable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCreatedBy()(*string)
    GetCreatedByIdentity()(IdentitySetable)
    GetCreatedTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetId()(*string)
    GetIsDefault()(*bool)
    GetIsShared()(*bool)
    GetLastModifiedBy()(*string)
    GetLastModifiedByIdentity()(IdentitySetable)
    GetLastModifiedTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLinks()(NotebookLinksable)
    GetName()(*string)
    GetOdataType()(*string)
    GetSectionGroupsUrl()(*string)
    GetSectionsUrl()(*string)
    GetSelf()(*string)
    GetUserRole()(*OnenoteUserRole)
    SetCreatedBy(value *string)()
    SetCreatedByIdentity(value IdentitySetable)()
    SetCreatedTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetId(value *string)()
    SetIsDefault(value *bool)()
    SetIsShared(value *bool)()
    SetLastModifiedBy(value *string)()
    SetLastModifiedByIdentity(value IdentitySetable)()
    SetLastModifiedTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLinks(value NotebookLinksable)()
    SetName(value *string)()
    SetOdataType(value *string)()
    SetSectionGroupsUrl(value *string)()
    SetSectionsUrl(value *string)()
    SetSelf(value *string)()
    SetUserRole(value *OnenoteUserRole)()
}
