package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CrossTenantAccessPolicyTargetConfiguration 
type CrossTenantAccessPolicyTargetConfiguration struct {
    // Defines whether access is allowed or blocked. The possible values are: allowed, blocked, unknownFutureValue.
    accessType *CrossTenantAccessPolicyTargetConfigurationAccessType
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The OdataType property
    odataType *string
    // Specifies whether to target users, groups, or applications with this rule.
    targets []CrossTenantAccessPolicyTargetable
}
// NewCrossTenantAccessPolicyTargetConfiguration instantiates a new crossTenantAccessPolicyTargetConfiguration and sets the default values.
func NewCrossTenantAccessPolicyTargetConfiguration()(*CrossTenantAccessPolicyTargetConfiguration) {
    m := &CrossTenantAccessPolicyTargetConfiguration{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCrossTenantAccessPolicyTargetConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCrossTenantAccessPolicyTargetConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCrossTenantAccessPolicyTargetConfiguration(), nil
}
// GetAccessType gets the accessType property value. Defines whether access is allowed or blocked. The possible values are: allowed, blocked, unknownFutureValue.
func (m *CrossTenantAccessPolicyTargetConfiguration) GetAccessType()(*CrossTenantAccessPolicyTargetConfigurationAccessType) {
    return m.accessType
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CrossTenantAccessPolicyTargetConfiguration) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CrossTenantAccessPolicyTargetConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["accessType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCrossTenantAccessPolicyTargetConfigurationAccessType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessType(val.(*CrossTenantAccessPolicyTargetConfigurationAccessType))
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
    res["targets"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCrossTenantAccessPolicyTargetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CrossTenantAccessPolicyTargetable, len(val))
            for i, v := range val {
                res[i] = v.(CrossTenantAccessPolicyTargetable)
            }
            m.SetTargets(res)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *CrossTenantAccessPolicyTargetConfiguration) GetOdataType()(*string) {
    return m.odataType
}
// GetTargets gets the targets property value. Specifies whether to target users, groups, or applications with this rule.
func (m *CrossTenantAccessPolicyTargetConfiguration) GetTargets()([]CrossTenantAccessPolicyTargetable) {
    return m.targets
}
// Serialize serializes information the current object
func (m *CrossTenantAccessPolicyTargetConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAccessType() != nil {
        cast := (*m.GetAccessType()).String()
        err := writer.WriteStringValue("accessType", &cast)
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
    if m.GetTargets() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTargets()))
        for i, v := range m.GetTargets() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("targets", cast)
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
// SetAccessType sets the accessType property value. Defines whether access is allowed or blocked. The possible values are: allowed, blocked, unknownFutureValue.
func (m *CrossTenantAccessPolicyTargetConfiguration) SetAccessType(value *CrossTenantAccessPolicyTargetConfigurationAccessType)() {
    m.accessType = value
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CrossTenantAccessPolicyTargetConfiguration) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *CrossTenantAccessPolicyTargetConfiguration) SetOdataType(value *string)() {
    m.odataType = value
}
// SetTargets sets the targets property value. Specifies whether to target users, groups, or applications with this rule.
func (m *CrossTenantAccessPolicyTargetConfiguration) SetTargets(value []CrossTenantAccessPolicyTargetable)() {
    m.targets = value
}
