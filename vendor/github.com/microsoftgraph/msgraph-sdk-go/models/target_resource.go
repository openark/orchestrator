package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TargetResource 
type TargetResource struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Indicates the visible name defined for the resource. Typically specified when the resource is created.
    displayName *string
    // When type is set to Group, this indicates the group type. Possible values are: unifiedGroups, azureAD, and unknownFutureValue
    groupType *GroupType
    // Indicates the unique ID of the resource.
    id *string
    // Indicates name, old value and new value of each attribute that changed. Property values depend on the operation type.
    modifiedProperties []ModifiedPropertyable
    // The OdataType property
    odataType *string
    // Describes the resource type.  Example values include Application, Group, ServicePrincipal, and User.
    typeEscaped *string
    // When type is set to User, this includes the user name that initiated the action; null for other types.
    userPrincipalName *string
}
// NewTargetResource instantiates a new targetResource and sets the default values.
func NewTargetResource()(*TargetResource) {
    m := &TargetResource{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTargetResourceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTargetResourceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTargetResource(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TargetResource) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDisplayName gets the displayName property value. Indicates the visible name defined for the resource. Typically specified when the resource is created.
func (m *TargetResource) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TargetResource) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["groupType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseGroupType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupType(val.(*GroupType))
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
    res["modifiedProperties"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateModifiedPropertyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ModifiedPropertyable, len(val))
            for i, v := range val {
                res[i] = v.(ModifiedPropertyable)
            }
            m.SetModifiedProperties(res)
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
    res["userPrincipalName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserPrincipalName(val)
        }
        return nil
    }
    return res
}
// GetGroupType gets the groupType property value. When type is set to Group, this indicates the group type. Possible values are: unifiedGroups, azureAD, and unknownFutureValue
func (m *TargetResource) GetGroupType()(*GroupType) {
    return m.groupType
}
// GetId gets the id property value. Indicates the unique ID of the resource.
func (m *TargetResource) GetId()(*string) {
    return m.id
}
// GetModifiedProperties gets the modifiedProperties property value. Indicates name, old value and new value of each attribute that changed. Property values depend on the operation type.
func (m *TargetResource) GetModifiedProperties()([]ModifiedPropertyable) {
    return m.modifiedProperties
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *TargetResource) GetOdataType()(*string) {
    return m.odataType
}
// GetType gets the type property value. Describes the resource type.  Example values include Application, Group, ServicePrincipal, and User.
func (m *TargetResource) GetType()(*string) {
    return m.typeEscaped
}
// GetUserPrincipalName gets the userPrincipalName property value. When type is set to User, this includes the user name that initiated the action; null for other types.
func (m *TargetResource) GetUserPrincipalName()(*string) {
    return m.userPrincipalName
}
// Serialize serializes information the current object
func (m *TargetResource) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetGroupType() != nil {
        cast := (*m.GetGroupType()).String()
        err := writer.WriteStringValue("groupType", &cast)
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
    if m.GetModifiedProperties() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetModifiedProperties()))
        for i, v := range m.GetModifiedProperties() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("modifiedProperties", cast)
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
        err := writer.WriteStringValue("userPrincipalName", m.GetUserPrincipalName())
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
func (m *TargetResource) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDisplayName sets the displayName property value. Indicates the visible name defined for the resource. Typically specified when the resource is created.
func (m *TargetResource) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetGroupType sets the groupType property value. When type is set to Group, this indicates the group type. Possible values are: unifiedGroups, azureAD, and unknownFutureValue
func (m *TargetResource) SetGroupType(value *GroupType)() {
    m.groupType = value
}
// SetId sets the id property value. Indicates the unique ID of the resource.
func (m *TargetResource) SetId(value *string)() {
    m.id = value
}
// SetModifiedProperties sets the modifiedProperties property value. Indicates name, old value and new value of each attribute that changed. Property values depend on the operation type.
func (m *TargetResource) SetModifiedProperties(value []ModifiedPropertyable)() {
    m.modifiedProperties = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *TargetResource) SetOdataType(value *string)() {
    m.odataType = value
}
// SetType sets the type property value. Describes the resource type.  Example values include Application, Group, ServicePrincipal, and User.
func (m *TargetResource) SetType(value *string)() {
    m.typeEscaped = value
}
// SetUserPrincipalName sets the userPrincipalName property value. When type is set to User, this includes the user name that initiated the action; null for other types.
func (m *TargetResource) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
