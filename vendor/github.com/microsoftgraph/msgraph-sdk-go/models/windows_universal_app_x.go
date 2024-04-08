package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsUniversalAppX 
type WindowsUniversalAppX struct {
    MobileLobApp
    // Contains properties for Windows architecture.
    applicableArchitectures *WindowsArchitecture
    // Contains properties for Windows device type.
    applicableDeviceTypes *WindowsDeviceType
    // The collection of contained apps in the committed mobileAppContent of a windowsUniversalAppX app.
    committedContainedApps []MobileContainedAppable
    // The Identity Name.
    identityName *string
    // The Identity Publisher Hash.
    identityPublisherHash *string
    // The Identity Resource Identifier.
    identityResourceIdentifier *string
    // The identity version.
    identityVersion *string
    // Whether or not the app is a bundle.
    isBundle *bool
    // The minimum operating system required for a Windows mobile app.
    minimumSupportedOperatingSystem WindowsMinimumOperatingSystemable
}
// NewWindowsUniversalAppX instantiates a new WindowsUniversalAppX and sets the default values.
func NewWindowsUniversalAppX()(*WindowsUniversalAppX) {
    m := &WindowsUniversalAppX{
        MobileLobApp: *NewMobileLobApp(),
    }
    odataTypeValue := "#microsoft.graph.windowsUniversalAppX"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateWindowsUniversalAppXFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsUniversalAppXFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsUniversalAppX(), nil
}
// GetApplicableArchitectures gets the applicableArchitectures property value. Contains properties for Windows architecture.
func (m *WindowsUniversalAppX) GetApplicableArchitectures()(*WindowsArchitecture) {
    return m.applicableArchitectures
}
// GetApplicableDeviceTypes gets the applicableDeviceTypes property value. Contains properties for Windows device type.
func (m *WindowsUniversalAppX) GetApplicableDeviceTypes()(*WindowsDeviceType) {
    return m.applicableDeviceTypes
}
// GetCommittedContainedApps gets the committedContainedApps property value. The collection of contained apps in the committed mobileAppContent of a windowsUniversalAppX app.
func (m *WindowsUniversalAppX) GetCommittedContainedApps()([]MobileContainedAppable) {
    return m.committedContainedApps
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsUniversalAppX) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.MobileLobApp.GetFieldDeserializers()
    res["applicableArchitectures"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindowsArchitecture)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicableArchitectures(val.(*WindowsArchitecture))
        }
        return nil
    }
    res["applicableDeviceTypes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindowsDeviceType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicableDeviceTypes(val.(*WindowsDeviceType))
        }
        return nil
    }
    res["committedContainedApps"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMobileContainedAppFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MobileContainedAppable, len(val))
            for i, v := range val {
                res[i] = v.(MobileContainedAppable)
            }
            m.SetCommittedContainedApps(res)
        }
        return nil
    }
    res["identityName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentityName(val)
        }
        return nil
    }
    res["identityPublisherHash"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentityPublisherHash(val)
        }
        return nil
    }
    res["identityResourceIdentifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentityResourceIdentifier(val)
        }
        return nil
    }
    res["identityVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentityVersion(val)
        }
        return nil
    }
    res["isBundle"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsBundle(val)
        }
        return nil
    }
    res["minimumSupportedOperatingSystem"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWindowsMinimumOperatingSystemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumSupportedOperatingSystem(val.(WindowsMinimumOperatingSystemable))
        }
        return nil
    }
    return res
}
// GetIdentityName gets the identityName property value. The Identity Name.
func (m *WindowsUniversalAppX) GetIdentityName()(*string) {
    return m.identityName
}
// GetIdentityPublisherHash gets the identityPublisherHash property value. The Identity Publisher Hash.
func (m *WindowsUniversalAppX) GetIdentityPublisherHash()(*string) {
    return m.identityPublisherHash
}
// GetIdentityResourceIdentifier gets the identityResourceIdentifier property value. The Identity Resource Identifier.
func (m *WindowsUniversalAppX) GetIdentityResourceIdentifier()(*string) {
    return m.identityResourceIdentifier
}
// GetIdentityVersion gets the identityVersion property value. The identity version.
func (m *WindowsUniversalAppX) GetIdentityVersion()(*string) {
    return m.identityVersion
}
// GetIsBundle gets the isBundle property value. Whether or not the app is a bundle.
func (m *WindowsUniversalAppX) GetIsBundle()(*bool) {
    return m.isBundle
}
// GetMinimumSupportedOperatingSystem gets the minimumSupportedOperatingSystem property value. The minimum operating system required for a Windows mobile app.
func (m *WindowsUniversalAppX) GetMinimumSupportedOperatingSystem()(WindowsMinimumOperatingSystemable) {
    return m.minimumSupportedOperatingSystem
}
// Serialize serializes information the current object
func (m *WindowsUniversalAppX) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.MobileLobApp.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetApplicableArchitectures() != nil {
        cast := (*m.GetApplicableArchitectures()).String()
        err = writer.WriteStringValue("applicableArchitectures", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetApplicableDeviceTypes() != nil {
        cast := (*m.GetApplicableDeviceTypes()).String()
        err = writer.WriteStringValue("applicableDeviceTypes", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetCommittedContainedApps() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCommittedContainedApps()))
        for i, v := range m.GetCommittedContainedApps() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("committedContainedApps", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("identityName", m.GetIdentityName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("identityPublisherHash", m.GetIdentityPublisherHash())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("identityResourceIdentifier", m.GetIdentityResourceIdentifier())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("identityVersion", m.GetIdentityVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isBundle", m.GetIsBundle())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("minimumSupportedOperatingSystem", m.GetMinimumSupportedOperatingSystem())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetApplicableArchitectures sets the applicableArchitectures property value. Contains properties for Windows architecture.
func (m *WindowsUniversalAppX) SetApplicableArchitectures(value *WindowsArchitecture)() {
    m.applicableArchitectures = value
}
// SetApplicableDeviceTypes sets the applicableDeviceTypes property value. Contains properties for Windows device type.
func (m *WindowsUniversalAppX) SetApplicableDeviceTypes(value *WindowsDeviceType)() {
    m.applicableDeviceTypes = value
}
// SetCommittedContainedApps sets the committedContainedApps property value. The collection of contained apps in the committed mobileAppContent of a windowsUniversalAppX app.
func (m *WindowsUniversalAppX) SetCommittedContainedApps(value []MobileContainedAppable)() {
    m.committedContainedApps = value
}
// SetIdentityName sets the identityName property value. The Identity Name.
func (m *WindowsUniversalAppX) SetIdentityName(value *string)() {
    m.identityName = value
}
// SetIdentityPublisherHash sets the identityPublisherHash property value. The Identity Publisher Hash.
func (m *WindowsUniversalAppX) SetIdentityPublisherHash(value *string)() {
    m.identityPublisherHash = value
}
// SetIdentityResourceIdentifier sets the identityResourceIdentifier property value. The Identity Resource Identifier.
func (m *WindowsUniversalAppX) SetIdentityResourceIdentifier(value *string)() {
    m.identityResourceIdentifier = value
}
// SetIdentityVersion sets the identityVersion property value. The identity version.
func (m *WindowsUniversalAppX) SetIdentityVersion(value *string)() {
    m.identityVersion = value
}
// SetIsBundle sets the isBundle property value. Whether or not the app is a bundle.
func (m *WindowsUniversalAppX) SetIsBundle(value *bool)() {
    m.isBundle = value
}
// SetMinimumSupportedOperatingSystem sets the minimumSupportedOperatingSystem property value. The minimum operating system required for a Windows mobile app.
func (m *WindowsUniversalAppX) SetMinimumSupportedOperatingSystem(value WindowsMinimumOperatingSystemable)() {
    m.minimumSupportedOperatingSystem = value
}
