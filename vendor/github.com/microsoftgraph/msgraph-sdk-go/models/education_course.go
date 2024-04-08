package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EducationCourse 
type EducationCourse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Unique identifier for the course.
    courseNumber *string
    // Description of the course.
    description *string
    // Name of the course.
    displayName *string
    // ID of the course from the syncing system.
    externalId *string
    // The OdataType property
    odataType *string
    // Subject of the course.
    subject *string
}
// NewEducationCourse instantiates a new educationCourse and sets the default values.
func NewEducationCourse()(*EducationCourse) {
    m := &EducationCourse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateEducationCourseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEducationCourseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEducationCourse(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EducationCourse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCourseNumber gets the courseNumber property value. Unique identifier for the course.
func (m *EducationCourse) GetCourseNumber()(*string) {
    return m.courseNumber
}
// GetDescription gets the description property value. Description of the course.
func (m *EducationCourse) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. Name of the course.
func (m *EducationCourse) GetDisplayName()(*string) {
    return m.displayName
}
// GetExternalId gets the externalId property value. ID of the course from the syncing system.
func (m *EducationCourse) GetExternalId()(*string) {
    return m.externalId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EducationCourse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["courseNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCourseNumber(val)
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["externalId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalId(val)
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
    res["subject"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubject(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *EducationCourse) GetOdataType()(*string) {
    return m.odataType
}
// GetSubject gets the subject property value. Subject of the course.
func (m *EducationCourse) GetSubject()(*string) {
    return m.subject
}
// Serialize serializes information the current object
func (m *EducationCourse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("courseNumber", m.GetCourseNumber())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("externalId", m.GetExternalId())
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
        err := writer.WriteStringValue("subject", m.GetSubject())
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
func (m *EducationCourse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCourseNumber sets the courseNumber property value. Unique identifier for the course.
func (m *EducationCourse) SetCourseNumber(value *string)() {
    m.courseNumber = value
}
// SetDescription sets the description property value. Description of the course.
func (m *EducationCourse) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. Name of the course.
func (m *EducationCourse) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetExternalId sets the externalId property value. ID of the course from the syncing system.
func (m *EducationCourse) SetExternalId(value *string)() {
    m.externalId = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *EducationCourse) SetOdataType(value *string)() {
    m.odataType = value
}
// SetSubject sets the subject property value. Subject of the course.
func (m *EducationCourse) SetSubject(value *string)() {
    m.subject = value
}
