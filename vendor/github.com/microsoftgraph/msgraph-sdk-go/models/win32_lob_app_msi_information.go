package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Win32LobAppMsiInformation contains MSI app properties for a Win32 App.
type Win32LobAppMsiInformation struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The OdataType property
    odataType *string
    // Indicates the package type of an MSI Win32LobApp.
    packageType *Win32LobAppMsiPackageType
    // The MSI product code.
    productCode *string
    // The MSI product name.
    productName *string
    // The MSI product version.
    productVersion *string
    // The MSI publisher.
    publisher *string
    // Whether the MSI app requires the machine to reboot to complete installation.
    requiresReboot *bool
    // The MSI upgrade code.
    upgradeCode *string
}
// NewWin32LobAppMsiInformation instantiates a new win32LobAppMsiInformation and sets the default values.
func NewWin32LobAppMsiInformation()(*Win32LobAppMsiInformation) {
    m := &Win32LobAppMsiInformation{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateWin32LobAppMsiInformationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWin32LobAppMsiInformationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWin32LobAppMsiInformation(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Win32LobAppMsiInformation) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Win32LobAppMsiInformation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["packageType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWin32LobAppMsiPackageType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPackageType(val.(*Win32LobAppMsiPackageType))
        }
        return nil
    }
    res["productCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProductCode(val)
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
    res["productVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProductVersion(val)
        }
        return nil
    }
    res["publisher"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublisher(val)
        }
        return nil
    }
    res["requiresReboot"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequiresReboot(val)
        }
        return nil
    }
    res["upgradeCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpgradeCode(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *Win32LobAppMsiInformation) GetOdataType()(*string) {
    return m.odataType
}
// GetPackageType gets the packageType property value. Indicates the package type of an MSI Win32LobApp.
func (m *Win32LobAppMsiInformation) GetPackageType()(*Win32LobAppMsiPackageType) {
    return m.packageType
}
// GetProductCode gets the productCode property value. The MSI product code.
func (m *Win32LobAppMsiInformation) GetProductCode()(*string) {
    return m.productCode
}
// GetProductName gets the productName property value. The MSI product name.
func (m *Win32LobAppMsiInformation) GetProductName()(*string) {
    return m.productName
}
// GetProductVersion gets the productVersion property value. The MSI product version.
func (m *Win32LobAppMsiInformation) GetProductVersion()(*string) {
    return m.productVersion
}
// GetPublisher gets the publisher property value. The MSI publisher.
func (m *Win32LobAppMsiInformation) GetPublisher()(*string) {
    return m.publisher
}
// GetRequiresReboot gets the requiresReboot property value. Whether the MSI app requires the machine to reboot to complete installation.
func (m *Win32LobAppMsiInformation) GetRequiresReboot()(*bool) {
    return m.requiresReboot
}
// GetUpgradeCode gets the upgradeCode property value. The MSI upgrade code.
func (m *Win32LobAppMsiInformation) GetUpgradeCode()(*string) {
    return m.upgradeCode
}
// Serialize serializes information the current object
func (m *Win32LobAppMsiInformation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    if m.GetPackageType() != nil {
        cast := (*m.GetPackageType()).String()
        err := writer.WriteStringValue("packageType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("productCode", m.GetProductCode())
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
        err := writer.WriteStringValue("productVersion", m.GetProductVersion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("publisher", m.GetPublisher())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("requiresReboot", m.GetRequiresReboot())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("upgradeCode", m.GetUpgradeCode())
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
func (m *Win32LobAppMsiInformation) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *Win32LobAppMsiInformation) SetOdataType(value *string)() {
    m.odataType = value
}
// SetPackageType sets the packageType property value. Indicates the package type of an MSI Win32LobApp.
func (m *Win32LobAppMsiInformation) SetPackageType(value *Win32LobAppMsiPackageType)() {
    m.packageType = value
}
// SetProductCode sets the productCode property value. The MSI product code.
func (m *Win32LobAppMsiInformation) SetProductCode(value *string)() {
    m.productCode = value
}
// SetProductName sets the productName property value. The MSI product name.
func (m *Win32LobAppMsiInformation) SetProductName(value *string)() {
    m.productName = value
}
// SetProductVersion sets the productVersion property value. The MSI product version.
func (m *Win32LobAppMsiInformation) SetProductVersion(value *string)() {
    m.productVersion = value
}
// SetPublisher sets the publisher property value. The MSI publisher.
func (m *Win32LobAppMsiInformation) SetPublisher(value *string)() {
    m.publisher = value
}
// SetRequiresReboot sets the requiresReboot property value. Whether the MSI app requires the machine to reboot to complete installation.
func (m *Win32LobAppMsiInformation) SetRequiresReboot(value *bool)() {
    m.requiresReboot = value
}
// SetUpgradeCode sets the upgradeCode property value. The MSI upgrade code.
func (m *Win32LobAppMsiInformation) SetUpgradeCode(value *string)() {
    m.upgradeCode = value
}
