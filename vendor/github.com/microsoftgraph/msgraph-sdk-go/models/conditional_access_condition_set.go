package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ConditionalAccessConditionSet 
type ConditionalAccessConditionSet struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Applications and user actions included in and excluded from the policy. Required.
    applications ConditionalAccessApplicationsable
    // Client applications (service principals and workload identities) included in and excluded from the policy. Either users or clientApplications is required.
    clientApplications ConditionalAccessClientApplicationsable
    // Client application types included in the policy. Possible values are: all, browser, mobileAppsAndDesktopClients, exchangeActiveSync, easSupported, other. Required.
    clientAppTypes []ConditionalAccessClientApp
    // Devices in the policy.
    devices ConditionalAccessDevicesable
    // Locations included in and excluded from the policy.
    locations ConditionalAccessLocationsable
    // The OdataType property
    odataType *string
    // Platforms included in and excluded from the policy.
    platforms ConditionalAccessPlatformsable
    // Service principal risk levels included in the policy. Possible values are: low, medium, high, none, unknownFutureValue.
    servicePrincipalRiskLevels []RiskLevel
    // Sign-in risk levels included in the policy. Possible values are: low, medium, high, hidden, none, unknownFutureValue. Required.
    signInRiskLevels []RiskLevel
    // User risk levels included in the policy. Possible values are: low, medium, high, hidden, none, unknownFutureValue. Required.
    userRiskLevels []RiskLevel
    // Users, groups, and roles included in and excluded from the policy. Either users or clientApplications is required.
    users ConditionalAccessUsersable
}
// NewConditionalAccessConditionSet instantiates a new conditionalAccessConditionSet and sets the default values.
func NewConditionalAccessConditionSet()(*ConditionalAccessConditionSet) {
    m := &ConditionalAccessConditionSet{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateConditionalAccessConditionSetFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateConditionalAccessConditionSetFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewConditionalAccessConditionSet(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ConditionalAccessConditionSet) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetApplications gets the applications property value. Applications and user actions included in and excluded from the policy. Required.
func (m *ConditionalAccessConditionSet) GetApplications()(ConditionalAccessApplicationsable) {
    return m.applications
}
// GetClientApplications gets the clientApplications property value. Client applications (service principals and workload identities) included in and excluded from the policy. Either users or clientApplications is required.
func (m *ConditionalAccessConditionSet) GetClientApplications()(ConditionalAccessClientApplicationsable) {
    return m.clientApplications
}
// GetClientAppTypes gets the clientAppTypes property value. Client application types included in the policy. Possible values are: all, browser, mobileAppsAndDesktopClients, exchangeActiveSync, easSupported, other. Required.
func (m *ConditionalAccessConditionSet) GetClientAppTypes()([]ConditionalAccessClientApp) {
    return m.clientAppTypes
}
// GetDevices gets the devices property value. Devices in the policy.
func (m *ConditionalAccessConditionSet) GetDevices()(ConditionalAccessDevicesable) {
    return m.devices
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ConditionalAccessConditionSet) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["applications"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateConditionalAccessApplicationsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplications(val.(ConditionalAccessApplicationsable))
        }
        return nil
    }
    res["clientApplications"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateConditionalAccessClientApplicationsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClientApplications(val.(ConditionalAccessClientApplicationsable))
        }
        return nil
    }
    res["clientAppTypes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseConditionalAccessClientApp)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ConditionalAccessClientApp, len(val))
            for i, v := range val {
                res[i] = *(v.(*ConditionalAccessClientApp))
            }
            m.SetClientAppTypes(res)
        }
        return nil
    }
    res["devices"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateConditionalAccessDevicesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDevices(val.(ConditionalAccessDevicesable))
        }
        return nil
    }
    res["locations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateConditionalAccessLocationsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocations(val.(ConditionalAccessLocationsable))
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
    res["platforms"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateConditionalAccessPlatformsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPlatforms(val.(ConditionalAccessPlatformsable))
        }
        return nil
    }
    res["servicePrincipalRiskLevels"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseRiskLevel)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RiskLevel, len(val))
            for i, v := range val {
                res[i] = *(v.(*RiskLevel))
            }
            m.SetServicePrincipalRiskLevels(res)
        }
        return nil
    }
    res["signInRiskLevels"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseRiskLevel)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RiskLevel, len(val))
            for i, v := range val {
                res[i] = *(v.(*RiskLevel))
            }
            m.SetSignInRiskLevels(res)
        }
        return nil
    }
    res["userRiskLevels"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseRiskLevel)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RiskLevel, len(val))
            for i, v := range val {
                res[i] = *(v.(*RiskLevel))
            }
            m.SetUserRiskLevels(res)
        }
        return nil
    }
    res["users"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateConditionalAccessUsersFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUsers(val.(ConditionalAccessUsersable))
        }
        return nil
    }
    return res
}
// GetLocations gets the locations property value. Locations included in and excluded from the policy.
func (m *ConditionalAccessConditionSet) GetLocations()(ConditionalAccessLocationsable) {
    return m.locations
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ConditionalAccessConditionSet) GetOdataType()(*string) {
    return m.odataType
}
// GetPlatforms gets the platforms property value. Platforms included in and excluded from the policy.
func (m *ConditionalAccessConditionSet) GetPlatforms()(ConditionalAccessPlatformsable) {
    return m.platforms
}
// GetServicePrincipalRiskLevels gets the servicePrincipalRiskLevels property value. Service principal risk levels included in the policy. Possible values are: low, medium, high, none, unknownFutureValue.
func (m *ConditionalAccessConditionSet) GetServicePrincipalRiskLevels()([]RiskLevel) {
    return m.servicePrincipalRiskLevels
}
// GetSignInRiskLevels gets the signInRiskLevels property value. Sign-in risk levels included in the policy. Possible values are: low, medium, high, hidden, none, unknownFutureValue. Required.
func (m *ConditionalAccessConditionSet) GetSignInRiskLevels()([]RiskLevel) {
    return m.signInRiskLevels
}
// GetUserRiskLevels gets the userRiskLevels property value. User risk levels included in the policy. Possible values are: low, medium, high, hidden, none, unknownFutureValue. Required.
func (m *ConditionalAccessConditionSet) GetUserRiskLevels()([]RiskLevel) {
    return m.userRiskLevels
}
// GetUsers gets the users property value. Users, groups, and roles included in and excluded from the policy. Either users or clientApplications is required.
func (m *ConditionalAccessConditionSet) GetUsers()(ConditionalAccessUsersable) {
    return m.users
}
// Serialize serializes information the current object
func (m *ConditionalAccessConditionSet) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("applications", m.GetApplications())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("clientApplications", m.GetClientApplications())
        if err != nil {
            return err
        }
    }
    if m.GetClientAppTypes() != nil {
        err := writer.WriteCollectionOfStringValues("clientAppTypes", SerializeConditionalAccessClientApp(m.GetClientAppTypes()))
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("devices", m.GetDevices())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("locations", m.GetLocations())
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
        err := writer.WriteObjectValue("platforms", m.GetPlatforms())
        if err != nil {
            return err
        }
    }
    if m.GetServicePrincipalRiskLevels() != nil {
        err := writer.WriteCollectionOfStringValues("servicePrincipalRiskLevels", SerializeRiskLevel(m.GetServicePrincipalRiskLevels()))
        if err != nil {
            return err
        }
    }
    if m.GetSignInRiskLevels() != nil {
        err := writer.WriteCollectionOfStringValues("signInRiskLevels", SerializeRiskLevel(m.GetSignInRiskLevels()))
        if err != nil {
            return err
        }
    }
    if m.GetUserRiskLevels() != nil {
        err := writer.WriteCollectionOfStringValues("userRiskLevels", SerializeRiskLevel(m.GetUserRiskLevels()))
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("users", m.GetUsers())
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
func (m *ConditionalAccessConditionSet) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetApplications sets the applications property value. Applications and user actions included in and excluded from the policy. Required.
func (m *ConditionalAccessConditionSet) SetApplications(value ConditionalAccessApplicationsable)() {
    m.applications = value
}
// SetClientApplications sets the clientApplications property value. Client applications (service principals and workload identities) included in and excluded from the policy. Either users or clientApplications is required.
func (m *ConditionalAccessConditionSet) SetClientApplications(value ConditionalAccessClientApplicationsable)() {
    m.clientApplications = value
}
// SetClientAppTypes sets the clientAppTypes property value. Client application types included in the policy. Possible values are: all, browser, mobileAppsAndDesktopClients, exchangeActiveSync, easSupported, other. Required.
func (m *ConditionalAccessConditionSet) SetClientAppTypes(value []ConditionalAccessClientApp)() {
    m.clientAppTypes = value
}
// SetDevices sets the devices property value. Devices in the policy.
func (m *ConditionalAccessConditionSet) SetDevices(value ConditionalAccessDevicesable)() {
    m.devices = value
}
// SetLocations sets the locations property value. Locations included in and excluded from the policy.
func (m *ConditionalAccessConditionSet) SetLocations(value ConditionalAccessLocationsable)() {
    m.locations = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ConditionalAccessConditionSet) SetOdataType(value *string)() {
    m.odataType = value
}
// SetPlatforms sets the platforms property value. Platforms included in and excluded from the policy.
func (m *ConditionalAccessConditionSet) SetPlatforms(value ConditionalAccessPlatformsable)() {
    m.platforms = value
}
// SetServicePrincipalRiskLevels sets the servicePrincipalRiskLevels property value. Service principal risk levels included in the policy. Possible values are: low, medium, high, none, unknownFutureValue.
func (m *ConditionalAccessConditionSet) SetServicePrincipalRiskLevels(value []RiskLevel)() {
    m.servicePrincipalRiskLevels = value
}
// SetSignInRiskLevels sets the signInRiskLevels property value. Sign-in risk levels included in the policy. Possible values are: low, medium, high, hidden, none, unknownFutureValue. Required.
func (m *ConditionalAccessConditionSet) SetSignInRiskLevels(value []RiskLevel)() {
    m.signInRiskLevels = value
}
// SetUserRiskLevels sets the userRiskLevels property value. User risk levels included in the policy. Possible values are: low, medium, high, hidden, none, unknownFutureValue. Required.
func (m *ConditionalAccessConditionSet) SetUserRiskLevels(value []RiskLevel)() {
    m.userRiskLevels = value
}
// SetUsers sets the users property value. Users, groups, and roles included in and excluded from the policy. Either users or clientApplications is required.
func (m *ConditionalAccessConditionSet) SetUsers(value ConditionalAccessUsersable)() {
    m.users = value
}
