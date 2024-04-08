package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceComplianceActionItemable 
type DeviceComplianceActionItemable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActionType()(*DeviceComplianceActionType)
    GetGracePeriodHours()(*int32)
    GetNotificationMessageCCList()([]string)
    GetNotificationTemplateId()(*string)
    SetActionType(value *DeviceComplianceActionType)()
    SetGracePeriodHours(value *int32)()
    SetNotificationMessageCCList(value []string)()
    SetNotificationTemplateId(value *string)()
}
