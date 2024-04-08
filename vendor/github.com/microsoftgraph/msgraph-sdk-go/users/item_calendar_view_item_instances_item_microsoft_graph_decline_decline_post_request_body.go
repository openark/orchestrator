package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// ItemCalendarViewItemInstancesItemMicrosoftGraphDeclineDeclinePostRequestBody 
type ItemCalendarViewItemInstancesItemMicrosoftGraphDeclineDeclinePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The Comment property
    comment *string
    // The ProposedNewTime property
    proposedNewTime iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TimeSlotable
    // The SendResponse property
    sendResponse *bool
}
// NewItemCalendarViewItemInstancesItemMicrosoftGraphDeclineDeclinePostRequestBody instantiates a new ItemCalendarViewItemInstancesItemMicrosoftGraphDeclineDeclinePostRequestBody and sets the default values.
func NewItemCalendarViewItemInstancesItemMicrosoftGraphDeclineDeclinePostRequestBody()(*ItemCalendarViewItemInstancesItemMicrosoftGraphDeclineDeclinePostRequestBody) {
    m := &ItemCalendarViewItemInstancesItemMicrosoftGraphDeclineDeclinePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemCalendarViewItemInstancesItemMicrosoftGraphDeclineDeclinePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemCalendarViewItemInstancesItemMicrosoftGraphDeclineDeclinePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemCalendarViewItemInstancesItemMicrosoftGraphDeclineDeclinePostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemCalendarViewItemInstancesItemMicrosoftGraphDeclineDeclinePostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetComment gets the comment property value. The Comment property
func (m *ItemCalendarViewItemInstancesItemMicrosoftGraphDeclineDeclinePostRequestBody) GetComment()(*string) {
    return m.comment
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemCalendarViewItemInstancesItemMicrosoftGraphDeclineDeclinePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["comment"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComment(val)
        }
        return nil
    }
    res["proposedNewTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateTimeSlotFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProposedNewTime(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TimeSlotable))
        }
        return nil
    }
    res["sendResponse"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSendResponse(val)
        }
        return nil
    }
    return res
}
// GetProposedNewTime gets the proposedNewTime property value. The ProposedNewTime property
func (m *ItemCalendarViewItemInstancesItemMicrosoftGraphDeclineDeclinePostRequestBody) GetProposedNewTime()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TimeSlotable) {
    return m.proposedNewTime
}
// GetSendResponse gets the sendResponse property value. The SendResponse property
func (m *ItemCalendarViewItemInstancesItemMicrosoftGraphDeclineDeclinePostRequestBody) GetSendResponse()(*bool) {
    return m.sendResponse
}
// Serialize serializes information the current object
func (m *ItemCalendarViewItemInstancesItemMicrosoftGraphDeclineDeclinePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("comment", m.GetComment())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("proposedNewTime", m.GetProposedNewTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("sendResponse", m.GetSendResponse())
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemCalendarViewItemInstancesItemMicrosoftGraphDeclineDeclinePostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetComment sets the comment property value. The Comment property
func (m *ItemCalendarViewItemInstancesItemMicrosoftGraphDeclineDeclinePostRequestBody) SetComment(value *string)() {
    m.comment = value
}
// SetProposedNewTime sets the proposedNewTime property value. The ProposedNewTime property
func (m *ItemCalendarViewItemInstancesItemMicrosoftGraphDeclineDeclinePostRequestBody) SetProposedNewTime(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TimeSlotable)() {
    m.proposedNewTime = value
}
// SetSendResponse sets the sendResponse property value. The SendResponse property
func (m *ItemCalendarViewItemInstancesItemMicrosoftGraphDeclineDeclinePostRequestBody) SetSendResponse(value *bool)() {
    m.sendResponse = value
}
