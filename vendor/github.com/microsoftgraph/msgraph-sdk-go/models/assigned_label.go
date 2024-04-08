package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AssignedLabel 
type AssignedLabel struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The display name of the label. Read-only.
    displayName *string
    // The unique identifier of the label.
    labelId *string
    // The OdataType property
    odataType *string
}
// NewAssignedLabel instantiates a new assignedLabel and sets the default values.
func NewAssignedLabel()(*AssignedLabel) {
    m := &AssignedLabel{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAssignedLabelFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAssignedLabelFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAssignedLabel(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AssignedLabel) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDisplayName gets the displayName property value. The display name of the label. Read-only.
func (m *AssignedLabel) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AssignedLabel) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["labelId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLabelId(val)
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
// GetLabelId gets the labelId property value. The unique identifier of the label.
func (m *AssignedLabel) GetLabelId()(*string) {
    return m.labelId
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AssignedLabel) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *AssignedLabel) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("labelId", m.GetLabelId())
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
func (m *AssignedLabel) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDisplayName sets the displayName property value. The display name of the label. Read-only.
func (m *AssignedLabel) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetLabelId sets the labelId property value. The unique identifier of the label.
func (m *AssignedLabel) SetLabelId(value *string)() {
    m.labelId = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AssignedLabel) SetOdataType(value *string)() {
    m.odataType = value
}
