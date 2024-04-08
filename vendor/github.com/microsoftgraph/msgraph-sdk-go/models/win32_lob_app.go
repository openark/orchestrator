package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Win32LobApp 
type Win32LobApp struct {
    MobileLobApp
    // Contains properties for Windows architecture.
    applicableArchitectures *WindowsArchitecture
    // The command line to install this app
    installCommandLine *string
    // The install experience for this app.
    installExperience Win32LobAppInstallExperienceable
    // The value for the minimum CPU speed which is required to install this app.
    minimumCpuSpeedInMHz *int32
    // The value for the minimum free disk space which is required to install this app.
    minimumFreeDiskSpaceInMB *int32
    // The value for the minimum physical memory which is required to install this app.
    minimumMemoryInMB *int32
    // The value for the minimum number of processors which is required to install this app.
    minimumNumberOfProcessors *int32
    // The value for the minimum supported windows release.
    minimumSupportedWindowsRelease *string
    // The MSI details if this Win32 app is an MSI app.
    msiInformation Win32LobAppMsiInformationable
    // The return codes for post installation behavior.
    returnCodes []Win32LobAppReturnCodeable
    // The detection and requirement rules for this app.
    rules []Win32LobAppRuleable
    // The relative path of the setup file in the encrypted Win32LobApp package.
    setupFilePath *string
    // The command line to uninstall this app
    uninstallCommandLine *string
}
// NewWin32LobApp instantiates a new Win32LobApp and sets the default values.
func NewWin32LobApp()(*Win32LobApp) {
    m := &Win32LobApp{
        MobileLobApp: *NewMobileLobApp(),
    }
    odataTypeValue := "#microsoft.graph.win32LobApp"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateWin32LobAppFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWin32LobAppFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWin32LobApp(), nil
}
// GetApplicableArchitectures gets the applicableArchitectures property value. Contains properties for Windows architecture.
func (m *Win32LobApp) GetApplicableArchitectures()(*WindowsArchitecture) {
    return m.applicableArchitectures
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Win32LobApp) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["installCommandLine"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInstallCommandLine(val)
        }
        return nil
    }
    res["installExperience"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWin32LobAppInstallExperienceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInstallExperience(val.(Win32LobAppInstallExperienceable))
        }
        return nil
    }
    res["minimumCpuSpeedInMHz"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumCpuSpeedInMHz(val)
        }
        return nil
    }
    res["minimumFreeDiskSpaceInMB"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumFreeDiskSpaceInMB(val)
        }
        return nil
    }
    res["minimumMemoryInMB"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumMemoryInMB(val)
        }
        return nil
    }
    res["minimumNumberOfProcessors"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumNumberOfProcessors(val)
        }
        return nil
    }
    res["minimumSupportedWindowsRelease"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumSupportedWindowsRelease(val)
        }
        return nil
    }
    res["msiInformation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWin32LobAppMsiInformationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMsiInformation(val.(Win32LobAppMsiInformationable))
        }
        return nil
    }
    res["returnCodes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWin32LobAppReturnCodeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Win32LobAppReturnCodeable, len(val))
            for i, v := range val {
                res[i] = v.(Win32LobAppReturnCodeable)
            }
            m.SetReturnCodes(res)
        }
        return nil
    }
    res["rules"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWin32LobAppRuleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Win32LobAppRuleable, len(val))
            for i, v := range val {
                res[i] = v.(Win32LobAppRuleable)
            }
            m.SetRules(res)
        }
        return nil
    }
    res["setupFilePath"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSetupFilePath(val)
        }
        return nil
    }
    res["uninstallCommandLine"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUninstallCommandLine(val)
        }
        return nil
    }
    return res
}
// GetInstallCommandLine gets the installCommandLine property value. The command line to install this app
func (m *Win32LobApp) GetInstallCommandLine()(*string) {
    return m.installCommandLine
}
// GetInstallExperience gets the installExperience property value. The install experience for this app.
func (m *Win32LobApp) GetInstallExperience()(Win32LobAppInstallExperienceable) {
    return m.installExperience
}
// GetMinimumCpuSpeedInMHz gets the minimumCpuSpeedInMHz property value. The value for the minimum CPU speed which is required to install this app.
func (m *Win32LobApp) GetMinimumCpuSpeedInMHz()(*int32) {
    return m.minimumCpuSpeedInMHz
}
// GetMinimumFreeDiskSpaceInMB gets the minimumFreeDiskSpaceInMB property value. The value for the minimum free disk space which is required to install this app.
func (m *Win32LobApp) GetMinimumFreeDiskSpaceInMB()(*int32) {
    return m.minimumFreeDiskSpaceInMB
}
// GetMinimumMemoryInMB gets the minimumMemoryInMB property value. The value for the minimum physical memory which is required to install this app.
func (m *Win32LobApp) GetMinimumMemoryInMB()(*int32) {
    return m.minimumMemoryInMB
}
// GetMinimumNumberOfProcessors gets the minimumNumberOfProcessors property value. The value for the minimum number of processors which is required to install this app.
func (m *Win32LobApp) GetMinimumNumberOfProcessors()(*int32) {
    return m.minimumNumberOfProcessors
}
// GetMinimumSupportedWindowsRelease gets the minimumSupportedWindowsRelease property value. The value for the minimum supported windows release.
func (m *Win32LobApp) GetMinimumSupportedWindowsRelease()(*string) {
    return m.minimumSupportedWindowsRelease
}
// GetMsiInformation gets the msiInformation property value. The MSI details if this Win32 app is an MSI app.
func (m *Win32LobApp) GetMsiInformation()(Win32LobAppMsiInformationable) {
    return m.msiInformation
}
// GetReturnCodes gets the returnCodes property value. The return codes for post installation behavior.
func (m *Win32LobApp) GetReturnCodes()([]Win32LobAppReturnCodeable) {
    return m.returnCodes
}
// GetRules gets the rules property value. The detection and requirement rules for this app.
func (m *Win32LobApp) GetRules()([]Win32LobAppRuleable) {
    return m.rules
}
// GetSetupFilePath gets the setupFilePath property value. The relative path of the setup file in the encrypted Win32LobApp package.
func (m *Win32LobApp) GetSetupFilePath()(*string) {
    return m.setupFilePath
}
// GetUninstallCommandLine gets the uninstallCommandLine property value. The command line to uninstall this app
func (m *Win32LobApp) GetUninstallCommandLine()(*string) {
    return m.uninstallCommandLine
}
// Serialize serializes information the current object
func (m *Win32LobApp) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    {
        err = writer.WriteStringValue("installCommandLine", m.GetInstallCommandLine())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("installExperience", m.GetInstallExperience())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("minimumCpuSpeedInMHz", m.GetMinimumCpuSpeedInMHz())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("minimumFreeDiskSpaceInMB", m.GetMinimumFreeDiskSpaceInMB())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("minimumMemoryInMB", m.GetMinimumMemoryInMB())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("minimumNumberOfProcessors", m.GetMinimumNumberOfProcessors())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("minimumSupportedWindowsRelease", m.GetMinimumSupportedWindowsRelease())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("msiInformation", m.GetMsiInformation())
        if err != nil {
            return err
        }
    }
    if m.GetReturnCodes() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetReturnCodes()))
        for i, v := range m.GetReturnCodes() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("returnCodes", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRules() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRules()))
        for i, v := range m.GetRules() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("rules", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("setupFilePath", m.GetSetupFilePath())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("uninstallCommandLine", m.GetUninstallCommandLine())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetApplicableArchitectures sets the applicableArchitectures property value. Contains properties for Windows architecture.
func (m *Win32LobApp) SetApplicableArchitectures(value *WindowsArchitecture)() {
    m.applicableArchitectures = value
}
// SetInstallCommandLine sets the installCommandLine property value. The command line to install this app
func (m *Win32LobApp) SetInstallCommandLine(value *string)() {
    m.installCommandLine = value
}
// SetInstallExperience sets the installExperience property value. The install experience for this app.
func (m *Win32LobApp) SetInstallExperience(value Win32LobAppInstallExperienceable)() {
    m.installExperience = value
}
// SetMinimumCpuSpeedInMHz sets the minimumCpuSpeedInMHz property value. The value for the minimum CPU speed which is required to install this app.
func (m *Win32LobApp) SetMinimumCpuSpeedInMHz(value *int32)() {
    m.minimumCpuSpeedInMHz = value
}
// SetMinimumFreeDiskSpaceInMB sets the minimumFreeDiskSpaceInMB property value. The value for the minimum free disk space which is required to install this app.
func (m *Win32LobApp) SetMinimumFreeDiskSpaceInMB(value *int32)() {
    m.minimumFreeDiskSpaceInMB = value
}
// SetMinimumMemoryInMB sets the minimumMemoryInMB property value. The value for the minimum physical memory which is required to install this app.
func (m *Win32LobApp) SetMinimumMemoryInMB(value *int32)() {
    m.minimumMemoryInMB = value
}
// SetMinimumNumberOfProcessors sets the minimumNumberOfProcessors property value. The value for the minimum number of processors which is required to install this app.
func (m *Win32LobApp) SetMinimumNumberOfProcessors(value *int32)() {
    m.minimumNumberOfProcessors = value
}
// SetMinimumSupportedWindowsRelease sets the minimumSupportedWindowsRelease property value. The value for the minimum supported windows release.
func (m *Win32LobApp) SetMinimumSupportedWindowsRelease(value *string)() {
    m.minimumSupportedWindowsRelease = value
}
// SetMsiInformation sets the msiInformation property value. The MSI details if this Win32 app is an MSI app.
func (m *Win32LobApp) SetMsiInformation(value Win32LobAppMsiInformationable)() {
    m.msiInformation = value
}
// SetReturnCodes sets the returnCodes property value. The return codes for post installation behavior.
func (m *Win32LobApp) SetReturnCodes(value []Win32LobAppReturnCodeable)() {
    m.returnCodes = value
}
// SetRules sets the rules property value. The detection and requirement rules for this app.
func (m *Win32LobApp) SetRules(value []Win32LobAppRuleable)() {
    m.rules = value
}
// SetSetupFilePath sets the setupFilePath property value. The relative path of the setup file in the encrypted Win32LobApp package.
func (m *Win32LobApp) SetSetupFilePath(value *string)() {
    m.setupFilePath = value
}
// SetUninstallCommandLine sets the uninstallCommandLine property value. The command line to uninstall this app
func (m *Win32LobApp) SetUninstallCommandLine(value *string)() {
    m.uninstallCommandLine = value
}
