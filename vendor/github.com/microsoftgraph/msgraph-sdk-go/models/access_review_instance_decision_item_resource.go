package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccessReviewInstanceDecisionItemResource 
type AccessReviewInstanceDecisionItemResource struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Display name of the resource
    displayName *string
    // Identifier of the resource
    id *string
    // The OdataType property
    odataType *string
    // Type of resource. Types include: Group, ServicePrincipal, DirectoryRole, AzureRole, AccessPackageAssignmentPolicy.
    typeEscaped *string
}
// NewAccessReviewInstanceDecisionItemResource instantiates a new accessReviewInstanceDecisionItemResource and sets the default values.
func NewAccessReviewInstanceDecisionItemResource()(*AccessReviewInstanceDecisionItemResource) {
    m := &AccessReviewInstanceDecisionItemResource{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAccessReviewInstanceDecisionItemResourceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAccessReviewInstanceDecisionItemResourceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.accessReviewInstanceDecisionItemAccessPackageAssignmentPolicyResource":
                        return NewAccessReviewInstanceDecisionItemAccessPackageAssignmentPolicyResource(), nil
                    case "#microsoft.graph.accessReviewInstanceDecisionItemAzureRoleResource":
                        return NewAccessReviewInstanceDecisionItemAzureRoleResource(), nil
                    case "#microsoft.graph.accessReviewInstanceDecisionItemServicePrincipalResource":
                        return NewAccessReviewInstanceDecisionItemServicePrincipalResource(), nil
                }
            }
        }
    }
    return NewAccessReviewInstanceDecisionItemResource(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AccessReviewInstanceDecisionItemResource) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDisplayName gets the displayName property value. Display name of the resource
func (m *AccessReviewInstanceDecisionItemResource) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AccessReviewInstanceDecisionItemResource) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
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
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val)
        }
        return nil
    }
    return res
}
// GetId gets the id property value. Identifier of the resource
func (m *AccessReviewInstanceDecisionItemResource) GetId()(*string) {
    return m.id
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AccessReviewInstanceDecisionItemResource) GetOdataType()(*string) {
    return m.odataType
}
// GetType gets the type property value. Type of resource. Types include: Group, ServicePrincipal, DirectoryRole, AzureRole, AccessPackageAssignmentPolicy.
func (m *AccessReviewInstanceDecisionItemResource) GetType()(*string) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *AccessReviewInstanceDecisionItemResource) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("id", m.GetId())
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
        err := writer.WriteStringValue("type", m.GetType())
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
func (m *AccessReviewInstanceDecisionItemResource) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDisplayName sets the displayName property value. Display name of the resource
func (m *AccessReviewInstanceDecisionItemResource) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetId sets the id property value. Identifier of the resource
func (m *AccessReviewInstanceDecisionItemResource) SetId(value *string)() {
    m.id = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AccessReviewInstanceDecisionItemResource) SetOdataType(value *string)() {
    m.odataType = value
}
// SetType sets the type property value. Type of resource. Types include: Group, ServicePrincipal, DirectoryRole, AzureRole, AccessPackageAssignmentPolicy.
func (m *AccessReviewInstanceDecisionItemResource) SetType(value *string)() {
    m.typeEscaped = value
}
