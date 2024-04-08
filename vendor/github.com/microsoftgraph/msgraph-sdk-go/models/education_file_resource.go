package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EducationFileResource 
type EducationFileResource struct {
    EducationResource
    // Location on disk of the file resource.
    fileUrl *string
}
// NewEducationFileResource instantiates a new EducationFileResource and sets the default values.
func NewEducationFileResource()(*EducationFileResource) {
    m := &EducationFileResource{
        EducationResource: *NewEducationResource(),
    }
    odataTypeValue := "#microsoft.graph.educationFileResource"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateEducationFileResourceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEducationFileResourceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEducationFileResource(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EducationFileResource) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
// GetFileUrl gets the fileUrl property value. Location on disk of the file resource.
func (m *EducationFileResource) GetFileUrl()(*string) {
    return m.fileUrl
}
// Serialize serializes information the current object
func (m *EducationFileResource) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
// SetFileUrl sets the fileUrl property value. Location on disk of the file resource.
func (m *EducationFileResource) SetFileUrl(value *string)() {
    m.fileUrl = value
}
