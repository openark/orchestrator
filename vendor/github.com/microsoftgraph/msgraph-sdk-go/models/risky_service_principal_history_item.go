package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RiskyServicePrincipalHistoryItem 
type RiskyServicePrincipalHistoryItem struct {
    RiskyServicePrincipal
    // The activity related to service principal risk level change.
    activity RiskServicePrincipalActivityable
    // The identifier of the actor of the operation.
    initiatedBy *string
}
// NewRiskyServicePrincipalHistoryItem instantiates a new riskyServicePrincipalHistoryItem and sets the default values.
func NewRiskyServicePrincipalHistoryItem()(*RiskyServicePrincipalHistoryItem) {
    m := &RiskyServicePrincipalHistoryItem{
        RiskyServicePrincipal: *NewRiskyServicePrincipal(),
    }
    return m
}
// CreateRiskyServicePrincipalHistoryItemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRiskyServicePrincipalHistoryItemFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRiskyServicePrincipalHistoryItem(), nil
}
// GetActivity gets the activity property value. The activity related to service principal risk level change.
func (m *RiskyServicePrincipalHistoryItem) GetActivity()(RiskServicePrincipalActivityable) {
    return m.activity
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RiskyServicePrincipalHistoryItem) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.RiskyServicePrincipal.GetFieldDeserializers()
    res["activity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRiskServicePrincipalActivityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActivity(val.(RiskServicePrincipalActivityable))
        }
        return nil
    }
    res["initiatedBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInitiatedBy(val)
        }
        return nil
    }
    return res
}
// GetInitiatedBy gets the initiatedBy property value. The identifier of the actor of the operation.
func (m *RiskyServicePrincipalHistoryItem) GetInitiatedBy()(*string) {
    return m.initiatedBy
}
// Serialize serializes information the current object
func (m *RiskyServicePrincipalHistoryItem) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.RiskyServicePrincipal.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("activity", m.GetActivity())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("initiatedBy", m.GetInitiatedBy())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActivity sets the activity property value. The activity related to service principal risk level change.
func (m *RiskyServicePrincipalHistoryItem) SetActivity(value RiskServicePrincipalActivityable)() {
    m.activity = value
}
// SetInitiatedBy sets the initiatedBy property value. The identifier of the actor of the operation.
func (m *RiskyServicePrincipalHistoryItem) SetInitiatedBy(value *string)() {
    m.initiatedBy = value
}
