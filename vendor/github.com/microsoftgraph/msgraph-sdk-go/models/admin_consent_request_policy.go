package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AdminConsentRequestPolicy 
type AdminConsentRequestPolicy struct {
    Entity
    // Specifies whether the admin consent request feature is enabled or disabled. Required.
    isEnabled *bool
    // Specifies whether reviewers will receive notifications. Required.
    notifyReviewers *bool
    // Specifies whether reviewers will receive reminder emails. Required.
    remindersEnabled *bool
    // Specifies the duration the request is active before it automatically expires if no decision is applied.
    requestDurationInDays *int32
    // The list of reviewers for the admin consent. Required.
    reviewers []AccessReviewReviewerScopeable
    // Specifies the version of this policy. When the policy is updated, this version is updated. Read-only.
    version *int32
}
// NewAdminConsentRequestPolicy instantiates a new adminConsentRequestPolicy and sets the default values.
func NewAdminConsentRequestPolicy()(*AdminConsentRequestPolicy) {
    m := &AdminConsentRequestPolicy{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAdminConsentRequestPolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAdminConsentRequestPolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAdminConsentRequestPolicy(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AdminConsentRequestPolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["isEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsEnabled(val)
        }
        return nil
    }
    res["notifyReviewers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotifyReviewers(val)
        }
        return nil
    }
    res["remindersEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemindersEnabled(val)
        }
        return nil
    }
    res["requestDurationInDays"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestDurationInDays(val)
        }
        return nil
    }
    res["reviewers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAccessReviewReviewerScopeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessReviewReviewerScopeable, len(val))
            for i, v := range val {
                res[i] = v.(AccessReviewReviewerScopeable)
            }
            m.SetReviewers(res)
        }
        return nil
    }
    res["version"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersion(val)
        }
        return nil
    }
    return res
}
// GetIsEnabled gets the isEnabled property value. Specifies whether the admin consent request feature is enabled or disabled. Required.
func (m *AdminConsentRequestPolicy) GetIsEnabled()(*bool) {
    return m.isEnabled
}
// GetNotifyReviewers gets the notifyReviewers property value. Specifies whether reviewers will receive notifications. Required.
func (m *AdminConsentRequestPolicy) GetNotifyReviewers()(*bool) {
    return m.notifyReviewers
}
// GetRemindersEnabled gets the remindersEnabled property value. Specifies whether reviewers will receive reminder emails. Required.
func (m *AdminConsentRequestPolicy) GetRemindersEnabled()(*bool) {
    return m.remindersEnabled
}
// GetRequestDurationInDays gets the requestDurationInDays property value. Specifies the duration the request is active before it automatically expires if no decision is applied.
func (m *AdminConsentRequestPolicy) GetRequestDurationInDays()(*int32) {
    return m.requestDurationInDays
}
// GetReviewers gets the reviewers property value. The list of reviewers for the admin consent. Required.
func (m *AdminConsentRequestPolicy) GetReviewers()([]AccessReviewReviewerScopeable) {
    return m.reviewers
}
// GetVersion gets the version property value. Specifies the version of this policy. When the policy is updated, this version is updated. Read-only.
func (m *AdminConsentRequestPolicy) GetVersion()(*int32) {
    return m.version
}
// Serialize serializes information the current object
func (m *AdminConsentRequestPolicy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("isEnabled", m.GetIsEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("notifyReviewers", m.GetNotifyReviewers())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("remindersEnabled", m.GetRemindersEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("requestDurationInDays", m.GetRequestDurationInDays())
        if err != nil {
            return err
        }
    }
    if m.GetReviewers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetReviewers()))
        for i, v := range m.GetReviewers() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("reviewers", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("version", m.GetVersion())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetIsEnabled sets the isEnabled property value. Specifies whether the admin consent request feature is enabled or disabled. Required.
func (m *AdminConsentRequestPolicy) SetIsEnabled(value *bool)() {
    m.isEnabled = value
}
// SetNotifyReviewers sets the notifyReviewers property value. Specifies whether reviewers will receive notifications. Required.
func (m *AdminConsentRequestPolicy) SetNotifyReviewers(value *bool)() {
    m.notifyReviewers = value
}
// SetRemindersEnabled sets the remindersEnabled property value. Specifies whether reviewers will receive reminder emails. Required.
func (m *AdminConsentRequestPolicy) SetRemindersEnabled(value *bool)() {
    m.remindersEnabled = value
}
// SetRequestDurationInDays sets the requestDurationInDays property value. Specifies the duration the request is active before it automatically expires if no decision is applied.
func (m *AdminConsentRequestPolicy) SetRequestDurationInDays(value *int32)() {
    m.requestDurationInDays = value
}
// SetReviewers sets the reviewers property value. The list of reviewers for the admin consent. Required.
func (m *AdminConsentRequestPolicy) SetReviewers(value []AccessReviewReviewerScopeable)() {
    m.reviewers = value
}
// SetVersion sets the version property value. Specifies the version of this policy. When the policy is updated, this version is updated. Read-only.
func (m *AdminConsentRequestPolicy) SetVersion(value *int32)() {
    m.version = value
}
