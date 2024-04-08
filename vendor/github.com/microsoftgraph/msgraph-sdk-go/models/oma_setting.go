package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OmaSetting oMA Settings definition.
type OmaSetting struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Description.
    description *string
    // Display Name.
    displayName *string
    // The OdataType property
    odataType *string
    // OMA.
    omaUri *string
}
// NewOmaSetting instantiates a new omaSetting and sets the default values.
func NewOmaSetting()(*OmaSetting) {
    m := &OmaSetting{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateOmaSettingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOmaSettingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.omaSettingBase64":
                        return NewOmaSettingBase64(), nil
                    case "#microsoft.graph.omaSettingBoolean":
                        return NewOmaSettingBoolean(), nil
                    case "#microsoft.graph.omaSettingDateTime":
                        return NewOmaSettingDateTime(), nil
                    case "#microsoft.graph.omaSettingFloatingPoint":
                        return NewOmaSettingFloatingPoint(), nil
                    case "#microsoft.graph.omaSettingInteger":
                        return NewOmaSettingInteger(), nil
                    case "#microsoft.graph.omaSettingString":
                        return NewOmaSettingString(), nil
                    case "#microsoft.graph.omaSettingStringXml":
                        return NewOmaSettingStringXml(), nil
                }
            }
        }
    }
    return NewOmaSetting(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OmaSetting) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDescription gets the description property value. Description.
func (m *OmaSetting) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. Display Name.
func (m *OmaSetting) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OmaSetting) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
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
    res["omaUri"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOmaUri(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *OmaSetting) GetOdataType()(*string) {
    return m.odataType
}
// GetOmaUri gets the omaUri property value. OMA.
func (m *OmaSetting) GetOmaUri()(*string) {
    return m.omaUri
}
// Serialize serializes information the current object
func (m *OmaSetting) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
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
        err := writer.WriteStringValue("omaUri", m.GetOmaUri())
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
func (m *OmaSetting) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDescription sets the description property value. Description.
func (m *OmaSetting) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. Display Name.
func (m *OmaSetting) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *OmaSetting) SetOdataType(value *string)() {
    m.odataType = value
}
// SetOmaUri sets the omaUri property value. OMA.
func (m *OmaSetting) SetOmaUri(value *string)() {
    m.omaUri = value
}
