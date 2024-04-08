package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ReportRootable 
type ReportRootable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDailyPrintUsageByPrinter()([]PrintUsageByPrinterable)
    GetDailyPrintUsageByUser()([]PrintUsageByUserable)
    GetMonthlyPrintUsageByPrinter()([]PrintUsageByPrinterable)
    GetMonthlyPrintUsageByUser()([]PrintUsageByUserable)
    GetOdataType()(*string)
    GetSecurity()(SecurityReportsRootable)
    SetDailyPrintUsageByPrinter(value []PrintUsageByPrinterable)()
    SetDailyPrintUsageByUser(value []PrintUsageByUserable)()
    SetMonthlyPrintUsageByPrinter(value []PrintUsageByPrinterable)()
    SetMonthlyPrintUsageByUser(value []PrintUsageByUserable)()
    SetOdataType(value *string)()
    SetSecurity(value SecurityReportsRootable)()
}
