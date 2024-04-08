package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WorkforceIntegration 
type WorkforceIntegration struct {
    ChangeTrackedEntity
    // API version for the call back URL. Start with 1.
    apiVersion *int32
    // Name of the workforce integration.
    displayName *string
    // The workforce integration encryption resource.
    encryption WorkforceIntegrationEncryptionable
    // Indicates whether this workforce integration is currently active and available.
    isActive *bool
    // The Shifts entities supported for synchronous change notifications. Shifts will make a call back to the url provided on client changes on those entities added here. By default, no entities are supported for change notifications. Possible values are: none, shift, swapRequest, userShiftPreferences, openshift, openShiftRequest, offerShiftRequest, unknownFutureValue.
    supportedEntities *WorkforceIntegrationSupportedEntities
    // Workforce Integration URL for callbacks from the Shifts service.
    url *string
}
// NewWorkforceIntegration instantiates a new WorkforceIntegration and sets the default values.
func NewWorkforceIntegration()(*WorkforceIntegration) {
    m := &WorkforceIntegration{
        ChangeTrackedEntity: *NewChangeTrackedEntity(),
    }
    odataTypeValue := "#microsoft.graph.workforceIntegration"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateWorkforceIntegrationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWorkforceIntegrationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWorkforceIntegration(), nil
}
// GetApiVersion gets the apiVersion property value. API version for the call back URL. Start with 1.
func (m *WorkforceIntegration) GetApiVersion()(*int32) {
    return m.apiVersion
}
// GetDisplayName gets the displayName property value. Name of the workforce integration.
func (m *WorkforceIntegration) GetDisplayName()(*string) {
    return m.displayName
}
// GetEncryption gets the encryption property value. The workforce integration encryption resource.
func (m *WorkforceIntegration) GetEncryption()(WorkforceIntegrationEncryptionable) {
    return m.encryption
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WorkforceIntegration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ChangeTrackedEntity.GetFieldDeserializers()
    res["apiVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApiVersion(val)
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
    res["encryption"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWorkforceIntegrationEncryptionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEncryption(val.(WorkforceIntegrationEncryptionable))
        }
        return nil
    }
    res["isActive"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsActive(val)
        }
        return nil
    }
    res["supportedEntities"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWorkforceIntegrationSupportedEntities)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSupportedEntities(val.(*WorkforceIntegrationSupportedEntities))
        }
        return nil
    }
    res["url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUrl(val)
        }
        return nil
    }
    return res
}
// GetIsActive gets the isActive property value. Indicates whether this workforce integration is currently active and available.
func (m *WorkforceIntegration) GetIsActive()(*bool) {
    return m.isActive
}
// GetSupportedEntities gets the supportedEntities property value. The Shifts entities supported for synchronous change notifications. Shifts will make a call back to the url provided on client changes on those entities added here. By default, no entities are supported for change notifications. Possible values are: none, shift, swapRequest, userShiftPreferences, openshift, openShiftRequest, offerShiftRequest, unknownFutureValue.
func (m *WorkforceIntegration) GetSupportedEntities()(*WorkforceIntegrationSupportedEntities) {
    return m.supportedEntities
}
// GetUrl gets the url property value. Workforce Integration URL for callbacks from the Shifts service.
func (m *WorkforceIntegration) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *WorkforceIntegration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ChangeTrackedEntity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("apiVersion", m.GetApiVersion())
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
    {
        err = writer.WriteObjectValue("encryption", m.GetEncryption())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isActive", m.GetIsActive())
        if err != nil {
            return err
        }
    }
    if m.GetSupportedEntities() != nil {
        cast := (*m.GetSupportedEntities()).String()
        err = writer.WriteStringValue("supportedEntities", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("url", m.GetUrl())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetApiVersion sets the apiVersion property value. API version for the call back URL. Start with 1.
func (m *WorkforceIntegration) SetApiVersion(value *int32)() {
    m.apiVersion = value
}
// SetDisplayName sets the displayName property value. Name of the workforce integration.
func (m *WorkforceIntegration) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetEncryption sets the encryption property value. The workforce integration encryption resource.
func (m *WorkforceIntegration) SetEncryption(value WorkforceIntegrationEncryptionable)() {
    m.encryption = value
}
// SetIsActive sets the isActive property value. Indicates whether this workforce integration is currently active and available.
func (m *WorkforceIntegration) SetIsActive(value *bool)() {
    m.isActive = value
}
// SetSupportedEntities sets the supportedEntities property value. The Shifts entities supported for synchronous change notifications. Shifts will make a call back to the url provided on client changes on those entities added here. By default, no entities are supported for change notifications. Possible values are: none, shift, swapRequest, userShiftPreferences, openshift, openShiftRequest, offerShiftRequest, unknownFutureValue.
func (m *WorkforceIntegration) SetSupportedEntities(value *WorkforceIntegrationSupportedEntities)() {
    m.supportedEntities = value
}
// SetUrl sets the url property value. Workforce Integration URL for callbacks from the Shifts service.
func (m *WorkforceIntegration) SetUrl(value *string)() {
    m.url = value
}
