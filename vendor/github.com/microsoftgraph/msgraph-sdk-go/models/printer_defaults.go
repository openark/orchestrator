package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PrinterDefaults 
type PrinterDefaults struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The default color mode to use when printing the document. Valid values are described in the following table.
    colorMode *PrintColorMode
    // The default content (MIME) type to use when processing documents.
    contentType *string
    // The default number of copies printed per job.
    copiesPerJob *int32
    // The default resolution in DPI to use when printing the job.
    dpi *int32
    // The default duplex (double-sided) configuration to use when printing a document. Valid values are described in the following table.
    duplexMode *PrintDuplexMode
    // The default set of finishings to apply to print jobs. Valid values are described in the following table.
    finishings []PrintFinishing
    // The default fitPdfToPage setting. True to fit each page of a PDF document to a physical sheet of media; false to let the printer decide how to lay out impressions.
    fitPdfToPage *bool
    // The inputBin property
    inputBin *string
    // The default media (such as paper) color to print the document on.
    mediaColor *string
    // The default media size to use. Supports standard size names for ISO and ANSI media sizes. Valid values are listed in the printerCapabilities topic.
    mediaSize *string
    // The default media (such as paper) type to print the document on.
    mediaType *string
    // The default direction to lay out pages when multiple pages are being printed per sheet. Valid values are described in the following table.
    multipageLayout *PrintMultipageLayout
    // The OdataType property
    odataType *string
    // The default orientation to use when printing the document. Valid values are described in the following table.
    orientation *PrintOrientation
    // The default output bin to place completed prints into. See the printer's capabilities for a list of supported output bins.
    outputBin *string
    // The default number of document pages to print on each sheet.
    pagesPerSheet *int32
    // The default quality to use when printing the document. Valid values are described in the following table.
    quality *PrintQuality
    // Specifies how the printer scales the document data to fit the requested media. Valid values are described in the following table.
    scaling *PrintScaling
}
// NewPrinterDefaults instantiates a new printerDefaults and sets the default values.
func NewPrinterDefaults()(*PrinterDefaults) {
    m := &PrinterDefaults{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePrinterDefaultsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePrinterDefaultsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPrinterDefaults(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PrinterDefaults) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetColorMode gets the colorMode property value. The default color mode to use when printing the document. Valid values are described in the following table.
func (m *PrinterDefaults) GetColorMode()(*PrintColorMode) {
    return m.colorMode
}
// GetContentType gets the contentType property value. The default content (MIME) type to use when processing documents.
func (m *PrinterDefaults) GetContentType()(*string) {
    return m.contentType
}
// GetCopiesPerJob gets the copiesPerJob property value. The default number of copies printed per job.
func (m *PrinterDefaults) GetCopiesPerJob()(*int32) {
    return m.copiesPerJob
}
// GetDpi gets the dpi property value. The default resolution in DPI to use when printing the job.
func (m *PrinterDefaults) GetDpi()(*int32) {
    return m.dpi
}
// GetDuplexMode gets the duplexMode property value. The default duplex (double-sided) configuration to use when printing a document. Valid values are described in the following table.
func (m *PrinterDefaults) GetDuplexMode()(*PrintDuplexMode) {
    return m.duplexMode
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PrinterDefaults) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["colorMode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePrintColorMode)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetColorMode(val.(*PrintColorMode))
        }
        return nil
    }
    res["contentType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentType(val)
        }
        return nil
    }
    res["copiesPerJob"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCopiesPerJob(val)
        }
        return nil
    }
    res["dpi"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDpi(val)
        }
        return nil
    }
    res["duplexMode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePrintDuplexMode)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDuplexMode(val.(*PrintDuplexMode))
        }
        return nil
    }
    res["finishings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParsePrintFinishing)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrintFinishing, len(val))
            for i, v := range val {
                res[i] = *(v.(*PrintFinishing))
            }
            m.SetFinishings(res)
        }
        return nil
    }
    res["fitPdfToPage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFitPdfToPage(val)
        }
        return nil
    }
    res["inputBin"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInputBin(val)
        }
        return nil
    }
    res["mediaColor"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMediaColor(val)
        }
        return nil
    }
    res["mediaSize"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMediaSize(val)
        }
        return nil
    }
    res["mediaType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMediaType(val)
        }
        return nil
    }
    res["multipageLayout"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePrintMultipageLayout)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMultipageLayout(val.(*PrintMultipageLayout))
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["orientation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePrintOrientation)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrientation(val.(*PrintOrientation))
        }
        return nil
    }
    res["outputBin"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOutputBin(val)
        }
        return nil
    }
    res["pagesPerSheet"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPagesPerSheet(val)
        }
        return nil
    }
    res["quality"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePrintQuality)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQuality(val.(*PrintQuality))
        }
        return nil
    }
    res["scaling"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePrintScaling)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScaling(val.(*PrintScaling))
        }
        return nil
    }
    return res
}
// GetFinishings gets the finishings property value. The default set of finishings to apply to print jobs. Valid values are described in the following table.
func (m *PrinterDefaults) GetFinishings()([]PrintFinishing) {
    return m.finishings
}
// GetFitPdfToPage gets the fitPdfToPage property value. The default fitPdfToPage setting. True to fit each page of a PDF document to a physical sheet of media; false to let the printer decide how to lay out impressions.
func (m *PrinterDefaults) GetFitPdfToPage()(*bool) {
    return m.fitPdfToPage
}
// GetInputBin gets the inputBin property value. The inputBin property
func (m *PrinterDefaults) GetInputBin()(*string) {
    return m.inputBin
}
// GetMediaColor gets the mediaColor property value. The default media (such as paper) color to print the document on.
func (m *PrinterDefaults) GetMediaColor()(*string) {
    return m.mediaColor
}
// GetMediaSize gets the mediaSize property value. The default media size to use. Supports standard size names for ISO and ANSI media sizes. Valid values are listed in the printerCapabilities topic.
func (m *PrinterDefaults) GetMediaSize()(*string) {
    return m.mediaSize
}
// GetMediaType gets the mediaType property value. The default media (such as paper) type to print the document on.
func (m *PrinterDefaults) GetMediaType()(*string) {
    return m.mediaType
}
// GetMultipageLayout gets the multipageLayout property value. The default direction to lay out pages when multiple pages are being printed per sheet. Valid values are described in the following table.
func (m *PrinterDefaults) GetMultipageLayout()(*PrintMultipageLayout) {
    return m.multipageLayout
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *PrinterDefaults) GetOdataType()(*string) {
    return m.odataType
}
// GetOrientation gets the orientation property value. The default orientation to use when printing the document. Valid values are described in the following table.
func (m *PrinterDefaults) GetOrientation()(*PrintOrientation) {
    return m.orientation
}
// GetOutputBin gets the outputBin property value. The default output bin to place completed prints into. See the printer's capabilities for a list of supported output bins.
func (m *PrinterDefaults) GetOutputBin()(*string) {
    return m.outputBin
}
// GetPagesPerSheet gets the pagesPerSheet property value. The default number of document pages to print on each sheet.
func (m *PrinterDefaults) GetPagesPerSheet()(*int32) {
    return m.pagesPerSheet
}
// GetQuality gets the quality property value. The default quality to use when printing the document. Valid values are described in the following table.
func (m *PrinterDefaults) GetQuality()(*PrintQuality) {
    return m.quality
}
// GetScaling gets the scaling property value. Specifies how the printer scales the document data to fit the requested media. Valid values are described in the following table.
func (m *PrinterDefaults) GetScaling()(*PrintScaling) {
    return m.scaling
}
// Serialize serializes information the current object
func (m *PrinterDefaults) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetColorMode() != nil {
        cast := (*m.GetColorMode()).String()
        err := writer.WriteStringValue("colorMode", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("contentType", m.GetContentType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("copiesPerJob", m.GetCopiesPerJob())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("dpi", m.GetDpi())
        if err != nil {
            return err
        }
    }
    if m.GetDuplexMode() != nil {
        cast := (*m.GetDuplexMode()).String()
        err := writer.WriteStringValue("duplexMode", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetFinishings() != nil {
        err := writer.WriteCollectionOfStringValues("finishings", SerializePrintFinishing(m.GetFinishings()))
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("fitPdfToPage", m.GetFitPdfToPage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("inputBin", m.GetInputBin())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("mediaColor", m.GetMediaColor())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("mediaSize", m.GetMediaSize())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("mediaType", m.GetMediaType())
        if err != nil {
            return err
        }
    }
    if m.GetMultipageLayout() != nil {
        cast := (*m.GetMultipageLayout()).String()
        err := writer.WriteStringValue("multipageLayout", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    if m.GetOrientation() != nil {
        cast := (*m.GetOrientation()).String()
        err := writer.WriteStringValue("orientation", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("outputBin", m.GetOutputBin())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("pagesPerSheet", m.GetPagesPerSheet())
        if err != nil {
            return err
        }
    }
    if m.GetQuality() != nil {
        cast := (*m.GetQuality()).String()
        err := writer.WriteStringValue("quality", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetScaling() != nil {
        cast := (*m.GetScaling()).String()
        err := writer.WriteStringValue("scaling", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PrinterDefaults) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetColorMode sets the colorMode property value. The default color mode to use when printing the document. Valid values are described in the following table.
func (m *PrinterDefaults) SetColorMode(value *PrintColorMode)() {
    m.colorMode = value
}
// SetContentType sets the contentType property value. The default content (MIME) type to use when processing documents.
func (m *PrinterDefaults) SetContentType(value *string)() {
    m.contentType = value
}
// SetCopiesPerJob sets the copiesPerJob property value. The default number of copies printed per job.
func (m *PrinterDefaults) SetCopiesPerJob(value *int32)() {
    m.copiesPerJob = value
}
// SetDpi sets the dpi property value. The default resolution in DPI to use when printing the job.
func (m *PrinterDefaults) SetDpi(value *int32)() {
    m.dpi = value
}
// SetDuplexMode sets the duplexMode property value. The default duplex (double-sided) configuration to use when printing a document. Valid values are described in the following table.
func (m *PrinterDefaults) SetDuplexMode(value *PrintDuplexMode)() {
    m.duplexMode = value
}
// SetFinishings sets the finishings property value. The default set of finishings to apply to print jobs. Valid values are described in the following table.
func (m *PrinterDefaults) SetFinishings(value []PrintFinishing)() {
    m.finishings = value
}
// SetFitPdfToPage sets the fitPdfToPage property value. The default fitPdfToPage setting. True to fit each page of a PDF document to a physical sheet of media; false to let the printer decide how to lay out impressions.
func (m *PrinterDefaults) SetFitPdfToPage(value *bool)() {
    m.fitPdfToPage = value
}
// SetInputBin sets the inputBin property value. The inputBin property
func (m *PrinterDefaults) SetInputBin(value *string)() {
    m.inputBin = value
}
// SetMediaColor sets the mediaColor property value. The default media (such as paper) color to print the document on.
func (m *PrinterDefaults) SetMediaColor(value *string)() {
    m.mediaColor = value
}
// SetMediaSize sets the mediaSize property value. The default media size to use. Supports standard size names for ISO and ANSI media sizes. Valid values are listed in the printerCapabilities topic.
func (m *PrinterDefaults) SetMediaSize(value *string)() {
    m.mediaSize = value
}
// SetMediaType sets the mediaType property value. The default media (such as paper) type to print the document on.
func (m *PrinterDefaults) SetMediaType(value *string)() {
    m.mediaType = value
}
// SetMultipageLayout sets the multipageLayout property value. The default direction to lay out pages when multiple pages are being printed per sheet. Valid values are described in the following table.
func (m *PrinterDefaults) SetMultipageLayout(value *PrintMultipageLayout)() {
    m.multipageLayout = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *PrinterDefaults) SetOdataType(value *string)() {
    m.odataType = value
}
// SetOrientation sets the orientation property value. The default orientation to use when printing the document. Valid values are described in the following table.
func (m *PrinterDefaults) SetOrientation(value *PrintOrientation)() {
    m.orientation = value
}
// SetOutputBin sets the outputBin property value. The default output bin to place completed prints into. See the printer's capabilities for a list of supported output bins.
func (m *PrinterDefaults) SetOutputBin(value *string)() {
    m.outputBin = value
}
// SetPagesPerSheet sets the pagesPerSheet property value. The default number of document pages to print on each sheet.
func (m *PrinterDefaults) SetPagesPerSheet(value *int32)() {
    m.pagesPerSheet = value
}
// SetQuality sets the quality property value. The default quality to use when printing the document. Valid values are described in the following table.
func (m *PrinterDefaults) SetQuality(value *PrintQuality)() {
    m.quality = value
}
// SetScaling sets the scaling property value. Specifies how the printer scales the document data to fit the requested media. Valid values are described in the following table.
func (m *PrinterDefaults) SetScaling(value *PrintScaling)() {
    m.scaling = value
}
