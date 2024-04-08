package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ResourceAction set of allowed and not allowed actions for a resource.
type ResourceAction struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Allowed Actions
    allowedResourceActions []string
    // Not Allowed Actions.
    notAllowedResourceActions []string
    // The OdataType property
    odataType *string
}
// NewResourceAction instantiates a new resourceAction and sets the default values.
func NewResourceAction()(*ResourceAction) {
    m := &ResourceAction{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateResourceActionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateResourceActionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewResourceAction(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ResourceAction) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAllowedResourceActions gets the allowedResourceActions property value. Allowed Actions
func (m *ResourceAction) GetAllowedResourceActions()([]string) {
    return m.allowedResourceActions
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ResourceAction) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["allowedResourceActions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetAllowedResourceActions(res)
        }
        return nil
    }
    res["notAllowedResourceActions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetNotAllowedResourceActions(res)
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
// GetNotAllowedResourceActions gets the notAllowedResourceActions property value. Not Allowed Actions.
func (m *ResourceAction) GetNotAllowedResourceActions()([]string) {
    return m.notAllowedResourceActions
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ResourceAction) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *ResourceAction) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAllowedResourceActions() != nil {
        err := writer.WriteCollectionOfStringValues("allowedResourceActions", m.GetAllowedResourceActions())
        if err != nil {
            return err
        }
    }
    if m.GetNotAllowedResourceActions() != nil {
        err := writer.WriteCollectionOfStringValues("notAllowedResourceActions", m.GetNotAllowedResourceActions())
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
func (m *ResourceAction) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAllowedResourceActions sets the allowedResourceActions property value. Allowed Actions
func (m *ResourceAction) SetAllowedResourceActions(value []string)() {
    m.allowedResourceActions = value
}
// SetNotAllowedResourceActions sets the notAllowedResourceActions property value. Not Allowed Actions.
func (m *ResourceAction) SetNotAllowedResourceActions(value []string)() {
    m.notAllowedResourceActions = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ResourceAction) SetOdataType(value *string)() {
    m.odataType = value
}
