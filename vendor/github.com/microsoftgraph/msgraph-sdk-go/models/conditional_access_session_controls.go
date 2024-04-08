package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ConditionalAccessSessionControls 
type ConditionalAccessSessionControls struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Session control to enforce application restrictions. Only Exchange Online and Sharepoint Online support this session control.
    applicationEnforcedRestrictions ApplicationEnforcedRestrictionsSessionControlable
    // Session control to apply cloud app security.
    cloudAppSecurity CloudAppSecuritySessionControlable
    // Session control that determines whether it is acceptable for Azure AD to extend existing sessions based on information collected prior to an outage or not.
    disableResilienceDefaults *bool
    // The OdataType property
    odataType *string
    // Session control to define whether to persist cookies or not. All apps should be selected for this session control to work correctly.
    persistentBrowser PersistentBrowserSessionControlable
    // Session control to enforce signin frequency.
    signInFrequency SignInFrequencySessionControlable
}
// NewConditionalAccessSessionControls instantiates a new conditionalAccessSessionControls and sets the default values.
func NewConditionalAccessSessionControls()(*ConditionalAccessSessionControls) {
    m := &ConditionalAccessSessionControls{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateConditionalAccessSessionControlsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateConditionalAccessSessionControlsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewConditionalAccessSessionControls(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ConditionalAccessSessionControls) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetApplicationEnforcedRestrictions gets the applicationEnforcedRestrictions property value. Session control to enforce application restrictions. Only Exchange Online and Sharepoint Online support this session control.
func (m *ConditionalAccessSessionControls) GetApplicationEnforcedRestrictions()(ApplicationEnforcedRestrictionsSessionControlable) {
    return m.applicationEnforcedRestrictions
}
// GetCloudAppSecurity gets the cloudAppSecurity property value. Session control to apply cloud app security.
func (m *ConditionalAccessSessionControls) GetCloudAppSecurity()(CloudAppSecuritySessionControlable) {
    return m.cloudAppSecurity
}
// GetDisableResilienceDefaults gets the disableResilienceDefaults property value. Session control that determines whether it is acceptable for Azure AD to extend existing sessions based on information collected prior to an outage or not.
func (m *ConditionalAccessSessionControls) GetDisableResilienceDefaults()(*bool) {
    return m.disableResilienceDefaults
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ConditionalAccessSessionControls) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["applicationEnforcedRestrictions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateApplicationEnforcedRestrictionsSessionControlFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationEnforcedRestrictions(val.(ApplicationEnforcedRestrictionsSessionControlable))
        }
        return nil
    }
    res["cloudAppSecurity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCloudAppSecuritySessionControlFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCloudAppSecurity(val.(CloudAppSecuritySessionControlable))
        }
        return nil
    }
    res["disableResilienceDefaults"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisableResilienceDefaults(val)
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
    res["persistentBrowser"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePersistentBrowserSessionControlFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPersistentBrowser(val.(PersistentBrowserSessionControlable))
        }
        return nil
    }
    res["signInFrequency"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSignInFrequencySessionControlFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSignInFrequency(val.(SignInFrequencySessionControlable))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ConditionalAccessSessionControls) GetOdataType()(*string) {
    return m.odataType
}
// GetPersistentBrowser gets the persistentBrowser property value. Session control to define whether to persist cookies or not. All apps should be selected for this session control to work correctly.
func (m *ConditionalAccessSessionControls) GetPersistentBrowser()(PersistentBrowserSessionControlable) {
    return m.persistentBrowser
}
// GetSignInFrequency gets the signInFrequency property value. Session control to enforce signin frequency.
func (m *ConditionalAccessSessionControls) GetSignInFrequency()(SignInFrequencySessionControlable) {
    return m.signInFrequency
}
// Serialize serializes information the current object
func (m *ConditionalAccessSessionControls) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("applicationEnforcedRestrictions", m.GetApplicationEnforcedRestrictions())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("cloudAppSecurity", m.GetCloudAppSecurity())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("disableResilienceDefaults", m.GetDisableResilienceDefaults())
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
        err := writer.WriteObjectValue("persistentBrowser", m.GetPersistentBrowser())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("signInFrequency", m.GetSignInFrequency())
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
func (m *ConditionalAccessSessionControls) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetApplicationEnforcedRestrictions sets the applicationEnforcedRestrictions property value. Session control to enforce application restrictions. Only Exchange Online and Sharepoint Online support this session control.
func (m *ConditionalAccessSessionControls) SetApplicationEnforcedRestrictions(value ApplicationEnforcedRestrictionsSessionControlable)() {
    m.applicationEnforcedRestrictions = value
}
// SetCloudAppSecurity sets the cloudAppSecurity property value. Session control to apply cloud app security.
func (m *ConditionalAccessSessionControls) SetCloudAppSecurity(value CloudAppSecuritySessionControlable)() {
    m.cloudAppSecurity = value
}
// SetDisableResilienceDefaults sets the disableResilienceDefaults property value. Session control that determines whether it is acceptable for Azure AD to extend existing sessions based on information collected prior to an outage or not.
func (m *ConditionalAccessSessionControls) SetDisableResilienceDefaults(value *bool)() {
    m.disableResilienceDefaults = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ConditionalAccessSessionControls) SetOdataType(value *string)() {
    m.odataType = value
}
// SetPersistentBrowser sets the persistentBrowser property value. Session control to define whether to persist cookies or not. All apps should be selected for this session control to work correctly.
func (m *ConditionalAccessSessionControls) SetPersistentBrowser(value PersistentBrowserSessionControlable)() {
    m.persistentBrowser = value
}
// SetSignInFrequency sets the signInFrequency property value. Session control to enforce signin frequency.
func (m *ConditionalAccessSessionControls) SetSignInFrequency(value SignInFrequencySessionControlable)() {
    m.signInFrequency = value
}
