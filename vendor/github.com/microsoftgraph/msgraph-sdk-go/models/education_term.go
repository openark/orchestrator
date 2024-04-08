package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EducationTerm 
type EducationTerm struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Display name of the term.
    displayName *string
    // End of the term.
    endDate *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
    // ID of term in the syncing system.
    externalId *string
    // The OdataType property
    odataType *string
    // Start of the term.
    startDate *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
}
// NewEducationTerm instantiates a new educationTerm and sets the default values.
func NewEducationTerm()(*EducationTerm) {
    m := &EducationTerm{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateEducationTermFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEducationTermFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEducationTerm(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EducationTerm) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDisplayName gets the displayName property value. Display name of the term.
func (m *EducationTerm) GetDisplayName()(*string) {
    return m.displayName
}
// GetEndDate gets the endDate property value. End of the term.
func (m *EducationTerm) GetEndDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    return m.endDate
}
// GetExternalId gets the externalId property value. ID of term in the syncing system.
func (m *EducationTerm) GetExternalId()(*string) {
    return m.externalId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EducationTerm) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["endDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndDate(val)
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
    res["startDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartDate(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *EducationTerm) GetOdataType()(*string) {
    return m.odataType
}
// GetStartDate gets the startDate property value. Start of the term.
func (m *EducationTerm) GetStartDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    return m.startDate
}
// Serialize serializes information the current object
func (m *EducationTerm) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteDateOnlyValue("endDate", m.GetEndDate())
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
        err := writer.WriteDateOnlyValue("startDate", m.GetStartDate())
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
func (m *EducationTerm) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDisplayName sets the displayName property value. Display name of the term.
func (m *EducationTerm) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetEndDate sets the endDate property value. End of the term.
func (m *EducationTerm) SetEndDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    m.endDate = value
}
// SetExternalId sets the externalId property value. ID of term in the syncing system.
func (m *EducationTerm) SetExternalId(value *string)() {
    m.externalId = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *EducationTerm) SetOdataType(value *string)() {
    m.odataType = value
}
// SetStartDate sets the startDate property value. Start of the term.
func (m *EducationTerm) SetStartDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    m.startDate = value
}
