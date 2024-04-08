package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MobileAppAssignmentSettings abstract class to contain properties used to assign a mobile app to a group.
type MobileAppAssignmentSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The OdataType property
    odataType *string
}
// NewMobileAppAssignmentSettings instantiates a new mobileAppAssignmentSettings and sets the default values.
func NewMobileAppAssignmentSettings()(*MobileAppAssignmentSettings) {
    m := &MobileAppAssignmentSettings{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateMobileAppAssignmentSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMobileAppAssignmentSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.iosLobAppAssignmentSettings":
                        return NewIosLobAppAssignmentSettings(), nil
                    case "#microsoft.graph.iosStoreAppAssignmentSettings":
                        return NewIosStoreAppAssignmentSettings(), nil
                    case "#microsoft.graph.iosVppAppAssignmentSettings":
                        return NewIosVppAppAssignmentSettings(), nil
                    case "#microsoft.graph.macOsLobAppAssignmentSettings":
                        return NewMacOsLobAppAssignmentSettings(), nil
                    case "#microsoft.graph.microsoftStoreForBusinessAppAssignmentSettings":
                        return NewMicrosoftStoreForBusinessAppAssignmentSettings(), nil
                    case "#microsoft.graph.win32LobAppAssignmentSettings":
                        return NewWin32LobAppAssignmentSettings(), nil
                    case "#microsoft.graph.windowsAppXAppAssignmentSettings":
                        return NewWindowsAppXAppAssignmentSettings(), nil
                    case "#microsoft.graph.windowsUniversalAppXAppAssignmentSettings":
                        return NewWindowsUniversalAppXAppAssignmentSettings(), nil
                }
            }
        }
    }
    return NewMobileAppAssignmentSettings(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MobileAppAssignmentSettings) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MobileAppAssignmentSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
func (m *MobileAppAssignmentSettings) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *MobileAppAssignmentSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *MobileAppAssignmentSettings) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *MobileAppAssignmentSettings) SetOdataType(value *string)() {
    m.odataType = value
}
