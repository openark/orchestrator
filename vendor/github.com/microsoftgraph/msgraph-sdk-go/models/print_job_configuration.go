package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PrintJobConfiguration 
type PrintJobConfiguration struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Whether the printer should collate pages wehen printing multiple copies of a multi-page document.
    collate *bool
    // The color mode the printer should use to print the job. Valid values are described in the table below. Read-only.
    colorMode *PrintColorMode
    // The number of copies that should be printed. Read-only.
    copies *int32
    // The resolution to use when printing the job, expressed in dots per inch (DPI). Read-only.
    dpi *int32
    // The duplex mode the printer should use when printing the job. Valid values are described in the table below. Read-only.
    duplexMode *PrintDuplexMode
    // The orientation to use when feeding media into the printer. Valid values are described in the following table. Read-only.
    feedOrientation *PrinterFeedOrientation
    // Finishing processes to use when printing.
    finishings []PrintFinishing
    // The fitPdfToPage property
    fitPdfToPage *bool
    // The input bin (tray) to use when printing. See the printer's capabilities for a list of supported input bins.
    inputBin *string
    // The margin settings to use when printing.
    margin PrintMarginable
    // The media size to use when printing. Supports standard size names for ISO and ANSI media sizes.
    mediaSize *string
    // The mediaType property
    mediaType *string
    // The multipageLayout property
    multipageLayout *PrintMultipageLayout
    // The OdataType property
    odataType *string
    // The orientation property
    orientation *PrintOrientation
    // The outputBin property
    outputBin *string
    // The pageRanges property
    pageRanges []IntegerRangeable
    // The pagesPerSheet property
    pagesPerSheet *int32
    // The quality property
    quality *PrintQuality
    // The scaling property
    scaling *PrintScaling
}
// NewPrintJobConfiguration instantiates a new printJobConfiguration and sets the default values.
func NewPrintJobConfiguration()(*PrintJobConfiguration) {
    m := &PrintJobConfiguration{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePrintJobConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePrintJobConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPrintJobConfiguration(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PrintJobConfiguration) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCollate gets the collate property value. Whether the printer should collate pages wehen printing multiple copies of a multi-page document.
func (m *PrintJobConfiguration) GetCollate()(*bool) {
    return m.collate
}
// GetColorMode gets the colorMode property value. The color mode the printer should use to print the job. Valid values are described in the table below. Read-only.
func (m *PrintJobConfiguration) GetColorMode()(*PrintColorMode) {
    return m.colorMode
}
// GetCopies gets the copies property value. The number of copies that should be printed. Read-only.
func (m *PrintJobConfiguration) GetCopies()(*int32) {
    return m.copies
}
// GetDpi gets the dpi property value. The resolution to use when printing the job, expressed in dots per inch (DPI). Read-only.
func (m *PrintJobConfiguration) GetDpi()(*int32) {
    return m.dpi
}
// GetDuplexMode gets the duplexMode property value. The duplex mode the printer should use when printing the job. Valid values are described in the table below. Read-only.
func (m *PrintJobConfiguration) GetDuplexMode()(*PrintDuplexMode) {
    return m.duplexMode
}
// GetFeedOrientation gets the feedOrientation property value. The orientation to use when feeding media into the printer. Valid values are described in the following table. Read-only.
func (m *PrintJobConfiguration) GetFeedOrientation()(*PrinterFeedOrientation) {
    return m.feedOrientation
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PrintJobConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["collate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCollate(val)
        }
        return nil
    }
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
    res["copies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCopies(val)
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
    res["feedOrientation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePrinterFeedOrientation)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFeedOrientation(val.(*PrinterFeedOrientation))
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
    res["margin"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePrintMarginFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMargin(val.(PrintMarginable))
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
    res["pageRanges"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateIntegerRangeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IntegerRangeable, len(val))
            for i, v := range val {
                res[i] = v.(IntegerRangeable)
            }
            m.SetPageRanges(res)
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
// GetFinishings gets the finishings property value. Finishing processes to use when printing.
func (m *PrintJobConfiguration) GetFinishings()([]PrintFinishing) {
    return m.finishings
}
// GetFitPdfToPage gets the fitPdfToPage property value. The fitPdfToPage property
func (m *PrintJobConfiguration) GetFitPdfToPage()(*bool) {
    return m.fitPdfToPage
}
// GetInputBin gets the inputBin property value. The input bin (tray) to use when printing. See the printer's capabilities for a list of supported input bins.
func (m *PrintJobConfiguration) GetInputBin()(*string) {
    return m.inputBin
}
// GetMargin gets the margin property value. The margin settings to use when printing.
func (m *PrintJobConfiguration) GetMargin()(PrintMarginable) {
    return m.margin
}
// GetMediaSize gets the mediaSize property value. The media size to use when printing. Supports standard size names for ISO and ANSI media sizes.
func (m *PrintJobConfiguration) GetMediaSize()(*string) {
    return m.mediaSize
}
// GetMediaType gets the mediaType property value. The mediaType property
func (m *PrintJobConfiguration) GetMediaType()(*string) {
    return m.mediaType
}
// GetMultipageLayout gets the multipageLayout property value. The multipageLayout property
func (m *PrintJobConfiguration) GetMultipageLayout()(*PrintMultipageLayout) {
    return m.multipageLayout
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *PrintJobConfiguration) GetOdataType()(*string) {
    return m.odataType
}
// GetOrientation gets the orientation property value. The orientation property
func (m *PrintJobConfiguration) GetOrientation()(*PrintOrientation) {
    return m.orientation
}
// GetOutputBin gets the outputBin property value. The outputBin property
func (m *PrintJobConfiguration) GetOutputBin()(*string) {
    return m.outputBin
}
// GetPageRanges gets the pageRanges property value. The pageRanges property
func (m *PrintJobConfiguration) GetPageRanges()([]IntegerRangeable) {
    return m.pageRanges
}
// GetPagesPerSheet gets the pagesPerSheet property value. The pagesPerSheet property
func (m *PrintJobConfiguration) GetPagesPerSheet()(*int32) {
    return m.pagesPerSheet
}
// GetQuality gets the quality property value. The quality property
func (m *PrintJobConfiguration) GetQuality()(*PrintQuality) {
    return m.quality
}
// GetScaling gets the scaling property value. The scaling property
func (m *PrintJobConfiguration) GetScaling()(*PrintScaling) {
    return m.scaling
}
// Serialize serializes information the current object
func (m *PrintJobConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("collate", m.GetCollate())
        if err != nil {
            return err
        }
    }
    if m.GetColorMode() != nil {
        cast := (*m.GetColorMode()).String()
        err := writer.WriteStringValue("colorMode", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("copies", m.GetCopies())
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
    if m.GetFeedOrientation() != nil {
        cast := (*m.GetFeedOrientation()).String()
        err := writer.WriteStringValue("feedOrientation", &cast)
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
        err := writer.WriteObjectValue("margin", m.GetMargin())
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
    if m.GetPageRanges() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPageRanges()))
        for i, v := range m.GetPageRanges() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("pageRanges", cast)
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
func (m *PrintJobConfiguration) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCollate sets the collate property value. Whether the printer should collate pages wehen printing multiple copies of a multi-page document.
func (m *PrintJobConfiguration) SetCollate(value *bool)() {
    m.collate = value
}
// SetColorMode sets the colorMode property value. The color mode the printer should use to print the job. Valid values are described in the table below. Read-only.
func (m *PrintJobConfiguration) SetColorMode(value *PrintColorMode)() {
    m.colorMode = value
}
// SetCopies sets the copies property value. The number of copies that should be printed. Read-only.
func (m *PrintJobConfiguration) SetCopies(value *int32)() {
    m.copies = value
}
// SetDpi sets the dpi property value. The resolution to use when printing the job, expressed in dots per inch (DPI). Read-only.
func (m *PrintJobConfiguration) SetDpi(value *int32)() {
    m.dpi = value
}
// SetDuplexMode sets the duplexMode property value. The duplex mode the printer should use when printing the job. Valid values are described in the table below. Read-only.
func (m *PrintJobConfiguration) SetDuplexMode(value *PrintDuplexMode)() {
    m.duplexMode = value
}
// SetFeedOrientation sets the feedOrientation property value. The orientation to use when feeding media into the printer. Valid values are described in the following table. Read-only.
func (m *PrintJobConfiguration) SetFeedOrientation(value *PrinterFeedOrientation)() {
    m.feedOrientation = value
}
// SetFinishings sets the finishings property value. Finishing processes to use when printing.
func (m *PrintJobConfiguration) SetFinishings(value []PrintFinishing)() {
    m.finishings = value
}
// SetFitPdfToPage sets the fitPdfToPage property value. The fitPdfToPage property
func (m *PrintJobConfiguration) SetFitPdfToPage(value *bool)() {
    m.fitPdfToPage = value
}
// SetInputBin sets the inputBin property value. The input bin (tray) to use when printing. See the printer's capabilities for a list of supported input bins.
func (m *PrintJobConfiguration) SetInputBin(value *string)() {
    m.inputBin = value
}
// SetMargin sets the margin property value. The margin settings to use when printing.
func (m *PrintJobConfiguration) SetMargin(value PrintMarginable)() {
    m.margin = value
}
// SetMediaSize sets the mediaSize property value. The media size to use when printing. Supports standard size names for ISO and ANSI media sizes.
func (m *PrintJobConfiguration) SetMediaSize(value *string)() {
    m.mediaSize = value
}
// SetMediaType sets the mediaType property value. The mediaType property
func (m *PrintJobConfiguration) SetMediaType(value *string)() {
    m.mediaType = value
}
// SetMultipageLayout sets the multipageLayout property value. The multipageLayout property
func (m *PrintJobConfiguration) SetMultipageLayout(value *PrintMultipageLayout)() {
    m.multipageLayout = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *PrintJobConfiguration) SetOdataType(value *string)() {
    m.odataType = value
}
// SetOrientation sets the orientation property value. The orientation property
func (m *PrintJobConfiguration) SetOrientation(value *PrintOrientation)() {
    m.orientation = value
}
// SetOutputBin sets the outputBin property value. The outputBin property
func (m *PrintJobConfiguration) SetOutputBin(value *string)() {
    m.outputBin = value
}
// SetPageRanges sets the pageRanges property value. The pageRanges property
func (m *PrintJobConfiguration) SetPageRanges(value []IntegerRangeable)() {
    m.pageRanges = value
}
// SetPagesPerSheet sets the pagesPerSheet property value. The pagesPerSheet property
func (m *PrintJobConfiguration) SetPagesPerSheet(value *int32)() {
    m.pagesPerSheet = value
}
// SetQuality sets the quality property value. The quality property
func (m *PrintJobConfiguration) SetQuality(value *PrintQuality)() {
    m.quality = value
}
// SetScaling sets the scaling property value. The scaling property
func (m *PrintJobConfiguration) SetScaling(value *PrintScaling)() {
    m.scaling = value
}
