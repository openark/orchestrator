package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ServiceUpdateMessageViewpoint 
type ServiceUpdateMessageViewpoint struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Indicates whether the user archived the message.
    isArchived *bool
    // Indicates whether the user marked the message as favorite.
    isFavorited *bool
    // Indicates whether the user read the message.
    isRead *bool
    // The OdataType property
    odataType *string
}
// NewServiceUpdateMessageViewpoint instantiates a new serviceUpdateMessageViewpoint and sets the default values.
func NewServiceUpdateMessageViewpoint()(*ServiceUpdateMessageViewpoint) {
    m := &ServiceUpdateMessageViewpoint{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateServiceUpdateMessageViewpointFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateServiceUpdateMessageViewpointFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewServiceUpdateMessageViewpoint(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ServiceUpdateMessageViewpoint) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ServiceUpdateMessageViewpoint) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["isArchived"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsArchived(val)
        }
        return nil
    }
    res["isFavorited"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsFavorited(val)
        }
        return nil
    }
    res["isRead"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsRead(val)
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
// GetIsArchived gets the isArchived property value. Indicates whether the user archived the message.
func (m *ServiceUpdateMessageViewpoint) GetIsArchived()(*bool) {
    return m.isArchived
}
// GetIsFavorited gets the isFavorited property value. Indicates whether the user marked the message as favorite.
func (m *ServiceUpdateMessageViewpoint) GetIsFavorited()(*bool) {
    return m.isFavorited
}
// GetIsRead gets the isRead property value. Indicates whether the user read the message.
func (m *ServiceUpdateMessageViewpoint) GetIsRead()(*bool) {
    return m.isRead
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ServiceUpdateMessageViewpoint) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *ServiceUpdateMessageViewpoint) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("isArchived", m.GetIsArchived())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isFavorited", m.GetIsFavorited())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isRead", m.GetIsRead())
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
func (m *ServiceUpdateMessageViewpoint) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetIsArchived sets the isArchived property value. Indicates whether the user archived the message.
func (m *ServiceUpdateMessageViewpoint) SetIsArchived(value *bool)() {
    m.isArchived = value
}
// SetIsFavorited sets the isFavorited property value. Indicates whether the user marked the message as favorite.
func (m *ServiceUpdateMessageViewpoint) SetIsFavorited(value *bool)() {
    m.isFavorited = value
}
// SetIsRead sets the isRead property value. Indicates whether the user read the message.
func (m *ServiceUpdateMessageViewpoint) SetIsRead(value *bool)() {
    m.isRead = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ServiceUpdateMessageViewpoint) SetOdataType(value *string)() {
    m.odataType = value
}
