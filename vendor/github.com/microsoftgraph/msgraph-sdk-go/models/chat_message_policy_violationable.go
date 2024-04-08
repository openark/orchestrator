package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ChatMessagePolicyViolationable 
type ChatMessagePolicyViolationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDlpAction()(*ChatMessagePolicyViolationDlpActionTypes)
    GetJustificationText()(*string)
    GetOdataType()(*string)
    GetPolicyTip()(ChatMessagePolicyViolationPolicyTipable)
    GetUserAction()(*ChatMessagePolicyViolationUserActionTypes)
    GetVerdictDetails()(*ChatMessagePolicyViolationVerdictDetailsTypes)
    SetDlpAction(value *ChatMessagePolicyViolationDlpActionTypes)()
    SetJustificationText(value *string)()
    SetOdataType(value *string)()
    SetPolicyTip(value ChatMessagePolicyViolationPolicyTipable)()
    SetUserAction(value *ChatMessagePolicyViolationUserActionTypes)()
    SetVerdictDetails(value *ChatMessagePolicyViolationVerdictDetailsTypes)()
}
