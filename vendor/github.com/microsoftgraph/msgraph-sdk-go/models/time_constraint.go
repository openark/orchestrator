package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TimeConstraint 
type TimeConstraint struct {
    // The nature of the activity, optional. The possible values are: work, personal, unrestricted, or unknown.
    activityDomain *ActivityDomain
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The OdataType property
    odataType *string
    // The timeSlots property
    timeSlots []TimeSlotable
}
// NewTimeConstraint instantiates a new timeConstraint and sets the default values.
func NewTimeConstraint()(*TimeConstraint) {
    m := &TimeConstraint{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTimeConstraintFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTimeConstraintFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTimeConstraint(), nil
}
// GetActivityDomain gets the activityDomain property value. The nature of the activity, optional. The possible values are: work, personal, unrestricted, or unknown.
func (m *TimeConstraint) GetActivityDomain()(*ActivityDomain) {
    return m.activityDomain
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TimeConstraint) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TimeConstraint) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["activityDomain"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseActivityDomain)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActivityDomain(val.(*ActivityDomain))
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
    res["timeSlots"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTimeSlotFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TimeSlotable, len(val))
            for i, v := range val {
                res[i] = v.(TimeSlotable)
            }
            m.SetTimeSlots(res)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *TimeConstraint) GetOdataType()(*string) {
    return m.odataType
}
// GetTimeSlots gets the timeSlots property value. The timeSlots property
func (m *TimeConstraint) GetTimeSlots()([]TimeSlotable) {
    return m.timeSlots
}
// Serialize serializes information the current object
func (m *TimeConstraint) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetActivityDomain() != nil {
        cast := (*m.GetActivityDomain()).String()
        err := writer.WriteStringValue("activityDomain", &cast)
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
    if m.GetTimeSlots() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTimeSlots()))
        for i, v := range m.GetTimeSlots() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("timeSlots", cast)
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
// SetActivityDomain sets the activityDomain property value. The nature of the activity, optional. The possible values are: work, personal, unrestricted, or unknown.
func (m *TimeConstraint) SetActivityDomain(value *ActivityDomain)() {
    m.activityDomain = value
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TimeConstraint) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *TimeConstraint) SetOdataType(value *string)() {
    m.odataType = value
}
// SetTimeSlots sets the timeSlots property value. The timeSlots property
func (m *TimeConstraint) SetTimeSlots(value []TimeSlotable)() {
    m.timeSlots = value
}
