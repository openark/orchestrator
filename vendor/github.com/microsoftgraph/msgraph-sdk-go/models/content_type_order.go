package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ContentTypeOrder 
type ContentTypeOrder struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Whether this is the default Content Type
    defaultEscaped *bool
    // The OdataType property
    odataType *string
    // Specifies the position in which the Content Type appears in the selection UI.
    position *int32
}
// NewContentTypeOrder instantiates a new contentTypeOrder and sets the default values.
func NewContentTypeOrder()(*ContentTypeOrder) {
    m := &ContentTypeOrder{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateContentTypeOrderFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateContentTypeOrderFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewContentTypeOrder(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ContentTypeOrder) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDefault gets the default property value. Whether this is the default Content Type
func (m *ContentTypeOrder) GetDefault()(*bool) {
    return m.defaultEscaped
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ContentTypeOrder) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["default"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefault(val)
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
    res["position"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPosition(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ContentTypeOrder) GetOdataType()(*string) {
    return m.odataType
}
// GetPosition gets the position property value. Specifies the position in which the Content Type appears in the selection UI.
func (m *ContentTypeOrder) GetPosition()(*int32) {
    return m.position
}
// Serialize serializes information the current object
func (m *ContentTypeOrder) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("default", m.GetDefault())
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
        err := writer.WriteInt32Value("position", m.GetPosition())
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
func (m *ContentTypeOrder) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDefault sets the default property value. Whether this is the default Content Type
func (m *ContentTypeOrder) SetDefault(value *bool)() {
    m.defaultEscaped = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ContentTypeOrder) SetOdataType(value *string)() {
    m.odataType = value
}
// SetPosition sets the position property value. Specifies the position in which the Content Type appears in the selection UI.
func (m *ContentTypeOrder) SetPosition(value *int32)() {
    m.position = value
}
