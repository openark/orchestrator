package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ArchivedPrintJob 
type ArchivedPrintJob struct {
    // True if the job was acquired by a printer; false otherwise. Read-only.
    acquiredByPrinter *bool
    // The dateTimeOffset when the job was acquired by the printer, if any. Read-only.
    acquiredDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The dateTimeOffset when the job was completed, canceled or aborted. Read-only.
    completionDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The number of copies that were printed. Read-only.
    copiesPrinted *int32
    // The user who created the print job. Read-only.
    createdBy UserIdentityable
    // The dateTimeOffset when the job was created. Read-only.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The archived print job's GUID. Read-only.
    id *string
    // The OdataType property
    odataType *string
    // The printer ID that the job was queued for. Read-only.
    printerId *string
    // The processingState property
    processingState *PrintJobProcessingState
}
// NewArchivedPrintJob instantiates a new archivedPrintJob and sets the default values.
func NewArchivedPrintJob()(*ArchivedPrintJob) {
    m := &ArchivedPrintJob{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateArchivedPrintJobFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateArchivedPrintJobFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewArchivedPrintJob(), nil
}
// GetAcquiredByPrinter gets the acquiredByPrinter property value. True if the job was acquired by a printer; false otherwise. Read-only.
func (m *ArchivedPrintJob) GetAcquiredByPrinter()(*bool) {
    return m.acquiredByPrinter
}
// GetAcquiredDateTime gets the acquiredDateTime property value. The dateTimeOffset when the job was acquired by the printer, if any. Read-only.
func (m *ArchivedPrintJob) GetAcquiredDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.acquiredDateTime
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ArchivedPrintJob) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCompletionDateTime gets the completionDateTime property value. The dateTimeOffset when the job was completed, canceled or aborted. Read-only.
func (m *ArchivedPrintJob) GetCompletionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.completionDateTime
}
// GetCopiesPrinted gets the copiesPrinted property value. The number of copies that were printed. Read-only.
func (m *ArchivedPrintJob) GetCopiesPrinted()(*int32) {
    return m.copiesPrinted
}
// GetCreatedBy gets the createdBy property value. The user who created the print job. Read-only.
func (m *ArchivedPrintJob) GetCreatedBy()(UserIdentityable) {
    return m.createdBy
}
// GetCreatedDateTime gets the createdDateTime property value. The dateTimeOffset when the job was created. Read-only.
func (m *ArchivedPrintJob) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ArchivedPrintJob) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["acquiredByPrinter"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAcquiredByPrinter(val)
        }
        return nil
    }
    res["acquiredDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAcquiredDateTime(val)
        }
        return nil
    }
    res["completionDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompletionDateTime(val)
        }
        return nil
    }
    res["copiesPrinted"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCopiesPrinted(val)
        }
        return nil
    }
    res["createdBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedBy(val.(UserIdentityable))
        }
        return nil
    }
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
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
    res["printerId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrinterId(val)
        }
        return nil
    }
    res["processingState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePrintJobProcessingState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProcessingState(val.(*PrintJobProcessingState))
        }
        return nil
    }
    return res
}
// GetId gets the id property value. The archived print job's GUID. Read-only.
func (m *ArchivedPrintJob) GetId()(*string) {
    return m.id
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ArchivedPrintJob) GetOdataType()(*string) {
    return m.odataType
}
// GetPrinterId gets the printerId property value. The printer ID that the job was queued for. Read-only.
func (m *ArchivedPrintJob) GetPrinterId()(*string) {
    return m.printerId
}
// GetProcessingState gets the processingState property value. The processingState property
func (m *ArchivedPrintJob) GetProcessingState()(*PrintJobProcessingState) {
    return m.processingState
}
// Serialize serializes information the current object
func (m *ArchivedPrintJob) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("acquiredByPrinter", m.GetAcquiredByPrinter())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("acquiredDateTime", m.GetAcquiredDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("completionDateTime", m.GetCompletionDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("copiesPrinted", m.GetCopiesPrinted())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("createdBy", m.GetCreatedBy())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("id", m.GetId())
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
    {
        err := writer.WriteStringValue("printerId", m.GetPrinterId())
        if err != nil {
            return err
        }
    }
    if m.GetProcessingState() != nil {
        cast := (*m.GetProcessingState()).String()
        err := writer.WriteStringValue("processingState", &cast)
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
// SetAcquiredByPrinter sets the acquiredByPrinter property value. True if the job was acquired by a printer; false otherwise. Read-only.
func (m *ArchivedPrintJob) SetAcquiredByPrinter(value *bool)() {
    m.acquiredByPrinter = value
}
// SetAcquiredDateTime sets the acquiredDateTime property value. The dateTimeOffset when the job was acquired by the printer, if any. Read-only.
func (m *ArchivedPrintJob) SetAcquiredDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.acquiredDateTime = value
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ArchivedPrintJob) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCompletionDateTime sets the completionDateTime property value. The dateTimeOffset when the job was completed, canceled or aborted. Read-only.
func (m *ArchivedPrintJob) SetCompletionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.completionDateTime = value
}
// SetCopiesPrinted sets the copiesPrinted property value. The number of copies that were printed. Read-only.
func (m *ArchivedPrintJob) SetCopiesPrinted(value *int32)() {
    m.copiesPrinted = value
}
// SetCreatedBy sets the createdBy property value. The user who created the print job. Read-only.
func (m *ArchivedPrintJob) SetCreatedBy(value UserIdentityable)() {
    m.createdBy = value
}
// SetCreatedDateTime sets the createdDateTime property value. The dateTimeOffset when the job was created. Read-only.
func (m *ArchivedPrintJob) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetId sets the id property value. The archived print job's GUID. Read-only.
func (m *ArchivedPrintJob) SetId(value *string)() {
    m.id = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ArchivedPrintJob) SetOdataType(value *string)() {
    m.odataType = value
}
// SetPrinterId sets the printerId property value. The printer ID that the job was queued for. Read-only.
func (m *ArchivedPrintJob) SetPrinterId(value *string)() {
    m.printerId = value
}
// SetProcessingState sets the processingState property value. The processingState property
func (m *ArchivedPrintJob) SetProcessingState(value *PrintJobProcessingState)() {
    m.processingState = value
}
