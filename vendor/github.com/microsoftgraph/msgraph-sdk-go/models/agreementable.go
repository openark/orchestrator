package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Agreementable 
type Agreementable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAcceptances()([]AgreementAcceptanceable)
    GetDisplayName()(*string)
    GetFile()(AgreementFileable)
    GetFiles()([]AgreementFileLocalizationable)
    GetIsPerDeviceAcceptanceRequired()(*bool)
    GetIsViewingBeforeAcceptanceRequired()(*bool)
    GetTermsExpiration()(TermsExpirationable)
    GetUserReacceptRequiredFrequency()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    SetAcceptances(value []AgreementAcceptanceable)()
    SetDisplayName(value *string)()
    SetFile(value AgreementFileable)()
    SetFiles(value []AgreementFileLocalizationable)()
    SetIsPerDeviceAcceptanceRequired(value *bool)()
    SetIsViewingBeforeAcceptanceRequired(value *bool)()
    SetTermsExpiration(value TermsExpirationable)()
    SetUserReacceptRequiredFrequency(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)()
}
