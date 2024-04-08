package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WorkbookSessionInfo 
type WorkbookSessionInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Id of the workbook session.
    id *string
    // The OdataType property
    odataType *string
    // true for persistent session. false for non-persistent session (view mode)
    persistChanges *bool
}
// NewWorkbookSessionInfo instantiates a new workbookSessionInfo and sets the default values.
func NewWorkbookSessionInfo()(*WorkbookSessionInfo) {
    m := &WorkbookSessionInfo{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateWorkbookSessionInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWorkbookSessionInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWorkbookSessionInfo(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WorkbookSessionInfo) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WorkbookSessionInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["persistChanges"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPersistChanges(val)
        }
        return nil
    }
    return res
}
// GetId gets the id property value. Id of the workbook session.
func (m *WorkbookSessionInfo) GetId()(*string) {
    return m.id
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *WorkbookSessionInfo) GetOdataType()(*string) {
    return m.odataType
}
// GetPersistChanges gets the persistChanges property value. true for persistent session. false for non-persistent session (view mode)
func (m *WorkbookSessionInfo) GetPersistChanges()(*bool) {
    return m.persistChanges
}
// Serialize serializes information the current object
func (m *WorkbookSessionInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteBoolValue("persistChanges", m.GetPersistChanges())
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
func (m *WorkbookSessionInfo) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetId sets the id property value. Id of the workbook session.
func (m *WorkbookSessionInfo) SetId(value *string)() {
    m.id = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *WorkbookSessionInfo) SetOdataType(value *string)() {
    m.odataType = value
}
// SetPersistChanges sets the persistChanges property value. true for persistent session. false for non-persistent session (view mode)
func (m *WorkbookSessionInfo) SetPersistChanges(value *bool)() {
    m.persistChanges = value
}
