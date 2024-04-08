package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ComplianceManagementPartner compliance management partner for all platforms
type ComplianceManagementPartner struct {
    Entity
    // User groups which enroll Android devices through partner.
    androidEnrollmentAssignments []ComplianceManagementPartnerAssignmentable
    // Partner onboarded for Android devices.
    androidOnboarded *bool
    // Partner display name
    displayName *string
    // User groups which enroll ios devices through partner.
    iosEnrollmentAssignments []ComplianceManagementPartnerAssignmentable
    // Partner onboarded for ios devices.
    iosOnboarded *bool
    // Timestamp of last heartbeat after admin onboarded to the compliance management partner
    lastHeartbeatDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // User groups which enroll Mac devices through partner.
    macOsEnrollmentAssignments []ComplianceManagementPartnerAssignmentable
    // Partner onboarded for Mac devices.
    macOsOnboarded *bool
    // Partner state of this tenant.
    partnerState *DeviceManagementPartnerTenantState
}
// NewComplianceManagementPartner instantiates a new complianceManagementPartner and sets the default values.
func NewComplianceManagementPartner()(*ComplianceManagementPartner) {
    m := &ComplianceManagementPartner{
        Entity: *NewEntity(),
    }
    return m
}
// CreateComplianceManagementPartnerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateComplianceManagementPartnerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewComplianceManagementPartner(), nil
}
// GetAndroidEnrollmentAssignments gets the androidEnrollmentAssignments property value. User groups which enroll Android devices through partner.
func (m *ComplianceManagementPartner) GetAndroidEnrollmentAssignments()([]ComplianceManagementPartnerAssignmentable) {
    return m.androidEnrollmentAssignments
}
// GetAndroidOnboarded gets the androidOnboarded property value. Partner onboarded for Android devices.
func (m *ComplianceManagementPartner) GetAndroidOnboarded()(*bool) {
    return m.androidOnboarded
}
// GetDisplayName gets the displayName property value. Partner display name
func (m *ComplianceManagementPartner) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ComplianceManagementPartner) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["androidEnrollmentAssignments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateComplianceManagementPartnerAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ComplianceManagementPartnerAssignmentable, len(val))
            for i, v := range val {
                res[i] = v.(ComplianceManagementPartnerAssignmentable)
            }
            m.SetAndroidEnrollmentAssignments(res)
        }
        return nil
    }
    res["androidOnboarded"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAndroidOnboarded(val)
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
    res["iosEnrollmentAssignments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateComplianceManagementPartnerAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ComplianceManagementPartnerAssignmentable, len(val))
            for i, v := range val {
                res[i] = v.(ComplianceManagementPartnerAssignmentable)
            }
            m.SetIosEnrollmentAssignments(res)
        }
        return nil
    }
    res["iosOnboarded"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIosOnboarded(val)
        }
        return nil
    }
    res["lastHeartbeatDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastHeartbeatDateTime(val)
        }
        return nil
    }
    res["macOsEnrollmentAssignments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateComplianceManagementPartnerAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ComplianceManagementPartnerAssignmentable, len(val))
            for i, v := range val {
                res[i] = v.(ComplianceManagementPartnerAssignmentable)
            }
            m.SetMacOsEnrollmentAssignments(res)
        }
        return nil
    }
    res["macOsOnboarded"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMacOsOnboarded(val)
        }
        return nil
    }
    res["partnerState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementPartnerTenantState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPartnerState(val.(*DeviceManagementPartnerTenantState))
        }
        return nil
    }
    return res
}
// GetIosEnrollmentAssignments gets the iosEnrollmentAssignments property value. User groups which enroll ios devices through partner.
func (m *ComplianceManagementPartner) GetIosEnrollmentAssignments()([]ComplianceManagementPartnerAssignmentable) {
    return m.iosEnrollmentAssignments
}
// GetIosOnboarded gets the iosOnboarded property value. Partner onboarded for ios devices.
func (m *ComplianceManagementPartner) GetIosOnboarded()(*bool) {
    return m.iosOnboarded
}
// GetLastHeartbeatDateTime gets the lastHeartbeatDateTime property value. Timestamp of last heartbeat after admin onboarded to the compliance management partner
func (m *ComplianceManagementPartner) GetLastHeartbeatDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastHeartbeatDateTime
}
// GetMacOsEnrollmentAssignments gets the macOsEnrollmentAssignments property value. User groups which enroll Mac devices through partner.
func (m *ComplianceManagementPartner) GetMacOsEnrollmentAssignments()([]ComplianceManagementPartnerAssignmentable) {
    return m.macOsEnrollmentAssignments
}
// GetMacOsOnboarded gets the macOsOnboarded property value. Partner onboarded for Mac devices.
func (m *ComplianceManagementPartner) GetMacOsOnboarded()(*bool) {
    return m.macOsOnboarded
}
// GetPartnerState gets the partnerState property value. Partner state of this tenant.
func (m *ComplianceManagementPartner) GetPartnerState()(*DeviceManagementPartnerTenantState) {
    return m.partnerState
}
// Serialize serializes information the current object
func (m *ComplianceManagementPartner) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAndroidEnrollmentAssignments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAndroidEnrollmentAssignments()))
        for i, v := range m.GetAndroidEnrollmentAssignments() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("androidEnrollmentAssignments", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("androidOnboarded", m.GetAndroidOnboarded())
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
    if m.GetIosEnrollmentAssignments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetIosEnrollmentAssignments()))
        for i, v := range m.GetIosEnrollmentAssignments() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("iosEnrollmentAssignments", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("iosOnboarded", m.GetIosOnboarded())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastHeartbeatDateTime", m.GetLastHeartbeatDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetMacOsEnrollmentAssignments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMacOsEnrollmentAssignments()))
        for i, v := range m.GetMacOsEnrollmentAssignments() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("macOsEnrollmentAssignments", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("macOsOnboarded", m.GetMacOsOnboarded())
        if err != nil {
            return err
        }
    }
    if m.GetPartnerState() != nil {
        cast := (*m.GetPartnerState()).String()
        err = writer.WriteStringValue("partnerState", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAndroidEnrollmentAssignments sets the androidEnrollmentAssignments property value. User groups which enroll Android devices through partner.
func (m *ComplianceManagementPartner) SetAndroidEnrollmentAssignments(value []ComplianceManagementPartnerAssignmentable)() {
    m.androidEnrollmentAssignments = value
}
// SetAndroidOnboarded sets the androidOnboarded property value. Partner onboarded for Android devices.
func (m *ComplianceManagementPartner) SetAndroidOnboarded(value *bool)() {
    m.androidOnboarded = value
}
// SetDisplayName sets the displayName property value. Partner display name
func (m *ComplianceManagementPartner) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetIosEnrollmentAssignments sets the iosEnrollmentAssignments property value. User groups which enroll ios devices through partner.
func (m *ComplianceManagementPartner) SetIosEnrollmentAssignments(value []ComplianceManagementPartnerAssignmentable)() {
    m.iosEnrollmentAssignments = value
}
// SetIosOnboarded sets the iosOnboarded property value. Partner onboarded for ios devices.
func (m *ComplianceManagementPartner) SetIosOnboarded(value *bool)() {
    m.iosOnboarded = value
}
// SetLastHeartbeatDateTime sets the lastHeartbeatDateTime property value. Timestamp of last heartbeat after admin onboarded to the compliance management partner
func (m *ComplianceManagementPartner) SetLastHeartbeatDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastHeartbeatDateTime = value
}
// SetMacOsEnrollmentAssignments sets the macOsEnrollmentAssignments property value. User groups which enroll Mac devices through partner.
func (m *ComplianceManagementPartner) SetMacOsEnrollmentAssignments(value []ComplianceManagementPartnerAssignmentable)() {
    m.macOsEnrollmentAssignments = value
}
// SetMacOsOnboarded sets the macOsOnboarded property value. Partner onboarded for Mac devices.
func (m *ComplianceManagementPartner) SetMacOsOnboarded(value *bool)() {
    m.macOsOnboarded = value
}
// SetPartnerState sets the partnerState property value. Partner state of this tenant.
func (m *ComplianceManagementPartner) SetPartnerState(value *DeviceManagementPartnerTenantState)() {
    m.partnerState = value
}
