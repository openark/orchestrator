package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OnlineMeetingInfo 
type OnlineMeetingInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The ID of the conference.
    conferenceId *string
    // The external link that launches the online meeting. This is a URL that clients will launch into a browser and will redirect the user to join the meeting.
    joinUrl *string
    // The OdataType property
    odataType *string
    // All of the phone numbers associated with this conference.
    phones []Phoneable
    // The pre-formatted quickdial for this call.
    quickDial *string
    // The toll free numbers that can be used to join the conference.
    tollFreeNumbers []string
    // The toll number that can be used to join the conference.
    tollNumber *string
}
// NewOnlineMeetingInfo instantiates a new onlineMeetingInfo and sets the default values.
func NewOnlineMeetingInfo()(*OnlineMeetingInfo) {
    m := &OnlineMeetingInfo{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateOnlineMeetingInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOnlineMeetingInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOnlineMeetingInfo(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OnlineMeetingInfo) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetConferenceId gets the conferenceId property value. The ID of the conference.
func (m *OnlineMeetingInfo) GetConferenceId()(*string) {
    return m.conferenceId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OnlineMeetingInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["conferenceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConferenceId(val)
        }
        return nil
    }
    res["joinUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJoinUrl(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["phones"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePhoneFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Phoneable, len(val))
            for i, v := range val {
                res[i] = v.(Phoneable)
            }
            m.SetPhones(res)
        }
        return nil
    }
    res["quickDial"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQuickDial(val)
        }
        return nil
    }
    res["tollFreeNumbers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetTollFreeNumbers(res)
        }
        return nil
    }
    res["tollNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTollNumber(val)
        }
        return nil
    }
    return res
}
// GetJoinUrl gets the joinUrl property value. The external link that launches the online meeting. This is a URL that clients will launch into a browser and will redirect the user to join the meeting.
func (m *OnlineMeetingInfo) GetJoinUrl()(*string) {
    return m.joinUrl
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *OnlineMeetingInfo) GetOdataType()(*string) {
    return m.odataType
}
// GetPhones gets the phones property value. All of the phone numbers associated with this conference.
func (m *OnlineMeetingInfo) GetPhones()([]Phoneable) {
    return m.phones
}
// GetQuickDial gets the quickDial property value. The pre-formatted quickdial for this call.
func (m *OnlineMeetingInfo) GetQuickDial()(*string) {
    return m.quickDial
}
// GetTollFreeNumbers gets the tollFreeNumbers property value. The toll free numbers that can be used to join the conference.
func (m *OnlineMeetingInfo) GetTollFreeNumbers()([]string) {
    return m.tollFreeNumbers
}
// GetTollNumber gets the tollNumber property value. The toll number that can be used to join the conference.
func (m *OnlineMeetingInfo) GetTollNumber()(*string) {
    return m.tollNumber
}
// Serialize serializes information the current object
func (m *OnlineMeetingInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("conferenceId", m.GetConferenceId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("joinUrl", m.GetJoinUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    if m.GetPhones() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPhones()))
        for i, v := range m.GetPhones() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("phones", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("quickDial", m.GetQuickDial())
        if err != nil {
            return err
        }
    }
    if m.GetTollFreeNumbers() != nil {
        err := writer.WriteCollectionOfStringValues("tollFreeNumbers", m.GetTollFreeNumbers())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("tollNumber", m.GetTollNumber())
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
func (m *OnlineMeetingInfo) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetConferenceId sets the conferenceId property value. The ID of the conference.
func (m *OnlineMeetingInfo) SetConferenceId(value *string)() {
    m.conferenceId = value
}
// SetJoinUrl sets the joinUrl property value. The external link that launches the online meeting. This is a URL that clients will launch into a browser and will redirect the user to join the meeting.
func (m *OnlineMeetingInfo) SetJoinUrl(value *string)() {
    m.joinUrl = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *OnlineMeetingInfo) SetOdataType(value *string)() {
    m.odataType = value
}
// SetPhones sets the phones property value. All of the phone numbers associated with this conference.
func (m *OnlineMeetingInfo) SetPhones(value []Phoneable)() {
    m.phones = value
}
// SetQuickDial sets the quickDial property value. The pre-formatted quickdial for this call.
func (m *OnlineMeetingInfo) SetQuickDial(value *string)() {
    m.quickDial = value
}
// SetTollFreeNumbers sets the tollFreeNumbers property value. The toll free numbers that can be used to join the conference.
func (m *OnlineMeetingInfo) SetTollFreeNumbers(value []string)() {
    m.tollFreeNumbers = value
}
// SetTollNumber sets the tollNumber property value. The toll number that can be used to join the conference.
func (m *OnlineMeetingInfo) SetTollNumber(value *string)() {
    m.tollNumber = value
}
