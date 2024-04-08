package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AudioConferencing 
type AudioConferencing struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The conference id of the online meeting.
    conferenceId *string
    // A URL to the externally-accessible web page that contains dial-in information.
    dialinUrl *string
    // The OdataType property
    odataType *string
    // The tollFreeNumber property
    tollFreeNumber *string
    // List of toll-free numbers that are displayed in the meeting invite.
    tollFreeNumbers []string
    // The tollNumber property
    tollNumber *string
    // List of toll numbers that are displayed in the meeting invite.
    tollNumbers []string
}
// NewAudioConferencing instantiates a new audioConferencing and sets the default values.
func NewAudioConferencing()(*AudioConferencing) {
    m := &AudioConferencing{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAudioConferencingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAudioConferencingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAudioConferencing(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AudioConferencing) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetConferenceId gets the conferenceId property value. The conference id of the online meeting.
func (m *AudioConferencing) GetConferenceId()(*string) {
    return m.conferenceId
}
// GetDialinUrl gets the dialinUrl property value. A URL to the externally-accessible web page that contains dial-in information.
func (m *AudioConferencing) GetDialinUrl()(*string) {
    return m.dialinUrl
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AudioConferencing) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["dialinUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDialinUrl(val)
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
    res["tollFreeNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTollFreeNumber(val)
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
    res["tollNumbers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetTollNumbers(res)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AudioConferencing) GetOdataType()(*string) {
    return m.odataType
}
// GetTollFreeNumber gets the tollFreeNumber property value. The tollFreeNumber property
func (m *AudioConferencing) GetTollFreeNumber()(*string) {
    return m.tollFreeNumber
}
// GetTollFreeNumbers gets the tollFreeNumbers property value. List of toll-free numbers that are displayed in the meeting invite.
func (m *AudioConferencing) GetTollFreeNumbers()([]string) {
    return m.tollFreeNumbers
}
// GetTollNumber gets the tollNumber property value. The tollNumber property
func (m *AudioConferencing) GetTollNumber()(*string) {
    return m.tollNumber
}
// GetTollNumbers gets the tollNumbers property value. List of toll numbers that are displayed in the meeting invite.
func (m *AudioConferencing) GetTollNumbers()([]string) {
    return m.tollNumbers
}
// Serialize serializes information the current object
func (m *AudioConferencing) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("conferenceId", m.GetConferenceId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("dialinUrl", m.GetDialinUrl())
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
    {
        err := writer.WriteStringValue("tollFreeNumber", m.GetTollFreeNumber())
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
    if m.GetTollNumbers() != nil {
        err := writer.WriteCollectionOfStringValues("tollNumbers", m.GetTollNumbers())
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
func (m *AudioConferencing) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetConferenceId sets the conferenceId property value. The conference id of the online meeting.
func (m *AudioConferencing) SetConferenceId(value *string)() {
    m.conferenceId = value
}
// SetDialinUrl sets the dialinUrl property value. A URL to the externally-accessible web page that contains dial-in information.
func (m *AudioConferencing) SetDialinUrl(value *string)() {
    m.dialinUrl = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AudioConferencing) SetOdataType(value *string)() {
    m.odataType = value
}
// SetTollFreeNumber sets the tollFreeNumber property value. The tollFreeNumber property
func (m *AudioConferencing) SetTollFreeNumber(value *string)() {
    m.tollFreeNumber = value
}
// SetTollFreeNumbers sets the tollFreeNumbers property value. List of toll-free numbers that are displayed in the meeting invite.
func (m *AudioConferencing) SetTollFreeNumbers(value []string)() {
    m.tollFreeNumbers = value
}
// SetTollNumber sets the tollNumber property value. The tollNumber property
func (m *AudioConferencing) SetTollNumber(value *string)() {
    m.tollNumber = value
}
// SetTollNumbers sets the tollNumbers property value. List of toll numbers that are displayed in the meeting invite.
func (m *AudioConferencing) SetTollNumbers(value []string)() {
    m.tollNumbers = value
}
