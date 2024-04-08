package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsInformationProtectionApp app for Windows information protection
type WindowsInformationProtectionApp struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // If true, app is denied protection or exemption.
    denied *bool
    // The app's description.
    description *string
    // App display name.
    displayName *string
    // The OdataType property
    odataType *string
    // The product name.
    productName *string
    // The publisher name
    publisherName *string
}
// NewWindowsInformationProtectionApp instantiates a new windowsInformationProtectionApp and sets the default values.
func NewWindowsInformationProtectionApp()(*WindowsInformationProtectionApp) {
    m := &WindowsInformationProtectionApp{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateWindowsInformationProtectionAppFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsInformationProtectionAppFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.windowsInformationProtectionDesktopApp":
                        return NewWindowsInformationProtectionDesktopApp(), nil
                    case "#microsoft.graph.windowsInformationProtectionStoreApp":
                        return NewWindowsInformationProtectionStoreApp(), nil
                }
            }
        }
    }
    return NewWindowsInformationProtectionApp(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WindowsInformationProtectionApp) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDenied gets the denied property value. If true, app is denied protection or exemption.
func (m *WindowsInformationProtectionApp) GetDenied()(*bool) {
    return m.denied
}
// GetDescription gets the description property value. The app's description.
func (m *WindowsInformationProtectionApp) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. App display name.
func (m *WindowsInformationProtectionApp) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsInformationProtectionApp) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["denied"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDenied(val)
        }
        return nil
    }
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
    res["productName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProductName(val)
        }
        return nil
    }
    res["publisherName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublisherName(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *WindowsInformationProtectionApp) GetOdataType()(*string) {
    return m.odataType
}
// GetProductName gets the productName property value. The product name.
func (m *WindowsInformationProtectionApp) GetProductName()(*string) {
    return m.productName
}
// GetPublisherName gets the publisherName property value. The publisher name
func (m *WindowsInformationProtectionApp) GetPublisherName()(*string) {
    return m.publisherName
}
// Serialize serializes information the current object
func (m *WindowsInformationProtectionApp) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("denied", m.GetDenied())
        if err != nil {
            return err
        }
    }
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
        err := writer.WriteStringValue("productName", m.GetProductName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("publisherName", m.GetPublisherName())
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
func (m *WindowsInformationProtectionApp) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDenied sets the denied property value. If true, app is denied protection or exemption.
func (m *WindowsInformationProtectionApp) SetDenied(value *bool)() {
    m.denied = value
}
// SetDescription sets the description property value. The app's description.
func (m *WindowsInformationProtectionApp) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. App display name.
func (m *WindowsInformationProtectionApp) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *WindowsInformationProtectionApp) SetOdataType(value *string)() {
    m.odataType = value
}
// SetProductName sets the productName property value. The product name.
func (m *WindowsInformationProtectionApp) SetProductName(value *string)() {
    m.productName = value
}
// SetPublisherName sets the publisherName property value. The publisher name
func (m *WindowsInformationProtectionApp) SetPublisherName(value *string)() {
    m.publisherName = value
}
