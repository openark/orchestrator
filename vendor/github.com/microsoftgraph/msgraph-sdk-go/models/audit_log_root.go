package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AuditLogRoot 
type AuditLogRoot struct {
    Entity
    // The directoryAudits property
    directoryAudits []DirectoryAuditable
    // The provisioning property
    provisioning []ProvisioningObjectSummaryable
    // The signIns property
    signIns []SignInable
}
// NewAuditLogRoot instantiates a new AuditLogRoot and sets the default values.
func NewAuditLogRoot()(*AuditLogRoot) {
    m := &AuditLogRoot{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAuditLogRootFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAuditLogRootFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAuditLogRoot(), nil
}
// GetDirectoryAudits gets the directoryAudits property value. The directoryAudits property
func (m *AuditLogRoot) GetDirectoryAudits()([]DirectoryAuditable) {
    return m.directoryAudits
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AuditLogRoot) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["directoryAudits"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDirectoryAuditFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DirectoryAuditable, len(val))
            for i, v := range val {
                res[i] = v.(DirectoryAuditable)
            }
            m.SetDirectoryAudits(res)
        }
        return nil
    }
    res["provisioning"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateProvisioningObjectSummaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ProvisioningObjectSummaryable, len(val))
            for i, v := range val {
                res[i] = v.(ProvisioningObjectSummaryable)
            }
            m.SetProvisioning(res)
        }
        return nil
    }
    res["signIns"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSignInFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SignInable, len(val))
            for i, v := range val {
                res[i] = v.(SignInable)
            }
            m.SetSignIns(res)
        }
        return nil
    }
    return res
}
// GetProvisioning gets the provisioning property value. The provisioning property
func (m *AuditLogRoot) GetProvisioning()([]ProvisioningObjectSummaryable) {
    return m.provisioning
}
// GetSignIns gets the signIns property value. The signIns property
func (m *AuditLogRoot) GetSignIns()([]SignInable) {
    return m.signIns
}
// Serialize serializes information the current object
func (m *AuditLogRoot) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetDirectoryAudits() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDirectoryAudits()))
        for i, v := range m.GetDirectoryAudits() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("directoryAudits", cast)
        if err != nil {
            return err
        }
    }
    if m.GetProvisioning() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetProvisioning()))
        for i, v := range m.GetProvisioning() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("provisioning", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSignIns() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSignIns()))
        for i, v := range m.GetSignIns() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("signIns", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDirectoryAudits sets the directoryAudits property value. The directoryAudits property
func (m *AuditLogRoot) SetDirectoryAudits(value []DirectoryAuditable)() {
    m.directoryAudits = value
}
// SetProvisioning sets the provisioning property value. The provisioning property
func (m *AuditLogRoot) SetProvisioning(value []ProvisioningObjectSummaryable)() {
    m.provisioning = value
}
// SetSignIns sets the signIns property value. The signIns property
func (m *AuditLogRoot) SetSignIns(value []SignInable)() {
    m.signIns = value
}
