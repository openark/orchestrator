package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PrinterDefaultsable 
type PrinterDefaultsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetColorMode()(*PrintColorMode)
    GetContentType()(*string)
    GetCopiesPerJob()(*int32)
    GetDpi()(*int32)
    GetDuplexMode()(*PrintDuplexMode)
    GetFinishings()([]PrintFinishing)
    GetFitPdfToPage()(*bool)
    GetInputBin()(*string)
    GetMediaColor()(*string)
    GetMediaSize()(*string)
    GetMediaType()(*string)
    GetMultipageLayout()(*PrintMultipageLayout)
    GetOdataType()(*string)
    GetOrientation()(*PrintOrientation)
    GetOutputBin()(*string)
    GetPagesPerSheet()(*int32)
    GetQuality()(*PrintQuality)
    GetScaling()(*PrintScaling)
    SetColorMode(value *PrintColorMode)()
    SetContentType(value *string)()
    SetCopiesPerJob(value *int32)()
    SetDpi(value *int32)()
    SetDuplexMode(value *PrintDuplexMode)()
    SetFinishings(value []PrintFinishing)()
    SetFitPdfToPage(value *bool)()
    SetInputBin(value *string)()
    SetMediaColor(value *string)()
    SetMediaSize(value *string)()
    SetMediaType(value *string)()
    SetMultipageLayout(value *PrintMultipageLayout)()
    SetOdataType(value *string)()
    SetOrientation(value *PrintOrientation)()
    SetOutputBin(value *string)()
    SetPagesPerSheet(value *int32)()
    SetQuality(value *PrintQuality)()
    SetScaling(value *PrintScaling)()
}
