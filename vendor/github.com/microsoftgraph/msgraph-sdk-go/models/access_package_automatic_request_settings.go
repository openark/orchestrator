package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccessPackageAutomaticRequestSettings 
type AccessPackageAutomaticRequestSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The gracePeriodBeforeAccessRemoval property
    gracePeriodBeforeAccessRemoval *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration
    // The OdataType property
    odataType *string
    // The removeAccessWhenTargetLeavesAllowedTargets property
    removeAccessWhenTargetLeavesAllowedTargets *bool
    // If set to true, automatic assignments will be created for targets in the allowed target scope.
    requestAccessForAllowedTargets *bool
}
// NewAccessPackageAutomaticRequestSettings instantiates a new accessPackageAutomaticRequestSettings and sets the default values.
func NewAccessPackageAutomaticRequestSettings()(*AccessPackageAutomaticRequestSettings) {
    m := &AccessPackageAutomaticRequestSettings{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAccessPackageAutomaticRequestSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAccessPackageAutomaticRequestSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAccessPackageAutomaticRequestSettings(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AccessPackageAutomaticRequestSettings) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AccessPackageAutomaticRequestSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["gracePeriodBeforeAccessRemoval"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGracePeriodBeforeAccessRemoval(val)
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
    res["removeAccessWhenTargetLeavesAllowedTargets"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemoveAccessWhenTargetLeavesAllowedTargets(val)
        }
        return nil
    }
    res["requestAccessForAllowedTargets"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestAccessForAllowedTargets(val)
        }
        return nil
    }
    return res
}
// GetGracePeriodBeforeAccessRemoval gets the gracePeriodBeforeAccessRemoval property value. The gracePeriodBeforeAccessRemoval property
func (m *AccessPackageAutomaticRequestSettings) GetGracePeriodBeforeAccessRemoval()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    return m.gracePeriodBeforeAccessRemoval
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AccessPackageAutomaticRequestSettings) GetOdataType()(*string) {
    return m.odataType
}
// GetRemoveAccessWhenTargetLeavesAllowedTargets gets the removeAccessWhenTargetLeavesAllowedTargets property value. The removeAccessWhenTargetLeavesAllowedTargets property
func (m *AccessPackageAutomaticRequestSettings) GetRemoveAccessWhenTargetLeavesAllowedTargets()(*bool) {
    return m.removeAccessWhenTargetLeavesAllowedTargets
}
// GetRequestAccessForAllowedTargets gets the requestAccessForAllowedTargets property value. If set to true, automatic assignments will be created for targets in the allowed target scope.
func (m *AccessPackageAutomaticRequestSettings) GetRequestAccessForAllowedTargets()(*bool) {
    return m.requestAccessForAllowedTargets
}
// Serialize serializes information the current object
func (m *AccessPackageAutomaticRequestSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteISODurationValue("gracePeriodBeforeAccessRemoval", m.GetGracePeriodBeforeAccessRemoval())
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
        err := writer.WriteBoolValue("removeAccessWhenTargetLeavesAllowedTargets", m.GetRemoveAccessWhenTargetLeavesAllowedTargets())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("requestAccessForAllowedTargets", m.GetRequestAccessForAllowedTargets())
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
func (m *AccessPackageAutomaticRequestSettings) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetGracePeriodBeforeAccessRemoval sets the gracePeriodBeforeAccessRemoval property value. The gracePeriodBeforeAccessRemoval property
func (m *AccessPackageAutomaticRequestSettings) SetGracePeriodBeforeAccessRemoval(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    m.gracePeriodBeforeAccessRemoval = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AccessPackageAutomaticRequestSettings) SetOdataType(value *string)() {
    m.odataType = value
}
// SetRemoveAccessWhenTargetLeavesAllowedTargets sets the removeAccessWhenTargetLeavesAllowedTargets property value. The removeAccessWhenTargetLeavesAllowedTargets property
func (m *AccessPackageAutomaticRequestSettings) SetRemoveAccessWhenTargetLeavesAllowedTargets(value *bool)() {
    m.removeAccessWhenTargetLeavesAllowedTargets = value
}
// SetRequestAccessForAllowedTargets sets the requestAccessForAllowedTargets property value. If set to true, automatic assignments will be created for targets in the allowed target scope.
func (m *AccessPackageAutomaticRequestSettings) SetRequestAccessForAllowedTargets(value *bool)() {
    m.requestAccessForAllowedTargets = value
}
