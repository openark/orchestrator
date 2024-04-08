package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PrintJobConfigurationable 
type PrintJobConfigurationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCollate()(*bool)
    GetColorMode()(*PrintColorMode)
    GetCopies()(*int32)
    GetDpi()(*int32)
    GetDuplexMode()(*PrintDuplexMode)
    GetFeedOrientation()(*PrinterFeedOrientation)
    GetFinishings()([]PrintFinishing)
    GetFitPdfToPage()(*bool)
    GetInputBin()(*string)
    GetMargin()(PrintMarginable)
    GetMediaSize()(*string)
    GetMediaType()(*string)
    GetMultipageLayout()(*PrintMultipageLayout)
    GetOdataType()(*string)
    GetOrientation()(*PrintOrientation)
    GetOutputBin()(*string)
    GetPageRanges()([]IntegerRangeable)
    GetPagesPerSheet()(*int32)
    GetQuality()(*PrintQuality)
    GetScaling()(*PrintScaling)
    SetCollate(value *bool)()
    SetColorMode(value *PrintColorMode)()
    SetCopies(value *int32)()
    SetDpi(value *int32)()
    SetDuplexMode(value *PrintDuplexMode)()
    SetFeedOrientation(value *PrinterFeedOrientation)()
    SetFinishings(value []PrintFinishing)()
    SetFitPdfToPage(value *bool)()
    SetInputBin(value *string)()
    SetMargin(value PrintMarginable)()
    SetMediaSize(value *string)()
    SetMediaType(value *string)()
    SetMultipageLayout(value *PrintMultipageLayout)()
    SetOdataType(value *string)()
    SetOrientation(value *PrintOrientation)()
    SetOutputBin(value *string)()
    SetPageRanges(value []IntegerRangeable)()
    SetPagesPerSheet(value *int32)()
    SetQuality(value *PrintQuality)()
    SetScaling(value *PrintScaling)()
}
