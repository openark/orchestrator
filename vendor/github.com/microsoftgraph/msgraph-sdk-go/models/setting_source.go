package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SettingSource 
type SettingSource struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Not yet documented
    displayName *string
    // Not yet documented
    id *string
    // The OdataType property
    odataType *string
    // The sourceType property
    sourceType *SettingSourceType
}
// NewSettingSource instantiates a new settingSource and sets the default values.
func NewSettingSource()(*SettingSource) {
    m := &SettingSource{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSettingSourceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSettingSourceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSettingSource(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SettingSource) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDisplayName gets the displayName property value. Not yet documented
func (m *SettingSource) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SettingSource) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["sourceType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSettingSourceType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSourceType(val.(*SettingSourceType))
        }
        return nil
    }
    return res
}
// GetId gets the id property value. Not yet documented
func (m *SettingSource) GetId()(*string) {
    return m.id
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *SettingSource) GetOdataType()(*string) {
    return m.odataType
}
// GetSourceType gets the sourceType property value. The sourceType property
func (m *SettingSource) GetSourceType()(*SettingSourceType) {
    return m.sourceType
}
// Serialize serializes information the current object
func (m *SettingSource) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
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
    if m.GetSourceType() != nil {
        cast := (*m.GetSourceType()).String()
        err := writer.WriteStringValue("sourceType", &cast)
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
func (m *SettingSource) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDisplayName sets the displayName property value. Not yet documented
func (m *SettingSource) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetId sets the id property value. Not yet documented
func (m *SettingSource) SetId(value *string)() {
    m.id = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *SettingSource) SetOdataType(value *string)() {
    m.odataType = value
}
// SetSourceType sets the sourceType property value. The sourceType property
func (m *SettingSource) SetSourceType(value *SettingSourceType)() {
    m.sourceType = value
}
