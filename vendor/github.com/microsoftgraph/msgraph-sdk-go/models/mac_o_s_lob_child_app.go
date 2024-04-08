package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MacOSLobChildApp contains properties of a macOS .app in the package
type MacOSLobChildApp struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The build number of the app.
    buildNumber *string
    // The bundleId of the app.
    bundleId *string
    // The OdataType property
    odataType *string
    // The version number of the app.
    versionNumber *string
}
// NewMacOSLobChildApp instantiates a new macOSLobChildApp and sets the default values.
func NewMacOSLobChildApp()(*MacOSLobChildApp) {
    m := &MacOSLobChildApp{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateMacOSLobChildAppFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMacOSLobChildAppFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMacOSLobChildApp(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MacOSLobChildApp) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBuildNumber gets the buildNumber property value. The build number of the app.
func (m *MacOSLobChildApp) GetBuildNumber()(*string) {
    return m.buildNumber
}
// GetBundleId gets the bundleId property value. The bundleId of the app.
func (m *MacOSLobChildApp) GetBundleId()(*string) {
    return m.bundleId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MacOSLobChildApp) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["buildNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBuildNumber(val)
        }
        return nil
    }
    res["bundleId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBundleId(val)
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
    res["versionNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersionNumber(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *MacOSLobChildApp) GetOdataType()(*string) {
    return m.odataType
}
// GetVersionNumber gets the versionNumber property value. The version number of the app.
func (m *MacOSLobChildApp) GetVersionNumber()(*string) {
    return m.versionNumber
}
// Serialize serializes information the current object
func (m *MacOSLobChildApp) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("buildNumber", m.GetBuildNumber())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("bundleId", m.GetBundleId())
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
        err := writer.WriteStringValue("versionNumber", m.GetVersionNumber())
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
func (m *MacOSLobChildApp) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBuildNumber sets the buildNumber property value. The build number of the app.
func (m *MacOSLobChildApp) SetBuildNumber(value *string)() {
    m.buildNumber = value
}
// SetBundleId sets the bundleId property value. The bundleId of the app.
func (m *MacOSLobChildApp) SetBundleId(value *string)() {
    m.bundleId = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *MacOSLobChildApp) SetOdataType(value *string)() {
    m.odataType = value
}
// SetVersionNumber sets the versionNumber property value. The version number of the app.
func (m *MacOSLobChildApp) SetVersionNumber(value *string)() {
    m.versionNumber = value
}
