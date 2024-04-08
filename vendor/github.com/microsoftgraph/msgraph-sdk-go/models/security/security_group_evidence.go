package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SecurityGroupEvidence 
type SecurityGroupEvidence struct {
    AlertEvidence
    // The name of the security group.
    displayName *string
    // Unique identifier of the security group.
    securityGroupId *string
}
// NewSecurityGroupEvidence instantiates a new SecurityGroupEvidence and sets the default values.
func NewSecurityGroupEvidence()(*SecurityGroupEvidence) {
    m := &SecurityGroupEvidence{
        AlertEvidence: *NewAlertEvidence(),
    }
    return m
}
// CreateSecurityGroupEvidenceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSecurityGroupEvidenceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSecurityGroupEvidence(), nil
}
// GetDisplayName gets the displayName property value. The name of the security group.
func (m *SecurityGroupEvidence) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SecurityGroupEvidence) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AlertEvidence.GetFieldDeserializers()
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
    res["securityGroupId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecurityGroupId(val)
        }
        return nil
    }
    return res
}
// GetSecurityGroupId gets the securityGroupId property value. Unique identifier of the security group.
func (m *SecurityGroupEvidence) GetSecurityGroupId()(*string) {
    return m.securityGroupId
}
// Serialize serializes information the current object
func (m *SecurityGroupEvidence) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AlertEvidence.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("securityGroupId", m.GetSecurityGroupId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDisplayName sets the displayName property value. The name of the security group.
func (m *SecurityGroupEvidence) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetSecurityGroupId sets the securityGroupId property value. Unique identifier of the security group.
func (m *SecurityGroupEvidence) SetSecurityGroupId(value *string)() {
    m.securityGroupId = value
}
