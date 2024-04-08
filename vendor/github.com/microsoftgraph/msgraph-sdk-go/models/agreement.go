package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Agreement 
type Agreement struct {
    Entity
    // Read-only. Information about acceptances of this agreement.
    acceptances []AgreementAcceptanceable
    // Display name of the agreement. The display name is used for internal tracking of the agreement but is not shown to end users who view the agreement. Supports $filter (eq).
    displayName *string
    // Default PDF linked to this agreement.
    file AgreementFileable
    // PDFs linked to this agreement. This property is in the process of being deprecated. Use the  file property instead. Supports $expand.
    files []AgreementFileLocalizationable
    // Indicates whether end users are required to accept this agreement on every device that they access it from. The end user is required to register their device in Azure AD, if they haven't already done so. Supports $filter (eq).
    isPerDeviceAcceptanceRequired *bool
    // Indicates whether the user has to expand the agreement before accepting. Supports $filter (eq).
    isViewingBeforeAcceptanceRequired *bool
    // Expiration schedule and frequency of agreement for all users. Supports $filter (eq).
    termsExpiration TermsExpirationable
    // The duration after which the user must re-accept the terms of use. The value is represented in ISO 8601 format for durations. Supports $filter (eq).
    userReacceptRequiredFrequency *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration
}
// NewAgreement instantiates a new agreement and sets the default values.
func NewAgreement()(*Agreement) {
    m := &Agreement{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAgreementFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAgreementFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAgreement(), nil
}
// GetAcceptances gets the acceptances property value. Read-only. Information about acceptances of this agreement.
func (m *Agreement) GetAcceptances()([]AgreementAcceptanceable) {
    return m.acceptances
}
// GetDisplayName gets the displayName property value. Display name of the agreement. The display name is used for internal tracking of the agreement but is not shown to end users who view the agreement. Supports $filter (eq).
func (m *Agreement) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Agreement) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["acceptances"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAgreementAcceptanceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AgreementAcceptanceable, len(val))
            for i, v := range val {
                res[i] = v.(AgreementAcceptanceable)
            }
            m.SetAcceptances(res)
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
    res["file"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAgreementFileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFile(val.(AgreementFileable))
        }
        return nil
    }
    res["files"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAgreementFileLocalizationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AgreementFileLocalizationable, len(val))
            for i, v := range val {
                res[i] = v.(AgreementFileLocalizationable)
            }
            m.SetFiles(res)
        }
        return nil
    }
    res["isPerDeviceAcceptanceRequired"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsPerDeviceAcceptanceRequired(val)
        }
        return nil
    }
    res["isViewingBeforeAcceptanceRequired"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsViewingBeforeAcceptanceRequired(val)
        }
        return nil
    }
    res["termsExpiration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTermsExpirationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTermsExpiration(val.(TermsExpirationable))
        }
        return nil
    }
    res["userReacceptRequiredFrequency"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserReacceptRequiredFrequency(val)
        }
        return nil
    }
    return res
}
// GetFile gets the file property value. Default PDF linked to this agreement.
func (m *Agreement) GetFile()(AgreementFileable) {
    return m.file
}
// GetFiles gets the files property value. PDFs linked to this agreement. This property is in the process of being deprecated. Use the  file property instead. Supports $expand.
func (m *Agreement) GetFiles()([]AgreementFileLocalizationable) {
    return m.files
}
// GetIsPerDeviceAcceptanceRequired gets the isPerDeviceAcceptanceRequired property value. Indicates whether end users are required to accept this agreement on every device that they access it from. The end user is required to register their device in Azure AD, if they haven't already done so. Supports $filter (eq).
func (m *Agreement) GetIsPerDeviceAcceptanceRequired()(*bool) {
    return m.isPerDeviceAcceptanceRequired
}
// GetIsViewingBeforeAcceptanceRequired gets the isViewingBeforeAcceptanceRequired property value. Indicates whether the user has to expand the agreement before accepting. Supports $filter (eq).
func (m *Agreement) GetIsViewingBeforeAcceptanceRequired()(*bool) {
    return m.isViewingBeforeAcceptanceRequired
}
// GetTermsExpiration gets the termsExpiration property value. Expiration schedule and frequency of agreement for all users. Supports $filter (eq).
func (m *Agreement) GetTermsExpiration()(TermsExpirationable) {
    return m.termsExpiration
}
// GetUserReacceptRequiredFrequency gets the userReacceptRequiredFrequency property value. The duration after which the user must re-accept the terms of use. The value is represented in ISO 8601 format for durations. Supports $filter (eq).
func (m *Agreement) GetUserReacceptRequiredFrequency()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    return m.userReacceptRequiredFrequency
}
// Serialize serializes information the current object
func (m *Agreement) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAcceptances() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAcceptances()))
        for i, v := range m.GetAcceptances() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("acceptances", cast)
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
        err = writer.WriteObjectValue("file", m.GetFile())
        if err != nil {
            return err
        }
    }
    if m.GetFiles() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetFiles()))
        for i, v := range m.GetFiles() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("files", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isPerDeviceAcceptanceRequired", m.GetIsPerDeviceAcceptanceRequired())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isViewingBeforeAcceptanceRequired", m.GetIsViewingBeforeAcceptanceRequired())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("termsExpiration", m.GetTermsExpiration())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteISODurationValue("userReacceptRequiredFrequency", m.GetUserReacceptRequiredFrequency())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAcceptances sets the acceptances property value. Read-only. Information about acceptances of this agreement.
func (m *Agreement) SetAcceptances(value []AgreementAcceptanceable)() {
    m.acceptances = value
}
// SetDisplayName sets the displayName property value. Display name of the agreement. The display name is used for internal tracking of the agreement but is not shown to end users who view the agreement. Supports $filter (eq).
func (m *Agreement) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetFile sets the file property value. Default PDF linked to this agreement.
func (m *Agreement) SetFile(value AgreementFileable)() {
    m.file = value
}
// SetFiles sets the files property value. PDFs linked to this agreement. This property is in the process of being deprecated. Use the  file property instead. Supports $expand.
func (m *Agreement) SetFiles(value []AgreementFileLocalizationable)() {
    m.files = value
}
// SetIsPerDeviceAcceptanceRequired sets the isPerDeviceAcceptanceRequired property value. Indicates whether end users are required to accept this agreement on every device that they access it from. The end user is required to register their device in Azure AD, if they haven't already done so. Supports $filter (eq).
func (m *Agreement) SetIsPerDeviceAcceptanceRequired(value *bool)() {
    m.isPerDeviceAcceptanceRequired = value
}
// SetIsViewingBeforeAcceptanceRequired sets the isViewingBeforeAcceptanceRequired property value. Indicates whether the user has to expand the agreement before accepting. Supports $filter (eq).
func (m *Agreement) SetIsViewingBeforeAcceptanceRequired(value *bool)() {
    m.isViewingBeforeAcceptanceRequired = value
}
// SetTermsExpiration sets the termsExpiration property value. Expiration schedule and frequency of agreement for all users. Supports $filter (eq).
func (m *Agreement) SetTermsExpiration(value TermsExpirationable)() {
    m.termsExpiration = value
}
// SetUserReacceptRequiredFrequency sets the userReacceptRequiredFrequency property value. The duration after which the user must re-accept the terms of use. The value is represented in ISO 8601 format for durations. Supports $filter (eq).
func (m *Agreement) SetUserReacceptRequiredFrequency(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    m.userReacceptRequiredFrequency = value
}
