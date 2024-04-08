package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EmailFileAssessmentRequestable 
type EmailFileAssessmentRequestable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    ThreatAssessmentRequestable
    GetContentData()(*string)
    GetDestinationRoutingReason()(*MailDestinationRoutingReason)
    GetRecipientEmail()(*string)
    SetContentData(value *string)()
    SetDestinationRoutingReason(value *MailDestinationRoutingReason)()
    SetRecipientEmail(value *string)()
}
