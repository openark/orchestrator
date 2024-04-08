package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceCompliancePolicySettingState device Compilance Policy Setting State for a given device.
type DeviceCompliancePolicySettingState struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Current value of setting on device
    currentValue *string
    // Error code for the setting
    errorCode *int64
    // Error description
    errorDescription *string
    // Name of setting instance that is being reported.
    instanceDisplayName *string
    // The OdataType property
    odataType *string
    // The setting that is being reported
    setting *string
    // Localized/user friendly setting name that is being reported
    settingName *string
    // Contributing policies
    sources []SettingSourceable
    // The state property
    state *ComplianceStatus
    // UserEmail
    userEmail *string
    // UserId
    userId *string
    // UserName
    userName *string
    // UserPrincipalName.
    userPrincipalName *string
}
// NewDeviceCompliancePolicySettingState instantiates a new deviceCompliancePolicySettingState and sets the default values.
func NewDeviceCompliancePolicySettingState()(*DeviceCompliancePolicySettingState) {
    m := &DeviceCompliancePolicySettingState{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDeviceCompliancePolicySettingStateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceCompliancePolicySettingStateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceCompliancePolicySettingState(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceCompliancePolicySettingState) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCurrentValue gets the currentValue property value. Current value of setting on device
func (m *DeviceCompliancePolicySettingState) GetCurrentValue()(*string) {
    return m.currentValue
}
// GetErrorCode gets the errorCode property value. Error code for the setting
func (m *DeviceCompliancePolicySettingState) GetErrorCode()(*int64) {
    return m.errorCode
}
// GetErrorDescription gets the errorDescription property value. Error description
func (m *DeviceCompliancePolicySettingState) GetErrorDescription()(*string) {
    return m.errorDescription
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceCompliancePolicySettingState) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["currentValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCurrentValue(val)
        }
        return nil
    }
    res["errorCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorCode(val)
        }
        return nil
    }
    res["errorDescription"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorDescription(val)
        }
        return nil
    }
    res["instanceDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInstanceDisplayName(val)
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
    res["setting"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSetting(val)
        }
        return nil
    }
    res["settingName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingName(val)
        }
        return nil
    }
    res["sources"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSettingSourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SettingSourceable, len(val))
            for i, v := range val {
                res[i] = v.(SettingSourceable)
            }
            m.SetSources(res)
        }
        return nil
    }
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseComplianceStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(*ComplianceStatus))
        }
        return nil
    }
    res["userEmail"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserEmail(val)
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
    res["userName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserName(val)
        }
        return nil
    }
    res["userPrincipalName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserPrincipalName(val)
        }
        return nil
    }
    return res
}
// GetInstanceDisplayName gets the instanceDisplayName property value. Name of setting instance that is being reported.
func (m *DeviceCompliancePolicySettingState) GetInstanceDisplayName()(*string) {
    return m.instanceDisplayName
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *DeviceCompliancePolicySettingState) GetOdataType()(*string) {
    return m.odataType
}
// GetSetting gets the setting property value. The setting that is being reported
func (m *DeviceCompliancePolicySettingState) GetSetting()(*string) {
    return m.setting
}
// GetSettingName gets the settingName property value. Localized/user friendly setting name that is being reported
func (m *DeviceCompliancePolicySettingState) GetSettingName()(*string) {
    return m.settingName
}
// GetSources gets the sources property value. Contributing policies
func (m *DeviceCompliancePolicySettingState) GetSources()([]SettingSourceable) {
    return m.sources
}
// GetState gets the state property value. The state property
func (m *DeviceCompliancePolicySettingState) GetState()(*ComplianceStatus) {
    return m.state
}
// GetUserEmail gets the userEmail property value. UserEmail
func (m *DeviceCompliancePolicySettingState) GetUserEmail()(*string) {
    return m.userEmail
}
// GetUserId gets the userId property value. UserId
func (m *DeviceCompliancePolicySettingState) GetUserId()(*string) {
    return m.userId
}
// GetUserName gets the userName property value. UserName
func (m *DeviceCompliancePolicySettingState) GetUserName()(*string) {
    return m.userName
}
// GetUserPrincipalName gets the userPrincipalName property value. UserPrincipalName.
func (m *DeviceCompliancePolicySettingState) GetUserPrincipalName()(*string) {
    return m.userPrincipalName
}
// Serialize serializes information the current object
func (m *DeviceCompliancePolicySettingState) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("currentValue", m.GetCurrentValue())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("errorCode", m.GetErrorCode())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("errorDescription", m.GetErrorDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("instanceDisplayName", m.GetInstanceDisplayName())
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
        err := writer.WriteStringValue("setting", m.GetSetting())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("settingName", m.GetSettingName())
        if err != nil {
            return err
        }
    }
    if m.GetSources() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSources()))
        for i, v := range m.GetSources() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("sources", cast)
        if err != nil {
            return err
        }
    }
    if m.GetState() != nil {
        cast := (*m.GetState()).String()
        err := writer.WriteStringValue("state", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("userEmail", m.GetUserEmail())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("userId", m.GetUserId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("userName", m.GetUserName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("userPrincipalName", m.GetUserPrincipalName())
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
func (m *DeviceCompliancePolicySettingState) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCurrentValue sets the currentValue property value. Current value of setting on device
func (m *DeviceCompliancePolicySettingState) SetCurrentValue(value *string)() {
    m.currentValue = value
}
// SetErrorCode sets the errorCode property value. Error code for the setting
func (m *DeviceCompliancePolicySettingState) SetErrorCode(value *int64)() {
    m.errorCode = value
}
// SetErrorDescription sets the errorDescription property value. Error description
func (m *DeviceCompliancePolicySettingState) SetErrorDescription(value *string)() {
    m.errorDescription = value
}
// SetInstanceDisplayName sets the instanceDisplayName property value. Name of setting instance that is being reported.
func (m *DeviceCompliancePolicySettingState) SetInstanceDisplayName(value *string)() {
    m.instanceDisplayName = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DeviceCompliancePolicySettingState) SetOdataType(value *string)() {
    m.odataType = value
}
// SetSetting sets the setting property value. The setting that is being reported
func (m *DeviceCompliancePolicySettingState) SetSetting(value *string)() {
    m.setting = value
}
// SetSettingName sets the settingName property value. Localized/user friendly setting name that is being reported
func (m *DeviceCompliancePolicySettingState) SetSettingName(value *string)() {
    m.settingName = value
}
// SetSources sets the sources property value. Contributing policies
func (m *DeviceCompliancePolicySettingState) SetSources(value []SettingSourceable)() {
    m.sources = value
}
// SetState sets the state property value. The state property
func (m *DeviceCompliancePolicySettingState) SetState(value *ComplianceStatus)() {
    m.state = value
}
// SetUserEmail sets the userEmail property value. UserEmail
func (m *DeviceCompliancePolicySettingState) SetUserEmail(value *string)() {
    m.userEmail = value
}
// SetUserId sets the userId property value. UserId
func (m *DeviceCompliancePolicySettingState) SetUserId(value *string)() {
    m.userId = value
}
// SetUserName sets the userName property value. UserName
func (m *DeviceCompliancePolicySettingState) SetUserName(value *string)() {
    m.userName = value
}
// SetUserPrincipalName sets the userPrincipalName property value. UserPrincipalName.
func (m *DeviceCompliancePolicySettingState) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
