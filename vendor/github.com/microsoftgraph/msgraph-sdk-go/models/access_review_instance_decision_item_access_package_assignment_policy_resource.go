package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccessReviewInstanceDecisionItemAccessPackageAssignmentPolicyResource 
type AccessReviewInstanceDecisionItemAccessPackageAssignmentPolicyResource struct {
    AccessReviewInstanceDecisionItemResource
    // Display name of the access package to which access has been granted.
    accessPackageDisplayName *string
    // Identifier of the access package to which access has been granted.
    accessPackageId *string
}
// NewAccessReviewInstanceDecisionItemAccessPackageAssignmentPolicyResource instantiates a new AccessReviewInstanceDecisionItemAccessPackageAssignmentPolicyResource and sets the default values.
func NewAccessReviewInstanceDecisionItemAccessPackageAssignmentPolicyResource()(*AccessReviewInstanceDecisionItemAccessPackageAssignmentPolicyResource) {
    m := &AccessReviewInstanceDecisionItemAccessPackageAssignmentPolicyResource{
        AccessReviewInstanceDecisionItemResource: *NewAccessReviewInstanceDecisionItemResource(),
    }
    odataTypeValue := "#microsoft.graph.accessReviewInstanceDecisionItemAccessPackageAssignmentPolicyResource"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateAccessReviewInstanceDecisionItemAccessPackageAssignmentPolicyResourceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAccessReviewInstanceDecisionItemAccessPackageAssignmentPolicyResourceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAccessReviewInstanceDecisionItemAccessPackageAssignmentPolicyResource(), nil
}
// GetAccessPackageDisplayName gets the accessPackageDisplayName property value. Display name of the access package to which access has been granted.
func (m *AccessReviewInstanceDecisionItemAccessPackageAssignmentPolicyResource) GetAccessPackageDisplayName()(*string) {
    return m.accessPackageDisplayName
}
// GetAccessPackageId gets the accessPackageId property value. Identifier of the access package to which access has been granted.
func (m *AccessReviewInstanceDecisionItemAccessPackageAssignmentPolicyResource) GetAccessPackageId()(*string) {
    return m.accessPackageId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AccessReviewInstanceDecisionItemAccessPackageAssignmentPolicyResource) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AccessReviewInstanceDecisionItemResource.GetFieldDeserializers()
    res["accessPackageDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessPackageDisplayName(val)
        }
        return nil
    }
    res["accessPackageId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessPackageId(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *AccessReviewInstanceDecisionItemAccessPackageAssignmentPolicyResource) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AccessReviewInstanceDecisionItemResource.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("accessPackageDisplayName", m.GetAccessPackageDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("accessPackageId", m.GetAccessPackageId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccessPackageDisplayName sets the accessPackageDisplayName property value. Display name of the access package to which access has been granted.
func (m *AccessReviewInstanceDecisionItemAccessPackageAssignmentPolicyResource) SetAccessPackageDisplayName(value *string)() {
    m.accessPackageDisplayName = value
}
// SetAccessPackageId sets the accessPackageId property value. Identifier of the access package to which access has been granted.
func (m *AccessReviewInstanceDecisionItemAccessPackageAssignmentPolicyResource) SetAccessPackageId(value *string)() {
    m.accessPackageId = value
}
