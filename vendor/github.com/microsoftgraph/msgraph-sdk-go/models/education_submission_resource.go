package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EducationSubmissionResource 
type EducationSubmissionResource struct {
    Entity
    // Pointer to the assignment from which this resource was copied. If this is null, the student uploaded the resource.
    assignmentResourceUrl *string
    // Resource object.
    resource EducationResourceable
}
// NewEducationSubmissionResource instantiates a new educationSubmissionResource and sets the default values.
func NewEducationSubmissionResource()(*EducationSubmissionResource) {
    m := &EducationSubmissionResource{
        Entity: *NewEntity(),
    }
    return m
}
// CreateEducationSubmissionResourceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEducationSubmissionResourceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEducationSubmissionResource(), nil
}
// GetAssignmentResourceUrl gets the assignmentResourceUrl property value. Pointer to the assignment from which this resource was copied. If this is null, the student uploaded the resource.
func (m *EducationSubmissionResource) GetAssignmentResourceUrl()(*string) {
    return m.assignmentResourceUrl
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EducationSubmissionResource) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignmentResourceUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignmentResourceUrl(val)
        }
        return nil
    }
    res["resource"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEducationResourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResource(val.(EducationResourceable))
        }
        return nil
    }
    return res
}
// GetResource gets the resource property value. Resource object.
func (m *EducationSubmissionResource) GetResource()(EducationResourceable) {
    return m.resource
}
// Serialize serializes information the current object
func (m *EducationSubmissionResource) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("assignmentResourceUrl", m.GetAssignmentResourceUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("resource", m.GetResource())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAssignmentResourceUrl sets the assignmentResourceUrl property value. Pointer to the assignment from which this resource was copied. If this is null, the student uploaded the resource.
func (m *EducationSubmissionResource) SetAssignmentResourceUrl(value *string)() {
    m.assignmentResourceUrl = value
}
// SetResource sets the resource property value. Resource object.
func (m *EducationSubmissionResource) SetResource(value EducationResourceable)() {
    m.resource = value
}
