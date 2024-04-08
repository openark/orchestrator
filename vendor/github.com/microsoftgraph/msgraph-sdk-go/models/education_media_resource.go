package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EducationMediaResource 
type EducationMediaResource struct {
    EducationResource
    // Location of the file on shared point folder. Required
    fileUrl *string
}
// NewEducationMediaResource instantiates a new EducationMediaResource and sets the default values.
func NewEducationMediaResource()(*EducationMediaResource) {
    m := &EducationMediaResource{
        EducationResource: *NewEducationResource(),
    }
    odataTypeValue := "#microsoft.graph.educationMediaResource"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateEducationMediaResourceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEducationMediaResourceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEducationMediaResource(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EducationMediaResource) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.EducationResource.GetFieldDeserializers()
    res["fileUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFileUrl(val)
        }
        return nil
    }
    return res
}
// GetFileUrl gets the fileUrl property value. Location of the file on shared point folder. Required
func (m *EducationMediaResource) GetFileUrl()(*string) {
    return m.fileUrl
}
// Serialize serializes information the current object
func (m *EducationMediaResource) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.EducationResource.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("fileUrl", m.GetFileUrl())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetFileUrl sets the fileUrl property value. Location of the file on shared point folder. Required
func (m *EducationMediaResource) SetFileUrl(value *string)() {
    m.fileUrl = value
}
