package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ManagedAppRegistration the ManagedAppEntity is the base entity type for all other entity types under app management workflow.
type ManagedAppRegistration struct {
    Entity
    // The app package Identifier
    appIdentifier MobileAppIdentifierable
    // App version
    applicationVersion *string
    // Zero or more policys already applied on the registered app when it last synchronized with managment service.
    appliedPolicies []ManagedAppPolicyable
    // Date and time of creation
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Host device name
    deviceName *string
    // App management SDK generated tag, which helps relate apps hosted on the same device. Not guaranteed to relate apps in all conditions.
    deviceTag *string
    // Host device type
    deviceType *string
    // Zero or more reasons an app registration is flagged. E.g. app running on rooted device
    flaggedReasons []ManagedAppFlaggedReason
    // Zero or more policies admin intended for the app as of now.
    intendedPolicies []ManagedAppPolicyable
    // Date and time of last the app synced with management service.
    lastSyncDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // App management SDK version
    managementSdkVersion *string
    // Zero or more long running operations triggered on the app registration.
    operations []ManagedAppOperationable
    // Operating System version
    platformVersion *string
    // The user Id to who this app registration belongs.
    userId *string
    // Version of the entity.
    version *string
}
// NewManagedAppRegistration instantiates a new managedAppRegistration and sets the default values.
func NewManagedAppRegistration()(*ManagedAppRegistration) {
    m := &ManagedAppRegistration{
        Entity: *NewEntity(),
    }
    return m
}
// CreateManagedAppRegistrationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagedAppRegistrationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.androidManagedAppRegistration":
                        return NewAndroidManagedAppRegistration(), nil
                    case "#microsoft.graph.iosManagedAppRegistration":
                        return NewIosManagedAppRegistration(), nil
                }
            }
        }
    }
    return NewManagedAppRegistration(), nil
}
// GetAppIdentifier gets the appIdentifier property value. The app package Identifier
func (m *ManagedAppRegistration) GetAppIdentifier()(MobileAppIdentifierable) {
    return m.appIdentifier
}
// GetApplicationVersion gets the applicationVersion property value. App version
func (m *ManagedAppRegistration) GetApplicationVersion()(*string) {
    return m.applicationVersion
}
// GetAppliedPolicies gets the appliedPolicies property value. Zero or more policys already applied on the registered app when it last synchronized with managment service.
func (m *ManagedAppRegistration) GetAppliedPolicies()([]ManagedAppPolicyable) {
    return m.appliedPolicies
}
// GetCreatedDateTime gets the createdDateTime property value. Date and time of creation
func (m *ManagedAppRegistration) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDeviceName gets the deviceName property value. Host device name
func (m *ManagedAppRegistration) GetDeviceName()(*string) {
    return m.deviceName
}
// GetDeviceTag gets the deviceTag property value. App management SDK generated tag, which helps relate apps hosted on the same device. Not guaranteed to relate apps in all conditions.
func (m *ManagedAppRegistration) GetDeviceTag()(*string) {
    return m.deviceTag
}
// GetDeviceType gets the deviceType property value. Host device type
func (m *ManagedAppRegistration) GetDeviceType()(*string) {
    return m.deviceType
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagedAppRegistration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["appIdentifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMobileAppIdentifierFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppIdentifier(val.(MobileAppIdentifierable))
        }
        return nil
    }
    res["applicationVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationVersion(val)
        }
        return nil
    }
    res["appliedPolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagedAppPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedAppPolicyable, len(val))
            for i, v := range val {
                res[i] = v.(ManagedAppPolicyable)
            }
            m.SetAppliedPolicies(res)
        }
        return nil
    }
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["deviceName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceName(val)
        }
        return nil
    }
    res["deviceTag"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceTag(val)
        }
        return nil
    }
    res["deviceType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceType(val)
        }
        return nil
    }
    res["flaggedReasons"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseManagedAppFlaggedReason)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedAppFlaggedReason, len(val))
            for i, v := range val {
                res[i] = *(v.(*ManagedAppFlaggedReason))
            }
            m.SetFlaggedReasons(res)
        }
        return nil
    }
    res["intendedPolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagedAppPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedAppPolicyable, len(val))
            for i, v := range val {
                res[i] = v.(ManagedAppPolicyable)
            }
            m.SetIntendedPolicies(res)
        }
        return nil
    }
    res["lastSyncDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastSyncDateTime(val)
        }
        return nil
    }
    res["managementSdkVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagementSdkVersion(val)
        }
        return nil
    }
    res["operations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagedAppOperationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedAppOperationable, len(val))
            for i, v := range val {
                res[i] = v.(ManagedAppOperationable)
            }
            m.SetOperations(res)
        }
        return nil
    }
    res["platformVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPlatformVersion(val)
        }
        return nil
    }
    res["userId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserId(val)
        }
        return nil
    }
    res["version"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersion(val)
        }
        return nil
    }
    return res
}
// GetFlaggedReasons gets the flaggedReasons property value. Zero or more reasons an app registration is flagged. E.g. app running on rooted device
func (m *ManagedAppRegistration) GetFlaggedReasons()([]ManagedAppFlaggedReason) {
    return m.flaggedReasons
}
// GetIntendedPolicies gets the intendedPolicies property value. Zero or more policies admin intended for the app as of now.
func (m *ManagedAppRegistration) GetIntendedPolicies()([]ManagedAppPolicyable) {
    return m.intendedPolicies
}
// GetLastSyncDateTime gets the lastSyncDateTime property value. Date and time of last the app synced with management service.
func (m *ManagedAppRegistration) GetLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastSyncDateTime
}
// GetManagementSdkVersion gets the managementSdkVersion property value. App management SDK version
func (m *ManagedAppRegistration) GetManagementSdkVersion()(*string) {
    return m.managementSdkVersion
}
// GetOperations gets the operations property value. Zero or more long running operations triggered on the app registration.
func (m *ManagedAppRegistration) GetOperations()([]ManagedAppOperationable) {
    return m.operations
}
// GetPlatformVersion gets the platformVersion property value. Operating System version
func (m *ManagedAppRegistration) GetPlatformVersion()(*string) {
    return m.platformVersion
}
// GetUserId gets the userId property value. The user Id to who this app registration belongs.
func (m *ManagedAppRegistration) GetUserId()(*string) {
    return m.userId
}
// GetVersion gets the version property value. Version of the entity.
func (m *ManagedAppRegistration) GetVersion()(*string) {
    return m.version
}
// Serialize serializes information the current object
func (m *ManagedAppRegistration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("appIdentifier", m.GetAppIdentifier())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("applicationVersion", m.GetApplicationVersion())
        if err != nil {
            return err
        }
    }
    if m.GetAppliedPolicies() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAppliedPolicies()))
        for i, v := range m.GetAppliedPolicies() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("appliedPolicies", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceName", m.GetDeviceName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceTag", m.GetDeviceTag())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceType", m.GetDeviceType())
        if err != nil {
            return err
        }
    }
    if m.GetFlaggedReasons() != nil {
        err = writer.WriteCollectionOfStringValues("flaggedReasons", SerializeManagedAppFlaggedReason(m.GetFlaggedReasons()))
        if err != nil {
            return err
        }
    }
    if m.GetIntendedPolicies() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetIntendedPolicies()))
        for i, v := range m.GetIntendedPolicies() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("intendedPolicies", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastSyncDateTime", m.GetLastSyncDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("managementSdkVersion", m.GetManagementSdkVersion())
        if err != nil {
            return err
        }
    }
    if m.GetOperations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetOperations()))
        for i, v := range m.GetOperations() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("operations", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("platformVersion", m.GetPlatformVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userId", m.GetUserId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("version", m.GetVersion())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAppIdentifier sets the appIdentifier property value. The app package Identifier
func (m *ManagedAppRegistration) SetAppIdentifier(value MobileAppIdentifierable)() {
    m.appIdentifier = value
}
// SetApplicationVersion sets the applicationVersion property value. App version
func (m *ManagedAppRegistration) SetApplicationVersion(value *string)() {
    m.applicationVersion = value
}
// SetAppliedPolicies sets the appliedPolicies property value. Zero or more policys already applied on the registered app when it last synchronized with managment service.
func (m *ManagedAppRegistration) SetAppliedPolicies(value []ManagedAppPolicyable)() {
    m.appliedPolicies = value
}
// SetCreatedDateTime sets the createdDateTime property value. Date and time of creation
func (m *ManagedAppRegistration) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDeviceName sets the deviceName property value. Host device name
func (m *ManagedAppRegistration) SetDeviceName(value *string)() {
    m.deviceName = value
}
// SetDeviceTag sets the deviceTag property value. App management SDK generated tag, which helps relate apps hosted on the same device. Not guaranteed to relate apps in all conditions.
func (m *ManagedAppRegistration) SetDeviceTag(value *string)() {
    m.deviceTag = value
}
// SetDeviceType sets the deviceType property value. Host device type
func (m *ManagedAppRegistration) SetDeviceType(value *string)() {
    m.deviceType = value
}
// SetFlaggedReasons sets the flaggedReasons property value. Zero or more reasons an app registration is flagged. E.g. app running on rooted device
func (m *ManagedAppRegistration) SetFlaggedReasons(value []ManagedAppFlaggedReason)() {
    m.flaggedReasons = value
}
// SetIntendedPolicies sets the intendedPolicies property value. Zero or more policies admin intended for the app as of now.
func (m *ManagedAppRegistration) SetIntendedPolicies(value []ManagedAppPolicyable)() {
    m.intendedPolicies = value
}
// SetLastSyncDateTime sets the lastSyncDateTime property value. Date and time of last the app synced with management service.
func (m *ManagedAppRegistration) SetLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastSyncDateTime = value
}
// SetManagementSdkVersion sets the managementSdkVersion property value. App management SDK version
func (m *ManagedAppRegistration) SetManagementSdkVersion(value *string)() {
    m.managementSdkVersion = value
}
// SetOperations sets the operations property value. Zero or more long running operations triggered on the app registration.
func (m *ManagedAppRegistration) SetOperations(value []ManagedAppOperationable)() {
    m.operations = value
}
// SetPlatformVersion sets the platformVersion property value. Operating System version
func (m *ManagedAppRegistration) SetPlatformVersion(value *string)() {
    m.platformVersion = value
}
// SetUserId sets the userId property value. The user Id to who this app registration belongs.
func (m *ManagedAppRegistration) SetUserId(value *string)() {
    m.userId = value
}
// SetVersion sets the version property value. Version of the entity.
func (m *ManagedAppRegistration) SetVersion(value *string)() {
    m.version = value
}
