package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SwapShiftsChangeRequest 
type SwapShiftsChangeRequest struct {
    OfferShiftRequest
    // ShiftId for the recipient user with whom the request is to swap.
    recipientShiftId *string
}
// NewSwapShiftsChangeRequest instantiates a new SwapShiftsChangeRequest and sets the default values.
func NewSwapShiftsChangeRequest()(*SwapShiftsChangeRequest) {
    m := &SwapShiftsChangeRequest{
        OfferShiftRequest: *NewOfferShiftRequest(),
    }
    odataTypeValue := "#microsoft.graph.swapShiftsChangeRequest"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateSwapShiftsChangeRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSwapShiftsChangeRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSwapShiftsChangeRequest(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SwapShiftsChangeRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.OfferShiftRequest.GetFieldDeserializers()
    res["recipientShiftId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecipientShiftId(val)
        }
        return nil
    }
    return res
}
// GetRecipientShiftId gets the recipientShiftId property value. ShiftId for the recipient user with whom the request is to swap.
func (m *SwapShiftsChangeRequest) GetRecipientShiftId()(*string) {
    return m.recipientShiftId
}
// Serialize serializes information the current object
func (m *SwapShiftsChangeRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.OfferShiftRequest.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("recipientShiftId", m.GetRecipientShiftId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetRecipientShiftId sets the recipientShiftId property value. ShiftId for the recipient user with whom the request is to swap.
func (m *SwapShiftsChangeRequest) SetRecipientShiftId(value *string)() {
    m.recipientShiftId = value
}
