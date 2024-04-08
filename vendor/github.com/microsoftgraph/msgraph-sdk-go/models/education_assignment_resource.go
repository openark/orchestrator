package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EducationAssignmentResource 
type EducationAssignmentResource struct {
    Entity
    // Indicates whether this resource should be copied to each student submission for modification and submission. Required
    distributeForStudentWork *bool
    // Resource object that has been associated with this assignment.
    resource EducationResourceable
}
// NewEducationAssignmentResource instantiates a new educationAssignmentResource and sets the default values.
func NewEducationAssignmentResource()(*EducationAssignmentResource) {
    m := &EducationAssignmentResource{
        Entity: *NewEntity(),
    }
    return m
}
// CreateEducationAssignmentResourceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEducationAssignmentResourceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEducationAssignmentResource(), nil
}
// GetDistributeForStudentWork gets the distributeForStudentWork property value. Indicates whether this resource should be copied to each student submission for modification and submission. Required
func (m *EducationAssignmentResource) GetDistributeForStudentWork()(*bool) {
    return m.distributeForStudentWork
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EducationAssignmentResource) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["distributeForStudentWork"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDistributeForStudentWork(val)
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
// GetResource gets the resource property value. Resource object that has been associated with this assignment.
func (m *EducationAssignmentResource) GetResource()(EducationResourceable) {
    return m.resource
}
// Serialize serializes information the current object
func (m *EducationAssignmentResource) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("distributeForStudentWork", m.GetDistributeForStudentWork())
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
// SetDistributeForStudentWork sets the distributeForStudentWork property value. Indicates whether this resource should be copied to each student submission for modification and submission. Required
func (m *EducationAssignmentResource) SetDistributeForStudentWork(value *bool)() {
    m.distributeForStudentWork = value
}
// SetResource sets the resource property value. Resource object that has been associated with this assignment.
func (m *EducationAssignmentResource) SetResource(value EducationResourceable)() {
    m.resource = value
}
