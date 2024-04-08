package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ResultTemplateOption 
type ResultTemplateOption struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Indicates whether search display layouts are enabled. If enabled, the user will get the result template to render the search results content in the resultTemplates property of the response. The result template is based on Adaptive Cards. Optional.
    enableResultTemplate *bool
    // The OdataType property
    odataType *string
}
// NewResultTemplateOption instantiates a new resultTemplateOption and sets the default values.
func NewResultTemplateOption()(*ResultTemplateOption) {
    m := &ResultTemplateOption{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateResultTemplateOptionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateResultTemplateOptionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewResultTemplateOption(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ResultTemplateOption) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetEnableResultTemplate gets the enableResultTemplate property value. Indicates whether search display layouts are enabled. If enabled, the user will get the result template to render the search results content in the resultTemplates property of the response. The result template is based on Adaptive Cards. Optional.
func (m *ResultTemplateOption) GetEnableResultTemplate()(*bool) {
    return m.enableResultTemplate
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ResultTemplateOption) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["enableResultTemplate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableResultTemplate(val)
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
func (m *ResultTemplateOption) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *ResultTemplateOption) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("enableResultTemplate", m.GetEnableResultTemplate())
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
func (m *ResultTemplateOption) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetEnableResultTemplate sets the enableResultTemplate property value. Indicates whether search display layouts are enabled. If enabled, the user will get the result template to render the search results content in the resultTemplates property of the response. The result template is based on Adaptive Cards. Optional.
func (m *ResultTemplateOption) SetEnableResultTemplate(value *bool)() {
    m.enableResultTemplate = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ResultTemplateOption) SetOdataType(value *string)() {
    m.odataType = value
}
