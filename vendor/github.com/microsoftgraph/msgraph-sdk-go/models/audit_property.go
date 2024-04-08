package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AuditProperty a class containing the properties for Audit Property.
type AuditProperty struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Display name.
    displayName *string
    // New value.
    newValue *string
    // The OdataType property
    odataType *string
    // Old value.
    oldValue *string
}
// NewAuditProperty instantiates a new auditProperty and sets the default values.
func NewAuditProperty()(*AuditProperty) {
    m := &AuditProperty{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAuditPropertyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAuditPropertyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAuditProperty(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AuditProperty) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDisplayName gets the displayName property value. Display name.
func (m *AuditProperty) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AuditProperty) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["newValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNewValue(val)
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
    res["oldValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOldValue(val)
        }
        return nil
    }
    return res
}
// GetNewValue gets the newValue property value. New value.
func (m *AuditProperty) GetNewValue()(*string) {
    return m.newValue
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AuditProperty) GetOdataType()(*string) {
    return m.odataType
}
// GetOldValue gets the oldValue property value. Old value.
func (m *AuditProperty) GetOldValue()(*string) {
    return m.oldValue
}
// Serialize serializes information the current object
func (m *AuditProperty) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("newValue", m.GetNewValue())
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
        err := writer.WriteStringValue("oldValue", m.GetOldValue())
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
func (m *AuditProperty) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDisplayName sets the displayName property value. Display name.
func (m *AuditProperty) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetNewValue sets the newValue property value. New value.
func (m *AuditProperty) SetNewValue(value *string)() {
    m.newValue = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AuditProperty) SetOdataType(value *string)() {
    m.odataType = value
}
// SetOldValue sets the oldValue property value. Old value.
func (m *AuditProperty) SetOldValue(value *string)() {
    m.oldValue = value
}
