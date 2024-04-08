package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Device 
type Device struct {
    DirectoryObject
    // true if the account is enabled; otherwise, false. Required. Default is true.  Supports $filter (eq, ne, not, in). Only callers in Global Administrator and Cloud Device Administrator roles can set this property.
    accountEnabled *bool
    // For internal use only. Not nullable. Supports $filter (eq, not, ge, le).
    alternativeSecurityIds []AlternativeSecurityIdable
    // The timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only. Supports $filter (eq, ne, not, ge, le, and eq on null values) and $orderBy.
    approximateLastSignInDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The timestamp when the device is no longer deemed compliant. The timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
    complianceExpirationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Unique identifier set by Azure Device Registration Service at the time of registration. This is an alternate key that can be used to reference the device object. Supports $filter (eq, ne, not, startsWith).
    deviceId *string
    // For internal use only. Set to null.
    deviceMetadata *string
    // For internal use only.
    deviceVersion *int32
    // The display name for the device. Required. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values), $search, and $orderBy.
    displayName *string
    // The collection of open extensions defined for the device. Read-only. Nullable.
    extensions []Extensionable
    // true if the device complies with Mobile Device Management (MDM) policies; otherwise, false. Read-only. This can only be updated by Intune for any device OS type or by an approved MDM app for Windows OS devices. Supports $filter (eq, ne, not).
    isCompliant *bool
    // true if the device is managed by a Mobile Device Management (MDM) app; otherwise, false. This can only be updated by Intune for any device OS type or by an approved MDM app for Windows OS devices. Supports $filter (eq, ne, not).
    isManaged *bool
    // Application identifier used to register device into MDM. Read-only. Supports $filter (eq, ne, not, startsWith).
    mdmAppId *string
    // Groups and administrative units that this device is a member of. Read-only. Nullable. Supports $expand.
    memberOf []DirectoryObjectable
    // The last time at which the object was synced with the on-premises directory. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z Read-only. Supports $filter (eq, ne, not, ge, le, in).
    onPremisesLastSyncDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // true if this object is synced from an on-premises directory; false if this object was originally synced from an on-premises directory but is no longer synced; null if this object has never been synced from an on-premises directory (default). Read-only. Supports $filter (eq, ne, not, in, and eq on null values).
    onPremisesSyncEnabled *bool
    // The type of operating system on the device. Required. Supports $filter (eq, ne, not, ge, le, startsWith, and eq on null values).
    operatingSystem *string
    // The version of the operating system on the device. Required. Supports $filter (eq, ne, not, ge, le, startsWith, and eq on null values).
    operatingSystemVersion *string
    // For internal use only. Not nullable. Supports $filter (eq, not, ge, le, startsWith,/$count eq 0, /$count ne 0).
    physicalIds []string
    // The profile type of the device. Possible values: RegisteredDevice (default), SecureVM, Printer, Shared, IoT.
    profileType *string
    // The user that cloud joined the device or registered their personal device. The registered owner is set at the time of registration. Currently, there can be only one owner. Read-only. Nullable. Supports $expand.
    registeredOwners []DirectoryObjectable
    // Collection of registered users of the device. For cloud joined devices and registered personal devices, registered users are set to the same value as registered owners at the time of registration. Read-only. Nullable. Supports $expand.
    registeredUsers []DirectoryObjectable
    // List of labels applied to the device by the system. Supports $filter (/$count eq 0, /$count ne 0).
    systemLabels []string
    // Groups and administrative units that the device is a member of. This operation is transitive. Supports $expand.
    transitiveMemberOf []DirectoryObjectable
    // Type of trust for the joined device. Read-only. Possible values:  Workplace (indicates bring your own personal devices), AzureAd (Cloud only joined devices), ServerAd (on-premises domain joined devices joined to Azure AD). For more details, see Introduction to device management in Azure Active Directory
    trustType *string
}
// NewDevice instantiates a new Device and sets the default values.
func NewDevice()(*Device) {
    m := &Device{
        DirectoryObject: *NewDirectoryObject(),
    }
    odataTypeValue := "#microsoft.graph.device"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateDeviceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDevice(), nil
}
// GetAccountEnabled gets the accountEnabled property value. true if the account is enabled; otherwise, false. Required. Default is true.  Supports $filter (eq, ne, not, in). Only callers in Global Administrator and Cloud Device Administrator roles can set this property.
func (m *Device) GetAccountEnabled()(*bool) {
    return m.accountEnabled
}
// GetAlternativeSecurityIds gets the alternativeSecurityIds property value. For internal use only. Not nullable. Supports $filter (eq, not, ge, le).
func (m *Device) GetAlternativeSecurityIds()([]AlternativeSecurityIdable) {
    return m.alternativeSecurityIds
}
// GetApproximateLastSignInDateTime gets the approximateLastSignInDateTime property value. The timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only. Supports $filter (eq, ne, not, ge, le, and eq on null values) and $orderBy.
func (m *Device) GetApproximateLastSignInDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.approximateLastSignInDateTime
}
// GetComplianceExpirationDateTime gets the complianceExpirationDateTime property value. The timestamp when the device is no longer deemed compliant. The timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *Device) GetComplianceExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.complianceExpirationDateTime
}
// GetDeviceId gets the deviceId property value. Unique identifier set by Azure Device Registration Service at the time of registration. This is an alternate key that can be used to reference the device object. Supports $filter (eq, ne, not, startsWith).
func (m *Device) GetDeviceId()(*string) {
    return m.deviceId
}
// GetDeviceMetadata gets the deviceMetadata property value. For internal use only. Set to null.
func (m *Device) GetDeviceMetadata()(*string) {
    return m.deviceMetadata
}
// GetDeviceVersion gets the deviceVersion property value. For internal use only.
func (m *Device) GetDeviceVersion()(*int32) {
    return m.deviceVersion
}
// GetDisplayName gets the displayName property value. The display name for the device. Required. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values), $search, and $orderBy.
func (m *Device) GetDisplayName()(*string) {
    return m.displayName
}
// GetExtensions gets the extensions property value. The collection of open extensions defined for the device. Read-only. Nullable.
func (m *Device) GetExtensions()([]Extensionable) {
    return m.extensions
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Device) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DirectoryObject.GetFieldDeserializers()
    res["accountEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccountEnabled(val)
        }
        return nil
    }
    res["alternativeSecurityIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAlternativeSecurityIdFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AlternativeSecurityIdable, len(val))
            for i, v := range val {
                res[i] = v.(AlternativeSecurityIdable)
            }
            m.SetAlternativeSecurityIds(res)
        }
        return nil
    }
    res["approximateLastSignInDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApproximateLastSignInDateTime(val)
        }
        return nil
    }
    res["complianceExpirationDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComplianceExpirationDateTime(val)
        }
        return nil
    }
    res["deviceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceId(val)
        }
        return nil
    }
    res["deviceMetadata"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceMetadata(val)
        }
        return nil
    }
    res["deviceVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceVersion(val)
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
    res["extensions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateExtensionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Extensionable, len(val))
            for i, v := range val {
                res[i] = v.(Extensionable)
            }
            m.SetExtensions(res)
        }
        return nil
    }
    res["isCompliant"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsCompliant(val)
        }
        return nil
    }
    res["isManaged"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsManaged(val)
        }
        return nil
    }
    res["mdmAppId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMdmAppId(val)
        }
        return nil
    }
    res["memberOf"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDirectoryObjectFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DirectoryObjectable, len(val))
            for i, v := range val {
                res[i] = v.(DirectoryObjectable)
            }
            m.SetMemberOf(res)
        }
        return nil
    }
    res["onPremisesLastSyncDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnPremisesLastSyncDateTime(val)
        }
        return nil
    }
    res["onPremisesSyncEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnPremisesSyncEnabled(val)
        }
        return nil
    }
    res["operatingSystem"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOperatingSystem(val)
        }
        return nil
    }
    res["operatingSystemVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOperatingSystemVersion(val)
        }
        return nil
    }
    res["physicalIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetPhysicalIds(res)
        }
        return nil
    }
    res["profileType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProfileType(val)
        }
        return nil
    }
    res["registeredOwners"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDirectoryObjectFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DirectoryObjectable, len(val))
            for i, v := range val {
                res[i] = v.(DirectoryObjectable)
            }
            m.SetRegisteredOwners(res)
        }
        return nil
    }
    res["registeredUsers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDirectoryObjectFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DirectoryObjectable, len(val))
            for i, v := range val {
                res[i] = v.(DirectoryObjectable)
            }
            m.SetRegisteredUsers(res)
        }
        return nil
    }
    res["systemLabels"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetSystemLabels(res)
        }
        return nil
    }
    res["transitiveMemberOf"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDirectoryObjectFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DirectoryObjectable, len(val))
            for i, v := range val {
                res[i] = v.(DirectoryObjectable)
            }
            m.SetTransitiveMemberOf(res)
        }
        return nil
    }
    res["trustType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTrustType(val)
        }
        return nil
    }
    return res
}
// GetIsCompliant gets the isCompliant property value. true if the device complies with Mobile Device Management (MDM) policies; otherwise, false. Read-only. This can only be updated by Intune for any device OS type or by an approved MDM app for Windows OS devices. Supports $filter (eq, ne, not).
func (m *Device) GetIsCompliant()(*bool) {
    return m.isCompliant
}
// GetIsManaged gets the isManaged property value. true if the device is managed by a Mobile Device Management (MDM) app; otherwise, false. This can only be updated by Intune for any device OS type or by an approved MDM app for Windows OS devices. Supports $filter (eq, ne, not).
func (m *Device) GetIsManaged()(*bool) {
    return m.isManaged
}
// GetMdmAppId gets the mdmAppId property value. Application identifier used to register device into MDM. Read-only. Supports $filter (eq, ne, not, startsWith).
func (m *Device) GetMdmAppId()(*string) {
    return m.mdmAppId
}
// GetMemberOf gets the memberOf property value. Groups and administrative units that this device is a member of. Read-only. Nullable. Supports $expand.
func (m *Device) GetMemberOf()([]DirectoryObjectable) {
    return m.memberOf
}
// GetOnPremisesLastSyncDateTime gets the onPremisesLastSyncDateTime property value. The last time at which the object was synced with the on-premises directory. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z Read-only. Supports $filter (eq, ne, not, ge, le, in).
func (m *Device) GetOnPremisesLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.onPremisesLastSyncDateTime
}
// GetOnPremisesSyncEnabled gets the onPremisesSyncEnabled property value. true if this object is synced from an on-premises directory; false if this object was originally synced from an on-premises directory but is no longer synced; null if this object has never been synced from an on-premises directory (default). Read-only. Supports $filter (eq, ne, not, in, and eq on null values).
func (m *Device) GetOnPremisesSyncEnabled()(*bool) {
    return m.onPremisesSyncEnabled
}
// GetOperatingSystem gets the operatingSystem property value. The type of operating system on the device. Required. Supports $filter (eq, ne, not, ge, le, startsWith, and eq on null values).
func (m *Device) GetOperatingSystem()(*string) {
    return m.operatingSystem
}
// GetOperatingSystemVersion gets the operatingSystemVersion property value. The version of the operating system on the device. Required. Supports $filter (eq, ne, not, ge, le, startsWith, and eq on null values).
func (m *Device) GetOperatingSystemVersion()(*string) {
    return m.operatingSystemVersion
}
// GetPhysicalIds gets the physicalIds property value. For internal use only. Not nullable. Supports $filter (eq, not, ge, le, startsWith,/$count eq 0, /$count ne 0).
func (m *Device) GetPhysicalIds()([]string) {
    return m.physicalIds
}
// GetProfileType gets the profileType property value. The profile type of the device. Possible values: RegisteredDevice (default), SecureVM, Printer, Shared, IoT.
func (m *Device) GetProfileType()(*string) {
    return m.profileType
}
// GetRegisteredOwners gets the registeredOwners property value. The user that cloud joined the device or registered their personal device. The registered owner is set at the time of registration. Currently, there can be only one owner. Read-only. Nullable. Supports $expand.
func (m *Device) GetRegisteredOwners()([]DirectoryObjectable) {
    return m.registeredOwners
}
// GetRegisteredUsers gets the registeredUsers property value. Collection of registered users of the device. For cloud joined devices and registered personal devices, registered users are set to the same value as registered owners at the time of registration. Read-only. Nullable. Supports $expand.
func (m *Device) GetRegisteredUsers()([]DirectoryObjectable) {
    return m.registeredUsers
}
// GetSystemLabels gets the systemLabels property value. List of labels applied to the device by the system. Supports $filter (/$count eq 0, /$count ne 0).
func (m *Device) GetSystemLabels()([]string) {
    return m.systemLabels
}
// GetTransitiveMemberOf gets the transitiveMemberOf property value. Groups and administrative units that the device is a member of. This operation is transitive. Supports $expand.
func (m *Device) GetTransitiveMemberOf()([]DirectoryObjectable) {
    return m.transitiveMemberOf
}
// GetTrustType gets the trustType property value. Type of trust for the joined device. Read-only. Possible values:  Workplace (indicates bring your own personal devices), AzureAd (Cloud only joined devices), ServerAd (on-premises domain joined devices joined to Azure AD). For more details, see Introduction to device management in Azure Active Directory
func (m *Device) GetTrustType()(*string) {
    return m.trustType
}
// Serialize serializes information the current object
func (m *Device) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DirectoryObject.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("accountEnabled", m.GetAccountEnabled())
        if err != nil {
            return err
        }
    }
    if m.GetAlternativeSecurityIds() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAlternativeSecurityIds()))
        for i, v := range m.GetAlternativeSecurityIds() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("alternativeSecurityIds", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("approximateLastSignInDateTime", m.GetApproximateLastSignInDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("complianceExpirationDateTime", m.GetComplianceExpirationDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceId", m.GetDeviceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceMetadata", m.GetDeviceMetadata())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("deviceVersion", m.GetDeviceVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetExtensions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetExtensions()))
        for i, v := range m.GetExtensions() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("extensions", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isCompliant", m.GetIsCompliant())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isManaged", m.GetIsManaged())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("mdmAppId", m.GetMdmAppId())
        if err != nil {
            return err
        }
    }
    if m.GetMemberOf() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMemberOf()))
        for i, v := range m.GetMemberOf() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("memberOf", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("onPremisesLastSyncDateTime", m.GetOnPremisesLastSyncDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("onPremisesSyncEnabled", m.GetOnPremisesSyncEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("operatingSystem", m.GetOperatingSystem())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("operatingSystemVersion", m.GetOperatingSystemVersion())
        if err != nil {
            return err
        }
    }
    if m.GetPhysicalIds() != nil {
        err = writer.WriteCollectionOfStringValues("physicalIds", m.GetPhysicalIds())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("profileType", m.GetProfileType())
        if err != nil {
            return err
        }
    }
    if m.GetRegisteredOwners() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRegisteredOwners()))
        for i, v := range m.GetRegisteredOwners() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("registeredOwners", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRegisteredUsers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRegisteredUsers()))
        for i, v := range m.GetRegisteredUsers() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("registeredUsers", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSystemLabels() != nil {
        err = writer.WriteCollectionOfStringValues("systemLabels", m.GetSystemLabels())
        if err != nil {
            return err
        }
    }
    if m.GetTransitiveMemberOf() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTransitiveMemberOf()))
        for i, v := range m.GetTransitiveMemberOf() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("transitiveMemberOf", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("trustType", m.GetTrustType())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccountEnabled sets the accountEnabled property value. true if the account is enabled; otherwise, false. Required. Default is true.  Supports $filter (eq, ne, not, in). Only callers in Global Administrator and Cloud Device Administrator roles can set this property.
func (m *Device) SetAccountEnabled(value *bool)() {
    m.accountEnabled = value
}
// SetAlternativeSecurityIds sets the alternativeSecurityIds property value. For internal use only. Not nullable. Supports $filter (eq, not, ge, le).
func (m *Device) SetAlternativeSecurityIds(value []AlternativeSecurityIdable)() {
    m.alternativeSecurityIds = value
}
// SetApproximateLastSignInDateTime sets the approximateLastSignInDateTime property value. The timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only. Supports $filter (eq, ne, not, ge, le, and eq on null values) and $orderBy.
func (m *Device) SetApproximateLastSignInDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.approximateLastSignInDateTime = value
}
// SetComplianceExpirationDateTime sets the complianceExpirationDateTime property value. The timestamp when the device is no longer deemed compliant. The timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *Device) SetComplianceExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.complianceExpirationDateTime = value
}
// SetDeviceId sets the deviceId property value. Unique identifier set by Azure Device Registration Service at the time of registration. This is an alternate key that can be used to reference the device object. Supports $filter (eq, ne, not, startsWith).
func (m *Device) SetDeviceId(value *string)() {
    m.deviceId = value
}
// SetDeviceMetadata sets the deviceMetadata property value. For internal use only. Set to null.
func (m *Device) SetDeviceMetadata(value *string)() {
    m.deviceMetadata = value
}
// SetDeviceVersion sets the deviceVersion property value. For internal use only.
func (m *Device) SetDeviceVersion(value *int32)() {
    m.deviceVersion = value
}
// SetDisplayName sets the displayName property value. The display name for the device. Required. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values), $search, and $orderBy.
func (m *Device) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetExtensions sets the extensions property value. The collection of open extensions defined for the device. Read-only. Nullable.
func (m *Device) SetExtensions(value []Extensionable)() {
    m.extensions = value
}
// SetIsCompliant sets the isCompliant property value. true if the device complies with Mobile Device Management (MDM) policies; otherwise, false. Read-only. This can only be updated by Intune for any device OS type or by an approved MDM app for Windows OS devices. Supports $filter (eq, ne, not).
func (m *Device) SetIsCompliant(value *bool)() {
    m.isCompliant = value
}
// SetIsManaged sets the isManaged property value. true if the device is managed by a Mobile Device Management (MDM) app; otherwise, false. This can only be updated by Intune for any device OS type or by an approved MDM app for Windows OS devices. Supports $filter (eq, ne, not).
func (m *Device) SetIsManaged(value *bool)() {
    m.isManaged = value
}
// SetMdmAppId sets the mdmAppId property value. Application identifier used to register device into MDM. Read-only. Supports $filter (eq, ne, not, startsWith).
func (m *Device) SetMdmAppId(value *string)() {
    m.mdmAppId = value
}
// SetMemberOf sets the memberOf property value. Groups and administrative units that this device is a member of. Read-only. Nullable. Supports $expand.
func (m *Device) SetMemberOf(value []DirectoryObjectable)() {
    m.memberOf = value
}
// SetOnPremisesLastSyncDateTime sets the onPremisesLastSyncDateTime property value. The last time at which the object was synced with the on-premises directory. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z Read-only. Supports $filter (eq, ne, not, ge, le, in).
func (m *Device) SetOnPremisesLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.onPremisesLastSyncDateTime = value
}
// SetOnPremisesSyncEnabled sets the onPremisesSyncEnabled property value. true if this object is synced from an on-premises directory; false if this object was originally synced from an on-premises directory but is no longer synced; null if this object has never been synced from an on-premises directory (default). Read-only. Supports $filter (eq, ne, not, in, and eq on null values).
func (m *Device) SetOnPremisesSyncEnabled(value *bool)() {
    m.onPremisesSyncEnabled = value
}
// SetOperatingSystem sets the operatingSystem property value. The type of operating system on the device. Required. Supports $filter (eq, ne, not, ge, le, startsWith, and eq on null values).
func (m *Device) SetOperatingSystem(value *string)() {
    m.operatingSystem = value
}
// SetOperatingSystemVersion sets the operatingSystemVersion property value. The version of the operating system on the device. Required. Supports $filter (eq, ne, not, ge, le, startsWith, and eq on null values).
func (m *Device) SetOperatingSystemVersion(value *string)() {
    m.operatingSystemVersion = value
}
// SetPhysicalIds sets the physicalIds property value. For internal use only. Not nullable. Supports $filter (eq, not, ge, le, startsWith,/$count eq 0, /$count ne 0).
func (m *Device) SetPhysicalIds(value []string)() {
    m.physicalIds = value
}
// SetProfileType sets the profileType property value. The profile type of the device. Possible values: RegisteredDevice (default), SecureVM, Printer, Shared, IoT.
func (m *Device) SetProfileType(value *string)() {
    m.profileType = value
}
// SetRegisteredOwners sets the registeredOwners property value. The user that cloud joined the device or registered their personal device. The registered owner is set at the time of registration. Currently, there can be only one owner. Read-only. Nullable. Supports $expand.
func (m *Device) SetRegisteredOwners(value []DirectoryObjectable)() {
    m.registeredOwners = value
}
// SetRegisteredUsers sets the registeredUsers property value. Collection of registered users of the device. For cloud joined devices and registered personal devices, registered users are set to the same value as registered owners at the time of registration. Read-only. Nullable. Supports $expand.
func (m *Device) SetRegisteredUsers(value []DirectoryObjectable)() {
    m.registeredUsers = value
}
// SetSystemLabels sets the systemLabels property value. List of labels applied to the device by the system. Supports $filter (/$count eq 0, /$count ne 0).
func (m *Device) SetSystemLabels(value []string)() {
    m.systemLabels = value
}
// SetTransitiveMemberOf sets the transitiveMemberOf property value. Groups and administrative units that the device is a member of. This operation is transitive. Supports $expand.
func (m *Device) SetTransitiveMemberOf(value []DirectoryObjectable)() {
    m.transitiveMemberOf = value
}
// SetTrustType sets the trustType property value. Type of trust for the joined device. Read-only. Possible values:  Workplace (indicates bring your own personal devices), AzureAd (Cloud only joined devices), ServerAd (on-premises domain joined devices joined to Azure AD). For more details, see Introduction to device management in Azure Active Directory
func (m *Device) SetTrustType(value *string)() {
    m.trustType = value
}
