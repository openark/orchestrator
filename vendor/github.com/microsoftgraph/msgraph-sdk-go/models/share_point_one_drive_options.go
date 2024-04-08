package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SharePointOneDriveOptions 
type SharePointOneDriveOptions struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The includeContent property
    includeContent *SearchContent
    // The OdataType property
    odataType *string
}
// NewSharePointOneDriveOptions instantiates a new sharePointOneDriveOptions and sets the default values.
func NewSharePointOneDriveOptions()(*SharePointOneDriveOptions) {
    m := &SharePointOneDriveOptions{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSharePointOneDriveOptionsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSharePointOneDriveOptionsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSharePointOneDriveOptions(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SharePointOneDriveOptions) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SharePointOneDriveOptions) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["includeContent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSearchContent)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIncludeContent(val.(*SearchContent))
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
// GetIncludeContent gets the includeContent property value. The includeContent property
func (m *SharePointOneDriveOptions) GetIncludeContent()(*SearchContent) {
    return m.includeContent
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *SharePointOneDriveOptions) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *SharePointOneDriveOptions) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetIncludeContent() != nil {
        cast := (*m.GetIncludeContent()).String()
        err := writer.WriteStringValue("includeContent", &cast)
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
func (m *SharePointOneDriveOptions) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetIncludeContent sets the includeContent property value. The includeContent property
func (m *SharePointOneDriveOptions) SetIncludeContent(value *SearchContent)() {
    m.includeContent = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *SharePointOneDriveOptions) SetOdataType(value *string)() {
    m.odataType = value
}
