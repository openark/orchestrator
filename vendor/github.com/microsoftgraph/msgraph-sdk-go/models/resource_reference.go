package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ResourceReference 
type ResourceReference struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The item's unique identifier.
    id *string
    // The OdataType property
    odataType *string
    // A string value that can be used to classify the item, such as 'microsoft.graph.driveItem'
    typeEscaped *string
    // A URL leading to the referenced item.
    webUrl *string
}
// NewResourceReference instantiates a new resourceReference and sets the default values.
func NewResourceReference()(*ResourceReference) {
    m := &ResourceReference{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateResourceReferenceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateResourceReferenceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewResourceReference(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ResourceReference) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ResourceReference) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val)
        }
        return nil
    }
    res["webUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWebUrl(val)
        }
        return nil
    }
    return res
}
// GetId gets the id property value. The item's unique identifier.
func (m *ResourceReference) GetId()(*string) {
    return m.id
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ResourceReference) GetOdataType()(*string) {
    return m.odataType
}
// GetType gets the type property value. A string value that can be used to classify the item, such as 'microsoft.graph.driveItem'
func (m *ResourceReference) GetType()(*string) {
    return m.typeEscaped
}
// GetWebUrl gets the webUrl property value. A URL leading to the referenced item.
func (m *ResourceReference) GetWebUrl()(*string) {
    return m.webUrl
}
// Serialize serializes information the current object
func (m *ResourceReference) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteStringValue("type", m.GetType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("webUrl", m.GetWebUrl())
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
func (m *ResourceReference) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetId sets the id property value. The item's unique identifier.
func (m *ResourceReference) SetId(value *string)() {
    m.id = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ResourceReference) SetOdataType(value *string)() {
    m.odataType = value
}
// SetType sets the type property value. A string value that can be used to classify the item, such as 'microsoft.graph.driveItem'
func (m *ResourceReference) SetType(value *string)() {
    m.typeEscaped = value
}
// SetWebUrl sets the webUrl property value. A URL leading to the referenced item.
func (m *ResourceReference) SetWebUrl(value *string)() {
    m.webUrl = value
}
