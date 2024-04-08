package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PatternedRecurrence 
type PatternedRecurrence struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The OdataType property
    odataType *string
    // The frequency of an event.  For access reviews: Do not specify this property for a one-time access review.  Only interval, dayOfMonth, and type (weekly, absoluteMonthly) properties of recurrencePattern are supported.
    pattern RecurrencePatternable
    // The duration of an event.
    rangeEscaped RecurrenceRangeable
}
// NewPatternedRecurrence instantiates a new patternedRecurrence and sets the default values.
func NewPatternedRecurrence()(*PatternedRecurrence) {
    m := &PatternedRecurrence{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePatternedRecurrenceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePatternedRecurrenceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPatternedRecurrence(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PatternedRecurrence) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PatternedRecurrence) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["pattern"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRecurrencePatternFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPattern(val.(RecurrencePatternable))
        }
        return nil
    }
    res["range"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRecurrenceRangeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRange(val.(RecurrenceRangeable))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *PatternedRecurrence) GetOdataType()(*string) {
    return m.odataType
}
// GetPattern gets the pattern property value. The frequency of an event.  For access reviews: Do not specify this property for a one-time access review.  Only interval, dayOfMonth, and type (weekly, absoluteMonthly) properties of recurrencePattern are supported.
func (m *PatternedRecurrence) GetPattern()(RecurrencePatternable) {
    return m.pattern
}
// GetRange gets the range property value. The duration of an event.
func (m *PatternedRecurrence) GetRange()(RecurrenceRangeable) {
    return m.rangeEscaped
}
// Serialize serializes information the current object
func (m *PatternedRecurrence) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("pattern", m.GetPattern())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("range", m.GetRange())
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
func (m *PatternedRecurrence) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *PatternedRecurrence) SetOdataType(value *string)() {
    m.odataType = value
}
// SetPattern sets the pattern property value. The frequency of an event.  For access reviews: Do not specify this property for a one-time access review.  Only interval, dayOfMonth, and type (weekly, absoluteMonthly) properties of recurrencePattern are supported.
func (m *PatternedRecurrence) SetPattern(value RecurrencePatternable)() {
    m.pattern = value
}
// SetRange sets the range property value. The duration of an event.
func (m *PatternedRecurrence) SetRange(value RecurrenceRangeable)() {
    m.rangeEscaped = value
}
