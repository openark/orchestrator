package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OfferShiftRequest 
type OfferShiftRequest struct {
    ScheduleChangeRequest
    // The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    recipientActionDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Custom message sent by recipient of the offer shift request.
    recipientActionMessage *string
    // User ID of the recipient of the offer shift request.
    recipientUserId *string
    // User ID of the sender of the offer shift request.
    senderShiftId *string
}
// NewOfferShiftRequest instantiates a new OfferShiftRequest and sets the default values.
func NewOfferShiftRequest()(*OfferShiftRequest) {
    m := &OfferShiftRequest{
        ScheduleChangeRequest: *NewScheduleChangeRequest(),
    }
    odataTypeValue := "#microsoft.graph.offerShiftRequest"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateOfferShiftRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOfferShiftRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.swapShiftsChangeRequest":
                        return NewSwapShiftsChangeRequest(), nil
                }
            }
        }
    }
    return NewOfferShiftRequest(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OfferShiftRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ScheduleChangeRequest.GetFieldDeserializers()
    res["recipientActionDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecipientActionDateTime(val)
        }
        return nil
    }
    res["recipientActionMessage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecipientActionMessage(val)
        }
        return nil
    }
    res["recipientUserId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecipientUserId(val)
        }
        return nil
    }
    res["senderShiftId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSenderShiftId(val)
        }
        return nil
    }
    return res
}
// GetRecipientActionDateTime gets the recipientActionDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *OfferShiftRequest) GetRecipientActionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.recipientActionDateTime
}
// GetRecipientActionMessage gets the recipientActionMessage property value. Custom message sent by recipient of the offer shift request.
func (m *OfferShiftRequest) GetRecipientActionMessage()(*string) {
    return m.recipientActionMessage
}
// GetRecipientUserId gets the recipientUserId property value. User ID of the recipient of the offer shift request.
func (m *OfferShiftRequest) GetRecipientUserId()(*string) {
    return m.recipientUserId
}
// GetSenderShiftId gets the senderShiftId property value. User ID of the sender of the offer shift request.
func (m *OfferShiftRequest) GetSenderShiftId()(*string) {
    return m.senderShiftId
}
// Serialize serializes information the current object
func (m *OfferShiftRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ScheduleChangeRequest.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("recipientActionMessage", m.GetRecipientActionMessage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("recipientUserId", m.GetRecipientUserId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("senderShiftId", m.GetSenderShiftId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetRecipientActionDateTime sets the recipientActionDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *OfferShiftRequest) SetRecipientActionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.recipientActionDateTime = value
}
// SetRecipientActionMessage sets the recipientActionMessage property value. Custom message sent by recipient of the offer shift request.
func (m *OfferShiftRequest) SetRecipientActionMessage(value *string)() {
    m.recipientActionMessage = value
}
// SetRecipientUserId sets the recipientUserId property value. User ID of the recipient of the offer shift request.
func (m *OfferShiftRequest) SetRecipientUserId(value *string)() {
    m.recipientUserId = value
}
// SetSenderShiftId sets the senderShiftId property value. User ID of the sender of the offer shift request.
func (m *OfferShiftRequest) SetSenderShiftId(value *string)() {
    m.senderShiftId = value
}
