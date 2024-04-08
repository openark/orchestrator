package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TextColumn 
type TextColumn struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Whether to allow multiple lines of text.
    allowMultipleLines *bool
    // Whether updates to this column should replace existing text, or append to it.
    appendChangesToExistingText *bool
    // The size of the text box.
    linesForEditing *int32
    // The maximum number of characters for the value.
    maxLength *int32
    // The OdataType property
    odataType *string
    // The type of text being stored. Must be one of plain or richText
    textType *string
}
// NewTextColumn instantiates a new textColumn and sets the default values.
func NewTextColumn()(*TextColumn) {
    m := &TextColumn{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTextColumnFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTextColumnFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTextColumn(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TextColumn) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAllowMultipleLines gets the allowMultipleLines property value. Whether to allow multiple lines of text.
func (m *TextColumn) GetAllowMultipleLines()(*bool) {
    return m.allowMultipleLines
}
// GetAppendChangesToExistingText gets the appendChangesToExistingText property value. Whether updates to this column should replace existing text, or append to it.
func (m *TextColumn) GetAppendChangesToExistingText()(*bool) {
    return m.appendChangesToExistingText
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TextColumn) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["allowMultipleLines"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowMultipleLines(val)
        }
        return nil
    }
    res["appendChangesToExistingText"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppendChangesToExistingText(val)
        }
        return nil
    }
    res["linesForEditing"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLinesForEditing(val)
        }
        return nil
    }
    res["maxLength"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxLength(val)
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
    res["textType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTextType(val)
        }
        return nil
    }
    return res
}
// GetLinesForEditing gets the linesForEditing property value. The size of the text box.
func (m *TextColumn) GetLinesForEditing()(*int32) {
    return m.linesForEditing
}
// GetMaxLength gets the maxLength property value. The maximum number of characters for the value.
func (m *TextColumn) GetMaxLength()(*int32) {
    return m.maxLength
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *TextColumn) GetOdataType()(*string) {
    return m.odataType
}
// GetTextType gets the textType property value. The type of text being stored. Must be one of plain or richText
func (m *TextColumn) GetTextType()(*string) {
    return m.textType
}
// Serialize serializes information the current object
func (m *TextColumn) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("allowMultipleLines", m.GetAllowMultipleLines())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("appendChangesToExistingText", m.GetAppendChangesToExistingText())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("linesForEditing", m.GetLinesForEditing())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("maxLength", m.GetMaxLength())
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
        err := writer.WriteStringValue("textType", m.GetTextType())
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
func (m *TextColumn) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAllowMultipleLines sets the allowMultipleLines property value. Whether to allow multiple lines of text.
func (m *TextColumn) SetAllowMultipleLines(value *bool)() {
    m.allowMultipleLines = value
}
// SetAppendChangesToExistingText sets the appendChangesToExistingText property value. Whether updates to this column should replace existing text, or append to it.
func (m *TextColumn) SetAppendChangesToExistingText(value *bool)() {
    m.appendChangesToExistingText = value
}
// SetLinesForEditing sets the linesForEditing property value. The size of the text box.
func (m *TextColumn) SetLinesForEditing(value *int32)() {
    m.linesForEditing = value
}
// SetMaxLength sets the maxLength property value. The maximum number of characters for the value.
func (m *TextColumn) SetMaxLength(value *int32)() {
    m.maxLength = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *TextColumn) SetOdataType(value *string)() {
    m.odataType = value
}
// SetTextType sets the textType property value. The type of text being stored. Must be one of plain or richText
func (m *TextColumn) SetTextType(value *string)() {
    m.textType = value
}
