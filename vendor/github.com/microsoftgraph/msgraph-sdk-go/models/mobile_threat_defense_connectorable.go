package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MobileThreatDefenseConnectorable 
type MobileThreatDefenseConnectorable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowPartnerToCollectIOSApplicationMetadata()(*bool)
    GetAllowPartnerToCollectIOSPersonalApplicationMetadata()(*bool)
    GetAndroidDeviceBlockedOnMissingPartnerData()(*bool)
    GetAndroidEnabled()(*bool)
    GetAndroidMobileApplicationManagementEnabled()(*bool)
    GetIosDeviceBlockedOnMissingPartnerData()(*bool)
    GetIosEnabled()(*bool)
    GetIosMobileApplicationManagementEnabled()(*bool)
    GetLastHeartbeatDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetMicrosoftDefenderForEndpointAttachEnabled()(*bool)
    GetPartnerState()(*MobileThreatPartnerTenantState)
    GetPartnerUnresponsivenessThresholdInDays()(*int32)
    GetPartnerUnsupportedOsVersionBlocked()(*bool)
    GetWindowsDeviceBlockedOnMissingPartnerData()(*bool)
    GetWindowsEnabled()(*bool)
    SetAllowPartnerToCollectIOSApplicationMetadata(value *bool)()
    SetAllowPartnerToCollectIOSPersonalApplicationMetadata(value *bool)()
    SetAndroidDeviceBlockedOnMissingPartnerData(value *bool)()
    SetAndroidEnabled(value *bool)()
    SetAndroidMobileApplicationManagementEnabled(value *bool)()
    SetIosDeviceBlockedOnMissingPartnerData(value *bool)()
    SetIosEnabled(value *bool)()
    SetIosMobileApplicationManagementEnabled(value *bool)()
    SetLastHeartbeatDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetMicrosoftDefenderForEndpointAttachEnabled(value *bool)()
    SetPartnerState(value *MobileThreatPartnerTenantState)()
    SetPartnerUnresponsivenessThresholdInDays(value *int32)()
    SetPartnerUnsupportedOsVersionBlocked(value *bool)()
    SetWindowsDeviceBlockedOnMissingPartnerData(value *bool)()
    SetWindowsEnabled(value *bool)()
}
