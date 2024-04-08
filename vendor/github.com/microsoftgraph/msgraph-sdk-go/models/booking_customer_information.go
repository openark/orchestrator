package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BookingCustomerInformation 
type BookingCustomerInformation struct {
    BookingCustomerInformationBase
    // The ID of the bookingCustomer for this appointment. If no ID is specified when an appointment is created, then a new bookingCustomer object is created. Once set, you should consider the customerId immutable.
    customerId *string
    // It consists of the list of custom questions and answers given by the customer as part of the appointment
    customQuestionAnswers []BookingQuestionAnswerable
    // The SMTP address of the bookingCustomer who is booking the appointment
    emailAddress *string
    // Represents location information for the bookingCustomer who is booking the appointment.
    location Locationable
    // The customer's name.
    name *string
    // Notes from the customer associated with this appointment. You can get the value only when reading this bookingAppointment by its ID. You can set this property only when initially creating an appointment with a new customer. After that point, the value is computed from the customer represented by the customerId.
    notes *string
    // The customer's phone number.
    phone *string
    // The time zone of the customer. For a list of possible values, see dateTimeTimeZone.
    timeZone *string
}
// NewBookingCustomerInformation instantiates a new BookingCustomerInformation and sets the default values.
func NewBookingCustomerInformation()(*BookingCustomerInformation) {
    m := &BookingCustomerInformation{
        BookingCustomerInformationBase: *NewBookingCustomerInformationBase(),
    }
    odataTypeValue := "#microsoft.graph.bookingCustomerInformation"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateBookingCustomerInformationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateBookingCustomerInformationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBookingCustomerInformation(), nil
}
// GetCustomerId gets the customerId property value. The ID of the bookingCustomer for this appointment. If no ID is specified when an appointment is created, then a new bookingCustomer object is created. Once set, you should consider the customerId immutable.
func (m *BookingCustomerInformation) GetCustomerId()(*string) {
    return m.customerId
}
// GetCustomQuestionAnswers gets the customQuestionAnswers property value. It consists of the list of custom questions and answers given by the customer as part of the appointment
func (m *BookingCustomerInformation) GetCustomQuestionAnswers()([]BookingQuestionAnswerable) {
    return m.customQuestionAnswers
}
// GetEmailAddress gets the emailAddress property value. The SMTP address of the bookingCustomer who is booking the appointment
func (m *BookingCustomerInformation) GetEmailAddress()(*string) {
    return m.emailAddress
}
// GetFieldDeserializers the deserialization information for the current model
func (m *BookingCustomerInformation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BookingCustomerInformationBase.GetFieldDeserializers()
    res["customerId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomerId(val)
        }
        return nil
    }
    res["customQuestionAnswers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateBookingQuestionAnswerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]BookingQuestionAnswerable, len(val))
            for i, v := range val {
                res[i] = v.(BookingQuestionAnswerable)
            }
            m.SetCustomQuestionAnswers(res)
        }
        return nil
    }
    res["emailAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmailAddress(val)
        }
        return nil
    }
    res["location"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateLocationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocation(val.(Locationable))
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["notes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotes(val)
        }
        return nil
    }
    res["phone"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPhone(val)
        }
        return nil
    }
    res["timeZone"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTimeZone(val)
        }
        return nil
    }
    return res
}
// GetLocation gets the location property value. Represents location information for the bookingCustomer who is booking the appointment.
func (m *BookingCustomerInformation) GetLocation()(Locationable) {
    return m.location
}
// GetName gets the name property value. The customer's name.
func (m *BookingCustomerInformation) GetName()(*string) {
    return m.name
}
// GetNotes gets the notes property value. Notes from the customer associated with this appointment. You can get the value only when reading this bookingAppointment by its ID. You can set this property only when initially creating an appointment with a new customer. After that point, the value is computed from the customer represented by the customerId.
func (m *BookingCustomerInformation) GetNotes()(*string) {
    return m.notes
}
// GetPhone gets the phone property value. The customer's phone number.
func (m *BookingCustomerInformation) GetPhone()(*string) {
    return m.phone
}
// GetTimeZone gets the timeZone property value. The time zone of the customer. For a list of possible values, see dateTimeTimeZone.
func (m *BookingCustomerInformation) GetTimeZone()(*string) {
    return m.timeZone
}
// Serialize serializes information the current object
func (m *BookingCustomerInformation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BookingCustomerInformationBase.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("customerId", m.GetCustomerId())
        if err != nil {
            return err
        }
    }
    if m.GetCustomQuestionAnswers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCustomQuestionAnswers()))
        for i, v := range m.GetCustomQuestionAnswers() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("customQuestionAnswers", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("emailAddress", m.GetEmailAddress())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("location", m.GetLocation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("notes", m.GetNotes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("phone", m.GetPhone())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("timeZone", m.GetTimeZone())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCustomerId sets the customerId property value. The ID of the bookingCustomer for this appointment. If no ID is specified when an appointment is created, then a new bookingCustomer object is created. Once set, you should consider the customerId immutable.
func (m *BookingCustomerInformation) SetCustomerId(value *string)() {
    m.customerId = value
}
// SetCustomQuestionAnswers sets the customQuestionAnswers property value. It consists of the list of custom questions and answers given by the customer as part of the appointment
func (m *BookingCustomerInformation) SetCustomQuestionAnswers(value []BookingQuestionAnswerable)() {
    m.customQuestionAnswers = value
}
// SetEmailAddress sets the emailAddress property value. The SMTP address of the bookingCustomer who is booking the appointment
func (m *BookingCustomerInformation) SetEmailAddress(value *string)() {
    m.emailAddress = value
}
// SetLocation sets the location property value. Represents location information for the bookingCustomer who is booking the appointment.
func (m *BookingCustomerInformation) SetLocation(value Locationable)() {
    m.location = value
}
// SetName sets the name property value. The customer's name.
func (m *BookingCustomerInformation) SetName(value *string)() {
    m.name = value
}
// SetNotes sets the notes property value. Notes from the customer associated with this appointment. You can get the value only when reading this bookingAppointment by its ID. You can set this property only when initially creating an appointment with a new customer. After that point, the value is computed from the customer represented by the customerId.
func (m *BookingCustomerInformation) SetNotes(value *string)() {
    m.notes = value
}
// SetPhone sets the phone property value. The customer's phone number.
func (m *BookingCustomerInformation) SetPhone(value *string)() {
    m.phone = value
}
// SetTimeZone sets the timeZone property value. The time zone of the customer. For a list of possible values, see dateTimeTimeZone.
func (m *BookingCustomerInformation) SetTimeZone(value *string)() {
    m.timeZone = value
}
