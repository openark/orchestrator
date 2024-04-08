package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EducationWordResource 
type EducationWordResource struct {
    EducationResource
    // Location of the file on disk.
    fileUrl *string
}
// NewEducationWordResource instantiates a new EducationWordResource and sets the default values.
func NewEducationWordResource()(*EducationWordResource) {
    m := &EducationWordResource{
        EducationResource: *NewEducationResource(),
    }
    odataTypeValue := "#microsoft.graph.educationWordResource"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateEducationWordResourceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEducationWordResourceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEducationWordResource(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EducationWordResource) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
// GetFileUrl gets the fileUrl property value. Location of the file on disk.
func (m *EducationWordResource) GetFileUrl()(*string) {
    return m.fileUrl
}
// Serialize serializes information the current object
func (m *EducationWordResource) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
// SetFileUrl sets the fileUrl property value. Location of the file on disk.
func (m *EducationWordResource) SetFileUrl(value *string)() {
    m.fileUrl = value
}
