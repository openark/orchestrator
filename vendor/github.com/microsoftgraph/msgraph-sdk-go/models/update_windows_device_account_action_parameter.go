package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UpdateWindowsDeviceAccountActionParameter 
type UpdateWindowsDeviceAccountActionParameter struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Not yet documented
    calendarSyncEnabled *bool
    // Not yet documented
    deviceAccount WindowsDeviceAccountable
    // Not yet documented
    deviceAccountEmail *string
    // Not yet documented
    exchangeServer *string
    // The OdataType property
    odataType *string
    // Not yet documented
    passwordRotationEnabled *bool
    // Not yet documented
    sessionInitiationProtocalAddress *string
}
// NewUpdateWindowsDeviceAccountActionParameter instantiates a new updateWindowsDeviceAccountActionParameter and sets the default values.
func NewUpdateWindowsDeviceAccountActionParameter()(*UpdateWindowsDeviceAccountActionParameter) {
    m := &UpdateWindowsDeviceAccountActionParameter{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateUpdateWindowsDeviceAccountActionParameterFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUpdateWindowsDeviceAccountActionParameterFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUpdateWindowsDeviceAccountActionParameter(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UpdateWindowsDeviceAccountActionParameter) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCalendarSyncEnabled gets the calendarSyncEnabled property value. Not yet documented
func (m *UpdateWindowsDeviceAccountActionParameter) GetCalendarSyncEnabled()(*bool) {
    return m.calendarSyncEnabled
}
// GetDeviceAccount gets the deviceAccount property value. Not yet documented
func (m *UpdateWindowsDeviceAccountActionParameter) GetDeviceAccount()(WindowsDeviceAccountable) {
    return m.deviceAccount
}
// GetDeviceAccountEmail gets the deviceAccountEmail property value. Not yet documented
func (m *UpdateWindowsDeviceAccountActionParameter) GetDeviceAccountEmail()(*string) {
    return m.deviceAccountEmail
}
// GetExchangeServer gets the exchangeServer property value. Not yet documented
func (m *UpdateWindowsDeviceAccountActionParameter) GetExchangeServer()(*string) {
    return m.exchangeServer
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UpdateWindowsDeviceAccountActionParameter) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["calendarSyncEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCalendarSyncEnabled(val)
        }
        return nil
    }
    res["deviceAccount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWindowsDeviceAccountFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceAccount(val.(WindowsDeviceAccountable))
        }
        return nil
    }
    res["deviceAccountEmail"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceAccountEmail(val)
        }
        return nil
    }
    res["exchangeServer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExchangeServer(val)
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
    res["passwordRotationEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordRotationEnabled(val)
        }
        return nil
    }
    res["sessionInitiationProtocalAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSessionInitiationProtocalAddress(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *UpdateWindowsDeviceAccountActionParameter) GetOdataType()(*string) {
    return m.odataType
}
// GetPasswordRotationEnabled gets the passwordRotationEnabled property value. Not yet documented
func (m *UpdateWindowsDeviceAccountActionParameter) GetPasswordRotationEnabled()(*bool) {
    return m.passwordRotationEnabled
}
// GetSessionInitiationProtocalAddress gets the sessionInitiationProtocalAddress property value. Not yet documented
func (m *UpdateWindowsDeviceAccountActionParameter) GetSessionInitiationProtocalAddress()(*string) {
    return m.sessionInitiationProtocalAddress
}
// Serialize serializes information the current object
func (m *UpdateWindowsDeviceAccountActionParameter) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("calendarSyncEnabled", m.GetCalendarSyncEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("deviceAccount", m.GetDeviceAccount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("deviceAccountEmail", m.GetDeviceAccountEmail())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("exchangeServer", m.GetExchangeServer())
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
        err := writer.WriteBoolValue("passwordRotationEnabled", m.GetPasswordRotationEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("sessionInitiationProtocalAddress", m.GetSessionInitiationProtocalAddress())
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
func (m *UpdateWindowsDeviceAccountActionParameter) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCalendarSyncEnabled sets the calendarSyncEnabled property value. Not yet documented
func (m *UpdateWindowsDeviceAccountActionParameter) SetCalendarSyncEnabled(value *bool)() {
    m.calendarSyncEnabled = value
}
// SetDeviceAccount sets the deviceAccount property value. Not yet documented
func (m *UpdateWindowsDeviceAccountActionParameter) SetDeviceAccount(value WindowsDeviceAccountable)() {
    m.deviceAccount = value
}
// SetDeviceAccountEmail sets the deviceAccountEmail property value. Not yet documented
func (m *UpdateWindowsDeviceAccountActionParameter) SetDeviceAccountEmail(value *string)() {
    m.deviceAccountEmail = value
}
// SetExchangeServer sets the exchangeServer property value. Not yet documented
func (m *UpdateWindowsDeviceAccountActionParameter) SetExchangeServer(value *string)() {
    m.exchangeServer = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *UpdateWindowsDeviceAccountActionParameter) SetOdataType(value *string)() {
    m.odataType = value
}
// SetPasswordRotationEnabled sets the passwordRotationEnabled property value. Not yet documented
func (m *UpdateWindowsDeviceAccountActionParameter) SetPasswordRotationEnabled(value *bool)() {
    m.passwordRotationEnabled = value
}
// SetSessionInitiationProtocalAddress sets the sessionInitiationProtocalAddress property value. Not yet documented
func (m *UpdateWindowsDeviceAccountActionParameter) SetSessionInitiationProtocalAddress(value *string)() {
    m.sessionInitiationProtocalAddress = value
}
