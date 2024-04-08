package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RequestSchedule 
type RequestSchedule struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // When the eligible or active assignment expires.
    expiration ExpirationPatternable
    // The OdataType property
    odataType *string
    // The frequency of the  eligible or active assignment. This property is currently unsupported in PIM.
    recurrence PatternedRecurrenceable
    // When the  eligible or active assignment becomes active.
    startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// NewRequestSchedule instantiates a new requestSchedule and sets the default values.
func NewRequestSchedule()(*RequestSchedule) {
    m := &RequestSchedule{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRequestScheduleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRequestScheduleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRequestSchedule(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RequestSchedule) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetExpiration gets the expiration property value. When the eligible or active assignment expires.
func (m *RequestSchedule) GetExpiration()(ExpirationPatternable) {
    return m.expiration
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RequestSchedule) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["expiration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateExpirationPatternFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpiration(val.(ExpirationPatternable))
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
    res["recurrence"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePatternedRecurrenceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecurrence(val.(PatternedRecurrenceable))
        }
        return nil
    }
    res["startDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartDateTime(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *RequestSchedule) GetOdataType()(*string) {
    return m.odataType
}
// GetRecurrence gets the recurrence property value. The frequency of the  eligible or active assignment. This property is currently unsupported in PIM.
func (m *RequestSchedule) GetRecurrence()(PatternedRecurrenceable) {
    return m.recurrence
}
// GetStartDateTime gets the startDateTime property value. When the  eligible or active assignment becomes active.
func (m *RequestSchedule) GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.startDateTime
}
// Serialize serializes information the current object
func (m *RequestSchedule) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("expiration", m.GetExpiration())
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
        err := writer.WriteObjectValue("recurrence", m.GetRecurrence())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("startDateTime", m.GetStartDateTime())
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
func (m *RequestSchedule) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetExpiration sets the expiration property value. When the eligible or active assignment expires.
func (m *RequestSchedule) SetExpiration(value ExpirationPatternable)() {
    m.expiration = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *RequestSchedule) SetOdataType(value *string)() {
    m.odataType = value
}
// SetRecurrence sets the recurrence property value. The frequency of the  eligible or active assignment. This property is currently unsupported in PIM.
func (m *RequestSchedule) SetRecurrence(value PatternedRecurrenceable)() {
    m.recurrence = value
}
// SetStartDateTime sets the startDateTime property value. When the  eligible or active assignment becomes active.
func (m *RequestSchedule) SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.startDateTime = value
}
