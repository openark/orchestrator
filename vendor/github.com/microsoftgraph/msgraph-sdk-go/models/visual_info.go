package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// VisualInfo 
type VisualInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Optional. JSON object used to represent an icon which represents the application used to generate the activity
    attribution ImageInfoable
    // Optional. Background color used to render the activity in the UI - brand color for the application source of the activity. Must be a valid hex color
    backgroundColor *string
    // Optional. Custom piece of data - JSON object used to provide custom content to render the activity in the Windows Shell UI
    content Jsonable
    // Optional. Longer text description of the user's unique activity (example: document name, first sentence, and/or metadata)
    description *string
    // Required. Short text description of the user's unique activity (for example, document name in cases where an activity refers to document creation)
    displayText *string
    // The OdataType property
    odataType *string
}
// NewVisualInfo instantiates a new visualInfo and sets the default values.
func NewVisualInfo()(*VisualInfo) {
    m := &VisualInfo{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateVisualInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateVisualInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVisualInfo(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *VisualInfo) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAttribution gets the attribution property value. Optional. JSON object used to represent an icon which represents the application used to generate the activity
func (m *VisualInfo) GetAttribution()(ImageInfoable) {
    return m.attribution
}
// GetBackgroundColor gets the backgroundColor property value. Optional. Background color used to render the activity in the UI - brand color for the application source of the activity. Must be a valid hex color
func (m *VisualInfo) GetBackgroundColor()(*string) {
    return m.backgroundColor
}
// GetContent gets the content property value. Optional. Custom piece of data - JSON object used to provide custom content to render the activity in the Windows Shell UI
func (m *VisualInfo) GetContent()(Jsonable) {
    return m.content
}
// GetDescription gets the description property value. Optional. Longer text description of the user's unique activity (example: document name, first sentence, and/or metadata)
func (m *VisualInfo) GetDescription()(*string) {
    return m.description
}
// GetDisplayText gets the displayText property value. Required. Short text description of the user's unique activity (for example, document name in cases where an activity refers to document creation)
func (m *VisualInfo) GetDisplayText()(*string) {
    return m.displayText
}
// GetFieldDeserializers the deserialization information for the current model
func (m *VisualInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["attribution"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateImageInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAttribution(val.(ImageInfoable))
        }
        return nil
    }
    res["backgroundColor"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBackgroundColor(val)
        }
        return nil
    }
    res["content"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContent(val.(Jsonable))
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["displayText"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayText(val)
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
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *VisualInfo) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *VisualInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("attribution", m.GetAttribution())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("backgroundColor", m.GetBackgroundColor())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("content", m.GetContent())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayText", m.GetDisplayText())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *VisualInfo) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAttribution sets the attribution property value. Optional. JSON object used to represent an icon which represents the application used to generate the activity
func (m *VisualInfo) SetAttribution(value ImageInfoable)() {
    m.attribution = value
}
// SetBackgroundColor sets the backgroundColor property value. Optional. Background color used to render the activity in the UI - brand color for the application source of the activity. Must be a valid hex color
func (m *VisualInfo) SetBackgroundColor(value *string)() {
    m.backgroundColor = value
}
// SetContent sets the content property value. Optional. Custom piece of data - JSON object used to provide custom content to render the activity in the Windows Shell UI
func (m *VisualInfo) SetContent(value Jsonable)() {
    m.content = value
}
// SetDescription sets the description property value. Optional. Longer text description of the user's unique activity (example: document name, first sentence, and/or metadata)
func (m *VisualInfo) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayText sets the displayText property value. Required. Short text description of the user's unique activity (for example, document name in cases where an activity refers to document creation)
func (m *VisualInfo) SetDisplayText(value *string)() {
    m.displayText = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *VisualInfo) SetOdataType(value *string)() {
    m.odataType = value
}
