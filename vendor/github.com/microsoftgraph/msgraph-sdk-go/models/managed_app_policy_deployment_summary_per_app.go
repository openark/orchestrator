package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ManagedAppPolicyDeploymentSummaryPerApp represents policy deployment summary per app.
type ManagedAppPolicyDeploymentSummaryPerApp struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Number of users the policy is applied.
    configurationAppliedUserCount *int32
    // Deployment of an app.
    mobileAppIdentifier MobileAppIdentifierable
    // The OdataType property
    odataType *string
}
// NewManagedAppPolicyDeploymentSummaryPerApp instantiates a new managedAppPolicyDeploymentSummaryPerApp and sets the default values.
func NewManagedAppPolicyDeploymentSummaryPerApp()(*ManagedAppPolicyDeploymentSummaryPerApp) {
    m := &ManagedAppPolicyDeploymentSummaryPerApp{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateManagedAppPolicyDeploymentSummaryPerAppFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagedAppPolicyDeploymentSummaryPerAppFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewManagedAppPolicyDeploymentSummaryPerApp(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ManagedAppPolicyDeploymentSummaryPerApp) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetConfigurationAppliedUserCount gets the configurationAppliedUserCount property value. Number of users the policy is applied.
func (m *ManagedAppPolicyDeploymentSummaryPerApp) GetConfigurationAppliedUserCount()(*int32) {
    return m.configurationAppliedUserCount
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagedAppPolicyDeploymentSummaryPerApp) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["configurationAppliedUserCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfigurationAppliedUserCount(val)
        }
        return nil
    }
    res["mobileAppIdentifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMobileAppIdentifierFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMobileAppIdentifier(val.(MobileAppIdentifierable))
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
    return res
}
// GetMobileAppIdentifier gets the mobileAppIdentifier property value. Deployment of an app.
func (m *ManagedAppPolicyDeploymentSummaryPerApp) GetMobileAppIdentifier()(MobileAppIdentifierable) {
    return m.mobileAppIdentifier
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ManagedAppPolicyDeploymentSummaryPerApp) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *ManagedAppPolicyDeploymentSummaryPerApp) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("configurationAppliedUserCount", m.GetConfigurationAppliedUserCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("mobileAppIdentifier", m.GetMobileAppIdentifier())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ManagedAppPolicyDeploymentSummaryPerApp) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetConfigurationAppliedUserCount sets the configurationAppliedUserCount property value. Number of users the policy is applied.
func (m *ManagedAppPolicyDeploymentSummaryPerApp) SetConfigurationAppliedUserCount(value *int32)() {
    m.configurationAppliedUserCount = value
}
// SetMobileAppIdentifier sets the mobileAppIdentifier property value. Deployment of an app.
func (m *ManagedAppPolicyDeploymentSummaryPerApp) SetMobileAppIdentifier(value MobileAppIdentifierable)() {
    m.mobileAppIdentifier = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ManagedAppPolicyDeploymentSummaryPerApp) SetOdataType(value *string)() {
    m.odataType = value
}
