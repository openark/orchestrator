package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TermsAndConditionsAcceptanceStatus a termsAndConditionsAcceptanceStatus entity represents the acceptance status of a given Terms and Conditions (T&C) policy by a given user. Users must accept the most up-to-date version of the terms in order to retain access to the Company Portal.
type TermsAndConditionsAcceptanceStatus struct {
    Entity
    // DateTime when the terms were last accepted by the user.
    acceptedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Most recent version number of the T&C accepted by the user.
    acceptedVersion *int32
    // Navigation link to the terms and conditions that are assigned.
    termsAndConditions TermsAndConditionsable
    // Display name of the user whose acceptance the entity represents.
    userDisplayName *string
    // The userPrincipalName of the User that accepted the term.
    userPrincipalName *string
}
// NewTermsAndConditionsAcceptanceStatus instantiates a new termsAndConditionsAcceptanceStatus and sets the default values.
func NewTermsAndConditionsAcceptanceStatus()(*TermsAndConditionsAcceptanceStatus) {
    m := &TermsAndConditionsAcceptanceStatus{
        Entity: *NewEntity(),
    }
    return m
}
// CreateTermsAndConditionsAcceptanceStatusFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTermsAndConditionsAcceptanceStatusFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTermsAndConditionsAcceptanceStatus(), nil
}
// GetAcceptedDateTime gets the acceptedDateTime property value. DateTime when the terms were last accepted by the user.
func (m *TermsAndConditionsAcceptanceStatus) GetAcceptedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.acceptedDateTime
}
// GetAcceptedVersion gets the acceptedVersion property value. Most recent version number of the T&C accepted by the user.
func (m *TermsAndConditionsAcceptanceStatus) GetAcceptedVersion()(*int32) {
    return m.acceptedVersion
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TermsAndConditionsAcceptanceStatus) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["acceptedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAcceptedDateTime(val)
        }
        return nil
    }
    res["acceptedVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAcceptedVersion(val)
        }
        return nil
    }
    res["termsAndConditions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTermsAndConditionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTermsAndConditions(val.(TermsAndConditionsable))
        }
        return nil
    }
    res["userDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserDisplayName(val)
        }
        return nil
    }
    res["userPrincipalName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserPrincipalName(val)
        }
        return nil
    }
    return res
}
// GetTermsAndConditions gets the termsAndConditions property value. Navigation link to the terms and conditions that are assigned.
func (m *TermsAndConditionsAcceptanceStatus) GetTermsAndConditions()(TermsAndConditionsable) {
    return m.termsAndConditions
}
// GetUserDisplayName gets the userDisplayName property value. Display name of the user whose acceptance the entity represents.
func (m *TermsAndConditionsAcceptanceStatus) GetUserDisplayName()(*string) {
    return m.userDisplayName
}
// GetUserPrincipalName gets the userPrincipalName property value. The userPrincipalName of the User that accepted the term.
func (m *TermsAndConditionsAcceptanceStatus) GetUserPrincipalName()(*string) {
    return m.userPrincipalName
}
// Serialize serializes information the current object
func (m *TermsAndConditionsAcceptanceStatus) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("acceptedDateTime", m.GetAcceptedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("acceptedVersion", m.GetAcceptedVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("termsAndConditions", m.GetTermsAndConditions())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userDisplayName", m.GetUserDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userPrincipalName", m.GetUserPrincipalName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAcceptedDateTime sets the acceptedDateTime property value. DateTime when the terms were last accepted by the user.
func (m *TermsAndConditionsAcceptanceStatus) SetAcceptedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.acceptedDateTime = value
}
// SetAcceptedVersion sets the acceptedVersion property value. Most recent version number of the T&C accepted by the user.
func (m *TermsAndConditionsAcceptanceStatus) SetAcceptedVersion(value *int32)() {
    m.acceptedVersion = value
}
// SetTermsAndConditions sets the termsAndConditions property value. Navigation link to the terms and conditions that are assigned.
func (m *TermsAndConditionsAcceptanceStatus) SetTermsAndConditions(value TermsAndConditionsable)() {
    m.termsAndConditions = value
}
// SetUserDisplayName sets the userDisplayName property value. Display name of the user whose acceptance the entity represents.
func (m *TermsAndConditionsAcceptanceStatus) SetUserDisplayName(value *string)() {
    m.userDisplayName = value
}
// SetUserPrincipalName sets the userPrincipalName property value. The userPrincipalName of the User that accepted the term.
func (m *TermsAndConditionsAcceptanceStatus) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
