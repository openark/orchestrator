package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PrintMargin 
type PrintMargin struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The margin in microns from the bottom edge.
    bottom *int32
    // The margin in microns from the left edge.
    left *int32
    // The OdataType property
    odataType *string
    // The margin in microns from the right edge.
    right *int32
    // The margin in microns from the top edge.
    top *int32
}
// NewPrintMargin instantiates a new printMargin and sets the default values.
func NewPrintMargin()(*PrintMargin) {
    m := &PrintMargin{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePrintMarginFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePrintMarginFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPrintMargin(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PrintMargin) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBottom gets the bottom property value. The margin in microns from the bottom edge.
func (m *PrintMargin) GetBottom()(*int32) {
    return m.bottom
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PrintMargin) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["bottom"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBottom(val)
        }
        return nil
    }
    res["left"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLeft(val)
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
    res["right"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRight(val)
        }
        return nil
    }
    res["top"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTop(val)
        }
        return nil
    }
    return res
}
// GetLeft gets the left property value. The margin in microns from the left edge.
func (m *PrintMargin) GetLeft()(*int32) {
    return m.left
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *PrintMargin) GetOdataType()(*string) {
    return m.odataType
}
// GetRight gets the right property value. The margin in microns from the right edge.
func (m *PrintMargin) GetRight()(*int32) {
    return m.right
}
// GetTop gets the top property value. The margin in microns from the top edge.
func (m *PrintMargin) GetTop()(*int32) {
    return m.top
}
// Serialize serializes information the current object
func (m *PrintMargin) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("bottom", m.GetBottom())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("left", m.GetLeft())
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
        err := writer.WriteInt32Value("right", m.GetRight())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("top", m.GetTop())
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
func (m *PrintMargin) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBottom sets the bottom property value. The margin in microns from the bottom edge.
func (m *PrintMargin) SetBottom(value *int32)() {
    m.bottom = value
}
// SetLeft sets the left property value. The margin in microns from the left edge.
func (m *PrintMargin) SetLeft(value *int32)() {
    m.left = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *PrintMargin) SetOdataType(value *string)() {
    m.odataType = value
}
// SetRight sets the right property value. The margin in microns from the right edge.
func (m *PrintMargin) SetRight(value *int32)() {
    m.right = value
}
// SetTop sets the top property value. The margin in microns from the top edge.
func (m *PrintMargin) SetTop(value *int32)() {
    m.top = value
}
