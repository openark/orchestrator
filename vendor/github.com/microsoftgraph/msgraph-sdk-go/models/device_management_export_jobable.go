package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementExportJobable 
type DeviceManagementExportJobable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetFilter()(*string)
    GetFormat()(*DeviceManagementReportFileFormat)
    GetLocalizationType()(*DeviceManagementExportJobLocalizationType)
    GetReportName()(*string)
    GetRequestDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetSelect()([]string)
    GetSnapshotId()(*string)
    GetStatus()(*DeviceManagementReportStatus)
    GetUrl()(*string)
    SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetFilter(value *string)()
    SetFormat(value *DeviceManagementReportFileFormat)()
    SetLocalizationType(value *DeviceManagementExportJobLocalizationType)()
    SetReportName(value *string)()
    SetRequestDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetSelect(value []string)()
    SetSnapshotId(value *string)()
    SetStatus(value *DeviceManagementReportStatus)()
    SetUrl(value *string)()
}
